//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

func ExampleRestorePointCollectionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewRestorePointCollectionsClient(con,
		"<subscription-id>")
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		armcompute.RestorePointCollection{
			Resource: armcompute.Resource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"myTag1": to.StringPtr("tagValue1"),
				},
			},
			Properties: &armcompute.RestorePointCollectionProperties{
				Source: &armcompute.RestorePointCollectionSourceProperties{
					ID: to.StringPtr("<id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RestorePointCollection.ID: %s\n", *res.ID)
}

func ExampleRestorePointCollectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewRestorePointCollectionsClient(con,
		"<subscription-id>")
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("RestorePointCollection.ID: %s\n", *res.ID)
}

func ExampleRestorePointCollectionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewRestorePointCollectionsClient(con,
		"<subscription-id>")
	pager := client.List("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RestorePointCollection.ID: %s\n", *v.ID)
		}
	}
}

func ExampleRestorePointCollectionsClient_ListAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewRestorePointCollectionsClient(con,
		"<subscription-id>")
	pager := client.ListAll(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("RestorePointCollection.ID: %s\n", *v.ID)
		}
	}
}
