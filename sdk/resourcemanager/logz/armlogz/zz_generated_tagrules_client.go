//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

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
	"strings"
)

// TagRulesClient contains the methods for the TagRules group.
// Don't use this type directly, use NewTagRulesClient() instead.
type TagRulesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTagRulesClient creates a new instance of TagRulesClient with the specified values.
func NewTagRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TagRulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TagRulesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Create or update a tag rule set for a given monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TagRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesCreateOrUpdateOptions) (TagRulesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagRulesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagRulesClient) createOrUpdateHandleResponse(resp *http.Response) (TagRulesCreateOrUpdateResponse, error) {
	result := TagRulesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return TagRulesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *TagRulesClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete a tag rule set for a given monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TagRulesClient) Delete(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesDeleteOptions) (TagRulesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagRulesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return TagRulesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *TagRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *TagRulesClient) deleteHandleResponse(resp *http.Response) (TagRulesDeleteResponse, error) {
	result := TagRulesDeleteResponse{RawResponse: resp}
	if val := resp.Header.Get("location"); val != "" {
		result.Location = &val
	}
	return result, nil
}

// deleteHandleError handles the Delete error response.
func (client *TagRulesClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get a tag rule set for a given monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TagRulesClient) Get(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesGetOptions) (TagRulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagRulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TagRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TagRulesClient) getHandleResponse(resp *http.Response) (TagRulesGetResponse, error) {
	result := TagRulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRules); err != nil {
		return TagRulesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TagRulesClient) getHandleError(resp *http.Response) error {
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

// List - List the tag rules for a given monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TagRulesClient) List(resourceGroupName string, monitorName string, options *TagRulesListOptions) *TagRulesListPager {
	return &TagRulesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp TagRulesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MonitoringTagRulesListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *TagRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *TagRulesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logz/monitors/{monitorName}/tagRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TagRulesClient) listHandleResponse(resp *http.Response) (TagRulesListResponse, error) {
	result := TagRulesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringTagRulesListResponse); err != nil {
		return TagRulesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *TagRulesClient) listHandleError(resp *http.Response) error {
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
