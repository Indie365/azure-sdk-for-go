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
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineRunCommandsClient contains the methods for the VirtualMachineRunCommands group.
// Don't use this type directly, use NewVirtualMachineRunCommandsClient() instead.
type VirtualMachineRunCommandsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineRunCommandsClient creates a new instance of VirtualMachineRunCommandsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineRunCommandsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineRunCommandsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &VirtualMachineRunCommandsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - The operation to create or update the run command.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// vmName - The name of the virtual machine where the run command should be created or updated.
// runCommandName - The name of the virtual machine run command.
// runCommand - Parameters supplied to the Create Virtual Machine RunCommand operation.
// options - VirtualMachineRunCommandsBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineRunCommandsClient.BeginCreateOrUpdate
// method.
func (client *VirtualMachineRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineRunCommandsBeginCreateOrUpdateOptions) (VirtualMachineRunCommandsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
	if err != nil {
		return VirtualMachineRunCommandsCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualMachineRunCommandsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineRunCommandsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualMachineRunCommandsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineRunCommandsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - The operation to create or update the run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineRunCommandsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineRunCommandsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineRunCommandsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineRunCommandsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, runCommand)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineRunCommandsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - The operation to delete the run command.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// vmName - The name of the virtual machine where the run command should be deleted.
// runCommandName - The name of the virtual machine run command.
// options - VirtualMachineRunCommandsBeginDeleteOptions contains the optional parameters for the VirtualMachineRunCommandsClient.BeginDelete
// method.
func (client *VirtualMachineRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsBeginDeleteOptions) (VirtualMachineRunCommandsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vmName, runCommandName, options)
	if err != nil {
		return VirtualMachineRunCommandsDeletePollerResponse{}, err
	}
	result := VirtualMachineRunCommandsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineRunCommandsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualMachineRunCommandsDeletePollerResponse{}, err
	}
	result.Poller = &VirtualMachineRunCommandsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The operation to delete the run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineRunCommandsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmName, runCommandName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachineRunCommandsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualMachineRunCommandsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets specific run command for a subscription in a location.
