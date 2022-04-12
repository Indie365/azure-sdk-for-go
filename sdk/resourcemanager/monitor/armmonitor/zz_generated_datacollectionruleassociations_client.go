//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// DataCollectionRuleAssociationsClient contains the methods for the DataCollectionRuleAssociations group.
// Don't use this type directly, use NewDataCollectionRuleAssociationsClient() instead.
type DataCollectionRuleAssociationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataCollectionRuleAssociationsClient creates a new instance of DataCollectionRuleAssociationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataCollectionRuleAssociationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataCollectionRuleAssociationsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DataCollectionRuleAssociationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Creates or updates an association.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The identifier of the resource.
// associationName - The name of the association. The name is case insensitive.
// options - DataCollectionRuleAssociationsClientCreateOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.Create
// method.
func (client *DataCollectionRuleAssociationsClient) Create(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientCreateOptions) (DataCollectionRuleAssociationsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceURI, associationName, options)
	if err != nil {
		return DataCollectionRuleAssociationsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionRuleAssociationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataCollectionRuleAssociationsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *DataCollectionRuleAssociationsClient) createCreateRequest(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/dataCollectionRuleAssociations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *DataCollectionRuleAssociationsClient) createHandleResponse(resp *http.Response) (DataCollectionRuleAssociationsClientCreateResponse, error) {
	result := DataCollectionRuleAssociationsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleAssociationProxyOnlyResource); err != nil {
		return DataCollectionRuleAssociationsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an association.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The identifier of the resource.
// associationName - The name of the association. The name is case insensitive.
// options - DataCollectionRuleAssociationsClientDeleteOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.Delete
// method.
func (client *DataCollectionRuleAssociationsClient) Delete(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientDeleteOptions) (DataCollectionRuleAssociationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceURI, associationName, options)
	if err != nil {
		return DataCollectionRuleAssociationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionRuleAssociationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataCollectionRuleAssociationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataCollectionRuleAssociationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataCollectionRuleAssociationsClient) deleteCreateRequest(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/dataCollectionRuleAssociations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the specified association.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The identifier of the resource.
// associationName - The name of the association. The name is case insensitive.
// options - DataCollectionRuleAssociationsClientGetOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.Get
// method.
func (client *DataCollectionRuleAssociationsClient) Get(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientGetOptions) (DataCollectionRuleAssociationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, associationName, options)
	if err != nil {
		return DataCollectionRuleAssociationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionRuleAssociationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataCollectionRuleAssociationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataCollectionRuleAssociationsClient) getCreateRequest(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/dataCollectionRuleAssociations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataCollectionRuleAssociationsClient) getHandleResponse(resp *http.Response) (DataCollectionRuleAssociationsClientGetResponse, error) {
	result := DataCollectionRuleAssociationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleAssociationProxyOnlyResource); err != nil {
		return DataCollectionRuleAssociationsClientGetResponse{}, err
	}
	return result, nil
}

// ListByDataCollectionEndpoint - Lists associations for the specified data collection endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dataCollectionEndpointName - The name of the data collection endpoint. The name is case insensitive.
// options - DataCollectionRuleAssociationsClientListByDataCollectionEndpointOptions contains the optional parameters for
// the DataCollectionRuleAssociationsClient.ListByDataCollectionEndpoint method.
func (client *DataCollectionRuleAssociationsClient) ListByDataCollectionEndpoint(resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionRuleAssociationsClientListByDataCollectionEndpointOptions) *runtime.Pager[DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse] {
	return runtime.NewPager(runtime.PageProcessor[DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse]{
		More: func(page DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse) (DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataCollectionEndpointCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataCollectionEndpointHandleResponse(resp)
		},
	})
}

// listByDataCollectionEndpointCreateRequest creates the ListByDataCollectionEndpoint request.
func (client *DataCollectionRuleAssociationsClient) listByDataCollectionEndpointCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionRuleAssociationsClientListByDataCollectionEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}/associations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByDataCollectionEndpointHandleResponse handles the ListByDataCollectionEndpoint response.
func (client *DataCollectionRuleAssociationsClient) listByDataCollectionEndpointHandleResponse(resp *http.Response) (DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse, error) {
	result := DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleAssociationProxyOnlyResourceListResult); err != nil {
		return DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse{}, err
	}
	return result, nil
}

// ListByResource - Lists associations for the specified resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceURI - The identifier of the resource.
// options - DataCollectionRuleAssociationsClientListByResourceOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.ListByResource
// method.
func (client *DataCollectionRuleAssociationsClient) ListByResource(resourceURI string, options *DataCollectionRuleAssociationsClientListByResourceOptions) *runtime.Pager[DataCollectionRuleAssociationsClientListByResourceResponse] {
	return runtime.NewPager(runtime.PageProcessor[DataCollectionRuleAssociationsClientListByResourceResponse]{
		More: func(page DataCollectionRuleAssociationsClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRuleAssociationsClientListByResourceResponse) (DataCollectionRuleAssociationsClientListByResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceCreateRequest(ctx, resourceURI, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataCollectionRuleAssociationsClientListByResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataCollectionRuleAssociationsClientListByResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataCollectionRuleAssociationsClientListByResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceHandleResponse(resp)
		},
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *DataCollectionRuleAssociationsClient) listByResourceCreateRequest(ctx context.Context, resourceURI string, options *DataCollectionRuleAssociationsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/dataCollectionRuleAssociations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *DataCollectionRuleAssociationsClient) listByResourceHandleResponse(resp *http.Response) (DataCollectionRuleAssociationsClientListByResourceResponse, error) {
	result := DataCollectionRuleAssociationsClientListByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleAssociationProxyOnlyResourceListResult); err != nil {
		return DataCollectionRuleAssociationsClientListByResourceResponse{}, err
	}
	return result, nil
}

// ListByRule - Lists associations for the specified data collection rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// dataCollectionRuleName - The name of the data collection rule. The name is case insensitive.
// options - DataCollectionRuleAssociationsClientListByRuleOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.ListByRule
// method.
func (client *DataCollectionRuleAssociationsClient) ListByRule(resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRuleAssociationsClientListByRuleOptions) *runtime.Pager[DataCollectionRuleAssociationsClientListByRuleResponse] {
	return runtime.NewPager(runtime.PageProcessor[DataCollectionRuleAssociationsClientListByRuleResponse]{
		More: func(page DataCollectionRuleAssociationsClientListByRuleResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRuleAssociationsClientListByRuleResponse) (DataCollectionRuleAssociationsClientListByRuleResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRuleCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataCollectionRuleAssociationsClientListByRuleResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataCollectionRuleAssociationsClientListByRuleResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataCollectionRuleAssociationsClientListByRuleResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRuleHandleResponse(resp)
		},
	})
}

// listByRuleCreateRequest creates the ListByRule request.
func (client *DataCollectionRuleAssociationsClient) listByRuleCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRuleAssociationsClientListByRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionRules/{dataCollectionRuleName}/associations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionRuleName == "" {
		return nil, errors.New("parameter dataCollectionRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionRuleName}", url.PathEscape(dataCollectionRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByRuleHandleResponse handles the ListByRule response.
func (client *DataCollectionRuleAssociationsClient) listByRuleHandleResponse(resp *http.Response) (DataCollectionRuleAssociationsClientListByRuleResponse, error) {
	result := DataCollectionRuleAssociationsClientListByRuleResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionRuleAssociationProxyOnlyResourceListResult); err != nil {
		return DataCollectionRuleAssociationsClientListByRuleResponse{}, err
	}
	return result, nil
}
