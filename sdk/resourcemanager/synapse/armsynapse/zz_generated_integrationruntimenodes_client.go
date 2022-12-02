//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IntegrationRuntimeNodesClient contains the methods for the IntegrationRuntimeNodes group.
// Don't use this type directly, use NewIntegrationRuntimeNodesClient() instead.
type IntegrationRuntimeNodesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationRuntimeNodesClient creates a new instance of IntegrationRuntimeNodesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationRuntimeNodesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationRuntimeNodesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationRuntimeNodesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Delete - Delete an integration runtime node
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// integrationRuntimeName - Integration runtime name
// nodeName - Integration runtime node name
// options - IntegrationRuntimeNodesClientDeleteOptions contains the optional parameters for the IntegrationRuntimeNodesClient.Delete
// method.
func (client *IntegrationRuntimeNodesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientDeleteOptions) (IntegrationRuntimeNodesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, nodeName, options)
	if err != nil {
		return IntegrationRuntimeNodesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeNodesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationRuntimeNodesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IntegrationRuntimeNodesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationRuntimeNodesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an integration runtime node
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// integrationRuntimeName - Integration runtime name
// nodeName - Integration runtime node name
// options - IntegrationRuntimeNodesClientGetOptions contains the optional parameters for the IntegrationRuntimeNodesClient.Get
// method.
func (client *IntegrationRuntimeNodesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientGetOptions) (IntegrationRuntimeNodesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, nodeName, options)
	if err != nil {
		return IntegrationRuntimeNodesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeNodesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeNodesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimeNodesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimeNodesClient) getHandleResponse(resp *http.Response) (IntegrationRuntimeNodesClientGetResponse, error) {
	result := IntegrationRuntimeNodesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SelfHostedIntegrationRuntimeNode); err != nil {
		return IntegrationRuntimeNodesClientGetResponse{}, err
	}
	return result, nil
}

// Update - Create an integration runtime node
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// integrationRuntimeName - Integration runtime name
// nodeName - Integration runtime node name
// updateIntegrationRuntimeNodeRequest - The parameters for updating an integration runtime node.
// options - IntegrationRuntimeNodesClientUpdateOptions contains the optional parameters for the IntegrationRuntimeNodesClient.Update
// method.
func (client *IntegrationRuntimeNodesClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest UpdateIntegrationRuntimeNodeRequest, options *IntegrationRuntimeNodesClientUpdateOptions) (IntegrationRuntimeNodesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, nodeName, updateIntegrationRuntimeNodeRequest, options)
	if err != nil {
		return IntegrationRuntimeNodesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeNodesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeNodesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *IntegrationRuntimeNodesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest UpdateIntegrationRuntimeNodeRequest, options *IntegrationRuntimeNodesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateIntegrationRuntimeNodeRequest)
}

// updateHandleResponse handles the Update response.
func (client *IntegrationRuntimeNodesClient) updateHandleResponse(resp *http.Response) (IntegrationRuntimeNodesClientUpdateResponse, error) {
	result := IntegrationRuntimeNodesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SelfHostedIntegrationRuntimeNode); err != nil {
		return IntegrationRuntimeNodesClientUpdateResponse{}, err
	}
	return result, nil
}