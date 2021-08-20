//go:build go1.13
// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// DataFlowDebugSessionCreatePoller provides polling facilities until the operation reaches a terminal state.
type DataFlowDebugSessionCreatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DataFlowDebugSessionCreateResponse will be returned.
	FinalResponse(ctx context.Context) (DataFlowDebugSessionCreateResponse, error)
}

type dataFlowDebugSessionCreatePoller struct {
	pt *armcore.LROPoller
}

func (p *dataFlowDebugSessionCreatePoller) Done() bool {
	return p.pt.Done()
}

func (p *dataFlowDebugSessionCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *dataFlowDebugSessionCreatePoller) FinalResponse(ctx context.Context) (DataFlowDebugSessionCreateResponse, error) {
	respType := DataFlowDebugSessionCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.CreateDataFlowDebugSessionResponse)
	if err != nil {
		return DataFlowDebugSessionCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *dataFlowDebugSessionCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *dataFlowDebugSessionCreatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (DataFlowDebugSessionCreateResponse, error) {
	respType := DataFlowDebugSessionCreateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.CreateDataFlowDebugSessionResponse)
	if err != nil {
		return DataFlowDebugSessionCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// DataFlowDebugSessionExecuteCommandPoller provides polling facilities until the operation reaches a terminal state.
type DataFlowDebugSessionExecuteCommandPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DataFlowDebugSessionExecuteCommandResponse will be returned.
	FinalResponse(ctx context.Context) (DataFlowDebugSessionExecuteCommandResponse, error)
}

type dataFlowDebugSessionExecuteCommandPoller struct {
	pt *armcore.LROPoller
}

func (p *dataFlowDebugSessionExecuteCommandPoller) Done() bool {
	return p.pt.Done()
}

func (p *dataFlowDebugSessionExecuteCommandPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *dataFlowDebugSessionExecuteCommandPoller) FinalResponse(ctx context.Context) (DataFlowDebugSessionExecuteCommandResponse, error) {
	respType := DataFlowDebugSessionExecuteCommandResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DataFlowDebugCommandResponse)
	if err != nil {
		return DataFlowDebugSessionExecuteCommandResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *dataFlowDebugSessionExecuteCommandPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *dataFlowDebugSessionExecuteCommandPoller) pollUntilDone(ctx context.Context, freq time.Duration) (DataFlowDebugSessionExecuteCommandResponse, error) {
	respType := DataFlowDebugSessionExecuteCommandResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.DataFlowDebugCommandResponse)
	if err != nil {
		return DataFlowDebugSessionExecuteCommandResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// IntegrationRuntimeObjectMetadataRefreshPoller provides polling facilities until the operation reaches a terminal state.
type IntegrationRuntimeObjectMetadataRefreshPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final IntegrationRuntimeObjectMetadataRefreshResponse will be returned.
	FinalResponse(ctx context.Context) (IntegrationRuntimeObjectMetadataRefreshResponse, error)
}

type integrationRuntimeObjectMetadataRefreshPoller struct {
	pt *armcore.LROPoller
}

func (p *integrationRuntimeObjectMetadataRefreshPoller) Done() bool {
	return p.pt.Done()
}

func (p *integrationRuntimeObjectMetadataRefreshPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *integrationRuntimeObjectMetadataRefreshPoller) FinalResponse(ctx context.Context) (IntegrationRuntimeObjectMetadataRefreshResponse, error) {
	respType := IntegrationRuntimeObjectMetadataRefreshResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.SsisObjectMetadataStatusResponse)
	if err != nil {
		return IntegrationRuntimeObjectMetadataRefreshResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *integrationRuntimeObjectMetadataRefreshPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *integrationRuntimeObjectMetadataRefreshPoller) pollUntilDone(ctx context.Context, freq time.Duration) (IntegrationRuntimeObjectMetadataRefreshResponse, error) {
	respType := IntegrationRuntimeObjectMetadataRefreshResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.SsisObjectMetadataStatusResponse)
	if err != nil {
		return IntegrationRuntimeObjectMetadataRefreshResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// IntegrationRuntimesStartPoller provides polling facilities until the operation reaches a terminal state.
type IntegrationRuntimesStartPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final IntegrationRuntimesStartResponse will be returned.
	FinalResponse(ctx context.Context) (IntegrationRuntimesStartResponse, error)
}

type integrationRuntimesStartPoller struct {
	pt *armcore.LROPoller
}

func (p *integrationRuntimesStartPoller) Done() bool {
	return p.pt.Done()
}

func (p *integrationRuntimesStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *integrationRuntimesStartPoller) FinalResponse(ctx context.Context) (IntegrationRuntimesStartResponse, error) {
	respType := IntegrationRuntimesStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.IntegrationRuntimeStatusResponse)
	if err != nil {
		return IntegrationRuntimesStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *integrationRuntimesStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *integrationRuntimesStartPoller) pollUntilDone(ctx context.Context, freq time.Duration) (IntegrationRuntimesStartResponse, error) {
	respType := IntegrationRuntimesStartResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.IntegrationRuntimeStatusResponse)
	if err != nil {
		return IntegrationRuntimesStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// IntegrationRuntimesStopPoller provides polling facilities until the operation reaches a terminal state.
type IntegrationRuntimesStopPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final IntegrationRuntimesStopResponse will be returned.
	FinalResponse(ctx context.Context) (IntegrationRuntimesStopResponse, error)
}

type integrationRuntimesStopPoller struct {
	pt *armcore.LROPoller
}

func (p *integrationRuntimesStopPoller) Done() bool {
	return p.pt.Done()
}

func (p *integrationRuntimesStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *integrationRuntimesStopPoller) FinalResponse(ctx context.Context) (IntegrationRuntimesStopResponse, error) {
	respType := IntegrationRuntimesStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return IntegrationRuntimesStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *integrationRuntimesStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *integrationRuntimesStopPoller) pollUntilDone(ctx context.Context, freq time.Duration) (IntegrationRuntimesStopResponse, error) {
	respType := IntegrationRuntimesStopResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return IntegrationRuntimesStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// TriggersStartPoller provides polling facilities until the operation reaches a terminal state.
type TriggersStartPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final TriggersStartResponse will be returned.
	FinalResponse(ctx context.Context) (TriggersStartResponse, error)
}

type triggersStartPoller struct {
	pt *armcore.LROPoller
}

func (p *triggersStartPoller) Done() bool {
	return p.pt.Done()
}

func (p *triggersStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *triggersStartPoller) FinalResponse(ctx context.Context) (TriggersStartResponse, error) {
	respType := TriggersStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return TriggersStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *triggersStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *triggersStartPoller) pollUntilDone(ctx context.Context, freq time.Duration) (TriggersStartResponse, error) {
	respType := TriggersStartResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return TriggersStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// TriggersStopPoller provides polling facilities until the operation reaches a terminal state.
type TriggersStopPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final TriggersStopResponse will be returned.
	FinalResponse(ctx context.Context) (TriggersStopResponse, error)
}

type triggersStopPoller struct {
	pt *armcore.LROPoller
}

func (p *triggersStopPoller) Done() bool {
	return p.pt.Done()
}

func (p *triggersStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *triggersStopPoller) FinalResponse(ctx context.Context) (TriggersStopResponse, error) {
	respType := TriggersStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return TriggersStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *triggersStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *triggersStopPoller) pollUntilDone(ctx context.Context, freq time.Duration) (TriggersStopResponse, error) {
	respType := TriggersStopResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return TriggersStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// TriggersSubscribeToEventsPoller provides polling facilities until the operation reaches a terminal state.
type TriggersSubscribeToEventsPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final TriggersSubscribeToEventsResponse will be returned.
	FinalResponse(ctx context.Context) (TriggersSubscribeToEventsResponse, error)
}

type triggersSubscribeToEventsPoller struct {
	pt *armcore.LROPoller
}

func (p *triggersSubscribeToEventsPoller) Done() bool {
	return p.pt.Done()
}

func (p *triggersSubscribeToEventsPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *triggersSubscribeToEventsPoller) FinalResponse(ctx context.Context) (TriggersSubscribeToEventsResponse, error) {
	respType := TriggersSubscribeToEventsResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.TriggerSubscriptionOperationStatus)
	if err != nil {
		return TriggersSubscribeToEventsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *triggersSubscribeToEventsPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *triggersSubscribeToEventsPoller) pollUntilDone(ctx context.Context, freq time.Duration) (TriggersSubscribeToEventsResponse, error) {
	respType := TriggersSubscribeToEventsResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.TriggerSubscriptionOperationStatus)
	if err != nil {
		return TriggersSubscribeToEventsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// TriggersUnsubscribeFromEventsPoller provides polling facilities until the operation reaches a terminal state.
type TriggersUnsubscribeFromEventsPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final TriggersUnsubscribeFromEventsResponse will be returned.
	FinalResponse(ctx context.Context) (TriggersUnsubscribeFromEventsResponse, error)
}

type triggersUnsubscribeFromEventsPoller struct {
	pt *armcore.LROPoller
}

func (p *triggersUnsubscribeFromEventsPoller) Done() bool {
	return p.pt.Done()
}

func (p *triggersUnsubscribeFromEventsPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *triggersUnsubscribeFromEventsPoller) FinalResponse(ctx context.Context) (TriggersUnsubscribeFromEventsResponse, error) {
	respType := TriggersUnsubscribeFromEventsResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.TriggerSubscriptionOperationStatus)
	if err != nil {
		return TriggersUnsubscribeFromEventsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *triggersUnsubscribeFromEventsPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *triggersUnsubscribeFromEventsPoller) pollUntilDone(ctx context.Context, freq time.Duration) (TriggersUnsubscribeFromEventsResponse, error) {
	respType := TriggersUnsubscribeFromEventsResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.TriggerSubscriptionOperationStatus)
	if err != nil {
		return TriggersUnsubscribeFromEventsResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
