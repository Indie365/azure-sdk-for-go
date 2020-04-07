package synapseapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/synapse/mgmt/2019-06-01-preview/synapse"
	"github.com/Azure/go-autorest/autorest"
)

// BigDataPoolsClientAPI contains the set of methods on the BigDataPoolsClient type.
type BigDataPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, bigDataPoolName string, bigDataPoolInfo synapse.BigDataPoolResourceInfo, force *bool) (result synapse.BigDataPoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, bigDataPoolName string) (result synapse.BigDataPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, bigDataPoolName string) (result synapse.BigDataPoolResourceInfo, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.BigDataPoolResourceInfoListResultPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.BigDataPoolResourceInfoListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, bigDataPoolName string, bigDataPoolPatchInfo synapse.BigDataPoolPatchInfo) (result synapse.BigDataPoolResourceInfo, err error)
}

var _ BigDataPoolsClientAPI = (*synapse.BigDataPoolsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	CheckNameAvailability(ctx context.Context, request synapse.CheckNameAvailabilityRequest) (result synapse.CheckNameAvailabilityResponse, err error)
	GetAzureAsyncHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, operationID string) (result synapse.SetObject, err error)
	GetLocationHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, operationID string) (result autorest.Response, err error)
	List(ctx context.Context) (result synapse.ListAvailableRpOperation, err error)
}

var _ OperationsClientAPI = (*synapse.OperationsClient)(nil)

// IPFirewallRulesClientAPI contains the set of methods on the IPFirewallRulesClient type.
type IPFirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string, IPFirewallRuleInfo synapse.IPFirewallRuleInfo) (result synapse.IPFirewallRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string) (result synapse.IPFirewallRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleName string) (result synapse.IPFirewallRuleInfo, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.IPFirewallRuleInfoListResultPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.IPFirewallRuleInfoListResultIterator, err error)
	ReplaceAll(ctx context.Context, resourceGroupName string, workspaceName string, request synapse.ReplaceAllIPFirewallRulesRequest) (result synapse.IPFirewallRulesReplaceAllFuture, err error)
}

var _ IPFirewallRulesClientAPI = (*synapse.IPFirewallRulesClient)(nil)

// SQLPoolsClientAPI contains the set of methods on the SQLPoolsClient type.
type SQLPoolsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, SQLPoolInfo synapse.SQLPool) (result synapse.SQLPoolsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPool, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.SQLPoolInfoListResultPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.SQLPoolInfoListResultIterator, err error)
	Pause(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolsPauseFuture, err error)
	Rename(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters synapse.ResourceMoveDefinition) (result autorest.Response, err error)
	Resume(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolsResumeFuture, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, SQLPoolInfo synapse.SQLPoolPatchInfo) (result synapse.SQLPool, err error)
}

var _ SQLPoolsClientAPI = (*synapse.SQLPoolsClient)(nil)

// SQLPoolMetadataSyncConfigsClientAPI contains the set of methods on the SQLPoolMetadataSyncConfigsClient type.
type SQLPoolMetadataSyncConfigsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, metadataSyncConfiguration synapse.MetadataSyncConfig) (result synapse.MetadataSyncConfig, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.MetadataSyncConfig, err error)
}

var _ SQLPoolMetadataSyncConfigsClientAPI = (*synapse.SQLPoolMetadataSyncConfigsClient)(nil)

// SQLPoolOperationResultsClientAPI contains the set of methods on the SQLPoolOperationResultsClient type.
type SQLPoolOperationResultsClientAPI interface {
	GetLocationHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, operationID string) (result synapse.SetObject, err error)
}

var _ SQLPoolOperationResultsClientAPI = (*synapse.SQLPoolOperationResultsClient)(nil)

// SQLPoolGeoBackupPoliciesClientAPI contains the set of methods on the SQLPoolGeoBackupPoliciesClient type.
type SQLPoolGeoBackupPoliciesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.GeoBackupPolicy, err error)
}

var _ SQLPoolGeoBackupPoliciesClientAPI = (*synapse.SQLPoolGeoBackupPoliciesClient)(nil)

// SQLPoolDataWarehouseUserActivitiesClientAPI contains the set of methods on the SQLPoolDataWarehouseUserActivitiesClient type.
type SQLPoolDataWarehouseUserActivitiesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.DataWarehouseUserActivities, err error)
}

var _ SQLPoolDataWarehouseUserActivitiesClientAPI = (*synapse.SQLPoolDataWarehouseUserActivitiesClient)(nil)

