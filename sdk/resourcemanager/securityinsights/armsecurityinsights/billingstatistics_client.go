//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// BillingStatisticsClient contains the methods for the BillingStatistics group.
// Don't use this type directly, use NewBillingStatisticsClient() instead.
type BillingStatisticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBillingStatisticsClient creates a new instance of BillingStatisticsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBillingStatisticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BillingStatisticsClient, error) {
	cl, err := arm.NewClient(moduleName+".BillingStatisticsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BillingStatisticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a billing statistic
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - billingStatisticName - The name of the billing statistic
//   - options - BillingStatisticsClientGetOptions contains the optional parameters for the BillingStatisticsClient.Get method.
func (client *BillingStatisticsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, billingStatisticName string, options *BillingStatisticsClientGetOptions) (BillingStatisticsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, billingStatisticName, options)
	if err != nil {
		return BillingStatisticsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BillingStatisticsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BillingStatisticsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BillingStatisticsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, billingStatisticName string, options *BillingStatisticsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/billingStatistics/{billingStatisticName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if billingStatisticName == "" {
		return nil, errors.New("parameter billingStatisticName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingStatisticName}", url.PathEscape(billingStatisticName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BillingStatisticsClient) getHandleResponse(resp *http.Response) (BillingStatisticsClientGetResponse, error) {
	result := BillingStatisticsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return BillingStatisticsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all Microsoft Sentinel billing statistics.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - BillingStatisticsClientListOptions contains the optional parameters for the BillingStatisticsClient.NewListPager
//     method.
func (client *BillingStatisticsClient) NewListPager(resourceGroupName string, workspaceName string, options *BillingStatisticsClientListOptions) *runtime.Pager[BillingStatisticsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BillingStatisticsClientListResponse]{
		More: func(page BillingStatisticsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BillingStatisticsClientListResponse) (BillingStatisticsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BillingStatisticsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BillingStatisticsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BillingStatisticsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BillingStatisticsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *BillingStatisticsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/billingStatistics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BillingStatisticsClient) listHandleResponse(resp *http.Response) (BillingStatisticsClientListResponse, error) {
	result := BillingStatisticsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingStatisticList); err != nil {
		return BillingStatisticsClientListResponse{}, err
	}
	return result, nil
}
