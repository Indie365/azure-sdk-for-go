//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsubscriptions

const (
	moduleName    = "hybrid20200901/armsubscriptions"
	moduleVersion = "v1.0.0"
)

// ResourceNameStatus - Is the resource name Allowed or Reserved
type ResourceNameStatus string

const (
	ResourceNameStatusAllowed  ResourceNameStatus = "Allowed"
	ResourceNameStatusReserved ResourceNameStatus = "Reserved"
)

// PossibleResourceNameStatusValues returns the possible values for the ResourceNameStatus const type.
func PossibleResourceNameStatusValues() []ResourceNameStatus {
	return []ResourceNameStatus{
		ResourceNameStatusAllowed,
		ResourceNameStatusReserved,
	}
}

// SpendingLimit - The subscription spending limit.
type SpendingLimit string

const (
	SpendingLimitOn               SpendingLimit = "On"
	SpendingLimitOff              SpendingLimit = "Off"
	SpendingLimitCurrentPeriodOff SpendingLimit = "CurrentPeriodOff"
)

// PossibleSpendingLimitValues returns the possible values for the SpendingLimit const type.
func PossibleSpendingLimitValues() []SpendingLimit {
	return []SpendingLimit{
		SpendingLimitOn,
		SpendingLimitOff,
		SpendingLimitCurrentPeriodOff,
	}
}

// SubscriptionState - The subscription state. Possible values are Enabled, Warned, PastDue, Disabled, and Deleted.
type SubscriptionState string

const (
	SubscriptionStateEnabled  SubscriptionState = "Enabled"
	SubscriptionStateWarned   SubscriptionState = "Warned"
	SubscriptionStatePastDue  SubscriptionState = "PastDue"
	SubscriptionStateDisabled SubscriptionState = "Disabled"
	SubscriptionStateDeleted  SubscriptionState = "Deleted"
)

// PossibleSubscriptionStateValues returns the possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{
		SubscriptionStateEnabled,
		SubscriptionStateWarned,
		SubscriptionStatePastDue,
		SubscriptionStateDisabled,
		SubscriptionStateDeleted,
	}
}