// SQLPoolRestorePointsClientAPI contains the set of methods on the SQLPoolRestorePointsClient type.
type SQLPoolRestorePointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters synapse.CreateSQLPoolRestorePointDefinition) (result synapse.SQLPoolRestorePointsCreateFuture, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.RestorePointListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.RestorePointListResultIterator, err error)
}

var _ SQLPoolRestorePointsClientAPI = (*synapse.SQLPoolRestorePointsClient)(nil)

// SQLPoolReplicationLinksClientAPI contains the set of methods on the SQLPoolReplicationLinksClient type.
type SQLPoolReplicationLinksClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.ReplicationLinkListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.ReplicationLinkListResultIterator, err error)
}

var _ SQLPoolReplicationLinksClientAPI = (*synapse.SQLPoolReplicationLinksClient)(nil)

// SQLPoolTransparentDataEncryptionsClientAPI contains the set of methods on the SQLPoolTransparentDataEncryptionsClient type.
type SQLPoolTransparentDataEncryptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters synapse.TransparentDataEncryption) (result synapse.TransparentDataEncryption, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.TransparentDataEncryption, err error)
}

var _ SQLPoolTransparentDataEncryptionsClientAPI = (*synapse.SQLPoolTransparentDataEncryptionsClient)(nil)

// SQLPoolBlobAuditingPoliciesClientAPI contains the set of methods on the SQLPoolBlobAuditingPoliciesClient type.
type SQLPoolBlobAuditingPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters synapse.SQLPoolBlobAuditingPolicy) (result synapse.SQLPoolBlobAuditingPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolBlobAuditingPolicy, err error)
}

var _ SQLPoolBlobAuditingPoliciesClientAPI = (*synapse.SQLPoolBlobAuditingPoliciesClient)(nil)

// SQLPoolOperationsClientAPI contains the set of methods on the SQLPoolOperationsClient type.
type SQLPoolOperationsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolBlobAuditingPolicySQLPoolOperationListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolBlobAuditingPolicySQLPoolOperationListResultIterator, err error)
}

var _ SQLPoolOperationsClientAPI = (*synapse.SQLPoolOperationsClient)(nil)

// SQLPoolUsagesClientAPI contains the set of methods on the SQLPoolUsagesClient type.
type SQLPoolUsagesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolUsageListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolUsageListResultIterator, err error)
}

var _ SQLPoolUsagesClientAPI = (*synapse.SQLPoolUsagesClient)(nil)

// SQLPoolSensitivityLabelsClientAPI contains the set of methods on the SQLPoolSensitivityLabelsClient type.
type SQLPoolSensitivityLabelsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string, parameters synapse.SensitivityLabel) (result synapse.SensitivityLabel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error)
	DisableRecommendation(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error)
	EnableRecommendation(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, columnName string) (result autorest.Response, err error)
	ListCurrent(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, filter string) (result synapse.SensitivityLabelListResultPage, err error)
	ListCurrentComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, filter string) (result synapse.SensitivityLabelListResultIterator, err error)
	ListRecommended(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result synapse.SensitivityLabelListResultPage, err error)
	ListRecommendedComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, includeDisabledRecommendations *bool, skipToken string, filter string) (result synapse.SensitivityLabelListResultIterator, err error)
}

var _ SQLPoolSensitivityLabelsClientAPI = (*synapse.SQLPoolSensitivityLabelsClient)(nil)

// SQLPoolSchemasClientAPI contains the set of methods on the SQLPoolSchemasClient type.
type SQLPoolSchemasClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, filter string) (result synapse.SQLPoolSchemaListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, filter string) (result synapse.SQLPoolSchemaListResultIterator, err error)
}

var _ SQLPoolSchemasClientAPI = (*synapse.SQLPoolSchemasClient)(nil)

// SQLPoolTablesClientAPI contains the set of methods on the SQLPoolTablesClient type.
type SQLPoolTablesClientAPI interface {
	ListBySchema(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, filter string) (result synapse.SQLPoolTableListResultPage, err error)
	ListBySchemaComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, filter string) (result synapse.SQLPoolTableListResultIterator, err error)
}

var _ SQLPoolTablesClientAPI = (*synapse.SQLPoolTablesClient)(nil)

// SQLPoolTableColumnsClientAPI contains the set of methods on the SQLPoolTableColumnsClient type.
type SQLPoolTableColumnsClientAPI interface {
	ListByTableName(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, filter string) (result synapse.SQLPoolColumnListResultPage, err error)
	ListByTableNameComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, schemaName string, tableName string, filter string) (result synapse.SQLPoolColumnListResultIterator, err error)
}

var _ SQLPoolTableColumnsClientAPI = (*synapse.SQLPoolTableColumnsClient)(nil)

