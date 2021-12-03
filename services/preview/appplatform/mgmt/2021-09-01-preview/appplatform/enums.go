package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AppResourceProvisioningState enumerates the values for app resource provisioning state.
type AppResourceProvisioningState string

const (
	// AppResourceProvisioningStateCreating ...
	AppResourceProvisioningStateCreating AppResourceProvisioningState = "Creating"
	// AppResourceProvisioningStateFailed ...
	AppResourceProvisioningStateFailed AppResourceProvisioningState = "Failed"
	// AppResourceProvisioningStateSucceeded ...
	AppResourceProvisioningStateSucceeded AppResourceProvisioningState = "Succeeded"
	// AppResourceProvisioningStateUpdating ...
	AppResourceProvisioningStateUpdating AppResourceProvisioningState = "Updating"
)

// PossibleAppResourceProvisioningStateValues returns an array of possible values for the AppResourceProvisioningState const type.
func PossibleAppResourceProvisioningStateValues() []AppResourceProvisioningState {
	return []AppResourceProvisioningState{AppResourceProvisioningStateCreating, AppResourceProvisioningStateFailed, AppResourceProvisioningStateSucceeded, AppResourceProvisioningStateUpdating}
}

// ConfigServerState enumerates the values for config server state.
type ConfigServerState string

const (
	// ConfigServerStateDeleted ...
	ConfigServerStateDeleted ConfigServerState = "Deleted"
	// ConfigServerStateFailed ...
	ConfigServerStateFailed ConfigServerState = "Failed"
	// ConfigServerStateNotAvailable ...
	ConfigServerStateNotAvailable ConfigServerState = "NotAvailable"
	// ConfigServerStateSucceeded ...
	ConfigServerStateSucceeded ConfigServerState = "Succeeded"
	// ConfigServerStateUpdating ...
	ConfigServerStateUpdating ConfigServerState = "Updating"
)

