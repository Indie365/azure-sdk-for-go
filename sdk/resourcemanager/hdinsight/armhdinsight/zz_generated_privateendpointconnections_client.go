//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndpointConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Approve or reject a private endpoint connection manually.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// privateEndpointConnectionName - The name of the private endpoint connection.
// parameters - The private endpoint connection create or update request.
// options - PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginCreateOrUpdate
// method.
func (client *PrivateEndpointConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, privateEndpointConnectionName, parameters, options)
	if err != nil {
		return PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateEndpointConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return PrivateEndpointConnectionsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PrivateEndpointConnectionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Approve or reject a private endpoint connection manually.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, privateEndpointConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateEndpointConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, parameters PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specific private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// privateEndpointConnectionName - The name of the private endpoint connection.
// options - PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (PrivateEndpointConnectionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientDeletePollerResponse{}, err
	}
	result := PrivateEndpointConnectionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateEndpointConnectionsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return PrivateEndpointConnectionsClientDeletePollerResponse{}, err
	}
	result.Poller = &PrivateEndpointConnectionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specific private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specific private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// privateEndpointConnectionName - The name of the private endpoint connection.
// options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, privateEndpointConnectionName, options)
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
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
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

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByCluster - Lists the private endpoint connections for a HDInsight cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster.
// options - PrivateEndpointConnectionsClientListByClusterOptions contains the optional parameters for the PrivateEndpointConnectionsClient.ListByCluster
// method.
func (client *PrivateEndpointConnectionsClient) ListByCluster(resourceGroupName string, clusterName string, options *PrivateEndpointConnectionsClientListByClusterOptions) *PrivateEndpointConnectionsClientListByClusterPager {
	return &PrivateEndpointConnectionsClientListByClusterPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
		},
		advancer: func(ctx context.Context, resp PrivateEndpointConnectionsClientListByClusterResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointConnectionListResult.NextLink)
		},
	}
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *PrivateEndpointConnectionsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *PrivateEndpointConnectionsClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusters/{clusterName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
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

// listByClusterHandleResponse handles the ListByCluster response.
func (client *PrivateEndpointConnectionsClient) listByClusterHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByClusterResponse, error) {
	result := PrivateEndpointConnectionsClientListByClusterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByClusterResponse{}, err
	}
	return result, nil
}
