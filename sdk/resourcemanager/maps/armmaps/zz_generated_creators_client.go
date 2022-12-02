//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

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
	"strings"
)

// CreatorsClient contains the methods for the Creators group.
// Don't use this type directly, use NewCreatorsClient() instead.
type CreatorsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCreatorsClient creates a new instance of CreatorsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCreatorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CreatorsClient, error) {
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
	client := &CreatorsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Maps Creator resource. Creator resource will manage Azure resources required to populate
// a custom set of mapping data. It requires an account to exist before it can be created.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// creatorName - The name of the Maps Creator instance.
// creatorResource - The new or updated parameters for the Creator resource.
// options - CreatorsClientCreateOrUpdateOptions contains the optional parameters for the CreatorsClient.CreateOrUpdate method.
func (client *CreatorsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, creatorName string, creatorResource Creator, options *CreatorsClientCreateOrUpdateOptions) (CreatorsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, creatorName, creatorResource, options)
	if err != nil {
		return CreatorsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CreatorsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CreatorsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CreatorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, creatorName string, creatorResource Creator, options *CreatorsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/creators/{creatorName}"
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
	if creatorName == "" {
		return nil, errors.New("parameter creatorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{creatorName}", url.PathEscape(creatorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, creatorResource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CreatorsClient) createOrUpdateHandleResponse(resp *http.Response) (CreatorsClientCreateOrUpdateResponse, error) {
	result := CreatorsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Creator); err != nil {
		return CreatorsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Maps Creator resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// creatorName - The name of the Maps Creator instance.
// options - CreatorsClientDeleteOptions contains the optional parameters for the CreatorsClient.Delete method.
func (client *CreatorsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, creatorName string, options *CreatorsClientDeleteOptions) (CreatorsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, creatorName, options)
	if err != nil {
		return CreatorsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CreatorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CreatorsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CreatorsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CreatorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, creatorName string, options *CreatorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/creators/{creatorName}"
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
	if creatorName == "" {
		return nil, errors.New("parameter creatorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{creatorName}", url.PathEscape(creatorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Maps Creator resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// creatorName - The name of the Maps Creator instance.
// options - CreatorsClientGetOptions contains the optional parameters for the CreatorsClient.Get method.
func (client *CreatorsClient) Get(ctx context.Context, resourceGroupName string, accountName string, creatorName string, options *CreatorsClientGetOptions) (CreatorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, creatorName, options)
	if err != nil {
		return CreatorsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CreatorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CreatorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CreatorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, creatorName string, options *CreatorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/creators/{creatorName}"
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
	if creatorName == "" {
		return nil, errors.New("parameter creatorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{creatorName}", url.PathEscape(creatorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CreatorsClient) getHandleResponse(resp *http.Response) (CreatorsClientGetResponse, error) {
	result := CreatorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Creator); err != nil {
		return CreatorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - Get all Creator instances for an Azure Maps Account
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// options - CreatorsClientListByAccountOptions contains the optional parameters for the CreatorsClient.ListByAccount method.
func (client *CreatorsClient) NewListByAccountPager(resourceGroupName string, accountName string, options *CreatorsClientListByAccountOptions) *runtime.Pager[CreatorsClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[CreatorsClientListByAccountResponse]{
		More: func(page CreatorsClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CreatorsClientListByAccountResponse) (CreatorsClientListByAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CreatorsClientListByAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CreatorsClientListByAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CreatorsClientListByAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAccountHandleResponse(resp)
		},
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *CreatorsClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *CreatorsClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/creators"
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
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *CreatorsClient) listByAccountHandleResponse(resp *http.Response) (CreatorsClientListByAccountResponse, error) {
	result := CreatorsClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreatorList); err != nil {
		return CreatorsClientListByAccountResponse{}, err
	}
	return result, nil
}

// Update - Updates the Maps Creator resource. Only a subset of the parameters may be updated after creation, such as Tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// accountName - The name of the Maps Account.
// creatorName - The name of the Maps Creator instance.
// creatorUpdateParameters - The update parameters for Maps Creator.
// options - CreatorsClientUpdateOptions contains the optional parameters for the CreatorsClient.Update method.
func (client *CreatorsClient) Update(ctx context.Context, resourceGroupName string, accountName string, creatorName string, creatorUpdateParameters CreatorUpdateParameters, options *CreatorsClientUpdateOptions) (CreatorsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, creatorName, creatorUpdateParameters, options)
	if err != nil {
		return CreatorsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CreatorsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CreatorsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CreatorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, creatorName string, creatorUpdateParameters CreatorUpdateParameters, options *CreatorsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maps/accounts/{accountName}/creators/{creatorName}"
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
	if creatorName == "" {
		return nil, errors.New("parameter creatorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{creatorName}", url.PathEscape(creatorName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, creatorUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *CreatorsClient) updateHandleResponse(resp *http.Response) (CreatorsClientUpdateResponse, error) {
	result := CreatorsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Creator); err != nil {
		return CreatorsClientUpdateResponse{}, err
	}
	return result, nil
}