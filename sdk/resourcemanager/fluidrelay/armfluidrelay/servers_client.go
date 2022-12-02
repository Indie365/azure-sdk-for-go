//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfluidrelay

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

// ServersClient contains the methods for the FluidRelayServers group.
// Don't use this type directly, use NewServersClient() instead.
type ServersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServersClient creates a new instance of ServersClient with the specified values.
// subscriptionID - The subscription id (GUID) for this resource.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServersClient, error) {
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
	client := &ServersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update a Fluid Relay server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroup - The resource group containing the resource.
// fluidRelayServerName - The Fluid Relay server resource name.
// resource - The details of the Fluid Relay server resource.
// options - ServersClientCreateOrUpdateOptions contains the optional parameters for the ServersClient.CreateOrUpdate method.
func (client *ServersClient) CreateOrUpdate(ctx context.Context, resourceGroup string, fluidRelayServerName string, resource Server, options *ServersClientCreateOrUpdateOptions) (ServersClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroup, fluidRelayServerName, resource, options)
	if err != nil {
		return ServersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServersClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, resource Server, options *ServersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServersClient) createOrUpdateHandleResponse(resp *http.Response) (ServersClientCreateOrUpdateResponse, error) {
	result := ServersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Server); err != nil {
		return ServersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Fluid Relay server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroup - The resource group containing the resource.
// fluidRelayServerName - The Fluid Relay server resource name.
// options - ServersClientDeleteOptions contains the optional parameters for the ServersClient.Delete method.
func (client *ServersClient) Delete(ctx context.Context, resourceGroup string, fluidRelayServerName string, options *ServersClientDeleteOptions) (ServersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroup, fluidRelayServerName, options)
	if err != nil {
		return ServersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ServersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ServersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServersClient) deleteCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, options *ServersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Fluid Relay server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroup - The resource group containing the resource.
// fluidRelayServerName - The Fluid Relay server resource name.
// options - ServersClientGetOptions contains the optional parameters for the ServersClient.Get method.
func (client *ServersClient) Get(ctx context.Context, resourceGroup string, fluidRelayServerName string, options *ServersClientGetOptions) (ServersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroup, fluidRelayServerName, options)
	if err != nil {
		return ServersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServersClient) getCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, options *ServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
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
func (client *ServersClient) getHandleResponse(resp *http.Response) (ServersClientGetResponse, error) {
	result := ServersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Server); err != nil {
		return ServersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Fluid Relay servers in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroup - The resource group containing the resource.
// options - ServersClientListByResourceGroupOptions contains the optional parameters for the ServersClient.ListByResourceGroup
// method.
func (client *ServersClient) NewListByResourceGroupPager(resourceGroup string, options *ServersClientListByResourceGroupOptions) *runtime.Pager[ServersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServersClientListByResourceGroupResponse]{
		More: func(page ServersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServersClientListByResourceGroupResponse) (ServersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroup, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroup string, options *ServersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServersClient) listByResourceGroupHandleResponse(resp *http.Response) (ServersClientListByResourceGroupResponse, error) {
	result := ServersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerList); err != nil {
		return ServersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all Fluid Relay servers in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// options - ServersClientListBySubscriptionOptions contains the optional parameters for the ServersClient.ListBySubscription
// method.
func (client *ServersClient) NewListBySubscriptionPager(options *ServersClientListBySubscriptionOptions) *runtime.Pager[ServersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServersClientListBySubscriptionResponse]{
		More: func(page ServersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServersClientListBySubscriptionResponse) (ServersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ServersClient) listBySubscriptionCreateRequest(ctx context.Context, options *ServersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.FluidRelay/fluidRelayServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ServersClient) listBySubscriptionHandleResponse(resp *http.Response) (ServersClientListBySubscriptionResponse, error) {
	result := ServersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerList); err != nil {
		return ServersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListKeys - Get primary and secondary key for this server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroup - The resource group containing the resource.
// fluidRelayServerName - The Fluid Relay server resource name.
// options - ServersClientListKeysOptions contains the optional parameters for the ServersClient.ListKeys method.
func (client *ServersClient) ListKeys(ctx context.Context, resourceGroup string, fluidRelayServerName string, options *ServersClientListKeysOptions) (ServersClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroup, fluidRelayServerName, options)
	if err != nil {
		return ServersClientListKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServersClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *ServersClient) listKeysCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, options *ServersClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *ServersClient) listKeysHandleResponse(resp *http.Response) (ServersClientListKeysResponse, error) {
	result := ServersClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerKeys); err != nil {
		return ServersClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKey - Regenerate the primary or secondary key for this server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroup - The resource group containing the resource.
// fluidRelayServerName - The Fluid Relay server resource name.
// parameters - The details of which keys to generate.
// options - ServersClientRegenerateKeyOptions contains the optional parameters for the ServersClient.RegenerateKey method.
func (client *ServersClient) RegenerateKey(ctx context.Context, resourceGroup string, fluidRelayServerName string, parameters RegenerateKeyRequest, options *ServersClientRegenerateKeyOptions) (ServersClientRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroup, fluidRelayServerName, parameters, options)
	if err != nil {
		return ServersClientRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersClientRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServersClientRegenerateKeyResponse{}, runtime.NewResponseError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *ServersClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, parameters RegenerateKeyRequest, options *ServersClientRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *ServersClient) regenerateKeyHandleResponse(resp *http.Response) (ServersClientRegenerateKeyResponse, error) {
	result := ServersClientRegenerateKeyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerKeys); err != nil {
		return ServersClientRegenerateKeyResponse{}, err
	}
	return result, nil
}

// Update - Update a Fluid Relay server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceGroup - The resource group containing the resource.
// fluidRelayServerName - The Fluid Relay server resource name.
// resource - The details of the Fluid Relay server resource included in update calls.
// options - ServersClientUpdateOptions contains the optional parameters for the ServersClient.Update method.
func (client *ServersClient) Update(ctx context.Context, resourceGroup string, fluidRelayServerName string, resource ServerUpdate, options *ServersClientUpdateOptions) (ServersClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroup, fluidRelayServerName, resource, options)
	if err != nil {
		return ServersClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServersClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServersClient) updateCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, resource ServerUpdate, options *ServersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// updateHandleResponse handles the Update response.
func (client *ServersClient) updateHandleResponse(resp *http.Response) (ServersClientUpdateResponse, error) {
	result := ServersClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Server); err != nil {
		return ServersClientUpdateResponse{}, err
	}
	return result, nil
}