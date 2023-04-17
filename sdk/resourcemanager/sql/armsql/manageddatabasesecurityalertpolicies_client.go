//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql

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

// ManagedDatabaseSecurityAlertPoliciesClient contains the methods for the ManagedDatabaseSecurityAlertPolicies group.
// Don't use this type directly, use NewManagedDatabaseSecurityAlertPoliciesClient() instead.
type ManagedDatabaseSecurityAlertPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedDatabaseSecurityAlertPoliciesClient creates a new instance of ManagedDatabaseSecurityAlertPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedDatabaseSecurityAlertPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedDatabaseSecurityAlertPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedDatabaseSecurityAlertPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedDatabaseSecurityAlertPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a database's security alert policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the managed database for which the security alert policy is defined.
//   - securityAlertPolicyName - The name of the security alert policy.
//   - parameters - The database security alert policy.
//   - options - ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ManagedDatabaseSecurityAlertPoliciesClient.CreateOrUpdate
//     method.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, parameters ManagedDatabaseSecurityAlertPolicy, options *ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateOptions) (ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, securityAlertPolicyName, parameters, options)
	if err != nil {
		return ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, parameters ManagedDatabaseSecurityAlertPolicy, options *ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse, error) {
	result := ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseSecurityAlertPolicy); err != nil {
		return ManagedDatabaseSecurityAlertPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a managed database's security alert policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the managed database for which the security alert policy is defined.
//   - securityAlertPolicyName - The name of the security alert policy.
//   - options - ManagedDatabaseSecurityAlertPoliciesClientGetOptions contains the optional parameters for the ManagedDatabaseSecurityAlertPoliciesClient.Get
//     method.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, options *ManagedDatabaseSecurityAlertPoliciesClientGetOptions) (ManagedDatabaseSecurityAlertPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, securityAlertPolicyName, options)
	if err != nil {
		return ManagedDatabaseSecurityAlertPoliciesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedDatabaseSecurityAlertPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedDatabaseSecurityAlertPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, securityAlertPolicyName SecurityAlertPolicyName, options *ManagedDatabaseSecurityAlertPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) getHandleResponse(resp *http.Response) (ManagedDatabaseSecurityAlertPoliciesClientGetResponse, error) {
	result := ManagedDatabaseSecurityAlertPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseSecurityAlertPolicy); err != nil {
		return ManagedDatabaseSecurityAlertPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets a list of managed database's security alert policies.
//
// Generated from API version 2020-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the managed database for which the security alert policies are defined.
//   - options - ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseOptions contains the optional parameters for the ManagedDatabaseSecurityAlertPoliciesClient.NewListByDatabasePager
//     method.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) NewListByDatabasePager(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseOptions) *runtime.Pager[ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse]{
		More: func(page ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse) (ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/securityAlertPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedDatabaseSecurityAlertPoliciesClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse, error) {
	result := ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseSecurityAlertPolicyListResult); err != nil {
		return ManagedDatabaseSecurityAlertPoliciesClientListByDatabaseResponse{}, err
	}
	return result, nil
}
