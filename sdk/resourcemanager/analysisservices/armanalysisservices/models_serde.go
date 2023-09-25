//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CheckServerNameAvailabilityParameters.
func (c CheckServerNameAvailabilityParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckServerNameAvailabilityParameters.
func (c *CheckServerNameAvailabilityParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
				err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &c.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CheckServerNameAvailabilityResult.
func (c CheckServerNameAvailabilityResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "message", c.Message)
	populate(objectMap, "nameAvailable", c.NameAvailable)
	populate(objectMap, "reason", c.Reason)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CheckServerNameAvailabilityResult.
func (c *CheckServerNameAvailabilityResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "message":
				err = unpopulate(val, "Message", &c.Message)
			delete(rawMsg, key)
		case "nameAvailable":
				err = unpopulate(val, "NameAvailable", &c.NameAvailable)
			delete(rawMsg, key)
		case "reason":
				err = unpopulate(val, "Reason", &c.Reason)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorAdditionalInfo.
func (e ErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateAny(objectMap, "info", e.Info)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorAdditionalInfo.
func (e *ErrorAdditionalInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "info":
				err = unpopulate(val, "Info", &e.Info)
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

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "httpStatusCode", e.HTTPStatusCode)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "subCode", e.SubCode)
	populate(objectMap, "target", e.Target)
	populate(objectMap, "timeStamp", e.TimeStamp)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetail.
func (e *ErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInfo":
				err = unpopulate(val, "AdditionalInfo", &e.AdditionalInfo)
			delete(rawMsg, key)
		case "code":
				err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
				err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "httpStatusCode":
				err = unpopulate(val, "HTTPStatusCode", &e.HTTPStatusCode)
			delete(rawMsg, key)
		case "message":
				err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "subCode":
				err = unpopulate(val, "SubCode", &e.SubCode)
			delete(rawMsg, key)
		case "target":
				err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		case "timeStamp":
				err = unpopulate(val, "TimeStamp", &e.TimeStamp)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
				err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GatewayDetails.
func (g GatewayDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "dmtsClusterUri", g.DmtsClusterURI)
	populate(objectMap, "gatewayObjectId", g.GatewayObjectID)
	populate(objectMap, "gatewayResourceId", g.GatewayResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GatewayDetails.
func (g *GatewayDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "dmtsClusterUri":
				err = unpopulate(val, "DmtsClusterURI", &g.DmtsClusterURI)
			delete(rawMsg, key)
		case "gatewayObjectId":
				err = unpopulate(val, "GatewayObjectID", &g.GatewayObjectID)
			delete(rawMsg, key)
		case "gatewayResourceId":
				err = unpopulate(val, "GatewayResourceID", &g.GatewayResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GatewayListStatusError.
func (g GatewayListStatusError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", g.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GatewayListStatusError.
func (g *GatewayListStatusError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
				err = unpopulate(val, "Error", &g.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GatewayListStatusLive.
func (g GatewayListStatusLive) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["status"] = 0
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GatewayListStatusLive.
func (g *GatewayListStatusLive) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "status":
				err = unpopulate(val, "Status", &g.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IPv4FirewallRule.
func (i IPv4FirewallRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "firewallRuleName", i.FirewallRuleName)
	populate(objectMap, "rangeEnd", i.RangeEnd)
	populate(objectMap, "rangeStart", i.RangeStart)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IPv4FirewallRule.
func (i *IPv4FirewallRule) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "firewallRuleName":
				err = unpopulate(val, "FirewallRuleName", &i.FirewallRuleName)
			delete(rawMsg, key)
		case "rangeEnd":
				err = unpopulate(val, "RangeEnd", &i.RangeEnd)
			delete(rawMsg, key)
		case "rangeStart":
				err = unpopulate(val, "RangeStart", &i.RangeStart)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IPv4FirewallSettings.
func (i IPv4FirewallSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "enablePowerBIService", i.EnablePowerBIService)
	populate(objectMap, "firewallRules", i.FirewallRules)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IPv4FirewallSettings.
func (i *IPv4FirewallSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "enablePowerBIService":
				err = unpopulate(val, "EnablePowerBIService", &i.EnablePowerBIService)
			delete(rawMsg, key)
		case "firewallRules":
				err = unpopulate(val, "FirewallRules", &i.FirewallRules)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LogSpecifications.
func (l LogSpecifications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "blobDuration", l.BlobDuration)
	populate(objectMap, "displayName", l.DisplayName)
	populate(objectMap, "name", l.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogSpecifications.
func (l *LogSpecifications) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "blobDuration":
				err = unpopulate(val, "BlobDuration", &l.BlobDuration)
			delete(rawMsg, key)
		case "displayName":
				err = unpopulate(val, "DisplayName", &l.DisplayName)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &l.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricDimensions.
func (m MetricDimensions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "name", m.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricDimensions.
func (m *MetricDimensions) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
				err = unpopulate(val, "DisplayName", &m.DisplayName)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricSpecifications.
func (m MetricSpecifications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "aggregationType", m.AggregationType)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricSpecifications.
func (m *MetricSpecifications) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "aggregationType":
				err = unpopulate(val, "AggregationType", &m.AggregationType)
			delete(rawMsg, key)
		case "dimensions":
				err = unpopulate(val, "Dimensions", &m.Dimensions)
			delete(rawMsg, key)
		case "displayDescription":
				err = unpopulate(val, "DisplayDescription", &m.DisplayDescription)
			delete(rawMsg, key)
		case "displayName":
				err = unpopulate(val, "DisplayName", &m.DisplayName)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "unit":
				err = unpopulate(val, "Unit", &m.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	populate(objectMap, "properties", o.Properties)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
				err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "origin":
				err = unpopulate(val, "Origin", &o.Origin)
			delete(rawMsg, key)
		case "properties":
				err = unpopulate(val, "Properties", &o.Properties)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
				err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
				err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
				err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
				err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
				err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
				err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationProperties.
func (o OperationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "serviceSpecification", o.ServiceSpecification)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationProperties.
func (o *OperationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "serviceSpecification":
				err = unpopulate(val, "ServiceSpecification", &o.ServiceSpecification)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationPropertiesServiceSpecification.
func (o OperationPropertiesServiceSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "logSpecifications", o.LogSpecifications)
	populate(objectMap, "metricSpecifications", o.MetricSpecifications)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationPropertiesServiceSpecification.
func (o *OperationPropertiesServiceSpecification) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "logSpecifications":
				err = unpopulate(val, "LogSpecifications", &o.LogSpecifications)
			delete(rawMsg, key)
		case "metricSpecifications":
				err = unpopulate(val, "MetricSpecifications", &o.MetricSpecifications)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationStatus.
func (o OperationStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "endTime", o.EndTime)
	populate(objectMap, "error", o.Error)
	populate(objectMap, "id", o.ID)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "startTime", o.StartTime)
	populate(objectMap, "status", o.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationStatus.
func (o *OperationStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "endTime":
				err = unpopulate(val, "EndTime", &o.EndTime)
			delete(rawMsg, key)
		case "error":
				err = unpopulate(val, "Error", &o.Error)
			delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &o.ID)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		case "startTime":
				err = unpopulate(val, "StartTime", &o.StartTime)
			delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &o.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "sku", r.SKU)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Resource.
func (r *Resource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
				err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "location":
				err = unpopulate(val, "Location", &r.Location)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "sku":
				err = unpopulate(val, "SKU", &r.SKU)
			delete(rawMsg, key)
		case "tags":
				err = unpopulate(val, "Tags", &r.Tags)
			delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &r.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceSKU.
func (r ResourceSKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "capacity", r.Capacity)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tier", r.Tier)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceSKU.
func (r *ResourceSKU) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "capacity":
				err = unpopulate(val, "Capacity", &r.Capacity)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "tier":
				err = unpopulate(val, "Tier", &r.Tier)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUDetailsForExistingResource.
func (s SKUDetailsForExistingResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "resourceType", s.ResourceType)
	populate(objectMap, "sku", s.SKU)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKUDetailsForExistingResource.
func (s *SKUDetailsForExistingResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "resourceType":
				err = unpopulate(val, "ResourceType", &s.ResourceType)
			delete(rawMsg, key)
		case "sku":
				err = unpopulate(val, "SKU", &s.SKU)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUEnumerationForExistingResourceResult.
func (s SKUEnumerationForExistingResourceResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKUEnumerationForExistingResourceResult.
func (s *SKUEnumerationForExistingResourceResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
				err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUEnumerationForNewResourceResult.
func (s SKUEnumerationForNewResourceResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKUEnumerationForNewResourceResult.
func (s *SKUEnumerationForNewResourceResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
				err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Server.
func (s Server) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Server.
func (s *Server) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
				err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "location":
				err = unpopulate(val, "Location", &s.Location)
			delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "properties":
				err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "sku":
				err = unpopulate(val, "SKU", &s.SKU)
			delete(rawMsg, key)
		case "tags":
				err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		case "type":
				err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServerAdministrators.
func (s ServerAdministrators) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "members", s.Members)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerAdministrators.
func (s *ServerAdministrators) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "members":
				err = unpopulate(val, "Members", &s.Members)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServerMutableProperties.
func (s ServerMutableProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "asAdministrators", s.AsAdministrators)
	populate(objectMap, "backupBlobContainerUri", s.BackupBlobContainerURI)
	populate(objectMap, "gatewayDetails", s.GatewayDetails)
	populate(objectMap, "ipV4FirewallSettings", s.IPV4FirewallSettings)
	populate(objectMap, "managedMode", s.ManagedMode)
	populate(objectMap, "querypoolConnectionMode", s.QuerypoolConnectionMode)
	populate(objectMap, "serverMonitorMode", s.ServerMonitorMode)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerMutableProperties.
func (s *ServerMutableProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "asAdministrators":
				err = unpopulate(val, "AsAdministrators", &s.AsAdministrators)
			delete(rawMsg, key)
		case "backupBlobContainerUri":
				err = unpopulate(val, "BackupBlobContainerURI", &s.BackupBlobContainerURI)
			delete(rawMsg, key)
		case "gatewayDetails":
				err = unpopulate(val, "GatewayDetails", &s.GatewayDetails)
			delete(rawMsg, key)
		case "ipV4FirewallSettings":
				err = unpopulate(val, "IPV4FirewallSettings", &s.IPV4FirewallSettings)
			delete(rawMsg, key)
		case "managedMode":
				err = unpopulate(val, "ManagedMode", &s.ManagedMode)
			delete(rawMsg, key)
		case "querypoolConnectionMode":
				err = unpopulate(val, "QuerypoolConnectionMode", &s.QuerypoolConnectionMode)
			delete(rawMsg, key)
		case "serverMonitorMode":
				err = unpopulate(val, "ServerMonitorMode", &s.ServerMonitorMode)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServerProperties.
func (s ServerProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "asAdministrators", s.AsAdministrators)
	populate(objectMap, "backupBlobContainerUri", s.BackupBlobContainerURI)
	populate(objectMap, "gatewayDetails", s.GatewayDetails)
	populate(objectMap, "ipV4FirewallSettings", s.IPV4FirewallSettings)
	populate(objectMap, "managedMode", s.ManagedMode)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "querypoolConnectionMode", s.QuerypoolConnectionMode)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "serverFullName", s.ServerFullName)
	populate(objectMap, "serverMonitorMode", s.ServerMonitorMode)
	populate(objectMap, "state", s.State)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerProperties.
func (s *ServerProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "asAdministrators":
				err = unpopulate(val, "AsAdministrators", &s.AsAdministrators)
			delete(rawMsg, key)
		case "backupBlobContainerUri":
				err = unpopulate(val, "BackupBlobContainerURI", &s.BackupBlobContainerURI)
			delete(rawMsg, key)
		case "gatewayDetails":
				err = unpopulate(val, "GatewayDetails", &s.GatewayDetails)
			delete(rawMsg, key)
		case "ipV4FirewallSettings":
				err = unpopulate(val, "IPV4FirewallSettings", &s.IPV4FirewallSettings)
			delete(rawMsg, key)
		case "managedMode":
				err = unpopulate(val, "ManagedMode", &s.ManagedMode)
			delete(rawMsg, key)
		case "provisioningState":
				err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "querypoolConnectionMode":
				err = unpopulate(val, "QuerypoolConnectionMode", &s.QuerypoolConnectionMode)
			delete(rawMsg, key)
		case "sku":
				err = unpopulate(val, "SKU", &s.SKU)
			delete(rawMsg, key)
		case "serverFullName":
				err = unpopulate(val, "ServerFullName", &s.ServerFullName)
			delete(rawMsg, key)
		case "serverMonitorMode":
				err = unpopulate(val, "ServerMonitorMode", &s.ServerMonitorMode)
			delete(rawMsg, key)
		case "state":
				err = unpopulate(val, "State", &s.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServerUpdateParameters.
func (s ServerUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServerUpdateParameters.
func (s *ServerUpdateParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
				err = unpopulate(val, "Properties", &s.Properties)
			delete(rawMsg, key)
		case "sku":
				err = unpopulate(val, "SKU", &s.SKU)
			delete(rawMsg, key)
		case "tags":
				err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Servers.
func (s Servers) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Servers.
func (s *Servers) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
				err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}

