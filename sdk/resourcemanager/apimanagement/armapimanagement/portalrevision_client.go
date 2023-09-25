//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PortalRevisionClient contains the methods for the PortalRevision group.
// Don't use this type directly, use NewPortalRevisionClient() instead.
type PortalRevisionClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewPortalRevisionClient creates a new instance of PortalRevisionClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPortalRevisionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PortalRevisionClient, error) {
	cl, err := arm.NewClient(moduleName+".PortalRevisionClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PortalRevisionClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new developer portal's revision by running the portal's publishing. The isCurrent property
// indicates if the revision is publicly accessible.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalRevisionID - Portal revision identifier. Must be unique in the current API Management service instance.
//   - options - PortalRevisionClientBeginCreateOrUpdateOptions contains the optional parameters for the PortalRevisionClient.BeginCreateOrUpdate
//     method.
func (client *PortalRevisionClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, parameters PortalRevisionContract, options *PortalRevisionClientBeginCreateOrUpdateOptions) (*runtime.Poller[PortalRevisionClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, portalRevisionID, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PortalRevisionClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PortalRevisionClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new developer portal's revision by running the portal's publishing. The isCurrent property indicates
// if the revision is publicly accessible.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *PortalRevisionClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, parameters PortalRevisionContract, options *PortalRevisionClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, portalRevisionID, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PortalRevisionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, parameters PortalRevisionContract, options *PortalRevisionClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalRevisions/{portalRevisionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalRevisionID == "" {
		return nil, errors.New("parameter portalRevisionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalRevisionId}", url.PathEscape(portalRevisionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

// Get - Gets the developer portal's revision specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalRevisionID - Portal revision identifier. Must be unique in the current API Management service instance.
//   - options - PortalRevisionClientGetOptions contains the optional parameters for the PortalRevisionClient.Get method.
func (client *PortalRevisionClient) Get(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, options *PortalRevisionClientGetOptions) (PortalRevisionClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, portalRevisionID, options)
	if err != nil {
		return PortalRevisionClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PortalRevisionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PortalRevisionClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PortalRevisionClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, options *PortalRevisionClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalRevisions/{portalRevisionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalRevisionID == "" {
		return nil, errors.New("parameter portalRevisionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalRevisionId}", url.PathEscape(portalRevisionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PortalRevisionClient) getHandleResponse(resp *http.Response) (PortalRevisionClientGetResponse, error) {
	result := PortalRevisionClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalRevisionContract); err != nil {
		return PortalRevisionClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the developer portal revision specified by its identifier.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalRevisionID - Portal revision identifier. Must be unique in the current API Management service instance.
//   - options - PortalRevisionClientGetEntityTagOptions contains the optional parameters for the PortalRevisionClient.GetEntityTag
//     method.
func (client *PortalRevisionClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, options *PortalRevisionClientGetEntityTagOptions) (PortalRevisionClientGetEntityTagResponse, error) {
	var err error
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, portalRevisionID, options)
	if err != nil {
		return PortalRevisionClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PortalRevisionClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PortalRevisionClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *PortalRevisionClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, options *PortalRevisionClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalRevisions/{portalRevisionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalRevisionID == "" {
		return nil, errors.New("parameter portalRevisionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalRevisionId}", url.PathEscape(portalRevisionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *PortalRevisionClient) getEntityTagHandleResponse(resp *http.Response) (PortalRevisionClientGetEntityTagResponse, error) {
	result := PortalRevisionClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	result.Success = resp.StatusCode >= 200 && resp.StatusCode < 300
	return result, nil
}

// NewListByServicePager - Lists developer portal's revisions.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - PortalRevisionClientListByServiceOptions contains the optional parameters for the PortalRevisionClient.NewListByServicePager
//     method.
func (client *PortalRevisionClient) NewListByServicePager(resourceGroupName string, serviceName string, options *PortalRevisionClientListByServiceOptions) (*runtime.Pager[PortalRevisionClientListByServiceResponse]) {
	return runtime.NewPager(runtime.PagingHandler[PortalRevisionClientListByServiceResponse]{
		More: func(page PortalRevisionClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PortalRevisionClientListByServiceResponse) (PortalRevisionClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PortalRevisionClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PortalRevisionClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PortalRevisionClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PortalRevisionClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PortalRevisionClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalRevisions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PortalRevisionClient) listByServiceHandleResponse(resp *http.Response) (PortalRevisionClientListByServiceResponse, error) {
	result := PortalRevisionClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalRevisionCollection); err != nil {
		return PortalRevisionClientListByServiceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the description of specified portal revision or makes it current.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - portalRevisionID - Portal revision identifier. Must be unique in the current API Management service instance.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - PortalRevisionClientBeginUpdateOptions contains the optional parameters for the PortalRevisionClient.BeginUpdate
//     method.
func (client *PortalRevisionClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, ifMatch string, parameters PortalRevisionContract, options *PortalRevisionClientBeginUpdateOptions) (*runtime.Poller[PortalRevisionClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serviceName, portalRevisionID, ifMatch, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PortalRevisionClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PortalRevisionClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates the description of specified portal revision or makes it current.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
func (client *PortalRevisionClient) update(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, ifMatch string, parameters PortalRevisionContract, options *PortalRevisionClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, portalRevisionID, ifMatch, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *PortalRevisionClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, portalRevisionID string, ifMatch string, parameters PortalRevisionContract, options *PortalRevisionClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalRevisions/{portalRevisionId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if portalRevisionID == "" {
		return nil, errors.New("parameter portalRevisionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{portalRevisionId}", url.PathEscape(portalRevisionID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
	return nil, err
}
	return req, nil
}

