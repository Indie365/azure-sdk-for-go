//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import "time"

// AvailabilityStatus - availabilityStatus of a resource.
type AvailabilityStatus struct {
	// Azure Resource Manager Identity for the availabilityStatuses resource.
	ID *string `json:"id,omitempty"`

	// Azure Resource Manager geo location of the resource.
	Location *string `json:"location,omitempty"`

	// current.
	Name *string `json:"name,omitempty"`

	// Properties of availability state.
	Properties *AvailabilityStatusProperties `json:"properties,omitempty"`

	// Microsoft.ResourceHealth/AvailabilityStatuses.
	Type *string `json:"type,omitempty"`
}

// AvailabilityStatusListResult - The List availabilityStatus operation response.
type AvailabilityStatusListResult struct {
	// REQUIRED; The list of availabilityStatuses.
	Value []*AvailabilityStatus `json:"value,omitempty"`

	// The URI to fetch the next page of availabilityStatuses. Call ListNext() with this URI to fetch the next page of availabilityStatuses.
	NextLink *string `json:"nextLink,omitempty"`
}

// AvailabilityStatusProperties - Properties of availability state.
type AvailabilityStatusProperties struct {
	// Availability status of the resource. When it is null, this availabilityStatus object represents an availability impacting
	// event
	AvailabilityState *AvailabilityStateValues `json:"availabilityState,omitempty"`

	// Details of the availability status.
	DetailedStatus *string `json:"detailedStatus,omitempty"`

	// In case of an availability impacting event, it describes the category of a PlatformInitiated health impacting event. Examples
	// are Planned, Unplanned etc.
	HealthEventCategory *string `json:"healthEventCategory,omitempty"`

	// In case of an availability impacting event, it describes where the health impacting event was originated. Examples are
	// PlatformInitiated, UserInitiated etc.
	HealthEventCause *string `json:"healthEventCause,omitempty"`

	// It is a unique Id that identifies the event
	HealthEventID *string `json:"healthEventId,omitempty"`

	// In case of an availability impacting event, it describes when the health impacting event was originated. Examples are Lifecycle,
	// Downtime, Fault Analysis etc.
	HealthEventType *string `json:"healthEventType,omitempty"`

	// Timestamp for when last change in health status occurred.
	OccurredTime *time.Time `json:"occurredTime,omitempty"`

	// Chronicity of the availability transition.
	ReasonChronicity *ReasonChronicityTypes `json:"reasonChronicity,omitempty"`

	// When the resource's availabilityState is Unavailable, it describes where the health impacting event was originated. Examples
	// are planned, unplanned, user initiated or an outage etc.
	ReasonType *string `json:"reasonType,omitempty"`

	// An annotation describing a change in the availabilityState to Available from Unavailable with a reasonType of type Unplanned
	RecentlyResolved *AvailabilityStatusPropertiesRecentlyResolved `json:"recentlyResolved,omitempty"`

	// Lists actions the user can take based on the current availabilityState of the resource.
	RecommendedActions []*RecommendedAction `json:"recommendedActions,omitempty"`

	// Timestamp for when the health was last checked.
	ReportedTime *time.Time `json:"reportedTime,omitempty"`

	// When the resource's availabilityState is Unavailable and the reasonType is not User Initiated, it provides the date and
	// time for when the issue is expected to be resolved.
	ResolutionETA *time.Time `json:"resolutionETA,omitempty"`

	// When the resource's availabilityState is Unavailable, it provides the Timestamp for when the health impacting event was
	// received.
	RootCauseAttributionTime *time.Time `json:"rootCauseAttributionTime,omitempty"`

	// Lists the service impacting events that may be affecting the health of the resource.
	ServiceImpactingEvents []*ServiceImpactingEvent `json:"serviceImpactingEvents,omitempty"`

	// Summary description of the availability status.
	Summary *string `json:"summary,omitempty"`

	// Title description of the availability status.
	Title *string `json:"title,omitempty"`
}

// AvailabilityStatusPropertiesRecentlyResolved - An annotation describing a change in the availabilityState to Available
// from Unavailable with a reasonType of type Unplanned
type AvailabilityStatusPropertiesRecentlyResolved struct {
	// Timestamp when the availabilityState changes to Available.
	ResolvedTime *time.Time `json:"resolvedTime,omitempty"`

	// Brief description of cause of the resource becoming unavailable.
	UnavailabilitySummary *string `json:"unavailabilitySummary,omitempty"`

	// Timestamp for when the availabilityState changed to Unavailable
	UnavailableOccurredTime *time.Time `json:"unavailableOccurredTime,omitempty"`
}

