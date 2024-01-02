//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

// CheckServerNameAvailabilityParameters - Details of server name request body.
type CheckServerNameAvailabilityParameters struct {
	// Name for checking availability.
	Name *string

	// The resource type of azure analysis services.
	Type *string
}

// CheckServerNameAvailabilityResult - The checking result of server name availability.
type CheckServerNameAvailabilityResult struct {
	// The detailed message of the request unavailability.
	Message *string

	// Indicator of available of the server name.
	NameAvailable *bool

	// The reason of unavailability.
	Reason *string
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

	// READ-ONLY; The http status code
	HTTPStatusCode *int32

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error sub code
	SubCode *int32

	// READ-ONLY; The error target.
	Target *string

	// READ-ONLY; the timestamp for the error.
	TimeStamp *string
}

// ErrorResponse - Describes the format of Error response.
type ErrorResponse struct {
	// The error object
	Error *ErrorDetail
}

// GatewayDetails - The gateway details.
type GatewayDetails struct {
	// Gateway resource to be associated with the server.
	GatewayResourceID *string

	// READ-ONLY; Uri of the DMTS cluster.
	DmtsClusterURI *string

	// READ-ONLY; Gateway object id from in the DMTS cluster for the gateway resource.
	GatewayObjectID *string
}

// GatewayListStatusError - Status of gateway is error.
type GatewayListStatusError struct {
	// Error of the list gateway status.
	Error *ErrorDetail
}

// GatewayListStatusLive - Status of gateway is live.
type GatewayListStatusLive struct {
	// Live message of list gateway. Status: 0 - Live
	Status *int32
}

// IPv4FirewallRule - The detail of firewall rule.
type IPv4FirewallRule struct {
	// The rule name.
	FirewallRuleName *string

	// The end range of IPv4.
	RangeEnd *string

	// The start range of IPv4.
	RangeStart *string
}

// IPv4FirewallSettings - An array of firewall rules.
type IPv4FirewallSettings struct {
	// The indicator of enabling PBI service.
	EnablePowerBIService *bool

	// An array of firewall rules.
	FirewallRules []*IPv4FirewallRule
}

// LogSpecifications - The log metric specification for exposing performance metrics to shoebox.
type LogSpecifications struct {
	// READ-ONLY; The blob duration for the log.
	BlobDuration *string

	// READ-ONLY; The displayed name of log.
	DisplayName *string

	// READ-ONLY; The name of metric.
	Name *string
}

// MetricDimensions - Metric dimension.
type MetricDimensions struct {
	// READ-ONLY; Dimension display name.
	DisplayName *string

	// READ-ONLY; Dimension name.
	Name *string
}

// MetricSpecifications - Available operation metric specification for exposing performance metrics to shoebox.
type MetricSpecifications struct {
	// READ-ONLY; The aggregation type of metric.
	AggregationType *string

	// READ-ONLY; The dimensions of metric.
	Dimensions []*MetricDimensions

	// READ-ONLY; The displayed description of metric.
	DisplayDescription *string

	// READ-ONLY; The displayed name of metric.
	DisplayName *string

	// READ-ONLY; The name of metric.
	Name *string

	// READ-ONLY; The unit of the metric.
	Unit *string
}

// Operation - A Consumption REST API operation.
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Additional properties to expose performance metrics to shoebox.
	Properties *OperationProperties

	// READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string

	// READ-ONLY; The origin
	Origin *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// READ-ONLY; Description of the operation object.
	Description *string

	// READ-ONLY; Operation type: Read, write, delete, etc.
	Operation *string

	// READ-ONLY; Service provider: Microsoft.Consumption.
	Provider *string

	// READ-ONLY; Resource on which the operation is performed: UsageDetail, etc.
	Resource *string
}

// OperationListResult - Result of listing consumption operations. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string

	// READ-ONLY; List of analysis services operations supported by the Microsoft.AnalysisServices resource provider.
	Value []*Operation
}

// OperationProperties - Additional properties to expose performance metrics to shoebox.
type OperationProperties struct {
	// Performance metrics to shoebox.
	ServiceSpecification *OperationPropertiesServiceSpecification
}

// OperationPropertiesServiceSpecification - Performance metrics to shoebox.
type OperationPropertiesServiceSpecification struct {
	// READ-ONLY; The log specifications.
	LogSpecifications []*LogSpecifications

	// READ-ONLY; The metric specifications.
	MetricSpecifications []*MetricSpecifications
}

// OperationStatus - The status of operation.
type OperationStatus struct {
	// The end time of the operation.
	EndTime *string

	// The error detail of the operation if any.
	Error *ErrorDetail

	// The operation Id.
	ID *string

	// The operation name.
	Name *string

	// The start time of the operation.
	StartTime *string

	// The status of the operation.
	Status *string
}

