//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmsi

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// SystemAssignedIdentitiesClientGetByScopeResponse contains the response from method SystemAssignedIdentitiesClient.GetByScope.
type SystemAssignedIdentitiesClientGetByScopeResponse struct {
	SystemAssignedIdentity
}

// UserAssignedIdentitiesClientCreateOrUpdateResponse contains the response from method UserAssignedIdentitiesClient.CreateOrUpdate.
type UserAssignedIdentitiesClientCreateOrUpdateResponse struct {
	Identity
}

// UserAssignedIdentitiesClientDeleteResponse contains the response from method UserAssignedIdentitiesClient.Delete.
type UserAssignedIdentitiesClientDeleteResponse struct {
	// placeholder for future response values
}

// UserAssignedIdentitiesClientGetResponse contains the response from method UserAssignedIdentitiesClient.Get.
type UserAssignedIdentitiesClientGetResponse struct {
	Identity
}

// UserAssignedIdentitiesClientListAssociatedResourcesResponse contains the response from method UserAssignedIdentitiesClient.ListAssociatedResources.
type UserAssignedIdentitiesClientListAssociatedResourcesResponse struct {
	AssociatedResourcesListResult
}

// UserAssignedIdentitiesClientListByResourceGroupResponse contains the response from method UserAssignedIdentitiesClient.ListByResourceGroup.
type UserAssignedIdentitiesClientListByResourceGroupResponse struct {
	UserAssignedIdentitiesListResult
}

// UserAssignedIdentitiesClientListBySubscriptionResponse contains the response from method UserAssignedIdentitiesClient.ListBySubscription.
type UserAssignedIdentitiesClientListBySubscriptionResponse struct {
	UserAssignedIdentitiesListResult
}

// UserAssignedIdentitiesClientUpdateResponse contains the response from method UserAssignedIdentitiesClient.Update.
type UserAssignedIdentitiesClientUpdateResponse struct {
	Identity
}
