//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrelay

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

// WCFRelaysClient contains the methods for the WCFRelays group.
// Don't use this type directly, use NewWCFRelaysClient() instead.
type WCFRelaysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWCFRelaysClient creates a new instance of WCFRelaysClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWCFRelaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WCFRelaysClient, error) {
	cl, err := arm.NewClient(moduleName+".WCFRelaysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WCFRelaysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a WCF relay. This operation is idempotent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - parameters - Parameters supplied to create a WCF relay.
//   - options - WCFRelaysClientCreateOrUpdateOptions contains the optional parameters for the WCFRelaysClient.CreateOrUpdate
//     method.
func (client *WCFRelaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, parameters WcfRelay, options *WCFRelaysClientCreateOrUpdateOptions) (WCFRelaysClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, relayName, parameters, options)
	if err != nil {
		return WCFRelaysClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WCFRelaysClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WCFRelaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, parameters WcfRelay, options *WCFRelaysClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WCFRelaysClient) createOrUpdateHandleResponse(resp *http.Response) (WCFRelaysClientCreateOrUpdateResponse, error) {
	result := WCFRelaysClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WcfRelay); err != nil {
		return WCFRelaysClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAuthorizationRule - Creates or updates an authorization rule for a WCF relay.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - authorizationRuleName - The authorization rule name.
//   - parameters - The authorization rule parameters.
//   - options - WCFRelaysClientCreateOrUpdateAuthorizationRuleOptions contains the optional parameters for the WCFRelaysClient.CreateOrUpdateAuthorizationRule
//     method.
func (client *WCFRelaysClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, parameters AuthorizationRule, options *WCFRelaysClientCreateOrUpdateAuthorizationRuleOptions) (WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse, error) {
	req, err := client.createOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, relayName, authorizationRuleName, parameters, options)
	if err != nil {
		return WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateAuthorizationRuleHandleResponse(resp)
}

// createOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *WCFRelaysClient) createOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, parameters AuthorizationRule, options *WCFRelaysClientCreateOrUpdateAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *WCFRelaysClient) createOrUpdateAuthorizationRuleHandleResponse(resp *http.Response) (WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse, error) {
	result := WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRule); err != nil {
		return WCFRelaysClientCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a WCF relay.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - options - WCFRelaysClientDeleteOptions contains the optional parameters for the WCFRelaysClient.Delete method.
func (client *WCFRelaysClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, options *WCFRelaysClientDeleteOptions) (WCFRelaysClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, relayName, options)
	if err != nil {
		return WCFRelaysClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WCFRelaysClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WCFRelaysClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WCFRelaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, options *WCFRelaysClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAuthorizationRule - Deletes a WCF relay authorization rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - authorizationRuleName - The authorization rule name.
//   - options - WCFRelaysClientDeleteAuthorizationRuleOptions contains the optional parameters for the WCFRelaysClient.DeleteAuthorizationRule
//     method.
func (client *WCFRelaysClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, options *WCFRelaysClientDeleteAuthorizationRuleOptions) (WCFRelaysClientDeleteAuthorizationRuleResponse, error) {
	req, err := client.deleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, relayName, authorizationRuleName, options)
	if err != nil {
		return WCFRelaysClientDeleteAuthorizationRuleResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientDeleteAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WCFRelaysClientDeleteAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return WCFRelaysClientDeleteAuthorizationRuleResponse{}, nil
}

// deleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *WCFRelaysClient) deleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, options *WCFRelaysClientDeleteAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the description for the specified WCF relay.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - options - WCFRelaysClientGetOptions contains the optional parameters for the WCFRelaysClient.Get method.
func (client *WCFRelaysClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, options *WCFRelaysClientGetOptions) (WCFRelaysClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, relayName, options)
	if err != nil {
		return WCFRelaysClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WCFRelaysClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WCFRelaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, options *WCFRelaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WCFRelaysClient) getHandleResponse(resp *http.Response) (WCFRelaysClientGetResponse, error) {
	result := WCFRelaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WcfRelay); err != nil {
		return WCFRelaysClientGetResponse{}, err
	}
	return result, nil
}

// GetAuthorizationRule - Get authorizationRule for a WCF relay by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - authorizationRuleName - The authorization rule name.
//   - options - WCFRelaysClientGetAuthorizationRuleOptions contains the optional parameters for the WCFRelaysClient.GetAuthorizationRule
//     method.
func (client *WCFRelaysClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, options *WCFRelaysClientGetAuthorizationRuleOptions) (WCFRelaysClientGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, relayName, authorizationRuleName, options)
	if err != nil {
		return WCFRelaysClientGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientGetAuthorizationRuleResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WCFRelaysClientGetAuthorizationRuleResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *WCFRelaysClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, options *WCFRelaysClientGetAuthorizationRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *WCFRelaysClient) getAuthorizationRuleHandleResponse(resp *http.Response) (WCFRelaysClientGetAuthorizationRuleResponse, error) {
	result := WCFRelaysClientGetAuthorizationRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRule); err != nil {
		return WCFRelaysClientGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// NewListAuthorizationRulesPager - Authorization rules for a WCF relay.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - options - WCFRelaysClientListAuthorizationRulesOptions contains the optional parameters for the WCFRelaysClient.NewListAuthorizationRulesPager
//     method.
func (client *WCFRelaysClient) NewListAuthorizationRulesPager(resourceGroupName string, namespaceName string, relayName string, options *WCFRelaysClientListAuthorizationRulesOptions) *runtime.Pager[WCFRelaysClientListAuthorizationRulesResponse] {
	return runtime.NewPager(runtime.PagingHandler[WCFRelaysClientListAuthorizationRulesResponse]{
		More: func(page WCFRelaysClientListAuthorizationRulesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WCFRelaysClientListAuthorizationRulesResponse) (WCFRelaysClientListAuthorizationRulesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, relayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WCFRelaysClientListAuthorizationRulesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WCFRelaysClientListAuthorizationRulesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WCFRelaysClientListAuthorizationRulesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAuthorizationRulesHandleResponse(resp)
		},
	})
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *WCFRelaysClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, options *WCFRelaysClientListAuthorizationRulesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *WCFRelaysClient) listAuthorizationRulesHandleResponse(resp *http.Response) (WCFRelaysClientListAuthorizationRulesResponse, error) {
	result := WCFRelaysClientListAuthorizationRulesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AuthorizationRuleListResult); err != nil {
		return WCFRelaysClientListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// NewListByNamespacePager - Lists the WCF relays within the namespace.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - options - WCFRelaysClientListByNamespaceOptions contains the optional parameters for the WCFRelaysClient.NewListByNamespacePager
//     method.
func (client *WCFRelaysClient) NewListByNamespacePager(resourceGroupName string, namespaceName string, options *WCFRelaysClientListByNamespaceOptions) *runtime.Pager[WCFRelaysClientListByNamespaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WCFRelaysClientListByNamespaceResponse]{
		More: func(page WCFRelaysClientListByNamespaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WCFRelaysClientListByNamespaceResponse) (WCFRelaysClientListByNamespaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WCFRelaysClientListByNamespaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WCFRelaysClientListByNamespaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WCFRelaysClientListByNamespaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByNamespaceHandleResponse(resp)
		},
	})
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *WCFRelaysClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *WCFRelaysClientListByNamespaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *WCFRelaysClient) listByNamespaceHandleResponse(resp *http.Response) (WCFRelaysClientListByNamespaceResponse, error) {
	result := WCFRelaysClientListByNamespaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WcfRelaysListResult); err != nil {
		return WCFRelaysClientListByNamespaceResponse{}, err
	}
	return result, nil
}

// ListKeys - Primary and secondary connection strings to the WCF relay.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - authorizationRuleName - The authorization rule name.
//   - options - WCFRelaysClientListKeysOptions contains the optional parameters for the WCFRelaysClient.ListKeys method.
func (client *WCFRelaysClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, options *WCFRelaysClientListKeysOptions) (WCFRelaysClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, relayName, authorizationRuleName, options)
	if err != nil {
		return WCFRelaysClientListKeysResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WCFRelaysClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *WCFRelaysClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, options *WCFRelaysClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *WCFRelaysClient) listKeysHandleResponse(resp *http.Response) (WCFRelaysClientListKeysResponse, error) {
	result := WCFRelaysClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return WCFRelaysClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKeys - Regenerates the primary or secondary connection strings to the WCF relay.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - namespaceName - The namespace name
//   - relayName - The relay name.
//   - authorizationRuleName - The authorization rule name.
//   - parameters - Parameters supplied to regenerate authorization rule.
//   - options - WCFRelaysClientRegenerateKeysOptions contains the optional parameters for the WCFRelaysClient.RegenerateKeys
//     method.
func (client *WCFRelaysClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *WCFRelaysClientRegenerateKeysOptions) (WCFRelaysClientRegenerateKeysResponse, error) {
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, relayName, authorizationRuleName, parameters, options)
	if err != nil {
		return WCFRelaysClientRegenerateKeysResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WCFRelaysClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WCFRelaysClientRegenerateKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeysHandleResponse(resp)
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *WCFRelaysClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, relayName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *WCFRelaysClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Relay/namespaces/{namespaceName}/wcfRelays/{relayName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if relayName == "" {
		return nil, errors.New("parameter relayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relayName}", url.PathEscape(relayName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *WCFRelaysClient) regenerateKeysHandleResponse(resp *http.Response) (WCFRelaysClientRegenerateKeysResponse, error) {
	result := WCFRelaysClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return WCFRelaysClientRegenerateKeysResponse{}, err
	}
	return result, nil
}
