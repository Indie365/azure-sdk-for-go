//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRoutePortsLocationsClient contains the methods for the ExpressRoutePortsLocations group.
// Don't use this type directly, use NewExpressRoutePortsLocationsClient() instead.
type ExpressRoutePortsLocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRoutePortsLocationsClient creates a new instance of ExpressRoutePortsLocationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRoutePortsLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExpressRoutePortsLocationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ExpressRoutePortsLocationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Retrieves a single ExpressRoutePort peering location, including the list of available bandwidths available at said
// peering location.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - Name of the requested ExpressRoutePort peering location.
// options - ExpressRoutePortsLocationsClientGetOptions contains the optional parameters for the ExpressRoutePortsLocationsClient.Get
// method.
func (client *ExpressRoutePortsLocationsClient) Get(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsClientGetOptions) (ExpressRoutePortsLocationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, options)
	if err != nil {
		return ExpressRoutePortsLocationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRoutePortsLocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRoutePortsLocationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRoutePortsLocationsClient) getCreateRequest(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRoutePortsLocationsClient) getHandleResponse(resp *http.Response) (ExpressRoutePortsLocationsClientGetResponse, error) {
	result := ExpressRoutePortsLocationsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortsLocation); err != nil {
		return ExpressRoutePortsLocationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Retrieves all ExpressRoutePort peering locations. Does not return available bandwidths for each location. Available
// bandwidths can only be obtained when retrieving a specific peering location.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ExpressRoutePortsLocationsClientListOptions contains the optional parameters for the ExpressRoutePortsLocationsClient.List
// method.
func (client *ExpressRoutePortsLocationsClient) List(options *ExpressRoutePortsLocationsClientListOptions) *ExpressRoutePortsLocationsClientListPager {
	return &ExpressRoutePortsLocationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExpressRoutePortsLocationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortsLocationListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRoutePortsLocationsClient) listCreateRequest(ctx context.Context, options *ExpressRoutePortsLocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRoutePortsLocationsClient) listHandleResponse(resp *http.Response) (ExpressRoutePortsLocationsClientListResponse, error) {
	result := ExpressRoutePortsLocationsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortsLocationListResult); err != nil {
		return ExpressRoutePortsLocationsClientListResponse{}, err
	}
	return result, nil
}
