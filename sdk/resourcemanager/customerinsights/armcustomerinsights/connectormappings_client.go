//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// ConnectorMappingsClient contains the methods for the ConnectorMappings group.
// Don't use this type directly, use NewConnectorMappingsClient() instead.
type ConnectorMappingsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewConnectorMappingsClient creates a new instance of ConnectorMappingsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectorMappingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectorMappingsClient, error) {
	cl, err := arm.NewClient(moduleName+".ConnectorMappingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectorMappingsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a connector mapping or updates an existing connector mapping in the connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - connectorName - The name of the connector.
//   - mappingName - The name of the connector mapping.
//   - parameters - Parameters supplied to the CreateOrUpdate Connector Mapping operation.
//   - options - ConnectorMappingsClientCreateOrUpdateOptions contains the optional parameters for the ConnectorMappingsClient.CreateOrUpdate
//     method.
func (client *ConnectorMappingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string, parameters ConnectorMappingResourceFormat, options *ConnectorMappingsClientCreateOrUpdateOptions) (ConnectorMappingsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, connectorName, mappingName, parameters, options)
	if err != nil {
		return ConnectorMappingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectorMappingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConnectorMappingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectorMappingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string, parameters ConnectorMappingResourceFormat, options *ConnectorMappingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectorMappingsClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectorMappingsClientCreateOrUpdateResponse, error) {
	result := ConnectorMappingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectorMappingResourceFormat); err != nil {
		return ConnectorMappingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a connector mapping in the connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - connectorName - The name of the connector.
//   - mappingName - The name of the connector mapping.
//   - options - ConnectorMappingsClientDeleteOptions contains the optional parameters for the ConnectorMappingsClient.Delete
//     method.
func (client *ConnectorMappingsClient) Delete(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string, options *ConnectorMappingsClientDeleteOptions) (ConnectorMappingsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, connectorName, mappingName, options)
	if err != nil {
		return ConnectorMappingsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectorMappingsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConnectorMappingsClientDeleteResponse{}, err
	}
	return ConnectorMappingsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectorMappingsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string, options *ConnectorMappingsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a connector mapping in the connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - connectorName - The name of the connector.
//   - mappingName - The name of the connector mapping.
//   - options - ConnectorMappingsClientGetOptions contains the optional parameters for the ConnectorMappingsClient.Get method.
func (client *ConnectorMappingsClient) Get(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string, options *ConnectorMappingsClientGetOptions) (ConnectorMappingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, connectorName, mappingName, options)
	if err != nil {
		return ConnectorMappingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectorMappingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectorMappingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectorMappingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectorName string, mappingName string, options *ConnectorMappingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings/{mappingName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectorMappingsClient) getHandleResponse(resp *http.Response) (ConnectorMappingsClientGetResponse, error) {
	result := ConnectorMappingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectorMappingResourceFormat); err != nil {
		return ConnectorMappingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByConnectorPager - Gets all the connector mappings in the specified connector.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - connectorName - The name of the connector.
//   - options - ConnectorMappingsClientListByConnectorOptions contains the optional parameters for the ConnectorMappingsClient.NewListByConnectorPager
//     method.
func (client *ConnectorMappingsClient) NewListByConnectorPager(resourceGroupName string, hubName string, connectorName string, options *ConnectorMappingsClientListByConnectorOptions) (*runtime.Pager[ConnectorMappingsClientListByConnectorResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConnectorMappingsClientListByConnectorResponse]{
		More: func(page ConnectorMappingsClientListByConnectorResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectorMappingsClientListByConnectorResponse) (ConnectorMappingsClientListByConnectorResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByConnectorCreateRequest(ctx, resourceGroupName, hubName, connectorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectorMappingsClientListByConnectorResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConnectorMappingsClientListByConnectorResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectorMappingsClientListByConnectorResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByConnectorHandleResponse(resp)
		},
	})
}

// listByConnectorCreateRequest creates the ListByConnector request.
func (client *ConnectorMappingsClient) listByConnectorCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectorName string, options *ConnectorMappingsClientListByConnectorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/connectors/{connectorName}/mappings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if connectorName == "" {
		return nil, errors.New("parameter connectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectorName}", url.PathEscape(connectorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByConnectorHandleResponse handles the ListByConnector response.
func (client *ConnectorMappingsClient) listByConnectorHandleResponse(resp *http.Response) (ConnectorMappingsClientListByConnectorResponse, error) {
	result := ConnectorMappingsClientListByConnectorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectorMappingListResult); err != nil {
		return ConnectorMappingsClientListByConnectorResponse{}, err
	}
	return result, nil
}

