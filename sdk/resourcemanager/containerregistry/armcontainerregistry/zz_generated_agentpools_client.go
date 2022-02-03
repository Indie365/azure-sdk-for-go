//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

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

// AgentPoolsClient contains the methods for the AgentPools group.
// Don't use this type directly, use NewAgentPoolsClient() instead.
type AgentPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAgentPoolsClient creates a new instance of AgentPoolsClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAgentPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AgentPoolsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AgentPoolsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Creates an agent pool for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// agentPoolName - The name of the agent pool.
// agentPool - The parameters of an agent pool that needs to scheduled.
// options - AgentPoolsClientBeginCreateOptions contains the optional parameters for the AgentPoolsClient.BeginCreate method.
func (client *AgentPoolsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool, options *AgentPoolsClientBeginCreateOptions) (AgentPoolsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, agentPoolName, agentPool, options)
	if err != nil {
		return AgentPoolsClientCreatePollerResponse{}, err
	}
	result := AgentPoolsClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AgentPoolsClient.Create", "", resp, client.pl)
	if err != nil {
		return AgentPoolsClientCreatePollerResponse{}, err
	}
	result.Poller = &AgentPoolsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates an agent pool for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AgentPoolsClient) create(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool, options *AgentPoolsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, agentPool, options)
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

// createCreateRequest creates the Create request.
func (client *AgentPoolsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool, options *AgentPoolsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, agentPool)
}

// BeginDelete - Deletes a specified agent pool resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// agentPoolName - The name of the agent pool.
// options - AgentPoolsClientBeginDeleteOptions contains the optional parameters for the AgentPoolsClient.BeginDelete method.
func (client *AgentPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (AgentPoolsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientDeletePollerResponse{}, err
	}
	result := AgentPoolsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AgentPoolsClient.Delete", "", resp, client.pl)
	if err != nil {
		return AgentPoolsClientDeletePollerResponse{}, err
	}
	result.Poller = &AgentPoolsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a specified agent pool resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AgentPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, options)
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
func (client *AgentPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the detailed information for a given agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// agentPoolName - The name of the agent pool.
// options - AgentPoolsClientGetOptions contains the optional parameters for the AgentPoolsClient.Get method.
func (client *AgentPoolsClient) Get(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetOptions) (AgentPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgentPoolsClient) getHandleResponse(resp *http.Response) (AgentPoolsClientGetResponse, error) {
	result := AgentPoolsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPool); err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	return result, nil
}

// GetQueueStatus - Gets the count of queued runs for a given agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// agentPoolName - The name of the agent pool.
// options - AgentPoolsClientGetQueueStatusOptions contains the optional parameters for the AgentPoolsClient.GetQueueStatus
// method.
func (client *AgentPoolsClient) GetQueueStatus(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetQueueStatusOptions) (AgentPoolsClientGetQueueStatusResponse, error) {
	req, err := client.getQueueStatusCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetQueueStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AgentPoolsClientGetQueueStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentPoolsClientGetQueueStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getQueueStatusHandleResponse(resp)
}

// getQueueStatusCreateRequest creates the GetQueueStatus request.
func (client *AgentPoolsClient) getQueueStatusCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetQueueStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}/listQueueStatus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getQueueStatusHandleResponse handles the GetQueueStatus response.
func (client *AgentPoolsClient) getQueueStatusHandleResponse(resp *http.Response) (AgentPoolsClientGetQueueStatusResponse, error) {
	result := AgentPoolsClientGetQueueStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolQueueStatus); err != nil {
		return AgentPoolsClientGetQueueStatusResponse{}, err
	}
	return result, nil
}

// List - Lists all the agent pools for a specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// options - AgentPoolsClientListOptions contains the optional parameters for the AgentPoolsClient.List method.
func (client *AgentPoolsClient) List(resourceGroupName string, registryName string, options *AgentPoolsClientListOptions) *AgentPoolsClientListPager {
	return &AgentPoolsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		advancer: func(ctx context.Context, resp AgentPoolsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AgentPoolListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AgentPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *AgentPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AgentPoolsClient) listHandleResponse(resp *http.Response) (AgentPoolsClientListResponse, error) {
	result := AgentPoolsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolListResult); err != nil {
		return AgentPoolsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an agent pool with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// agentPoolName - The name of the agent pool.
// updateParameters - The parameters for updating an agent pool.
// options - AgentPoolsClientBeginUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginUpdate method.
func (client *AgentPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters, options *AgentPoolsClientBeginUpdateOptions) (AgentPoolsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, agentPoolName, updateParameters, options)
	if err != nil {
		return AgentPoolsClientUpdatePollerResponse{}, err
	}
	result := AgentPoolsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("AgentPoolsClient.Update", "", resp, client.pl)
	if err != nil {
		return AgentPoolsClientUpdatePollerResponse{}, err
	}
	result.Poller = &AgentPoolsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an agent pool with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AgentPoolsClient) update(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters, options *AgentPoolsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, updateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *AgentPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters, options *AgentPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, updateParameters)
}
