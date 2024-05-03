//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

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

// IotConnectorsClient contains the methods for the IotConnectors group.
// Don't use this type directly, use NewIotConnectorsClient() instead.
type IotConnectorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIotConnectorsClient creates a new instance of IotConnectorsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIotConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IotConnectorsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IotConnectorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an IoT Connector resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - iotConnectorName - The name of IoT Connector resource.
//   - iotConnector - The parameters for creating or updating an IoT Connectors resource.
//   - options - IotConnectorsClientBeginCreateOrUpdateOptions contains the optional parameters for the IotConnectorsClient.BeginCreateOrUpdate
//     method.
func (client *IotConnectorsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[IotConnectorsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, iotConnectorName, iotConnector, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IotConnectorsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IotConnectorsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an IoT Connector resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31
func (client *IotConnectorsClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "IotConnectorsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, iotConnectorName, iotConnector, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IotConnectorsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector IotConnector, options *IotConnectorsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, iotConnector); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - iotConnectorName - The name of IoT Connector resource.
//   - workspaceName - The name of workspace resource.
//   - options - IotConnectorsClientBeginDeleteOptions contains the optional parameters for the IotConnectorsClient.BeginDelete
//     method.
func (client *IotConnectorsClient) BeginDelete(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsClientBeginDeleteOptions) (*runtime.Poller[IotConnectorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, iotConnectorName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IotConnectorsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IotConnectorsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31
func (client *IotConnectorsClient) deleteOperation(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "IotConnectorsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, iotConnectorName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IotConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *IotConnectorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - iotConnectorName - The name of IoT Connector resource.
//   - options - IotConnectorsClientGetOptions contains the optional parameters for the IotConnectorsClient.Get method.
func (client *IotConnectorsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *IotConnectorsClientGetOptions) (IotConnectorsClientGetResponse, error) {
	var err error
	const operationName = "IotConnectorsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, iotConnectorName, options)
	if err != nil {
		return IotConnectorsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IotConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IotConnectorsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IotConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *IotConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotConnectorsClient) getHandleResponse(resp *http.Response) (IotConnectorsClientGetResponse, error) {
	result := IotConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IotConnector); err != nil {
		return IotConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists all IoT Connectors for the given workspace
//
// Generated from API version 2024-03-31
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - options - IotConnectorsClientListByWorkspaceOptions contains the optional parameters for the IotConnectorsClient.NewListByWorkspacePager
//     method.
func (client *IotConnectorsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *IotConnectorsClientListByWorkspaceOptions) *runtime.Pager[IotConnectorsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[IotConnectorsClientListByWorkspaceResponse]{
		More: func(page IotConnectorsClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IotConnectorsClientListByWorkspaceResponse) (IotConnectorsClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IotConnectorsClient.NewListByWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return IotConnectorsClientListByWorkspaceResponse{}, err
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *IotConnectorsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *IotConnectorsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *IotConnectorsClient) listByWorkspaceHandleResponse(resp *http.Response) (IotConnectorsClientListByWorkspaceResponse, error) {
	result := IotConnectorsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IotConnectorCollection); err != nil {
		return IotConnectorsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - iotConnectorName - The name of IoT Connector resource.
//   - workspaceName - The name of workspace resource.
//   - iotConnectorPatchResource - The parameters for updating an IoT Connector.
//   - options - IotConnectorsClientBeginUpdateOptions contains the optional parameters for the IotConnectorsClient.BeginUpdate
//     method.
func (client *IotConnectorsClient) BeginUpdate(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsClientBeginUpdateOptions) (*runtime.Poller[IotConnectorsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, iotConnectorName, workspaceName, iotConnectorPatchResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IotConnectorsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IotConnectorsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch an IoT Connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-31
func (client *IotConnectorsClient) update(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "IotConnectorsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, iotConnectorName, workspaceName, iotConnectorPatchResource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *IotConnectorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource IotConnectorPatchResource, options *IotConnectorsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/iotconnectors/{iotConnectorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if iotConnectorName == "" {
		return nil, errors.New("parameter iotConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotConnectorName}", url.PathEscape(iotConnectorName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-31")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, iotConnectorPatchResource); err != nil {
		return nil, err
	}
	return req, nil
}
