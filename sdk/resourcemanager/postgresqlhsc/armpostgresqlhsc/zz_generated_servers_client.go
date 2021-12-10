//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlhsc

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

// ServersClient contains the methods for the Servers group.
// Don't use this type directly, use NewServersClient() instead.
type ServersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServersClient creates a new instance of ServersClient with the specified values.
func NewServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ServersClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets information about a server in server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) Get(ctx context.Context, resourceGroupName string, serverGroupName string, serverName string, options *ServersGetOptions) (ServersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverGroupName, serverName, options)
	if err != nil {
		return ServersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, serverName string, options *ServersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/servers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServersClient) getHandleResponse(resp *http.Response) (ServersGetResponse, error) {
	result := ServersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerGroupServer); err != nil {
		return ServersGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServersClient) getHandleError(resp *http.Response) error {
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

// ListByServerGroup - Lists servers of a server group.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) ListByServerGroup(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServersListByServerGroupOptions) (ServersListByServerGroupResponse, error) {
	req, err := client.listByServerGroupCreateRequest(ctx, resourceGroupName, serverGroupName, options)
	if err != nil {
		return ServersListByServerGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersListByServerGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServersListByServerGroupResponse{}, client.listByServerGroupHandleError(resp)
	}
	return client.listByServerGroupHandleResponse(resp)
}

// listByServerGroupCreateRequest creates the ListByServerGroup request.
func (client *ServersClient) listByServerGroupCreateRequest(ctx context.Context, resourceGroupName string, serverGroupName string, options *ServersListByServerGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/serverGroupsv2/{serverGroupName}/servers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverGroupName == "" {
		return nil, errors.New("parameter serverGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverGroupName}", url.PathEscape(serverGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-05-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerGroupHandleResponse handles the ListByServerGroup response.
func (client *ServersClient) listByServerGroupHandleResponse(resp *http.Response) (ServersListByServerGroupResponse, error) {
	result := ServersListByServerGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerGroupServerListResult); err != nil {
		return ServersListByServerGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByServerGroupHandleError handles the ListByServerGroup error response.
func (client *ServersClient) listByServerGroupHandleError(resp *http.Response) error {
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
