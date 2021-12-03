//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

const (
	module  = "armmysqlflexibleservers"
	version = "v0.2.1"
)

// ConfigurationSource - Source of the configuration.
type ConfigurationSource string

const (
	ConfigurationSourceSystemDefault ConfigurationSource = "system-default"
	ConfigurationSourceUserOverride  ConfigurationSource = "user-override"
)

// PossibleConfigurationSourceValues returns the possible values for the ConfigurationSource const type.
func PossibleConfigurationSourceValues() []ConfigurationSource {
	return []ConfigurationSource{
		ConfigurationSourceSystemDefault,
		ConfigurationSourceUserOverride,
	}
}

// ToPtr returns a *ConfigurationSource pointing to the current value.
func (c ConfigurationSource) ToPtr() *ConfigurationSource {
	return &c
}

// CreateMode - The mode to create a new MySQL server.
type CreateMode string

const (
	CreateModeDefault            CreateMode = "Default"
	CreateModeGeoRestore         CreateMode = "GeoRestore"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReplica            CreateMode = "Replica"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeGeoRestore,
		CreateModePointInTimeRestore,
		CreateModeReplica,
	}
}

// ToPtr returns a *CreateMode pointing to the current value.
func (c CreateMode) ToPtr() *CreateMode {
	return &c
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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// EnableStatusEnum - Enum to indicate whether value is 'Enabled' or 'Disabled'
type EnableStatusEnum string

const (
	EnableStatusEnumDisabled EnableStatusEnum = "Disabled"
	EnableStatusEnumEnabled  EnableStatusEnum = "Enabled"
)

// PossibleEnableStatusEnumValues returns the possible values for the EnableStatusEnum const type.
func PossibleEnableStatusEnumValues() []EnableStatusEnum {
	return []EnableStatusEnum{
		EnableStatusEnumDisabled,
		EnableStatusEnumEnabled,
	}
}

// ToPtr returns a *EnableStatusEnum pointing to the current value.
func (c EnableStatusEnum) ToPtr() *EnableStatusEnum {
	return &c
}

// HighAvailabilityMode - High availability mode for a server.
type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      HighAvailabilityMode = "Disabled"
	HighAvailabilityModeSameZone      HighAvailabilityMode = "SameZone"
	HighAvailabilityModeZoneRedundant HighAvailabilityMode = "ZoneRedundant"
)

// PossibleHighAvailabilityModeValues returns the possible values for the HighAvailabilityMode const type.
func PossibleHighAvailabilityModeValues() []HighAvailabilityMode {
	return []HighAvailabilityMode{
		HighAvailabilityModeDisabled,
		HighAvailabilityModeSameZone,
		HighAvailabilityModeZoneRedundant,
	}
}

// ToPtr returns a *HighAvailabilityMode pointing to the current value.
func (c HighAvailabilityMode) ToPtr() *HighAvailabilityMode {
	return &c
}

// HighAvailabilityState - The state of server high availability.
type HighAvailabilityState string

const (
	HighAvailabilityStateCreatingStandby HighAvailabilityState = "CreatingStandby"
	HighAvailabilityStateFailingOver     HighAvailabilityState = "FailingOver"
	HighAvailabilityStateHealthy         HighAvailabilityState = "Healthy"
	HighAvailabilityStateNotEnabled      HighAvailabilityState = "NotEnabled"
	HighAvailabilityStateRemovingStandby HighAvailabilityState = "RemovingStandby"
)

// PossibleHighAvailabilityStateValues returns the possible values for the HighAvailabilityState const type.
func PossibleHighAvailabilityStateValues() []HighAvailabilityState {
	return []HighAvailabilityState{
		HighAvailabilityStateCreatingStandby,
		HighAvailabilityStateFailingOver,
		HighAvailabilityStateHealthy,
		HighAvailabilityStateNotEnabled,
		HighAvailabilityStateRemovingStandby,
	}
}

// ToPtr returns a *HighAvailabilityState pointing to the current value.
func (c HighAvailabilityState) ToPtr() *HighAvailabilityState {
	return &c
}

// IsConfigPendingRestart - If is the configuration pending restart or not.
type IsConfigPendingRestart string

const (
	IsConfigPendingRestartFalse IsConfigPendingRestart = "False"
	IsConfigPendingRestartTrue  IsConfigPendingRestart = "True"
)

