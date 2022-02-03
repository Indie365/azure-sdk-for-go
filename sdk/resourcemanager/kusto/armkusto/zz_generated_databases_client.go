//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

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

// DatabasesClient contains the methods for the Databases group.
// Don't use this type directly, use NewDatabasesClient() instead.
type DatabasesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDatabasesClient creates a new instance of DatabasesClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DatabasesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DatabasesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// AddPrincipals - Add Database principals permissions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// databasePrincipalsToAdd - List of database principals to add.
// options - DatabasesClientAddPrincipalsOptions contains the optional parameters for the DatabasesClient.AddPrincipals method.
func (client *DatabasesClient) AddPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest, options *DatabasesClientAddPrincipalsOptions) (DatabasesClientAddPrincipalsResponse, error) {
	req, err := client.addPrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToAdd, options)
	if err != nil {
		return DatabasesClientAddPrincipalsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesClientAddPrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientAddPrincipalsResponse{}, runtime.NewResponseError(resp)
	}
	return client.addPrincipalsHandleResponse(resp)
}

// addPrincipalsCreateRequest creates the AddPrincipals request.
func (client *DatabasesClient) addPrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest, options *DatabasesClientAddPrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/addPrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, databasePrincipalsToAdd)
}

// addPrincipalsHandleResponse handles the AddPrincipals response.
func (client *DatabasesClient) addPrincipalsHandleResponse(resp *http.Response) (DatabasesClientAddPrincipalsResponse, error) {
	result := DatabasesClientAddPrincipalsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesClientAddPrincipalsResponse{}, err
	}
	return result, nil
}

// CheckNameAvailability - Checks that the databases resource name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// resourceName - The name of the resource.
// options - DatabasesClientCheckNameAvailabilityOptions contains the optional parameters for the DatabasesClient.CheckNameAvailability
// method.
func (client *DatabasesClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest, options *DatabasesClientCheckNameAvailabilityOptions) (DatabasesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, resourceName, options)
	if err != nil {
		return DatabasesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DatabasesClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest, options *DatabasesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/checkNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resourceName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DatabasesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (DatabasesClientCheckNameAvailabilityResponse, error) {
	result := DatabasesClientCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return DatabasesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// parameters - The database parameters supplied to the CreateOrUpdate operation.
// options - DatabasesClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabasesClient.BeginCreateOrUpdate
// method.
func (client *DatabasesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginCreateOrUpdateOptions) (DatabasesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return DatabasesClientCreateOrUpdatePollerResponse{}, err
	}
	result := DatabasesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DatabasesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return DatabasesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DatabasesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DatabasesClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabasesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the database with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// options - DatabasesClientBeginDeleteOptions contains the optional parameters for the DatabasesClient.BeginDelete method.
func (client *DatabasesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (DatabasesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesClientDeletePollerResponse{}, err
	}
	result := DatabasesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DatabasesClient.Delete", "", resp, client.pl)
	if err != nil {
		return DatabasesClientDeletePollerResponse{}, err
	}
	result.Poller = &DatabasesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the database with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DatabasesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DatabasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns a database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// options - DatabasesClientGetOptions contains the optional parameters for the DatabasesClient.Get method.
func (client *DatabasesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientGetOptions) (DatabasesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabasesClient) getHandleResponse(resp *http.Response) (DatabasesClientGetResponse, error) {
	result := DatabasesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DatabasesClientGetResponse{}, err
	}
	return result, nil
}

// ListByCluster - Returns the list of databases of the given Kusto cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// options - DatabasesClientListByClusterOptions contains the optional parameters for the DatabasesClient.ListByCluster method.
func (client *DatabasesClient) ListByCluster(ctx context.Context, resourceGroupName string, clusterName string, options *DatabasesClientListByClusterOptions) (DatabasesClientListByClusterResponse, error) {
	req, err := client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return DatabasesClientListByClusterResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesClientListByClusterResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientListByClusterResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByClusterHandleResponse(resp)
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *DatabasesClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *DatabasesClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *DatabasesClient) listByClusterHandleResponse(resp *http.Response) (DatabasesClientListByClusterResponse, error) {
	result := DatabasesClientListByClusterResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseListResult); err != nil {
		return DatabasesClientListByClusterResponse{}, err
	}
	return result, nil
}

// ListPrincipals - Returns a list of database principals of the given Kusto cluster and database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// options - DatabasesClientListPrincipalsOptions contains the optional parameters for the DatabasesClient.ListPrincipals
// method.
func (client *DatabasesClient) ListPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientListPrincipalsOptions) (DatabasesClientListPrincipalsResponse, error) {
	req, err := client.listPrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesClientListPrincipalsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesClientListPrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientListPrincipalsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listPrincipalsHandleResponse(resp)
}

// listPrincipalsCreateRequest creates the ListPrincipals request.
func (client *DatabasesClient) listPrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientListPrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/listPrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listPrincipalsHandleResponse handles the ListPrincipals response.
func (client *DatabasesClient) listPrincipalsHandleResponse(resp *http.Response) (DatabasesClientListPrincipalsResponse, error) {
	result := DatabasesClientListPrincipalsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesClientListPrincipalsResponse{}, err
	}
	return result, nil
}

// RemovePrincipals - Remove Database principals permissions.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// databasePrincipalsToRemove - List of database principals to remove.
// options - DatabasesClientRemovePrincipalsOptions contains the optional parameters for the DatabasesClient.RemovePrincipals
// method.
func (client *DatabasesClient) RemovePrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest, options *DatabasesClientRemovePrincipalsOptions) (DatabasesClientRemovePrincipalsResponse, error) {
	req, err := client.removePrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToRemove, options)
	if err != nil {
		return DatabasesClientRemovePrincipalsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatabasesClientRemovePrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientRemovePrincipalsResponse{}, runtime.NewResponseError(resp)
	}
	return client.removePrincipalsHandleResponse(resp)
}

// removePrincipalsCreateRequest creates the RemovePrincipals request.
func (client *DatabasesClient) removePrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest, options *DatabasesClientRemovePrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/removePrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, databasePrincipalsToRemove)
}

// removePrincipalsHandleResponse handles the RemovePrincipals response.
func (client *DatabasesClient) removePrincipalsHandleResponse(resp *http.Response) (DatabasesClientRemovePrincipalsResponse, error) {
	result := DatabasesClientRemovePrincipalsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesClientRemovePrincipalsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group containing the Kusto cluster.
// clusterName - The name of the Kusto cluster.
// databaseName - The name of the database in the Kusto cluster.
// parameters - The database parameters supplied to the Update operation.
// options - DatabasesClientBeginUpdateOptions contains the optional parameters for the DatabasesClient.BeginUpdate method.
func (client *DatabasesClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginUpdateOptions) (DatabasesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return DatabasesClientUpdatePollerResponse{}, err
	}
	result := DatabasesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DatabasesClient.Update", "", resp, client.pl)
	if err != nil {
		return DatabasesClientUpdatePollerResponse{}, err
	}
	result.Poller = &DatabasesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DatabasesClient) update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DatabasesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-27")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
