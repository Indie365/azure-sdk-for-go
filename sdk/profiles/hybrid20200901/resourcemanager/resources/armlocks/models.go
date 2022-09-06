//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlocks

// AuthorizationOperationsClientListOptions contains the optional parameters for the AuthorizationOperationsClient.List method.
type AuthorizationOperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ManagementLockListResult - The list of locks.
type ManagementLockListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of locks.
	Value []*ManagementLockObject `json:"value,omitempty"`
}

// ManagementLockObject - The lock information.
type ManagementLockObject struct {
	// REQUIRED; The properties of the lock.
	Properties *ManagementLockProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID of the lock.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the lock.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type of the lock - Microsoft.Authorization/locks.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ManagementLockOwner - Lock owner properties.
type ManagementLockOwner struct {
	// The application ID of the lock owner.
	ApplicationID *string `json:"applicationId,omitempty"`
}

// ManagementLockProperties - The lock properties.
type ManagementLockProperties struct {
	// REQUIRED; The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized
	// users are able to read and modify the resources, but not delete. ReadOnly means
	// authorized users can only read from a resource, but they can't modify or delete it.
	Level *LockLevel `json:"level,omitempty"`

	// Notes about the lock. Maximum of 512 characters.
	Notes *string `json:"notes,omitempty"`

	// The owners of the lock.
	Owners []*ManagementLockOwner `json:"owners,omitempty"`
}

// ManagementLocksClientCreateOrUpdateAtResourceGroupLevelOptions contains the optional parameters for the ManagementLocksClient.CreateOrUpdateAtResourceGroupLevel
// method.
type ManagementLocksClientCreateOrUpdateAtResourceGroupLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientCreateOrUpdateAtResourceLevelOptions contains the optional parameters for the ManagementLocksClient.CreateOrUpdateAtResourceLevel
// method.
type ManagementLocksClientCreateOrUpdateAtResourceLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientCreateOrUpdateAtSubscriptionLevelOptions contains the optional parameters for the ManagementLocksClient.CreateOrUpdateAtSubscriptionLevel
// method.
type ManagementLocksClientCreateOrUpdateAtSubscriptionLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientCreateOrUpdateByScopeOptions contains the optional parameters for the ManagementLocksClient.CreateOrUpdateByScope
// method.
type ManagementLocksClientCreateOrUpdateByScopeOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientDeleteAtResourceGroupLevelOptions contains the optional parameters for the ManagementLocksClient.DeleteAtResourceGroupLevel
// method.
type ManagementLocksClientDeleteAtResourceGroupLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientDeleteAtResourceLevelOptions contains the optional parameters for the ManagementLocksClient.DeleteAtResourceLevel
// method.
type ManagementLocksClientDeleteAtResourceLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientDeleteAtSubscriptionLevelOptions contains the optional parameters for the ManagementLocksClient.DeleteAtSubscriptionLevel
// method.
type ManagementLocksClientDeleteAtSubscriptionLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientDeleteByScopeOptions contains the optional parameters for the ManagementLocksClient.DeleteByScope
// method.
type ManagementLocksClientDeleteByScopeOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientGetAtResourceGroupLevelOptions contains the optional parameters for the ManagementLocksClient.GetAtResourceGroupLevel
// method.
type ManagementLocksClientGetAtResourceGroupLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientGetAtResourceLevelOptions contains the optional parameters for the ManagementLocksClient.GetAtResourceLevel
// method.
type ManagementLocksClientGetAtResourceLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientGetAtSubscriptionLevelOptions contains the optional parameters for the ManagementLocksClient.GetAtSubscriptionLevel
// method.
type ManagementLocksClientGetAtSubscriptionLevelOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientGetByScopeOptions contains the optional parameters for the ManagementLocksClient.GetByScope method.
type ManagementLocksClientGetByScopeOptions struct {
	// placeholder for future optional parameters
}

// ManagementLocksClientListAtResourceGroupLevelOptions contains the optional parameters for the ManagementLocksClient.ListAtResourceGroupLevel
// method.
type ManagementLocksClientListAtResourceGroupLevelOptions struct {
	// The filter to apply on the operation.
	Filter *string
}

// ManagementLocksClientListAtResourceLevelOptions contains the optional parameters for the ManagementLocksClient.ListAtResourceLevel
// method.
type ManagementLocksClientListAtResourceLevelOptions struct {
	// The filter to apply on the operation.
	Filter *string
}

// ManagementLocksClientListAtSubscriptionLevelOptions contains the optional parameters for the ManagementLocksClient.ListAtSubscriptionLevel
// method.
type ManagementLocksClientListAtSubscriptionLevelOptions struct {
	// The filter to apply on the operation.
	Filter *string
}

// ManagementLocksClientListByScopeOptions contains the optional parameters for the ManagementLocksClient.ListByScope method.
type ManagementLocksClientListByScopeOptions struct {
	// The filter to apply on the operation.
	Filter *string
}

// Operation - Microsoft.Authorization operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.Authorization
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of the request to list Microsoft.Authorization operations. It contains a list of operations
// and a URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Microsoft.Authorization operations.
	Value []*Operation `json:"value,omitempty"`
}
