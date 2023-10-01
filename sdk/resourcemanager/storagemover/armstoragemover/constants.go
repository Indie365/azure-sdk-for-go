//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragemover

const (
	moduleName    = "armstoragemover"
	moduleVersion = "v2.0.0"
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

// AgentStatus - The Agent status.
type AgentStatus string

const (
	AgentStatusExecuting         AgentStatus = "Executing"
	AgentStatusOffline           AgentStatus = "Offline"
	AgentStatusOnline            AgentStatus = "Online"
	AgentStatusRegistering       AgentStatus = "Registering"
	AgentStatusRequiresAttention AgentStatus = "RequiresAttention"
	AgentStatusUnregistering     AgentStatus = "Unregistering"
)

// PossibleAgentStatusValues returns the possible values for the AgentStatus const type.
func PossibleAgentStatusValues() []AgentStatus {
	return []AgentStatus{
		AgentStatusExecuting,
		AgentStatusOffline,
		AgentStatusOnline,
		AgentStatusRegistering,
		AgentStatusRequiresAttention,
		AgentStatusUnregistering,
	}
}

// CopyMode - Strategy to use for copy.
type CopyMode string

const (
	CopyModeAdditive CopyMode = "Additive"
	CopyModeMirror   CopyMode = "Mirror"
)

// PossibleCopyModeValues returns the possible values for the CopyMode const type.
func PossibleCopyModeValues() []CopyMode {
	return []CopyMode{
		CopyModeAdditive,
		CopyModeMirror,
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

// CredentialType - The Credentials type.
type CredentialType string

const (
	CredentialTypeAzureKeyVaultSmb CredentialType = "AzureKeyVaultSmb"
)

// PossibleCredentialTypeValues returns the possible values for the CredentialType const type.
func PossibleCredentialTypeValues() []CredentialType {
	return []CredentialType{
		CredentialTypeAzureKeyVaultSmb,
	}
}

// EndpointType - The Endpoint resource type.
type EndpointType string

const (
	EndpointTypeAzureStorageBlobContainer EndpointType = "AzureStorageBlobContainer"
	EndpointTypeAzureStorageSmbFileShare  EndpointType = "AzureStorageSmbFileShare"
	EndpointTypeNfsMount                  EndpointType = "NfsMount"
	EndpointTypeSmbMount                  EndpointType = "SmbMount"
)

// PossibleEndpointTypeValues returns the possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{
		EndpointTypeAzureStorageBlobContainer,
		EndpointTypeAzureStorageSmbFileShare,
		EndpointTypeNfsMount,
		EndpointTypeSmbMount,
	}
}

// JobRunScanStatus - The status of Agent's scanning of source.
type JobRunScanStatus string

const (
	JobRunScanStatusCompleted  JobRunScanStatus = "Completed"
	JobRunScanStatusNotStarted JobRunScanStatus = "NotStarted"
	JobRunScanStatusScanning   JobRunScanStatus = "Scanning"
)

// PossibleJobRunScanStatusValues returns the possible values for the JobRunScanStatus const type.
func PossibleJobRunScanStatusValues() []JobRunScanStatus {
	return []JobRunScanStatus{
		JobRunScanStatusCompleted,
		JobRunScanStatusNotStarted,
		JobRunScanStatusScanning,
	}
}

// JobRunStatus - The current status of the Job Run in a non-terminal state, if exists.
type JobRunStatus string

const (
	JobRunStatusCancelRequested JobRunStatus = "CancelRequested"
	JobRunStatusCanceled        JobRunStatus = "Canceled"
	JobRunStatusCanceling       JobRunStatus = "Canceling"
	JobRunStatusFailed          JobRunStatus = "Failed"
	JobRunStatusQueued          JobRunStatus = "Queued"
	JobRunStatusRunning         JobRunStatus = "Running"
	JobRunStatusStarted         JobRunStatus = "Started"
	JobRunStatusSucceeded       JobRunStatus = "Succeeded"
)

// PossibleJobRunStatusValues returns the possible values for the JobRunStatus const type.
func PossibleJobRunStatusValues() []JobRunStatus {
	return []JobRunStatus{
		JobRunStatusCancelRequested,
		JobRunStatusCanceled,
		JobRunStatusCanceling,
		JobRunStatusFailed,
		JobRunStatusQueued,
		JobRunStatusRunning,
		JobRunStatusStarted,
		JobRunStatusSucceeded,
	}
}

// NfsVersion - The NFS protocol version.
type NfsVersion string

const (
	NfsVersionNFSauto NfsVersion = "NFSauto"
	NfsVersionNFSv3   NfsVersion = "NFSv3"
	NfsVersionNFSv4   NfsVersion = "NFSv4"
)

// PossibleNfsVersionValues returns the possible values for the NfsVersion const type.
func PossibleNfsVersionValues() []NfsVersion {
	return []NfsVersion{
		NfsVersionNFSauto,
		NfsVersionNFSv3,
		NfsVersionNFSv4,
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

// ProvisioningState - The provisioning state of this resource.
type ProvisioningState string

const (
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateSucceeded,
	}
}
