//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationsmanagement

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

// ManagementAssociationsClient contains the methods for the ManagementAssociations group.
// Don't use this type directly, use NewManagementAssociationsClient() instead.
type ManagementAssociationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewManagementAssociationsClient creates a new instance of ManagementAssociationsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewManagementAssociationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementAssociationsClient, error) {
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
	client := &ManagementAssociationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the ManagementAssociation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-11-01-preview
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// providerName - Provider name for the parent resource.
// resourceType - Resource type for the parent resource
// resourceName - Parent resource name.
// managementAssociationName - User ManagementAssociation Name.
// parameters - The parameters required to create ManagementAssociation extension.
// options - ManagementAssociationsClientCreateOrUpdateOptions contains the optional parameters for the ManagementAssociationsClient.CreateOrUpdate
// method.
func (client *ManagementAssociationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, parameters ManagementAssociation, options *ManagementAssociationsClientCreateOrUpdateOptions) (ManagementAssociationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, managementAssociationName, parameters, options)
	if err != nil {
		return ManagementAssociationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagementAssociationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, parameters ManagementAssociation, options *ManagementAssociationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"
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
	if managementAssociationName == "" {
		return nil, errors.New("parameter managementAssociationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementAssociationName}", url.PathEscape(managementAssociationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagementAssociationsClient) createOrUpdateHandleResponse(resp *http.Response) (ManagementAssociationsClientCreateOrUpdateResponse, error) {
	result := ManagementAssociationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementAssociation); err != nil {
		return ManagementAssociationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the ManagementAssociation in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-11-01-preview
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// providerName - Provider name for the parent resource.
// resourceType - Resource type for the parent resource
// resourceName - Parent resource name.
// managementAssociationName - User ManagementAssociation Name.
// options - ManagementAssociationsClientDeleteOptions contains the optional parameters for the ManagementAssociationsClient.Delete
// method.
func (client *ManagementAssociationsClient) Delete(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, options *ManagementAssociationsClientDeleteOptions) (ManagementAssociationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, managementAssociationName, options)
	if err != nil {
		return ManagementAssociationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ManagementAssociationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagementAssociationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, options *ManagementAssociationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"
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
	if managementAssociationName == "" {
		return nil, errors.New("parameter managementAssociationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementAssociationName}", url.PathEscape(managementAssociationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the user ManagementAssociation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-11-01-preview
// resourceGroupName - The name of the resource group to get. The name is case insensitive.
// providerName - Provider name for the parent resource.
// resourceType - Resource type for the parent resource
// resourceName - Parent resource name.
// managementAssociationName - User ManagementAssociation Name.
// options - ManagementAssociationsClientGetOptions contains the optional parameters for the ManagementAssociationsClient.Get
// method.
func (client *ManagementAssociationsClient) Get(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, options *ManagementAssociationsClientGetOptions) (ManagementAssociationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, managementAssociationName, options)
	if err != nil {
		return ManagementAssociationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagementAssociationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string, options *ManagementAssociationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.OperationsManagement/ManagementAssociations/{managementAssociationName}"
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
	if managementAssociationName == "" {
		return nil, errors.New("parameter managementAssociationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementAssociationName}", url.PathEscape(managementAssociationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagementAssociationsClient) getHandleResponse(resp *http.Response) (ManagementAssociationsClientGetResponse, error) {
	result := ManagementAssociationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementAssociation); err != nil {
		return ManagementAssociationsClientGetResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieves the ManagementAssociations list.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2015-11-01-preview
// options - ManagementAssociationsClientListBySubscriptionOptions contains the optional parameters for the ManagementAssociationsClient.ListBySubscription
// method.
func (client *ManagementAssociationsClient) ListBySubscription(ctx context.Context, options *ManagementAssociationsClientListBySubscriptionOptions) (ManagementAssociationsClientListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ManagementAssociationsClientListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagementAssociationsClientListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementAssociationsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ManagementAssociationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ManagementAssociationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationsManagement/ManagementAssociations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ManagementAssociationsClient) listBySubscriptionHandleResponse(resp *http.Response) (ManagementAssociationsClientListBySubscriptionResponse, error) {
	result := ManagementAssociationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagementAssociationPropertiesList); err != nil {
		return ManagementAssociationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
