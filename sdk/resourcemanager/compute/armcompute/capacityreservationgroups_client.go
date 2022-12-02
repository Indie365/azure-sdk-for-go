//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

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

// CapacityReservationGroupsClient contains the methods for the CapacityReservationGroups group.
// Don't use this type directly, use NewCapacityReservationGroupsClient() instead.
type CapacityReservationGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCapacityReservationGroupsClient creates a new instance of CapacityReservationGroupsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCapacityReservationGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CapacityReservationGroupsClient, error) {
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
	client := &CapacityReservationGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - The operation to create or update a capacity reservation group. When updating a capacity reservation group,
// only tags may be modified. Please refer to https://aka.ms/CapacityReservation for more
// details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// capacityReservationGroupName - The name of the capacity reservation group.
// parameters - Parameters supplied to the Create capacity reservation Group.
// options - CapacityReservationGroupsClientCreateOrUpdateOptions contains the optional parameters for the CapacityReservationGroupsClient.CreateOrUpdate
// method.
func (client *CapacityReservationGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroup, options *CapacityReservationGroupsClientCreateOrUpdateOptions) (CapacityReservationGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, parameters, options)
	if err != nil {
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CapacityReservationGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroup, options *CapacityReservationGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CapacityReservationGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (CapacityReservationGroupsClientCreateOrUpdateResponse, error) {
	result := CapacityReservationGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The operation to delete a capacity reservation group. This operation is allowed only if all the associated resources
// are disassociated from the reservation group and all capacity reservations under
// the reservation group have also been deleted. Please refer to https://aka.ms/CapacityReservation for more details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// capacityReservationGroupName - The name of the capacity reservation group.
// options - CapacityReservationGroupsClientDeleteOptions contains the optional parameters for the CapacityReservationGroupsClient.Delete
// method.
func (client *CapacityReservationGroupsClient) Delete(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsClientDeleteOptions) (CapacityReservationGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, options)
	if err != nil {
		return CapacityReservationGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CapacityReservationGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CapacityReservationGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CapacityReservationGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation that retrieves information about a capacity reservation group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// capacityReservationGroupName - The name of the capacity reservation group.
// options - CapacityReservationGroupsClientGetOptions contains the optional parameters for the CapacityReservationGroupsClient.Get
// method.
func (client *CapacityReservationGroupsClient) Get(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsClientGetOptions) (CapacityReservationGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, options)
	if err != nil {
		return CapacityReservationGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacityReservationGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CapacityReservationGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, options *CapacityReservationGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CapacityReservationGroupsClient) getHandleResponse(resp *http.Response) (CapacityReservationGroupsClientGetResponse, error) {
	result := CapacityReservationGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all of the capacity reservation groups in the specified resource group. Use the nextLink
// property in the response to get the next page of capacity reservation groups.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// options - CapacityReservationGroupsClientListByResourceGroupOptions contains the optional parameters for the CapacityReservationGroupsClient.ListByResourceGroup
// method.
func (client *CapacityReservationGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *CapacityReservationGroupsClientListByResourceGroupOptions) *runtime.Pager[CapacityReservationGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacityReservationGroupsClientListByResourceGroupResponse]{
		More: func(page CapacityReservationGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacityReservationGroupsClientListByResourceGroupResponse) (CapacityReservationGroupsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CapacityReservationGroupsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CapacityReservationGroupsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CapacityReservationGroupsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CapacityReservationGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CapacityReservationGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CapacityReservationGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (CapacityReservationGroupsClientListByResourceGroupResponse, error) {
	result := CapacityReservationGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroupListResult); err != nil {
		return CapacityReservationGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all of the capacity reservation groups in the subscription. Use the nextLink property
// in the response to get the next page of capacity reservation groups.
// Generated from API version 2022-08-01
// options - CapacityReservationGroupsClientListBySubscriptionOptions contains the optional parameters for the CapacityReservationGroupsClient.ListBySubscription
// method.
func (client *CapacityReservationGroupsClient) NewListBySubscriptionPager(options *CapacityReservationGroupsClientListBySubscriptionOptions) *runtime.Pager[CapacityReservationGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CapacityReservationGroupsClientListBySubscriptionResponse]{
		More: func(page CapacityReservationGroupsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CapacityReservationGroupsClientListBySubscriptionResponse) (CapacityReservationGroupsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CapacityReservationGroupsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CapacityReservationGroupsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CapacityReservationGroupsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CapacityReservationGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CapacityReservationGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/capacityReservationGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CapacityReservationGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (CapacityReservationGroupsClientListBySubscriptionResponse, error) {
	result := CapacityReservationGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroupListResult); err != nil {
		return CapacityReservationGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - The operation to update a capacity reservation group. When updating a capacity reservation group, only tags may
// be modified.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-08-01
// resourceGroupName - The name of the resource group.
// capacityReservationGroupName - The name of the capacity reservation group.
// parameters - Parameters supplied to the Update capacity reservation Group operation.
// options - CapacityReservationGroupsClientUpdateOptions contains the optional parameters for the CapacityReservationGroupsClient.Update
// method.
func (client *CapacityReservationGroupsClient) Update(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroupUpdate, options *CapacityReservationGroupsClientUpdateOptions) (CapacityReservationGroupsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, capacityReservationGroupName, parameters, options)
	if err != nil {
		return CapacityReservationGroupsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacityReservationGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacityReservationGroupsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CapacityReservationGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, parameters CapacityReservationGroupUpdate, options *CapacityReservationGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if capacityReservationGroupName == "" {
		return nil, errors.New("parameter capacityReservationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{capacityReservationGroupName}", url.PathEscape(capacityReservationGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *CapacityReservationGroupsClient) updateHandleResponse(resp *http.Response) (CapacityReservationGroupsClientUpdateResponse, error) {
	result := CapacityReservationGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapacityReservationGroup); err != nil {
		return CapacityReservationGroupsClientUpdateResponse{}, err
	}
	return result, nil
}