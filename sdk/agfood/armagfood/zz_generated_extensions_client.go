// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagfood

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ExtensionsClient contains the methods for the Extensions group.
// Don't use this type directly, use NewExtensionsClient() instead.
type ExtensionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExtensionsClient creates a new instance of ExtensionsClient with the specified values.
func NewExtensionsClient(con *armcore.Connection, subscriptionID string) *ExtensionsClient {
	return &ExtensionsClient{con: con, subscriptionID: subscriptionID}
}

// Create - Install extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExtensionsClient) Create(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsCreateOptions) (ExtensionResponse, error) {
	req, err := client.createCreateRequest(ctx, extensionID, farmBeatsResourceName, resourceGroupName, options)
	if err != nil {
		return ExtensionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExtensionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return ExtensionResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ExtensionsClient) createCreateRequest(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}/extensions/{extensionId}"
	if extensionID == "" {
		return nil, errors.New("parameter extensionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionId}", url.PathEscape(extensionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ExtensionsClient) createHandleResponse(resp *azcore.Response) (ExtensionResponse, error) {
	var val *Extension
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExtensionResponse{}, err
	}
	return ExtensionResponse{RawResponse: resp.Response, Extension: val}, nil
}

// createHandleError handles the Create error response.
func (client *ExtensionsClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Uninstall extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExtensionsClient) Delete(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, extensionID, farmBeatsResourceName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExtensionsClient) deleteCreateRequest(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}/extensions/{extensionId}"
	if extensionID == "" {
		return nil, errors.New("parameter extensionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionId}", url.PathEscape(extensionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ExtensionsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Get installed extension details by extension id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExtensionsClient) Get(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsGetOptions) (ExtensionResponse, error) {
	req, err := client.getCreateRequest(ctx, extensionID, farmBeatsResourceName, resourceGroupName, options)
	if err != nil {
		return ExtensionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExtensionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExtensionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtensionsClient) getCreateRequest(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}/extensions/{extensionId}"
	if extensionID == "" {
		return nil, errors.New("parameter extensionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionId}", url.PathEscape(extensionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtensionsClient) getHandleResponse(resp *azcore.Response) (ExtensionResponse, error) {
	var val *Extension
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExtensionResponse{}, err
	}
	return ExtensionResponse{RawResponse: resp.Response, Extension: val}, nil
}

// getHandleError handles the Get error response.
func (client *ExtensionsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByFarmBeats - Get installed extensions details.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExtensionsClient) ListByFarmBeats(resourceGroupName string, farmBeatsResourceName string, options *ExtensionsListByFarmBeatsOptions) ExtensionListResponsePager {
	return &extensionListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByFarmBeatsCreateRequest(ctx, resourceGroupName, farmBeatsResourceName, options)
		},
		responder: client.listByFarmBeatsHandleResponse,
		errorer:   client.listByFarmBeatsHandleError,
		advancer: func(ctx context.Context, resp ExtensionListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExtensionListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByFarmBeatsCreateRequest creates the ListByFarmBeats request.
func (client *ExtensionsClient) listByFarmBeatsCreateRequest(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *ExtensionsListByFarmBeatsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	if options != nil && options.ExtensionIDs != nil {
		for _, qv := range options.ExtensionIDs {
			reqQP.Add("extensionIds", qv)
		}
	}
	if options != nil && options.ExtensionCategories != nil {
		for _, qv := range options.ExtensionCategories {
			reqQP.Add("extensionCategories", qv)
		}
	}
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByFarmBeatsHandleResponse handles the ListByFarmBeats response.
func (client *ExtensionsClient) listByFarmBeatsHandleResponse(resp *azcore.Response) (ExtensionListResponseResponse, error) {
	var val *ExtensionListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExtensionListResponseResponse{}, err
	}
	return ExtensionListResponseResponse{RawResponse: resp.Response, ExtensionListResponse: val}, nil
}

// listByFarmBeatsHandleError handles the ListByFarmBeats error response.
func (client *ExtensionsClient) listByFarmBeatsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Upgrade to latest extension.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ExtensionsClient) Update(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsUpdateOptions) (ExtensionResponse, error) {
	req, err := client.updateCreateRequest(ctx, extensionID, farmBeatsResourceName, resourceGroupName, options)
	if err != nil {
		return ExtensionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExtensionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExtensionResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ExtensionsClient) updateCreateRequest(ctx context.Context, extensionID string, farmBeatsResourceName string, resourceGroupName string, options *ExtensionsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}/extensions/{extensionId}"
	if extensionID == "" {
		return nil, errors.New("parameter extensionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionId}", url.PathEscape(extensionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ExtensionsClient) updateHandleResponse(resp *azcore.Response) (ExtensionResponse, error) {
	var val *Extension
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExtensionResponse{}, err
	}
	return ExtensionResponse{RawResponse: resp.Response, Extension: val}, nil
}

// updateHandleError handles the Update error response.
func (client *ExtensionsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
