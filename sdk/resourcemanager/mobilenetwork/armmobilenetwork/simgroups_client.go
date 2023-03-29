//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork

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

// SimGroupsClient contains the methods for the SimGroups group.
// Don't use this type directly, use NewSimGroupsClient() instead.
type SimGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSimGroupsClient creates a new instance of SimGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSimGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SimGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".SimGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SimGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - parameters - Parameters supplied to the create or update SIM group operation.
//   - options - SimGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the SimGroupsClient.BeginCreateOrUpdate
//     method.
func (client *SimGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimGroup, options *SimGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SimGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, simGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SimGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *SimGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimGroup, options *SimGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, simGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SimGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, parameters SimGroup, options *SimGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - options - SimGroupsClientBeginDeleteOptions contains the optional parameters for the SimGroupsClient.BeginDelete method.
func (client *SimGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, simGroupName string, options *SimGroupsClientBeginDeleteOptions) (*runtime.Poller[SimGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, simGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SimGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SimGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *SimGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, simGroupName string, options *SimGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, simGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SimGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, options *SimGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified SIM group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - options - SimGroupsClientGetOptions contains the optional parameters for the SimGroupsClient.Get method.
func (client *SimGroupsClient) Get(ctx context.Context, resourceGroupName string, simGroupName string, options *SimGroupsClientGetOptions) (SimGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, simGroupName, options)
	if err != nil {
		return SimGroupsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SimGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SimGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SimGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, options *SimGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SimGroupsClient) getHandleResponse(resp *http.Response) (SimGroupsClientGetResponse, error) {
	result := SimGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimGroup); err != nil {
		return SimGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all the SIM groups in a resource group.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SimGroupsClientListByResourceGroupOptions contains the optional parameters for the SimGroupsClient.NewListByResourceGroupPager
//     method.
func (client *SimGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *SimGroupsClientListByResourceGroupOptions) *runtime.Pager[SimGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SimGroupsClientListByResourceGroupResponse]{
		More: func(page SimGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SimGroupsClientListByResourceGroupResponse) (SimGroupsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SimGroupsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SimGroupsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SimGroupsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SimGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SimGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups"
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
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SimGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (SimGroupsClientListByResourceGroupResponse, error) {
	result := SimGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimGroupListResult); err != nil {
		return SimGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets all the SIM groups in a subscription.
//
// Generated from API version 2022-11-01
//   - options - SimGroupsClientListBySubscriptionOptions contains the optional parameters for the SimGroupsClient.NewListBySubscriptionPager
//     method.
func (client *SimGroupsClient) NewListBySubscriptionPager(options *SimGroupsClientListBySubscriptionOptions) *runtime.Pager[SimGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SimGroupsClientListBySubscriptionResponse]{
		More: func(page SimGroupsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SimGroupsClientListBySubscriptionResponse) (SimGroupsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SimGroupsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SimGroupsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SimGroupsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SimGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SimGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MobileNetwork/simGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SimGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (SimGroupsClientListBySubscriptionResponse, error) {
	result := SimGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimGroupListResult); err != nil {
		return SimGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates SIM group tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - simGroupName - The name of the SIM Group.
//   - parameters - Parameters supplied to update SIM group tags.
//   - options - SimGroupsClientUpdateTagsOptions contains the optional parameters for the SimGroupsClient.UpdateTags method.
func (client *SimGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, simGroupName string, parameters TagsObject, options *SimGroupsClientUpdateTagsOptions) (SimGroupsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, simGroupName, parameters, options)
	if err != nil {
		return SimGroupsClientUpdateTagsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SimGroupsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SimGroupsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SimGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, simGroupName string, parameters TagsObject, options *SimGroupsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/simGroups/{simGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if simGroupName == "" {
		return nil, errors.New("parameter simGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{simGroupName}", url.PathEscape(simGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SimGroupsClient) updateTagsHandleResponse(resp *http.Response) (SimGroupsClientUpdateTagsResponse, error) {
	result := SimGroupsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SimGroup); err != nil {
		return SimGroupsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
