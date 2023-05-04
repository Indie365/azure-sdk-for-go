//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeploymentmanager

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

// ArtifactSourcesClient contains the methods for the ArtifactSources group.
// Don't use this type directly, use NewArtifactSourcesClient() instead.
type ArtifactSourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewArtifactSourcesClient creates a new instance of ArtifactSourcesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewArtifactSourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ArtifactSourcesClient, error) {
	cl, err := arm.NewClient(moduleName+".ArtifactSourcesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ArtifactSourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Synchronously creates a new artifact source or updates an existing artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - artifactSourceName - The name of the artifact source.
//   - options - ArtifactSourcesClientCreateOrUpdateOptions contains the optional parameters for the ArtifactSourcesClient.CreateOrUpdate
//     method.
func (client *ArtifactSourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, artifactSourceName string, options *ArtifactSourcesClientCreateOrUpdateOptions) (ArtifactSourcesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, artifactSourceName, options)
	if err != nil {
		return ArtifactSourcesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ArtifactSourcesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ArtifactSourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, artifactSourceName string, options *ArtifactSourcesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources/{artifactSourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ArtifactSourceInfo != nil {
		return req, runtime.MarshalAsJSON(req, *options.ArtifactSourceInfo)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ArtifactSourcesClient) createOrUpdateHandleResponse(resp *http.Response) (ArtifactSourcesClientCreateOrUpdateResponse, error) {
	result := ArtifactSourcesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactSource); err != nil {
		return ArtifactSourcesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - artifactSourceName - The name of the artifact source.
//   - options - ArtifactSourcesClientDeleteOptions contains the optional parameters for the ArtifactSourcesClient.Delete method.
func (client *ArtifactSourcesClient) Delete(ctx context.Context, resourceGroupName string, artifactSourceName string, options *ArtifactSourcesClientDeleteOptions) (ArtifactSourcesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, artifactSourceName, options)
	if err != nil {
		return ArtifactSourcesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ArtifactSourcesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ArtifactSourcesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ArtifactSourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, artifactSourceName string, options *ArtifactSourcesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources/{artifactSourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an artifact source.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - artifactSourceName - The name of the artifact source.
//   - options - ArtifactSourcesClientGetOptions contains the optional parameters for the ArtifactSourcesClient.Get method.
func (client *ArtifactSourcesClient) Get(ctx context.Context, resourceGroupName string, artifactSourceName string, options *ArtifactSourcesClientGetOptions) (ArtifactSourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, artifactSourceName, options)
	if err != nil {
		return ArtifactSourcesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArtifactSourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ArtifactSourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, artifactSourceName string, options *ArtifactSourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources/{artifactSourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if artifactSourceName == "" {
		return nil, errors.New("parameter artifactSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{artifactSourceName}", url.PathEscape(artifactSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ArtifactSourcesClient) getHandleResponse(resp *http.Response) (ArtifactSourcesClientGetResponse, error) {
	result := ArtifactSourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactSource); err != nil {
		return ArtifactSourcesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the artifact sources in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ArtifactSourcesClientListOptions contains the optional parameters for the ArtifactSourcesClient.List method.
func (client *ArtifactSourcesClient) List(ctx context.Context, resourceGroupName string, options *ArtifactSourcesClientListOptions) (ArtifactSourcesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ArtifactSourcesClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ArtifactSourcesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ArtifactSourcesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ArtifactSourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ArtifactSourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DeploymentManager/artifactSources"
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
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ArtifactSourcesClient) listHandleResponse(resp *http.Response) (ArtifactSourcesClientListResponse, error) {
	result := ArtifactSourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArtifactSourceArray); err != nil {
		return ArtifactSourcesClientListResponse{}, err
	}
	return result, nil
}
