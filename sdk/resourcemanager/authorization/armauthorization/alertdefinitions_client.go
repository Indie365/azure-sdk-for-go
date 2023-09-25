//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// AlertDefinitionsClient contains the methods for the AlertDefinitions group.
// Don't use this type directly, use NewAlertDefinitionsClient() instead.
type AlertDefinitionsClient struct {
	internal *arm.Client
}

// NewAlertDefinitionsClient creates a new instance of AlertDefinitionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName+".AlertDefinitionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AlertDefinitionsClient{
	internal: cl,
	}
	return client, nil
}

// Get - Get the specified alert definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - scope - The scope of the alert definition. The scope can be any REST resource instance. For example, use '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/'
//     for a subscription,
//     '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource
//     group, and
//     '/providers/Microsoft.Subscription/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}'
//     for a resource.
//   - alertDefinitionID - The name of the alert definition to get.
//   - options - AlertDefinitionsClientGetOptions contains the optional parameters for the AlertDefinitionsClient.Get method.
func (client *AlertDefinitionsClient) Get(ctx context.Context, scope string, alertDefinitionID string, options *AlertDefinitionsClientGetOptions) (AlertDefinitionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, scope, alertDefinitionID, options)
	if err != nil {
		return AlertDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AlertDefinitionsClient) getCreateRequest(ctx context.Context, scope string, alertDefinitionID string, options *AlertDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementAlertDefinitions/{alertDefinitionId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	urlPath = strings.ReplaceAll(urlPath, "{alertDefinitionId}", alertDefinitionID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertDefinitionsClient) getHandleResponse(resp *http.Response) (AlertDefinitionsClientGetResponse, error) {
	result := AlertDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertDefinition); err != nil {
		return AlertDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets alert definitions for a resource scope.
//
// Generated from API version 2022-08-01-preview
//   - scope - The scope of the alert definition.
//   - options - AlertDefinitionsClientListForScopeOptions contains the optional parameters for the AlertDefinitionsClient.NewListForScopePager
//     method.
func (client *AlertDefinitionsClient) NewListForScopePager(scope string, options *AlertDefinitionsClientListForScopeOptions) (*runtime.Pager[AlertDefinitionsClientListForScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[AlertDefinitionsClientListForScopeResponse]{
		More: func(page AlertDefinitionsClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AlertDefinitionsClientListForScopeResponse) (AlertDefinitionsClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AlertDefinitionsClientListForScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AlertDefinitionsClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AlertDefinitionsClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *AlertDefinitionsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *AlertDefinitionsClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementAlertDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *AlertDefinitionsClient) listForScopeHandleResponse(resp *http.Response) (AlertDefinitionsClientListForScopeResponse, error) {
	result := AlertDefinitionsClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertDefinitionListResult); err != nil {
		return AlertDefinitionsClientListForScopeResponse{}, err
	}
	return result, nil
}

