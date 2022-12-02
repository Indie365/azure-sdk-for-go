//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// BillingMetersClient contains the methods for the BillingMeters group.
// Don't use this type directly, use NewBillingMetersClient() instead.
type BillingMetersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewBillingMetersClient creates a new instance of BillingMetersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewBillingMetersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BillingMetersClient, error) {
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
	client := &BillingMetersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get all billingMeters for a location.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01-preview
// location - The name of Azure region.
// options - BillingMetersClientGetOptions contains the optional parameters for the BillingMetersClient.Get method.
func (client *BillingMetersClient) Get(ctx context.Context, location string, options *BillingMetersClientGetOptions) (BillingMetersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, options)
	if err != nil {
		return BillingMetersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BillingMetersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingMetersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BillingMetersClient) getCreateRequest(ctx context.Context, location string, options *BillingMetersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.App/locations/{location}/billingMeters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BillingMetersClient) getHandleResponse(resp *http.Response) (BillingMetersClientGetResponse, error) {
	result := BillingMetersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingMeterCollection); err != nil {
		return BillingMetersClientGetResponse{}, err
	}
	return result, nil
}