//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearningservices

// ComputeClientCreateOrUpdateResponse contains the response from method ComputeClient.CreateOrUpdate.
type ComputeClientCreateOrUpdateResponse struct {
	ComputeResource
}

// ComputeClientDeleteResponse contains the response from method ComputeClient.Delete.
type ComputeClientDeleteResponse struct {
	// placeholder for future response values
}

// ComputeClientGetResponse contains the response from method ComputeClient.Get.
type ComputeClientGetResponse struct {
	ComputeResource
}

// ComputeClientListKeysResponse contains the response from method ComputeClient.ListKeys.
type ComputeClientListKeysResponse struct {
	ComputeSecretsClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ComputeClientListKeysResponse.
func (c *ComputeClientListKeysResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalComputeSecretsClassification(data)
	if err != nil {
		return err
	}
	c.ComputeSecretsClassification = res
	return nil
}

// ComputeClientListNodesResponse contains the response from method ComputeClient.ListNodes.
type ComputeClientListNodesResponse struct {
	AmlComputeNodesInformation
}

// ComputeClientListResponse contains the response from method ComputeClient.List.
type ComputeClientListResponse struct {
	PaginatedComputeResourcesList
}

// ComputeClientRestartResponse contains the response from method ComputeClient.Restart.
type ComputeClientRestartResponse struct {
	// placeholder for future response values
}

// ComputeClientStartResponse contains the response from method ComputeClient.Start.
type ComputeClientStartResponse struct {
	// placeholder for future response values
}

// ComputeClientStopResponse contains the response from method ComputeClient.Stop.
type ComputeClientStopResponse struct {
	// placeholder for future response values
}

// ComputeClientUpdateResponse contains the response from method ComputeClient.Update.
type ComputeClientUpdateResponse struct {
	ComputeResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
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
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientListResponse contains the response from method PrivateLinkResourcesClient.List.
type PrivateLinkResourcesClientListResponse struct {
	PrivateLinkResourceListResult
}

// QuotasClientListResponse contains the response from method QuotasClient.List.
type QuotasClientListResponse struct {
	ListWorkspaceQuotas
}

// QuotasClientUpdateResponse contains the response from method QuotasClient.Update.
type QuotasClientUpdateResponse struct {
	UpdateWorkspaceQuotasResult
}

// UsagesClientListResponse contains the response from method UsagesClient.List.
type UsagesClientListResponse struct {
	ListUsagesResult
}

// VirtualMachineSizesClientListResponse contains the response from method VirtualMachineSizesClient.List.
type VirtualMachineSizesClientListResponse struct {
	VirtualMachineSizeListResult
}

// WorkspaceConnectionsClientCreateResponse contains the response from method WorkspaceConnectionsClient.Create.
type WorkspaceConnectionsClientCreateResponse struct {
	WorkspaceConnection
}

// WorkspaceConnectionsClientDeleteResponse contains the response from method WorkspaceConnectionsClient.Delete.
type WorkspaceConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspaceConnectionsClientGetResponse contains the response from method WorkspaceConnectionsClient.Get.
type WorkspaceConnectionsClientGetResponse struct {
	WorkspaceConnection
}

// WorkspaceConnectionsClientListResponse contains the response from method WorkspaceConnectionsClient.List.
type WorkspaceConnectionsClientListResponse struct {
	PaginatedWorkspaceConnectionsList
}

// WorkspaceFeaturesClientListResponse contains the response from method WorkspaceFeaturesClient.List.
type WorkspaceFeaturesClientListResponse struct {
	ListAmlUserFeatureResult
}

// WorkspaceSKUsClientListResponse contains the response from method WorkspaceSKUsClient.List.
type WorkspaceSKUsClientListResponse struct {
	SKUListResult
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.CreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.Delete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientDiagnoseResponse contains the response from method WorkspacesClient.Diagnose.
type WorkspacesClientDiagnoseResponse struct {
	DiagnoseResponseResult
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

// WorkspacesClientListKeysResponse contains the response from method WorkspacesClient.ListKeys.
type WorkspacesClientListKeysResponse struct {
	ListWorkspaceKeysResult
}

// WorkspacesClientListNotebookAccessTokenResponse contains the response from method WorkspacesClient.ListNotebookAccessToken.
type WorkspacesClientListNotebookAccessTokenResponse struct {
	NotebookAccessTokenResult
}

// WorkspacesClientListNotebookKeysResponse contains the response from method WorkspacesClient.ListNotebookKeys.
type WorkspacesClientListNotebookKeysResponse struct {
	ListNotebookKeysResult
}

// WorkspacesClientListOutboundNetworkDependenciesEndpointsResponse contains the response from method WorkspacesClient.ListOutboundNetworkDependenciesEndpoints.
type WorkspacesClientListOutboundNetworkDependenciesEndpointsResponse struct {
	ExternalFQDNResponse
}

// WorkspacesClientListStorageAccountKeysResponse contains the response from method WorkspacesClient.ListStorageAccountKeys.
type WorkspacesClientListStorageAccountKeysResponse struct {
	ListStorageAccountKeysResult
}

// WorkspacesClientPrepareNotebookResponse contains the response from method WorkspacesClient.PrepareNotebook.
type WorkspacesClientPrepareNotebookResponse struct {
	NotebookResourceInfo
}

// WorkspacesClientResyncKeysResponse contains the response from method WorkspacesClient.ResyncKeys.
type WorkspacesClientResyncKeysResponse struct {
	// placeholder for future response values
}

// WorkspacesClientUpdateResponse contains the response from method WorkspacesClient.Update.
type WorkspacesClientUpdateResponse struct {
	Workspace
}
