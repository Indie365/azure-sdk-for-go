//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

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

// ExternalSecuritySolutionsClient contains the methods for the ExternalSecuritySolutions group.
// Don't use this type directly, use NewExternalSecuritySolutionsClient() instead.
type ExternalSecuritySolutionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExternalSecuritySolutionsClient creates a new instance of ExternalSecuritySolutionsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExternalSecuritySolutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExternalSecuritySolutionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ExternalSecuritySolutionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExternalSecuritySolutionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a specific external Security Solution.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - externalSecuritySolutionsName - Name of an external security solution.
//   - options - ExternalSecuritySolutionsClientGetOptions contains the optional parameters for the ExternalSecuritySolutionsClient.Get
//     method.
func (client *ExternalSecuritySolutionsClient) Get(ctx context.Context, resourceGroupName string, ascLocation string, externalSecuritySolutionsName string, options *ExternalSecuritySolutionsClientGetOptions) (ExternalSecuritySolutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ascLocation, externalSecuritySolutionsName, options)
	if err != nil {
		return ExternalSecuritySolutionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExternalSecuritySolutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExternalSecuritySolutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExternalSecuritySolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ascLocation string, externalSecuritySolutionsName string, options *ExternalSecuritySolutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions/{externalSecuritySolutionsName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if externalSecuritySolutionsName == "" {
		return nil, errors.New("parameter externalSecuritySolutionsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalSecuritySolutionsName}", url.PathEscape(externalSecuritySolutionsName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExternalSecuritySolutionsClient) getHandleResponse(resp *http.Response) (ExternalSecuritySolutionsClientGetResponse, error) {
	result := ExternalSecuritySolutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalSecuritySolution); err != nil {
		return ExternalSecuritySolutionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of external security solutions for the subscription.
//
// Generated from API version 2020-01-01
//   - options - ExternalSecuritySolutionsClientListOptions contains the optional parameters for the ExternalSecuritySolutionsClient.NewListPager
//     method.
func (client *ExternalSecuritySolutionsClient) NewListPager(options *ExternalSecuritySolutionsClientListOptions) *runtime.Pager[ExternalSecuritySolutionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExternalSecuritySolutionsClientListResponse]{
		More: func(page ExternalSecuritySolutionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExternalSecuritySolutionsClientListResponse) (ExternalSecuritySolutionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExternalSecuritySolutionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExternalSecuritySolutionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExternalSecuritySolutionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExternalSecuritySolutionsClient) listCreateRequest(ctx context.Context, options *ExternalSecuritySolutionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/externalSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExternalSecuritySolutionsClient) listHandleResponse(resp *http.Response) (ExternalSecuritySolutionsClientListResponse, error) {
	result := ExternalSecuritySolutionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalSecuritySolutionList); err != nil {
		return ExternalSecuritySolutionsClientListResponse{}, err
	}
	return result, nil
}

// NewListByHomeRegionPager - Gets a list of external Security Solutions for the subscription and location.
//
// Generated from API version 2020-01-01
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - options - ExternalSecuritySolutionsClientListByHomeRegionOptions contains the optional parameters for the ExternalSecuritySolutionsClient.NewListByHomeRegionPager
//     method.
func (client *ExternalSecuritySolutionsClient) NewListByHomeRegionPager(ascLocation string, options *ExternalSecuritySolutionsClientListByHomeRegionOptions) *runtime.Pager[ExternalSecuritySolutionsClientListByHomeRegionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExternalSecuritySolutionsClientListByHomeRegionResponse]{
		More: func(page ExternalSecuritySolutionsClientListByHomeRegionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExternalSecuritySolutionsClientListByHomeRegionResponse) (ExternalSecuritySolutionsClientListByHomeRegionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHomeRegionCreateRequest(ctx, ascLocation, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExternalSecuritySolutionsClientListByHomeRegionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExternalSecuritySolutionsClientListByHomeRegionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExternalSecuritySolutionsClientListByHomeRegionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHomeRegionHandleResponse(resp)
		},
	})
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *ExternalSecuritySolutionsClient) listByHomeRegionCreateRequest(ctx context.Context, ascLocation string, options *ExternalSecuritySolutionsClientListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *ExternalSecuritySolutionsClient) listByHomeRegionHandleResponse(resp *http.Response) (ExternalSecuritySolutionsClientListByHomeRegionResponse, error) {
	result := ExternalSecuritySolutionsClientListByHomeRegionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalSecuritySolutionList); err != nil {
		return ExternalSecuritySolutionsClientListByHomeRegionResponse{}, err
	}
	return result, nil
}
