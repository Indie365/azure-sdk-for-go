// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ServersClient contains the methods for the Servers group.
// Don't use this type directly, use NewServersClient() instead.
type ServersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewServersClient creates a new instance of ServersClient with the specified values.
func NewServersClient(con *armcore.Connection, subscriptionID string) *ServersClient {
	return &ServersClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreate - Creates a new server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) BeginCreate(ctx context.Context, resourceGroupName string, serverName string, parameters Server, options *ServersBeginCreateOptions) (ServersCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return ServersCreatePollerResponse{}, err
	}
	result := ServersCreatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServersClient.Create", "", resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return ServersCreatePollerResponse{}, err
	}
	poller := &serversCreatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersCreateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreate creates a new ServersCreatePoller from the specified resume token.
// token - The value must come from a previous call to ServersCreatePoller.ResumeToken().
func (client *ServersClient) ResumeCreate(ctx context.Context, token string) (ServersCreatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServersClient.Create", token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return ServersCreatePollerResponse{}, err
	}
	poller := &serversCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServersCreatePollerResponse{}, err
	}
	result := ServersCreatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersCreateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Create - Creates a new server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) create(ctx context.Context, resourceGroupName string, serverName string, parameters Server, options *ServersBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ServersClient) createCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters Server, options *ServersBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createHandleError handles the Create error response.
func (client *ServersClient) createHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginDeleteOptions) (ServersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServersDeletePollerResponse{}, err
	}
	result := ServersDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServersClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServersDeletePollerResponse{}, err
	}
	poller := &serversDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new ServersDeletePoller from the specified resume token.
// token - The value must come from a previous call to ServersDeletePoller.ResumeToken().
func (client *ServersClient) ResumeDelete(ctx context.Context, token string) (ServersDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServersClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ServersDeletePollerResponse{}, err
	}
	poller := &serversDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServersDeletePollerResponse{}, err
	}
	result := ServersDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServersClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets information about a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServersGetOptions) (ServersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServersGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ServersGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ServersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServersClient) getHandleResponse(resp *azcore.Response) (ServersGetResponse, error) {
	result := ServersGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Server); err != nil {
		return ServersGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServersClient) getHandleError(resp *azcore.Response) error {
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

// List - List all the servers in a given subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) List(options *ServersListOptions) ServersListPager {
	return &serversListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ServersListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServersClient) listCreateRequest(ctx context.Context, options *ServersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/flexibleServers"
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
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServersClient) listHandleResponse(resp *azcore.Response) (ServersListResponse, error) {
	result := ServersListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerListResult); err != nil {
		return ServersListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ServersClient) listHandleError(resp *azcore.Response) error {
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

// ListByResourceGroup - List all the servers in a given resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) ListByResourceGroup(resourceGroupName string, options *ServersListByResourceGroupOptions) ServersListByResourceGroupPager {
	return &serversListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ServersListByResourceGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ServerListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServersListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers"
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
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServersClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ServersListByResourceGroupResponse, error) {
	result := ServersListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ServerListResult); err != nil {
		return ServersListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ServersClient) listByResourceGroupHandleError(resp *azcore.Response) error {
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

// BeginRestart - Restarts a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) BeginRestart(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginRestartOptions) (ServersRestartPollerResponse, error) {
	resp, err := client.restart(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServersRestartPollerResponse{}, err
	}
	result := ServersRestartPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServersClient.Restart", "", resp, client.con.Pipeline(), client.restartHandleError)
	if err != nil {
		return ServersRestartPollerResponse{}, err
	}
	poller := &serversRestartPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersRestartResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRestart creates a new ServersRestartPoller from the specified resume token.
// token - The value must come from a previous call to ServersRestartPoller.ResumeToken().
func (client *ServersClient) ResumeRestart(ctx context.Context, token string) (ServersRestartPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServersClient.Restart", token, client.con.Pipeline(), client.restartHandleError)
	if err != nil {
		return ServersRestartPollerResponse{}, err
	}
	poller := &serversRestartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServersRestartPollerResponse{}, err
	}
	result := ServersRestartPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersRestartResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Restart - Restarts a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) restart(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginRestartOptions) (*azcore.Response, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.restartHandleError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ServersClient) restartCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginRestartOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, req.MarshalAsJSON(*options.Parameters)
	}
	return req, nil
}

// restartHandleError handles the Restart error response.
func (client *ServersClient) restartHandleError(resp *azcore.Response) error {
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

// BeginStart - Starts a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) BeginStart(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginStartOptions) (ServersStartPollerResponse, error) {
	resp, err := client.start(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServersStartPollerResponse{}, err
	}
	result := ServersStartPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServersClient.Start", "", resp, client.con.Pipeline(), client.startHandleError)
	if err != nil {
		return ServersStartPollerResponse{}, err
	}
	poller := &serversStartPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersStartResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStart creates a new ServersStartPoller from the specified resume token.
// token - The value must come from a previous call to ServersStartPoller.ResumeToken().
func (client *ServersClient) ResumeStart(ctx context.Context, token string) (ServersStartPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServersClient.Start", token, client.con.Pipeline(), client.startHandleError)
	if err != nil {
		return ServersStartPollerResponse{}, err
	}
	poller := &serversStartPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServersStartPollerResponse{}, err
	}
	result := ServersStartPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersStartResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Start - Starts a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) start(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginStartOptions) (*azcore.Response, error) {
	req, err := client.startCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.startHandleError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *ServersClient) startCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginStartOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// startHandleError handles the Start error response.
func (client *ServersClient) startHandleError(resp *azcore.Response) error {
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

// BeginStop - Stops a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) BeginStop(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginStopOptions) (ServersStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServersStopPollerResponse{}, err
	}
	result := ServersStopPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServersClient.Stop", "", resp, client.con.Pipeline(), client.stopHandleError)
	if err != nil {
		return ServersStopPollerResponse{}, err
	}
	poller := &serversStopPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersStopResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStop creates a new ServersStopPoller from the specified resume token.
// token - The value must come from a previous call to ServersStopPoller.ResumeToken().
func (client *ServersClient) ResumeStop(ctx context.Context, token string) (ServersStopPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServersClient.Stop", token, client.con.Pipeline(), client.stopHandleError)
	if err != nil {
		return ServersStopPollerResponse{}, err
	}
	poller := &serversStopPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServersStopPollerResponse{}, err
	}
	result := ServersStopPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersStopResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Stop - Stops a server.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) stop(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginStopOptions) (*azcore.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.stopHandleError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *ServersClient) stopCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServersBeginStopOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// stopHandleError handles the Stop error response.
func (client *ServersClient) stopHandleError(resp *azcore.Response) error {
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

// BeginUpdate - Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, parameters ServerForUpdate, options *ServersBeginUpdateOptions) (ServersUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return ServersUpdatePollerResponse{}, err
	}
	result := ServersUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ServersClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ServersUpdatePollerResponse{}, err
	}
	poller := &serversUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ServersUpdatePoller from the specified resume token.
// token - The value must come from a previous call to ServersUpdatePoller.ResumeToken().
func (client *ServersClient) ResumeUpdate(ctx context.Context, token string) (ServersUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ServersClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ServersUpdatePollerResponse{}, err
	}
	poller := &serversUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ServersUpdatePollerResponse{}, err
	}
	result := ServersUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ServersUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
// If the operation fails it returns the *CloudError error type.
func (client *ServersClient) update(ctx context.Context, resourceGroupName string, serverName string, parameters ServerForUpdate, options *ServersBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ServersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters ServerForUpdate, options *ServersBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleError handles the Update error response.
func (client *ServersClient) updateHandleError(resp *azcore.Response) error {
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
