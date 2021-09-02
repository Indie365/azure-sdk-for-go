//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RoleEligibilitySchedulesClient contains the methods for the RoleEligibilitySchedules group.
// Don't use this type directly, use NewRoleEligibilitySchedulesClient() instead.
type RoleEligibilitySchedulesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRoleEligibilitySchedulesClient creates a new instance of RoleEligibilitySchedulesClient with the specified values.
func NewRoleEligibilitySchedulesClient(con *arm.Connection) *RoleEligibilitySchedulesClient {
	return &RoleEligibilitySchedulesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Get the specified role eligibility schedule for a resource scope
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilitySchedulesClient) Get(ctx context.Context, scope string, roleEligibilityScheduleName string, options *RoleEligibilitySchedulesGetOptions) (RoleEligibilitySchedulesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleEligibilityScheduleName, options)
	if err != nil {
		return RoleEligibilitySchedulesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilitySchedulesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilitySchedulesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleEligibilitySchedulesClient) getCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleName string, options *RoleEligibilitySchedulesGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilitySchedules/{roleEligibilityScheduleName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleName}", url.PathEscape(roleEligibilityScheduleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleEligibilitySchedulesClient) getHandleResponse(resp *http.Response) (RoleEligibilitySchedulesGetResponse, error) {
	result := RoleEligibilitySchedulesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilitySchedule); err != nil {
		return RoleEligibilitySchedulesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleEligibilitySchedulesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListForScope - Gets role eligibility schedules for a resource scope.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilitySchedulesClient) ListForScope(scope string, options *RoleEligibilitySchedulesListForScopeOptions) *RoleEligibilitySchedulesListForScopePager {
	return &RoleEligibilitySchedulesListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleEligibilitySchedulesListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleEligibilityScheduleListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleEligibilitySchedulesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleEligibilitySchedulesListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilitySchedules"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleEligibilitySchedulesClient) listForScopeHandleResponse(resp *http.Response) (RoleEligibilitySchedulesListForScopeResponse, error) {
	result := RoleEligibilitySchedulesListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleListResult); err != nil {
		return RoleEligibilitySchedulesListForScopeResponse{}, err
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleEligibilitySchedulesClient) listForScopeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
