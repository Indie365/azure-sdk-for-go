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

// DatabaseColumnsClient contains the methods for the DatabaseColumns group.
// Don't use this type directly, use NewDatabaseColumnsClient() instead.
type DatabaseColumnsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDatabaseColumnsClient creates a new instance of DatabaseColumnsClient with the specified values.
func NewDatabaseColumnsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DatabaseColumnsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DatabaseColumnsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get database column
// If the operation fails it returns a generic error.
func (client *DatabaseColumnsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *DatabaseColumnsGetOptions) (DatabaseColumnsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return DatabaseColumnsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseColumnsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseColumnsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabaseColumnsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *DatabaseColumnsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
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
func (client *DatabaseColumnsClient) getHandleResponse(resp *http.Response) (DatabaseColumnsGetResponse, error) {
	result := DatabaseColumnsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumn); err != nil {
		return DatabaseColumnsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DatabaseColumnsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabase - List database columns
// If the operation fails it returns a generic error.
func (client *DatabaseColumnsClient) ListByDatabase(resourceGroupName string, serverName string, databaseName string, options *DatabaseColumnsListByDatabaseOptions) *DatabaseColumnsListByDatabasePager {
	return &DatabaseColumnsListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp DatabaseColumnsListByDatabaseResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatabaseColumnListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DatabaseColumnsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DatabaseColumnsListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/columns"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
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
func (client *DatabaseColumnsClient) listByDatabaseHandleResponse(resp *http.Response) (DatabaseColumnsListByDatabaseResponse, error) {
	result := DatabaseColumnsListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumnListResult); err != nil {
		return DatabaseColumnsListByDatabaseResponse{}, err
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *DatabaseColumnsClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByTable - List database columns
// If the operation fails it returns a generic error.
func (client *DatabaseColumnsClient) ListByTable(resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, options *DatabaseColumnsListByTableOptions) *DatabaseColumnsListByTablePager {
	return &DatabaseColumnsListByTablePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByTableCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, options)
		},
		advancer: func(ctx context.Context, resp DatabaseColumnsListByTableResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatabaseColumnListResult.NextLink)
		},
	}
}

// listByTableCreateRequest creates the ListByTable request.
func (client *DatabaseColumnsClient) listByTableCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, options *DatabaseColumnsListByTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
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
func (client *DatabaseColumnsClient) listByTableHandleResponse(resp *http.Response) (DatabaseColumnsListByTableResponse, error) {
	result := DatabaseColumnsListByTableResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumnListResult); err != nil {
		return DatabaseColumnsListByTableResponse{}, err
	}
	return result, nil
}

// listByTableHandleError handles the ListByTable error response.
func (client *DatabaseColumnsClient) listByTableHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
