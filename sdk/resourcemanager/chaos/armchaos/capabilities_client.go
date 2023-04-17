//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos

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

// CapabilitiesClient contains the methods for the Capabilities group.
// Don't use this type directly, use NewCapabilitiesClient() instead.
type CapabilitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCapabilitiesClient creates a new instance of CapabilitiesClient with the specified values.
//   - subscriptionID - GUID that represents an Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCapabilitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CapabilitiesClient, error) {
	cl, err := arm.NewClient(moduleName+".CapabilitiesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CapabilitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Capability resource that extends a Target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - targetName - String that represents a Target resource name.
//   - capabilityName - String that represents a Capability resource name.
//   - capability - Capability resource to be created or updated.
//   - options - CapabilitiesClientCreateOrUpdateOptions contains the optional parameters for the CapabilitiesClient.CreateOrUpdate
//     method.
func (client *CapabilitiesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, capability Capability, options *CapabilitiesClientCreateOrUpdateOptions) (CapabilitiesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName, capability, options)
	if err != nil {
		return CapabilitiesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapabilitiesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapabilitiesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CapabilitiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, capability Capability, options *CapabilitiesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities/{capabilityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	if capabilityName == "" {
		return nil, errors.New("parameter capabilityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityName}", url.PathEscape(capabilityName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, capability)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CapabilitiesClient) createOrUpdateHandleResponse(resp *http.Response) (CapabilitiesClientCreateOrUpdateResponse, error) {
	result := CapabilitiesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Capability); err != nil {
		return CapabilitiesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Capability that extends a Target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - targetName - String that represents a Target resource name.
//   - capabilityName - String that represents a Capability resource name.
//   - options - CapabilitiesClientDeleteOptions contains the optional parameters for the CapabilitiesClient.Delete method.
func (client *CapabilitiesClient) Delete(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, options *CapabilitiesClientDeleteOptions) (CapabilitiesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName, options)
	if err != nil {
		return CapabilitiesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapabilitiesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CapabilitiesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CapabilitiesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CapabilitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, options *CapabilitiesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities/{capabilityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	if capabilityName == "" {
		return nil, errors.New("parameter capabilityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityName}", url.PathEscape(capabilityName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Capability resource that extends a Target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-09-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - targetName - String that represents a Target resource name.
//   - capabilityName - String that represents a Capability resource name.
//   - options - CapabilitiesClientGetOptions contains the optional parameters for the CapabilitiesClient.Get method.
func (client *CapabilitiesClient) Get(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, options *CapabilitiesClientGetOptions) (CapabilitiesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, capabilityName, options)
	if err != nil {
		return CapabilitiesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CapabilitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapabilitiesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CapabilitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, capabilityName string, options *CapabilitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities/{capabilityName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	if capabilityName == "" {
		return nil, errors.New("parameter capabilityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capabilityName}", url.PathEscape(capabilityName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapabilitiesClient) getHandleResponse(resp *http.Response) (CapabilitiesClientGetResponse, error) {
	result := CapabilitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Capability); err != nil {
		return CapabilitiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of Capability resources that extend a Target resource..
//
// Generated from API version 2021-09-15-preview
//   - resourceGroupName - String that represents an Azure resource group.
//   - parentProviderNamespace - String that represents a resource provider namespace.
//   - parentResourceType - String that represents a resource type.
//   - parentResourceName - String that represents a resource name.
//   - targetName - String that represents a Target resource name.
//   - options - CapabilitiesClientListOptions contains the optional parameters for the CapabilitiesClient.NewListPager method.
func (client *CapabilitiesClient) NewListPager(resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *CapabilitiesClientListOptions) *runtime.Pager[CapabilitiesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapabilitiesClientListResponse]{
		More: func(page CapabilitiesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapabilitiesClientListResponse) (CapabilitiesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, parentProviderNamespace, parentResourceType, parentResourceName, targetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CapabilitiesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CapabilitiesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CapabilitiesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CapabilitiesClient) listCreateRequest(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *CapabilitiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentProviderNamespace == "" {
		return nil, errors.New("parameter parentProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentProviderNamespace}", url.PathEscape(parentProviderNamespace))
	if parentResourceType == "" {
		return nil, errors.New("parameter parentResourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceType}", url.PathEscape(parentResourceType))
	if parentResourceName == "" {
		return nil, errors.New("parameter parentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentResourceName}", url.PathEscape(parentResourceName))
	if targetName == "" {
		return nil, errors.New("parameter targetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetName}", url.PathEscape(targetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-15-preview")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CapabilitiesClient) listHandleResponse(resp *http.Response) (CapabilitiesClientListResponse, error) {
	result := CapabilitiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilityListResult); err != nil {
		return CapabilitiesClientListResponse{}, err
	}
	return result, nil
}
