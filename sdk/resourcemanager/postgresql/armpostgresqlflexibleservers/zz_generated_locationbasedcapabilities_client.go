//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

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

// LocationBasedCapabilitiesClient contains the methods for the LocationBasedCapabilities group.
// Don't use this type directly, use NewLocationBasedCapabilitiesClient() instead.
type LocationBasedCapabilitiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationBasedCapabilitiesClient creates a new instance of LocationBasedCapabilitiesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationBasedCapabilitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LocationBasedCapabilitiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &LocationBasedCapabilitiesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Execute - Get capabilities at specified location in a given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The name of the location.
// options - LocationBasedCapabilitiesClientExecuteOptions contains the optional parameters for the LocationBasedCapabilitiesClient.Execute
// method.
func (client *LocationBasedCapabilitiesClient) Execute(locationName string, options *LocationBasedCapabilitiesClientExecuteOptions) *LocationBasedCapabilitiesClientExecutePager {
	return &LocationBasedCapabilitiesClientExecutePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.executeCreateRequest(ctx, locationName, options)
		},
		advancer: func(ctx context.Context, resp LocationBasedCapabilitiesClientExecuteResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CapabilitiesListResult.NextLink)
		},
	}
}

// executeCreateRequest creates the Execute request.
func (client *LocationBasedCapabilitiesClient) executeCreateRequest(ctx context.Context, locationName string, options *LocationBasedCapabilitiesClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/locations/{locationName}/capabilities"
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *LocationBasedCapabilitiesClient) executeHandleResponse(resp *http.Response) (LocationBasedCapabilitiesClientExecuteResponse, error) {
	result := LocationBasedCapabilitiesClientExecuteResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilitiesListResult); err != nil {
		return LocationBasedCapabilitiesClientExecuteResponse{}, err
	}
	return result, nil
}
