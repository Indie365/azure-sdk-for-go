//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TestJobStreamsClient contains the methods for the TestJobStreams group.
// Don't use this type directly, use NewTestJobStreamsClient() instead.
type TestJobStreamsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewTestJobStreamsClient creates a new instance of TestJobStreamsClient with the specified values.
func NewTestJobStreamsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *TestJobStreamsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &TestJobStreamsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Retrieve a test job stream of the test job identified by runbook name and stream id.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TestJobStreamsClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string, options *TestJobStreamsGetOptions) (TestJobStreamsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, jobStreamID, options)
	if err != nil {
		return TestJobStreamsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TestJobStreamsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TestJobStreamsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TestJobStreamsClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, jobStreamID string, options *TestJobStreamsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams/{jobStreamId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	if jobStreamID == "" {
		return nil, errors.New("parameter jobStreamID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobStreamId}", url.PathEscape(jobStreamID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TestJobStreamsClient) getHandleResponse(resp *http.Response) (TestJobStreamsGetResponse, error) {
	result := TestJobStreamsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStream); err != nil {
		return TestJobStreamsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TestJobStreamsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByTestJob - Retrieve a list of test job streams identified by runbook name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *TestJobStreamsClient) ListByTestJob(resourceGroupName string, automationAccountName string, runbookName string, options *TestJobStreamsListByTestJobOptions) *TestJobStreamsListByTestJobPager {
	return &TestJobStreamsListByTestJobPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByTestJobCreateRequest(ctx, resourceGroupName, automationAccountName, runbookName, options)
		},
		advancer: func(ctx context.Context, resp TestJobStreamsListByTestJobResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.JobStreamListResult.NextLink)
		},
	}
}

// listByTestJobCreateRequest creates the ListByTestJob request.
func (client *TestJobStreamsClient) listByTestJobCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, runbookName string, options *TestJobStreamsListByTestJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/runbooks/{runbookName}/draft/testJob/streams"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if runbookName == "" {
		return nil, errors.New("parameter runbookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runbookName}", url.PathEscape(runbookName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2018-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByTestJobHandleResponse handles the ListByTestJob response.
func (client *TestJobStreamsClient) listByTestJobHandleResponse(resp *http.Response) (TestJobStreamsListByTestJobResponse, error) {
	result := TestJobStreamsListByTestJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobStreamListResult); err != nil {
		return TestJobStreamsListByTestJobResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByTestJobHandleError handles the ListByTestJob error response.
func (client *TestJobStreamsClient) listByTestJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
