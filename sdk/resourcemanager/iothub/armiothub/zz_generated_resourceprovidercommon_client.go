//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

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

// ResourceProviderCommonClient contains the methods for the ResourceProviderCommon group.
// Don't use this type directly, use NewResourceProviderCommonClient() instead.
type ResourceProviderCommonClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceProviderCommonClient creates a new instance of ResourceProviderCommonClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceProviderCommonClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourceProviderCommonClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ResourceProviderCommonClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// GetSubscriptionQuota - Get the number of free and paid iot hubs in the subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - ResourceProviderCommonClientGetSubscriptionQuotaOptions contains the optional parameters for the ResourceProviderCommonClient.GetSubscriptionQuota
// method.
func (client *ResourceProviderCommonClient) GetSubscriptionQuota(ctx context.Context, options *ResourceProviderCommonClientGetSubscriptionQuotaOptions) (ResourceProviderCommonClientGetSubscriptionQuotaResponse, error) {
	req, err := client.getSubscriptionQuotaCreateRequest(ctx, options)
	if err != nil {
		return ResourceProviderCommonClientGetSubscriptionQuotaResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceProviderCommonClientGetSubscriptionQuotaResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceProviderCommonClientGetSubscriptionQuotaResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSubscriptionQuotaHandleResponse(resp)
}

// getSubscriptionQuotaCreateRequest creates the GetSubscriptionQuota request.
func (client *ResourceProviderCommonClient) getSubscriptionQuotaCreateRequest(ctx context.Context, options *ResourceProviderCommonClientGetSubscriptionQuotaOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Devices/usages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSubscriptionQuotaHandleResponse handles the GetSubscriptionQuota response.
func (client *ResourceProviderCommonClient) getSubscriptionQuotaHandleResponse(resp *http.Response) (ResourceProviderCommonClientGetSubscriptionQuotaResponse, error) {
	result := ResourceProviderCommonClientGetSubscriptionQuotaResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserSubscriptionQuotaListResult); err != nil {
		return ResourceProviderCommonClientGetSubscriptionQuotaResponse{}, err
	}
	return result, nil
}
