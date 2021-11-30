//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PolicyDefinitionsClient contains the methods for the PolicyDefinitions group.
// Don't use this type directly, use NewPolicyDefinitionsClient() instead.
type PolicyDefinitionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPolicyDefinitionsClient creates a new instance of PolicyDefinitionsClient with the specified values.
func NewPolicyDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PolicyDefinitionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PolicyDefinitionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - This operation creates or updates a policy definition in the given subscription with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters PolicyDefinition, options *PolicyDefinitionsCreateOrUpdateOptions) (PolicyDefinitionsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, policyDefinitionName, parameters, options)
	if err != nil {
		return PolicyDefinitionsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyDefinitionsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return PolicyDefinitionsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PolicyDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, policyDefinitionName string, parameters PolicyDefinition, options *PolicyDefinitionsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PolicyDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (PolicyDefinitionsCreateOrUpdateResponse, error) {
	result := PolicyDefinitionsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinition); err != nil {
		return PolicyDefinitionsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PolicyDefinitionsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// CreateOrUpdateAtManagementGroup - This operation creates or updates a policy definition in the given management group with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, parameters PolicyDefinition, options *PolicyDefinitionsCreateOrUpdateAtManagementGroupOptions) (PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse, error) {
	req, err := client.createOrUpdateAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, parameters, options)
	if err != nil {
		return PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse{}, client.createOrUpdateAtManagementGroupHandleError(resp)
	}
	return client.createOrUpdateAtManagementGroupHandleResponse(resp)
}

// createOrUpdateAtManagementGroupCreateRequest creates the CreateOrUpdateAtManagementGroup request.
func (client *PolicyDefinitionsClient) createOrUpdateAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, parameters PolicyDefinition, options *PolicyDefinitionsCreateOrUpdateAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateAtManagementGroupHandleResponse handles the CreateOrUpdateAtManagementGroup response.
func (client *PolicyDefinitionsClient) createOrUpdateAtManagementGroupHandleResponse(resp *http.Response) (PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse, error) {
	result := PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinition); err != nil {
		return PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateAtManagementGroupHandleError handles the CreateOrUpdateAtManagementGroup error response.
func (client *PolicyDefinitionsClient) createOrUpdateAtManagementGroupHandleError(resp *http.Response) error {
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

// Delete - This operation deletes the policy definition in the given subscription with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) Delete(ctx context.Context, policyDefinitionName string, options *PolicyDefinitionsDeleteOptions) (PolicyDefinitionsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return PolicyDefinitionsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyDefinitionsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PolicyDefinitionsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return PolicyDefinitionsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PolicyDefinitionsClient) deleteCreateRequest(ctx context.Context, policyDefinitionName string, options *PolicyDefinitionsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PolicyDefinitionsClient) deleteHandleError(resp *http.Response) error {
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

// DeleteAtManagementGroup - This operation deletes the policy definition in the given management group with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) DeleteAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, options *PolicyDefinitionsDeleteAtManagementGroupOptions) (PolicyDefinitionsDeleteAtManagementGroupResponse, error) {
	req, err := client.deleteAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, options)
	if err != nil {
		return PolicyDefinitionsDeleteAtManagementGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyDefinitionsDeleteAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PolicyDefinitionsDeleteAtManagementGroupResponse{}, client.deleteAtManagementGroupHandleError(resp)
	}
	return PolicyDefinitionsDeleteAtManagementGroupResponse{RawResponse: resp}, nil
}

