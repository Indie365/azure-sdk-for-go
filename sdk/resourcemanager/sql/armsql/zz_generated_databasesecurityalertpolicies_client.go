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

// DatabaseSecurityAlertPoliciesClient contains the methods for the DatabaseSecurityAlertPolicies group.
// Don't use this type directly, use NewDatabaseSecurityAlertPoliciesClient() instead.
type DatabaseSecurityAlertPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDatabaseSecurityAlertPoliciesClient creates a new instance of DatabaseSecurityAlertPoliciesClient with the specified values.
func NewDatabaseSecurityAlertPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DatabaseSecurityAlertPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DatabaseSecurityAlertPoliciesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates a database's security alert policy.
// If the operation fails it returns a generic error.
func (client *DatabaseSecurityAlertPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, parameters DatabaseSecurityAlertPolicy, options *DatabaseSecurityAlertPoliciesCreateOrUpdateOptions) (DatabaseSecurityAlertPoliciesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, securityAlertPolicyName, parameters, options)
	if err != nil {
		return DatabaseSecurityAlertPoliciesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseSecurityAlertPoliciesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DatabaseSecurityAlertPoliciesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabaseSecurityAlertPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, parameters DatabaseSecurityAlertPolicy, options *DatabaseSecurityAlertPoliciesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}"
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
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DatabaseSecurityAlertPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (DatabaseSecurityAlertPoliciesCreateOrUpdateResponse, error) {
	result := DatabaseSecurityAlertPoliciesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSecurityAlertPolicy); err != nil {
		return DatabaseSecurityAlertPoliciesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DatabaseSecurityAlertPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a database's security alert policy.
// If the operation fails it returns a generic error.
func (client *DatabaseSecurityAlertPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, options *DatabaseSecurityAlertPoliciesGetOptions) (DatabaseSecurityAlertPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, securityAlertPolicyName, options)
	if err != nil {
		return DatabaseSecurityAlertPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabaseSecurityAlertPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseSecurityAlertPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabaseSecurityAlertPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, options *DatabaseSecurityAlertPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}"
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
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
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
func (client *DatabaseSecurityAlertPoliciesClient) getHandleResponse(resp *http.Response) (DatabaseSecurityAlertPoliciesGetResponse, error) {
	result := DatabaseSecurityAlertPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSecurityAlertPolicy); err != nil {
		return DatabaseSecurityAlertPoliciesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DatabaseSecurityAlertPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByDatabase - Gets a list of database's security alert policies.
// If the operation fails it returns a generic error.
func (client *DatabaseSecurityAlertPoliciesClient) ListByDatabase(resourceGroupName string, serverName string, databaseName string, options *DatabaseSecurityAlertPoliciesListByDatabaseOptions) *DatabaseSecurityAlertPoliciesListByDatabasePager {
	return &DatabaseSecurityAlertPoliciesListByDatabasePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
		},
		advancer: func(ctx context.Context, resp DatabaseSecurityAlertPoliciesListByDatabaseResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatabaseSecurityAlertListResult.NextLink)
		},
	}
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DatabaseSecurityAlertPoliciesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DatabaseSecurityAlertPoliciesListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies"
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
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DatabaseSecurityAlertPoliciesClient) listByDatabaseHandleResponse(resp *http.Response) (DatabaseSecurityAlertPoliciesListByDatabaseResponse, error) {
	result := DatabaseSecurityAlertPoliciesListByDatabaseResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseSecurityAlertListResult); err != nil {
		return DatabaseSecurityAlertPoliciesListByDatabaseResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByDatabaseHandleError handles the ListByDatabase error response.
func (client *DatabaseSecurityAlertPoliciesClient) listByDatabaseHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
