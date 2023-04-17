//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlogz

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

// SubAccountTagRulesClient contains the methods for the SubAccountTagRules group.
// Don't use this type directly, use NewSubAccountTagRulesClient() instead.
type SubAccountTagRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSubAccountTagRulesClient creates a new instance of SubAccountTagRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubAccountTagRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SubAccountTagRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".SubAccountTagRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubAccountTagRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a tag rule set for a given sub account resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - subAccountName - Sub Account resource name
//   - options - SubAccountTagRulesClientCreateOrUpdateOptions contains the optional parameters for the SubAccountTagRulesClient.CreateOrUpdate
//     method.
func (client *SubAccountTagRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, subAccountName string, ruleSetName string, options *SubAccountTagRulesClientCreateOrUpdateOptions) (SubAccountTagRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, subAccountName, ruleSetName, options)
	if err != nil {
		return SubAccountTagRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubAccountTagRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubAccountTagRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SubAccountTagRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, subAccountName string, ruleSetName string, options *SubAccountTagRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/accounts/{subAccountName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if subAccountName == "" {
		return nil, errors.New("parameter subAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subAccountName}", url.PathEscape(subAccountName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SubAccountTagRulesClient) createOrUpdateHandleResponse(resp *http.Response) (SubAccountTagRulesClientCreateOrUpdateResponse, error) {
	result := SubAccountTagRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return SubAccountTagRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - subAccountName - Sub Account resource name
//   - options - SubAccountTagRulesClientDeleteOptions contains the optional parameters for the SubAccountTagRulesClient.Delete
//     method.
func (client *SubAccountTagRulesClient) Delete(ctx context.Context, resourceGroupName string, monitorName string, subAccountName string, ruleSetName string, options *SubAccountTagRulesClientDeleteOptions) (SubAccountTagRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, subAccountName, ruleSetName, options)
	if err != nil {
		return SubAccountTagRulesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubAccountTagRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return SubAccountTagRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *SubAccountTagRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, subAccountName string, ruleSetName string, options *SubAccountTagRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/accounts/{subAccountName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if subAccountName == "" {
		return nil, errors.New("parameter subAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subAccountName}", url.PathEscape(subAccountName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *SubAccountTagRulesClient) deleteHandleResponse(resp *http.Response) (SubAccountTagRulesClientDeleteResponse, error) {
	result := SubAccountTagRulesClientDeleteResponse{}
	if val := resp.Header.Get("location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// Get - Get a tag rule set for a given monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - subAccountName - Sub Account resource name
//   - options - SubAccountTagRulesClientGetOptions contains the optional parameters for the SubAccountTagRulesClient.Get method.
func (client *SubAccountTagRulesClient) Get(ctx context.Context, resourceGroupName string, monitorName string, subAccountName string, ruleSetName string, options *SubAccountTagRulesClientGetOptions) (SubAccountTagRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, subAccountName, ruleSetName, options)
	if err != nil {
		return SubAccountTagRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubAccountTagRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubAccountTagRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubAccountTagRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, subAccountName string, ruleSetName string, options *SubAccountTagRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/accounts/{subAccountName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if subAccountName == "" {
		return nil, errors.New("parameter subAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subAccountName}", url.PathEscape(subAccountName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubAccountTagRulesClient) getHandleResponse(resp *http.Response) (SubAccountTagRulesClientGetResponse, error) {
	result := SubAccountTagRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return SubAccountTagRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the tag rules for a given sub account resource.
//
// Generated from API version 2020-10-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - subAccountName - Sub Account resource name
//   - options - SubAccountTagRulesClientListOptions contains the optional parameters for the SubAccountTagRulesClient.NewListPager
//     method.
func (client *SubAccountTagRulesClient) NewListPager(resourceGroupName string, monitorName string, subAccountName string, options *SubAccountTagRulesClientListOptions) *runtime.Pager[SubAccountTagRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubAccountTagRulesClientListResponse]{
		More: func(page SubAccountTagRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubAccountTagRulesClientListResponse) (SubAccountTagRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, monitorName, subAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SubAccountTagRulesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SubAccountTagRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SubAccountTagRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SubAccountTagRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, subAccountName string, options *SubAccountTagRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/accounts/{subAccountName}/tagRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if subAccountName == "" {
		return nil, errors.New("parameter subAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subAccountName}", url.PathEscape(subAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubAccountTagRulesClient) listHandleResponse(resp *http.Response) (SubAccountTagRulesClientListResponse, error) {
	result := SubAccountTagRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRulesListResponse); err != nil {
		return SubAccountTagRulesClientListResponse{}, err
	}
	return result, nil
}
