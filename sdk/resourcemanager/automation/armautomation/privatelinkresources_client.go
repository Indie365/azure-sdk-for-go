//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	internal *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateLinkResourcesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
	internal: cl,
	}
	return client, nil
}

// NewAutomationPager - Gets the private link resources that need to be created for Automation account.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - PrivateLinkResourcesClientAutomationOptions contains the optional parameters for the PrivateLinkResourcesClient.NewAutomationPager
//     method.
func (client *PrivateLinkResourcesClient) NewAutomationPager(resourceGroupName string, automationAccountName string, options *PrivateLinkResourcesClientAutomationOptions) (*runtime.Pager[PrivateLinkResourcesClientAutomationResponse]) {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientAutomationResponse]{
		More: func(page PrivateLinkResourcesClientAutomationResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientAutomationResponse) (PrivateLinkResourcesClientAutomationResponse, error) {
			req, err := client.automationCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			if err != nil {
				return PrivateLinkResourcesClientAutomationResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateLinkResourcesClientAutomationResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkResourcesClientAutomationResponse{}, runtime.NewResponseError(resp)
			}
			return client.automationHandleResponse(resp)
		},
	})
}

// automationCreateRequest creates the Automation request.
func (client *PrivateLinkResourcesClient) automationCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *PrivateLinkResourcesClientAutomationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// automationHandleResponse handles the Automation response.
func (client *PrivateLinkResourcesClient) automationHandleResponse(resp *http.Response) (PrivateLinkResourcesClientAutomationResponse, error) {
	result := PrivateLinkResourcesClientAutomationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientAutomationResponse{}, err
	}
	return result, nil
}

