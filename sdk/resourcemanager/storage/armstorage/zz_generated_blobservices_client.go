//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

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

// BlobServicesClient contains the methods for the BlobServices group.
// Don't use this type directly, use NewBlobServicesClient() instead.
type BlobServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBlobServicesClient creates a new instance of BlobServicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBlobServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *BlobServicesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &BlobServicesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetServiceProperties - Gets the properties of a storage account’s Blob service, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// options - BlobServicesClientGetServicePropertiesOptions contains the optional parameters for the BlobServicesClient.GetServiceProperties
// method.
func (client *BlobServicesClient) GetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesClientGetServicePropertiesOptions) (BlobServicesClientGetServicePropertiesResponse, error) {
	req, err := client.getServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return BlobServicesClientGetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobServicesClientGetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobServicesClientGetServicePropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.getServicePropertiesHandleResponse(resp)
}

// getServicePropertiesCreateRequest creates the GetServiceProperties request.
func (client *BlobServicesClient) getServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesClientGetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{BlobServicesName}", url.PathEscape("default"))
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

// getServicePropertiesHandleResponse handles the GetServiceProperties response.
func (client *BlobServicesClient) getServicePropertiesHandleResponse(resp *http.Response) (BlobServicesClientGetServicePropertiesResponse, error) {
	result := BlobServicesClientGetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobServiceProperties); err != nil {
		return BlobServicesClientGetServicePropertiesResponse{}, err
	}
	return result, nil
}

// List - List blob services of storage account. It returns a collection of one object named default.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// options - BlobServicesClientListOptions contains the optional parameters for the BlobServicesClient.List method.
func (client *BlobServicesClient) List(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesClientListOptions) (BlobServicesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return BlobServicesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobServicesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobServicesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *BlobServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *BlobServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
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
func (client *BlobServicesClient) listHandleResponse(resp *http.Response) (BlobServicesClientListResponse, error) {
	result := BlobServicesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobServiceItems); err != nil {
		return BlobServicesClientListResponse{}, err
	}
	return result, nil
}

// SetServiceProperties - Sets the properties of a storage account’s Blob service, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// parameters - The properties of a storage account’s Blob service, including properties for Storage Analytics and CORS (Cross-Origin
// Resource Sharing) rules.
// options - BlobServicesClientSetServicePropertiesOptions contains the optional parameters for the BlobServicesClient.SetServiceProperties
// method.
func (client *BlobServicesClient) SetServiceProperties(ctx context.Context, resourceGroupName string, accountName string, parameters BlobServiceProperties, options *BlobServicesClientSetServicePropertiesOptions) (BlobServicesClientSetServicePropertiesResponse, error) {
	req, err := client.setServicePropertiesCreateRequest(ctx, resourceGroupName, accountName, parameters, options)
	if err != nil {
		return BlobServicesClientSetServicePropertiesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BlobServicesClientSetServicePropertiesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BlobServicesClientSetServicePropertiesResponse{}, runtime.NewResponseError(resp)
	}
	return client.setServicePropertiesHandleResponse(resp)
}

// setServicePropertiesCreateRequest creates the SetServiceProperties request.
func (client *BlobServicesClient) setServicePropertiesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, parameters BlobServiceProperties, options *BlobServicesClientSetServicePropertiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{BlobServicesName}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// setServicePropertiesHandleResponse handles the SetServiceProperties response.
func (client *BlobServicesClient) setServicePropertiesHandleResponse(resp *http.Response) (BlobServicesClientSetServicePropertiesResponse, error) {
	result := BlobServicesClientSetServicePropertiesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BlobServiceProperties); err != nil {
		return BlobServicesClientSetServicePropertiesResponse{}, err
	}
	return result, nil
}
