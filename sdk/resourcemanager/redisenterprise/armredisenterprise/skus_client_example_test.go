//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9565a97e3efbeb3691c9100d5473b0a884c1b786/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseSkusList.json
func ExampleSKUsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSKUsClient().NewListPager("westus2", nil)
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
		// page.RegionSKUDetails = armredisenterprise.RegionSKUDetails{
		// 	Value: []*armredisenterprise.RegionSKUDetail{
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUName("Enterprise_E1")),
		// 			},
		// 		},
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseE10),
		// 			},
		// 		},
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseE20),
		// 			},
		// 		},
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseE50),
		// 			},
		// 		},
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseE100),
		// 			},
		// 		},
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
		// 			},
		// 		},
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF700),
		// 			},
		// 		},
		// 		{
		// 			LocationInfo: &armredisenterprise.LocationInfo{
		// 				Capabilities: []*armredisenterprise.Capability{
		// 					{
		// 						Name: to.Ptr("ZoneAware"),
		// 						Value: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("FDsAvailable"),
		// 						Value: to.Ptr(true),
		// 				}},
		// 				Location: to.Ptr("westus2"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Cache/redisEnterprise"),
		// 			SKUDetails: &armredisenterprise.SKUDetail{
		// 				Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF1500),
		// 			},
		// 	}},
		// }
	}
}
