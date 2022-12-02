//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabricks

const (
	moduleName    = "armdatabricks"
	moduleVersion = "v0.6.0"
)

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

// CustomParameterType - Provisioning status of the workspace.
type CustomParameterType string

const (
	CustomParameterTypeBool   CustomParameterType = "Bool"
	CustomParameterTypeObject CustomParameterType = "Object"
	CustomParameterTypeString CustomParameterType = "String"
)

// PossibleCustomParameterTypeValues returns the possible values for the CustomParameterType const type.
func PossibleCustomParameterTypeValues() []CustomParameterType {
	return []CustomParameterType{
		CustomParameterTypeBool,
		CustomParameterTypeObject,
		CustomParameterTypeString,
	}
}

// EncryptionKeySource - The encryption keySource (provider). Possible values (case-insensitive): Microsoft.Keyvault
type EncryptionKeySource string

const (
	EncryptionKeySourceMicrosoftKeyvault EncryptionKeySource = "Microsoft.Keyvault"
)

// PossibleEncryptionKeySourceValues returns the possible values for the EncryptionKeySource const type.
func PossibleEncryptionKeySourceValues() []EncryptionKeySource {
	return []EncryptionKeySource{
		EncryptionKeySourceMicrosoftKeyvault,
	}
}

// KeySource - The encryption keySource (provider). Possible values (case-insensitive): Default, Microsoft.Keyvault
type KeySource string

const (
	KeySourceDefault           KeySource = "Default"
	KeySourceMicrosoftKeyvault KeySource = "Microsoft.Keyvault"
)

// PossibleKeySourceValues returns the possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{
		KeySourceDefault,
		KeySourceMicrosoftKeyvault,
	}
}

// PeeringProvisioningState - The current provisioning state.
type PeeringProvisioningState string

const (
	PeeringProvisioningStateDeleting  PeeringProvisioningState = "Deleting"
	PeeringProvisioningStateFailed    PeeringProvisioningState = "Failed"
	PeeringProvisioningStateSucceeded PeeringProvisioningState = "Succeeded"
	PeeringProvisioningStateUpdating  PeeringProvisioningState = "Updating"
)

// PossiblePeeringProvisioningStateValues returns the possible values for the PeeringProvisioningState const type.
func PossiblePeeringProvisioningStateValues() []PeeringProvisioningState {
	return []PeeringProvisioningState{
		PeeringProvisioningStateDeleting,
		PeeringProvisioningStateFailed,
		PeeringProvisioningStateSucceeded,
		PeeringProvisioningStateUpdating,
	}
}

// PeeringState - The status of the virtual network peering.
type PeeringState string

const (
	PeeringStateConnected    PeeringState = "Connected"
	PeeringStateDisconnected PeeringState = "Disconnected"
	PeeringStateInitiated    PeeringState = "Initiated"
)

// PossiblePeeringStateValues returns the possible values for the PeeringState const type.
func PossiblePeeringStateValues() []PeeringState {
	return []PeeringState{
		PeeringStateConnected,
		PeeringStateDisconnected,
		PeeringStateInitiated,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating  PrivateEndpointConnectionProvisioningState = "Updating"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
		PrivateEndpointConnectionProvisioningStateUpdating,
	}
}

// PrivateLinkServiceConnectionStatus - The status of a private endpoint connection
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusDisconnected,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
	}
}

// ProvisioningState - Provisioning status of the workspace.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreated   ProvisioningState = "Created"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateReady     ProvisioningState = "Ready"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateReady,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - The network access type for accessing workspace. Set value to disabled to access workspace only via
// private link.
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

// RequiredNsgRules - Gets or sets a value indicating whether data plane (clusters) to control plane communication happen
// over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'.
// 'NoAzureServiceRules' value is for internal use only.
type RequiredNsgRules string

const (
	RequiredNsgRulesAllRules               RequiredNsgRules = "AllRules"
	RequiredNsgRulesNoAzureDatabricksRules RequiredNsgRules = "NoAzureDatabricksRules"
	RequiredNsgRulesNoAzureServiceRules    RequiredNsgRules = "NoAzureServiceRules"
)

// PossibleRequiredNsgRulesValues returns the possible values for the RequiredNsgRules const type.
func PossibleRequiredNsgRulesValues() []RequiredNsgRules {
	return []RequiredNsgRules{
		RequiredNsgRulesAllRules,
		RequiredNsgRulesNoAzureDatabricksRules,
		RequiredNsgRulesNoAzureServiceRules,
	}
}