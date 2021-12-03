//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// CollectionPartitionClient contains the methods for the CollectionPartition group.
// Don't use this type directly, use NewCollectionPartitionClient() instead.
type CollectionPartitionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCollectionPartitionClient creates a new instance of CollectionPartitionClient with the specified values.
func NewCollectionPartitionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CollectionPartitionClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CollectionPartitionClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListMetrics - Retrieves the metrics determined by the given filter for the given collection, split by partition.
// If the operation fails it returns a generic error.
func (client *CollectionPartitionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionListMetricsOptions) (CollectionPartitionListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter, options)
	if err != nil {
		return CollectionPartitionListMetricsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CollectionPartitionListMetricsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CollectionPartitionListMetricsResponse{}, client.listMetricsHandleError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionPartitionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics"
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
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionPartitionClient) listMetricsHandleResponse(resp *http.Response) (CollectionPartitionListMetricsResponse, error) {
	result := CollectionPartitionListMetricsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartitionMetricListResult); err != nil {
		return CollectionPartitionListMetricsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listMetricsHandleError handles the ListMetrics error response.
func (client *CollectionPartitionClient) listMetricsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListUsages - Retrieves the usages (most recent storage data) for the given collection, split by partition.
// If the operation fails it returns a generic error.
func (client *CollectionPartitionClient) ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionPartitionListUsagesOptions) (CollectionPartitionListUsagesResponse, error) {
	req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, options)
	if err != nil {
		return CollectionPartitionListUsagesResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CollectionPartitionListUsagesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CollectionPartitionListUsagesResponse{}, client.listUsagesHandleError(resp)
	}
	return client.listUsagesHandleResponse(resp)
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *CollectionPartitionClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionPartitionListUsagesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/usages"
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
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *CollectionPartitionClient) listUsagesHandleResponse(resp *http.Response) (CollectionPartitionListUsagesResponse, error) {
	result := CollectionPartitionListUsagesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartitionUsagesResult); err != nil {
		return CollectionPartitionListUsagesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listUsagesHandleError handles the ListUsages error response.
func (client *CollectionPartitionClient) listUsagesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
