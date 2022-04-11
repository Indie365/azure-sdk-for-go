//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight

const (
	moduleName    = "armhdinsight"
	moduleVersion = "v0.3.0"
)

// AsyncOperationState - The async operation state.
type AsyncOperationState string

const (
	AsyncOperationStateFailed     AsyncOperationState = "Failed"
	AsyncOperationStateInProgress AsyncOperationState = "InProgress"
	AsyncOperationStateSucceeded  AsyncOperationState = "Succeeded"
)

// PossibleAsyncOperationStateValues returns the possible values for the AsyncOperationState const type.
func PossibleAsyncOperationStateValues() []AsyncOperationState {
	return []AsyncOperationState{
		AsyncOperationStateFailed,
		AsyncOperationStateInProgress,
		AsyncOperationStateSucceeded,
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

type DaysOfWeek string

const (
	DaysOfWeekFriday    DaysOfWeek = "Friday"
	DaysOfWeekMonday    DaysOfWeek = "Monday"
	DaysOfWeekSaturday  DaysOfWeek = "Saturday"
	DaysOfWeekSunday    DaysOfWeek = "Sunday"
	DaysOfWeekThursday  DaysOfWeek = "Thursday"
	DaysOfWeekTuesday   DaysOfWeek = "Tuesday"
	DaysOfWeekWednesday DaysOfWeek = "Wednesday"
)

// PossibleDaysOfWeekValues returns the possible values for the DaysOfWeek const type.
func PossibleDaysOfWeekValues() []DaysOfWeek {
	return []DaysOfWeek{
		DaysOfWeekFriday,
		DaysOfWeekMonday,
		DaysOfWeekSaturday,
		DaysOfWeekSunday,
		DaysOfWeekThursday,
		DaysOfWeekTuesday,
		DaysOfWeekWednesday,
	}
}

// DirectoryType - The directory type.
type DirectoryType string

const (
	DirectoryTypeActiveDirectory DirectoryType = "ActiveDirectory"
)

// PossibleDirectoryTypeValues returns the possible values for the DirectoryType const type.
func PossibleDirectoryTypeValues() []DirectoryType {
	return []DirectoryType{
		DirectoryTypeActiveDirectory,
	}
}

// FilterMode - The filtering mode. Effectively this can enabling or disabling the VM sizes in a particular set.
type FilterMode string

const (
	FilterModeDefault   FilterMode = "Default"
	FilterModeExclude   FilterMode = "Exclude"
	FilterModeInclude   FilterMode = "Include"
	FilterModeRecommend FilterMode = "Recommend"
)

// PossibleFilterModeValues returns the possible values for the FilterMode const type.
func PossibleFilterModeValues() []FilterMode {
	return []FilterMode{
		FilterModeDefault,
		FilterModeExclude,
		FilterModeInclude,
		FilterModeRecommend,
	}
}

// HDInsightClusterProvisioningState - The provisioning state, which only appears in the response.
type HDInsightClusterProvisioningState string

const (
	HDInsightClusterProvisioningStateCanceled   HDInsightClusterProvisioningState = "Canceled"
	HDInsightClusterProvisioningStateDeleting   HDInsightClusterProvisioningState = "Deleting"
	HDInsightClusterProvisioningStateFailed     HDInsightClusterProvisioningState = "Failed"
	HDInsightClusterProvisioningStateInProgress HDInsightClusterProvisioningState = "InProgress"
	HDInsightClusterProvisioningStateSucceeded  HDInsightClusterProvisioningState = "Succeeded"
)

// PossibleHDInsightClusterProvisioningStateValues returns the possible values for the HDInsightClusterProvisioningState const type.
func PossibleHDInsightClusterProvisioningStateValues() []HDInsightClusterProvisioningState {
	return []HDInsightClusterProvisioningState{
		HDInsightClusterProvisioningStateCanceled,
		HDInsightClusterProvisioningStateDeleting,
		HDInsightClusterProvisioningStateFailed,
		HDInsightClusterProvisioningStateInProgress,
		HDInsightClusterProvisioningStateSucceeded,
	}
}

// JSONWebKeyEncryptionAlgorithm - Algorithm identifier for encryption, default RSA-OAEP.
type JSONWebKeyEncryptionAlgorithm string

const (
	JSONWebKeyEncryptionAlgorithmRSA15      JSONWebKeyEncryptionAlgorithm = "RSA1_5"
	JSONWebKeyEncryptionAlgorithmRSAOAEP    JSONWebKeyEncryptionAlgorithm = "RSA-OAEP"
	JSONWebKeyEncryptionAlgorithmRSAOAEP256 JSONWebKeyEncryptionAlgorithm = "RSA-OAEP-256"
)

// PossibleJSONWebKeyEncryptionAlgorithmValues returns the possible values for the JSONWebKeyEncryptionAlgorithm const type.
func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
	return []JSONWebKeyEncryptionAlgorithm{
		JSONWebKeyEncryptionAlgorithmRSA15,
		JSONWebKeyEncryptionAlgorithmRSAOAEP,
		JSONWebKeyEncryptionAlgorithmRSAOAEP256,
	}
}

// OSType - The type of operating system.
type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

// PossibleOSTypeValues returns the possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{
		OSTypeLinux,
		OSTypeWindows,
	}
}

