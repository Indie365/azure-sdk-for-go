//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnotificationhubs

import (
	"context"
	"net/http"
	"time"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
)

// NamespacesCheckAvailabilityResponse contains the response from method Namespaces.CheckAvailability.
type NamespacesCheckAvailabilityResponse struct {
	NamespacesCheckAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCheckAvailabilityResult contains the result from method Namespaces.CheckAvailability.
type NamespacesCheckAvailabilityResult struct {
	CheckAvailabilityResult
}

// NamespacesCreateOrUpdateAuthorizationRuleResponse contains the response from method Namespaces.CreateOrUpdateAuthorizationRule.
type NamespacesCreateOrUpdateAuthorizationRuleResponse struct {
	NamespacesCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateAuthorizationRuleResult contains the result from method Namespaces.CreateOrUpdateAuthorizationRule.
type NamespacesCreateOrUpdateAuthorizationRuleResult struct {
	SharedAccessAuthorizationRuleResource
}

// NamespacesCreateOrUpdateResponse contains the response from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdateResponse struct {
	NamespacesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesCreateOrUpdateResult contains the result from method Namespaces.CreateOrUpdate.
type NamespacesCreateOrUpdateResult struct {
	NamespaceResource
}

// NamespacesDeleteAuthorizationRuleResponse contains the response from method Namespaces.DeleteAuthorizationRule.
type NamespacesDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesDeletePollerResponse contains the response from method Namespaces.Delete.
type NamespacesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *NamespacesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
func (l NamespacesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (NamespacesDeleteResponse, error) {
	respType := NamespacesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a NamespacesDeletePollerResponse from the provided client and resume token.
func (l *NamespacesDeletePollerResponse) Resume(ctx context.Context, client *NamespacesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("NamespacesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &NamespacesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// NamespacesDeleteResponse contains the response from method Namespaces.Delete.
type NamespacesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetAuthorizationRuleResponse contains the response from method Namespaces.GetAuthorizationRule.
type NamespacesGetAuthorizationRuleResponse struct {
	NamespacesGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetAuthorizationRuleResult contains the result from method Namespaces.GetAuthorizationRule.
type NamespacesGetAuthorizationRuleResult struct {
	SharedAccessAuthorizationRuleResource
}

// NamespacesGetResponse contains the response from method Namespaces.Get.
type NamespacesGetResponse struct {
	NamespacesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesGetResult contains the result from method Namespaces.Get.
type NamespacesGetResult struct {
	NamespaceResource
}

// NamespacesListAllResponse contains the response from method Namespaces.ListAll.
type NamespacesListAllResponse struct {
	NamespacesListAllResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListAllResult contains the result from method Namespaces.ListAll.
type NamespacesListAllResult struct {
	NamespaceListResult
}

// NamespacesListAuthorizationRulesResponse contains the response from method Namespaces.ListAuthorizationRules.
type NamespacesListAuthorizationRulesResponse struct {
	NamespacesListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListAuthorizationRulesResult contains the result from method Namespaces.ListAuthorizationRules.
type NamespacesListAuthorizationRulesResult struct {
	SharedAccessAuthorizationRuleListResult
}

// NamespacesListKeysResponse contains the response from method Namespaces.ListKeys.
type NamespacesListKeysResponse struct {
	NamespacesListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListKeysResult contains the result from method Namespaces.ListKeys.
type NamespacesListKeysResult struct {
	SharedAccessAuthorizationRuleListResult
}

// NamespacesListResponse contains the response from method Namespaces.List.
type NamespacesListResponse struct {
	NamespacesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesListResult contains the result from method Namespaces.List.
type NamespacesListResult struct {
	NamespaceListResult
}

// NamespacesPatchResponse contains the response from method Namespaces.Patch.
type NamespacesPatchResponse struct {
	NamespacesPatchResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesPatchResult contains the result from method Namespaces.Patch.
type NamespacesPatchResult struct {
	NamespaceResource
}

// NamespacesRegenerateKeysResponse contains the response from method Namespaces.RegenerateKeys.
type NamespacesRegenerateKeysResponse struct {
	NamespacesRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NamespacesRegenerateKeysResult contains the result from method Namespaces.RegenerateKeys.
type NamespacesRegenerateKeysResult struct {
	ResourceListKeys
}

// NotificationHubsCheckNotificationHubAvailabilityResponse contains the response from method NotificationHubs.CheckNotificationHubAvailability.
type NotificationHubsCheckNotificationHubAvailabilityResponse struct {
	NotificationHubsCheckNotificationHubAvailabilityResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsCheckNotificationHubAvailabilityResult contains the result from method NotificationHubs.CheckNotificationHubAvailability.
type NotificationHubsCheckNotificationHubAvailabilityResult struct {
	CheckAvailabilityResult
}

// NotificationHubsCreateOrUpdateAuthorizationRuleResponse contains the response from method NotificationHubs.CreateOrUpdateAuthorizationRule.
type NotificationHubsCreateOrUpdateAuthorizationRuleResponse struct {
	NotificationHubsCreateOrUpdateAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsCreateOrUpdateAuthorizationRuleResult contains the result from method NotificationHubs.CreateOrUpdateAuthorizationRule.
type NotificationHubsCreateOrUpdateAuthorizationRuleResult struct {
	SharedAccessAuthorizationRuleResource
}

// NotificationHubsCreateOrUpdateResponse contains the response from method NotificationHubs.CreateOrUpdate.
type NotificationHubsCreateOrUpdateResponse struct {
	NotificationHubsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsCreateOrUpdateResult contains the result from method NotificationHubs.CreateOrUpdate.
type NotificationHubsCreateOrUpdateResult struct {
	NotificationHubResource
}

// NotificationHubsDebugSendResponse contains the response from method NotificationHubs.DebugSend.
type NotificationHubsDebugSendResponse struct {
	NotificationHubsDebugSendResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsDebugSendResult contains the result from method NotificationHubs.DebugSend.
type NotificationHubsDebugSendResult struct {
	DebugSendResponse
}

// NotificationHubsDeleteAuthorizationRuleResponse contains the response from method NotificationHubs.DeleteAuthorizationRule.
type NotificationHubsDeleteAuthorizationRuleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsDeleteResponse contains the response from method NotificationHubs.Delete.
type NotificationHubsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsGetAuthorizationRuleResponse contains the response from method NotificationHubs.GetAuthorizationRule.
type NotificationHubsGetAuthorizationRuleResponse struct {
	NotificationHubsGetAuthorizationRuleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsGetAuthorizationRuleResult contains the result from method NotificationHubs.GetAuthorizationRule.
type NotificationHubsGetAuthorizationRuleResult struct {
	SharedAccessAuthorizationRuleResource
}

// NotificationHubsGetPnsCredentialsResponse contains the response from method NotificationHubs.GetPnsCredentials.
type NotificationHubsGetPnsCredentialsResponse struct {
	NotificationHubsGetPnsCredentialsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsGetPnsCredentialsResult contains the result from method NotificationHubs.GetPnsCredentials.
type NotificationHubsGetPnsCredentialsResult struct {
	PnsCredentialsResource
}

// NotificationHubsGetResponse contains the response from method NotificationHubs.Get.
type NotificationHubsGetResponse struct {
	NotificationHubsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsGetResult contains the result from method NotificationHubs.Get.
type NotificationHubsGetResult struct {
	NotificationHubResource
}

// NotificationHubsListAuthorizationRulesResponse contains the response from method NotificationHubs.ListAuthorizationRules.
type NotificationHubsListAuthorizationRulesResponse struct {
	NotificationHubsListAuthorizationRulesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsListAuthorizationRulesResult contains the result from method NotificationHubs.ListAuthorizationRules.
type NotificationHubsListAuthorizationRulesResult struct {
	SharedAccessAuthorizationRuleListResult
}

// NotificationHubsListKeysResponse contains the response from method NotificationHubs.ListKeys.
type NotificationHubsListKeysResponse struct {
	NotificationHubsListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsListKeysResult contains the result from method NotificationHubs.ListKeys.
type NotificationHubsListKeysResult struct {
	ResourceListKeys
}

// NotificationHubsListResponse contains the response from method NotificationHubs.List.
type NotificationHubsListResponse struct {
	NotificationHubsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsListResult contains the result from method NotificationHubs.List.
type NotificationHubsListResult struct {
	NotificationHubListResult
}

// NotificationHubsPatchResponse contains the response from method NotificationHubs.Patch.
type NotificationHubsPatchResponse struct {
	NotificationHubsPatchResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsPatchResult contains the result from method NotificationHubs.Patch.
type NotificationHubsPatchResult struct {
	NotificationHubResource
}

// NotificationHubsRegenerateKeysResponse contains the response from method NotificationHubs.RegenerateKeys.
type NotificationHubsRegenerateKeysResponse struct {
	NotificationHubsRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NotificationHubsRegenerateKeysResult contains the result from method NotificationHubs.RegenerateKeys.
type NotificationHubsRegenerateKeysResult struct {
	ResourceListKeys
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}
