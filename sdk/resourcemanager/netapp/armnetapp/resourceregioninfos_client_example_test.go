//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/RegionInfos_List.json
func ExampleResourceRegionInfosClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewResourceRegionInfosClient().NewListPager("eastus", nil)
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
		// page.RegionInfosList = armnetapp.RegionInfosList{
		// 	Value: []*armnetapp.RegionInfoResource{
		// 		{
		// 			Name: to.Ptr("eastus/default"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/regionInfos"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/regionInfos/default"),
		// 			Properties: &armnetapp.RegionInfo{
		// 				AvailabilityZoneMappings: []*armnetapp.RegionInfoAvailabilityZoneMappingsItem{
		// 					{
		// 						AvailabilityZone: to.Ptr("1"),
		// 						IsAvailable: to.Ptr(true),
		// 					},
		// 					{
		// 						AvailabilityZone: to.Ptr("2"),
		// 						IsAvailable: to.Ptr(true),
		// 					},
		// 					{
		// 						AvailabilityZone: to.Ptr("3"),
		// 						IsAvailable: to.Ptr(true),
		// 				}},
		// 				StorageToNetworkProximity: to.Ptr(armnetapp.RegionStorageToNetworkProximityT2),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/RegionInfos_Get.json
func ExampleResourceRegionInfosClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceRegionInfosClient().Get(ctx, "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegionInfoResource = armnetapp.RegionInfoResource{
	// 	Name: to.Ptr("eastus/default"),
	// 	Type: to.Ptr("Microsoft.NetApp/locations/regionInfos"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/regionInfos/default"),
	// 	Properties: &armnetapp.RegionInfo{
	// 		AvailabilityZoneMappings: []*armnetapp.RegionInfoAvailabilityZoneMappingsItem{
	// 			{
	// 				AvailabilityZone: to.Ptr("1"),
	// 				IsAvailable: to.Ptr(true),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("2"),
	// 				IsAvailable: to.Ptr(true),
	// 			},
	// 			{
	// 				AvailabilityZone: to.Ptr("3"),
	// 				IsAvailable: to.Ptr(true),
	// 		}},
	// 		StorageToNetworkProximity: to.Ptr(armnetapp.RegionStorageToNetworkProximityT2),
	// 	},
	// }
}
