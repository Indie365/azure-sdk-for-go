//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

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

// ServerParametersClient contains the methods for the ServerParameters group.
// Don't use this type directly, use NewServerParametersClient() instead.
type ServerParametersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerParametersClient creates a new instance of ServerParametersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerParametersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerParametersClient, error) {
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
	client := &ServerParametersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginListUpdateConfigurations - Update a list of configurations in a given server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-12-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// value - The parameters for updating a list of server configuration.
// options - ServerParametersClientBeginListUpdateConfigurationsOptions contains the optional parameters for the ServerParametersClient.BeginListUpdateConfigurations
// method.
func (client *ServerParametersClient) BeginListUpdateConfigurations(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult, options *ServerParametersClientBeginListUpdateConfigurationsOptions) (*runtime.Poller[ServerParametersClientListUpdateConfigurationsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listUpdateConfigurations(ctx, resourceGroupName, serverName, value, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ServerParametersClientListUpdateConfigurationsResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServerParametersClientListUpdateConfigurationsResponse](options.ResumeToken, client.pl, nil)
	}
}

// ListUpdateConfigurations - Update a list of configurations in a given server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-12-01
func (client *ServerParametersClient) listUpdateConfigurations(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult, options *ServerParametersClientBeginListUpdateConfigurationsOptions) (*http.Response, error) {
	req, err := client.listUpdateConfigurationsCreateRequest(ctx, resourceGroupName, serverName, value, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listUpdateConfigurationsCreateRequest creates the ListUpdateConfigurations request.
func (client *ServerParametersClient) listUpdateConfigurationsCreateRequest(ctx context.Context, resourceGroupName string, serverName string, value ConfigurationListResult, options *ServerParametersClientBeginListUpdateConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/updateConfigurations"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, value)
}
