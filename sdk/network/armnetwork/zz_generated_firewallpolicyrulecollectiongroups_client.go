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

// FirewallPolicyRuleCollectionGroupsClient contains the methods for the FirewallPolicyRuleCollectionGroups group.
// Don't use this type directly, use NewFirewallPolicyRuleCollectionGroupsClient() instead.
type FirewallPolicyRuleCollectionGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFirewallPolicyRuleCollectionGroupsClient creates a new instance of FirewallPolicyRuleCollectionGroupsClient with the specified values.
func NewFirewallPolicyRuleCollectionGroupsClient(con *armcore.Connection, subscriptionID string) *FirewallPolicyRuleCollectionGroupsClient {
	return &FirewallPolicyRuleCollectionGroupsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleCollectionGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsBeginCreateOrUpdateOptions) (FirewallPolicyRuleCollectionGroupPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupPollerResponse{}, err
	}
	result := FirewallPolicyRuleCollectionGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPolicyRuleCollectionGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupPollerResponse{}, err
	}
	poller := &firewallPolicyRuleCollectionGroupPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FirewallPolicyRuleCollectionGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new FirewallPolicyRuleCollectionGroupPoller from the specified resume token.
// token - The value must come from a previous call to FirewallPolicyRuleCollectionGroupPoller.ResumeToken().
func (client *FirewallPolicyRuleCollectionGroupsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (FirewallPolicyRuleCollectionGroupPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPolicyRuleCollectionGroupsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupPollerResponse{}, err
	}
	poller := &firewallPolicyRuleCollectionGroupPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupPollerResponse{}, err
	}
	result := FirewallPolicyRuleCollectionGroupPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (FirewallPolicyRuleCollectionGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
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
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
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
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FirewallPolicyRuleCollectionGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleCollectionGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPolicyRuleCollectionGroupsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client *FirewallPolicyRuleCollectionGroupsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPolicyRuleCollectionGroupsClient.Delete", token, client.deleteHandleError)
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

// Delete - Deletes the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
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
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
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
func (client *FirewallPolicyRuleCollectionGroupsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified FirewallPolicyRuleCollectionGroup.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleCollectionGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsGetOptions) (FirewallPolicyRuleCollectionGroupResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FirewallPolicyRuleCollectionGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FirewallPolicyRuleCollectionGroupResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FirewallPolicyRuleCollectionGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if ruleCollectionGroupName == "" {
		return nil, errors.New("parameter ruleCollectionGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
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
func (client *FirewallPolicyRuleCollectionGroupsClient) getHandleResponse(resp *azcore.Response) (FirewallPolicyRuleCollectionGroupResponse, error) {
	var val *FirewallPolicyRuleCollectionGroup
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FirewallPolicyRuleCollectionGroupResponse{}, err
	}
	return FirewallPolicyRuleCollectionGroupResponse{RawResponse: resp.Response, FirewallPolicyRuleCollectionGroup: val}, nil
}

// getHandleError handles the Get error response.
func (client *FirewallPolicyRuleCollectionGroupsClient) getHandleError(resp *azcore.Response) error {
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

// List - Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
// If the operation fails it returns the *CloudError error type.
func (client *FirewallPolicyRuleCollectionGroupsClient) List(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsListOptions) FirewallPolicyRuleCollectionGroupListResultPager {
	return &firewallPolicyRuleCollectionGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp FirewallPolicyRuleCollectionGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyRuleCollectionGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *FirewallPolicyRuleCollectionGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
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
func (client *FirewallPolicyRuleCollectionGroupsClient) listHandleResponse(resp *azcore.Response) (FirewallPolicyRuleCollectionGroupListResultResponse, error) {
	var val *FirewallPolicyRuleCollectionGroupListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FirewallPolicyRuleCollectionGroupListResultResponse{}, err
	}
	return FirewallPolicyRuleCollectionGroupListResultResponse{RawResponse: resp.Response, FirewallPolicyRuleCollectionGroupListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *FirewallPolicyRuleCollectionGroupsClient) listHandleError(resp *azcore.Response) error {
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
