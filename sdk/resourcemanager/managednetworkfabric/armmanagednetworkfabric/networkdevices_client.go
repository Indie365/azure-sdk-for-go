//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetworkfabric

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

// NetworkDevicesClient contains the methods for the NetworkDevices group.
// Don't use this type directly, use NewNetworkDevicesClient() instead.
type NetworkDevicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkDevicesClient creates a new instance of NetworkDevicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkDevicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkDevicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkDevicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a Network Device resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginCreateOptions contains the optional parameters for the NetworkDevicesClient.BeginCreate
//     method.
func (client *NetworkDevicesClient) BeginCreate(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevice, options *NetworkDevicesClientBeginCreateOptions) (*runtime.Poller[NetworkDevicesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkDevicesClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a Network Device resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) create(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevice, options *NetworkDevicesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkDevicesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *NetworkDevicesClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevice, options *NetworkDevicesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - options - NetworkDevicesClientBeginDeleteOptions contains the optional parameters for the NetworkDevicesClient.BeginDelete
//     method.
func (client *NetworkDevicesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginDeleteOptions) (*runtime.Poller[NetworkDevicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkDeviceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkDevicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkDevicesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkDeviceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkDevicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the Network Device resource details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - options - NetworkDevicesClientGetOptions contains the optional parameters for the NetworkDevicesClient.Get method.
func (client *NetworkDevicesClient) Get(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientGetOptions) (NetworkDevicesClientGetResponse, error) {
	var err error
	const operationName = "NetworkDevicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkDeviceName, options)
	if err != nil {
		return NetworkDevicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkDevicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkDevicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkDevicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkDevicesClient) getHandleResponse(resp *http.Response) (NetworkDevicesClientGetResponse, error) {
	result := NetworkDevicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkDevice); err != nil {
		return NetworkDevicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the Network Device resources in a given resource group.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - NetworkDevicesClientListByResourceGroupOptions contains the optional parameters for the NetworkDevicesClient.NewListByResourceGroupPager
//     method.
func (client *NetworkDevicesClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkDevicesClientListByResourceGroupOptions) *runtime.Pager[NetworkDevicesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkDevicesClientListByResourceGroupResponse]{
		More: func(page NetworkDevicesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkDevicesClientListByResourceGroupResponse) (NetworkDevicesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkDevicesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return NetworkDevicesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkDevicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkDevicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkDevicesClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkDevicesClientListByResourceGroupResponse, error) {
	result := NetworkDevicesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkDevicesListResult); err != nil {
		return NetworkDevicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the Network Device resources in a given subscription.
//
// Generated from API version 2023-06-15
//   - options - NetworkDevicesClientListBySubscriptionOptions contains the optional parameters for the NetworkDevicesClient.NewListBySubscriptionPager
//     method.
func (client *NetworkDevicesClient) NewListBySubscriptionPager(options *NetworkDevicesClientListBySubscriptionOptions) *runtime.Pager[NetworkDevicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkDevicesClientListBySubscriptionResponse]{
		More: func(page NetworkDevicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkDevicesClientListBySubscriptionResponse) (NetworkDevicesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkDevicesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return NetworkDevicesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkDevicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkDevicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/networkDevices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkDevicesClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkDevicesClientListBySubscriptionResponse, error) {
	result := NetworkDevicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkDevicesListResult); err != nil {
		return NetworkDevicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginReboot - Reboot the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginRebootOptions contains the optional parameters for the NetworkDevicesClient.BeginReboot
//     method.
func (client *NetworkDevicesClient) BeginReboot(ctx context.Context, resourceGroupName string, networkDeviceName string, body RebootProperties, options *NetworkDevicesClientBeginRebootOptions) (*runtime.Poller[NetworkDevicesClientRebootResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reboot(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientRebootResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkDevicesClientRebootResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Reboot - Reboot the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) reboot(ctx context.Context, resourceGroupName string, networkDeviceName string, body RebootProperties, options *NetworkDevicesClientBeginRebootOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkDevicesClient.BeginReboot"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rebootCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// rebootCreateRequest creates the Reboot request.
func (client *NetworkDevicesClient) rebootCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body RebootProperties, options *NetworkDevicesClientBeginRebootOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/reboot"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginRefreshConfiguration - Refreshes the configuration the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - options - NetworkDevicesClientBeginRefreshConfigurationOptions contains the optional parameters for the NetworkDevicesClient.BeginRefreshConfiguration
//     method.
func (client *NetworkDevicesClient) BeginRefreshConfiguration(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginRefreshConfigurationOptions) (*runtime.Poller[NetworkDevicesClientRefreshConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refreshConfiguration(ctx, resourceGroupName, networkDeviceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientRefreshConfigurationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkDevicesClientRefreshConfigurationResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RefreshConfiguration - Refreshes the configuration the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) refreshConfiguration(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginRefreshConfigurationOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkDevicesClient.BeginRefreshConfiguration"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.refreshConfigurationCreateRequest(ctx, resourceGroupName, networkDeviceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// refreshConfigurationCreateRequest creates the RefreshConfiguration request.
func (client *NetworkDevicesClient) refreshConfigurationCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, options *NetworkDevicesClientBeginRefreshConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/refreshConfiguration"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update certain properties of the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Network Device properties to update.
//   - options - NetworkDevicesClientBeginUpdateOptions contains the optional parameters for the NetworkDevicesClient.BeginUpdate
//     method.
func (client *NetworkDevicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevicePatchParameters, options *NetworkDevicesClientBeginUpdateOptions) (*runtime.Poller[NetworkDevicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkDevicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update certain properties of the Network Device resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) update(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevicePatchParameters, options *NetworkDevicesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkDevicesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *NetworkDevicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body NetworkDevicePatchParameters, options *NetworkDevicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateAdministrativeState - Updates the Administrative state of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginUpdateAdministrativeStateOptions contains the optional parameters for the NetworkDevicesClient.BeginUpdateAdministrativeState
//     method.
func (client *NetworkDevicesClient) BeginUpdateAdministrativeState(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateDeviceAdministrativeState, options *NetworkDevicesClientBeginUpdateAdministrativeStateOptions) (*runtime.Poller[NetworkDevicesClientUpdateAdministrativeStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateAdministrativeState(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientUpdateAdministrativeStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkDevicesClientUpdateAdministrativeStateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateAdministrativeState - Updates the Administrative state of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) updateAdministrativeState(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateDeviceAdministrativeState, options *NetworkDevicesClientBeginUpdateAdministrativeStateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkDevicesClient.BeginUpdateAdministrativeState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateAdministrativeStateCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateAdministrativeStateCreateRequest creates the UpdateAdministrativeState request.
func (client *NetworkDevicesClient) updateAdministrativeStateCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateDeviceAdministrativeState, options *NetworkDevicesClientBeginUpdateAdministrativeStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/updateAdministrativeState"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpgrade - Upgrades the version of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkDeviceName - Name of the Network Device.
//   - body - Request payload.
//   - options - NetworkDevicesClientBeginUpgradeOptions contains the optional parameters for the NetworkDevicesClient.BeginUpgrade
//     method.
func (client *NetworkDevicesClient) BeginUpgrade(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateVersion, options *NetworkDevicesClientBeginUpgradeOptions) (*runtime.Poller[NetworkDevicesClientUpgradeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgrade(ctx, resourceGroupName, networkDeviceName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkDevicesClientUpgradeResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkDevicesClientUpgradeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Upgrade - Upgrades the version of the Network Device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkDevicesClient) upgrade(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateVersion, options *NetworkDevicesClientBeginUpgradeOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkDevicesClient.BeginUpgrade"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.upgradeCreateRequest(ctx, resourceGroupName, networkDeviceName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// upgradeCreateRequest creates the Upgrade request.
func (client *NetworkDevicesClient) upgradeCreateRequest(ctx context.Context, resourceGroupName string, networkDeviceName string, body UpdateVersion, options *NetworkDevicesClientBeginUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkDevices/{networkDeviceName}/upgrade"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkDeviceName == "" {
		return nil, errors.New("parameter networkDeviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkDeviceName}", url.PathEscape(networkDeviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
