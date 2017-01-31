package commerce

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/satori/uuid"
	"github.com/shopspring/decimal"
	"net/http"
)

// AggregationGranularity enumerates the values for aggregation granularity.
type AggregationGranularity string

const (
	// Daily specifies the daily state for aggregation granularity.
	Daily AggregationGranularity = "Daily"
	// Hourly specifies the hourly state for aggregation granularity.
	Hourly AggregationGranularity = "Hourly"
)

// ErrorResponse is describes the format of Error response.
type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// InfoField is key-value pairs of instance details in the legacy format.
type InfoField struct {
	Project *string `json:"project,omitempty"`
}

// MeterInfo is detailed information about the meter.
type MeterInfo struct {
	MeterID          *uuid.UUID           `json:"MeterId,omitempty"`
	MeterName        *string              `json:"MeterName,omitempty"`
	MeterCategory    *string              `json:"MeterCategory,omitempty"`
	MeterSubCategory *string              `json:"MeterSubCategory,omitempty"`
	Unit             *string              `json:"Unit,omitempty"`
	MeterRates       *map[string]*float64 `json:"MeterRates,omitempty"`
	EffectiveDate    *date.Time           `json:"EffectiveDate,omitempty"`
	IncludedQuantity *float64             `json:"IncludedQuantity,omitempty"`
}

// MonetaryCommitment is
type MonetaryCommitment struct {
	EffectiveDate    *date.Time                   `json:"EffectiveDate,omitempty"`
	ExcludedMeterIds *[]uuid.UUID                 `json:"ExcludedMeterIds,omitempty"`
	TieredDiscount   *map[string]*decimal.Decimal `json:"TieredDiscount,omitempty"`
}

// MonetaryCredit is
type MonetaryCredit struct {
	EffectiveDate    *date.Time       `json:"EffectiveDate,omitempty"`
	ExcludedMeterIds *[]uuid.UUID     `json:"ExcludedMeterIds,omitempty"`
	Credit           *decimal.Decimal `json:"Credit,omitempty"`
}

// OfferTermInfo is describes the offer term.
type OfferTermInfo struct {
	EffectiveDate    *date.Time   `json:"EffectiveDate,omitempty"`
	ExcludedMeterIds *[]uuid.UUID `json:"ExcludedMeterIds,omitempty"`
}

// RateCardQueryParameters is parameters that are used in the odata $filter
// query parameter for providing RateCard information.
type RateCardQueryParameters struct {
	OfferDurableID *string `json:"OfferDurableId,omitempty"`
	Currency       *string `json:"Currency,omitempty"`
	Locale         *string `json:"Locale,omitempty"`
	RegionInfo     *string `json:"RegionInfo,omitempty"`
}

// RecurringCharge is
type RecurringCharge struct {
	EffectiveDate    *date.Time   `json:"EffectiveDate,omitempty"`
	ExcludedMeterIds *[]uuid.UUID `json:"ExcludedMeterIds,omitempty"`
	RecurringCharge  *int32       `json:"RecurringCharge,omitempty"`
}

// ResourceRateCardInfo is price and Metadata information for resources
type ResourceRateCardInfo struct {
	autorest.Response `json:"-"`
	Currency          *string          `json:"Currency,omitempty"`
	Locale            *string          `json:"Locale,omitempty"`
	IsTaxIncluded     *bool            `json:"IsTaxIncluded,omitempty"`
	MeterRegion       *string          `json:"MeterRegion,omitempty"`
	Tags              *[]string        `json:"Tags,omitempty"`
	OfferTerms        *[]OfferTermInfo `json:"OfferTerms,omitempty"`
	Meters            *[]MeterInfo     `json:"Meters,omitempty"`
}

// UsageAggregation is describes the usageAggregation.
type UsageAggregation struct {
	ID           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Type         *string `json:"type,omitempty"`
	*UsageSample `json:"properties,omitempty"`
}

// UsageAggregationListResult is the Get UsageAggregates operation response.
type UsageAggregationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]UsageAggregation `json:"value,omitempty"`
	NextLink          *string             `json:"nextLink,omitempty"`
}

// UsageAggregationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client UsageAggregationListResult) UsageAggregationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// UsageSample is describes a sample of the usageAggregation.
type UsageSample struct {
	SubscriptionID   *uuid.UUID              `json:"subscriptionId,omitempty"`
	MeterID          *string                 `json:"meterId,omitempty"`
	UsageStartTime   *date.Time              `json:"usageStartTime,omitempty"`
	UsageEndTime     *date.Time              `json:"usageEndTime,omitempty"`
	Quantity         *map[string]interface{} `json:"quantity,omitempty"`
	Unit             *string                 `json:"unit,omitempty"`
	MeterName        *string                 `json:"meterName,omitempty"`
	MeterCategory    *string                 `json:"meterCategory,omitempty"`
	MeterSubCategory *string                 `json:"meterSubCategory,omitempty"`
	MeterRegion      *string                 `json:"meterRegion,omitempty"`
	InfoFields       *InfoField              `json:"infoFields,omitempty"`
	InstanceData     *string                 `json:"instanceData,omitempty"`
}
