//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtrafficmanager

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

// ProfilesClient contains the methods for the Profiles group.
// Don't use this type directly, use NewProfilesClient() instead.
type ProfilesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProfilesClient creates a new instance of ProfilesClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProfilesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProfilesClient, error) {
	cl, err := arm.NewClient(moduleName+".ProfilesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProfilesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckTrafficManagerRelativeDNSNameAvailability - Checks the availability of a Traffic Manager Relative DNS name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-08-01
//   - parameters - The Traffic Manager name parameters supplied to the CheckTrafficManagerNameAvailability operation.
//   - options - ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityOptions contains the optional parameters for the
//     ProfilesClient.CheckTrafficManagerRelativeDNSNameAvailability method.
func (client *ProfilesClient) CheckTrafficManagerRelativeDNSNameAvailability(ctx context.Context, parameters CheckTrafficManagerRelativeDNSNameAvailabilityParameters, options *ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityOptions) (ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityResponse, error) {
	req, err := client.checkTrafficManagerRelativeDNSNameAvailabilityCreateRequest(ctx, parameters, options)
	if err != nil {
		return ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkTrafficManagerRelativeDNSNameAvailabilityHandleResponse(resp)
}

// checkTrafficManagerRelativeDNSNameAvailabilityCreateRequest creates the CheckTrafficManagerRelativeDNSNameAvailability request.
func (client *ProfilesClient) checkTrafficManagerRelativeDNSNameAvailabilityCreateRequest(ctx context.Context, parameters CheckTrafficManagerRelativeDNSNameAvailabilityParameters, options *ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Network/checkTrafficManagerNameAvailability"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkTrafficManagerRelativeDNSNameAvailabilityHandleResponse handles the CheckTrafficManagerRelativeDNSNameAvailability response.
func (client *ProfilesClient) checkTrafficManagerRelativeDNSNameAvailabilityHandleResponse(resp *http.Response) (ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityResponse, error) {
	result := ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NameAvailability); err != nil {
		return ProfilesClientCheckTrafficManagerRelativeDNSNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - Create or update a Traffic Manager profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-08-01
//   - resourceGroupName - The name of the resource group containing the Traffic Manager profile.
//   - profileName - The name of the Traffic Manager profile.
//   - parameters - The Traffic Manager profile parameters supplied to the CreateOrUpdate operation.
//   - options - ProfilesClientCreateOrUpdateOptions contains the optional parameters for the ProfilesClient.CreateOrUpdate method.
func (client *ProfilesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesClientCreateOrUpdateOptions) (ProfilesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, profileName, parameters, options)
	if err != nil {
		return ProfilesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProfilesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProfilesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProfilesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProfilesClient) createOrUpdateHandleResponse(resp *http.Response) (ProfilesClientCreateOrUpdateResponse, error) {
	result := ProfilesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Traffic Manager profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-08-01
//   - resourceGroupName - The name of the resource group containing the Traffic Manager profile to be deleted.
//   - profileName - The name of the Traffic Manager profile to be deleted.
//   - options - ProfilesClientDeleteOptions contains the optional parameters for the ProfilesClient.Delete method.
func (client *ProfilesClient) Delete(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesClientDeleteOptions) (ProfilesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return ProfilesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProfilesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProfilesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ProfilesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ProfilesClient) deleteHandleResponse(resp *http.Response) (ProfilesClientDeleteResponse, error) {
	result := ProfilesClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeleteOperationResult); err != nil {
		return ProfilesClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Gets a Traffic Manager profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-08-01
//   - resourceGroupName - The name of the resource group containing the Traffic Manager profile.
//   - profileName - The name of the Traffic Manager profile.
//   - options - ProfilesClientGetOptions contains the optional parameters for the ProfilesClient.Get method.
func (client *ProfilesClient) Get(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesClientGetOptions) (ProfilesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProfilesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProfilesClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *ProfilesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProfilesClient) getHandleResponse(resp *http.Response) (ProfilesClientGetResponse, error) {
	result := ProfilesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all Traffic Manager profiles within a resource group.
//
// Generated from API version 2018-08-01
//   - resourceGroupName - The name of the resource group containing the Traffic Manager profiles to be listed.
//   - options - ProfilesClientListByResourceGroupOptions contains the optional parameters for the ProfilesClient.NewListByResourceGroupPager
//     method.
func (client *ProfilesClient) NewListByResourceGroupPager(resourceGroupName string, options *ProfilesClientListByResourceGroupOptions) *runtime.Pager[ProfilesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProfilesClientListByResourceGroupResponse]{
		More: func(page ProfilesClientListByResourceGroupResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ProfilesClientListByResourceGroupResponse) (ProfilesClientListByResourceGroupResponse, error) {
			req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return ProfilesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProfilesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProfilesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ProfilesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ProfilesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles"
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
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ProfilesClient) listByResourceGroupHandleResponse(resp *http.Response) (ProfilesClientListByResourceGroupResponse, error) {
	result := ProfilesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all Traffic Manager profiles within a subscription.
//
// Generated from API version 2018-08-01
//   - options - ProfilesClientListBySubscriptionOptions contains the optional parameters for the ProfilesClient.NewListBySubscriptionPager
//     method.
func (client *ProfilesClient) NewListBySubscriptionPager(options *ProfilesClientListBySubscriptionOptions) *runtime.Pager[ProfilesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProfilesClientListBySubscriptionResponse]{
		More: func(page ProfilesClientListBySubscriptionResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ProfilesClientListBySubscriptionResponse) (ProfilesClientListBySubscriptionResponse, error) {
			req, err := client.listBySubscriptionCreateRequest(ctx, options)
			if err != nil {
				return ProfilesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProfilesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProfilesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProfilesClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProfilesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficmanagerprofiles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProfilesClient) listBySubscriptionHandleResponse(resp *http.Response) (ProfilesClientListBySubscriptionResponse, error) {
	result := ProfilesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProfileListResult); err != nil {
		return ProfilesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a Traffic Manager profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-08-01
//   - resourceGroupName - The name of the resource group containing the Traffic Manager profile.
//   - profileName - The name of the Traffic Manager profile.
//   - parameters - The Traffic Manager profile parameters supplied to the Update operation.
//   - options - ProfilesClientUpdateOptions contains the optional parameters for the ProfilesClient.Update method.
func (client *ProfilesClient) Update(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesClientUpdateOptions) (ProfilesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, profileName, parameters, options)
	if err != nil {
		return ProfilesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProfilesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProfilesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProfilesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, profileName string, parameters Profile, options *ProfilesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ProfilesClient) updateHandleResponse(resp *http.Response) (ProfilesClientUpdateResponse, error) {
	result := ProfilesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Profile); err != nil {
		return ProfilesClientUpdateResponse{}, err
	}
	return result, nil
}
