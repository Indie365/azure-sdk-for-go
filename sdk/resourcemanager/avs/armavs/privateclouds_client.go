//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// PrivateCloudsClient contains the methods for the PrivateClouds group.
// Don't use this type directly, use NewPrivateCloudsClient() instead.
type PrivateCloudsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateCloudsClient creates a new instance of PrivateCloudsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateCloudsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateCloudsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateCloudsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - privateCloud - The private cloud
//   - options - PrivateCloudsClientBeginCreateOrUpdateOptions contains the optional parameters for the PrivateCloudsClient.BeginCreateOrUpdate
//     method.
func (client *PrivateCloudsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateCloudsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, privateCloud, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateCloudsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateCloudsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *PrivateCloudsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateCloudsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, privateCloud, options)
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
func (client *PrivateCloudsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloud PrivateCloud, options *PrivateCloudsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, privateCloud); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - PrivateCloudsClientBeginDeleteOptions contains the optional parameters for the PrivateCloudsClient.BeginDelete
//     method.
func (client *PrivateCloudsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginDeleteOptions) (*runtime.Poller[PrivateCloudsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateCloudsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateCloudsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *PrivateCloudsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateCloudsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, options)
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
func (client *PrivateCloudsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - PrivateCloudsClientGetOptions contains the optional parameters for the PrivateCloudsClient.Get method.
func (client *PrivateCloudsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientGetOptions) (PrivateCloudsClientGetResponse, error) {
	var err error
	const operationName = "PrivateCloudsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateCloudsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateCloudsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateCloudsClient) getHandleResponse(resp *http.Response) (PrivateCloudsClientGetResponse, error) {
	result := PrivateCloudsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloud); err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List private clouds in a resource group
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - PrivateCloudsClientListOptions contains the optional parameters for the PrivateCloudsClient.NewListPager method.
func (client *PrivateCloudsClient) NewListPager(resourceGroupName string, options *PrivateCloudsClientListOptions) *runtime.Pager[PrivateCloudsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateCloudsClientListResponse]{
		More: func(page PrivateCloudsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateCloudsClientListResponse) (PrivateCloudsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateCloudsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return PrivateCloudsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PrivateCloudsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateCloudsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds"
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
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateCloudsClient) listHandleResponse(resp *http.Response) (PrivateCloudsClientListResponse, error) {
	result := PrivateCloudsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloudList); err != nil {
		return PrivateCloudsClientListResponse{}, err
	}
	return result, nil
}

// ListAdminCredentials - List the admin credentials for the private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - PrivateCloudsClientListAdminCredentialsOptions contains the optional parameters for the PrivateCloudsClient.ListAdminCredentials
//     method.
func (client *PrivateCloudsClient) ListAdminCredentials(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientListAdminCredentialsOptions) (PrivateCloudsClientListAdminCredentialsResponse, error) {
	var err error
	const operationName = "PrivateCloudsClient.ListAdminCredentials"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAdminCredentialsCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return PrivateCloudsClientListAdminCredentialsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateCloudsClientListAdminCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateCloudsClientListAdminCredentialsResponse{}, err
	}
	resp, err := client.listAdminCredentialsHandleResponse(httpResp)
	return resp, err
}

// listAdminCredentialsCreateRequest creates the ListAdminCredentials request.
func (client *PrivateCloudsClient) listAdminCredentialsCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientListAdminCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/listAdminCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAdminCredentialsHandleResponse handles the ListAdminCredentials response.
func (client *PrivateCloudsClient) listAdminCredentialsHandleResponse(resp *http.Response) (PrivateCloudsClientListAdminCredentialsResponse, error) {
	result := PrivateCloudsClientListAdminCredentialsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdminCredentials); err != nil {
		return PrivateCloudsClientListAdminCredentialsResponse{}, err
	}
	return result, nil
}

