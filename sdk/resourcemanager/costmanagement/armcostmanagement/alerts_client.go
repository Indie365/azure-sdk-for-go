//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

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

// AlertsClient contains the methods for the Alerts group.
// Don't use this type directly, use NewAlertsClient() instead.
type AlertsClient struct {
	internal *arm.Client
}

// NewAlertsClient creates a new instance of AlertsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAlertsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AlertsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AlertsClient{
		internal: cl,
	}
	return client, nil
}

// Dismiss - Dismisses the specified alert
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - scope - The scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for subscription
//     scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
//     resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
//     scope,
//     '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
//     for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
//     for invoiceSection scope, and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
//   - alertID - Alert ID
//   - parameters - Parameters supplied to the Dismiss Alert operation.
//   - options - AlertsClientDismissOptions contains the optional parameters for the AlertsClient.Dismiss method.
func (client *AlertsClient) Dismiss(ctx context.Context, scope string, alertID string, parameters DismissAlertPayload, options *AlertsClientDismissOptions) (AlertsClientDismissResponse, error) {
	var err error
	const operationName = "AlertsClient.Dismiss"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.dismissCreateRequest(ctx, scope, alertID, parameters, options)
	if err != nil {
		return AlertsClientDismissResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertsClientDismissResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertsClientDismissResponse{}, err
	}
	resp, err := client.dismissHandleResponse(httpResp)
	return resp, err
}

// dismissCreateRequest creates the Dismiss request.
func (client *AlertsClient) dismissCreateRequest(ctx context.Context, scope string, alertID string, parameters DismissAlertPayload, options *AlertsClientDismissOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/alerts/{alertId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", alertID)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// dismissHandleResponse handles the Dismiss response.
func (client *AlertsClient) dismissHandleResponse(resp *http.Response) (AlertsClientDismissResponse, error) {
	result := AlertsClientDismissResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return AlertsClientDismissResponse{}, err
	}
	return result, nil
}

// Get - Gets the alert for the scope by alert ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - scope - The scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for subscription
//     scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
//     resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
//     scope,
//     '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
//     for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
//     for invoiceSection scope, and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
//   - alertID - Alert ID
//   - options - AlertsClientGetOptions contains the optional parameters for the AlertsClient.Get method.
func (client *AlertsClient) Get(ctx context.Context, scope string, alertID string, options *AlertsClientGetOptions) (AlertsClientGetResponse, error) {
	var err error
	const operationName = "AlertsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, alertID, options)
	if err != nil {
		return AlertsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AlertsClient) getCreateRequest(ctx context.Context, scope string, alertID string, options *AlertsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/alerts/{alertId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	urlPath = strings.ReplaceAll(urlPath, "{alertId}", alertID)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AlertsClient) getHandleResponse(resp *http.Response) (AlertsClientGetResponse, error) {
	result := AlertsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Alert); err != nil {
		return AlertsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the alerts for scope defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - scope - The scope associated with alerts operations. This includes '/subscriptions/{subscriptionId}/' for subscription
//     scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
//     resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
//     scope,
//     '/providers/Microsoft.Management/managementGroups/{managementGroupId} for Management Group scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
//     for billingProfile scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
//     for invoiceSection scope, and
//     '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
//   - options - AlertsClientListOptions contains the optional parameters for the AlertsClient.List method.
func (client *AlertsClient) List(ctx context.Context, scope string, options *AlertsClientListOptions) (AlertsClientListResponse, error) {
	var err error
	const operationName = "AlertsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, scope, options)
	if err != nil {
		return AlertsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *AlertsClient) listCreateRequest(ctx context.Context, scope string, options *AlertsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/alerts"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AlertsClient) listHandleResponse(resp *http.Response) (AlertsClientListResponse, error) {
	result := AlertsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsResult); err != nil {
		return AlertsClientListResponse{}, err
	}
	return result, nil
}

// ListExternal - Lists the Alerts for external cloud provider type defined.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-10-01
//   - externalCloudProviderType - The external cloud provider type associated with dimension/query operations. This includes
//     'externalSubscriptions' for linked account and 'externalBillingAccounts' for consolidated account.
//   - externalCloudProviderID - This can be '{externalSubscriptionId}' for linked account or '{externalBillingAccountId}' for
//     consolidated account used with dimension/query operations.
//   - options - AlertsClientListExternalOptions contains the optional parameters for the AlertsClient.ListExternal method.
func (client *AlertsClient) ListExternal(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, options *AlertsClientListExternalOptions) (AlertsClientListExternalResponse, error) {
	var err error
	const operationName = "AlertsClient.ListExternal"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listExternalCreateRequest(ctx, externalCloudProviderType, externalCloudProviderID, options)
	if err != nil {
		return AlertsClientListExternalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AlertsClientListExternalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AlertsClientListExternalResponse{}, err
	}
	resp, err := client.listExternalHandleResponse(httpResp)
	return resp, err
}

// listExternalCreateRequest creates the ListExternal request.
func (client *AlertsClient) listExternalCreateRequest(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, options *AlertsClientListExternalOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/{externalCloudProviderType}/{externalCloudProviderId}/alerts"
	if externalCloudProviderType == "" {
		return nil, errors.New("parameter externalCloudProviderType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderType}", url.PathEscape(string(externalCloudProviderType)))
	if externalCloudProviderID == "" {
		return nil, errors.New("parameter externalCloudProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderId}", url.PathEscape(externalCloudProviderID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listExternalHandleResponse handles the ListExternal response.
func (client *AlertsClient) listExternalHandleResponse(resp *http.Response) (AlertsClientListExternalResponse, error) {
	result := AlertsClientListExternalResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AlertsResult); err != nil {
		return AlertsClientListExternalResponse{}, err
	}
	return result, nil
}