// If the operation fails it returns a generic error.
// location - The location upon which run commands is queried.
// commandID - The command id.
// options - VirtualMachineRunCommandsGetOptions contains the optional parameters for the VirtualMachineRunCommandsClient.Get
// method.
func (client *VirtualMachineRunCommandsClient) Get(ctx context.Context, location string, commandID string, options *VirtualMachineRunCommandsGetOptions) (VirtualMachineRunCommandsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, commandID, options)
	if err != nil {
		return VirtualMachineRunCommandsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineRunCommandsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineRunCommandsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineRunCommandsClient) getCreateRequest(ctx context.Context, location string, commandID string, options *VirtualMachineRunCommandsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands/{commandId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if commandID == "" {
		return nil, errors.New("parameter commandID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{commandId}", url.PathEscape(commandID))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineRunCommandsClient) getHandleResponse(resp *http.Response) (VirtualMachineRunCommandsGetResponse, error) {
	result := VirtualMachineRunCommandsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunCommandDocument); err != nil {
		return VirtualMachineRunCommandsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineRunCommandsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// GetByVirtualMachine - The operation to get the run command.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// vmName - The name of the virtual machine containing the run command.
// runCommandName - The name of the virtual machine run command.
// options - VirtualMachineRunCommandsGetByVirtualMachineOptions contains the optional parameters for the VirtualMachineRunCommandsClient.GetByVirtualMachine
// method.
func (client *VirtualMachineRunCommandsClient) GetByVirtualMachine(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsGetByVirtualMachineOptions) (VirtualMachineRunCommandsGetByVirtualMachineResponse, error) {
	req, err := client.getByVirtualMachineCreateRequest(ctx, resourceGroupName, vmName, runCommandName, options)
	if err != nil {
		return VirtualMachineRunCommandsGetByVirtualMachineResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineRunCommandsGetByVirtualMachineResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineRunCommandsGetByVirtualMachineResponse{}, client.getByVirtualMachineHandleError(resp)
	}
	return client.getByVirtualMachineHandleResponse(resp)
}

// getByVirtualMachineCreateRequest creates the GetByVirtualMachine request.
func (client *VirtualMachineRunCommandsClient) getByVirtualMachineCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, options *VirtualMachineRunCommandsGetByVirtualMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// getByVirtualMachineHandleResponse handles the GetByVirtualMachine response.
func (client *VirtualMachineRunCommandsClient) getByVirtualMachineHandleResponse(resp *http.Response) (VirtualMachineRunCommandsGetByVirtualMachineResponse, error) {
	result := VirtualMachineRunCommandsGetByVirtualMachineResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommand); err != nil {
		return VirtualMachineRunCommandsGetByVirtualMachineResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByVirtualMachineHandleError handles the GetByVirtualMachine error response.
func (client *VirtualMachineRunCommandsClient) getByVirtualMachineHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all available run commands for a subscription in a location.
// If the operation fails it returns a generic error.
// location - The location upon which run commands is queried.
// options - VirtualMachineRunCommandsListOptions contains the optional parameters for the VirtualMachineRunCommandsClient.List
// method.
func (client *VirtualMachineRunCommandsClient) List(location string, options *VirtualMachineRunCommandsListOptions) *VirtualMachineRunCommandsListPager {
	return &VirtualMachineRunCommandsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachineRunCommandsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RunCommandListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VirtualMachineRunCommandsClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineRunCommandsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/runCommands"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineRunCommandsClient) listHandleResponse(resp *http.Response) (VirtualMachineRunCommandsListResponse, error) {
	result := VirtualMachineRunCommandsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RunCommandListResult); err != nil {
		return VirtualMachineRunCommandsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineRunCommandsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByVirtualMachine - The operation to get all run commands of a Virtual Machine.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// vmName - The name of the virtual machine containing the run command.
// options - VirtualMachineRunCommandsListByVirtualMachineOptions contains the optional parameters for the VirtualMachineRunCommandsClient.ListByVirtualMachine
// method.
func (client *VirtualMachineRunCommandsClient) ListByVirtualMachine(resourceGroupName string, vmName string, options *VirtualMachineRunCommandsListByVirtualMachineOptions) *VirtualMachineRunCommandsListByVirtualMachinePager {
	return &VirtualMachineRunCommandsListByVirtualMachinePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByVirtualMachineCreateRequest(ctx, resourceGroupName, vmName, options)
		},
		advancer: func(ctx context.Context, resp VirtualMachineRunCommandsListByVirtualMachineResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineRunCommandsListResult.NextLink)
		},
	}
}

// listByVirtualMachineCreateRequest creates the ListByVirtualMachine request.
func (client *VirtualMachineRunCommandsClient) listByVirtualMachineCreateRequest(ctx context.Context, resourceGroupName string, vmName string, options *VirtualMachineRunCommandsListByVirtualMachineOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listByVirtualMachineHandleResponse handles the ListByVirtualMachine response.
func (client *VirtualMachineRunCommandsClient) listByVirtualMachineHandleResponse(resp *http.Response) (VirtualMachineRunCommandsListByVirtualMachineResponse, error) {
	result := VirtualMachineRunCommandsListByVirtualMachineResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommandsListResult); err != nil {
		return VirtualMachineRunCommandsListByVirtualMachineResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByVirtualMachineHandleError handles the ListByVirtualMachine error response.
func (client *VirtualMachineRunCommandsClient) listByVirtualMachineHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - The operation to update the run command.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// vmName - The name of the virtual machine where the run command should be updated.
// runCommandName - The name of the virtual machine run command.
// runCommand - Parameters supplied to the Update Virtual Machine RunCommand operation.
// options - VirtualMachineRunCommandsBeginUpdateOptions contains the optional parameters for the VirtualMachineRunCommandsClient.BeginUpdate
// method.
func (client *VirtualMachineRunCommandsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineRunCommandsBeginUpdateOptions) (VirtualMachineRunCommandsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
	if err != nil {
		return VirtualMachineRunCommandsUpdatePollerResponse{}, err
	}
	result := VirtualMachineRunCommandsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualMachineRunCommandsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return VirtualMachineRunCommandsUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualMachineRunCommandsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The operation to update the run command.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineRunCommandsClient) update(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineRunCommandsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmName, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineRunCommandsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineRunCommandsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if runCommandName == "" {
		return nil, errors.New("parameter runCommandName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
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
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, runtime.MarshalAsJSON(req, runCommand)
}

// updateHandleError handles the Update error response.
func (client *VirtualMachineRunCommandsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
