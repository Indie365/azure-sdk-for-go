package billing

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DetachPaymentMethodEligibilityErrorCode enumerates the values for detach payment method eligibility error
// code.
type DetachPaymentMethodEligibilityErrorCode string

const (
	// AzureSubscriptions ...
	AzureSubscriptions DetachPaymentMethodEligibilityErrorCode = "AzureSubscriptions"
	// OutstandingCharges ...
	OutstandingCharges DetachPaymentMethodEligibilityErrorCode = "OutstandingCharges"
	// PendingCharges ...
	PendingCharges DetachPaymentMethodEligibilityErrorCode = "PendingCharges"
	// RecurringCharges ...
	RecurringCharges DetachPaymentMethodEligibilityErrorCode = "RecurringCharges"
	// ReservedInstances ...
	ReservedInstances DetachPaymentMethodEligibilityErrorCode = "ReservedInstances"
)

// PossibleDetachPaymentMethodEligibilityErrorCodeValues returns an array of possible values for the DetachPaymentMethodEligibilityErrorCode const type.
func PossibleDetachPaymentMethodEligibilityErrorCodeValues() []DetachPaymentMethodEligibilityErrorCode {
	return []DetachPaymentMethodEligibilityErrorCode{AzureSubscriptions, OutstandingCharges, PendingCharges, RecurringCharges, ReservedInstances}
}

// PaymentMethodFamily enumerates the values for payment method family.
type PaymentMethodFamily string

const (
	// CheckWire ...
	CheckWire PaymentMethodFamily = "CheckWire"
	// CreditCard ...
	CreditCard PaymentMethodFamily = "CreditCard"
)

// PossiblePaymentMethodFamilyValues returns an array of possible values for the PaymentMethodFamily const type.
func PossiblePaymentMethodFamilyValues() []PaymentMethodFamily {
	return []PaymentMethodFamily{CheckWire, CreditCard}
}

// PaymentMethodStatus enumerates the values for payment method status.
type PaymentMethodStatus string

const (
	// Active ...
	Active PaymentMethodStatus = "active"
	// Inactive ...
	Inactive PaymentMethodStatus = "inactive"
)

// PossiblePaymentMethodStatusValues returns an array of possible values for the PaymentMethodStatus const type.
func PossiblePaymentMethodStatusValues() []PaymentMethodStatus {
	return []PaymentMethodStatus{Active, Inactive}
}
