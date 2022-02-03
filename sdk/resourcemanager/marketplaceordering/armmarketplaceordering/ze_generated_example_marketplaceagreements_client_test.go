//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmarketplaceordering_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering"
)

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/GetMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		armmarketplaceordering.OfferType("virtualmachine"),
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientGetResult)
}

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SetMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		armmarketplaceordering.OfferType("virtualmachine"),
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		armmarketplaceordering.AgreementTerms{
			Properties: &armmarketplaceordering.AgreementProperties{
				Accepted:             to.BoolPtr(false),
				LicenseTextLink:      to.StringPtr("<license-text-link>"),
				MarketplaceTermsLink: to.StringPtr("<marketplace-terms-link>"),
				Plan:                 to.StringPtr("<plan>"),
				PrivacyPolicyLink:    to.StringPtr("<privacy-policy-link>"),
				Product:              to.StringPtr("<product>"),
				Publisher:            to.StringPtr("<publisher>"),
				RetrieveDatetime:     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T11:33:07.12132Z"); return t }()),
				Signature:            to.StringPtr("<signature>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientCreateResult)
}

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SignMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Sign() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.Sign(ctx,
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientSignResult)
}

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/CancelMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.Cancel(ctx,
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientCancelResult)
}

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/GetAgreementMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_GetAgreement() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.GetAgreement(ctx,
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientGetAgreementResult)
}

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/ListMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientListResult)
}
