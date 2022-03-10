//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceupdate

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

// InstancesClient contains the methods for the Instances group.
// Don't use this type directly, use NewInstancesClient() instead.
type InstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInstancesClient creates a new instance of InstancesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *InstancesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &InstancesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Creates or updates instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - Account name.
// instanceName - Instance name.
// instance - Instance details.
// options - InstancesClientBeginCreateOptions contains the optional parameters for the InstancesClient.BeginCreate method.
func (client *InstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance, options *InstancesClientBeginCreateOptions) (InstancesClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, accountName, instanceName, instance, options)
	if err != nil {
		return InstancesClientCreatePollerResponse{}, err
	}
	result := InstancesClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstancesClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return InstancesClientCreatePollerResponse{}, err
	}
	result.Poller = &InstancesClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates or updates instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstancesClient) create(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance, options *InstancesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, instanceName, instance, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *InstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance Instance, options *InstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, instance)
}

// BeginDelete - Deletes instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - Account name.
// instanceName - Instance name.
// options - InstancesClientBeginDeleteOptions contains the optional parameters for the InstancesClient.BeginDelete method.
func (client *InstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientBeginDeleteOptions) (InstancesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, instanceName, options)
	if err != nil {
		return InstancesClientDeletePollerResponse{}, err
	}
	result := InstancesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("InstancesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return InstancesClientDeletePollerResponse{}, err
	}
	result.Poller = &InstancesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *InstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, instanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns instance details for the given instance and account name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - Account name.
// instanceName - Instance name.
// options - InstancesClientGetOptions contains the optional parameters for the InstancesClient.Get method.
func (client *InstancesClient) Get(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientGetOptions) (InstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, instanceName, options)
	if err != nil {
		return InstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InstancesClient) getHandleResponse(resp *http.Response) (InstancesClientGetResponse, error) {
	result := InstancesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instance); err != nil {
		return InstancesClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks whether instance exists.
// resourceGroupName - The resource group name.
// accountName - Account name.
// instanceName - Instance name.
// options - InstancesClientHeadOptions contains the optional parameters for the InstancesClient.Head method.
func (client *InstancesClient) Head(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientHeadOptions) (InstancesClientHeadResponse, error) {
	req, err := client.headCreateRequest(ctx, resourceGroupName, accountName, instanceName, options)
	if err != nil {
		return InstancesClientHeadResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstancesClientHeadResponse{}, err
	}
	result := InstancesClientHeadResponse{RawResponse: resp}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// headCreateRequest creates the Head request.
func (client *InstancesClient) headCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, options *InstancesClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// ListByAccount - Returns instances for the given account name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - Account name.
// options - InstancesClientListByAccountOptions contains the optional parameters for the InstancesClient.ListByAccount method.
func (client *InstancesClient) ListByAccount(resourceGroupName string, accountName string, options *InstancesClientListByAccountOptions) *InstancesClientListByAccountPager {
	return &InstancesClientListByAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp InstancesClientListByAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.InstanceList.NextLink)
		},
	}
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *InstancesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *InstancesClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *InstancesClient) listByAccountHandleResponse(resp *http.Response) (InstancesClientListByAccountResponse, error) {
	result := InstancesClientListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.InstanceList); err != nil {
		return InstancesClientListByAccountResponse{}, err
	}
	return result, nil
}

// Update - Updates instance's tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - Account name.
// instanceName - Instance name.
// tagUpdatePayload - Updated tags.
// options - InstancesClientUpdateOptions contains the optional parameters for the InstancesClient.Update method.
func (client *InstancesClient) Update(ctx context.Context, resourceGroupName string, accountName string, instanceName string, tagUpdatePayload TagUpdate, options *InstancesClientUpdateOptions) (InstancesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, instanceName, tagUpdatePayload, options)
	if err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InstancesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *InstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, instanceName string, tagUpdatePayload TagUpdate, options *InstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeviceUpdate/accounts/{accountName}/instances/{instanceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if instanceName == "" {
		return nil, errors.New("parameter instanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceName}", url.PathEscape(instanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, tagUpdatePayload)
}

// updateHandleResponse handles the Update response.
func (client *InstancesClient) updateHandleResponse(resp *http.Response) (InstancesClientUpdateResponse, error) {
	result := InstancesClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Instance); err != nil {
		return InstancesClientUpdateResponse{}, err
	}
	return result, nil
}
