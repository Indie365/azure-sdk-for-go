//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azsecrets

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type BackupSecretResult.
func (b BackupSecretResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", b.Value, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BackupSecretResult.
func (b *BackupSecretResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = runtime.DecodeByteArray(string(val), &b.Value, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeletedSecret.
func (d DeletedSecret) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", d.Attributes)
	populate(objectMap, "contentType", d.ContentType)
	populateTimeUnix(objectMap, "deletedDate", d.DeletedDate)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "kid", d.KeyID)
	populate(objectMap, "managed", d.Managed)
	populate(objectMap, "recoveryId", d.RecoveryID)
	populateTimeUnix(objectMap, "scheduledPurgeDate", d.ScheduledPurgeDate)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedSecret.
func (d *DeletedSecret) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &d.Attributes)
			delete(rawMsg, key)
		case "contentType":
			err = unpopulate(val, "ContentType", &d.ContentType)
			delete(rawMsg, key)
		case "deletedDate":
			err = unpopulateTimeUnix(val, "DeletedDate", &d.DeletedDate)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "kid":
			err = unpopulate(val, "KeyID", &d.KeyID)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &d.Managed)
			delete(rawMsg, key)
		case "recoveryId":
			err = unpopulate(val, "RecoveryID", &d.RecoveryID)
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			err = unpopulateTimeUnix(val, "ScheduledPurgeDate", &d.ScheduledPurgeDate)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &d.Tags)
			delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type DeletedSecretProperties.
func (d DeletedSecretProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", d.Attributes)
	populate(objectMap, "contentType", d.ContentType)
	populateTimeUnix(objectMap, "deletedDate", d.DeletedDate)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "managed", d.Managed)
	populate(objectMap, "recoveryId", d.RecoveryID)
	populateTimeUnix(objectMap, "scheduledPurgeDate", d.ScheduledPurgeDate)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedSecretProperties.
func (d *DeletedSecretProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &d.Attributes)
			delete(rawMsg, key)
		case "contentType":
			err = unpopulate(val, "ContentType", &d.ContentType)
			delete(rawMsg, key)
		case "deletedDate":
			err = unpopulateTimeUnix(val, "DeletedDate", &d.DeletedDate)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &d.Managed)
			delete(rawMsg, key)
		case "recoveryId":
			err = unpopulate(val, "RecoveryID", &d.RecoveryID)
			delete(rawMsg, key)
		case "scheduledPurgeDate":
			err = unpopulateTimeUnix(val, "ScheduledPurgeDate", &d.ScheduledPurgeDate)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &d.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DeletedSecretPropertiesListResult.
func (d DeletedSecretPropertiesListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", d.NextLink)
	populate(objectMap, "value", d.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeletedSecretPropertiesListResult.
func (d *DeletedSecretPropertiesListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &d.NextLink)
			delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type RestoreSecretParameters.
func (r RestoreSecretParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "value", r.SecretBackup, runtime.Base64URLFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RestoreSecretParameters.
func (r *RestoreSecretParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = runtime.DecodeByteArray(string(val), &r.SecretBackup, runtime.Base64URLFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Secret.
func (s Secret) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", s.Attributes)
	populate(objectMap, "contentType", s.ContentType)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "kid", s.KeyID)
	populate(objectMap, "managed", s.Managed)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Secret.
func (s *Secret) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &s.Attributes)
			delete(rawMsg, key)
		case "contentType":
			err = unpopulate(val, "ContentType", &s.ContentType)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "kid":
			err = unpopulate(val, "KeyID", &s.KeyID)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &s.Managed)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type SecretAttributes.
func (s SecretAttributes) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeUnix(objectMap, "created", s.Created)
	populate(objectMap, "enabled", s.Enabled)
	populateTimeUnix(objectMap, "exp", s.Expires)
	populateTimeUnix(objectMap, "nbf", s.NotBefore)
	populate(objectMap, "recoverableDays", s.RecoverableDays)
	populate(objectMap, "recoveryLevel", s.RecoveryLevel)
	populateTimeUnix(objectMap, "updated", s.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecretAttributes.
func (s *SecretAttributes) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "created":
			err = unpopulateTimeUnix(val, "Created", &s.Created)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, "Enabled", &s.Enabled)
			delete(rawMsg, key)
		case "exp":
			err = unpopulateTimeUnix(val, "Expires", &s.Expires)
			delete(rawMsg, key)
		case "nbf":
			err = unpopulateTimeUnix(val, "NotBefore", &s.NotBefore)
			delete(rawMsg, key)
		case "recoverableDays":
			err = unpopulate(val, "RecoverableDays", &s.RecoverableDays)
			delete(rawMsg, key)
		case "recoveryLevel":
			err = unpopulate(val, "RecoveryLevel", &s.RecoveryLevel)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeUnix(val, "Updated", &s.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SecretProperties.
func (s SecretProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "attributes", s.Attributes)
	populate(objectMap, "contentType", s.ContentType)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "managed", s.Managed)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecretProperties.
func (s *SecretProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "attributes":
			err = unpopulate(val, "Attributes", &s.Attributes)
			delete(rawMsg, key)
		case "contentType":
			err = unpopulate(val, "ContentType", &s.ContentType)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "managed":
			err = unpopulate(val, "Managed", &s.Managed)
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

// MarshalJSON implements the json.Marshaller interface for type SecretPropertiesListResult.
func (s SecretPropertiesListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecretPropertiesListResult.
func (s *SecretPropertiesListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &s.NextLink)
			delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type SetSecretParameters.
func (s SetSecretParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "contentType", s.ContentType)
	populate(objectMap, "attributes", s.SecretAttributes)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SetSecretParameters.
func (s *SetSecretParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, "ContentType", &s.ContentType)
			delete(rawMsg, key)
		case "attributes":
			err = unpopulate(val, "SecretAttributes", &s.SecretAttributes)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
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

// MarshalJSON implements the json.Marshaller interface for type UpdateSecretPropertiesParameters.
func (u UpdateSecretPropertiesParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "contentType", u.ContentType)
	populate(objectMap, "attributes", u.SecretAttributes)
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateSecretPropertiesParameters.
func (u *UpdateSecretPropertiesParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, "ContentType", &u.ContentType)
			delete(rawMsg, key)
		case "attributes":
			err = unpopulate(val, "SecretAttributes", &u.SecretAttributes)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &u.Tags)
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

func populateByteArray(m map[string]any, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
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
