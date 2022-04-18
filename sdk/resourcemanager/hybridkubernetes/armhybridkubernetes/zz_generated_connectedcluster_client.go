//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridkubernetes

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

// ConnectedClusterClient contains the methods for the ConnectedCluster group.
// Don't use this type directly, use NewConnectedClusterClient() instead.
type ConnectedClusterClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectedClusterClient creates a new instance of ConnectedClusterClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConnectedClusterClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedClusterClient, error) {
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
	client := &ConnectedClusterClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - API to register a new Kubernetes cluster and create a tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the Kubernetes cluster on which get is called.
// connectedCluster - Parameters supplied to Create a Connected Cluster.
// options - ConnectedClusterClientBeginCreateOptions contains the optional parameters for the ConnectedClusterClient.BeginCreate
// method.
func (client *ConnectedClusterClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster ConnectedCluster, options *ConnectedClusterClientBeginCreateOptions) (*armruntime.Poller[ConnectedClusterClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, clusterName, connectedCluster, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ConnectedClusterClientCreateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ConnectedClusterClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - API to register a new Kubernetes cluster and create a tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConnectedClusterClient) create(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster ConnectedCluster, options *ConnectedClusterClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, connectedCluster, options)
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

// createCreateRequest creates the Create request.
func (client *ConnectedClusterClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster ConnectedCluster, options *ConnectedClusterClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, connectedCluster)
}

// BeginDelete - Delete a connected cluster, removing the tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the Kubernetes cluster on which get is called.
// options - ConnectedClusterClientBeginDeleteOptions contains the optional parameters for the ConnectedClusterClient.BeginDelete
// method.
func (client *ConnectedClusterClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientBeginDeleteOptions) (*armruntime.Poller[ConnectedClusterClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ConnectedClusterClientDeleteResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ConnectedClusterClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete a connected cluster, removing the tracked resource in Azure Resource Manager (ARM).
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ConnectedClusterClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
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
func (client *ConnectedClusterClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns the properties of the specified connected cluster, including name, identity, properties, and additional cluster
// details.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the Kubernetes cluster on which get is called.
// options - ConnectedClusterClientGetOptions contains the optional parameters for the ConnectedClusterClient.Get method.
func (client *ConnectedClusterClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientGetOptions) (ConnectedClusterClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ConnectedClusterClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedClusterClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedClusterClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectedClusterClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConnectedClusterClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
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

// getHandleResponse handles the Get response.
func (client *ConnectedClusterClient) getHandleResponse(resp *http.Response) (ConnectedClusterClientGetResponse, error) {
	result := ConnectedClusterClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedCluster); err != nil {
		return ConnectedClusterClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - API to enumerate registered connected K8s clusters under a Resource Group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - ConnectedClusterClientListByResourceGroupOptions contains the optional parameters for the ConnectedClusterClient.ListByResourceGroup
// method.
func (client *ConnectedClusterClient) NewListByResourceGroupPager(resourceGroupName string, options *ConnectedClusterClientListByResourceGroupOptions) *runtime.Pager[ConnectedClusterClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[ConnectedClusterClientListByResourceGroupResponse]{
		More: func(page ConnectedClusterClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedClusterClientListByResourceGroupResponse) (ConnectedClusterClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectedClusterClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConnectedClusterClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectedClusterClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ConnectedClusterClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ConnectedClusterClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters"
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
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ConnectedClusterClient) listByResourceGroupHandleResponse(resp *http.Response) (ConnectedClusterClientListByResourceGroupResponse, error) {
	result := ConnectedClusterClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedClusterList); err != nil {
		return ConnectedClusterClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - API to enumerate registered connected K8s clusters under a Subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - ConnectedClusterClientListBySubscriptionOptions contains the optional parameters for the ConnectedClusterClient.ListBySubscription
// method.
func (client *ConnectedClusterClient) NewListBySubscriptionPager(options *ConnectedClusterClientListBySubscriptionOptions) *runtime.Pager[ConnectedClusterClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[ConnectedClusterClientListBySubscriptionResponse]{
		More: func(page ConnectedClusterClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedClusterClientListBySubscriptionResponse) (ConnectedClusterClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectedClusterClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConnectedClusterClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectedClusterClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ConnectedClusterClient) listBySubscriptionCreateRequest(ctx context.Context, options *ConnectedClusterClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Kubernetes/connectedClusters"
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ConnectedClusterClient) listBySubscriptionHandleResponse(resp *http.Response) (ConnectedClusterClientListBySubscriptionResponse, error) {
	result := ConnectedClusterClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedClusterList); err != nil {
		return ConnectedClusterClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListClusterUserCredential - Gets cluster user credentials of the connected cluster with a specified resource group and
// name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the Kubernetes cluster on which get is called.
// properties - ListClusterUserCredential properties
// options - ConnectedClusterClientListClusterUserCredentialOptions contains the optional parameters for the ConnectedClusterClient.ListClusterUserCredential
// method.
func (client *ConnectedClusterClient) ListClusterUserCredential(ctx context.Context, resourceGroupName string, clusterName string, properties ListClusterUserCredentialProperties, options *ConnectedClusterClientListClusterUserCredentialOptions) (ConnectedClusterClientListClusterUserCredentialResponse, error) {
	req, err := client.listClusterUserCredentialCreateRequest(ctx, resourceGroupName, clusterName, properties, options)
	if err != nil {
		return ConnectedClusterClientListClusterUserCredentialResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedClusterClientListClusterUserCredentialResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedClusterClientListClusterUserCredentialResponse{}, runtime.NewResponseError(resp)
	}
	return client.listClusterUserCredentialHandleResponse(resp)
}

// listClusterUserCredentialCreateRequest creates the ListClusterUserCredential request.
func (client *ConnectedClusterClient) listClusterUserCredentialCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, properties ListClusterUserCredentialProperties, options *ConnectedClusterClientListClusterUserCredentialOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}/listClusterUserCredential"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// listClusterUserCredentialHandleResponse handles the ListClusterUserCredential response.
func (client *ConnectedClusterClient) listClusterUserCredentialHandleResponse(resp *http.Response) (ConnectedClusterClientListClusterUserCredentialResponse, error) {
	result := ConnectedClusterClientListClusterUserCredentialResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialResults); err != nil {
		return ConnectedClusterClientListClusterUserCredentialResponse{}, err
	}
	return result, nil
}

// Update - API to update certain properties of the connected cluster resource
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// clusterName - The name of the Kubernetes cluster on which get is called.
// connectedClusterPatch - Parameters supplied to update Connected Cluster.
// options - ConnectedClusterClientUpdateOptions contains the optional parameters for the ConnectedClusterClient.Update method.
func (client *ConnectedClusterClient) Update(ctx context.Context, resourceGroupName string, clusterName string, connectedClusterPatch ConnectedClusterPatch, options *ConnectedClusterClientUpdateOptions) (ConnectedClusterClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, connectedClusterPatch, options)
	if err != nil {
		return ConnectedClusterClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectedClusterClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectedClusterClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ConnectedClusterClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, connectedClusterPatch ConnectedClusterPatch, options *ConnectedClusterClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedClusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, connectedClusterPatch)
}

// updateHandleResponse handles the Update response.
func (client *ConnectedClusterClient) updateHandleResponse(resp *http.Response) (ConnectedClusterClientUpdateResponse, error) {
	result := ConnectedClusterClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedCluster); err != nil {
		return ConnectedClusterClientUpdateResponse{}, err
	}
	return result, nil
}
