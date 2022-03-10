//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// ServicesClient contains the methods for the Services group.
// Don't use this type directly, use NewServicesClient() instead.
type ServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServicesClient creates a new instance of ServicesClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ServicesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CheckNameAvailability - Checks that the resource name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
// location - the region
// availabilityParameters - Parameters supplied to the operation.
// options - ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
// method.
func (client *ServicesClient) CheckNameAvailability(ctx context.Context, location string, availabilityParameters NameAvailabilityParameters, options *ServicesClientCheckNameAvailabilityOptions) (ServicesClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, availabilityParameters, options)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ServicesClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, availabilityParameters NameAvailabilityParameters, options *ServicesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, availabilityParameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ServicesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ServicesClientCheckNameAvailabilityResponse, error) {
	result := ServicesClientCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return ServicesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create a new Service or update an exiting Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// resource - Parameters for the create or update operation
// options - ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate
// method.
func (client *ServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, resource ServiceResource, options *ServicesClientBeginCreateOrUpdateOptions) (ServicesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, resource, options)
	if err != nil {
		return ServicesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ServicesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServicesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ServicesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServicesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create a new Service or update an exiting Service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, resource ServiceResource, options *ServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, resource, options)
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
func (client *ServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, resource ServiceResource, options *ServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resource)
}

// BeginDelete - Operation to delete a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServicesClientBeginDeleteOptions contains the optional parameters for the ServicesClient.BeginDelete method.
func (client *ServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginDeleteOptions) (ServicesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientDeletePollerResponse{}, err
	}
	result := ServicesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServicesClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ServicesClientDeletePollerResponse{}, err
	}
	result.Poller = &ServicesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Operation to delete a Service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, options)
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
func (client *ServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// DisableTestEndpoint - Disable test endpoint functionality for a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServicesClientDisableTestEndpointOptions contains the optional parameters for the ServicesClient.DisableTestEndpoint
// method.
func (client *ServicesClient) DisableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientDisableTestEndpointOptions) (ServicesClientDisableTestEndpointResponse, error) {
	req, err := client.disableTestEndpointCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientDisableTestEndpointResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientDisableTestEndpointResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientDisableTestEndpointResponse{}, runtime.NewResponseError(resp)
	}
	return ServicesClientDisableTestEndpointResponse{RawResponse: resp}, nil
}

// disableTestEndpointCreateRequest creates the DisableTestEndpoint request.
func (client *ServicesClient) disableTestEndpointCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientDisableTestEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/disableTestEndpoint"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// EnableTestEndpoint - Enable test endpoint functionality for a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServicesClientEnableTestEndpointOptions contains the optional parameters for the ServicesClient.EnableTestEndpoint
// method.
func (client *ServicesClient) EnableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientEnableTestEndpointOptions) (ServicesClientEnableTestEndpointResponse, error) {
	req, err := client.enableTestEndpointCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientEnableTestEndpointResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientEnableTestEndpointResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientEnableTestEndpointResponse{}, runtime.NewResponseError(resp)
	}
	return client.enableTestEndpointHandleResponse(resp)
}

// enableTestEndpointCreateRequest creates the EnableTestEndpoint request.
func (client *ServicesClient) enableTestEndpointCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientEnableTestEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/enableTestEndpoint"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// enableTestEndpointHandleResponse handles the EnableTestEndpoint response.
func (client *ServicesClient) enableTestEndpointHandleResponse(resp *http.Response) (ServicesClientEnableTestEndpointResponse, error) {
	result := ServicesClientEnableTestEndpointResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestKeys); err != nil {
		return ServicesClientEnableTestEndpointResponse{}, err
	}
	return result, nil
}

// Get - Get a Service and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
func (client *ServicesClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientGetOptions) (ServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServicesClient) getHandleResponse(resp *http.Response) (ServicesClientGetResponse, error) {
	result := ServicesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServicesClientGetResponse{}, err
	}
	return result, nil
}

