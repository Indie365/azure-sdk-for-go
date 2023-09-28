//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

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

// ExportPipelinesClient contains the methods for the ExportPipelines group.
// Don't use this type directly, use NewExportPipelinesClient() instead.
type ExportPipelinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExportPipelinesClient creates a new instance of ExportPipelinesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExportPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExportPipelinesClient, error) {
	cl, err := arm.NewClient(moduleName+".ExportPipelinesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExportPipelinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates an export pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - exportPipelineName - The name of the export pipeline.
//   - exportPipelineCreateParameters - The parameters for creating an export pipeline.
//   - options - ExportPipelinesClientBeginCreateOptions contains the optional parameters for the ExportPipelinesClient.BeginCreate
//     method.
func (client *ExportPipelinesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, exportPipelineCreateParameters ExportPipeline, options *ExportPipelinesClientBeginCreateOptions) (*runtime.Poller[ExportPipelinesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, exportPipelineName, exportPipelineCreateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExportPipelinesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ExportPipelinesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates an export pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *ExportPipelinesClient) create(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, exportPipelineCreateParameters ExportPipeline, options *ExportPipelinesClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ExportPipelinesClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, exportPipelineName, exportPipelineCreateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ExportPipelinesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, exportPipelineCreateParameters ExportPipeline, options *ExportPipelinesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines/{exportPipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if exportPipelineName == "" {
		return nil, errors.New("parameter exportPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportPipelineName}", url.PathEscape(exportPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, exportPipelineCreateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an export pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - exportPipelineName - The name of the export pipeline.
//   - options - ExportPipelinesClientBeginDeleteOptions contains the optional parameters for the ExportPipelinesClient.BeginDelete
//     method.
func (client *ExportPipelinesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientBeginDeleteOptions) (*runtime.Poller[ExportPipelinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, exportPipelineName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ExportPipelinesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ExportPipelinesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an export pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *ExportPipelinesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ExportPipelinesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, exportPipelineName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExportPipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines/{exportPipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if exportPipelineName == "" {
		return nil, errors.New("parameter exportPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportPipelineName}", url.PathEscape(exportPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the export pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - exportPipelineName - The name of the export pipeline.
//   - options - ExportPipelinesClientGetOptions contains the optional parameters for the ExportPipelinesClient.Get method.
func (client *ExportPipelinesClient) Get(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientGetOptions) (ExportPipelinesClientGetResponse, error) {
	var err error
	const operationName = "ExportPipelinesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, exportPipelineName, options)
	if err != nil {
		return ExportPipelinesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExportPipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExportPipelinesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExportPipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, exportPipelineName string, options *ExportPipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines/{exportPipelineName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if exportPipelineName == "" {
		return nil, errors.New("parameter exportPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{exportPipelineName}", url.PathEscape(exportPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExportPipelinesClient) getHandleResponse(resp *http.Response) (ExportPipelinesClientGetResponse, error) {
	result := ExportPipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExportPipeline); err != nil {
		return ExportPipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all export pipelines for the specified container registry.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - registryName - The name of the container registry.
//   - options - ExportPipelinesClientListOptions contains the optional parameters for the ExportPipelinesClient.NewListPager
//     method.
func (client *ExportPipelinesClient) NewListPager(resourceGroupName string, registryName string, options *ExportPipelinesClientListOptions) *runtime.Pager[ExportPipelinesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExportPipelinesClientListResponse]{
		More: func(page ExportPipelinesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExportPipelinesClientListResponse) (ExportPipelinesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExportPipelinesClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExportPipelinesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExportPipelinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExportPipelinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ExportPipelinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ExportPipelinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/exportPipelines"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExportPipelinesClient) listHandleResponse(resp *http.Response) (ExportPipelinesClientListResponse, error) {
	result := ExportPipelinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExportPipelineListResult); err != nil {
		return ExportPipelinesClientListResponse{}, err
	}
	return result, nil
}
