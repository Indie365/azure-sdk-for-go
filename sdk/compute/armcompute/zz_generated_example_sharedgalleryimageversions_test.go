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
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

func ExampleSharedGalleryImageVersionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewSharedGalleryImageVersionsClient(con,
		"<subscription-id>")
	pager := client.List("<location>",
		"<gallery-unique-name>",
		"<gallery-image-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}

func ExampleSharedGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewSharedGalleryImageVersionsClient(con,
		"<subscription-id>")
	_, err = client.Get(ctx,
		"<location>",
		"<gallery-unique-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
