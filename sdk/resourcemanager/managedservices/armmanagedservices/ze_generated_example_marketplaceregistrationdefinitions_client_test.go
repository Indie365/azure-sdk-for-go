//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetMarketplaceRegistrationDefinitions.json
func ExampleMarketplaceRegistrationDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagedservices.NewMarketplaceRegistrationDefinitionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("subscription/0afefe50-734e-4610-8a82-a144ahf49dea",
		&armmanagedservices.MarketplaceRegistrationDefinitionsClientListOptions{Filter: to.Ptr("planIdentifier eq 'publisher.offerIdentifier.planName.version'")})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetMarketplaceRegistrationDefinition.json
func ExampleMarketplaceRegistrationDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagedservices.NewMarketplaceRegistrationDefinitionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"subscription/0afefe50-734e-4610-8a82-a144ahf49dea",
		"publisher.product.planName.version",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}