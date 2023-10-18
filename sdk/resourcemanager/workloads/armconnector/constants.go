//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconnector

const (
	moduleName    = "armconnector"
	moduleVersion = "v0.1.0"
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

// BackupType - The type of backup, VM, SQL or HANA.
type BackupType string

const (
	BackupTypeHANA BackupType = "HANA"
	BackupTypeSQL  BackupType = "SQL"
	BackupTypeVM   BackupType = "VM"
)

// PossibleBackupTypeValues returns the possible values for the BackupType const type.
func PossibleBackupTypeValues() []BackupType {
	return []BackupType{
		BackupTypeHANA,
		BackupTypeSQL,
		BackupTypeVM,
	}
}

// ConnectorProvisioningState - Defines the provisioning states.
type ConnectorProvisioningState string

const (
	ConnectorProvisioningStateCanceled  ConnectorProvisioningState = "Canceled"
	ConnectorProvisioningStateCreating  ConnectorProvisioningState = "Creating"
	ConnectorProvisioningStateDeleting  ConnectorProvisioningState = "Deleting"
	ConnectorProvisioningStateFailed    ConnectorProvisioningState = "Failed"
	ConnectorProvisioningStateSucceeded ConnectorProvisioningState = "Succeeded"
	ConnectorProvisioningStateUpdating  ConnectorProvisioningState = "Updating"
)

// PossibleConnectorProvisioningStateValues returns the possible values for the ConnectorProvisioningState const type.
func PossibleConnectorProvisioningStateValues() []ConnectorProvisioningState {
	return []ConnectorProvisioningState{
		ConnectorProvisioningStateCanceled,
		ConnectorProvisioningStateCreating,
		ConnectorProvisioningStateDeleting,
		ConnectorProvisioningStateFailed,
		ConnectorProvisioningStateSucceeded,
		ConnectorProvisioningStateUpdating,
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

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
	}
}

// IAASVMPolicyType - The policy type.
type IAASVMPolicyType string

const (
	IAASVMPolicyTypeInvalid IAASVMPolicyType = "Invalid"
	IAASVMPolicyTypeV1      IAASVMPolicyType = "V1"
	IAASVMPolicyTypeV2      IAASVMPolicyType = "V2"
)

// PossibleIAASVMPolicyTypeValues returns the possible values for the IAASVMPolicyType const type.
func PossibleIAASVMPolicyTypeValues() []IAASVMPolicyType {
	return []IAASVMPolicyType{
		IAASVMPolicyTypeInvalid,
		IAASVMPolicyTypeV1,
		IAASVMPolicyTypeV2,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (only None, UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone         ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

type MonthOfYear string

const (
	MonthOfYearApril     MonthOfYear = "April"
	MonthOfYearAugust    MonthOfYear = "August"
	MonthOfYearDecember  MonthOfYear = "December"
	MonthOfYearFebruary  MonthOfYear = "February"
	MonthOfYearInvalid   MonthOfYear = "Invalid"
	MonthOfYearJanuary   MonthOfYear = "January"
	MonthOfYearJuly      MonthOfYear = "July"
	MonthOfYearJune      MonthOfYear = "June"
	MonthOfYearMarch     MonthOfYear = "March"
	MonthOfYearMay       MonthOfYear = "May"
	MonthOfYearNovember  MonthOfYear = "November"
	MonthOfYearOctober   MonthOfYear = "October"
	MonthOfYearSeptember MonthOfYear = "September"
)

// PossibleMonthOfYearValues returns the possible values for the MonthOfYear const type.
func PossibleMonthOfYearValues() []MonthOfYear {
	return []MonthOfYear{
		MonthOfYearApril,
		MonthOfYearAugust,
		MonthOfYearDecember,
		MonthOfYearFebruary,
		MonthOfYearInvalid,
		MonthOfYearJanuary,
		MonthOfYearJuly,
		MonthOfYearJune,
		MonthOfYearMarch,
		MonthOfYearMay,
		MonthOfYearNovember,
		MonthOfYearOctober,
		MonthOfYearSeptember,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
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

// PolicyType - Type of backup policy type
type PolicyType string

const (
	PolicyTypeCopyOnlyFull         PolicyType = "CopyOnlyFull"
	PolicyTypeDifferential         PolicyType = "Differential"
	PolicyTypeFull                 PolicyType = "Full"
	PolicyTypeIncremental          PolicyType = "Incremental"
	PolicyTypeInvalid              PolicyType = "Invalid"
	PolicyTypeLog                  PolicyType = "Log"
	PolicyTypeSnapshotCopyOnlyFull PolicyType = "SnapshotCopyOnlyFull"
	PolicyTypeSnapshotFull         PolicyType = "SnapshotFull"
)

// PossiblePolicyTypeValues returns the possible values for the PolicyType const type.
func PossiblePolicyTypeValues() []PolicyType {
	return []PolicyType{
		PolicyTypeCopyOnlyFull,
		PolicyTypeDifferential,
		PolicyTypeFull,
		PolicyTypeIncremental,
		PolicyTypeInvalid,
		PolicyTypeLog,
		PolicyTypeSnapshotCopyOnlyFull,
		PolicyTypeSnapshotFull,
	}
}

// RetentionDurationType - Retention duration type: days/weeks/months/years Used only if TieringMode is set to TierAfter
type RetentionDurationType string

const (
	RetentionDurationTypeDays    RetentionDurationType = "Days"
	RetentionDurationTypeInvalid RetentionDurationType = "Invalid"
	RetentionDurationTypeMonths  RetentionDurationType = "Months"
	RetentionDurationTypeWeeks   RetentionDurationType = "Weeks"
	RetentionDurationTypeYears   RetentionDurationType = "Years"
)

// PossibleRetentionDurationTypeValues returns the possible values for the RetentionDurationType const type.
func PossibleRetentionDurationTypeValues() []RetentionDurationType {
	return []RetentionDurationType{
		RetentionDurationTypeDays,
		RetentionDurationTypeInvalid,
		RetentionDurationTypeMonths,
		RetentionDurationTypeWeeks,
		RetentionDurationTypeYears,
	}
}

// RetentionScheduleFormat - Retention schedule format type for monthly retention policy.
type RetentionScheduleFormat string

const (
	RetentionScheduleFormatDaily   RetentionScheduleFormat = "Daily"
	RetentionScheduleFormatInvalid RetentionScheduleFormat = "Invalid"
	RetentionScheduleFormatWeekly  RetentionScheduleFormat = "Weekly"
)

// PossibleRetentionScheduleFormatValues returns the possible values for the RetentionScheduleFormat const type.
func PossibleRetentionScheduleFormatValues() []RetentionScheduleFormat {
	return []RetentionScheduleFormat{
		RetentionScheduleFormatDaily,
		RetentionScheduleFormatInvalid,
		RetentionScheduleFormatWeekly,
	}
}

// SSLCryptoProvider - Specify the crypto provider being used (commoncrypto/openssl). If this argument is not provided, it
// is automatically determined by searching in the configuration files.
type SSLCryptoProvider string

const (
	SSLCryptoProviderCommoncrypto SSLCryptoProvider = "commoncrypto"
	SSLCryptoProviderOpenssl      SSLCryptoProvider = "openssl"
)

// PossibleSSLCryptoProviderValues returns the possible values for the SSLCryptoProvider const type.
func PossibleSSLCryptoProviderValues() []SSLCryptoProvider {
	return []SSLCryptoProvider{
		SSLCryptoProviderCommoncrypto,
		SSLCryptoProviderOpenssl,
	}
}

// ScheduleRunType - Frequency of the schedule operation of this policy.
type ScheduleRunType string

const (
	ScheduleRunTypeDaily   ScheduleRunType = "Daily"
	ScheduleRunTypeHourly  ScheduleRunType = "Hourly"
	ScheduleRunTypeInvalid ScheduleRunType = "Invalid"
	ScheduleRunTypeWeekly  ScheduleRunType = "Weekly"
)

// PossibleScheduleRunTypeValues returns the possible values for the ScheduleRunType const type.
func PossibleScheduleRunTypeValues() []ScheduleRunType {
	return []ScheduleRunType{
		ScheduleRunTypeDaily,
		ScheduleRunTypeHourly,
		ScheduleRunTypeInvalid,
		ScheduleRunTypeWeekly,
	}
}

// TieringMode - Tiering Mode to control automatic tiering of recovery points. Supported values are:
// 1. TierRecommended: Tier all recovery points recommended to be tiered
// 2. TierAfter: Tier all recovery points after a fixed period, as specified in duration + durationType below.
// 3. DoNotTier: Do not tier any recovery points
type TieringMode string

const (
	TieringModeDoNotTier       TieringMode = "DoNotTier"
	TieringModeInvalid         TieringMode = "Invalid"
	TieringModeTierAfter       TieringMode = "TierAfter"
	TieringModeTierRecommended TieringMode = "TierRecommended"
)

// PossibleTieringModeValues returns the possible values for the TieringMode const type.
func PossibleTieringModeValues() []TieringMode {
	return []TieringMode{
		TieringModeDoNotTier,
		TieringModeInvalid,
		TieringModeTierAfter,
		TieringModeTierRecommended,
	}
}

// VaultType - The vault type, whether it is existing or has to be created.
type VaultType string

const (
	VaultTypeExisting VaultType = "Existing"
	VaultTypeNew      VaultType = "New"
)

// PossibleVaultTypeValues returns the possible values for the VaultType const type.
func PossibleVaultTypeValues() []VaultType {
	return []VaultType{
		VaultTypeExisting,
		VaultTypeNew,
	}
}

type WeekOfMonth string

const (
	WeekOfMonthFirst   WeekOfMonth = "First"
	WeekOfMonthFourth  WeekOfMonth = "Fourth"
	WeekOfMonthInvalid WeekOfMonth = "Invalid"
	WeekOfMonthLast    WeekOfMonth = "Last"
	WeekOfMonthSecond  WeekOfMonth = "Second"
	WeekOfMonthThird   WeekOfMonth = "Third"
)

// PossibleWeekOfMonthValues returns the possible values for the WeekOfMonth const type.
func PossibleWeekOfMonthValues() []WeekOfMonth {
	return []WeekOfMonth{
		WeekOfMonthFirst,
		WeekOfMonthFourth,
		WeekOfMonthInvalid,
		WeekOfMonthLast,
		WeekOfMonthSecond,
		WeekOfMonthThird,
	}
}

// WorkloadType - Type of workload for the backup management
type WorkloadType string

const (
	WorkloadTypeAzureFileShare    WorkloadType = "AzureFileShare"
	WorkloadTypeAzureSQLDb        WorkloadType = "AzureSqlDb"
	WorkloadTypeClient            WorkloadType = "Client"
	WorkloadTypeExchange          WorkloadType = "Exchange"
	WorkloadTypeFileFolder        WorkloadType = "FileFolder"
	WorkloadTypeGenericDataSource WorkloadType = "GenericDataSource"
	WorkloadTypeInvalid           WorkloadType = "Invalid"
	WorkloadTypeSAPAseDatabase    WorkloadType = "SAPAseDatabase"
	WorkloadTypeSAPHanaDBInstance WorkloadType = "SAPHanaDBInstance"
	WorkloadTypeSAPHanaDatabase   WorkloadType = "SAPHanaDatabase"
	WorkloadTypeSQLDB             WorkloadType = "SQLDB"
	WorkloadTypeSQLDataBase       WorkloadType = "SQLDataBase"
	WorkloadTypeSharepoint        WorkloadType = "Sharepoint"
	WorkloadTypeSystemState       WorkloadType = "SystemState"
	WorkloadTypeVM                WorkloadType = "VM"
	WorkloadTypeVMwareVM          WorkloadType = "VMwareVM"
)

// PossibleWorkloadTypeValues returns the possible values for the WorkloadType const type.
func PossibleWorkloadTypeValues() []WorkloadType {
	return []WorkloadType{
		WorkloadTypeAzureFileShare,
		WorkloadTypeAzureSQLDb,
		WorkloadTypeClient,
		WorkloadTypeExchange,
		WorkloadTypeFileFolder,
		WorkloadTypeGenericDataSource,
		WorkloadTypeInvalid,
		WorkloadTypeSAPAseDatabase,
		WorkloadTypeSAPHanaDBInstance,
		WorkloadTypeSAPHanaDatabase,
		WorkloadTypeSQLDB,
		WorkloadTypeSQLDataBase,
		WorkloadTypeSharepoint,
		WorkloadTypeSystemState,
		WorkloadTypeVM,
		WorkloadTypeVMwareVM,
	}
}
