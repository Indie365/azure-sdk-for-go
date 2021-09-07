//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AccessReviewDefaultSettingsClient contains the methods for the AccessReviewDefaultSettings group.
// Don't use this type directly, use NewAccessReviewDefaultSettingsClient() instead.
type AccessReviewDefaultSettingsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAccessReviewDefaultSettingsClient creates a new instance of AccessReviewDefaultSettingsClient with the specified values.
func NewAccessReviewDefaultSettingsClient(con *arm.Connection, subscriptionID string) *AccessReviewDefaultSettingsClient {
	return &AccessReviewDefaultSettingsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Get access review default settings for the subscription
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewDefaultSettingsClient) Get(ctx context.Context, options *AccessReviewDefaultSettingsGetOptions) (AccessReviewDefaultSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return AccessReviewDefaultSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewDefaultSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewDefaultSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AccessReviewDefaultSettingsClient) getCreateRequest(ctx context.Context, options *AccessReviewDefaultSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessReviewDefaultSettingsClient) getHandleResponse(resp *http.Response) (AccessReviewDefaultSettingsGetResponse, error) {
	result := AccessReviewDefaultSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDefaultSettings); err != nil {
		return AccessReviewDefaultSettingsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AccessReviewDefaultSettingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Put - Get access review default settings for the subscription
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewDefaultSettingsClient) Put(ctx context.Context, properties AccessReviewScheduleSettings, options *AccessReviewDefaultSettingsPutOptions) (AccessReviewDefaultSettingsPutResponse, error) {
	req, err := client.putCreateRequest(ctx, properties, options)
	if err != nil {
		return AccessReviewDefaultSettingsPutResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewDefaultSettingsPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewDefaultSettingsPutResponse{}, client.putHandleError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *AccessReviewDefaultSettingsClient) putCreateRequest(ctx context.Context, properties AccessReviewScheduleSettings, options *AccessReviewDefaultSettingsPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// putHandleResponse handles the Put response.
func (client *AccessReviewDefaultSettingsClient) putHandleResponse(resp *http.Response) (AccessReviewDefaultSettingsPutResponse, error) {
	result := AccessReviewDefaultSettingsPutResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDefaultSettings); err != nil {
		return AccessReviewDefaultSettingsPutResponse{}, err
	}
	return result, nil
}

// putHandleError handles the Put error response.
func (client *AccessReviewDefaultSettingsClient) putHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
