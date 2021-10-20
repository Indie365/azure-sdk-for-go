//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// CustomAssessmentAutomationsClient contains the methods for the CustomAssessmentAutomations group.
// Don't use this type directly, use NewCustomAssessmentAutomationsClient() instead.
type CustomAssessmentAutomationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCustomAssessmentAutomationsClient creates a new instance of CustomAssessmentAutomationsClient with the specified values.
func NewCustomAssessmentAutomationsClient(con *arm.Connection, subscriptionID string) *CustomAssessmentAutomationsClient {
	return &CustomAssessmentAutomationsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Creates or updates a custom assessment automation for the provided subscription. Please note that providing an existing custom assessment automation
// will replace the existing record.
// If the operation fails it returns the *CloudError error type.
func (client *CustomAssessmentAutomationsClient) Create(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, customAssessmentAutomationBody CustomAssessmentAutomation, options *CustomAssessmentAutomationsCreateOptions) (CustomAssessmentAutomationsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, customAssessmentAutomationName, customAssessmentAutomationBody, options)
	if err != nil {
		return CustomAssessmentAutomationsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomAssessmentAutomationsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CustomAssessmentAutomationsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *CustomAssessmentAutomationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, customAssessmentAutomationBody CustomAssessmentAutomation, options *CustomAssessmentAutomationsCreateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customAssessmentAutomationBody)
}

// createHandleResponse handles the Create response.
func (client *CustomAssessmentAutomationsClient) createHandleResponse(resp *http.Response) (CustomAssessmentAutomationsCreateResponse, error) {
	result := CustomAssessmentAutomationsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomation); err != nil {
		return CustomAssessmentAutomationsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *CustomAssessmentAutomationsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a custom assessment automation by name for a provided subscription
// If the operation fails it returns the *CloudError error type.
func (client *CustomAssessmentAutomationsClient) Delete(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsDeleteOptions) (CustomAssessmentAutomationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, customAssessmentAutomationName, options)
	if err != nil {
		return CustomAssessmentAutomationsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomAssessmentAutomationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CustomAssessmentAutomationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return CustomAssessmentAutomationsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomAssessmentAutomationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsDeleteOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CustomAssessmentAutomationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a single custom assessment automation by name for the provided subscription and resource group.
// If the operation fails it returns the *CloudError error type.
func (client *CustomAssessmentAutomationsClient) Get(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsGetOptions) (CustomAssessmentAutomationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, customAssessmentAutomationName, options)
	if err != nil {
		return CustomAssessmentAutomationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomAssessmentAutomationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomAssessmentAutomationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomAssessmentAutomationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, customAssessmentAutomationName string, options *CustomAssessmentAutomationsGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomAssessmentAutomationsClient) getHandleResponse(resp *http.Response) (CustomAssessmentAutomationsGetResponse, error) {
	result := CustomAssessmentAutomationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomation); err != nil {
		return CustomAssessmentAutomationsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CustomAssessmentAutomationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - List custom assessment automations by provided subscription and resource group
// If the operation fails it returns the *CloudError error type.
func (client *CustomAssessmentAutomationsClient) ListByResourceGroup(resourceGroupName string, options *CustomAssessmentAutomationsListByResourceGroupOptions) *CustomAssessmentAutomationsListByResourceGroupPager {
	return &CustomAssessmentAutomationsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CustomAssessmentAutomationsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomAssessmentAutomationsListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomAssessmentAutomationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomAssessmentAutomationsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customAssessmentAutomations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomAssessmentAutomationsClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomAssessmentAutomationsListByResourceGroupResponse, error) {
	result := CustomAssessmentAutomationsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomationsListResult); err != nil {
		return CustomAssessmentAutomationsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *CustomAssessmentAutomationsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - List custom assessment automations by provided subscription
// If the operation fails it returns the *CloudError error type.
func (client *CustomAssessmentAutomationsClient) ListBySubscription(options *CustomAssessmentAutomationsListBySubscriptionOptions) *CustomAssessmentAutomationsListBySubscriptionPager {
	return &CustomAssessmentAutomationsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomAssessmentAutomationsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomAssessmentAutomationsListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomAssessmentAutomationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomAssessmentAutomationsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/customAssessmentAutomations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomAssessmentAutomationsClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomAssessmentAutomationsListBySubscriptionResponse, error) {
	result := CustomAssessmentAutomationsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomAssessmentAutomationsListResult); err != nil {
		return CustomAssessmentAutomationsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *CustomAssessmentAutomationsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
