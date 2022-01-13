//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// VirtualNetworkRulesClient contains the methods for the VirtualNetworkRules group.
// Don't use this type directly, use NewVirtualNetworkRulesClient() instead.
type VirtualNetworkRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualNetworkRulesClient creates a new instance of VirtualNetworkRulesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualNetworkRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualNetworkRulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VirtualNetworkRulesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an existing virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// virtualNetworkRuleName - The name of the virtual network rule.
// parameters - The requested virtual Network Rule Resource state.
// options - VirtualNetworkRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkRulesClient.BeginCreateOrUpdate
// method.
func (client *VirtualNetworkRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesClientBeginCreateOrUpdateOptions) (VirtualNetworkRulesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, virtualNetworkRuleName, parameters, options)
	if err != nil {
		return VirtualNetworkRulesClientCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualNetworkRulesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkRulesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return VirtualNetworkRulesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkRulesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an existing virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, parameters, options)
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
func (client *VirtualNetworkRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the virtual network rule with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// virtualNetworkRuleName - The name of the virtual network rule.
// options - VirtualNetworkRulesClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkRulesClient.BeginDelete
// method.
func (client *VirtualNetworkRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientBeginDeleteOptions) (VirtualNetworkRulesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesClientDeletePollerResponse{}, err
	}
	result := VirtualNetworkRulesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualNetworkRulesClient.Delete", "", resp, client.pl)
	if err != nil {
		return VirtualNetworkRulesClientDeletePollerResponse{}, err
	}
	result.Poller = &VirtualNetworkRulesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the virtual network rule with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualNetworkRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
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
func (client *VirtualNetworkRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// virtualNetworkRuleName - The name of the virtual network rule.
// options - VirtualNetworkRulesClientGetOptions contains the optional parameters for the VirtualNetworkRulesClient.Get method.
func (client *VirtualNetworkRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientGetOptions) (VirtualNetworkRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkRulesClient) getHandleResponse(resp *http.Response) (VirtualNetworkRulesClientGetResponse, error) {
	result := VirtualNetworkRulesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Gets a list of virtual network rules in a server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - VirtualNetworkRulesClientListByServerOptions contains the optional parameters for the VirtualNetworkRulesClient.ListByServer
// method.
func (client *VirtualNetworkRulesClient) ListByServer(resourceGroupName string, serverName string, options *VirtualNetworkRulesClientListByServerOptions) *VirtualNetworkRulesClientListByServerPager {
	return &VirtualNetworkRulesClientListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp VirtualNetworkRulesClientListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkRuleListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *VirtualNetworkRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *VirtualNetworkRulesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/virtualNetworkRules"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *VirtualNetworkRulesClient) listByServerHandleResponse(resp *http.Response) (VirtualNetworkRulesClientListByServerResponse, error) {
	result := VirtualNetworkRulesClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRuleListResult); err != nil {
		return VirtualNetworkRulesClientListByServerResponse{}, err
	}
	return result, nil
}
