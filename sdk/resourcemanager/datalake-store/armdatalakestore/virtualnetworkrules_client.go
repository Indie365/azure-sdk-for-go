//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatalakestore

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

// VirtualNetworkRulesClient contains the methods for the VirtualNetworkRules group.
// Don't use this type directly, use NewVirtualNetworkRulesClient() instead.
type VirtualNetworkRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworkRulesClient creates a new instance of VirtualNetworkRulesClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworkRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualNetworkRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworkRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the specified virtual network rule. During update, the virtual network rule with the
// specified name will be replaced with this new virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-11-01
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Store account.
//   - virtualNetworkRuleName - The name of the virtual network rule to create or update.
//   - parameters - Parameters supplied to create or update the virtual network rule.
//   - options - VirtualNetworkRulesClientCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkRulesClient.CreateOrUpdate
//     method.
func (client *VirtualNetworkRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, parameters CreateOrUpdateVirtualNetworkRuleParameters, options *VirtualNetworkRulesClientCreateOrUpdateOptions) (VirtualNetworkRulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, parameters, options)
	if err != nil {
		return VirtualNetworkRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, parameters CreateOrUpdateVirtualNetworkRuleParameters, options *VirtualNetworkRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualNetworkRulesClient) createOrUpdateHandleResponse(resp *http.Response) (VirtualNetworkRulesClientCreateOrUpdateResponse, error) {
	result := VirtualNetworkRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified virtual network rule from the specified Data Lake Store account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-11-01
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Store account.
//   - virtualNetworkRuleName - The name of the virtual network rule to delete.
//   - options - VirtualNetworkRulesClientDeleteOptions contains the optional parameters for the VirtualNetworkRulesClient.Delete
//     method.
func (client *VirtualNetworkRulesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientDeleteOptions) (VirtualNetworkRulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return VirtualNetworkRulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return VirtualNetworkRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified Data Lake Store virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-11-01
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Store account.
//   - virtualNetworkRuleName - The name of the virtual network rule to retrieve.
//   - options - VirtualNetworkRulesClientGetOptions contains the optional parameters for the VirtualNetworkRulesClient.Get method.
func (client *VirtualNetworkRulesClient) Get(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientGetOptions) (VirtualNetworkRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkRulesClient) getHandleResponse(resp *http.Response) (VirtualNetworkRulesClientGetResponse, error) {
	result := VirtualNetworkRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - Lists the Data Lake Store virtual network rules within the specified Data Lake Store account.
//
// Generated from API version 2016-11-01
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Store account.
//   - options - VirtualNetworkRulesClientListByAccountOptions contains the optional parameters for the VirtualNetworkRulesClient.NewListByAccountPager
//     method.
func (client *VirtualNetworkRulesClient) NewListByAccountPager(resourceGroupName string, accountName string, options *VirtualNetworkRulesClientListByAccountOptions) *runtime.Pager[VirtualNetworkRulesClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkRulesClientListByAccountResponse]{
		More: func(page VirtualNetworkRulesClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkRulesClientListByAccountResponse) (VirtualNetworkRulesClientListByAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworkRulesClientListByAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworkRulesClientListByAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworkRulesClientListByAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAccountHandleResponse(resp)
		},
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *VirtualNetworkRulesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *VirtualNetworkRulesClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *VirtualNetworkRulesClient) listByAccountHandleResponse(resp *http.Response) (VirtualNetworkRulesClientListByAccountResponse, error) {
	result := VirtualNetworkRulesClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRuleListResult); err != nil {
		return VirtualNetworkRulesClientListByAccountResponse{}, err
	}
	return result, nil
}

// Update - Updates the specified virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-11-01
//   - resourceGroupName - The name of the Azure resource group.
//   - accountName - The name of the Data Lake Store account.
//   - virtualNetworkRuleName - The name of the virtual network rule to update.
//   - options - VirtualNetworkRulesClientUpdateOptions contains the optional parameters for the VirtualNetworkRulesClient.Update
//     method.
func (client *VirtualNetworkRulesClient) Update(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientUpdateOptions) (VirtualNetworkRulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VirtualNetworkRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeStore/accounts/{accountName}/virtualNetworkRules/{virtualNetworkRuleName}"
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
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *VirtualNetworkRulesClient) updateHandleResponse(resp *http.Response) (VirtualNetworkRulesClientUpdateResponse, error) {
	result := VirtualNetworkRulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesClientUpdateResponse{}, err
	}
	return result, nil
}
