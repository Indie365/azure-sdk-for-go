// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	common "github.com/Azure/azure-amqp-common-go/v3"

	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

const sessionFilterName = "com.microsoft:session-filter"
const defaultLinkRxBuffer = 2048

type (
	// Receiver provides connection, session and link handling for a receiving to an entity path
	Receiver struct {
		namespace          *Namespace
		client             *amqp.Client
		clientMu           sync.RWMutex
		session            *amqp.Session
		receiver           *amqp.Receiver
		entityPath         string
		doneListening      func()
		Name               string
		useSessions        bool
		sessionID          *string
		lastError          error
		lastErrorMu        sync.RWMutex
		mode               ReceiveMode
		prefetch           uint32
		DefaultDisposition DispositionAction
		Closed             bool
		cancelAuthRefresh  func() <-chan struct{}
	}

	// ReceiverOption provides a structure for configuring receivers
	ReceiverOption func(receiver *Receiver) error

	// ListenerHandleImpl provides the ability to close or listen to the close of a Receiver
	ListenerHandleImpl struct {
		r   *Receiver
		ctx context.Context
	}

	// LegacyReceiver is actually implemented by *Receiver
	LegacyReceiver interface {
		ReceiveOner

		Session() *amqp.Session
		SessionID() *string
		Client() *amqp.Client

		IssueCredit(credit uint32) error
		DrainCredit(ctx context.Context) error
		Listen(ctx context.Context, handler Handler) ListenerHandle
		CompleteMessage(ctx context.Context, message *Message) error
		AbandonMessage(ctx context.Context, message *Message) error
	}
)

// ReceiverWithSession configures a Receiver to use a session
func ReceiverWithSession(sessionID *string) ReceiverOption {
	return func(r *Receiver) error {
		r.sessionID = sessionID
		r.useSessions = true
		return nil
	}
}

// ReceiverWithReceiveMode configures a Receiver to use the specified receive mode
func ReceiverWithReceiveMode(mode ReceiveMode) ReceiverOption {
	return func(r *Receiver) error {
		r.mode = mode
		return nil
	}
}

// ReceiverWithPrefetchCount configures the receiver to attempt to fetch the number of messages specified by the prefect
// at one time.
//
// The default is 1 message at a time.
//
// Caution: Using PeekLock, messages have a set lock timeout, which can be renewed. By setting a high prefetch count, a
// local queue of messages could build up and cause message locks to expire before the message lands in the handler. If
// this happens, the message disposition will fail and will be re-queued and processed again.
func ReceiverWithPrefetchCount(prefetch uint32) ReceiverOption {
	return func(receiver *Receiver) error {
		receiver.prefetch = prefetch
		return nil
	}
}

// NewReceiver creates a new Service Bus message listener given an AMQP client and an entity path
func (ns *Namespace) NewReceiver(ctx context.Context, entityPath string, opts ...ReceiverOption) (LegacyReceiver, error) {
	ctx, span := ns.startSpanFromContext(ctx, "sb.Namespace.NewReceiver")
	defer span.End()

	r := &Receiver{
		namespace:  ns,
		entityPath: entityPath,
		mode:       PeekLockMode,
		prefetch:   1,
	}

	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}

	err := r.newSessionAndLink(ctx)
	if err != nil {
		_ = r.Close(ctx)
		return nil, err
	}

	return r, nil
}

// Close will close the AMQP session and link of the Receiver
func (r *Receiver) Close(ctx context.Context) error {
	ctx, span := r.startConsumerSpanFromContext(ctx, "sb.Receiver.Close")
	defer span.End()

	r.clientMu.Lock()
	defer r.clientMu.Unlock()

	return r.close(ctx)
}

// closes the session.  callers *must* hold the client write lock before calling!
func (r *Receiver) close(ctx context.Context) error {
	if r.doneListening != nil {
		r.doneListening()
	}

	if r.cancelAuthRefresh != nil {
		<-r.cancelAuthRefresh()
	}

	r.Closed = true

	var lastErr error
	if r.receiver != nil {
		lastErr = r.receiver.Close(ctx)
		if lastErr != nil {
			tab.For(ctx).Error(lastErr)
		}
	}

	if r.session != nil {
		if err := r.session.Close(ctx); err != nil {
			tab.For(ctx).Error(err)
			lastErr = err
		}
	}

	if r.client != nil {
		if err := r.client.Close(); err != nil {
			tab.For(ctx).Error(err)
			lastErr = err
		}
	}

	r.receiver = nil
	r.session = nil
	r.client = nil

	return lastErr
}

