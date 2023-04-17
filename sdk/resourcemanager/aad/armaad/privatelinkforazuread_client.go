//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armaad

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

// PrivateLinkForAzureAdClient contains the methods for the PrivateLinkForAzureAd group.
// Don't use this type directly, use NewPrivateLinkForAzureAdClient() instead.
type PrivateLinkForAzureAdClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkForAzureAdClient creates a new instance of PrivateLinkForAzureAdClient with the specified values.
//   - subscriptionID - Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkForAzureAdClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkForAzureAdClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateLinkForAzureAdClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkForAzureAdClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a private link policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - Name of an Azure resource group.
//   - policyName - The name of the private link policy in Azure AD.
//   - privateLinkPolicy - private link Policy supplied to the operation.
//   - options - PrivateLinkForAzureAdClientBeginCreateOptions contains the optional parameters for the PrivateLinkForAzureAdClient.BeginCreate
//     method.
func (client *PrivateLinkForAzureAdClient) BeginCreate(ctx context.Context, resourceGroupName string, policyName string, privateLinkPolicy PrivateLinkPolicy, options *PrivateLinkForAzureAdClientBeginCreateOptions) (*runtime.Poller[PrivateLinkForAzureAdClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, policyName, privateLinkPolicy, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PrivateLinkForAzureAdClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PrivateLinkForAzureAdClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a private link policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *PrivateLinkForAzureAdClient) create(ctx context.Context, resourceGroupName string, policyName string, privateLinkPolicy PrivateLinkPolicy, options *PrivateLinkForAzureAdClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, policyName, privateLinkPolicy, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *PrivateLinkForAzureAdClient) createCreateRequest(ctx context.Context, resourceGroupName string, policyName string, privateLinkPolicy PrivateLinkPolicy, options *PrivateLinkForAzureAdClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.aadiam/privateLinkForAzureAd/{policyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, privateLinkPolicy)
}

// Delete - Deletes a private link policy. When operation completes, status code 200 returned without content.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - Name of an Azure resource group.
//   - policyName - The name of the private link policy in Azure AD.
//   - options - PrivateLinkForAzureAdClientDeleteOptions contains the optional parameters for the PrivateLinkForAzureAdClient.Delete
//     method.
func (client *PrivateLinkForAzureAdClient) Delete(ctx context.Context, resourceGroupName string, policyName string, options *PrivateLinkForAzureAdClientDeleteOptions) (PrivateLinkForAzureAdClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return PrivateLinkForAzureAdClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkForAzureAdClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PrivateLinkForAzureAdClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PrivateLinkForAzureAdClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateLinkForAzureAdClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *PrivateLinkForAzureAdClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.aadiam/privateLinkForAzureAd/{policyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a private link policy with a given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - Name of an Azure resource group.
//   - policyName - The name of the private link policy in Azure AD.
//   - options - PrivateLinkForAzureAdClientGetOptions contains the optional parameters for the PrivateLinkForAzureAdClient.Get
//     method.
func (client *PrivateLinkForAzureAdClient) Get(ctx context.Context, resourceGroupName string, policyName string, options *PrivateLinkForAzureAdClientGetOptions) (PrivateLinkForAzureAdClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return PrivateLinkForAzureAdClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkForAzureAdClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkForAzureAdClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkForAzureAdClient) getCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *PrivateLinkForAzureAdClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.aadiam/privateLinkForAzureAd/{policyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkForAzureAdClient) getHandleResponse(resp *http.Response) (PrivateLinkForAzureAdClientGetResponse, error) {
	result := PrivateLinkForAzureAdClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkPolicy); err != nil {
		return PrivateLinkForAzureAdClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Operation to return the list of Private Link Policies For AzureAD scoped to the resourceGroup.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - Name of an Azure resource group.
//   - options - PrivateLinkForAzureAdClientListOptions contains the optional parameters for the PrivateLinkForAzureAdClient.NewListPager
//     method.
func (client *PrivateLinkForAzureAdClient) NewListPager(resourceGroupName string, options *PrivateLinkForAzureAdClientListOptions) *runtime.Pager[PrivateLinkForAzureAdClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkForAzureAdClientListResponse]{
		More: func(page PrivateLinkForAzureAdClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkForAzureAdClientListResponse) (PrivateLinkForAzureAdClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkForAzureAdClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateLinkForAzureAdClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkForAzureAdClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinkForAzureAdClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkForAzureAdClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.aadiam/privateLinkForAzureAd"
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
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkForAzureAdClient) listHandleResponse(resp *http.Response) (PrivateLinkForAzureAdClientListResponse, error) {
	result := PrivateLinkForAzureAdClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkPolicyListResult); err != nil {
		return PrivateLinkForAzureAdClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all Private Link Policies For AzureAD in the given subscription.
//
// Generated from API version 2020-03-01
//   - options - PrivateLinkForAzureAdClientListBySubscriptionOptions contains the optional parameters for the PrivateLinkForAzureAdClient.NewListBySubscriptionPager
//     method.
func (client *PrivateLinkForAzureAdClient) NewListBySubscriptionPager(options *PrivateLinkForAzureAdClientListBySubscriptionOptions) *runtime.Pager[PrivateLinkForAzureAdClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkForAzureAdClientListBySubscriptionResponse]{
		More: func(page PrivateLinkForAzureAdClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkForAzureAdClientListBySubscriptionResponse) (PrivateLinkForAzureAdClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateLinkForAzureAdClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateLinkForAzureAdClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkForAzureAdClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PrivateLinkForAzureAdClient) listBySubscriptionCreateRequest(ctx context.Context, options *PrivateLinkForAzureAdClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/microsoft.aadiam/privateLinkForAzureAd"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PrivateLinkForAzureAdClient) listBySubscriptionHandleResponse(resp *http.Response) (PrivateLinkForAzureAdClientListBySubscriptionResponse, error) {
	result := PrivateLinkForAzureAdClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkPolicyListResult); err != nil {
		return PrivateLinkForAzureAdClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates private link policy tags with specified values.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - Name of an Azure resource group.
//   - policyName - The name of the private link policy in Azure AD.
//   - options - PrivateLinkForAzureAdClientUpdateOptions contains the optional parameters for the PrivateLinkForAzureAdClient.Update
//     method.
func (client *PrivateLinkForAzureAdClient) Update(ctx context.Context, resourceGroupName string, policyName string, options *PrivateLinkForAzureAdClientUpdateOptions) (PrivateLinkForAzureAdClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, policyName, options)
	if err != nil {
		return PrivateLinkForAzureAdClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkForAzureAdClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkForAzureAdClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PrivateLinkForAzureAdClient) updateCreateRequest(ctx context.Context, resourceGroupName string, policyName string, options *PrivateLinkForAzureAdClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/microsoft.aadiam/privateLinkForAzureAd/{policyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.PrivateLinkPolicy != nil {
		return req, runtime.MarshalAsJSON(req, *options.PrivateLinkPolicy)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *PrivateLinkForAzureAdClient) updateHandleResponse(resp *http.Response) (PrivateLinkForAzureAdClientUpdateResponse, error) {
	result := PrivateLinkForAzureAdClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkPolicy); err != nil {
		return PrivateLinkForAzureAdClientUpdateResponse{}, err
	}
	return result, nil
}
