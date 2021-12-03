package netapp

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActiveDirectoryStatus enumerates the values for active directory status.
type ActiveDirectoryStatus string

const (
	// ActiveDirectoryStatusCreated Active Directory created but not in use
	ActiveDirectoryStatusCreated ActiveDirectoryStatus = "Created"
	// ActiveDirectoryStatusDeleted Active Directory Deleted
	ActiveDirectoryStatusDeleted ActiveDirectoryStatus = "Deleted"
	// ActiveDirectoryStatusError Error with the Active Directory
	ActiveDirectoryStatusError ActiveDirectoryStatus = "Error"
	// ActiveDirectoryStatusInUse Active Directory in use by SMB Volume
	ActiveDirectoryStatusInUse ActiveDirectoryStatus = "InUse"
	// ActiveDirectoryStatusUpdating Active Directory Updating
	ActiveDirectoryStatusUpdating ActiveDirectoryStatus = "Updating"
)

// PossibleActiveDirectoryStatusValues returns an array of possible values for the ActiveDirectoryStatus const type.
func PossibleActiveDirectoryStatusValues() []ActiveDirectoryStatus {
	return []ActiveDirectoryStatus{ActiveDirectoryStatusCreated, ActiveDirectoryStatusDeleted, ActiveDirectoryStatusError, ActiveDirectoryStatusInUse, ActiveDirectoryStatusUpdating}
}

// ApplicationType enumerates the values for application type.
type ApplicationType string

const (
	// ApplicationTypeSAPHANA ...
	ApplicationTypeSAPHANA ApplicationType = "SAP-HANA"
)

// PossibleApplicationTypeValues returns an array of possible values for the ApplicationType const type.
func PossibleApplicationTypeValues() []ApplicationType {
	return []ApplicationType{ApplicationTypeSAPHANA}
}

// AvsDataStore enumerates the values for avs data store.
type AvsDataStore string

const (
	// AvsDataStoreDisabled avsDataStore is disabled
	AvsDataStoreDisabled AvsDataStore = "Disabled"
	// AvsDataStoreEnabled avsDataStore is enabled
	AvsDataStoreEnabled AvsDataStore = "Enabled"
)

// PossibleAvsDataStoreValues returns an array of possible values for the AvsDataStore const type.
func PossibleAvsDataStoreValues() []AvsDataStore {
	return []AvsDataStore{AvsDataStoreDisabled, AvsDataStoreEnabled}
}

// BackupType enumerates the values for backup type.
type BackupType string

const (
	// BackupTypeManual Manual backup
	BackupTypeManual BackupType = "Manual"
	// BackupTypeScheduled Scheduled backup
	BackupTypeScheduled BackupType = "Scheduled"
)

// PossibleBackupTypeValues returns an array of possible values for the BackupType const type.
func PossibleBackupTypeValues() []BackupType {
	return []BackupType{BackupTypeManual, BackupTypeScheduled}
}

// CheckNameResourceTypes enumerates the values for check name resource types.
type CheckNameResourceTypes string

