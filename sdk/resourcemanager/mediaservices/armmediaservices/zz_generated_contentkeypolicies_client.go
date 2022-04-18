//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ContentKeyPoliciesClient contains the methods for the ContentKeyPolicies group.
// Don't use this type directly, use NewContentKeyPoliciesClient() instead.
type ContentKeyPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewContentKeyPoliciesClient creates a new instance of ContentKeyPoliciesClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewContentKeyPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContentKeyPoliciesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ContentKeyPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Content Key Policy in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// contentKeyPolicyName - The Content Key Policy name.
// parameters - The request parameters
// options - ContentKeyPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ContentKeyPoliciesClient.CreateOrUpdate
// method.
func (client *ContentKeyPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy, options *ContentKeyPoliciesClientCreateOrUpdateOptions) (ContentKeyPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, contentKeyPolicyName, parameters, options)
	if err != nil {
		return ContentKeyPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentKeyPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ContentKeyPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContentKeyPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy, options *ContentKeyPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}"
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
	if contentKeyPolicyName == "" {
		return nil, errors.New("parameter contentKeyPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentKeyPolicyName}", url.PathEscape(contentKeyPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ContentKeyPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ContentKeyPoliciesClientCreateOrUpdateResponse, error) {
	result := ContentKeyPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentKeyPolicy); err != nil {
		return ContentKeyPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Content Key Policy in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// contentKeyPolicyName - The Content Key Policy name.
// options - ContentKeyPoliciesClientDeleteOptions contains the optional parameters for the ContentKeyPoliciesClient.Delete
// method.
func (client *ContentKeyPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, options *ContentKeyPoliciesClientDeleteOptions) (ContentKeyPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, contentKeyPolicyName, options)
	if err != nil {
		return ContentKeyPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentKeyPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ContentKeyPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ContentKeyPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContentKeyPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, options *ContentKeyPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}"
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
	if contentKeyPolicyName == "" {
		return nil, errors.New("parameter contentKeyPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentKeyPolicyName}", url.PathEscape(contentKeyPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the details of a Content Key Policy in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// contentKeyPolicyName - The Content Key Policy name.
// options - ContentKeyPoliciesClientGetOptions contains the optional parameters for the ContentKeyPoliciesClient.Get method.
func (client *ContentKeyPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, options *ContentKeyPoliciesClientGetOptions) (ContentKeyPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, contentKeyPolicyName, options)
	if err != nil {
		return ContentKeyPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentKeyPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContentKeyPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContentKeyPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, options *ContentKeyPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}"
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
	if contentKeyPolicyName == "" {
		return nil, errors.New("parameter contentKeyPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentKeyPolicyName}", url.PathEscape(contentKeyPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContentKeyPoliciesClient) getHandleResponse(resp *http.Response) (ContentKeyPoliciesClientGetResponse, error) {
	result := ContentKeyPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentKeyPolicy); err != nil {
		return ContentKeyPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// GetPolicyPropertiesWithSecrets - Get a Content Key Policy including secret values
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// contentKeyPolicyName - The Content Key Policy name.
// options - ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsOptions contains the optional parameters for the ContentKeyPoliciesClient.GetPolicyPropertiesWithSecrets
// method.
func (client *ContentKeyPoliciesClient) GetPolicyPropertiesWithSecrets(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, options *ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsOptions) (ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsResponse, error) {
	req, err := client.getPolicyPropertiesWithSecretsCreateRequest(ctx, resourceGroupName, accountName, contentKeyPolicyName, options)
	if err != nil {
		return ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPolicyPropertiesWithSecretsHandleResponse(resp)
}

// getPolicyPropertiesWithSecretsCreateRequest creates the GetPolicyPropertiesWithSecrets request.
func (client *ContentKeyPoliciesClient) getPolicyPropertiesWithSecretsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, options *ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}/getPolicyPropertiesWithSecrets"
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
	if contentKeyPolicyName == "" {
		return nil, errors.New("parameter contentKeyPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentKeyPolicyName}", url.PathEscape(contentKeyPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPolicyPropertiesWithSecretsHandleResponse handles the GetPolicyPropertiesWithSecrets response.
func (client *ContentKeyPoliciesClient) getPolicyPropertiesWithSecretsHandleResponse(resp *http.Response) (ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsResponse, error) {
	result := ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentKeyPolicyProperties); err != nil {
		return ContentKeyPoliciesClientGetPolicyPropertiesWithSecretsResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the Content Key Policies in the account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// options - ContentKeyPoliciesClientListOptions contains the optional parameters for the ContentKeyPoliciesClient.List method.
func (client *ContentKeyPoliciesClient) NewListPager(resourceGroupName string, accountName string, options *ContentKeyPoliciesClientListOptions) *runtime.Pager[ContentKeyPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[ContentKeyPoliciesClientListResponse]{
		More: func(page ContentKeyPoliciesClientListResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContentKeyPoliciesClientListResponse) (ContentKeyPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.ODataNextLink)
			}
			if err != nil {
				return ContentKeyPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ContentKeyPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ContentKeyPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ContentKeyPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *ContentKeyPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies"
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
	reqQP.Set("api-version", "2021-11-01")
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContentKeyPoliciesClient) listHandleResponse(resp *http.Response) (ContentKeyPoliciesClientListResponse, error) {
	result := ContentKeyPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentKeyPolicyCollection); err != nil {
		return ContentKeyPoliciesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing Content Key Policy in the Media Services account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// contentKeyPolicyName - The Content Key Policy name.
// parameters - The request parameters
// options - ContentKeyPoliciesClientUpdateOptions contains the optional parameters for the ContentKeyPoliciesClient.Update
// method.
func (client *ContentKeyPoliciesClient) Update(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy, options *ContentKeyPoliciesClientUpdateOptions) (ContentKeyPoliciesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, contentKeyPolicyName, parameters, options)
	if err != nil {
		return ContentKeyPoliciesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContentKeyPoliciesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContentKeyPoliciesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ContentKeyPoliciesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, contentKeyPolicyName string, parameters ContentKeyPolicy, options *ContentKeyPoliciesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/contentKeyPolicies/{contentKeyPolicyName}"
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
	if contentKeyPolicyName == "" {
		return nil, errors.New("parameter contentKeyPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentKeyPolicyName}", url.PathEscape(contentKeyPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ContentKeyPoliciesClient) updateHandleResponse(resp *http.Response) (ContentKeyPoliciesClientUpdateResponse, error) {
	result := ContentKeyPoliciesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentKeyPolicy); err != nil {
		return ContentKeyPoliciesClientUpdateResponse{}, err
	}
	return result, nil
}
