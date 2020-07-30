package containerinstance

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/containerinstance/mgmt/2017-10-01-preview/containerinstance"

// AzureFileVolume the properties of the Azure File volume. Azure File shares are mounted as volumes.
type AzureFileVolume struct {
	// ShareName - The name of the Azure File share to be mounted as a volume.
	ShareName *string `json:"shareName,omitempty"`
	// ReadOnly - The flag indicating whether the Azure File shared mounted as a volume is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// StorageAccountName - The name of the storage account that contains the Azure File share.
	StorageAccountName *string `json:"storageAccountName,omitempty"`
	// StorageAccountKey - The storage account access key used to access the Azure File share.
	StorageAccountKey *string `json:"storageAccountKey,omitempty"`
}

// Container a container instance.
type Container struct {
	// Name - The user-provided name of the container instance.
	Name *string `json:"name,omitempty"`
	// ContainerProperties - The properties of the container instance.
	*ContainerProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Container.
func (c Container) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.Name != nil {
		objectMap["name"] = c.Name
	}
	if c.ContainerProperties != nil {
		objectMap["properties"] = c.ContainerProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Container struct.
func (c *Container) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				c.Name = &name
			}
		case "properties":
			if v != nil {
				var containerProperties ContainerProperties
				err = json.Unmarshal(*v, &containerProperties)
				if err != nil {
					return err
				}
				c.ContainerProperties = &containerProperties
			}
		}
	}

	return nil
}

// ContainerGroup a container group.
type ContainerGroup struct {
	autorest.Response         `json:"-"`
	*ContainerGroupProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The resource id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type.
	Type *string `json:"type,omitempty"`
	// Location - The resource location.
	Location *string `json:"location,omitempty"`
	// Tags - The resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ContainerGroup.
func (cg ContainerGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cg.ContainerGroupProperties != nil {
		objectMap["properties"] = cg.ContainerGroupProperties
	}
	if cg.Location != nil {
		objectMap["location"] = cg.Location
	}
	if cg.Tags != nil {
		objectMap["tags"] = cg.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ContainerGroup struct.
func (cg *ContainerGroup) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var containerGroupProperties ContainerGroupProperties
				err = json.Unmarshal(*v, &containerGroupProperties)
				if err != nil {
					return err
				}
				cg.ContainerGroupProperties = &containerGroupProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cg.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cg.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cg.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				cg.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				cg.Tags = tags
			}
		}
	}

	return nil
}

// ContainerGroupListResult the container group list response that contains the container group properties.
type ContainerGroupListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of container groups.
	Value *[]ContainerGroup `json:"value,omitempty"`
	// NextLink - The URI to fetch the next page of container groups.
	NextLink *string `json:"nextLink,omitempty"`
}

