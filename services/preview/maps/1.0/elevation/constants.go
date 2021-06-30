// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package elevation

const telemetryInfo = "azsdk-go-elevation/<version>"
type GeographicResourceLocation string

const (
	// GeographicResourceLocationEu - Used to access an Azure Maps Creator resource in Europe
	GeographicResourceLocationEu GeographicResourceLocation = "eu"
	// GeographicResourceLocationUs - Used to access an Azure Maps Creator resource in the United States
	GeographicResourceLocationUs GeographicResourceLocation = "us"
)

// PossibleGeographicResourceLocationValues returns the possible values for the GeographicResourceLocation const type.
func PossibleGeographicResourceLocationValues() []GeographicResourceLocation {
	return []GeographicResourceLocation{	
		GeographicResourceLocationEu,
		GeographicResourceLocationUs,
	}
}

// ToPtr returns a *GeographicResourceLocation pointing to the current value.
func (c GeographicResourceLocation) ToPtr() *GeographicResourceLocation {
	return &c
}

// Geography - This parameter specifies where the Azure Maps Creator resource is located. Valid values are us and eu.
type Geography string

const (
	GeographyEu Geography = "eu"
	GeographyUs Geography = "us"
)

// PossibleGeographyValues returns the possible values for the Geography const type.
func PossibleGeographyValues() []Geography {
	return []Geography{	
		GeographyEu,
		GeographyUs,
	}
}

// ToPtr returns a *Geography pointing to the current value.
func (c Geography) ToPtr() *Geography {
	return &c
}

type ResponseFormat string

const (
	// ResponseFormatJSON - [The JavaScript Object Notation Data Interchange Format](https://tools.ietf.org/html/rfc8259)
	ResponseFormatJSON ResponseFormat = "json"
)

// PossibleResponseFormatValues returns the possible values for the ResponseFormat const type.
func PossibleResponseFormatValues() []ResponseFormat {
	return []ResponseFormat{	
		ResponseFormatJSON,
	}
}

// ToPtr returns a *ResponseFormat pointing to the current value.
func (c ResponseFormat) ToPtr() *ResponseFormat {
	return &c
}

