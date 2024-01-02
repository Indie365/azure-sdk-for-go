//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscheduler

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

// JobCollectionsClient contains the methods for the JobCollections group.
// Don't use this type directly, use NewJobCollectionsClient() instead.
type JobCollectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewJobCollectionsClient creates a new instance of JobCollectionsClient with the specified values.
//   - subscriptionID - The subscription id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewJobCollectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*JobCollectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &JobCollectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Provisions a new job collection or updates an existing job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - resourceGroupName - The resource group name.
//   - jobCollectionName - The job collection name.
//   - jobCollection - The job collection definition.
//   - options - JobCollectionsClientCreateOrUpdateOptions contains the optional parameters for the JobCollectionsClient.CreateOrUpdate
//     method.
func (client *JobCollectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsClientCreateOrUpdateOptions) (JobCollectionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "JobCollectionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, jobCollectionName, jobCollection, options)
	if err != nil {
		return JobCollectionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobCollectionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return JobCollectionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *JobCollectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	if err := runtime.MarshalAsJSON(req, jobCollection); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *JobCollectionsClient) createOrUpdateHandleResponse(resp *http.Response) (JobCollectionsClientCreateOrUpdateResponse, error) {
	result := JobCollectionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionDefinition); err != nil {
		return JobCollectionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - resourceGroupName - The resource group name.
//   - jobCollectionName - The job collection name.
//   - options - JobCollectionsClientBeginDeleteOptions contains the optional parameters for the JobCollectionsClient.BeginDelete
//     method.
func (client *JobCollectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginDeleteOptions) (*runtime.Poller[JobCollectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, jobCollectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[JobCollectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[JobCollectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
func (client *JobCollectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "JobCollectionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *JobCollectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginDisable - Disables all of the jobs in the job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - resourceGroupName - The resource group name.
//   - jobCollectionName - The job collection name.
//   - options - JobCollectionsClientBeginDisableOptions contains the optional parameters for the JobCollectionsClient.BeginDisable
//     method.
func (client *JobCollectionsClient) BeginDisable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginDisableOptions) (*runtime.Poller[JobCollectionsClientDisableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.disable(ctx, resourceGroupName, jobCollectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[JobCollectionsClientDisableResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[JobCollectionsClientDisableResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Disable - Disables all of the jobs in the job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
func (client *JobCollectionsClient) disable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginDisableOptions) (*http.Response, error) {
	var err error
	const operationName = "JobCollectionsClient.BeginDisable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.disableCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// disableCreateRequest creates the Disable request.
func (client *JobCollectionsClient) disableCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginDisableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/disable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginEnable - Enables all of the jobs in the job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - resourceGroupName - The resource group name.
//   - jobCollectionName - The job collection name.
//   - options - JobCollectionsClientBeginEnableOptions contains the optional parameters for the JobCollectionsClient.BeginEnable
//     method.
func (client *JobCollectionsClient) BeginEnable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginEnableOptions) (*runtime.Poller[JobCollectionsClientEnableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.enable(ctx, resourceGroupName, jobCollectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[JobCollectionsClientEnableResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[JobCollectionsClientEnableResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Enable - Enables all of the jobs in the job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
func (client *JobCollectionsClient) enable(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginEnableOptions) (*http.Response, error) {
	var err error
	const operationName = "JobCollectionsClient.BeginEnable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.enableCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// enableCreateRequest creates the Enable request.
func (client *JobCollectionsClient) enableCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientBeginEnableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/enable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - resourceGroupName - The resource group name.
//   - jobCollectionName - The job collection name.
//   - options - JobCollectionsClientGetOptions contains the optional parameters for the JobCollectionsClient.Get method.
func (client *JobCollectionsClient) Get(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientGetOptions) (JobCollectionsClientGetResponse, error) {
	var err error
	const operationName = "JobCollectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobCollectionName, options)
	if err != nil {
		return JobCollectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobCollectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobCollectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *JobCollectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, options *JobCollectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *JobCollectionsClient) getHandleResponse(resp *http.Response) (JobCollectionsClientGetResponse, error) {
	result := JobCollectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionDefinition); err != nil {
		return JobCollectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all job collections under specified resource group.
//
// Generated from API version 2016-03-01
//   - resourceGroupName - The resource group name.
//   - options - JobCollectionsClientListByResourceGroupOptions contains the optional parameters for the JobCollectionsClient.NewListByResourceGroupPager
//     method.
func (client *JobCollectionsClient) NewListByResourceGroupPager(resourceGroupName string, options *JobCollectionsClientListByResourceGroupOptions) *runtime.Pager[JobCollectionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobCollectionsClientListByResourceGroupResponse]{
		More: func(page JobCollectionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobCollectionsClientListByResourceGroupResponse) (JobCollectionsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "JobCollectionsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return JobCollectionsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *JobCollectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *JobCollectionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *JobCollectionsClient) listByResourceGroupHandleResponse(resp *http.Response) (JobCollectionsClientListByResourceGroupResponse, error) {
	result := JobCollectionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionListResult); err != nil {
		return JobCollectionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all job collections under specified subscription.
//
// Generated from API version 2016-03-01
//   - options - JobCollectionsClientListBySubscriptionOptions contains the optional parameters for the JobCollectionsClient.NewListBySubscriptionPager
//     method.
func (client *JobCollectionsClient) NewListBySubscriptionPager(options *JobCollectionsClientListBySubscriptionOptions) *runtime.Pager[JobCollectionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[JobCollectionsClientListBySubscriptionResponse]{
		More: func(page JobCollectionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *JobCollectionsClientListBySubscriptionResponse) (JobCollectionsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "JobCollectionsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return JobCollectionsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *JobCollectionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *JobCollectionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Scheduler/jobCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *JobCollectionsClient) listBySubscriptionHandleResponse(resp *http.Response) (JobCollectionsClientListBySubscriptionResponse, error) {
	result := JobCollectionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionListResult); err != nil {
		return JobCollectionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Patch - Patches an existing job collection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-03-01
//   - resourceGroupName - The resource group name.
//   - jobCollectionName - The job collection name.
//   - jobCollection - The job collection definition.
//   - options - JobCollectionsClientPatchOptions contains the optional parameters for the JobCollectionsClient.Patch method.
func (client *JobCollectionsClient) Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsClientPatchOptions) (JobCollectionsClientPatchResponse, error) {
	var err error
	const operationName = "JobCollectionsClient.Patch"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.patchCreateRequest(ctx, resourceGroupName, jobCollectionName, jobCollection, options)
	if err != nil {
		return JobCollectionsClientPatchResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return JobCollectionsClientPatchResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return JobCollectionsClientPatchResponse{}, err
	}
	resp, err := client.patchHandleResponse(httpResp)
	return resp, err
}

// patchCreateRequest creates the Patch request.
func (client *JobCollectionsClient) patchCreateRequest(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition, options *JobCollectionsClientPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobCollectionName == "" {
		return nil, errors.New("parameter jobCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobCollectionName}", url.PathEscape(jobCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	if err := runtime.MarshalAsJSON(req, jobCollection); err != nil {
		return nil, err
	}
	return req, nil
}

// patchHandleResponse handles the Patch response.
func (client *JobCollectionsClient) patchHandleResponse(resp *http.Response) (JobCollectionsClientPatchResponse, error) {
	result := JobCollectionsClientPatchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobCollectionDefinition); err != nil {
		return JobCollectionsClientPatchResponse{}, err
	}
	return result, nil
}
