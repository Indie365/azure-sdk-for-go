//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

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

// ConfigurationProfilesClient contains the methods for the ConfigurationProfiles group.
// Don't use this type directly, use NewConfigurationProfilesClient() instead.
type ConfigurationProfilesClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewConfigurationProfilesClient creates a new instance of ConfigurationProfilesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationProfilesClient, error) {
	cl, err := arm.NewClient(moduleName+".ConfigurationProfilesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationProfilesClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - configurationProfileName - Name of the configuration profile.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parameters - Parameters supplied to create or update configuration profile.
//   - options - ConfigurationProfilesClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationProfilesClient.CreateOrUpdate
//     method.
func (client *ConfigurationProfilesClient) CreateOrUpdate(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfile, options *ConfigurationProfilesClientCreateOrUpdateOptions) (ConfigurationProfilesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, configurationProfileName, resourceGroupName, parameters, options)
	if err != nil {
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationProfilesClient) createOrUpdateCreateRequest(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfile, options *ConfigurationProfilesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConfigurationProfilesClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationProfilesClientCreateOrUpdateResponse, error) {
	result := ConfigurationProfilesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfile); err != nil {
		return ConfigurationProfilesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - configurationProfileName - Name of the configuration profile
//   - options - ConfigurationProfilesClientDeleteOptions contains the optional parameters for the ConfigurationProfilesClient.Delete
//     method.
func (client *ConfigurationProfilesClient) Delete(ctx context.Context, resourceGroupName string, configurationProfileName string, options *ConfigurationProfilesClientDeleteOptions) (ConfigurationProfilesClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configurationProfileName, options)
	if err != nil {
		return ConfigurationProfilesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfilesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfilesClientDeleteResponse{}, err
	}
	return ConfigurationProfilesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileName string, options *ConfigurationProfilesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get information about a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - configurationProfileName - The configuration profile name.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ConfigurationProfilesClientGetOptions contains the optional parameters for the ConfigurationProfilesClient.Get
//     method.
func (client *ConfigurationProfilesClient) Get(ctx context.Context, configurationProfileName string, resourceGroupName string, options *ConfigurationProfilesClientGetOptions) (ConfigurationProfilesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, configurationProfileName, resourceGroupName, options)
	if err != nil {
		return ConfigurationProfilesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfilesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationProfilesClient) getCreateRequest(ctx context.Context, configurationProfileName string, resourceGroupName string, options *ConfigurationProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
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
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationProfilesClient) getHandleResponse(resp *http.Response) (ConfigurationProfilesClientGetResponse, error) {
	result := ConfigurationProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfile); err != nil {
		return ConfigurationProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve a list of configuration profile within a given resource group
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ConfigurationProfilesClientListByResourceGroupOptions contains the optional parameters for the ConfigurationProfilesClient.NewListByResourceGroupPager
//     method.
func (client *ConfigurationProfilesClient) NewListByResourceGroupPager(resourceGroupName string, options *ConfigurationProfilesClientListByResourceGroupOptions) (*runtime.Pager[ConfigurationProfilesClientListByResourceGroupResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationProfilesClientListByResourceGroupResponse]{
		More: func(page ConfigurationProfilesClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfilesClientListByResourceGroupResponse) (ConfigurationProfilesClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return ConfigurationProfilesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationProfilesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfilesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConfigurationProfilesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationProfilesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles"
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
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConfigurationProfilesClient) listByResourceGroupHandleResponse(resp *http.Response) (ConfigurationProfilesClientListByResourceGroupResponse, error) {
	result := ConfigurationProfilesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileList); err != nil {
		return ConfigurationProfilesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Retrieve a list of configuration profile within a subscription
//
// Generated from API version 2022-05-04
//   - options - ConfigurationProfilesClientListBySubscriptionOptions contains the optional parameters for the ConfigurationProfilesClient.NewListBySubscriptionPager
//     method.
func (client *ConfigurationProfilesClient) NewListBySubscriptionPager(options *ConfigurationProfilesClientListBySubscriptionOptions) (*runtime.Pager[ConfigurationProfilesClientListBySubscriptionResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationProfilesClientListBySubscriptionResponse]{
		More: func(page ConfigurationProfilesClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfilesClientListBySubscriptionResponse) (ConfigurationProfilesClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return ConfigurationProfilesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationProfilesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfilesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConfigurationProfilesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConfigurationProfilesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Automanage/configurationProfiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConfigurationProfilesClient) listBySubscriptionHandleResponse(resp *http.Response) (ConfigurationProfilesClientListBySubscriptionResponse, error) {
	result := ConfigurationProfilesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileList); err != nil {
		return ConfigurationProfilesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates a configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - configurationProfileName - Name of the configuration profile.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - parameters - Parameters supplied to update configuration profile.
//   - options - ConfigurationProfilesClientUpdateOptions contains the optional parameters for the ConfigurationProfilesClient.Update
//     method.
func (client *ConfigurationProfilesClient) Update(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfileUpdate, options *ConfigurationProfilesClientUpdateOptions) (ConfigurationProfilesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, configurationProfileName, resourceGroupName, parameters, options)
	if err != nil {
		return ConfigurationProfilesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfilesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfilesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ConfigurationProfilesClient) updateCreateRequest(ctx context.Context, configurationProfileName string, resourceGroupName string, parameters ConfigurationProfileUpdate, options *ConfigurationProfilesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}"
	if configurationProfileName == "" {
		return nil, errors.New("parameter configurationProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileName}", url.PathEscape(configurationProfileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ConfigurationProfilesClient) updateHandleResponse(resp *http.Response) (ConfigurationProfilesClientUpdateResponse, error) {
	result := ConfigurationProfilesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfile); err != nil {
		return ConfigurationProfilesClientUpdateResponse{}, err
	}
	return result, nil
}

