package advisor

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.21.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// Category enumerates the values for category.
type Category string

const (
	// Cost specifies the cost state for category.
	Cost Category = "Cost"
	// HighAvailability specifies the high availability state for category.
	HighAvailability Category = "HighAvailability"
	// Performance specifies the performance state for category.
	Performance Category = "Performance"
	// Security specifies the security state for category.
	Security Category = "Security"
)

// Impact enumerates the values for impact.
type Impact string

const (
	// High specifies the high state for impact.
	High Impact = "High"
	// Low specifies the low state for impact.
	Low Impact = "Low"
	// Medium specifies the medium state for impact.
	Medium Impact = "Medium"
)

// Risk enumerates the values for risk.
type Risk string

const (
	// Error specifies the error state for risk.
	Error Risk = "Error"
	// None specifies the none state for risk.
	None Risk = "None"
	// Warning specifies the warning state for risk.
	Warning Risk = "Warning"
)

// OperationDisplayInfo is the operation supported by Advisor.
type OperationDisplayInfo struct {
	Description *string `json:"description,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
}

// OperationEntity is the operation supported by Advisor.
type OperationEntity struct {
	Name    *string               `json:"name,omitempty"`
	Display *OperationDisplayInfo `json:"display,omitempty"`
}

// OperationEntityListResult is the list of Advisor operations.
type OperationEntityListResult struct {
	autorest.Response `json:"-"`
	NextLink          *string            `json:"nextLink,omitempty"`
	Value             *[]OperationEntity `json:"value,omitempty"`
}

// OperationEntityListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationEntityListResult) OperationEntityListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// RecommendationProperties is the properties of the recommendation.
type RecommendationProperties struct {
	Category             Category                            `json:"category,omitempty"`
	Impact               Impact                              `json:"impact,omitempty"`
	ImpactedField        *string                             `json:"impactedField,omitempty"`
	ImpactedValue        *string                             `json:"impactedValue,omitempty"`
	LastUpdated          *date.Time                          `json:"lastUpdated,omitempty"`
	Metadata             *map[string]*map[string]interface{} `json:"metadata,omitempty"`
	RecommendationTypeID *string                             `json:"recommendationTypeId,omitempty"`
	Risk                 Risk                                `json:"risk,omitempty"`
	ShortDescription     *ShortDescription                   `json:"shortDescription,omitempty"`
	SuppressionIds       *[]uuid.UUID                        `json:"suppressionIds,omitempty"`
}

// Resource is an Azure resource.
type Resource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// ResourceRecommendationBase is advisor Recommendation.
type ResourceRecommendationBase struct {
	autorest.Response         `json:"-"`
	ID                        *string `json:"id,omitempty"`
	Name                      *string `json:"name,omitempty"`
	Type                      *string `json:"type,omitempty"`
	*RecommendationProperties `json:"properties,omitempty"`
}

// ResourceRecommendationBaseListResult is the list of Advisor recommendations.
type ResourceRecommendationBaseListResult struct {
	autorest.Response `json:"-"`
	NextLink          *string                       `json:"nextLink,omitempty"`
	Value             *[]ResourceRecommendationBase `json:"value,omitempty"`
}

// ResourceRecommendationBaseListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResourceRecommendationBaseListResult) ResourceRecommendationBaseListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ShortDescription is a summary of the recommendation.
type ShortDescription struct {
	Problem  *string `json:"problem,omitempty"`
	Solution *string `json:"solution,omitempty"`
}

// SuppressionContract is the details of the snoozed or dismissed rule; for example, the duration, name, and GUID
// associated with the rule.
type SuppressionContract struct {
	autorest.Response      `json:"-"`
	ID                     *string `json:"id,omitempty"`
	Name                   *string `json:"name,omitempty"`
	Type                   *string `json:"type,omitempty"`
	*SuppressionProperties `json:"properties,omitempty"`
}

// SuppressionContractListResult is the list of Advisor suppressions.
type SuppressionContractListResult struct {
	autorest.Response `json:"-"`
	NextLink          *string                `json:"nextLink,omitempty"`
	Value             *[]SuppressionContract `json:"value,omitempty"`
}

// SuppressionContractListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client SuppressionContractListResult) SuppressionContractListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SuppressionProperties is the properties of the suppression.
type SuppressionProperties struct {
	SuppressionID *string `json:"suppressionId,omitempty"`
	TTL           *string `json:"ttl,omitempty"`
}
