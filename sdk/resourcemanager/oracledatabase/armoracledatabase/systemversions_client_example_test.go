//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/systemVersions_listByLocation.json
func ExampleSystemVersionsClient_NewListByLocationPager_listExadataSystemVersionsByTheProvidedFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSystemVersionsClient().NewListByLocationPager("eastus", nil)
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
		// page.SystemVersionListResult = armoracledatabase.SystemVersionListResult{
		// 	Value: []*armoracledatabase.SystemVersion{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/systemVersion/22.1.7.0.0.230113"),
		// 			Properties: &armoracledatabase.SystemVersionProperties{
		// 				SystemVersion: to.Ptr("22.1.7.0.0.230113"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/systemVersions_get.json
func ExampleSystemVersionsClient_Get_getExadataSystemVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSystemVersionsClient().Get(ctx, "eastus", "22.x", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SystemVersion = armoracledatabase.SystemVersion{
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/systemVersion/22.1.7.0.0.230113"),
	// 	Properties: &armoracledatabase.SystemVersionProperties{
	// 		SystemVersion: to.Ptr("22.1.7.0.0.230113"),
	// 	},
	// }
}
