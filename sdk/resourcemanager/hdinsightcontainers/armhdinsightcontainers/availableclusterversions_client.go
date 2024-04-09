//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

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

// AvailableClusterVersionsClient contains the methods for the AvailableClusterVersions group.
// Don't use this type directly, use NewAvailableClusterVersionsClient() instead.
type AvailableClusterVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAvailableClusterVersionsClient creates a new instance of AvailableClusterVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailableClusterVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableClusterVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableClusterVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByLocationPager - Returns a list of available cluster versions.
//
// Generated from API version 2023-11-01-preview
//   - location - The name of the Azure region.
//   - options - AvailableClusterVersionsClientListByLocationOptions contains the optional parameters for the AvailableClusterVersionsClient.NewListByLocationPager
//     method.
func (client *AvailableClusterVersionsClient) NewListByLocationPager(location string, options *AvailableClusterVersionsClientListByLocationOptions) *runtime.Pager[AvailableClusterVersionsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailableClusterVersionsClientListByLocationResponse]{
		More: func(page AvailableClusterVersionsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableClusterVersionsClientListByLocationResponse) (AvailableClusterVersionsClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvailableClusterVersionsClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return AvailableClusterVersionsClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *AvailableClusterVersionsClient) listByLocationCreateRequest(ctx context.Context, location string, options *AvailableClusterVersionsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/availableClusterVersions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *AvailableClusterVersionsClient) listByLocationHandleResponse(resp *http.Response) (AvailableClusterVersionsClientListByLocationResponse, error) {
	result := AvailableClusterVersionsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterVersionsListResult); err != nil {
		return AvailableClusterVersionsClientListByLocationResponse{}, err
	}
	return result, nil
}
