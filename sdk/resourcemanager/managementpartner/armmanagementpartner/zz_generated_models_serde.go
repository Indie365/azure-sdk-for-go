//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementpartner

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON implements the json.Unmarshaller interface for type PartnerProperties.
func (p *PartnerProperties) UnmarshalJSON(data []byte) error {
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
		case "objectId":
			err = unpopulate(val, "ObjectID", &p.ObjectID)
			delete(rawMsg, key)
		case "partnerId":
			err = unpopulate(val, "PartnerID", &p.PartnerID)
			delete(rawMsg, key)
		case "partnerName":
			err = unpopulate(val, "PartnerName", &p.PartnerName)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &p.State)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &p.TenantID)
			delete(rawMsg, key)
		case "updatedTime":
			err = unpopulateTimeRFC3339(val, "UpdatedTime", &p.UpdatedTime)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &p.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
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