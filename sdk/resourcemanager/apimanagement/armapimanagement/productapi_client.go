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

// ProductAPIClient contains the methods for the ProductAPI group.
// Don't use this type directly, use NewProductAPIClient() instead.
type ProductAPIClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProductAPIClient creates a new instance of ProductAPIClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProductAPIClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProductAPIClient, error) {
	cl, err := arm.NewClient(moduleName+".ProductAPIClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProductAPIClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckEntityExists - Checks that API entity specified by identifier is associated with the Product entity.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - options - ProductAPIClientCheckEntityExistsOptions contains the optional parameters for the ProductAPIClient.CheckEntityExists
//     method.
func (client *ProductAPIClient) CheckEntityExists(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *ProductAPIClientCheckEntityExistsOptions) (ProductAPIClientCheckEntityExistsResponse, error) {
	req, err := client.checkEntityExistsCreateRequest(ctx, resourceGroupName, serviceName, productID, apiID, options)
	if err != nil {
		return ProductAPIClientCheckEntityExistsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductAPIClientCheckEntityExistsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return ProductAPIClientCheckEntityExistsResponse{}, runtime.NewResponseError(resp)
	}
	return ProductAPIClientCheckEntityExistsResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// checkEntityExistsCreateRequest creates the CheckEntityExists request.
func (client *ProductAPIClient) checkEntityExistsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *ProductAPIClientCheckEntityExistsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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

// CreateOrUpdate - Adds an API to the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - options - ProductAPIClientCreateOrUpdateOptions contains the optional parameters for the ProductAPIClient.CreateOrUpdate
//     method.
func (client *ProductAPIClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *ProductAPIClientCreateOrUpdateOptions) (ProductAPIClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, productID, apiID, options)
	if err != nil {
		return ProductAPIClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductAPIClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProductAPIClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProductAPIClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *ProductAPIClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProductAPIClient) createOrUpdateHandleResponse(resp *http.Response) (ProductAPIClientCreateOrUpdateResponse, error) {
	result := ProductAPIClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIContract); err != nil {
		return ProductAPIClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified API from the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - options - ProductAPIClientDeleteOptions contains the optional parameters for the ProductAPIClient.Delete method.
func (client *ProductAPIClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *ProductAPIClientDeleteOptions) (ProductAPIClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, productID, apiID, options)
	if err != nil {
		return ProductAPIClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductAPIClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProductAPIClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProductAPIClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProductAPIClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiID string, options *ProductAPIClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByProductPager - Lists a collection of the APIs associated with a product.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - options - ProductAPIClientListByProductOptions contains the optional parameters for the ProductAPIClient.NewListByProductPager
//     method.
func (client *ProductAPIClient) NewListByProductPager(resourceGroupName string, serviceName string, productID string, options *ProductAPIClientListByProductOptions) *runtime.Pager[ProductAPIClientListByProductResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductAPIClientListByProductResponse]{
		More: func(page ProductAPIClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductAPIClientListByProductResponse) (ProductAPIClientListByProductResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, productID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProductAPIClientListByProductResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProductAPIClientListByProductResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProductAPIClientListByProductResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProductHandleResponse(resp)
		},
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *ProductAPIClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *ProductAPIClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apis"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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

// listByProductHandleResponse handles the ListByProduct response.
func (client *ProductAPIClient) listByProductHandleResponse(resp *http.Response) (ProductAPIClientListByProductResponse, error) {
	result := ProductAPIClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APICollection); err != nil {
		return ProductAPIClientListByProductResponse{}, err
	}
	return result, nil
}
