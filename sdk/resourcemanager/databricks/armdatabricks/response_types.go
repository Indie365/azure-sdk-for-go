//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabricks

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Resource Provider operations. It contains a list of operations and a URL link to get the
// next set of results.
	OperationListResult
}

// OutboundNetworkDependenciesEndpointsClientListResponse contains the response from method OutboundNetworkDependenciesEndpointsClient.List.
type OutboundNetworkDependenciesEndpointsClientListResponse struct {
	// Collection of outbound network dependency endpoints
	OutboundEnvironmentEndpointArray []*OutboundEnvironmentEndpoint
}

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreate.
type PrivateEndpointConnectionsClientCreateResponse struct {
	// The private endpoint connection of a workspace
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The private endpoint connection of a workspace
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.NewListPager.
type PrivateEndpointConnectionsClientListResponse struct {
	// List of private link connections.
	PrivateEndpointConnectionsList
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// The group information for creating a private endpoint on a workspace
	GroupIDInformation
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.NewListPager.
type PrivateLinkResourcesClientListResponse struct {
	// The available private link resources for a workspace
	PrivateLinkResourcesList
}

// VNetPeeringClientCreateOrUpdateResponse contains the response from method VNetPeeringClient.BeginCreateOrUpdate.
type VNetPeeringClientCreateOrUpdateResponse struct {
	// Peerings in a VirtualNetwork resource
	VirtualNetworkPeering
}

// VNetPeeringClientDeleteResponse contains the response from method VNetPeeringClient.BeginDelete.
type VNetPeeringClientDeleteResponse struct {
	// placeholder for future response values
}

// VNetPeeringClientGetResponse contains the response from method VNetPeeringClient.Get.
type VNetPeeringClientGetResponse struct {
	// Peerings in a VirtualNetwork resource
	VirtualNetworkPeering
}

// VNetPeeringClientListByWorkspaceResponse contains the response from method VNetPeeringClient.NewListByWorkspacePager.
type VNetPeeringClientListByWorkspaceResponse struct {
	// Gets all virtual network peerings under a workspace.
	VirtualNetworkPeeringList
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.BeginCreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	// Information about workspace.
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.BeginDelete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	// Information about workspace.
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.NewListByResourceGroupPager.
type WorkspacesClientListByResourceGroupResponse struct {
	// List of workspaces.
	WorkspaceListResult
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.NewListBySubscriptionPager.
type WorkspacesClientListBySubscriptionResponse struct {
	// List of workspaces.
	WorkspaceListResult
}

// WorkspacesClientUpdateResponse contains the response from method WorkspacesClient.BeginUpdate.
type WorkspacesClientUpdateResponse struct {
	// Information about workspace.
	Workspace
}

