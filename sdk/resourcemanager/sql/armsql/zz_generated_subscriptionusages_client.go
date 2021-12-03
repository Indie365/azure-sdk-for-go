//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionUsagesClient contains the methods for the SubscriptionUsages group.
// Don't use this type directly, use NewSubscriptionUsagesClient() instead.
type SubscriptionUsagesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSubscriptionUsagesClient creates a new instance of SubscriptionUsagesClient with the specified values.
func NewSubscriptionUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SubscriptionUsagesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SubscriptionUsagesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a subscription usage metric.
// If the operation fails it returns a generic error.
func (client *SubscriptionUsagesClient) Get(ctx context.Context, locationName string, usageName string, options *SubscriptionUsagesGetOptions) (SubscriptionUsagesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, usageName, options)
	if err != nil {
		return SubscriptionUsagesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionUsagesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionUsagesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubscriptionUsagesClient) getCreateRequest(ctx context.Context, locationName string, usageName string, options *SubscriptionUsagesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages/{usageName}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if usageName == "" {
		return nil, errors.New("parameter usageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{usageName}", url.PathEscape(usageName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionUsagesClient) getHandleResponse(resp *http.Response) (SubscriptionUsagesGetResponse, error) {
	result := SubscriptionUsagesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUsage); err != nil {
		return SubscriptionUsagesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SubscriptionUsagesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByLocation - Gets all subscription usage metrics in a given location.
// If the operation fails it returns a generic error.
func (client *SubscriptionUsagesClient) ListByLocation(locationName string, options *SubscriptionUsagesListByLocationOptions) *SubscriptionUsagesListByLocationPager {
	return &SubscriptionUsagesListByLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLocationCreateRequest(ctx, locationName, options)
		},
		advancer: func(ctx context.Context, resp SubscriptionUsagesListByLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SubscriptionUsageListResult.NextLink)
		},
	}
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *SubscriptionUsagesClient) listByLocationCreateRequest(ctx context.Context, locationName string, options *SubscriptionUsagesListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/usages"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *SubscriptionUsagesClient) listByLocationHandleResponse(resp *http.Response) (SubscriptionUsagesListByLocationResponse, error) {
	result := SubscriptionUsagesListByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionUsageListResult); err != nil {
		return SubscriptionUsagesListByLocationResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *SubscriptionUsagesClient) listByLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
