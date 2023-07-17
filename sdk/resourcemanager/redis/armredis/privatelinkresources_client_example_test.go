//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheListPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_NewListByRedisCachePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByRedisCachePager("rgtest01", "cacheTest01", nil)
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
		// page.PrivateLinkResourceListResult = armredis.PrivateLinkResourceListResult{
		// 	Value: []*armredis.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("redisCache"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/rgtest01/providers/Microsoft.Cache/Redis/cacheTest01/privateLinkResources/redisCache"),
		// 			Properties: &armredis.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("redisCache"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("redisCache")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.redis.cache.windows.net")},
		// 					},
		// 			}},
		// 		}
	}
}
