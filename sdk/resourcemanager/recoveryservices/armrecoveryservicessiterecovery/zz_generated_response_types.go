//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

// MigrationRecoveryPointsClientGetResponse contains the response from method MigrationRecoveryPointsClient.Get.
type MigrationRecoveryPointsClientGetResponse struct {
	MigrationRecoveryPoint
}

// MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse contains the response from method MigrationRecoveryPointsClient.ListByReplicationMigrationItems.
type MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse struct {
	MigrationRecoveryPointCollection
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsDiscoveryCollection
}

// RecoveryPointsClientGetResponse contains the response from method RecoveryPointsClient.Get.
type RecoveryPointsClientGetResponse struct {
	RecoveryPoint
}

// RecoveryPointsClientListByReplicationProtectedItemsResponse contains the response from method RecoveryPointsClient.ListByReplicationProtectedItems.
type RecoveryPointsClientListByReplicationProtectedItemsResponse struct {
	RecoveryPointCollection
}

// ReplicationAlertSettingsClientCreateResponse contains the response from method ReplicationAlertSettingsClient.Create.
type ReplicationAlertSettingsClientCreateResponse struct {
	Alert
}

// ReplicationAlertSettingsClientGetResponse contains the response from method ReplicationAlertSettingsClient.Get.
type ReplicationAlertSettingsClientGetResponse struct {
	Alert
}

// ReplicationAlertSettingsClientListResponse contains the response from method ReplicationAlertSettingsClient.List.
type ReplicationAlertSettingsClientListResponse struct {
	AlertCollection
}

// ReplicationAppliancesClientListResponse contains the response from method ReplicationAppliancesClient.List.
type ReplicationAppliancesClientListResponse struct {
	ApplianceCollection
}

// ReplicationEligibilityResultsClientGetResponse contains the response from method ReplicationEligibilityResultsClient.Get.
type ReplicationEligibilityResultsClientGetResponse struct {
	ReplicationEligibilityResults
}

// ReplicationEligibilityResultsClientListResponse contains the response from method ReplicationEligibilityResultsClient.List.
type ReplicationEligibilityResultsClientListResponse struct {
	ReplicationEligibilityResultsCollection
}

// ReplicationEventsClientGetResponse contains the response from method ReplicationEventsClient.Get.
type ReplicationEventsClientGetResponse struct {
	Event
}

// ReplicationEventsClientListResponse contains the response from method ReplicationEventsClient.List.
type ReplicationEventsClientListResponse struct {
	EventCollection
}

// ReplicationFabricsClientCheckConsistencyResponse contains the response from method ReplicationFabricsClient.CheckConsistency.
type ReplicationFabricsClientCheckConsistencyResponse struct {
	Fabric
}

// ReplicationFabricsClientCreateResponse contains the response from method ReplicationFabricsClient.Create.
type ReplicationFabricsClientCreateResponse struct {
	Fabric
}

// ReplicationFabricsClientDeleteResponse contains the response from method ReplicationFabricsClient.Delete.
type ReplicationFabricsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationFabricsClientGetResponse contains the response from method ReplicationFabricsClient.Get.
type ReplicationFabricsClientGetResponse struct {
	Fabric
}

// ReplicationFabricsClientListResponse contains the response from method ReplicationFabricsClient.List.
type ReplicationFabricsClientListResponse struct {
	FabricCollection
}

// ReplicationFabricsClientMigrateToAADResponse contains the response from method ReplicationFabricsClient.MigrateToAAD.
type ReplicationFabricsClientMigrateToAADResponse struct {
	// placeholder for future response values
}

// ReplicationFabricsClientPurgeResponse contains the response from method ReplicationFabricsClient.Purge.
type ReplicationFabricsClientPurgeResponse struct {
	// placeholder for future response values
}

// ReplicationFabricsClientReassociateGatewayResponse contains the response from method ReplicationFabricsClient.ReassociateGateway.
type ReplicationFabricsClientReassociateGatewayResponse struct {
	Fabric
}

// ReplicationFabricsClientRenewCertificateResponse contains the response from method ReplicationFabricsClient.RenewCertificate.
type ReplicationFabricsClientRenewCertificateResponse struct {
	Fabric
}

// ReplicationJobsClientCancelResponse contains the response from method ReplicationJobsClient.Cancel.
type ReplicationJobsClientCancelResponse struct {
	Job
}

// ReplicationJobsClientExportResponse contains the response from method ReplicationJobsClient.Export.
type ReplicationJobsClientExportResponse struct {
	Job
}

