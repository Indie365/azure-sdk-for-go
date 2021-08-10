// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// ClustersCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ClustersCreateOrUpdatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final ClustersCreateOrUpdateResponse will be returned.
	FinalResponse(ctx context.Context) (ClustersCreateOrUpdateResponse, error)
}

type clustersCreateOrUpdatePoller struct {
	pt *armcore.LROPoller
}

func (p *clustersCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

func (p *clustersCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *clustersCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ClustersCreateOrUpdateResponse, error) {
	respType := ClustersCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Cluster)
	if err != nil {
		return ClustersCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *clustersCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *clustersCreateOrUpdatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (ClustersCreateOrUpdateResponse, error) {
	respType := ClustersCreateOrUpdateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return ClustersCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ClustersDeletePoller provides polling facilities until the operation reaches a terminal state.
type ClustersDeletePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final ClustersDeleteResponse will be returned.
	FinalResponse(ctx context.Context) (ClustersDeleteResponse, error)
}

type clustersDeletePoller struct {
	pt *armcore.LROPoller
}

func (p *clustersDeletePoller) Done() bool {
	return p.pt.Done()
}

func (p *clustersDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *clustersDeletePoller) FinalResponse(ctx context.Context) (ClustersDeleteResponse, error) {
	respType := ClustersDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ClustersDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *clustersDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *clustersDeletePoller) pollUntilDone(ctx context.Context, freq time.Duration) (ClustersDeleteResponse, error) {
	respType := ClustersDeleteResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return ClustersDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ClustersUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ClustersUpdatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final ClustersUpdateResponse will be returned.
	FinalResponse(ctx context.Context) (ClustersUpdateResponse, error)
}

type clustersUpdatePoller struct {
	pt *armcore.LROPoller
}

func (p *clustersUpdatePoller) Done() bool {
	return p.pt.Done()
}

func (p *clustersUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *clustersUpdatePoller) FinalResponse(ctx context.Context) (ClustersUpdateResponse, error) {
	respType := ClustersUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Cluster)
	if err != nil {
		return ClustersUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *clustersUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *clustersUpdatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (ClustersUpdateResponse, error) {
	respType := ClustersUpdateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.Cluster)
	if err != nil {
		return ClustersUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// NamespacesCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type NamespacesCreateOrUpdatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final NamespacesCreateOrUpdateResponse will be returned.
	FinalResponse(ctx context.Context) (NamespacesCreateOrUpdateResponse, error)
}

type namespacesCreateOrUpdatePoller struct {
	pt *armcore.LROPoller
}

func (p *namespacesCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

func (p *namespacesCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *namespacesCreateOrUpdatePoller) FinalResponse(ctx context.Context) (NamespacesCreateOrUpdateResponse, error) {
	respType := NamespacesCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.EHNamespace)
	if err != nil {
		return NamespacesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *namespacesCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *namespacesCreateOrUpdatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (NamespacesCreateOrUpdateResponse, error) {
	respType := NamespacesCreateOrUpdateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.EHNamespace)
	if err != nil {
		return NamespacesCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// NamespacesDeletePoller provides polling facilities until the operation reaches a terminal state.
type NamespacesDeletePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final NamespacesDeleteResponse will be returned.
	FinalResponse(ctx context.Context) (NamespacesDeleteResponse, error)
}

type namespacesDeletePoller struct {
	pt *armcore.LROPoller
}

func (p *namespacesDeletePoller) Done() bool {
	return p.pt.Done()
}

func (p *namespacesDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *namespacesDeletePoller) FinalResponse(ctx context.Context) (NamespacesDeleteResponse, error) {
	respType := NamespacesDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return NamespacesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *namespacesDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *namespacesDeletePoller) pollUntilDone(ctx context.Context, freq time.Duration) (NamespacesDeleteResponse, error) {
	respType := NamespacesDeleteResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return NamespacesDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// PrivateEndpointConnectionsDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsDeletePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final PrivateEndpointConnectionsDeleteResponse will be returned.
	FinalResponse(ctx context.Context) (PrivateEndpointConnectionsDeleteResponse, error)
}

type privateEndpointConnectionsDeletePoller struct {
	pt *armcore.LROPoller
}

func (p *privateEndpointConnectionsDeletePoller) Done() bool {
	return p.pt.Done()
}

func (p *privateEndpointConnectionsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *privateEndpointConnectionsDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return PrivateEndpointConnectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *privateEndpointConnectionsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *privateEndpointConnectionsDeletePoller) pollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return PrivateEndpointConnectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
