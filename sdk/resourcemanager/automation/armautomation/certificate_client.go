//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation

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

// CertificateClient contains the methods for the Certificate group.
// Don't use this type directly, use NewCertificateClient() instead.
type CertificateClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCertificateClient creates a new instance of CertificateClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCertificateClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CertificateClient, error) {
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
	client := &CertificateClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create a certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// certificateName - The parameters supplied to the create or update certificate operation.
// parameters - The parameters supplied to the create or update certificate operation.
// options - CertificateClientCreateOrUpdateOptions contains the optional parameters for the CertificateClient.CreateOrUpdate
// method.
func (client *CertificateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateClientCreateOrUpdateOptions) (CertificateClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, parameters, options)
	if err != nil {
		return CertificateClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CertificateClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CertificateClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateCreateOrUpdateParameters, options *CertificateClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CertificateClient) createOrUpdateHandleResponse(resp *http.Response) (CertificateClientCreateOrUpdateResponse, error) {
	result := CertificateClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// certificateName - The name of certificate.
// options - CertificateClientDeleteOptions contains the optional parameters for the CertificateClient.Delete method.
func (client *CertificateClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateClientDeleteOptions) (CertificateClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, options)
	if err != nil {
		return CertificateClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificateClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CertificateClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CertificateClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the certificate identified by certificate name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// certificateName - The name of certificate.
// options - CertificateClientGetOptions contains the optional parameters for the CertificateClient.Get method.
func (client *CertificateClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateClientGetOptions) (CertificateClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, options)
	if err != nil {
		return CertificateClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificateClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CertificateClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, options *CertificateClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CertificateClient) getHandleResponse(resp *http.Response) (CertificateClientGetResponse, error) {
	result := CertificateClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of certificates.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - CertificateClientListByAutomationAccountOptions contains the optional parameters for the CertificateClient.ListByAutomationAccount
// method.
func (client *CertificateClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *CertificateClientListByAutomationAccountOptions) *runtime.Pager[CertificateClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[CertificateClientListByAutomationAccountResponse]{
		More: func(page CertificateClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CertificateClientListByAutomationAccountResponse) (CertificateClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CertificateClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CertificateClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CertificateClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *CertificateClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *CertificateClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *CertificateClient) listByAutomationAccountHandleResponse(resp *http.Response) (CertificateClientListByAutomationAccountResponse, error) {
	result := CertificateClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CertificateListResult); err != nil {
		return CertificateClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// Update - Update a certificate.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// certificateName - The parameters supplied to the update certificate operation.
// parameters - The parameters supplied to the update certificate operation.
// options - CertificateClientUpdateOptions contains the optional parameters for the CertificateClient.Update method.
func (client *CertificateClient) Update(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateUpdateParameters, options *CertificateClientUpdateOptions) (CertificateClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, automationAccountName, certificateName, parameters, options)
	if err != nil {
		return CertificateClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CertificateClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CertificateClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *CertificateClient) updateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, certificateName string, parameters CertificateUpdateParameters, options *CertificateClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/certificates/{certificateName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if certificateName == "" {
		return nil, errors.New("parameter certificateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{certificateName}", url.PathEscape(certificateName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *CertificateClient) updateHandleResponse(resp *http.Response) (CertificateClientUpdateResponse, error) {
	result := CertificateClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Certificate); err != nil {
		return CertificateClientUpdateResponse{}, err
	}
	return result, nil
}