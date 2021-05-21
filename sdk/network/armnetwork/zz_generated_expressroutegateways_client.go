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
	"time"
)

// ExpressRouteGatewaysClient contains the methods for the ExpressRouteGateways group.
// Don't use this type directly, use NewExpressRouteGatewaysClient() instead.
type ExpressRouteGatewaysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteGatewaysClient creates a new instance of ExpressRouteGatewaysClient with the specified values.
func NewExpressRouteGatewaysClient(con *armcore.Connection, subscriptionID string) *ExpressRouteGatewaysClient {
	return &ExpressRouteGatewaysClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a ExpressRoute gateway in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway, options *ExpressRouteGatewaysBeginCreateOrUpdateOptions) (ExpressRouteGatewayPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, expressRouteGatewayName, putExpressRouteGatewayParameters, options)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	result := ExpressRouteGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	poller := &expressRouteGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteGatewayPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteGatewayPoller.ResumeToken().
func (client *ExpressRouteGatewaysClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ExpressRouteGatewayPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteGatewaysClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	poller := &expressRouteGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	result := ExpressRouteGatewayPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a ExpressRoute gateway in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway, options *ExpressRouteGatewaysBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, putExpressRouteGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway, options *ExpressRouteGatewaysBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(putExpressRouteGatewayParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteGatewaysClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified ExpressRoute gateway in a resource group. An ExpressRoute gateway resource can only be deleted when there are no
// connection subresources.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteGatewaysClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ExpressRouteGatewaysClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteGatewaysClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified ExpressRoute gateway in a resource group. An ExpressRoute gateway resource can only be deleted when there are no connection
// subresources.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, options)
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
func (client *ExpressRouteGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ExpressRouteGatewaysClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Fetches the details of a ExpressRoute gateway in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysGetOptions) (ExpressRouteGatewayResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return ExpressRouteGatewayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteGatewayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteGatewayResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
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
func (client *ExpressRouteGatewaysClient) getHandleResponse(resp *azcore.Response) (ExpressRouteGatewayResponse, error) {
	var val *ExpressRouteGateway
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteGatewayResponse{}, err
	}
	return ExpressRouteGatewayResponse{RawResponse: resp.Response, ExpressRouteGateway: val}, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteGatewaysClient) getHandleError(resp *azcore.Response) error {
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

// ListByResourceGroup - Lists ExpressRoute gateways in a given resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ExpressRouteGatewaysListByResourceGroupOptions) (ExpressRouteGatewayListResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ExpressRouteGatewayListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteGatewayListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteGatewayListResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRouteGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteGatewaysListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways"
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
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRouteGatewaysClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ExpressRouteGatewayListResponse, error) {
	var val *ExpressRouteGatewayList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteGatewayListResponse{}, err
	}
	return ExpressRouteGatewayListResponse{RawResponse: resp.Response, ExpressRouteGatewayList: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRouteGatewaysClient) listByResourceGroupHandleError(resp *azcore.Response) error {
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

// ListBySubscription - Lists ExpressRoute gateways under a given subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) ListBySubscription(ctx context.Context, options *ExpressRouteGatewaysListBySubscriptionOptions) (ExpressRouteGatewayListResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ExpressRouteGatewayListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteGatewayListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRouteGatewayListResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ExpressRouteGatewaysClient) listBySubscriptionCreateRequest(ctx context.Context, options *ExpressRouteGatewaysListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways"
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ExpressRouteGatewaysClient) listBySubscriptionHandleResponse(resp *azcore.Response) (ExpressRouteGatewayListResponse, error) {
	var val *ExpressRouteGatewayList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ExpressRouteGatewayListResponse{}, err
	}
	return ExpressRouteGatewayListResponse{RawResponse: resp.Response, ExpressRouteGatewayList: val}, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ExpressRouteGatewaysClient) listBySubscriptionHandleError(resp *azcore.Response) error {
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

// BeginUpdateTags - Updates express route gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject, options *ExpressRouteGatewaysBeginUpdateTagsOptions) (ExpressRouteGatewayPollerResponse, error) {
	resp, err := client.updateTags(ctx, resourceGroupName, expressRouteGatewayName, expressRouteGatewayParameters, options)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	result := ExpressRouteGatewayPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteGatewaysClient.UpdateTags", "azure-async-operation", resp, client.updateTagsHandleError)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	poller := &expressRouteGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdateTags creates a new ExpressRouteGatewayPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteGatewayPoller.ResumeToken().
func (client *ExpressRouteGatewaysClient) ResumeUpdateTags(ctx context.Context, token string) (ExpressRouteGatewayPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteGatewaysClient.UpdateTags", token, client.updateTagsHandleError)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	poller := &expressRouteGatewayPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ExpressRouteGatewayPollerResponse{}, err
	}
	result := ExpressRouteGatewayPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ExpressRouteGatewayResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// UpdateTags - Updates express route gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) updateTags(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject, options *ExpressRouteGatewaysBeginUpdateTagsOptions) (*azcore.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, expressRouteGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateTagsHandleError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRouteGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject, options *ExpressRouteGatewaysBeginUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(expressRouteGatewayParameters)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRouteGatewaysClient) updateTagsHandleError(resp *azcore.Response) error {
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
