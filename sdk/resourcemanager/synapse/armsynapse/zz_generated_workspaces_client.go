//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// WorkspacesClient contains the methods for the Workspaces group.
// Don't use this type directly, use NewWorkspacesClient() instead.
type WorkspacesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspacesClient creates a new instance of WorkspacesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspacesClient, error) {
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
	client := &WorkspacesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// workspaceInfo - Workspace create or update request properties
// options - WorkspacesClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspacesClient.BeginCreateOrUpdate
// method.
func (client *WorkspacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceInfo Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspacesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, workspaceInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[WorkspacesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[WorkspacesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *WorkspacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceInfo Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, workspaceInfo, options)
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
func (client *WorkspacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceInfo Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, workspaceInfo)
}

// BeginDelete - Deletes a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspacesClientBeginDeleteOptions contains the optional parameters for the WorkspacesClient.BeginDelete method.
func (client *WorkspacesClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*runtime.Poller[WorkspacesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[WorkspacesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[WorkspacesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *WorkspacesClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, options)
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
func (client *WorkspacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
func (client *WorkspacesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (WorkspacesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspacesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspacesClient) getHandleResponse(resp *http.Response) (WorkspacesClientGetResponse, error) {
	result := WorkspacesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workspace); err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of workspaces in a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// options - WorkspacesClientListOptions contains the optional parameters for the WorkspacesClient.List method.
func (client *WorkspacesClient) NewListPager(options *WorkspacesClientListOptions) *runtime.Pager[WorkspacesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspacesClientListResponse]{
		More: func(page WorkspacesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspacesClientListResponse) (WorkspacesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspacesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspacesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspacesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspacesClient) listCreateRequest(ctx context.Context, options *WorkspacesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Synapse/workspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspacesClient) listHandleResponse(resp *http.Response) (WorkspacesClientListResponse, error) {
	result := WorkspacesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceInfoListResult); err != nil {
		return WorkspacesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Returns a list of workspaces in a resource group
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - WorkspacesClientListByResourceGroupOptions contains the optional parameters for the WorkspacesClient.ListByResourceGroup
// method.
func (client *WorkspacesClient) NewListByResourceGroupPager(resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) *runtime.Pager[WorkspacesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspacesClientListByResourceGroupResponse]{
		More: func(page WorkspacesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspacesClientListByResourceGroupResponse) (WorkspacesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspacesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return WorkspacesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspacesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WorkspacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WorkspacesClient) listByResourceGroupHandleResponse(resp *http.Response) (WorkspacesClientListByResourceGroupResponse, error) {
	result := WorkspacesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceInfoListResult); err != nil {
		return WorkspacesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// workspacePatchInfo - Workspace patch request properties
// options - WorkspacesClientBeginUpdateOptions contains the optional parameters for the WorkspacesClient.BeginUpdate method.
func (client *WorkspacesClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchInfo WorkspacePatchInfo, options *WorkspacesClientBeginUpdateOptions) (*runtime.Poller[WorkspacesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, workspacePatchInfo, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[WorkspacesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[WorkspacesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates a workspace
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *WorkspacesClient) update(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchInfo WorkspacePatchInfo, options *WorkspacesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, workspacePatchInfo, options)
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
func (client *WorkspacesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchInfo WorkspacePatchInfo, options *WorkspacesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, workspacePatchInfo)
}