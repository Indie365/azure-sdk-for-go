package quota

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// LimitObjectType enumerates the values for limit object type.
type LimitObjectType string

const (
	// LimitObjectTypeLimitJSONObject ...
	LimitObjectTypeLimitJSONObject LimitObjectType = "LimitJsonObject"
)

// PossibleLimitObjectTypeValues returns an array of possible values for the LimitObjectType const type.
func PossibleLimitObjectTypeValues() []LimitObjectType {
	return []LimitObjectType{LimitObjectTypeLimitJSONObject}
}

// LimitType enumerates the values for limit type.
type LimitType string

const (
	// LimitTypeLimitValue ...
	LimitTypeLimitValue LimitType = "LimitValue"
)

// PossibleLimitTypeValues returns an array of possible values for the LimitType const type.
func PossibleLimitTypeValues() []LimitType {
	return []LimitType{LimitTypeLimitValue}
}

// LimitTypes enumerates the values for limit types.
type LimitTypes string

const (
	// Independent ...
	Independent LimitTypes = "Independent"
	// Shared ...
	Shared LimitTypes = "Shared"
)

// PossibleLimitTypesValues returns an array of possible values for the LimitTypes const type.
func PossibleLimitTypesValues() []LimitTypes {
	return []LimitTypes{Independent, Shared}
}

// RequestState enumerates the values for request state.
type RequestState string

const (
	// Accepted ...
	Accepted RequestState = "Accepted"
	// Failed ...
	Failed RequestState = "Failed"
	// InProgress ...
	InProgress RequestState = "InProgress"
	// Invalid ...
	Invalid RequestState = "Invalid"
	// Succeeded ...
	Succeeded RequestState = "Succeeded"
)

// PossibleRequestStateValues returns an array of possible values for the RequestState const type.
func PossibleRequestStateValues() []RequestState {
	return []RequestState{Accepted, Failed, InProgress, Invalid, Succeeded}
}

// UsagesTypes enumerates the values for usages types.
type UsagesTypes string

const (
	// Combined ...
	Combined UsagesTypes = "Combined"
	// Individual ...
	Individual UsagesTypes = "Individual"
)

// PossibleUsagesTypesValues returns an array of possible values for the UsagesTypes const type.
func PossibleUsagesTypesValues() []UsagesTypes {
	return []UsagesTypes{Combined, Individual}
}
