//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

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

// ModuleClient contains the methods for the Module group.
// Don't use this type directly, use NewModuleClient() instead.
type ModuleClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewModuleClient creates a new instance of ModuleClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewModuleClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ModuleClient, error) {
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
	client := &ModuleClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update the module identified by module name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// moduleName - The name of module.
// parameters - The create or update parameters for module.
// options - ModuleClientCreateOrUpdateOptions contains the optional parameters for the ModuleClient.CreateOrUpdate method.
func (client *ModuleClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleCreateOrUpdateParameters, options *ModuleClientCreateOrUpdateOptions) (ModuleClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, parameters, options)
	if err != nil {
		return ModuleClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ModuleClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ModuleClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ModuleClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleCreateOrUpdateParameters, options *ModuleClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ModuleClient) createOrUpdateHandleResponse(resp *http.Response) (ModuleClientCreateOrUpdateResponse, error) {
	result := ModuleClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return ModuleClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the module by name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// moduleName - The module name.
// options - ModuleClientDeleteOptions contains the optional parameters for the ModuleClient.Delete method.
func (client *ModuleClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, options *ModuleClientDeleteOptions) (ModuleClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, options)
	if err != nil {
		return ModuleClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ModuleClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ModuleClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ModuleClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ModuleClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, options *ModuleClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the module identified by module name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// moduleName - The module name.
// options - ModuleClientGetOptions contains the optional parameters for the ModuleClient.Get method.
func (client *ModuleClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, options *ModuleClientGetOptions) (ModuleClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, options)
	if err != nil {
		return ModuleClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ModuleClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ModuleClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ModuleClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, options *ModuleClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ModuleClient) getHandleResponse(resp *http.Response) (ModuleClientGetResponse, error) {
	result := ModuleClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return ModuleClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of modules.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - ModuleClientListByAutomationAccountOptions contains the optional parameters for the ModuleClient.ListByAutomationAccount
// method.
func (client *ModuleClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *ModuleClientListByAutomationAccountOptions) *runtime.Pager[ModuleClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[ModuleClientListByAutomationAccountResponse]{
		More: func(page ModuleClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ModuleClientListByAutomationAccountResponse) (ModuleClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ModuleClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ModuleClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ModuleClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *ModuleClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *ModuleClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *ModuleClient) listByAutomationAccountHandleResponse(resp *http.Response) (ModuleClientListByAutomationAccountResponse, error) {
	result := ModuleClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModuleListResult); err != nil {
		return ModuleClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update the module identified by module name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// moduleName - The name of module.
// parameters - The update parameters for module.
// options - ModuleClientUpdateOptions contains the optional parameters for the ModuleClient.Update method.
func (client *ModuleClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleUpdateParameters, options *ModuleClientUpdateOptions) (ModuleClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, parameters, options)
	if err != nil {
		return ModuleClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ModuleClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ModuleClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ModuleClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, parameters ModuleUpdateParameters, options *ModuleClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ModuleClient) updateHandleResponse(resp *http.Response) (ModuleClientUpdateResponse, error) {
	result := ModuleClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return ModuleClientUpdateResponse{}, err
	}
	return result, nil
}