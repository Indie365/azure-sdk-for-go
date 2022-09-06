//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// LoadBalancersClient contains the methods for the LoadBalancers group.
// Don't use this type directly, use NewLoadBalancersClient() instead.
type LoadBalancersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLoadBalancersClient creates a new instance of LoadBalancersClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLoadBalancersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancersClient, error) {
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
	client := &LoadBalancersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// parameters - Parameters supplied to the create or update load balancer operation.
// options - LoadBalancersClientBeginCreateOrUpdateOptions contains the optional parameters for the LoadBalancersClient.BeginCreateOrUpdate
// method.
func (client *LoadBalancersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*runtime.Poller[LoadBalancersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LoadBalancersClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *LoadBalancersClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoadBalancersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// options - LoadBalancersClientBeginDeleteOptions contains the optional parameters for the LoadBalancersClient.BeginDelete
// method.
func (client *LoadBalancersClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*runtime.Poller[LoadBalancersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LoadBalancersClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *LoadBalancersClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LoadBalancersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// options - LoadBalancersClientGetOptions contains the optional parameters for the LoadBalancersClient.Get method.
func (client *LoadBalancersClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (LoadBalancersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LoadBalancersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancersClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancersClient) getHandleResponse(resp *http.Response) (LoadBalancersClientGetResponse, error) {
	result := LoadBalancersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancer); err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the load balancers in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// options - LoadBalancersClientListOptions contains the optional parameters for the LoadBalancersClient.List method.
func (client *LoadBalancersClient) NewListPager(resourceGroupName string, options *LoadBalancersClientListOptions) *runtime.Pager[LoadBalancersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancersClientListResponse]{
		More: func(page LoadBalancersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancersClientListResponse) (LoadBalancersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadBalancersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LoadBalancersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadBalancersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *LoadBalancersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers"
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
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancersClient) listHandleResponse(resp *http.Response) (LoadBalancersClientListResponse, error) {
	result := LoadBalancersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerListResult); err != nil {
		return LoadBalancersClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all the load balancers in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// options - LoadBalancersClientListAllOptions contains the optional parameters for the LoadBalancersClient.ListAll method.
func (client *LoadBalancersClient) NewListAllPager(options *LoadBalancersClientListAllOptions) *runtime.Pager[LoadBalancersClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancersClientListAllResponse]{
		More: func(page LoadBalancersClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancersClientListAllResponse) (LoadBalancersClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LoadBalancersClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LoadBalancersClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LoadBalancersClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *LoadBalancersClient) listAllCreateRequest(ctx context.Context, options *LoadBalancersClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *LoadBalancersClient) listAllHandleResponse(resp *http.Response) (LoadBalancersClientListAllResponse, error) {
	result := LoadBalancersClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerListResult); err != nil {
		return LoadBalancersClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates a load balancer tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// parameters - Parameters supplied to update load balancer tags.
// options - LoadBalancersClientBeginUpdateTagsOptions contains the optional parameters for the LoadBalancersClient.BeginUpdateTags
// method.
func (client *LoadBalancersClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersClientBeginUpdateTagsOptions) (*runtime.Poller[LoadBalancersClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, loadBalancerName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LoadBalancersClientUpdateTagsResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LoadBalancersClientUpdateTagsResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdateTags - Updates a load balancer tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-11-01
func (client *LoadBalancersClient) updateTags(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *LoadBalancersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
