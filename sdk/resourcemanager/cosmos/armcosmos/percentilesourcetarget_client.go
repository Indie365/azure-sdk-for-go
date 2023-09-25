//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PercentileSourceTargetClient contains the methods for the PercentileSourceTarget group.
// Don't use this type directly, use NewPercentileSourceTargetClient() instead.
type PercentileSourceTargetClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewPercentileSourceTargetClient creates a new instance of PercentileSourceTargetClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPercentileSourceTargetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PercentileSourceTargetClient, error) {
	cl, err := arm.NewClient(moduleName+".PercentileSourceTargetClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PercentileSourceTargetClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// NewListMetricsPager - Retrieves the metrics determined by the given filter for the given account, source and target region.
// This url is only for PBS and Replication Latency data
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - sourceRegion - Source region from which data is written. Cosmos DB region, with spaces between words and each word capitalized.
//   - targetRegion - Target region to which data is written. Cosmos DB region, with spaces between words and each word capitalized.
//   - filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
//     name.value (name of the metric, can have an or of multiple names), startTime, endTime,
//     and timeGrain. The supported operator is eq.
//   - options - PercentileSourceTargetClientListMetricsOptions contains the optional parameters for the PercentileSourceTargetClient.NewListMetricsPager
//     method.
func (client *PercentileSourceTargetClient) NewListMetricsPager(resourceGroupName string, accountName string, sourceRegion string, targetRegion string, filter string, options *PercentileSourceTargetClientListMetricsOptions) (*runtime.Pager[PercentileSourceTargetClientListMetricsResponse]) {
	return runtime.NewPager(runtime.PagingHandler[PercentileSourceTargetClientListMetricsResponse]{
		More: func(page PercentileSourceTargetClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PercentileSourceTargetClientListMetricsResponse) (PercentileSourceTargetClientListMetricsResponse, error) {
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, sourceRegion, targetRegion, filter, options)
			if err != nil {
				return PercentileSourceTargetClientListMetricsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PercentileSourceTargetClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PercentileSourceTargetClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *PercentileSourceTargetClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, sourceRegion string, targetRegion string, filter string, options *PercentileSourceTargetClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sourceRegion/{sourceRegion}/targetRegion/{targetRegion}/percentile/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if sourceRegion == "" {
		return nil, errors.New("parameter sourceRegion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sourceRegion}", url.PathEscape(sourceRegion))
	if targetRegion == "" {
		return nil, errors.New("parameter targetRegion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetRegion}", url.PathEscape(targetRegion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *PercentileSourceTargetClient) listMetricsHandleResponse(resp *http.Response) (PercentileSourceTargetClientListMetricsResponse, error) {
	result := PercentileSourceTargetClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PercentileMetricListResult); err != nil {
		return PercentileSourceTargetClientListMetricsResponse{}, err
	}
	return result, nil
}

