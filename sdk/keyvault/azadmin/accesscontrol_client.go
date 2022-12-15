//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azadmin

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccessControlClient contains the methods for the AccessControl group.
// Don't use this type directly, use NewAccessControlClient() instead.
type AccessControlClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// CreateRoleAssignment - Creates a role assignment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.4-preview.1
// scope - The scope of the role assignment to create.
// roleAssignmentName - The name of the role assignment to create. It can be any valid GUID.
// parameters - Parameters for the role assignment.
// options - AccessControlClientCreateRoleAssignmentOptions contains the optional parameters for the AccessControlClient.CreateRoleAssignment
// method.
func (client *AccessControlClient) CreateRoleAssignment(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *AccessControlClientCreateRoleAssignmentOptions) (AccessControlClientCreateRoleAssignmentResponse, error) {
	req, err := client.createRoleAssignmentCreateRequest(ctx, scope, roleAssignmentName, parameters, options)
	if err != nil {
		return AccessControlClientCreateRoleAssignmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessControlClientCreateRoleAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return AccessControlClientCreateRoleAssignmentResponse{}, runtime.NewResponseError(resp)
	}
	return client.createRoleAssignmentHandleResponse(resp)
}

// createRoleAssignmentCreateRequest creates the CreateRoleAssignment request.
func (client *AccessControlClient) createRoleAssignmentCreateRequest(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *AccessControlClientCreateRoleAssignmentOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createRoleAssignmentHandleResponse handles the CreateRoleAssignment response.
func (client *AccessControlClient) createRoleAssignmentHandleResponse(resp *http.Response) (AccessControlClientCreateRoleAssignmentResponse, error) {
	result := AccessControlClientCreateRoleAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return AccessControlClientCreateRoleAssignmentResponse{}, err
	}
	return result, nil
}

// DeleteRoleAssignment - Deletes a role assignment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.4-preview.1
// scope - The scope of the role assignment to delete.
// roleAssignmentName - The name of the role assignment to delete.
// options - AccessControlClientDeleteRoleAssignmentOptions contains the optional parameters for the AccessControlClient.DeleteRoleAssignment
// method.
func (client *AccessControlClient) DeleteRoleAssignment(ctx context.Context, scope string, roleAssignmentName string, options *AccessControlClientDeleteRoleAssignmentOptions) (AccessControlClientDeleteRoleAssignmentResponse, error) {
	req, err := client.deleteRoleAssignmentCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return AccessControlClientDeleteRoleAssignmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessControlClientDeleteRoleAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessControlClientDeleteRoleAssignmentResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteRoleAssignmentHandleResponse(resp)
}

// deleteRoleAssignmentCreateRequest creates the DeleteRoleAssignment request.
func (client *AccessControlClient) deleteRoleAssignmentCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *AccessControlClientDeleteRoleAssignmentOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteRoleAssignmentHandleResponse handles the DeleteRoleAssignment response.
func (client *AccessControlClient) deleteRoleAssignmentHandleResponse(resp *http.Response) (AccessControlClientDeleteRoleAssignmentResponse, error) {
	result := AccessControlClientDeleteRoleAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return AccessControlClientDeleteRoleAssignmentResponse{}, err
	}
	return result, nil
}

