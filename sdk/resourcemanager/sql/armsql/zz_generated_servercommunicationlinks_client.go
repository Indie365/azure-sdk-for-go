//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerCommunicationLinksClient contains the methods for the ServerCommunicationLinks group.
// Don't use this type directly, use NewServerCommunicationLinksClient() instead.
type ServerCommunicationLinksClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerCommunicationLinksClient creates a new instance of ServerCommunicationLinksClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerCommunicationLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerCommunicationLinksClient, error) {
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
	client := &ServerCommunicationLinksClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a server communication link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// communicationLinkName - The name of the server communication link.
// parameters - The required parameters for creating a server communication link.
// options - ServerCommunicationLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerCommunicationLinksClient.BeginCreateOrUpdate
// method.
func (client *ServerCommunicationLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, parameters ServerCommunicationLink, options *ServerCommunicationLinksClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerCommunicationLinksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, communicationLinkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ServerCommunicationLinksClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ServerCommunicationLinksClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates a server communication link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
func (client *ServerCommunicationLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, parameters ServerCommunicationLink, options *ServerCommunicationLinksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, communicationLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerCommunicationLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, parameters ServerCommunicationLink, options *ServerCommunicationLinksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}"
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
	if communicationLinkName == "" {
		return nil, errors.New("parameter communicationLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationLinkName}", url.PathEscape(communicationLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Deletes a server communication link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// communicationLinkName - The name of the server communication link.
// options - ServerCommunicationLinksClientDeleteOptions contains the optional parameters for the ServerCommunicationLinksClient.Delete
// method.
func (client *ServerCommunicationLinksClient) Delete(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksClientDeleteOptions) (ServerCommunicationLinksClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, communicationLinkName, options)
	if err != nil {
		return ServerCommunicationLinksClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerCommunicationLinksClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerCommunicationLinksClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ServerCommunicationLinksClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerCommunicationLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}"
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
	if communicationLinkName == "" {
		return nil, errors.New("parameter communicationLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationLinkName}", url.PathEscape(communicationLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Returns a server communication link.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// communicationLinkName - The name of the server communication link.
// options - ServerCommunicationLinksClientGetOptions contains the optional parameters for the ServerCommunicationLinksClient.Get
// method.
func (client *ServerCommunicationLinksClient) Get(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksClientGetOptions) (ServerCommunicationLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, communicationLinkName, options)
	if err != nil {
		return ServerCommunicationLinksClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerCommunicationLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerCommunicationLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerCommunicationLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}"
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
	if communicationLinkName == "" {
		return nil, errors.New("parameter communicationLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationLinkName}", url.PathEscape(communicationLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerCommunicationLinksClient) getHandleResponse(resp *http.Response) (ServerCommunicationLinksClientGetResponse, error) {
	result := ServerCommunicationLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerCommunicationLink); err != nil {
		return ServerCommunicationLinksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of server communication links.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2014-04-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - ServerCommunicationLinksClientListByServerOptions contains the optional parameters for the ServerCommunicationLinksClient.ListByServer
// method.
func (client *ServerCommunicationLinksClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerCommunicationLinksClientListByServerOptions) *runtime.Pager[ServerCommunicationLinksClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerCommunicationLinksClientListByServerResponse]{
		More: func(page ServerCommunicationLinksClientListByServerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ServerCommunicationLinksClientListByServerResponse) (ServerCommunicationLinksClientListByServerResponse, error) {
			req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			if err != nil {
				return ServerCommunicationLinksClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServerCommunicationLinksClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerCommunicationLinksClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerCommunicationLinksClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerCommunicationLinksClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerCommunicationLinksClient) listByServerHandleResponse(resp *http.Response) (ServerCommunicationLinksClientListByServerResponse, error) {
	result := ServerCommunicationLinksClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerCommunicationLinkListResult); err != nil {
		return ServerCommunicationLinksClientListByServerResponse{}, err
	}
	return result, nil
}