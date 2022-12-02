//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmediaservices

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AssetsClient contains the methods for the Assets group.
// Don't use this type directly, use NewAssetsClient() instead.
type AssetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAssetsClient creates a new instance of AssetsClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAssetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssetsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AssetsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an Asset in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// parameters - The request parameters
// options - AssetsClientCreateOrUpdateOptions contains the optional parameters for the AssetsClient.CreateOrUpdate method.
func (client *AssetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsClientCreateOrUpdateOptions) (AssetsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, assetName, parameters, options)
	if err != nil {
		return AssetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AssetsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AssetsClient) createOrUpdateHandleResponse(resp *http.Response) (AssetsClientCreateOrUpdateResponse, error) {
	result := AssetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Asset); err != nil {
		return AssetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Asset in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// options - AssetsClientDeleteOptions contains the optional parameters for the AssetsClient.Delete method.
func (client *AssetsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientDeleteOptions) (AssetsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AssetsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AssetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the details of an Asset in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// options - AssetsClientGetOptions contains the optional parameters for the AssetsClient.Get method.
func (client *AssetsClient) Get(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientGetOptions) (AssetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssetsClient) getHandleResponse(resp *http.Response) (AssetsClientGetResponse, error) {
	result := AssetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Asset); err != nil {
		return AssetsClientGetResponse{}, err
	}
	return result, nil
}

// GetEncryptionKey - Gets the Asset storage encryption keys used to decrypt content created by version 2 of the Media Services
// API
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// options - AssetsClientGetEncryptionKeyOptions contains the optional parameters for the AssetsClient.GetEncryptionKey method.
func (client *AssetsClient) GetEncryptionKey(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientGetEncryptionKeyOptions) (AssetsClientGetEncryptionKeyResponse, error) {
	req, err := client.getEncryptionKeyCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsClientGetEncryptionKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsClientGetEncryptionKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsClientGetEncryptionKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEncryptionKeyHandleResponse(resp)
}

// getEncryptionKeyCreateRequest creates the GetEncryptionKey request.
func (client *AssetsClient) getEncryptionKeyCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientGetEncryptionKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/getEncryptionKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEncryptionKeyHandleResponse handles the GetEncryptionKey response.
func (client *AssetsClient) getEncryptionKeyHandleResponse(resp *http.Response) (AssetsClientGetEncryptionKeyResponse, error) {
	result := AssetsClientGetEncryptionKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageEncryptedAssetDecryptionData); err != nil {
		return AssetsClientGetEncryptionKeyResponse{}, err
	}
	return result, nil
}

// NewListPager - List Assets in the Media Services account with optional filtering and ordering
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// options - AssetsClientListOptions contains the optional parameters for the AssetsClient.List method.
func (client *AssetsClient) NewListPager(resourceGroupName string, accountName string, options *AssetsClientListOptions) *runtime.Pager[AssetsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssetsClientListResponse]{
		More: func(page AssetsClientListResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssetsClientListResponse) (AssetsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.ODataNextLink)
			}
			if err != nil {
				return AssetsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AssetsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssetsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AssetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AssetsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssetsClient) listHandleResponse(resp *http.Response) (AssetsClientListResponse, error) {
	result := AssetsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetCollection); err != nil {
		return AssetsClientListResponse{}, err
	}
	return result, nil
}

// ListContainerSas - Lists storage container URLs with shared access signatures (SAS) for uploading and downloading Asset
// content. The signatures are derived from the storage account keys.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// parameters - The request parameters
// options - AssetsClientListContainerSasOptions contains the optional parameters for the AssetsClient.ListContainerSas method.
func (client *AssetsClient) ListContainerSas(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters ListContainerSasInput, options *AssetsClientListContainerSasOptions) (AssetsClientListContainerSasResponse, error) {
	req, err := client.listContainerSasCreateRequest(ctx, resourceGroupName, accountName, assetName, parameters, options)
	if err != nil {
		return AssetsClientListContainerSasResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsClientListContainerSasResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsClientListContainerSasResponse{}, runtime.NewResponseError(resp)
	}
	return client.listContainerSasHandleResponse(resp)
}

// listContainerSasCreateRequest creates the ListContainerSas request.
func (client *AssetsClient) listContainerSasCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters ListContainerSasInput, options *AssetsClientListContainerSasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/listContainerSas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listContainerSasHandleResponse handles the ListContainerSas response.
func (client *AssetsClient) listContainerSasHandleResponse(resp *http.Response) (AssetsClientListContainerSasResponse, error) {
	result := AssetsClientListContainerSasResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetContainerSas); err != nil {
		return AssetsClientListContainerSasResponse{}, err
	}
	return result, nil
}

// ListStreamingLocators - Lists Streaming Locators which are associated with this asset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// options - AssetsClientListStreamingLocatorsOptions contains the optional parameters for the AssetsClient.ListStreamingLocators
// method.
func (client *AssetsClient) ListStreamingLocators(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientListStreamingLocatorsOptions) (AssetsClientListStreamingLocatorsResponse, error) {
	req, err := client.listStreamingLocatorsCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsClientListStreamingLocatorsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsClientListStreamingLocatorsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsClientListStreamingLocatorsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listStreamingLocatorsHandleResponse(resp)
}

// listStreamingLocatorsCreateRequest creates the ListStreamingLocators request.
func (client *AssetsClient) listStreamingLocatorsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsClientListStreamingLocatorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/listStreamingLocators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listStreamingLocatorsHandleResponse handles the ListStreamingLocators response.
func (client *AssetsClient) listStreamingLocatorsHandleResponse(resp *http.Response) (AssetsClientListStreamingLocatorsResponse, error) {
	result := AssetsClientListStreamingLocatorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListStreamingLocatorsResponse); err != nil {
		return AssetsClientListStreamingLocatorsResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing Asset in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// assetName - The Asset name.
// parameters - The request parameters
// options - AssetsClientUpdateOptions contains the optional parameters for the AssetsClient.Update method.
func (client *AssetsClient) Update(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsClientUpdateOptions) (AssetsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, assetName, parameters, options)
	if err != nil {
		return AssetsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AssetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AssetsClient) updateHandleResponse(resp *http.Response) (AssetsClientUpdateResponse, error) {
	result := AssetsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Asset); err != nil {
		return AssetsClientUpdateResponse{}, err
	}
	return result, nil
}