//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

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

// CheckVirtualNetworkSubnetUsageClient contains the methods for the CheckVirtualNetworkSubnetUsage group.
// Don't use this type directly, use NewCheckVirtualNetworkSubnetUsageClient() instead.
type CheckVirtualNetworkSubnetUsageClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCheckVirtualNetworkSubnetUsageClient creates a new instance of CheckVirtualNetworkSubnetUsageClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCheckVirtualNetworkSubnetUsageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CheckVirtualNetworkSubnetUsageClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CheckVirtualNetworkSubnetUsageClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Execute - Get virtual network subnet usage for a given vNet resource id.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The name of the location.
// parameters - The required parameters for creating or updating a server.
// options - CheckVirtualNetworkSubnetUsageClientExecuteOptions contains the optional parameters for the CheckVirtualNetworkSubnetUsageClient.Execute
// method.
func (client *CheckVirtualNetworkSubnetUsageClient) Execute(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *CheckVirtualNetworkSubnetUsageClientExecuteOptions) (CheckVirtualNetworkSubnetUsageClientExecuteResponse, error) {
	req, err := client.executeCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, runtime.NewResponseError(resp)
	}
	return client.executeHandleResponse(resp)
}

// executeCreateRequest creates the Execute request.
func (client *CheckVirtualNetworkSubnetUsageClient) executeCreateRequest(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *CheckVirtualNetworkSubnetUsageClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/checkVirtualNetworkSubnetUsage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// executeHandleResponse handles the Execute response.
func (client *CheckVirtualNetworkSubnetUsageClient) executeHandleResponse(resp *http.Response) (CheckVirtualNetworkSubnetUsageClientExecuteResponse, error) {
	result := CheckVirtualNetworkSubnetUsageClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkSubnetUsageResult); err != nil {
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	return result, nil
}
