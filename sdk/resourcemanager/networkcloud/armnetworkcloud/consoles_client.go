//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud

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

// ConsolesClient contains the methods for the Consoles group.
// Don't use this type directly, use NewConsolesClient() instead.
type ConsolesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConsolesClient creates a new instance of ConsolesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConsolesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConsolesClient, error) {
	cl, err := arm.NewClient(moduleName+".ConsolesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConsolesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new virtual machine console or update the properties of the existing virtual machine console.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-12-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - consoleName - The name of the virtual machine console.
//   - consoleParameters - The request body.
//   - options - ConsolesClientBeginCreateOrUpdateOptions contains the optional parameters for the ConsolesClient.BeginCreateOrUpdate
//     method.
func (client *ConsolesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleParameters Console, options *ConsolesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ConsolesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualMachineName, consoleName, consoleParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConsolesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ConsolesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new virtual machine console or update the properties of the existing virtual machine console.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-12-preview
func (client *ConsolesClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleParameters Console, options *ConsolesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualMachineName, consoleName, consoleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConsolesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleParameters Console, options *ConsolesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/consoles/{consoleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	if consoleName == "" {
		return nil, errors.New("parameter consoleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{consoleName}", url.PathEscape(consoleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, consoleParameters)
}

// BeginDelete - Delete the provided virtual machine console.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-12-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - consoleName - The name of the virtual machine console.
//   - options - ConsolesClientBeginDeleteOptions contains the optional parameters for the ConsolesClient.BeginDelete method.
func (client *ConsolesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, options *ConsolesClientBeginDeleteOptions) (*runtime.Poller[ConsolesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualMachineName, consoleName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConsolesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ConsolesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided virtual machine console.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-12-preview
func (client *ConsolesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, options *ConsolesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualMachineName, consoleName, options)
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
func (client *ConsolesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, options *ConsolesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/consoles/{consoleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	if consoleName == "" {
		return nil, errors.New("parameter consoleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{consoleName}", url.PathEscape(consoleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of the provided virtual machine console.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-12-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - consoleName - The name of the virtual machine console.
//   - options - ConsolesClientGetOptions contains the optional parameters for the ConsolesClient.Get method.
func (client *ConsolesClient) Get(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, options *ConsolesClientGetOptions) (ConsolesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualMachineName, consoleName, options)
	if err != nil {
		return ConsolesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConsolesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConsolesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConsolesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, options *ConsolesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/consoles/{consoleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	if consoleName == "" {
		return nil, errors.New("parameter consoleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{consoleName}", url.PathEscape(consoleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConsolesClient) getHandleResponse(resp *http.Response) (ConsolesClientGetResponse, error) {
	result := ConsolesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Console); err != nil {
		return ConsolesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of virtual machine consoles in the provided resource group.
//
// Generated from API version 2022-12-12-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - options - ConsolesClientListByResourceGroupOptions contains the optional parameters for the ConsolesClient.NewListByResourceGroupPager
//     method.
func (client *ConsolesClient) NewListByResourceGroupPager(resourceGroupName string, virtualMachineName string, options *ConsolesClientListByResourceGroupOptions) *runtime.Pager[ConsolesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConsolesClientListByResourceGroupResponse]{
		More: func(page ConsolesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConsolesClientListByResourceGroupResponse) (ConsolesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, virtualMachineName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConsolesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConsolesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConsolesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConsolesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, options *ConsolesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/consoles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConsolesClient) listByResourceGroupHandleResponse(resp *http.Response) (ConsolesClientListByResourceGroupResponse, error) {
	result := ConsolesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsoleList); err != nil {
		return ConsolesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch the properties of the provided virtual machine console, or update the tags associated with the virtual
// machine console. Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-12-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - virtualMachineName - The name of the virtual machine.
//   - consoleName - The name of the virtual machine console.
//   - consoleUpdateParameters - The request body.
//   - options - ConsolesClientBeginUpdateOptions contains the optional parameters for the ConsolesClient.BeginUpdate method.
func (client *ConsolesClient) BeginUpdate(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleUpdateParameters ConsolePatchParameters, options *ConsolesClientBeginUpdateOptions) (*runtime.Poller[ConsolesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, virtualMachineName, consoleName, consoleUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConsolesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ConsolesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Patch the properties of the provided virtual machine console, or update the tags associated with the virtual machine
// console. Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-12-preview
func (client *ConsolesClient) update(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleUpdateParameters ConsolePatchParameters, options *ConsolesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualMachineName, consoleName, consoleUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ConsolesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleUpdateParameters ConsolePatchParameters, options *ConsolesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/virtualMachines/{virtualMachineName}/consoles/{consoleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualMachineName == "" {
		return nil, errors.New("parameter virtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualMachineName}", url.PathEscape(virtualMachineName))
	if consoleName == "" {
		return nil, errors.New("parameter consoleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{consoleName}", url.PathEscape(consoleName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, consoleUpdateParameters)
}