// deleteAtManagementGroupCreateRequest creates the DeleteAtManagementGroup request.
func (client *PolicyDefinitionsClient) deleteAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, options *PolicyDefinitionsDeleteAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAtManagementGroupHandleError handles the DeleteAtManagementGroup error response.
func (client *PolicyDefinitionsClient) deleteAtManagementGroupHandleError(resp *http.Response) error {
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

// Get - This operation retrieves the policy definition in the given subscription with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) Get(ctx context.Context, policyDefinitionName string, options *PolicyDefinitionsGetOptions) (PolicyDefinitionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return PolicyDefinitionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyDefinitionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyDefinitionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PolicyDefinitionsClient) getCreateRequest(ctx context.Context, policyDefinitionName string, options *PolicyDefinitionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PolicyDefinitionsClient) getHandleResponse(resp *http.Response) (PolicyDefinitionsGetResponse, error) {
	result := PolicyDefinitionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinition); err != nil {
		return PolicyDefinitionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PolicyDefinitionsClient) getHandleError(resp *http.Response) error {
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

// GetAtManagementGroup - This operation retrieves the policy definition in the given management group with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) GetAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, options *PolicyDefinitionsGetAtManagementGroupOptions) (PolicyDefinitionsGetAtManagementGroupResponse, error) {
	req, err := client.getAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, options)
	if err != nil {
		return PolicyDefinitionsGetAtManagementGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyDefinitionsGetAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyDefinitionsGetAtManagementGroupResponse{}, client.getAtManagementGroupHandleError(resp)
	}
	return client.getAtManagementGroupHandleResponse(resp)
}

// getAtManagementGroupCreateRequest creates the GetAtManagementGroup request.
func (client *PolicyDefinitionsClient) getAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, options *PolicyDefinitionsGetAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getAtManagementGroupHandleResponse handles the GetAtManagementGroup response.
func (client *PolicyDefinitionsClient) getAtManagementGroupHandleResponse(resp *http.Response) (PolicyDefinitionsGetAtManagementGroupResponse, error) {
	result := PolicyDefinitionsGetAtManagementGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinition); err != nil {
		return PolicyDefinitionsGetAtManagementGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getAtManagementGroupHandleError handles the GetAtManagementGroup error response.
func (client *PolicyDefinitionsClient) getAtManagementGroupHandleError(resp *http.Response) error {
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

// GetBuiltIn - This operation retrieves the built-in policy definition with the given name.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) GetBuiltIn(ctx context.Context, policyDefinitionName string, options *PolicyDefinitionsGetBuiltInOptions) (PolicyDefinitionsGetBuiltInResponse, error) {
	req, err := client.getBuiltInCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return PolicyDefinitionsGetBuiltInResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyDefinitionsGetBuiltInResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyDefinitionsGetBuiltInResponse{}, client.getBuiltInHandleError(resp)
	}
	return client.getBuiltInHandleResponse(resp)
}

// getBuiltInCreateRequest creates the GetBuiltIn request.
func (client *PolicyDefinitionsClient) getBuiltInCreateRequest(ctx context.Context, policyDefinitionName string, options *PolicyDefinitionsGetBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getBuiltInHandleResponse handles the GetBuiltIn response.
func (client *PolicyDefinitionsClient) getBuiltInHandleResponse(resp *http.Response) (PolicyDefinitionsGetBuiltInResponse, error) {
	result := PolicyDefinitionsGetBuiltInResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinition); err != nil {
		return PolicyDefinitionsGetBuiltInResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getBuiltInHandleError handles the GetBuiltIn error response.
func (client *PolicyDefinitionsClient) getBuiltInHandleError(resp *http.Response) error {
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

// List - This operation retrieves a list of all the policy definitions in a given subscription that match the optional given $filter. Valid values for
// $filter are: 'atExactScope()', 'policyType -eq {value}' or
// 'category eq '{value}''. If $filter is not provided, the unfiltered list includes all policy definitions associated with the subscription, including
// those that apply directly or from management groups
// that contain the given subscription. If $filter=atExactScope() is provided, the returned list only includes all policy definitions that at the given
// subscription. If $filter='policyType -eq {value}'
// is provided, the returned list only includes all policy definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn,
// Custom, and Static. If $filter='category -eq
// {value}' is provided, the returned list only includes all policy definitions whose category match the {value}.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) List(options *PolicyDefinitionsListOptions) *PolicyDefinitionsListPager {
	return &PolicyDefinitionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PolicyDefinitionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyDefinitionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PolicyDefinitionsClient) listCreateRequest(ctx context.Context, options *PolicyDefinitionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PolicyDefinitionsClient) listHandleResponse(resp *http.Response) (PolicyDefinitionsListResponse, error) {
	result := PolicyDefinitionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinitionListResult); err != nil {
		return PolicyDefinitionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PolicyDefinitionsClient) listHandleError(resp *http.Response) error {
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

// ListBuiltIn - This operation retrieves a list of all the built-in policy definitions that match the optional given $filter. If $filter='policyType -eq
// {value}' is provided, the returned list only includes all
// built-in policy definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom, and Static. If $filter='category
// -eq {value}' is provided, the returned list
// only includes all built-in policy definitions whose category match the {value}.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) ListBuiltIn(options *PolicyDefinitionsListBuiltInOptions) *PolicyDefinitionsListBuiltInPager {
	return &PolicyDefinitionsListBuiltInPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBuiltInCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PolicyDefinitionsListBuiltInResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyDefinitionListResult.NextLink)
		},
	}
}

