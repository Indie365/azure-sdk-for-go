//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecuritydevops

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

// GitHubConnectorStatsClient contains the methods for the GitHubConnectorStats group.
// Don't use this type directly, use NewGitHubConnectorStatsClient() instead.
type GitHubConnectorStatsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGitHubConnectorStatsClient creates a new instance of GitHubConnectorStatsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGitHubConnectorStatsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GitHubConnectorStatsClient, error) {
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
	client := &GitHubConnectorStatsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns the summary of the GitHub Connector Stats.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// gitHubConnectorName - Name of the GitHub Connector.
// options - GitHubConnectorStatsClientGetOptions contains the optional parameters for the GitHubConnectorStatsClient.Get
// method.
func (client *GitHubConnectorStatsClient) Get(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *GitHubConnectorStatsClientGetOptions) (GitHubConnectorStatsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gitHubConnectorName, options)
	if err != nil {
		return GitHubConnectorStatsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GitHubConnectorStatsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GitHubConnectorStatsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GitHubConnectorStatsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *GitHubConnectorStatsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}/stats"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gitHubConnectorName == "" {
		return nil, errors.New("parameter gitHubConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gitHubConnectorName}", url.PathEscape(gitHubConnectorName))
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
func (client *GitHubConnectorStatsClient) getHandleResponse(resp *http.Response) (GitHubConnectorStatsClientGetResponse, error) {
	result := GitHubConnectorStatsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubConnectorStatsListResponse); err != nil {
		return GitHubConnectorStatsClientGetResponse{}, err
	}
	return result, nil
}
