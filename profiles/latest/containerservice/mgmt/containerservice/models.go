//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package containerservice

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/containerservice/mgmt/2022-07-01/containerservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AgentPoolMode = original.AgentPoolMode

const (
	System AgentPoolMode = original.System
	User   AgentPoolMode = original.User
)

type AgentPoolType = original.AgentPoolType

const (
	AvailabilitySet         AgentPoolType = original.AvailabilitySet
	VirtualMachineScaleSets AgentPoolType = original.VirtualMachineScaleSets
)

type Code = original.Code

const (
	Running Code = original.Running
	Stopped Code = original.Stopped
)

type ConnectionStatus = original.ConnectionStatus

const (
	Approved     ConnectionStatus = original.Approved
	Disconnected ConnectionStatus = original.Disconnected
	Pending      ConnectionStatus = original.Pending
	Rejected     ConnectionStatus = original.Rejected
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type Expander = original.Expander

const (
	LeastWaste Expander = original.LeastWaste
	MostPods   Expander = original.MostPods
	Priority   Expander = original.Priority
	Random     Expander = original.Random
)

type ExtendedLocationTypes = original.ExtendedLocationTypes

const (
	EdgeZone ExtendedLocationTypes = original.EdgeZone
)

type Format = original.Format

const (
	Azure Format = original.Azure
	Exec  Format = original.Exec
)

type GPUInstanceProfile = original.GPUInstanceProfile

const (
	MIG1g GPUInstanceProfile = original.MIG1g
	MIG2g GPUInstanceProfile = original.MIG2g
	MIG3g GPUInstanceProfile = original.MIG3g
	MIG4g GPUInstanceProfile = original.MIG4g
	MIG7g GPUInstanceProfile = original.MIG7g
)

type IPFamily = original.IPFamily

const (
	IPv4 IPFamily = original.IPv4
	IPv6 IPFamily = original.IPv6
)

type KeyVaultNetworkAccessTypes = original.KeyVaultNetworkAccessTypes

const (
	Private KeyVaultNetworkAccessTypes = original.Private
	Public  KeyVaultNetworkAccessTypes = original.Public
)

type KubeletDiskType = original.KubeletDiskType

const (
	OS        KubeletDiskType = original.OS
	Temporary KubeletDiskType = original.Temporary
)

type LicenseType = original.LicenseType

const (
	None          LicenseType = original.None
	WindowsServer LicenseType = original.WindowsServer
)

type LoadBalancerSku = original.LoadBalancerSku

const (
	Basic    LoadBalancerSku = original.Basic
	Standard LoadBalancerSku = original.Standard
)

type ManagedClusterPodIdentityProvisioningState = original.ManagedClusterPodIdentityProvisioningState

const (
	Assigned ManagedClusterPodIdentityProvisioningState = original.Assigned
	Deleting ManagedClusterPodIdentityProvisioningState = original.Deleting
	Failed   ManagedClusterPodIdentityProvisioningState = original.Failed
	Updating ManagedClusterPodIdentityProvisioningState = original.Updating
)

type ManagedClusterSKUName = original.ManagedClusterSKUName

const (
	ManagedClusterSKUNameBasic ManagedClusterSKUName = original.ManagedClusterSKUNameBasic
)

type ManagedClusterSKUTier = original.ManagedClusterSKUTier

const (
	Free ManagedClusterSKUTier = original.Free
	Paid ManagedClusterSKUTier = original.Paid
)

type NetworkMode = original.NetworkMode

const (
	Bridge      NetworkMode = original.Bridge
	Transparent NetworkMode = original.Transparent
)

type NetworkPlugin = original.NetworkPlugin

const (
	NetworkPluginAzure   NetworkPlugin = original.NetworkPluginAzure
	NetworkPluginKubenet NetworkPlugin = original.NetworkPluginKubenet
	NetworkPluginNone    NetworkPlugin = original.NetworkPluginNone
)

type NetworkPolicy = original.NetworkPolicy

const (
	NetworkPolicyAzure  NetworkPolicy = original.NetworkPolicyAzure
	NetworkPolicyCalico NetworkPolicy = original.NetworkPolicyCalico
)

type OSDiskType = original.OSDiskType

const (
	Ephemeral OSDiskType = original.Ephemeral
	Managed   OSDiskType = original.Managed
)

type OSSKU = original.OSSKU

const (
	CBLMariner  OSSKU = original.CBLMariner
	Ubuntu      OSSKU = original.Ubuntu
	Windows2019 OSSKU = original.Windows2019
	Windows2022 OSSKU = original.Windows2022
)

type OSType = original.OSType

const (
	Linux   OSType = original.Linux
	Windows OSType = original.Windows
)

type OutboundType = original.OutboundType

const (
	LoadBalancer           OutboundType = original.LoadBalancer
	ManagedNATGateway      OutboundType = original.ManagedNATGateway
	UserAssignedNATGateway OutboundType = original.UserAssignedNATGateway
	UserDefinedRouting     OutboundType = original.UserDefinedRouting
)

type PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningState

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateCreating
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateDeleting
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateFailed
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = original.PrivateEndpointConnectionProvisioningStateSucceeded
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	Disabled PublicNetworkAccess = original.Disabled
	Enabled  PublicNetworkAccess = original.Enabled
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = original.ResourceIdentityTypeUserAssigned
)

