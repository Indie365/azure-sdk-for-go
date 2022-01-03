//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbooksList2.json
func ExampleWorkbooksClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(armapplicationinsights.CategoryTypeWorkbook,
		&armapplicationinsights.WorkbooksListBySubscriptionOptions{Tags: []string{},
			CanFetchContent: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Workbook.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbooksList.json
func ExampleWorkbooksClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		armapplicationinsights.CategoryTypeWorkbook,
		&armapplicationinsights.WorkbooksListByResourceGroupOptions{Tags: []string{},
			SourceID:        to.StringPtr("<source-id>"),
			CanFetchContent: nil,
		})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Workbook.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookGet.json
func ExampleWorkbooksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Workbook.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookDelete.json
func ExampleWorkbooksClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookAdd.json
func ExampleWorkbooksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.Workbook{
			WorkbookResource: armapplicationinsights.WorkbookResource{
				TrackedResource: armapplicationinsights.TrackedResource{
					Location: to.StringPtr("<location>"),
					Tags: map[string]*string{
						"TagSample01": to.StringPtr("sample01"),
						"TagSample02": to.StringPtr("sample02"),
					},
				},
				Kind: armapplicationinsights.KindShared.ToPtr(),
			},
			Properties: &armapplicationinsights.WorkbookProperties{
				Description:    to.StringPtr("<description>"),
				Category:       to.StringPtr("<category>"),
				DisplayName:    to.StringPtr("<display-name>"),
				SerializedData: to.StringPtr("<serialized-data>"),
			},
		},
		&armapplicationinsights.WorkbooksCreateOrUpdateOptions{SourceID: to.StringPtr("<source-id>")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Workbook.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookManagedUpdate.json
func ExampleWorkbooksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<resource-name>",
		&armapplicationinsights.WorkbooksUpdateOptions{SourceID: to.StringPtr("<source-id>"),
			WorkbookUpdateParameters: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Workbook.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookRevisionsList.json
func ExampleWorkbooksClient_RevisionsList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	pager := client.RevisionsList("<resource-group-name>",
		"<resource-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Workbook.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookRevisionGet.json
func ExampleWorkbooksClient_RevisionGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewWorkbooksClient("<subscription-id>", cred, nil)
	res, err := client.RevisionGet(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<revision-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Workbook.ID: %s\n", *res.ID)
}
