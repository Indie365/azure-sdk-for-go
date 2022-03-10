//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscriptions

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// TenantsClient contains the methods for the Tenants group.
// Don't use this type directly, use NewTenantsClient() instead.
type TenantsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewTenantsClient creates a new instance of TenantsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTenantsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *TenantsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &TenantsClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets the tenants for your account.
// If the operation fails it returns an *azcore.ResponseError type.
// options - TenantsClientListOptions contains the optional parameters for the TenantsClient.List method.
func (client *TenantsClient) List(options *TenantsClientListOptions) *TenantsClientListPager {
	return &TenantsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp TenantsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.TenantListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *TenantsClient) listCreateRequest(ctx context.Context, options *TenantsClientListOptions) (*policy.Request, error) {
	urlPath := "/tenants"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TenantsClient) listHandleResponse(resp *http.Response) (TenantsClientListResponse, error) {
	result := TenantsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantListResult); err != nil {
		return TenantsClientListResponse{}, err
	}
	return result, nil
}
