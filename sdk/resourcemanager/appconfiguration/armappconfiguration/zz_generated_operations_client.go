//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

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

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
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
	client := &OperationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Checks whether the configuration store name is available for use.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// checkNameAvailabilityParameters - The object containing information for the availability request.
// options - OperationsClientCheckNameAvailabilityOptions contains the optional parameters for the OperationsClient.CheckNameAvailability
// method.
func (client *OperationsClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *OperationsClientCheckNameAvailabilityOptions) (OperationsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityParameters, options)
	if err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *OperationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *OperationsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityParameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *OperationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (OperationsClientCheckNameAvailabilityResponse, error) {
	result := OperationsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityStatus); err != nil {
		return OperationsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the operations available from this provider.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// options - OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
func (client *OperationsClient) NewListPager(options *OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationsClientListResponse]{
		More: func(page OperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationsClientListResponse) (OperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AppConfiguration/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationDefinitionListResult); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}

// RegionalCheckNameAvailability - Checks whether the configuration store name is available for use.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// location - The location in which uniqueness will be verified.
// checkNameAvailabilityParameters - The object containing information for the availability request.
// options - OperationsClientRegionalCheckNameAvailabilityOptions contains the optional parameters for the OperationsClient.RegionalCheckNameAvailability
// method.
func (client *OperationsClient) RegionalCheckNameAvailability(ctx context.Context, location string, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *OperationsClientRegionalCheckNameAvailabilityOptions) (OperationsClientRegionalCheckNameAvailabilityResponse, error) {
	req, err := client.regionalCheckNameAvailabilityCreateRequest(ctx, location, checkNameAvailabilityParameters, options)
	if err != nil {
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.regionalCheckNameAvailabilityHandleResponse(resp)
}

// regionalCheckNameAvailabilityCreateRequest creates the RegionalCheckNameAvailability request.
func (client *OperationsClient) regionalCheckNameAvailabilityCreateRequest(ctx context.Context, location string, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *OperationsClientRegionalCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AppConfiguration/locations/{location}/checkNameAvailability"
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
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, checkNameAvailabilityParameters)
}

// regionalCheckNameAvailabilityHandleResponse handles the RegionalCheckNameAvailability response.
func (client *OperationsClient) regionalCheckNameAvailabilityHandleResponse(resp *http.Response) (OperationsClientRegionalCheckNameAvailabilityResponse, error) {
	result := OperationsClientRegionalCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailabilityStatus); err != nil {
		return OperationsClientRegionalCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}