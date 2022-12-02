//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// ExpressRouteCircuitConnectionsClient contains the methods for the ExpressRouteCircuitConnections group.
// Don't use this type directly, use NewExpressRouteCircuitConnectionsClient() instead.
type ExpressRouteCircuitConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRouteCircuitConnectionsClient creates a new instance of ExpressRouteCircuitConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExpressRouteCircuitConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteCircuitConnectionsClient, error) {
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
	client := &ExpressRouteCircuitConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Express Route Circuit Connection in the specified express route circuits.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// connectionName - The name of the express route circuit connection.
// expressRouteCircuitConnectionParameters - Parameters supplied to the create or update express route circuit connection
// operation.
// options - ExpressRouteCircuitConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ExpressRouteCircuitConnectionsClient.BeginCreateOrUpdate
// method.
func (client *ExpressRouteCircuitConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, expressRouteCircuitConnectionParameters ExpressRouteCircuitConnection, options *ExpressRouteCircuitConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ExpressRouteCircuitConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, circuitName, peeringName, connectionName, expressRouteCircuitConnectionParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ExpressRouteCircuitConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCircuitConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Express Route Circuit Connection in the specified express route circuits.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *ExpressRouteCircuitConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, expressRouteCircuitConnectionParameters ExpressRouteCircuitConnection, options *ExpressRouteCircuitConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, peeringName, connectionName, expressRouteCircuitConnectionParameters, options)
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
func (client *ExpressRouteCircuitConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, expressRouteCircuitConnectionParameters ExpressRouteCircuitConnection, options *ExpressRouteCircuitConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, expressRouteCircuitConnectionParameters)
}

// BeginDelete - Deletes the specified Express Route Circuit Connection from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// connectionName - The name of the express route circuit connection.
// options - ExpressRouteCircuitConnectionsClientBeginDeleteOptions contains the optional parameters for the ExpressRouteCircuitConnectionsClient.BeginDelete
// method.
func (client *ExpressRouteCircuitConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *ExpressRouteCircuitConnectionsClientBeginDeleteOptions) (*runtime.Poller[ExpressRouteCircuitConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, circuitName, peeringName, connectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ExpressRouteCircuitConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExpressRouteCircuitConnectionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified Express Route Circuit Connection from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
func (client *ExpressRouteCircuitConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *ExpressRouteCircuitConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, circuitName, peeringName, connectionName, options)
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
func (client *ExpressRouteCircuitConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *ExpressRouteCircuitConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Express Route Circuit Connection from the specified express route circuit.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// circuitName - The name of the express route circuit.
// peeringName - The name of the peering.
// connectionName - The name of the express route circuit connection.
// options - ExpressRouteCircuitConnectionsClientGetOptions contains the optional parameters for the ExpressRouteCircuitConnectionsClient.Get
// method.
func (client *ExpressRouteCircuitConnectionsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *ExpressRouteCircuitConnectionsClientGetOptions) (ExpressRouteCircuitConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, peeringName, connectionName, options)
	if err != nil {
		return ExpressRouteCircuitConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteCircuitConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteCircuitConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteCircuitConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *ExpressRouteCircuitConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteCircuitConnectionsClient) getHandleResponse(resp *http.Response) (ExpressRouteCircuitConnectionsClientGetResponse, error) {
	result := ExpressRouteCircuitConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitConnection); err != nil {
		return ExpressRouteCircuitConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all global reach connections associated with a private peering in an express route circuit.
// Generated from API version 2022-05-01
// resourceGroupName - The name of the resource group.
// circuitName - The name of the circuit.
// peeringName - The name of the peering.
// options - ExpressRouteCircuitConnectionsClientListOptions contains the optional parameters for the ExpressRouteCircuitConnectionsClient.List
// method.
func (client *ExpressRouteCircuitConnectionsClient) NewListPager(resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitConnectionsClientListOptions) *runtime.Pager[ExpressRouteCircuitConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRouteCircuitConnectionsClientListResponse]{
		More: func(page ExpressRouteCircuitConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteCircuitConnectionsClientListResponse) (ExpressRouteCircuitConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRouteCircuitConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExpressRouteCircuitConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRouteCircuitConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRouteCircuitConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/connections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteCircuitConnectionsClient) listHandleResponse(resp *http.Response) (ExpressRouteCircuitConnectionsClientListResponse, error) {
	result := ExpressRouteCircuitConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteCircuitConnectionListResult); err != nil {
		return ExpressRouteCircuitConnectionsClientListResponse{}, err
	}
	return result, nil
}