// listBuiltInCreateRequest creates the ListBuiltIn request.
func (client *PolicyDefinitionsClient) listBuiltInCreateRequest(ctx context.Context, options *PolicyDefinitionsListBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policyDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBuiltInHandleResponse handles the ListBuiltIn response.
func (client *PolicyDefinitionsClient) listBuiltInHandleResponse(resp *http.Response) (PolicyDefinitionsListBuiltInResponse, error) {
	result := PolicyDefinitionsListBuiltInResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinitionListResult); err != nil {
		return PolicyDefinitionsListBuiltInResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBuiltInHandleError handles the ListBuiltIn error response.
func (client *PolicyDefinitionsClient) listBuiltInHandleError(resp *http.Response) error {
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

// ListByManagementGroup - This operation retrieves a list of all the policy definitions in a given management group that match the optional given $filter.
// Valid values for $filter are: 'atExactScope()', 'policyType -eq
// {value}' or 'category eq '{value}''. If $filter is not provided, the unfiltered list includes all policy definitions associated with the management group,
// including those that apply directly or from
// management groups that contain the given management group. If $filter=atExactScope() is provided, the returned list only includes all policy definitions
// that at the given management group. If
// $filter='policyType -eq {value}' is provided, the returned list only includes all policy definitions whose type match the {value}. Possible policyType
// values are NotSpecified, BuiltIn, Custom, and
// Static. If $filter='category -eq {value}' is provided, the returned list only includes all policy definitions whose category match the {value}.
// If the operation fails it returns the *CloudError error type.
func (client *PolicyDefinitionsClient) ListByManagementGroup(managementGroupID string, options *PolicyDefinitionsListByManagementGroupOptions) *PolicyDefinitionsListByManagementGroupPager {
	return &PolicyDefinitionsListByManagementGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByManagementGroupCreateRequest(ctx, managementGroupID, options)
		},
		advancer: func(ctx context.Context, resp PolicyDefinitionsListByManagementGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PolicyDefinitionListResult.NextLink)
		},
	}
}

// listByManagementGroupCreateRequest creates the ListByManagementGroup request.
func (client *PolicyDefinitionsClient) listByManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *PolicyDefinitionsListByManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByManagementGroupHandleResponse handles the ListByManagementGroup response.
func (client *PolicyDefinitionsClient) listByManagementGroupHandleResponse(resp *http.Response) (PolicyDefinitionsListByManagementGroupResponse, error) {
	result := PolicyDefinitionsListByManagementGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyDefinitionListResult); err != nil {
		return PolicyDefinitionsListByManagementGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByManagementGroupHandleError handles the ListByManagementGroup error response.
func (client *PolicyDefinitionsClient) listByManagementGroupHandleError(resp *http.Response) error {
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
