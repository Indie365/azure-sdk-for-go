//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdnsresolver

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
	"strconv"
	"strings"
)

// OutboundEndpointsClient contains the methods for the OutboundEndpoints group.
// Don't use this type directly, use NewOutboundEndpointsClient() instead.
type OutboundEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOutboundEndpointsClient creates a new instance of OutboundEndpointsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOutboundEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OutboundEndpointsClient, error) {
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
	client := &OutboundEndpointsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an outbound endpoint for a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// outboundEndpointName - The name of the outbound endpoint for the DNS resolver.
// parameters - Parameters supplied to the CreateOrUpdate operation.
// options - OutboundEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the OutboundEndpointsClient.BeginCreateOrUpdate
// method.
func (client *OutboundEndpointsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, parameters OutboundEndpoint, options *OutboundEndpointsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[OutboundEndpointsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, dnsResolverName, outboundEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[OutboundEndpointsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[OutboundEndpointsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates an outbound endpoint for a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OutboundEndpointsClient) createOrUpdate(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, parameters OutboundEndpoint, options *OutboundEndpointsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dnsResolverName, outboundEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OutboundEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, parameters OutboundEndpoint, options *OutboundEndpointsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	if outboundEndpointName == "" {
		return nil, errors.New("parameter outboundEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundEndpointName}", url.PathEscape(outboundEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an outbound endpoint for a DNS resolver. WARNING: This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// outboundEndpointName - The name of the outbound endpoint for the DNS resolver.
// options - OutboundEndpointsClientBeginDeleteOptions contains the optional parameters for the OutboundEndpointsClient.BeginDelete
// method.
func (client *OutboundEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, options *OutboundEndpointsClientBeginDeleteOptions) (*armruntime.Poller[OutboundEndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dnsResolverName, outboundEndpointName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[OutboundEndpointsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[OutboundEndpointsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an outbound endpoint for a DNS resolver. WARNING: This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OutboundEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, options *OutboundEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dnsResolverName, outboundEndpointName, options)
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
func (client *OutboundEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, options *OutboundEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	if outboundEndpointName == "" {
		return nil, errors.New("parameter outboundEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundEndpointName}", url.PathEscape(outboundEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets properties of an outbound endpoint for a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// outboundEndpointName - The name of the outbound endpoint for the DNS resolver.
// options - OutboundEndpointsClientGetOptions contains the optional parameters for the OutboundEndpointsClient.Get method.
func (client *OutboundEndpointsClient) Get(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, options *OutboundEndpointsClientGetOptions) (OutboundEndpointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dnsResolverName, outboundEndpointName, options)
	if err != nil {
		return OutboundEndpointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutboundEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OutboundEndpointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OutboundEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, options *OutboundEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	if outboundEndpointName == "" {
		return nil, errors.New("parameter outboundEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundEndpointName}", url.PathEscape(outboundEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OutboundEndpointsClient) getHandleResponse(resp *http.Response) (OutboundEndpointsClientGetResponse, error) {
	result := OutboundEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEndpoint); err != nil {
		return OutboundEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists outbound endpoints for a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// options - OutboundEndpointsClientListOptions contains the optional parameters for the OutboundEndpointsClient.List method.
func (client *OutboundEndpointsClient) NewListPager(resourceGroupName string, dnsResolverName string, options *OutboundEndpointsClientListOptions) *runtime.Pager[OutboundEndpointsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[OutboundEndpointsClientListResponse]{
		More: func(page OutboundEndpointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OutboundEndpointsClientListResponse) (OutboundEndpointsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, dnsResolverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OutboundEndpointsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OutboundEndpointsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OutboundEndpointsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OutboundEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, options *OutboundEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OutboundEndpointsClient) listHandleResponse(resp *http.Response) (OutboundEndpointsClientListResponse, error) {
	result := OutboundEndpointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEndpointListResult); err != nil {
		return OutboundEndpointsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an outbound endpoint for a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dnsResolverName - The name of the DNS resolver.
// outboundEndpointName - The name of the outbound endpoint for the DNS resolver.
// parameters - Parameters supplied to the Update operation.
// options - OutboundEndpointsClientBeginUpdateOptions contains the optional parameters for the OutboundEndpointsClient.BeginUpdate
// method.
func (client *OutboundEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, parameters OutboundEndpointPatch, options *OutboundEndpointsClientBeginUpdateOptions) (*armruntime.Poller[OutboundEndpointsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, dnsResolverName, outboundEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[OutboundEndpointsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[OutboundEndpointsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Updates an outbound endpoint for a DNS resolver.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OutboundEndpointsClient) update(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, parameters OutboundEndpointPatch, options *OutboundEndpointsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dnsResolverName, outboundEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OutboundEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dnsResolverName string, outboundEndpointName string, parameters OutboundEndpointPatch, options *OutboundEndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dnsResolverName == "" {
		return nil, errors.New("parameter dnsResolverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsResolverName}", url.PathEscape(dnsResolverName))
	if outboundEndpointName == "" {
		return nil, errors.New("parameter outboundEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outboundEndpointName}", url.PathEscape(outboundEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
