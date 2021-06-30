// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package search

import "encoding/json"

func unmarshalGeoJSONGeometryClassification(rawMsg json.RawMessage) (GeoJSONGeometryClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b GeoJSONGeometryClassification
	switch m["type"] {
	case string(GeoJSONObjectTypeGeoJSONGeometryCollection):
		b = &GeoJSONGeometryCollection{}
	case string(GeoJSONObjectTypeGeoJSONLineString):
		b = &GeoJSONLineString{}
	case string(GeoJSONObjectTypeGeoJSONMultiLineString):
		b = &GeoJSONMultiLineString{}
	case string(GeoJSONObjectTypeGeoJSONMultiPoint):
		b = &GeoJSONMultiPoint{}
	case string(GeoJSONObjectTypeGeoJSONMultiPolygon):
		b = &GeoJSONMultiPolygon{}
	case string(GeoJSONObjectTypeGeoJSONPoint):
		b = &GeoJSONPoint{}
	case string(GeoJSONObjectTypeGeoJSONPolygon):
		b = &GeoJSONPolygon{}
	default:
		b = &GeoJSONGeometry{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalGeoJSONGeometryClassificationArray(rawMsg json.RawMessage) ([]GeoJSONGeometryClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]GeoJSONGeometryClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalGeoJSONGeometryClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalGeoJSONObjectClassification(rawMsg json.RawMessage) (GeoJSONObjectClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b GeoJSONObjectClassification
	switch m["type"] {
	case string(GeoJSONObjectTypeGeoJSONFeature):
		b = &GeoJSONFeature{}
	case string(GeoJSONObjectTypeGeoJSONFeatureCollection):
		b = &GeoJSONFeatureCollection{}
	case "GeoJsonGeometry":
		b = &GeoJSONGeometry{}
	case string(GeoJSONObjectTypeGeoJSONGeometryCollection):
		b = &GeoJSONGeometryCollection{}
	case string(GeoJSONObjectTypeGeoJSONLineString):
		b = &GeoJSONLineString{}
	case string(GeoJSONObjectTypeGeoJSONMultiLineString):
		b = &GeoJSONMultiLineString{}
	case string(GeoJSONObjectTypeGeoJSONMultiPoint):
		b = &GeoJSONMultiPoint{}
	case string(GeoJSONObjectTypeGeoJSONMultiPolygon):
		b = &GeoJSONMultiPolygon{}
	case string(GeoJSONObjectTypeGeoJSONPoint):
		b = &GeoJSONPoint{}
	case string(GeoJSONObjectTypeGeoJSONPolygon):
		b = &GeoJSONPolygon{}
	default:
		b = &GeoJSONObject{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalGeoJSONObjectClassificationArray(rawMsg json.RawMessage) ([]GeoJSONObjectClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]GeoJSONObjectClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalGeoJSONObjectClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

