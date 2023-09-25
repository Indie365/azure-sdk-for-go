//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakestore

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
	cl, err := arm.NewClient(moduleName+".LocationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LocationsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// GetCapability - Gets subscription-level properties and limits for Data Lake Store specified by resource location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-11-01
//   - location - The resource location without whitespace.
//   - options - LocationsClientGetCapabilityOptions contains the optional parameters for the LocationsClient.GetCapability method.
func (client *LocationsClient) GetCapability(ctx context.Context, location string, options *LocationsClientGetCapabilityOptions) (LocationsClientGetCapabilityResponse, error) {
	var err error
	req, err := client.getCapabilityCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsClientGetCapabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationsClientGetCapabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return LocationsClientGetCapabilityResponse{}, err
	}
	resp, err := client.getCapabilityHandleResponse(httpResp)
	return resp, err
}

// getCapabilityCreateRequest creates the GetCapability request.
func (client *LocationsClient) getCapabilityCreateRequest(ctx context.Context, location string, options *LocationsClientGetCapabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/capability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getCapabilityHandleResponse handles the GetCapability response.
func (client *LocationsClient) getCapabilityHandleResponse(resp *http.Response) (LocationsClientGetCapabilityResponse, error) {
	result := LocationsClientGetCapabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityInformation); err != nil {
		return LocationsClientGetCapabilityResponse{}, err
	}
	return result, nil
}

// NewGetUsagePager - Gets the current usage count and the limit for the resources of the location under the subscription.
//
// Generated from API version 2016-11-01
//   - location - The resource location without whitespace.
//   - options - LocationsClientGetUsageOptions contains the optional parameters for the LocationsClient.NewGetUsagePager method.
func (client *LocationsClient) NewGetUsagePager(location string, options *LocationsClientGetUsageOptions) (*runtime.Pager[LocationsClientGetUsageResponse]) {
	return runtime.NewPager(runtime.PagingHandler[LocationsClientGetUsageResponse]{
		More: func(page LocationsClientGetUsageResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *LocationsClientGetUsageResponse) (LocationsClientGetUsageResponse, error) {
			req, err := client.getUsageCreateRequest(ctx, location, options)
			if err != nil {
				return LocationsClientGetUsageResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LocationsClientGetUsageResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationsClientGetUsageResponse{}, runtime.NewResponseError(resp)
			}
			return client.getUsageHandleResponse(resp)
		},
	})
}

// getUsageCreateRequest creates the GetUsage request.
func (client *LocationsClient) getUsageCreateRequest(ctx context.Context, location string, options *LocationsClientGetUsageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataLakeStore/locations/{location}/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUsageHandleResponse handles the GetUsage response.
func (client *LocationsClient) getUsageHandleResponse(resp *http.Response) (LocationsClientGetUsageResponse, error) {
	result := LocationsClientGetUsageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return LocationsClientGetUsageResponse{}, err
	}
	return result, nil
}

