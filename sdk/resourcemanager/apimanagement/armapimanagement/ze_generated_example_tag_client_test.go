//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiOperationTags.json
func ExampleTagClient_ListByOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	pager := client.ListByOperation("<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		&armapimanagement.TagClientListByOperationOptions{Filter: nil,
			Top:  nil,
			Skip: nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiOperationTag.json
func ExampleTagClient_GetEntityStateByOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	_, err = client.GetEntityStateByOperation(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiOperationTag.json
func ExampleTagClient_GetByOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	res, err := client.GetByOperation(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagClientGetByOperationResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiOperationTag.json
func ExampleTagClient_AssignToOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	res, err := client.AssignToOperation(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagClientAssignToOperationResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiOperationTag.json
func ExampleTagClient_DetachFromOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	_, err = client.DetachFromOperation(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductTags.json
func ExampleTagClient_ListByProduct() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	pager := client.ListByProduct("<resource-group-name>",
		"<service-name>",
		"<product-id>",
		&armapimanagement.TagClientListByProductOptions{Filter: nil,
			Top:  nil,
			Skip: nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadProductTag.json
func ExampleTagClient_GetEntityStateByProduct() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	_, err = client.GetEntityStateByProduct(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<product-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetProductTag.json
func ExampleTagClient_GetByProduct() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	res, err := client.GetByProduct(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<product-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagClientGetByProductResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateProductTag.json
func ExampleTagClient_AssignToProduct() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	res, err := client.AssignToProduct(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<product-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagClientAssignToProductResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteProductTag.json
func ExampleTagClient_DetachFromProduct() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	_, err = client.DetachFromProduct(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<product-id>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListTags.json
func ExampleTagClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	pager := client.ListByService("<resource-group-name>",
		"<service-name>",
		&armapimanagement.TagClientListByServiceOptions{Filter: nil,
			Top:   nil,
			Skip:  nil,
			Scope: nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadTag.json
func ExampleTagClient_GetEntityState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	_, err = client.GetEntityState(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetTag.json
func ExampleTagClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<tag-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagClientGetResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateTag.json
func ExampleTagClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<tag-id>",
		armapimanagement.TagCreateUpdateParameters{
			Properties: &armapimanagement.TagContractProperties{
				DisplayName: to.StringPtr("<display-name>"),
			},
		},
		&armapimanagement.TagClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateTag.json
func ExampleTagClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<tag-id>",
		"<if-match>",
		armapimanagement.TagCreateUpdateParameters{
			Properties: &armapimanagement.TagContractProperties{
				DisplayName: to.StringPtr("<display-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TagClientUpdateResult)
}

// x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteTag.json
func ExampleTagClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapimanagement.NewTagClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<tag-id>",
		"<if-match>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
