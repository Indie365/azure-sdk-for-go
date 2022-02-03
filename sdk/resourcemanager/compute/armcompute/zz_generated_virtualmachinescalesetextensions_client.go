//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// VirtualMachineScaleSetExtensionsClient contains the methods for the VirtualMachineScaleSetExtensions group.
// Don't use this type directly, use NewVirtualMachineScaleSetExtensionsClient() instead.
type VirtualMachineScaleSetExtensionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineScaleSetExtensionsClient creates a new instance of VirtualMachineScaleSetExtensionsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineScaleSetExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineScaleSetExtensionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VirtualMachineScaleSetExtensionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - The operation to create or update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set where the extension should be create or updated.
// vmssExtensionName - The name of the VM scale set extension.
// extensionParameters - Parameters supplied to the Create VM scale set Extension operation.
// options - VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.BeginCreateOrUpdate
// method.
func (client *VirtualMachineScaleSetExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions) (VirtualMachineScaleSetExtensionsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetExtensionsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetExtensionsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetExtensionsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineScaleSetExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtension, options *VirtualMachineScaleSetExtensionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}

// BeginDelete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set where the extension should be deleted.
// vmssExtensionName - The name of the VM scale set extension.
// options - VirtualMachineScaleSetExtensionsClientBeginDeleteOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.BeginDelete
// method.
func (client *VirtualMachineScaleSetExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientBeginDeleteOptions) (VirtualMachineScaleSetExtensionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientDeletePollerResponse{}, err
	}
	result := VirtualMachineScaleSetExtensionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetExtensionsClient.Delete", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientDeletePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetExtensionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
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
func (client *VirtualMachineScaleSetExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - The operation to get the extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set containing the extension.
// vmssExtensionName - The name of the VM scale set extension.
// options - VirtualMachineScaleSetExtensionsClientGetOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.Get
// method.
func (client *VirtualMachineScaleSetExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientGetOptions) (VirtualMachineScaleSetExtensionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, options *VirtualMachineScaleSetExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetExtensionsClient) getHandleResponse(resp *http.Response) (VirtualMachineScaleSetExtensionsClientGetResponse, error) {
	result := VirtualMachineScaleSetExtensionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineScaleSetExtension); err != nil {
		return VirtualMachineScaleSetExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of all extensions in a VM scale set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set containing the extension.
// options - VirtualMachineScaleSetExtensionsClientListOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.List
// method.
func (client *VirtualMachineScaleSetExtensionsClient) List(resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsClientListOptions) *VirtualMachineScaleSetExtensionsClientListPager {
	return &VirtualMachineScaleSetExtensionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachineScaleSetExtensionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineScaleSetExtensionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetExtensionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetExtensionsClient) listHandleResponse(resp *http.Response) (VirtualMachineScaleSetExtensionsClientListResponse, error) {
	result := VirtualMachineScaleSetExtensionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineScaleSetExtensionListResult); err != nil {
		return VirtualMachineScaleSetExtensionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set where the extension should be updated.
// vmssExtensionName - The name of the VM scale set extension.
// extensionParameters - Parameters supplied to the Update VM scale set Extension operation.
// options - VirtualMachineScaleSetExtensionsClientBeginUpdateOptions contains the optional parameters for the VirtualMachineScaleSetExtensionsClient.BeginUpdate
// method.
func (client *VirtualMachineScaleSetExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsClientBeginUpdateOptions) (VirtualMachineScaleSetExtensionsClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientUpdatePollerResponse{}, err
	}
	result := VirtualMachineScaleSetExtensionsClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineScaleSetExtensionsClient.Update", "", resp, client.pl)
	if err != nil {
		return VirtualMachineScaleSetExtensionsClientUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineScaleSetExtensionsClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The operation to update an extension.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualMachineScaleSetExtensionsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, vmssExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, vmssExtensionName string, extensionParameters VirtualMachineScaleSetExtensionUpdate, options *VirtualMachineScaleSetExtensionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensions/{vmssExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if vmssExtensionName == "" {
		return nil, errors.New("parameter vmssExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssExtensionName}", url.PathEscape(vmssExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}
