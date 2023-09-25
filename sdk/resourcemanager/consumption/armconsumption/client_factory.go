//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential azcore.TokenCredential
	options *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: 	subscriptionID,		credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAggregatedCostClient() *AggregatedCostClient {
	subClient, _ := NewAggregatedCostClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBalancesClient() *BalancesClient {
	subClient, _ := NewBalancesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBudgetsClient() *BudgetsClient {
	subClient, _ := NewBudgetsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewChargesClient() *ChargesClient {
	subClient, _ := NewChargesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCreditsClient() *CreditsClient {
	subClient, _ := NewCreditsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEventsClient() *EventsClient {
	subClient, _ := NewEventsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLotsClient() *LotsClient {
	subClient, _ := NewLotsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMarketplacesClient() *MarketplacesClient {
	subClient, _ := NewMarketplacesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPriceSheetClient() *PriceSheetClient {
	subClient, _ := NewPriceSheetClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReservationRecommendationDetailsClient() *ReservationRecommendationDetailsClient {
	subClient, _ := NewReservationRecommendationDetailsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReservationRecommendationsClient() *ReservationRecommendationsClient {
	subClient, _ := NewReservationRecommendationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReservationTransactionsClient() *ReservationTransactionsClient {
	subClient, _ := NewReservationTransactionsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReservationsDetailsClient() *ReservationsDetailsClient {
	subClient, _ := NewReservationsDetailsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReservationsSummariesClient() *ReservationsSummariesClient {
	subClient, _ := NewReservationsSummariesClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTagsClient() *TagsClient {
	subClient, _ := NewTagsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUsageDetailsClient() *UsageDetailsClient {
	subClient, _ := NewUsageDetailsClient(c.credential, c.options)
	return subClient
}

