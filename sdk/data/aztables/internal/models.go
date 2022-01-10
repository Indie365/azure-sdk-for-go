//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

import (
	"encoding/json"
	"encoding/xml"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AccessPolicy - An Access policy.
type AccessPolicy struct {
	// REQUIRED; The datetime that the policy expires.
	Expiry *time.Time `xml:"Expiry"`

	// REQUIRED; The permissions for the acl policy.
	Permission *string `xml:"Permission"`

	// REQUIRED; The start datetime from which the policy is active.
	Start *time.Time `xml:"Start"`
}

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(&a),
		Expiry: (*timeRFC3339)(a.Expiry),
		Start: (*timeRFC3339)(a.Start),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// CorsRule - CORS is an HTTP feature that enables a web application running under one domain to access resources in another
// domain. Web browsers implement a security restriction known as same-origin policy that
// prevents a web page from calling APIs in a different domain; CORS provides a secure way to allow one domain (the origin
// domain) to call APIs in another domain.
type CorsRule struct {
	// REQUIRED; The request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `xml:"AllowedHeaders"`

	// REQUIRED; The methods (HTTP request verbs) that the origin domain may use for a CORS request. (comma separated)
	AllowedMethods *string `xml:"AllowedMethods"`

	// REQUIRED; The origin domains that are permitted to make a request against the service via CORS. The origin domain is the
// domain from which the request originates. Note that the origin must be an exact
// case-sensitive match with the origin that the user age sends to the service. You can also use the wildcard character '*'
// to allow all origin domains to make requests via CORS.
	AllowedOrigins *string `xml:"AllowedOrigins"`

	// REQUIRED; The response headers that may be sent in the response to the CORS request and exposed by the browser to the request
// issuer.
	ExposedHeaders *string `xml:"ExposedHeaders"`

	// REQUIRED; The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int32 `xml:"MaxAgeInSeconds"`
}

type GeoReplication struct {
	// REQUIRED; A GMT date/time value, to the second. All primary writes preceding this value are guaranteed to be available
// for read operations at the secondary. Primary writes after this point in time may or may
// not be available for reads.
	LastSyncTime *time.Time `xml:"LastSyncTime"`

	// REQUIRED; The status of the secondary location.
	Status *GeoReplicationStatusType `xml:"Status"`
}

// MarshalXML implements the xml.Marshaller interface for type GeoReplication.
func (g GeoReplication) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias: (*alias)(&g),
		LastSyncTime: (*timeRFC1123)(g.LastSyncTime),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type GeoReplication.
func (g *GeoReplication) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias GeoReplication
	aux := &struct {
		*alias
		LastSyncTime *timeRFC1123 `xml:"LastSyncTime"`
	}{
		alias: (*alias)(g),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	g.LastSyncTime = (*time.Time)(aux.LastSyncTime)
	return nil
}

// Logging - Azure Analytics Logging settings.
type Logging struct {
	// REQUIRED; Indicates whether all delete requests should be logged.
	Delete *bool `xml:"Delete"`

	// REQUIRED; Indicates whether all read requests should be logged.
	Read *bool `xml:"Read"`

	// REQUIRED; The retention policy.
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// REQUIRED; The version of Analytics to configure.
	Version *string `xml:"Version"`

	// REQUIRED; Indicates whether all write requests should be logged.
	Write *bool `xml:"Write"`
}

type Metrics struct {
	// REQUIRED; Indicates whether metrics are enabled for the Table service.
	Enabled *bool `xml:"Enabled"`

	// Indicates whether metrics should generate summary statistics for called API operations.
	IncludeAPIs *bool `xml:"IncludeAPIs"`

	// The retention policy.
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// The version of Analytics to configure.
	Version *string `xml:"Version"`
}

// QueryOptions contains a group of parameters for the TableClient.Query method.
type QueryOptions struct {
	// OData filter expression.
	Filter *string
	// Specifies the media type for the response.
	Format *ODataMetadataFormat
	// Select expression using OData notation. Limits the columns on each record to just those requested, e.g. "$select=PolicyAssignmentId,
// ResourceId".
	Select *string
	// Maximum number of records to return.
	Top *int32
}

// RetentionPolicy - The retention policy.
type RetentionPolicy struct {
	// REQUIRED; Indicates whether a retention policy is enabled for the service.
	Enabled *bool `xml:"Enabled"`

	// Indicates the number of days that metrics or logging or soft-deleted data should be retained. All data older than this
// value will be deleted.
	Days *int32 `xml:"Days"`
}

// ServiceClientGetPropertiesOptions contains the optional parameters for the ServiceClient.GetProperties method.
type ServiceClientGetPropertiesOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ServiceClientGetStatisticsOptions contains the optional parameters for the ServiceClient.GetStatistics method.
type ServiceClientGetStatisticsOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// ServiceClientSetPropertiesOptions contains the optional parameters for the ServiceClient.SetProperties method.
type ServiceClientSetPropertiesOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// SignedIdentifier - A signed identifier.
type SignedIdentifier struct {
	// REQUIRED; The access policy.
	AccessPolicy *AccessPolicy `xml:"AccessPolicy"`

	// REQUIRED; A unique id.
	ID *string `xml:"Id"`
}

// TableClientCreateOptions contains the optional parameters for the TableClient.Create method.
type TableClientCreateOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// Specifies whether the response should include the inserted entity in the payload. Possible values are return-no-content
// and return-content.
	ResponsePreference *ResponseFormat
}

// TableClientDeleteEntityOptions contains the optional parameters for the TableClient.DeleteEntity method.
type TableClientDeleteEntityOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableClientDeleteOptions contains the optional parameters for the TableClient.Delete method.
type TableClientDeleteOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
}