// ContainerGroupListResultIterator provides access to a complete listing of ContainerGroup values.
type ContainerGroupListResultIterator struct {
	i    int
	page ContainerGroupListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ContainerGroupListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerGroupListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ContainerGroupListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ContainerGroupListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ContainerGroupListResultIterator) Response() ContainerGroupListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ContainerGroupListResultIterator) Value() ContainerGroup {
	if !iter.page.NotDone() {
		return ContainerGroup{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ContainerGroupListResultIterator type.
func NewContainerGroupListResultIterator(page ContainerGroupListResultPage) ContainerGroupListResultIterator {
	return ContainerGroupListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (cglr ContainerGroupListResult) IsEmpty() bool {
	return cglr.Value == nil || len(*cglr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (cglr ContainerGroupListResult) hasNextLink() bool {
	return cglr.NextLink != nil && len(*cglr.NextLink) != 0
}

// containerGroupListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (cglr ContainerGroupListResult) containerGroupListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !cglr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(cglr.NextLink)))
}

// ContainerGroupListResultPage contains a page of ContainerGroup values.
type ContainerGroupListResultPage struct {
	fn   func(context.Context, ContainerGroupListResult) (ContainerGroupListResult, error)
	cglr ContainerGroupListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ContainerGroupListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerGroupListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.cglr)
		if err != nil {
			return err
		}
		page.cglr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ContainerGroupListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ContainerGroupListResultPage) NotDone() bool {
	return !page.cglr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ContainerGroupListResultPage) Response() ContainerGroupListResult {
	return page.cglr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ContainerGroupListResultPage) Values() []ContainerGroup {
	if page.cglr.IsEmpty() {
		return nil
	}
	return *page.cglr.Value
}

// Creates a new instance of the ContainerGroupListResultPage type.
func NewContainerGroupListResultPage(getNextPage func(context.Context, ContainerGroupListResult) (ContainerGroupListResult, error)) ContainerGroupListResultPage {
	return ContainerGroupListResultPage{fn: getNextPage}
}

// ContainerGroupProperties ...
type ContainerGroupProperties struct {
	// ProvisioningState - READ-ONLY; The provisioning state of the container group. This only appears in the response.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// Containers - The containers within the container group.
	Containers *[]Container `json:"containers,omitempty"`
	// ImageRegistryCredentials - The image registry credentials by which the container group is created from.
	ImageRegistryCredentials *[]ImageRegistryCredential `json:"imageRegistryCredentials,omitempty"`
	// RestartPolicy - Restart policy for all containers within the container group.
	// - `Always` Always restart
	// - `OnFailure` Restart on failure
	// - `Never` Never restart
	// . Possible values include: 'Always', 'OnFailure', 'Never'
	RestartPolicy ContainerGroupRestartPolicy `json:"restartPolicy,omitempty"`
	// IPAddress - The IP address type of the container group.
	IPAddress *IPAddress `json:"ipAddress,omitempty"`
	// OsType - The operating system type required by the containers in the container group. Possible values include: 'Windows', 'Linux'
	OsType OperatingSystemTypes `json:"osType,omitempty"`
	// Volumes - The list of volumes that can be mounted by containers in this container group.
	Volumes *[]Volume `json:"volumes,omitempty"`
	// InstanceView - READ-ONLY; The instance view of the container group. Only valid in response.
	InstanceView *ContainerGroupPropertiesInstanceView `json:"instanceView,omitempty"`
}

// MarshalJSON is the custom marshaler for ContainerGroupProperties.
func (cg ContainerGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cg.Containers != nil {
		objectMap["containers"] = cg.Containers
	}
	if cg.ImageRegistryCredentials != nil {
		objectMap["imageRegistryCredentials"] = cg.ImageRegistryCredentials
	}
	if cg.RestartPolicy != "" {
		objectMap["restartPolicy"] = cg.RestartPolicy
	}
	if cg.IPAddress != nil {
		objectMap["ipAddress"] = cg.IPAddress
	}
	if cg.OsType != "" {
		objectMap["osType"] = cg.OsType
	}
	if cg.Volumes != nil {
		objectMap["volumes"] = cg.Volumes
	}
	return json.Marshal(objectMap)
}

// ContainerGroupPropertiesInstanceView the instance view of the container group. Only valid in response.
type ContainerGroupPropertiesInstanceView struct {
	// Events - READ-ONLY; The events of this container group.
	Events *[]Event `json:"events,omitempty"`
	// State - READ-ONLY; The state of the container group. Only valid in response.
	State *string `json:"state,omitempty"`
}

// ContainerPort the port exposed on the container instance.
type ContainerPort struct {
	// Protocol - The protocol associated with the port. Possible values include: 'ContainerNetworkProtocolTCP', 'ContainerNetworkProtocolUDP'
	Protocol ContainerNetworkProtocol `json:"protocol,omitempty"`
	// Port - The port number exposed within the container group.
	Port *int32 `json:"port,omitempty"`
}

// ContainerProperties the container instance properties.
type ContainerProperties struct {
	// Image - The name of the image used to create the container instance.
	Image *string `json:"image,omitempty"`
	// Command - The commands to execute within the container instance in exec form.
	Command *[]string `json:"command,omitempty"`
	// Ports - The exposed ports on the container instance.
	Ports *[]ContainerPort `json:"ports,omitempty"`
	// EnvironmentVariables - The environment variables to set in the container instance.
	EnvironmentVariables *[]EnvironmentVariable `json:"environmentVariables,omitempty"`
	// InstanceView - READ-ONLY; The instance view of the container instance. Only valid in response.
	InstanceView *ContainerPropertiesInstanceView `json:"instanceView,omitempty"`
	// Resources - The resource requirements of the container instance.
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// VolumeMounts - The volume mounts available to the container instance.
	VolumeMounts *[]VolumeMount `json:"volumeMounts,omitempty"`
}

// MarshalJSON is the custom marshaler for ContainerProperties.
func (cp ContainerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.Image != nil {
		objectMap["image"] = cp.Image
	}
	if cp.Command != nil {
		objectMap["command"] = cp.Command
	}
	if cp.Ports != nil {
		objectMap["ports"] = cp.Ports
	}
	if cp.EnvironmentVariables != nil {
		objectMap["environmentVariables"] = cp.EnvironmentVariables
	}
	if cp.Resources != nil {
		objectMap["resources"] = cp.Resources
	}
	if cp.VolumeMounts != nil {
		objectMap["volumeMounts"] = cp.VolumeMounts
	}
	return json.Marshal(objectMap)
}

// ContainerPropertiesInstanceView the instance view of the container instance. Only valid in response.
type ContainerPropertiesInstanceView struct {
	// RestartCount - READ-ONLY; The number of times that the container instance has been restarted.
	RestartCount *int32 `json:"restartCount,omitempty"`
	// CurrentState - READ-ONLY; Current container instance state.
	CurrentState *ContainerState `json:"currentState,omitempty"`
	// PreviousState - READ-ONLY; Previous container instance state.
	PreviousState *ContainerState `json:"previousState,omitempty"`
	// Events - READ-ONLY; The events of the container instance.
	Events *[]Event `json:"events,omitempty"`
}

// ContainerState the container instance state.
type ContainerState struct {
	// State - The state of the container instance.
	State *string `json:"state,omitempty"`
	// StartTime - The date-time when the container instance state started.
	StartTime *date.Time `json:"startTime,omitempty"`
	// ExitCode - The container instance exit codes correspond to those from the `docker run` command.
	ExitCode *int32 `json:"exitCode,omitempty"`
	// FinishTime - The date-time when the container instance state finished.
	FinishTime *date.Time `json:"finishTime,omitempty"`
	// DetailStatus - The human-readable status of the container instance state.
	DetailStatus *string `json:"detailStatus,omitempty"`
}

// EnvironmentVariable the environment variable to set within the container instance.
type EnvironmentVariable struct {
	// Name - The name of the environment variable.
	Name *string `json:"name,omitempty"`
	// Value - The value of the environment variable.
	Value *string `json:"value,omitempty"`
}

// Event a container group or container instance event.
type Event struct {
	// Count - The count of the event.
	Count *int32 `json:"count,omitempty"`
	// FirstTimestamp - The date-time of the earliest logged event.
	FirstTimestamp *date.Time `json:"firstTimestamp,omitempty"`
	// LastTimestamp - The date-time of the latest logged event.
	LastTimestamp *date.Time `json:"lastTimestamp,omitempty"`
	// Name - The event name.
	Name *string `json:"name,omitempty"`
	// Message - The event message.
	Message *string `json:"message,omitempty"`
	// Type - The event type.
	Type *string `json:"type,omitempty"`
}

// ImageRegistryCredential image registry credential.
type ImageRegistryCredential struct {
	// Server - The Docker image registry server without a protocol such as "http" and "https".
	Server *string `json:"server,omitempty"`
	// Username - The username for the private registry.
	Username *string `json:"username,omitempty"`
	// Password - The password for the private registry.
	Password *string `json:"password,omitempty"`
}

// IPAddress IP address for the container group.
type IPAddress struct {
	// Ports - The list of ports exposed on the container group.
	Ports *[]Port `json:"ports,omitempty"`
	// Type - Specifies if the IP is exposed to the public internet.
	Type *string `json:"type,omitempty"`
	// IP - The IP exposed to the public internet.
	IP *string `json:"ip,omitempty"`
}

// Logs the logs.
type Logs struct {
	autorest.Response `json:"-"`
	// Content - The content of the log.
	Content *string `json:"content,omitempty"`
}

// Operation an operation for Azure Container Instance service.
type Operation struct {
	// Name - The name of the operation.
	Name *string `json:"name,omitempty"`
	// Display - The display information of the operation.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - The intended executor of the operation. Possible values include: 'User', 'System'
	Origin OperationsOrigin `json:"origin,omitempty"`
}

// OperationDisplay the display information of the operation.
type OperationDisplay struct {
	// Provider - The name of the provider of the operation.
	Provider *string `json:"provider,omitempty"`
	// Resource - The name of the resource type of the operation.
	Resource *string `json:"resource,omitempty"`
	// Operation - The friendly name of the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult the operation list response that contains all operations for Azure Container Instance
// service.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of operations.
	Value *[]Operation `json:"value,omitempty"`
}

// Port the port exposed on the container group.
type Port struct {
	// Protocol - The protocol associated with the port. Possible values include: 'TCP', 'UDP'
	Protocol ContainerGroupNetworkProtocol `json:"protocol,omitempty"`
	// Port - The port number.
	Port *int32 `json:"port,omitempty"`
}

// Resource the Resource model definition.
type Resource struct {
	// ID - READ-ONLY; The resource id.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type.
	Type *string `json:"type,omitempty"`
	// Location - The resource location.
	Location *string `json:"location,omitempty"`
	// Tags - The resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceLimits the resource limits.
type ResourceLimits struct {
	// MemoryInGB - The memory limit in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
	// CPU - The CPU limit of this container instance.
	CPU *float64 `json:"cpu,omitempty"`
}

// ResourceRequests the resource requests.
type ResourceRequests struct {
	// MemoryInGB - The memory request in GB of this container instance.
	MemoryInGB *float64 `json:"memoryInGB,omitempty"`
	// CPU - The CPU request of this container instance.
	CPU *float64 `json:"cpu,omitempty"`
}

// ResourceRequirements the resource requirements.
type ResourceRequirements struct {
	// Requests - The resource requests of this container instance.
	Requests *ResourceRequests `json:"requests,omitempty"`
	// Limits - The resource limits of this container instance.
	Limits *ResourceLimits `json:"limits,omitempty"`
}

// Volume the properties of the volume.
type Volume struct {
	// Name - The name of the volume.
	Name *string `json:"name,omitempty"`
	// AzureFile - The name of the Azure File volume.
	AzureFile *AzureFileVolume `json:"azureFile,omitempty"`
	// EmptyDir - The empty directory volume.
	EmptyDir interface{} `json:"emptyDir,omitempty"`
}

// VolumeMount the properties of the volume mount.
type VolumeMount struct {
	// Name - The name of the volume mount.
	Name *string `json:"name,omitempty"`
	// MountPath - The path within the container where the volume should be mounted. Must not contain colon (:).
	MountPath *string `json:"mountPath,omitempty"`
	// ReadOnly - The flag indicating whether the volume mount is read-only.
	ReadOnly *bool `json:"readOnly,omitempty"`
}
