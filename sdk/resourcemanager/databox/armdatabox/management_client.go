//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabox

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagementClient contains the methods for the DataBoxManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - The Subscription Id
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagementClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// Mitigate - Request to mitigate for a given job
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-12-01
//   - jobName - The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters
//     in length and use any alphanumeric and underscore only
//   - resourceGroupName - The Resource Group Name
//   - mitigateJobRequest - Mitigation Request
//   - options - ManagementClientMitigateOptions contains the optional parameters for the ManagementClient.Mitigate method.
func (client *ManagementClient) Mitigate(ctx context.Context, jobName string, resourceGroupName string, mitigateJobRequest MitigateJobRequest, options *ManagementClientMitigateOptions) (ManagementClientMitigateResponse, error) {
	var err error
	req, err := client.mitigateCreateRequest(ctx, jobName, resourceGroupName, mitigateJobRequest, options)
	if err != nil {
		return ManagementClientMitigateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientMitigateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ManagementClientMitigateResponse{}, err
	}
	return ManagementClientMitigateResponse{}, nil
}

// mitigateCreateRequest creates the Mitigate request.
func (client *ManagementClient) mitigateCreateRequest(ctx context.Context, jobName string, resourceGroupName string, mitigateJobRequest MitigateJobRequest, options *ManagementClientMitigateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBox/jobs/{jobName}/mitigate"
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, mitigateJobRequest); err != nil {
	return nil, err
}
	return req, nil
}

