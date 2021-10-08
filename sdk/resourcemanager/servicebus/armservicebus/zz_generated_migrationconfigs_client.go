//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicebus

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// MigrationConfigsClient contains the methods for the MigrationConfigs group.
// Don't use this type directly, use NewMigrationConfigsClient() instead.
type MigrationConfigsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMigrationConfigsClient creates a new instance of MigrationConfigsClient with the specified values.
func NewMigrationConfigsClient(con *arm.Connection, subscriptionID string) *MigrationConfigsClient {
	return &MigrationConfigsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CompleteMigration - This operation Completes Migration of entities by pointing the connection strings to Premium namespace and any entities created after
// the operation will be under Premium Namespace. CompleteMigration
// operation will fail when entity migration is in-progress.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MigrationConfigsClient) CompleteMigration(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsCompleteMigrationOptions) (MigrationConfigsCompleteMigrationResponse, error) {
	req, err := client.completeMigrationCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsCompleteMigrationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsCompleteMigrationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationConfigsCompleteMigrationResponse{}, client.completeMigrationHandleError(resp)
	}
	return MigrationConfigsCompleteMigrationResponse{RawResponse: resp}, nil
}

// completeMigrationCreateRequest creates the CompleteMigration request.
func (client *MigrationConfigsClient) completeMigrationCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsCompleteMigrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}/upgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// completeMigrationHandleError handles the CompleteMigration error response.
func (client *MigrationConfigsClient) completeMigrationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateAndStartMigration - Creates Migration configuration and starts migration of entities from Standard to Premium namespace
// If the operation fails it returns the *ErrorResponse error type.
func (client *MigrationConfigsClient) BeginCreateAndStartMigration(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, parameters MigrationConfigProperties, options *MigrationConfigsBeginCreateAndStartMigrationOptions) (MigrationConfigsCreateAndStartMigrationPollerResponse, error) {
	resp, err := client.createAndStartMigration(ctx, resourceGroupName, namespaceName, configName, parameters, options)
	if err != nil {
		return MigrationConfigsCreateAndStartMigrationPollerResponse{}, err
	}
	result := MigrationConfigsCreateAndStartMigrationPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MigrationConfigsClient.CreateAndStartMigration", "", resp, client.pl, client.createAndStartMigrationHandleError)
	if err != nil {
		return MigrationConfigsCreateAndStartMigrationPollerResponse{}, err
	}
	result.Poller = &MigrationConfigsCreateAndStartMigrationPoller{
		pt: pt,
	}
	return result, nil
}

// CreateAndStartMigration - Creates Migration configuration and starts migration of entities from Standard to Premium namespace
// If the operation fails it returns the *ErrorResponse error type.
func (client *MigrationConfigsClient) createAndStartMigration(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, parameters MigrationConfigProperties, options *MigrationConfigsBeginCreateAndStartMigrationOptions) (*http.Response, error) {
	req, err := client.createAndStartMigrationCreateRequest(ctx, resourceGroupName, namespaceName, configName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createAndStartMigrationHandleError(resp)
	}
	return resp, nil
}

// createAndStartMigrationCreateRequest creates the CreateAndStartMigration request.
func (client *MigrationConfigsClient) createAndStartMigrationCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, parameters MigrationConfigProperties, options *MigrationConfigsBeginCreateAndStartMigrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createAndStartMigrationHandleError handles the CreateAndStartMigration error response.
func (client *MigrationConfigsClient) createAndStartMigrationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a MigrationConfiguration
// If the operation fails it returns the *ErrorResponse error type.
func (client *MigrationConfigsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsDeleteOptions) (MigrationConfigsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return MigrationConfigsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return MigrationConfigsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MigrationConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *MigrationConfigsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves Migration Config
// If the operation fails it returns the *ErrorResponse error type.
func (client *MigrationConfigsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsGetOptions) (MigrationConfigsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationConfigsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MigrationConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MigrationConfigsClient) getHandleResponse(resp *http.Response) (MigrationConfigsGetResponse, error) {
	result := MigrationConfigsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationConfigProperties); err != nil {
		return MigrationConfigsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MigrationConfigsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all migrationConfigurations
// If the operation fails it returns the *ErrorResponse error type.
func (client *MigrationConfigsClient) List(resourceGroupName string, namespaceName string, options *MigrationConfigsListOptions) *MigrationConfigsListPager {
	return &MigrationConfigsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		advancer: func(ctx context.Context, resp MigrationConfigsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MigrationConfigListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MigrationConfigsClient) listCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *MigrationConfigsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MigrationConfigsClient) listHandleResponse(resp *http.Response) (MigrationConfigsListResponse, error) {
	result := MigrationConfigsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationConfigListResult); err != nil {
		return MigrationConfigsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MigrationConfigsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Revert - This operation reverts Migration
// If the operation fails it returns the *ErrorResponse error type.
func (client *MigrationConfigsClient) Revert(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsRevertOptions) (MigrationConfigsRevertResponse, error) {
	req, err := client.revertCreateRequest(ctx, resourceGroupName, namespaceName, configName, options)
	if err != nil {
		return MigrationConfigsRevertResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationConfigsRevertResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationConfigsRevertResponse{}, client.revertHandleError(resp)
	}
	return MigrationConfigsRevertResponse{RawResponse: resp}, nil
}

// revertCreateRequest creates the Revert request.
func (client *MigrationConfigsClient) revertCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, configName MigrationConfigurationName, options *MigrationConfigsRevertOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/migrationConfigurations/{configName}/revert"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if configName == "" {
		return nil, errors.New("parameter configName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configName}", url.PathEscape(string(configName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// revertHandleError handles the Revert error response.
func (client *MigrationConfigsClient) revertHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
