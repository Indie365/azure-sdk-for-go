//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

// AgentPoolsClientCreateOrUpdateResponse contains the response from method AgentPoolsClient.CreateOrUpdate.
type AgentPoolsClientCreateOrUpdateResponse struct {
	AgentPool
}

// AgentPoolsClientDeleteResponse contains the response from method AgentPoolsClient.Delete.
type AgentPoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentPoolsClientGetAvailableAgentPoolVersionsResponse contains the response from method AgentPoolsClient.GetAvailableAgentPoolVersions.
type AgentPoolsClientGetAvailableAgentPoolVersionsResponse struct {
	AgentPoolAvailableVersions
}

// AgentPoolsClientGetResponse contains the response from method AgentPoolsClient.Get.
type AgentPoolsClientGetResponse struct {
	AgentPool
}

// AgentPoolsClientGetUpgradeProfileResponse contains the response from method AgentPoolsClient.GetUpgradeProfile.
type AgentPoolsClientGetUpgradeProfileResponse struct {
	AgentPoolUpgradeProfile
}

// AgentPoolsClientListResponse contains the response from method AgentPoolsClient.List.
type AgentPoolsClientListResponse struct {
	AgentPoolListResult
}

// AgentPoolsClientUpgradeNodeImageVersionResponse contains the response from method AgentPoolsClient.UpgradeNodeImageVersion.
type AgentPoolsClientUpgradeNodeImageVersionResponse struct {
	AgentPool
}

// MaintenanceConfigurationsClientCreateOrUpdateResponse contains the response from method MaintenanceConfigurationsClient.CreateOrUpdate.
type MaintenanceConfigurationsClientCreateOrUpdateResponse struct {
	MaintenanceConfiguration
}

// MaintenanceConfigurationsClientDeleteResponse contains the response from method MaintenanceConfigurationsClient.Delete.
type MaintenanceConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// MaintenanceConfigurationsClientGetResponse contains the response from method MaintenanceConfigurationsClient.Get.
type MaintenanceConfigurationsClientGetResponse struct {
	MaintenanceConfiguration
}

// MaintenanceConfigurationsClientListByManagedClusterResponse contains the response from method MaintenanceConfigurationsClient.ListByManagedCluster.
type MaintenanceConfigurationsClientListByManagedClusterResponse struct {
	MaintenanceConfigurationListResult
}

// ManagedClustersClientCreateOrUpdateResponse contains the response from method ManagedClustersClient.CreateOrUpdate.
type ManagedClustersClientCreateOrUpdateResponse struct {
	ManagedCluster
}

// ManagedClustersClientDeleteResponse contains the response from method ManagedClustersClient.Delete.
type ManagedClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientGetAccessProfileResponse contains the response from method ManagedClustersClient.GetAccessProfile.
type ManagedClustersClientGetAccessProfileResponse struct {
	ManagedClusterAccessProfile
}

// ManagedClustersClientGetCommandResultResponse contains the response from method ManagedClustersClient.GetCommandResult.
type ManagedClustersClientGetCommandResultResponse struct {
	RunCommandResult
}

// ManagedClustersClientGetOSOptionsResponse contains the response from method ManagedClustersClient.GetOSOptions.
type ManagedClustersClientGetOSOptionsResponse struct {
	OSOptionProfile
}

// ManagedClustersClientGetResponse contains the response from method ManagedClustersClient.Get.
type ManagedClustersClientGetResponse struct {
	ManagedCluster
}

// ManagedClustersClientGetUpgradeProfileResponse contains the response from method ManagedClustersClient.GetUpgradeProfile.
type ManagedClustersClientGetUpgradeProfileResponse struct {
	ManagedClusterUpgradeProfile
}

// ManagedClustersClientListByResourceGroupResponse contains the response from method ManagedClustersClient.ListByResourceGroup.
type ManagedClustersClientListByResourceGroupResponse struct {
	ManagedClusterListResult
}

// ManagedClustersClientListClusterAdminCredentialsResponse contains the response from method ManagedClustersClient.ListClusterAdminCredentials.
type ManagedClustersClientListClusterAdminCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListClusterMonitoringUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterMonitoringUserCredentials.
type ManagedClustersClientListClusterMonitoringUserCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListClusterUserCredentialsResponse contains the response from method ManagedClustersClient.ListClusterUserCredentials.
type ManagedClustersClientListClusterUserCredentialsResponse struct {
	CredentialResults
}

// ManagedClustersClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method ManagedClustersClient.ListOutboundNetworkDependenciesEndpoints.
type ManagedClustersClientListOutboundNetworkDependenciesEndpointsResponse struct {
	OutboundEnvironmentEndpointCollection
}

// ManagedClustersClientListResponse contains the response from method ManagedClustersClient.List.
type ManagedClustersClientListResponse struct {
	ManagedClusterListResult
}

// ManagedClustersClientResetAADProfileResponse contains the response from method ManagedClustersClient.ResetAADProfile.
type ManagedClustersClientResetAADProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientResetServicePrincipalProfileResponse contains the response from method ManagedClustersClient.ResetServicePrincipalProfile.
type ManagedClustersClientResetServicePrincipalProfileResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRotateClusterCertificatesResponse contains the response from method ManagedClustersClient.RotateClusterCertificates.
type ManagedClustersClientRotateClusterCertificatesResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientRunCommandResponse contains the response from method ManagedClustersClient.RunCommand.
type ManagedClustersClientRunCommandResponse struct {
	RunCommandResult
}

// ManagedClustersClientStartResponse contains the response from method ManagedClustersClient.Start.
type ManagedClustersClientStartResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientStopResponse contains the response from method ManagedClustersClient.Stop.
type ManagedClustersClientStopResponse struct {
	// placeholder for future response values
}

// ManagedClustersClientUpdateTagsResponse contains the response from method ManagedClustersClient.UpdateTags.
type ManagedClustersClientUpdateTagsResponse struct {
	ManagedCluster
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.List.
type PrivateEndpointConnectionsClientListResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesListResult
}

// ResolvePrivateLinkServiceIDClientPOSTResponse contains the response from method ResolvePrivateLinkServiceIDClient.POST.
type ResolvePrivateLinkServiceIDClientPOSTResponse struct {
	PrivateLinkResource
}

// SnapshotsClientCreateOrUpdateResponse contains the response from method SnapshotsClient.CreateOrUpdate.
type SnapshotsClientCreateOrUpdateResponse struct {
	Snapshot
}

// SnapshotsClientDeleteResponse contains the response from method SnapshotsClient.Delete.
type SnapshotsClientDeleteResponse struct {
	// placeholder for future response values
}

// SnapshotsClientGetResponse contains the response from method SnapshotsClient.Get.
type SnapshotsClientGetResponse struct {
	Snapshot
}

// SnapshotsClientListByResourceGroupResponse contains the response from method SnapshotsClient.ListByResourceGroup.
type SnapshotsClientListByResourceGroupResponse struct {
	SnapshotListResult
}

// SnapshotsClientListResponse contains the response from method SnapshotsClient.List.
type SnapshotsClientListResponse struct {
	SnapshotListResult
}

// SnapshotsClientUpdateTagsResponse contains the response from method SnapshotsClient.UpdateTags.
type SnapshotsClientUpdateTagsResponse struct {
	Snapshot
}
