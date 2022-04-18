//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift

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

// OpenShiftClustersClient contains the methods for the OpenShiftClusters group.
// Don't use this type directly, use NewOpenShiftClustersClient() instead.
type OpenShiftClustersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOpenShiftClustersClient creates a new instance of OpenShiftClustersClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOpenShiftClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OpenShiftClustersClient, error) {
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
	client := &OpenShiftClustersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// parameters - The OpenShift cluster resource.
// options - OpenShiftClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the OpenShiftClustersClient.BeginCreateOrUpdate
// method.
func (client *OpenShiftClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftCluster, options *OpenShiftClustersClientBeginCreateOrUpdateOptions) (*armruntime.Poller[OpenShiftClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[OpenShiftClustersClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[OpenShiftClustersClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OpenShiftClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftCluster, options *OpenShiftClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
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
func (client *OpenShiftClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftCluster, options *OpenShiftClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - The operation returns nothing.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// options - OpenShiftClustersClientBeginDeleteOptions contains the optional parameters for the OpenShiftClustersClient.BeginDelete
// method.
func (client *OpenShiftClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientBeginDeleteOptions) (*armruntime.Poller[OpenShiftClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[OpenShiftClustersClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[OpenShiftClustersClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - The operation returns nothing.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OpenShiftClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OpenShiftClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// options - OpenShiftClustersClientGetOptions contains the optional parameters for the OpenShiftClustersClient.Get method.
func (client *OpenShiftClustersClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientGetOptions) (OpenShiftClustersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return OpenShiftClustersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenShiftClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OpenShiftClustersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OpenShiftClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OpenShiftClustersClient) getHandleResponse(resp *http.Response) (OpenShiftClustersClientGetResponse, error) {
	result := OpenShiftClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftCluster); err != nil {
		return OpenShiftClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation returns properties of each OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// options - OpenShiftClustersClientListOptions contains the optional parameters for the OpenShiftClustersClient.List method.
func (client *OpenShiftClustersClient) NewListPager(options *OpenShiftClustersClientListOptions) *runtime.Pager[OpenShiftClustersClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[OpenShiftClustersClientListResponse]{
		More: func(page OpenShiftClustersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OpenShiftClustersClientListResponse) (OpenShiftClustersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OpenShiftClustersClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OpenShiftClustersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OpenShiftClustersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OpenShiftClustersClient) listCreateRequest(ctx context.Context, options *OpenShiftClustersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.RedHatOpenShift/openShiftClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OpenShiftClustersClient) listHandleResponse(resp *http.Response) (OpenShiftClustersClientListResponse, error) {
	result := OpenShiftClustersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftClusterList); err != nil {
		return OpenShiftClustersClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The operation returns properties of each OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - OpenShiftClustersClientListByResourceGroupOptions contains the optional parameters for the OpenShiftClustersClient.ListByResourceGroup
// method.
func (client *OpenShiftClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *OpenShiftClustersClientListByResourceGroupOptions) *runtime.Pager[OpenShiftClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[OpenShiftClustersClientListByResourceGroupResponse]{
		More: func(page OpenShiftClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OpenShiftClustersClientListByResourceGroupResponse) (OpenShiftClustersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OpenShiftClustersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OpenShiftClustersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OpenShiftClustersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OpenShiftClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OpenShiftClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters"
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
	reqQP.Set("api-version", "2020-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OpenShiftClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (OpenShiftClustersClientListByResourceGroupResponse, error) {
	result := OpenShiftClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftClusterList); err != nil {
		return OpenShiftClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListCredentials - The operation returns the credentials.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// options - OpenShiftClustersClientListCredentialsOptions contains the optional parameters for the OpenShiftClustersClient.ListCredentials
// method.
func (client *OpenShiftClustersClient) ListCredentials(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientListCredentialsOptions) (OpenShiftClustersClientListCredentialsResponse, error) {
	req, err := client.listCredentialsCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return OpenShiftClustersClientListCredentialsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OpenShiftClustersClientListCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OpenShiftClustersClientListCredentialsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listCredentialsHandleResponse(resp)
}

// listCredentialsCreateRequest creates the ListCredentials request.
func (client *OpenShiftClustersClient) listCredentialsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientListCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}/listCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listCredentialsHandleResponse handles the ListCredentials response.
func (client *OpenShiftClustersClient) listCredentialsHandleResponse(resp *http.Response) (OpenShiftClustersClientListCredentialsResponse, error) {
	result := OpenShiftClustersClientListCredentialsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftClusterCredentials); err != nil {
		return OpenShiftClustersClientListCredentialsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the OpenShift cluster resource.
// parameters - The OpenShift cluster resource.
// options - OpenShiftClustersClientBeginUpdateOptions contains the optional parameters for the OpenShiftClustersClient.BeginUpdate
// method.
func (client *OpenShiftClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftClusterUpdate, options *OpenShiftClustersClientBeginUpdateOptions) (*armruntime.Poller[OpenShiftClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[OpenShiftClustersClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[OpenShiftClustersClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OpenShiftClustersClient) update(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftClusterUpdate, options *OpenShiftClustersClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *OpenShiftClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftClusterUpdate, options *OpenShiftClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-04-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}