// NewListInSubscriptionPager - List private clouds in a subscription
//
// Generated from API version 2023-03-01
//   - options - PrivateCloudsClientListInSubscriptionOptions contains the optional parameters for the PrivateCloudsClient.NewListInSubscriptionPager
//     method.
func (client *PrivateCloudsClient) NewListInSubscriptionPager(options *PrivateCloudsClientListInSubscriptionOptions) *runtime.Pager[PrivateCloudsClientListInSubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateCloudsClientListInSubscriptionResponse]{
		More: func(page PrivateCloudsClientListInSubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateCloudsClientListInSubscriptionResponse) (PrivateCloudsClientListInSubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateCloudsClient.NewListInSubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listInSubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PrivateCloudsClientListInSubscriptionResponse{}, err
			}
			return client.listInSubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listInSubscriptionCreateRequest creates the ListInSubscription request.
func (client *PrivateCloudsClient) listInSubscriptionCreateRequest(ctx context.Context, options *PrivateCloudsClientListInSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AVS/privateClouds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listInSubscriptionHandleResponse handles the ListInSubscription response.
func (client *PrivateCloudsClient) listInSubscriptionHandleResponse(resp *http.Response) (PrivateCloudsClientListInSubscriptionResponse, error) {
	result := PrivateCloudsClientListInSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloudList); err != nil {
		return PrivateCloudsClientListInSubscriptionResponse{}, err
	}
	return result, nil
}

// BeginRotateNsxtPassword - Rotate the NSX-T Manager password
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - PrivateCloudsClientBeginRotateNsxtPasswordOptions contains the optional parameters for the PrivateCloudsClient.BeginRotateNsxtPassword
//     method.
func (client *PrivateCloudsClient) BeginRotateNsxtPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateNsxtPasswordOptions) (*runtime.Poller[PrivateCloudsClientRotateNsxtPasswordResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.rotateNsxtPassword(ctx, resourceGroupName, privateCloudName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateCloudsClientRotateNsxtPasswordResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateCloudsClientRotateNsxtPasswordResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RotateNsxtPassword - Rotate the NSX-T Manager password
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *PrivateCloudsClient) rotateNsxtPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateNsxtPasswordOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateCloudsClient.BeginRotateNsxtPassword"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rotateNsxtPasswordCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// rotateNsxtPasswordCreateRequest creates the RotateNsxtPassword request.
func (client *PrivateCloudsClient) rotateNsxtPasswordCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateNsxtPasswordOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/rotateNsxtPassword"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginRotateVcenterPassword - Rotate the vCenter password
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - PrivateCloudsClientBeginRotateVcenterPasswordOptions contains the optional parameters for the PrivateCloudsClient.BeginRotateVcenterPassword
//     method.
func (client *PrivateCloudsClient) BeginRotateVcenterPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateVcenterPasswordOptions) (*runtime.Poller[PrivateCloudsClientRotateVcenterPasswordResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.rotateVcenterPassword(ctx, resourceGroupName, privateCloudName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateCloudsClientRotateVcenterPasswordResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateCloudsClientRotateVcenterPasswordResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RotateVcenterPassword - Rotate the vCenter password
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *PrivateCloudsClient) rotateVcenterPassword(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateVcenterPasswordOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateCloudsClient.BeginRotateVcenterPassword"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rotateVcenterPasswordCreateRequest(ctx, resourceGroupName, privateCloudName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// rotateVcenterPasswordCreateRequest creates the RotateVcenterPassword request.
func (client *PrivateCloudsClient) rotateVcenterPasswordCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *PrivateCloudsClientBeginRotateVcenterPasswordOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/rotateVcenterPassword"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - privateCloudUpdate - The private cloud properties to be updated
//   - options - PrivateCloudsClientBeginUpdateOptions contains the optional parameters for the PrivateCloudsClient.BeginUpdate
//     method.
func (client *PrivateCloudsClient) BeginUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsClientBeginUpdateOptions) (*runtime.Poller[PrivateCloudsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, privateCloudName, privateCloudUpdate, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateCloudsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateCloudsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *PrivateCloudsClient) update(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateCloudsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateCloudName, privateCloudUpdate, options)
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
func (client *PrivateCloudsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, privateCloudUpdate PrivateCloudUpdate, options *PrivateCloudsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, privateCloudUpdate); err != nil {
		return nil, err
	}
	return req, nil
}