// TableClientGetAccessPolicyOptions contains the optional parameters for the TableClient.GetAccessPolicy method.
type TableClientGetAccessPolicyOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableClientInsertEntityOptions contains the optional parameters for the TableClient.InsertEntity method.
type TableClientInsertEntityOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// Specifies whether the response should include the inserted entity in the payload. Possible values are return-no-content
// and return-content.
	ResponsePreference *ResponseFormat
	// The properties for the table entity.
	TableEntityProperties map[string]interface{}
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableClientMergeEntityOptions contains the optional parameters for the TableClient.MergeEntity method.
type TableClientMergeEntityOptions struct {
	// Match condition for an entity to be updated. If specified and a matching entity is not found, an error will be raised.
// To force an unconditional update, set to the wildcard character (*). If not
// specified, an insert will be performed when no existing entity is found to update and a merge will be performed if an existing
// entity is found.
	IfMatch *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The properties for the table entity.
	TableEntityProperties map[string]interface{}
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableClientQueryEntitiesOptions contains the optional parameters for the TableClient.QueryEntities method.
type TableClientQueryEntitiesOptions struct {
	// An entity query continuation token from a previous call.
	NextPartitionKey *string
	// An entity query continuation token from a previous call.
	NextRowKey *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableClientQueryEntityWithPartitionAndRowKeyOptions contains the optional parameters for the TableClient.QueryEntityWithPartitionAndRowKey
// method.
type TableClientQueryEntityWithPartitionAndRowKeyOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableClientQueryOptions contains the optional parameters for the TableClient.Query method.
type TableClientQueryOptions struct {
	// A table query continuation token from a previous call.
	NextTableName *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
}

// TableClientSetAccessPolicyOptions contains the optional parameters for the TableClient.SetAccessPolicy method.
type TableClientSetAccessPolicyOptions struct {
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The acls for the table.
	TableACL []*SignedIdentifier
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableClientUpdateEntityOptions contains the optional parameters for the TableClient.UpdateEntity method.
type TableClientUpdateEntityOptions struct {
	// Match condition for an entity to be updated. If specified and a matching entity is not found, an error will be raised.
// To force an unconditional update, set to the wildcard character (*). If not
// specified, an insert will be performed when no existing entity is found to update and a replace will be performed if an
// existing entity is found.
	IfMatch *string
	// Provides a client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when analytics
// logging is enabled.
	RequestID *string
	// The properties for the table entity.
	TableEntityProperties map[string]interface{}
	// The timeout parameter is expressed in seconds.
	Timeout *int32
}

// TableEntityQueryResponse - The properties for the table entity query response.
type TableEntityQueryResponse struct {
	// The metadata response of the table.
	ODataMetadata *string `json:"odata.metadata,omitempty"`

	// List of table entities.
	Value []map[string]interface{} `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TableEntityQueryResponse.
func (t TableEntityQueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "odata.metadata", t.ODataMetadata)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// TableProperties - The properties for creating a table.
type TableProperties struct {
	// The name of the table to create.
	TableName *string `json:"TableName,omitempty"`
}

// TableQueryResponse - The properties for the table query response.
type TableQueryResponse struct {
	// The metadata response of the table.
	ODataMetadata *string `json:"odata.metadata,omitempty"`

	// List of tables.
	Value []*TableResponseProperties `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TableQueryResponse.
func (t TableQueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "odata.metadata", t.ODataMetadata)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

// TableResponse - The response for a single table.
type TableResponse struct {
	// The edit link of the table.
	ODataEditLink *string `json:"odata.editLink,omitempty"`

	// The id of the table.
	ODataID *string `json:"odata.id,omitempty"`

	// The metadata response of the table.
	ODataMetadata *string `json:"odata.metadata,omitempty"`

	// The odata type of the table.
	ODataType *string `json:"odata.type,omitempty"`

	// The name of the table.
	TableName *string `json:"TableName,omitempty"`
}

// TableResponseProperties - The properties for the table response.
type TableResponseProperties struct {
	// The edit link of the table.
	ODataEditLink *string `json:"odata.editLink,omitempty"`

	// The id of the table.
	ODataID *string `json:"odata.id,omitempty"`

	// The odata type of the table.
	ODataType *string `json:"odata.type,omitempty"`

	// The name of the table.
	TableName *string `json:"TableName,omitempty"`
}

// TableServiceError - Table Service error.
type TableServiceError struct {
	// The error message.
	Message *string `json:"Message,omitempty"`
}

// TableServiceProperties - Table Service Properties.
type TableServiceProperties struct {
	// The set of CORS rules.
	Cors []*CorsRule `xml:"Cors>CorsRule"`

	// A summary of request statistics grouped by API in hourly aggregates for tables.
	HourMetrics *Metrics `xml:"HourMetrics"`

	// Azure Analytics Logging settings.
	Logging *Logging `xml:"Logging"`

	// A summary of request statistics grouped by API in minute aggregates for tables.
	MinuteMetrics *Metrics `xml:"MinuteMetrics"`
}

// MarshalXML implements the xml.Marshaller interface for type TableServiceProperties.
func (t TableServiceProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "StorageServiceProperties"
	type alias TableServiceProperties
	aux := &struct {
		*alias
		Cors *[]*CorsRule `xml:"Cors>CorsRule"`
	}{
		alias: (*alias)(&t),
	}
	if t.Cors != nil {
		aux.Cors = &t.Cors
	}
	return e.EncodeElement(aux, start)
}

// TableServiceStats - Stats for the service.
type TableServiceStats struct {
	// Geo-Replication information for the Secondary Storage Service.
	GeoReplication *GeoReplication `xml:"GeoReplication"`
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

