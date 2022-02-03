//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications"
)

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getApplicationDefinition.json
func ExampleApplicationDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<application-definition-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationDefinitionsClientGetResult)
}

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/deleteApplicationDefinition.json
func ExampleApplicationDefinitionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationDefinitionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<application-definition-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplicationDefinition.json
func ExampleApplicationDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationDefinitionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<application-definition-name>",
		armmanagedapplications.ApplicationDefinition{
			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
				Description: to.StringPtr("<description>"),
				Authorizations: []*armmanagedapplications.ApplicationAuthorization{
					{
						PrincipalID:      to.StringPtr("<principal-id>"),
						RoleDefinitionID: to.StringPtr("<role-definition-id>"),
					}},
				DisplayName:    to.StringPtr("<display-name>"),
				LockLevel:      armmanagedapplications.ApplicationLockLevelNone.ToPtr(),
				PackageFileURI: to.StringPtr("<package-file-uri>"),
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
	log.Printf("Response result: %#v\n", res.ApplicationDefinitionsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateApplicationDefinition.json
func ExampleApplicationDefinitionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationDefinitionsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<application-definition-name>",
		armmanagedapplications.ApplicationDefinitionPatchable{
			Tags: map[string]*string{
				"department": to.StringPtr("Finance"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ApplicationDefinitionsClientUpdateResult)
}

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listApplicationDefinitionsByResourceGroup.json
func ExampleApplicationDefinitionsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationDefinitionsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
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

// x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listApplicationDefinitionsByResourceGroup.json
func ExampleApplicationDefinitionsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedapplications.NewApplicationDefinitionsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(nil)
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
