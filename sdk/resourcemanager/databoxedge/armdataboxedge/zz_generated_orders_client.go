//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// OrdersClient contains the methods for the Orders group.
// Don't use this type directly, use NewOrdersClient() instead.
type OrdersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOrdersClient creates a new instance of OrdersClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOrdersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OrdersClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &OrdersClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates an order.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The order details of a device.
// resourceGroupName - The resource group name.
// order - The order to be created or updated.
// options - OrdersClientBeginCreateOrUpdateOptions contains the optional parameters for the OrdersClient.BeginCreateOrUpdate
// method.
func (client *OrdersClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, resourceGroupName string, order Order, options *OrdersClientBeginCreateOrUpdateOptions) (OrdersClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, deviceName, resourceGroupName, order, options)
	if err != nil {
		return OrdersClientCreateOrUpdatePollerResponse{}, err
	}
	result := OrdersClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OrdersClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return OrdersClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &OrdersClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an order.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OrdersClient) createOrUpdate(ctx context.Context, deviceName string, resourceGroupName string, order Order, options *OrdersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, resourceGroupName, order, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OrdersClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, order Order, options *OrdersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, order)
}

// BeginDelete - Deletes the order related to the device.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// resourceGroupName - The resource group name.
// options - OrdersClientBeginDeleteOptions contains the optional parameters for the OrdersClient.BeginDelete method.
func (client *OrdersClient) BeginDelete(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientBeginDeleteOptions) (OrdersClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return OrdersClientDeletePollerResponse{}, err
	}
	result := OrdersClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OrdersClient.Delete", "", resp, client.pl)
	if err != nil {
		return OrdersClientDeletePollerResponse{}, err
	}
	result.Poller = &OrdersClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the order related to the device.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OrdersClient) deleteOperation(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, resourceGroupName, options)
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
func (client *OrdersClient) deleteCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a specific order by name.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// resourceGroupName - The resource group name.
// options - OrdersClientGetOptions contains the optional parameters for the OrdersClient.Get method.
func (client *OrdersClient) Get(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientGetOptions) (OrdersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return OrdersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrdersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrdersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OrdersClient) getCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrdersClient) getHandleResponse(resp *http.Response) (OrdersClientGetResponse, error) {
	result := OrdersClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Order); err != nil {
		return OrdersClientGetResponse{}, err
	}
	return result, nil
}

// ListByDataBoxEdgeDevice - Lists all the orders related to a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name.
// resourceGroupName - The resource group name.
// options - OrdersClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the OrdersClient.ListByDataBoxEdgeDevice
// method.
func (client *OrdersClient) ListByDataBoxEdgeDevice(deviceName string, resourceGroupName string, options *OrdersClientListByDataBoxEdgeDeviceOptions) *OrdersClientListByDataBoxEdgeDevicePager {
	return &OrdersClientListByDataBoxEdgeDevicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp OrdersClientListByDataBoxEdgeDeviceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OrderList.NextLink)
		},
	}
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *OrdersClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *OrdersClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (OrdersClientListByDataBoxEdgeDeviceResponse, error) {
	result := OrdersClientListByDataBoxEdgeDeviceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderList); err != nil {
		return OrdersClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}

// ListDCAccessCode - Gets the DCAccess Code
// If the operation fails it returns an *azcore.ResponseError type.
// deviceName - The device name
// resourceGroupName - The resource group name.
// options - OrdersClientListDCAccessCodeOptions contains the optional parameters for the OrdersClient.ListDCAccessCode method.
func (client *OrdersClient) ListDCAccessCode(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientListDCAccessCodeOptions) (OrdersClientListDCAccessCodeResponse, error) {
	req, err := client.listDCAccessCodeCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return OrdersClientListDCAccessCodeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrdersClientListDCAccessCodeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrdersClientListDCAccessCodeResponse{}, runtime.NewResponseError(resp)
	}
	return client.listDCAccessCodeHandleResponse(resp)
}

// listDCAccessCodeCreateRequest creates the ListDCAccessCode request.
func (client *OrdersClient) listDCAccessCodeCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientListDCAccessCodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default/listDCAccessCode"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listDCAccessCodeHandleResponse handles the ListDCAccessCode response.
func (client *OrdersClient) listDCAccessCodeHandleResponse(resp *http.Response) (OrdersClientListDCAccessCodeResponse, error) {
	result := OrdersClientListDCAccessCodeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DCAccessCode); err != nil {
		return OrdersClientListDCAccessCodeResponse{}, err
	}
	return result, nil
}
