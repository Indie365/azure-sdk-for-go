//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsqlvirtualmachine

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

// AvailabilityGroupListenersClient contains the methods for the AvailabilityGroupListeners group.
// Don't use this type directly, use NewAvailabilityGroupListenersClient() instead.
type AvailabilityGroupListenersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAvailabilityGroupListenersClient creates a new instance of AvailabilityGroupListenersClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailabilityGroupListenersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailabilityGroupListenersClient, error) {
	cl, err := arm.NewClient(moduleName+".AvailabilityGroupListenersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailabilityGroupListenersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an availability group listener.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - availabilityGroupListenerName - Name of the availability group listener.
//   - parameters - The availability group listener.
//   - options - AvailabilityGroupListenersClientBeginCreateOrUpdateOptions contains the optional parameters for the AvailabilityGroupListenersClient.BeginCreateOrUpdate
//     method.
func (client *AvailabilityGroupListenersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, parameters AvailabilityGroupListener, options *AvailabilityGroupListenersClientBeginCreateOrUpdateOptions) (*runtime.Poller[AvailabilityGroupListenersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, sqlVirtualMachineGroupName, availabilityGroupListenerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AvailabilityGroupListenersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AvailabilityGroupListenersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an availability group listener.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
func (client *AvailabilityGroupListenersClient) createOrUpdate(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, parameters AvailabilityGroupListener, options *AvailabilityGroupListenersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, availabilityGroupListenerName, parameters, options)
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
func (client *AvailabilityGroupListenersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, parameters AvailabilityGroupListener, options *AvailabilityGroupListenersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if availabilityGroupListenerName == "" {
		return nil, errors.New("parameter availabilityGroupListenerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilityGroupListenerName}", url.PathEscape(availabilityGroupListenerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an availability group listener.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - availabilityGroupListenerName - Name of the availability group listener.
//   - options - AvailabilityGroupListenersClientBeginDeleteOptions contains the optional parameters for the AvailabilityGroupListenersClient.BeginDelete
//     method.
func (client *AvailabilityGroupListenersClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, options *AvailabilityGroupListenersClientBeginDeleteOptions) (*runtime.Poller[AvailabilityGroupListenersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlVirtualMachineGroupName, availabilityGroupListenerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AvailabilityGroupListenersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[AvailabilityGroupListenersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an availability group listener.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
func (client *AvailabilityGroupListenersClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, options *AvailabilityGroupListenersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, availabilityGroupListenerName, options)
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
func (client *AvailabilityGroupListenersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, options *AvailabilityGroupListenersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if availabilityGroupListenerName == "" {
		return nil, errors.New("parameter availabilityGroupListenerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilityGroupListenerName}", url.PathEscape(availabilityGroupListenerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an availability group listener.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - availabilityGroupListenerName - Name of the availability group listener.
//   - options - AvailabilityGroupListenersClientGetOptions contains the optional parameters for the AvailabilityGroupListenersClient.Get
//     method.
func (client *AvailabilityGroupListenersClient) Get(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, options *AvailabilityGroupListenersClientGetOptions) (AvailabilityGroupListenersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, availabilityGroupListenerName, options)
	if err != nil {
		return AvailabilityGroupListenersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvailabilityGroupListenersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AvailabilityGroupListenersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AvailabilityGroupListenersClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, availabilityGroupListenerName string, options *AvailabilityGroupListenersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners/{availabilityGroupListenerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if availabilityGroupListenerName == "" {
		return nil, errors.New("parameter availabilityGroupListenerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{availabilityGroupListenerName}", url.PathEscape(availabilityGroupListenerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AvailabilityGroupListenersClient) getHandleResponse(resp *http.Response) (AvailabilityGroupListenersClientGetResponse, error) {
	result := AvailabilityGroupListenersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityGroupListener); err != nil {
		return AvailabilityGroupListenersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByGroupPager - Lists all availability group listeners in a SQL virtual machine group.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - sqlVirtualMachineGroupName - Name of the SQL virtual machine group.
//   - options - AvailabilityGroupListenersClientListByGroupOptions contains the optional parameters for the AvailabilityGroupListenersClient.NewListByGroupPager
//     method.
func (client *AvailabilityGroupListenersClient) NewListByGroupPager(resourceGroupName string, sqlVirtualMachineGroupName string, options *AvailabilityGroupListenersClientListByGroupOptions) *runtime.Pager[AvailabilityGroupListenersClientListByGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailabilityGroupListenersClientListByGroupResponse]{
		More: func(page AvailabilityGroupListenersClientListByGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailabilityGroupListenersClientListByGroupResponse) (AvailabilityGroupListenersClientListByGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByGroupCreateRequest(ctx, resourceGroupName, sqlVirtualMachineGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailabilityGroupListenersClientListByGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AvailabilityGroupListenersClientListByGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailabilityGroupListenersClientListByGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByGroupHandleResponse(resp)
		},
	})
}

// listByGroupCreateRequest creates the ListByGroup request.
func (client *AvailabilityGroupListenersClient) listByGroupCreateRequest(ctx context.Context, resourceGroupName string, sqlVirtualMachineGroupName string, options *AvailabilityGroupListenersClientListByGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/{sqlVirtualMachineGroupName}/availabilityGroupListeners"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlVirtualMachineGroupName == "" {
		return nil, errors.New("parameter sqlVirtualMachineGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlVirtualMachineGroupName}", url.PathEscape(sqlVirtualMachineGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByGroupHandleResponse handles the ListByGroup response.
func (client *AvailabilityGroupListenersClient) listByGroupHandleResponse(resp *http.Response) (AvailabilityGroupListenersClientListByGroupResponse, error) {
	result := AvailabilityGroupListenersClientListByGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityGroupListenerListResult); err != nil {
		return AvailabilityGroupListenersClientListByGroupResponse{}, err
	}
	return result, nil
}