// Recover will attempt to close the current session and link, then rebuild them
func (r *Receiver) Recover(ctx context.Context) error {
	ctx, span := r.startConsumerSpanFromContext(ctx, "sb.Receiver.Recover")
	defer span.End()

	// we expect the Sender, session or client is in an error state, ignore errors
	closeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	closeCtx = tab.NewContext(closeCtx, span)
	defer cancel()
	// we must close then rebuild the session/link atomically
	r.clientMu.Lock()
	defer r.clientMu.Unlock()
	_ = r.close(closeCtx)
	return r.newSessionAndLink(ctx)
}

// ReceiveOne will receive one message from the link
func (r *Receiver) ReceiveOne(ctx context.Context, handler Handler) error {
	ctx, span := r.startConsumerSpanFromContext(ctx, "sb.Receiver.ReceiveOne")
	defer span.End()

	err := r.listenForMessage(ctx, newAmqpAdapterHandler(r, handler))
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	return nil
}

// IssueCredit adds credits to the link that will be sent
// in the next flow frame.
func (r *Receiver) IssueCredit(credits uint32) error {
	return r.receiver.IssueCredit(credits)
}

// DrainCredit will cause the next flow frame to have drain set to
// true, with the current credits in the link.
func (r *Receiver) DrainCredit(ctx context.Context) error {
	return r.receiver.DrainCredit(ctx)
}

// Listen start a listener for messages sent to the entity path
func (r *Receiver) Listen(ctx context.Context, handler Handler) ListenerHandle {
	ctx, done := context.WithCancel(ctx)
	r.doneListening = done

	ctx, span := r.startConsumerSpanFromContext(ctx, "sb.Receiver.Listen")
	defer span.End()

	go r.listenForMessages(ctx, newAmqpAdapterHandler(r, handler))

	return &ListenerHandleImpl{
		r:   r,
		ctx: ctx,
	}
}

// CompleteMessage will complete a message, causing it to be removed
// from the remote queue or subscription.
// NOTE: added to make mocking a bit simpler for the processor.
func (r *Receiver) CompleteMessage(ctx context.Context, message *Message) error {
	return message.Complete(ctx)
}

// AbandonMessage will abandon a message, causing it to be available to
// the remote queue or subscription.
// NOTE: added to make mocking a bit simpler for the processor.
func (r *Receiver) AbandonMessage(ctx context.Context, message *Message) error {
	return message.Abandon(ctx)
}

// Session returns the underlying *amqp.Session struct in `r.session.Session`.
func (r *Receiver) Session() *amqp.Session {
	return r.session
}

// Client returns the underlying *amqp.Client struct in `r.client`.
func (r *Receiver) Client() *amqp.Client {
	return r.client
}

// SessionID returns the sessionID in `r.sessionID`.
func (r *Receiver) SessionID() *string {
	return r.sessionID
}

func (r *Receiver) listenForMessages(ctx context.Context, handler amqpHandler) {
	ctx, span := r.startConsumerSpanFromContext(ctx, "sb.Receiver.listenForMessages")
	defer span.End()

	for {
		err := r.listenForMessage(ctx, handler)
		if err == nil {
			continue
		}

		select {
		case <-ctx.Done():
			tab.For(ctx).Debug("context done")
			return
		default:
			_, retryErr := common.Retry(10, 10*time.Second, func() (interface{}, error) {
				ctx, sp := r.startConsumerSpanFromContext(ctx, "sb.Receiver.listenForMessages.tryRecover")
				defer sp.End()

				tab.For(ctx).Debug("recovering connection")
				err := r.Recover(ctx)
				if err == nil {
					tab.For(ctx).Debug("recovered connection")
					return nil, nil
				}

				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				default:
					return nil, common.Retryable(err.Error())
				}
			})

			if retryErr != nil {
				tab.For(ctx).Debug("retried, but error was unrecoverable")
				r.setLastError(retryErr)
				if err := r.Close(ctx); err != nil {
					tab.For(ctx).Error(err)
				}
				return
			}
		}
	}
}

func (r *Receiver) setLastError(err error) {
	r.lastErrorMu.Lock()
	r.lastError = err
	r.lastErrorMu.Unlock()
}

