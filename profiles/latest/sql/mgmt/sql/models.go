// +build go1.9

// Copyright 2017 Microsoft Corporation
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
// commit ID: 714052a3359963ba46703fe033cf9dd4838bea11

package sql

import original "github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"

type FirewallRulesClient = original.FirewallRulesClient
type AuthenticationType = original.AuthenticationType

const (
	ADPassword	AuthenticationType	= original.ADPassword
	SQL		AuthenticationType	= original.SQL
)

type CheckNameAvailabilityReason = original.CheckNameAvailabilityReason

const (
	AlreadyExists	CheckNameAvailabilityReason	= original.AlreadyExists
	Invalid		CheckNameAvailabilityReason	= original.Invalid
)

type CreateMode = original.CreateMode

const (
	Copy				CreateMode	= original.Copy
	Default				CreateMode	= original.Default
	NonReadableSecondary		CreateMode	= original.NonReadableSecondary
	OnlineSecondary			CreateMode	= original.OnlineSecondary
	PointInTimeRestore		CreateMode	= original.PointInTimeRestore
	Recovery			CreateMode	= original.Recovery
	Restore				CreateMode	= original.Restore
	RestoreLongTermRetentionBackup	CreateMode	= original.RestoreLongTermRetentionBackup
)

type DatabaseEdition = original.DatabaseEdition

const (
	Basic		DatabaseEdition	= original.Basic
	Business	DatabaseEdition	= original.Business
	DataWarehouse	DatabaseEdition	= original.DataWarehouse
	Free		DatabaseEdition	= original.Free
	Premium		DatabaseEdition	= original.Premium
	Standard	DatabaseEdition	= original.Standard
	Stretch		DatabaseEdition	= original.Stretch
	System		DatabaseEdition	= original.System
	System2		DatabaseEdition	= original.System2
	Web		DatabaseEdition	= original.Web
)

type ElasticPoolEdition = original.ElasticPoolEdition

const (
	ElasticPoolEditionBasic		ElasticPoolEdition	= original.ElasticPoolEditionBasic
	ElasticPoolEditionPremium	ElasticPoolEdition	= original.ElasticPoolEditionPremium
	ElasticPoolEditionStandard	ElasticPoolEdition	= original.ElasticPoolEditionStandard
)

type ElasticPoolState = original.ElasticPoolState

const (
	Creating	ElasticPoolState	= original.Creating
	Disabled	ElasticPoolState	= original.Disabled
	Ready		ElasticPoolState	= original.Ready
)

type ReadScale = original.ReadScale

const (
	ReadScaleDisabled	ReadScale	= original.ReadScaleDisabled
	ReadScaleEnabled	ReadScale	= original.ReadScaleEnabled
)

type RecommendedIndexAction = original.RecommendedIndexAction

const (
	Create	RecommendedIndexAction	= original.Create
	Drop	RecommendedIndexAction	= original.Drop
	Rebuild	RecommendedIndexAction	= original.Rebuild
)

type RecommendedIndexState = original.RecommendedIndexState

const (
	Active		RecommendedIndexState	= original.Active
	Blocked		RecommendedIndexState	= original.Blocked
	Executing	RecommendedIndexState	= original.Executing
	Expired		RecommendedIndexState	= original.Expired
	Ignored		RecommendedIndexState	= original.Ignored
	Pending		RecommendedIndexState	= original.Pending
	PendingRevert	RecommendedIndexState	= original.PendingRevert
	Reverted	RecommendedIndexState	= original.Reverted
	Reverting	RecommendedIndexState	= original.Reverting
	Success		RecommendedIndexState	= original.Success
	Verifying	RecommendedIndexState	= original.Verifying
)

type RecommendedIndexType = original.RecommendedIndexType

const (
	CLUSTERED		RecommendedIndexType	= original.CLUSTERED
	CLUSTEREDCOLUMNSTORE	RecommendedIndexType	= original.CLUSTEREDCOLUMNSTORE
	COLUMNSTORE		RecommendedIndexType	= original.COLUMNSTORE
	NONCLUSTERED		RecommendedIndexType	= original.NONCLUSTERED
)

type ReplicationRole = original.ReplicationRole

