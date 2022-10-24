//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/profiles/hybrid20200901/internal"
	"net/http"
	"net/url"
	"strings"
)

// ObjectReplicationPoliciesClient contains the methods for the ObjectReplicationPolicies group.
// Don't use this type directly, use NewObjectReplicationPoliciesClient() instead.
type ObjectReplicationPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewObjectReplicationPoliciesClient creates a new instance of ObjectReplicationPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewObjectReplicationPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ObjectReplicationPoliciesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(hybrid20200901.ModuleName, hybrid20200901.ModuleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ObjectReplicationPoliciesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update the object replication policy of the storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// objectReplicationPolicyID - The ID of object replication policy or 'default' if the policy ID is unknown.
// properties - The object replication policy set to a storage account. A unique policy ID will be created if absent.
// options - ObjectReplicationPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ObjectReplicationPoliciesClient.CreateOrUpdate
// method.
func (client *ObjectReplicationPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, properties ObjectReplicationPolicy, options *ObjectReplicationPoliciesClientCreateOrUpdateOptions) (ObjectReplicationPoliciesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, properties, options)
	if err != nil {
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ObjectReplicationPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, properties ObjectReplicationPolicy, options *ObjectReplicationPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ObjectReplicationPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ObjectReplicationPoliciesClientCreateOrUpdateResponse, error) {
	result := ObjectReplicationPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicy); err != nil {
		return ObjectReplicationPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the object replication policy associated with the specified storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// objectReplicationPolicyID - The ID of object replication policy or 'default' if the policy ID is unknown.
// options - ObjectReplicationPoliciesClientDeleteOptions contains the optional parameters for the ObjectReplicationPoliciesClient.Delete
// method.
func (client *ObjectReplicationPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesClientDeleteOptions) (ObjectReplicationPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, options)
	if err != nil {
		return ObjectReplicationPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectReplicationPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ObjectReplicationPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ObjectReplicationPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ObjectReplicationPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the object replication policy of the storage account by policy ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// objectReplicationPolicyID - The ID of object replication policy or 'default' if the policy ID is unknown.
// options - ObjectReplicationPoliciesClientGetOptions contains the optional parameters for the ObjectReplicationPoliciesClient.Get
// method.
func (client *ObjectReplicationPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesClientGetOptions) (ObjectReplicationPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, objectReplicationPolicyID, options)
	if err != nil {
		return ObjectReplicationPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ObjectReplicationPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ObjectReplicationPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ObjectReplicationPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, objectReplicationPolicyID string, options *ObjectReplicationPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies/{objectReplicationPolicyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if objectReplicationPolicyID == "" {
		return nil, errors.New("parameter objectReplicationPolicyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{objectReplicationPolicyId}", url.PathEscape(objectReplicationPolicyID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ObjectReplicationPoliciesClient) getHandleResponse(resp *http.Response) (ObjectReplicationPoliciesClientGetResponse, error) {
	result := ObjectReplicationPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicy); err != nil {
		return ObjectReplicationPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the object replication policies associated with the storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// accountName - The name of the storage account within the specified resource group. Storage account names must be between
// 3 and 24 characters in length and use numbers and lower-case letters only.
// options - ObjectReplicationPoliciesClientListOptions contains the optional parameters for the ObjectReplicationPoliciesClient.List
// method.
func (client *ObjectReplicationPoliciesClient) NewListPager(resourceGroupName string, accountName string, options *ObjectReplicationPoliciesClientListOptions) *runtime.Pager[ObjectReplicationPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ObjectReplicationPoliciesClientListResponse]{
		More: func(page ObjectReplicationPoliciesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ObjectReplicationPoliciesClientListResponse) (ObjectReplicationPoliciesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return ObjectReplicationPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ObjectReplicationPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ObjectReplicationPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ObjectReplicationPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *ObjectReplicationPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/objectReplicationPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ObjectReplicationPoliciesClient) listHandleResponse(resp *http.Response) (ObjectReplicationPoliciesClientListResponse, error) {
	result := ObjectReplicationPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ObjectReplicationPolicies); err != nil {
		return ObjectReplicationPoliciesClientListResponse{}, err
	}
	return result, nil
}
