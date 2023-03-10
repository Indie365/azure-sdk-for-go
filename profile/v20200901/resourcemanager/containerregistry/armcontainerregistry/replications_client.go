//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry

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

// ReplicationsClient contains the methods for the Replications group.
// Don't use this type directly, use NewReplicationsClient() instead.
type ReplicationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationsClient creates a new instance of ReplicationsClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+".ReplicationsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a replication for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - replicationName - The name of the replication.
//   - replication - The parameters for creating a replication.
//   - options - ReplicationsClientBeginCreateOptions contains the optional parameters for the ReplicationsClient.BeginCreate
//     method.
func (client *ReplicationsClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsClientBeginCreateOptions) (*runtime.Poller[ReplicationsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, replicationName, replication, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationsClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a replication for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *ReplicationsClient) create(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, replicationName, replication, options)
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

// createCreateRequest creates the Create request.
func (client *ReplicationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replication Replication, options *ReplicationsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
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
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, replication)
}

// BeginDelete - Deletes a replication from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - replicationName - The name of the replication.
//   - options - ReplicationsClientBeginDeleteOptions contains the optional parameters for the ReplicationsClient.BeginDelete
//     method.
func (client *ReplicationsClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsClientBeginDeleteOptions) (*runtime.Poller[ReplicationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, replicationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a replication from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *ReplicationsClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
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
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the properties of the specified replication.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - replicationName - The name of the replication.
//   - options - ReplicationsClientGetOptions contains the optional parameters for the ReplicationsClient.Get method.
func (client *ReplicationsClient) Get(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsClientGetOptions) (ReplicationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, replicationName, options)
	if err != nil {
		return ReplicationsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, options *ReplicationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
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
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationsClient) getHandleResponse(resp *http.Response) (ReplicationsClientGetResponse, error) {
	result := ReplicationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Replication); err != nil {
		return ReplicationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the replications for the specified container registry.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - options - ReplicationsClientListOptions contains the optional parameters for the ReplicationsClient.NewListPager method.
func (client *ReplicationsClient) NewListPager(resourceGroupName string, registryName string, options *ReplicationsClientListOptions) *runtime.Pager[ReplicationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationsClientListResponse]{
		More: func(page ReplicationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationsClientListResponse) (ReplicationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ReplicationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationsClient) listHandleResponse(resp *http.Response) (ReplicationsClientListResponse, error) {
	result := ReplicationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationListResult); err != nil {
		return ReplicationsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a replication for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - registryName - The name of the container registry.
//   - replicationName - The name of the replication.
//   - replicationUpdateParameters - The parameters for updating a replication.
//   - options - ReplicationsClientBeginUpdateOptions contains the optional parameters for the ReplicationsClient.BeginUpdate
//     method.
func (client *ReplicationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsClientBeginUpdateOptions) (*runtime.Poller[ReplicationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, registryName, replicationName, replicationUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a replication for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *ReplicationsClient) update(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, replicationName, replicationUpdateParameters, options)
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

// updateCreateRequest creates the Update request.
func (client *ReplicationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, replicationName string, replicationUpdateParameters ReplicationUpdateParameters, options *ReplicationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/replications/{replicationName}"
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
	if replicationName == "" {
		return nil, errors.New("parameter replicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationName}", url.PathEscape(replicationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, replicationUpdateParameters)
}