// List - Handles requests to list all resources in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - ServicesClientListOptions contains the optional parameters for the ServicesClient.List method.
func (client *ServicesClient) List(resourceGroupName string, options *ServicesClientListOptions) *ServicesClientListPager {
	return &ServicesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ServicesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServiceResourceList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring"
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
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServicesClient) listHandleResponse(resp *http.Response) (ServicesClientListResponse, error) {
	result := ServicesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceList); err != nil {
		return ServicesClientListResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Handles requests to list all resources in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ServicesClientListBySubscriptionOptions contains the optional parameters for the ServicesClient.ListBySubscription
// method.
func (client *ServicesClient) ListBySubscription(options *ServicesClientListBySubscriptionOptions) *ServicesClientListBySubscriptionPager {
	return &ServicesClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ServicesClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServiceResourceList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/Spring"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (ServicesClientListBySubscriptionResponse, error) {
	result := ServicesClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceList); err != nil {
		return ServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListTestKeys - List test keys for a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServicesClientListTestKeysOptions contains the optional parameters for the ServicesClient.ListTestKeys method.
func (client *ServicesClient) ListTestKeys(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientListTestKeysOptions) (ServicesClientListTestKeysResponse, error) {
	req, err := client.listTestKeysCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientListTestKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientListTestKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientListTestKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listTestKeysHandleResponse(resp)
}

// listTestKeysCreateRequest creates the ListTestKeys request.
func (client *ServicesClient) listTestKeysCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientListTestKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/listTestKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listTestKeysHandleResponse handles the ListTestKeys response.
func (client *ServicesClient) listTestKeysHandleResponse(resp *http.Response) (ServicesClientListTestKeysResponse, error) {
	result := ServicesClientListTestKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestKeys); err != nil {
		return ServicesClientListTestKeysResponse{}, err
	}
	return result, nil
}

// RegenerateTestKey - Regenerate a test key for a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// regenerateTestKeyRequest - Parameters for the operation
// options - ServicesClientRegenerateTestKeyOptions contains the optional parameters for the ServicesClient.RegenerateTestKey
// method.
func (client *ServicesClient) RegenerateTestKey(ctx context.Context, resourceGroupName string, serviceName string, regenerateTestKeyRequest RegenerateTestKeyRequestPayload, options *ServicesClientRegenerateTestKeyOptions) (ServicesClientRegenerateTestKeyResponse, error) {
	req, err := client.regenerateTestKeyCreateRequest(ctx, resourceGroupName, serviceName, regenerateTestKeyRequest, options)
	if err != nil {
		return ServicesClientRegenerateTestKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServicesClientRegenerateTestKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServicesClientRegenerateTestKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateTestKeyHandleResponse(resp)
}

// regenerateTestKeyCreateRequest creates the RegenerateTestKey request.
func (client *ServicesClient) regenerateTestKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, regenerateTestKeyRequest RegenerateTestKeyRequestPayload, options *ServicesClientRegenerateTestKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/regenerateTestKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, regenerateTestKeyRequest)
}

// regenerateTestKeyHandleResponse handles the RegenerateTestKey response.
func (client *ServicesClient) regenerateTestKeyHandleResponse(resp *http.Response) (ServicesClientRegenerateTestKeyResponse, error) {
	result := ServicesClientRegenerateTestKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestKeys); err != nil {
		return ServicesClientRegenerateTestKeyResponse{}, err
	}
	return result, nil
}

// BeginStart - Start a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServicesClientBeginStartOptions contains the optional parameters for the ServicesClient.BeginStart method.
func (client *ServicesClient) BeginStart(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginStartOptions) (ServicesClientStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientStartPollerResponse{}, err
	}
	result := ServicesClientStartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServicesClient.Start", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ServicesClientStartPollerResponse{}, err
	}
	result.Poller = &ServicesClientStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - Start a Service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) start(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ServicesClient) startCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStop - Stop a Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ServicesClientBeginStopOptions contains the optional parameters for the ServicesClient.BeginStop method.
func (client *ServicesClient) BeginStop(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginStopOptions) (ServicesClientStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServicesClientStopPollerResponse{}, err
	}
	result := ServicesClientStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServicesClient.Stop", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ServicesClientStopPollerResponse{}, err
	}
	result.Poller = &ServicesClientStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Stop a Service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) stop(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ServicesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServicesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUpdate - Operation to update an exiting Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// resource - Parameters for the update operation
// options - ServicesClientBeginUpdateOptions contains the optional parameters for the ServicesClient.BeginUpdate method.
func (client *ServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, resource ServiceResource, options *ServicesClientBeginUpdateOptions) (ServicesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serviceName, resource, options)
	if err != nil {
		return ServicesClientUpdatePollerResponse{}, err
	}
	result := ServicesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServicesClient.Update", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ServicesClientUpdatePollerResponse{}, err
	}
	result.Poller = &ServicesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Operation to update an exiting Service.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServicesClient) update(ctx context.Context, resourceGroupName string, serviceName string, resource ServiceResource, options *ServicesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, resource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, resource ServiceResource, options *ServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, resource)
}
