//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfilesListByBillingAccount.json
func ExampleBillingProfilesClient_ListByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingProfilesClient(cred, nil)
	pager := client.ListByBillingAccount("<billing-account-name>",
		&armbilling.BillingProfilesListByBillingAccountOptions{Expand: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("BillingProfile.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfile.json
func ExampleBillingProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingProfilesClient(cred, nil)
	res, err := client.Get(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		&armbilling.BillingProfilesGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingProfile.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
func ExampleBillingProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingProfilesClient(cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		armbilling.BillingProfile{
			Properties: &armbilling.BillingProfileProperties{
				BillTo: &armbilling.AddressDetails{
					AddressLine1: to.StringPtr("<address-line1>"),
					City:         to.StringPtr("<city>"),
					Country:      to.StringPtr("<country>"),
					FirstName:    to.StringPtr("<first-name>"),
					LastName:     to.StringPtr("<last-name>"),
					PostalCode:   to.StringPtr("<postal-code>"),
					Region:       to.StringPtr("<region>"),
				},
				DisplayName: to.StringPtr("<display-name>"),
				EnabledAzurePlans: []*armbilling.AzurePlan{
					{
						SKUID: to.StringPtr("<skuid>"),
					},
					{
						SKUID: to.StringPtr("<skuid>"),
					}},
				InvoiceEmailOptIn: to.BoolPtr(true),
				PoNumber:          to.StringPtr("<po-number>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingProfile.ID: %s\n", *res.ID)
}
