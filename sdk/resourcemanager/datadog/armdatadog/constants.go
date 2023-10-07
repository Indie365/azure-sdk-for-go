//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

const (
	moduleName    = "armdatadog"
	moduleVersion = "v1.2.0"
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

type LiftrResourceCategories string

const (
	LiftrResourceCategoriesMonitorLogs LiftrResourceCategories = "MonitorLogs"
	LiftrResourceCategoriesUnknown     LiftrResourceCategories = "Unknown"
)

// PossibleLiftrResourceCategoriesValues returns the possible values for the LiftrResourceCategories const type.
func PossibleLiftrResourceCategoriesValues() []LiftrResourceCategories {
	return []LiftrResourceCategories{
		LiftrResourceCategoriesMonitorLogs,
		LiftrResourceCategoriesUnknown,
	}
}

// ManagedIdentityTypes - Specifies the identity type of the Datadog Monitor. At this time the only allowed value is 'SystemAssigned'.
type ManagedIdentityTypes string

const (
	ManagedIdentityTypesSystemAssigned ManagedIdentityTypes = "SystemAssigned"
	ManagedIdentityTypesUserAssigned   ManagedIdentityTypes = "UserAssigned"
)

// PossibleManagedIdentityTypesValues returns the possible values for the ManagedIdentityTypes const type.
func PossibleManagedIdentityTypesValues() []ManagedIdentityTypes {
	return []ManagedIdentityTypes{
		ManagedIdentityTypesSystemAssigned,
		ManagedIdentityTypesUserAssigned,
	}
}

// MarketplaceSubscriptionStatus - Flag specifying the Marketplace Subscription Status of the resource. If payment is not
// made in time, the resource will go in Suspended state.
type MarketplaceSubscriptionStatus string

const (
	MarketplaceSubscriptionStatusActive       MarketplaceSubscriptionStatus = "Active"
	MarketplaceSubscriptionStatusProvisioning MarketplaceSubscriptionStatus = "Provisioning"
	MarketplaceSubscriptionStatusSuspended    MarketplaceSubscriptionStatus = "Suspended"
	MarketplaceSubscriptionStatusUnsubscribed MarketplaceSubscriptionStatus = "Unsubscribed"
)

// PossibleMarketplaceSubscriptionStatusValues returns the possible values for the MarketplaceSubscriptionStatus const type.
func PossibleMarketplaceSubscriptionStatusValues() []MarketplaceSubscriptionStatus {
	return []MarketplaceSubscriptionStatus{
		MarketplaceSubscriptionStatusActive,
		MarketplaceSubscriptionStatusProvisioning,
		MarketplaceSubscriptionStatusSuspended,
		MarketplaceSubscriptionStatusUnsubscribed,
	}
}

// MonitoringStatus - Flag specifying if the resource monitoring is enabled or disabled.
type MonitoringStatus string

const (
	MonitoringStatusDisabled MonitoringStatus = "Disabled"
	MonitoringStatusEnabled  MonitoringStatus = "Enabled"
)

// PossibleMonitoringStatusValues returns the possible values for the MonitoringStatus const type.
func PossibleMonitoringStatusValues() []MonitoringStatus {
	return []MonitoringStatus{
		MonitoringStatusDisabled,
		MonitoringStatusEnabled,
	}
}

// Operation - The operation for the patch on the resource.
type Operation string

const (
	OperationActive         Operation = "Active"
	OperationAddBegin       Operation = "AddBegin"
	OperationAddComplete    Operation = "AddComplete"
	OperationDeleteBegin    Operation = "DeleteBegin"
	OperationDeleteComplete Operation = "DeleteComplete"
)

// PossibleOperationValues returns the possible values for the Operation const type.
func PossibleOperationValues() []Operation {
	return []Operation{
		OperationActive,
		OperationAddBegin,
		OperationAddComplete,
		OperationDeleteBegin,
		OperationDeleteComplete,
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateNotSpecified,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// SingleSignOnStates - Various states of the SSO resource
type SingleSignOnStates string

const (
	SingleSignOnStatesDisable  SingleSignOnStates = "Disable"
	SingleSignOnStatesEnable   SingleSignOnStates = "Enable"
	SingleSignOnStatesExisting SingleSignOnStates = "Existing"
	SingleSignOnStatesInitial  SingleSignOnStates = "Initial"
)

// PossibleSingleSignOnStatesValues returns the possible values for the SingleSignOnStates const type.
func PossibleSingleSignOnStatesValues() []SingleSignOnStates {
	return []SingleSignOnStates{
		SingleSignOnStatesDisable,
		SingleSignOnStatesEnable,
		SingleSignOnStatesExisting,
		SingleSignOnStatesInitial,
	}
}

// Status - The state of monitoring.
type Status string

const (
	StatusActive     Status = "Active"
	StatusDeleting   Status = "Deleting"
	StatusFailed     Status = "Failed"
	StatusInProgress Status = "InProgress"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusActive,
		StatusDeleting,
		StatusFailed,
		StatusInProgress,
	}
}

// TagAction - Valid actions for a filtering tag. Exclusion takes priority over inclusion.
type TagAction string

const (
	TagActionExclude TagAction = "Exclude"
	TagActionInclude TagAction = "Include"
)

// PossibleTagActionValues returns the possible values for the TagAction const type.
func PossibleTagActionValues() []TagAction {
	return []TagAction{
		TagActionExclude,
		TagActionInclude,
	}
}
