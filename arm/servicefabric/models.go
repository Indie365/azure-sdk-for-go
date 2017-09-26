package servicefabric

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ClusterState enumerates the values for cluster state.
type ClusterState string

const (
	// AutoScale specifies the auto scale state for cluster state.
	AutoScale ClusterState = "AutoScale"
	// BaselineUpgrade specifies the baseline upgrade state for cluster state.
	BaselineUpgrade ClusterState = "BaselineUpgrade"
	// Deploying specifies the deploying state for cluster state.
	Deploying ClusterState = "Deploying"
	// EnforcingClusterVersion specifies the enforcing cluster version state for cluster state.
	EnforcingClusterVersion ClusterState = "EnforcingClusterVersion"
	// Ready specifies the ready state for cluster state.
	Ready ClusterState = "Ready"
	// UpdatingInfrastructure specifies the updating infrastructure state for cluster state.
	UpdatingInfrastructure ClusterState = "UpdatingInfrastructure"
	// UpdatingUserCertificate specifies the updating user certificate state for cluster state.
	UpdatingUserCertificate ClusterState = "UpdatingUserCertificate"
	// UpdatingUserConfiguration specifies the updating user configuration state for cluster state.
	UpdatingUserConfiguration ClusterState = "UpdatingUserConfiguration"
	// UpgradeServiceUnreachable specifies the upgrade service unreachable state for cluster state.
	UpgradeServiceUnreachable ClusterState = "UpgradeServiceUnreachable"
	// WaitingForNodes specifies the waiting for nodes state for cluster state.
	WaitingForNodes ClusterState = "WaitingForNodes"
)

// DurabilityLevel enumerates the values for durability level.
type DurabilityLevel string

const (
	// Bronze specifies the bronze state for durability level.
	Bronze DurabilityLevel = "Bronze"
	// Gold specifies the gold state for durability level.
	Gold DurabilityLevel = "Gold"
	// Silver specifies the silver state for durability level.
	Silver DurabilityLevel = "Silver"
)

// Environment enumerates the values for environment.
type Environment string

const (
	// Linux specifies the linux state for environment.
	Linux Environment = "Linux"
	// Windows specifies the windows state for environment.
	Windows Environment = "Windows"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled specifies the canceled state for provisioning state.
	Canceled ProvisioningState = "Canceled"
	// Failed specifies the failed state for provisioning state.
	Failed ProvisioningState = "Failed"
	// Succeeded specifies the succeeded state for provisioning state.
	Succeeded ProvisioningState = "Succeeded"
	// Updating specifies the updating state for provisioning state.
	Updating ProvisioningState = "Updating"
)

// ReliabilityLevel enumerates the values for reliability level.
type ReliabilityLevel string

const (
	// ReliabilityLevelBronze specifies the reliability level bronze state for reliability level.
	ReliabilityLevelBronze ReliabilityLevel = "Bronze"
	// ReliabilityLevelGold specifies the reliability level gold state for reliability level.
	ReliabilityLevelGold ReliabilityLevel = "Gold"
	// ReliabilityLevelSilver specifies the reliability level silver state for reliability level.
	ReliabilityLevelSilver ReliabilityLevel = "Silver"
)

// ReliabilityLevel1 enumerates the values for reliability level 1.
type ReliabilityLevel1 string

const (
	// ReliabilityLevel1Bronze specifies the reliability level 1 bronze state for reliability level 1.
	ReliabilityLevel1Bronze ReliabilityLevel1 = "Bronze"
	// ReliabilityLevel1Gold specifies the reliability level 1 gold state for reliability level 1.
	ReliabilityLevel1Gold ReliabilityLevel1 = "Gold"
	// ReliabilityLevel1Platinum specifies the reliability level 1 platinum state for reliability level 1.
	ReliabilityLevel1Platinum ReliabilityLevel1 = "Platinum"
	// ReliabilityLevel1Silver specifies the reliability level 1 silver state for reliability level 1.
	ReliabilityLevel1Silver ReliabilityLevel1 = "Silver"
)

// UpgradeMode enumerates the values for upgrade mode.
type UpgradeMode string

const (
	// Automatic specifies the automatic state for upgrade mode.
	Automatic UpgradeMode = "Automatic"
	// Manual specifies the manual state for upgrade mode.
	Manual UpgradeMode = "Manual"
)

