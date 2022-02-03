//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagementGroupSubscriptionsClient contains the methods for the ManagementGroupSubscriptions group.
// Don't use this type directly, use NewManagementGroupSubscriptionsClient() instead.
type ManagementGroupSubscriptionsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewManagementGroupSubscriptionsClient creates a new instance of ManagementGroupSubscriptionsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagementGroupSubscriptionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ManagementGroupSubscriptionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ManagementGroupSubscriptionsClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Create - Associates existing subscription with the management group.
// If the operation fails it returns an *azcore.ResponseError type.
// groupID - Management Group ID.
// subscriptionID - Subscription ID.
// options - ManagementGroupSubscriptionsClientCreateOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.Create
// method.
func (client *ManagementGroupSubscriptionsClient) Create(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientCreateOptions) (ManagementGroupSubscriptionsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, groupID, subscriptionID, options)
	if err != nil {
		return ManagementGroupSubscriptionsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementGroupSubscriptionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementGroupSubscriptionsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ManagementGroupSubscriptionsClient) createCreateRequest(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.CacheControl != nil {
		req.Raw().Header.Set("Cache-Control", *options.CacheControl)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ManagementGroupSubscriptionsClient) createHandleResponse(resp *http.Response) (ManagementGroupSubscriptionsClientCreateResponse, error) {
	result := ManagementGroupSubscriptionsClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUnderManagementGroup); err != nil {
		return ManagementGroupSubscriptionsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - De-associates subscription from the management group.
// If the operation fails it returns an *azcore.ResponseError type.
// groupID - Management Group ID.
// subscriptionID - Subscription ID.
// options - ManagementGroupSubscriptionsClientDeleteOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.Delete
// method.
func (client *ManagementGroupSubscriptionsClient) Delete(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientDeleteOptions) (ManagementGroupSubscriptionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, groupID, subscriptionID, options)
	if err != nil {
		return ManagementGroupSubscriptionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementGroupSubscriptionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ManagementGroupSubscriptionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ManagementGroupSubscriptionsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagementGroupSubscriptionsClient) deleteCreateRequest(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.CacheControl != nil {
		req.Raw().Header.Set("Cache-Control", *options.CacheControl)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetSubscription - Retrieves details about given subscription which is associated with the management group.
// If the operation fails it returns an *azcore.ResponseError type.
// groupID - Management Group ID.
// subscriptionID - Subscription ID.
// options - ManagementGroupSubscriptionsClientGetSubscriptionOptions contains the optional parameters for the ManagementGroupSubscriptionsClient.GetSubscription
// method.
func (client *ManagementGroupSubscriptionsClient) GetSubscription(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientGetSubscriptionOptions) (ManagementGroupSubscriptionsClientGetSubscriptionResponse, error) {
	req, err := client.getSubscriptionCreateRequest(ctx, groupID, subscriptionID, options)
	if err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSubscriptionHandleResponse(resp)
}

// getSubscriptionCreateRequest creates the GetSubscription request.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionCreateRequest(ctx context.Context, groupID string, subscriptionID string, options *ManagementGroupSubscriptionsClientGetSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.CacheControl != nil {
		req.Raw().Header.Set("Cache-Control", *options.CacheControl)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSubscriptionHandleResponse handles the GetSubscription response.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionHandleResponse(resp *http.Response) (ManagementGroupSubscriptionsClientGetSubscriptionResponse, error) {
	result := ManagementGroupSubscriptionsClientGetSubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUnderManagementGroup); err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionResponse{}, err
	}
	return result, nil
}

// GetSubscriptionsUnderManagementGroup - Retrieves details about all subscriptions which are associated with the management
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// groupID - Management Group ID.
// options - ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions contains the optional parameters
// for the ManagementGroupSubscriptionsClient.GetSubscriptionsUnderManagementGroup method.
func (client *ManagementGroupSubscriptionsClient) GetSubscriptionsUnderManagementGroup(groupID string, options *ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions) *ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupPager {
	return &ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getSubscriptionsUnderManagementGroupCreateRequest(ctx, groupID, options)
		},
		advancer: func(ctx context.Context, resp ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListSubscriptionUnderManagementGroup.NextLink)
		},
	}
}

// getSubscriptionsUnderManagementGroupCreateRequest creates the GetSubscriptionsUnderManagementGroup request.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionsUnderManagementGroupCreateRequest(ctx context.Context, groupID string, options *ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions"
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSubscriptionsUnderManagementGroupHandleResponse handles the GetSubscriptionsUnderManagementGroup response.
func (client *ManagementGroupSubscriptionsClient) getSubscriptionsUnderManagementGroupHandleResponse(resp *http.Response) (ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse, error) {
	result := ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListSubscriptionUnderManagementGroup); err != nil {
		return ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse{}, err
	}
	return result, nil
}
