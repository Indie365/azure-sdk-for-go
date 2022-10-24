//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/profiles/hybrid20200901/internal"
	"net/http"
	"strconv"
	"strings"
)

// MetricsClient contains the methods for the Metrics group.
// Don't use this type directly, use NewMetricsClient() instead.
type MetricsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewMetricsClient creates a new instance of MetricsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMetricsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*MetricsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(hybrid20200901.ModuleName, hybrid20200901.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &MetricsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// List - Lists the metric values for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-01-01
// resourceURI - The identifier of the resource.
// options - MetricsClientListOptions contains the optional parameters for the MetricsClient.List method.
func (client *MetricsClient) List(ctx context.Context, resourceURI string, options *MetricsClientListOptions) (MetricsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return MetricsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MetricsClient) listCreateRequest(ctx context.Context, resourceURI string, options *MetricsClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metrics"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("orderby", *options.Orderby)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	reqQP.Set("api-version", "2018-01-01")
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricsClient) listHandleResponse(resp *http.Response) (MetricsClientListResponse, error) {
	result := MetricsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Response); err != nil {
		return MetricsClientListResponse{}, err
	}
	return result, nil
}