// ReplicationJobsClientGetResponse contains the response from method ReplicationJobsClient.Get.
type ReplicationJobsClientGetResponse struct {
	Job
}

// ReplicationJobsClientListResponse contains the response from method ReplicationJobsClient.List.
type ReplicationJobsClientListResponse struct {
	JobCollection
}

// ReplicationJobsClientRestartResponse contains the response from method ReplicationJobsClient.Restart.
type ReplicationJobsClientRestartResponse struct {
	Job
}

// ReplicationJobsClientResumeResponse contains the response from method ReplicationJobsClient.Resume.
type ReplicationJobsClientResumeResponse struct {
	Job
}

// ReplicationLogicalNetworksClientGetResponse contains the response from method ReplicationLogicalNetworksClient.Get.
type ReplicationLogicalNetworksClientGetResponse struct {
	LogicalNetwork
}

// ReplicationLogicalNetworksClientListByReplicationFabricsResponse contains the response from method ReplicationLogicalNetworksClient.ListByReplicationFabrics.
type ReplicationLogicalNetworksClientListByReplicationFabricsResponse struct {
	LogicalNetworkCollection
}

// ReplicationMigrationItemsClientCreateResponse contains the response from method ReplicationMigrationItemsClient.Create.
type ReplicationMigrationItemsClientCreateResponse struct {
	MigrationItem
}

// ReplicationMigrationItemsClientDeleteResponse contains the response from method ReplicationMigrationItemsClient.Delete.
type ReplicationMigrationItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationMigrationItemsClientGetResponse contains the response from method ReplicationMigrationItemsClient.Get.
type ReplicationMigrationItemsClientGetResponse struct {
	MigrationItem
}

// ReplicationMigrationItemsClientListByReplicationProtectionContainersResponse contains the response from method ReplicationMigrationItemsClient.ListByReplicationProtectionContainers.
type ReplicationMigrationItemsClientListByReplicationProtectionContainersResponse struct {
	MigrationItemCollection
}

// ReplicationMigrationItemsClientListResponse contains the response from method ReplicationMigrationItemsClient.List.
type ReplicationMigrationItemsClientListResponse struct {
	MigrationItemCollection
}

// ReplicationMigrationItemsClientMigrateResponse contains the response from method ReplicationMigrationItemsClient.Migrate.
type ReplicationMigrationItemsClientMigrateResponse struct {
	MigrationItem
}

// ReplicationMigrationItemsClientResyncResponse contains the response from method ReplicationMigrationItemsClient.Resync.
type ReplicationMigrationItemsClientResyncResponse struct {
	MigrationItem
}

// ReplicationMigrationItemsClientTestMigrateCleanupResponse contains the response from method ReplicationMigrationItemsClient.TestMigrateCleanup.
type ReplicationMigrationItemsClientTestMigrateCleanupResponse struct {
	MigrationItem
}

// ReplicationMigrationItemsClientTestMigrateResponse contains the response from method ReplicationMigrationItemsClient.TestMigrate.
type ReplicationMigrationItemsClientTestMigrateResponse struct {
	MigrationItem
}

// ReplicationMigrationItemsClientUpdateResponse contains the response from method ReplicationMigrationItemsClient.Update.
type ReplicationMigrationItemsClientUpdateResponse struct {
	MigrationItem
}

// ReplicationNetworkMappingsClientCreateResponse contains the response from method ReplicationNetworkMappingsClient.Create.
type ReplicationNetworkMappingsClientCreateResponse struct {
	NetworkMapping
}

// ReplicationNetworkMappingsClientDeleteResponse contains the response from method ReplicationNetworkMappingsClient.Delete.
type ReplicationNetworkMappingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationNetworkMappingsClientGetResponse contains the response from method ReplicationNetworkMappingsClient.Get.
type ReplicationNetworkMappingsClientGetResponse struct {
	NetworkMapping
}

// ReplicationNetworkMappingsClientListByReplicationNetworksResponse contains the response from method ReplicationNetworkMappingsClient.ListByReplicationNetworks.
type ReplicationNetworkMappingsClientListByReplicationNetworksResponse struct {
	NetworkMappingCollection
}

// ReplicationNetworkMappingsClientListResponse contains the response from method ReplicationNetworkMappingsClient.List.
type ReplicationNetworkMappingsClientListResponse struct {
	NetworkMappingCollection
}

