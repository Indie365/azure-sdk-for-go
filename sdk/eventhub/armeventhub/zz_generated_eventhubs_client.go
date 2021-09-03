// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// EventHubsClient contains the methods for the EventHubs group.
// Don't use this type directly, use NewEventHubsClient() instead.
type EventHubsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewEventHubsClient creates a new instance of EventHubsClient with the specified values.
func NewEventHubsClient(con *armcore.Connection, subscriptionID string) *EventHubsClient {
	return &EventHubsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a new Event Hub as a nested resource within a Namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters Eventhub, options *EventHubsCreateOrUpdateOptions) (EventHubsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, parameters, options)
	if err != nil {
		return EventHubsCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EventHubsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EventHubsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters Eventhub, options *EventHubsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EventHubsClient) createOrUpdateHandleResponse(resp *azcore.Response) (EventHubsCreateOrUpdateResponse, error) {
	result := EventHubsCreateOrUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Eventhub); err != nil {
		return EventHubsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *EventHubsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CreateOrUpdateAuthorizationRule - Creates or updates an AuthorizationRule for the specified Event Hub. Creation/update of the AuthorizationRule will
// take a few seconds to take effect.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters AuthorizationRule, options *EventHubsCreateOrUpdateAuthorizationRuleOptions) (EventHubsCreateOrUpdateAuthorizationRuleResponse, error) {
	req, err := client.createOrUpdateAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, parameters, options)
	if err != nil {
		return EventHubsCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EventHubsCreateOrUpdateAuthorizationRuleResponse{}, client.createOrUpdateAuthorizationRuleHandleError(resp)
	}
	return client.createOrUpdateAuthorizationRuleHandleResponse(resp)
}

// createOrUpdateAuthorizationRuleCreateRequest creates the CreateOrUpdateAuthorizationRule request.
func (client *EventHubsClient) createOrUpdateAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters AuthorizationRule, options *EventHubsCreateOrUpdateAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateAuthorizationRuleHandleResponse handles the CreateOrUpdateAuthorizationRule response.
func (client *EventHubsClient) createOrUpdateAuthorizationRuleHandleResponse(resp *azcore.Response) (EventHubsCreateOrUpdateAuthorizationRuleResponse, error) {
	result := EventHubsCreateOrUpdateAuthorizationRuleResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AuthorizationRule); err != nil {
		return EventHubsCreateOrUpdateAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// createOrUpdateAuthorizationRuleHandleError handles the CreateOrUpdateAuthorizationRule error response.
func (client *EventHubsClient) createOrUpdateAuthorizationRuleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes an Event Hub from the specified Namespace and resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsDeleteOptions) (EventHubsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, options)
	if err != nil {
		return EventHubsDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return EventHubsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return EventHubsDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EventHubsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *EventHubsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// DeleteAuthorizationRule - Deletes an Event Hub AuthorizationRule.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsDeleteAuthorizationRuleOptions) (EventHubsDeleteAuthorizationRuleResponse, error) {
	req, err := client.deleteAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, options)
	if err != nil {
		return EventHubsDeleteAuthorizationRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsDeleteAuthorizationRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return EventHubsDeleteAuthorizationRuleResponse{}, client.deleteAuthorizationRuleHandleError(resp)
	}
	return EventHubsDeleteAuthorizationRuleResponse{RawResponse: resp.Response}, nil
}

// deleteAuthorizationRuleCreateRequest creates the DeleteAuthorizationRule request.
func (client *EventHubsClient) deleteAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsDeleteAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAuthorizationRuleHandleError handles the DeleteAuthorizationRule error response.
func (client *EventHubsClient) deleteAuthorizationRuleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets an Event Hubs description for the specified Event Hub.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsGetOptions) (EventHubsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, options)
	if err != nil {
		return EventHubsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EventHubsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EventHubsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EventHubsClient) getHandleResponse(resp *azcore.Response) (EventHubsGetResponse, error) {
	result := EventHubsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Eventhub); err != nil {
		return EventHubsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *EventHubsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetAuthorizationRule - Gets an AuthorizationRule for an Event Hub by rule name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsGetAuthorizationRuleOptions) (EventHubsGetAuthorizationRuleResponse, error) {
	req, err := client.getAuthorizationRuleCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, options)
	if err != nil {
		return EventHubsGetAuthorizationRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsGetAuthorizationRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EventHubsGetAuthorizationRuleResponse{}, client.getAuthorizationRuleHandleError(resp)
	}
	return client.getAuthorizationRuleHandleResponse(resp)
}

