//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// DisksClient contains the methods for the Disks group.
// Don't use this type directly, use NewDisksClient() instead.
type DisksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDisksClient creates a new instance of DisksClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDisksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DisksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DisksClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a disk.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
// characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
// characters.
// disk - Disk object supplied in the body of the Put disk operation.
// options - DisksClientBeginCreateOrUpdateOptions contains the optional parameters for the DisksClient.BeginCreateOrUpdate
// method.
func (client *DisksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (DisksClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return DisksClientCreateOrUpdatePollerResponse{}, err
	}
	result := DisksClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DisksClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return DisksClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DisksClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a disk.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DisksClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DisksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskName string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, disk)
}

// BeginDelete - Deletes a disk.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
// characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
// characters.
// options - DisksClientBeginDeleteOptions contains the optional parameters for the DisksClient.BeginDelete method.
func (client *DisksClient) BeginDelete(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginDeleteOptions) (DisksClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return DisksClientDeletePollerResponse{}, err
	}
	result := DisksClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DisksClient.Delete", "", resp, client.pl)
	if err != nil {
		return DisksClientDeletePollerResponse{}, err
	}
	result.Poller = &DisksClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a disk.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DisksClient) deleteOperation(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskName, options)
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
func (client *DisksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about a disk.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
// characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
// characters.
// options - DisksClientGetOptions contains the optional parameters for the DisksClient.Get method.
func (client *DisksClient) Get(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientGetOptions) (DisksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DisksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DisksClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisksClient) getHandleResponse(resp *http.Response) (DisksClientGetResponse, error) {
	result := DisksClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Disk); err != nil {
		return DisksClientGetResponse{}, err
	}
	return result, nil
}

// BeginGrantAccess - Grants access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
// characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
// characters.
// grantAccessData - Access data object supplied in the body of the get disk access operation.
// options - DisksClientBeginGrantAccessOptions contains the optional parameters for the DisksClient.BeginGrantAccess method.
func (client *DisksClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksClientBeginGrantAccessOptions) (DisksClientGrantAccessPollerResponse, error) {
	resp, err := client.grantAccess(ctx, resourceGroupName, diskName, grantAccessData, options)
	if err != nil {
		return DisksClientGrantAccessPollerResponse{}, err
	}
	result := DisksClientGrantAccessPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DisksClient.GrantAccess", "location", resp, client.pl)
	if err != nil {
		return DisksClientGrantAccessPollerResponse{}, err
	}
	result.Poller = &DisksClientGrantAccessPoller{
		pt: pt,
	}
	return result, nil
}

// GrantAccess - Grants access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DisksClient) grantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksClientBeginGrantAccessOptions) (*http.Response, error) {
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, diskName, grantAccessData, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *DisksClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, diskName string, grantAccessData GrantAccessData, options *DisksClientBeginGrantAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/beginGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, grantAccessData)
}

// List - Lists all the disks under a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DisksClientListOptions contains the optional parameters for the DisksClient.List method.
func (client *DisksClient) List(options *DisksClientListOptions) *DisksClientListPager {
	return &DisksClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DisksClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DisksClient) listCreateRequest(ctx context.Context, options *DisksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/disks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisksClient) listHandleResponse(resp *http.Response) (DisksClientListResponse, error) {
	result := DisksClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskList); err != nil {
		return DisksClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all the disks under a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - DisksClientListByResourceGroupOptions contains the optional parameters for the DisksClient.ListByResourceGroup
// method.
func (client *DisksClient) ListByResourceGroup(resourceGroupName string, options *DisksClientListByResourceGroupOptions) *DisksClientListByResourceGroupPager {
	return &DisksClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DisksClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DiskList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DisksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DisksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DisksClient) listByResourceGroupHandleResponse(resp *http.Response) (DisksClientListByResourceGroupResponse, error) {
	result := DisksClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskList); err != nil {
		return DisksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginRevokeAccess - Revokes access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
// characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
// characters.
// options - DisksClientBeginRevokeAccessOptions contains the optional parameters for the DisksClient.BeginRevokeAccess method.
func (client *DisksClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginRevokeAccessOptions) (DisksClientRevokeAccessPollerResponse, error) {
	resp, err := client.revokeAccess(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return DisksClientRevokeAccessPollerResponse{}, err
	}
	result := DisksClientRevokeAccessPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DisksClient.RevokeAccess", "location", resp, client.pl)
	if err != nil {
		return DisksClientRevokeAccessPollerResponse{}, err
	}
	result.Poller = &DisksClientRevokeAccessPoller{
		pt: pt,
	}
	return result, nil
}

// RevokeAccess - Revokes access to a disk.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DisksClient) revokeAccess(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginRevokeAccessOptions) (*http.Response, error) {
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, diskName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *DisksClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, diskName string, options *DisksClientBeginRevokeAccessOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}/endGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdate - Updates (patches) a disk.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// diskName - The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported
// characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80
// characters.
// disk - Disk object supplied in the body of the Patch disk operation.
// options - DisksClientBeginUpdateOptions contains the optional parameters for the DisksClient.BeginUpdate method.
func (client *DisksClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksClientBeginUpdateOptions) (DisksClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return DisksClientUpdatePollerResponse{}, err
	}
	result := DisksClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DisksClient.Update", "", resp, client.pl)
	if err != nil {
		return DisksClientUpdatePollerResponse{}, err
	}
	result.Poller = &DisksClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates (patches) a disk.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DisksClient) update(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskName, disk, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DisksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskName string, disk DiskUpdate, options *DisksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/disks/{diskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskName == "" {
		return nil, errors.New("parameter diskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskName}", url.PathEscape(diskName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, disk)
}
