// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRoutePortsLocationsClient contains the methods for the ExpressRoutePortsLocations group.
// Don't use this type directly, use NewExpressRoutePortsLocationsClient() instead.
type ExpressRoutePortsLocationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRoutePortsLocationsClient creates a new instance of ExpressRoutePortsLocationsClient with the specified values.
func NewExpressRoutePortsLocationsClient(con *armcore.Connection, subscriptionID string) *ExpressRoutePortsLocationsClient {
	return &ExpressRoutePortsLocationsClient{con: con, subscriptionID: subscriptionID}
}

// Get - Retrieves a single ExpressRoutePort peering location, including the list of available bandwidths available at said peering location.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsLocationsClient) Get(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsGetOptions) (ExpressRoutePortsLocationResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, options)
	if err != nil {
		return ExpressRoutePortsLocationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRoutePortsLocationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRoutePortsLocationResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRoutePortsLocationsClient) getCreateRequest(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRoutePortsLocationsClient) getHandleResponse(resp *azcore.Response) (ExpressRoutePortsLocationResponse, error) {
	var val *ExpressRoutePortsLocation
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRoutePortsLocationResponse{}, err
	}
	return ExpressRoutePortsLocationResponse{RawResponse: resp.Response, ExpressRoutePortsLocation: val}, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRoutePortsLocationsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Retrieves all ExpressRoutePort peering locations. Does not return available bandwidths for each location. Available bandwidths can only be obtained
// when retrieving a specific peering location.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsLocationsClient) List(options *ExpressRoutePortsLocationsListOptions) ExpressRoutePortsLocationListResultPager {
	return &expressRoutePortsLocationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ExpressRoutePortsLocationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortsLocationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRoutePortsLocationsClient) listCreateRequest(ctx context.Context, options *ExpressRoutePortsLocationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRoutePortsLocationsClient) listHandleResponse(resp *azcore.Response) (ExpressRoutePortsLocationListResultResponse, error) {
	var val *ExpressRoutePortsLocationListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRoutePortsLocationListResultResponse{}, err
	}
	return ExpressRoutePortsLocationListResultResponse{RawResponse: resp.Response, ExpressRoutePortsLocationListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ExpressRoutePortsLocationsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
