// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// ContentItemClient contains the methods for the ContentItem group.
// Don't use this type directly, use NewContentItemClient() instead.
type ContentItemClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewContentItemClient creates a new instance of ContentItemClient with the specified values.
func NewContentItemClient(con *armcore.Connection, subscriptionID string) *ContentItemClient {
	return &ContentItemClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates a new developer portal's content item specified by the provided content type.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ContentItemClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemCreateOrUpdateOptions) (ContentItemCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, options)
	if err != nil {
		return ContentItemCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ContentItemCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return ContentItemCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContentItemClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ContentItemClient) createOrUpdateHandleResponse(resp *azcore.Response) (ContentItemCreateOrUpdateResponse, error) {
	result := ContentItemCreateOrUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.ContentItemContract); err != nil {
		return ContentItemCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ContentItemClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Removes the specified developer portal's content item.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ContentItemClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, ifMatch string, options *ContentItemDeleteOptions) (ContentItemDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, ifMatch, options)
	if err != nil {
		return ContentItemDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ContentItemDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return ContentItemDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ContentItemDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContentItemClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, ifMatch string, options *ContentItemDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ContentItemClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Returns the developer portal's content item specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ContentItemClient) Get(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemGetOptions) (ContentItemGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, options)
	if err != nil {
		return ContentItemGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ContentItemGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ContentItemGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContentItemClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContentItemClient) getHandleResponse(resp *azcore.Response) (ContentItemGetResponse, error) {
	result := ContentItemGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.ContentItemContract); err != nil {
		return ContentItemGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ContentItemClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetEntityTag - Returns the entity state (ETag) version of the developer portal's content item specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ContentItemClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemGetEntityTagOptions) (ContentItemGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, contentItemID, options)
	if err != nil {
		return ContentItemGetEntityTagResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ContentItemGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *ContentItemClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, contentItemID string, options *ContentItemGetEntityTagOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems/{contentItemId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if contentItemID == "" {
		return nil, errors.New("parameter contentItemID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentItemId}", url.PathEscape(contentItemID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *ContentItemClient) getEntityTagHandleResponse(resp *azcore.Response) (ContentItemGetEntityTagResponse, error) {
	result := ContentItemGetEntityTagResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists developer portal's content items specified by the provided content type.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ContentItemClient) ListByService(resourceGroupName string, serviceName string, contentTypeID string, options *ContentItemListByServiceOptions) ContentItemListByServicePager {
	return &contentItemListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, options)
		},
		advancer: func(ctx context.Context, resp ContentItemListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ContentItemCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *ContentItemClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, options *ContentItemListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}/contentItems"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *ContentItemClient) listByServiceHandleResponse(resp *azcore.Response) (ContentItemListByServiceResponse, error) {
	result := ContentItemListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ContentItemCollection); err != nil {
		return ContentItemListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *ContentItemClient) listByServiceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