// UpgradeMode1 enumerates the values for upgrade mode 1.
type UpgradeMode1 string

const (
	// UpgradeMode1Automatic specifies the upgrade mode 1 automatic state for upgrade mode 1.
	UpgradeMode1Automatic UpgradeMode1 = "Automatic"
	// UpgradeMode1Manual specifies the upgrade mode 1 manual state for upgrade mode 1.
	UpgradeMode1Manual UpgradeMode1 = "Manual"
)

// X509StoreName enumerates the values for x509 store name.
type X509StoreName string

const (
	// AddressBook specifies the address book state for x509 store name.
	AddressBook X509StoreName = "AddressBook"
	// AuthRoot specifies the auth root state for x509 store name.
	AuthRoot X509StoreName = "AuthRoot"
	// CertificateAuthority specifies the certificate authority state for x509 store name.
	CertificateAuthority X509StoreName = "CertificateAuthority"
	// Disallowed specifies the disallowed state for x509 store name.
	Disallowed X509StoreName = "Disallowed"
	// My specifies the my state for x509 store name.
	My X509StoreName = "My"
	// Root specifies the root state for x509 store name.
	Root X509StoreName = "Root"
	// TrustedPeople specifies the trusted people state for x509 store name.
	TrustedPeople X509StoreName = "TrustedPeople"
	// TrustedPublisher specifies the trusted publisher state for x509 store name.
	TrustedPublisher X509StoreName = "TrustedPublisher"
)

// AvailableOperationDisplay is operation supported by ServiceFabric resource provider
type AvailableOperationDisplay struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// AzureActiveDirectory is the settings to enable AAD authentication on the cluster
type AzureActiveDirectory struct {
	TenantID           *string `json:"tenantId,omitempty"`
	ClusterApplication *string `json:"clusterApplication,omitempty"`
	ClientApplication  *string `json:"clientApplication,omitempty"`
}

// CertificateDescription is certificate details
type CertificateDescription struct {
	Thumbprint          *string       `json:"thumbprint,omitempty"`
	ThumbprintSecondary *string       `json:"thumbprintSecondary,omitempty"`
	X509StoreName       X509StoreName `json:"x509StoreName,omitempty"`
}

// ClientCertificateCommonName is client certificate details using common name
type ClientCertificateCommonName struct {
	IsAdmin                     *bool   `json:"isAdmin,omitempty"`
	CertificateCommonName       *string `json:"certificateCommonName,omitempty"`
	CertificateIssuerThumbprint *string `json:"certificateIssuerThumbprint,omitempty"`
}

// ClientCertificateThumbprint is client certificate details using thumbprint
type ClientCertificateThumbprint struct {
	IsAdmin               *bool   `json:"isAdmin,omitempty"`
	CertificateThumbprint *string `json:"certificateThumbprint,omitempty"`
}

