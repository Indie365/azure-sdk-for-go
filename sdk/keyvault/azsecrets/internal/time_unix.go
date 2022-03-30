//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal



import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"strings"
	"time"
)



type timeUnix time.Time

func (t timeUnix) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix())
}

func (t *timeUnix) UnmarshalJSON(data []byte) error {
	var seconds int64
	if err := json.Unmarshal(data, &seconds); err != nil {
		return err
	}
	*t = timeUnix(time.Unix(seconds, 0))
	return nil
}

func (t timeUnix) String() string {
	return fmt.Sprintf("%d", time.Time(t).Unix())
}

func populateTimeUnix(m map[string]interface{}, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*timeUnix)(t)
}

func unpopulateTimeUnix(data json.RawMessage, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux timeUnix
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*t = (*time.Time)(&aux)
	return nil
}
