//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdigitaltwins

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AzureDataExplorerConnectionProperties.
func (a AzureDataExplorerConnectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "adxDatabaseName", a.AdxDatabaseName)
	populate(objectMap, "adxEndpointUri", a.AdxEndpointURI)
	populate(objectMap, "adxResourceId", a.AdxResourceID)
	populate(objectMap, "adxTableName", a.AdxTableName)
	objectMap["connectionType"] = ConnectionTypeAzureDataExplorer
	populate(objectMap, "eventHubConsumerGroup", a.EventHubConsumerGroup)
	populate(objectMap, "eventHubEndpointUri", a.EventHubEndpointURI)
	populate(objectMap, "eventHubEntityPath", a.EventHubEntityPath)
	populate(objectMap, "eventHubNamespaceResourceId", a.EventHubNamespaceResourceID)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureDataExplorerConnectionProperties.
func (a *AzureDataExplorerConnectionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "adxDatabaseName":
			err = unpopulate(val, "AdxDatabaseName", &a.AdxDatabaseName)
			delete(rawMsg, key)
		case "adxEndpointUri":
			err = unpopulate(val, "AdxEndpointURI", &a.AdxEndpointURI)
			delete(rawMsg, key)
		case "adxResourceId":
			err = unpopulate(val, "AdxResourceID", &a.AdxResourceID)
			delete(rawMsg, key)
		case "adxTableName":
			err = unpopulate(val, "AdxTableName", &a.AdxTableName)
			delete(rawMsg, key)
		case "connectionType":
			err = unpopulate(val, "ConnectionType", &a.ConnectionType)
			delete(rawMsg, key)
		case "eventHubConsumerGroup":
			err = unpopulate(val, "EventHubConsumerGroup", &a.EventHubConsumerGroup)
			delete(rawMsg, key)
		case "eventHubEndpointUri":
			err = unpopulate(val, "EventHubEndpointURI", &a.EventHubEndpointURI)
			delete(rawMsg, key)
		case "eventHubEntityPath":
			err = unpopulate(val, "EventHubEntityPath", &a.EventHubEntityPath)
			delete(rawMsg, key)
		case "eventHubNamespaceResourceId":
			err = unpopulate(val, "EventHubNamespaceResourceID", &a.EventHubNamespaceResourceID)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &a.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConnectionProperties.
func (c ConnectionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupIds", c.GroupIDs)
	populate(objectMap, "privateEndpoint", c.PrivateEndpoint)
	populate(objectMap, "privateLinkServiceConnectionState", c.PrivateLinkServiceConnectionState)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Description.
func (d Description) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "identity", d.Identity)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "systemData", d.SystemData)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type EndpointResource.
func (e EndpointResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "systemData", e.SystemData)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EndpointResource.
func (e *EndpointResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &e.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &e.Name)
			delete(rawMsg, key)
		case "properties":
			e.Properties, err = unmarshalEndpointResourcePropertiesClassification(val)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &e.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EndpointResourceProperties.
