//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblockchain

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

// MembersClient contains the methods for the BlockchainMembers group.
// Don't use this type directly, use NewMembersClient() instead.
type MembersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMembersClient creates a new instance of MembersClient with the specified values.
// subscriptionID - Gets the subscription Id which uniquely identifies the Microsoft Azure subscription. The subscription
// ID is part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMembersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MembersClient, error) {
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
	client := &MembersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Create a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientBeginCreateOptions contains the optional parameters for the MembersClient.BeginCreate method.
func (client *MembersClient) BeginCreate(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientBeginCreateOptions) (*runtime.Poller[MembersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, blockchainMemberName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[MembersClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[MembersClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Create a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
func (client *MembersClient) create(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *MembersClient) createCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.BlockchainMember != nil {
		return req, runtime.MarshalAsJSON(req, *options.BlockchainMember)
	}
	return req, nil
}

// BeginDelete - Delete a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientBeginDeleteOptions contains the optional parameters for the MembersClient.BeginDelete method.
func (client *MembersClient) BeginDelete(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientBeginDeleteOptions) (*runtime.Poller[MembersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, blockchainMemberName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[MembersClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[MembersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
func (client *MembersClient) deleteOperation(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MembersClient) deleteCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get details about a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientGetOptions contains the optional parameters for the MembersClient.Get method.
func (client *MembersClient) Get(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientGetOptions) (MembersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
	if err != nil {
		return MembersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MembersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MembersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MembersClient) getCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
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
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MembersClient) getHandleResponse(resp *http.Response) (MembersClientGetResponse, error) {
	result := MembersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Member); err != nil {
		return MembersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the blockchain members for a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientListOptions contains the optional parameters for the MembersClient.List method.
func (client *MembersClient) NewListPager(resourceGroupName string, options *MembersClientListOptions) *runtime.Pager[MembersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MembersClientListResponse]{
		More: func(page MembersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MembersClientListResponse) (MembersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MembersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MembersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MembersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MembersClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *MembersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers"
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
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MembersClient) listHandleResponse(resp *http.Response) (MembersClientListResponse, error) {
	result := MembersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MemberCollection); err != nil {
		return MembersClientListResponse{}, err
	}
	return result, nil
}

// ListAPIKeys - Lists the API keys for a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientListAPIKeysOptions contains the optional parameters for the MembersClient.ListAPIKeys method.
func (client *MembersClient) ListAPIKeys(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientListAPIKeysOptions) (MembersClientListAPIKeysResponse, error) {
	req, err := client.listAPIKeysCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
	if err != nil {
		return MembersClientListAPIKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MembersClientListAPIKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MembersClientListAPIKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listAPIKeysHandleResponse(resp)
}

// listAPIKeysCreateRequest creates the ListAPIKeys request.
func (client *MembersClient) listAPIKeysCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientListAPIKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/listApiKeys"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAPIKeysHandleResponse handles the ListAPIKeys response.
func (client *MembersClient) listAPIKeysHandleResponse(resp *http.Response) (MembersClientListAPIKeysResponse, error) {
	result := MembersClientListAPIKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyCollection); err != nil {
		return MembersClientListAPIKeysResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Lists the blockchain members for a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// options - MembersClientListAllOptions contains the optional parameters for the MembersClient.ListAll method.
func (client *MembersClient) NewListAllPager(options *MembersClientListAllOptions) *runtime.Pager[MembersClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[MembersClientListAllResponse]{
		More: func(page MembersClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MembersClientListAllResponse) (MembersClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MembersClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MembersClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MembersClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *MembersClient) listAllCreateRequest(ctx context.Context, options *MembersClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/blockchainMembers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *MembersClient) listAllHandleResponse(resp *http.Response) (MembersClientListAllResponse, error) {
	result := MembersClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MemberCollection); err != nil {
		return MembersClientListAllResponse{}, err
	}
	return result, nil
}

// NewListConsortiumMembersPager - Lists the consortium members for a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientListConsortiumMembersOptions contains the optional parameters for the MembersClient.ListConsortiumMembers
// method.
func (client *MembersClient) NewListConsortiumMembersPager(blockchainMemberName string, resourceGroupName string, options *MembersClientListConsortiumMembersOptions) *runtime.Pager[MembersClientListConsortiumMembersResponse] {
	return runtime.NewPager(runtime.PagingHandler[MembersClientListConsortiumMembersResponse]{
		More: func(page MembersClientListConsortiumMembersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MembersClientListConsortiumMembersResponse) (MembersClientListConsortiumMembersResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listConsortiumMembersCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MembersClientListConsortiumMembersResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MembersClientListConsortiumMembersResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MembersClientListConsortiumMembersResponse{}, runtime.NewResponseError(resp)
			}
			return client.listConsortiumMembersHandleResponse(resp)
		},
	})
}

// listConsortiumMembersCreateRequest creates the ListConsortiumMembers request.
func (client *MembersClient) listConsortiumMembersCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientListConsortiumMembersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/consortiumMembers"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
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
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConsortiumMembersHandleResponse handles the ListConsortiumMembers response.
func (client *MembersClient) listConsortiumMembersHandleResponse(resp *http.Response) (MembersClientListConsortiumMembersResponse, error) {
	result := MembersClientListConsortiumMembersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsortiumMemberCollection); err != nil {
		return MembersClientListConsortiumMembersResponse{}, err
	}
	return result, nil
}

// ListRegenerateAPIKeys - Regenerate the API keys for a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientListRegenerateAPIKeysOptions contains the optional parameters for the MembersClient.ListRegenerateAPIKeys
// method.
func (client *MembersClient) ListRegenerateAPIKeys(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientListRegenerateAPIKeysOptions) (MembersClientListRegenerateAPIKeysResponse, error) {
	req, err := client.listRegenerateAPIKeysCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
	if err != nil {
		return MembersClientListRegenerateAPIKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MembersClientListRegenerateAPIKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MembersClientListRegenerateAPIKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listRegenerateAPIKeysHandleResponse(resp)
}

// listRegenerateAPIKeysCreateRequest creates the ListRegenerateAPIKeys request.
func (client *MembersClient) listRegenerateAPIKeysCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientListRegenerateAPIKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}/regenerateApiKeys"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.APIKey != nil {
		return req, runtime.MarshalAsJSON(req, *options.APIKey)
	}
	return req, nil
}

// listRegenerateAPIKeysHandleResponse handles the ListRegenerateAPIKeys response.
func (client *MembersClient) listRegenerateAPIKeysHandleResponse(resp *http.Response) (MembersClientListRegenerateAPIKeysResponse, error) {
	result := MembersClientListRegenerateAPIKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIKeyCollection); err != nil {
		return MembersClientListRegenerateAPIKeysResponse{}, err
	}
	return result, nil
}

// Update - Update a blockchain member.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01-preview
// blockchainMemberName - Blockchain member name.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// options - MembersClientUpdateOptions contains the optional parameters for the MembersClient.Update method.
func (client *MembersClient) Update(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientUpdateOptions) (MembersClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, blockchainMemberName, resourceGroupName, options)
	if err != nil {
		return MembersClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MembersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MembersClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *MembersClient) updateCreateRequest(ctx context.Context, blockchainMemberName string, resourceGroupName string, options *MembersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Blockchain/blockchainMembers/{blockchainMemberName}"
	if blockchainMemberName == "" {
		return nil, errors.New("parameter blockchainMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blockchainMemberName}", url.PathEscape(blockchainMemberName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.BlockchainMember != nil {
		return req, runtime.MarshalAsJSON(req, *options.BlockchainMember)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MembersClient) updateHandleResponse(resp *http.Response) (MembersClientUpdateResponse, error) {
	result := MembersClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Member); err != nil {
		return MembersClientUpdateResponse{}, err
	}
	return result, nil
}