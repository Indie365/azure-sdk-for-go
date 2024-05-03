//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PartnerNamespacesClient contains the methods for the PartnerNamespaces group.
// Don't use this type directly, use NewPartnerNamespacesClient() instead.
type PartnerNamespacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPartnerNamespacesClient creates a new instance of PartnerNamespacesClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPartnerNamespacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PartnerNamespacesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PartnerNamespacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates a new partner namespace with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerNamespaceName - Name of the partner namespace.
//   - partnerNamespaceInfo - PartnerNamespace information.
//   - options - PartnerNamespacesClientBeginCreateOrUpdateOptions contains the optional parameters for the PartnerNamespacesClient.BeginCreateOrUpdate
//     method.
func (client *PartnerNamespacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceInfo PartnerNamespace, options *PartnerNamespacesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PartnerNamespacesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, partnerNamespaceName, partnerNamespaceInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PartnerNamespacesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PartnerNamespacesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Asynchronously creates a new partner namespace with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *PartnerNamespacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceInfo PartnerNamespace, options *PartnerNamespacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PartnerNamespacesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerNamespaceName, partnerNamespaceInfo, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PartnerNamespacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceInfo PartnerNamespace, options *PartnerNamespacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, partnerNamespaceInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete existing partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerNamespaceName - Name of the partner namespace.
//   - options - PartnerNamespacesClientBeginDeleteOptions contains the optional parameters for the PartnerNamespacesClient.BeginDelete
//     method.
func (client *PartnerNamespacesClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *PartnerNamespacesClientBeginDeleteOptions) (*runtime.Poller[PartnerNamespacesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, partnerNamespaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PartnerNamespacesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PartnerNamespacesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete existing partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *PartnerNamespacesClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *PartnerNamespacesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PartnerNamespacesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerNamespaceName, options)
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
func (client *PartnerNamespacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *PartnerNamespacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of a partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerNamespaceName - Name of the partner namespace.
//   - options - PartnerNamespacesClientGetOptions contains the optional parameters for the PartnerNamespacesClient.Get method.
func (client *PartnerNamespacesClient) Get(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *PartnerNamespacesClientGetOptions) (PartnerNamespacesClientGetResponse, error) {
	var err error
	const operationName = "PartnerNamespacesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerNamespaceName, options)
	if err != nil {
		return PartnerNamespacesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerNamespacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerNamespacesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PartnerNamespacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *PartnerNamespacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerNamespacesClient) getHandleResponse(resp *http.Response) (PartnerNamespacesClientGetResponse, error) {
	result := PartnerNamespacesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerNamespace); err != nil {
		return PartnerNamespacesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the partner namespaces under a resource group.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - options - PartnerNamespacesClientListByResourceGroupOptions contains the optional parameters for the PartnerNamespacesClient.NewListByResourceGroupPager
//     method.
func (client *PartnerNamespacesClient) NewListByResourceGroupPager(resourceGroupName string, options *PartnerNamespacesClientListByResourceGroupOptions) *runtime.Pager[PartnerNamespacesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerNamespacesClientListByResourceGroupResponse]{
		More: func(page PartnerNamespacesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerNamespacesClientListByResourceGroupResponse) (PartnerNamespacesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PartnerNamespacesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return PartnerNamespacesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PartnerNamespacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PartnerNamespacesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PartnerNamespacesClient) listByResourceGroupHandleResponse(resp *http.Response) (PartnerNamespacesClientListByResourceGroupResponse, error) {
	result := PartnerNamespacesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerNamespacesListResult); err != nil {
		return PartnerNamespacesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all the partner namespaces under an Azure subscription.
//
// Generated from API version 2024-06-01-preview
//   - options - PartnerNamespacesClientListBySubscriptionOptions contains the optional parameters for the PartnerNamespacesClient.NewListBySubscriptionPager
//     method.
func (client *PartnerNamespacesClient) NewListBySubscriptionPager(options *PartnerNamespacesClientListBySubscriptionOptions) *runtime.Pager[PartnerNamespacesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerNamespacesClientListBySubscriptionResponse]{
		More: func(page PartnerNamespacesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerNamespacesClientListBySubscriptionResponse) (PartnerNamespacesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PartnerNamespacesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PartnerNamespacesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PartnerNamespacesClient) listBySubscriptionCreateRequest(ctx context.Context, options *PartnerNamespacesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/partnerNamespaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PartnerNamespacesClient) listBySubscriptionHandleResponse(resp *http.Response) (PartnerNamespacesClientListBySubscriptionResponse, error) {
	result := PartnerNamespacesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerNamespacesListResult); err != nil {
		return PartnerNamespacesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListSharedAccessKeys - List the two keys used to publish to a partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerNamespaceName - Name of the partner namespace.
//   - options - PartnerNamespacesClientListSharedAccessKeysOptions contains the optional parameters for the PartnerNamespacesClient.ListSharedAccessKeys
//     method.
func (client *PartnerNamespacesClient) ListSharedAccessKeys(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *PartnerNamespacesClientListSharedAccessKeysOptions) (PartnerNamespacesClientListSharedAccessKeysResponse, error) {
	var err error
	const operationName = "PartnerNamespacesClient.ListSharedAccessKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSharedAccessKeysCreateRequest(ctx, resourceGroupName, partnerNamespaceName, options)
	if err != nil {
		return PartnerNamespacesClientListSharedAccessKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerNamespacesClientListSharedAccessKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerNamespacesClientListSharedAccessKeysResponse{}, err
	}
	resp, err := client.listSharedAccessKeysHandleResponse(httpResp)
	return resp, err
}

// listSharedAccessKeysCreateRequest creates the ListSharedAccessKeys request.
func (client *PartnerNamespacesClient) listSharedAccessKeysCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, options *PartnerNamespacesClientListSharedAccessKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSharedAccessKeysHandleResponse handles the ListSharedAccessKeys response.
func (client *PartnerNamespacesClient) listSharedAccessKeysHandleResponse(resp *http.Response) (PartnerNamespacesClientListSharedAccessKeysResponse, error) {
	result := PartnerNamespacesClientListSharedAccessKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerNamespaceSharedAccessKeys); err != nil {
		return PartnerNamespacesClientListSharedAccessKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKey - Regenerate a shared access key for a partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerNamespaceName - Name of the partner namespace.
//   - regenerateKeyRequest - Request body to regenerate key.
//   - options - PartnerNamespacesClientRegenerateKeyOptions contains the optional parameters for the PartnerNamespacesClient.RegenerateKey
//     method.
func (client *PartnerNamespacesClient) RegenerateKey(ctx context.Context, resourceGroupName string, partnerNamespaceName string, regenerateKeyRequest PartnerNamespaceRegenerateKeyRequest, options *PartnerNamespacesClientRegenerateKeyOptions) (PartnerNamespacesClientRegenerateKeyResponse, error) {
	var err error
	const operationName = "PartnerNamespacesClient.RegenerateKey"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, partnerNamespaceName, regenerateKeyRequest, options)
	if err != nil {
		return PartnerNamespacesClientRegenerateKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerNamespacesClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerNamespacesClientRegenerateKeyResponse{}, err
	}
	resp, err := client.regenerateKeyHandleResponse(httpResp)
	return resp, err
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *PartnerNamespacesClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, regenerateKeyRequest PartnerNamespaceRegenerateKeyRequest, options *PartnerNamespacesClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, regenerateKeyRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *PartnerNamespacesClient) regenerateKeyHandleResponse(resp *http.Response) (PartnerNamespacesClientRegenerateKeyResponse, error) {
	result := PartnerNamespacesClientRegenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartnerNamespaceSharedAccessKeys); err != nil {
		return PartnerNamespacesClientRegenerateKeyResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Asynchronously updates a partner namespace with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerNamespaceName - Name of the partner namespace.
//   - partnerNamespaceUpdateParameters - Partner namespace update information.
//   - options - PartnerNamespacesClientBeginUpdateOptions contains the optional parameters for the PartnerNamespacesClient.BeginUpdate
//     method.
func (client *PartnerNamespacesClient) BeginUpdate(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceUpdateParameters PartnerNamespaceUpdateParameters, options *PartnerNamespacesClientBeginUpdateOptions) (*runtime.Poller[PartnerNamespacesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, partnerNamespaceName, partnerNamespaceUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PartnerNamespacesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PartnerNamespacesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Asynchronously updates a partner namespace with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
func (client *PartnerNamespacesClient) update(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceUpdateParameters PartnerNamespaceUpdateParameters, options *PartnerNamespacesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PartnerNamespacesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, partnerNamespaceName, partnerNamespaceUpdateParameters, options)
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
func (client *PartnerNamespacesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, partnerNamespaceName string, partnerNamespaceUpdateParameters PartnerNamespaceUpdateParameters, options *PartnerNamespacesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerNamespaces/{partnerNamespaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerNamespaceName == "" {
		return nil, errors.New("parameter partnerNamespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerNamespaceName}", url.PathEscape(partnerNamespaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, partnerNamespaceUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
