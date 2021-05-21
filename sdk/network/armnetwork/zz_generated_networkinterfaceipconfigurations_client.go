// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// NetworkInterfaceIPConfigurationsClient contains the methods for the NetworkInterfaceIPConfigurations group.
// Don't use this type directly, use NewNetworkInterfaceIPConfigurationsClient() instead.
type NetworkInterfaceIPConfigurationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkInterfaceIPConfigurationsClient creates a new instance of NetworkInterfaceIPConfigurationsClient with the specified values.
func NewNetworkInterfaceIPConfigurationsClient(con *armcore.Connection, subscriptionID string) *NetworkInterfaceIPConfigurationsClient {
	return &NetworkInterfaceIPConfigurationsClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the specified network interface ip configuration.
// If the operation fails it returns the *CloudError error type.
func (client *NetworkInterfaceIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *NetworkInterfaceIPConfigurationsGetOptions) (NetworkInterfaceIPConfigurationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkInterfaceName, ipConfigurationName, options)
	if err != nil {
		return NetworkInterfaceIPConfigurationResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NetworkInterfaceIPConfigurationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NetworkInterfaceIPConfigurationResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkInterfaceIPConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *NetworkInterfaceIPConfigurationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if ipConfigurationName == "" {
		return nil, errors.New("parameter ipConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkInterfaceIPConfigurationsClient) getHandleResponse(resp *azcore.Response) (NetworkInterfaceIPConfigurationResponse, error) {
	var val *NetworkInterfaceIPConfiguration
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkInterfaceIPConfigurationResponse{}, err
	}
	return NetworkInterfaceIPConfigurationResponse{RawResponse: resp.Response, NetworkInterfaceIPConfiguration: val}, nil
}

// getHandleError handles the Get error response.
func (client *NetworkInterfaceIPConfigurationsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Get all ip configurations in a network interface.
// If the operation fails it returns the *CloudError error type.
func (client *NetworkInterfaceIPConfigurationsClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceIPConfigurationsListOptions) NetworkInterfaceIPConfigurationListResultPager {
	return &networkInterfaceIPConfigurationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp NetworkInterfaceIPConfigurationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceIPConfigurationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *NetworkInterfaceIPConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceIPConfigurationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkInterfaceIPConfigurationsClient) listHandleResponse(resp *azcore.Response) (NetworkInterfaceIPConfigurationListResultResponse, error) {
	var val *NetworkInterfaceIPConfigurationListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return NetworkInterfaceIPConfigurationListResultResponse{}, err
	}
	return NetworkInterfaceIPConfigurationListResultResponse{RawResponse: resp.Response, NetworkInterfaceIPConfigurationListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *NetworkInterfaceIPConfigurationsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
