package containerserviceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-10-01/containerservice"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerservice.OperationListResult, err error)
}

var _ OperationsClientAPI = (*containerservice.OperationsClient)(nil)

// ManagedClustersClientAPI contains the set of methods on the ManagedClustersClient type.
type ManagedClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedCluster) (result containerservice.ManagedClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedCluster, err error)
	GetAccessProfile(ctx context.Context, resourceGroupName string, resourceName string, roleName string) (result containerservice.ManagedClusterAccessProfile, err error)
	GetCommandResult(ctx context.Context, resourceGroupName string, resourceName string, commandID string) (result containerservice.RunCommandResult, err error)
	GetOSOptions(ctx context.Context, location string, resourceType string) (result containerservice.OSOptionProfile, err error)
	GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClusterUpgradeProfile, err error)
	List(ctx context.Context) (result containerservice.ManagedClusterListResultPage, err error)
	ListComplete(ctx context.Context) (result containerservice.ManagedClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.ManagedClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerservice.ManagedClusterListResultIterator, err error)
	ListClusterAdminCredentials(ctx context.Context, resourceGroupName string, resourceName string, serverFqdn string) (result containerservice.CredentialResults, err error)
	ListClusterMonitoringUserCredentials(ctx context.Context, resourceGroupName string, resourceName string, serverFqdn string) (result containerservice.CredentialResults, err error)
	ListClusterUserCredentials(ctx context.Context, resourceGroupName string, resourceName string, serverFqdn string) (result containerservice.CredentialResults, err error)
	ListOutboundNetworkDependenciesEndpoints(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.OutboundEnvironmentEndpointCollectionPage, err error)
	ListOutboundNetworkDependenciesEndpointsComplete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.OutboundEnvironmentEndpointCollectionIterator, err error)
	ResetAADProfile(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedClusterAADProfile) (result containerservice.ManagedClustersResetAADProfileFuture, err error)
	ResetServicePrincipalProfile(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.ManagedClusterServicePrincipalProfile) (result containerservice.ManagedClustersResetServicePrincipalProfileFuture, err error)
	RotateClusterCertificates(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClustersRotateClusterCertificatesFuture, err error)
	RunCommand(ctx context.Context, resourceGroupName string, resourceName string, requestPayload containerservice.RunCommandRequest) (result containerservice.ManagedClustersRunCommandFuture, err error)
	Start(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClustersStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.ManagedClustersStopFuture, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.TagsObject) (result containerservice.ManagedClustersUpdateTagsFuture, err error)
}

var _ ManagedClustersClientAPI = (*containerservice.ManagedClustersClient)(nil)

// MaintenanceConfigurationsClientAPI contains the set of methods on the MaintenanceConfigurationsClient type.
type MaintenanceConfigurationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, configName string, parameters containerservice.MaintenanceConfiguration) (result containerservice.MaintenanceConfiguration, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, configName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, configName string) (result containerservice.MaintenanceConfiguration, err error)
	ListByManagedCluster(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.MaintenanceConfigurationListResultPage, err error)
	ListByManagedClusterComplete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.MaintenanceConfigurationListResultIterator, err error)
}

var _ MaintenanceConfigurationsClientAPI = (*containerservice.MaintenanceConfigurationsClient)(nil)

// AgentPoolsClientAPI contains the set of methods on the AgentPoolsClient type.
type AgentPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, parameters containerservice.AgentPool) (result containerservice.AgentPoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result containerservice.AgentPoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result containerservice.AgentPool, err error)
	GetAvailableAgentPoolVersions(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.AgentPoolAvailableVersions, err error)
	GetUpgradeProfile(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result containerservice.AgentPoolUpgradeProfile, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.AgentPoolListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.AgentPoolListResultIterator, err error)
	UpgradeNodeImageVersion(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string) (result containerservice.AgentPoolsUpgradeNodeImageVersionFuture, err error)
}

var _ AgentPoolsClientAPI = (*containerservice.AgentPoolsClient)(nil)

// PrivateEndpointConnectionsClientAPI contains the set of methods on the PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result containerservice.PrivateEndpointConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string) (result containerservice.PrivateEndpointConnection, err error)
	List(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.PrivateEndpointConnectionListResult, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, privateEndpointConnectionName string, parameters containerservice.PrivateEndpointConnection) (result containerservice.PrivateEndpointConnection, err error)
}

var _ PrivateEndpointConnectionsClientAPI = (*containerservice.PrivateEndpointConnectionsClient)(nil)

// PrivateLinkResourcesClientAPI contains the set of methods on the PrivateLinkResourcesClient type.
type PrivateLinkResourcesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.PrivateLinkResourcesListResult, err error)
}

var _ PrivateLinkResourcesClientAPI = (*containerservice.PrivateLinkResourcesClient)(nil)

// ResolvePrivateLinkServiceIDClientAPI contains the set of methods on the ResolvePrivateLinkServiceIDClient type.
type ResolvePrivateLinkServiceIDClientAPI interface {
	POST(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.PrivateLinkResource) (result containerservice.PrivateLinkResource, err error)
}

var _ ResolvePrivateLinkServiceIDClientAPI = (*containerservice.ResolvePrivateLinkServiceIDClient)(nil)

// SnapshotsClientAPI contains the set of methods on the SnapshotsClient type.
type SnapshotsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.Snapshot) (result containerservice.Snapshot, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result containerservice.Snapshot, err error)
	List(ctx context.Context) (result containerservice.SnapshotListResultPage, err error)
	ListComplete(ctx context.Context) (result containerservice.SnapshotListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerservice.SnapshotListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerservice.SnapshotListResultIterator, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.TagsObject) (result containerservice.Snapshot, err error)
}

var _ SnapshotsClientAPI = (*containerservice.SnapshotsClient)(nil)
