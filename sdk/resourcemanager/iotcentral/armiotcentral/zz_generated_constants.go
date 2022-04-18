//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotcentral

const (
	moduleName    = "armiotcentral"
	moduleVersion = "v0.4.0"
)

// AppSKU - The name of the SKU.
type AppSKU string

const (
	AppSKUST0 AppSKU = "ST0"
	AppSKUST1 AppSKU = "ST1"
	AppSKUST2 AppSKU = "ST2"
)

// PossibleAppSKUValues returns the possible values for the AppSKU const type.
func PossibleAppSKUValues() []AppSKU {
	return []AppSKU{
		AppSKUST0,
		AppSKUST1,
		AppSKUST2,
	}
}

// AppState - The current state of the application.
type AppState string

const (
	AppStateCreated   AppState = "created"
	AppStateSuspended AppState = "suspended"
)

// PossibleAppStateValues returns the possible values for the AppState const type.
func PossibleAppStateValues() []AppState {
	return []AppState{
		AppStateCreated,
		AppStateSuspended,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// NetworkAction - Whether to allow or deny network traffic.
type NetworkAction string

const (
	NetworkActionAllow NetworkAction = "Allow"
	NetworkActionDeny  NetworkAction = "Deny"
)

// PossibleNetworkActionValues returns the possible values for the NetworkAction const type.
func PossibleNetworkActionValues() []NetworkAction {
	return []NetworkAction{
		NetworkActionAllow,
		NetworkActionDeny,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningState - The provisioning state of the application.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Whether requests from the public network are allowed.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// SystemAssignedServiceIdentityType - Type of managed service identity (either system assigned, or none).
type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           SystemAssignedServiceIdentityType = "None"
	SystemAssignedServiceIdentityTypeSystemAssigned SystemAssignedServiceIdentityType = "SystemAssigned"
)

// PossibleSystemAssignedServiceIdentityTypeValues returns the possible values for the SystemAssignedServiceIdentityType const type.
func PossibleSystemAssignedServiceIdentityTypeValues() []SystemAssignedServiceIdentityType {
	return []SystemAssignedServiceIdentityType{
		SystemAssignedServiceIdentityTypeNone,
		SystemAssignedServiceIdentityTypeSystemAssigned,
	}
}
