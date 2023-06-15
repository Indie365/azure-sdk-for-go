//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

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

// AgentPoolsClient contains the methods for the AgentPools group.
// Don't use this type directly, use NewAgentPoolsClient() instead.
type AgentPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAgentPoolsClient creates a new instance of AgentPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAgentPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AgentPoolsClient, error) {
	cl, err := arm.NewClient(moduleName+".AgentPoolsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AgentPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates an agent pool for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - agentPoolName - The name of the agent pool.
//   - agentPool - The parameters of an agent pool that needs to scheduled.
//   - options - AgentPoolsClientBeginCreateOptions contains the optional parameters for the AgentPoolsClient.BeginCreate method.
func (client *AgentPoolsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool, options *AgentPoolsClientBeginCreateOptions) (*runtime.Poller[AgentPoolsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, agentPoolName, agentPool, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AgentPoolsClientCreateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates an agent pool for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *AgentPoolsClient) create(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool, options *AgentPoolsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, agentPool, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *AgentPoolsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, agentPool AgentPool, options *AgentPoolsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, agentPool); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a specified agent pool resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientBeginDeleteOptions contains the optional parameters for the AgentPoolsClient.BeginDelete method.
func (client *AgentPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*runtime.Poller[AgentPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, agentPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AgentPoolsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a specified agent pool resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *AgentPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AgentPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the detailed information for a given agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientGetOptions contains the optional parameters for the AgentPoolsClient.Get method.
func (client *AgentPoolsClient) Get(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetOptions) (AgentPoolsClientGetResponse, error) {
	var err error
	const operationName = "AgentPoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AgentPoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AgentPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AgentPoolsClient) getHandleResponse(resp *http.Response) (AgentPoolsClientGetResponse, error) {
	result := AgentPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPool); err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	return result, nil
}

// GetQueueStatus - Gets the count of queued runs for a given agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientGetQueueStatusOptions contains the optional parameters for the AgentPoolsClient.GetQueueStatus
//     method.
func (client *AgentPoolsClient) GetQueueStatus(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetQueueStatusOptions) (AgentPoolsClientGetQueueStatusResponse, error) {
	var err error
	const operationName = "AgentPoolsClient.GetQueueStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getQueueStatusCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetQueueStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetQueueStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AgentPoolsClientGetQueueStatusResponse{}, err
	}
	resp, err := client.getQueueStatusHandleResponse(httpResp)
	return resp, err
}

// getQueueStatusCreateRequest creates the GetQueueStatus request.
func (client *AgentPoolsClient) getQueueStatusCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, options *AgentPoolsClientGetQueueStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}/listQueueStatus"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getQueueStatusHandleResponse handles the GetQueueStatus response.
func (client *AgentPoolsClient) getQueueStatusHandleResponse(resp *http.Response) (AgentPoolsClientGetQueueStatusResponse, error) {
	result := AgentPoolsClientGetQueueStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolQueueStatus); err != nil {
		return AgentPoolsClientGetQueueStatusResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the agent pools for a specified container registry.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - AgentPoolsClientListOptions contains the optional parameters for the AgentPoolsClient.NewListPager method.
func (client *AgentPoolsClient) NewListPager(resourceGroupName string, registryName string, options *AgentPoolsClientListOptions) *runtime.Pager[AgentPoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AgentPoolsClientListResponse]{
		More: func(page AgentPoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AgentPoolsClientListResponse) (AgentPoolsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AgentPoolsClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AgentPoolsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AgentPoolsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AgentPoolsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AgentPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *AgentPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AgentPoolsClient) listHandleResponse(resp *http.Response) (AgentPoolsClientListResponse, error) {
	result := AgentPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolListResult); err != nil {
		return AgentPoolsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an agent pool with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - agentPoolName - The name of the agent pool.
//   - updateParameters - The parameters for updating an agent pool.
//   - options - AgentPoolsClientBeginUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginUpdate method.
func (client *AgentPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters, options *AgentPoolsClientBeginUpdateOptions) (*runtime.Poller[AgentPoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, agentPoolName, updateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[AgentPoolsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an agent pool with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01-preview
func (client *AgentPoolsClient) update(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters, options *AgentPoolsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AgentPoolsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, agentPoolName, updateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *AgentPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, agentPoolName string, updateParameters AgentPoolUpdateParameters, options *AgentPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/agentPools/{agentPoolName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
