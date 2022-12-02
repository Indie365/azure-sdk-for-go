//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// ComponentsClient contains the methods for the Components group.
// Don't use this type directly, use NewComponentsClient() instead.
type ComponentsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewComponentsClient creates a new instance of ComponentsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewComponentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ComponentsClient, error) {
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
	client := &ComponentsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates (or updates) an Application Insights component. Note: You cannot specify a different value for
// InstrumentationKey nor AppId in the Put operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// insightProperties - Properties that need to be specified to create an Application Insights component.
// options - ComponentsClientCreateOrUpdateOptions contains the optional parameters for the ComponentsClient.CreateOrUpdate
// method.
func (client *ComponentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, insightProperties Component, options *ComponentsClientCreateOrUpdateOptions) (ComponentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, insightProperties, options)
	if err != nil {
		return ComponentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComponentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ComponentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, insightProperties Component, options *ComponentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, insightProperties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ComponentsClient) createOrUpdateHandleResponse(resp *http.Response) (ComponentsClientCreateOrUpdateResponse, error) {
	result := ComponentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Component); err != nil {
		return ComponentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - ComponentsClientDeleteOptions contains the optional parameters for the ComponentsClient.Delete method.
func (client *ComponentsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentsClientDeleteOptions) (ComponentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ComponentsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ComponentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ComponentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ComponentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - ComponentsClientGetOptions contains the optional parameters for the ComponentsClient.Get method.
func (client *ComponentsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentsClientGetOptions) (ComponentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ComponentsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComponentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ComponentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComponentsClient) getHandleResponse(resp *http.Response) (ComponentsClientGetResponse, error) {
	result := ComponentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Component); err != nil {
		return ComponentsClientGetResponse{}, err
	}
	return result, nil
}

// GetPurgeStatus - Get status for an ongoing purge operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// purgeID - In a purge status request, this is the Id of the operation the status of which is returned.
// options - ComponentsClientGetPurgeStatusOptions contains the optional parameters for the ComponentsClient.GetPurgeStatus
// method.
func (client *ComponentsClient) GetPurgeStatus(ctx context.Context, resourceGroupName string, resourceName string, purgeID string, options *ComponentsClientGetPurgeStatusOptions) (ComponentsClientGetPurgeStatusResponse, error) {
	req, err := client.getPurgeStatusCreateRequest(ctx, resourceGroupName, resourceName, purgeID, options)
	if err != nil {
		return ComponentsClientGetPurgeStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentsClientGetPurgeStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComponentsClientGetPurgeStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPurgeStatusHandleResponse(resp)
}

// getPurgeStatusCreateRequest creates the GetPurgeStatus request.
func (client *ComponentsClient) getPurgeStatusCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, purgeID string, options *ComponentsClientGetPurgeStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/operations/{purgeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if purgeID == "" {
		return nil, errors.New("parameter purgeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{purgeId}", url.PathEscape(purgeID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPurgeStatusHandleResponse handles the GetPurgeStatus response.
func (client *ComponentsClient) getPurgeStatusHandleResponse(resp *http.Response) (ComponentsClientGetPurgeStatusResponse, error) {
	result := ComponentsClientGetPurgeStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPurgeStatusResponse); err != nil {
		return ComponentsClientGetPurgeStatusResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of all Application Insights components within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// options - ComponentsClientListOptions contains the optional parameters for the ComponentsClient.List method.
func (client *ComponentsClient) NewListPager(options *ComponentsClientListOptions) *runtime.Pager[ComponentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ComponentsClientListResponse]{
		More: func(page ComponentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ComponentsClientListResponse) (ComponentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ComponentsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ComponentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ComponentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ComponentsClient) listCreateRequest(ctx context.Context, options *ComponentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/components"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ComponentsClient) listHandleResponse(resp *http.Response) (ComponentsClientListResponse, error) {
	result := ComponentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentListResult); err != nil {
		return ComponentsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of Application Insights components within a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ComponentsClientListByResourceGroupOptions contains the optional parameters for the ComponentsClient.ListByResourceGroup
// method.
func (client *ComponentsClient) NewListByResourceGroupPager(resourceGroupName string, options *ComponentsClientListByResourceGroupOptions) *runtime.Pager[ComponentsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ComponentsClientListByResourceGroupResponse]{
		More: func(page ComponentsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ComponentsClientListByResourceGroupResponse) (ComponentsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ComponentsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ComponentsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ComponentsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ComponentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ComponentsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ComponentsClient) listByResourceGroupHandleResponse(resp *http.Response) (ComponentsClientListByResourceGroupResponse, error) {
	result := ComponentsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentListResult); err != nil {
		return ComponentsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Purge - Purges data in an Application Insights component by a set of user-defined filters.
// In order to manage system resources, purge requests are throttled at 50 requests per hour. You should batch the execution
// of purge requests by sending a single command whose predicate includes all
// user identities that require purging. Use the in operator to specify multiple identities. You should run the query prior
// to using for a purge request to verify that the results are expected.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// body - Describes the body of a request to purge data in a single table of an Application Insights component
// options - ComponentsClientPurgeOptions contains the optional parameters for the ComponentsClient.Purge method.
func (client *ComponentsClient) Purge(ctx context.Context, resourceGroupName string, resourceName string, body ComponentPurgeBody, options *ComponentsClientPurgeOptions) (ComponentsClientPurgeResponse, error) {
	req, err := client.purgeCreateRequest(ctx, resourceGroupName, resourceName, body, options)
	if err != nil {
		return ComponentsClientPurgeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentsClientPurgeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return ComponentsClientPurgeResponse{}, runtime.NewResponseError(resp)
	}
	return client.purgeHandleResponse(resp)
}

// purgeCreateRequest creates the Purge request.
func (client *ComponentsClient) purgeCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, body ComponentPurgeBody, options *ComponentsClientPurgeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/purge"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// purgeHandleResponse handles the Purge response.
func (client *ComponentsClient) purgeHandleResponse(resp *http.Response) (ComponentsClientPurgeResponse, error) {
	result := ComponentsClientPurgeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPurgeResponse); err != nil {
		return ComponentsClientPurgeResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates an existing component's tags. To update other fields use the CreateOrUpdate method.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-02-02
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// componentTags - Updated tag information to set into the component instance.
// options - ComponentsClientUpdateTagsOptions contains the optional parameters for the ComponentsClient.UpdateTags method.
func (client *ComponentsClient) UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, componentTags TagsResource, options *ComponentsClientUpdateTagsOptions) (ComponentsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, resourceName, componentTags, options)
	if err != nil {
		return ComponentsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComponentsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ComponentsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, componentTags TagsResource, options *ComponentsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-02-02")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, componentTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ComponentsClient) updateTagsHandleResponse(resp *http.Response) (ComponentsClientUpdateTagsResponse, error) {
	result := ComponentsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Component); err != nil {
		return ComponentsClientUpdateTagsResponse{}, err
	}
	return result, nil
}