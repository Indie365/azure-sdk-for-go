//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabricks

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// OutboundNetworkDependenciesEndpointsClientListResponse contains the response from method OutboundNetworkDependenciesEndpointsClient.List.
type OutboundNetworkDependenciesEndpointsClientListResponse struct {
	// Collection of outbound network dependency endpoints
	OutboundEnvironmentEndpointArray []*OutboundEnvironmentEndpoint
}

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.Create.
type PrivateEndpointConnectionsClientCreateResponse struct {
	PrivateEndpointConnection
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
	PrivateEndpointConnectionsList
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	GroupIDInformation
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourcesList
}

// VNetPeeringClientCreateOrUpdateResponse contains the response from method VNetPeeringClient.CreateOrUpdate.
type VNetPeeringClientCreateOrUpdateResponse struct {
	VirtualNetworkPeering
}

// VNetPeeringClientDeleteResponse contains the response from method VNetPeeringClient.Delete.
type VNetPeeringClientDeleteResponse struct {
	// placeholder for future response values
}

// VNetPeeringClientGetResponse contains the response from method VNetPeeringClient.Get.
type VNetPeeringClientGetResponse struct {
	VirtualNetworkPeering
}

// VNetPeeringClientListByWorkspaceResponse contains the response from method VNetPeeringClient.ListByWorkspace.
type VNetPeeringClientListByWorkspaceResponse struct {
	VirtualNetworkPeeringList
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.CreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.Delete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.ListByResourceGroup.
type WorkspacesClientListByResourceGroupResponse struct {
	WorkspaceListResult
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.ListBySubscription.
type WorkspacesClientListBySubscriptionResponse struct {
	WorkspaceListResult
}

// WorkspacesClientUpdateResponse contains the response from method WorkspacesClient.Update.
type WorkspacesClientUpdateResponse struct {
	Workspace
}