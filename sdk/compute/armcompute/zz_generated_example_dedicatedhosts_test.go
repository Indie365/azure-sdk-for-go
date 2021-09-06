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

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

func ExampleDedicatedHostsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewDedicatedHostsClient(con,
		"<subscription-id>")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<host-group-name>",
		"<host-name>",
		armcompute.DedicatedHost{
			Resource: armcompute.Resource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"department": to.StringPtr("HR"),
				},
			},
			Properties: &armcompute.DedicatedHostProperties{
				PlatformFaultDomain: to.Int32Ptr(1),
			},
			SKU: &armcompute.SKU{
				Name: to.StringPtr("<name>"),
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
	log.Printf("DedicatedHost.ID: %s\n", *res.ID)
}

func ExampleDedicatedHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewDedicatedHostsClient(con,
		"<subscription-id>")
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<host-group-name>",
		"<host-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DedicatedHost.ID: %s\n", *res.ID)
}
