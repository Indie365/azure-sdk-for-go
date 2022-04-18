//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationsClient, error) {
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
	client := &LocationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckQuotaAvailability - Return quota for subscription by region
// If the operation fails it returns an *azcore.ResponseError type.
// location - Azure region
// options - LocationsClientCheckQuotaAvailabilityOptions contains the optional parameters for the LocationsClient.CheckQuotaAvailability
// method.
func (client *LocationsClient) CheckQuotaAvailability(ctx context.Context, location string, options *LocationsClientCheckQuotaAvailabilityOptions) (LocationsClientCheckQuotaAvailabilityResponse, error) {
	req, err := client.checkQuotaAvailabilityCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsClientCheckQuotaAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsClientCheckQuotaAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsClientCheckQuotaAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkQuotaAvailabilityHandleResponse(resp)
}

// checkQuotaAvailabilityCreateRequest creates the CheckQuotaAvailability request.
func (client *LocationsClient) checkQuotaAvailabilityCreateRequest(ctx context.Context, location string, options *LocationsClientCheckQuotaAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AVS/locations/{location}/checkQuotaAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// checkQuotaAvailabilityHandleResponse handles the CheckQuotaAvailability response.
func (client *LocationsClient) checkQuotaAvailabilityHandleResponse(resp *http.Response) (LocationsClientCheckQuotaAvailabilityResponse, error) {
	result := LocationsClientCheckQuotaAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Quota); err != nil {
		return LocationsClientCheckQuotaAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckTrialAvailability - Return trial status for subscription by region
// If the operation fails it returns an *azcore.ResponseError type.
// location - Azure region
// options - LocationsClientCheckTrialAvailabilityOptions contains the optional parameters for the LocationsClient.CheckTrialAvailability
// method.
func (client *LocationsClient) CheckTrialAvailability(ctx context.Context, location string, options *LocationsClientCheckTrialAvailabilityOptions) (LocationsClientCheckTrialAvailabilityResponse, error) {
	req, err := client.checkTrialAvailabilityCreateRequest(ctx, location, options)
	if err != nil {
		return LocationsClientCheckTrialAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsClientCheckTrialAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsClientCheckTrialAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkTrialAvailabilityHandleResponse(resp)
}

// checkTrialAvailabilityCreateRequest creates the CheckTrialAvailability request.
func (client *LocationsClient) checkTrialAvailabilityCreateRequest(ctx context.Context, location string, options *LocationsClientCheckTrialAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AVS/locations/{location}/checkTrialAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// checkTrialAvailabilityHandleResponse handles the CheckTrialAvailability response.
func (client *LocationsClient) checkTrialAvailabilityHandleResponse(resp *http.Response) (LocationsClientCheckTrialAvailabilityResponse, error) {
	result := LocationsClientCheckTrialAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Trial); err != nil {
		return LocationsClientCheckTrialAvailabilityResponse{}, err
	}
	return result, nil
}
