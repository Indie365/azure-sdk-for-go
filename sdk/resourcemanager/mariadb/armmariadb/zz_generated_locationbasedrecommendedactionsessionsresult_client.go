//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

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

// LocationBasedRecommendedActionSessionsResultClient contains the methods for the LocationBasedRecommendedActionSessionsResult group.
// Don't use this type directly, use NewLocationBasedRecommendedActionSessionsResultClient() instead.
type LocationBasedRecommendedActionSessionsResultClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationBasedRecommendedActionSessionsResultClient creates a new instance of LocationBasedRecommendedActionSessionsResultClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationBasedRecommendedActionSessionsResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationBasedRecommendedActionSessionsResultClient, error) {
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
	client := &LocationBasedRecommendedActionSessionsResultClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Recommendation action session operation result.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// locationName - The name of the location.
// operationID - The operation identifier.
// options - LocationBasedRecommendedActionSessionsResultClientListOptions contains the optional parameters for the LocationBasedRecommendedActionSessionsResultClient.List
// method.
func (client *LocationBasedRecommendedActionSessionsResultClient) NewListPager(locationName string, operationID string, options *LocationBasedRecommendedActionSessionsResultClientListOptions) *runtime.Pager[LocationBasedRecommendedActionSessionsResultClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocationBasedRecommendedActionSessionsResultClientListResponse]{
		More: func(page LocationBasedRecommendedActionSessionsResultClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationBasedRecommendedActionSessionsResultClientListResponse) (LocationBasedRecommendedActionSessionsResultClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, locationName, operationID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LocationBasedRecommendedActionSessionsResultClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LocationBasedRecommendedActionSessionsResultClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationBasedRecommendedActionSessionsResultClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LocationBasedRecommendedActionSessionsResultClient) listCreateRequest(ctx context.Context, locationName string, operationID string, options *LocationBasedRecommendedActionSessionsResultClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMariaDB/locations/{locationName}/recommendedActionSessionsOperationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationBasedRecommendedActionSessionsResultClient) listHandleResponse(resp *http.Response) (LocationBasedRecommendedActionSessionsResultClientListResponse, error) {
	result := LocationBasedRecommendedActionSessionsResultClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendationActionsResultList); err != nil {
		return LocationBasedRecommendedActionSessionsResultClientListResponse{}, err
	}
	return result, nil
}