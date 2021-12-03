//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// AutoProvisioningSettingsClient contains the methods for the AutoProvisioningSettings group.
// Don't use this type directly, use NewAutoProvisioningSettingsClient() instead.
type AutoProvisioningSettingsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAutoProvisioningSettingsClient creates a new instance of AutoProvisioningSettingsClient with the specified values.
func NewAutoProvisioningSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AutoProvisioningSettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AutoProvisioningSettingsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Details of a specific setting
// If the operation fails it returns the *CloudError error type.
func (client *AutoProvisioningSettingsClient) Create(ctx context.Context, settingName string, setting AutoProvisioningSetting, options *AutoProvisioningSettingsCreateOptions) (AutoProvisioningSettingsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, settingName, setting, options)
	if err != nil {
		return AutoProvisioningSettingsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoProvisioningSettingsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoProvisioningSettingsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *AutoProvisioningSettingsClient) createCreateRequest(ctx context.Context, settingName string, setting AutoProvisioningSetting, options *AutoProvisioningSettingsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings/{settingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if settingName == "" {
		return nil, errors.New("parameter settingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingName}", url.PathEscape(settingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, setting)
}

// createHandleResponse handles the Create response.
func (client *AutoProvisioningSettingsClient) createHandleResponse(resp *http.Response) (AutoProvisioningSettingsCreateResponse, error) {
	result := AutoProvisioningSettingsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutoProvisioningSetting); err != nil {
		return AutoProvisioningSettingsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *AutoProvisioningSettingsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Details of a specific setting
// If the operation fails it returns the *CloudError error type.
func (client *AutoProvisioningSettingsClient) Get(ctx context.Context, settingName string, options *AutoProvisioningSettingsGetOptions) (AutoProvisioningSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, settingName, options)
	if err != nil {
		return AutoProvisioningSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AutoProvisioningSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AutoProvisioningSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AutoProvisioningSettingsClient) getCreateRequest(ctx context.Context, settingName string, options *AutoProvisioningSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings/{settingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if settingName == "" {
		return nil, errors.New("parameter settingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingName}", url.PathEscape(settingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AutoProvisioningSettingsClient) getHandleResponse(resp *http.Response) (AutoProvisioningSettingsGetResponse, error) {
	result := AutoProvisioningSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutoProvisioningSetting); err != nil {
		return AutoProvisioningSettingsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AutoProvisioningSettingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Exposes the auto provisioning settings of the subscriptions
// If the operation fails it returns the *CloudError error type.
func (client *AutoProvisioningSettingsClient) List(options *AutoProvisioningSettingsListOptions) *AutoProvisioningSettingsListPager {
	return &AutoProvisioningSettingsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp AutoProvisioningSettingsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AutoProvisioningSettingList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AutoProvisioningSettingsClient) listCreateRequest(ctx context.Context, options *AutoProvisioningSettingsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/autoProvisioningSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AutoProvisioningSettingsClient) listHandleResponse(resp *http.Response) (AutoProvisioningSettingsListResponse, error) {
	result := AutoProvisioningSettingsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutoProvisioningSettingList); err != nil {
		return AutoProvisioningSettingsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AutoProvisioningSettingsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