const (
	ReplicationRoleCopy			ReplicationRole	= original.ReplicationRoleCopy
	ReplicationRoleNonReadableSecondary	ReplicationRole	= original.ReplicationRoleNonReadableSecondary
	ReplicationRolePrimary			ReplicationRole	= original.ReplicationRolePrimary
	ReplicationRoleSecondary		ReplicationRole	= original.ReplicationRoleSecondary
	ReplicationRoleSource			ReplicationRole	= original.ReplicationRoleSource
)

type ReplicationState = original.ReplicationState

const (
	CATCHUP		ReplicationState	= original.CATCHUP
	PENDING		ReplicationState	= original.PENDING
	SEEDING		ReplicationState	= original.SEEDING
	SUSPENDED	ReplicationState	= original.SUSPENDED
)

type SampleName = original.SampleName

const (
	AdventureWorksLT SampleName = original.AdventureWorksLT
)

type SecurityAlertPolicyEmailAccountAdmins = original.SecurityAlertPolicyEmailAccountAdmins

const (
	SecurityAlertPolicyEmailAccountAdminsDisabled	SecurityAlertPolicyEmailAccountAdmins	= original.SecurityAlertPolicyEmailAccountAdminsDisabled
	SecurityAlertPolicyEmailAccountAdminsEnabled	SecurityAlertPolicyEmailAccountAdmins	= original.SecurityAlertPolicyEmailAccountAdminsEnabled
)

type SecurityAlertPolicyState = original.SecurityAlertPolicyState

const (
	SecurityAlertPolicyStateDisabled	SecurityAlertPolicyState	= original.SecurityAlertPolicyStateDisabled
	SecurityAlertPolicyStateEnabled		SecurityAlertPolicyState	= original.SecurityAlertPolicyStateEnabled
	SecurityAlertPolicyStateNew		SecurityAlertPolicyState	= original.SecurityAlertPolicyStateNew
)

type SecurityAlertPolicyUseServerDefault = original.SecurityAlertPolicyUseServerDefault

const (
	SecurityAlertPolicyUseServerDefaultDisabled	SecurityAlertPolicyUseServerDefault	= original.SecurityAlertPolicyUseServerDefaultDisabled
	SecurityAlertPolicyUseServerDefaultEnabled	SecurityAlertPolicyUseServerDefault	= original.SecurityAlertPolicyUseServerDefaultEnabled
)

type ServiceObjectiveName = original.ServiceObjectiveName

const (
	ServiceObjectiveNameBasic	ServiceObjectiveName	= original.ServiceObjectiveNameBasic
	ServiceObjectiveNameElasticPool	ServiceObjectiveName	= original.ServiceObjectiveNameElasticPool
	ServiceObjectiveNameP1		ServiceObjectiveName	= original.ServiceObjectiveNameP1
	ServiceObjectiveNameP11		ServiceObjectiveName	= original.ServiceObjectiveNameP11
	ServiceObjectiveNameP15		ServiceObjectiveName	= original.ServiceObjectiveNameP15
	ServiceObjectiveNameP2		ServiceObjectiveName	= original.ServiceObjectiveNameP2
	ServiceObjectiveNameP3		ServiceObjectiveName	= original.ServiceObjectiveNameP3
	ServiceObjectiveNameP4		ServiceObjectiveName	= original.ServiceObjectiveNameP4
	ServiceObjectiveNameP6		ServiceObjectiveName	= original.ServiceObjectiveNameP6
	ServiceObjectiveNameS0		ServiceObjectiveName	= original.ServiceObjectiveNameS0
	ServiceObjectiveNameS1		ServiceObjectiveName	= original.ServiceObjectiveNameS1
	ServiceObjectiveNameS2		ServiceObjectiveName	= original.ServiceObjectiveNameS2
	ServiceObjectiveNameS3		ServiceObjectiveName	= original.ServiceObjectiveNameS3
	ServiceObjectiveNameSystem	ServiceObjectiveName	= original.ServiceObjectiveNameSystem
	ServiceObjectiveNameSystem2	ServiceObjectiveName	= original.ServiceObjectiveNameSystem2
)

type StorageKeyType = original.StorageKeyType

