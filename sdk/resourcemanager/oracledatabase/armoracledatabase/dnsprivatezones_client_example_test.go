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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateZones_listByLocation.json
func ExampleDNSPrivateZonesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDNSPrivateZonesClient().NewListByLocationPager("eastus", nil)
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
		// page.DNSPrivateZoneListResult = armoracledatabase.DNSPrivateZoneListResult{
		// 	Value: []*armoracledatabase.DNSPrivateZone{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/locations/dnsPrivateZones"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dnsPrivateZones/zone1"),
		// 			Properties: &armoracledatabase.DNSPrivateZoneProperties{
		// 				IsProtected: to.Ptr(false),
		// 				LifecycleState: to.Ptr(armoracledatabase.DNSPrivateZonesLifecycleStateActive),
		// 				Ocid: to.Ptr("zone-id-1"),
		// 				Self: to.Ptr("https://api.example.com/zone1"),
		// 				Serial: to.Ptr[int32](12345),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-26T12:34:56.000Z"); return t}()),
		// 				Version: to.Ptr("1.0.0.0"),
		// 				ViewID: to.Ptr("view-id-1"),
		// 				ZoneType: to.Ptr(armoracledatabase.ZoneTypePrimary),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Oracle.Database/locations/dnsPrivateZones"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dnsPrivateZones/zone2"),
		// 			Properties: &armoracledatabase.DNSPrivateZoneProperties{
		// 				IsProtected: to.Ptr(true),
		// 				LifecycleState: to.Ptr(armoracledatabase.DNSPrivateZonesLifecycleStateCreating),
		// 				Ocid: to.Ptr("zone-id-2"),
		// 				Self: to.Ptr("https://api.example.com/zone2"),
		// 				Serial: to.Ptr[int32](54321),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-27T09:45:00.000Z"); return t}()),
		// 				Version: to.Ptr("2.0.0.0"),
		// 				ViewID: to.Ptr("view-id-2"),
		// 				ZoneType: to.Ptr(armoracledatabase.ZoneTypeSecondary),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dnsPrivateZones_get.json
func ExampleDNSPrivateZonesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDNSPrivateZonesClient().Get(ctx, "eastus", "example-dns-private-zone", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DNSPrivateZone = armoracledatabase.DNSPrivateZone{
	// 	Type: to.Ptr("Oracle.Database/locations/dnsPrivateZones"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Oracle.Database/locations/eastus/dnsPrivateZones/example-dns-private-zone"),
	// 	Properties: &armoracledatabase.DNSPrivateZoneProperties{
	// 		IsProtected: to.Ptr(false),
	// 		LifecycleState: to.Ptr(armoracledatabase.DNSPrivateZonesLifecycleStateActive),
	// 		Ocid: to.Ptr("your-zone-id"),
	// 		Self: to.Ptr("https://api.example.com/your-dns-private-zone"),
	// 		Serial: to.Ptr[int32](12345),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-26T12:34:56.000Z"); return t}()),
	// 		Version: to.Ptr("1.0.0.0"),
	// 		ViewID: to.Ptr("your-view-id"),
	// 		ZoneType: to.Ptr(armoracledatabase.ZoneTypePrimary),
	// 	},
	// }
}