// Cluster is the cluster resource
type Cluster struct {
	autorest.Response  `json:"-"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Type               *string             `json:"type,omitempty"`
	Location           *string             `json:"location,omitempty"`
	Tags               *map[string]*string `json:"tags,omitempty"`
	*ClusterProperties `json:"properties,omitempty"`
}

// ClusterCodeVersionsListResult is the list results of the ServiceFabric runtime versions
type ClusterCodeVersionsListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ClusterCodeVersionsResult `json:"value,omitempty"`
	NextLink          *string                      `json:"nextLink,omitempty"`
}

// ClusterCodeVersionsListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ClusterCodeVersionsListResult) ClusterCodeVersionsListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ClusterCodeVersionsResult is the result of the ServiceFabric runtime versions
type ClusterCodeVersionsResult struct {
	autorest.Response      `json:"-"`
	ID                     *string `json:"id,omitempty"`
	Name                   *string `json:"name,omitempty"`
	Type                   *string `json:"type,omitempty"`
	*ClusterVersionDetails `json:"properties,omitempty"`
}

// ClusterHealthPolicy is defines a health policy used to evaluate the health of the cluster or of a cluster node.
type ClusterHealthPolicy struct {
	MaxPercentUnhealthyNodes        *int32 `json:"maxPercentUnhealthyNodes,omitempty"`
	MaxPercentUnhealthyApplications *int32 `json:"maxPercentUnhealthyApplications,omitempty"`
}

// ClusterListResult is cluster list results
type ClusterListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Cluster `json:"value,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// ClusterListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ClusterListResult) ClusterListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ClusterProperties is the cluster resource properties
type ClusterProperties struct {
	AvailableClusterVersions        *[]ClusterVersionDetails         `json:"availableClusterVersions,omitempty"`
	ClusterID                       *string                          `json:"clusterId,omitempty"`
	ClusterState                    ClusterState                     `json:"clusterState,omitempty"`
	ClusterEndpoint                 *string                          `json:"clusterEndpoint,omitempty"`
	ClusterCodeVersion              *string                          `json:"clusterCodeVersion,omitempty"`
	Certificate                     *CertificateDescription          `json:"certificate,omitempty"`
	ReliabilityLevel                ReliabilityLevel                 `json:"reliabilityLevel,omitempty"`
	UpgradeMode                     UpgradeMode                      `json:"upgradeMode,omitempty"`
	ClientCertificateThumbprints    *[]ClientCertificateThumbprint   `json:"clientCertificateThumbprints,omitempty"`
	ClientCertificateCommonNames    *[]ClientCertificateCommonName   `json:"clientCertificateCommonNames,omitempty"`
	FabricSettings                  *[]SettingsSectionDescription    `json:"fabricSettings,omitempty"`
	ReverseProxyCertificate         *CertificateDescription          `json:"reverseProxyCertificate,omitempty"`
	ManagementEndpoint              *string                          `json:"managementEndpoint,omitempty"`
	NodeTypes                       *[]NodeTypeDescription           `json:"nodeTypes,omitempty"`
	AzureActiveDirectory            *AzureActiveDirectory            `json:"azureActiveDirectory,omitempty"`
	ProvisioningState               ProvisioningState                `json:"provisioningState,omitempty"`
	VMImage                         *string                          `json:"vmImage,omitempty"`
	DiagnosticsStorageAccountConfig *DiagnosticsStorageAccountConfig `json:"diagnosticsStorageAccountConfig,omitempty"`
	UpgradeDescription              *ClusterUpgradePolicy            `json:"upgradeDescription,omitempty"`
}

// ClusterPropertiesUpdateParameters is the cluster resource properties can be updated
type ClusterPropertiesUpdateParameters struct {
	ReliabilityLevel             ReliabilityLevel               `json:"reliabilityLevel,omitempty"`
	UpgradeMode                  UpgradeMode                    `json:"upgradeMode,omitempty"`
	ClusterCodeVersion           *string                        `json:"clusterCodeVersion,omitempty"`
	Certificate                  *CertificateDescription        `json:"certificate,omitempty"`
	ClientCertificateThumbprints *[]ClientCertificateThumbprint `json:"clientCertificateThumbprints,omitempty"`
	ClientCertificateCommonNames *[]ClientCertificateCommonName `json:"clientCertificateCommonNames,omitempty"`
	FabricSettings               *[]SettingsSectionDescription  `json:"fabricSettings,omitempty"`
	ReverseProxyCertificate      *CertificateDescription        `json:"reverseProxyCertificate,omitempty"`
	NodeTypes                    *[]NodeTypeDescription         `json:"nodeTypes,omitempty"`
	UpgradeDescription           *ClusterUpgradePolicy          `json:"upgradeDescription,omitempty"`
}

// ClusterUpdateParameters is cluster update request
type ClusterUpdateParameters struct {
	*ClusterPropertiesUpdateParameters `json:"properties,omitempty"`
	Tags                               *map[string]*string `json:"tags,omitempty"`
}

// ClusterUpgradeDeltaHealthPolicy is delta health policy for the cluster
type ClusterUpgradeDeltaHealthPolicy struct {
	MaxPercentDeltaUnhealthyNodes              *int32 `json:"maxPercentDeltaUnhealthyNodes,omitempty"`
	MaxPercentUpgradeDomainDeltaUnhealthyNodes *int32 `json:"maxPercentUpgradeDomainDeltaUnhealthyNodes,omitempty"`
	MaxPercentDeltaUnhealthyApplications       *int32 `json:"maxPercentDeltaUnhealthyApplications,omitempty"`
}

// ClusterUpgradePolicy is cluster upgrade policy
type ClusterUpgradePolicy struct {
	OverrideUserUpgradePolicy     *bool                            `json:"overrideUserUpgradePolicy,omitempty"`
	ForceRestart                  *bool                            `json:"forceRestart,omitempty"`
	UpgradeReplicaSetCheckTimeout *string                          `json:"upgradeReplicaSetCheckTimeout,omitempty"`
	HealthCheckWaitDuration       *string                          `json:"healthCheckWaitDuration,omitempty"`
	HealthCheckStableDuration     *string                          `json:"healthCheckStableDuration,omitempty"`
	HealthCheckRetryTimeout       *string                          `json:"healthCheckRetryTimeout,omitempty"`
	UpgradeTimeout                *string                          `json:"upgradeTimeout,omitempty"`
	UpgradeDomainTimeout          *string                          `json:"upgradeDomainTimeout,omitempty"`
	HealthPolicy                  *ClusterHealthPolicy             `json:"healthPolicy,omitempty"`
	DeltaHealthPolicy             *ClusterUpgradeDeltaHealthPolicy `json:"deltaHealthPolicy,omitempty"`
}

// ClusterVersionDetails is the detail of the ServiceFabric runtime version result
type ClusterVersionDetails struct {
	CodeVersion      *string     `json:"codeVersion,omitempty"`
	SupportExpiryUtc *string     `json:"supportExpiryUtc,omitempty"`
	Environment      Environment `json:"environment,omitempty"`
}

// DiagnosticsStorageAccountConfig is diagnostics storage account config
type DiagnosticsStorageAccountConfig struct {
	StorageAccountName      *string `json:"storageAccountName,omitempty"`
	ProtectedAccountKeyName *string `json:"protectedAccountKeyName,omitempty"`
	BlobEndpoint            *string `json:"blobEndpoint,omitempty"`
	QueueEndpoint           *string `json:"queueEndpoint,omitempty"`
	TableEndpoint           *string `json:"tableEndpoint,omitempty"`
}

// EndpointRangeDescription is port range details
type EndpointRangeDescription struct {
	StartPort *int32 `json:"startPort,omitempty"`
	EndPort   *int32 `json:"endPort,omitempty"`
}

// ErrorModel is the structure of the error
type ErrorModel struct {
	Error *ErrorModelError `json:"error,omitempty"`
}

// ErrorModelError is the error detail
type ErrorModelError struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NodeTypeDescription is describes a node type in the cluster, each node type represents sub set of nodes in the
// cluster
type NodeTypeDescription struct {
	Name                         *string                   `json:"name,omitempty"`
	PlacementProperties          *map[string]*string       `json:"placementProperties,omitempty"`
	Capacities                   *map[string]*string       `json:"capacities,omitempty"`
	ClientConnectionEndpointPort *int32                    `json:"clientConnectionEndpointPort,omitempty"`
	HTTPGatewayEndpointPort      *int32                    `json:"httpGatewayEndpointPort,omitempty"`
	DurabilityLevel              DurabilityLevel           `json:"durabilityLevel,omitempty"`
	ApplicationPorts             *EndpointRangeDescription `json:"applicationPorts,omitempty"`
	EphemeralPorts               *EndpointRangeDescription `json:"ephemeralPorts,omitempty"`
	IsPrimary                    *bool                     `json:"isPrimary,omitempty"`
	VMInstanceCount              *int32                    `json:"vmInstanceCount,omitempty"`
	ReverseProxyEndpointPort     *int32                    `json:"reverseProxyEndpointPort,omitempty"`
}

// OperationListResult is result of the request to list ServiceFabric operations. It contains a list of operations and
// a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]OperationResult `json:"value,omitempty"`
	NextLink          *string            `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// OperationResult is available operation list result
type OperationResult struct {
	Name     *string                    `json:"name,omitempty"`
	Display  *AvailableOperationDisplay `json:"display,omitempty"`
	Origin   *string                    `json:"origin,omitempty"`
	NextLink *string                    `json:"nextLink,omitempty"`
}

// Resource is the resource model definition.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// SettingsParameterDescription is serviceFabric settings under sections
type SettingsParameterDescription struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// SettingsSectionDescription is serviceFabric section settings
type SettingsSectionDescription struct {
	Name       *string                         `json:"name,omitempty"`
	Parameters *[]SettingsParameterDescription `json:"parameters,omitempty"`
}
