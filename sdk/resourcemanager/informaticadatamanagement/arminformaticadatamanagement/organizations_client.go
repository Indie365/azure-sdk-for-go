//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package arminformaticadatamanagement

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

// OrganizationsClient contains the methods for the Organizations group.
// Don't use this type directly, use NewOrganizationsClient() instead.
type OrganizationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOrganizationsClient creates a new instance of OrganizationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOrganizationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrganizationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OrganizationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a InformaticaOrganizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - resource - Resource create parameters.
//   - options - OrganizationsClientBeginCreateOrUpdateOptions contains the optional parameters for the OrganizationsClient.BeginCreateOrUpdate
//     method.
func (client *OrganizationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, resource InformaticaOrganizationResource, options *OrganizationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[OrganizationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, organizationName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrganizationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrganizationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a InformaticaOrganizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
func (client *OrganizationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, organizationName string, resource InformaticaOrganizationResource, options *OrganizationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "OrganizationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, organizationName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OrganizationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, resource InformaticaOrganizationResource, options *OrganizationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a InformaticaOrganizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - options - OrganizationsClientBeginDeleteOptions contains the optional parameters for the OrganizationsClient.BeginDelete
//     method.
func (client *OrganizationsClient) BeginDelete(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientBeginDeleteOptions) (*runtime.Poller[OrganizationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, organizationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrganizationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrganizationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a InformaticaOrganizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
func (client *OrganizationsClient) deleteOperation(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "OrganizationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OrganizationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a InformaticaOrganizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - options - OrganizationsClientGetOptions contains the optional parameters for the OrganizationsClient.Get method.
func (client *OrganizationsClient) Get(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientGetOptions) (OrganizationsClientGetResponse, error) {
	var err error
	const operationName = "OrganizationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrganizationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrganizationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OrganizationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrganizationsClient) getHandleResponse(resp *http.Response) (OrganizationsClientGetResponse, error) {
	result := OrganizationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaOrganizationResource); err != nil {
		return OrganizationsClientGetResponse{}, err
	}
	return result, nil
}

// GetAllServerlessRuntimes - Gets all serverless runtime resources in a given informatica organization resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - options - OrganizationsClientGetAllServerlessRuntimesOptions contains the optional parameters for the OrganizationsClient.GetAllServerlessRuntimes
//     method.
func (client *OrganizationsClient) GetAllServerlessRuntimes(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientGetAllServerlessRuntimesOptions) (OrganizationsClientGetAllServerlessRuntimesResponse, error) {
	var err error
	const operationName = "OrganizationsClient.GetAllServerlessRuntimes"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAllServerlessRuntimesCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationsClientGetAllServerlessRuntimesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrganizationsClientGetAllServerlessRuntimesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrganizationsClientGetAllServerlessRuntimesResponse{}, err
	}
	resp, err := client.getAllServerlessRuntimesHandleResponse(httpResp)
	return resp, err
}

// getAllServerlessRuntimesCreateRequest creates the GetAllServerlessRuntimes request.
func (client *OrganizationsClient) getAllServerlessRuntimesCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientGetAllServerlessRuntimesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/getAllServerlessRuntimes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAllServerlessRuntimesHandleResponse handles the GetAllServerlessRuntimes response.
func (client *OrganizationsClient) getAllServerlessRuntimesHandleResponse(resp *http.Response) (OrganizationsClientGetAllServerlessRuntimesResponse, error) {
	result := OrganizationsClientGetAllServerlessRuntimesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaServerlessRuntimeResourceList); err != nil {
		return OrganizationsClientGetAllServerlessRuntimesResponse{}, err
	}
	return result, nil
}

// GetServerlessMetadata - Gets Metadata of the serverless runtime environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - options - OrganizationsClientGetServerlessMetadataOptions contains the optional parameters for the OrganizationsClient.GetServerlessMetadata
//     method.
func (client *OrganizationsClient) GetServerlessMetadata(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientGetServerlessMetadataOptions) (OrganizationsClientGetServerlessMetadataResponse, error) {
	var err error
	const operationName = "OrganizationsClient.GetServerlessMetadata"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getServerlessMetadataCreateRequest(ctx, resourceGroupName, organizationName, options)
	if err != nil {
		return OrganizationsClientGetServerlessMetadataResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrganizationsClientGetServerlessMetadataResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrganizationsClientGetServerlessMetadataResponse{}, err
	}
	resp, err := client.getServerlessMetadataHandleResponse(httpResp)
	return resp, err
}

// getServerlessMetadataCreateRequest creates the GetServerlessMetadata request.
func (client *OrganizationsClient) getServerlessMetadataCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, options *OrganizationsClientGetServerlessMetadataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}/getServerlessMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getServerlessMetadataHandleResponse handles the GetServerlessMetadata response.
func (client *OrganizationsClient) getServerlessMetadataHandleResponse(resp *http.Response) (OrganizationsClientGetServerlessMetadataResponse, error) {
	result := OrganizationsClientGetServerlessMetadataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerlessMetadataResponse); err != nil {
		return OrganizationsClientGetServerlessMetadataResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List InformaticaOrganizationResource resources by resource group
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - OrganizationsClientListByResourceGroupOptions contains the optional parameters for the OrganizationsClient.NewListByResourceGroupPager
//     method.
func (client *OrganizationsClient) NewListByResourceGroupPager(resourceGroupName string, options *OrganizationsClientListByResourceGroupOptions) *runtime.Pager[OrganizationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrganizationsClientListByResourceGroupResponse]{
		More: func(page OrganizationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrganizationsClientListByResourceGroupResponse) (OrganizationsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OrganizationsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return OrganizationsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OrganizationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OrganizationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OrganizationsClient) listByResourceGroupHandleResponse(resp *http.Response) (OrganizationsClientListByResourceGroupResponse, error) {
	result := OrganizationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaOrganizationResourceListResult); err != nil {
		return OrganizationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List InformaticaOrganizationResource resources by subscription ID
//
// Generated from API version 2024-05-08
//   - options - OrganizationsClientListBySubscriptionOptions contains the optional parameters for the OrganizationsClient.NewListBySubscriptionPager
//     method.
func (client *OrganizationsClient) NewListBySubscriptionPager(options *OrganizationsClientListBySubscriptionOptions) *runtime.Pager[OrganizationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrganizationsClientListBySubscriptionResponse]{
		More: func(page OrganizationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrganizationsClientListBySubscriptionResponse) (OrganizationsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OrganizationsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return OrganizationsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OrganizationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *OrganizationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Informatica.DataManagement/organizations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *OrganizationsClient) listBySubscriptionHandleResponse(resp *http.Response) (OrganizationsClientListBySubscriptionResponse, error) {
	result := OrganizationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaOrganizationResourceListResult); err != nil {
		return OrganizationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a InformaticaOrganizationResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-08
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - organizationName - Name of the Organizations resource
//   - properties - The resource properties to be updated.
//   - options - OrganizationsClientUpdateOptions contains the optional parameters for the OrganizationsClient.Update method.
func (client *OrganizationsClient) Update(ctx context.Context, resourceGroupName string, organizationName string, properties InformaticaOrganizationResourceUpdate, options *OrganizationsClientUpdateOptions) (OrganizationsClientUpdateResponse, error) {
	var err error
	const operationName = "OrganizationsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, organizationName, properties, options)
	if err != nil {
		return OrganizationsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrganizationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrganizationsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *OrganizationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, organizationName string, properties InformaticaOrganizationResourceUpdate, options *OrganizationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Informatica.DataManagement/organizations/{organizationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if organizationName == "" {
		return nil, errors.New("parameter organizationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{organizationName}", url.PathEscape(organizationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-08")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *OrganizationsClient) updateHandleResponse(resp *http.Response) (OrganizationsClientUpdateResponse, error) {
	result := OrganizationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InformaticaOrganizationResource); err != nil {
		return OrganizationsClientUpdateResponse{}, err
	}
	return result, nil
}
