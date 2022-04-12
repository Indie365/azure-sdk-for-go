//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AccessProfile.
func (a AccessProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateByteArray(objectMap, "kubeConfig", a.KubeConfig, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AccessProfile.
func (a *AccessProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "kubeConfig":
			err = runtime.DecodeByteArray(string(val), &a.KubeConfig, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AgentPoolAvailableVersionsProperties.
func (a AgentPoolAvailableVersionsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "agentPoolVersions", a.AgentPoolVersions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AgentPoolListResult.
func (a AgentPoolListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AgentPoolUpgradeProfileProperties.
func (a AgentPoolUpgradeProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "kubernetesVersion", a.KubernetesVersion)
	populate(objectMap, "latestNodeImageVersion", a.LatestNodeImageVersion)
	populate(objectMap, "osType", a.OSType)
	populate(objectMap, "upgrades", a.Upgrades)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CloudErrorBody.
func (c CloudErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", c.Code)
	populate(objectMap, "details", c.Details)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "target", c.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CommandResultProperties.
func (c CommandResultProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "exitCode", c.ExitCode)
	populateTimeRFC3339(objectMap, "finishedAt", c.FinishedAt)
	populate(objectMap, "logs", c.Logs)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "reason", c.Reason)
	populateTimeRFC3339(objectMap, "startedAt", c.StartedAt)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CommandResultProperties.
func (c *CommandResultProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "exitCode":
			err = unpopulate(val, &c.ExitCode)
			delete(rawMsg, key)
		case "finishedAt":
			err = unpopulateTimeRFC3339(val, &c.FinishedAt)
			delete(rawMsg, key)
		case "logs":
			err = unpopulate(val, &c.Logs)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &c.ProvisioningState)
			delete(rawMsg, key)
		case "reason":
			err = unpopulate(val, &c.Reason)
			delete(rawMsg, key)
		case "startedAt":
			err = unpopulateTimeRFC3339(val, &c.StartedAt)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CredentialResult.
func (c CredentialResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", c.Name)
	populateByteArray(objectMap, "value", c.Value, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CredentialResult.
func (c *CredentialResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, &c.Name)
			delete(rawMsg, key)
		case "value":
			err = runtime.DecodeByteArray(string(val), &c.Value, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CredentialResults.
func (c CredentialResults) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "kubeconfigs", c.Kubeconfigs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EndpointDependency.
func (e EndpointDependency) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "domainName", e.DomainName)
	populate(objectMap, "endpointDetails", e.EndpointDetails)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type KubeletConfig.
func (k KubeletConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedUnsafeSysctls", k.AllowedUnsafeSysctls)
	populate(objectMap, "cpuCfsQuota", k.CPUCfsQuota)
	populate(objectMap, "cpuCfsQuotaPeriod", k.CPUCfsQuotaPeriod)
	populate(objectMap, "cpuManagerPolicy", k.CPUManagerPolicy)
	populate(objectMap, "containerLogMaxFiles", k.ContainerLogMaxFiles)
	populate(objectMap, "containerLogMaxSizeMB", k.ContainerLogMaxSizeMB)
	populate(objectMap, "failSwapOn", k.FailSwapOn)
	populate(objectMap, "imageGcHighThreshold", k.ImageGcHighThreshold)
	populate(objectMap, "imageGcLowThreshold", k.ImageGcLowThreshold)
	populate(objectMap, "podMaxPids", k.PodMaxPids)
	populate(objectMap, "topologyManagerPolicy", k.TopologyManagerPolicy)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MaintenanceConfigurationListResult.
func (m MaintenanceConfigurationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MaintenanceConfigurationProperties.
func (m MaintenanceConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "notAllowedTime", m.NotAllowedTime)
	populate(objectMap, "timeInWeek", m.TimeInWeek)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedCluster.
func (m ManagedCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedLocation", m.ExtendedLocation)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "identity", m.Identity)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "sku", m.SKU)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterAADProfile.
func (m ManagedClusterAADProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "adminGroupObjectIDs", m.AdminGroupObjectIDs)
	populate(objectMap, "clientAppID", m.ClientAppID)
	populate(objectMap, "enableAzureRBAC", m.EnableAzureRBAC)
	populate(objectMap, "managed", m.Managed)
	populate(objectMap, "serverAppID", m.ServerAppID)
	populate(objectMap, "serverAppSecret", m.ServerAppSecret)
	populate(objectMap, "tenantID", m.TenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterAPIServerAccessProfile.
func (m ManagedClusterAPIServerAccessProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizedIPRanges", m.AuthorizedIPRanges)
	populate(objectMap, "disableRunCommand", m.DisableRunCommand)
	populate(objectMap, "enablePrivateCluster", m.EnablePrivateCluster)
	populate(objectMap, "enablePrivateClusterPublicFQDN", m.EnablePrivateClusterPublicFQDN)
	populate(objectMap, "privateDNSZone", m.PrivateDNSZone)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterAccessProfile.
func (m ManagedClusterAccessProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "systemData", m.SystemData)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterAddonProfile.
func (m ManagedClusterAddonProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "config", m.Config)
	populate(objectMap, "enabled", m.Enabled)
	populate(objectMap, "identity", m.Identity)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterAgentPoolProfile.
func (m ManagedClusterAgentPoolProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "availabilityZones", m.AvailabilityZones)
	populate(objectMap, "count", m.Count)
	populate(objectMap, "creationData", m.CreationData)
	populate(objectMap, "enableAutoScaling", m.EnableAutoScaling)
	populate(objectMap, "enableEncryptionAtHost", m.EnableEncryptionAtHost)
	populate(objectMap, "enableFIPS", m.EnableFIPS)
	populate(objectMap, "enableNodePublicIP", m.EnableNodePublicIP)
	populate(objectMap, "enableUltraSSD", m.EnableUltraSSD)
	populate(objectMap, "gpuInstanceProfile", m.GpuInstanceProfile)
	populate(objectMap, "kubeletConfig", m.KubeletConfig)
	populate(objectMap, "kubeletDiskType", m.KubeletDiskType)
	populate(objectMap, "linuxOSConfig", m.LinuxOSConfig)
	populate(objectMap, "maxCount", m.MaxCount)
	populate(objectMap, "maxPods", m.MaxPods)
	populate(objectMap, "minCount", m.MinCount)
	populate(objectMap, "mode", m.Mode)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "nodeImageVersion", m.NodeImageVersion)
	populate(objectMap, "nodeLabels", m.NodeLabels)
	populate(objectMap, "nodePublicIPPrefixID", m.NodePublicIPPrefixID)
	populate(objectMap, "nodeTaints", m.NodeTaints)
	populate(objectMap, "osDiskSizeGB", m.OSDiskSizeGB)
	populate(objectMap, "osDiskType", m.OSDiskType)
	populate(objectMap, "osSKU", m.OSSKU)
	populate(objectMap, "osType", m.OSType)
	populate(objectMap, "orchestratorVersion", m.OrchestratorVersion)
	populate(objectMap, "podSubnetID", m.PodSubnetID)
	populate(objectMap, "powerState", m.PowerState)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "proximityPlacementGroupID", m.ProximityPlacementGroupID)
	populate(objectMap, "scaleDownMode", m.ScaleDownMode)
	populate(objectMap, "scaleSetEvictionPolicy", m.ScaleSetEvictionPolicy)
	populate(objectMap, "scaleSetPriority", m.ScaleSetPriority)
	populate(objectMap, "spotMaxPrice", m.SpotMaxPrice)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "upgradeSettings", m.UpgradeSettings)
	populate(objectMap, "vmSize", m.VMSize)
	populate(objectMap, "vnetSubnetID", m.VnetSubnetID)
	populate(objectMap, "workloadRuntime", m.WorkloadRuntime)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterAgentPoolProfileProperties.
func (m ManagedClusterAgentPoolProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "availabilityZones", m.AvailabilityZones)
	populate(objectMap, "count", m.Count)
	populate(objectMap, "creationData", m.CreationData)
	populate(objectMap, "enableAutoScaling", m.EnableAutoScaling)
	populate(objectMap, "enableEncryptionAtHost", m.EnableEncryptionAtHost)
	populate(objectMap, "enableFIPS", m.EnableFIPS)
	populate(objectMap, "enableNodePublicIP", m.EnableNodePublicIP)
	populate(objectMap, "enableUltraSSD", m.EnableUltraSSD)
	populate(objectMap, "gpuInstanceProfile", m.GpuInstanceProfile)
	populate(objectMap, "kubeletConfig", m.KubeletConfig)
	populate(objectMap, "kubeletDiskType", m.KubeletDiskType)
	populate(objectMap, "linuxOSConfig", m.LinuxOSConfig)
	populate(objectMap, "maxCount", m.MaxCount)
	populate(objectMap, "maxPods", m.MaxPods)
	populate(objectMap, "minCount", m.MinCount)
	populate(objectMap, "mode", m.Mode)
	populate(objectMap, "nodeImageVersion", m.NodeImageVersion)
	populate(objectMap, "nodeLabels", m.NodeLabels)
	populate(objectMap, "nodePublicIPPrefixID", m.NodePublicIPPrefixID)
	populate(objectMap, "nodeTaints", m.NodeTaints)
	populate(objectMap, "osDiskSizeGB", m.OSDiskSizeGB)
	populate(objectMap, "osDiskType", m.OSDiskType)
	populate(objectMap, "osSKU", m.OSSKU)
	populate(objectMap, "osType", m.OSType)
	populate(objectMap, "orchestratorVersion", m.OrchestratorVersion)
	populate(objectMap, "podSubnetID", m.PodSubnetID)
	populate(objectMap, "powerState", m.PowerState)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "proximityPlacementGroupID", m.ProximityPlacementGroupID)
	populate(objectMap, "scaleDownMode", m.ScaleDownMode)
	populate(objectMap, "scaleSetEvictionPolicy", m.ScaleSetEvictionPolicy)
	populate(objectMap, "scaleSetPriority", m.ScaleSetPriority)
	populate(objectMap, "spotMaxPrice", m.SpotMaxPrice)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "upgradeSettings", m.UpgradeSettings)
	populate(objectMap, "vmSize", m.VMSize)
	populate(objectMap, "vnetSubnetID", m.VnetSubnetID)
	populate(objectMap, "workloadRuntime", m.WorkloadRuntime)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterHTTPProxyConfig.
func (m ManagedClusterHTTPProxyConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "httpProxy", m.HTTPProxy)
	populate(objectMap, "httpsProxy", m.HTTPSProxy)
	populate(objectMap, "noProxy", m.NoProxy)
	populate(objectMap, "trustedCa", m.TrustedCa)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterIdentity.
func (m ManagedClusterIdentity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "principalId", m.PrincipalID)
	populate(objectMap, "tenantId", m.TenantID)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "userAssignedIdentities", m.UserAssignedIdentities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterListResult.
func (m ManagedClusterListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", m.NextLink)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterLoadBalancerProfile.
func (m ManagedClusterLoadBalancerProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allocatedOutboundPorts", m.AllocatedOutboundPorts)
	populate(objectMap, "effectiveOutboundIPs", m.EffectiveOutboundIPs)
	populate(objectMap, "enableMultipleStandardLoadBalancers", m.EnableMultipleStandardLoadBalancers)
	populate(objectMap, "idleTimeoutInMinutes", m.IdleTimeoutInMinutes)
	populate(objectMap, "managedOutboundIPs", m.ManagedOutboundIPs)
	populate(objectMap, "outboundIPPrefixes", m.OutboundIPPrefixes)
	populate(objectMap, "outboundIPs", m.OutboundIPs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterLoadBalancerProfileOutboundIPPrefixes.
func (m ManagedClusterLoadBalancerProfileOutboundIPPrefixes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "publicIPPrefixes", m.PublicIPPrefixes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterLoadBalancerProfileOutboundIPs.
func (m ManagedClusterLoadBalancerProfileOutboundIPs) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "publicIPs", m.PublicIPs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterNATGatewayProfile.
func (m ManagedClusterNATGatewayProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "effectiveOutboundIPs", m.EffectiveOutboundIPs)
	populate(objectMap, "idleTimeoutInMinutes", m.IdleTimeoutInMinutes)
	populate(objectMap, "managedOutboundIPProfile", m.ManagedOutboundIPProfile)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterPodIdentityException.
func (m ManagedClusterPodIdentityException) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", m.Name)
	populate(objectMap, "namespace", m.Namespace)
	populate(objectMap, "podLabels", m.PodLabels)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterPodIdentityProfile.
func (m ManagedClusterPodIdentityProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowNetworkPluginKubenet", m.AllowNetworkPluginKubenet)
	populate(objectMap, "enabled", m.Enabled)
	populate(objectMap, "userAssignedIdentities", m.UserAssignedIdentities)
	populate(objectMap, "userAssignedIdentityExceptions", m.UserAssignedIdentityExceptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterPodIdentityProvisioningErrorBody.
func (m ManagedClusterPodIdentityProvisioningErrorBody) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", m.Code)
	populate(objectMap, "details", m.Details)
	populate(objectMap, "message", m.Message)
	populate(objectMap, "target", m.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterPoolUpgradeProfile.
func (m ManagedClusterPoolUpgradeProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "kubernetesVersion", m.KubernetesVersion)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "osType", m.OSType)
	populate(objectMap, "upgrades", m.Upgrades)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterProperties.
func (m ManagedClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "aadProfile", m.AADProfile)
	populate(objectMap, "apiServerAccessProfile", m.APIServerAccessProfile)
	populate(objectMap, "addonProfiles", m.AddonProfiles)
	populate(objectMap, "agentPoolProfiles", m.AgentPoolProfiles)
	populate(objectMap, "autoScalerProfile", m.AutoScalerProfile)
	populate(objectMap, "autoUpgradeProfile", m.AutoUpgradeProfile)
	populate(objectMap, "azurePortalFQDN", m.AzurePortalFQDN)
	populate(objectMap, "dnsPrefix", m.DNSPrefix)
	populate(objectMap, "disableLocalAccounts", m.DisableLocalAccounts)
	populate(objectMap, "diskEncryptionSetID", m.DiskEncryptionSetID)
	populate(objectMap, "enablePodSecurityPolicy", m.EnablePodSecurityPolicy)
	populate(objectMap, "enableRBAC", m.EnableRBAC)
	populate(objectMap, "fqdn", m.Fqdn)
	populate(objectMap, "fqdnSubdomain", m.FqdnSubdomain)
	populate(objectMap, "httpProxyConfig", m.HTTPProxyConfig)
	populate(objectMap, "identityProfile", m.IdentityProfile)
	populate(objectMap, "kubernetesVersion", m.KubernetesVersion)
	populate(objectMap, "linuxProfile", m.LinuxProfile)
	populate(objectMap, "maxAgentPools", m.MaxAgentPools)
	populate(objectMap, "networkProfile", m.NetworkProfile)
	populate(objectMap, "nodeResourceGroup", m.NodeResourceGroup)
	populate(objectMap, "podIdentityProfile", m.PodIdentityProfile)
	populate(objectMap, "powerState", m.PowerState)
	populate(objectMap, "privateFQDN", m.PrivateFQDN)
	populate(objectMap, "privateLinkResources", m.PrivateLinkResources)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", m.PublicNetworkAccess)
	populate(objectMap, "securityProfile", m.SecurityProfile)
	populate(objectMap, "servicePrincipalProfile", m.ServicePrincipalProfile)
	populate(objectMap, "windowsProfile", m.WindowsProfile)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ManagedClusterUpgradeProfileProperties.
func (m ManagedClusterUpgradeProfileProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "agentPoolProfiles", m.AgentPoolProfiles)
	populate(objectMap, "controlPlaneProfile", m.ControlPlaneProfile)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkProfile.
func (n NetworkProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "dnsServiceIP", n.DNSServiceIP)
	populate(objectMap, "dockerBridgeCidr", n.DockerBridgeCidr)
	populate(objectMap, "ipFamilies", n.IPFamilies)
	populate(objectMap, "loadBalancerProfile", n.LoadBalancerProfile)
	populate(objectMap, "loadBalancerSku", n.LoadBalancerSKU)
	populate(objectMap, "natGatewayProfile", n.NatGatewayProfile)
	populate(objectMap, "networkMode", n.NetworkMode)
	populate(objectMap, "networkPlugin", n.NetworkPlugin)
	populate(objectMap, "networkPolicy", n.NetworkPolicy)
	populate(objectMap, "outboundType", n.OutboundType)
	populate(objectMap, "podCidr", n.PodCidr)
	populate(objectMap, "podCidrs", n.PodCidrs)
	populate(objectMap, "serviceCidr", n.ServiceCidr)
	populate(objectMap, "serviceCidrs", n.ServiceCidrs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OSOptionPropertyList.
func (o OSOptionPropertyList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "osOptionPropertyList", o.OSOptionPropertyList)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEnvironmentEndpoint.
func (o OutboundEnvironmentEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", o.Category)
	populate(objectMap, "endpoints", o.Endpoints)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OutboundEnvironmentEndpointCollection.
func (o OutboundEnvironmentEndpointCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateEndpointConnectionListResult.
func (p PrivateEndpointConnectionListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResource.
func (p PrivateLinkResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "privateLinkServiceID", p.PrivateLinkServiceID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourcesListResult.
func (p PrivateLinkResourcesListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", p.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SSHConfiguration.
func (s SSHConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "publicKeys", s.PublicKeys)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Snapshot.
func (s Snapshot) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SnapshotListResult.
func (s SnapshotListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagsObject.
func (t TagsObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TimeInWeek.
func (t TimeInWeek) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "day", t.Day)
	populate(objectMap, "hourSlots", t.HourSlots)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TimeSpan.
func (t TimeSpan) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "end", t.End)
	populateTimeRFC3339(objectMap, "start", t.Start)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TimeSpan.
func (t *TimeSpan) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "end":
			err = unpopulateTimeRFC3339(val, &t.End)
			delete(rawMsg, key)
		case "start":
			err = unpopulateTimeRFC3339(val, &t.Start)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
