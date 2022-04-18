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

// ComponentAvailableFeaturesClient contains the methods for the ComponentAvailableFeatures group.
// Don't use this type directly, use NewComponentAvailableFeaturesClient() instead.
type ComponentAvailableFeaturesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewComponentAvailableFeaturesClient creates a new instance of ComponentAvailableFeaturesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewComponentAvailableFeaturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ComponentAvailableFeaturesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ComponentAvailableFeaturesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Returns all available features of the application insights component.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the Application Insights component resource.
// options - ComponentAvailableFeaturesClientGetOptions contains the optional parameters for the ComponentAvailableFeaturesClient.Get
// method.
func (client *ComponentAvailableFeaturesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentAvailableFeaturesClientGetOptions) (ComponentAvailableFeaturesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ComponentAvailableFeaturesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ComponentAvailableFeaturesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ComponentAvailableFeaturesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ComponentAvailableFeaturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ComponentAvailableFeaturesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/getavailablebillingfeatures"
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
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ComponentAvailableFeaturesClient) getHandleResponse(resp *http.Response) (ComponentAvailableFeaturesClientGetResponse, error) {
	result := ComponentAvailableFeaturesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAvailableFeatures); err != nil {
		return ComponentAvailableFeaturesClientGetResponse{}, err
	}
	return result, nil
}
