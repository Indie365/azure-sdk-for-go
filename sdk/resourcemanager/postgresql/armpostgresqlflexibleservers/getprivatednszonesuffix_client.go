//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// GetPrivateDNSZoneSuffixClient contains the methods for the GetPrivateDNSZoneSuffix group.
// Don't use this type directly, use NewGetPrivateDNSZoneSuffixClient() instead.
type GetPrivateDNSZoneSuffixClient struct {
	internal *arm.Client
}

// NewGetPrivateDNSZoneSuffixClient creates a new instance of GetPrivateDNSZoneSuffixClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGetPrivateDNSZoneSuffixClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*GetPrivateDNSZoneSuffixClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GetPrivateDNSZoneSuffixClient{
		internal: cl,
	}
	return client, nil
}

// Execute - Get private DNS zone suffix in the cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - options - GetPrivateDNSZoneSuffixClientExecuteOptions contains the optional parameters for the GetPrivateDNSZoneSuffixClient.Execute
//     method.
func (client *GetPrivateDNSZoneSuffixClient) Execute(ctx context.Context, options *GetPrivateDNSZoneSuffixClientExecuteOptions) (GetPrivateDNSZoneSuffixClientExecuteResponse, error) {
	var err error
	const operationName = "GetPrivateDNSZoneSuffixClient.Execute"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.executeCreateRequest(ctx, options)
	if err != nil {
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, err
	}
	resp, err := client.executeHandleResponse(httpResp)
	return resp, err
}

// executeCreateRequest creates the Execute request.
func (client *GetPrivateDNSZoneSuffixClient) executeCreateRequest(ctx context.Context, options *GetPrivateDNSZoneSuffixClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DBforPostgreSQL/getPrivateDnsZoneSuffix"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *GetPrivateDNSZoneSuffixClient) executeHandleResponse(resp *http.Response) (GetPrivateDNSZoneSuffixClientExecuteResponse, error) {
	result := GetPrivateDNSZoneSuffixClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return GetPrivateDNSZoneSuffixClientExecuteResponse{}, err
	}
	return result, nil
}
