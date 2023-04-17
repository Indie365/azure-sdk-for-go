//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement

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

// ServiceClient contains the methods for the APIManagementService group.
// Don't use this type directly, use NewServiceClient() instead.
type ServiceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceClient creates a new instance of ServiceClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceClient, error) {
	cl, err := arm.NewClient(moduleName+".ServiceClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginApplyNetworkConfigurationUpdates - Updates the Microsoft.ApiManagement resource running in the Virtual network to
// pick the updated DNS changes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - options - ServiceClientBeginApplyNetworkConfigurationUpdatesOptions contains the optional parameters for the ServiceClient.BeginApplyNetworkConfigurationUpdates
//     method.
func (client *ServiceClient) BeginApplyNetworkConfigurationUpdates(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientBeginApplyNetworkConfigurationUpdatesOptions) (*runtime.Poller[ServiceClientApplyNetworkConfigurationUpdatesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.applyNetworkConfigurationUpdates(ctx, resourceGroupName, serviceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceClientApplyNetworkConfigurationUpdatesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientApplyNetworkConfigurationUpdatesResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ApplyNetworkConfigurationUpdates - Updates the Microsoft.ApiManagement resource running in the Virtual network to pick
// the updated DNS changes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *ServiceClient) applyNetworkConfigurationUpdates(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientBeginApplyNetworkConfigurationUpdatesOptions) (*http.Response, error) {
	req, err := client.applyNetworkConfigurationUpdatesCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// applyNetworkConfigurationUpdatesCreateRequest creates the ApplyNetworkConfigurationUpdates request.
func (client *ServiceClient) applyNetworkConfigurationUpdatesCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientBeginApplyNetworkConfigurationUpdatesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/applynetworkconfigurationupdates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// BeginBackup - Creates a backup of the API Management service to the given Azure Storage Account. This is long running operation
// and could take several minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - parameters - Parameters supplied to the ApiManagementService_Backup operation.
//   - options - ServiceClientBeginBackupOptions contains the optional parameters for the ServiceClient.BeginBackup method.
func (client *ServiceClient) BeginBackup(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceBackupRestoreParameters, options *ServiceClientBeginBackupOptions) (*runtime.Poller[ServiceClientBackupResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.backup(ctx, resourceGroupName, serviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceClientBackupResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientBackupResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Backup - Creates a backup of the API Management service to the given Azure Storage Account. This is long running operation
// and could take several minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *ServiceClient) backup(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceBackupRestoreParameters, options *ServiceClientBeginBackupOptions) (*http.Response, error) {
	req, err := client.backupCreateRequest(ctx, resourceGroupName, serviceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// backupCreateRequest creates the Backup request.
func (client *ServiceClient) backupCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceBackupRestoreParameters, options *ServiceClientBeginBackupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/backup"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// CheckNameAvailability - Checks availability and correctness of a name for an API Management service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - parameters - Parameters supplied to the CheckNameAvailability operation.
//   - options - ServiceClientCheckNameAvailabilityOptions contains the optional parameters for the ServiceClient.CheckNameAvailability
//     method.
func (client *ServiceClient) CheckNameAvailability(ctx context.Context, parameters ServiceCheckNameAvailabilityParameters, options *ServiceClientCheckNameAvailabilityOptions) (ServiceClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, parameters, options)
	if err != nil {
		return ServiceClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *ServiceClient) checkNameAvailabilityCreateRequest(ctx context.Context, parameters ServiceCheckNameAvailabilityParameters, options *ServiceClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *ServiceClient) checkNameAvailabilityHandleResponse(resp *http.Response) (ServiceClientCheckNameAvailabilityResponse, error) {
	result := ServiceClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceNameAvailabilityResult); err != nil {
		return ServiceClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates an API Management service. This is long running operation and could take several
// minutes to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - parameters - Parameters supplied to the CreateOrUpdate API Management service operation.
//   - options - ServiceClientBeginCreateOrUpdateOptions contains the optional parameters for the ServiceClient.BeginCreateOrUpdate
//     method.
func (client *ServiceClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceResource, options *ServiceClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServiceClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an API Management service. This is long running operation and could take several minutes
// to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *ServiceClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceResource, options *ServiceClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, parameters, options)
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
func (client *ServiceClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceResource, options *ServiceClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes an existing API Management service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - options - ServiceClientBeginDeleteOptions contains the optional parameters for the ServiceClient.BeginDelete method.
func (client *ServiceClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientBeginDeleteOptions) (*runtime.Poller[ServiceClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an existing API Management service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *ServiceClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, options)
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
func (client *ServiceClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an API Management service resource description.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - options - ServiceClientGetOptions contains the optional parameters for the ServiceClient.Get method.
func (client *ServiceClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientGetOptions) (ServiceClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceClient) getHandleResponse(resp *http.Response) (ServiceClientGetResponse, error) {
	result := ServiceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceResource); err != nil {
		return ServiceClientGetResponse{}, err
	}
	return result, nil
}

// GetDomainOwnershipIdentifier - Get the custom domain ownership identifier for an API Management service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - options - ServiceClientGetDomainOwnershipIdentifierOptions contains the optional parameters for the ServiceClient.GetDomainOwnershipIdentifier
//     method.
func (client *ServiceClient) GetDomainOwnershipIdentifier(ctx context.Context, options *ServiceClientGetDomainOwnershipIdentifierOptions) (ServiceClientGetDomainOwnershipIdentifierResponse, error) {
	req, err := client.getDomainOwnershipIdentifierCreateRequest(ctx, options)
	if err != nil {
		return ServiceClientGetDomainOwnershipIdentifierResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetDomainOwnershipIdentifierResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetDomainOwnershipIdentifierResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDomainOwnershipIdentifierHandleResponse(resp)
}

// getDomainOwnershipIdentifierCreateRequest creates the GetDomainOwnershipIdentifier request.
func (client *ServiceClient) getDomainOwnershipIdentifierCreateRequest(ctx context.Context, options *ServiceClientGetDomainOwnershipIdentifierOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/getDomainOwnershipIdentifier"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDomainOwnershipIdentifierHandleResponse handles the GetDomainOwnershipIdentifier response.
func (client *ServiceClient) getDomainOwnershipIdentifierHandleResponse(resp *http.Response) (ServiceClientGetDomainOwnershipIdentifierResponse, error) {
	result := ServiceClientGetDomainOwnershipIdentifierResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceGetDomainOwnershipIdentifierResult); err != nil {
		return ServiceClientGetDomainOwnershipIdentifierResponse{}, err
	}
	return result, nil
}

// GetSsoToken - Gets the Single-Sign-On token for the API Management Service which is valid for 5 Minutes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - options - ServiceClientGetSsoTokenOptions contains the optional parameters for the ServiceClient.GetSsoToken method.
func (client *ServiceClient) GetSsoToken(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientGetSsoTokenOptions) (ServiceClientGetSsoTokenResponse, error) {
	req, err := client.getSsoTokenCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ServiceClientGetSsoTokenResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceClientGetSsoTokenResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceClientGetSsoTokenResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSsoTokenHandleResponse(resp)
}

// getSsoTokenCreateRequest creates the GetSsoToken request.
func (client *ServiceClient) getSsoTokenCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ServiceClientGetSsoTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/getssotoken"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSsoTokenHandleResponse handles the GetSsoToken response.
func (client *ServiceClient) getSsoTokenHandleResponse(resp *http.Response) (ServiceClientGetSsoTokenResponse, error) {
	result := ServiceClientGetSsoTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceGetSsoTokenResult); err != nil {
		return ServiceClientGetSsoTokenResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all API Management services within an Azure subscription.
//
// Generated from API version 2021-08-01
//   - options - ServiceClientListOptions contains the optional parameters for the ServiceClient.NewListPager method.
func (client *ServiceClient) NewListPager(options *ServiceClientListOptions) *runtime.Pager[ServiceClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceClientListResponse]{
		More: func(page ServiceClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceClientListResponse) (ServiceClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServiceClient) listCreateRequest(ctx context.Context, options *ServiceClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ApiManagement/service"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceClient) listHandleResponse(resp *http.Response) (ServiceClientListResponse, error) {
	result := ServiceClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceListResult); err != nil {
		return ServiceClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all API Management services within a resource group.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - options - ServiceClientListByResourceGroupOptions contains the optional parameters for the ServiceClient.NewListByResourceGroupPager
//     method.
func (client *ServiceClient) NewListByResourceGroupPager(resourceGroupName string, options *ServiceClientListByResourceGroupOptions) *runtime.Pager[ServiceClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceClientListByResourceGroupResponse]{
		More: func(page ServiceClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceClientListByResourceGroupResponse) (ServiceClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServiceClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ServiceClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service"
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
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServiceClient) listByResourceGroupHandleResponse(resp *http.Response) (ServiceClientListByResourceGroupResponse, error) {
	result := ServiceClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceListResult); err != nil {
		return ServiceClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginRestore - Restores a backup of an API Management service created using the ApiManagementService_Backup operation on
// the current service. This is a long running operation and could take several minutes to
// complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - parameters - Parameters supplied to the Restore API Management service from backup operation.
//   - options - ServiceClientBeginRestoreOptions contains the optional parameters for the ServiceClient.BeginRestore method.
func (client *ServiceClient) BeginRestore(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceBackupRestoreParameters, options *ServiceClientBeginRestoreOptions) (*runtime.Poller[ServiceClientRestoreResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restore(ctx, resourceGroupName, serviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceClientRestoreResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientRestoreResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Restore - Restores a backup of an API Management service created using the ApiManagementService_Backup operation on the
// current service. This is a long running operation and could take several minutes to
// complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *ServiceClient) restore(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceBackupRestoreParameters, options *ServiceClientBeginRestoreOptions) (*http.Response, error) {
	req, err := client.restoreCreateRequest(ctx, resourceGroupName, serviceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// restoreCreateRequest creates the Restore request.
func (client *ServiceClient) restoreCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceBackupRestoreParameters, options *ServiceClientBeginRestoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/restore"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginUpdate - Updates an existing API Management service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - parameters - Parameters supplied to the CreateOrUpdate API Management service operation.
//   - options - ServiceClientBeginUpdateOptions contains the optional parameters for the ServiceClient.BeginUpdate method.
func (client *ServiceClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceUpdateParameters, options *ServiceClientBeginUpdateOptions) (*runtime.Poller[ServiceClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serviceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServiceClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServiceClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates an existing API Management service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *ServiceClient) update(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceUpdateParameters, options *ServiceClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ServiceClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, parameters ServiceUpdateParameters, options *ServiceClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
