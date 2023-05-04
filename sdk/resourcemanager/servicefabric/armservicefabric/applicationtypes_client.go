//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicefabric

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

// ApplicationTypesClient contains the methods for the ApplicationTypes group.
// Don't use this type directly, use NewApplicationTypesClient() instead.
type ApplicationTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationTypesClient creates a new instance of ApplicationTypesClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationTypesClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationTypesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Service Fabric application type name resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - applicationTypeName - The name of the application type name resource.
//   - parameters - The application type name resource.
//   - options - ApplicationTypesClientCreateOrUpdateOptions contains the optional parameters for the ApplicationTypesClient.CreateOrUpdate
//     method.
func (client *ApplicationTypesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters ApplicationTypeResource, options *ApplicationTypesClientCreateOrUpdateOptions) (ApplicationTypesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, parameters, options)
	if err != nil {
		return ApplicationTypesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationTypesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationTypesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationTypesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, parameters ApplicationTypeResource, options *ApplicationTypesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplicationTypesClient) createOrUpdateHandleResponse(resp *http.Response) (ApplicationTypesClientCreateOrUpdateResponse, error) {
	result := ApplicationTypesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeResource); err != nil {
		return ApplicationTypesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete a Service Fabric application type name resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - applicationTypeName - The name of the application type name resource.
//   - options - ApplicationTypesClientBeginDeleteOptions contains the optional parameters for the ApplicationTypesClient.BeginDelete
//     method.
func (client *ApplicationTypesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientBeginDeleteOptions) (*runtime.Poller[ApplicationTypesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, applicationTypeName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ApplicationTypesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ApplicationTypesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a Service Fabric application type name resource with the specified name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *ApplicationTypesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationTypesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Service Fabric application type name resource created or in the process of being created in the Service Fabric
// cluster resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - applicationTypeName - The name of the application type name resource.
//   - options - ApplicationTypesClientGetOptions contains the optional parameters for the ApplicationTypesClient.Get method.
func (client *ApplicationTypesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientGetOptions) (ApplicationTypesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, applicationTypeName, options)
	if err != nil {
		return ApplicationTypesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationTypesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationTypesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, applicationTypeName string, options *ApplicationTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes/{applicationTypeName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *ApplicationTypesClient) getHandleResponse(resp *http.Response) (ApplicationTypesClientGetResponse, error) {
	result := ApplicationTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeResource); err != nil {
		return ApplicationTypesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all application type name resources created or in the process of being created in the Service Fabric cluster
// resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group.
//   - clusterName - The name of the cluster resource.
//   - options - ApplicationTypesClientListOptions contains the optional parameters for the ApplicationTypesClient.List method.
func (client *ApplicationTypesClient) List(ctx context.Context, resourceGroupName string, clusterName string, options *ApplicationTypesClientListOptions) (ApplicationTypesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ApplicationTypesClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationTypesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ApplicationTypesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ApplicationTypesClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ApplicationTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applicationTypes"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *ApplicationTypesClient) listHandleResponse(resp *http.Response) (ApplicationTypesClientListResponse, error) {
	result := ApplicationTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationTypeResourceList); err != nil {
		return ApplicationTypesClientListResponse{}, err
	}
	return result, nil
}
