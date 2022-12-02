//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnotificationhubs

// ClientCheckNotificationHubAvailabilityResponse contains the response from method Client.CheckNotificationHubAvailability.
type ClientCheckNotificationHubAvailabilityResponse struct {
	CheckAvailabilityResult
}

// ClientCreateOrUpdateAuthorizationRuleResponse contains the response from method Client.CreateOrUpdateAuthorizationRule.
type ClientCreateOrUpdateAuthorizationRuleResponse struct {
	SharedAccessAuthorizationRuleResource
}

// ClientCreateOrUpdateResponse contains the response from method Client.CreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	NotificationHubResource
}

// ClientDebugSendResponse contains the response from method Client.DebugSend.
type ClientDebugSendResponse struct {
	DebugSendResponse
}

// ClientDeleteAuthorizationRuleResponse contains the response from method Client.DeleteAuthorizationRule.
type ClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetAuthorizationRuleResponse contains the response from method Client.GetAuthorizationRule.
type ClientGetAuthorizationRuleResponse struct {
	SharedAccessAuthorizationRuleResource
}

// ClientGetPnsCredentialsResponse contains the response from method Client.GetPnsCredentials.
type ClientGetPnsCredentialsResponse struct {
	PnsCredentialsResource
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	NotificationHubResource
}

// ClientListAuthorizationRulesResponse contains the response from method Client.ListAuthorizationRules.
type ClientListAuthorizationRulesResponse struct {
	SharedAccessAuthorizationRuleListResult
}

// ClientListKeysResponse contains the response from method Client.ListKeys.
type ClientListKeysResponse struct {
	ResourceListKeys
}

// ClientListResponse contains the response from method Client.List.
type ClientListResponse struct {
	NotificationHubListResult
}

// ClientPatchResponse contains the response from method Client.Patch.
type ClientPatchResponse struct {
	NotificationHubResource
}

// ClientRegenerateKeysResponse contains the response from method Client.RegenerateKeys.
type ClientRegenerateKeysResponse struct {
	ResourceListKeys
}

// NamespacesClientCheckAvailabilityResponse contains the response from method NamespacesClient.CheckAvailability.
type NamespacesClientCheckAvailabilityResponse struct {
	CheckAvailabilityResult
}

// NamespacesClientCreateOrUpdateAuthorizationRuleResponse contains the response from method NamespacesClient.CreateOrUpdateAuthorizationRule.
type NamespacesClientCreateOrUpdateAuthorizationRuleResponse struct {
	SharedAccessAuthorizationRuleResource
}

// NamespacesClientCreateOrUpdateResponse contains the response from method NamespacesClient.CreateOrUpdate.
type NamespacesClientCreateOrUpdateResponse struct {
	NamespaceResource
}

// NamespacesClientDeleteAuthorizationRuleResponse contains the response from method NamespacesClient.DeleteAuthorizationRule.
type NamespacesClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// NamespacesClientDeleteResponse contains the response from method NamespacesClient.Delete.
type NamespacesClientDeleteResponse struct {
	// placeholder for future response values
}

// NamespacesClientGetAuthorizationRuleResponse contains the response from method NamespacesClient.GetAuthorizationRule.
type NamespacesClientGetAuthorizationRuleResponse struct {
	SharedAccessAuthorizationRuleResource
}

// NamespacesClientGetResponse contains the response from method NamespacesClient.Get.
type NamespacesClientGetResponse struct {
	NamespaceResource
}

// NamespacesClientListAllResponse contains the response from method NamespacesClient.ListAll.
type NamespacesClientListAllResponse struct {
	NamespaceListResult
}

// NamespacesClientListAuthorizationRulesResponse contains the response from method NamespacesClient.ListAuthorizationRules.
type NamespacesClientListAuthorizationRulesResponse struct {
	SharedAccessAuthorizationRuleListResult
}

// NamespacesClientListKeysResponse contains the response from method NamespacesClient.ListKeys.
type NamespacesClientListKeysResponse struct {
	ResourceListKeys
}

// NamespacesClientListResponse contains the response from method NamespacesClient.List.
type NamespacesClientListResponse struct {
	NamespaceListResult
}

// NamespacesClientPatchResponse contains the response from method NamespacesClient.Patch.
type NamespacesClientPatchResponse struct {
	NamespaceResource
}

// NamespacesClientRegenerateKeysResponse contains the response from method NamespacesClient.RegenerateKeys.
type NamespacesClientRegenerateKeysResponse struct {
	ResourceListKeys
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}