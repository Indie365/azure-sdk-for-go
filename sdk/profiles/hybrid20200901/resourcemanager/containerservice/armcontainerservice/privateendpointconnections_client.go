//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
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
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginDelete - Deletes the private endpoint connection in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01
// resourceGroupName - The name of the resource group.
// resourceName - The name of the managed cluster resource.
// privateEndpointConnectionName - The name of the private endpoint connection.
// options - PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PrivateEndpointConnectionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[PrivateEndpointConnectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the private endpoint connection in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01
func (client *PrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the private endpoint connection by managed cluster and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01
// resourceGroupName - The name of the resource group.
// resourceName - The name of the managed cluster resource.
// privateEndpointConnectionName - The name of the private endpoint connection.
// options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of private endpoint connections in the specified managed cluster. The operation returns properties of
// each private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01
// resourceGroupName - The name of the resource group.
// resourceName - The name of the managed cluster resource.
// options - PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.List
// method.
func (client *PrivateEndpointConnectionsClient) List(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateEndpointConnectionsClientListOptions) (PrivateEndpointConnectionsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *PrivateEndpointConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateEndpointConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateEndpointConnectionsClient) listHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListResponse, error) {
	result := PrivateEndpointConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates a private endpoint connection in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01
// resourceGroupName - The name of the resource group.
// resourceName - The name of the managed cluster resource.
// privateEndpointConnectionName - The name of the private endpoint connection.
// parameters - Parameters supplied to the Update a private endpoint connection operation.
// options - PrivateEndpointConnectionsClientUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Update
// method.
func (client *PrivateEndpointConnectionsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, options *PrivateEndpointConnectionsClientUpdateOptions) (PrivateEndpointConnectionsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, privateEndpointConnectionName, parameters, options)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PrivateEndpointConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, options *PrivateEndpointConnectionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *PrivateEndpointConnectionsClient) updateHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientUpdateResponse, error) {
	result := PrivateEndpointConnectionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientUpdateResponse{}, err
	}
	return result, nil
}
