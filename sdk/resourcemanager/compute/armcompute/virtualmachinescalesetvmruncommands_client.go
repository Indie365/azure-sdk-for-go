//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

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

// VirtualMachineScaleSetVMRunCommandsClient contains the methods for the VirtualMachineScaleSetVMRunCommands group.
// Don't use this type directly, use NewVirtualMachineScaleSetVMRunCommandsClient() instead.
type VirtualMachineScaleSetVMRunCommandsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineScaleSetVMRunCommandsClient creates a new instance of VirtualMachineScaleSetVMRunCommandsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineScaleSetVMRunCommandsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineScaleSetVMRunCommandsClient, error) {
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
	client := &VirtualMachineScaleSetVMRunCommandsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// runCommand - Parameters supplied to the Create Virtual Machine RunCommand operation.
// options - VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions contains the optional parameters for the
// VirtualMachineScaleSetVMRunCommandsClient.BeginCreateOrUpdate method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineScaleSetVMRunCommandsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - The operation to create or update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
func (client *VirtualMachineScaleSetVMRunCommandsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
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
func (client *VirtualMachineScaleSetVMRunCommandsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
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
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, runtime.MarshalAsJSON(req, runCommand)
}

// BeginDelete - The operation to delete the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// options - VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.BeginDelete
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineScaleSetVMRunCommandsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineScaleSetVMRunCommandsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineScaleSetVMRunCommandsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - The operation to delete the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
func (client *VirtualMachineScaleSetVMRunCommandsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
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
func (client *VirtualMachineScaleSetVMRunCommandsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
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
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// Get - The operation to get the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// options - VirtualMachineScaleSetVMRunCommandsClientGetOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.Get
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientGetOptions) (VirtualMachineScaleSetVMRunCommandsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, options)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
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
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) getHandleResponse(resp *http.Response) (VirtualMachineScaleSetVMRunCommandsClientGetResponse, error) {
	result := VirtualMachineScaleSetVMRunCommandsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommand); err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation to get all run commands of an instance in Virtual Machine Scaleset.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// options - VirtualMachineScaleSetVMRunCommandsClientListOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.List
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) NewListPager(resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMRunCommandsClientListOptions) *runtime.Pager[VirtualMachineScaleSetVMRunCommandsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualMachineScaleSetVMRunCommandsClientListResponse]{
		More: func(page VirtualMachineScaleSetVMRunCommandsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualMachineScaleSetVMRunCommandsClientListResponse) (VirtualMachineScaleSetVMRunCommandsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualMachineScaleSetVMRunCommandsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualMachineScaleSetVMRunCommandsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualMachineScaleSetVMRunCommandsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, options *VirtualMachineScaleSetVMRunCommandsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
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
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) listHandleResponse(resp *http.Response) (VirtualMachineScaleSetVMRunCommandsClientListResponse, error) {
	result := VirtualMachineScaleSetVMRunCommandsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineRunCommandsListResult); err != nil {
		return VirtualMachineScaleSetVMRunCommandsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// vmScaleSetName - The name of the VM scale set.
// instanceID - The instance ID of the virtual machine.
// runCommandName - The name of the virtual machine run command.
// runCommand - Parameters supplied to the Update Virtual Machine RunCommand operation.
// options - VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions contains the optional parameters for the VirtualMachineScaleSetVMRunCommandsClient.BeginUpdate
// method.
func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineScaleSetVMRunCommandsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineScaleSetVMRunCommandsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineScaleSetVMRunCommandsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - The operation to update the VMSS VM run command.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
func (client *VirtualMachineScaleSetVMRunCommandsClient) update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceID, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceID string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
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
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, runtime.MarshalAsJSON(req, runCommand)
}