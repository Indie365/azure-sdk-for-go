//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/69ece3818b8b0929b43a07c3fe25716427734882/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PrivateEndPointConnections_ListByFactory.json
func ExamplePrivateEndPointConnectionsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndPointConnectionsClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.PrivateEndpointConnectionListResponse = armdatafactory.PrivateEndpointConnectionListResponse{
		// 	Value: []*armdatafactory.PrivateEndpointConnectionResource{
		// 		{
		// 			Name: to.Ptr("factories"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/privateEndpoints/myPrivateEndpoint"),
		// 			Properties: &armdatafactory.RemotePrivateEndpointConnection{
		// 				PrivateEndpoint: &armdatafactory.ArmIDWrapper{
		// 					ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/managedVirtualNetworks/myPrivateEndpoint"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armdatafactory.PrivateLinkConnectionState{
		// 					Description: to.Ptr("Approved by admin."),
		// 					ActionsRequired: to.Ptr("exampleActionsRequired"),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