// SQLPoolConnectionPoliciesClientAPI contains the set of methods on the SQLPoolConnectionPoliciesClient type.
type SQLPoolConnectionPoliciesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolConnectionPolicy, err error)
}

var _ SQLPoolConnectionPoliciesClientAPI = (*synapse.SQLPoolConnectionPoliciesClient)(nil)

// SQLPoolVulnerabilityAssessmentsClientAPI contains the set of methods on the SQLPoolVulnerabilityAssessmentsClient type.
type SQLPoolVulnerabilityAssessmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters synapse.SQLPoolVulnerabilityAssessment) (result synapse.SQLPoolVulnerabilityAssessment, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolVulnerabilityAssessment, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolVulnerabilityAssessmentListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolVulnerabilityAssessmentListResultIterator, err error)
}

var _ SQLPoolVulnerabilityAssessmentsClientAPI = (*synapse.SQLPoolVulnerabilityAssessmentsClient)(nil)

// SQLPoolVulnerabilityAssessmentScansClientAPI contains the set of methods on the SQLPoolVulnerabilityAssessmentScansClient type.
type SQLPoolVulnerabilityAssessmentScansClientAPI interface {
	Export(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, scanID string) (result synapse.SQLPoolVulnerabilityAssessmentScansExport, err error)
	InitiateScan(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, scanID string) (result synapse.SQLPoolVulnerabilityAssessmentScansInitiateScanFuture, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.VulnerabilityAssessmentScanRecordListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.VulnerabilityAssessmentScanRecordListResultIterator, err error)
}

var _ SQLPoolVulnerabilityAssessmentScansClientAPI = (*synapse.SQLPoolVulnerabilityAssessmentScansClient)(nil)

// SQLPoolSecurityAlertPoliciesClientAPI contains the set of methods on the SQLPoolSecurityAlertPoliciesClient type.
type SQLPoolSecurityAlertPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, parameters synapse.SQLPoolSecurityAlertPolicy) (result synapse.SQLPoolSecurityAlertPolicy, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result synapse.SQLPoolSecurityAlertPolicy, err error)
}

var _ SQLPoolSecurityAlertPoliciesClientAPI = (*synapse.SQLPoolSecurityAlertPoliciesClient)(nil)

// SQLPoolVulnerabilityAssessmentRuleBaselinesClientAPI contains the set of methods on the SQLPoolVulnerabilityAssessmentRuleBaselinesClient type.
type SQLPoolVulnerabilityAssessmentRuleBaselinesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, ruleID string, baselineName synapse.VulnerabilityAssessmentPolicyBaselineName, parameters synapse.SQLPoolVulnerabilityAssessmentRuleBaseline) (result synapse.SQLPoolVulnerabilityAssessmentRuleBaseline, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, ruleID string, baselineName synapse.VulnerabilityAssessmentPolicyBaselineName) (result autorest.Response, err error)
}

var _ SQLPoolVulnerabilityAssessmentRuleBaselinesClientAPI = (*synapse.SQLPoolVulnerabilityAssessmentRuleBaselinesClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceInfo synapse.Workspace) (result synapse.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.Workspace, err error)
	List(ctx context.Context) (result synapse.WorkspaceInfoListResultPage, err error)
	ListComplete(ctx context.Context) (result synapse.WorkspaceInfoListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result synapse.WorkspaceInfoListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result synapse.WorkspaceInfoListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchInfo synapse.WorkspacePatchInfo) (result synapse.WorkspacesUpdateFuture, err error)
}

var _ WorkspacesClientAPI = (*synapse.WorkspacesClient)(nil)

// WorkspaceAadAdminsClientAPI contains the set of methods on the WorkspaceAadAdminsClient type.
type WorkspaceAadAdminsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, aadAdminInfo synapse.WorkspaceAadAdminInfo) (result synapse.WorkspaceAadAdminsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.WorkspaceAadAdminsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.WorkspaceAadAdminInfo, err error)
}

var _ WorkspaceAadAdminsClientAPI = (*synapse.WorkspaceAadAdminsClient)(nil)

// WorkspaceManagedIdentitySQLControlSettingsClientAPI contains the set of methods on the WorkspaceManagedIdentitySQLControlSettingsClient type.
type WorkspaceManagedIdentitySQLControlSettingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, managedIdentitySQLControlSettings synapse.ManagedIdentitySQLControlSettingsModel) (result synapse.ManagedIdentitySQLControlSettingsModel, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.ManagedIdentitySQLControlSettingsModel, err error)
}

