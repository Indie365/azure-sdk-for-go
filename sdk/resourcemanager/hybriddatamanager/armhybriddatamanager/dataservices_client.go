//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybriddatamanager

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

// DataServicesClient contains the methods for the DataServices group.
// Don't use this type directly, use NewDataServicesClient() instead.
type DataServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataServicesClient creates a new instance of DataServicesClient with the specified values.
//   - subscriptionID - The Subscription Id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataServicesClient, error) {
	cl, err := arm.NewClient(moduleName+".DataServicesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the data service that matches the data service name given.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-06-01
//   - dataServiceName - The name of the data service that is being queried.
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - DataServicesClientGetOptions contains the optional parameters for the DataServicesClient.Get method.
func (client *DataServicesClient) Get(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, options *DataServicesClientGetOptions) (DataServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, dataServiceName, resourceGroupName, dataManagerName, options)
	if err != nil {
		return DataServicesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataServicesClient) getCreateRequest(ctx context.Context, dataServiceName string, resourceGroupName string, dataManagerName string, options *DataServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices/{dataServiceName}"
	if dataServiceName == "" {
		return nil, errors.New("parameter dataServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataServiceName}", url.PathEscape(dataServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataServicesClient) getHandleResponse(resp *http.Response) (DataServicesClientGetResponse, error) {
	result := DataServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataService); err != nil {
		return DataServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataManagerPager - This method gets all the data services.
//
// Generated from API version 2019-06-01
//   - resourceGroupName - The Resource Group Name
//   - dataManagerName - The name of the DataManager Resource within the specified resource group. DataManager names must be between
//     3 and 24 characters in length and use any alphanumeric and underscore only
//   - options - DataServicesClientListByDataManagerOptions contains the optional parameters for the DataServicesClient.NewListByDataManagerPager
//     method.
func (client *DataServicesClient) NewListByDataManagerPager(resourceGroupName string, dataManagerName string, options *DataServicesClientListByDataManagerOptions) *runtime.Pager[DataServicesClientListByDataManagerResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataServicesClientListByDataManagerResponse]{
		More: func(page DataServicesClientListByDataManagerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataServicesClientListByDataManagerResponse) (DataServicesClientListByDataManagerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataManagerCreateRequest(ctx, resourceGroupName, dataManagerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataServicesClientListByDataManagerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DataServicesClientListByDataManagerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataServicesClientListByDataManagerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataManagerHandleResponse(resp)
		},
	})
}

// listByDataManagerCreateRequest creates the ListByDataManager request.
func (client *DataServicesClient) listByDataManagerCreateRequest(ctx context.Context, resourceGroupName string, dataManagerName string, options *DataServicesClientListByDataManagerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridData/dataManagers/{dataManagerName}/dataServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataManagerName == "" {
		return nil, errors.New("parameter dataManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataManagerName}", url.PathEscape(dataManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataManagerHandleResponse handles the ListByDataManager response.
func (client *DataServicesClient) listByDataManagerHandleResponse(resp *http.Response) (DataServicesClientListByDataManagerResponse, error) {
	result := DataServicesClientListByDataManagerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataServiceList); err != nil {
		return DataServicesClientListByDataManagerResponse{}, err
	}
	return result, nil
}
