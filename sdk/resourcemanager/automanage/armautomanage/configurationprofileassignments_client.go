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

// ConfigurationProfileAssignmentsClient contains the methods for the ConfigurationProfileAssignments group.
// Don't use this type directly, use NewConfigurationProfileAssignmentsClient() instead.
type ConfigurationProfileAssignmentsClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewConfigurationProfileAssignmentsClient creates a new instance of ConfigurationProfileAssignmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationProfileAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationProfileAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName+".ConfigurationProfileAssignmentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationProfileAssignmentsClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates an association between a VM and Automanage configuration profile
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - configurationProfileAssignmentName - Name of the configuration profile assignment. Only default is supported.
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmName - The name of the virtual machine.
//   - parameters - Parameters supplied to the create or update configuration profile assignment.
//   - options - ConfigurationProfileAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.CreateOrUpdate
//     method.
func (client *ConfigurationProfileAssignmentsClient) CreateOrUpdate(ctx context.Context, configurationProfileAssignmentName string, resourceGroupName string, vmName string, parameters ConfigurationProfileAssignment, options *ConfigurationProfileAssignmentsClientCreateOrUpdateOptions) (ConfigurationProfileAssignmentsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, configurationProfileAssignmentName, resourceGroupName, vmName, parameters, options)
	if err != nil {
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationProfileAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, configurationProfileAssignmentName string, resourceGroupName string, vmName string, parameters ConfigurationProfileAssignment, options *ConfigurationProfileAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
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
func (client *ConfigurationProfileAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientCreateOrUpdateResponse, error) {
	result := ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignment); err != nil {
		return ConfigurationProfileAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - configurationProfileAssignmentName - Name of the configuration profile assignment
//   - vmName - The name of the virtual machine.
//   - options - ConfigurationProfileAssignmentsClientDeleteOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.Delete
//     method.
func (client *ConfigurationProfileAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientDeleteOptions) (ConfigurationProfileAssignmentsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, configurationProfileAssignmentName, vmName, options)
	if err != nil {
		return ConfigurationProfileAssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfileAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfileAssignmentsClientDeleteResponse{}, err
	}
	return ConfigurationProfileAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationProfileAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
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

// Get - Get information about a configuration profile assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - configurationProfileAssignmentName - The configuration profile assignment name.
//   - vmName - The name of the virtual machine.
//   - options - ConfigurationProfileAssignmentsClientGetOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.Get
//     method.
func (client *ConfigurationProfileAssignmentsClient) Get(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientGetOptions) (ConfigurationProfileAssignmentsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, configurationProfileAssignmentName, vmName, options)
	if err != nil {
		return ConfigurationProfileAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationProfileAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationProfileAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationProfileAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, configurationProfileAssignmentName string, vmName string, options *ConfigurationProfileAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments/{configurationProfileAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configurationProfileAssignmentName == "" {
		return nil, errors.New("parameter configurationProfileAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationProfileAssignmentName}", url.PathEscape(configurationProfileAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
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
func (client *ConfigurationProfileAssignmentsClient) getHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientGetResponse, error) {
	result := ConfigurationProfileAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignment); err != nil {
		return ConfigurationProfileAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get list of configuration profile assignments
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ConfigurationProfileAssignmentsClientListOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.NewListPager
//     method.
func (client *ConfigurationProfileAssignmentsClient) NewListPager(resourceGroupName string, options *ConfigurationProfileAssignmentsClientListOptions) (*runtime.Pager[ConfigurationProfileAssignmentsClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationProfileAssignmentsClientListResponse]{
		More: func(page ConfigurationProfileAssignmentsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfileAssignmentsClientListResponse) (ConfigurationProfileAssignmentsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfileAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConfigurationProfileAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ConfigurationProfileAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfileAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listHandleResponse handles the List response.
func (client *ConfigurationProfileAssignmentsClient) listHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientListResponse, error) {
	result := ConfigurationProfileAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignmentList); err != nil {
		return ConfigurationProfileAssignmentsClientListResponse{}, err
	}
	return result, nil
}

// NewListByClusterNamePager - Get list of configuration profile assignments
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Arc machine.
//   - options - ConfigurationProfileAssignmentsClientListByClusterNameOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.NewListByClusterNamePager
//     method.
func (client *ConfigurationProfileAssignmentsClient) NewListByClusterNamePager(resourceGroupName string, clusterName string, options *ConfigurationProfileAssignmentsClientListByClusterNameOptions) (*runtime.Pager[ConfigurationProfileAssignmentsClientListByClusterNameResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationProfileAssignmentsClientListByClusterNameResponse]{
		More: func(page ConfigurationProfileAssignmentsClientListByClusterNameResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfileAssignmentsClientListByClusterNameResponse) (ConfigurationProfileAssignmentsClientListByClusterNameResponse, error) {
			req, err := client.listByClusterNameCreateRequest(ctx, resourceGroupName, clusterName, options)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListByClusterNameResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListByClusterNameResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfileAssignmentsClientListByClusterNameResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterNameHandleResponse(resp)
		},
	})
}

// listByClusterNameCreateRequest creates the ListByClusterName request.
func (client *ConfigurationProfileAssignmentsClient) listByClusterNameCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConfigurationProfileAssignmentsClientListByClusterNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHci/clusters/{clusterName}/providers/Microsoft.Automanage/configurationProfileAssignments"
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
	reqQP.Set("api-version", "2022-05-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterNameHandleResponse handles the ListByClusterName response.
func (client *ConfigurationProfileAssignmentsClient) listByClusterNameHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientListByClusterNameResponse, error) {
	result := ConfigurationProfileAssignmentsClientListByClusterNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignmentList); err != nil {
		return ConfigurationProfileAssignmentsClientListByClusterNameResponse{}, err
	}
	return result, nil
}

// NewListByMachineNamePager - Get list of configuration profile assignments
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the Arc machine.
//   - options - ConfigurationProfileAssignmentsClientListByMachineNameOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.NewListByMachineNamePager
//     method.
func (client *ConfigurationProfileAssignmentsClient) NewListByMachineNamePager(resourceGroupName string, machineName string, options *ConfigurationProfileAssignmentsClientListByMachineNameOptions) (*runtime.Pager[ConfigurationProfileAssignmentsClientListByMachineNameResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationProfileAssignmentsClientListByMachineNameResponse]{
		More: func(page ConfigurationProfileAssignmentsClientListByMachineNameResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfileAssignmentsClientListByMachineNameResponse) (ConfigurationProfileAssignmentsClientListByMachineNameResponse, error) {
			req, err := client.listByMachineNameCreateRequest(ctx, resourceGroupName, machineName, options)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListByMachineNameResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListByMachineNameResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfileAssignmentsClientListByMachineNameResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByMachineNameHandleResponse(resp)
		},
	})
}

// listByMachineNameCreateRequest creates the ListByMachineName request.
func (client *ConfigurationProfileAssignmentsClient) listByMachineNameCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *ConfigurationProfileAssignmentsClientListByMachineNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.Automanage/configurationProfileAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
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

// listByMachineNameHandleResponse handles the ListByMachineName response.
func (client *ConfigurationProfileAssignmentsClient) listByMachineNameHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientListByMachineNameResponse, error) {
	result := ConfigurationProfileAssignmentsClientListByMachineNameResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignmentList); err != nil {
		return ConfigurationProfileAssignmentsClientListByMachineNameResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get list of configuration profile assignments under a given subscription
//
// Generated from API version 2022-05-04
//   - options - ConfigurationProfileAssignmentsClientListBySubscriptionOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.NewListBySubscriptionPager
//     method.
func (client *ConfigurationProfileAssignmentsClient) NewListBySubscriptionPager(options *ConfigurationProfileAssignmentsClientListBySubscriptionOptions) (*runtime.Pager[ConfigurationProfileAssignmentsClientListBySubscriptionResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationProfileAssignmentsClientListBySubscriptionResponse]{
		More: func(page ConfigurationProfileAssignmentsClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfileAssignmentsClientListBySubscriptionResponse) (ConfigurationProfileAssignmentsClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConfigurationProfileAssignmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConfigurationProfileAssignmentsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Automanage/configurationProfileAssignments"
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
func (client *ConfigurationProfileAssignmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientListBySubscriptionResponse, error) {
	result := ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignmentList); err != nil {
		return ConfigurationProfileAssignmentsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListByVirtualMachinesPager - Get list of configuration profile assignments
//
// Generated from API version 2022-05-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vmName - The name of the virtual machine.
//   - options - ConfigurationProfileAssignmentsClientListByVirtualMachinesOptions contains the optional parameters for the ConfigurationProfileAssignmentsClient.NewListByVirtualMachinesPager
//     method.
func (client *ConfigurationProfileAssignmentsClient) NewListByVirtualMachinesPager(resourceGroupName string, vmName string, options *ConfigurationProfileAssignmentsClientListByVirtualMachinesOptions) (*runtime.Pager[ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse]{
		More: func(page ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse) (ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse, error) {
			req, err := client.listByVirtualMachinesCreateRequest(ctx, resourceGroupName, vmName, options)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByVirtualMachinesHandleResponse(resp)
		},
	})
}

// listByVirtualMachinesCreateRequest creates the ListByVirtualMachines request.
func (client *ConfigurationProfileAssignmentsClient) listByVirtualMachinesCreateRequest(ctx context.Context, resourceGroupName string, vmName string, options *ConfigurationProfileAssignmentsClientListByVirtualMachinesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.Automanage/configurationProfileAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
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

// listByVirtualMachinesHandleResponse handles the ListByVirtualMachines response.
func (client *ConfigurationProfileAssignmentsClient) listByVirtualMachinesHandleResponse(resp *http.Response) (ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse, error) {
	result := ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationProfileAssignmentList); err != nil {
		return ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse{}, err
	}
	return result, nil
}

