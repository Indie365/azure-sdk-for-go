//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridkubernetes

import "time"

// ConnectedCluster - Represents a connected cluster.
type ConnectedCluster struct {
	// REQUIRED; The identity of the connected cluster.
	Identity *ConnectedClusterIdentity

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Describes the connected cluster resource properties.
	Properties *ConnectedClusterProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ConnectedClusterClientBeginCreateOptions contains the optional parameters for the ConnectedClusterClient.BeginCreate method.
type ConnectedClusterClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConnectedClusterClientBeginDeleteOptions contains the optional parameters for the ConnectedClusterClient.BeginDelete method.
type ConnectedClusterClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConnectedClusterClientGetOptions contains the optional parameters for the ConnectedClusterClient.Get method.
type ConnectedClusterClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterClientListByResourceGroupOptions contains the optional parameters for the ConnectedClusterClient.NewListByResourceGroupPager
// method.
type ConnectedClusterClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterClientListBySubscriptionOptions contains the optional parameters for the ConnectedClusterClient.NewListBySubscriptionPager
// method.
type ConnectedClusterClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterClientListClusterUserCredentialOptions contains the optional parameters for the ConnectedClusterClient.ListClusterUserCredential
// method.
type ConnectedClusterClientListClusterUserCredentialOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterClientUpdateOptions contains the optional parameters for the ConnectedClusterClient.Update method.
type ConnectedClusterClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ConnectedClusterIdentity - Identity for the connected cluster.
type ConnectedClusterIdentity struct {
	// REQUIRED; The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity.
	// The type 'None' means no identity is assigned to the connected cluster.
	Type *ResourceIdentityType

	// READ-ONLY; The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
	PrincipalID *string

	// READ-ONLY; The tenant id associated with the connected cluster. This property will only be provided for a system assigned
	// identity.
	TenantID *string
}

// ConnectedClusterList - The paginated list of connected Clusters
type ConnectedClusterList struct {
	// The link to fetch the next page of connected cluster
	NextLink *string

	// The list of connected clusters
	Value []*ConnectedCluster
}

// ConnectedClusterPatch - Object containing updates for patch operations.
type ConnectedClusterPatch struct {
	// Describes the connected cluster resource properties that can be updated during PATCH operation.
	Properties any

	// Resource tags.
	Tags map[string]*string
}

// ConnectedClusterProperties - Properties of the connected cluster.
type ConnectedClusterProperties struct {
	// REQUIRED; Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate *string

	// The Kubernetes distribution running on this connected cluster.
	Distribution *string

	// The infrastructure on which the Kubernetes cluster represented by this connected cluster is running on.
	Infrastructure *string

	// Provisioning state of the connected cluster resource.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Version of the agent running on the connected cluster resource
	AgentVersion *string

	// READ-ONLY; Represents the connectivity status of the connected cluster.
	ConnectivityStatus *ConnectivityStatus

	// READ-ONLY; The Kubernetes version of the connected cluster resource
	KubernetesVersion *string

	// READ-ONLY; Time representing the last instance when heart beat was received from the cluster
	LastConnectivityTime *time.Time

	// READ-ONLY; Expiration time of the managed identity certificate
	ManagedIdentityCertificateExpirationTime *time.Time

	// READ-ONLY; Connected cluster offering
	Offering *string

	// READ-ONLY; Number of CPU cores present in the connected cluster resource
	TotalCoreCount *int32

	// READ-ONLY; Number of nodes present in the connected cluster resource
	TotalNodeCount *int32
}

// CredentialResult - The credential result response.
type CredentialResult struct {
	// READ-ONLY; The name of the credential.
	Name *string

	// READ-ONLY; Base64-encoded Kubernetes configuration file.
	Value []byte
}

// CredentialResults - The list of credential result response.
type CredentialResults struct {
	// READ-ONLY; Contains the REP (rendezvous endpoint) and “Sender” access token.
	HybridConnectionConfig *HybridConnectionConfig

	// READ-ONLY; Base64-encoded Kubernetes configuration file.
	Kubeconfigs []*CredentialResult
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// HybridConnectionConfig - Contains the REP (rendezvous endpoint) and “Sender” access token.
type HybridConnectionConfig struct {
	// READ-ONLY; Timestamp when this token will be expired.
	ExpirationTime *int64

	// READ-ONLY; Name of the connection
	HybridConnectionName *string

	// READ-ONLY; Name of the relay.
	Relay *string

	// READ-ONLY; Sender access token
	Token *string
}

type ListClusterUserCredentialProperties struct {
	// REQUIRED; The mode of client authentication.
	AuthenticationMethod *AuthenticationMethod

	// REQUIRED; Boolean value to indicate whether the request is for client side proxy or not
	ClientProxy *bool
}

// Operation - The Connected cluster API operation
type Operation struct {
	// READ-ONLY; The object that represents the operation.
	Display *OperationDisplay

	// READ-ONLY; Operation name: {Microsoft.Kubernetes}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string

	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.connectedClusters
	Provider *string

	// Connected Cluster Resource on which the operation is performed
	Resource *string
}

// OperationList - The paginated list of connected cluster API operations.
type OperationList struct {
	// The link to fetch the next page of connected cluster API operations.
	NextLink *string

	// READ-ONLY; The list of connected cluster API operations.
	Value []*Operation
}

// OperationsClientGetOptions contains the optional parameters for the OperationsClient.NewGetPager method.
type OperationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource modification (UTC).
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *LastModifiedByType
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}
