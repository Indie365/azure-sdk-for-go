//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

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

// AccountFiltersClient contains the methods for the AccountFilters group.
// Don't use this type directly, use NewAccountFiltersClient() instead.
type AccountFiltersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAccountFiltersClient creates a new instance of AccountFiltersClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAccountFiltersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AccountFiltersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &AccountFiltersClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates an Account Filter in the Media Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// filterName - The Account Filter name
// parameters - The request parameters
// options - AccountFiltersClientCreateOrUpdateOptions contains the optional parameters for the AccountFiltersClient.CreateOrUpdate
// method.
func (client *AccountFiltersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, filterName string, parameters AccountFilter, options *AccountFiltersClientCreateOrUpdateOptions) (AccountFiltersClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, filterName, parameters, options)
	if err != nil {
		return AccountFiltersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountFiltersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AccountFiltersClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AccountFiltersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, filterName string, parameters AccountFilter, options *AccountFiltersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AccountFiltersClient) createOrUpdateHandleResponse(resp *http.Response) (AccountFiltersClientCreateOrUpdateResponse, error) {
	result := AccountFiltersClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountFilter); err != nil {
		return AccountFiltersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Account Filter in the Media Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// filterName - The Account Filter name
// options - AccountFiltersClientDeleteOptions contains the optional parameters for the AccountFiltersClient.Delete method.
func (client *AccountFiltersClient) Delete(ctx context.Context, resourceGroupName string, accountName string, filterName string, options *AccountFiltersClientDeleteOptions) (AccountFiltersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, filterName, options)
	if err != nil {
		return AccountFiltersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountFiltersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AccountFiltersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AccountFiltersClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccountFiltersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, filterName string, options *AccountFiltersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the details of an Account Filter in the Media Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// filterName - The Account Filter name
// options - AccountFiltersClientGetOptions contains the optional parameters for the AccountFiltersClient.Get method.
func (client *AccountFiltersClient) Get(ctx context.Context, resourceGroupName string, accountName string, filterName string, options *AccountFiltersClientGetOptions) (AccountFiltersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, filterName, options)
	if err != nil {
		return AccountFiltersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountFiltersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountFiltersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccountFiltersClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, filterName string, options *AccountFiltersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccountFiltersClient) getHandleResponse(resp *http.Response) (AccountFiltersClientGetResponse, error) {
	result := AccountFiltersClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountFilter); err != nil {
		return AccountFiltersClientGetResponse{}, err
	}
	return result, nil
}

// List - List Account Filters in the Media Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// options - AccountFiltersClientListOptions contains the optional parameters for the AccountFiltersClient.List method.
func (client *AccountFiltersClient) List(resourceGroupName string, accountName string, options *AccountFiltersClientListOptions) *AccountFiltersClientListPager {
	return &AccountFiltersClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp AccountFiltersClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccountFilterCollection.ODataNextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AccountFiltersClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AccountFiltersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccountFiltersClient) listHandleResponse(resp *http.Response) (AccountFiltersClientListResponse, error) {
	result := AccountFiltersClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountFilterCollection); err != nil {
		return AccountFiltersClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing Account Filter in the Media Services account.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// filterName - The Account Filter name
// parameters - The request parameters
// options - AccountFiltersClientUpdateOptions contains the optional parameters for the AccountFiltersClient.Update method.
func (client *AccountFiltersClient) Update(ctx context.Context, resourceGroupName string, accountName string, filterName string, parameters AccountFilter, options *AccountFiltersClientUpdateOptions) (AccountFiltersClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, filterName, parameters, options)
	if err != nil {
		return AccountFiltersClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccountFiltersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccountFiltersClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AccountFiltersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, filterName string, parameters AccountFilter, options *AccountFiltersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/accountFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AccountFiltersClient) updateHandleResponse(resp *http.Response) (AccountFiltersClientUpdateResponse, error) {
	result := AccountFiltersClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountFilter); err != nil {
		return AccountFiltersClientUpdateResponse{}, err
	}
	return result, nil
}