// DeleteRoleDefinition - Deletes a custom role definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.4-preview.1
// scope - The scope of the role definition to delete. Managed HSM only supports '/'.
// roleDefinitionName - The name (GUID) of the role definition to delete.
// options - AccessControlClientDeleteRoleDefinitionOptions contains the optional parameters for the AccessControlClient.DeleteRoleDefinition
// method.
func (client *AccessControlClient) DeleteRoleDefinition(ctx context.Context, scope string, roleDefinitionName string, options *AccessControlClientDeleteRoleDefinitionOptions) (AccessControlClientDeleteRoleDefinitionResponse, error) {
	req, err := client.deleteRoleDefinitionCreateRequest(ctx, scope, roleDefinitionName, options)
	if err != nil {
		return AccessControlClientDeleteRoleDefinitionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessControlClientDeleteRoleDefinitionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessControlClientDeleteRoleDefinitionResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteRoleDefinitionHandleResponse(resp)
}

// deleteRoleDefinitionCreateRequest creates the DeleteRoleDefinition request.
func (client *AccessControlClient) deleteRoleDefinitionCreateRequest(ctx context.Context, scope string, roleDefinitionName string, options *AccessControlClientDeleteRoleDefinitionOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteRoleDefinitionHandleResponse handles the DeleteRoleDefinition response.
func (client *AccessControlClient) deleteRoleDefinitionHandleResponse(resp *http.Response) (AccessControlClientDeleteRoleDefinitionResponse, error) {
	result := AccessControlClientDeleteRoleDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return AccessControlClientDeleteRoleDefinitionResponse{}, err
	}
	return result, nil
}

// GetRoleAssignment - Get the specified role assignment.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.4-preview.1
// scope - The scope of the role assignment.
// roleAssignmentName - The name of the role assignment to get.
// options - AccessControlClientGetRoleAssignmentOptions contains the optional parameters for the AccessControlClient.GetRoleAssignment
// method.
func (client *AccessControlClient) GetRoleAssignment(ctx context.Context, scope string, roleAssignmentName string, options *AccessControlClientGetRoleAssignmentOptions) (AccessControlClientGetRoleAssignmentResponse, error) {
	req, err := client.getRoleAssignmentCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return AccessControlClientGetRoleAssignmentResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessControlClientGetRoleAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessControlClientGetRoleAssignmentResponse{}, runtime.NewResponseError(resp)
	}
	return client.getRoleAssignmentHandleResponse(resp)
}

// getRoleAssignmentCreateRequest creates the GetRoleAssignment request.
func (client *AccessControlClient) getRoleAssignmentCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *AccessControlClientGetRoleAssignmentOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", url.PathEscape(roleAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRoleAssignmentHandleResponse handles the GetRoleAssignment response.
func (client *AccessControlClient) getRoleAssignmentHandleResponse(resp *http.Response) (AccessControlClientGetRoleAssignmentResponse, error) {
	result := AccessControlClientGetRoleAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignment); err != nil {
		return AccessControlClientGetRoleAssignmentResponse{}, err
	}
	return result, nil
}

// GetRoleDefinition - Get the specified role definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.4-preview.1
// scope - The scope of the role definition to get. Managed HSM only supports '/'.
// roleDefinitionName - The name of the role definition to get.
// options - AccessControlClientGetRoleDefinitionOptions contains the optional parameters for the AccessControlClient.GetRoleDefinition
// method.
func (client *AccessControlClient) GetRoleDefinition(ctx context.Context, scope string, roleDefinitionName string, options *AccessControlClientGetRoleDefinitionOptions) (AccessControlClientGetRoleDefinitionResponse, error) {
	req, err := client.getRoleDefinitionCreateRequest(ctx, scope, roleDefinitionName, options)
	if err != nil {
		return AccessControlClientGetRoleDefinitionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessControlClientGetRoleDefinitionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessControlClientGetRoleDefinitionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getRoleDefinitionHandleResponse(resp)
}

// getRoleDefinitionCreateRequest creates the GetRoleDefinition request.
func (client *AccessControlClient) getRoleDefinitionCreateRequest(ctx context.Context, scope string, roleDefinitionName string, options *AccessControlClientGetRoleDefinitionOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getRoleDefinitionHandleResponse handles the GetRoleDefinition response.
func (client *AccessControlClient) getRoleDefinitionHandleResponse(resp *http.Response) (AccessControlClientGetRoleDefinitionResponse, error) {
	result := AccessControlClientGetRoleDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return AccessControlClientGetRoleDefinitionResponse{}, err
	}
	return result, nil
}

// NewListRoleAssignmentsPager - Gets role assignments for a scope.
// Generated from API version 7.4-preview.1
// scope - The scope of the role assignments.
// options - AccessControlClientListRoleAssignmentsOptions contains the optional parameters for the AccessControlClient.ListRoleAssignments
// method.
func (client *AccessControlClient) NewListRoleAssignmentsPager(scope string, options *AccessControlClientListRoleAssignmentsOptions) *runtime.Pager[AccessControlClientListRoleAssignmentsResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessControlClientListRoleAssignmentsResponse]{
		More: func(page AccessControlClientListRoleAssignmentsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessControlClientListRoleAssignmentsResponse) (AccessControlClientListRoleAssignmentsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listRoleAssignmentsCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessControlClientListRoleAssignmentsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccessControlClientListRoleAssignmentsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessControlClientListRoleAssignmentsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listRoleAssignmentsHandleResponse(resp)
		},
	})
}

