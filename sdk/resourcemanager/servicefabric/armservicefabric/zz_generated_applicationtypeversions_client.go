//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

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

// ApplicationTypeVersionsClient contains the methods for the ApplicationTypeVersions group.
// Don't use this type directly, use NewApplicationTypeVersionsClient() instead.
type ApplicationTypeVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationTypeVersionsClient creates a new instance of ApplicationTypeVersionsClient with the specified values.
// subscriptionID - The customer subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewApplicationTypeVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationTypeVersionsClient, error) {
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
	client := &ApplicationTypeVersionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a Service Fabric application type version resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster resource.
// applicationTypeName - The name of the application type name resource.
// version - The application type version.
// parameters - The application type version resource.
// options - ApplicationTypeVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationTypeVersionsClient.BeginCreateOrUpdate
// method.
func (client *ApplicationTypeVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource, options *ApplicationTypeVersionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ApplicationTypeVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, applicationTypeName, version, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ApplicationTypeVersionsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationTypeVersionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a Service Fabric application type version resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *ApplicationTypeVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource, options *ApplicationTypeVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, version, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationTypeVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, parameters ApplicationTypeVersionResource, options *ApplicationTypeVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Delete a Service Fabric application type version resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster resource.
// applicationTypeName - The name of the application type name resource.
// version - The application type version.
// options - ApplicationTypeVersionsClientBeginDeleteOptions contains the optional parameters for the ApplicationTypeVersionsClient.BeginDelete
// method.
func (client *ApplicationTypeVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsClientBeginDeleteOptions) (*runtime.Poller[ApplicationTypeVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, applicationTypeName, version, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ApplicationTypeVersionsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationTypeVersionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a Service Fabric application type version resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
func (client *ApplicationTypeVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, version, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationTypeVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Service Fabric application type version resource created or in the process of being created in the Service
// Fabric application type name resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster resource.
// applicationTypeName - The name of the application type name resource.
// version - The application type version.
// options - ApplicationTypeVersionsClientGetOptions contains the optional parameters for the ApplicationTypeVersionsClient.Get
// method.
func (client *ApplicationTypeVersionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsClientGetOptions) (ApplicationTypeVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, version, options)
	if err != nil {
		return ApplicationTypeVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationTypeVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationTypeVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationTypeVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, version string, options *ApplicationTypeVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationTypeVersionsClient) getHandleResponse(resp *http.Response) (ApplicationTypeVersionsClientGetResponse, error) {
	result := ApplicationTypeVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeVersionResource); err != nil {
		return ApplicationTypeVersionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all application type version resources created or in the process of being created in the Service Fabric application
// type name resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-06-01
// resourceGroupName - The name of the resource group.
// clusterName - The name of the cluster resource.
// applicationTypeName - The name of the application type name resource.
// options - ApplicationTypeVersionsClientListOptions contains the optional parameters for the ApplicationTypeVersionsClient.List
// method.
func (client *ApplicationTypeVersionsClient) List(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypeVersionsClientListOptions) (ApplicationTypeVersionsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, options)
	if err != nil {
		return ApplicationTypeVersionsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationTypeVersionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationTypeVersionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ApplicationTypeVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypeVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if applicationTypeName == "" {
		return nil, errors.New("parameter applicationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationTypeName}", url.PathEscape(applicationTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationTypeVersionsClient) listHandleResponse(resp *http.Response) (ApplicationTypeVersionsClientListResponse, error) {
	result := ApplicationTypeVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeVersionResourceList); err != nil {
		return ApplicationTypeVersionsClientListResponse{}, err
	}
	return result, nil
}