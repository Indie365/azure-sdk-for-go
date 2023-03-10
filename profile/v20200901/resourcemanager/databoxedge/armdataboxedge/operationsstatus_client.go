//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OperationsStatusClient contains the methods for the OperationsStatus group.
// Don't use this type directly, use NewOperationsStatusClient() instead.
type OperationsStatusClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationsStatusClient creates a new instance of OperationsStatusClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsStatusClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+".OperationsStatusClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsStatusClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the details of a specified job on a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - deviceName - The device name.
//   - name - The job name.
//   - resourceGroupName - The resource group name.
//   - options - OperationsStatusClientGetOptions contains the optional parameters for the OperationsStatusClient.Get method.
func (client *OperationsStatusClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *OperationsStatusClientGetOptions) (OperationsStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return OperationsStatusClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OperationsStatusClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *OperationsStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/operationsStatus/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationsStatusClient) getHandleResponse(resp *http.Response) (OperationsStatusClientGetResponse, error) {
	result := OperationsStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return OperationsStatusClientGetResponse{}, err
	}
	return result, nil
}
