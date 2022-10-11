//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRouteProviderPortsLocationClient contains the methods for the ExpressRouteProviderPortsLocation group.
// Don't use this type directly, use NewExpressRouteProviderPortsLocationClient() instead.
type ExpressRouteProviderPortsLocationClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteProviderPortsLocationClient creates a new instance of ExpressRouteProviderPortsLocationClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteProviderPortsLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteProviderPortsLocationClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRouteProviderPortsLocationClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Retrieves all the ExpressRouteProviderPorts in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// options - ExpressRouteProviderPortsLocationClientListOptions contains the optional parameters for the ExpressRouteProviderPortsLocationClient.List
// method.
func (client *ExpressRouteProviderPortsLocationClient) List(ctx context.Context, options *ExpressRouteProviderPortsLocationClientListOptions) (ExpressRouteProviderPortsLocationClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return ExpressRouteProviderPortsLocationClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteProviderPortsLocationClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteProviderPortsLocationClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ExpressRouteProviderPortsLocationClient) listCreateRequest(ctx context.Context, options *ExpressRouteProviderPortsLocationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteProviderPorts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteProviderPortsLocationClient) listHandleResponse(resp *http.Response) (ExpressRouteProviderPortsLocationClientListResponse, error) {
	result := ExpressRouteProviderPortsLocationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteProviderPortListResult); err != nil {
		return ExpressRouteProviderPortsLocationClientListResponse{}, err
	}
	return result, nil
}
