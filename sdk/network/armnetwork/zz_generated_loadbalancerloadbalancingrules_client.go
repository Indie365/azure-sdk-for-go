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

// LoadBalancerLoadBalancingRulesClient contains the methods for the LoadBalancerLoadBalancingRules group.
// Don't use this type directly, use NewLoadBalancerLoadBalancingRulesClient() instead.
type LoadBalancerLoadBalancingRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerLoadBalancingRulesClient creates a new instance of LoadBalancerLoadBalancingRulesClient with the specified values.
func NewLoadBalancerLoadBalancingRulesClient(con *armcore.Connection, subscriptionID string) *LoadBalancerLoadBalancingRulesClient {
	return &LoadBalancerLoadBalancingRulesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the specified load balancer load balancing rule.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerLoadBalancingRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesGetOptions) (LoadBalancingRuleResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, loadBalancingRuleName, options)
	if err != nil {
		return LoadBalancingRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoadBalancingRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoadBalancingRuleResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerLoadBalancingRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, loadBalancingRuleName string, options *LoadBalancerLoadBalancingRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if loadBalancingRuleName == "" {
		return nil, errors.New("parameter loadBalancingRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancingRuleName}", url.PathEscape(loadBalancingRuleName))
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
func (client *LoadBalancerLoadBalancingRulesClient) getHandleResponse(resp *azcore.Response) (LoadBalancingRuleResponse, error) {
	var val *LoadBalancingRule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancingRuleResponse{}, err
	}
	return LoadBalancingRuleResponse{RawResponse: resp.Response, LoadBalancingRule: val}, nil
}

// getHandleError handles the Get error response.
func (client *LoadBalancerLoadBalancingRulesClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets all the load balancing rules in a load balancer.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerLoadBalancingRulesClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesListOptions) LoadBalancerLoadBalancingRuleListResultPager {
	return &loadBalancerLoadBalancingRuleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp LoadBalancerLoadBalancingRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerLoadBalancingRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancerLoadBalancingRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerLoadBalancingRulesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
func (client *LoadBalancerLoadBalancingRulesClient) listHandleResponse(resp *azcore.Response) (LoadBalancerLoadBalancingRuleListResultResponse, error) {
	var val *LoadBalancerLoadBalancingRuleListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerLoadBalancingRuleListResultResponse{}, err
	}
	return LoadBalancerLoadBalancingRuleListResultResponse{RawResponse: resp.Response, LoadBalancerLoadBalancingRuleListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancerLoadBalancingRulesClient) listHandleError(resp *azcore.Response) error {
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
