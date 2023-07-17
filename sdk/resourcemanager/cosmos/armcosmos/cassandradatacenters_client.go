//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos

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

// CassandraDataCentersClient contains the methods for the CassandraDataCenters group.
// Don't use this type directly, use NewCassandraDataCentersClient() instead.
type CassandraDataCentersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCassandraDataCentersClient creates a new instance of CassandraDataCentersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCassandraDataCentersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CassandraDataCentersClient, error) {
	cl, err := arm.NewClient(moduleName+".CassandraDataCentersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CassandraDataCentersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateUpdate - Create or update a managed Cassandra data center. When updating, overwrite all properties. To update
// only some properties, use PATCH.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - dataCenterName - Data center name in a managed Cassandra cluster.
//   - body - Parameters specifying the managed Cassandra data center.
//   - options - CassandraDataCentersClientBeginCreateUpdateOptions contains the optional parameters for the CassandraDataCentersClient.BeginCreateUpdate
//     method.
func (client *CassandraDataCentersClient) BeginCreateUpdate(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource, options *CassandraDataCentersClientBeginCreateUpdateOptions) (*runtime.Poller[CassandraDataCentersClientCreateUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createUpdate(ctx, resourceGroupName, clusterName, dataCenterName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CassandraDataCentersClientCreateUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[CassandraDataCentersClientCreateUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateUpdate - Create or update a managed Cassandra data center. When updating, overwrite all properties. To update only
// some properties, use PATCH.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *CassandraDataCentersClient) createUpdate(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource, options *CassandraDataCentersClientBeginCreateUpdateOptions) (*http.Response, error) {
	req, err := client.createUpdateCreateRequest(ctx, resourceGroupName, clusterName, dataCenterName, body, options)
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

// createUpdateCreateRequest creates the CreateUpdate request.
func (client *CassandraDataCentersClient) createUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource, options *CassandraDataCentersClientBeginCreateUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}"
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
	if dataCenterName == "" {
		return nil, errors.New("parameter dataCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCenterName}", url.PathEscape(dataCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete a managed Cassandra data center.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - dataCenterName - Data center name in a managed Cassandra cluster.
//   - options - CassandraDataCentersClientBeginDeleteOptions contains the optional parameters for the CassandraDataCentersClient.BeginDelete
//     method.
func (client *CassandraDataCentersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, options *CassandraDataCentersClientBeginDeleteOptions) (*runtime.Poller[CassandraDataCentersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, dataCenterName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CassandraDataCentersClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[CassandraDataCentersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a managed Cassandra data center.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *CassandraDataCentersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, options *CassandraDataCentersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, dataCenterName, options)
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
func (client *CassandraDataCentersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, options *CassandraDataCentersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}"
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
	if dataCenterName == "" {
		return nil, errors.New("parameter dataCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCenterName}", url.PathEscape(dataCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the properties of a managed Cassandra data center.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - dataCenterName - Data center name in a managed Cassandra cluster.
//   - options - CassandraDataCentersClientGetOptions contains the optional parameters for the CassandraDataCentersClient.Get
//     method.
func (client *CassandraDataCentersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, options *CassandraDataCentersClientGetOptions) (CassandraDataCentersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, dataCenterName, options)
	if err != nil {
		return CassandraDataCentersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CassandraDataCentersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CassandraDataCentersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CassandraDataCentersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, options *CassandraDataCentersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}"
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
	if dataCenterName == "" {
		return nil, errors.New("parameter dataCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCenterName}", url.PathEscape(dataCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CassandraDataCentersClient) getHandleResponse(resp *http.Response) (CassandraDataCentersClientGetResponse, error) {
	result := CassandraDataCentersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCenterResource); err != nil {
		return CassandraDataCentersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all data centers in a particular managed Cassandra cluster.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - options - CassandraDataCentersClientListOptions contains the optional parameters for the CassandraDataCentersClient.NewListPager
//     method.
func (client *CassandraDataCentersClient) NewListPager(resourceGroupName string, clusterName string, options *CassandraDataCentersClientListOptions) *runtime.Pager[CassandraDataCentersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CassandraDataCentersClientListResponse]{
		More: func(page CassandraDataCentersClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CassandraDataCentersClientListResponse) (CassandraDataCentersClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, clusterName, options)
			if err != nil {
				return CassandraDataCentersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CassandraDataCentersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CassandraDataCentersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CassandraDataCentersClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *CassandraDataCentersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CassandraDataCentersClient) listHandleResponse(resp *http.Response) (CassandraDataCentersClientListResponse, error) {
	result := CassandraDataCentersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListDataCenters); err != nil {
		return CassandraDataCentersClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update some of the properties of a managed Cassandra data center.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Managed Cassandra cluster name.
//   - dataCenterName - Data center name in a managed Cassandra cluster.
//   - body - Parameters to provide for specifying the managed Cassandra data center.
//   - options - CassandraDataCentersClientBeginUpdateOptions contains the optional parameters for the CassandraDataCentersClient.BeginUpdate
//     method.
func (client *CassandraDataCentersClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource, options *CassandraDataCentersClientBeginUpdateOptions) (*runtime.Poller[CassandraDataCentersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, dataCenterName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CassandraDataCentersClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[CassandraDataCentersClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update some of the properties of a managed Cassandra data center.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
func (client *CassandraDataCentersClient) update(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource, options *CassandraDataCentersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, dataCenterName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *CassandraDataCentersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, dataCenterName string, body DataCenterResource, options *CassandraDataCentersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/cassandraClusters/{clusterName}/dataCenters/{dataCenterName}"
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
	if dataCenterName == "" {
		return nil, errors.New("parameter dataCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCenterName}", url.PathEscape(dataCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