// getAuthorizationRuleCreateRequest creates the GetAuthorizationRule request.
func (client *EventHubsClient) getAuthorizationRuleCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsGetAuthorizationRuleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getAuthorizationRuleHandleResponse handles the GetAuthorizationRule response.
func (client *EventHubsClient) getAuthorizationRuleHandleResponse(resp *azcore.Response) (EventHubsGetAuthorizationRuleResponse, error) {
	result := EventHubsGetAuthorizationRuleResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AuthorizationRule); err != nil {
		return EventHubsGetAuthorizationRuleResponse{}, err
	}
	return result, nil
}

// getAuthorizationRuleHandleError handles the GetAuthorizationRule error response.
func (client *EventHubsClient) getAuthorizationRuleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListAuthorizationRules - Gets the authorization rules for an Event Hub.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsListAuthorizationRulesOptions) EventHubsListAuthorizationRulesPager {
	return &eventHubsListAuthorizationRulesPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAuthorizationRulesCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, options)
		},
		advancer: func(ctx context.Context, resp EventHubsListAuthorizationRulesResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AuthorizationRuleListResult.NextLink)
		},
	}
}

// listAuthorizationRulesCreateRequest creates the ListAuthorizationRules request.
func (client *EventHubsClient) listAuthorizationRulesCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *EventHubsListAuthorizationRulesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAuthorizationRulesHandleResponse handles the ListAuthorizationRules response.
func (client *EventHubsClient) listAuthorizationRulesHandleResponse(resp *azcore.Response) (EventHubsListAuthorizationRulesResponse, error) {
	result := EventHubsListAuthorizationRulesResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AuthorizationRuleListResult); err != nil {
		return EventHubsListAuthorizationRulesResponse{}, err
	}
	return result, nil
}

// listAuthorizationRulesHandleError handles the ListAuthorizationRules error response.
func (client *EventHubsClient) listAuthorizationRulesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByNamespace - Gets all the Event Hubs in a Namespace.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) ListByNamespace(resourceGroupName string, namespaceName string, options *EventHubsListByNamespaceOptions) EventHubsListByNamespacePager {
	return &eventHubsListByNamespacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByNamespaceCreateRequest(ctx, resourceGroupName, namespaceName, options)
		},
		advancer: func(ctx context.Context, resp EventHubsListByNamespaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EventHubListResult.NextLink)
		},
	}
}

// listByNamespaceCreateRequest creates the ListByNamespace request.
func (client *EventHubsClient) listByNamespaceCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *EventHubsListByNamespaceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByNamespaceHandleResponse handles the ListByNamespace response.
func (client *EventHubsClient) listByNamespaceHandleResponse(resp *azcore.Response) (EventHubsListByNamespaceResponse, error) {
	result := EventHubsListByNamespaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.EventHubListResult); err != nil {
		return EventHubsListByNamespaceResponse{}, err
	}
	return result, nil
}

// listByNamespaceHandleError handles the ListByNamespace error response.
func (client *EventHubsClient) listByNamespaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListKeys - Gets the ACS and SAS connection strings for the Event Hub.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsListKeysOptions) (EventHubsListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, options)
	if err != nil {
		return EventHubsListKeysResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsListKeysResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EventHubsListKeysResponse{}, client.listKeysHandleError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *EventHubsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *EventHubsListKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *EventHubsClient) listKeysHandleResponse(resp *azcore.Response) (EventHubsListKeysResponse, error) {
	result := EventHubsListKeysResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AccessKeys); err != nil {
		return EventHubsListKeysResponse{}, err
	}
	return result, nil
}

// listKeysHandleError handles the ListKeys error response.
func (client *EventHubsClient) listKeysHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// RegenerateKeys - Regenerates the ACS and SAS connection strings for the Event Hub.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EventHubsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *EventHubsRegenerateKeysOptions) (EventHubsRegenerateKeysResponse, error) {
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, authorizationRuleName, parameters, options)
	if err != nil {
		return EventHubsRegenerateKeysResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return EventHubsRegenerateKeysResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return EventHubsRegenerateKeysResponse{}, client.regenerateKeysHandleError(resp)
	}
	return client.regenerateKeysHandleResponse(resp)
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *EventHubsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters RegenerateAccessKeyParameters, options *EventHubsRegenerateKeysOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}/regenerateKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if authorizationRuleName == "" {
		return nil, errors.New("parameter authorizationRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{authorizationRuleName}", url.PathEscape(authorizationRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-01-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *EventHubsClient) regenerateKeysHandleResponse(resp *azcore.Response) (EventHubsRegenerateKeysResponse, error) {
	result := EventHubsRegenerateKeysResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AccessKeys); err != nil {
		return EventHubsRegenerateKeysResponse{}, err
	}
	return result, nil
}

// regenerateKeysHandleError handles the RegenerateKeys error response.
func (client *EventHubsClient) regenerateKeysHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