// listRoleAssignmentsCreateRequest creates the ListRoleAssignments request.
func (client *AccessControlClient) listRoleAssignmentsCreateRequest(ctx context.Context, scope string, options *AccessControlClientListRoleAssignmentsOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRoleAssignmentsHandleResponse handles the ListRoleAssignments response.
func (client *AccessControlClient) listRoleAssignmentsHandleResponse(resp *http.Response) (AccessControlClientListRoleAssignmentsResponse, error) {
	result := AccessControlClientListRoleAssignmentsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentListResult); err != nil {
		return AccessControlClientListRoleAssignmentsResponse{}, err
	}
	return result, nil
}

// NewListRoleDefinitionsPager - Get all role definitions that are applicable at scope and above.
// Generated from API version 7.4-preview.1
// scope - The scope of the role definition.
// options - AccessControlClientListRoleDefinitionsOptions contains the optional parameters for the AccessControlClient.ListRoleDefinitions
// method.
func (client *AccessControlClient) NewListRoleDefinitionsPager(scope string, options *AccessControlClientListRoleDefinitionsOptions) *runtime.Pager[AccessControlClientListRoleDefinitionsResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessControlClientListRoleDefinitionsResponse]{
		More: func(page AccessControlClientListRoleDefinitionsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessControlClientListRoleDefinitionsResponse) (AccessControlClientListRoleDefinitionsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listRoleDefinitionsCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AccessControlClientListRoleDefinitionsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AccessControlClientListRoleDefinitionsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AccessControlClientListRoleDefinitionsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listRoleDefinitionsHandleResponse(resp)
		},
	})
}

// listRoleDefinitionsCreateRequest creates the ListRoleDefinitions request.
func (client *AccessControlClient) listRoleDefinitionsCreateRequest(ctx context.Context, scope string, options *AccessControlClientListRoleDefinitionsOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRoleDefinitionsHandleResponse handles the ListRoleDefinitions response.
func (client *AccessControlClient) listRoleDefinitionsHandleResponse(resp *http.Response) (AccessControlClientListRoleDefinitionsResponse, error) {
	result := AccessControlClientListRoleDefinitionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinitionListResult); err != nil {
		return AccessControlClientListRoleDefinitionsResponse{}, err
	}
	return result, nil
}

// SetRoleDefinition - Creates or updates a custom role definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 7.4-preview.1
// scope - The scope of the role definition to create or update. Managed HSM only supports '/'.
// roleDefinitionName - The name of the role definition to create or update. It can be any valid GUID.
// parameters - Parameters for the role definition.
// options - AccessControlClientSetRoleDefinitionOptions contains the optional parameters for the AccessControlClient.SetRoleDefinition
// method.
func (client *AccessControlClient) SetRoleDefinition(ctx context.Context, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters, options *AccessControlClientSetRoleDefinitionOptions) (AccessControlClientSetRoleDefinitionResponse, error) {
	req, err := client.setRoleDefinitionCreateRequest(ctx, scope, roleDefinitionName, parameters, options)
	if err != nil {
		return AccessControlClientSetRoleDefinitionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessControlClientSetRoleDefinitionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return AccessControlClientSetRoleDefinitionResponse{}, runtime.NewResponseError(resp)
	}
	return client.setRoleDefinitionHandleResponse(resp)
}

// setRoleDefinitionCreateRequest creates the SetRoleDefinition request.
func (client *AccessControlClient) setRoleDefinitionCreateRequest(ctx context.Context, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters, options *AccessControlClientSetRoleDefinitionOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleDefinitionName == "" {
		return nil, errors.New("parameter roleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleDefinitionName}", url.PathEscape(roleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.4-preview.1")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// setRoleDefinitionHandleResponse handles the SetRoleDefinition response.
func (client *AccessControlClient) setRoleDefinitionHandleResponse(resp *http.Response) (AccessControlClientSetRoleDefinitionResponse, error) {
	result := AccessControlClientSetRoleDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleDefinition); err != nil {
		return AccessControlClientSetRoleDefinitionResponse{}, err
	}
	return result, nil
}
