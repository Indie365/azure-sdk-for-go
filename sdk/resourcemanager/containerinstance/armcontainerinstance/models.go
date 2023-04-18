//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerinstance

import "time"

// AzureFileVolume - The properties of the Azure File volume. Azure File shares are mounted as volumes.
type AzureFileVolume struct {
	// REQUIRED; The name of the Azure File share to be mounted as a volume.
	ShareName *string

	// REQUIRED; The name of the storage account that contains the Azure File share.
	StorageAccountName *string

	// The flag indicating whether the Azure File shared mounted as a volume is read-only.
	ReadOnly *bool

	// The storage account access key used to access the Azure File share.
	StorageAccountKey *string
}

// CachedImages - The cached image and OS type.
type CachedImages struct {
	// REQUIRED; The cached image name.
	Image *string

	// REQUIRED; The OS type of the cached image.
	OSType *string
}

// CachedImagesListResult - The response containing cached images.
type CachedImagesListResult struct {
	// The URI to fetch the next page of cached images.
	NextLink *string

	// The list of cached images.
	Value []*CachedImages
}

// Capabilities - The regional capabilities.
type Capabilities struct {
	// READ-ONLY; The supported capabilities.
	Capabilities *CapabilitiesCapabilities

	// READ-ONLY; The GPU sku that this capability describes.
	Gpu *string

	// READ-ONLY; The ip address type that this capability describes.
	IPAddressType *string

	// READ-ONLY; The resource location.
	Location *string

	// READ-ONLY; The OS type that this capability describes.
	OSType *string

	// READ-ONLY; The resource type that this capability describes.
	ResourceType *string
}

// CapabilitiesCapabilities - The supported capabilities.
type CapabilitiesCapabilities struct {
	// READ-ONLY; The maximum allowed CPU request in cores.
	MaxCPU *float32

	// READ-ONLY; The maximum allowed GPU count.
	MaxGpuCount *float32

	// READ-ONLY; The maximum allowed memory request in GB.
	MaxMemoryInGB *float32
}

// CapabilitiesListResult - The response containing list of capabilities.
type CapabilitiesListResult struct {
	// The URI to fetch the next page of capabilities.
	NextLink *string

	// The list of capabilities.
	Value []*Capabilities
}

// Container - A container instance.
type Container struct {
	// REQUIRED; The user-provided name of the container instance.
	Name *string

	// REQUIRED; The properties of the container instance.
	Properties *ContainerProperties
}

// ContainerAttachResponse - The information for the output stream from container attach.
type ContainerAttachResponse struct {
	// The password to the output stream from the attach. Send as an Authorization header value when connecting to the websocketUri.
	Password *string

	// The uri for the output stream from the attach.
	WebSocketURI *string
}

// ContainerExec - The container execution command, for liveness or readiness probe
type ContainerExec struct {
	// The commands to execute within the container.
	Command []*string
}

// ContainerExecRequest - The container exec request.
type ContainerExecRequest struct {
	// The command to be executed.
	Command *string

	// The size of the terminal.
	TerminalSize *ContainerExecRequestTerminalSize
}

// ContainerExecRequestTerminalSize - The size of the terminal.
type ContainerExecRequestTerminalSize struct {
	// The column size of the terminal
	Cols *int32

	// The row size of the terminal
	Rows *int32
}

// ContainerExecResponse - The information for the container exec command.
type ContainerExecResponse struct {
	// The password to start the exec command.
	Password *string

	// The uri for the exec websocket.
	WebSocketURI *string
}

