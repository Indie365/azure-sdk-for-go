//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DscNodeClient contains the methods for the DscNode group.
// Don't use this type directly, use NewDscNodeClient() instead.
type DscNodeClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDscNodeClient creates a new instance of DscNodeClient with the specified values.
func NewDscNodeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DscNodeClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DscNodeClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Delete - Delete the dsc node identified by node id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscNodeClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeDeleteOptions) (DscNodeDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, options)
	if err != nil {
		return DscNodeDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscNodeDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscNodeDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DscNodeDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DscNodeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DscNodeClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieve the dsc node identified by node id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscNodeClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeGetOptions) (DscNodeGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, options)
	if err != nil {
		return DscNodeGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscNodeGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscNodeGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DscNodeClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *DscNodeGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DscNodeClient) getHandleResponse(resp *http.Response) (DscNodeGetResponse, error) {
	result := DscNodeGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNode); err != nil {
		return DscNodeGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DscNodeClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByAutomationAccount - Retrieve a list of dsc nodes.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscNodeClient) ListByAutomationAccount(resourceGroupName string, automationAccountName string, options *DscNodeListByAutomationAccountOptions) *DscNodeListByAutomationAccountPager {
	return &DscNodeListByAutomationAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
		},
		advancer: func(ctx context.Context, resp DscNodeListByAutomationAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DscNodeListResult.NextLink)
		},
	}
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *DscNodeClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *DscNodeListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Inlinecount != nil {
		reqQP.Set("$inlinecount", *options.Inlinecount)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *DscNodeClient) listByAutomationAccountHandleResponse(resp *http.Response) (DscNodeListByAutomationAccountResponse, error) {
	result := DscNodeListByAutomationAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNodeListResult); err != nil {
		return DscNodeListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *DscNodeClient) listByAutomationAccountHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Update the dsc node.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DscNodeClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, dscNodeUpdateParameters DscNodeUpdateParameters, options *DscNodeUpdateOptions) (DscNodeUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, dscNodeUpdateParameters, options)
	if err != nil {
		return DscNodeUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DscNodeUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DscNodeUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DscNodeClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, dscNodeUpdateParameters DscNodeUpdateParameters, options *DscNodeUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dscNodeUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *DscNodeClient) updateHandleResponse(resp *http.Response) (DscNodeUpdateResponse, error) {
	result := DscNodeUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNode); err != nil {
		return DscNodeUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *DscNodeClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
