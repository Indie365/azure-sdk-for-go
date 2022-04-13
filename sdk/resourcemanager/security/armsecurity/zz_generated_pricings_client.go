//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// PricingsClient contains the methods for the Pricings group.
// Don't use this type directly, use NewPricingsClient() instead.
type PricingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPricingsClient creates a new instance of PricingsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPricingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PricingsClient, error) {
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
	client := &PricingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a provided Security Center pricing configuration in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// pricingName - name of the pricing configuration
// options - PricingsClientGetOptions contains the optional parameters for the PricingsClient.Get method.
func (client *PricingsClient) Get(ctx context.Context, pricingName string, options *PricingsClientGetOptions) (PricingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, pricingName, options)
	if err != nil {
		return PricingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PricingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PricingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PricingsClient) getCreateRequest(ctx context.Context, pricingName string, options *PricingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/pricings/{pricingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if pricingName == "" {
		return nil, errors.New("parameter pricingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pricingName}", url.PathEscape(pricingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PricingsClient) getHandleResponse(resp *http.Response) (PricingsClientGetResponse, error) {
	result := PricingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pricing); err != nil {
		return PricingsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists Security Center pricing configurations in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PricingsClientListOptions contains the optional parameters for the PricingsClient.List method.
func (client *PricingsClient) List(ctx context.Context, options *PricingsClientListOptions) (PricingsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return PricingsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PricingsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PricingsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *PricingsClient) listCreateRequest(ctx context.Context, options *PricingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/pricings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PricingsClient) listHandleResponse(resp *http.Response) (PricingsClientListResponse, error) {
	result := PricingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PricingList); err != nil {
		return PricingsClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates a provided Security Center pricing configuration in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// pricingName - name of the pricing configuration
// pricing - Pricing object
// options - PricingsClientUpdateOptions contains the optional parameters for the PricingsClient.Update method.
func (client *PricingsClient) Update(ctx context.Context, pricingName string, pricing Pricing, options *PricingsClientUpdateOptions) (PricingsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, pricingName, pricing, options)
	if err != nil {
		return PricingsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PricingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PricingsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PricingsClient) updateCreateRequest(ctx context.Context, pricingName string, pricing Pricing, options *PricingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/pricings/{pricingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if pricingName == "" {
		return nil, errors.New("parameter pricingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pricingName}", url.PathEscape(pricingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, pricing)
}

// updateHandleResponse handles the Update response.
func (client *PricingsClient) updateHandleResponse(resp *http.Response) (PricingsClientUpdateResponse, error) {
	result := PricingsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pricing); err != nil {
		return PricingsClientUpdateResponse{}, err
	}
	return result, nil
}
