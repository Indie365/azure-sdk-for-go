//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// DataCollectionRuleAssociationsClient contains the methods for the DataCollectionRuleAssociations group.
// Don't use this type directly, use NewDataCollectionRuleAssociationsClient() instead.
type DataCollectionRuleAssociationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDataCollectionRuleAssociationsClient creates a new instance of DataCollectionRuleAssociationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDataCollectionRuleAssociationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataCollectionRuleAssociationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DataCollectionRuleAssociationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates or updates an association.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceURI - The identifier of the resource.
//   - associationName - The name of the association. The name is case insensitive.
//   - options - DataCollectionRuleAssociationsClientCreateOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.Create
//     method.
func (client *DataCollectionRuleAssociationsClient) Create(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientCreateOptions) (DataCollectionRuleAssociationsClientCreateResponse, error) {
	var err error
	const operationName = "DataCollectionRuleAssociationsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceURI, associationName, options)
	if err != nil {
		return DataCollectionRuleAssociationsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRuleAssociationsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DataCollectionRuleAssociationsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *DataCollectionRuleAssociationsClient) createCreateRequest(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/dataCollectionRuleAssociations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
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
//
// Generated from API version 2022-06-01
//   - resourceURI - The identifier of the resource.
//   - associationName - The name of the association. The name is case insensitive.
//   - options - DataCollectionRuleAssociationsClientDeleteOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.Delete
//     method.
func (client *DataCollectionRuleAssociationsClient) Delete(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientDeleteOptions) (DataCollectionRuleAssociationsClientDeleteResponse, error) {
	var err error
	const operationName = "DataCollectionRuleAssociationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, associationName, options)
	if err != nil {
		return DataCollectionRuleAssociationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRuleAssociationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DataCollectionRuleAssociationsClientDeleteResponse{}, err
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the specified association.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceURI - The identifier of the resource.
//   - associationName - The name of the association. The name is case insensitive.
//   - options - DataCollectionRuleAssociationsClientGetOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.Get
//     method.
func (client *DataCollectionRuleAssociationsClient) Get(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientGetOptions) (DataCollectionRuleAssociationsClientGetResponse, error) {
	var err error
	const operationName = "DataCollectionRuleAssociationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, associationName, options)
	if err != nil {
		return DataCollectionRuleAssociationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DataCollectionRuleAssociationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DataCollectionRuleAssociationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DataCollectionRuleAssociationsClient) getCreateRequest(ctx context.Context, resourceURI string, associationName string, options *DataCollectionRuleAssociationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/dataCollectionRuleAssociations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
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

// NewListByDataCollectionEndpointPager - Lists associations for the specified data collection endpoint.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataCollectionEndpointName - The name of the data collection endpoint. The name is case insensitive.
//   - options - DataCollectionRuleAssociationsClientListByDataCollectionEndpointOptions contains the optional parameters for
//     the DataCollectionRuleAssociationsClient.NewListByDataCollectionEndpointPager method.
func (client *DataCollectionRuleAssociationsClient) NewListByDataCollectionEndpointPager(resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionRuleAssociationsClientListByDataCollectionEndpointOptions) *runtime.Pager[DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse]{
		More: func(page DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse) (DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataCollectionRuleAssociationsClient.NewListByDataCollectionEndpointPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDataCollectionEndpointCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
			}, nil)
			if err != nil {
				return DataCollectionRuleAssociationsClientListByDataCollectionEndpointResponse{}, err
			}
			return client.listByDataCollectionEndpointHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
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

// NewListByResourcePager - Lists associations for the specified resource.
//
// Generated from API version 2022-06-01
//   - resourceURI - The identifier of the resource.
//   - options - DataCollectionRuleAssociationsClientListByResourceOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.NewListByResourcePager
//     method.
func (client *DataCollectionRuleAssociationsClient) NewListByResourcePager(resourceURI string, options *DataCollectionRuleAssociationsClientListByResourceOptions) *runtime.Pager[DataCollectionRuleAssociationsClientListByResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataCollectionRuleAssociationsClientListByResourceResponse]{
		More: func(page DataCollectionRuleAssociationsClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRuleAssociationsClientListByResourceResponse) (DataCollectionRuleAssociationsClientListByResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataCollectionRuleAssociationsClient.NewListByResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceCreateRequest(ctx, resourceURI, options)
			}, nil)
			if err != nil {
				return DataCollectionRuleAssociationsClientListByResourceResponse{}, err
			}
			return client.listByResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *DataCollectionRuleAssociationsClient) listByResourceCreateRequest(ctx context.Context, resourceURI string, options *DataCollectionRuleAssociationsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/dataCollectionRuleAssociations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
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

// NewListByRulePager - Lists associations for the specified data collection rule.
//
// Generated from API version 2022-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - dataCollectionRuleName - The name of the data collection rule. The name is case insensitive.
//   - options - DataCollectionRuleAssociationsClientListByRuleOptions contains the optional parameters for the DataCollectionRuleAssociationsClient.NewListByRulePager
//     method.
func (client *DataCollectionRuleAssociationsClient) NewListByRulePager(resourceGroupName string, dataCollectionRuleName string, options *DataCollectionRuleAssociationsClientListByRuleOptions) *runtime.Pager[DataCollectionRuleAssociationsClientListByRuleResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataCollectionRuleAssociationsClientListByRuleResponse]{
		More: func(page DataCollectionRuleAssociationsClientListByRuleResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataCollectionRuleAssociationsClientListByRuleResponse) (DataCollectionRuleAssociationsClientListByRuleResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DataCollectionRuleAssociationsClient.NewListByRulePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByRuleCreateRequest(ctx, resourceGroupName, dataCollectionRuleName, options)
			}, nil)
			if err != nil {
				return DataCollectionRuleAssociationsClientListByRuleResponse{}, err
			}
			return client.listByRuleHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
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
