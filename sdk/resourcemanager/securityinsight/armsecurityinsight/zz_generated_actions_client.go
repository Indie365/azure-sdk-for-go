//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsight

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ActionsClient contains the methods for the Actions group.
// Don't use this type directly, use NewActionsClient() instead.
type ActionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewActionsClient creates a new instance of ActionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewActionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ActionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ActionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates the action of alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// ruleID - Alert rule ID
// actionID - Action ID
// action - The action
// options - ActionsClientCreateOrUpdateOptions contains the optional parameters for the ActionsClient.CreateOrUpdate method.
func (client *ActionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, action ActionRequest, options *ActionsClientCreateOrUpdateOptions) (ActionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, actionID, action, options)
	if err != nil {
		return ActionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ActionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ActionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, action ActionRequest, options *ActionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}"
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
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	if actionID == "" {
		return nil, errors.New("parameter actionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionId}", url.PathEscape(actionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, action)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ActionsClient) createOrUpdateHandleResponse(resp *http.Response) (ActionsClientCreateOrUpdateResponse, error) {
	result := ActionsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActionResponse); err != nil {
		return ActionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the action of alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// ruleID - Alert rule ID
// actionID - Action ID
// options - ActionsClientDeleteOptions contains the optional parameters for the ActionsClient.Delete method.
func (client *ActionsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, options *ActionsClientDeleteOptions) (ActionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, actionID, options)
	if err != nil {
		return ActionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ActionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ActionsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ActionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, options *ActionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}"
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
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	if actionID == "" {
		return nil, errors.New("parameter actionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionId}", url.PathEscape(actionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the action of alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// ruleID - Alert rule ID
// actionID - Action ID
// options - ActionsClientGetOptions contains the optional parameters for the ActionsClient.Get method.
func (client *ActionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, options *ActionsClientGetOptions) (ActionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, actionID, options)
	if err != nil {
		return ActionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ActionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, options *ActionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}"
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
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	if actionID == "" {
		return nil, errors.New("parameter actionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{actionId}", url.PathEscape(actionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActionsClient) getHandleResponse(resp *http.Response) (ActionsClientGetResponse, error) {
	result := ActionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActionResponse); err != nil {
		return ActionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByAlertRule - Gets all actions of alert rule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// ruleID - Alert rule ID
// options - ActionsClientListByAlertRuleOptions contains the optional parameters for the ActionsClient.ListByAlertRule method.
func (client *ActionsClient) ListByAlertRule(resourceGroupName string, workspaceName string, ruleID string, options *ActionsClientListByAlertRuleOptions) *ActionsClientListByAlertRulePager {
	return &ActionsClientListByAlertRulePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAlertRuleCreateRequest(ctx, resourceGroupName, workspaceName, ruleID, options)
		},
		advancer: func(ctx context.Context, resp ActionsClientListByAlertRuleResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ActionsList.NextLink)
		},
	}
}

// listByAlertRuleCreateRequest creates the ListByAlertRule request.
func (client *ActionsClient) listByAlertRuleCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, options *ActionsClientListByAlertRuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions"
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
	if ruleID == "" {
		return nil, errors.New("parameter ruleID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleId}", url.PathEscape(ruleID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAlertRuleHandleResponse handles the ListByAlertRule response.
func (client *ActionsClient) listByAlertRuleHandleResponse(resp *http.Response) (ActionsClientListByAlertRuleResponse, error) {
	result := ActionsClientListByAlertRuleResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActionsList); err != nil {
		return ActionsClientListByAlertRuleResponse{}, err
	}
	return result, nil
}