type ScaleDownMode = original.ScaleDownMode

const (
	Deallocate ScaleDownMode = original.Deallocate
	Delete     ScaleDownMode = original.Delete
)

type ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicy

const (
	ScaleSetEvictionPolicyDeallocate ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicyDeallocate
	ScaleSetEvictionPolicyDelete     ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicyDelete
)

type ScaleSetPriority = original.ScaleSetPriority

const (
	Regular ScaleSetPriority = original.Regular
	Spot    ScaleSetPriority = original.Spot
)

type SnapshotType = original.SnapshotType

const (
	NodePool SnapshotType = original.NodePool
)

type StorageProfileTypes = original.StorageProfileTypes

const (
	ManagedDisks   StorageProfileTypes = original.ManagedDisks
	StorageAccount StorageProfileTypes = original.StorageAccount
)

type UpgradeChannel = original.UpgradeChannel

const (
	UpgradeChannelNodeImage UpgradeChannel = original.UpgradeChannelNodeImage
	UpgradeChannelNone      UpgradeChannel = original.UpgradeChannelNone
	UpgradeChannelPatch     UpgradeChannel = original.UpgradeChannelPatch
	UpgradeChannelRapid     UpgradeChannel = original.UpgradeChannelRapid
	UpgradeChannelStable    UpgradeChannel = original.UpgradeChannelStable
)

type VMSizeTypes = original.VMSizeTypes

