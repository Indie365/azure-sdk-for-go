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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/065033d1c4087a2b009e71c0b3f0666718354ebd/specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("rgtest01", "cachetest01", nil)
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
		// page.PrivateEndpointConnectionListResult = armredis.PrivateEndpointConnectionListResult{
		// 	Value: []*armredis.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("pectest01"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Cache/Redis/cachetest01/privateEndpointConnections/pectest01"),
		// 			Properties: &armredis.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armredis.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armredis.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pectest01"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Cache/Redis/cachetest01/privateEndpointConnections/pectest01"),
		// 			Properties: &armredis.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armredis.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armredis.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/065033d1c4087a2b009e71c0b3f0666718354ebd/specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheGetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "rgtest01", "cachetest01", "pectest01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armredis.PrivateEndpointConnection{
	// 	Name: to.Ptr("pectest01"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Cache/Redis/cachetest01/privateEndpointConnections/pectest01"),
	// 	Properties: &armredis.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armredis.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armredis.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/065033d1c4087a2b009e71c0b3f0666718354ebd/specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCachePutPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_BeginPut() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginPut(ctx, "rgtest01", "cachetest01", "pectest01", armredis.PrivateEndpointConnection{
		Properties: &armredis.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armredis.PrivateEndpointConnection{
	// 	Name: to.Ptr("pectest01"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Cache/Redis/cachetest01/privateEndpointConnections/pectest01"),
	// 	Properties: &armredis.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armredis.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/rgtest01/providers/Microsoft.Network/privateEndpoints/petest01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armredis.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/065033d1c4087a2b009e71c0b3f0666718354ebd/specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheDeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionsClient().Delete(ctx, "rgtest01", "cachetest01", "pectest01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
