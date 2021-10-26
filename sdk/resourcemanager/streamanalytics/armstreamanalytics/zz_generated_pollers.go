//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ClustersCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ClustersCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClustersCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ClustersCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClustersCreateOrUpdateResponse will be returned.
func (p *ClustersCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ClustersCreateOrUpdateResponse, error) {
	respType := ClustersCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Cluster)
	if err != nil {
		return ClustersCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClustersCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ClustersDeletePoller provides polling facilities until the operation reaches a terminal state.
type ClustersDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClustersDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ClustersDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClustersDeleteResponse will be returned.
func (p *ClustersDeletePoller) FinalResponse(ctx context.Context) (ClustersDeleteResponse, error) {
	respType := ClustersDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ClustersDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClustersDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ClustersUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ClustersUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClustersUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ClustersUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClustersUpdateResponse will be returned.
func (p *ClustersUpdatePoller) FinalResponse(ctx context.Context) (ClustersUpdateResponse, error) {
	respType := ClustersUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Cluster)
	if err != nil {
		return ClustersUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClustersUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// FunctionsTestPoller provides polling facilities until the operation reaches a terminal state.
type FunctionsTestPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *FunctionsTestPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *FunctionsTestPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final FunctionsTestResponse will be returned.
func (p *FunctionsTestPoller) FinalResponse(ctx context.Context) (FunctionsTestResponse, error) {
	respType := FunctionsTestResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ResourceTestStatus)
	if err != nil {
		return FunctionsTestResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *FunctionsTestPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// InputsTestPoller provides polling facilities until the operation reaches a terminal state.
type InputsTestPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *InputsTestPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *InputsTestPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final InputsTestResponse will be returned.
func (p *InputsTestPoller) FinalResponse(ctx context.Context) (InputsTestResponse, error) {
	respType := InputsTestResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ResourceTestStatus)
	if err != nil {
		return InputsTestResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *InputsTestPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// OutputsTestPoller provides polling facilities until the operation reaches a terminal state.
type OutputsTestPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *OutputsTestPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *OutputsTestPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final OutputsTestResponse will be returned.
func (p *OutputsTestPoller) FinalResponse(ctx context.Context) (OutputsTestResponse, error) {
	respType := OutputsTestResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ResourceTestStatus)
	if err != nil {
		return OutputsTestResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *OutputsTestPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointsDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointsDeleteResponse will be returned.
func (p *PrivateEndpointsDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointsDeleteResponse, error) {
	respType := PrivateEndpointsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return PrivateEndpointsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StreamingJobsCreateOrReplacePoller provides polling facilities until the operation reaches a terminal state.
type StreamingJobsCreateOrReplacePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StreamingJobsCreateOrReplacePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StreamingJobsCreateOrReplacePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StreamingJobsCreateOrReplaceResponse will be returned.
func (p *StreamingJobsCreateOrReplacePoller) FinalResponse(ctx context.Context) (StreamingJobsCreateOrReplaceResponse, error) {
	respType := StreamingJobsCreateOrReplaceResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.StreamingJob)
	if err != nil {
		return StreamingJobsCreateOrReplaceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StreamingJobsCreateOrReplacePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StreamingJobsDeletePoller provides polling facilities until the operation reaches a terminal state.
type StreamingJobsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StreamingJobsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StreamingJobsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StreamingJobsDeleteResponse will be returned.
func (p *StreamingJobsDeletePoller) FinalResponse(ctx context.Context) (StreamingJobsDeleteResponse, error) {
	respType := StreamingJobsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StreamingJobsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StreamingJobsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StreamingJobsStartPoller provides polling facilities until the operation reaches a terminal state.
type StreamingJobsStartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StreamingJobsStartPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StreamingJobsStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StreamingJobsStartResponse will be returned.
func (p *StreamingJobsStartPoller) FinalResponse(ctx context.Context) (StreamingJobsStartResponse, error) {
	respType := StreamingJobsStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StreamingJobsStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StreamingJobsStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// StreamingJobsStopPoller provides polling facilities until the operation reaches a terminal state.
type StreamingJobsStopPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *StreamingJobsStopPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *StreamingJobsStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final StreamingJobsStopResponse will be returned.
func (p *StreamingJobsStopPoller) FinalResponse(ctx context.Context) (StreamingJobsStopResponse, error) {
	respType := StreamingJobsStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return StreamingJobsStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *StreamingJobsStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionsSampleInputPoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionsSampleInputPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionsSampleInputPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionsSampleInputPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionsSampleInputResponse will be returned.
func (p *SubscriptionsSampleInputPoller) FinalResponse(ctx context.Context) (SubscriptionsSampleInputResponse, error) {
	respType := SubscriptionsSampleInputResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.SampleInputResult)
	if err != nil {
		return SubscriptionsSampleInputResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionsSampleInputPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionsTestInputPoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionsTestInputPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionsTestInputPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionsTestInputPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionsTestInputResponse will be returned.
func (p *SubscriptionsTestInputPoller) FinalResponse(ctx context.Context) (SubscriptionsTestInputResponse, error) {
	respType := SubscriptionsTestInputResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.TestDatasourceResult)
	if err != nil {
		return SubscriptionsTestInputResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionsTestInputPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionsTestOutputPoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionsTestOutputPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionsTestOutputPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionsTestOutputPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionsTestOutputResponse will be returned.
func (p *SubscriptionsTestOutputPoller) FinalResponse(ctx context.Context) (SubscriptionsTestOutputResponse, error) {
	respType := SubscriptionsTestOutputResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.TestDatasourceResult)
	if err != nil {
		return SubscriptionsTestOutputResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionsTestOutputPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubscriptionsTestQueryPoller provides polling facilities until the operation reaches a terminal state.
type SubscriptionsTestQueryPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubscriptionsTestQueryPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SubscriptionsTestQueryPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubscriptionsTestQueryResponse will be returned.
func (p *SubscriptionsTestQueryPoller) FinalResponse(ctx context.Context) (SubscriptionsTestQueryResponse, error) {
	respType := SubscriptionsTestQueryResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.QueryTestingResult)
	if err != nil {
		return SubscriptionsTestQueryResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubscriptionsTestQueryPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
