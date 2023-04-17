//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetProviderOperationsRP.json
func ExampleProviderOperationsMetadataClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderOperationsMetadataClient().Get(ctx, "resourceProviderNamespace", &armauthorization.ProviderOperationsMetadataClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProviderOperationsMetadata = armauthorization.ProviderOperationsMetadata{
	// 	Name: to.Ptr("name"),
	// 	Type: to.Ptr("type"),
	// 	DisplayName: to.Ptr("displayName"),
	// 	ID: to.Ptr("id"),
	// 	ResourceTypes: []*armauthorization.ResourceType{
	// 		{
	// 			Name: to.Ptr("name"),
	// 			DisplayName: to.Ptr("name"),
	// 			Operations: []*armauthorization.ProviderOperation{
	// 			},
	// 	}},
	// 	Operations: []*armauthorization.ProviderOperation{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/53b1affe357b3bfbb53721d0a2002382a046d3b0/specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetAllProviderOperations.json
func ExampleProviderOperationsMetadataClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProviderOperationsMetadataClient().NewListPager(&armauthorization.ProviderOperationsMetadataClientListOptions{Expand: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProviderOperationsMetadataListResult = armauthorization.ProviderOperationsMetadataListResult{
		// 	Value: []*armauthorization.ProviderOperationsMetadata{
		// 		{
		// 			Name: to.Ptr("name"),
		// 			Type: to.Ptr("type"),
		// 			DisplayName: to.Ptr("displayName"),
		// 			ID: to.Ptr("id"),
		// 			ResourceTypes: []*armauthorization.ResourceType{
		// 				{
		// 					Name: to.Ptr("name"),
		// 					DisplayName: to.Ptr("name"),
		// 					Operations: []*armauthorization.ProviderOperation{
		// 					},
		// 			}},
		// 			Operations: []*armauthorization.ProviderOperation{
		// 			},
		// 	}},
		// }
	}
}
