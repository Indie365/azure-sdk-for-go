//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// IncidentTasksClient contains the methods for the IncidentTasks group.
// Don't use this type directly, use NewIncidentTasksClient() instead.
type IncidentTasksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIncidentTasksClient creates a new instance of IncidentTasksClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIncidentTasksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IncidentTasksClient, error) {
	cl, err := arm.NewClient(moduleName+".IncidentTasksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IncidentTasksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the incident task.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - incidentTaskID - Incident task ID
//   - incidentTask - The incident task
//   - options - IncidentTasksClientCreateOrUpdateOptions contains the optional parameters for the IncidentTasksClient.CreateOrUpdate
//     method.
func (client *IncidentTasksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentTaskID string, incidentTask IncidentTask, options *IncidentTasksClientCreateOrUpdateOptions) (IncidentTasksClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incidentTaskID, incidentTask, options)
	if err != nil {
		return IncidentTasksClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentTasksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IncidentTasksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IncidentTasksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentTaskID string, incidentTask IncidentTask, options *IncidentTasksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/tasks/{incidentTaskId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if incidentTaskID == "" {
		return nil, errors.New("parameter incidentTaskID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentTaskId}", url.PathEscape(incidentTaskID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, incidentTask); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IncidentTasksClient) createOrUpdateHandleResponse(resp *http.Response) (IncidentTasksClientCreateOrUpdateResponse, error) {
	result := IncidentTasksClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentTask); err != nil {
		return IncidentTasksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the incident task.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - incidentTaskID - Incident task ID
//   - options - IncidentTasksClientDeleteOptions contains the optional parameters for the IncidentTasksClient.Delete method.
func (client *IncidentTasksClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentTaskID string, options *IncidentTasksClientDeleteOptions) (IncidentTasksClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incidentTaskID, options)
	if err != nil {
		return IncidentTasksClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentTasksClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IncidentTasksClientDeleteResponse{}, err
	}
	return IncidentTasksClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IncidentTasksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentTaskID string, options *IncidentTasksClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/tasks/{incidentTaskId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if incidentTaskID == "" {
		return nil, errors.New("parameter incidentTaskID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentTaskId}", url.PathEscape(incidentTaskID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an incident task.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - incidentTaskID - Incident task ID
//   - options - IncidentTasksClientGetOptions contains the optional parameters for the IncidentTasksClient.Get method.
func (client *IncidentTasksClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentTaskID string, options *IncidentTasksClientGetOptions) (IncidentTasksClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, incidentTaskID, options)
	if err != nil {
		return IncidentTasksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IncidentTasksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IncidentTasksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IncidentTasksClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentTaskID string, options *IncidentTasksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/tasks/{incidentTaskId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	if incidentTaskID == "" {
		return nil, errors.New("parameter incidentTaskID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentTaskId}", url.PathEscape(incidentTaskID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IncidentTasksClient) getHandleResponse(resp *http.Response) (IncidentTasksClientGetResponse, error) {
	result := IncidentTasksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentTask); err != nil {
		return IncidentTasksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all incident tasks.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - incidentID - Incident ID
//   - options - IncidentTasksClientListOptions contains the optional parameters for the IncidentTasksClient.NewListPager method.
func (client *IncidentTasksClient) NewListPager(resourceGroupName string, workspaceName string, incidentID string, options *IncidentTasksClientListOptions) *runtime.Pager[IncidentTasksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IncidentTasksClientListResponse]{
		More: func(page IncidentTasksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IncidentTasksClientListResponse) (IncidentTasksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, incidentID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return IncidentTasksClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return IncidentTasksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IncidentTasksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IncidentTasksClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *IncidentTasksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/incidents/{incidentId}/tasks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if incidentID == "" {
		return nil, errors.New("parameter incidentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{incidentId}", url.PathEscape(incidentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IncidentTasksClient) listHandleResponse(resp *http.Response) (IncidentTasksClientListResponse, error) {
	result := IncidentTasksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IncidentTaskList); err != nil {
		return IncidentTasksClientListResponse{}, err
	}
	return result, nil
}
