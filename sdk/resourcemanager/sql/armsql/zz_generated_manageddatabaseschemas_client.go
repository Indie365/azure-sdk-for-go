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

// ManagedDatabaseSchemasClient contains the methods for the ManagedDatabaseSchemas group.
// Don't use this type directly, use NewManagedDatabaseSchemasClient() instead.
type ManagedDatabaseSchemasClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedDatabaseSchemasClient creates a new instance of ManagedDatabaseSchemasClient with the specified values.
func NewManagedDatabaseSchemasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedDatabaseSchemasClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ManagedDatabaseSchemasClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get managed database schema
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseSchemasClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, options *ManagedDatabaseSchemasGetOptions) (ManagedDatabaseSchemasGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, options)
	if err != nil {
		return ManagedDatabaseSchemasGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedDatabaseSchemasGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseSchemasGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseSchemasClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, options *ManagedDatabaseSchemasGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}"
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
func (client *ManagedDatabaseSchemasClient) getHandleResponse(resp *http.Response) (ManagedDatabaseSchemasGetResponse, error) {
	result := ManagedDatabaseSchemasGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSchema); err != nil {
		return ManagedDatabaseSchemasGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedDatabaseSchemasClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabase - List managed database schemas
// If the operation fails it returns a generic error.
func (client *ManagedDatabaseSchemasClient) ListByDatabase(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSchemasListByDatabaseOptions) *ManagedDatabaseSchemasListByDatabasePager {
	return &ManagedDatabaseSchemasListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp ManagedDatabaseSchemasListByDatabaseResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatabaseSchemaListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedDatabaseSchemasClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSchemasListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedDatabaseSchemasClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseSchemasListByDatabaseResponse, error) {
	result := ManagedDatabaseSchemasListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSchemaListResult); err != nil {
		return ManagedDatabaseSchemasListByDatabaseResponse{}, err
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *ManagedDatabaseSchemasClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
