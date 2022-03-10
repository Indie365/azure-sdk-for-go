//go:build go1.16
// +build go1.16

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
func NewLocationBasedRecommendedActionSessionsResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LocationBasedRecommendedActionSessionsResultClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &LocationBasedRecommendedActionSessionsResultClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Recommendation action session operation result.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The name of the location.
// operationID - The operation identifier.
// options - LocationBasedRecommendedActionSessionsResultClientListOptions contains the optional parameters for the LocationBasedRecommendedActionSessionsResultClient.List
// method.
func (client *LocationBasedRecommendedActionSessionsResultClient) List(locationName string, operationID string, options *LocationBasedRecommendedActionSessionsResultClientListOptions) *LocationBasedRecommendedActionSessionsResultClientListPager {
	return &LocationBasedRecommendedActionSessionsResultClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, locationName, operationID, options)
		},
		advancer: func(ctx context.Context, resp LocationBasedRecommendedActionSessionsResultClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RecommendationActionsResultList.NextLink)
		},
	}
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
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationBasedRecommendedActionSessionsResultClient) listHandleResponse(resp *http.Response) (LocationBasedRecommendedActionSessionsResultClientListResponse, error) {
	result := LocationBasedRecommendedActionSessionsResultClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendationActionsResultList); err != nil {
		return LocationBasedRecommendedActionSessionsResultClientListResponse{}, err
	}
	return result, nil
}