// Resource - Represents an instance of an Analysis Services resource.
type Resource struct {
	// REQUIRED; Location of the Analysis Services resource.
	Location *string

	// REQUIRED; The SKU of the Analysis Services resource.
	SKU *ResourceSKU

	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the Analysis Services resource.
	ID *string

	// READ-ONLY; The name of the Analysis Services resource.
	Name *string

	// READ-ONLY; The type of the Analysis Services resource.
	Type *string
}

// ResourceSKU - Represents the SKU name and Azure pricing tier for Analysis Services resource.
type ResourceSKU struct {
	// REQUIRED; Name of the SKU level.
	Name *string

	// The number of instances in the read only query pool.
	Capacity *int32

	// The name of the Azure pricing tier to which the SKU applies.
	Tier *SKUTier
}

// SKUDetailsForExistingResource - An object that represents SKU details for existing resources.
type SKUDetailsForExistingResource struct {
	// The resource type.
	ResourceType *string

	// The SKU in SKU details for existing resources.
	SKU *ResourceSKU
}

// SKUEnumerationForExistingResourceResult - An object that represents enumerating SKUs for existing resources.
type SKUEnumerationForExistingResourceResult struct {
	// The collection of available SKUs for existing resources.
	Value []*SKUDetailsForExistingResource
}

// SKUEnumerationForNewResourceResult - An object that represents enumerating SKUs for new resources.
type SKUEnumerationForNewResourceResult struct {
	// The collection of available SKUs for new resources.
	Value []*ResourceSKU
}

// Server - Represents an instance of an Analysis Services resource.
type Server struct {
	// REQUIRED; Location of the Analysis Services resource.
	Location *string

	// REQUIRED; The SKU of the Analysis Services resource.
	SKU *ResourceSKU

	// Properties of the provision operation request.
	Properties *ServerProperties

	// Key-value pairs of additional resource provisioning properties.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the Analysis Services resource.
	ID *string

	// READ-ONLY; The name of the Analysis Services resource.
	Name *string

	// READ-ONLY; The type of the Analysis Services resource.
	Type *string
}

// ServerAdministrators - An array of administrator user identities.
type ServerAdministrators struct {
	// An array of administrator user identities.
	Members []*string
}

// ServerMutableProperties - An object that represents a set of mutable Analysis Services resource properties.
type ServerMutableProperties struct {
	// A collection of AS server administrators
	AsAdministrators *ServerAdministrators

	// The SAS container URI to the backup container.
	BackupBlobContainerURI *string

	// The gateway details configured for the AS server.
	GatewayDetails *GatewayDetails

	// The firewall settings for the AS server.
	IPV4FirewallSettings *IPv4FirewallSettings

	// The managed mode of the server (0 = not managed, 1 = managed).
	ManagedMode *ManagedMode

	// How the read-write server's participation in the query pool is controlled.
	// It can have the following values: * readOnly - indicates that the read-write server is intended not to participate in query
	// operations
	// * all - indicates that the read-write server can participate in query operations
	// Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode *ConnectionMode

	// The server monitor mode for AS server
	ServerMonitorMode *ServerMonitorMode
}

// ServerProperties - Properties of Analysis Services resource.
type ServerProperties struct {
	// A collection of AS server administrators
	AsAdministrators *ServerAdministrators

	// The SAS container URI to the backup container.
	BackupBlobContainerURI *string

	// The gateway details configured for the AS server.
	GatewayDetails *GatewayDetails

	// The firewall settings for the AS server.
	IPV4FirewallSettings *IPv4FirewallSettings

	// The managed mode of the server (0 = not managed, 1 = managed).
	ManagedMode *ManagedMode

	// How the read-write server's participation in the query pool is controlled.
	// It can have the following values: * readOnly - indicates that the read-write server is intended not to participate in query
	// operations
	// * all - indicates that the read-write server can participate in query operations
	// Specifying readOnly when capacity is 1 results in error.
	QuerypoolConnectionMode *ConnectionMode

	// The SKU of the Analysis Services resource.
	SKU *ResourceSKU

	// The server monitor mode for AS server
	ServerMonitorMode *ServerMonitorMode

	// READ-ONLY; The current deployment state of Analysis Services resource. The provisioningState is to indicate states for
	// resource provisioning.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The full name of the Analysis Services resource.
	ServerFullName *string

	// READ-ONLY; The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
	State *State
}

// ServerUpdateParameters - Provision request specification
type ServerUpdateParameters struct {
	// Properties of the provision operation request.
	Properties *ServerMutableProperties

	// The SKU of the Analysis Services resource.
	SKU *ResourceSKU

	// Key-value pairs of additional provisioning properties.
	Tags map[string]*string
}

// Servers - An array of Analysis Services resources.
type Servers struct {
	// REQUIRED; An array of Analysis Services resources.
	Value []*Server
}
