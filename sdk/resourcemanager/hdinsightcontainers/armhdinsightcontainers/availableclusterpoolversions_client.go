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

// AvailableClusterPoolVersionsClient contains the methods for the AvailableClusterPoolVersions group.
// Don't use this type directly, use NewAvailableClusterPoolVersionsClient() instead.
type AvailableClusterPoolVersionsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewAvailableClusterPoolVersionsClient creates a new instance of AvailableClusterPoolVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailableClusterPoolVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableClusterPoolVersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".AvailableClusterPoolVersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableClusterPoolVersionsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// NewListByLocationPager - Returns a list of available cluster pool versions.
//
// Generated from API version 2023-06-01-preview
//   - location - The name of the Azure region.
//   - options - AvailableClusterPoolVersionsClientListByLocationOptions contains the optional parameters for the AvailableClusterPoolVersionsClient.NewListByLocationPager
//     method.
func (client *AvailableClusterPoolVersionsClient) NewListByLocationPager(location string, options *AvailableClusterPoolVersionsClientListByLocationOptions) (*runtime.Pager[AvailableClusterPoolVersionsClientListByLocationResponse]) {
	return runtime.NewPager(runtime.PagingHandler[AvailableClusterPoolVersionsClientListByLocationResponse]{
		More: func(page AvailableClusterPoolVersionsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableClusterPoolVersionsClientListByLocationResponse) (AvailableClusterPoolVersionsClientListByLocationResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByLocationCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailableClusterPoolVersionsClientListByLocationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AvailableClusterPoolVersionsClientListByLocationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailableClusterPoolVersionsClientListByLocationResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByLocationHandleResponse(resp)
		},
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *AvailableClusterPoolVersionsClient) listByLocationCreateRequest(ctx context.Context, location string, options *AvailableClusterPoolVersionsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/locations/{location}/availableClusterPoolVersions"
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
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *AvailableClusterPoolVersionsClient) listByLocationHandleResponse(resp *http.Response) (AvailableClusterPoolVersionsClientListByLocationResponse, error) {
	result := AvailableClusterPoolVersionsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterPoolVersionsListResult); err != nil {
		return AvailableClusterPoolVersionsClientListByLocationResponse{}, err
	}
	return result, nil
}