// PossibleConfigServerStateValues returns an array of possible values for the ConfigServerState const type.
func PossibleConfigServerStateValues() []ConfigServerState {
	return []ConfigServerState{ConfigServerStateDeleted, ConfigServerStateFailed, ConfigServerStateNotAvailable, ConfigServerStateSucceeded, ConfigServerStateUpdating}
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

// DeploymentResourceProvisioningState enumerates the values for deployment resource provisioning state.
type DeploymentResourceProvisioningState string

const (
	// DeploymentResourceProvisioningStateCreating ...
	DeploymentResourceProvisioningStateCreating DeploymentResourceProvisioningState = "Creating"
	// DeploymentResourceProvisioningStateFailed ...
	DeploymentResourceProvisioningStateFailed DeploymentResourceProvisioningState = "Failed"
	// DeploymentResourceProvisioningStateSucceeded ...
	DeploymentResourceProvisioningStateSucceeded DeploymentResourceProvisioningState = "Succeeded"
	// DeploymentResourceProvisioningStateUpdating ...
	DeploymentResourceProvisioningStateUpdating DeploymentResourceProvisioningState = "Updating"
)

// PossibleDeploymentResourceProvisioningStateValues returns an array of possible values for the DeploymentResourceProvisioningState const type.
func PossibleDeploymentResourceProvisioningStateValues() []DeploymentResourceProvisioningState {
	return []DeploymentResourceProvisioningState{DeploymentResourceProvisioningStateCreating, DeploymentResourceProvisioningStateFailed, DeploymentResourceProvisioningStateSucceeded, DeploymentResourceProvisioningStateUpdating}
}

// DeploymentResourceStatus enumerates the values for deployment resource status.
type DeploymentResourceStatus string

const (
	// DeploymentResourceStatusAllocating ...
	DeploymentResourceStatusAllocating DeploymentResourceStatus = "Allocating"
	// DeploymentResourceStatusCompiling ...
	DeploymentResourceStatusCompiling DeploymentResourceStatus = "Compiling"
	// DeploymentResourceStatusFailed ...
	DeploymentResourceStatusFailed DeploymentResourceStatus = "Failed"
	// DeploymentResourceStatusRunning ...
	DeploymentResourceStatusRunning DeploymentResourceStatus = "Running"
	// DeploymentResourceStatusStopped ...
	DeploymentResourceStatusStopped DeploymentResourceStatus = "Stopped"
	// DeploymentResourceStatusUnknown ...
	DeploymentResourceStatusUnknown DeploymentResourceStatus = "Unknown"
	// DeploymentResourceStatusUpgrading ...
	DeploymentResourceStatusUpgrading DeploymentResourceStatus = "Upgrading"
)

// PossibleDeploymentResourceStatusValues returns an array of possible values for the DeploymentResourceStatus const type.
func PossibleDeploymentResourceStatusValues() []DeploymentResourceStatus {
	return []DeploymentResourceStatus{DeploymentResourceStatusAllocating, DeploymentResourceStatusCompiling, DeploymentResourceStatusFailed, DeploymentResourceStatusRunning, DeploymentResourceStatusStopped, DeploymentResourceStatusUnknown, DeploymentResourceStatusUpgrading}
}

// ManagedIdentityType enumerates the values for managed identity type.
type ManagedIdentityType string

const (
	// ManagedIdentityTypeNone ...
	ManagedIdentityTypeNone ManagedIdentityType = "None"
	// ManagedIdentityTypeSystemAssigned ...
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
	// ManagedIdentityTypeSystemAssignedUserAssigned ...
	ManagedIdentityTypeSystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	// ManagedIdentityTypeUserAssigned ...
	ManagedIdentityTypeUserAssigned ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns an array of possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{ManagedIdentityTypeNone, ManagedIdentityTypeSystemAssigned, ManagedIdentityTypeSystemAssignedUserAssigned, ManagedIdentityTypeUserAssigned}
}

// MonitoringSettingState enumerates the values for monitoring setting state.
type MonitoringSettingState string

const (
	// MonitoringSettingStateFailed ...
	MonitoringSettingStateFailed MonitoringSettingState = "Failed"
	// MonitoringSettingStateNotAvailable ...
	MonitoringSettingStateNotAvailable MonitoringSettingState = "NotAvailable"
	// MonitoringSettingStateSucceeded ...
	MonitoringSettingStateSucceeded MonitoringSettingState = "Succeeded"
	// MonitoringSettingStateUpdating ...
	MonitoringSettingStateUpdating MonitoringSettingState = "Updating"
)

// PossibleMonitoringSettingStateValues returns an array of possible values for the MonitoringSettingState const type.
func PossibleMonitoringSettingStateValues() []MonitoringSettingState {
	return []MonitoringSettingState{MonitoringSettingStateFailed, MonitoringSettingStateNotAvailable, MonitoringSettingStateSucceeded, MonitoringSettingStateUpdating}
}

// PowerState enumerates the values for power state.
type PowerState string

const (
	// PowerStateRunning ...
	PowerStateRunning PowerState = "Running"
	// PowerStateStopped ...
	PowerStateStopped PowerState = "Stopped"
)

// PossiblePowerStateValues returns an array of possible values for the PowerState const type.
func PossiblePowerStateValues() []PowerState {
	return []PowerState{PowerStateRunning, PowerStateStopped}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMoved ...
	ProvisioningStateMoved ProvisioningState = "Moved"
	// ProvisioningStateMoveFailed ...
	ProvisioningStateMoveFailed ProvisioningState = "MoveFailed"
	// ProvisioningStateMoving ...
	ProvisioningStateMoving ProvisioningState = "Moving"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCreating, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateMoved, ProvisioningStateMoveFailed, ProvisioningStateMoving, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// ResourceSkuRestrictionsReasonCode enumerates the values for resource sku restrictions reason code.
type ResourceSkuRestrictionsReasonCode string

const (
	// ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ...
	ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	// ResourceSkuRestrictionsReasonCodeQuotaID ...
	ResourceSkuRestrictionsReasonCodeQuotaID ResourceSkuRestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSkuRestrictionsReasonCodeValues returns an array of possible values for the ResourceSkuRestrictionsReasonCode const type.
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return []ResourceSkuRestrictionsReasonCode{ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription, ResourceSkuRestrictionsReasonCodeQuotaID}
}

// ResourceSkuRestrictionsType enumerates the values for resource sku restrictions type.
type ResourceSkuRestrictionsType string

const (
	// ResourceSkuRestrictionsTypeLocation ...
	ResourceSkuRestrictionsTypeLocation ResourceSkuRestrictionsType = "Location"
	// ResourceSkuRestrictionsTypeZone ...
	ResourceSkuRestrictionsTypeZone ResourceSkuRestrictionsType = "Zone"
)

// PossibleResourceSkuRestrictionsTypeValues returns an array of possible values for the ResourceSkuRestrictionsType const type.
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return []ResourceSkuRestrictionsType{ResourceSkuRestrictionsTypeLocation, ResourceSkuRestrictionsTypeZone}
}

// RuntimeVersion enumerates the values for runtime version.
type RuntimeVersion string

const (
	// RuntimeVersionJava11 ...
	RuntimeVersionJava11 RuntimeVersion = "Java_11"
	// RuntimeVersionJava8 ...
	RuntimeVersionJava8 RuntimeVersion = "Java_8"
	// RuntimeVersionNetCore31 ...
	RuntimeVersionNetCore31 RuntimeVersion = "NetCore_31"
)

// PossibleRuntimeVersionValues returns an array of possible values for the RuntimeVersion const type.
func PossibleRuntimeVersionValues() []RuntimeVersion {
	return []RuntimeVersion{RuntimeVersionJava11, RuntimeVersionJava8, RuntimeVersionNetCore31}
}

// SkuScaleType enumerates the values for sku scale type.
type SkuScaleType string

const (
	// SkuScaleTypeAutomatic ...
	SkuScaleTypeAutomatic SkuScaleType = "Automatic"
	// SkuScaleTypeManual ...
	SkuScaleTypeManual SkuScaleType = "Manual"
	// SkuScaleTypeNone ...
	SkuScaleTypeNone SkuScaleType = "None"
)

// PossibleSkuScaleTypeValues returns an array of possible values for the SkuScaleType const type.
func PossibleSkuScaleTypeValues() []SkuScaleType {
	return []SkuScaleType{SkuScaleTypeAutomatic, SkuScaleTypeManual, SkuScaleTypeNone}
}

// StorageType enumerates the values for storage type.
type StorageType string

const (
	// StorageTypeStorageAccount ...
	StorageTypeStorageAccount StorageType = "StorageAccount"
	// StorageTypeStorageProperties ...
	StorageTypeStorageProperties StorageType = "StorageProperties"
)

// PossibleStorageTypeValues returns an array of possible values for the StorageType const type.
func PossibleStorageTypeValues() []StorageType {
	return []StorageType{StorageTypeStorageAccount, StorageTypeStorageProperties}
}

// SupportedRuntimePlatform enumerates the values for supported runtime platform.
type SupportedRuntimePlatform string

const (
	// SupportedRuntimePlatformJava ...
	SupportedRuntimePlatformJava SupportedRuntimePlatform = "Java"
	// SupportedRuntimePlatformNETCore ...
	SupportedRuntimePlatformNETCore SupportedRuntimePlatform = ".NET Core"
)

// PossibleSupportedRuntimePlatformValues returns an array of possible values for the SupportedRuntimePlatform const type.
func PossibleSupportedRuntimePlatformValues() []SupportedRuntimePlatform {
	return []SupportedRuntimePlatform{SupportedRuntimePlatformJava, SupportedRuntimePlatformNETCore}
}

// SupportedRuntimeValue enumerates the values for supported runtime value.
type SupportedRuntimeValue string

const (
	// SupportedRuntimeValueJava11 ...
	SupportedRuntimeValueJava11 SupportedRuntimeValue = "Java_11"
	// SupportedRuntimeValueJava8 ...
	SupportedRuntimeValueJava8 SupportedRuntimeValue = "Java_8"
	// SupportedRuntimeValueNetCore31 ...
	SupportedRuntimeValueNetCore31 SupportedRuntimeValue = "NetCore_31"
)

// PossibleSupportedRuntimeValueValues returns an array of possible values for the SupportedRuntimeValue const type.
func PossibleSupportedRuntimeValueValues() []SupportedRuntimeValue {
	return []SupportedRuntimeValue{SupportedRuntimeValueJava11, SupportedRuntimeValueJava8, SupportedRuntimeValueNetCore31}
}

// TestKeyType enumerates the values for test key type.
type TestKeyType string

const (
	// TestKeyTypePrimary ...
	TestKeyTypePrimary TestKeyType = "Primary"
	// TestKeyTypeSecondary ...
	TestKeyTypeSecondary TestKeyType = "Secondary"
)

// PossibleTestKeyTypeValues returns an array of possible values for the TestKeyType const type.
func PossibleTestKeyTypeValues() []TestKeyType {
	return []TestKeyType{TestKeyTypePrimary, TestKeyTypeSecondary}
}

// TrafficDirection enumerates the values for traffic direction.
type TrafficDirection string

const (
	// TrafficDirectionInbound ...
	TrafficDirectionInbound TrafficDirection = "Inbound"
	// TrafficDirectionOutbound ...
	TrafficDirectionOutbound TrafficDirection = "Outbound"
)

// PossibleTrafficDirectionValues returns an array of possible values for the TrafficDirection const type.
func PossibleTrafficDirectionValues() []TrafficDirection {
	return []TrafficDirection{TrafficDirectionInbound, TrafficDirectionOutbound}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeAzureFileVolume ...
	TypeAzureFileVolume Type = "AzureFileVolume"
	// TypeCustomPersistentDiskProperties ...
	TypeCustomPersistentDiskProperties Type = "CustomPersistentDiskProperties"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeAzureFileVolume, TypeCustomPersistentDiskProperties}
}

