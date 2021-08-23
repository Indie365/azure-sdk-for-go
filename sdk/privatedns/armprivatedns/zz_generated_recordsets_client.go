// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armprivatedns

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

// RecordSetsClient contains the methods for the RecordSets group.
// Don't use this type directly, use NewRecordSetsClient() instead.
type RecordSetsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRecordSetsClient creates a new instance of RecordSetsClient with the specified values.
func NewRecordSetsClient(con *armcore.Connection, subscriptionID string) *RecordSetsClient {
	return &RecordSetsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a record set within a Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *RecordSetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsCreateOrUpdateOptions) (RecordSetResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, parameters, options)
	if err != nil {
		return RecordSetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RecordSetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return RecordSetResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RecordSetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if relativeRecordSetName == "" {
		return nil, errors.New("parameter relativeRecordSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
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
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RecordSetsClient) createOrUpdateHandleResponse(resp *azcore.Response) (RecordSetResponse, error) {
	var val *RecordSet
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RecordSetResponse{}, err
	}
	return RecordSetResponse{RawResponse: resp.Response, RecordSet: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *RecordSetsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// Delete - Deletes a record set from a Private DNS zone. This operation cannot be undone.
// If the operation fails it returns the *CloudError error type.
func (client *RecordSetsClient) Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, options)
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
func (client *RecordSetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if relativeRecordSetName == "" {
		return nil, errors.New("parameter relativeRecordSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
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
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *RecordSetsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets a record set.
// If the operation fails it returns the *CloudError error type.
func (client *RecordSetsClient) Get(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsGetOptions) (RecordSetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, options)
	if err != nil {
		return RecordSetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RecordSetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RecordSetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecordSetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, options *RecordSetsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if relativeRecordSetName == "" {
		return nil, errors.New("parameter relativeRecordSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
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
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecordSetsClient) getHandleResponse(resp *azcore.Response) (RecordSetResponse, error) {
	var val *RecordSet
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RecordSetResponse{}, err
	}
	return RecordSetResponse{RawResponse: resp.Response, RecordSet: val}, nil
}

// getHandleError handles the Get error response.
func (client *RecordSetsClient) getHandleError(resp *azcore.Response) error {
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

// List - Lists all record sets in a Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *RecordSetsClient) List(resourceGroupName string, privateZoneName string, options *RecordSetsListOptions) RecordSetListResultPager {
	return &recordSetListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, privateZoneName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp RecordSetListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RecordSetListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *RecordSetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, options *RecordSetsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/ALL"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Recordsetnamesuffix != nil {
		reqQP.Set("$recordsetnamesuffix", *options.Recordsetnamesuffix)
	}
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecordSetsClient) listHandleResponse(resp *azcore.Response) (RecordSetListResultResponse, error) {
	var val *RecordSetListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RecordSetListResultResponse{}, err
	}
	return RecordSetListResultResponse{RawResponse: resp.Response, RecordSetListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *RecordSetsClient) listHandleError(resp *azcore.Response) error {
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

// ListByType - Lists the record sets of a specified type in a Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *RecordSetsClient) ListByType(resourceGroupName string, privateZoneName string, recordType RecordType, options *RecordSetsListByTypeOptions) RecordSetListResultPager {
	return &recordSetListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByTypeCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, options)
		},
		responder: client.listByTypeHandleResponse,
		errorer:   client.listByTypeHandleError,
		advancer: func(ctx context.Context, resp RecordSetListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RecordSetListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByTypeCreateRequest creates the ListByType request.
func (client *RecordSetsClient) listByTypeCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, options *RecordSetsListByTypeOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
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
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Recordsetnamesuffix != nil {
		reqQP.Set("$recordsetnamesuffix", *options.Recordsetnamesuffix)
	}
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByTypeHandleResponse handles the ListByType response.
func (client *RecordSetsClient) listByTypeHandleResponse(resp *azcore.Response) (RecordSetListResultResponse, error) {
	var val *RecordSetListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RecordSetListResultResponse{}, err
	}
	return RecordSetListResultResponse{RawResponse: resp.Response, RecordSetListResult: val}, nil
}

// listByTypeHandleError handles the ListByType error response.
func (client *RecordSetsClient) listByTypeHandleError(resp *azcore.Response) error {
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

// Update - Updates a record set within a Private DNS zone.
// If the operation fails it returns the *CloudError error type.
func (client *RecordSetsClient) Update(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsUpdateOptions) (RecordSetResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateZoneName, recordType, relativeRecordSetName, parameters, options)
	if err != nil {
		return RecordSetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RecordSetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RecordSetResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RecordSetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, options *RecordSetsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if recordType == "" {
		return nil, errors.New("parameter recordType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recordType}", url.PathEscape(string(recordType)))
	if relativeRecordSetName == "" {
		return nil, errors.New("parameter relativeRecordSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relativeRecordSetName}", relativeRecordSetName)
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
	reqQP.Set("api-version", "2020-06-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *RecordSetsClient) updateHandleResponse(resp *azcore.Response) (RecordSetResponse, error) {
	var val *RecordSet
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RecordSetResponse{}, err
	}
	return RecordSetResponse{RawResponse: resp.Response, RecordSet: val}, nil
}

// updateHandleError handles the Update error response.
func (client *RecordSetsClient) updateHandleError(resp *azcore.Response) error {
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
