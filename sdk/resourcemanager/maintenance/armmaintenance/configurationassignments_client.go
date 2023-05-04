//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmaintenance

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

// ConfigurationAssignmentsClient contains the methods for the ConfigurationAssignments group.
// Don't use this type directly, use NewConfigurationAssignmentsClient() instead.
type ConfigurationAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationAssignmentsClient creates a new instance of ConfigurationAssignmentsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName+".ConfigurationAssignmentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Register configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Configuration assignment name
//   - configurationAssignment - The configurationAssignment
//   - options - ConfigurationAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationAssignmentsClient.CreateOrUpdate
//     method.
func (client *ConfigurationAssignmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsClientCreateOrUpdateOptions) (ConfigurationAssignmentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, configurationAssignment, options)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, configurationAssignment)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConfigurationAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientCreateOrUpdateResponse, error) {
	result := ConfigurationAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateParent - Register configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Configuration assignment name
//   - configurationAssignment - The configurationAssignment
//   - options - ConfigurationAssignmentsClientCreateOrUpdateParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.CreateOrUpdateParent
//     method.
func (client *ConfigurationAssignmentsClient) CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsClientCreateOrUpdateParentOptions) (ConfigurationAssignmentsClientCreateOrUpdateParentResponse, error) {
	req, err := client.createOrUpdateParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, configurationAssignment, options)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateParentHandleResponse(resp)
}

// createOrUpdateParentCreateRequest creates the CreateOrUpdateParent request.
func (client *ConfigurationAssignmentsClient) createOrUpdateParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsClientCreateOrUpdateParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, configurationAssignment)
}

// createOrUpdateParentHandleResponse handles the CreateOrUpdateParent response.
func (client *ConfigurationAssignmentsClient) createOrUpdateParentHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientCreateOrUpdateParentResponse, error) {
	result := ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, err
	}
	return result, nil
}

// Delete - Unregister configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Unique configuration assignment name
//   - options - ConfigurationAssignmentsClientDeleteOptions contains the optional parameters for the ConfigurationAssignmentsClient.Delete
//     method.
func (client *ConfigurationAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientDeleteOptions) (ConfigurationAssignmentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConfigurationAssignmentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ConfigurationAssignmentsClient) deleteHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientDeleteResponse, error) {
	result := ConfigurationAssignmentsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientDeleteResponse{}, err
	}
	return result, nil
}

// DeleteParent - Unregister configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Unique configuration assignment name
//   - options - ConfigurationAssignmentsClientDeleteParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.DeleteParent
//     method.
func (client *ConfigurationAssignmentsClient) DeleteParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientDeleteParentOptions) (ConfigurationAssignmentsClientDeleteParentResponse, error) {
	req, err := client.deleteParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteParentResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteParentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConfigurationAssignmentsClientDeleteParentResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteParentHandleResponse(resp)
}

// deleteParentCreateRequest creates the DeleteParent request.
func (client *ConfigurationAssignmentsClient) deleteParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientDeleteParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteParentHandleResponse handles the DeleteParent response.
func (client *ConfigurationAssignmentsClient) deleteParentHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientDeleteParentResponse, error) {
	result := ConfigurationAssignmentsClientDeleteParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientDeleteParentResponse{}, err
	}
	return result, nil
}

// NewListPager - List configurationAssignments for resource.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - ConfigurationAssignmentsClientListOptions contains the optional parameters for the ConfigurationAssignmentsClient.NewListPager
//     method.
func (client *ConfigurationAssignmentsClient) NewListPager(resourceGroupName string, providerName string, resourceType string, resourceName string, options *ConfigurationAssignmentsClientListOptions) *runtime.Pager[ConfigurationAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationAssignmentsClientListResponse]{
		More: func(page ConfigurationAssignmentsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationAssignmentsClientListResponse) (ConfigurationAssignmentsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, options)
			if err != nil {
				return ConfigurationAssignmentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConfigurationAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, options *ConfigurationAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConfigurationAssignmentsClient) listHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientListResponse, error) {
	result := ConfigurationAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListConfigurationAssignmentsResult); err != nil {
		return ConfigurationAssignmentsClientListResponse{}, err
	}
	return result, nil
}

// NewListParentPager - List configurationAssignments for resource.
//
// Generated from API version 2021-05-01
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - ConfigurationAssignmentsClientListParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.NewListParentPager
//     method.
func (client *ConfigurationAssignmentsClient) NewListParentPager(resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *ConfigurationAssignmentsClientListParentOptions) *runtime.Pager[ConfigurationAssignmentsClientListParentResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationAssignmentsClientListParentResponse]{
		More: func(page ConfigurationAssignmentsClientListParentResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationAssignmentsClientListParentResponse) (ConfigurationAssignmentsClientListParentResponse, error) {
			req, err := client.listParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, options)
			if err != nil {
				return ConfigurationAssignmentsClientListParentResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationAssignmentsClientListParentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationAssignmentsClientListParentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listParentHandleResponse(resp)
		},
	})
}

// listParentCreateRequest creates the ListParent request.
func (client *ConfigurationAssignmentsClient) listParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *ConfigurationAssignmentsClientListParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listParentHandleResponse handles the ListParent response.
func (client *ConfigurationAssignmentsClient) listParentHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientListParentResponse, error) {
	result := ConfigurationAssignmentsClientListParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListConfigurationAssignmentsResult); err != nil {
		return ConfigurationAssignmentsClientListParentResponse{}, err
	}
	return result, nil
}
