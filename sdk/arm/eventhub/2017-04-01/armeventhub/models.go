// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Namespace/EventHub Connection String
type AccessKeys struct {
	// Primary connection string of the alias if GEO DR is enabled
	AliasPrimaryConnectionString *string `json:"aliasPrimaryConnectionString,omitempty"`

	// Secondary connection string of the alias if GEO DR is enabled
	AliasSecondaryConnectionString *string `json:"aliasSecondaryConnectionString,omitempty"`

	// A string that describes the AuthorizationRule.
	KeyName *string `json:"keyName,omitempty"`

	// Primary connection string of the created namespace AuthorizationRule.
	PrimaryConnectionString *string `json:"primaryConnectionString,omitempty"`

	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string `json:"primaryKey,omitempty"`

	// Secondary connection string of the created namespace AuthorizationRule.
	SecondaryConnectionString *string `json:"secondaryConnectionString,omitempty"`

	// A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// AccessKeysResponse is the response envelope for operations that return a AccessKeys type.
type AccessKeysResponse struct {
	// Namespace/EventHub Connection String
	AccessKeys *AccessKeys

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Single item in List or Get Alias(Disaster Recovery configuration) operation
type ArmDisasterRecovery struct {
	Resource
	// Properties required to the Create Or Update Alias(Disaster Recovery configurations)
	Properties *ArmDisasterRecoveryProperties `json:"properties,omitempty"`
}

// The result of the List Alias(Disaster Recovery configuration) operation.
type ArmDisasterRecoveryListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of Alias(Disaster Recovery configuration)
	NextLink *string `json:"nextLink,omitempty"`

	// List of Alias(Disaster Recovery configurations)
	Value *[]ArmDisasterRecovery `json:"value,omitempty"`
}

// ArmDisasterRecoveryListResultResponse is the response envelope for operations that return a ArmDisasterRecoveryListResult
// type.
type ArmDisasterRecoveryListResultResponse struct {
	// The result of the List Alias(Disaster Recovery configuration) operation.
	ArmDisasterRecoveryListResult *ArmDisasterRecoveryListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties required to the Create Or Update Alias(Disaster Recovery configurations)
type ArmDisasterRecoveryProperties struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string `json:"alternateName,omitempty"`

	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace *string `json:"partnerNamespace,omitempty"`

	// Number of entities pending to be replicated.
	PendingReplicationOperationsCount *int64 `json:"pendingReplicationOperationsCount,omitempty"`

	// Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
	ProvisioningState *ProvisioningStateDR `json:"provisioningState,omitempty"`

	// role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
	Role *RoleDisasterRecovery `json:"role,omitempty"`
}

// ArmDisasterRecoveryResponse is the response envelope for operations that return a ArmDisasterRecovery type.
type ArmDisasterRecoveryResponse struct {
	// Single item in List or Get Alias(Disaster Recovery configuration) operation
	ArmDisasterRecovery *ArmDisasterRecovery

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Single item in a List or Get AuthorizationRule operation
type AuthorizationRule struct {
	Resource
	// Properties supplied to create or update AuthorizationRule
	Properties *AuthorizationRuleProperties `json:"properties,omitempty"`
}

// The response from the List namespace operation.
type AuthorizationRuleListResult struct {
	// Link to the next set of results. Not empty if Value contains an incomplete list of Authorization Rules
	NextLink *string `json:"nextLink,omitempty"`

	// Result of the List Authorization Rules operation.
	Value *[]AuthorizationRule `json:"value,omitempty"`
}

// AuthorizationRuleListResultResponse is the response envelope for operations that return a AuthorizationRuleListResult type.
type AuthorizationRuleListResultResponse struct {
	// The response from the List namespace operation.
	AuthorizationRuleListResult *AuthorizationRuleListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties supplied to create or update AuthorizationRule
type AuthorizationRuleProperties struct {
	// The rights associated with the rule.
	Rights *[]AccessRights `json:"rights,omitempty"`
}

// AuthorizationRuleResponse is the response envelope for operations that return a AuthorizationRule type.
type AuthorizationRuleResponse struct {
	// Single item in a List or Get AuthorizationRule operation
	AuthorizationRule *AuthorizationRule

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties to configure capture description for eventhub
type CaptureDescription struct {
	// Properties of Destination where capture will be stored. (Storage Account, Blob Names)
	Destination *Destination `json:"destination,omitempty"`

	// A value that indicates whether capture description is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in
	// New API Version
	Encoding *EncodingCaptureDescription `json:"encoding,omitempty"`

	// The time window allows you to set the frequency with which the capture to Azure Blobs will happen, value should between
	// 60 to 900 seconds
	IntervalInSeconds *int32 `json:"intervalInSeconds,omitempty"`

	// The size window defines the amount of data built up in your Event Hub before an capture operation, value should be between
	// 10485760 to 524288000 bytes
	SizeLimitInBytes *int32 `json:"sizeLimitInBytes,omitempty"`

	// A value that indicates whether to Skip Empty Archives
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty"`
}

// Parameter supplied to check Namespace name availability operation
type CheckNameAvailabilityParameter struct {
	// Name to check the namespace name availability
	Name *string `json:"name,omitempty"`
}

// The Result of the CheckNameAvailability operation
type CheckNameAvailabilityResult struct {
	// The detailed info regarding the reason associated with the Namespace.
	Message *string `json:"message,omitempty"`

	// Value indicating Namespace is availability, true if the Namespace is available; otherwise, false.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// The reason for unavailability of a Namespace.
	Reason *UnavailableReason `json:"reason,omitempty"`
}

// CheckNameAvailabilityResultResponse is the response envelope for operations that return a CheckNameAvailabilityResult type.
type CheckNameAvailabilityResultResponse struct {
	// The Result of the CheckNameAvailability operation
	CheckNameAvailabilityResult *CheckNameAvailabilityResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Single item in List or Get Consumer group operation
type ConsumerGroup struct {
	Resource
	// Single item in List or Get Consumer group operation
	Properties *ConsumerGroupProperties `json:"properties,omitempty"`
}

// The result to the List Consumer Group operation.
type ConsumerGroupListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of Consumer Group
	NextLink *string `json:"nextLink,omitempty"`

	// Result of the List Consumer Group operation.
	Value *[]ConsumerGroup `json:"value,omitempty"`
}

// ConsumerGroupListResultResponse is the response envelope for operations that return a ConsumerGroupListResult type.
type ConsumerGroupListResultResponse struct {
	// The result to the List Consumer Group operation.
	ConsumerGroupListResult *ConsumerGroupListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Single item in List or Get Consumer group operation
type ConsumerGroupProperties struct {
	// Exact time the message was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The exact time the message was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	// User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be used to store
	// descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ConsumerGroupProperties.
func (c ConsumerGroupProperties) MarshalJSON() ([]byte, error) {
	type alias ConsumerGroupProperties
	aux := &struct {
		*alias
		CreatedAt *timeRFC3339 `json:"createdAt"`
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias:     (*alias)(&c),
		CreatedAt: (*timeRFC3339)(c.CreatedAt),
		UpdatedAt: (*timeRFC3339)(c.UpdatedAt),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConsumerGroupProperties.
func (c *ConsumerGroupProperties) UnmarshalJSON(data []byte) error {
	type alias ConsumerGroupProperties
	aux := &struct {
		*alias
		CreatedAt *timeRFC3339 `json:"createdAt"`
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	c.CreatedAt = (*time.Time)(aux.CreatedAt)
	c.UpdatedAt = (*time.Time)(aux.UpdatedAt)
	return nil
}

// ConsumerGroupResponse is the response envelope for operations that return a ConsumerGroup type.
type ConsumerGroupResponse struct {
	// Single item in List or Get Consumer group operation
	ConsumerGroup *ConsumerGroup

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConsumerGroupsListByEventHubOptions contains the optional parameters for the ConsumerGroups.ListByEventHub method.
type ConsumerGroupsListByEventHubOptions struct {
	// Skip is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skip parameter that specifies a starting point to use for subsequent calls.
	Skip *int32
	// May be used to limit the number of results to the most recent N usageDetails.
	Top *int32
}

// Capture storage details for capture description
type Destination struct {
	// Name for capture destination
	Name *string `json:"name,omitempty"`

	// Properties describing the storage account, blob container and archive name format for capture destination
	Properties *DestinationProperties `json:"properties,omitempty"`
}

// Properties describing the storage account, blob container and archive name format for capture destination
type DestinationProperties struct {
	// Blob naming convention for archive, e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}.
	// Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order
	ArchiveNameFormat *string `json:"archiveNameFormat,omitempty"`

	// Blob container Name
	BlobContainer *string `json:"blobContainer,omitempty"`

	// Resource id of the storage account to be used to create the blobs
	StorageAccountResourceID *string `json:"storageAccountResourceId,omitempty"`
}

// Single Namespace item in List or Get Operation
type EhNamespace struct {
	TrackedResource
	// Namespace properties supplied for create namespace operation.
	Properties *EhNamespaceProperties `json:"properties,omitempty"`

	// Properties of sku resource
	Sku *Sku `json:"sku,omitempty"`
}

// The response of the List Namespace operation
type EhNamespaceListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of namespaces.
	NextLink *string `json:"nextLink,omitempty"`

	// Result of the List Namespace operation
	Value *[]EhNamespace `json:"value,omitempty"`
}

// EhNamespaceListResultResponse is the response envelope for operations that return a EhNamespaceListResult type.
type EhNamespaceListResultResponse struct {
	// The response of the List Namespace operation
	EhNamespaceListResult *EhNamespaceListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EhNamespacePollerResponse is the response envelope for operations that asynchronously return a EhNamespace type.
type EhNamespacePollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*EhNamespaceResponse, error)

	// Poller contains an initialized poller.
	Poller EhNamespacePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Namespace properties supplied for create namespace operation.
type EhNamespaceProperties struct {
	// The time the Namespace was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled *bool `json:"isAutoInflateEnabled,omitempty"`

	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled *bool `json:"kafkaEnabled,omitempty"`

	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if
	// AutoInflateEnabled = true)
	MaximumThroughputUnits *int32 `json:"maximumThroughputUnits,omitempty"`

	// Identifier for Azure Insights metrics.
	MetricID *string `json:"metricId,omitempty"`

	// Provisioning state of the Namespace.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `json:"serviceBusEndpoint,omitempty"`

	// The time the Namespace was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type EhNamespaceProperties.
func (e EhNamespaceProperties) MarshalJSON() ([]byte, error) {
	type alias EhNamespaceProperties
	aux := &struct {
		*alias
		CreatedAt *timeRFC3339 `json:"createdAt"`
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias:     (*alias)(&e),
		CreatedAt: (*timeRFC3339)(e.CreatedAt),
		UpdatedAt: (*timeRFC3339)(e.UpdatedAt),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EhNamespaceProperties.
func (e *EhNamespaceProperties) UnmarshalJSON(data []byte) error {
	type alias EhNamespaceProperties
	aux := &struct {
		*alias
		CreatedAt *timeRFC3339 `json:"createdAt"`
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias: (*alias)(e),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	e.CreatedAt = (*time.Time)(aux.CreatedAt)
	e.UpdatedAt = (*time.Time)(aux.UpdatedAt)
	return nil
}

// EhNamespaceResponse is the response envelope for operations that return a EhNamespace type.
type EhNamespaceResponse struct {
	// Single Namespace item in List or Get Operation
	EhNamespace *EhNamespace

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Error response indicates Event Hub service is not able to process the incoming request. The reason is provided in the error
// message.
type ErrorResponse struct {
	// Error code.
	Code *string `json:"code,omitempty"`

	// Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
func (e ErrorResponse) Error() string {
	msg := ""
	if e.Code != nil {
		msg += fmt.Sprintf("Code: %v\n", *e.Code)
	}
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// The result of the List EventHubs operation.
type EventHubListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of EventHubs.
	NextLink *string `json:"nextLink,omitempty"`

	// Result of the List EventHubs operation.
	Value *[]Eventhub `json:"value,omitempty"`
}

// EventHubListResultResponse is the response envelope for operations that return a EventHubListResult type.
type EventHubListResultResponse struct {
	// The result of the List EventHubs operation.
	EventHubListResult *EventHubListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// EventHubsListByNamespaceOptions contains the optional parameters for the EventHubs.ListByNamespace method.
type EventHubsListByNamespaceOptions struct {
	// Skip is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element will include a skip parameter that specifies a starting point to use for subsequent calls.
	Skip *int32
	// May be used to limit the number of results to the most recent N usageDetails.
	Top *int32
}

// Single item in List or Get Event Hub operation
type Eventhub struct {
	Resource
	// Properties supplied to the Create Or Update Event Hub operation.
	Properties *EventhubProperties `json:"properties,omitempty"`
}

// Properties supplied to the Create Or Update Event Hub operation.
type EventhubProperties struct {
	// Properties of capture description
	CaptureDescription *CaptureDescription `json:"captureDescription,omitempty"`

	// Exact time the Event Hub was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *int64 `json:"messageRetentionInDays,omitempty"`

	// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *int64 `json:"partitionCount,omitempty"`

	// Current number of shards on the Event Hub.
	PartitionIDs *[]string `json:"partitionIds,omitempty"`

	// Enumerates the possible values for the status of the Event Hub.
	Status *EntityStatus `json:"status,omitempty"`

	// The exact time the message was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type EventhubProperties.
func (e EventhubProperties) MarshalJSON() ([]byte, error) {
	type alias EventhubProperties
	aux := &struct {
		*alias
		CreatedAt *timeRFC3339 `json:"createdAt"`
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias:     (*alias)(&e),
		CreatedAt: (*timeRFC3339)(e.CreatedAt),
		UpdatedAt: (*timeRFC3339)(e.UpdatedAt),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventhubProperties.
func (e *EventhubProperties) UnmarshalJSON(data []byte) error {
	type alias EventhubProperties
	aux := &struct {
		*alias
		CreatedAt *timeRFC3339 `json:"createdAt"`
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias: (*alias)(e),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	e.CreatedAt = (*time.Time)(aux.CreatedAt)
	e.UpdatedAt = (*time.Time)(aux.UpdatedAt)
	return nil
}

// EventhubResponse is the response envelope for operations that return a Eventhub type.
type EventhubResponse struct {
	// Single item in List or Get Event Hub operation
	Eventhub *Eventhub

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HTTPPollerResponse contains the asynchronous HTTP response from the call to the service endpoint.
type HTTPPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received
	PollUntilDone func(ctx context.Context, frequency time.Duration) (*http.Response, error)

	// Poller contains an initialized poller.
	Poller HTTPPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Messaging Plan for the namespace
type MessagingPlan struct {
	TrackedResource
	Properties *MessagingPlanProperties `json:"properties,omitempty"`
}

type MessagingPlanProperties struct {
	// revision number
	Revision *int64 `json:"revision,omitempty"`

	// Selected event hub unit
	SelectedEventHubUnit *int32 `json:"selectedEventHubUnit,omitempty"`

	// Sku type
	Sku *int32 `json:"sku,omitempty"`

	// The exact time the messaging plan was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MessagingPlanProperties.
func (m MessagingPlanProperties) MarshalJSON() ([]byte, error) {
	type alias MessagingPlanProperties
	aux := &struct {
		*alias
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias:     (*alias)(&m),
		UpdatedAt: (*timeRFC3339)(m.UpdatedAt),
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MessagingPlanProperties.
func (m *MessagingPlanProperties) UnmarshalJSON(data []byte) error {
	type alias MessagingPlanProperties
	aux := &struct {
		*alias
		UpdatedAt *timeRFC3339 `json:"updatedAt"`
	}{
		alias: (*alias)(m),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	m.UpdatedAt = (*time.Time)(aux.UpdatedAt)
	return nil
}

// MessagingPlanResponse is the response envelope for operations that return a MessagingPlan type.
type MessagingPlanResponse struct {
	// Messaging Plan for the namespace
	MessagingPlan *MessagingPlan

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Messaging Region
type MessagingRegions struct {
	TrackedResource
	// Properties of Messaging Region
	Properties *MessagingRegionsProperties `json:"properties,omitempty"`
}

// The response of the List MessagingRegions operation.
type MessagingRegionsListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of MessagingRegions.
	NextLink *string `json:"nextLink,omitempty"`

	// Result of the List MessagingRegions type.
	Value *[]MessagingRegions `json:"value,omitempty"`
}

// MessagingRegionsListResultResponse is the response envelope for operations that return a MessagingRegionsListResult type.
type MessagingRegionsListResultResponse struct {
	// The response of the List MessagingRegions operation.
	MessagingRegionsListResult *MessagingRegionsListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Properties of Messaging Region
type MessagingRegionsProperties struct {
	// Region code
	Code *string `json:"code,omitempty"`

	// Full name of the region
	FullName *string `json:"fullName,omitempty"`
}

// Description of NetworkRuleSet resource.
type NetworkRuleSet struct {
	Resource
	// NetworkRuleSet properties
	Properties *NetworkRuleSetProperties `json:"properties,omitempty"`
}

// The response of the List NetworkRuleSet operation
type NetworkRuleSetListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of NetworkRuleSet.
	NextLink *string `json:"nextLink,omitempty"`

	// Result of the List NetworkRuleSet operation.
	Value *[]NetworkRuleSet `json:"value,omitempty"`
}

// NetworkRuleSetListResultResponse is the response envelope for operations that return a NetworkRuleSetListResult type.
type NetworkRuleSetListResultResponse struct {
	// The response of the List NetworkRuleSet operation
	NetworkRuleSetListResult *NetworkRuleSetListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// NetworkRuleSet properties
type NetworkRuleSetProperties struct {
	// Default Action for Network Rule Set
	DefaultAction *DefaultAction `json:"defaultAction,omitempty"`

	// List of IpRules
	IPRules *[]NwRuleSetIPRules `json:"ipRules,omitempty"`

	// List VirtualNetwork Rules
	VirtualNetworkRules *[]NwRuleSetVirtualNetworkRules `json:"virtualNetworkRules,omitempty"`
}

// NetworkRuleSetResponse is the response envelope for operations that return a NetworkRuleSet type.
type NetworkRuleSetResponse struct {
	// Description of NetworkRuleSet resource.
	NetworkRuleSet *NetworkRuleSet

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Description of NetWorkRuleSet - IpRules resource.
type NwRuleSetIPRules struct {
	// The IP Filter Action
	Action *string `json:"action,omitempty"`

	// IP Mask
	IPMask *string `json:"ipMask,omitempty"`
}

// Description of VirtualNetworkRules - NetworkRules resource.
type NwRuleSetVirtualNetworkRules struct {
	// Value that indicates whether to ignore missing VNet Service Endpoint
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`

	// Subnet properties
	Subnet *Subnet `json:"subnet,omitempty"`
}

// A Event Hub REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.EventHub
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Invoice, etc.
	Resource *string `json:"resource,omitempty"`
}

// Result of the request to list Event Hub operations. It contains a list of operations and a URL link to get the next set
// of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Event Hub operations supported by the Microsoft.EventHub resource provider.
	Value *[]Operation `json:"value,omitempty"`
}

// OperationListResultResponse is the response envelope for operations that return a OperationListResult type.
type OperationListResultResponse struct {
	// Result of the request to list Event Hub operations. It contains a list of operations and a URL link to get the next set
	// of results.
	OperationListResult *OperationListResult

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Parameters supplied to the Regenerate Authorization Rule operation, specifies which key needs to be reset.
type RegenerateAccessKeyParameters struct {
	// Optional, if the key value provided, is set for KeyType or autogenerated Key value set for keyType
	Key *string `json:"key,omitempty"`

	// The access key to regenerate.
	KeyType *KeyType `json:"keyType,omitempty"`
}

// The resource definition.
type Resource struct {
	// Resource ID.
	ID *string `json:"id,omitempty"`

	// Resource name.
	Name *string `json:"name,omitempty"`

	// Resource type.
	Type *string `json:"type,omitempty"`
}

// SKU parameters supplied to the create namespace operation
type Sku struct {
	// The Event Hubs throughput units, value should be 0 to 20 throughput units.
	Capacity *int32 `json:"capacity,omitempty"`

	// Name of this SKU.
	Name *SkuName `json:"name,omitempty"`

	// The billing tier of this particular SKU.
	Tier *SkuTier `json:"tier,omitempty"`
}

// Properties supplied for Subnet
type Subnet struct {
	// Resource ID of Virtual Network Subnet
	ID *string `json:"id,omitempty"`
}

// Definition of resource.
type TrackedResource struct {
	Resource
	// Resource location.
	Location *string `json:"location,omitempty"`

	// Resource tags.
	Tags *map[string]string `json:"tags,omitempty"`
}
