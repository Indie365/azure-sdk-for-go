//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfrontdoor

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

// FrontDoorsClient contains the methods for the FrontDoors group.
// Don't use this type directly, use NewFrontDoorsClient() instead.
type FrontDoorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFrontDoorsClient creates a new instance of FrontDoorsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFrontDoorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FrontDoorsClient, error) {
	cl, err := arm.NewClient(moduleName+".FrontDoorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FrontDoorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new Front Door with a Front Door name under the specified subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - frontDoorName - Name of the Front Door which is globally unique.
//   - frontDoorParameters - Front Door properties needed to create a new Front Door.
//   - options - FrontDoorsClientBeginCreateOrUpdateOptions contains the optional parameters for the FrontDoorsClient.BeginCreateOrUpdate
//     method.
func (client *FrontDoorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters FrontDoor, options *FrontDoorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[FrontDoorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, frontDoorName, frontDoorParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FrontDoorsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FrontDoorsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new Front Door with a Front Door name under the specified subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
func (client *FrontDoorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters FrontDoor, options *FrontDoorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, frontDoorName, frontDoorParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FrontDoorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters FrontDoor, options *FrontDoorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, frontDoorParameters)
}

// BeginDelete - Deletes an existing Front Door with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - frontDoorName - Name of the Front Door which is globally unique.
//   - options - FrontDoorsClientBeginDeleteOptions contains the optional parameters for the FrontDoorsClient.BeginDelete method.
func (client *FrontDoorsClient) BeginDelete(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientBeginDeleteOptions) (*runtime.Poller[FrontDoorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, frontDoorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FrontDoorsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[FrontDoorsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an existing Front Door with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
func (client *FrontDoorsClient) deleteOperation(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, frontDoorName, options)
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
func (client *FrontDoorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Front Door with the specified Front Door name under the specified subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - frontDoorName - Name of the Front Door which is globally unique.
//   - options - FrontDoorsClientGetOptions contains the optional parameters for the FrontDoorsClient.Get method.
func (client *FrontDoorsClient) Get(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientGetOptions) (FrontDoorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, frontDoorName, options)
	if err != nil {
		return FrontDoorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FrontDoorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FrontDoorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FrontDoorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, options *FrontDoorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FrontDoorsClient) getHandleResponse(resp *http.Response) (FrontDoorsClientGetResponse, error) {
	result := FrontDoorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FrontDoor); err != nil {
		return FrontDoorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the Front Doors within an Azure subscription.
//
// Generated from API version 2020-05-01
//   - options - FrontDoorsClientListOptions contains the optional parameters for the FrontDoorsClient.NewListPager method.
func (client *FrontDoorsClient) NewListPager(options *FrontDoorsClientListOptions) *runtime.Pager[FrontDoorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FrontDoorsClientListResponse]{
		More: func(page FrontDoorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FrontDoorsClientListResponse) (FrontDoorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FrontDoorsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FrontDoorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FrontDoorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FrontDoorsClient) listCreateRequest(ctx context.Context, options *FrontDoorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/frontDoors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FrontDoorsClient) listHandleResponse(resp *http.Response) (FrontDoorsClientListResponse, error) {
	result := FrontDoorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return FrontDoorsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all of the Front Doors within a resource group under a subscription.
//
// Generated from API version 2020-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - options - FrontDoorsClientListByResourceGroupOptions contains the optional parameters for the FrontDoorsClient.NewListByResourceGroupPager
//     method.
func (client *FrontDoorsClient) NewListByResourceGroupPager(resourceGroupName string, options *FrontDoorsClientListByResourceGroupOptions) *runtime.Pager[FrontDoorsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[FrontDoorsClientListByResourceGroupResponse]{
		More: func(page FrontDoorsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FrontDoorsClientListByResourceGroupResponse) (FrontDoorsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FrontDoorsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FrontDoorsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FrontDoorsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FrontDoorsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FrontDoorsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors"
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
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FrontDoorsClient) listByResourceGroupHandleResponse(resp *http.Response) (FrontDoorsClientListByResourceGroupResponse, error) {
	result := FrontDoorsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return FrontDoorsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ValidateCustomDomain - Validates the custom domain mapping to ensure it maps to the correct Front Door endpoint in DNS.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-05-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - frontDoorName - Name of the Front Door which is globally unique.
//   - customDomainProperties - Custom domain to be validated.
//   - options - FrontDoorsClientValidateCustomDomainOptions contains the optional parameters for the FrontDoorsClient.ValidateCustomDomain
//     method.
func (client *FrontDoorsClient) ValidateCustomDomain(ctx context.Context, resourceGroupName string, frontDoorName string, customDomainProperties ValidateCustomDomainInput, options *FrontDoorsClientValidateCustomDomainOptions) (FrontDoorsClientValidateCustomDomainResponse, error) {
	req, err := client.validateCustomDomainCreateRequest(ctx, resourceGroupName, frontDoorName, customDomainProperties, options)
	if err != nil {
		return FrontDoorsClientValidateCustomDomainResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FrontDoorsClientValidateCustomDomainResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FrontDoorsClientValidateCustomDomainResponse{}, runtime.NewResponseError(resp)
	}
	return client.validateCustomDomainHandleResponse(resp)
}

// validateCustomDomainCreateRequest creates the ValidateCustomDomain request.
func (client *FrontDoorsClient) validateCustomDomainCreateRequest(ctx context.Context, resourceGroupName string, frontDoorName string, customDomainProperties ValidateCustomDomainInput, options *FrontDoorsClientValidateCustomDomainOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/frontDoors/{frontDoorName}/validateCustomDomain"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if frontDoorName == "" {
		return nil, errors.New("parameter frontDoorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{frontDoorName}", url.PathEscape(frontDoorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, customDomainProperties)
}

// validateCustomDomainHandleResponse handles the ValidateCustomDomain response.
func (client *FrontDoorsClient) validateCustomDomainHandleResponse(resp *http.Response) (FrontDoorsClientValidateCustomDomainResponse, error) {
	result := FrontDoorsClientValidateCustomDomainResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ValidateCustomDomainOutput); err != nil {
		return FrontDoorsClientValidateCustomDomainResponse{}, err
	}
	return result, nil
}
