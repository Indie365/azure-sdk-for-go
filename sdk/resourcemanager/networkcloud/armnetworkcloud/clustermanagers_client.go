//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

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

// ClusterManagersClient contains the methods for the ClusterManagers group.
// Don't use this type directly, use NewClusterManagersClient() instead.
type ClusterManagersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClusterManagersClient creates a new instance of ClusterManagersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClusterManagersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClusterManagersClient, error) {
	cl, err := arm.NewClient(moduleName+".ClusterManagersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClusterManagersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new cluster manager or update properties of the cluster manager if it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterManagerName - The name of the cluster manager.
//   - clusterManagerParameters - The request body.
//   - options - ClusterManagersClientBeginCreateOrUpdateOptions contains the optional parameters for the ClusterManagersClient.BeginCreateOrUpdate
//     method.
func (client *ClusterManagersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterManagerName string, clusterManagerParameters ClusterManager, options *ClusterManagersClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClusterManagersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterManagerName, clusterManagerParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClusterManagersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClusterManagersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new cluster manager or update properties of the cluster manager if it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *ClusterManagersClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterManagerName string, clusterManagerParameters ClusterManager, options *ClusterManagersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterManagerName, clusterManagerParameters, options)
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
func (client *ClusterManagersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterManagerName string, clusterManagerParameters ClusterManager, options *ClusterManagersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusterManagers/{clusterManagerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterManagerName == "" {
		return nil, errors.New("parameter clusterManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterManagerName}", url.PathEscape(clusterManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterManagerParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the provided cluster manager.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterManagerName - The name of the cluster manager.
//   - options - ClusterManagersClientBeginDeleteOptions contains the optional parameters for the ClusterManagersClient.BeginDelete
//     method.
func (client *ClusterManagersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterManagerName string, options *ClusterManagersClientBeginDeleteOptions) (*runtime.Poller[ClusterManagersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterManagerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClusterManagersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClusterManagersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided cluster manager.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *ClusterManagersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterManagerName string, options *ClusterManagersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterManagerName, options)
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
func (client *ClusterManagersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterManagerName string, options *ClusterManagersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusterManagers/{clusterManagerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterManagerName == "" {
		return nil, errors.New("parameter clusterManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterManagerName}", url.PathEscape(clusterManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of the provided cluster manager.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterManagerName - The name of the cluster manager.
//   - options - ClusterManagersClientGetOptions contains the optional parameters for the ClusterManagersClient.Get method.
func (client *ClusterManagersClient) Get(ctx context.Context, resourceGroupName string, clusterManagerName string, options *ClusterManagersClientGetOptions) (ClusterManagersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterManagerName, options)
	if err != nil {
		return ClusterManagersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClusterManagersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClusterManagersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ClusterManagersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterManagerName string, options *ClusterManagersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusterManagers/{clusterManagerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterManagerName == "" {
		return nil, errors.New("parameter clusterManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterManagerName}", url.PathEscape(clusterManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClusterManagersClient) getHandleResponse(resp *http.Response) (ClusterManagersClientGetResponse, error) {
	result := ClusterManagersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterManager); err != nil {
		return ClusterManagersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of cluster managers in the provided resource group.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ClusterManagersClientListByResourceGroupOptions contains the optional parameters for the ClusterManagersClient.NewListByResourceGroupPager
//     method.
func (client *ClusterManagersClient) NewListByResourceGroupPager(resourceGroupName string, options *ClusterManagersClientListByResourceGroupOptions) *runtime.Pager[ClusterManagersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterManagersClientListByResourceGroupResponse]{
		More: func(page ClusterManagersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterManagersClientListByResourceGroupResponse) (ClusterManagersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClusterManagersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClusterManagersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClusterManagersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ClusterManagersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClusterManagersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusterManagers"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ClusterManagersClient) listByResourceGroupHandleResponse(resp *http.Response) (ClusterManagersClientListByResourceGroupResponse, error) {
	result := ClusterManagersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterManagerList); err != nil {
		return ClusterManagersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of cluster managers in the provided subscription.
//
// Generated from API version 2023-05-01-preview
//   - options - ClusterManagersClientListBySubscriptionOptions contains the optional parameters for the ClusterManagersClient.NewListBySubscriptionPager
//     method.
func (client *ClusterManagersClient) NewListBySubscriptionPager(options *ClusterManagersClientListBySubscriptionOptions) *runtime.Pager[ClusterManagersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterManagersClientListBySubscriptionResponse]{
		More: func(page ClusterManagersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterManagersClientListBySubscriptionResponse) (ClusterManagersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClusterManagersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClusterManagersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClusterManagersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ClusterManagersClient) listBySubscriptionCreateRequest(ctx context.Context, options *ClusterManagersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/clusterManagers"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ClusterManagersClient) listBySubscriptionHandleResponse(resp *http.Response) (ClusterManagersClientListBySubscriptionResponse, error) {
	result := ClusterManagersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterManagerList); err != nil {
		return ClusterManagersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patch properties of the provided cluster manager, or update the tags assigned to the cluster manager. Properties
// and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterManagerName - The name of the cluster manager.
//   - clusterManagerUpdateParameters - The request body.
//   - options - ClusterManagersClientUpdateOptions contains the optional parameters for the ClusterManagersClient.Update method.
func (client *ClusterManagersClient) Update(ctx context.Context, resourceGroupName string, clusterManagerName string, clusterManagerUpdateParameters ClusterManagerPatchParameters, options *ClusterManagersClientUpdateOptions) (ClusterManagersClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterManagerName, clusterManagerUpdateParameters, options)
	if err != nil {
		return ClusterManagersClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClusterManagersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClusterManagersClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ClusterManagersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterManagerName string, clusterManagerUpdateParameters ClusterManagerPatchParameters, options *ClusterManagersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusterManagers/{clusterManagerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterManagerName == "" {
		return nil, errors.New("parameter clusterManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterManagerName}", url.PathEscape(clusterManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterManagerUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ClusterManagersClient) updateHandleResponse(resp *http.Response) (ClusterManagersClientUpdateResponse, error) {
	result := ClusterManagersClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterManager); err != nil {
		return ClusterManagersClientUpdateResponse{}, err
	}
	return result, nil
}