// AvailabilityStatusesClientGetByResourceOptions contains the optional parameters for the AvailabilityStatusesClient.GetByResource
// method.
type AvailabilityStatusesClientGetByResourceOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListByResourceGroupOptions contains the optional parameters for the AvailabilityStatusesClient.ListByResourceGroup
// method.
type AvailabilityStatusesClientListByResourceGroupOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListBySubscriptionIDOptions contains the optional parameters for the AvailabilityStatusesClient.ListBySubscriptionID
// method.
type AvailabilityStatusesClientListBySubscriptionIDOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// AvailabilityStatusesClientListOptions contains the optional parameters for the AvailabilityStatusesClient.List method.
type AvailabilityStatusesClientListOptions struct {
	// Setting $expand=recommendedactions in url query expands the recommendedactions in the response.
	Expand *string
	// The filter to apply on the operation. For more information please see https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
	Filter *string
}

// ErrorResponse - Error details.
type ErrorResponse struct {
	// The error object.
	Error *ErrorResponseError `json:"error,omitempty"`
}

// ErrorResponseError - The error object.
type ErrorResponseError struct {
	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details *string `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ImpactedRegion - Object of impacted region.
type ImpactedRegion struct {
	// The impacted region id.
	ID *string `json:"id,omitempty"`

	// The impacted region name.
	Name *string `json:"name,omitempty"`
}

// ImpactedResourceStatus - impactedResource with health status
type ImpactedResourceStatus struct {
	// Properties of impacted resource status.
	Properties *ImpactedResourceStatusProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ImpactedResourceStatusProperties - Properties of impacted resource status.
type ImpactedResourceStatusProperties struct {
	// Impacted resource status of the resource.
	AvailabilityState *AvailabilityStateValues `json:"availabilityState,omitempty"`

	// Timestamp for when last change in health status occurred.
	OccurredTime *time.Time `json:"occurredTime,omitempty"`

	// When the resource's availabilityState is Unavailable, it describes where the health impacting event was originated.
	ReasonType *ReasonTypeValues `json:"reasonType,omitempty"`

	// Summary description of the impacted resource status.
	Summary *string `json:"summary,omitempty"`

	// Title description of the impacted resource status.
	Title *string `json:"title,omitempty"`
}

// Operation available in the resourcehealth resource provider.
type Operation struct {
	// Properties of the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Name of the operation.
	Name *string `json:"name,omitempty"`
}

// OperationDisplay - Properties of the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string `json:"description,omitempty"`

	// Operation name.
	Operation *string `json:"operation,omitempty"`

	// Provider name.
	Provider *string `json:"provider,omitempty"`

	// Resource name.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Lists the operations response.
type OperationListResult struct {
	// REQUIRED; List of operations available in the resourcehealth resource provider.
	Value []*Operation `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// RecommendedAction - Lists actions the user can take based on the current availabilityState of the resource.
type RecommendedAction struct {
	// Recommended action.
	Action *string `json:"action,omitempty"`

	// Link to the action
	ActionURL *string `json:"actionUrl,omitempty"`

	// Substring of action, it describes which text should host the action url.
	ActionURLText *string `json:"actionUrlText,omitempty"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ServiceImpactingEvent - Lists the service impacting events that may be affecting the health of the resource.
type ServiceImpactingEvent struct {
	// Correlation id for the event
	CorrelationID *string `json:"correlationId,omitempty"`

	// Timestamp for when the event started.
	EventStartTime *time.Time `json:"eventStartTime,omitempty"`

	// Timestamp for when event was submitted/detected.
	EventStatusLastModifiedTime *time.Time `json:"eventStatusLastModifiedTime,omitempty"`

	// Properties of the service impacting event.
	IncidentProperties *ServiceImpactingEventIncidentProperties `json:"incidentProperties,omitempty"`

	// Status of the service impacting event.
	Status *ServiceImpactingEventStatus `json:"status,omitempty"`
}

// ServiceImpactingEventIncidentProperties - Properties of the service impacting event.
type ServiceImpactingEventIncidentProperties struct {
	// Type of Event.
	IncidentType *string `json:"incidentType,omitempty"`

	// Region impacted by the event.
	Region *string `json:"region,omitempty"`

	// Service impacted by the event.
	Service *string `json:"service,omitempty"`

	// Title of the incident.
	Title *string `json:"title,omitempty"`
}

// ServiceImpactingEventStatus - Status of the service impacting event.
type ServiceImpactingEventStatus struct {
	// Current status of the event
	Value *string `json:"value,omitempty"`
}

// StatusBanner - Banner type of emerging issue.
type StatusBanner struct {
	// The cloud type of this banner.
	Cloud *string `json:"cloud,omitempty"`

	// The last time modified on this banner.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`

	// The details of banner.
	Message *string `json:"message,omitempty"`

	// The banner title.
	Title *string `json:"title,omitempty"`
}