// ReplicationNetworkMappingsClientUpdateResponse contains the response from method ReplicationNetworkMappingsClient.Update.
type ReplicationNetworkMappingsClientUpdateResponse struct {
	NetworkMapping
}

// ReplicationNetworksClientGetResponse contains the response from method ReplicationNetworksClient.Get.
type ReplicationNetworksClientGetResponse struct {
	Network
}

// ReplicationNetworksClientListByReplicationFabricsResponse contains the response from method ReplicationNetworksClient.ListByReplicationFabrics.
type ReplicationNetworksClientListByReplicationFabricsResponse struct {
	NetworkCollection
}

// ReplicationNetworksClientListResponse contains the response from method ReplicationNetworksClient.List.
type ReplicationNetworksClientListResponse struct {
	NetworkCollection
}

// ReplicationPoliciesClientCreateResponse contains the response from method ReplicationPoliciesClient.Create.
type ReplicationPoliciesClientCreateResponse struct {
	Policy
}

// ReplicationPoliciesClientDeleteResponse contains the response from method ReplicationPoliciesClient.Delete.
type ReplicationPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationPoliciesClientGetResponse contains the response from method ReplicationPoliciesClient.Get.
type ReplicationPoliciesClientGetResponse struct {
	Policy
}

// ReplicationPoliciesClientListResponse contains the response from method ReplicationPoliciesClient.List.
type ReplicationPoliciesClientListResponse struct {
	PolicyCollection
}

// ReplicationPoliciesClientUpdateResponse contains the response from method ReplicationPoliciesClient.Update.
type ReplicationPoliciesClientUpdateResponse struct {
	Policy
}

// ReplicationProtectableItemsClientGetResponse contains the response from method ReplicationProtectableItemsClient.Get.
type ReplicationProtectableItemsClientGetResponse struct {
	ProtectableItem
}

// ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse contains the response from method ReplicationProtectableItemsClient.ListByReplicationProtectionContainers.
type ReplicationProtectableItemsClientListByReplicationProtectionContainersResponse struct {
	ProtectableItemCollection
}

// ReplicationProtectedItemsClientAddDisksResponse contains the response from method ReplicationProtectedItemsClient.AddDisks.
type ReplicationProtectedItemsClientAddDisksResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientApplyRecoveryPointResponse contains the response from method ReplicationProtectedItemsClient.ApplyRecoveryPoint.
type ReplicationProtectedItemsClientApplyRecoveryPointResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientCreateResponse contains the response from method ReplicationProtectedItemsClient.Create.
type ReplicationProtectedItemsClientCreateResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientDeleteResponse contains the response from method ReplicationProtectedItemsClient.Delete.
type ReplicationProtectedItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationProtectedItemsClientFailoverCancelResponse contains the response from method ReplicationProtectedItemsClient.FailoverCancel.
type ReplicationProtectedItemsClientFailoverCancelResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientFailoverCommitResponse contains the response from method ReplicationProtectedItemsClient.FailoverCommit.
type ReplicationProtectedItemsClientFailoverCommitResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientGetResponse contains the response from method ReplicationProtectedItemsClient.Get.
type ReplicationProtectedItemsClientGetResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientListByReplicationProtectionContainersResponse contains the response from method ReplicationProtectedItemsClient.ListByReplicationProtectionContainers.
type ReplicationProtectedItemsClientListByReplicationProtectionContainersResponse struct {
	ReplicationProtectedItemCollection
}

// ReplicationProtectedItemsClientListResponse contains the response from method ReplicationProtectedItemsClient.List.
type ReplicationProtectedItemsClientListResponse struct {
	ReplicationProtectedItemCollection
}

// ReplicationProtectedItemsClientPlannedFailoverResponse contains the response from method ReplicationProtectedItemsClient.PlannedFailover.
type ReplicationProtectedItemsClientPlannedFailoverResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientPurgeResponse contains the response from method ReplicationProtectedItemsClient.Purge.
type ReplicationProtectedItemsClientPurgeResponse struct {
	// placeholder for future response values
}

