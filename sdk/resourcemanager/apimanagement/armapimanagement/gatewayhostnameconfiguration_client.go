//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GatewayHostnameConfigurationClient contains the methods for the GatewayHostnameConfiguration group.
// Don't use this type directly, use NewGatewayHostnameConfigurationClient() instead.
type GatewayHostnameConfigurationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGatewayHostnameConfigurationClient creates a new instance of GatewayHostnameConfigurationClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGatewayHostnameConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GatewayHostnameConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName+".GatewayHostnameConfigurationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GatewayHostnameConfigurationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates of updates hostname configuration for a Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
//     'managed'
//   - hcID - Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
//   - options - GatewayHostnameConfigurationClientCreateOrUpdateOptions contains the optional parameters for the GatewayHostnameConfigurationClient.CreateOrUpdate
//     method.
func (client *GatewayHostnameConfigurationClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, parameters GatewayHostnameConfigurationContract, options *GatewayHostnameConfigurationClientCreateOrUpdateOptions) (GatewayHostnameConfigurationClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, hcID, parameters, options)
	if err != nil {
		return GatewayHostnameConfigurationClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewayHostnameConfigurationClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return GatewayHostnameConfigurationClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GatewayHostnameConfigurationClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, parameters GatewayHostnameConfigurationContract, options *GatewayHostnameConfigurationClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if hcID == "" {
		return nil, errors.New("parameter hcID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcId}", url.PathEscape(hcID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GatewayHostnameConfigurationClient) createOrUpdateHandleResponse(resp *http.Response) (GatewayHostnameConfigurationClientCreateOrUpdateResponse, error) {
	result := GatewayHostnameConfigurationClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayHostnameConfigurationContract); err != nil {
		return GatewayHostnameConfigurationClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified hostname configuration from the specified Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
//     'managed'
//   - hcID - Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - GatewayHostnameConfigurationClientDeleteOptions contains the optional parameters for the GatewayHostnameConfigurationClient.Delete
//     method.
func (client *GatewayHostnameConfigurationClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, ifMatch string, options *GatewayHostnameConfigurationClientDeleteOptions) (GatewayHostnameConfigurationClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, hcID, ifMatch, options)
	if err != nil {
		return GatewayHostnameConfigurationClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewayHostnameConfigurationClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return GatewayHostnameConfigurationClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return GatewayHostnameConfigurationClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewayHostnameConfigurationClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, ifMatch string, options *GatewayHostnameConfigurationClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if hcID == "" {
		return nil, errors.New("parameter hcID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcId}", url.PathEscape(hcID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get details of a hostname configuration
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
//     'managed'
//   - hcID - Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
//   - options - GatewayHostnameConfigurationClientGetOptions contains the optional parameters for the GatewayHostnameConfigurationClient.Get
//     method.
func (client *GatewayHostnameConfigurationClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, options *GatewayHostnameConfigurationClientGetOptions) (GatewayHostnameConfigurationClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, hcID, options)
	if err != nil {
		return GatewayHostnameConfigurationClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewayHostnameConfigurationClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayHostnameConfigurationClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GatewayHostnameConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, options *GatewayHostnameConfigurationClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if hcID == "" {
		return nil, errors.New("parameter hcID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcId}", url.PathEscape(hcID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GatewayHostnameConfigurationClient) getHandleResponse(resp *http.Response) (GatewayHostnameConfigurationClientGetResponse, error) {
	result := GatewayHostnameConfigurationClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayHostnameConfigurationContract); err != nil {
		return GatewayHostnameConfigurationClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Checks that hostname configuration entity specified by identifier exists for specified Gateway entity.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
//     'managed'
//   - hcID - Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
//   - options - GatewayHostnameConfigurationClientGetEntityTagOptions contains the optional parameters for the GatewayHostnameConfigurationClient.GetEntityTag
//     method.
func (client *GatewayHostnameConfigurationClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, options *GatewayHostnameConfigurationClientGetEntityTagOptions) (GatewayHostnameConfigurationClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, hcID, options)
	if err != nil {
		return GatewayHostnameConfigurationClientGetEntityTagResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GatewayHostnameConfigurationClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GatewayHostnameConfigurationClientGetEntityTagResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *GatewayHostnameConfigurationClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, hcID string, options *GatewayHostnameConfigurationClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations/{hcId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if hcID == "" {
		return nil, errors.New("parameter hcID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hcId}", url.PathEscape(hcID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *GatewayHostnameConfigurationClient) getEntityTagHandleResponse(resp *http.Response) (GatewayHostnameConfigurationClientGetEntityTagResponse, error) {
	result := GatewayHostnameConfigurationClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists the collection of hostname configurations for the specified gateway.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - gatewayID - Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value
//     'managed'
//   - options - GatewayHostnameConfigurationClientListByServiceOptions contains the optional parameters for the GatewayHostnameConfigurationClient.NewListByServicePager
//     method.
func (client *GatewayHostnameConfigurationClient) NewListByServicePager(resourceGroupName string, serviceName string, gatewayID string, options *GatewayHostnameConfigurationClientListByServiceOptions) *runtime.Pager[GatewayHostnameConfigurationClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[GatewayHostnameConfigurationClientListByServiceResponse]{
		More: func(page GatewayHostnameConfigurationClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GatewayHostnameConfigurationClientListByServiceResponse) (GatewayHostnameConfigurationClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GatewayHostnameConfigurationClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return GatewayHostnameConfigurationClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GatewayHostnameConfigurationClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *GatewayHostnameConfigurationClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, options *GatewayHostnameConfigurationClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/hostnameConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *GatewayHostnameConfigurationClient) listByServiceHandleResponse(resp *http.Response) (GatewayHostnameConfigurationClientListByServiceResponse, error) {
	result := GatewayHostnameConfigurationClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GatewayHostnameConfigurationCollection); err != nil {
		return GatewayHostnameConfigurationClientListByServiceResponse{}, err
	}
	return result, nil
}
