//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// Python2PackageClient contains the methods for the Python2Package group.
// Don't use this type directly, use NewPython2PackageClient() instead.
type Python2PackageClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewPython2PackageClient creates a new instance of Python2PackageClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPython2PackageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Python2PackageClient, error) {
	cl, err := arm.NewClient(moduleName+".Python2PackageClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Python2PackageClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update the python 2 package identified by package name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The name of python package.
//   - parameters - The create or update parameters for python package.
//   - options - Python2PackageClientCreateOrUpdateOptions contains the optional parameters for the Python2PackageClient.CreateOrUpdate
//     method.
func (client *Python2PackageClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageCreateParameters, options *Python2PackageClientCreateOrUpdateOptions) (Python2PackageClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, parameters, options)
	if err != nil {
		return Python2PackageClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python2PackageClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return Python2PackageClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *Python2PackageClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageCreateParameters, options *Python2PackageClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *Python2PackageClient) createOrUpdateHandleResponse(resp *http.Response) (Python2PackageClientCreateOrUpdateResponse, error) {
	result := Python2PackageClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return Python2PackageClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the python 2 package by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The python package name.
//   - options - Python2PackageClientDeleteOptions contains the optional parameters for the Python2PackageClient.Delete method.
func (client *Python2PackageClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python2PackageClientDeleteOptions) (Python2PackageClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, options)
	if err != nil {
		return Python2PackageClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python2PackageClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Python2PackageClientDeleteResponse{}, err
	}
	return Python2PackageClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *Python2PackageClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python2PackageClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the python 2 package identified by package name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The python package name.
//   - options - Python2PackageClientGetOptions contains the optional parameters for the Python2PackageClient.Get method.
func (client *Python2PackageClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python2PackageClientGetOptions) (Python2PackageClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, options)
	if err != nil {
		return Python2PackageClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python2PackageClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Python2PackageClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *Python2PackageClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, options *Python2PackageClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *Python2PackageClient) getHandleResponse(resp *http.Response) (Python2PackageClientGetResponse, error) {
	result := Python2PackageClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return Python2PackageClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of python 2 packages.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - Python2PackageClientListByAutomationAccountOptions contains the optional parameters for the Python2PackageClient.NewListByAutomationAccountPager
//     method.
func (client *Python2PackageClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *Python2PackageClientListByAutomationAccountOptions) (*runtime.Pager[Python2PackageClientListByAutomationAccountResponse]) {
	return runtime.NewPager(runtime.PagingHandler[Python2PackageClientListByAutomationAccountResponse]{
		More: func(page Python2PackageClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *Python2PackageClientListByAutomationAccountResponse) (Python2PackageClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return Python2PackageClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return Python2PackageClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return Python2PackageClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *Python2PackageClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *Python2PackageClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *Python2PackageClient) listByAutomationAccountHandleResponse(resp *http.Response) (Python2PackageClientListByAutomationAccountResponse, error) {
	result := Python2PackageClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ModuleListResult); err != nil {
		return Python2PackageClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update the python 2 package identified by package name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - packageName - The name of python package.
//   - parameters - The update parameters for python package.
//   - options - Python2PackageClientUpdateOptions contains the optional parameters for the Python2PackageClient.Update method.
func (client *Python2PackageClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageUpdateParameters, options *Python2PackageClientUpdateOptions) (Python2PackageClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, packageName, parameters, options)
	if err != nil {
		return Python2PackageClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Python2PackageClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Python2PackageClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *Python2PackageClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, packageName string, parameters PythonPackageUpdateParameters, options *Python2PackageClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/python2Packages/{packageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *Python2PackageClient) updateHandleResponse(resp *http.Response) (Python2PackageClientUpdateResponse, error) {
	result := Python2PackageClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Module); err != nil {
		return Python2PackageClientUpdateResponse{}, err
	}
	return result, nil
}