const (
	StandardA1          VMSizeTypes = original.StandardA1
	StandardA10         VMSizeTypes = original.StandardA10
	StandardA11         VMSizeTypes = original.StandardA11
	StandardA1V2        VMSizeTypes = original.StandardA1V2
	StandardA2          VMSizeTypes = original.StandardA2
	StandardA2mV2       VMSizeTypes = original.StandardA2mV2
	StandardA2V2        VMSizeTypes = original.StandardA2V2
	StandardA3          VMSizeTypes = original.StandardA3
	StandardA4          VMSizeTypes = original.StandardA4
	StandardA4mV2       VMSizeTypes = original.StandardA4mV2
	StandardA4V2        VMSizeTypes = original.StandardA4V2
	StandardA5          VMSizeTypes = original.StandardA5
	StandardA6          VMSizeTypes = original.StandardA6
	StandardA7          VMSizeTypes = original.StandardA7
	StandardA8          VMSizeTypes = original.StandardA8
	StandardA8mV2       VMSizeTypes = original.StandardA8mV2
	StandardA8V2        VMSizeTypes = original.StandardA8V2
	StandardA9          VMSizeTypes = original.StandardA9
	StandardB2ms        VMSizeTypes = original.StandardB2ms
	StandardB2s         VMSizeTypes = original.StandardB2s
	StandardB4ms        VMSizeTypes = original.StandardB4ms
	StandardB8ms        VMSizeTypes = original.StandardB8ms
	StandardD1          VMSizeTypes = original.StandardD1
	StandardD11         VMSizeTypes = original.StandardD11
	StandardD11V2       VMSizeTypes = original.StandardD11V2
	StandardD11V2Promo  VMSizeTypes = original.StandardD11V2Promo
	StandardD12         VMSizeTypes = original.StandardD12
	StandardD12V2       VMSizeTypes = original.StandardD12V2
	StandardD12V2Promo  VMSizeTypes = original.StandardD12V2Promo
	StandardD13         VMSizeTypes = original.StandardD13
	StandardD13V2       VMSizeTypes = original.StandardD13V2
	StandardD13V2Promo  VMSizeTypes = original.StandardD13V2Promo
	StandardD14         VMSizeTypes = original.StandardD14
	StandardD14V2       VMSizeTypes = original.StandardD14V2
	StandardD14V2Promo  VMSizeTypes = original.StandardD14V2Promo
	StandardD15V2       VMSizeTypes = original.StandardD15V2
	StandardD16sV3      VMSizeTypes = original.StandardD16sV3
	StandardD16V3       VMSizeTypes = original.StandardD16V3
	StandardD1V2        VMSizeTypes = original.StandardD1V2
	StandardD2          VMSizeTypes = original.StandardD2
	StandardD2sV3       VMSizeTypes = original.StandardD2sV3
	StandardD2V2        VMSizeTypes = original.StandardD2V2
	StandardD2V2Promo   VMSizeTypes = original.StandardD2V2Promo
	StandardD2V3        VMSizeTypes = original.StandardD2V3
	StandardD3          VMSizeTypes = original.StandardD3
	StandardD32sV3      VMSizeTypes = original.StandardD32sV3
	StandardD32V3       VMSizeTypes = original.StandardD32V3
	StandardD3V2        VMSizeTypes = original.StandardD3V2
	StandardD3V2Promo   VMSizeTypes = original.StandardD3V2Promo
	StandardD4          VMSizeTypes = original.StandardD4
	StandardD4sV3       VMSizeTypes = original.StandardD4sV3
	StandardD4V2        VMSizeTypes = original.StandardD4V2
	StandardD4V2Promo   VMSizeTypes = original.StandardD4V2Promo
	StandardD4V3        VMSizeTypes = original.StandardD4V3
	StandardD5V2        VMSizeTypes = original.StandardD5V2
	StandardD5V2Promo   VMSizeTypes = original.StandardD5V2Promo
	StandardD64sV3      VMSizeTypes = original.StandardD64sV3
	StandardD64V3       VMSizeTypes = original.StandardD64V3
	StandardD8sV3       VMSizeTypes = original.StandardD8sV3
	StandardD8V3        VMSizeTypes = original.StandardD8V3
	StandardDS1         VMSizeTypes = original.StandardDS1
	StandardDS11        VMSizeTypes = original.StandardDS11
	StandardDS11V2      VMSizeTypes = original.StandardDS11V2
	StandardDS11V2Promo VMSizeTypes = original.StandardDS11V2Promo
	StandardDS12        VMSizeTypes = original.StandardDS12
	StandardDS12V2      VMSizeTypes = original.StandardDS12V2
	StandardDS12V2Promo VMSizeTypes = original.StandardDS12V2Promo
	StandardDS13        VMSizeTypes = original.StandardDS13
	StandardDS132V2     VMSizeTypes = original.StandardDS132V2
	StandardDS134V2     VMSizeTypes = original.StandardDS134V2
	StandardDS13V2      VMSizeTypes = original.StandardDS13V2
	StandardDS13V2Promo VMSizeTypes = original.StandardDS13V2Promo
	StandardDS14        VMSizeTypes = original.StandardDS14
	StandardDS144V2     VMSizeTypes = original.StandardDS144V2
	StandardDS148V2     VMSizeTypes = original.StandardDS148V2
	StandardDS14V2      VMSizeTypes = original.StandardDS14V2
	StandardDS14V2Promo VMSizeTypes = original.StandardDS14V2Promo
	StandardDS15V2      VMSizeTypes = original.StandardDS15V2
	StandardDS1V2       VMSizeTypes = original.StandardDS1V2
	StandardDS2         VMSizeTypes = original.StandardDS2
	StandardDS2V2       VMSizeTypes = original.StandardDS2V2
	StandardDS2V2Promo  VMSizeTypes = original.StandardDS2V2Promo
	StandardDS3         VMSizeTypes = original.StandardDS3
	StandardDS3V2       VMSizeTypes = original.StandardDS3V2
	StandardDS3V2Promo  VMSizeTypes = original.StandardDS3V2Promo
	StandardDS4         VMSizeTypes = original.StandardDS4
	StandardDS4V2       VMSizeTypes = original.StandardDS4V2
	StandardDS4V2Promo  VMSizeTypes = original.StandardDS4V2Promo
	StandardDS5V2       VMSizeTypes = original.StandardDS5V2
	StandardDS5V2Promo  VMSizeTypes = original.StandardDS5V2Promo
	StandardE16sV3      VMSizeTypes = original.StandardE16sV3
	StandardE16V3       VMSizeTypes = original.StandardE16V3
	StandardE2sV3       VMSizeTypes = original.StandardE2sV3
	StandardE2V3        VMSizeTypes = original.StandardE2V3
	StandardE3216sV3    VMSizeTypes = original.StandardE3216sV3
	StandardE328sV3     VMSizeTypes = original.StandardE328sV3
	StandardE32sV3      VMSizeTypes = original.StandardE32sV3
	StandardE32V3       VMSizeTypes = original.StandardE32V3
	StandardE4sV3       VMSizeTypes = original.StandardE4sV3
	StandardE4V3        VMSizeTypes = original.StandardE4V3
	StandardE6416sV3    VMSizeTypes = original.StandardE6416sV3
	StandardE6432sV3    VMSizeTypes = original.StandardE6432sV3
	StandardE64sV3      VMSizeTypes = original.StandardE64sV3
	StandardE64V3       VMSizeTypes = original.StandardE64V3
	StandardE8sV3       VMSizeTypes = original.StandardE8sV3
	StandardE8V3        VMSizeTypes = original.StandardE8V3
	StandardF1          VMSizeTypes = original.StandardF1
	StandardF16         VMSizeTypes = original.StandardF16
	StandardF16s        VMSizeTypes = original.StandardF16s
	StandardF16sV2      VMSizeTypes = original.StandardF16sV2
	StandardF1s         VMSizeTypes = original.StandardF1s
	StandardF2          VMSizeTypes = original.StandardF2
	StandardF2s         VMSizeTypes = original.StandardF2s
	StandardF2sV2       VMSizeTypes = original.StandardF2sV2
	StandardF32sV2      VMSizeTypes = original.StandardF32sV2
	StandardF4          VMSizeTypes = original.StandardF4
	StandardF4s         VMSizeTypes = original.StandardF4s
	StandardF4sV2       VMSizeTypes = original.StandardF4sV2
	StandardF64sV2      VMSizeTypes = original.StandardF64sV2
	StandardF72sV2      VMSizeTypes = original.StandardF72sV2
	StandardF8          VMSizeTypes = original.StandardF8
	StandardF8s         VMSizeTypes = original.StandardF8s
	StandardF8sV2       VMSizeTypes = original.StandardF8sV2
	StandardG1          VMSizeTypes = original.StandardG1
	StandardG2          VMSizeTypes = original.StandardG2
	StandardG3          VMSizeTypes = original.StandardG3
	StandardG4          VMSizeTypes = original.StandardG4
	StandardG5          VMSizeTypes = original.StandardG5
	StandardGS1         VMSizeTypes = original.StandardGS1
	StandardGS2         VMSizeTypes = original.StandardGS2
	StandardGS3         VMSizeTypes = original.StandardGS3
	StandardGS4         VMSizeTypes = original.StandardGS4
	StandardGS44        VMSizeTypes = original.StandardGS44
	StandardGS48        VMSizeTypes = original.StandardGS48
	StandardGS5         VMSizeTypes = original.StandardGS5
	StandardGS516       VMSizeTypes = original.StandardGS516
	StandardGS58        VMSizeTypes = original.StandardGS58
	StandardH16         VMSizeTypes = original.StandardH16
	StandardH16m        VMSizeTypes = original.StandardH16m
	StandardH16mr       VMSizeTypes = original.StandardH16mr
	StandardH16r        VMSizeTypes = original.StandardH16r
	StandardH8          VMSizeTypes = original.StandardH8
	StandardH8m         VMSizeTypes = original.StandardH8m
	StandardL16s        VMSizeTypes = original.StandardL16s
	StandardL32s        VMSizeTypes = original.StandardL32s
	StandardL4s         VMSizeTypes = original.StandardL4s
	StandardL8s         VMSizeTypes = original.StandardL8s
	StandardM12832ms    VMSizeTypes = original.StandardM12832ms
	StandardM12864ms    VMSizeTypes = original.StandardM12864ms
	StandardM128ms      VMSizeTypes = original.StandardM128ms
	StandardM128s       VMSizeTypes = original.StandardM128s
	StandardM6416ms     VMSizeTypes = original.StandardM6416ms
	StandardM6432ms     VMSizeTypes = original.StandardM6432ms
	StandardM64ms       VMSizeTypes = original.StandardM64ms
	StandardM64s        VMSizeTypes = original.StandardM64s
	StandardNC12        VMSizeTypes = original.StandardNC12
	StandardNC12sV2     VMSizeTypes = original.StandardNC12sV2
	StandardNC12sV3     VMSizeTypes = original.StandardNC12sV3
	StandardNC24        VMSizeTypes = original.StandardNC24
	StandardNC24r       VMSizeTypes = original.StandardNC24r
	StandardNC24rsV2    VMSizeTypes = original.StandardNC24rsV2
	StandardNC24rsV3    VMSizeTypes = original.StandardNC24rsV3
	StandardNC24sV2     VMSizeTypes = original.StandardNC24sV2
	StandardNC24sV3     VMSizeTypes = original.StandardNC24sV3
	StandardNC6         VMSizeTypes = original.StandardNC6
	StandardNC6sV2      VMSizeTypes = original.StandardNC6sV2
	StandardNC6sV3      VMSizeTypes = original.StandardNC6sV3
	StandardND12s       VMSizeTypes = original.StandardND12s
	StandardND24rs      VMSizeTypes = original.StandardND24rs
	StandardND24s       VMSizeTypes = original.StandardND24s
	StandardND6s        VMSizeTypes = original.StandardND6s
	StandardNV12        VMSizeTypes = original.StandardNV12
	StandardNV24        VMSizeTypes = original.StandardNV24
	StandardNV6         VMSizeTypes = original.StandardNV6
)

