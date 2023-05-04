//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// OutboundFirewallRulesClient contains the methods for the OutboundFirewallRules group.
// Don't use this type directly, use NewOutboundFirewallRulesClient() instead.
type OutboundFirewallRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOutboundFirewallRulesClient creates a new instance of OutboundFirewallRulesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOutboundFirewallRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OutboundFirewallRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".OutboundFirewallRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OutboundFirewallRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a outbound firewall rule with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - OutboundFirewallRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the OutboundFirewallRulesClient.BeginCreateOrUpdate
//     method.
func (client *OutboundFirewallRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, parameters OutboundFirewallRule, options *OutboundFirewallRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[OutboundFirewallRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, outboundRuleFqdn, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OutboundFirewallRulesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OutboundFirewallRulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a outbound firewall rule with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
func (client *OutboundFirewallRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, parameters OutboundFirewallRule, options *OutboundFirewallRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, outboundRuleFqdn, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OutboundFirewallRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, parameters OutboundFirewallRule, options *OutboundFirewallRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules/{outboundRuleFqdn}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if outboundRuleFqdn == "" {
		return nil, errors.New("parameter outboundRuleFqdn cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleFqdn}", url.PathEscape(outboundRuleFqdn))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a outbound firewall rule with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - OutboundFirewallRulesClientBeginDeleteOptions contains the optional parameters for the OutboundFirewallRulesClient.BeginDelete
//     method.
func (client *OutboundFirewallRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesClientBeginDeleteOptions) (*runtime.Poller[OutboundFirewallRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, outboundRuleFqdn, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[OutboundFirewallRulesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[OutboundFirewallRulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a outbound firewall rule with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
func (client *OutboundFirewallRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, outboundRuleFqdn, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OutboundFirewallRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules/{outboundRuleFqdn}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if outboundRuleFqdn == "" {
		return nil, errors.New("parameter outboundRuleFqdn cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleFqdn}", url.PathEscape(outboundRuleFqdn))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets an outbound firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - OutboundFirewallRulesClientGetOptions contains the optional parameters for the OutboundFirewallRulesClient.Get
//     method.
func (client *OutboundFirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesClientGetOptions) (OutboundFirewallRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, outboundRuleFqdn, options)
	if err != nil {
		return OutboundFirewallRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OutboundFirewallRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OutboundFirewallRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OutboundFirewallRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, outboundRuleFqdn string, options *OutboundFirewallRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules/{outboundRuleFqdn}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if outboundRuleFqdn == "" {
		return nil, errors.New("parameter outboundRuleFqdn cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundRuleFqdn}", url.PathEscape(outboundRuleFqdn))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OutboundFirewallRulesClient) getHandleResponse(resp *http.Response) (OutboundFirewallRulesClientGetResponse, error) {
	result := OutboundFirewallRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundFirewallRule); err != nil {
		return OutboundFirewallRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets all outbound firewall rules on a server.
//
// Generated from API version 2021-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - OutboundFirewallRulesClientListByServerOptions contains the optional parameters for the OutboundFirewallRulesClient.NewListByServerPager
//     method.
func (client *OutboundFirewallRulesClient) NewListByServerPager(resourceGroupName string, serverName string, options *OutboundFirewallRulesClientListByServerOptions) *runtime.Pager[OutboundFirewallRulesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[OutboundFirewallRulesClientListByServerResponse]{
		More: func(page OutboundFirewallRulesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OutboundFirewallRulesClientListByServerResponse) (OutboundFirewallRulesClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OutboundFirewallRulesClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OutboundFirewallRulesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OutboundFirewallRulesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *OutboundFirewallRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *OutboundFirewallRulesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/outboundFirewallRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *OutboundFirewallRulesClient) listByServerHandleResponse(resp *http.Response) (OutboundFirewallRulesClientListByServerResponse, error) {
	result := OutboundFirewallRulesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundFirewallRuleListResult); err != nil {
		return OutboundFirewallRulesClientListByServerResponse{}, err
	}
	return result, nil
}
