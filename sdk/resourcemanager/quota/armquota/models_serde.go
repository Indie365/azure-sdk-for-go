//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armquota

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CommonResourceProperties.
func (c CommonResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CommonResourceProperties.
func (c *CommonResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type CreateGenericQuotaRequestParameters.
func (c CreateGenericQuotaRequestParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CreateGenericQuotaRequestParameters.
func (c *CreateGenericQuotaRequestParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &c.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CurrentQuotaLimitBase.
func (c CurrentQuotaLimitBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CurrentQuotaLimitBase.
func (c *CurrentQuotaLimitBase) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &c.Properties)
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

// MarshalJSON implements the json.Marshaller interface for type CurrentUsagesBase.
func (c CurrentUsagesBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CurrentUsagesBase.
func (c *CurrentUsagesBase) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &c.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &c.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &c.Properties)
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

// MarshalJSON implements the json.Marshaller interface for type ExceptionResponse.
func (e ExceptionResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExceptionResponse.
func (e *ExceptionResponse) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type LimitJSONObject.
func (l LimitJSONObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["limitObjectType"] = l.LimitObjectType
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LimitJSONObject.
func (l *LimitJSONObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limitObjectType":
			err = unpopulate(val, "LimitObjectType", &l.LimitObjectType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LimitObject.
func (l LimitObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["limitObjectType"] = LimitTypeLimitValue
	populate(objectMap, "limitType", l.LimitType)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LimitObject.
func (l *LimitObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limitObjectType":
			err = unpopulate(val, "LimitObjectType", &l.LimitObjectType)
			delete(rawMsg, key)
		case "limitType":
			err = unpopulate(val, "LimitType", &l.LimitType)
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

// MarshalJSON implements the json.Marshaller interface for type Limits.
func (l Limits) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Limits.
func (l *Limits) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &l.NextLink)
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

// MarshalJSON implements the json.Marshaller interface for type LimitsResponse.
func (l LimitsResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LimitsResponse.
func (l *LimitsResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &l.NextLink)
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

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationList.
func (o *OperationList) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type OperationResponse.
func (o OperationResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "origin", o.Origin)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationResponse.
func (o *OperationResponse) UnmarshalJSON(data []byte) error {
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
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "isQuotaApplicable", p.IsQuotaApplicable)
	populate(objectMap, "limit", p.Limit)
	populate(objectMap, "name", p.Name)
	populateAny(objectMap, "properties", p.Properties)
	populate(objectMap, "quotaPeriod", p.QuotaPeriod)
	populate(objectMap, "resourceType", p.ResourceType)
	populate(objectMap, "unit", p.Unit)
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
		case "isQuotaApplicable":
			err = unpopulate(val, "IsQuotaApplicable", &p.IsQuotaApplicable)
			delete(rawMsg, key)
		case "limit":
			p.Limit, err = unmarshalLimitJSONObjectClassification(val)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, "QuotaPeriod", &p.QuotaPeriod)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &p.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &p.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RequestDetails.
func (r RequestDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestDetails.
func (r *RequestDetails) UnmarshalJSON(data []byte) error {
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
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
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

// MarshalJSON implements the json.Marshaller interface for type RequestDetailsList.
func (r RequestDetailsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestDetailsList.
func (r *RequestDetailsList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &r.NextLink)
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

// MarshalJSON implements the json.Marshaller interface for type RequestOneResourceProperties.
func (r RequestOneResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "currentValue", r.CurrentValue)
	populate(objectMap, "error", r.Error)
	populate(objectMap, "isQuotaApplicable", r.IsQuotaApplicable)
	populate(objectMap, "limit", r.Limit)
	populate(objectMap, "message", r.Message)
	populate(objectMap, "name", r.Name)
	populateAny(objectMap, "properties", r.Properties)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "quotaPeriod", r.QuotaPeriod)
	populateTimeRFC3339(objectMap, "requestSubmitTime", r.RequestSubmitTime)
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "unit", r.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestOneResourceProperties.
func (r *RequestOneResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currentValue":
			err = unpopulate(val, "CurrentValue", &r.CurrentValue)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, "Error", &r.Error)
			delete(rawMsg, key)
		case "isQuotaApplicable":
			err = unpopulate(val, "IsQuotaApplicable", &r.IsQuotaApplicable)
			delete(rawMsg, key)
		case "limit":
			err = unpopulate(val, "Limit", &r.Limit)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &r.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, "QuotaPeriod", &r.QuotaPeriod)
			delete(rawMsg, key)
		case "requestSubmitTime":
			err = unpopulateTimeRFC3339(val, "RequestSubmitTime", &r.RequestSubmitTime)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &r.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &r.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RequestOneResourceSubmitResponse.
func (r RequestOneResourceSubmitResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestOneResourceSubmitResponse.
func (r *RequestOneResourceSubmitResponse) UnmarshalJSON(data []byte) error {
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
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
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

// MarshalJSON implements the json.Marshaller interface for type RequestProperties.
func (r RequestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", r.Error)
	populate(objectMap, "message", r.Message)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populateTimeRFC3339(objectMap, "requestSubmitTime", r.RequestSubmitTime)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestProperties.
func (r *RequestProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &r.Error)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &r.Message)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "requestSubmitTime":
			err = unpopulateTimeRFC3339(val, "RequestSubmitTime", &r.RequestSubmitTime)
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

// MarshalJSON implements the json.Marshaller interface for type RequestStatusDetails.
func (r RequestStatusDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "limit", r.Limit)
	populate(objectMap, "message", r.Message)
	populate(objectMap, "name", r.Name)
	populateAny(objectMap, "properties", r.Properties)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "quotaPeriod", r.QuotaPeriod)
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "unit", r.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestStatusDetails.
func (r *RequestStatusDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limit":
			err = unpopulate(val, "Limit", &r.Limit)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &r.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &r.ProvisioningState)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, "QuotaPeriod", &r.QuotaPeriod)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &r.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &r.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RequestSubmitResponse.
func (r RequestSubmitResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestSubmitResponse.
func (r *RequestSubmitResponse) UnmarshalJSON(data []byte) error {
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
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
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

// MarshalJSON implements the json.Marshaller interface for type RequestSubmitResponse202.
func (r RequestSubmitResponse202) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestSubmitResponse202.
func (r *RequestSubmitResponse202) UnmarshalJSON(data []byte) error {
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
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &r.Properties)
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

// MarshalJSON implements the json.Marshaller interface for type ResourceName.
func (r ResourceName) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "localizedValue", r.LocalizedValue)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ResourceName.
func (r *ResourceName) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "localizedValue":
			err = unpopulate(val, "LocalizedValue", &r.LocalizedValue)
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

// MarshalJSON implements the json.Marshaller interface for type ServiceError.
func (s ServiceError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", s.Code)
	populate(objectMap, "details", s.Details)
	populate(objectMap, "message", s.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceError.
func (s *ServiceError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &s.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &s.Details)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ServiceErrorDetail.
func (s ServiceErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", s.Code)
	populate(objectMap, "message", s.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceErrorDetail.
func (s *ServiceErrorDetail) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &s.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SubRequest.
func (s SubRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "limit", s.Limit)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "resourceType", s.ResourceType)
	populate(objectMap, "subRequestId", s.SubRequestID)
	populate(objectMap, "unit", s.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubRequest.
func (s *SubRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limit":
			s.Limit, err = unmarshalLimitJSONObjectClassification(val)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &s.ResourceType)
			delete(rawMsg, key)
		case "subRequestId":
			err = unpopulate(val, "SubRequestID", &s.SubRequestID)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &s.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UsagesLimits.
func (u UsagesLimits) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", u.NextLink)
	populate(objectMap, "value", u.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UsagesLimits.
func (u *UsagesLimits) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &u.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &u.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UsagesObject.
func (u UsagesObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "usagesType", u.UsagesType)
	populate(objectMap, "value", u.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UsagesObject.
func (u *UsagesObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "usagesType":
			err = unpopulate(val, "UsagesType", &u.UsagesType)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &u.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UsagesProperties.
func (u UsagesProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "isQuotaApplicable", u.IsQuotaApplicable)
	populate(objectMap, "name", u.Name)
	populateAny(objectMap, "properties", u.Properties)
	populate(objectMap, "quotaPeriod", u.QuotaPeriod)
	populate(objectMap, "resourceType", u.ResourceType)
	populate(objectMap, "unit", u.Unit)
	populate(objectMap, "usages", u.Usages)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UsagesProperties.
func (u *UsagesProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "isQuotaApplicable":
			err = unpopulate(val, "IsQuotaApplicable", &u.IsQuotaApplicable)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &u.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &u.Properties)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, "QuotaPeriod", &u.QuotaPeriod)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &u.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, "Unit", &u.Unit)
			delete(rawMsg, key)
		case "usages":
			err = unpopulate(val, "Usages", &u.Usages)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
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