const (
	// CheckNameResourceTypesMicrosoftNetAppnetAppAccounts ...
	CheckNameResourceTypesMicrosoftNetAppnetAppAccounts CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	// CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools ...
	CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	// CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes ...
	CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	// CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots ...
	CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

// PossibleCheckNameResourceTypesValues returns an array of possible values for the CheckNameResourceTypes const type.
func PossibleCheckNameResourceTypesValues() []CheckNameResourceTypes {
	return []CheckNameResourceTypes{CheckNameResourceTypesMicrosoftNetAppnetAppAccounts, CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools, CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes, CheckNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots}
}

// CheckQuotaNameResourceTypes enumerates the values for check quota name resource types.
type CheckQuotaNameResourceTypes string

const (
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts"
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools"
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes"
	// CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots ...
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckQuotaNameResourceTypes = "Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"
)

// PossibleCheckQuotaNameResourceTypesValues returns an array of possible values for the CheckQuotaNameResourceTypes const type.
func PossibleCheckQuotaNameResourceTypesValues() []CheckQuotaNameResourceTypes {
	return []CheckQuotaNameResourceTypes{CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts, CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools, CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes, CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots}
}

// ChownMode enumerates the values for chown mode.
type ChownMode string

const (
	// ChownModeRestricted ...
	ChownModeRestricted ChownMode = "Restricted"
	// ChownModeUnrestricted ...
	ChownModeUnrestricted ChownMode = "Unrestricted"
)

// PossibleChownModeValues returns an array of possible values for the ChownMode const type.
func PossibleChownModeValues() []ChownMode {
	return []ChownMode{ChownModeRestricted, ChownModeUnrestricted}
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

// EncryptionType enumerates the values for encryption type.
type EncryptionType string

const (
	// EncryptionTypeDouble EncryptionType Double, volumes will use double encryption at rest
	EncryptionTypeDouble EncryptionType = "Double"
	// EncryptionTypeSingle EncryptionType Single, volumes will use single encryption at rest
	EncryptionTypeSingle EncryptionType = "Single"
)

// PossibleEncryptionTypeValues returns an array of possible values for the EncryptionType const type.
func PossibleEncryptionTypeValues() []EncryptionType {
	return []EncryptionType{EncryptionTypeDouble, EncryptionTypeSingle}
}

// EndpointType enumerates the values for endpoint type.
type EndpointType string

const (
	// EndpointTypeDst ...
	EndpointTypeDst EndpointType = "dst"
	// EndpointTypeSrc ...
	EndpointTypeSrc EndpointType = "src"
)

// PossibleEndpointTypeValues returns an array of possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{EndpointTypeDst, EndpointTypeSrc}
}

// InAvailabilityReasonType enumerates the values for in availability reason type.
type InAvailabilityReasonType string

const (
	// InAvailabilityReasonTypeAlreadyExists ...
	InAvailabilityReasonTypeAlreadyExists InAvailabilityReasonType = "AlreadyExists"
	// InAvailabilityReasonTypeInvalid ...
	InAvailabilityReasonTypeInvalid InAvailabilityReasonType = "Invalid"
)

// PossibleInAvailabilityReasonTypeValues returns an array of possible values for the InAvailabilityReasonType const type.
func PossibleInAvailabilityReasonTypeValues() []InAvailabilityReasonType {
	return []InAvailabilityReasonType{InAvailabilityReasonTypeAlreadyExists, InAvailabilityReasonTypeInvalid}
}

// MetricAggregationType enumerates the values for metric aggregation type.
type MetricAggregationType string

const (
	// MetricAggregationTypeAverage ...
	MetricAggregationTypeAverage MetricAggregationType = "Average"
)

// PossibleMetricAggregationTypeValues returns an array of possible values for the MetricAggregationType const type.
func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return []MetricAggregationType{MetricAggregationTypeAverage}
}

// MirrorState enumerates the values for mirror state.
type MirrorState string

const (
	// MirrorStateBroken ...
	MirrorStateBroken MirrorState = "Broken"
	// MirrorStateMirrored ...
	MirrorStateMirrored MirrorState = "Mirrored"
	// MirrorStateUninitialized ...
	MirrorStateUninitialized MirrorState = "Uninitialized"
)

// PossibleMirrorStateValues returns an array of possible values for the MirrorState const type.
func PossibleMirrorStateValues() []MirrorState {
	return []MirrorState{MirrorStateBroken, MirrorStateMirrored, MirrorStateUninitialized}
}

// NetworkFeatures enumerates the values for network features.
type NetworkFeatures string

const (
	// NetworkFeaturesBasic Basic network feature.
	NetworkFeaturesBasic NetworkFeatures = "Basic"
	// NetworkFeaturesStandard Standard network feature.
	NetworkFeaturesStandard NetworkFeatures = "Standard"
)

// PossibleNetworkFeaturesValues returns an array of possible values for the NetworkFeatures const type.
func PossibleNetworkFeaturesValues() []NetworkFeatures {
	return []NetworkFeatures{NetworkFeaturesBasic, NetworkFeaturesStandard}
}

// QosType enumerates the values for qos type.
type QosType string

const (
	// QosTypeAuto qos type Auto
	QosTypeAuto QosType = "Auto"
	// QosTypeManual qos type Manual
	QosTypeManual QosType = "Manual"
)

