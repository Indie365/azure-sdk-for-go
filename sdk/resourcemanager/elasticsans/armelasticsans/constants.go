//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsans

const (
	moduleName    = "armelasticsans"
	moduleVersion = "v0.1.1"
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

// EncryptionType - The type of key used to encrypt the data of the disk.
type EncryptionType string

const (
	// EncryptionTypeEncryptionAtRestWithPlatformKey - Volume is encrypted at rest with Platform managed key. It is the default
	// encryption type.
	EncryptionTypeEncryptionAtRestWithPlatformKey EncryptionType = "EncryptionAtRestWithPlatformKey"
)

// PossibleEncryptionTypeValues returns the possible values for the EncryptionType const type.
func PossibleEncryptionTypeValues() []EncryptionType {
	return []EncryptionType{
		EncryptionTypeEncryptionAtRestWithPlatformKey,
	}
}

// OperationalStatus - Operational status of the resource.
type OperationalStatus string

const (
	OperationalStatusHealthy            OperationalStatus = "Healthy"
	OperationalStatusInvalid            OperationalStatus = "Invalid"
	OperationalStatusRunning            OperationalStatus = "Running"
	OperationalStatusStopped            OperationalStatus = "Stopped"
	OperationalStatusStoppedDeallocated OperationalStatus = "Stopped (deallocated)"
	OperationalStatusUnhealthy          OperationalStatus = "Unhealthy"
	OperationalStatusUnknown            OperationalStatus = "Unknown"
	OperationalStatusUpdating           OperationalStatus = "Updating"
)

// PossibleOperationalStatusValues returns the possible values for the OperationalStatus const type.
func PossibleOperationalStatusValues() []OperationalStatus {
	return []OperationalStatus{
		OperationalStatusHealthy,
		OperationalStatusInvalid,
		OperationalStatusRunning,
		OperationalStatusStopped,
		OperationalStatusStoppedDeallocated,
		OperationalStatusUnhealthy,
		OperationalStatusUnknown,
		OperationalStatusUpdating,
	}
}

// ProvisioningStates - Provisioning state of the iSCSI Target.
type ProvisioningStates string

const (
	ProvisioningStatesCanceled  ProvisioningStates = "Canceled"
	ProvisioningStatesCreating  ProvisioningStates = "Creating"
	ProvisioningStatesDeleting  ProvisioningStates = "Deleting"
	ProvisioningStatesFailed    ProvisioningStates = "Failed"
	ProvisioningStatesInvalid   ProvisioningStates = "Invalid"
	ProvisioningStatesPending   ProvisioningStates = "Pending"
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
	ProvisioningStatesUpdating  ProvisioningStates = "Updating"
)

// PossibleProvisioningStatesValues returns the possible values for the ProvisioningStates const type.
func PossibleProvisioningStatesValues() []ProvisioningStates {
	return []ProvisioningStates{
		ProvisioningStatesCanceled,
		ProvisioningStatesCreating,
		ProvisioningStatesDeleting,
		ProvisioningStatesFailed,
		ProvisioningStatesInvalid,
		ProvisioningStatesPending,
		ProvisioningStatesSucceeded,
		ProvisioningStatesUpdating,
	}
}

// SKUName - The sku name.
type SKUName string

const (
	// SKUNamePremiumLRS - Premium locally redundant storage
	SKUNamePremiumLRS SKUName = "Premium_LRS"
	// SKUNamePremiumZRS - Premium zone redundant storage
	SKUNamePremiumZRS SKUName = "Premium_ZRS"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePremiumLRS,
		SKUNamePremiumZRS,
	}
}

// SKUTier - The sku tier.
type SKUTier string

const (
	// SKUTierPremium - Premium Tier
	SKUTierPremium SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierPremium,
	}
}

// State - Gets the state of virtual network rule.
type State string

const (
	StateProvisioning         State = "provisioning"
	StateDeprovisioning       State = "deprovisioning"
	StateSucceeded            State = "succeeded"
	StateFailed               State = "failed"
	StateNetworkSourceDeleted State = "networkSourceDeleted"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateProvisioning,
		StateDeprovisioning,
		StateSucceeded,
		StateFailed,
		StateNetworkSourceDeleted,
	}
}

// StorageTargetType - Storage Target type.
type StorageTargetType string

const (
	StorageTargetTypeIscsi StorageTargetType = "Iscsi"
	StorageTargetTypeNone  StorageTargetType = "None"
)

// PossibleStorageTargetTypeValues returns the possible values for the StorageTargetType const type.
func PossibleStorageTargetTypeValues() []StorageTargetType {
	return []StorageTargetType{
		StorageTargetTypeIscsi,
		StorageTargetTypeNone,
	}
}
