//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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

// RolesClient contains the methods for the Roles group.
// Don't use this type directly, use NewRolesClient() instead.
type RolesClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewRolesClient creates a new instance of RolesClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRolesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RolesClient, error) {
	cl, err := arm.NewClient(moduleName+".RolesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RolesClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// NewListByHubPager - Gets all the roles for the hub.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - options - RolesClientListByHubOptions contains the optional parameters for the RolesClient.NewListByHubPager method.
func (client *RolesClient) NewListByHubPager(resourceGroupName string, hubName string, options *RolesClientListByHubOptions) (*runtime.Pager[RolesClientListByHubResponse]) {
	return runtime.NewPager(runtime.PagingHandler[RolesClientListByHubResponse]{
		More: func(page RolesClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RolesClientListByHubResponse) (RolesClientListByHubResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RolesClientListByHubResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RolesClientListByHubResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RolesClientListByHubResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHubHandleResponse(resp)
		},
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *RolesClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *RolesClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/roles"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *RolesClient) listByHubHandleResponse(resp *http.Response) (RolesClientListByHubResponse, error) {
	result := RolesClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleListResult); err != nil {
		return RolesClientListByHubResponse{}, err
	}
	return result, nil
}

