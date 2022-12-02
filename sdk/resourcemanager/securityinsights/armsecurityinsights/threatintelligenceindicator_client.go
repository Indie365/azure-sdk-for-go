//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights

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

// ThreatIntelligenceIndicatorClient contains the methods for the ThreatIntelligenceIndicator group.
// Don't use this type directly, use NewThreatIntelligenceIndicatorClient() instead.
type ThreatIntelligenceIndicatorClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewThreatIntelligenceIndicatorClient creates a new instance of ThreatIntelligenceIndicatorClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewThreatIntelligenceIndicatorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ThreatIntelligenceIndicatorClient, error) {
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
	client := &ThreatIntelligenceIndicatorClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// AppendTags - Append tags to a threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// name - Threat intelligence indicator name field.
// threatIntelligenceAppendTags - The threat intelligence append tags request body
// options - ThreatIntelligenceIndicatorClientAppendTagsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.AppendTags
// method.
func (client *ThreatIntelligenceIndicatorClient) AppendTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags ThreatIntelligenceAppendTags, options *ThreatIntelligenceIndicatorClientAppendTagsOptions) (ThreatIntelligenceIndicatorClientAppendTagsResponse, error) {
	req, err := client.appendTagsCreateRequest(ctx, resourceGroupName, workspaceName, name, threatIntelligenceAppendTags, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, runtime.NewResponseError(resp)
	}
	return ThreatIntelligenceIndicatorClientAppendTagsResponse{}, nil
}

// appendTagsCreateRequest creates the AppendTags request.
func (client *ThreatIntelligenceIndicatorClient) appendTagsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceAppendTags ThreatIntelligenceAppendTags, options *ThreatIntelligenceIndicatorClientAppendTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}/appendTags"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, threatIntelligenceAppendTags)
}

// Create - Update a threat Intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// name - Threat intelligence indicator name field.
// threatIntelligenceProperties - Properties of threat intelligence indicators to create and update.
// options - ThreatIntelligenceIndicatorClientCreateOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Create
// method.
func (client *ThreatIntelligenceIndicatorClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateOptions) (ThreatIntelligenceIndicatorClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, name, threatIntelligenceProperties, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ThreatIntelligenceIndicatorClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *ThreatIntelligenceIndicatorClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, threatIntelligenceProperties)
}

// createHandleResponse handles the Create response.
func (client *ThreatIntelligenceIndicatorClient) createHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientCreateResponse, error) {
	result := ThreatIntelligenceIndicatorClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientCreateResponse{}, err
	}
	return result, nil
}

// CreateIndicator - Create a new threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// threatIntelligenceProperties - Properties of threat intelligence indicators to create and update.
// options - ThreatIntelligenceIndicatorClientCreateIndicatorOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.CreateIndicator
// method.
func (client *ThreatIntelligenceIndicatorClient) CreateIndicator(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateIndicatorOptions) (ThreatIntelligenceIndicatorClientCreateIndicatorResponse, error) {
	req, err := client.createIndicatorCreateRequest(ctx, resourceGroupName, workspaceName, threatIntelligenceProperties, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, runtime.NewResponseError(resp)
	}
	return client.createIndicatorHandleResponse(resp)
}

// createIndicatorCreateRequest creates the CreateIndicator request.
func (client *ThreatIntelligenceIndicatorClient) createIndicatorCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceProperties ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientCreateIndicatorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/createIndicator"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, threatIntelligenceProperties)
}

// createIndicatorHandleResponse handles the CreateIndicator response.
func (client *ThreatIntelligenceIndicatorClient) createIndicatorHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientCreateIndicatorResponse, error) {
	result := ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientCreateIndicatorResponse{}, err
	}
	return result, nil
}

