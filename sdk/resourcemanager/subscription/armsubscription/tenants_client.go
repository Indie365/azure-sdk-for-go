//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscription

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// TenantsClient contains the methods for the Tenants group.
// Don't use this type directly, use NewTenantsClient() instead.
type TenantsClient struct {
	internal *arm.Client
}

// NewTenantsClient creates a new instance of TenantsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTenantsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TenantsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Gets the tenants for your account.
//
// Generated from API version 2016-06-01
//   - options - TenantsClientListOptions contains the optional parameters for the TenantsClient.NewListPager method.
func (client *TenantsClient) NewListPager(options *TenantsClientListOptions) *runtime.Pager[TenantsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TenantsClientListResponse]{
		More: func(page TenantsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TenantsClientListResponse) (TenantsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TenantsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TenantsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TenantsClient) listCreateRequest(ctx context.Context, options *TenantsClientListOptions) (*policy.Request, error) {
	urlPath := "/tenants"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TenantsClient) listHandleResponse(resp *http.Response) (TenantsClientListResponse, error) {
	result := TenantsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantListResult); err != nil {
		return TenantsClientListResponse{}, err
	}
	return result, nil
}