type WeekDay = original.WeekDay

const (
	Friday    WeekDay = original.Friday
	Monday    WeekDay = original.Monday
	Saturday  WeekDay = original.Saturday
	Sunday    WeekDay = original.Sunday
	Thursday  WeekDay = original.Thursday
	Tuesday   WeekDay = original.Tuesday
	Wednesday WeekDay = original.Wednesday
)

type WorkloadRuntime = original.WorkloadRuntime

const (
	OCIContainer WorkloadRuntime = original.OCIContainer
	WasmWasi     WorkloadRuntime = original.WasmWasi
)

type AccessProfile = original.AccessProfile
type AgentPool = original.AgentPool
type AgentPoolAvailableVersions = original.AgentPoolAvailableVersions
type AgentPoolAvailableVersionsProperties = original.AgentPoolAvailableVersionsProperties
type AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem = original.AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem
type AgentPoolListResult = original.AgentPoolListResult
type AgentPoolListResultIterator = original.AgentPoolListResultIterator
type AgentPoolListResultPage = original.AgentPoolListResultPage
type AgentPoolUpgradeProfile = original.AgentPoolUpgradeProfile
type AgentPoolUpgradeProfileProperties = original.AgentPoolUpgradeProfileProperties
type AgentPoolUpgradeProfilePropertiesUpgradesItem = original.AgentPoolUpgradeProfilePropertiesUpgradesItem
type AgentPoolUpgradeSettings = original.AgentPoolUpgradeSettings
type AgentPoolsClient = original.AgentPoolsClient
type AgentPoolsCreateOrUpdateFuture = original.AgentPoolsCreateOrUpdateFuture
type AgentPoolsDeleteFuture = original.AgentPoolsDeleteFuture
type AgentPoolsUpgradeNodeImageVersionFuture = original.AgentPoolsUpgradeNodeImageVersionFuture
type AzureEntityResource = original.AzureEntityResource
type AzureKeyVaultKms = original.AzureKeyVaultKms
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type CommandResultProperties = original.CommandResultProperties
type CreationData = original.CreationData
type CredentialResult = original.CredentialResult
type CredentialResults = original.CredentialResults
type DiagnosticsProfile = original.DiagnosticsProfile
type EndpointDependency = original.EndpointDependency
type EndpointDetail = original.EndpointDetail
type ExtendedLocation = original.ExtendedLocation
type KubeletConfig = original.KubeletConfig
type LinuxOSConfig = original.LinuxOSConfig
type LinuxProfile = original.LinuxProfile
type MaintenanceConfiguration = original.MaintenanceConfiguration
type MaintenanceConfigurationListResult = original.MaintenanceConfigurationListResult
type MaintenanceConfigurationListResultIterator = original.MaintenanceConfigurationListResultIterator
type MaintenanceConfigurationListResultPage = original.MaintenanceConfigurationListResultPage
type MaintenanceConfigurationProperties = original.MaintenanceConfigurationProperties
type MaintenanceConfigurationsClient = original.MaintenanceConfigurationsClient
type ManagedCluster = original.ManagedCluster
type ManagedClusterAADProfile = original.ManagedClusterAADProfile
type ManagedClusterAPIServerAccessProfile = original.ManagedClusterAPIServerAccessProfile
type ManagedClusterAccessProfile = original.ManagedClusterAccessProfile
type ManagedClusterAddonProfile = original.ManagedClusterAddonProfile
type ManagedClusterAddonProfileIdentity = original.ManagedClusterAddonProfileIdentity
type ManagedClusterAgentPoolProfile = original.ManagedClusterAgentPoolProfile
type ManagedClusterAgentPoolProfileProperties = original.ManagedClusterAgentPoolProfileProperties
type ManagedClusterAutoUpgradeProfile = original.ManagedClusterAutoUpgradeProfile
type ManagedClusterHTTPProxyConfig = original.ManagedClusterHTTPProxyConfig
type ManagedClusterIdentity = original.ManagedClusterIdentity
type ManagedClusterIdentityUserAssignedIdentitiesValue = original.ManagedClusterIdentityUserAssignedIdentitiesValue
type ManagedClusterListResult = original.ManagedClusterListResult
type ManagedClusterListResultIterator = original.ManagedClusterListResultIterator
type ManagedClusterListResultPage = original.ManagedClusterListResultPage
type ManagedClusterLoadBalancerProfile = original.ManagedClusterLoadBalancerProfile
type ManagedClusterLoadBalancerProfileManagedOutboundIPs = original.ManagedClusterLoadBalancerProfileManagedOutboundIPs
type ManagedClusterLoadBalancerProfileOutboundIPPrefixes = original.ManagedClusterLoadBalancerProfileOutboundIPPrefixes
type ManagedClusterLoadBalancerProfileOutboundIPs = original.ManagedClusterLoadBalancerProfileOutboundIPs
type ManagedClusterManagedOutboundIPProfile = original.ManagedClusterManagedOutboundIPProfile
type ManagedClusterNATGatewayProfile = original.ManagedClusterNATGatewayProfile
type ManagedClusterPodIdentity = original.ManagedClusterPodIdentity
type ManagedClusterPodIdentityException = original.ManagedClusterPodIdentityException
type ManagedClusterPodIdentityProfile = original.ManagedClusterPodIdentityProfile
type ManagedClusterPodIdentityProvisioningError = original.ManagedClusterPodIdentityProvisioningError
type ManagedClusterPodIdentityProvisioningErrorBody = original.ManagedClusterPodIdentityProvisioningErrorBody
type ManagedClusterPodIdentityProvisioningInfo = original.ManagedClusterPodIdentityProvisioningInfo
type ManagedClusterPoolUpgradeProfile = original.ManagedClusterPoolUpgradeProfile
type ManagedClusterPoolUpgradeProfileUpgradesItem = original.ManagedClusterPoolUpgradeProfileUpgradesItem
type ManagedClusterProperties = original.ManagedClusterProperties
type ManagedClusterPropertiesAutoScalerProfile = original.ManagedClusterPropertiesAutoScalerProfile
type ManagedClusterSKU = original.ManagedClusterSKU
type ManagedClusterSecurityProfile = original.ManagedClusterSecurityProfile
type ManagedClusterSecurityProfileDefender = original.ManagedClusterSecurityProfileDefender
type ManagedClusterSecurityProfileDefenderSecurityMonitoring = original.ManagedClusterSecurityProfileDefenderSecurityMonitoring
type ManagedClusterServicePrincipalProfile = original.ManagedClusterServicePrincipalProfile
type ManagedClusterStorageProfile = original.ManagedClusterStorageProfile
type ManagedClusterStorageProfileDiskCSIDriver = original.ManagedClusterStorageProfileDiskCSIDriver
type ManagedClusterStorageProfileFileCSIDriver = original.ManagedClusterStorageProfileFileCSIDriver
type ManagedClusterStorageProfileSnapshotController = original.ManagedClusterStorageProfileSnapshotController
type ManagedClusterUpgradeProfile = original.ManagedClusterUpgradeProfile
type ManagedClusterUpgradeProfileProperties = original.ManagedClusterUpgradeProfileProperties
type ManagedClusterWindowsProfile = original.ManagedClusterWindowsProfile
type ManagedClustersClient = original.ManagedClustersClient
type ManagedClustersCreateOrUpdateFuture = original.ManagedClustersCreateOrUpdateFuture
type ManagedClustersDeleteFuture = original.ManagedClustersDeleteFuture
type ManagedClustersResetAADProfileFuture = original.ManagedClustersResetAADProfileFuture
type ManagedClustersResetServicePrincipalProfileFuture = original.ManagedClustersResetServicePrincipalProfileFuture
type ManagedClustersRotateClusterCertificatesFuture = original.ManagedClustersRotateClusterCertificatesFuture
type ManagedClustersRunCommandFuture = original.ManagedClustersRunCommandFuture
type ManagedClustersStartFuture = original.ManagedClustersStartFuture
type ManagedClustersStopFuture = original.ManagedClustersStopFuture
type ManagedClustersUpdateTagsFuture = original.ManagedClustersUpdateTagsFuture
type MasterProfile = original.MasterProfile
type NetworkProfile = original.NetworkProfile
type OSOptionProfile = original.OSOptionProfile
type OSOptionProperty = original.OSOptionProperty
type OSOptionPropertyList = original.OSOptionPropertyList
type OperationListResult = original.OperationListResult
type OperationValue = original.OperationValue
type OperationValueDisplay = original.OperationValueDisplay
type OperationsClient = original.OperationsClient
type OutboundEnvironmentEndpoint = original.OutboundEnvironmentEndpoint
type OutboundEnvironmentEndpointCollection = original.OutboundEnvironmentEndpointCollection
type OutboundEnvironmentEndpointCollectionIterator = original.OutboundEnvironmentEndpointCollectionIterator
type OutboundEnvironmentEndpointCollectionPage = original.OutboundEnvironmentEndpointCollectionPage
type PowerState = original.PowerState
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionListResult = original.PrivateEndpointConnectionListResult
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateLinkResource = original.PrivateLinkResource
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type PrivateLinkResourcesListResult = original.PrivateLinkResourcesListResult
type PrivateLinkServiceConnectionState = original.PrivateLinkServiceConnectionState
type ProxyResource = original.ProxyResource
type ResolvePrivateLinkServiceIDClient = original.ResolvePrivateLinkServiceIDClient
type Resource = original.Resource
type ResourceReference = original.ResourceReference
type RunCommandRequest = original.RunCommandRequest
type RunCommandResult = original.RunCommandResult
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type Snapshot = original.Snapshot
type SnapshotListResult = original.SnapshotListResult
type SnapshotListResultIterator = original.SnapshotListResultIterator
type SnapshotListResultPage = original.SnapshotListResultPage
type SnapshotProperties = original.SnapshotProperties
type SnapshotsClient = original.SnapshotsClient
type SubResource = original.SubResource
type SysctlConfig = original.SysctlConfig
type SystemData = original.SystemData
type TagsObject = original.TagsObject
type TimeInWeek = original.TimeInWeek
type TimeSpan = original.TimeSpan
type TrackedResource = original.TrackedResource
type UserAssignedIdentity = original.UserAssignedIdentity
type VMDiagnostics = original.VMDiagnostics
type WindowsGmsaProfile = original.WindowsGmsaProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAgentPoolListResultIterator(page AgentPoolListResultPage) AgentPoolListResultIterator {
	return original.NewAgentPoolListResultIterator(page)
}
func NewAgentPoolListResultPage(cur AgentPoolListResult, getNextPage func(context.Context, AgentPoolListResult) (AgentPoolListResult, error)) AgentPoolListResultPage {
	return original.NewAgentPoolListResultPage(cur, getNextPage)
}
func NewAgentPoolsClient(subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClient(subscriptionID)
}
func NewAgentPoolsClientWithBaseURI(baseURI string, subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMaintenanceConfigurationListResultIterator(page MaintenanceConfigurationListResultPage) MaintenanceConfigurationListResultIterator {
	return original.NewMaintenanceConfigurationListResultIterator(page)
}
func NewMaintenanceConfigurationListResultPage(cur MaintenanceConfigurationListResult, getNextPage func(context.Context, MaintenanceConfigurationListResult) (MaintenanceConfigurationListResult, error)) MaintenanceConfigurationListResultPage {
	return original.NewMaintenanceConfigurationListResultPage(cur, getNextPage)
}
func NewMaintenanceConfigurationsClient(subscriptionID string) MaintenanceConfigurationsClient {
	return original.NewMaintenanceConfigurationsClient(subscriptionID)
}
func NewMaintenanceConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) MaintenanceConfigurationsClient {
	return original.NewMaintenanceConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagedClusterListResultIterator(page ManagedClusterListResultPage) ManagedClusterListResultIterator {
	return original.NewManagedClusterListResultIterator(page)
}
func NewManagedClusterListResultPage(cur ManagedClusterListResult, getNextPage func(context.Context, ManagedClusterListResult) (ManagedClusterListResult, error)) ManagedClusterListResultPage {
	return original.NewManagedClusterListResultPage(cur, getNextPage)
}
func NewManagedClustersClient(subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClient(subscriptionID)
}
func NewManagedClustersClientWithBaseURI(baseURI string, subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOutboundEnvironmentEndpointCollectionIterator(page OutboundEnvironmentEndpointCollectionPage) OutboundEnvironmentEndpointCollectionIterator {
	return original.NewOutboundEnvironmentEndpointCollectionIterator(page)
}
func NewOutboundEnvironmentEndpointCollectionPage(cur OutboundEnvironmentEndpointCollection, getNextPage func(context.Context, OutboundEnvironmentEndpointCollection) (OutboundEnvironmentEndpointCollection, error)) OutboundEnvironmentEndpointCollectionPage {
	return original.NewOutboundEnvironmentEndpointCollectionPage(cur, getNextPage)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewResolvePrivateLinkServiceIDClient(subscriptionID string) ResolvePrivateLinkServiceIDClient {
	return original.NewResolvePrivateLinkServiceIDClient(subscriptionID)
}
func NewResolvePrivateLinkServiceIDClientWithBaseURI(baseURI string, subscriptionID string) ResolvePrivateLinkServiceIDClient {
	return original.NewResolvePrivateLinkServiceIDClientWithBaseURI(baseURI, subscriptionID)
}
func NewSnapshotListResultIterator(page SnapshotListResultPage) SnapshotListResultIterator {
	return original.NewSnapshotListResultIterator(page)
}
func NewSnapshotListResultPage(cur SnapshotListResult, getNextPage func(context.Context, SnapshotListResult) (SnapshotListResult, error)) SnapshotListResultPage {
	return original.NewSnapshotListResultPage(cur, getNextPage)
}
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAgentPoolModeValues() []AgentPoolMode {
	return original.PossibleAgentPoolModeValues()
}
func PossibleAgentPoolTypeValues() []AgentPoolType {
	return original.PossibleAgentPoolTypeValues()
}
func PossibleCodeValues() []Code {
	return original.PossibleCodeValues()
}
func PossibleConnectionStatusValues() []ConnectionStatus {
	return original.PossibleConnectionStatusValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleExpanderValues() []Expander {
	return original.PossibleExpanderValues()
}
func PossibleExtendedLocationTypesValues() []ExtendedLocationTypes {
	return original.PossibleExtendedLocationTypesValues()
}
func PossibleFormatValues() []Format {
	return original.PossibleFormatValues()
}
func PossibleGPUInstanceProfileValues() []GPUInstanceProfile {
	return original.PossibleGPUInstanceProfileValues()
}
func PossibleIPFamilyValues() []IPFamily {
	return original.PossibleIPFamilyValues()
}
func PossibleKeyVaultNetworkAccessTypesValues() []KeyVaultNetworkAccessTypes {
	return original.PossibleKeyVaultNetworkAccessTypesValues()
}
func PossibleKubeletDiskTypeValues() []KubeletDiskType {
	return original.PossibleKubeletDiskTypeValues()
}
func PossibleLicenseTypeValues() []LicenseType {
	return original.PossibleLicenseTypeValues()
}
func PossibleLoadBalancerSkuValues() []LoadBalancerSku {
	return original.PossibleLoadBalancerSkuValues()
}
func PossibleManagedClusterPodIdentityProvisioningStateValues() []ManagedClusterPodIdentityProvisioningState {
	return original.PossibleManagedClusterPodIdentityProvisioningStateValues()
}
func PossibleManagedClusterSKUNameValues() []ManagedClusterSKUName {
	return original.PossibleManagedClusterSKUNameValues()
}
func PossibleManagedClusterSKUTierValues() []ManagedClusterSKUTier {
	return original.PossibleManagedClusterSKUTierValues()
}
func PossibleNetworkModeValues() []NetworkMode {
	return original.PossibleNetworkModeValues()
}
func PossibleNetworkPluginValues() []NetworkPlugin {
	return original.PossibleNetworkPluginValues()
}
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return original.PossibleNetworkPolicyValues()
}
func PossibleOSDiskTypeValues() []OSDiskType {
	return original.PossibleOSDiskTypeValues()
}
func PossibleOSSKUValues() []OSSKU {
	return original.PossibleOSSKUValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossibleOutboundTypeValues() []OutboundType {
	return original.PossibleOutboundTypeValues()
}
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return original.PossiblePrivateEndpointConnectionProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleScaleDownModeValues() []ScaleDownMode {
	return original.PossibleScaleDownModeValues()
}
func PossibleScaleSetEvictionPolicyValues() []ScaleSetEvictionPolicy {
	return original.PossibleScaleSetEvictionPolicyValues()
}
func PossibleScaleSetPriorityValues() []ScaleSetPriority {
	return original.PossibleScaleSetPriorityValues()
}
func PossibleSnapshotTypeValues() []SnapshotType {
	return original.PossibleSnapshotTypeValues()
}
func PossibleStorageProfileTypesValues() []StorageProfileTypes {
	return original.PossibleStorageProfileTypesValues()
}
func PossibleUpgradeChannelValues() []UpgradeChannel {
	return original.PossibleUpgradeChannelValues()
}
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return original.PossibleVMSizeTypesValues()
}
func PossibleWeekDayValues() []WeekDay {
	return original.PossibleWeekDayValues()
}
func PossibleWorkloadRuntimeValues() []WorkloadRuntime {
	return original.PossibleWorkloadRuntimeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
