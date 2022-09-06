//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlocks

// AuthorizationOperationsClientListResponse contains the response from method AuthorizationOperationsClient.List.
type AuthorizationOperationsClientListResponse struct {
	OperationListResult
}

// ManagementLocksClientCreateOrUpdateAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.CreateOrUpdateAtResourceGroupLevel.
type ManagementLocksClientCreateOrUpdateAtResourceGroupLevelResponse struct {
	ManagementLockObject
}

// ManagementLocksClientCreateOrUpdateAtResourceLevelResponse contains the response from method ManagementLocksClient.CreateOrUpdateAtResourceLevel.
type ManagementLocksClientCreateOrUpdateAtResourceLevelResponse struct {
	ManagementLockObject
}

// ManagementLocksClientCreateOrUpdateAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.CreateOrUpdateAtSubscriptionLevel.
type ManagementLocksClientCreateOrUpdateAtSubscriptionLevelResponse struct {
	ManagementLockObject
}

// ManagementLocksClientCreateOrUpdateByScopeResponse contains the response from method ManagementLocksClient.CreateOrUpdateByScope.
type ManagementLocksClientCreateOrUpdateByScopeResponse struct {
	ManagementLockObject
}

// ManagementLocksClientDeleteAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.DeleteAtResourceGroupLevel.
type ManagementLocksClientDeleteAtResourceGroupLevelResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientDeleteAtResourceLevelResponse contains the response from method ManagementLocksClient.DeleteAtResourceLevel.
type ManagementLocksClientDeleteAtResourceLevelResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientDeleteAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.DeleteAtSubscriptionLevel.
type ManagementLocksClientDeleteAtSubscriptionLevelResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientDeleteByScopeResponse contains the response from method ManagementLocksClient.DeleteByScope.
type ManagementLocksClientDeleteByScopeResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientGetAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.GetAtResourceGroupLevel.
type ManagementLocksClientGetAtResourceGroupLevelResponse struct {
	ManagementLockObject
}

// ManagementLocksClientGetAtResourceLevelResponse contains the response from method ManagementLocksClient.GetAtResourceLevel.
type ManagementLocksClientGetAtResourceLevelResponse struct {
	ManagementLockObject
}

// ManagementLocksClientGetAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.GetAtSubscriptionLevel.
type ManagementLocksClientGetAtSubscriptionLevelResponse struct {
	ManagementLockObject
}

// ManagementLocksClientGetByScopeResponse contains the response from method ManagementLocksClient.GetByScope.
type ManagementLocksClientGetByScopeResponse struct {
	ManagementLockObject
}

// ManagementLocksClientListAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.ListAtResourceGroupLevel.
type ManagementLocksClientListAtResourceGroupLevelResponse struct {
	ManagementLockListResult
}

// ManagementLocksClientListAtResourceLevelResponse contains the response from method ManagementLocksClient.ListAtResourceLevel.
type ManagementLocksClientListAtResourceLevelResponse struct {
	ManagementLockListResult
}

// ManagementLocksClientListAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.ListAtSubscriptionLevel.
type ManagementLocksClientListAtSubscriptionLevelResponse struct {
	ManagementLockListResult
}

// ManagementLocksClientListByScopeResponse contains the response from method ManagementLocksClient.ListByScope.
type ManagementLocksClientListByScopeResponse struct {
	ManagementLockListResult
}
