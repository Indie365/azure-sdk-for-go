package quota

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ResourceQueryMethod enumerates the values for resource query method.
type ResourceQueryMethod string

const (
	// GET ...
	GET ResourceQueryMethod = "GET"
	// POST ...
	POST ResourceQueryMethod = "POST"
)

// PossibleResourceQueryMethodValues returns an array of possible values for the ResourceQueryMethod const type.
func PossibleResourceQueryMethodValues() []ResourceQueryMethod {
	return []ResourceQueryMethod{GET, POST}
}

// ResourceQueryType enumerates the values for resource query type.
type ResourceQueryType string

const (
	// ARG ...
	ARG ResourceQueryType = "ARG"
	// RestAPI ...
	RestAPI ResourceQueryType = "RestAPI"
)

// PossibleResourceQueryTypeValues returns an array of possible values for the ResourceQueryType const type.
func PossibleResourceQueryTypeValues() []ResourceQueryType {
	return []ResourceQueryType{ARG, RestAPI}
}

// ResourceType enumerates the values for resource type.
type ResourceType string

const (
	// Dedicated ...
	Dedicated ResourceType = "dedicated"
	// LowPriority ...
	LowPriority ResourceType = "lowPriority"
	// ServiceSpecific ...
	ServiceSpecific ResourceType = "serviceSpecific"
	// Shared ...
	Shared ResourceType = "shared"
	// Standard ...
	Standard ResourceType = "standard"
)

// PossibleResourceTypeValues returns an array of possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{Dedicated, LowPriority, ServiceSpecific, Shared, Standard}
}
