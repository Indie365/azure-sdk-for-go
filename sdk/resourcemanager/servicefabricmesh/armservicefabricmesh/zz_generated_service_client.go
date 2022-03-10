//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

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

// ServiceClient contains the methods for the Service group.
// Don't use this type directly, use NewServiceClient() instead.
type ServiceClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceClient creates a new instance of ServiceClient with the specified values.
// subscriptionID - The customer subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServiceClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ServiceClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets the information about the service resource with the given name. The information include the description and
// other properties of the service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// serviceResourceName - The identity of the service.
// options - ServiceClientGetOptions contains the optional parameters for the ServiceClient.Get method.
func (client *ServiceClient) Get(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, options *ServiceClientGetOptions) (ServiceClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationResourceName, serviceResourceName, options)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, options *ServiceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services/{serviceResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceResourceName}", serviceResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceClient) getHandleResponse(resp *http.Response) (ServiceClientGetResponse, error) {
	result := ServiceClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceDescription); err != nil {
		return ServiceClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets the information about all services of an application resource. The information include the description and
// other properties of the Service.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Azure resource group name
// applicationResourceName - The identity of the application.
// options - ServiceClientListOptions contains the optional parameters for the ServiceClient.List method.
func (client *ServiceClient) List(resourceGroupName string, applicationResourceName string, options *ServiceClientListOptions) *ServiceClientListPager {
	return &ServiceClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, applicationResourceName, options)
		},
		advancer: func(ctx context.Context, resp ServiceClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServiceResourceDescriptionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ServiceClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationResourceName string, options *ServiceClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabricMesh/applications/{applicationResourceName}/services"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationResourceName}", applicationResourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceClient) listHandleResponse(resp *http.Response) (ServiceClientListResponse, error) {
	result := ServiceClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResourceDescriptionList); err != nil {
		return ServiceClientListResponse{}, err
	}
	return result, nil
}
