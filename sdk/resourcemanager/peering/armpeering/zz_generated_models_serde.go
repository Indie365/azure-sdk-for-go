//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ConnectionMonitorTestProperties.
func (c ConnectionMonitorTestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "destination", c.Destination)
	populate(objectMap, "destinationPort", c.DestinationPort)
	populate(objectMap, "isTestSuccessful", c.IsTestSuccessful)
	populate(objectMap, "path", c.Path)
	populate(objectMap, "provisioningState", c.ProvisioningState)
	populate(objectMap, "sourceAgent", c.SourceAgent)
	populate(objectMap, "testFrequencyInSec", c.TestFrequencyInSec)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LocationPropertiesDirect.
func (l LocationPropertiesDirect) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "bandwidthOffers", l.BandwidthOffers)
	populate(objectMap, "peeringFacilities", l.PeeringFacilities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LocationPropertiesExchange.
func (l LocationPropertiesExchange) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "peeringFacilities", l.PeeringFacilities)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LogAnalyticsWorkspaceProperties.
func (l LogAnalyticsWorkspaceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connectedAgents", l.ConnectedAgents)
	populate(objectMap, "key", l.Key)
	populate(objectMap, "workspaceID", l.WorkspaceID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PeerAsnProperties.
func (p PeerAsnProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "errorMessage", p.ErrorMessage)
	populate(objectMap, "peerAsn", p.PeerAsn)
	populate(objectMap, "peerContactDetail", p.PeerContactDetail)
	populate(objectMap, "peerName", p.PeerName)
	populate(objectMap, "validationState", p.ValidationState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Peering.
func (p Peering) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "kind", p.Kind)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PropertiesDirect.
func (p PropertiesDirect) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connections", p.Connections)
	populate(objectMap, "directPeeringType", p.DirectPeeringType)
	populate(objectMap, "peerAsn", p.PeerAsn)
	populate(objectMap, "useForPeeringService", p.UseForPeeringService)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type PropertiesExchange.
func (p PropertiesExchange) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "connections", p.Connections)
	populate(objectMap, "peerAsn", p.PeerAsn)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTags.
func (r ResourceTags) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", r.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Service.
func (s Service) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServicePrefixEvent.
func (s ServicePrefixEvent) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "eventDescription", s.EventDescription)
	populate(objectMap, "eventLevel", s.EventLevel)
	populate(objectMap, "eventSummary", s.EventSummary)
	populateTimeRFC3339(objectMap, "eventTimestamp", s.EventTimestamp)
	populate(objectMap, "eventType", s.EventType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServicePrefixEvent.
func (s *ServicePrefixEvent) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "eventDescription":
			err = unpopulate(val, "EventDescription", &s.EventDescription)
			delete(rawMsg, key)
		case "eventLevel":
			err = unpopulate(val, "EventLevel", &s.EventLevel)
			delete(rawMsg, key)
		case "eventSummary":
			err = unpopulate(val, "EventSummary", &s.EventSummary)
			delete(rawMsg, key)
		case "eventTimestamp":
			err = unpopulateTimeRFC3339(val, "EventTimestamp", &s.EventTimestamp)
			delete(rawMsg, key)
		case "eventType":
			err = unpopulate(val, "EventType", &s.EventType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServicePrefixProperties.
func (s ServicePrefixProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "errorMessage", s.ErrorMessage)
	populate(objectMap, "events", s.Events)
	populate(objectMap, "learnedType", s.LearnedType)
	populate(objectMap, "peeringServicePrefixKey", s.PeeringServicePrefixKey)
	populate(objectMap, "prefix", s.Prefix)
	populate(objectMap, "prefixValidationState", s.PrefixValidationState)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceProviderProperties.
func (s ServiceProviderProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "peeringLocations", s.PeeringLocations)
	populate(objectMap, "serviceProviderName", s.ServiceProviderName)
	return json.Marshal(objectMap)
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