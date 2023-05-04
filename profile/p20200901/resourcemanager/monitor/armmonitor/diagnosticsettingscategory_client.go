//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsCategoryClient contains the methods for the DiagnosticSettingsCategory group.
// Don't use this type directly, use NewDiagnosticSettingsCategoryClient() instead.
type DiagnosticSettingsCategoryClient struct {
	internal *arm.Client
}

// NewDiagnosticSettingsCategoryClient creates a new instance of DiagnosticSettingsCategoryClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDiagnosticSettingsCategoryClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*DiagnosticSettingsCategoryClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armmonitor.DiagnosticSettingsCategoryClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DiagnosticSettingsCategoryClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets the diagnostic settings category for the specified resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-05-01-preview
//   - resourceURI - The identifier of the resource.
//   - name - The name of the diagnostic setting.
//   - options - DiagnosticSettingsCategoryClientGetOptions contains the optional parameters for the DiagnosticSettingsCategoryClient.Get
//     method.
func (client *DiagnosticSettingsCategoryClient) Get(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsCategoryClientGetOptions) (DiagnosticSettingsCategoryClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, name, options)
	if err != nil {
		return DiagnosticSettingsCategoryClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsCategoryClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsCategoryClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiagnosticSettingsCategoryClient) getCreateRequest(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsCategoryClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettingsCategories/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiagnosticSettingsCategoryClient) getHandleResponse(resp *http.Response) (DiagnosticSettingsCategoryClientGetResponse, error) {
	result := DiagnosticSettingsCategoryClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsCategoryResource); err != nil {
		return DiagnosticSettingsCategoryClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the diagnostic settings categories for the specified resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-05-01-preview
//   - resourceURI - The identifier of the resource.
//   - options - DiagnosticSettingsCategoryClientListOptions contains the optional parameters for the DiagnosticSettingsCategoryClient.List
//     method.
func (client *DiagnosticSettingsCategoryClient) List(ctx context.Context, resourceURI string, options *DiagnosticSettingsCategoryClientListOptions) (DiagnosticSettingsCategoryClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return DiagnosticSettingsCategoryClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsCategoryClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsCategoryClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DiagnosticSettingsCategoryClient) listCreateRequest(ctx context.Context, resourceURI string, options *DiagnosticSettingsCategoryClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettingsCategories"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiagnosticSettingsCategoryClient) listHandleResponse(resp *http.Response) (DiagnosticSettingsCategoryClientListResponse, error) {
	result := DiagnosticSettingsCategoryClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticSettingsCategoryResourceCollection); err != nil {
		return DiagnosticSettingsCategoryClientListResponse{}, err
	}
	return result, nil
}