var _ WorkspaceManagedIdentitySQLControlSettingsClientAPI = (*synapse.WorkspaceManagedIdentitySQLControlSettingsClient)(nil)

// IntegrationRuntimesClientAPI contains the set of methods on the IntegrationRuntimesClient type.
type IntegrationRuntimesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, integrationRuntime synapse.IntegrationRuntimeResource, ifMatch string) (result synapse.IntegrationRuntimeResource, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, ifNoneMatch string) (result synapse.IntegrationRuntimeResource, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.IntegrationRuntimeListResponsePage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result synapse.IntegrationRuntimeListResponseIterator, err error)
	Start(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result synapse.IntegrationRuntimeStatusResponse, err error)
	Stop(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, updateIntegrationRuntimeRequest synapse.UpdateIntegrationRuntimeRequest) (result synapse.IntegrationRuntimeResource, err error)
	Upgrade(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result autorest.Response, err error)
}

var _ IntegrationRuntimesClientAPI = (*synapse.IntegrationRuntimesClient)(nil)

// IntegrationRuntimeNodeIPAddressClientAPI contains the set of methods on the IntegrationRuntimeNodeIPAddressClient type.
type IntegrationRuntimeNodeIPAddressClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string) (result synapse.IntegrationRuntimeNodeIPAddress, err error)
}

var _ IntegrationRuntimeNodeIPAddressClientAPI = (*synapse.IntegrationRuntimeNodeIPAddressClient)(nil)

// IntegrationRuntimeObjectMetadataClientAPI contains the set of methods on the IntegrationRuntimeObjectMetadataClient type.
type IntegrationRuntimeObjectMetadataClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, getMetadataRequest *synapse.GetSsisObjectMetadataRequest) (result synapse.SsisObjectMetadataListResponse, err error)
	Refresh(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result synapse.SsisObjectMetadataStatusResponse, err error)
}

var _ IntegrationRuntimeObjectMetadataClientAPI = (*synapse.IntegrationRuntimeObjectMetadataClient)(nil)

// IntegrationRuntimeNodesClientAPI contains the set of methods on the IntegrationRuntimeNodesClient type.
type IntegrationRuntimeNodesClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string) (result synapse.SelfHostedIntegrationRuntimeNode, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest synapse.UpdateIntegrationRuntimeNodeRequest) (result synapse.SelfHostedIntegrationRuntimeNode, err error)
}

var _ IntegrationRuntimeNodesClientAPI = (*synapse.IntegrationRuntimeNodesClient)(nil)

// IntegrationRuntimeCredentialsClientAPI contains the set of methods on the IntegrationRuntimeCredentialsClient type.
type IntegrationRuntimeCredentialsClientAPI interface {
	Sync(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result autorest.Response, err error)
}

var _ IntegrationRuntimeCredentialsClientAPI = (*synapse.IntegrationRuntimeCredentialsClient)(nil)

// IntegrationRuntimeConnectionInfosClientAPI contains the set of methods on the IntegrationRuntimeConnectionInfosClient type.
type IntegrationRuntimeConnectionInfosClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result synapse.IntegrationRuntimeConnectionInfo, err error)
}

var _ IntegrationRuntimeConnectionInfosClientAPI = (*synapse.IntegrationRuntimeConnectionInfosClient)(nil)

// IntegrationRuntimeAuthKeysClientAPI contains the set of methods on the IntegrationRuntimeAuthKeysClient type.
type IntegrationRuntimeAuthKeysClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result synapse.IntegrationRuntimeAuthKeys, err error)
	Regenerate(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, regenerateKeyParameters synapse.IntegrationRuntimeRegenerateKeyParameters) (result synapse.IntegrationRuntimeAuthKeys, err error)
}

var _ IntegrationRuntimeAuthKeysClientAPI = (*synapse.IntegrationRuntimeAuthKeysClient)(nil)

// IntegrationRuntimeMonitoringDataClientAPI contains the set of methods on the IntegrationRuntimeMonitoringDataClient type.
type IntegrationRuntimeMonitoringDataClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result synapse.IntegrationRuntimeMonitoringData, err error)
}

var _ IntegrationRuntimeMonitoringDataClientAPI = (*synapse.IntegrationRuntimeMonitoringDataClient)(nil)

// IntegrationRuntimeStatusClientAPI contains the set of methods on the IntegrationRuntimeStatusClient type.
type IntegrationRuntimeStatusClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string) (result synapse.IntegrationRuntimeStatusResponse, err error)
}

var _ IntegrationRuntimeStatusClientAPI = (*synapse.IntegrationRuntimeStatusClient)(nil)
