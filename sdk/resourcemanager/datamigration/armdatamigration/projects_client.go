//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatamigration

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ProjectsClient contains the methods for the Projects group.
// Don't use this type directly, use NewProjectsClient() instead.
type ProjectsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProjectsClient creates a new instance of ProjectsClient with the specified values.
//   - subscriptionID - Identifier of the subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProjectsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProjectsClient, error) {
	cl, err := arm.NewClient(moduleName+".ProjectsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProjectsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - The project resource is a nested resource representing a stored migration project. The PUT method creates
// a new project or updates an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-30
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - projectName - Name of the project
//   - parameters - Information about the project
//   - options - ProjectsClientCreateOrUpdateOptions contains the optional parameters for the ProjectsClient.CreateOrUpdate method.
func (client *ProjectsClient) CreateOrUpdate(ctx context.Context, groupName string, serviceName string, projectName string, parameters Project, options *ProjectsClientCreateOrUpdateOptions) (ProjectsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, groupName, serviceName, projectName, parameters, options)
	if err != nil {
		return ProjectsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProjectsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProjectsClient) createOrUpdateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, parameters Project, options *ProjectsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProjectsClient) createOrUpdateHandleResponse(resp *http.Response) (ProjectsClientCreateOrUpdateResponse, error) {
	result := ProjectsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The project resource is a nested resource representing a stored migration project. The DELETE method deletes a
// project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-30
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - projectName - Name of the project
//   - options - ProjectsClientDeleteOptions contains the optional parameters for the ProjectsClient.Delete method.
func (client *ProjectsClient) Delete(ctx context.Context, groupName string, serviceName string, projectName string, options *ProjectsClientDeleteOptions) (ProjectsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, groupName, serviceName, projectName, options)
	if err != nil {
		return ProjectsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProjectsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProjectsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProjectsClient) deleteCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, options *ProjectsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	if options != nil && options.DeleteRunningTasks != nil {
		reqQP.Set("deleteRunningTasks", strconv.FormatBool(*options.DeleteRunningTasks))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The project resource is a nested resource representing a stored migration project. The GET method retrieves information
// about a project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-30
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - projectName - Name of the project
//   - options - ProjectsClientGetOptions contains the optional parameters for the ProjectsClient.Get method.
func (client *ProjectsClient) Get(ctx context.Context, groupName string, serviceName string, projectName string, options *ProjectsClientGetOptions) (ProjectsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, groupName, serviceName, projectName, options)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProjectsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProjectsClient) getCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, options *ProjectsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProjectsClient) getHandleResponse(resp *http.Response) (ProjectsClientGetResponse, error) {
	result := ProjectsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The project resource is a nested resource representing a stored migration project. This method returns a
// list of projects owned by a service resource.
//
// Generated from API version 2021-06-30
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - options - ProjectsClientListOptions contains the optional parameters for the ProjectsClient.NewListPager method.
func (client *ProjectsClient) NewListPager(groupName string, serviceName string, options *ProjectsClientListOptions) *runtime.Pager[ProjectsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProjectsClientListResponse]{
		More: func(page ProjectsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProjectsClientListResponse) (ProjectsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, groupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProjectsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProjectsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProjectsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ProjectsClient) listCreateRequest(ctx context.Context, groupName string, serviceName string, options *ProjectsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProjectsClient) listHandleResponse(resp *http.Response) (ProjectsClientListResponse, error) {
	result := ProjectsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectList); err != nil {
		return ProjectsClientListResponse{}, err
	}
	return result, nil
}

// Update - The project resource is a nested resource representing a stored migration project. The PATCH method updates an
// existing project.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-30
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - projectName - Name of the project
//   - parameters - Information about the project
//   - options - ProjectsClientUpdateOptions contains the optional parameters for the ProjectsClient.Update method.
func (client *ProjectsClient) Update(ctx context.Context, groupName string, serviceName string, projectName string, parameters Project, options *ProjectsClientUpdateOptions) (ProjectsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, groupName, serviceName, projectName, parameters, options)
	if err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProjectsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProjectsClient) updateCreateRequest(ctx context.Context, groupName string, serviceName string, projectName string, parameters Project, options *ProjectsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ProjectsClient) updateHandleResponse(resp *http.Response) (ProjectsClientUpdateResponse, error) {
	result := ProjectsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Project); err != nil {
		return ProjectsClientUpdateResponse{}, err
	}
	return result, nil
}
