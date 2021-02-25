// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package netapp

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/netapp/mgmt/2020-11-01/netapp"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActiveDirectoryStatus = original.ActiveDirectoryStatus

const (
	Created  ActiveDirectoryStatus = original.Created
	Deleted  ActiveDirectoryStatus = original.Deleted
	Error    ActiveDirectoryStatus = original.Error
	InUse    ActiveDirectoryStatus = original.InUse
	Updating ActiveDirectoryStatus = original.Updating
)

type CheckNameResourceTypes = original.CheckNameResourceTypes

const (
	MicrosoftNetAppnetAppAccounts                              CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccounts
	MicrosoftNetAppnetAppAccountscapacityPools                 CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccountscapacityPools
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumes          CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccountscapacityPoolsvolumes
	MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckNameResourceTypes = original.MicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots
)

type CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypes

const (
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts                              CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccounts
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools                 CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPools
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes          CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumes
	CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots CheckQuotaNameResourceTypes = original.CheckQuotaNameResourceTypesMicrosoftNetAppnetAppAccountscapacityPoolsvolumessnapshots
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type EndpointType = original.EndpointType

const (
	Dst EndpointType = original.Dst
	Src EndpointType = original.Src
)

type InAvailabilityReasonType = original.InAvailabilityReasonType

const (
	AlreadyExists InAvailabilityReasonType = original.AlreadyExists
	Invalid       InAvailabilityReasonType = original.Invalid
)

type KeySource = original.KeySource

const (
	MicrosoftNetApp KeySource = original.MicrosoftNetApp
)

type MirrorState = original.MirrorState

const (
	Broken        MirrorState = original.Broken
	Mirrored      MirrorState = original.Mirrored
	Uninitialized MirrorState = original.Uninitialized
)

type QosType = original.QosType

const (
	Auto   QosType = original.Auto
	Manual QosType = original.Manual
)

type RelationshipStatus = original.RelationshipStatus

const (
	Idle         RelationshipStatus = original.Idle
	Transferring RelationshipStatus = original.Transferring
)

type ReplicationSchedule = original.ReplicationSchedule

const (
	OneZerominutely ReplicationSchedule = original.OneZerominutely
	Daily           ReplicationSchedule = original.Daily
	Hourly          ReplicationSchedule = original.Hourly
)

type SecurityStyle = original.SecurityStyle

const (
	Ntfs SecurityStyle = original.Ntfs
	Unix SecurityStyle = original.Unix
)

type ServiceLevel = original.ServiceLevel

const (
	Premium  ServiceLevel = original.Premium
	Standard ServiceLevel = original.Standard
	Ultra    ServiceLevel = original.Ultra
)

type Account = original.Account
type AccountBackupsClient = original.AccountBackupsClient
type AccountBackupsDeleteFuture = original.AccountBackupsDeleteFuture
type AccountEncryption = original.AccountEncryption
type AccountList = original.AccountList
type AccountListIterator = original.AccountListIterator
type AccountListPage = original.AccountListPage
type AccountPatch = original.AccountPatch
type AccountProperties = original.AccountProperties
type AccountsClient = original.AccountsClient
type AccountsCreateOrUpdateFuture = original.AccountsCreateOrUpdateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type AccountsUpdateFuture = original.AccountsUpdateFuture
type ActiveDirectory = original.ActiveDirectory
type AuthorizeRequest = original.AuthorizeRequest
type Backup = original.Backup
type BackupPatch = original.BackupPatch
type BackupPoliciesClient = original.BackupPoliciesClient
type BackupPoliciesCreateFuture = original.BackupPoliciesCreateFuture
type BackupPoliciesDeleteFuture = original.BackupPoliciesDeleteFuture
type BackupPoliciesList = original.BackupPoliciesList
type BackupPolicy = original.BackupPolicy
type BackupPolicyDetails = original.BackupPolicyDetails
type BackupPolicyPatch = original.BackupPolicyPatch
type BackupPolicyProperties = original.BackupPolicyProperties
type BackupProperties = original.BackupProperties
type BackupsClient = original.BackupsClient
type BackupsCreateFuture = original.BackupsCreateFuture
type BackupsDeleteFuture = original.BackupsDeleteFuture
type BackupsList = original.BackupsList
type BackupsUpdateFuture = original.BackupsUpdateFuture
type BaseClient = original.BaseClient
type BreakReplicationRequest = original.BreakReplicationRequest
type CapacityPool = original.CapacityPool
type CapacityPoolList = original.CapacityPoolList
type CapacityPoolListIterator = original.CapacityPoolListIterator
type CapacityPoolListPage = original.CapacityPoolListPage
type CapacityPoolPatch = original.CapacityPoolPatch
type CheckAvailabilityResponse = original.CheckAvailabilityResponse
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type DailySchedule = original.DailySchedule
type Dimension = original.Dimension
type ExportPolicyRule = original.ExportPolicyRule
type HourlySchedule = original.HourlySchedule
type MetricSpecification = original.MetricSpecification
type MonthlySchedule = original.MonthlySchedule
type MountTarget = original.MountTarget
type MountTargetProperties = original.MountTargetProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PoolChangeRequest = original.PoolChangeRequest
type PoolPatchProperties = original.PoolPatchProperties
type PoolProperties = original.PoolProperties
type PoolsClient = original.PoolsClient
type PoolsCreateOrUpdateFuture = original.PoolsCreateOrUpdateFuture
type PoolsDeleteFuture = original.PoolsDeleteFuture
type PoolsUpdateFuture = original.PoolsUpdateFuture
type QuotaAvailabilityRequest = original.QuotaAvailabilityRequest
type ReplicationObject = original.ReplicationObject
type ReplicationStatus = original.ReplicationStatus
type ResourceClient = original.ResourceClient
type ResourceNameAvailabilityRequest = original.ResourceNameAvailabilityRequest
type ServiceSpecification = original.ServiceSpecification
type Snapshot = original.Snapshot
type SnapshotPoliciesClient = original.SnapshotPoliciesClient
type SnapshotPoliciesDeleteFuture = original.SnapshotPoliciesDeleteFuture
type SnapshotPoliciesList = original.SnapshotPoliciesList
type SnapshotPoliciesUpdateFuture = original.SnapshotPoliciesUpdateFuture
type SnapshotPolicy = original.SnapshotPolicy
type SnapshotPolicyDetails = original.SnapshotPolicyDetails
type SnapshotPolicyPatch = original.SnapshotPolicyPatch
type SnapshotPolicyProperties = original.SnapshotPolicyProperties
type SnapshotPolicyVolumeList = original.SnapshotPolicyVolumeList
type SnapshotProperties = original.SnapshotProperties
type SnapshotsClient = original.SnapshotsClient
type SnapshotsCreateFuture = original.SnapshotsCreateFuture
type SnapshotsDeleteFuture = original.SnapshotsDeleteFuture
type SnapshotsList = original.SnapshotsList
type SnapshotsUpdateFuture = original.SnapshotsUpdateFuture
type SystemData = original.SystemData
type Vault = original.Vault
type VaultList = original.VaultList
type VaultProperties = original.VaultProperties
type VaultsClient = original.VaultsClient
type Volume = original.Volume
type VolumeBackupProperties = original.VolumeBackupProperties
type VolumeBackups = original.VolumeBackups
type VolumeList = original.VolumeList
type VolumeListIterator = original.VolumeListIterator
type VolumeListPage = original.VolumeListPage
type VolumePatch = original.VolumePatch
type VolumePatchProperties = original.VolumePatchProperties
type VolumePatchPropertiesDataProtection = original.VolumePatchPropertiesDataProtection
type VolumePatchPropertiesExportPolicy = original.VolumePatchPropertiesExportPolicy
type VolumeProperties = original.VolumeProperties
type VolumePropertiesDataProtection = original.VolumePropertiesDataProtection
type VolumePropertiesExportPolicy = original.VolumePropertiesExportPolicy
type VolumeRevert = original.VolumeRevert
type VolumeSnapshotProperties = original.VolumeSnapshotProperties
type VolumesAuthorizeReplicationFuture = original.VolumesAuthorizeReplicationFuture
type VolumesBreakReplicationFuture = original.VolumesBreakReplicationFuture
type VolumesClient = original.VolumesClient
type VolumesCreateOrUpdateFuture = original.VolumesCreateOrUpdateFuture
type VolumesDeleteFuture = original.VolumesDeleteFuture
type VolumesDeleteReplicationFuture = original.VolumesDeleteReplicationFuture
type VolumesPoolChangeFuture = original.VolumesPoolChangeFuture
type VolumesReInitializeReplicationFuture = original.VolumesReInitializeReplicationFuture
type VolumesResyncReplicationFuture = original.VolumesResyncReplicationFuture
type VolumesRevertFuture = original.VolumesRevertFuture
type VolumesUpdateFuture = original.VolumesUpdateFuture
type WeeklySchedule = original.WeeklySchedule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountBackupsClient(subscriptionID string) AccountBackupsClient {
	return original.NewAccountBackupsClient(subscriptionID)
}
func NewAccountBackupsClientWithBaseURI(baseURI string, subscriptionID string) AccountBackupsClient {
	return original.NewAccountBackupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAccountListIterator(page AccountListPage) AccountListIterator {
	return original.NewAccountListIterator(page)
}
func NewAccountListPage(cur AccountList, getNextPage func(context.Context, AccountList) (AccountList, error)) AccountListPage {
	return original.NewAccountListPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBackupPoliciesClient(subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClient(subscriptionID)
}
func NewBackupPoliciesClientWithBaseURI(baseURI string, subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewBackupsClient(subscriptionID string) BackupsClient {
	return original.NewBackupsClient(subscriptionID)
}
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return original.NewBackupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCapacityPoolListIterator(page CapacityPoolListPage) CapacityPoolListIterator {
	return original.NewCapacityPoolListIterator(page)
}
func NewCapacityPoolListPage(cur CapacityPoolList, getNextPage func(context.Context, CapacityPoolList) (CapacityPoolList, error)) CapacityPoolListPage {
	return original.NewCapacityPoolListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPoolsClient(subscriptionID string) PoolsClient {
	return original.NewPoolsClient(subscriptionID)
}
func NewPoolsClientWithBaseURI(baseURI string, subscriptionID string) PoolsClient {
	return original.NewPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewResourceClient(subscriptionID string) ResourceClient {
	return original.NewResourceClient(subscriptionID)
}
func NewResourceClientWithBaseURI(baseURI string, subscriptionID string) ResourceClient {
	return original.NewResourceClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotPoliciesClient(subscriptionID string) SnapshotPoliciesClient {
	return original.NewSnapshotPoliciesClient(subscriptionID)
}
func NewSnapshotPoliciesClientWithBaseURI(baseURI string, subscriptionID string) SnapshotPoliciesClient {
	return original.NewSnapshotPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVaultsClient(subscriptionID string) VaultsClient {
	return original.NewVaultsClient(subscriptionID)
}
func NewVaultsClientWithBaseURI(baseURI string, subscriptionID string) VaultsClient {
	return original.NewVaultsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVolumeListIterator(page VolumeListPage) VolumeListIterator {
	return original.NewVolumeListIterator(page)
}
func NewVolumeListPage(cur VolumeList, getNextPage func(context.Context, VolumeList) (VolumeList, error)) VolumeListPage {
	return original.NewVolumeListPage(cur, getNextPage)
}
func NewVolumesClient(subscriptionID string) VolumesClient {
	return original.NewVolumesClient(subscriptionID)
}
func NewVolumesClientWithBaseURI(baseURI string, subscriptionID string) VolumesClient {
	return original.NewVolumesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActiveDirectoryStatusValues() []ActiveDirectoryStatus {
	return original.PossibleActiveDirectoryStatusValues()
}
func PossibleCheckNameResourceTypesValues() []CheckNameResourceTypes {
	return original.PossibleCheckNameResourceTypesValues()
}
func PossibleCheckQuotaNameResourceTypesValues() []CheckQuotaNameResourceTypes {
	return original.PossibleCheckQuotaNameResourceTypesValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleEndpointTypeValues() []EndpointType {
	return original.PossibleEndpointTypeValues()
}
func PossibleInAvailabilityReasonTypeValues() []InAvailabilityReasonType {
	return original.PossibleInAvailabilityReasonTypeValues()
}
func PossibleKeySourceValues() []KeySource {
	return original.PossibleKeySourceValues()
}
func PossibleMirrorStateValues() []MirrorState {
	return original.PossibleMirrorStateValues()
}
func PossibleQosTypeValues() []QosType {
	return original.PossibleQosTypeValues()
}
func PossibleRelationshipStatusValues() []RelationshipStatus {
	return original.PossibleRelationshipStatusValues()
}
func PossibleReplicationScheduleValues() []ReplicationSchedule {
	return original.PossibleReplicationScheduleValues()
}
func PossibleSecurityStyleValues() []SecurityStyle {
	return original.PossibleSecurityStyleValues()
}
func PossibleServiceLevelValues() []ServiceLevel {
	return original.PossibleServiceLevelValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
