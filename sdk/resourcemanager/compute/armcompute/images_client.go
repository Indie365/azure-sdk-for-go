//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// ImagesClient contains the methods for the Images group.
// Don't use this type directly, use NewImagesClient() instead.
type ImagesClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewImagesClient creates a new instance of ImagesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImagesClient, error) {
	cl, err := arm.NewClient(moduleName+".ImagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImagesClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update an image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - imageName - The name of the image.
//   - parameters - Parameters supplied to the Create Image operation.
//   - options - ImagesClientBeginCreateOrUpdateOptions contains the optional parameters for the ImagesClient.BeginCreateOrUpdate
//     method.
func (client *ImagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ImagesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, imageName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ImagesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ImagesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update an image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ImagesClient) createOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ImagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, imageName string, parameters Image, options *ImagesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

// BeginDelete - Deletes an Image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - imageName - The name of the image.
//   - options - ImagesClientBeginDeleteOptions contains the optional parameters for the ImagesClient.BeginDelete method.
func (client *ImagesClient) BeginDelete(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientBeginDeleteOptions) (*runtime.Poller[ImagesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, imageName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ImagesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ImagesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an Image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ImagesClient) deleteOperation(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, imageName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ImagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - imageName - The name of the image.
//   - options - ImagesClientGetOptions contains the optional parameters for the ImagesClient.Get method.
func (client *ImagesClient) Get(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientGetOptions) (ImagesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, imageName, options)
	if err != nil {
		return ImagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ImagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ImagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, imageName string, options *ImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImagesClient) getHandleResponse(resp *http.Response) (ImagesClientGetResponse, error) {
	result := ImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Image); err != nil {
		return ImagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of Images in the subscription. Use nextLink property in the response to get the next page
// of Images. Do this till nextLink is null to fetch all the Images.
//
// Generated from API version 2023-07-01
//   - options - ImagesClientListOptions contains the optional parameters for the ImagesClient.NewListPager method.
func (client *ImagesClient) NewListPager(options *ImagesClientListOptions) (*runtime.Pager[ImagesClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ImagesClientListResponse]{
		More: func(page ImagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImagesClientListResponse) (ImagesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ImagesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ImagesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImagesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ImagesClient) listCreateRequest(ctx context.Context, options *ImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/images"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ImagesClient) listHandleResponse(resp *http.Response) (ImagesClientListResponse, error) {
	result := ImagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageListResult); err != nil {
		return ImagesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets the list of images under a resource group. Use nextLink property in the response to
// get the next page of Images. Do this till nextLink is null to fetch all the Images.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - options - ImagesClientListByResourceGroupOptions contains the optional parameters for the ImagesClient.NewListByResourceGroupPager
//     method.
func (client *ImagesClient) NewListByResourceGroupPager(resourceGroupName string, options *ImagesClientListByResourceGroupOptions) (*runtime.Pager[ImagesClientListByResourceGroupResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ImagesClientListByResourceGroupResponse]{
		More: func(page ImagesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImagesClientListByResourceGroupResponse) (ImagesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ImagesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ImagesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImagesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ImagesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ImagesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ImagesClient) listByResourceGroupHandleResponse(resp *http.Response) (ImagesClientListByResourceGroupResponse, error) {
	result := ImagesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageListResult); err != nil {
		return ImagesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group.
//   - imageName - The name of the image.
//   - parameters - Parameters supplied to the Update Image operation.
//   - options - ImagesClientBeginUpdateOptions contains the optional parameters for the ImagesClient.BeginUpdate method.
func (client *ImagesClient) BeginUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesClientBeginUpdateOptions) (*runtime.Poller[ImagesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, imageName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ImagesClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ImagesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update an image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ImagesClient) update(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, imageName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ImagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, imageName string, parameters ImageUpdate, options *ImagesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

