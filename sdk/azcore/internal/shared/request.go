//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"golang.org/x/net/http/httpguts"
)

// Policy represents an extensibility point for the Pipeline that can mutate the specified
// Request and react to the received Response.
type Policy interface {
	// Do applies the policy to the specified Request.  When implementing a Policy, mutate the
	// request before calling req.Next() to move on to the next policy, and respond to the result
	// before returning to the caller.
	Do(req *Request) (*http.Response, error)
}

// Request is an abstraction over the creation of an HTTP request as it passes through the pipeline.
// Don't use this type directly, use NewRequest() instead.
type Request struct {
	req      *http.Request
	body     io.ReadSeekCloser
	policies []Policy
	values   opValues
}

type opValues map[reflect.Type]interface{}

// Set adds/changes a value
func (ov opValues) set(value interface{}) {
	ov[reflect.TypeOf(value)] = value
}

// Get looks for a value set by SetValue first
func (ov opValues) get(value interface{}) bool {
	v, ok := ov[reflect.ValueOf(value).Elem().Type()]
	if ok {
		reflect.ValueOf(value).Elem().Set(reflect.ValueOf(v))
	}
	return ok
}

// NewRequest creates a new Request with the specified input.
func NewRequest(ctx context.Context, httpMethod string, endpoint string) (*Request, error) {
	req, err := http.NewRequestWithContext(ctx, httpMethod, endpoint, nil)
	if err != nil {
		return nil, err
	}
	if req.URL.Host == "" {
		return nil, errors.New("no Host in request URL")
	}
	if !(req.URL.Scheme == "http" || req.URL.Scheme == "https") {
		return nil, fmt.Errorf("unsupported protocol scheme %s", req.URL.Scheme)
	}
	// check copied from Transport.roundTrip()
	for k, vv := range req.Header {
		if !httpguts.ValidHeaderFieldName(k) {
			req.Body.Close()
			return nil, fmt.Errorf("invalid header field name %q", k)
		}
		for _, v := range vv {
			if !httpguts.ValidHeaderFieldValue(v) {
				req.Body.Close()
				return nil, fmt.Errorf("invalid header field value %q for key %v", v, k)
			}
		}
	}
	return &Request{req: req}, nil
}

// SetPolicies assigns the policies to the Request (INTERNAL ONLY).
func SetPolicies(req *Request, policies []Policy) {
	req.policies = policies
}

// Body returns the original body specified when the Request was created.
func (req *Request) Body() io.ReadSeekCloser {
	return req.body
}

// Raw returns the underlying HTTP request.
func (req *Request) Raw() *http.Request {
	return req.req
}

// Next calls the next policy in the pipeline.
// If there are no more policies, nil and an error are returned.
// This method is intended to be called from pipeline policies.
// To send a request through a pipeline call Pipeline.Do().
func (req *Request) Next() (*http.Response, error) {
	if len(req.policies) == 0 {
		return nil, errors.New("no more policies")
	}
	nextPolicy := req.policies[0]
	nextReq := *req
	nextReq.policies = nextReq.policies[1:]
	return nextPolicy.Do(&nextReq)
}

// SetOperationValue adds/changes a mutable key/value associated with a single operation.
func (req *Request) SetOperationValue(value interface{}) {
	if req.values == nil {
		req.values = opValues{}
	}
	req.values.set(value)
}

// OperationValue looks for a value set by SetOperationValue().
func (req *Request) OperationValue(value interface{}) bool {
	if req.values == nil {
		return false
	}
	return req.values.get(value)
}

// SetBody sets the specified ReadSeekCloser as the HTTP request body.
func (req *Request) SetBody(body io.ReadSeekCloser, contentType string) error {
	// Set the body and content length.
	size, err := body.Seek(0, io.SeekEnd) // Seek to the end to get the stream's size
	if err != nil {
		return err
	}
	if size == 0 {
		body.Close()
		return nil
	}
	_, err = body.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	req.Raw().GetBody = func() (io.ReadCloser, error) {
		_, err := body.Seek(0, io.SeekStart) // Seek back to the beginning of the stream
		return body, err
	}
	// keep a copy of the original body.  this is to handle cases
	// where req.Body is replaced, e.g. httputil.DumpRequest and friends.
	req.body = body
	req.req.Body = body
	req.req.ContentLength = size
	req.req.Header.Set(HeaderContentType, contentType)
	req.req.Header.Set(HeaderContentLength, strconv.FormatInt(size, 10))
	return nil
}

// SkipBodyDownload will disable automatic downloading of the response body.
func (req *Request) SkipBodyDownload() {
	req.SetOperationValue(BodyDownloadPolicyOpValues{Skip: true})
}

// RewindBody seeks the request's Body stream back to the beginning so it can be resent when retrying an operation.
func (req *Request) RewindBody() error {
	if req.body != nil {
		// Reset the stream back to the beginning and restore the body
		_, err := req.body.Seek(0, io.SeekStart)
		req.req.Body = req.body
		return err
	}
	return nil
}

// Close closes the request body.
func (req *Request) Close() error {
	if req.body == nil {
		return nil
	}
	return req.body.Close()
}

// Clone returns a deep copy of the request with its context changed to ctx.
func (req *Request) Clone(ctx context.Context) *Request {
	r2 := Request{}
	r2 = *req
	r2.req = req.req.Clone(ctx)
	return &r2
}