// PossibleQosTypeValues returns an array of possible values for the QosType const type.
func PossibleQosTypeValues() []QosType {
	return []QosType{QosTypeAuto, QosTypeManual}
}

// RelationshipStatus enumerates the values for relationship status.
type RelationshipStatus string

const (
	// RelationshipStatusIdle ...
	RelationshipStatusIdle RelationshipStatus = "Idle"
	// RelationshipStatusTransferring ...
	RelationshipStatusTransferring RelationshipStatus = "Transferring"
)

// PossibleRelationshipStatusValues returns an array of possible values for the RelationshipStatus const type.
func PossibleRelationshipStatusValues() []RelationshipStatus {
	return []RelationshipStatus{RelationshipStatusIdle, RelationshipStatusTransferring}
}

// ReplicationSchedule enumerates the values for replication schedule.
type ReplicationSchedule string

const (
	// ReplicationSchedule10minutely ...
	ReplicationSchedule10minutely ReplicationSchedule = "_10minutely"
	// ReplicationScheduleDaily ...
	ReplicationScheduleDaily ReplicationSchedule = "daily"
	// ReplicationScheduleHourly ...
	ReplicationScheduleHourly ReplicationSchedule = "hourly"
)

// PossibleReplicationScheduleValues returns an array of possible values for the ReplicationSchedule const type.
func PossibleReplicationScheduleValues() []ReplicationSchedule {
	return []ReplicationSchedule{ReplicationSchedule10minutely, ReplicationScheduleDaily, ReplicationScheduleHourly}
}

// SecurityStyle enumerates the values for security style.
type SecurityStyle string

const (
	// SecurityStyleNtfs ...
	SecurityStyleNtfs SecurityStyle = "ntfs"
	// SecurityStyleUnix ...
	SecurityStyleUnix SecurityStyle = "unix"
)

// PossibleSecurityStyleValues returns an array of possible values for the SecurityStyle const type.
func PossibleSecurityStyleValues() []SecurityStyle {
	return []SecurityStyle{SecurityStyleNtfs, SecurityStyleUnix}
}

// ServiceLevel enumerates the values for service level.
type ServiceLevel string

const (
	// ServiceLevelPremium Premium service level
	ServiceLevelPremium ServiceLevel = "Premium"
	// ServiceLevelStandard Standard service level
	ServiceLevelStandard ServiceLevel = "Standard"
	// ServiceLevelStandardZRS Zone redundant storage service level
	ServiceLevelStandardZRS ServiceLevel = "StandardZRS"
	// ServiceLevelUltra Ultra service level
	ServiceLevelUltra ServiceLevel = "Ultra"
)

// PossibleServiceLevelValues returns an array of possible values for the ServiceLevel const type.
func PossibleServiceLevelValues() []ServiceLevel {
	return []ServiceLevel{ServiceLevelPremium, ServiceLevelStandard, ServiceLevelStandardZRS, ServiceLevelUltra}
}

// VolumeStorageToNetworkProximity enumerates the values for volume storage to network proximity.
type VolumeStorageToNetworkProximity string

const (
	// VolumeStorageToNetworkProximityDefault Basic storage to network connectivity.
	VolumeStorageToNetworkProximityDefault VolumeStorageToNetworkProximity = "Default"
	// VolumeStorageToNetworkProximityT1 Standard T1 storage to network connectivity.
	VolumeStorageToNetworkProximityT1 VolumeStorageToNetworkProximity = "T1"
	// VolumeStorageToNetworkProximityT2 Standard T2 storage to network connectivity.
	VolumeStorageToNetworkProximityT2 VolumeStorageToNetworkProximity = "T2"
)

// PossibleVolumeStorageToNetworkProximityValues returns an array of possible values for the VolumeStorageToNetworkProximity const type.
func PossibleVolumeStorageToNetworkProximityValues() []VolumeStorageToNetworkProximity {
	return []VolumeStorageToNetworkProximity{VolumeStorageToNetworkProximityDefault, VolumeStorageToNetworkProximityT1, VolumeStorageToNetworkProximityT2}
}
