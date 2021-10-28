//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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
	"strconv"
	"strings"
)

// JobVersionsClient contains the methods for the JobVersions group.
// Don't use this type directly, use NewJobVersionsClient() instead.
type JobVersionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewJobVersionsClient creates a new instance of JobVersionsClient with the specified values.
func NewJobVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *JobVersionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &JobVersionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a job version.
// If the operation fails it returns a generic error.
func (client *JobVersionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, options *JobVersionsGetOptions) (JobVersionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, jobVersion, options)
	if err != nil {
		return JobVersionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return JobVersionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return JobVersionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *JobVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, options *JobVersionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/versions/{jobVersion}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	urlPath = strings.ReplaceAll(urlPath, "{jobVersion}", url.PathEscape(strconv.FormatInt(int64(jobVersion), 10)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobVersionsClient) getHandleResponse(resp *http.Response) (JobVersionsGetResponse, error) {
	result := JobVersionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobVersion); err != nil {
		return JobVersionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *JobVersionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByJob - Gets all versions of a job.
// If the operation fails it returns a generic error.
func (client *JobVersionsClient) ListByJob(resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobVersionsListByJobOptions) *JobVersionsListByJobPager {
	return &JobVersionsListByJobPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByJobCreateRequest(ctx, resourceGroupName, serverName, jobAgentName, jobName, options)
		},
		advancer: func(ctx context.Context, resp JobVersionsListByJobResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobVersionListResult.NextLink)
		},
	}
}

// listByJobCreateRequest creates the ListByJob request.
func (client *JobVersionsClient) listByJobCreateRequest(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *JobVersionsListByJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/jobAgents/{jobAgentName}/jobs/{jobName}/versions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if jobAgentName == "" {
		return nil, errors.New("parameter jobAgentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobAgentName}", url.PathEscape(jobAgentName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByJobHandleResponse handles the ListByJob response.
func (client *JobVersionsClient) listByJobHandleResponse(resp *http.Response) (JobVersionsListByJobResponse, error) {
	result := JobVersionsListByJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobVersionListResult); err != nil {
		return JobVersionsListByJobResponse{}, err
	}
	return result, nil
}

// listByJobHandleError handles the ListByJob error response.
func (client *JobVersionsClient) listByJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
