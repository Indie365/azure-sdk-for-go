// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// DataCollectionRulesClient contains the methods for the DataCollectionRules group.
// Don't use this type directly, use NewDataCollectionRulesClient() instead.
type DataCollectionRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDataCollectionRulesClient creates a new instance of DataCollectionRulesClient with the specified values.
func NewDataCollectionRulesClient(con *armcore.Connection, subscriptionID string) *DataCollectionRulesClient {
	return &DataCollectionRulesClient{con: con, subscriptionID: subscriptionID}
}

// Create - Creates or updates a data collection rule.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionRulesClient) Create(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesCreateOptions) (DataCollectionRuleResourceResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
	if err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return DataCollectionRuleResourceResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *DataCollectionRulesClient) createCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, req.MarshalAsJSON(*options.Body)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *DataCollectionRulesClient) createHandleResponse(resp *azcore.Response) (DataCollectionRuleResourceResponse, error) {
	var val *DataCollectionRuleResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	return DataCollectionRuleResourceResponse{RawResponse: resp.Response, DataCollectionRuleResource: val}, nil
}

// createHandleError handles the Create error response.
func (client *DataCollectionRulesClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes a data collection rule.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionRulesClient) Delete(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
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
func (client *DataCollectionRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DataCollectionRulesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Returns the specified data collection rule.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionRulesClient) Get(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesGetOptions) (DataCollectionRuleResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
	if err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DataCollectionRuleResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataCollectionRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataCollectionRulesClient) getHandleResponse(resp *azcore.Response) (DataCollectionRuleResourceResponse, error) {
	var val *DataCollectionRuleResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	return DataCollectionRuleResourceResponse{RawResponse: resp.Response, DataCollectionRuleResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *DataCollectionRulesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Lists all data collection rules in the specified resource group.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionRulesClient) ListByResourceGroup(resourceGroupName string, options *DataCollectionRulesListByResourceGroupOptions) DataCollectionRuleResourceListResultPager {
	return &dataCollectionRuleResourceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp DataCollectionRuleResourceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DataCollectionRuleResourceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DataCollectionRulesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DataCollectionRulesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DataCollectionRulesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (DataCollectionRuleResourceListResultResponse, error) {
	var val *DataCollectionRuleResourceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataCollectionRuleResourceListResultResponse{}, err
	}
	return DataCollectionRuleResourceListResultResponse{RawResponse: resp.Response, DataCollectionRuleResourceListResult: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DataCollectionRulesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListBySubscription - Lists all data collection rules in the specified subscription.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionRulesClient) ListBySubscription(options *DataCollectionRulesListBySubscriptionOptions) DataCollectionRuleResourceListResultPager {
	return &dataCollectionRuleResourceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.listBySubscriptionHandleResponse,
		errorer:   client.listBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp DataCollectionRuleResourceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DataCollectionRuleResourceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DataCollectionRulesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DataCollectionRulesListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/dataCollectionRules"
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
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DataCollectionRulesClient) listBySubscriptionHandleResponse(resp *azcore.Response) (DataCollectionRuleResourceListResultResponse, error) {
	var val *DataCollectionRuleResourceListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataCollectionRuleResourceListResultResponse{}, err
	}
	return DataCollectionRuleResourceListResultResponse{RawResponse: resp.Response, DataCollectionRuleResourceListResult: val}, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *DataCollectionRulesClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Updates part of a data collection rule.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionRulesClient) Update(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesUpdateOptions) (DataCollectionRuleResourceResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
	if err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DataCollectionRuleResourceResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DataCollectionRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRulesUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, req.MarshalAsJSON(*options.Body)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DataCollectionRulesClient) updateHandleResponse(resp *azcore.Response) (DataCollectionRuleResourceResponse, error) {
	var val *DataCollectionRuleResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataCollectionRuleResourceResponse{}, err
	}
	return DataCollectionRuleResourceResponse{RawResponse: resp.Response, DataCollectionRuleResource: val}, nil
}

// updateHandleError handles the Update error response.
func (client *DataCollectionRulesClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
