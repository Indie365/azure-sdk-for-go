//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommerce

import "time"

// ErrorResponse - Describes the format of Error response.
type ErrorResponse struct {
	// Error code
	Code *string `json:"code,omitempty"`

	// Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// InfoField - Key-value pairs of instance details in the legacy format.
type InfoField struct {
	// Identifies the name of the instance provisioned by the user.
	Project *string `json:"project,omitempty"`
}

// MeterInfo - Detailed information about the meter.
type MeterInfo struct {
	// Indicates the date from which the meter rate is effective.
	EffectiveDate *time.Time `json:"EffectiveDate,omitempty"`

	// The resource quantity that is included in the offer at no cost. Consumption beyond this quantity will be charged.
	IncludedQuantity *float32 `json:"IncludedQuantity,omitempty"`

	// The category of the meter, e.g., 'Cloud services', 'Networking', etc..
	MeterCategory *string `json:"MeterCategory,omitempty"`

	// The unique identifier of the resource.
	MeterID *string `json:"MeterId,omitempty"`

	// The name of the meter, within the given meter category
	MeterName *string `json:"MeterName,omitempty"`

	// The list of key/value pairs for the meter rates, in the format 'key':'value' where key = the meter quantity, and value
	// = the corresponding price
	MeterRates map[string]*float32 `json:"MeterRates,omitempty"`

	// The region in which the Azure service is available.
	MeterRegion *string `json:"MeterRegion,omitempty"`

	// The subcategory of the meter, e.g., 'A6 Cloud services', 'ExpressRoute (IXP)', etc..
	MeterSubCategory *string `json:"MeterSubCategory,omitempty"`

	// Provides additional meter data. 'Third Party' indicates a meter with no discount. Blanks indicate First Party.
	MeterTags []*string `json:"MeterTags,omitempty"`

	// The unit in which the meter consumption is charged, e.g., 'Hours', 'GB', etc.
	Unit *string `json:"Unit,omitempty"`
}

// MonetaryCommitment - Indicates that a monetary commitment is required for this offer
type MonetaryCommitment struct {
	// REQUIRED; Name of the offer term
	Name *OfferTermInfo `json:"Name,omitempty"`

	// Indicates the date from which the offer term is effective.
	EffectiveDate *time.Time `json:"EffectiveDate,omitempty"`

	// An array of meter ids that are excluded from the given offer terms.
	ExcludedMeterIDs []*string `json:"ExcludedMeterIds,omitempty"`

	// The list of key/value pairs for the tiered meter rates, in the format 'key':'value' where key = price, and value = the
	// corresponding discount percentage. This field is used only by offer terms of type
	// 'Monetary Commitment'.
	TieredDiscount map[string]*float64 `json:"TieredDiscount,omitempty"`
}

// GetOfferTermInfoAutoGenerated implements the OfferTermInfoAutoGeneratedClassification interface for type MonetaryCommitment.
func (m *MonetaryCommitment) GetOfferTermInfoAutoGenerated() *OfferTermInfoAutoGenerated {
	return &OfferTermInfoAutoGenerated{
		Name:          m.Name,
		EffectiveDate: m.EffectiveDate,
	}
}

// MonetaryCredit - Indicates that this is a monetary credit offer.
type MonetaryCredit struct {
	// REQUIRED; Name of the offer term
	Name *OfferTermInfo `json:"Name,omitempty"`

	// The amount of credit provided under the terms of the given offer level.
	Credit *float64 `json:"Credit,omitempty"`

	// Indicates the date from which the offer term is effective.
	EffectiveDate *time.Time `json:"EffectiveDate,omitempty"`

	// An array of meter ids that are excluded from the given offer terms.
	ExcludedMeterIDs []*string `json:"ExcludedMeterIds,omitempty"`
}

// GetOfferTermInfoAutoGenerated implements the OfferTermInfoAutoGeneratedClassification interface for type MonetaryCredit.
func (m *MonetaryCredit) GetOfferTermInfoAutoGenerated() *OfferTermInfoAutoGenerated {
	return &OfferTermInfoAutoGenerated{
		Name:          m.Name,
		EffectiveDate: m.EffectiveDate,
	}
}

// OfferTermInfoAutoGeneratedClassification provides polymorphic access to related types.
// Call the interface's GetOfferTermInfoAutoGenerated() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MonetaryCommitment, *MonetaryCredit, *OfferTermInfoAutoGenerated, *RecurringCharge
type OfferTermInfoAutoGeneratedClassification interface {
	// GetOfferTermInfoAutoGenerated returns the OfferTermInfoAutoGenerated content of the underlying type.
	GetOfferTermInfoAutoGenerated() *OfferTermInfoAutoGenerated
}

// OfferTermInfoAutoGenerated - Describes the offer term.
type OfferTermInfoAutoGenerated struct {
	// REQUIRED; Name of the offer term
	Name *OfferTermInfo `json:"Name,omitempty"`

	// Indicates the date from which the offer term is effective.
	EffectiveDate *time.Time `json:"EffectiveDate,omitempty"`
}

// GetOfferTermInfoAutoGenerated implements the OfferTermInfoAutoGeneratedClassification interface for type OfferTermInfoAutoGenerated.
func (o *OfferTermInfoAutoGenerated) GetOfferTermInfoAutoGenerated() *OfferTermInfoAutoGenerated {
	return o
}

// RateCardClientGetOptions contains the optional parameters for the RateCardClient.Get method.
type RateCardClientGetOptions struct {
	// placeholder for future optional parameters
}

// RateCardQueryParameters - Parameters that are used in the odata $filter query parameter for providing RateCard information.
type RateCardQueryParameters struct {
	// REQUIRED; The currency in which the rates need to be provided.
	Currency *string `json:"Currency,omitempty"`

	// REQUIRED; The culture in which the resource metadata needs to be localized.
	Locale *string `json:"Locale,omitempty"`

	// REQUIRED; The Offer ID parameter consists of the 'MS-AZR-' prefix, plus the Offer ID number (e.g., MS-AZR-0026P). See https://azure.microsoft.com/en-us/support/legal/offer-details/
	// for more information on the
	// list of available Offer IDs, country/region availability, and billing currency.
	OfferDurableID *string `json:"OfferDurableId,omitempty"`

	// REQUIRED; 2 letter ISO code where the offer was purchased.
	RegionInfo *string `json:"RegionInfo,omitempty"`
}

// RecurringCharge - Indicates a recurring charge is present for this offer.
type RecurringCharge struct {
	// REQUIRED; Name of the offer term
	Name *OfferTermInfo `json:"Name,omitempty"`

	// Indicates the date from which the offer term is effective.
	EffectiveDate *time.Time `json:"EffectiveDate,omitempty"`

	// The amount of recurring charge as per the offer term.
	RecurringCharge *int32 `json:"RecurringCharge,omitempty"`
}

// GetOfferTermInfoAutoGenerated implements the OfferTermInfoAutoGeneratedClassification interface for type RecurringCharge.
func (r *RecurringCharge) GetOfferTermInfoAutoGenerated() *OfferTermInfoAutoGenerated {
	return &OfferTermInfoAutoGenerated{
		Name:          r.Name,
		EffectiveDate: r.EffectiveDate,
	}
}

// ResourceRateCardInfo - Price and Metadata information for resources
type ResourceRateCardInfo struct {
	// The currency in which the rates are provided.
	Currency *string `json:"Currency,omitempty"`

	// All rates are pretax, so this will always be returned as 'false'.
	IsTaxIncluded *bool `json:"IsTaxIncluded,omitempty"`

	// The culture in which the resource information is localized.
	Locale *string `json:"Locale,omitempty"`

	// A list of meters.
	Meters []*MeterInfo `json:"Meters,omitempty"`

	// A list of offer terms.
	OfferTerms []OfferTermInfoAutoGeneratedClassification `json:"OfferTerms,omitempty"`
}

// UsageAggregatesClientListOptions contains the optional parameters for the UsageAggregatesClient.List method.
type UsageAggregatesClientListOptions struct {
	// Daily (default) returns the data in daily granularity, Hourly returns the data in hourly granularity.
	AggregationGranularity *AggregationGranularity
	// Used when a continuation token string is provided in the response body of the previous call, enabling paging through a
	// large result set. If not present, the data is retrieved from the beginning of the
	// day/hour (based on the granularity) passed in.
	ContinuationToken *string
	// True returns usage data in instance-level detail, false causes server-side aggregation with fewer details. For example,
	// if you have 3 website instances, by default you will get 3 line items for
	// website consumption. If you specify showDetails = false, the data will be aggregated as a single line item for website
	// consumption within the time period (for the given subscriptionId, meterId,
	// usageStartTime and usageEndTime).
	ShowDetails *bool
}

// UsageAggregation - Describes the usageAggregation.
type UsageAggregation struct {
	// Unique Id for the usage aggregate.
	ID *string `json:"id,omitempty"`

	// Name of the usage aggregate.
	Name *string `json:"name,omitempty"`

	// Usage data.
	Properties *UsageSample `json:"properties,omitempty"`

	// Type of the resource being returned.
	Type *string `json:"type,omitempty"`
}

// UsageAggregationListResult - The Get UsageAggregates operation response.
type UsageAggregationListResult struct {
	// Gets or sets the link to the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// Gets or sets details for the requested aggregation.
	Value []*UsageAggregation `json:"value,omitempty"`
}

// UsageSample - Describes a sample of the usageAggregation.
type UsageSample struct {
	// Key-value pairs of instance details (legacy format).
	InfoFields *InfoField `json:"infoFields,omitempty"`

	// Key-value pairs of instance details represented as a string.
	InstanceData *string `json:"instanceData,omitempty"`

	// Category of the consumed resource.
	MeterCategory *string `json:"meterCategory,omitempty"`

	// Unique ID for the resource that was consumed (aka ResourceID).
	MeterID *string `json:"meterId,omitempty"`

	// Friendly name of the resource being consumed.
	MeterName *string `json:"meterName,omitempty"`

	// Region of the meterId used for billing purposes
	MeterRegion *string `json:"meterRegion,omitempty"`

	// Sub-category of the consumed resource.
	MeterSubCategory *string `json:"meterSubCategory,omitempty"`

	// The amount of the resource consumption that occurred in this time frame.
	Quantity *float32 `json:"quantity,omitempty"`

	// The subscription identifier for the Azure user.
	SubscriptionID *string `json:"subscriptionId,omitempty"`

	// The unit in which the usage for this resource is being counted, e.g. Hours, GB.
	Unit *string `json:"unit,omitempty"`

	// UTC end time for the usage bucket to which this usage aggregate belongs.
	UsageEndTime *time.Time `json:"usageEndTime,omitempty"`

	// UTC start time for the usage bucket to which this usage aggregate belongs.
	UsageStartTime *time.Time `json:"usageStartTime,omitempty"`
}
