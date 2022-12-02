//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// LocationBasedRecommendedActionSessionsOperationStatusClient contains the methods for the LocationBasedRecommendedActionSessionsOperationStatus group.
// Don't use this type directly, use NewLocationBasedRecommendedActionSessionsOperationStatusClient() instead.
type LocationBasedRecommendedActionSessionsOperationStatusClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLocationBasedRecommendedActionSessionsOperationStatusClient creates a new instance of LocationBasedRecommendedActionSessionsOperationStatusClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLocationBasedRecommendedActionSessionsOperationStatusClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationBasedRecommendedActionSessionsOperationStatusClient, error) {
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
	client := &LocationBasedRecommendedActionSessionsOperationStatusClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Recommendation action session operation status.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// locationName - The name of the location.
// operationID - The operation identifier.
// options - LocationBasedRecommendedActionSessionsOperationStatusClientGetOptions contains the optional parameters for the
// LocationBasedRecommendedActionSessionsOperationStatusClient.Get method.
func (client *LocationBasedRecommendedActionSessionsOperationStatusClient) Get(ctx context.Context, locationName string, operationID string, options *LocationBasedRecommendedActionSessionsOperationStatusClientGetOptions) (LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, operationID, options)
	if err != nil {
		return LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LocationBasedRecommendedActionSessionsOperationStatusClient) getCreateRequest(ctx context.Context, locationName string, operationID string, options *LocationBasedRecommendedActionSessionsOperationStatusClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/recommendedActionSessionsAzureAsyncOperation/{operationId}"
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

// getHandleResponse handles the Get response.
func (client *LocationBasedRecommendedActionSessionsOperationStatusClient) getHandleResponse(resp *http.Response) (LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse, error) {
	result := LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendedActionSessionsOperationStatus); err != nil {
		return LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse{}, err
	}
	return result, nil
}