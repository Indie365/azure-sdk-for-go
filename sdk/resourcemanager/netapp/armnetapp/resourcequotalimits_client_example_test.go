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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/QuotaLimits_List.json
func ExampleResourceQuotaLimitsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewResourceQuotaLimitsClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("eastus", nil)
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
		// page.SubscriptionQuotaItemList = armnetapp.SubscriptionQuotaItemList{
		// 	Value: []*armnetapp.SubscriptionQuotaItem{
		// 		{
		// 			Name: to.Ptr("eastus/accountsPerSubscription"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/accountsPerSubscription"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](10),
		// 				Current: to.Ptr[int32](10),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eastus/poolsPerAccount"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/poolsPerAccount"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](25),
		// 				Current: to.Ptr[int32](25),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eastus/volumesPerPool"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/volumesPerPool"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](500),
		// 				Current: to.Ptr[int32](500),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eastus/snapshotsPerVolume"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/snapshotsPerVolume"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](255),
		// 				Current: to.Ptr[int32](255),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eastus/totalTiBsPerSubscription"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/totalTiBsPerSubscription"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](25),
		// 				Current: to.Ptr[int32](1000),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eastus/totalDPVolumesPerSubscription"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/totalDPVolumesPerSubscription"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](10),
		// 				Current: to.Ptr[int32](10),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eastus/totalVolumesPerSubscription"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/totalVolumesPerSubscription"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](500),
		// 				Current: to.Ptr[int32](500),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("eastus/totalCoolAccessVolumesPerSubscription"),
		// 			Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/totalCoolAccessVolumesPerSubscription"),
		// 			Properties: &armnetapp.SubscriptionQuotaItemProperties{
		// 				Default: to.Ptr[int32](10),
		// 				Current: to.Ptr[int32](10),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/QuotaLimits_Get.json
func ExampleResourceQuotaLimitsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewResourceQuotaLimitsClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "eastus", "totalCoolAccessVolumesPerSubscription", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionQuotaItem = armnetapp.SubscriptionQuotaItem{
	// 	Name: to.Ptr("eastus/totalCoolAccessVolumesPerSubscription"),
	// 	Type: to.Ptr("Microsoft.NetApp/locations/quotaLimits"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/providers/Microsoft.NetApp/locations/eastus/quotaLimits/totalCoolAccessVolumesPerSubscription"),
	// 	Properties: &armnetapp.SubscriptionQuotaItemProperties{
	// 		Default: to.Ptr[int32](10),
	// 		Current: to.Ptr[int32](10),
	// 	},
	// }
}
