//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CustomAssessmentAutomationsClient contains the methods for the CustomAssessmentAutomations group.
// Don't use this type directly, use NewCustomAssessmentAutomationsClient() instead.
type CustomAssessmentAutomationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomAssessmentAutomationsClient creates a new instance of CustomAssessmentAutomationsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomAssessmentAutomationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomAssessmentAutomationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomAssessmentAutomationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates or updates a custom assessment automation for the provided subscription. Please note that providing an
// existing custom assessment automation will replace the existing record.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - customAssessmentAutomationName - Name of the Custom Assessment Automation.
//   - customAssessmentAutomationBody - Custom Assessment Automation body
//   - options - CustomAssessmentAutomationsClientCreateOptions contains the optional parameters for the CustomAssessmentAutomationsClient.Create
//     method.
func (client *CustomAssessmentAutomationsClient) Create(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, customAssessmentAutomationBody CustomAssessmentAutomationRequest, options *CustomAssessmentAutomationsClientCreateOptions) (CustomAssessmentAutomationsClientCreateResponse, error) {
	var err error
	const operationName = "CustomAssessmentAutomationsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, customAssessmentAutomationName, customAssessmentAutomationBody, options)
	if err != nil {
		return CustomAssessmentAutomationsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomAssessmentAutomationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CustomAssessmentAutomationsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *CustomAssessmentAutomationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, customAssessmentAutomationBody CustomAssessmentAutomationRequest, options *CustomAssessmentAutomationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations/{customAssessmentAutomationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customAssessmentAutomationName == "" {
		return nil, errors.New("parameter customAssessmentAutomationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customAssessmentAutomationName}", url.PathEscape(customAssessmentAutomationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, customAssessmentAutomationBody); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *CustomAssessmentAutomationsClient) createHandleResponse(resp *http.Response) (CustomAssessmentAutomationsClientCreateResponse, error) {
	result := CustomAssessmentAutomationsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomation); err != nil {
		return CustomAssessmentAutomationsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a custom assessment automation by name for a provided subscription
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - customAssessmentAutomationName - Name of the Custom Assessment Automation.
//   - options - CustomAssessmentAutomationsClientDeleteOptions contains the optional parameters for the CustomAssessmentAutomationsClient.Delete
//     method.
func (client *CustomAssessmentAutomationsClient) Delete(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsClientDeleteOptions) (CustomAssessmentAutomationsClientDeleteResponse, error) {
	var err error
	const operationName = "CustomAssessmentAutomationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, customAssessmentAutomationName, options)
	if err != nil {
		return CustomAssessmentAutomationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomAssessmentAutomationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CustomAssessmentAutomationsClientDeleteResponse{}, err
	}
	return CustomAssessmentAutomationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomAssessmentAutomationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations/{customAssessmentAutomationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customAssessmentAutomationName == "" {
		return nil, errors.New("parameter customAssessmentAutomationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customAssessmentAutomationName}", url.PathEscape(customAssessmentAutomationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a single custom assessment automation by name for the provided subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - customAssessmentAutomationName - Name of the Custom Assessment Automation.
//   - options - CustomAssessmentAutomationsClientGetOptions contains the optional parameters for the CustomAssessmentAutomationsClient.Get
//     method.
func (client *CustomAssessmentAutomationsClient) Get(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsClientGetOptions) (CustomAssessmentAutomationsClientGetResponse, error) {
	var err error
	const operationName = "CustomAssessmentAutomationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, customAssessmentAutomationName, options)
	if err != nil {
		return CustomAssessmentAutomationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomAssessmentAutomationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CustomAssessmentAutomationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CustomAssessmentAutomationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations/{customAssessmentAutomationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customAssessmentAutomationName == "" {
		return nil, errors.New("parameter customAssessmentAutomationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customAssessmentAutomationName}", url.PathEscape(customAssessmentAutomationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomAssessmentAutomationsClient) getHandleResponse(resp *http.Response) (CustomAssessmentAutomationsClientGetResponse, error) {
	result := CustomAssessmentAutomationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomation); err != nil {
		return CustomAssessmentAutomationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List custom assessment automations by provided subscription and resource group
//
// Generated from API version 2021-07-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - options - CustomAssessmentAutomationsClientListByResourceGroupOptions contains the optional parameters for the CustomAssessmentAutomationsClient.NewListByResourceGroupPager
//     method.
func (client *CustomAssessmentAutomationsClient) NewListByResourceGroupPager(resourceGroupName string, options *CustomAssessmentAutomationsClientListByResourceGroupOptions) *runtime.Pager[CustomAssessmentAutomationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomAssessmentAutomationsClientListByResourceGroupResponse]{
		More: func(page CustomAssessmentAutomationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomAssessmentAutomationsClientListByResourceGroupResponse) (CustomAssessmentAutomationsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomAssessmentAutomationsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CustomAssessmentAutomationsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomAssessmentAutomationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomAssessmentAutomationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomAssessmentAutomationsClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomAssessmentAutomationsClientListByResourceGroupResponse, error) {
	result := CustomAssessmentAutomationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomationsListResult); err != nil {
		return CustomAssessmentAutomationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List custom assessment automations by provided subscription
//
// Generated from API version 2021-07-01-preview
//   - options - CustomAssessmentAutomationsClientListBySubscriptionOptions contains the optional parameters for the CustomAssessmentAutomationsClient.NewListBySubscriptionPager
//     method.
func (client *CustomAssessmentAutomationsClient) NewListBySubscriptionPager(options *CustomAssessmentAutomationsClientListBySubscriptionOptions) *runtime.Pager[CustomAssessmentAutomationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomAssessmentAutomationsClientListBySubscriptionResponse]{
		More: func(page CustomAssessmentAutomationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomAssessmentAutomationsClientListBySubscriptionResponse) (CustomAssessmentAutomationsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomAssessmentAutomationsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CustomAssessmentAutomationsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomAssessmentAutomationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomAssessmentAutomationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/customAssessmentAutomations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomAssessmentAutomationsClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomAssessmentAutomationsClientListBySubscriptionResponse, error) {
	result := CustomAssessmentAutomationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomationsListResult); err != nil {
		return CustomAssessmentAutomationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
