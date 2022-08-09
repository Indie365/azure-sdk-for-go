//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azquery

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BatchQueryRequest.
func (b BatchQueryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "body", b.Body)
	populate(objectMap, "headers", b.Headers)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "method", b.Method)
	populate(objectMap, "path", b.Path)
	populate(objectMap, "workspace", b.Workspace)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchQueryRequest.
func (b *BatchQueryRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "body":
				err = unpopulate(val, "Body", &b.Body)
				delete(rawMsg, key)
		case "headers":
				err = unpopulate(val, "Headers", &b.Headers)
				delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &b.ID)
				delete(rawMsg, key)
		case "method":
				err = unpopulate(val, "Method", &b.Method)
				delete(rawMsg, key)
		case "path":
				err = unpopulate(val, "Path", &b.Path)
				delete(rawMsg, key)
		case "workspace":
				err = unpopulate(val, "Workspace", &b.Workspace)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchQueryResponse.
func (b BatchQueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "body", b.Body)
	populate(objectMap, "headers", b.Headers)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "status", b.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchQueryResponse.
func (b *BatchQueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "body":
				err = unpopulate(val, "Body", &b.Body)
				delete(rawMsg, key)
		case "headers":
				err = unpopulate(val, "Headers", &b.Headers)
				delete(rawMsg, key)
		case "id":
				err = unpopulate(val, "ID", &b.ID)
				delete(rawMsg, key)
		case "status":
				err = unpopulate(val, "Status", &b.Status)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchQueryResults.
func (b BatchQueryResults) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", b.Error)
	populate(objectMap, "render", &b.Render)
	populate(objectMap, "statistics", &b.Statistics)
	populate(objectMap, "tables", b.Tables)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchQueryResults.
func (b *BatchQueryResults) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
				err = unpopulate(val, "Error", &b.Error)
				delete(rawMsg, key)
		case "render":
				err = unpopulate(val, "Render", &b.Render)
				delete(rawMsg, key)
		case "statistics":
				err = unpopulate(val, "Statistics", &b.Statistics)
				delete(rawMsg, key)
		case "tables":
				err = unpopulate(val, "Tables", &b.Tables)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchRequest.
func (b BatchRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "requests", b.Requests)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchRequest.
func (b *BatchRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "requests":
				err = unpopulate(val, "Requests", &b.Requests)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BatchResponse.
func (b BatchResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "responses", b.Responses)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BatchResponse.
func (b *BatchResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "responses":
				err = unpopulate(val, "Responses", &b.Responses)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Body.
func (b Body) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "query", b.Query)
	populate(objectMap, "timespan", b.Timespan)
	populate(objectMap, "workspaces", b.Workspaces)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Body.
func (b *Body) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "query":
				err = unpopulate(val, "Query", &b.Query)
				delete(rawMsg, key)
		case "timespan":
				err = unpopulate(val, "Timespan", &b.Timespan)
				delete(rawMsg, key)
		case "workspaces":
				err = unpopulate(val, "Workspaces", &b.Workspaces)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Column.
func (c Column) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", c.Name)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Column.
func (c *Column) UnmarshalJSON(data []byte) error {
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

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalProperties", &e.AdditionalProperties)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "resources", e.Resources)
	populate(objectMap, "target", e.Target)
	populate(objectMap, "value", e.Value)
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
		case "additionalProperties":
				err = unpopulate(val, "AdditionalProperties", &e.AdditionalProperties)
				delete(rawMsg, key)
		case "code":
				err = unpopulate(val, "Code", &e.Code)
				delete(rawMsg, key)
		case "message":
				err = unpopulate(val, "Message", &e.Message)
				delete(rawMsg, key)
		case "resources":
				err = unpopulate(val, "Resources", &e.Resources)
				delete(rawMsg, key)
		case "target":
				err = unpopulate(val, "Target", &e.Target)
				delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type ErrorInfo.
func (e ErrorInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalProperties", &e.AdditionalProperties)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "innererror", e.Innererror)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorInfo.
func (e *ErrorInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
				err = unpopulate(val, "AdditionalProperties", &e.AdditionalProperties)
				delete(rawMsg, key)
		case "code":
				err = unpopulate(val, "Code", &e.Code)
				delete(rawMsg, key)
		case "details":
				err = unpopulate(val, "Details", &e.Details)
				delete(rawMsg, key)
		case "innererror":
				err = unpopulate(val, "Innererror", &e.Innererror)
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

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
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

// MarshalJSON implements the json.Marshaller interface for type Results.
func (r Results) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", r.Error)
	populate(objectMap, "render", &r.Render)
	populate(objectMap, "statistics", &r.Statistics)
	populate(objectMap, "tables", r.Tables)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Results.
func (r *Results) UnmarshalJSON(data []byte) error {
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
		case "render":
				err = unpopulate(val, "Render", &r.Render)
				delete(rawMsg, key)
		case "statistics":
				err = unpopulate(val, "Statistics", &r.Statistics)
				delete(rawMsg, key)
		case "tables":
				err = unpopulate(val, "Tables", &r.Tables)
				delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Table.
func (t Table) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "columns", t.Columns)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "rows", t.Rows)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Table.
func (t *Table) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "columns":
				err = unpopulate(val, "Columns", &t.Columns)
				delete(rawMsg, key)
		case "name":
				err = unpopulate(val, "Name", &t.Name)
				delete(rawMsg, key)
		case "rows":
				err = unpopulate(val, "Rows", &t.Rows)
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

