// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package timezone

import "net/http"

// IanaIDArrayResponse is the response envelope for operations that return a []*IanaID type.
type IanaIDArrayResponse struct {
	// This object is returned from a successful Timezone Enum IANA call
	IanaIDArray []*IanaID

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TimezoneByCoordinatesResultResponse is the response envelope for operations that return a TimezoneByCoordinatesResult type.
type TimezoneByCoordinatesResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// This object is returned from a successful Timezone By Coordinates call
	TimezoneByCoordinatesResult *TimezoneByCoordinatesResult
}

// TimezoneByIDResultResponse is the response envelope for operations that return a TimezoneByIDResult type.
type TimezoneByIDResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// This object is returned from a successful Timezone By ID call
	TimezoneByIDResult *TimezoneByIDResult
}

// TimezoneEnumWindowArrayResponse is the response envelope for operations that return a []*TimezoneEnumWindow type.
type TimezoneEnumWindowArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// This object is returned from a successful Timezone Enum Windows call
	TimezoneEnumWindowArray []*TimezoneEnumWindow
}

// TimezoneIanaVersionResultResponse is the response envelope for operations that return a TimezoneIanaVersionResult type.
type TimezoneIanaVersionResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// This object is returned from a successful Timezone IANA Version call
	TimezoneIanaVersionResult *TimezoneIanaVersionResult
}

