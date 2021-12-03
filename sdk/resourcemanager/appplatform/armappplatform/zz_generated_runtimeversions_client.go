//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// RuntimeVersionsClient contains the methods for the RuntimeVersions group.
// Don't use this type directly, use NewRuntimeVersionsClient() instead.
type RuntimeVersionsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRuntimeVersionsClient creates a new instance of RuntimeVersionsClient with the specified values.
func NewRuntimeVersionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *RuntimeVersionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RuntimeVersionsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListRuntimeVersions - Lists all of the available runtime versions supported by Microsoft.AppPlatform provider.
// If the operation fails it returns the *CloudError error type.
func (client *RuntimeVersionsClient) ListRuntimeVersions(ctx context.Context, options *RuntimeVersionsListRuntimeVersionsOptions) (RuntimeVersionsListRuntimeVersionsResponse, error) {
	req, err := client.listRuntimeVersionsCreateRequest(ctx, options)
	if err != nil {
		return RuntimeVersionsListRuntimeVersionsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RuntimeVersionsListRuntimeVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RuntimeVersionsListRuntimeVersionsResponse{}, client.listRuntimeVersionsHandleError(resp)
	}
	return client.listRuntimeVersionsHandleResponse(resp)
}

// listRuntimeVersionsCreateRequest creates the ListRuntimeVersions request.
func (client *RuntimeVersionsClient) listRuntimeVersionsCreateRequest(ctx context.Context, options *RuntimeVersionsListRuntimeVersionsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppPlatform/runtimeVersions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listRuntimeVersionsHandleResponse handles the ListRuntimeVersions response.
func (client *RuntimeVersionsClient) listRuntimeVersionsHandleResponse(resp *http.Response) (RuntimeVersionsListRuntimeVersionsResponse, error) {
	result := RuntimeVersionsListRuntimeVersionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableRuntimeVersions); err != nil {
		return RuntimeVersionsListRuntimeVersionsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listRuntimeVersionsHandleError handles the ListRuntimeVersions error response.
func (client *RuntimeVersionsClient) listRuntimeVersionsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
