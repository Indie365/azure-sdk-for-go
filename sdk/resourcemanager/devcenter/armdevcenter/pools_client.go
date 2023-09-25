//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PoolsClient contains the methods for the Pools group.
// Don't use this type directly, use NewPoolsClient() instead.
type PoolsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewPoolsClient creates a new instance of PoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PoolsClient, error) {
	cl, err := arm.NewClient(moduleName+".PoolsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PoolsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a machine pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - poolName - Name of the pool.
//   - body - Represents a machine pool
//   - options - PoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the PoolsClient.BeginCreateOrUpdate
//     method.
func (client *PoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, projectName string, poolName string, body Pool, options *PoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, projectName, poolName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PoolsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PoolsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a machine pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *PoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, projectName string, poolName string, body Pool, options *PoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, projectName, poolName, body, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, body Pool, options *PoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
	return req, nil
}

// BeginDelete - Deletes a machine pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - poolName - Name of the pool.
//   - options - PoolsClientBeginDeleteOptions contains the optional parameters for the PoolsClient.BeginDelete method.
func (client *PoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientBeginDeleteOptions) (*runtime.Poller[PoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, projectName, poolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PoolsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PoolsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a machine pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *PoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, projectName, poolName, options)
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
func (client *PoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a machine pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - poolName - Name of the pool.
//   - options - PoolsClientGetOptions contains the optional parameters for the PoolsClient.Get method.
func (client *PoolsClient) Get(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientGetOptions) (PoolsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, projectName, poolName, options)
	if err != nil {
		return PoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PoolsClient) getHandleResponse(resp *http.Response) (PoolsClientGetResponse, error) {
	result := PoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pool); err != nil {
		return PoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProjectPager - Lists pools for a project
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - options - PoolsClientListByProjectOptions contains the optional parameters for the PoolsClient.NewListByProjectPager method.
func (client *PoolsClient) NewListByProjectPager(resourceGroupName string, projectName string, options *PoolsClientListByProjectOptions) (*runtime.Pager[PoolsClientListByProjectResponse]) {
	return runtime.NewPager(runtime.PagingHandler[PoolsClientListByProjectResponse]{
		More: func(page PoolsClientListByProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PoolsClientListByProjectResponse) (PoolsClientListByProjectResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PoolsClientListByProjectResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PoolsClientListByProjectResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PoolsClientListByProjectResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProjectHandleResponse(resp)
		},
	})
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *PoolsClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *PoolsClientListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *PoolsClient) listByProjectHandleResponse(resp *http.Response) (PoolsClientListByProjectResponse, error) {
	result := PoolsClientListByProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PoolListResult); err != nil {
		return PoolsClientListByProjectResponse{}, err
	}
	return result, nil
}

// BeginRunHealthChecks - Triggers a refresh of the pool status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - poolName - Name of the pool.
//   - options - PoolsClientBeginRunHealthChecksOptions contains the optional parameters for the PoolsClient.BeginRunHealthChecks
//     method.
func (client *PoolsClient) BeginRunHealthChecks(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientBeginRunHealthChecksOptions) (*runtime.Poller[PoolsClientRunHealthChecksResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.runHealthChecks(ctx, resourceGroupName, projectName, poolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PoolsClientRunHealthChecksResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PoolsClientRunHealthChecksResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RunHealthChecks - Triggers a refresh of the pool status.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *PoolsClient) runHealthChecks(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientBeginRunHealthChecksOptions) (*http.Response, error) {
	var err error
	req, err := client.runHealthChecksCreateRequest(ctx, resourceGroupName, projectName, poolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// runHealthChecksCreateRequest creates the RunHealthChecks request.
func (client *PoolsClient) runHealthChecksCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, options *PoolsClientBeginRunHealthChecksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}/runHealthChecks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Partially updates a machine pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - projectName - The name of the project.
//   - poolName - Name of the pool.
//   - body - Represents a machine pool
//   - options - PoolsClientBeginUpdateOptions contains the optional parameters for the PoolsClient.BeginUpdate method.
func (client *PoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, projectName string, poolName string, body PoolUpdate, options *PoolsClientBeginUpdateOptions) (*runtime.Poller[PoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, projectName, poolName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PoolsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PoolsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Partially updates a machine pool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *PoolsClient) update(ctx context.Context, resourceGroupName string, projectName string, poolName string, body PoolUpdate, options *PoolsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, projectName, poolName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *PoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, projectName string, poolName string, body PoolUpdate, options *PoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/pools/{poolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
	return req, nil
}

