//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridnetwork

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

// DevicesClient contains the methods for the Devices group.
// Don't use this type directly, use NewDevicesClient() instead.
type DevicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDevicesClient creates a new instance of DevicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDevicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevicesClient, error) {
	cl, err := arm.NewClient(moduleName+".DevicesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DevicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deviceName - Resource name for the device resource.
//   - parameters - Parameters supplied to the create or update device operation.
//   - options - DevicesClientBeginCreateOrUpdateOptions contains the optional parameters for the DevicesClient.BeginCreateOrUpdate
//     method.
func (client *DevicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, deviceName string, parameters Device, options *DevicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DevicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, deviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevicesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
func (client *DevicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, deviceName string, parameters Device, options *DevicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, deviceName, parameters, options)
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
func (client *DevicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, parameters Device, options *DevicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/devices/{deviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deviceName - The name of the device resource.
//   - options - DevicesClientBeginDeleteOptions contains the optional parameters for the DevicesClient.BeginDelete method.
func (client *DevicesClient) BeginDelete(ctx context.Context, resourceGroupName string, deviceName string, options *DevicesClientBeginDeleteOptions) (*runtime.Poller[DevicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, deviceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DevicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevicesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
func (client *DevicesClient) deleteOperation(ctx context.Context, resourceGroupName string, deviceName string, options *DevicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, deviceName, options)
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
func (client *DevicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, options *DevicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/devices/{deviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deviceName - The name of the device resource.
//   - options - DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
func (client *DevicesClient) Get(ctx context.Context, resourceGroupName string, deviceName string, options *DevicesClientGetOptions) (DevicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, deviceName, options)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DevicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, options *DevicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/devices/{deviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevicesClient) getHandleResponse(resp *http.Response) (DevicesClientGetResponse, error) {
	result := DevicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Device); err != nil {
		return DevicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all the device resource in a resource group.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DevicesClientListByResourceGroupOptions contains the optional parameters for the DevicesClient.NewListByResourceGroupPager
//     method.
func (client *DevicesClient) NewListByResourceGroupPager(resourceGroupName string, options *DevicesClientListByResourceGroupOptions) *runtime.Pager[DevicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevicesClientListByResourceGroupResponse]{
		More: func(page DevicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevicesClientListByResourceGroupResponse) (DevicesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevicesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DevicesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevicesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DevicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DevicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/devices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DevicesClient) listByResourceGroupHandleResponse(resp *http.Response) (DevicesClientListByResourceGroupResponse, error) {
	result := DevicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceListResult); err != nil {
		return DevicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all the devices in a subscription.
//
// Generated from API version 2021-05-01
//   - options - DevicesClientListBySubscriptionOptions contains the optional parameters for the DevicesClient.NewListBySubscriptionPager
//     method.
func (client *DevicesClient) NewListBySubscriptionPager(options *DevicesClientListBySubscriptionOptions) *runtime.Pager[DevicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevicesClientListBySubscriptionResponse]{
		More: func(page DevicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevicesClientListBySubscriptionResponse) (DevicesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevicesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DevicesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevicesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DevicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DevicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/devices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DevicesClient) listBySubscriptionHandleResponse(resp *http.Response) (DevicesClientListBySubscriptionResponse, error) {
	result := DevicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceListResult); err != nil {
		return DevicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListRegistrationKey - List the registration key for the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deviceName - The name of the device resource.
//   - options - DevicesClientListRegistrationKeyOptions contains the optional parameters for the DevicesClient.ListRegistrationKey
//     method.
func (client *DevicesClient) ListRegistrationKey(ctx context.Context, resourceGroupName string, deviceName string, options *DevicesClientListRegistrationKeyOptions) (DevicesClientListRegistrationKeyResponse, error) {
	req, err := client.listRegistrationKeyCreateRequest(ctx, resourceGroupName, deviceName, options)
	if err != nil {
		return DevicesClientListRegistrationKeyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DevicesClientListRegistrationKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevicesClientListRegistrationKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.listRegistrationKeyHandleResponse(resp)
}

// listRegistrationKeyCreateRequest creates the ListRegistrationKey request.
func (client *DevicesClient) listRegistrationKeyCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, options *DevicesClientListRegistrationKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/devices/{deviceName}/listRegistrationKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRegistrationKeyHandleResponse handles the ListRegistrationKey response.
func (client *DevicesClient) listRegistrationKeyHandleResponse(resp *http.Response) (DevicesClientListRegistrationKeyResponse, error) {
	result := DevicesClientListRegistrationKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceRegistrationKey); err != nil {
		return DevicesClientListRegistrationKeyResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates device tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deviceName - The name of the device resource.
//   - parameters - Parameters supplied to the update device tags operation.
//   - options - DevicesClientUpdateTagsOptions contains the optional parameters for the DevicesClient.UpdateTags method.
func (client *DevicesClient) UpdateTags(ctx context.Context, resourceGroupName string, deviceName string, parameters TagsObject, options *DevicesClientUpdateTagsOptions) (DevicesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, deviceName, parameters, options)
	if err != nil {
		return DevicesClientUpdateTagsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DevicesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevicesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *DevicesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, parameters TagsObject, options *DevicesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/devices/{deviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *DevicesClient) updateTagsHandleResponse(resp *http.Response) (DevicesClientUpdateTagsResponse, error) {
	result := DevicesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Device); err != nil {
		return DevicesClientUpdateTagsResponse{}, err
	}
	return result, nil
}
