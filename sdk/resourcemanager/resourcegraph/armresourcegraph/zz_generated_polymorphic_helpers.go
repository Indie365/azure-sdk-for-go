//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

import "encoding/json"

func unmarshalFacetClassification(rawMsg json.RawMessage) (FacetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FacetClassification
	switch m["resultType"] {
	case "FacetError":
		b = &FacetError{}
	case "FacetResult":
		b = &FacetResult{}
	default:
		b = &Facet{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFacetClassificationArray(rawMsg json.RawMessage) ([]FacetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FacetClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFacetClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalFacetClassificationMap(rawMsg json.RawMessage) (map[string]FacetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]FacetClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalFacetClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}
