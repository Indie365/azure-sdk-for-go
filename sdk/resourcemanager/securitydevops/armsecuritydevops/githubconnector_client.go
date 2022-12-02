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

// GitHubConnectorClient contains the methods for the GitHubConnector group.
// Don't use this type directly, use NewGitHubConnectorClient() instead.
type GitHubConnectorClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGitHubConnectorClient creates a new instance of GitHubConnectorClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGitHubConnectorClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GitHubConnectorClient, error) {
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
	client := &GitHubConnectorClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a monitored GitHub Connector resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// gitHubConnectorName - Name of the GitHub Connector.
// gitHubConnector - Connector resource payload.
// options - GitHubConnectorClientBeginCreateOrUpdateOptions contains the optional parameters for the GitHubConnectorClient.BeginCreateOrUpdate
// method.
func (client *GitHubConnectorClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector GitHubConnector, options *GitHubConnectorClientBeginCreateOrUpdateOptions) (*runtime.Poller[GitHubConnectorClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, gitHubConnectorName, gitHubConnector, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[GitHubConnectorClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[GitHubConnectorClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update a monitored GitHub Connector resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *GitHubConnectorClient) createOrUpdate(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector GitHubConnector, options *GitHubConnectorClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gitHubConnectorName, gitHubConnector, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GitHubConnectorClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector GitHubConnector, options *GitHubConnectorClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, gitHubConnector)
}

// BeginDelete - Delete monitored GitHub Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// gitHubConnectorName - Name of the GitHub Connector.
// options - GitHubConnectorClientBeginDeleteOptions contains the optional parameters for the GitHubConnectorClient.BeginDelete
// method.
func (client *GitHubConnectorClient) BeginDelete(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *GitHubConnectorClientBeginDeleteOptions) (*runtime.Poller[GitHubConnectorClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, gitHubConnectorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GitHubConnectorClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GitHubConnectorClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete monitored GitHub Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *GitHubConnectorClient) deleteOperation(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *GitHubConnectorClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gitHubConnectorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GitHubConnectorClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *GitHubConnectorClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a monitored GitHub Connector resource for a given ID.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// gitHubConnectorName - Name of the GitHub Connector.
// options - GitHubConnectorClientGetOptions contains the optional parameters for the GitHubConnectorClient.Get method.
func (client *GitHubConnectorClient) Get(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *GitHubConnectorClientGetOptions) (GitHubConnectorClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gitHubConnectorName, options)
	if err != nil {
		return GitHubConnectorClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GitHubConnectorClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GitHubConnectorClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GitHubConnectorClient) getCreateRequest(ctx context.Context, resourceGroupName string, gitHubConnectorName string, options *GitHubConnectorClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}"
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
func (client *GitHubConnectorClient) getHandleResponse(resp *http.Response) (GitHubConnectorClientGetResponse, error) {
	result := GitHubConnectorClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubConnector); err != nil {
		return GitHubConnectorClientGetResponse{}, err
	}
	return result, nil
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - GitHubConnectorClientListByResourceGroupOptions contains the optional parameters for the GitHubConnectorClient.ListByResourceGroup
// method.
func (client *GitHubConnectorClient) NewListByResourceGroupPager(resourceGroupName string, options *GitHubConnectorClientListByResourceGroupOptions) *runtime.Pager[GitHubConnectorClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[GitHubConnectorClientListByResourceGroupResponse]{
		More: func(page GitHubConnectorClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GitHubConnectorClientListByResourceGroupResponse) (GitHubConnectorClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GitHubConnectorClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GitHubConnectorClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GitHubConnectorClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *GitHubConnectorClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *GitHubConnectorClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *GitHubConnectorClient) listByResourceGroupHandleResponse(resp *http.Response) (GitHubConnectorClientListByResourceGroupResponse, error) {
	result := GitHubConnectorClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubConnectorListResponse); err != nil {
		return GitHubConnectorClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns a list of monitored GitHub Connectors.
// Generated from API version 2022-09-01-preview
// options - GitHubConnectorClientListBySubscriptionOptions contains the optional parameters for the GitHubConnectorClient.ListBySubscription
// method.
func (client *GitHubConnectorClient) NewListBySubscriptionPager(options *GitHubConnectorClientListBySubscriptionOptions) *runtime.Pager[GitHubConnectorClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[GitHubConnectorClientListBySubscriptionResponse]{
		More: func(page GitHubConnectorClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GitHubConnectorClientListBySubscriptionResponse) (GitHubConnectorClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GitHubConnectorClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GitHubConnectorClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GitHubConnectorClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *GitHubConnectorClient) listBySubscriptionCreateRequest(ctx context.Context, options *GitHubConnectorClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.SecurityDevOps/gitHubConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *GitHubConnectorClient) listBySubscriptionHandleResponse(resp *http.Response) (GitHubConnectorClientListBySubscriptionResponse, error) {
	result := GitHubConnectorClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubConnectorListResponse); err != nil {
		return GitHubConnectorClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update monitored GitHub Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// gitHubConnectorName - Name of the GitHub Connector.
// gitHubConnector - Connector resource payload.
// options - GitHubConnectorClientBeginUpdateOptions contains the optional parameters for the GitHubConnectorClient.BeginUpdate
// method.
func (client *GitHubConnectorClient) BeginUpdate(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector GitHubConnector, options *GitHubConnectorClientBeginUpdateOptions) (*runtime.Poller[GitHubConnectorClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, gitHubConnectorName, gitHubConnector, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GitHubConnectorClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[GitHubConnectorClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update monitored GitHub Connector details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *GitHubConnectorClient) update(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector GitHubConnector, options *GitHubConnectorClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, gitHubConnectorName, gitHubConnector, options)
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

// updateCreateRequest creates the Update request.
func (client *GitHubConnectorClient) updateCreateRequest(ctx context.Context, resourceGroupName string, gitHubConnectorName string, gitHubConnector GitHubConnector, options *GitHubConnectorClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, gitHubConnector)
}