// TypeBasicCertificateProperties enumerates the values for type basic certificate properties.
type TypeBasicCertificateProperties string

const (
	// TypeBasicCertificatePropertiesTypeCertificateProperties ...
	TypeBasicCertificatePropertiesTypeCertificateProperties TypeBasicCertificateProperties = "CertificateProperties"
	// TypeBasicCertificatePropertiesTypeContentCertificate ...
	TypeBasicCertificatePropertiesTypeContentCertificate TypeBasicCertificateProperties = "ContentCertificate"
	// TypeBasicCertificatePropertiesTypeKeyVaultCertificate ...
	TypeBasicCertificatePropertiesTypeKeyVaultCertificate TypeBasicCertificateProperties = "KeyVaultCertificate"
)

// PossibleTypeBasicCertificatePropertiesValues returns an array of possible values for the TypeBasicCertificateProperties const type.
func PossibleTypeBasicCertificatePropertiesValues() []TypeBasicCertificateProperties {
	return []TypeBasicCertificateProperties{TypeBasicCertificatePropertiesTypeCertificateProperties, TypeBasicCertificatePropertiesTypeContentCertificate, TypeBasicCertificatePropertiesTypeKeyVaultCertificate}
}

// UserSourceType enumerates the values for user source type.
type UserSourceType string

const (
	// UserSourceTypeContainer ...
	UserSourceTypeContainer UserSourceType = "Container"
	// UserSourceTypeJar ...
	UserSourceTypeJar UserSourceType = "Jar"
	// UserSourceTypeNetCoreZip ...
	UserSourceTypeNetCoreZip UserSourceType = "NetCoreZip"
	// UserSourceTypeSource ...
	UserSourceTypeSource UserSourceType = "Source"
)

// PossibleUserSourceTypeValues returns an array of possible values for the UserSourceType const type.
func PossibleUserSourceTypeValues() []UserSourceType {
	return []UserSourceType{UserSourceTypeContainer, UserSourceTypeJar, UserSourceTypeNetCoreZip, UserSourceTypeSource}
}
