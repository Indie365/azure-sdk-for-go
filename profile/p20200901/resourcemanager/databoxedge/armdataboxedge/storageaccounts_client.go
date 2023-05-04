//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// StorageAccountsClient contains the methods for the StorageAccounts group.
// Don't use this type directly, use NewStorageAccountsClient() instead.
type StorageAccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageAccountsClient creates a new instance of StorageAccountsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageAccountsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armdataboxedge.StorageAccountsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageAccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new StorageAccount or updates an existing StorageAccount on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - deviceName - The device name.
//   - storageAccountName - The StorageAccount name.
//   - resourceGroupName - The resource group name.
//   - storageAccount - The StorageAccount properties.
//   - options - StorageAccountsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageAccountsClient.BeginCreateOrUpdate
//     method.
func (client *StorageAccountsClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, storageAccount StorageAccount, options *StorageAccountsClientBeginCreateOrUpdateOptions) (*runtime.Poller[StorageAccountsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, storageAccountName, resourceGroupName, storageAccount, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StorageAccountsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[StorageAccountsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new StorageAccount or updates an existing StorageAccount on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
func (client *StorageAccountsClient) createOrUpdate(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, storageAccount StorageAccount, options *StorageAccountsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, storageAccountName, resourceGroupName, storageAccount, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StorageAccountsClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, storageAccount StorageAccount, options *StorageAccountsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, storageAccount)
}

// BeginDelete - Deletes the StorageAccount on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - deviceName - The device name.
//   - storageAccountName - The StorageAccount name.
//   - resourceGroupName - The resource group name.
//   - options - StorageAccountsClientBeginDeleteOptions contains the optional parameters for the StorageAccountsClient.BeginDelete
//     method.
func (client *StorageAccountsClient) BeginDelete(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *StorageAccountsClientBeginDeleteOptions) (*runtime.Poller[StorageAccountsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, storageAccountName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[StorageAccountsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[StorageAccountsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the StorageAccount on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
func (client *StorageAccountsClient) deleteOperation(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *StorageAccountsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, storageAccountName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageAccountsClient) deleteCreateRequest(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *StorageAccountsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a StorageAccount by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - deviceName - The device name.
//   - storageAccountName - The storage account name.
//   - resourceGroupName - The resource group name.
//   - options - StorageAccountsClientGetOptions contains the optional parameters for the StorageAccountsClient.Get method.
func (client *StorageAccountsClient) Get(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *StorageAccountsClientGetOptions) (StorageAccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, storageAccountName, resourceGroupName, options)
	if err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageAccountsClient) getCreateRequest(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *StorageAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts/{storageAccountName}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
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
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageAccountsClient) getHandleResponse(resp *http.Response) (StorageAccountsClientGetResponse, error) {
	result := StorageAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccount); err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataBoxEdgeDevicePager - Lists all the storage accounts in a Data Box Edge/Data Box Gateway device.
//
// Generated from API version 2019-08-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - StorageAccountsClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the StorageAccountsClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *StorageAccountsClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *StorageAccountsClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[StorageAccountsClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountsClientListByDataBoxEdgeDeviceResponse]{
		More: func(page StorageAccountsClientListByDataBoxEdgeDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountsClientListByDataBoxEdgeDeviceResponse) (StorageAccountsClientListByDataBoxEdgeDeviceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageAccountsClientListByDataBoxEdgeDeviceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return StorageAccountsClientListByDataBoxEdgeDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageAccountsClientListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *StorageAccountsClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *StorageAccountsClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccounts"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *StorageAccountsClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (StorageAccountsClientListByDataBoxEdgeDeviceResponse, error) {
	result := StorageAccountsClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountList); err != nil {
		return StorageAccountsClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}
