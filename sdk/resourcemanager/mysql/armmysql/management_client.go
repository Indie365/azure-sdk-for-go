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

// ManagementClient contains the methods for the MySQLManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagementClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateRecommendedActionSession - Create recommendation action session for the advisor.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - advisorName - The advisor name for recommendation action.
//   - databaseName - The name of the database.
//   - options - ManagementClientBeginCreateRecommendedActionSessionOptions contains the optional parameters for the ManagementClient.BeginCreateRecommendedActionSession
//     method.
func (client *ManagementClient) BeginCreateRecommendedActionSession(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string, options *ManagementClientBeginCreateRecommendedActionSessionOptions) (*runtime.Poller[ManagementClientCreateRecommendedActionSessionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createRecommendedActionSession(ctx, resourceGroupName, serverName, advisorName, databaseName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ManagementClientCreateRecommendedActionSessionResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ManagementClientCreateRecommendedActionSessionResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateRecommendedActionSession - Create recommendation action session for the advisor.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *ManagementClient) createRecommendedActionSession(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string, options *ManagementClientBeginCreateRecommendedActionSessionOptions) (*http.Response, error) {
	req, err := client.createRecommendedActionSessionCreateRequest(ctx, resourceGroupName, serverName, advisorName, databaseName, options)
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

// createRecommendedActionSessionCreateRequest creates the CreateRecommendedActionSession request.
func (client *ManagementClient) createRecommendedActionSessionCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string, options *ManagementClientBeginCreateRecommendedActionSessionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/advisors/{advisorName}/createRecommendedActionSession"
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
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	reqQP.Set("databaseName", databaseName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// ResetQueryPerformanceInsightData - Reset data for Query Performance Insight.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ManagementClientResetQueryPerformanceInsightDataOptions contains the optional parameters for the ManagementClient.ResetQueryPerformanceInsightData
//     method.
func (client *ManagementClient) ResetQueryPerformanceInsightData(ctx context.Context, resourceGroupName string, serverName string, options *ManagementClientResetQueryPerformanceInsightDataOptions) (ManagementClientResetQueryPerformanceInsightDataResponse, error) {
	req, err := client.resetQueryPerformanceInsightDataCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ManagementClientResetQueryPerformanceInsightDataResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientResetQueryPerformanceInsightDataResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagementClientResetQueryPerformanceInsightDataResponse{}, runtime.NewResponseError(resp)
	}
	return client.resetQueryPerformanceInsightDataHandleResponse(resp)
}

// resetQueryPerformanceInsightDataCreateRequest creates the ResetQueryPerformanceInsightData request.
func (client *ManagementClient) resetQueryPerformanceInsightDataCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ManagementClientResetQueryPerformanceInsightDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/resetQueryPerformanceInsightData"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// resetQueryPerformanceInsightDataHandleResponse handles the ResetQueryPerformanceInsightData response.
func (client *ManagementClient) resetQueryPerformanceInsightDataHandleResponse(resp *http.Response) (ManagementClientResetQueryPerformanceInsightDataResponse, error) {
	result := ManagementClientResetQueryPerformanceInsightDataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryPerformanceInsightResetDataResult); err != nil {
		return ManagementClientResetQueryPerformanceInsightDataResponse{}, err
	}
	return result, nil
}
