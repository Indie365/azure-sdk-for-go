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

// TaskRunsClient contains the methods for the TaskRuns group.
// Don't use this type directly, use NewTaskRunsClient() instead.
type TaskRunsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTaskRunsClient creates a new instance of TaskRunsClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTaskRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TaskRunsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &TaskRunsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Creates a task run for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// taskRunName - The name of the task run.
// taskRun - The parameters of a run that needs to scheduled.
// options - TaskRunsClientBeginCreateOptions contains the optional parameters for the TaskRunsClient.BeginCreate method.
func (client *TaskRunsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, taskRun TaskRun, options *TaskRunsClientBeginCreateOptions) (TaskRunsClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, taskRunName, taskRun, options)
	if err != nil {
		return TaskRunsClientCreatePollerResponse{}, err
	}
	result := TaskRunsClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TaskRunsClient.Create", "", resp, client.pl)
	if err != nil {
		return TaskRunsClientCreatePollerResponse{}, err
	}
	result.Poller = &TaskRunsClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a task run for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TaskRunsClient) create(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, taskRun TaskRun, options *TaskRunsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, taskRunName, taskRun, options)
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
func (client *TaskRunsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, taskRun TaskRun, options *TaskRunsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/taskRuns/{taskRunName}"
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
	if taskRunName == "" {
		return nil, errors.New("parameter taskRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskRunName}", url.PathEscape(taskRunName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, taskRun)
}

// BeginDelete - Deletes a specified task run resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// taskRunName - The name of the task run.
// options - TaskRunsClientBeginDeleteOptions contains the optional parameters for the TaskRunsClient.BeginDelete method.
func (client *TaskRunsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *TaskRunsClientBeginDeleteOptions) (TaskRunsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, taskRunName, options)
	if err != nil {
		return TaskRunsClientDeletePollerResponse{}, err
	}
	result := TaskRunsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TaskRunsClient.Delete", "", resp, client.pl)
	if err != nil {
		return TaskRunsClientDeletePollerResponse{}, err
	}
	result.Poller = &TaskRunsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a specified task run resource.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TaskRunsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *TaskRunsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, taskRunName, options)
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
func (client *TaskRunsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *TaskRunsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/taskRuns/{taskRunName}"
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
	if taskRunName == "" {
		return nil, errors.New("parameter taskRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskRunName}", url.PathEscape(taskRunName))
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

// Get - Gets the detailed information for a given task run.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// taskRunName - The name of the task run.
// options - TaskRunsClientGetOptions contains the optional parameters for the TaskRunsClient.Get method.
func (client *TaskRunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *TaskRunsClientGetOptions) (TaskRunsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, taskRunName, options)
	if err != nil {
		return TaskRunsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TaskRunsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TaskRunsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TaskRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *TaskRunsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/taskRuns/{taskRunName}"
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
	if taskRunName == "" {
		return nil, errors.New("parameter taskRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskRunName}", url.PathEscape(taskRunName))
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
func (client *TaskRunsClient) getHandleResponse(resp *http.Response) (TaskRunsClientGetResponse, error) {
	result := TaskRunsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskRun); err != nil {
		return TaskRunsClientGetResponse{}, err
	}
	return result, nil
}

// GetDetails - Gets the detailed information for a given task run that includes all secrets.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// taskRunName - The name of the task run.
// options - TaskRunsClientGetDetailsOptions contains the optional parameters for the TaskRunsClient.GetDetails method.
func (client *TaskRunsClient) GetDetails(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *TaskRunsClientGetDetailsOptions) (TaskRunsClientGetDetailsResponse, error) {
	req, err := client.getDetailsCreateRequest(ctx, resourceGroupName, registryName, taskRunName, options)
	if err != nil {
		return TaskRunsClientGetDetailsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TaskRunsClientGetDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TaskRunsClientGetDetailsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDetailsHandleResponse(resp)
}

// getDetailsCreateRequest creates the GetDetails request.
func (client *TaskRunsClient) getDetailsCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, options *TaskRunsClientGetDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/taskRuns/{taskRunName}/listDetails"
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
	if taskRunName == "" {
		return nil, errors.New("parameter taskRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskRunName}", url.PathEscape(taskRunName))
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

// getDetailsHandleResponse handles the GetDetails response.
func (client *TaskRunsClient) getDetailsHandleResponse(resp *http.Response) (TaskRunsClientGetDetailsResponse, error) {
	result := TaskRunsClientGetDetailsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskRun); err != nil {
		return TaskRunsClientGetDetailsResponse{}, err
	}
	return result, nil
}

// List - Lists all the task runs for a specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// options - TaskRunsClientListOptions contains the optional parameters for the TaskRunsClient.List method.
func (client *TaskRunsClient) List(resourceGroupName string, registryName string, options *TaskRunsClientListOptions) *TaskRunsClientListPager {
	return &TaskRunsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		advancer: func(ctx context.Context, resp TaskRunsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TaskRunListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *TaskRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *TaskRunsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/taskRuns"
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
func (client *TaskRunsClient) listHandleResponse(resp *http.Response) (TaskRunsClientListResponse, error) {
	result := TaskRunsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskRunListResult); err != nil {
		return TaskRunsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a task run with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// taskRunName - The name of the task run.
// updateParameters - The parameters for updating a task run.
// options - TaskRunsClientBeginUpdateOptions contains the optional parameters for the TaskRunsClient.BeginUpdate method.
func (client *TaskRunsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, updateParameters TaskRunUpdateParameters, options *TaskRunsClientBeginUpdateOptions) (TaskRunsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, taskRunName, updateParameters, options)
	if err != nil {
		return TaskRunsClientUpdatePollerResponse{}, err
	}
	result := TaskRunsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("TaskRunsClient.Update", "", resp, client.pl)
	if err != nil {
		return TaskRunsClientUpdatePollerResponse{}, err
	}
	result.Poller = &TaskRunsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a task run with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *TaskRunsClient) update(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, updateParameters TaskRunUpdateParameters, options *TaskRunsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, taskRunName, updateParameters, options)
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
func (client *TaskRunsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, taskRunName string, updateParameters TaskRunUpdateParameters, options *TaskRunsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/taskRuns/{taskRunName}"
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
	if taskRunName == "" {
		return nil, errors.New("parameter taskRunName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskRunName}", url.PathEscape(taskRunName))
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