const (
	SharedAccessKey		StorageKeyType	= original.SharedAccessKey
	StorageAccessKey	StorageKeyType	= original.StorageAccessKey
)

type TransparentDataEncryptionActivityStatus = original.TransparentDataEncryptionActivityStatus

const (
	Decrypting	TransparentDataEncryptionActivityStatus	= original.Decrypting
	Encrypting	TransparentDataEncryptionActivityStatus	= original.Encrypting
)

type TransparentDataEncryptionStatus = original.TransparentDataEncryptionStatus

const (
	TransparentDataEncryptionStatusDisabled	TransparentDataEncryptionStatus	= original.TransparentDataEncryptionStatusDisabled
	TransparentDataEncryptionStatusEnabled	TransparentDataEncryptionStatus	= original.TransparentDataEncryptionStatusEnabled
)

type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type Database = original.Database
type DatabaseListResult = original.DatabaseListResult
type DatabaseProperties = original.DatabaseProperties
type DatabaseSecurityAlertPolicy = original.DatabaseSecurityAlertPolicy
type DatabaseSecurityAlertPolicyProperties = original.DatabaseSecurityAlertPolicyProperties
type DatabaseUpdate = original.DatabaseUpdate
type ElasticPool = original.ElasticPool
type ElasticPoolActivity = original.ElasticPoolActivity
type ElasticPoolActivityListResult = original.ElasticPoolActivityListResult
type ElasticPoolActivityProperties = original.ElasticPoolActivityProperties
type ElasticPoolDatabaseActivity = original.ElasticPoolDatabaseActivity
type ElasticPoolDatabaseActivityListResult = original.ElasticPoolDatabaseActivityListResult
type ElasticPoolDatabaseActivityProperties = original.ElasticPoolDatabaseActivityProperties
type ElasticPoolListResult = original.ElasticPoolListResult
type ElasticPoolProperties = original.ElasticPoolProperties
type ElasticPoolUpdate = original.ElasticPoolUpdate
type ExportRequest = original.ExportRequest
type FirewallRule = original.FirewallRule
type FirewallRuleListResult = original.FirewallRuleListResult
type FirewallRuleProperties = original.FirewallRuleProperties
type ImportExportResponse = original.ImportExportResponse
type ImportExportResponseProperties = original.ImportExportResponseProperties
type ImportExtensionProperties = original.ImportExtensionProperties
type ImportExtensionRequest = original.ImportExtensionRequest
type ImportRequest = original.ImportRequest
type OperationImpact = original.OperationImpact
type ProxyResource = original.ProxyResource
type RecommendedElasticPool = original.RecommendedElasticPool
type RecommendedElasticPoolListMetricsResult = original.RecommendedElasticPoolListMetricsResult
type RecommendedElasticPoolListResult = original.RecommendedElasticPoolListResult
type RecommendedElasticPoolMetric = original.RecommendedElasticPoolMetric
type RecommendedElasticPoolProperties = original.RecommendedElasticPoolProperties
type RecommendedIndex = original.RecommendedIndex
type RecommendedIndexProperties = original.RecommendedIndexProperties
type ReplicationLink = original.ReplicationLink
type ReplicationLinkListResult = original.ReplicationLinkListResult
type ReplicationLinkProperties = original.ReplicationLinkProperties
type Resource = original.Resource
type ServiceTierAdvisor = original.ServiceTierAdvisor
type ServiceTierAdvisorListResult = original.ServiceTierAdvisorListResult
type ServiceTierAdvisorProperties = original.ServiceTierAdvisorProperties
type SloUsageMetric = original.SloUsageMetric
type TrackedResource = original.TrackedResource
type TransparentDataEncryption = original.TransparentDataEncryption
type TransparentDataEncryptionActivity = original.TransparentDataEncryptionActivity
type TransparentDataEncryptionActivityListResult = original.TransparentDataEncryptionActivityListResult
type TransparentDataEncryptionActivityProperties = original.TransparentDataEncryptionActivityProperties
type TransparentDataEncryptionProperties = original.TransparentDataEncryptionProperties
type ServiceTierAdvisorsClient = original.ServiceTierAdvisorsClient
type ElasticPoolActivitiesClient = original.ElasticPoolActivitiesClient
type ElasticPoolDatabaseActivitiesClient = original.ElasticPoolDatabaseActivitiesClient
type TransparentDataEncryptionActivitiesClient = original.TransparentDataEncryptionActivitiesClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type DatabasesClient = original.DatabasesClient
type RecommendedElasticPoolsClient = original.RecommendedElasticPoolsClient
type ReplicationLinksClient = original.ReplicationLinksClient
type DatabaseThreatDetectionPoliciesClient = original.DatabaseThreatDetectionPoliciesClient
type ElasticPoolsClient = original.ElasticPoolsClient
type ServersClient = original.ServersClient
type TransparentDataEncryptionsClient = original.TransparentDataEncryptionsClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewDatabasesClient(subscriptionID string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRecommendedElasticPoolsClient(subscriptionID string) RecommendedElasticPoolsClient {
	return original.NewRecommendedElasticPoolsClient(subscriptionID)
}
func NewRecommendedElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) RecommendedElasticPoolsClient {
	return original.NewRecommendedElasticPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewReplicationLinksClient(subscriptionID string) ReplicationLinksClient {
	return original.NewReplicationLinksClient(subscriptionID)
}
func NewReplicationLinksClientWithBaseURI(baseURI string, subscriptionID string) ReplicationLinksClient {
	return original.NewReplicationLinksClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseThreatDetectionPoliciesClient(subscriptionID string) DatabaseThreatDetectionPoliciesClient {
	return original.NewDatabaseThreatDetectionPoliciesClient(subscriptionID)
}
func NewDatabaseThreatDetectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseThreatDetectionPoliciesClient {
	return original.NewDatabaseThreatDetectionPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewElasticPoolsClient(subscriptionID string) ElasticPoolsClient {
	return original.NewElasticPoolsClient(subscriptionID)
}
func NewElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolsClient {
	return original.NewElasticPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServersClient(subscriptionID string) ServersClient {
	return original.NewServersClient(subscriptionID)
}
func NewServersClientWithBaseURI(baseURI string, subscriptionID string) ServersClient {
	return original.NewServersClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransparentDataEncryptionsClient(subscriptionID string) TransparentDataEncryptionsClient {
	return original.NewTransparentDataEncryptionsClient(subscriptionID)
}
func NewTransparentDataEncryptionsClientWithBaseURI(baseURI string, subscriptionID string) TransparentDataEncryptionsClient {
	return original.NewTransparentDataEncryptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFirewallRulesClient(subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClient(subscriptionID)
}
func NewFirewallRulesClientWithBaseURI(baseURI string, subscriptionID string) FirewallRulesClient {
	return original.NewFirewallRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceTierAdvisorsClient(subscriptionID string) ServiceTierAdvisorsClient {
	return original.NewServiceTierAdvisorsClient(subscriptionID)
}
func NewServiceTierAdvisorsClientWithBaseURI(baseURI string, subscriptionID string) ServiceTierAdvisorsClient {
	return original.NewServiceTierAdvisorsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewElasticPoolActivitiesClient(subscriptionID string) ElasticPoolActivitiesClient {
	return original.NewElasticPoolActivitiesClient(subscriptionID)
}
func NewElasticPoolActivitiesClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolActivitiesClient {
	return original.NewElasticPoolActivitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewElasticPoolDatabaseActivitiesClient(subscriptionID string) ElasticPoolDatabaseActivitiesClient {
	return original.NewElasticPoolDatabaseActivitiesClient(subscriptionID)
}
func NewElasticPoolDatabaseActivitiesClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolDatabaseActivitiesClient {
	return original.NewElasticPoolDatabaseActivitiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransparentDataEncryptionActivitiesClient(subscriptionID string) TransparentDataEncryptionActivitiesClient {
	return original.NewTransparentDataEncryptionActivitiesClient(subscriptionID)
}
func NewTransparentDataEncryptionActivitiesClientWithBaseURI(baseURI string, subscriptionID string) TransparentDataEncryptionActivitiesClient {
	return original.NewTransparentDataEncryptionActivitiesClientWithBaseURI(baseURI, subscriptionID)
}
