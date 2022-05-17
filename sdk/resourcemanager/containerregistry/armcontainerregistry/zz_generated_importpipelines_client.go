//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

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

// ImportPipelinesClient contains the methods for the ImportPipelines group.
// Don't use this type directly, use NewImportPipelinesClient() instead.
type ImportPipelinesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewImportPipelinesClient creates a new instance of ImportPipelinesClient with the specified values.
// subscriptionID - The Microsoft Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewImportPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImportPipelinesClient, error) {
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
	client := &ImportPipelinesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates an import pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01-preview
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// importPipelineName - The name of the import pipeline.
// importPipelineCreateParameters - The parameters for creating an import pipeline.
// options - ImportPipelinesClientBeginCreateOptions contains the optional parameters for the ImportPipelinesClient.BeginCreate
// method.
func (client *ImportPipelinesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline, options *ImportPipelinesClientBeginCreateOptions) (*runtime.Poller[ImportPipelinesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, registryName, importPipelineName, importPipelineCreateParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ImportPipelinesClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ImportPipelinesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates an import pipeline for a container registry with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01-preview
func (client *ImportPipelinesClient) create(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline, options *ImportPipelinesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, importPipelineName, importPipelineCreateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ImportPipelinesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, importPipelineCreateParameters ImportPipeline, options *ImportPipelinesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if importPipelineName == "" {
		return nil, errors.New("parameter importPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importPipelineName}", url.PathEscape(importPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, importPipelineCreateParameters)
}

// BeginDelete - Deletes an import pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01-preview
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// importPipelineName - The name of the import pipeline.
// options - ImportPipelinesClientBeginDeleteOptions contains the optional parameters for the ImportPipelinesClient.BeginDelete
// method.
func (client *ImportPipelinesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientBeginDeleteOptions) (*runtime.Poller[ImportPipelinesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, importPipelineName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ImportPipelinesClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ImportPipelinesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes an import pipeline from a container registry.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01-preview
func (client *ImportPipelinesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, importPipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ImportPipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if importPipelineName == "" {
		return nil, errors.New("parameter importPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importPipelineName}", url.PathEscape(importPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the import pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01-preview
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// importPipelineName - The name of the import pipeline.
// options - ImportPipelinesClientGetOptions contains the optional parameters for the ImportPipelinesClient.Get method.
func (client *ImportPipelinesClient) Get(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientGetOptions) (ImportPipelinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, importPipelineName, options)
	if err != nil {
		return ImportPipelinesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ImportPipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ImportPipelinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ImportPipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, importPipelineName string, options *ImportPipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines/{importPipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if importPipelineName == "" {
		return nil, errors.New("parameter importPipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{importPipelineName}", url.PathEscape(importPipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ImportPipelinesClient) getHandleResponse(resp *http.Response) (ImportPipelinesClientGetResponse, error) {
	result := ImportPipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportPipeline); err != nil {
		return ImportPipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all import pipelines for the specified container registry.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01-preview
// resourceGroupName - The name of the resource group to which the container registry belongs.
// registryName - The name of the container registry.
// options - ImportPipelinesClientListOptions contains the optional parameters for the ImportPipelinesClient.List method.
func (client *ImportPipelinesClient) NewListPager(resourceGroupName string, registryName string, options *ImportPipelinesClientListOptions) *runtime.Pager[ImportPipelinesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImportPipelinesClientListResponse]{
		More: func(page ImportPipelinesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImportPipelinesClientListResponse) (ImportPipelinesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, registryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ImportPipelinesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ImportPipelinesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ImportPipelinesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ImportPipelinesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ImportPipelinesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/importPipelines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ImportPipelinesClient) listHandleResponse(resp *http.Response) (ImportPipelinesClientListResponse, error) {
	result := ImportPipelinesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImportPipelineListResult); err != nil {
		return ImportPipelinesClientListResponse{}, err
	}
	return result, nil
}