func (e EndpointResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authenticationType", e.AuthenticationType)
	populateTimeRFC3339(objectMap, "createdTime", e.CreatedTime)
	populate(objectMap, "deadLetterSecret", e.DeadLetterSecret)
	populate(objectMap, "deadLetterUri", e.DeadLetterURI)
	objectMap["endpointType"] = e.EndpointType
	populate(objectMap, "provisioningState", e.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EndpointResourceProperties.
func (e *EndpointResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authenticationType":
			err = unpopulate(val, "AuthenticationType", &e.AuthenticationType)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateTimeRFC3339(val, "CreatedTime", &e.CreatedTime)
			delete(rawMsg, key)
		case "deadLetterSecret":
			err = unpopulate(val, "DeadLetterSecret", &e.DeadLetterSecret)
			delete(rawMsg, key)
		case "deadLetterUri":
			err = unpopulate(val, "DeadLetterURI", &e.DeadLetterURI)
			delete(rawMsg, key)
		case "endpointType":
			err = unpopulate(val, "EndpointType", &e.EndpointType)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &e.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventGrid.
func (e EventGrid) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accessKey1", e.AccessKey1)
	populate(objectMap, "accessKey2", e.AccessKey2)
	populate(objectMap, "authenticationType", e.AuthenticationType)
	populateTimeRFC3339(objectMap, "createdTime", e.CreatedTime)
	populate(objectMap, "deadLetterSecret", e.DeadLetterSecret)
	populate(objectMap, "deadLetterUri", e.DeadLetterURI)
	objectMap["endpointType"] = EndpointTypeEventGrid
	populate(objectMap, "provisioningState", e.ProvisioningState)
	populate(objectMap, "TopicEndpoint", e.TopicEndpoint)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventGrid.
func (e *EventGrid) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accessKey1":
			err = unpopulate(val, "AccessKey1", &e.AccessKey1)
			delete(rawMsg, key)
		case "accessKey2":
			err = unpopulate(val, "AccessKey2", &e.AccessKey2)
			delete(rawMsg, key)
		case "authenticationType":
			err = unpopulate(val, "AuthenticationType", &e.AuthenticationType)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateTimeRFC3339(val, "CreatedTime", &e.CreatedTime)
			delete(rawMsg, key)
		case "deadLetterSecret":
			err = unpopulate(val, "DeadLetterSecret", &e.DeadLetterSecret)
			delete(rawMsg, key)
		case "deadLetterUri":
			err = unpopulate(val, "DeadLetterURI", &e.DeadLetterURI)
			delete(rawMsg, key)
		case "endpointType":
			err = unpopulate(val, "EndpointType", &e.EndpointType)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &e.ProvisioningState)
			delete(rawMsg, key)
		case "TopicEndpoint":
			err = unpopulate(val, "TopicEndpoint", &e.TopicEndpoint)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventHub.
func (e EventHub) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authenticationType", e.AuthenticationType)
	populate(objectMap, "connectionStringPrimaryKey", e.ConnectionStringPrimaryKey)
	populate(objectMap, "connectionStringSecondaryKey", e.ConnectionStringSecondaryKey)
	populateTimeRFC3339(objectMap, "createdTime", e.CreatedTime)
	populate(objectMap, "deadLetterSecret", e.DeadLetterSecret)
	populate(objectMap, "deadLetterUri", e.DeadLetterURI)
	objectMap["endpointType"] = EndpointTypeEventHub
	populate(objectMap, "endpointUri", e.EndpointURI)
	populate(objectMap, "entityPath", e.EntityPath)
	populate(objectMap, "provisioningState", e.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventHub.
func (e *EventHub) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authenticationType":
			err = unpopulate(val, "AuthenticationType", &e.AuthenticationType)
			delete(rawMsg, key)
		case "connectionStringPrimaryKey":
			err = unpopulate(val, "ConnectionStringPrimaryKey", &e.ConnectionStringPrimaryKey)
			delete(rawMsg, key)
		case "connectionStringSecondaryKey":
			err = unpopulate(val, "ConnectionStringSecondaryKey", &e.ConnectionStringSecondaryKey)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateTimeRFC3339(val, "CreatedTime", &e.CreatedTime)
			delete(rawMsg, key)
		case "deadLetterSecret":
			err = unpopulate(val, "DeadLetterSecret", &e.DeadLetterSecret)
			delete(rawMsg, key)
		case "deadLetterUri":
			err = unpopulate(val, "DeadLetterURI", &e.DeadLetterURI)
			delete(rawMsg, key)
		case "endpointType":
			err = unpopulate(val, "EndpointType", &e.EndpointType)
			delete(rawMsg, key)
		case "endpointUri":
			err = unpopulate(val, "EndpointURI", &e.EndpointURI)
			delete(rawMsg, key)
		case "entityPath":
			err = unpopulate(val, "EntityPath", &e.EntityPath)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &e.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PatchDescription.
func (p PatchDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "identity", p.Identity)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "tags", p.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdTime", p.CreatedTime)
	populate(objectMap, "hostName", p.HostName)
	populateTimeRFC3339(objectMap, "lastUpdatedTime", p.LastUpdatedTime)
	populate(objectMap, "privateEndpointConnections", p.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", p.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdTime":
			err = unpopulateTimeRFC3339(val, "CreatedTime", &p.CreatedTime)
			delete(rawMsg, key)
		case "hostName":
			err = unpopulate(val, "HostName", &p.HostName)
			delete(rawMsg, key)
		case "lastUpdatedTime":
			err = unpopulateTimeRFC3339(val, "LastUpdatedTime", &p.LastUpdatedTime)
			delete(rawMsg, key)
		case "privateEndpointConnections":
			err = unpopulate(val, "PrivateEndpointConnections", &p.PrivateEndpointConnections)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		case "publicNetworkAccess":
			err = unpopulate(val, "PublicNetworkAccess", &p.PublicNetworkAccess)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "identity", r.Identity)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "systemData", r.SystemData)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceBus.
func (s ServiceBus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authenticationType", s.AuthenticationType)
	populateTimeRFC3339(objectMap, "createdTime", s.CreatedTime)
	populate(objectMap, "deadLetterSecret", s.DeadLetterSecret)
	populate(objectMap, "deadLetterUri", s.DeadLetterURI)
	objectMap["endpointType"] = EndpointTypeServiceBus
	populate(objectMap, "endpointUri", s.EndpointURI)
	populate(objectMap, "entityPath", s.EntityPath)
	populate(objectMap, "primaryConnectionString", s.PrimaryConnectionString)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "secondaryConnectionString", s.SecondaryConnectionString)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceBus.
func (s *ServiceBus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "authenticationType":
			err = unpopulate(val, "AuthenticationType", &s.AuthenticationType)
			delete(rawMsg, key)
		case "createdTime":
			err = unpopulateTimeRFC3339(val, "CreatedTime", &s.CreatedTime)
			delete(rawMsg, key)
		case "deadLetterSecret":
			err = unpopulate(val, "DeadLetterSecret", &s.DeadLetterSecret)
			delete(rawMsg, key)
		case "deadLetterUri":
			err = unpopulate(val, "DeadLetterURI", &s.DeadLetterURI)
			delete(rawMsg, key)
		case "endpointType":
			err = unpopulate(val, "EndpointType", &s.EndpointType)
			delete(rawMsg, key)
		case "endpointUri":
			err = unpopulate(val, "EndpointURI", &s.EndpointURI)
			delete(rawMsg, key)
		case "entityPath":
			err = unpopulate(val, "EntityPath", &s.EntityPath)
			delete(rawMsg, key)
		case "primaryConnectionString":
			err = unpopulate(val, "PrimaryConnectionString", &s.PrimaryConnectionString)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "secondaryConnectionString":
			err = unpopulate(val, "SecondaryConnectionString", &s.SecondaryConnectionString)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
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
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TimeSeriesDatabaseConnection.
func (t TimeSeriesDatabaseConnection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "properties", t.Properties)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TimeSeriesDatabaseConnection.
func (t *TimeSeriesDatabaseConnection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &t.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &t.Name)
			delete(rawMsg, key)
		case "properties":
			t.Properties, err = unmarshalTimeSeriesDatabaseConnectionPropertiesClassification(val)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &t.SystemData)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &t.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
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

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}