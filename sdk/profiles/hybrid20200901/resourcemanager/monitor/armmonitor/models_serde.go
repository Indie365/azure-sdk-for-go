//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type DiagnosticSettings.
func (d DiagnosticSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "eventHubAuthorizationRuleId", d.EventHubAuthorizationRuleID)
	populate(objectMap, "eventHubName", d.EventHubName)
	populate(objectMap, "logAnalyticsDestinationType", d.LogAnalyticsDestinationType)
	populate(objectMap, "logs", d.Logs)
	populate(objectMap, "metrics", d.Metrics)
	populate(objectMap, "serviceBusRuleId", d.ServiceBusRuleID)
	populate(objectMap, "storageAccountId", d.StorageAccountID)
	populate(objectMap, "workspaceId", d.WorkspaceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiagnosticSettings.
func (d *DiagnosticSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "eventHubAuthorizationRuleId":
			err = unpopulate(val, "EventHubAuthorizationRuleID", &d.EventHubAuthorizationRuleID)
			delete(rawMsg, key)
		case "eventHubName":
			err = unpopulate(val, "EventHubName", &d.EventHubName)
			delete(rawMsg, key)
		case "logAnalyticsDestinationType":
			err = unpopulate(val, "LogAnalyticsDestinationType", &d.LogAnalyticsDestinationType)
			delete(rawMsg, key)
		case "logs":
			err = unpopulate(val, "Logs", &d.Logs)
			delete(rawMsg, key)
		case "metrics":
			err = unpopulate(val, "Metrics", &d.Metrics)
			delete(rawMsg, key)
		case "serviceBusRuleId":
			err = unpopulate(val, "ServiceBusRuleID", &d.ServiceBusRuleID)
			delete(rawMsg, key)
		case "storageAccountId":
			err = unpopulate(val, "StorageAccountID", &d.StorageAccountID)
			delete(rawMsg, key)
		case "workspaceId":
			err = unpopulate(val, "WorkspaceID", &d.WorkspaceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DiagnosticSettingsCategory.
func (d DiagnosticSettingsCategory) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "categoryType", d.CategoryType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiagnosticSettingsCategory.
func (d *DiagnosticSettingsCategory) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "categoryType":
			err = unpopulate(val, "CategoryType", &d.CategoryType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DiagnosticSettingsCategoryResource.
func (d DiagnosticSettingsCategoryResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiagnosticSettingsCategoryResource.
func (d *DiagnosticSettingsCategoryResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &d.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &d.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DiagnosticSettingsCategoryResourceCollection.
func (d DiagnosticSettingsCategoryResourceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiagnosticSettingsCategoryResourceCollection.
func (d *DiagnosticSettingsCategoryResourceCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &d.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DiagnosticSettingsResource.
func (d DiagnosticSettingsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiagnosticSettingsResource.
func (d *DiagnosticSettingsResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &d.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &d.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DiagnosticSettingsResourceCollection.
func (d DiagnosticSettingsResourceCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiagnosticSettingsResourceCollection.
func (d *DiagnosticSettingsResourceCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &d.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
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
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EventCategoryCollection.
func (e EventCategoryCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", e.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EventCategoryCollection.
func (e *EventCategoryCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &e.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LocalizableString.
func (l LocalizableString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "localizedValue", l.LocalizedValue)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LocalizableString.
func (l *LocalizableString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "localizedValue":
			err = unpopulate(val, "LocalizedValue", &l.LocalizedValue)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LogSettings.
func (l LogSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", l.Category)
	populate(objectMap, "enabled", l.Enabled)
	populate(objectMap, "retentionPolicy", l.RetentionPolicy)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LogSettings.
func (l *LogSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "category":
			err = unpopulate(val, "Category", &l.Category)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, "Enabled", &l.Enabled)
			delete(rawMsg, key)
		case "retentionPolicy":
			err = unpopulate(val, "RetentionPolicy", &l.RetentionPolicy)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetadataValue.
func (m MetadataValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", m.Name)
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetadataValue.
func (m *MetadataValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Metric.
func (m Metric) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "errorCode", m.ErrorCode)
	populate(objectMap, "errorMessage", m.ErrorMessage)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "timeseries", m.Timeseries)
	populate(objectMap, "type", m.Type)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Metric.
func (m *Metric) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayDescription":
			err = unpopulate(val, "DisplayDescription", &m.DisplayDescription)
			delete(rawMsg, key)
		case "errorCode":
			err = unpopulate(val, "ErrorCode", &m.ErrorCode)
			delete(rawMsg, key)
		case "errorMessage":
			err = unpopulate(val, "ErrorMessage", &m.ErrorMessage)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "timeseries":
			err = unpopulate(val, "Timeseries", &m.Timeseries)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &m.Type)
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

// MarshalJSON implements the json.Marshaller interface for type MetricAvailability.
func (m MetricAvailability) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "retention", m.Retention)
	populate(objectMap, "timeGrain", m.TimeGrain)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricAvailability.
func (m *MetricAvailability) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "retention":
			err = unpopulate(val, "Retention", &m.Retention)
			delete(rawMsg, key)
		case "timeGrain":
			err = unpopulate(val, "TimeGrain", &m.TimeGrain)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricDefinition.
func (m MetricDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", m.Category)
	populate(objectMap, "dimensions", m.Dimensions)
	populate(objectMap, "displayDescription", m.DisplayDescription)
	populate(objectMap, "id", m.ID)
	populate(objectMap, "isDimensionRequired", m.IsDimensionRequired)
	populate(objectMap, "metricAvailabilities", m.MetricAvailabilities)
	populate(objectMap, "metricClass", m.MetricClass)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "namespace", m.Namespace)
	populate(objectMap, "primaryAggregationType", m.PrimaryAggregationType)
	populate(objectMap, "resourceId", m.ResourceID)
	populate(objectMap, "supportedAggregationTypes", m.SupportedAggregationTypes)
	populate(objectMap, "unit", m.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricDefinition.
func (m *MetricDefinition) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "category":
			err = unpopulate(val, "Category", &m.Category)
			delete(rawMsg, key)
		case "dimensions":
			err = unpopulate(val, "Dimensions", &m.Dimensions)
			delete(rawMsg, key)
		case "displayDescription":
			err = unpopulate(val, "DisplayDescription", &m.DisplayDescription)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &m.ID)
			delete(rawMsg, key)
		case "isDimensionRequired":
			err = unpopulate(val, "IsDimensionRequired", &m.IsDimensionRequired)
			delete(rawMsg, key)
		case "metricAvailabilities":
			err = unpopulate(val, "MetricAvailabilities", &m.MetricAvailabilities)
			delete(rawMsg, key)
		case "metricClass":
			err = unpopulate(val, "MetricClass", &m.MetricClass)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &m.Name)
			delete(rawMsg, key)
		case "namespace":
			err = unpopulate(val, "Namespace", &m.Namespace)
			delete(rawMsg, key)
		case "primaryAggregationType":
			err = unpopulate(val, "PrimaryAggregationType", &m.PrimaryAggregationType)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &m.ResourceID)
			delete(rawMsg, key)
		case "supportedAggregationTypes":
			err = unpopulate(val, "SupportedAggregationTypes", &m.SupportedAggregationTypes)
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

// MarshalJSON implements the json.Marshaller interface for type MetricDefinitionCollection.
func (m MetricDefinitionCollection) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", m.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricDefinitionCollection.
func (m *MetricDefinitionCollection) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &m.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricSettings.
func (m MetricSettings) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "category", m.Category)
	populate(objectMap, "enabled", m.Enabled)
	populate(objectMap, "retentionPolicy", m.RetentionPolicy)
	populate(objectMap, "timeGrain", m.TimeGrain)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricSettings.
func (m *MetricSettings) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "category":
			err = unpopulate(val, "Category", &m.Category)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, "Enabled", &m.Enabled)
			delete(rawMsg, key)
		case "retentionPolicy":
			err = unpopulate(val, "RetentionPolicy", &m.RetentionPolicy)
			delete(rawMsg, key)
		case "timeGrain":
			err = unpopulate(val, "TimeGrain", &m.TimeGrain)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MetricValue.
func (m MetricValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "average", m.Average)
	populate(objectMap, "count", m.Count)
	populate(objectMap, "maximum", m.Maximum)
	populate(objectMap, "minimum", m.Minimum)
	populateTimeRFC3339(objectMap, "timeStamp", m.TimeStamp)
	populate(objectMap, "total", m.Total)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MetricValue.
func (m *MetricValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "average":
			err = unpopulate(val, "Average", &m.Average)
			delete(rawMsg, key)
		case "count":
			err = unpopulate(val, "Count", &m.Count)
			delete(rawMsg, key)
		case "maximum":
			err = unpopulate(val, "Maximum", &m.Maximum)
			delete(rawMsg, key)
		case "minimum":
			err = unpopulate(val, "Minimum", &m.Minimum)
			delete(rawMsg, key)
		case "timeStamp":
			err = unpopulateTimeRFC3339(val, "TimeStamp", &m.TimeStamp)
			delete(rawMsg, key)
		case "total":
			err = unpopulate(val, "Total", &m.Total)
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
	objectMap := make(map[string]interface{})
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
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
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
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
	objectMap := make(map[string]interface{})
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

// MarshalJSON implements the json.Marshaller interface for type ProxyOnlyResource.
func (p ProxyOnlyResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProxyOnlyResource.
func (p *ProxyOnlyResource) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Response.
func (r Response) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "cost", r.Cost)
	populate(objectMap, "interval", r.Interval)
	populate(objectMap, "namespace", r.Namespace)
	populate(objectMap, "resourceregion", r.Resourceregion)
	populate(objectMap, "timespan", r.Timespan)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Response.
func (r *Response) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cost":
			err = unpopulate(val, "Cost", &r.Cost)
			delete(rawMsg, key)
		case "interval":
			err = unpopulate(val, "Interval", &r.Interval)
			delete(rawMsg, key)
		case "namespace":
			err = unpopulate(val, "Namespace", &r.Namespace)
			delete(rawMsg, key)
		case "resourceregion":
			err = unpopulate(val, "Resourceregion", &r.Resourceregion)
			delete(rawMsg, key)
		case "timespan":
			err = unpopulate(val, "Timespan", &r.Timespan)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RetentionPolicy.
func (r RetentionPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "days", r.Days)
	populate(objectMap, "enabled", r.Enabled)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RetentionPolicy.
func (r *RetentionPolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "days":
			err = unpopulate(val, "Days", &r.Days)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, "Enabled", &r.Enabled)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TimeSeriesElement.
func (t TimeSeriesElement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "data", t.Data)
	populate(objectMap, "metadatavalues", t.Metadatavalues)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TimeSeriesElement.
func (t *TimeSeriesElement) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "data":
			err = unpopulate(val, "Data", &t.Data)
			delete(rawMsg, key)
		case "metadatavalues":
			err = unpopulate(val, "Metadatavalues", &t.Metadatavalues)
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
