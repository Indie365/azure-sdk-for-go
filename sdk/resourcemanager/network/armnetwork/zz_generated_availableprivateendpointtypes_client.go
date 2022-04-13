//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// AvailablePrivateEndpointTypesClient contains the methods for the AvailablePrivateEndpointTypes group.
// Don't use this type directly, use NewAvailablePrivateEndpointTypesClient() instead.
type AvailablePrivateEndpointTypesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailablePrivateEndpointTypesClient creates a new instance of AvailablePrivateEndpointTypesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailablePrivateEndpointTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailablePrivateEndpointTypesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AvailablePrivateEndpointTypesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in this region.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location of the domain name.
// options - AvailablePrivateEndpointTypesClientListOptions contains the optional parameters for the AvailablePrivateEndpointTypesClient.List
// method.
func (client *AvailablePrivateEndpointTypesClient) List(location string, options *AvailablePrivateEndpointTypesClientListOptions) *runtime.Pager[AvailablePrivateEndpointTypesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[AvailablePrivateEndpointTypesClientListResponse]{
		More: func(page AvailablePrivateEndpointTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailablePrivateEndpointTypesClientListResponse) (AvailablePrivateEndpointTypesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailablePrivateEndpointTypesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AvailablePrivateEndpointTypesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailablePrivateEndpointTypesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AvailablePrivateEndpointTypesClient) listCreateRequest(ctx context.Context, location string, options *AvailablePrivateEndpointTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailablePrivateEndpointTypesClient) listHandleResponse(resp *http.Response) (AvailablePrivateEndpointTypesClientListResponse, error) {
	result := AvailablePrivateEndpointTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailablePrivateEndpointTypesResult); err != nil {
		return AvailablePrivateEndpointTypesClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Returns all of the resource types that can be linked to a Private Endpoint in this subscription in
// this region.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location of the domain name.
// resourceGroupName - The name of the resource group.
// options - AvailablePrivateEndpointTypesClientListByResourceGroupOptions contains the optional parameters for the AvailablePrivateEndpointTypesClient.ListByResourceGroup
// method.
func (client *AvailablePrivateEndpointTypesClient) ListByResourceGroup(location string, resourceGroupName string, options *AvailablePrivateEndpointTypesClientListByResourceGroupOptions) *runtime.Pager[AvailablePrivateEndpointTypesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[AvailablePrivateEndpointTypesClientListByResourceGroupResponse]{
		More: func(page AvailablePrivateEndpointTypesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailablePrivateEndpointTypesClientListByResourceGroupResponse) (AvailablePrivateEndpointTypesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, location, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailablePrivateEndpointTypesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AvailablePrivateEndpointTypesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailablePrivateEndpointTypesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailablePrivateEndpointTypesClient) listByResourceGroupCreateRequest(ctx context.Context, location string, resourceGroupName string, options *AvailablePrivateEndpointTypesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availablePrivateEndpointTypes"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailablePrivateEndpointTypesClient) listByResourceGroupHandleResponse(resp *http.Response) (AvailablePrivateEndpointTypesClientListByResourceGroupResponse, error) {
	result := AvailablePrivateEndpointTypesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailablePrivateEndpointTypesResult); err != nil {
		return AvailablePrivateEndpointTypesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
