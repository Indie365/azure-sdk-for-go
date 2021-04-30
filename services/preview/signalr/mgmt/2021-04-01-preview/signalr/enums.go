package signalr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ACLAction enumerates the values for acl action.
type ACLAction string

const (
	// ACLActionAllow ...
	ACLActionAllow ACLAction = "Allow"
	// ACLActionDeny ...
	ACLActionDeny ACLAction = "Deny"
)

// PossibleACLActionValues returns an array of possible values for the ACLAction const type.
func PossibleACLActionValues() []ACLAction {
	return []ACLAction{ACLActionAllow, ACLActionDeny}
}

// FeatureFlags enumerates the values for feature flags.
type FeatureFlags string

const (
	// FeatureFlagsEnableConnectivityLogs ...
	FeatureFlagsEnableConnectivityLogs FeatureFlags = "EnableConnectivityLogs"
	// FeatureFlagsEnableLiveTrace ...
	FeatureFlagsEnableLiveTrace FeatureFlags = "EnableLiveTrace"
	// FeatureFlagsEnableMessagingLogs ...
	FeatureFlagsEnableMessagingLogs FeatureFlags = "EnableMessagingLogs"
	// FeatureFlagsServiceMode ...
	FeatureFlagsServiceMode FeatureFlags = "ServiceMode"
)

// PossibleFeatureFlagsValues returns an array of possible values for the FeatureFlags const type.
func PossibleFeatureFlagsValues() []FeatureFlags {
	return []FeatureFlags{FeatureFlagsEnableConnectivityLogs, FeatureFlagsEnableLiveTrace, FeatureFlagsEnableMessagingLogs, FeatureFlagsServiceMode}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// KeyTypePrimary ...
	KeyTypePrimary KeyType = "Primary"
	// KeyTypeSecondary ...
	KeyTypeSecondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{KeyTypePrimary, KeyTypeSecondary}
}

// ManagedIdentityType enumerates the values for managed identity type.
type ManagedIdentityType string

const (
	// ManagedIdentityTypeNone ...
	ManagedIdentityTypeNone ManagedIdentityType = "None"
	// ManagedIdentityTypeSystemAssigned ...
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
	// ManagedIdentityTypeUserAssigned ...
	ManagedIdentityTypeUserAssigned ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns an array of possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{ManagedIdentityTypeNone, ManagedIdentityTypeSystemAssigned, ManagedIdentityTypeUserAssigned}
}

// RequestType enumerates the values for request type.
type RequestType string

const (
	// RequestTypeClientConnection ...
	RequestTypeClientConnection RequestType = "ClientConnection"
	// RequestTypeRESTAPI ...
	RequestTypeRESTAPI RequestType = "RESTAPI"
	// RequestTypeServerConnection ...
	RequestTypeServerConnection RequestType = "ServerConnection"
	// RequestTypeTrace ...
	RequestTypeTrace RequestType = "Trace"
)

// PossibleRequestTypeValues returns an array of possible values for the RequestType const type.
func PossibleRequestTypeValues() []RequestType {
	return []RequestType{RequestTypeClientConnection, RequestTypeRESTAPI, RequestTypeServerConnection, RequestTypeTrace}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// PrivateLinkServiceConnectionStatus enumerates the values for private link service connection status.
type PrivateLinkServiceConnectionStatus string

const (
	// PrivateLinkServiceConnectionStatusApproved ...
	PrivateLinkServiceConnectionStatusApproved PrivateLinkServiceConnectionStatus = "Approved"
	// PrivateLinkServiceConnectionStatusDisconnected ...
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	// PrivateLinkServiceConnectionStatusPending ...
	PrivateLinkServiceConnectionStatusPending PrivateLinkServiceConnectionStatus = "Pending"
	// PrivateLinkServiceConnectionStatusRejected ...
	PrivateLinkServiceConnectionStatusRejected PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns an array of possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{PrivateLinkServiceConnectionStatusApproved, PrivateLinkServiceConnectionStatusDisconnected, PrivateLinkServiceConnectionStatusPending, PrivateLinkServiceConnectionStatusRejected}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMoving ...
	ProvisioningStateMoving ProvisioningState = "Moving"
	// ProvisioningStateRunning ...
	ProvisioningStateRunning ProvisioningState = "Running"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUnknown ...
	ProvisioningStateUnknown ProvisioningState = "Unknown"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCanceled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateMoving, ProvisioningStateRunning, ProvisioningStateSucceeded, ProvisioningStateUnknown, ProvisioningStateUpdating}
}

// ServiceKind enumerates the values for service kind.
type ServiceKind string

const (
	// ServiceKindRawWebSockets ...
	ServiceKindRawWebSockets ServiceKind = "RawWebSockets"
	// ServiceKindSignalR ...
	ServiceKindSignalR ServiceKind = "SignalR"
)

// PossibleServiceKindValues returns an array of possible values for the ServiceKind const type.
func PossibleServiceKindValues() []ServiceKind {
	return []ServiceKind{ServiceKindRawWebSockets, ServiceKindSignalR}
}

// SharedPrivateLinkResourceStatus enumerates the values for shared private link resource status.
type SharedPrivateLinkResourceStatus string

const (
	// SharedPrivateLinkResourceStatusApproved ...
	SharedPrivateLinkResourceStatusApproved SharedPrivateLinkResourceStatus = "Approved"
	// SharedPrivateLinkResourceStatusDisconnected ...
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	// SharedPrivateLinkResourceStatusPending ...
	SharedPrivateLinkResourceStatusPending SharedPrivateLinkResourceStatus = "Pending"
	// SharedPrivateLinkResourceStatusRejected ...
	SharedPrivateLinkResourceStatusRejected SharedPrivateLinkResourceStatus = "Rejected"
	// SharedPrivateLinkResourceStatusTimeout ...
	SharedPrivateLinkResourceStatusTimeout SharedPrivateLinkResourceStatus = "Timeout"
)

// PossibleSharedPrivateLinkResourceStatusValues returns an array of possible values for the SharedPrivateLinkResourceStatus const type.
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return []SharedPrivateLinkResourceStatus{SharedPrivateLinkResourceStatusApproved, SharedPrivateLinkResourceStatusDisconnected, SharedPrivateLinkResourceStatusPending, SharedPrivateLinkResourceStatusRejected, SharedPrivateLinkResourceStatusTimeout}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTier = "Basic"
	// SkuTierFree ...
	SkuTierFree SkuTier = "Free"
	// SkuTierPremium ...
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard ...
	SkuTierStandard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBasic, SkuTierFree, SkuTierPremium, SkuTierStandard}
}

// UpstreamAuthType enumerates the values for upstream auth type.
type UpstreamAuthType string

const (
	// UpstreamAuthTypeManagedIdentity ...
	UpstreamAuthTypeManagedIdentity UpstreamAuthType = "ManagedIdentity"
	// UpstreamAuthTypeNone ...
	UpstreamAuthTypeNone UpstreamAuthType = "None"
)

// PossibleUpstreamAuthTypeValues returns an array of possible values for the UpstreamAuthType const type.
func PossibleUpstreamAuthTypeValues() []UpstreamAuthType {
	return []UpstreamAuthType{UpstreamAuthTypeManagedIdentity, UpstreamAuthTypeNone}
}