// ReplicationProtectedItemsClientRemoveDisksResponse contains the response from method ReplicationProtectedItemsClient.RemoveDisks.
type ReplicationProtectedItemsClientRemoveDisksResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientRepairReplicationResponse contains the response from method ReplicationProtectedItemsClient.RepairReplication.
type ReplicationProtectedItemsClientRepairReplicationResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientReprotectResponse contains the response from method ReplicationProtectedItemsClient.Reprotect.
type ReplicationProtectedItemsClientReprotectResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientResolveHealthErrorsResponse contains the response from method ReplicationProtectedItemsClient.ResolveHealthErrors.
type ReplicationProtectedItemsClientResolveHealthErrorsResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientSwitchProviderResponse contains the response from method ReplicationProtectedItemsClient.SwitchProvider.
type ReplicationProtectedItemsClientSwitchProviderResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientTestFailoverCleanupResponse contains the response from method ReplicationProtectedItemsClient.TestFailoverCleanup.
type ReplicationProtectedItemsClientTestFailoverCleanupResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientTestFailoverResponse contains the response from method ReplicationProtectedItemsClient.TestFailover.
type ReplicationProtectedItemsClientTestFailoverResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientUnplannedFailoverResponse contains the response from method ReplicationProtectedItemsClient.UnplannedFailover.
type ReplicationProtectedItemsClientUnplannedFailoverResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientUpdateApplianceResponse contains the response from method ReplicationProtectedItemsClient.UpdateAppliance.
type ReplicationProtectedItemsClientUpdateApplianceResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientUpdateMobilityServiceResponse contains the response from method ReplicationProtectedItemsClient.UpdateMobilityService.
type ReplicationProtectedItemsClientUpdateMobilityServiceResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectedItemsClientUpdateResponse contains the response from method ReplicationProtectedItemsClient.Update.
type ReplicationProtectedItemsClientUpdateResponse struct {
	ReplicationProtectedItem
}

// ReplicationProtectionContainerMappingsClientCreateResponse contains the response from method ReplicationProtectionContainerMappingsClient.Create.
type ReplicationProtectionContainerMappingsClientCreateResponse struct {
	ProtectionContainerMapping
}

// ReplicationProtectionContainerMappingsClientDeleteResponse contains the response from method ReplicationProtectionContainerMappingsClient.Delete.
type ReplicationProtectionContainerMappingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationProtectionContainerMappingsClientGetResponse contains the response from method ReplicationProtectionContainerMappingsClient.Get.
type ReplicationProtectionContainerMappingsClientGetResponse struct {
	ProtectionContainerMapping
}

// ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse contains the response from method
// ReplicationProtectionContainerMappingsClient.ListByReplicationProtectionContainers.
type ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse struct {
	ProtectionContainerMappingCollection
}

// ReplicationProtectionContainerMappingsClientListResponse contains the response from method ReplicationProtectionContainerMappingsClient.List.
type ReplicationProtectionContainerMappingsClientListResponse struct {
	ProtectionContainerMappingCollection
}

// ReplicationProtectionContainerMappingsClientPurgeResponse contains the response from method ReplicationProtectionContainerMappingsClient.Purge.
type ReplicationProtectionContainerMappingsClientPurgeResponse struct {
	// placeholder for future response values
}

// ReplicationProtectionContainerMappingsClientUpdateResponse contains the response from method ReplicationProtectionContainerMappingsClient.Update.
type ReplicationProtectionContainerMappingsClientUpdateResponse struct {
	ProtectionContainerMapping
}

// ReplicationProtectionContainersClientCreateResponse contains the response from method ReplicationProtectionContainersClient.Create.
type ReplicationProtectionContainersClientCreateResponse struct {
	ProtectionContainer
}

// ReplicationProtectionContainersClientDeleteResponse contains the response from method ReplicationProtectionContainersClient.Delete.
type ReplicationProtectionContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationProtectionContainersClientDiscoverProtectableItemResponse contains the response from method ReplicationProtectionContainersClient.DiscoverProtectableItem.
type ReplicationProtectionContainersClientDiscoverProtectableItemResponse struct {
	ProtectionContainer
}

// ReplicationProtectionContainersClientGetResponse contains the response from method ReplicationProtectionContainersClient.Get.
type ReplicationProtectionContainersClientGetResponse struct {
	ProtectionContainer
}

// ReplicationProtectionContainersClientListByReplicationFabricsResponse contains the response from method ReplicationProtectionContainersClient.ListByReplicationFabrics.
type ReplicationProtectionContainersClientListByReplicationFabricsResponse struct {
	ProtectionContainerCollection
}

// ReplicationProtectionContainersClientListResponse contains the response from method ReplicationProtectionContainersClient.List.
type ReplicationProtectionContainersClientListResponse struct {
	ProtectionContainerCollection
}

