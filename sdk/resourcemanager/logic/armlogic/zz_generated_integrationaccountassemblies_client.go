//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationAccountAssembliesClient contains the methods for the IntegrationAccountAssemblies group.
// Don't use this type directly, use NewIntegrationAccountAssembliesClient() instead.
type IntegrationAccountAssembliesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationAccountAssembliesClient creates a new instance of IntegrationAccountAssembliesClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationAccountAssembliesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationAccountAssembliesClient, error) {
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
	client := &IntegrationAccountAssembliesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update an assembly for an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// assemblyArtifactName - The assembly artifact name.
// assemblyArtifact - The assembly artifact.
// options - IntegrationAccountAssembliesClientCreateOrUpdateOptions contains the optional parameters for the IntegrationAccountAssembliesClient.CreateOrUpdate
// method.
func (client *IntegrationAccountAssembliesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, assemblyArtifact AssemblyDefinition, options *IntegrationAccountAssembliesClientCreateOrUpdateOptions) (IntegrationAccountAssembliesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, assemblyArtifact, options)
	if err != nil {
		return IntegrationAccountAssembliesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return IntegrationAccountAssembliesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationAccountAssembliesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, assemblyArtifact AssemblyDefinition, options *IntegrationAccountAssembliesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, assemblyArtifact)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IntegrationAccountAssembliesClient) createOrUpdateHandleResponse(resp *http.Response) (IntegrationAccountAssembliesClientCreateOrUpdateResponse, error) {
	result := IntegrationAccountAssembliesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssemblyDefinition); err != nil {
		return IntegrationAccountAssembliesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete an assembly for an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// assemblyArtifactName - The assembly artifact name.
// options - IntegrationAccountAssembliesClientDeleteOptions contains the optional parameters for the IntegrationAccountAssembliesClient.Delete
// method.
func (client *IntegrationAccountAssembliesClient) Delete(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesClientDeleteOptions) (IntegrationAccountAssembliesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, options)
	if err != nil {
		return IntegrationAccountAssembliesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationAccountAssembliesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return IntegrationAccountAssembliesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationAccountAssembliesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an assembly for an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// assemblyArtifactName - The assembly artifact name.
// options - IntegrationAccountAssembliesClientGetOptions contains the optional parameters for the IntegrationAccountAssembliesClient.Get
// method.
func (client *IntegrationAccountAssembliesClient) Get(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesClientGetOptions) (IntegrationAccountAssembliesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, options)
	if err != nil {
		return IntegrationAccountAssembliesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAssembliesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationAccountAssembliesClient) getCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationAccountAssembliesClient) getHandleResponse(resp *http.Response) (IntegrationAccountAssembliesClientGetResponse, error) {
	result := IntegrationAccountAssembliesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssemblyDefinition); err != nil {
		return IntegrationAccountAssembliesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the assemblies for an integration account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// options - IntegrationAccountAssembliesClientListOptions contains the optional parameters for the IntegrationAccountAssembliesClient.List
// method.
func (client *IntegrationAccountAssembliesClient) NewListPager(resourceGroupName string, integrationAccountName string, options *IntegrationAccountAssembliesClientListOptions) *runtime.Pager[IntegrationAccountAssembliesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationAccountAssembliesClientListResponse]{
		More: func(page IntegrationAccountAssembliesClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *IntegrationAccountAssembliesClientListResponse) (IntegrationAccountAssembliesClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, integrationAccountName, options)
			if err != nil {
				return IntegrationAccountAssembliesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return IntegrationAccountAssembliesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return IntegrationAccountAssembliesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationAccountAssembliesClient) listCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, options *IntegrationAccountAssembliesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationAccountAssembliesClient) listHandleResponse(resp *http.Response) (IntegrationAccountAssembliesClientListResponse, error) {
	result := IntegrationAccountAssembliesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssemblyCollection); err != nil {
		return IntegrationAccountAssembliesClientListResponse{}, err
	}
	return result, nil
}

// ListContentCallbackURL - Get the content callback url for an integration account assembly.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-05-01
// resourceGroupName - The resource group name.
// integrationAccountName - The integration account name.
// assemblyArtifactName - The assembly artifact name.
// options - IntegrationAccountAssembliesClientListContentCallbackURLOptions contains the optional parameters for the IntegrationAccountAssembliesClient.ListContentCallbackURL
// method.
func (client *IntegrationAccountAssembliesClient) ListContentCallbackURL(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesClientListContentCallbackURLOptions) (IntegrationAccountAssembliesClientListContentCallbackURLResponse, error) {
	req, err := client.listContentCallbackURLCreateRequest(ctx, resourceGroupName, integrationAccountName, assemblyArtifactName, options)
	if err != nil {
		return IntegrationAccountAssembliesClientListContentCallbackURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationAccountAssembliesClientListContentCallbackURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationAccountAssembliesClientListContentCallbackURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.listContentCallbackURLHandleResponse(resp)
}

// listContentCallbackURLCreateRequest creates the ListContentCallbackURL request.
func (client *IntegrationAccountAssembliesClient) listContentCallbackURLCreateRequest(ctx context.Context, resourceGroupName string, integrationAccountName string, assemblyArtifactName string, options *IntegrationAccountAssembliesClientListContentCallbackURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/integrationAccounts/{integrationAccountName}/assemblies/{assemblyArtifactName}/listContentCallbackUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if integrationAccountName == "" {
		return nil, errors.New("parameter integrationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationAccountName}", url.PathEscape(integrationAccountName))
	if assemblyArtifactName == "" {
		return nil, errors.New("parameter assemblyArtifactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assemblyArtifactName}", url.PathEscape(assemblyArtifactName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listContentCallbackURLHandleResponse handles the ListContentCallbackURL response.
func (client *IntegrationAccountAssembliesClient) listContentCallbackURLHandleResponse(resp *http.Response) (IntegrationAccountAssembliesClientListContentCallbackURLResponse, error) {
	result := IntegrationAccountAssembliesClientListContentCallbackURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowTriggerCallbackURL); err != nil {
		return IntegrationAccountAssembliesClientListContentCallbackURLResponse{}, err
	}
	return result, nil
}