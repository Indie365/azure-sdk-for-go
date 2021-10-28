//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ManagedDatabaseColumnsClient contains the methods for the ManagedDatabaseColumns group.
// Don't use this type directly, use NewManagedDatabaseColumnsClient() instead.
type ManagedDatabaseColumnsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedDatabaseColumnsClient creates a new instance of ManagedDatabaseColumnsClient with the specified values.
func NewManagedDatabaseColumnsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedDatabaseColumnsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ManagedDatabaseColumnsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get managed database column
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseColumnsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseColumnsGetOptions) (ManagedDatabaseColumnsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return ManagedDatabaseColumnsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseColumnsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseColumnsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseColumnsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseColumnsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedDatabaseColumnsClient) getHandleResponse(resp *http.Response) (ManagedDatabaseColumnsGetResponse, error) {
	result := ManagedDatabaseColumnsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumn); err != nil {
		return ManagedDatabaseColumnsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedDatabaseColumnsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabase - List managed database columns
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseColumnsClient) ListByDatabase(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseColumnsListByDatabaseOptions) *ManagedDatabaseColumnsListByDatabasePager {
	return &ManagedDatabaseColumnsListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp ManagedDatabaseColumnsListByDatabaseResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatabaseColumnListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedDatabaseColumnsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseColumnsListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/columns"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Schema != nil {
		for _, qv := range options.Schema {
			reqQP.Add("schema", qv)
		}
	}
	if options != nil && options.Table != nil {
		for _, qv := range options.Table {
			reqQP.Add("table", qv)
		}
	}
	if options != nil && options.Column != nil {
		for _, qv := range options.Column {
			reqQP.Add("column", qv)
		}
	}
	if options != nil && options.OrderBy != nil {
		for _, qv := range options.OrderBy {
			reqQP.Add("orderBy", qv)
		}
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedDatabaseColumnsClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseColumnsListByDatabaseResponse, error) {
	result := ManagedDatabaseColumnsListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumnListResult); err != nil {
		return ManagedDatabaseColumnsListByDatabaseResponse{}, err
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *ManagedDatabaseColumnsClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByTable - List managed database columns
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseColumnsClient) ListByTable(resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, options *ManagedDatabaseColumnsListByTableOptions) *ManagedDatabaseColumnsListByTablePager {
	return &ManagedDatabaseColumnsListByTablePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByTableCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, options)
		},
		advancer: func(ctx context.Context, resp ManagedDatabaseColumnsListByTableResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatabaseColumnListResult.NextLink)
		},
	}
}

// listByTableCreateRequest creates the ListByTable request.
func (client *ManagedDatabaseColumnsClient) listByTableCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, options *ManagedDatabaseColumnsListByTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByTableHandleResponse handles the ListByTable response.
func (client *ManagedDatabaseColumnsClient) listByTableHandleResponse(resp *http.Response) (ManagedDatabaseColumnsListByTableResponse, error) {
	result := ManagedDatabaseColumnsListByTableResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumnListResult); err != nil {
		return ManagedDatabaseColumnsListByTableResponse{}, err
	}
	return result, nil
}

// listByTableHandleError handles the ListByTable error response.
func (client *ManagedDatabaseColumnsClient) listByTableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