// PossibleIsConfigPendingRestartValues returns the possible values for the IsConfigPendingRestart const type.
func PossibleIsConfigPendingRestartValues() []IsConfigPendingRestart {
	return []IsConfigPendingRestart{
		IsConfigPendingRestartFalse,
		IsConfigPendingRestartTrue,
	}
}

// ToPtr returns a *IsConfigPendingRestart pointing to the current value.
func (c IsConfigPendingRestart) ToPtr() *IsConfigPendingRestart {
	return &c
}

// IsDynamicConfig - If is the configuration dynamic.
type IsDynamicConfig string

const (
	IsDynamicConfigFalse IsDynamicConfig = "False"
	IsDynamicConfigTrue  IsDynamicConfig = "True"
)

// PossibleIsDynamicConfigValues returns the possible values for the IsDynamicConfig const type.
func PossibleIsDynamicConfigValues() []IsDynamicConfig {
	return []IsDynamicConfig{
		IsDynamicConfigFalse,
		IsDynamicConfigTrue,
	}
}

// ToPtr returns a *IsDynamicConfig pointing to the current value.
func (c IsDynamicConfig) ToPtr() *IsDynamicConfig {
	return &c
}

// IsReadOnly - If is the configuration read only.
type IsReadOnly string

const (
	IsReadOnlyFalse IsReadOnly = "False"
	IsReadOnlyTrue  IsReadOnly = "True"
)

// PossibleIsReadOnlyValues returns the possible values for the IsReadOnly const type.
func PossibleIsReadOnlyValues() []IsReadOnly {
	return []IsReadOnly{
		IsReadOnlyFalse,
		IsReadOnlyTrue,
	}
}

// ToPtr returns a *IsReadOnly pointing to the current value.
func (c IsReadOnly) ToPtr() *IsReadOnly {
	return &c
}

// ReplicationRole - The replication role.
type ReplicationRole string

const (
	ReplicationRoleNone    ReplicationRole = "None"
	ReplicationRoleReplica ReplicationRole = "Replica"
	ReplicationRoleSource  ReplicationRole = "Source"
)

// PossibleReplicationRoleValues returns the possible values for the ReplicationRole const type.
func PossibleReplicationRoleValues() []ReplicationRole {
	return []ReplicationRole{
		ReplicationRoleNone,
		ReplicationRoleReplica,
		ReplicationRoleSource,
	}
}

// ToPtr returns a *ReplicationRole pointing to the current value.
func (c ReplicationRole) ToPtr() *ReplicationRole {
	return &c
}

// SKUTier - The tier of the particular SKU, e.g. GeneralPurpose.
type SKUTier string

const (
	SKUTierBurstable       SKUTier = "Burstable"
	SKUTierGeneralPurpose  SKUTier = "GeneralPurpose"
	SKUTierMemoryOptimized SKUTier = "MemoryOptimized"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBurstable,
		SKUTierGeneralPurpose,
		SKUTierMemoryOptimized,
	}
}

// ToPtr returns a *SKUTier pointing to the current value.
func (c SKUTier) ToPtr() *SKUTier {
	return &c
}

// ServerState - The state of a server.
type ServerState string

const (
	ServerStateDisabled ServerState = "Disabled"
	ServerStateDropping ServerState = "Dropping"
	ServerStateReady    ServerState = "Ready"
	ServerStateStarting ServerState = "Starting"
	ServerStateStopped  ServerState = "Stopped"
	ServerStateStopping ServerState = "Stopping"
	ServerStateUpdating ServerState = "Updating"
)

// PossibleServerStateValues returns the possible values for the ServerState const type.
func PossibleServerStateValues() []ServerState {
	return []ServerState{
		ServerStateDisabled,
		ServerStateDropping,
		ServerStateReady,
		ServerStateStarting,
		ServerStateStopped,
		ServerStateStopping,
		ServerStateUpdating,
	}
}

// ToPtr returns a *ServerState pointing to the current value.
func (c ServerState) ToPtr() *ServerState {
	return &c
}

// ServerVersion - The version of a server.
type ServerVersion string

const (
	ServerVersionEight021 ServerVersion = "8.0.21"
	ServerVersionFive7    ServerVersion = "5.7"
)

// PossibleServerVersionValues returns the possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{
		ServerVersionEight021,
		ServerVersionFive7,
	}
}

// ToPtr returns a *ServerVersion pointing to the current value.
func (c ServerVersion) ToPtr() *ServerVersion {
	return &c
}
