//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

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

// SQLManagedInstancesClient contains the methods for the SQLManagedInstances group.
// Don't use this type directly, use NewSQLManagedInstancesClient() instead.
type SQLManagedInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLManagedInstancesClient creates a new instance of SQLManagedInstancesClient with the specified values.
//   - subscriptionID - The ID of the Azure subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLManagedInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLManagedInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLManagedInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates or replaces a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - sqlManagedInstance - The SQL Managed Instance to be created or updated.
//   - options - SQLManagedInstancesClientBeginCreateOptions contains the optional parameters for the SQLManagedInstancesClient.BeginCreate
//     method.
func (client *SQLManagedInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, sqlManagedInstance SQLManagedInstance, options *SQLManagedInstancesClientBeginCreateOptions) (*runtime.Poller[SQLManagedInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sqlManagedInstanceName, sqlManagedInstance, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SQLManagedInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SQLManagedInstancesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates or replaces a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
func (client *SQLManagedInstancesClient) create(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, sqlManagedInstance SQLManagedInstance, options *SQLManagedInstancesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "SQLManagedInstancesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, sqlManagedInstance, options)
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
func (client *SQLManagedInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, sqlManagedInstance SQLManagedInstance, options *SQLManagedInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sqlManagedInstance); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - options - SQLManagedInstancesClientBeginDeleteOptions contains the optional parameters for the SQLManagedInstancesClient.BeginDelete
//     method.
func (client *SQLManagedInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientBeginDeleteOptions) (*runtime.Poller[SQLManagedInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlManagedInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SQLManagedInstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SQLManagedInstancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
func (client *SQLManagedInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SQLManagedInstancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, options)
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
func (client *SQLManagedInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - options - SQLManagedInstancesClientGetOptions contains the optional parameters for the SQLManagedInstancesClient.Get method.
func (client *SQLManagedInstancesClient) Get(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientGetOptions) (SQLManagedInstancesClientGetResponse, error) {
	var err error
	const operationName = "SQLManagedInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, options)
	if err != nil {
		return SQLManagedInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLManagedInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLManagedInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLManagedInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, options *SQLManagedInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLManagedInstancesClient) getHandleResponse(resp *http.Response) (SQLManagedInstancesClientGetResponse, error) {
	result := SQLManagedInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstance); err != nil {
		return SQLManagedInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List sqlManagedInstance resources in the subscription
//
// Generated from API version 2022-03-01-preview
//   - options - SQLManagedInstancesClientListOptions contains the optional parameters for the SQLManagedInstancesClient.NewListPager
//     method.
func (client *SQLManagedInstancesClient) NewListPager(options *SQLManagedInstancesClientListOptions) *runtime.Pager[SQLManagedInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLManagedInstancesClientListResponse]{
		More: func(page SQLManagedInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLManagedInstancesClientListResponse) (SQLManagedInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLManagedInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SQLManagedInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SQLManagedInstancesClient) listCreateRequest(ctx context.Context, options *SQLManagedInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/sqlManagedInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLManagedInstancesClient) listHandleResponse(resp *http.Response) (SQLManagedInstancesClientListResponse, error) {
	result := SQLManagedInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstanceListResult); err != nil {
		return SQLManagedInstancesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all sqlManagedInstances in a resource group.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - options - SQLManagedInstancesClientListByResourceGroupOptions contains the optional parameters for the SQLManagedInstancesClient.NewListByResourceGroupPager
//     method.
func (client *SQLManagedInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *SQLManagedInstancesClientListByResourceGroupOptions) *runtime.Pager[SQLManagedInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLManagedInstancesClientListByResourceGroupResponse]{
		More: func(page SQLManagedInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLManagedInstancesClientListByResourceGroupResponse) (SQLManagedInstancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLManagedInstancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SQLManagedInstancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLManagedInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SQLManagedInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLManagedInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLManagedInstancesClientListByResourceGroupResponse, error) {
	result := SQLManagedInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstanceListResult); err != nil {
		return SQLManagedInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a SQL Managed Instance resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlManagedInstanceName - Name of SQL Managed Instance
//   - parameters - The SQL Managed Instance.
//   - options - SQLManagedInstancesClientUpdateOptions contains the optional parameters for the SQLManagedInstancesClient.Update
//     method.
func (client *SQLManagedInstancesClient) Update(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, parameters SQLManagedInstanceUpdate, options *SQLManagedInstancesClientUpdateOptions) (SQLManagedInstancesClientUpdateResponse, error) {
	var err error
	const operationName = "SQLManagedInstancesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlManagedInstanceName, parameters, options)
	if err != nil {
		return SQLManagedInstancesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLManagedInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLManagedInstancesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SQLManagedInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, parameters SQLManagedInstanceUpdate, options *SQLManagedInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlManagedInstances/{sqlManagedInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlManagedInstanceName == "" {
		return nil, errors.New("parameter sqlManagedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlManagedInstanceName}", url.PathEscape(sqlManagedInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SQLManagedInstancesClient) updateHandleResponse(resp *http.Response) (SQLManagedInstancesClientUpdateResponse, error) {
	result := SQLManagedInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLManagedInstance); err != nil {
		return SQLManagedInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
