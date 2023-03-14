//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragecache

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

// ManagementClient contains the methods for the StorageCacheManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
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
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckAmlFSSubnets - Check that subnets will be valid for AML file system create calls.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - options - ManagementClientCheckAmlFSSubnetsOptions contains the optional parameters for the ManagementClient.CheckAmlFSSubnets
//     method.
func (client *ManagementClient) CheckAmlFSSubnets(ctx context.Context, options *ManagementClientCheckAmlFSSubnetsOptions) (ManagementClientCheckAmlFSSubnetsResponse, error) {
	req, err := client.checkAmlFSSubnetsCreateRequest(ctx, options)
	if err != nil {
		return ManagementClientCheckAmlFSSubnetsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementClientCheckAmlFSSubnetsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientCheckAmlFSSubnetsResponse{}, runtime.NewResponseError(resp)
	}
	return ManagementClientCheckAmlFSSubnetsResponse{}, nil
}

// checkAmlFSSubnetsCreateRequest creates the CheckAmlFSSubnets request.
func (client *ManagementClient) checkAmlFSSubnetsCreateRequest(ctx context.Context, options *ManagementClientCheckAmlFSSubnetsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/checkAmlFSSubnets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.AmlFilesystemSubnetInfo != nil {
		return req, runtime.MarshalAsJSON(req, *options.AmlFilesystemSubnetInfo)
	}
	return req, nil
}

// GetRequiredAmlFSSubnetsSize - Get the number of available IP addresses needed for the AML file system information provided.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - options - ManagementClientGetRequiredAmlFSSubnetsSizeOptions contains the optional parameters for the ManagementClient.GetRequiredAmlFSSubnetsSize
//     method.
func (client *ManagementClient) GetRequiredAmlFSSubnetsSize(ctx context.Context, options *ManagementClientGetRequiredAmlFSSubnetsSizeOptions) (ManagementClientGetRequiredAmlFSSubnetsSizeResponse, error) {
	req, err := client.getRequiredAmlFSSubnetsSizeCreateRequest(ctx, options)
	if err != nil {
		return ManagementClientGetRequiredAmlFSSubnetsSizeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementClientGetRequiredAmlFSSubnetsSizeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientGetRequiredAmlFSSubnetsSizeResponse{}, runtime.NewResponseError(resp)
	}
	return client.getRequiredAmlFSSubnetsSizeHandleResponse(resp)
}

// getRequiredAmlFSSubnetsSizeCreateRequest creates the GetRequiredAmlFSSubnetsSize request.
func (client *ManagementClient) getRequiredAmlFSSubnetsSizeCreateRequest(ctx context.Context, options *ManagementClientGetRequiredAmlFSSubnetsSizeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StorageCache/getRequiredAmlFSSubnetsSize"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.RequiredAMLFilesystemSubnetsSizeInfo != nil {
		return req, runtime.MarshalAsJSON(req, *options.RequiredAMLFilesystemSubnetsSizeInfo)
	}
	return req, nil
}

// getRequiredAmlFSSubnetsSizeHandleResponse handles the GetRequiredAmlFSSubnetsSize response.
func (client *ManagementClient) getRequiredAmlFSSubnetsSizeHandleResponse(resp *http.Response) (ManagementClientGetRequiredAmlFSSubnetsSizeResponse, error) {
	result := ManagementClientGetRequiredAmlFSSubnetsSizeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RequiredAmlFilesystemSubnetsSize); err != nil {
		return ManagementClientGetRequiredAmlFSSubnetsSizeResponse{}, err
	}
	return result, nil
}
