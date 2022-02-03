//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ResourceSKUsClient contains the methods for the ResourceSKUs group.
// Don't use this type directly, use NewResourceSKUsClient() instead.
type ResourceSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourceSKUsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ResourceSKUsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets the list of Microsoft.CognitiveServices SKUs available for your Subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ResourceSKUsClientListOptions contains the optional parameters for the ResourceSKUsClient.List method.
func (client *ResourceSKUsClient) List(options *ResourceSKUsClientListOptions) *ResourceSKUsClientListPager {
	return &ResourceSKUsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ResourceSKUsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceSKUListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ResourceSKUsClient) listCreateRequest(ctx context.Context, options *ResourceSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceSKUsClient) listHandleResponse(resp *http.Response) (ResourceSKUsClientListResponse, error) {
	result := ResourceSKUsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUListResult); err != nil {
		return ResourceSKUsClientListResponse{}, err
	}
	return result, nil
}