// PrivateEndpointConnectionProvisioningState - The provisioning state, which only appears in the response.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCanceled   PrivateEndpointConnectionProvisioningState = "Canceled"
	PrivateEndpointConnectionProvisioningStateDeleting   PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed     PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateInProgress PrivateEndpointConnectionProvisioningState = "InProgress"
	PrivateEndpointConnectionProvisioningStateSucceeded  PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating   PrivateEndpointConnectionProvisioningState = "Updating"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCanceled,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateInProgress,
		PrivateEndpointConnectionProvisioningStateSucceeded,
		PrivateEndpointConnectionProvisioningStateUpdating,
	}
}

// PrivateIPAllocationMethod - The method that private IP address is allocated.
type PrivateIPAllocationMethod string

const (
	PrivateIPAllocationMethodDynamic PrivateIPAllocationMethod = "dynamic"
	PrivateIPAllocationMethodStatic  PrivateIPAllocationMethod = "static"
)

// PossiblePrivateIPAllocationMethodValues returns the possible values for the PrivateIPAllocationMethod const type.
func PossiblePrivateIPAllocationMethodValues() []PrivateIPAllocationMethod {
	return []PrivateIPAllocationMethod{
		PrivateIPAllocationMethodDynamic,
		PrivateIPAllocationMethodStatic,
	}
}

// PrivateLink - Indicates whether or not private link is enabled.
type PrivateLink string

const (
	PrivateLinkDisabled PrivateLink = "Disabled"
	PrivateLinkEnabled  PrivateLink = "Enabled"
)

// PossiblePrivateLinkValues returns the possible values for the PrivateLink const type.
func PossiblePrivateLinkValues() []PrivateLink {
	return []PrivateLink{
		PrivateLinkDisabled,
		PrivateLinkEnabled,
	}
}

// PrivateLinkConfigurationProvisioningState - The private link configuration provisioning state, which only appears in the
// response.
type PrivateLinkConfigurationProvisioningState string

const (
	PrivateLinkConfigurationProvisioningStateCanceled   PrivateLinkConfigurationProvisioningState = "Canceled"
	PrivateLinkConfigurationProvisioningStateDeleting   PrivateLinkConfigurationProvisioningState = "Deleting"
	PrivateLinkConfigurationProvisioningStateFailed     PrivateLinkConfigurationProvisioningState = "Failed"
	PrivateLinkConfigurationProvisioningStateInProgress PrivateLinkConfigurationProvisioningState = "InProgress"
	PrivateLinkConfigurationProvisioningStateSucceeded  PrivateLinkConfigurationProvisioningState = "Succeeded"
)

// PossiblePrivateLinkConfigurationProvisioningStateValues returns the possible values for the PrivateLinkConfigurationProvisioningState const type.
func PossiblePrivateLinkConfigurationProvisioningStateValues() []PrivateLinkConfigurationProvisioningState {
	return []PrivateLinkConfigurationProvisioningState{
		PrivateLinkConfigurationProvisioningStateCanceled,
		PrivateLinkConfigurationProvisioningStateDeleting,
		PrivateLinkConfigurationProvisioningStateFailed,
		PrivateLinkConfigurationProvisioningStateInProgress,
		PrivateLinkConfigurationProvisioningStateSucceeded,
	}
}

// PrivateLinkServiceConnectionStatus - The concrete private link service connection.
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusPending  PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected PrivateLinkServiceConnectionStatus = "Rejected"
	PrivateLinkServiceConnectionStatusRemoved  PrivateLinkServiceConnectionStatus = "Removed"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
		PrivateLinkServiceConnectionStatusRemoved,
	}
}

// ResourceIdentityType - The type of identity used for the cluster. The type 'SystemAssigned, UserAssigned' includes both
// an implicitly created identity and a set of user assigned identities.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// ResourceProviderConnection - The direction for the resource provider connection.
type ResourceProviderConnection string

const (
	ResourceProviderConnectionInbound  ResourceProviderConnection = "Inbound"
	ResourceProviderConnectionOutbound ResourceProviderConnection = "Outbound"
)

// PossibleResourceProviderConnectionValues returns the possible values for the ResourceProviderConnection const type.
func PossibleResourceProviderConnectionValues() []ResourceProviderConnection {
	return []ResourceProviderConnection{
		ResourceProviderConnectionInbound,
		ResourceProviderConnectionOutbound,
	}
}

type RoleName string

const (
	RoleNameWorkernode RoleName = "workernode"
)

// PossibleRoleNameValues returns the possible values for the RoleName const type.
func PossibleRoleNameValues() []RoleName {
	return []RoleName{
		RoleNameWorkernode,
	}
}

// Tier - The cluster tier.
type Tier string

const (
	TierPremium  Tier = "Premium"
	TierStandard Tier = "Standard"
)

// PossibleTierValues returns the possible values for the Tier const type.
func PossibleTierValues() []Tier {
	return []Tier{
		TierPremium,
		TierStandard,
	}
}
