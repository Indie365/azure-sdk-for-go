//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysql

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

// ServerAdministratorsClient contains the methods for the ServerAdministrators group.
// Don't use this type directly, use NewServerAdministratorsClient() instead.
type ServerAdministratorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerAdministratorsClient creates a new instance of ServerAdministratorsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerAdministratorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerAdministratorsClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerAdministratorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerAdministratorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite
// the existing administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - properties - The required parameters for creating or updating an AAD server administrator.
//   - options - ServerAdministratorsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerAdministratorsClient.BeginCreateOrUpdate
//     method.
func (client *ServerAdministratorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerAdministratorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, properties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerAdministratorsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerAdministratorsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or update active directory administrator on an existing server. The update action will overwrite
// the existing administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
func (client *ServerAdministratorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerAdministratorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, properties ServerAdministratorResource, options *ServerAdministratorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// BeginDelete - Deletes server active directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ServerAdministratorsClientBeginDeleteOptions contains the optional parameters for the ServerAdministratorsClient.BeginDelete
//     method.
func (client *ServerAdministratorsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientBeginDeleteOptions) (*runtime.Poller[ServerAdministratorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerAdministratorsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerAdministratorsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes server active directory administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
func (client *ServerAdministratorsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerAdministratorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about a AAD server administrator.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ServerAdministratorsClientGetOptions contains the optional parameters for the ServerAdministratorsClient.Get
//     method.
func (client *ServerAdministratorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientGetOptions) (ServerAdministratorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdministratorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerAdministratorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdministratorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAdministratorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators/activeDirectory"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerAdministratorsClient) getHandleResponse(resp *http.Response) (ServerAdministratorsClientGetResponse, error) {
	result := ServerAdministratorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAdministratorResource); err != nil {
		return ServerAdministratorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of server Administrators.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ServerAdministratorsClientListOptions contains the optional parameters for the ServerAdministratorsClient.NewListPager
//     method.
func (client *ServerAdministratorsClient) NewListPager(resourceGroupName string, serverName string, options *ServerAdministratorsClientListOptions) *runtime.Pager[ServerAdministratorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerAdministratorsClientListResponse]{
		More: func(page ServerAdministratorsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServerAdministratorsClientListResponse) (ServerAdministratorsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, serverName, options)
			if err != nil {
				return ServerAdministratorsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerAdministratorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerAdministratorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ServerAdministratorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdministratorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/administrators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServerAdministratorsClient) listHandleResponse(resp *http.Response) (ServerAdministratorsClientListResponse, error) {
	result := ServerAdministratorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerAdministratorResourceListResult); err != nil {
		return ServerAdministratorsClientListResponse{}, err
	}
	return result, nil
}