// ContainerGroup - A container group.
type ContainerGroup struct {
	// REQUIRED; The container group properties
	Properties *ContainerGroupPropertiesProperties

	// The identity of the container group, if configured.
	Identity *ContainerGroupIdentity

	// The resource location.
	Location *string

	// The resource tags.
	Tags map[string]*string

	// The zones for the container group.
	Zones []*string

	// READ-ONLY; The resource id.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// ContainerGroupDiagnostics - Container group diagnostic information.
type ContainerGroupDiagnostics struct {
	// Container group log analytics information.
	LogAnalytics *LogAnalytics
}

// ContainerGroupIdentity - Identity for the container group.
type ContainerGroupIdentity struct {
	// The type of identity used for the container group. The type 'SystemAssigned, UserAssigned' includes both an implicitly
	// created identity and a set of user assigned identities. The type 'None' will
	// remove any identities from the container group.
	Type *ResourceIdentityType

	// The list of user identities associated with the container group.
	UserAssignedIdentities map[string]*UserAssignedIdentities

	// READ-ONLY; The principal id of the container group identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant id associated with the container group. This property will only be provided for a system assigned
	// identity.
	TenantID *string
}

// ContainerGroupListResult - The container group list response that contains the container group properties.
type ContainerGroupListResult struct {
	// The URI to fetch the next page of container groups.
	NextLink *string

	// The list of container groups.
	Value []*ContainerGroup
}

// ContainerGroupProperties - The container group properties
type ContainerGroupProperties struct {
	// REQUIRED; The container group properties
	Properties *ContainerGroupPropertiesProperties

	// The identity of the container group, if configured.
	Identity *ContainerGroupIdentity
}

// ContainerGroupPropertiesInstanceView - The instance view of the container group. Only valid in response.
type ContainerGroupPropertiesInstanceView struct {
	// READ-ONLY; The events of this container group.
	Events []*Event

	// READ-ONLY; The state of the container group. Only valid in response.
	State *string
}

// ContainerGroupPropertiesProperties - The container group properties
type ContainerGroupPropertiesProperties struct {
	// REQUIRED; The containers within the container group.
	Containers []*Container

	// REQUIRED; The operating system type required by the containers in the container group.
	OSType *OperatingSystemTypes

	// The DNS config information for a container group.
	DNSConfig *DNSConfiguration

	// The diagnostic information for a container group.
	Diagnostics *ContainerGroupDiagnostics

	// The encryption properties for a container group.
	EncryptionProperties *EncryptionProperties

	// extensions used by virtual kubelet
	Extensions []*DeploymentExtensionSpec

	// The IP address type of the container group.
	IPAddress *IPAddress

	// The image registry credentials by which the container group is created from.
	ImageRegistryCredentials []*ImageRegistryCredential

	// The init containers for a container group.
	InitContainers []*InitContainerDefinition

	// Time in seconds in which a container group deployment would timeout and fail. The allowed maximum value is 1800 seconds.
	// If value is not provided, property is given maximum value by default.
	ProvisioningTimeoutInSeconds *int32

	// Restart policy for all containers within the container group.
	// * Always Always restart
	// * OnFailure Restart on failure
	// * Never Never restart
	RestartPolicy *ContainerGroupRestartPolicy

	// The SKU for a container group.
	SKU *ContainerGroupSKU

	// The subnet resource IDs for a container group.
	SubnetIDs []*ContainerGroupSubnetID

	// The list of volumes that can be mounted by containers in this container group.
	Volumes []*Volume

	// READ-ONLY; The instance view of the container group. Only valid in response.
	InstanceView *ContainerGroupPropertiesInstanceView

	// READ-ONLY; Flag indicating whether a custom value was provided for the provisioningTimeoutInSeconds property
	IsCustomProvisioningTimeout *IsCustomProvisioningTimeout

	// READ-ONLY; The provisioning state of the container group. This only appears in the response.
	ProvisioningState *string
}

// ContainerGroupSubnetID - Container group subnet information.
type ContainerGroupSubnetID struct {
	// REQUIRED; Resource ID of virtual network and subnet.
	ID *string

	// Friendly name for the subnet.
	Name *string
}

// ContainerGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainerGroupsClient.BeginCreateOrUpdate
// method.
type ContainerGroupsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContainerGroupsClientBeginDeleteOptions contains the optional parameters for the ContainerGroupsClient.BeginDelete method.
type ContainerGroupsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContainerGroupsClientBeginRestartOptions contains the optional parameters for the ContainerGroupsClient.BeginRestart method.
type ContainerGroupsClientBeginRestartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContainerGroupsClientBeginStartOptions contains the optional parameters for the ContainerGroupsClient.BeginStart method.
type ContainerGroupsClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ContainerGroupsClientGetOptions contains the optional parameters for the ContainerGroupsClient.Get method.
type ContainerGroupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the ContainerGroupsClient.GetOutboundNetworkDependenciesEndpoints
// method.
type ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsOptions struct {
	// placeholder for future optional parameters
}

// ContainerGroupsClientListByResourceGroupOptions contains the optional parameters for the ContainerGroupsClient.NewListByResourceGroupPager
// method.
type ContainerGroupsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ContainerGroupsClientListOptions contains the optional parameters for the ContainerGroupsClient.NewListPager method.
type ContainerGroupsClientListOptions struct {
	// placeholder for future optional parameters
}

// ContainerGroupsClientStopOptions contains the optional parameters for the ContainerGroupsClient.Stop method.
type ContainerGroupsClientStopOptions struct {
	// placeholder for future optional parameters
}

// ContainerGroupsClientUpdateOptions contains the optional parameters for the ContainerGroupsClient.Update method.
type ContainerGroupsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ContainerHTTPGet - The container Http Get settings, for liveness or readiness probe
type ContainerHTTPGet struct {
	// REQUIRED; The port number to probe.
	Port *int32

	// The HTTP headers.
	HTTPHeaders []*HTTPHeader

	// The path to probe.
	Path *string

	// The scheme.
	Scheme *Scheme
}

// ContainerPort - The port exposed on the container instance.
type ContainerPort struct {
	// REQUIRED; The port number exposed within the container group.
	Port *int32

	// The protocol associated with the port.
	Protocol *ContainerNetworkProtocol
}

// ContainerProbe - The container probe, for liveness or readiness
type ContainerProbe struct {
	// The execution command to probe
	Exec *ContainerExec

	// The failure threshold.
	FailureThreshold *int32

	// The Http Get settings to probe
	HTTPGet *ContainerHTTPGet

	// The initial delay seconds.
	InitialDelaySeconds *int32

	// The period seconds.
	PeriodSeconds *int32

	// The success threshold.
	SuccessThreshold *int32

	// The timeout seconds.
	TimeoutSeconds *int32
}

// ContainerProperties - The container instance properties.
type ContainerProperties struct {
	// REQUIRED; The name of the image used to create the container instance.
	Image *string

	// REQUIRED; The resource requirements of the container instance.
	Resources *ResourceRequirements

	// The commands to execute within the container instance in exec form.
	Command []*string

	// The environment variables to set in the container instance.
	EnvironmentVariables []*EnvironmentVariable

	// The liveness probe.
	LivenessProbe *ContainerProbe

	// The exposed ports on the container instance.
	Ports []*ContainerPort

	// The readiness probe.
	ReadinessProbe *ContainerProbe

	// The volume mounts available to the container instance.
	VolumeMounts []*VolumeMount

	// READ-ONLY; The instance view of the container instance. Only valid in response.
	InstanceView *ContainerPropertiesInstanceView
}

// ContainerPropertiesInstanceView - The instance view of the container instance. Only valid in response.
type ContainerPropertiesInstanceView struct {
	// READ-ONLY; Current container instance state.
	CurrentState *ContainerState

	// READ-ONLY; The events of the container instance.
	Events []*Event

	// READ-ONLY; Previous container instance state.
	PreviousState *ContainerState

	// READ-ONLY; The number of times that the container instance has been restarted.
	RestartCount *int32
}

// ContainerState - The container instance state.
type ContainerState struct {
	// READ-ONLY; The human-readable status of the container instance state.
	DetailStatus *string

	// READ-ONLY; The container instance exit codes correspond to those from the docker run command.
	ExitCode *int32

	// READ-ONLY; The date-time when the container instance state finished.
	FinishTime *time.Time

	// READ-ONLY; The date-time when the container instance state started.
	StartTime *time.Time

	// READ-ONLY; The state of the container instance.
	State *string
}

// ContainersClientAttachOptions contains the optional parameters for the ContainersClient.Attach method.
type ContainersClientAttachOptions struct {
	// placeholder for future optional parameters
}

// ContainersClientExecuteCommandOptions contains the optional parameters for the ContainersClient.ExecuteCommand method.
type ContainersClientExecuteCommandOptions struct {
	// placeholder for future optional parameters
}

// ContainersClientListLogsOptions contains the optional parameters for the ContainersClient.ListLogs method.
type ContainersClientListLogsOptions struct {
	// The number of lines to show from the tail of the container instance log. If not provided, all available logs are shown
	// up to 4mb.
	Tail *int32
	// If true, adds a timestamp at the beginning of every line of log output. If not provided, defaults to false.
	Timestamps *bool
}

// DNSConfiguration - DNS configuration for the container group.
type DNSConfiguration struct {
	// REQUIRED; The DNS servers for the container group.
	NameServers []*string

	// The DNS options for the container group.
	Options *string

	// The DNS search domains for hostname lookup in the container group.
	SearchDomains *string
}

// DeploymentExtensionSpec - Extension sidecars to be added to the deployment.
type DeploymentExtensionSpec struct {
	// REQUIRED; Name of the extension.
	Name *string

	// Extension specific properties
	Properties *DeploymentExtensionSpecProperties
}

// DeploymentExtensionSpecProperties - Extension specific properties
type DeploymentExtensionSpecProperties struct {
	// REQUIRED; Type of extension to be added.
	ExtensionType *string

	// REQUIRED; Version of the extension being used.
	Version *string

	// Protected settings for the extension.
	ProtectedSettings any

	// Settings for the extension.
	Settings any
}

// EncryptionProperties - The container group encryption properties.
type EncryptionProperties struct {
	// REQUIRED; The encryption key name.
	KeyName *string

	// REQUIRED; The encryption key version.
	KeyVersion *string

	// REQUIRED; The keyvault base url.
	VaultBaseURL *string

	// The keyvault managed identity.
	Identity *string
}

// EnvironmentVariable - The environment variable to set within the container instance.
type EnvironmentVariable struct {
	// REQUIRED; The name of the environment variable.
	Name *string

	// The value of the secure environment variable.
	SecureValue *string

	// The value of the environment variable.
	Value *string
}

// Event - A container group or container instance event.
type Event struct {
	// READ-ONLY; The count of the event.
	Count *int32

	// READ-ONLY; The date-time of the earliest logged event.
	FirstTimestamp *time.Time

	// READ-ONLY; The date-time of the latest logged event.
	LastTimestamp *time.Time

	// READ-ONLY; The event message.
	Message *string

	// READ-ONLY; The event name.
	Name *string

	// READ-ONLY; The event type.
	Type *string
}

// GitRepoVolume - Represents a volume that is populated with the contents of a git repository
type GitRepoVolume struct {
	// REQUIRED; Repository URL
	Repository *string

	// Target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository.
	// Otherwise, if specified, the volume will contain the git repository in
	// the subdirectory with the given name.
	Directory *string

	// Commit hash for the specified revision.
	Revision *string
}

// GpuResource - The GPU resource.
type GpuResource struct {
	// REQUIRED; The count of the GPU resource.
	Count *int32

	// REQUIRED; The SKU of the GPU resource.
	SKU *GpuSKU
}

// HTTPHeader - The HTTP header.
type HTTPHeader struct {
	// The header name.
	Name *string

	// The header value.
	Value *string
}

// IPAddress - IP address for the container group.
type IPAddress struct {
	// REQUIRED; The list of ports exposed on the container group.
	Ports []*Port

	// REQUIRED; Specifies if the IP is exposed to the public internet or private VNET.
	Type *ContainerGroupIPAddressType

	// The value representing the security enum. The 'Unsecure' value is the default value if not selected and means the object's
	// domain name label is not secured against subdomain takeover. The
	// 'TenantReuse' value is the default value if selected and means the object's domain name label can be reused within the
	// same tenant. The 'SubscriptionReuse' value means the object's domain name label
	// can be reused within the same subscription. The 'ResourceGroupReuse' value means the object's domain name label can be
	// reused within the same resource group. The 'NoReuse' value means the object's
	// domain name label cannot be reused within the same resource group, subscription, or tenant.
	AutoGeneratedDomainNameLabelScope *DNSNameLabelReusePolicy

	// The Dns name label for the IP.
	DNSNameLabel *string

	// The IP exposed to the public internet.
	IP *string

	// READ-ONLY; The FQDN for the IP.
	Fqdn *string
}

// ImageRegistryCredential - Image registry credential.
type ImageRegistryCredential struct {
	// REQUIRED; The Docker image registry server without a protocol such as "http" and "https".
	Server *string

	// The identity for the private registry.
	Identity *string

	// The identity URL for the private registry.
	IdentityURL *string

	// The password for the private registry.
	Password *string

	// The username for the private registry.
	Username *string
}

// InitContainerDefinition - The init container definition.
type InitContainerDefinition struct {
	// REQUIRED; The name for the init container.
	Name *string

	// REQUIRED; The properties for the init container.
	Properties *InitContainerPropertiesDefinition
}

// InitContainerPropertiesDefinition - The init container definition properties.
type InitContainerPropertiesDefinition struct {
	// The command to execute within the init container in exec form.
	Command []*string

	// The environment variables to set in the init container.
	EnvironmentVariables []*EnvironmentVariable

	// The image of the init container.
	Image *string

	// The volume mounts available to the init container.
	VolumeMounts []*VolumeMount

	// READ-ONLY; The instance view of the init container. Only valid in response.
	InstanceView *InitContainerPropertiesDefinitionInstanceView
}

// InitContainerPropertiesDefinitionInstanceView - The instance view of the init container. Only valid in response.
type InitContainerPropertiesDefinitionInstanceView struct {
	// READ-ONLY; The current state of the init container.
	CurrentState *ContainerState

	// READ-ONLY; The events of the init container.
	Events []*Event

	// READ-ONLY; The previous state of the init container.
	PreviousState *ContainerState

	// READ-ONLY; The number of times that the init container has been restarted.
	RestartCount *int32
}

// LocationClientListCachedImagesOptions contains the optional parameters for the LocationClient.NewListCachedImagesPager
// method.
type LocationClientListCachedImagesOptions struct {
	// placeholder for future optional parameters
}

// LocationClientListCapabilitiesOptions contains the optional parameters for the LocationClient.NewListCapabilitiesPager
// method.
type LocationClientListCapabilitiesOptions struct {
	// placeholder for future optional parameters
}

// LocationClientListUsageOptions contains the optional parameters for the LocationClient.NewListUsagePager method.
type LocationClientListUsageOptions struct {
	// placeholder for future optional parameters
}

// LogAnalytics - Container group log analytics information.
type LogAnalytics struct {
	// REQUIRED; The workspace id for log analytics
	WorkspaceID *string

	// REQUIRED; The workspace key for log analytics
	WorkspaceKey *string

	// The log type to be used.
	LogType *LogAnalyticsLogType

	// Metadata for log analytics.
	Metadata map[string]*string

	// The workspace resource id for log analytics
	WorkspaceResourceID *string
}

// Logs - The logs.
type Logs struct {
	// The content of the log.
	Content *string
}

// Operation - An operation for Azure Container Instance service.
type Operation struct {
	// REQUIRED; The display information of the operation.
	Display *OperationDisplay

	// REQUIRED; The name of the operation.
	Name *string

	// The intended executor of the operation.
	Origin *ContainerInstanceOperationsOrigin

	// The additional properties.
	Properties any
}

// OperationDisplay - The display information of the operation.
type OperationDisplay struct {
	// The description of the operation.
	Description *string

	// The friendly name of the operation.
	Operation *string

	// The name of the provider of the operation.
	Provider *string

	// The name of the resource type of the operation.
	Resource *string
}

// OperationListResult - The operation list response that contains all operations for Azure Container Instance service.
type OperationListResult struct {
	// The URI to fetch the next page of operations.
	NextLink *string

	// The list of operations.
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Port - The port exposed on the container group.
type Port struct {
	// REQUIRED; The port number.
	Port *int32

	// The protocol associated with the port.
	Protocol *ContainerGroupNetworkProtocol
}

// Resource - The Resource model definition.
type Resource struct {
	// The resource location.
	Location *string

	// The resource tags.
	Tags map[string]*string

	// The zones for the container group.
	Zones []*string

	// READ-ONLY; The resource id.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// ResourceLimits - The resource limits.
type ResourceLimits struct {
	// The CPU limit of this container instance.
	CPU *float64

	// The GPU limit of this container instance.
	Gpu *GpuResource

	// The memory limit in GB of this container instance.
	MemoryInGB *float64
}

// ResourceRequests - The resource requests.
type ResourceRequests struct {
	// REQUIRED; The CPU request of this container instance.
	CPU *float64

	// REQUIRED; The memory request in GB of this container instance.
	MemoryInGB *float64

	// The GPU request of this container instance.
	Gpu *GpuResource
}

// ResourceRequirements - The resource requirements.
type ResourceRequirements struct {
	// REQUIRED; The resource requests of this container instance.
	Requests *ResourceRequests

	// The resource limits of this container instance.
	Limits *ResourceLimits
}

// SubnetServiceAssociationLinkClientBeginDeleteOptions contains the optional parameters for the SubnetServiceAssociationLinkClient.BeginDelete
// method.
type SubnetServiceAssociationLinkClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// Usage - A single usage result
type Usage struct {
	// READ-ONLY; The current usage of the resource
	CurrentValue *int32

	// READ-ONLY; Id of the usage result
	ID *string

	// READ-ONLY; The maximum permitted usage of the resource.
	Limit *int32

	// READ-ONLY; The name object of the resource
	Name *UsageName

	// READ-ONLY; Unit of the usage result
	Unit *string
}

// UsageListResult - The response containing the usage data
type UsageListResult struct {
	// READ-ONLY; The usage data.
	Value []*Usage
}

// UsageName - The name object of the resource
type UsageName struct {
	// READ-ONLY; The localized name of the resource
	LocalizedValue *string

	// READ-ONLY; The name of the resource
	Value *string
}

// UserAssignedIdentities - The list of user identities associated with the container group. The user identity dictionary
// key references will be ARM resource ids in the form:
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
type UserAssignedIdentities struct {
	// READ-ONLY; The client id of user assigned identity.
	ClientID *string

	// READ-ONLY; The principal id of user assigned identity.
	PrincipalID *string
}

// Volume - The properties of the volume.
type Volume struct {
	// REQUIRED; The name of the volume.
	Name *string

	// The Azure File volume.
	AzureFile *AzureFileVolume

	// The empty directory volume.
	EmptyDir any

	// The git repo volume.
	GitRepo *GitRepoVolume

	// The secret volume.
	Secret map[string]*string
}

// VolumeMount - The properties of the volume mount.
type VolumeMount struct {
	// REQUIRED; The path within the container where the volume should be mounted. Must not contain colon (:).
	MountPath *string

	// REQUIRED; The name of the volume mount.
	Name *string

	// The flag indicating whether the volume mount is read-only.
	ReadOnly *bool
}