func (r *Receiver) listenForMessage(ctx context.Context, handler amqpHandler) error {
	ctx, span := r.startConsumerSpanFromContext(ctx, "sb.Receiver.listenForMessage")
	defer span.End()

	var receiver *amqp.Receiver
	r.clientMu.RLock()
	if r.receiver == nil {
		r.clientMu.RUnlock()
		return r.connClosedError(ctx)
	}
	receiver = r.receiver
	r.clientMu.RUnlock()
	err := receiver.HandleMessage(ctx, func(message *amqp.Message) error {
		return handler.Handle(ctx, message)
	})
	if err != nil {
		tab.For(ctx).Debug(err.Error())
		return err
	}
	return nil
}

func (r *Receiver) connClosedError(ctx context.Context) error {
	name := "Receiver"
	if r.Name != "" {
		name = r.Name
	}
	err := ErrConnectionClosed(name)
	tab.For(ctx).Error(err)
	return err
}

// newSessionAndLink will replace the session and link on the Receiver
// NOTE: this does *not* take the write lock, callers must hold it as required!
func (r *Receiver) newSessionAndLink(ctx context.Context) error {
	ctx, span := r.startConsumerSpanFromContext(ctx, "sb.Receiver.newSessionAndLink")
	defer span.End()

	client, err := r.namespace.newClient(ctx)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}
	r.client = client

	r.cancelAuthRefresh, err = r.namespace.negotiateClaim(ctx, client, r.entityPath)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	amqpSession, err := client.NewSession()
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}

	r.session = amqpSession

	receiveMode := amqp.ModeSecond
	if r.mode == ReceiveAndDeleteMode {
		receiveMode = amqp.ModeFirst
	}

	opts := []amqp.LinkOption{
		amqp.LinkSourceAddress(r.entityPath),
		amqp.LinkReceiverSettle(receiveMode),
		amqp.LinkWithManualCredits(),
		amqp.LinkCredit(defaultLinkRxBuffer),
	}

	if r.mode == ReceiveAndDeleteMode {
		opts = append(opts, amqp.LinkSenderSettle(amqp.ModeSettled))
	}

	sessionOpt, useSessionOpt := r.getSessionFilterLinkOption()
	if useSessionOpt {
		opts = append(opts, sessionOpt)
	}

	amqpReceiver, err := amqpSession.NewReceiver(opts...)
	if err != nil {
		tab.For(ctx).Error(err)
		return err
	}
	r.receiver = amqpReceiver
	if useSessionOpt {
		rawsid := r.receiver.LinkSourceFilterValue(sessionFilterName)
		if rawsid == nil && r.sessionID == nil {
			return errors.New("failed to create a receiver.  no unlocked sessions available")
		} else if rawsid != nil && r.sessionID != nil && rawsid != *r.sessionID {
			return fmt.Errorf("failed to create a receiver for session %s, it may be locked by another receiver", rawsid)
		} else if r.sessionID == nil {
			sid := rawsid.(string)
			r.sessionID = &sid
		}
	}
	return nil
}

func (r *Receiver) getSessionFilterLinkOption() (amqp.LinkOption, bool) {
	const code = uint64(0x00000137000000C)

	if !r.useSessions {
		return nil, false
	}

	if r.sessionID == nil {
		return amqp.LinkSourceFilter(sessionFilterName, code, nil), true
	}

	return amqp.LinkSourceFilter(sessionFilterName, code, r.sessionID), true
}

func messageID(msg *amqp.Message) interface{} {
	var id interface{} = "null"
	if msg.Properties != nil {
		id = msg.Properties.MessageID
	}
	return id
}

// ListenerHandle is the handle that you can use to wait on a `Listen`
// call to complete. It is extremely similar to a `context.Context`.
type ListenerHandle interface {
	Done() <-chan struct{}
	Err() error
}

// Done will close the channel when the listener has stopped
func (lc *ListenerHandleImpl) Done() <-chan struct{} {
	return lc.ctx.Done()
}

// Err will return the last error encountered
func (lc *ListenerHandleImpl) Err() error {
	lc.r.lastErrorMu.RLock()
	defer lc.r.lastErrorMu.RUnlock()
	if lc.r.lastError != nil {
		return lc.r.lastError
	}
	return lc.ctx.Err()
}