// Delete - Delete a threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// name - Threat intelligence indicator name field.
// options - ThreatIntelligenceIndicatorClientDeleteOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Delete
// method.
func (client *ThreatIntelligenceIndicatorClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientDeleteOptions) (ThreatIntelligenceIndicatorClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ThreatIntelligenceIndicatorClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ThreatIntelligenceIndicatorClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ThreatIntelligenceIndicatorClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - View a threat intelligence indicator by name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// name - Threat intelligence indicator name field.
// options - ThreatIntelligenceIndicatorClientGetOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.Get
// method.
func (client *ThreatIntelligenceIndicatorClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientGetOptions) (ThreatIntelligenceIndicatorClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ThreatIntelligenceIndicatorClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ThreatIntelligenceIndicatorClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *ThreatIntelligenceIndicatorClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ThreatIntelligenceIndicatorClient) getHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientGetResponse, error) {
	result := ThreatIntelligenceIndicatorClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientGetResponse{}, err
	}
	return result, nil
}

// NewQueryIndicatorsPager - Query threat intelligence indicators as per filtering criteria.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// threatIntelligenceFilteringCriteria - Filtering criteria for querying threat intelligence indicators.
// options - ThreatIntelligenceIndicatorClientQueryIndicatorsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.QueryIndicators
// method.
func (client *ThreatIntelligenceIndicatorClient) NewQueryIndicatorsPager(resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria ThreatIntelligenceFilteringCriteria, options *ThreatIntelligenceIndicatorClientQueryIndicatorsOptions) *runtime.Pager[ThreatIntelligenceIndicatorClientQueryIndicatorsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ThreatIntelligenceIndicatorClientQueryIndicatorsResponse]{
		More: func(page ThreatIntelligenceIndicatorClientQueryIndicatorsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ThreatIntelligenceIndicatorClientQueryIndicatorsResponse) (ThreatIntelligenceIndicatorClientQueryIndicatorsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.queryIndicatorsCreateRequest(ctx, resourceGroupName, workspaceName, threatIntelligenceFilteringCriteria, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}, runtime.NewResponseError(resp)
			}
			return client.queryIndicatorsHandleResponse(resp)
		},
	})
}

// queryIndicatorsCreateRequest creates the QueryIndicators request.
func (client *ThreatIntelligenceIndicatorClient) queryIndicatorsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, threatIntelligenceFilteringCriteria ThreatIntelligenceFilteringCriteria, options *ThreatIntelligenceIndicatorClientQueryIndicatorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/queryIndicators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, threatIntelligenceFilteringCriteria)
}

// queryIndicatorsHandleResponse handles the QueryIndicators response.
func (client *ThreatIntelligenceIndicatorClient) queryIndicatorsHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientQueryIndicatorsResponse, error) {
	result := ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThreatIntelligenceInformationList); err != nil {
		return ThreatIntelligenceIndicatorClientQueryIndicatorsResponse{}, err
	}
	return result, nil
}

// ReplaceTags - Replace tags added to a threat intelligence indicator.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// name - Threat intelligence indicator name field.
// threatIntelligenceReplaceTags - Tags in the threat intelligence indicator to be replaced.
// options - ThreatIntelligenceIndicatorClientReplaceTagsOptions contains the optional parameters for the ThreatIntelligenceIndicatorClient.ReplaceTags
// method.
func (client *ThreatIntelligenceIndicatorClient) ReplaceTags(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientReplaceTagsOptions) (ThreatIntelligenceIndicatorClientReplaceTagsResponse, error) {
	req, err := client.replaceTagsCreateRequest(ctx, resourceGroupName, workspaceName, name, threatIntelligenceReplaceTags, options)
	if err != nil {
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.replaceTagsHandleResponse(resp)
}

// replaceTagsCreateRequest creates the ReplaceTags request.
func (client *ThreatIntelligenceIndicatorClient) replaceTagsCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, threatIntelligenceReplaceTags ThreatIntelligenceIndicatorModel, options *ThreatIntelligenceIndicatorClientReplaceTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/{name}/replaceTags"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, threatIntelligenceReplaceTags)
}

// replaceTagsHandleResponse handles the ReplaceTags response.
func (client *ThreatIntelligenceIndicatorClient) replaceTagsHandleResponse(resp *http.Response) (ThreatIntelligenceIndicatorClientReplaceTagsResponse, error) {
	result := ThreatIntelligenceIndicatorClientReplaceTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return ThreatIntelligenceIndicatorClientReplaceTagsResponse{}, err
	}
	return result, nil
}