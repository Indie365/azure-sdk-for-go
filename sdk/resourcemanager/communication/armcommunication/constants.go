//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommunication

const (
	moduleName = "armcommunication"
	moduleVersion = "v2.1.0-beta.1"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{	
		ActionTypeInternal,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{	
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// CommunicationServicesProvisioningState - Provisioning state of the resource.
type CommunicationServicesProvisioningState string

const (
	CommunicationServicesProvisioningStateCanceled CommunicationServicesProvisioningState = "Canceled"
	CommunicationServicesProvisioningStateCreating CommunicationServicesProvisioningState = "Creating"
	CommunicationServicesProvisioningStateDeleting CommunicationServicesProvisioningState = "Deleting"
	CommunicationServicesProvisioningStateFailed CommunicationServicesProvisioningState = "Failed"
	CommunicationServicesProvisioningStateMoving CommunicationServicesProvisioningState = "Moving"
	CommunicationServicesProvisioningStateRunning CommunicationServicesProvisioningState = "Running"
	CommunicationServicesProvisioningStateSucceeded CommunicationServicesProvisioningState = "Succeeded"
	CommunicationServicesProvisioningStateUnknown CommunicationServicesProvisioningState = "Unknown"
	CommunicationServicesProvisioningStateUpdating CommunicationServicesProvisioningState = "Updating"
)

// PossibleCommunicationServicesProvisioningStateValues returns the possible values for the CommunicationServicesProvisioningState const type.
func PossibleCommunicationServicesProvisioningStateValues() []CommunicationServicesProvisioningState {
	return []CommunicationServicesProvisioningState{	
		CommunicationServicesProvisioningStateCanceled,
		CommunicationServicesProvisioningStateCreating,
		CommunicationServicesProvisioningStateDeleting,
		CommunicationServicesProvisioningStateFailed,
		CommunicationServicesProvisioningStateMoving,
		CommunicationServicesProvisioningStateRunning,
		CommunicationServicesProvisioningStateSucceeded,
		CommunicationServicesProvisioningStateUnknown,
		CommunicationServicesProvisioningStateUpdating,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication CreatedByType = "Application"
	CreatedByTypeKey CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser CreatedByType = "User"
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

// DomainManagement - Describes how a Domains resource is being managed.
type DomainManagement string

const (
	DomainManagementAzureManaged DomainManagement = "AzureManaged"
	DomainManagementCustomerManaged DomainManagement = "CustomerManaged"
	DomainManagementCustomerManagedInExchangeOnline DomainManagement = "CustomerManagedInExchangeOnline"
)

// PossibleDomainManagementValues returns the possible values for the DomainManagement const type.
func PossibleDomainManagementValues() []DomainManagement {
	return []DomainManagement{	
		DomainManagementAzureManaged,
		DomainManagementCustomerManaged,
		DomainManagementCustomerManagedInExchangeOnline,
	}
}

// DomainsProvisioningState - Provisioning state of the resource.
type DomainsProvisioningState string

const (
	DomainsProvisioningStateCanceled DomainsProvisioningState = "Canceled"
	DomainsProvisioningStateCreating DomainsProvisioningState = "Creating"
	DomainsProvisioningStateDeleting DomainsProvisioningState = "Deleting"
	DomainsProvisioningStateFailed DomainsProvisioningState = "Failed"
	DomainsProvisioningStateMoving DomainsProvisioningState = "Moving"
	DomainsProvisioningStateRunning DomainsProvisioningState = "Running"
	DomainsProvisioningStateSucceeded DomainsProvisioningState = "Succeeded"
	DomainsProvisioningStateUnknown DomainsProvisioningState = "Unknown"
	DomainsProvisioningStateUpdating DomainsProvisioningState = "Updating"
)

// PossibleDomainsProvisioningStateValues returns the possible values for the DomainsProvisioningState const type.
func PossibleDomainsProvisioningStateValues() []DomainsProvisioningState {
	return []DomainsProvisioningState{	
		DomainsProvisioningStateCanceled,
		DomainsProvisioningStateCreating,
		DomainsProvisioningStateDeleting,
		DomainsProvisioningStateFailed,
		DomainsProvisioningStateMoving,
		DomainsProvisioningStateRunning,
		DomainsProvisioningStateSucceeded,
		DomainsProvisioningStateUnknown,
		DomainsProvisioningStateUpdating,
	}
}

// EmailServicesProvisioningState - Provisioning state of the resource.
type EmailServicesProvisioningState string

const (
	EmailServicesProvisioningStateCanceled EmailServicesProvisioningState = "Canceled"
	EmailServicesProvisioningStateCreating EmailServicesProvisioningState = "Creating"
	EmailServicesProvisioningStateDeleting EmailServicesProvisioningState = "Deleting"
	EmailServicesProvisioningStateFailed EmailServicesProvisioningState = "Failed"
	EmailServicesProvisioningStateMoving EmailServicesProvisioningState = "Moving"
	EmailServicesProvisioningStateRunning EmailServicesProvisioningState = "Running"
	EmailServicesProvisioningStateSucceeded EmailServicesProvisioningState = "Succeeded"
	EmailServicesProvisioningStateUnknown EmailServicesProvisioningState = "Unknown"
	EmailServicesProvisioningStateUpdating EmailServicesProvisioningState = "Updating"
)

// PossibleEmailServicesProvisioningStateValues returns the possible values for the EmailServicesProvisioningState const type.
func PossibleEmailServicesProvisioningStateValues() []EmailServicesProvisioningState {
	return []EmailServicesProvisioningState{	
		EmailServicesProvisioningStateCanceled,
		EmailServicesProvisioningStateCreating,
		EmailServicesProvisioningStateDeleting,
		EmailServicesProvisioningStateFailed,
		EmailServicesProvisioningStateMoving,
		EmailServicesProvisioningStateRunning,
		EmailServicesProvisioningStateSucceeded,
		EmailServicesProvisioningStateUnknown,
		EmailServicesProvisioningStateUpdating,
	}
}

// KeyType - The keyType to regenerate. Must be either 'primary' or 'secondary'(case-insensitive).
type KeyType string

const (
	KeyTypePrimary KeyType = "Primary"
	KeyTypeSecondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{	
		KeyTypePrimary,
		KeyTypeSecondary,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{	
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem Origin = "system"
	OriginUser Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{	
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - Provisioning state of the resource. Unknown is the default state for Communication Services.
type ProvisioningState string

const (
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	ProvisioningStateCreating ProvisioningState = "Creating"
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	ProvisioningStateFailed ProvisioningState = "Failed"
	ProvisioningStateMoving ProvisioningState = "Moving"
	ProvisioningStateRunning ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown ProvisioningState = "Unknown"
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{	
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
		ProvisioningStateUpdating,
	}
}

// UserEngagementTracking - Describes whether user engagement tracking is enabled or disabled.
type UserEngagementTracking string

const (
	UserEngagementTrackingDisabled UserEngagementTracking = "Disabled"
	UserEngagementTrackingEnabled UserEngagementTracking = "Enabled"
)

// PossibleUserEngagementTrackingValues returns the possible values for the UserEngagementTracking const type.
func PossibleUserEngagementTrackingValues() []UserEngagementTracking {
	return []UserEngagementTracking{	
		UserEngagementTrackingDisabled,
		UserEngagementTrackingEnabled,
	}
}

// VerificationStatus - Status of the verification operation.
type VerificationStatus string

const (
	VerificationStatusCancellationRequested VerificationStatus = "CancellationRequested"
	VerificationStatusNotStarted VerificationStatus = "NotStarted"
	VerificationStatusVerificationFailed VerificationStatus = "VerificationFailed"
	VerificationStatusVerificationInProgress VerificationStatus = "VerificationInProgress"
	VerificationStatusVerificationRequested VerificationStatus = "VerificationRequested"
	VerificationStatusVerified VerificationStatus = "Verified"
)

// PossibleVerificationStatusValues returns the possible values for the VerificationStatus const type.
func PossibleVerificationStatusValues() []VerificationStatus {
	return []VerificationStatus{	
		VerificationStatusCancellationRequested,
		VerificationStatusNotStarted,
		VerificationStatusVerificationFailed,
		VerificationStatusVerificationInProgress,
		VerificationStatusVerificationRequested,
		VerificationStatusVerified,
	}
}

// VerificationType - Type of verification.
type VerificationType string

const (
	VerificationTypeDKIM VerificationType = "DKIM"
	VerificationTypeDKIM2 VerificationType = "DKIM2"
	VerificationTypeDMARC VerificationType = "DMARC"
	VerificationTypeDomain VerificationType = "Domain"
	VerificationTypeSPF VerificationType = "SPF"
)

// PossibleVerificationTypeValues returns the possible values for the VerificationType const type.
func PossibleVerificationTypeValues() []VerificationType {
	return []VerificationType{	
		VerificationTypeDKIM,
		VerificationTypeDKIM2,
		VerificationTypeDMARC,
		VerificationTypeDomain,
		VerificationTypeSPF,
	}
}

