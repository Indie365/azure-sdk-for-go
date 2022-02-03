//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor

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

// FrontDoorsClient contains the methods for the FrontDoors group.
// Don't use this type directly, use NewFrontDoorsClient() instead.
type FrontDoorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFrontDoorsClient creates a new instance of FrontDoorsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFrontDoorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FrontDoorsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &FrontDoorsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a new Front Door with a Front Door name under the specified subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// frontDoorParameters - Front Door properties needed to create a new Front Door.
// options - FrontDoorsClientBeginCreateOrUpdateOptions contains the optional parameters for the FrontDoorsClient.BeginCreateOrUpdate
// method.
func (client *FrontDoorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters FrontDoor, options *FrontDoorsClientBeginCreateOrUpdateOptions) (FrontDoorsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, frontDoorName, frontDoorParameters, options)
	if err != nil {
		return FrontDoorsClientCreateOrUpdatePollerResponse{}, err
	}
	result := FrontDoorsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("FrontDoorsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return FrontDoorsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &FrontDoorsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a new Front Door with a Front Door name under the specified subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FrontDoorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters FrontDoor, options *FrontDoorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, frontDoorName, frontDoorParameters, options)
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
func (client *FrontDoorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters FrontDoor, options *FrontDoorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, frontDoorParameters)
}

// BeginDelete - Deletes an existing Front Door with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// options - FrontDoorsClientBeginDeleteOptions contains the optional parameters for the FrontDoorsClient.BeginDelete method.
func (client *FrontDoorsClient) BeginDelete(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientBeginDeleteOptions) (FrontDoorsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, frontDoorName, options)
	if err != nil {
		return FrontDoorsClientDeletePollerResponse{}, err
	}
	result := FrontDoorsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("FrontDoorsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return FrontDoorsClientDeletePollerResponse{}, err
	}
	result.Poller = &FrontDoorsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing Front Door with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FrontDoorsClient) deleteOperation(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, frontDoorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FrontDoorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a Front Door with the specified Front Door name under the specified subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// options - FrontDoorsClientGetOptions contains the optional parameters for the FrontDoorsClient.Get method.
func (client *FrontDoorsClient) Get(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientGetOptions) (FrontDoorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, frontDoorName, options)
	if err != nil {
		return FrontDoorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FrontDoorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FrontDoorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FrontDoorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FrontDoorsClient) getHandleResponse(resp *http.Response) (FrontDoorsClientGetResponse, error) {
	result := FrontDoorsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.FrontDoor); err != nil {
		return FrontDoorsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all of the Front Doors within an Azure subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - FrontDoorsClientListOptions contains the optional parameters for the FrontDoorsClient.List method.
func (client *FrontDoorsClient) List(options *FrontDoorsClientListOptions) *FrontDoorsClientListPager {
	return &FrontDoorsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp FrontDoorsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *FrontDoorsClient) listCreateRequest(ctx context.Context, options *FrontDoorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/frontDoors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FrontDoorsClient) listHandleResponse(resp *http.Response) (FrontDoorsClientListResponse, error) {
	result := FrontDoorsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return FrontDoorsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all of the Front Doors within a resource group under a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// options - FrontDoorsClientListByResourceGroupOptions contains the optional parameters for the FrontDoorsClient.ListByResourceGroup
// method.
func (client *FrontDoorsClient) ListByResourceGroup(resourceGroupName string, options *FrontDoorsClientListByResourceGroupOptions) *FrontDoorsClientListByResourceGroupPager {
	return &FrontDoorsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp FrontDoorsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FrontDoorsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FrontDoorsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FrontDoorsClient) listByResourceGroupHandleResponse(resp *http.Response) (FrontDoorsClientListByResourceGroupResponse, error) {
	result := FrontDoorsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return FrontDoorsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ValidateCustomDomain - Validates the custom domain mapping to ensure it maps to the correct Front Door endpoint in DNS.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Name of the Resource group within the Azure subscription.
// frontDoorName - Name of the Front Door which is globally unique.
// customDomainProperties - Custom domain to be validated.
// options - FrontDoorsClientValidateCustomDomainOptions contains the optional parameters for the FrontDoorsClient.ValidateCustomDomain
// method.
func (client *FrontDoorsClient) ValidateCustomDomain(ctx context.Context, resourceGroupName string, frontDoorName string, customDomainProperties ValidateCustomDomainInput, options *FrontDoorsClientValidateCustomDomainOptions) (FrontDoorsClientValidateCustomDomainResponse, error) {
	req, err := client.validateCustomDomainCreateRequest(ctx, resourceGroupName, frontDoorName, customDomainProperties, options)
	if err != nil {
		return FrontDoorsClientValidateCustomDomainResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FrontDoorsClientValidateCustomDomainResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FrontDoorsClientValidateCustomDomainResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateCustomDomainHandleResponse(resp)
}

// validateCustomDomainCreateRequest creates the ValidateCustomDomain request.
func (client *FrontDoorsClient) validateCustomDomainCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, customDomainProperties ValidateCustomDomainInput, options *FrontDoorsClientValidateCustomDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/validateCustomDomain"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customDomainProperties)
}

// validateCustomDomainHandleResponse handles the ValidateCustomDomain response.
func (client *FrontDoorsClient) validateCustomDomainHandleResponse(resp *http.Response) (FrontDoorsClientValidateCustomDomainResponse, error) {
	result := FrontDoorsClientValidateCustomDomainResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateCustomDomainOutput); err != nil {
		return FrontDoorsClientValidateCustomDomainResponse{}, err
	}
	return result, nil
}
