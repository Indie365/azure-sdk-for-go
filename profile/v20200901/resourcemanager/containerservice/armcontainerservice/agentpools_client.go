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
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
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
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAgentPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AgentPoolsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+".AgentPoolsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AgentPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - parameters - Parameters supplied to the Create or Update an agent pool operation.
//   - options - AgentPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the AgentPoolsClient.BeginCreateOrUpdate
//     method.
func (client *AgentPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AgentPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, agentPoolName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AgentPoolsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
func (client *AgentPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AgentPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters AgentPool, options *AgentPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}"
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
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientBeginDeleteOptions contains the optional parameters for the AgentPoolsClient.BeginDelete method.
func (client *AgentPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*runtime.Poller[AgentPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, agentPoolName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AgentPoolsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the agent pool in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
func (client *AgentPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AgentPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}"
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
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the agent pool by managed cluster and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientGetOptions contains the optional parameters for the AgentPoolsClient.Get method.
func (client *AgentPoolsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetOptions) (AgentPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AgentPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}"
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
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *AgentPoolsClient) getHandleResponse(resp *http.Response) (AgentPoolsClientGetResponse, error) {
	result := AgentPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPool); err != nil {
		return AgentPoolsClientGetResponse{}, err
	}
	return result, nil
}

// GetAvailableAgentPoolVersions - Gets a list of supported versions for the specified agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - options - AgentPoolsClientGetAvailableAgentPoolVersionsOptions contains the optional parameters for the AgentPoolsClient.GetAvailableAgentPoolVersions
//     method.
func (client *AgentPoolsClient) GetAvailableAgentPoolVersions(ctx context.Context, resourceGroupName string, resourceName string, options *AgentPoolsClientGetAvailableAgentPoolVersionsOptions) (AgentPoolsClientGetAvailableAgentPoolVersionsResponse, error) {
	req, err := client.getAvailableAgentPoolVersionsCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAvailableAgentPoolVersionsHandleResponse(resp)
}

// getAvailableAgentPoolVersionsCreateRequest creates the GetAvailableAgentPoolVersions request.
func (client *AgentPoolsClient) getAvailableAgentPoolVersionsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AgentPoolsClientGetAvailableAgentPoolVersionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/availableAgentPoolVersions"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAvailableAgentPoolVersionsHandleResponse handles the GetAvailableAgentPoolVersions response.
func (client *AgentPoolsClient) getAvailableAgentPoolVersionsHandleResponse(resp *http.Response) (AgentPoolsClientGetAvailableAgentPoolVersionsResponse, error) {
	result := AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolAvailableVersions); err != nil {
		return AgentPoolsClientGetAvailableAgentPoolVersionsResponse{}, err
	}
	return result, nil
}

// GetUpgradeProfile - Gets the details of the upgrade profile for an agent pool with a specified resource group and managed
// cluster name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientGetUpgradeProfileOptions contains the optional parameters for the AgentPoolsClient.GetUpgradeProfile
//     method.
func (client *AgentPoolsClient) GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetUpgradeProfileOptions) (AgentPoolsClientGetUpgradeProfileResponse, error) {
	req, err := client.getUpgradeProfileCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return AgentPoolsClientGetUpgradeProfileResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AgentPoolsClientGetUpgradeProfileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AgentPoolsClientGetUpgradeProfileResponse{}, runtime.NewResponseError(resp)
	}
	return client.getUpgradeProfileHandleResponse(resp)
}

// getUpgradeProfileCreateRequest creates the GetUpgradeProfile request.
func (client *AgentPoolsClient) getUpgradeProfileCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientGetUpgradeProfileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeProfiles/default"
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
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getUpgradeProfileHandleResponse handles the GetUpgradeProfile response.
func (client *AgentPoolsClient) getUpgradeProfileHandleResponse(resp *http.Response) (AgentPoolsClientGetUpgradeProfileResponse, error) {
	result := AgentPoolsClientGetUpgradeProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolUpgradeProfile); err != nil {
		return AgentPoolsClientGetUpgradeProfileResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of agent pools in the specified managed cluster. The operation returns properties of each agent
// pool.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - options - AgentPoolsClientListOptions contains the optional parameters for the AgentPoolsClient.NewListPager method.
func (client *AgentPoolsClient) NewListPager(resourceGroupName string, resourceName string, options *AgentPoolsClientListOptions) *runtime.Pager[AgentPoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AgentPoolsClientListResponse]{
		More: func(page AgentPoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AgentPoolsClientListResponse) (AgentPoolsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
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
	})
}

// listCreateRequest creates the List request.
func (client *AgentPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AgentPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *AgentPoolsClient) listHandleResponse(resp *http.Response) (AgentPoolsClientListResponse, error) {
	result := AgentPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgentPoolListResult); err != nil {
		return AgentPoolsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpgradeNodeImageVersion - Upgrade node image version of an agent pool to the latest.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
//   - resourceGroupName - The name of the resource group.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - options - AgentPoolsClientBeginUpgradeNodeImageVersionOptions contains the optional parameters for the AgentPoolsClient.BeginUpgradeNodeImageVersion
//     method.
func (client *AgentPoolsClient) BeginUpgradeNodeImageVersion(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginUpgradeNodeImageVersionOptions) (*runtime.Poller[AgentPoolsClientUpgradeNodeImageVersionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgradeNodeImageVersion(ctx, resourceGroupName, resourceName, agentPoolName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AgentPoolsClientUpgradeNodeImageVersionResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[AgentPoolsClientUpgradeNodeImageVersionResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpgradeNodeImageVersion - Upgrade node image version of an agent pool to the latest.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01
func (client *AgentPoolsClient) upgradeNodeImageVersion(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginUpgradeNodeImageVersionOptions) (*http.Response, error) {
	req, err := client.upgradeNodeImageVersionCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// upgradeNodeImageVersionCreateRequest creates the UpgradeNodeImageVersion request.
func (client *AgentPoolsClient) upgradeNodeImageVersionCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *AgentPoolsClientBeginUpgradeNodeImageVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/upgradeNodeImageVersion"
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
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