// ReplicationProtectionContainersClientSwitchProtectionResponse contains the response from method ReplicationProtectionContainersClient.SwitchProtection.
type ReplicationProtectionContainersClientSwitchProtectionResponse struct {
	ProtectionContainer
}

// ReplicationProtectionIntentsClientCreateResponse contains the response from method ReplicationProtectionIntentsClient.Create.
type ReplicationProtectionIntentsClientCreateResponse struct {
	ReplicationProtectionIntent
}

// ReplicationProtectionIntentsClientGetResponse contains the response from method ReplicationProtectionIntentsClient.Get.
type ReplicationProtectionIntentsClientGetResponse struct {
	ReplicationProtectionIntent
}

// ReplicationProtectionIntentsClientListResponse contains the response from method ReplicationProtectionIntentsClient.List.
type ReplicationProtectionIntentsClientListResponse struct {
	ReplicationProtectionIntentCollection
}

// ReplicationRecoveryPlansClientCreateResponse contains the response from method ReplicationRecoveryPlansClient.Create.
type ReplicationRecoveryPlansClientCreateResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientDeleteResponse contains the response from method ReplicationRecoveryPlansClient.Delete.
type ReplicationRecoveryPlansClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationRecoveryPlansClientFailoverCancelResponse contains the response from method ReplicationRecoveryPlansClient.FailoverCancel.
type ReplicationRecoveryPlansClientFailoverCancelResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientFailoverCommitResponse contains the response from method ReplicationRecoveryPlansClient.FailoverCommit.
type ReplicationRecoveryPlansClientFailoverCommitResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientGetResponse contains the response from method ReplicationRecoveryPlansClient.Get.
type ReplicationRecoveryPlansClientGetResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientListResponse contains the response from method ReplicationRecoveryPlansClient.List.
type ReplicationRecoveryPlansClientListResponse struct {
	RecoveryPlanCollection
}

// ReplicationRecoveryPlansClientPlannedFailoverResponse contains the response from method ReplicationRecoveryPlansClient.PlannedFailover.
type ReplicationRecoveryPlansClientPlannedFailoverResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientReprotectResponse contains the response from method ReplicationRecoveryPlansClient.Reprotect.
type ReplicationRecoveryPlansClientReprotectResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientTestFailoverCleanupResponse contains the response from method ReplicationRecoveryPlansClient.TestFailoverCleanup.
type ReplicationRecoveryPlansClientTestFailoverCleanupResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientTestFailoverResponse contains the response from method ReplicationRecoveryPlansClient.TestFailover.
type ReplicationRecoveryPlansClientTestFailoverResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientUnplannedFailoverResponse contains the response from method ReplicationRecoveryPlansClient.UnplannedFailover.
type ReplicationRecoveryPlansClientUnplannedFailoverResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryPlansClientUpdateResponse contains the response from method ReplicationRecoveryPlansClient.Update.
type ReplicationRecoveryPlansClientUpdateResponse struct {
	RecoveryPlan
}

// ReplicationRecoveryServicesProvidersClientCreateResponse contains the response from method ReplicationRecoveryServicesProvidersClient.Create.
type ReplicationRecoveryServicesProvidersClientCreateResponse struct {
	RecoveryServicesProvider
}

// ReplicationRecoveryServicesProvidersClientDeleteResponse contains the response from method ReplicationRecoveryServicesProvidersClient.Delete.
type ReplicationRecoveryServicesProvidersClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationRecoveryServicesProvidersClientGetResponse contains the response from method ReplicationRecoveryServicesProvidersClient.Get.
type ReplicationRecoveryServicesProvidersClientGetResponse struct {
	RecoveryServicesProvider
}

// ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse contains the response from method ReplicationRecoveryServicesProvidersClient.ListByReplicationFabrics.
type ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse struct {
	RecoveryServicesProviderCollection
}

// ReplicationRecoveryServicesProvidersClientListResponse contains the response from method ReplicationRecoveryServicesProvidersClient.List.
type ReplicationRecoveryServicesProvidersClientListResponse struct {
	RecoveryServicesProviderCollection
}

// ReplicationRecoveryServicesProvidersClientPurgeResponse contains the response from method ReplicationRecoveryServicesProvidersClient.Purge.
type ReplicationRecoveryServicesProvidersClientPurgeResponse struct {
	// placeholder for future response values
}

// ReplicationRecoveryServicesProvidersClientRefreshProviderResponse contains the response from method ReplicationRecoveryServicesProvidersClient.RefreshProvider.
type ReplicationRecoveryServicesProvidersClientRefreshProviderResponse struct {
	RecoveryServicesProvider
}

