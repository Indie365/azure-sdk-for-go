package features

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// SubscriptionFeatureRegistrationApprovalType enumerates the values for subscription feature registration
// approval type.
type SubscriptionFeatureRegistrationApprovalType string

const (
	// ApprovalRequired ...
	ApprovalRequired SubscriptionFeatureRegistrationApprovalType = "ApprovalRequired"
	// AutoApproval ...
	AutoApproval SubscriptionFeatureRegistrationApprovalType = "AutoApproval"
	// NotSpecified ...
	NotSpecified SubscriptionFeatureRegistrationApprovalType = "NotSpecified"
)

// PossibleSubscriptionFeatureRegistrationApprovalTypeValues returns an array of possible values for the SubscriptionFeatureRegistrationApprovalType const type.
func PossibleSubscriptionFeatureRegistrationApprovalTypeValues() []SubscriptionFeatureRegistrationApprovalType {
	return []SubscriptionFeatureRegistrationApprovalType{ApprovalRequired, AutoApproval, NotSpecified}
}

// SubscriptionFeatureRegistrationState enumerates the values for subscription feature registration state.
type SubscriptionFeatureRegistrationState string

const (
	// SubscriptionFeatureRegistrationStateNotRegistered ...
	SubscriptionFeatureRegistrationStateNotRegistered SubscriptionFeatureRegistrationState = "NotRegistered"
	// SubscriptionFeatureRegistrationStateNotSpecified ...
	SubscriptionFeatureRegistrationStateNotSpecified SubscriptionFeatureRegistrationState = "NotSpecified"
	// SubscriptionFeatureRegistrationStatePending ...
	SubscriptionFeatureRegistrationStatePending SubscriptionFeatureRegistrationState = "Pending"
	// SubscriptionFeatureRegistrationStateRegistered ...
	SubscriptionFeatureRegistrationStateRegistered SubscriptionFeatureRegistrationState = "Registered"
	// SubscriptionFeatureRegistrationStateRegistering ...
	SubscriptionFeatureRegistrationStateRegistering SubscriptionFeatureRegistrationState = "Registering"
	// SubscriptionFeatureRegistrationStateUnregistered ...
	SubscriptionFeatureRegistrationStateUnregistered SubscriptionFeatureRegistrationState = "Unregistered"
	// SubscriptionFeatureRegistrationStateUnregistering ...
	SubscriptionFeatureRegistrationStateUnregistering SubscriptionFeatureRegistrationState = "Unregistering"
)

// PossibleSubscriptionFeatureRegistrationStateValues returns an array of possible values for the SubscriptionFeatureRegistrationState const type.
func PossibleSubscriptionFeatureRegistrationStateValues() []SubscriptionFeatureRegistrationState {
	return []SubscriptionFeatureRegistrationState{SubscriptionFeatureRegistrationStateNotRegistered, SubscriptionFeatureRegistrationStateNotSpecified, SubscriptionFeatureRegistrationStatePending, SubscriptionFeatureRegistrationStateRegistered, SubscriptionFeatureRegistrationStateRegistering, SubscriptionFeatureRegistrationStateUnregistered, SubscriptionFeatureRegistrationStateUnregistering}
}