// ReplicationStorageClassificationMappingsClientCreateResponse contains the response from method ReplicationStorageClassificationMappingsClient.Create.
type ReplicationStorageClassificationMappingsClientCreateResponse struct {
	StorageClassificationMapping
}

// ReplicationStorageClassificationMappingsClientDeleteResponse contains the response from method ReplicationStorageClassificationMappingsClient.Delete.
type ReplicationStorageClassificationMappingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationStorageClassificationMappingsClientGetResponse contains the response from method ReplicationStorageClassificationMappingsClient.Get.
type ReplicationStorageClassificationMappingsClientGetResponse struct {
	StorageClassificationMapping
}

// ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse contains the response from
// method ReplicationStorageClassificationMappingsClient.ListByReplicationStorageClassifications.
type ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse struct {
	StorageClassificationMappingCollection
}

// ReplicationStorageClassificationMappingsClientListResponse contains the response from method ReplicationStorageClassificationMappingsClient.List.
type ReplicationStorageClassificationMappingsClientListResponse struct {
	StorageClassificationMappingCollection
}

// ReplicationStorageClassificationsClientGetResponse contains the response from method ReplicationStorageClassificationsClient.Get.
type ReplicationStorageClassificationsClientGetResponse struct {
	StorageClassification
}

// ReplicationStorageClassificationsClientListByReplicationFabricsResponse contains the response from method ReplicationStorageClassificationsClient.ListByReplicationFabrics.
type ReplicationStorageClassificationsClientListByReplicationFabricsResponse struct {
	StorageClassificationCollection
}

// ReplicationStorageClassificationsClientListResponse contains the response from method ReplicationStorageClassificationsClient.List.
type ReplicationStorageClassificationsClientListResponse struct {
	StorageClassificationCollection
}

// ReplicationVaultHealthClientGetResponse contains the response from method ReplicationVaultHealthClient.Get.
type ReplicationVaultHealthClientGetResponse struct {
	VaultHealthDetails
}

// ReplicationVaultHealthClientRefreshResponse contains the response from method ReplicationVaultHealthClient.Refresh.
type ReplicationVaultHealthClientRefreshResponse struct {
	VaultHealthDetails
}

// ReplicationVaultSettingClientCreateResponse contains the response from method ReplicationVaultSettingClient.Create.
type ReplicationVaultSettingClientCreateResponse struct {
	VaultSetting
}

// ReplicationVaultSettingClientGetResponse contains the response from method ReplicationVaultSettingClient.Get.
type ReplicationVaultSettingClientGetResponse struct {
	VaultSetting
}

// ReplicationVaultSettingClientListResponse contains the response from method ReplicationVaultSettingClient.List.
type ReplicationVaultSettingClientListResponse struct {
	VaultSettingCollection
}

// ReplicationvCentersClientCreateResponse contains the response from method ReplicationvCentersClient.Create.
type ReplicationvCentersClientCreateResponse struct {
	VCenter
}

// ReplicationvCentersClientDeleteResponse contains the response from method ReplicationvCentersClient.Delete.
type ReplicationvCentersClientDeleteResponse struct {
	// placeholder for future response values
}

// ReplicationvCentersClientGetResponse contains the response from method ReplicationvCentersClient.Get.
type ReplicationvCentersClientGetResponse struct {
	VCenter
}

// ReplicationvCentersClientListByReplicationFabricsResponse contains the response from method ReplicationvCentersClient.ListByReplicationFabrics.
type ReplicationvCentersClientListByReplicationFabricsResponse struct {
	VCenterCollection
}

// ReplicationvCentersClientListResponse contains the response from method ReplicationvCentersClient.List.
type ReplicationvCentersClientListResponse struct {
	VCenterCollection
}

// ReplicationvCentersClientUpdateResponse contains the response from method ReplicationvCentersClient.Update.
type ReplicationvCentersClientUpdateResponse struct {
	VCenter
}

// SupportedOperatingSystemsClientGetResponse contains the response from method SupportedOperatingSystemsClient.Get.
type SupportedOperatingSystemsClientGetResponse struct {
	SupportedOperatingSystems
}

// TargetComputeSizesClientListByReplicationProtectedItemsResponse contains the response from method TargetComputeSizesClient.ListByReplicationProtectedItems.
type TargetComputeSizesClientListByReplicationProtectedItemsResponse struct {
	TargetComputeSizeCollection
}