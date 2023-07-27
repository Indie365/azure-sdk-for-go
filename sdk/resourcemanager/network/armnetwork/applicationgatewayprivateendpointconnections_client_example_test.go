//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ApplicationGatewayPrivateEndpointConnectionDelete.json
func ExampleApplicationGatewayPrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationGatewayPrivateEndpointConnectionsClient().BeginDelete(ctx, "rg1", "appgw", "connection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ApplicationGatewayPrivateEndpointConnectionUpdate.json
func ExampleApplicationGatewayPrivateEndpointConnectionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationGatewayPrivateEndpointConnectionsClient().BeginUpdate(ctx, "rg1", "appgw", "connection1", armnetwork.ApplicationGatewayPrivateEndpointConnection{
		Name: to.Ptr("connection1"),
		Properties: &armnetwork.ApplicationGatewayPrivateEndpointConnectionProperties{
			PrivateEndpoint: &armnetwork.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/subId2/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
			},
			PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
				Description: to.Ptr("approved it for some reason."),
				Status:      to.Ptr("Approved"),
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
	// res.ApplicationGatewayPrivateEndpointConnection = armnetwork.ApplicationGatewayPrivateEndpointConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkResources/testPlePeConnection"),
	// 	Name: to.Ptr("testPlePeConnection"),
	// 	Properties: &armnetwork.ApplicationGatewayPrivateEndpointConnectionProperties{
	// 		LinkIdentifier: to.Ptr("linkId"),
	// 		PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("approved it for some reason."),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ApplicationGatewayPrivateEndpointConnectionGet.json
func ExampleApplicationGatewayPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationGatewayPrivateEndpointConnectionsClient().Get(ctx, "rg1", "appgw", "connection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationGatewayPrivateEndpointConnection = armnetwork.ApplicationGatewayPrivateEndpointConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkResources/connection1"),
	// 	Name: to.Ptr("coonection1"),
	// 	Type: to.Ptr("Microsoft.Network/applicationGateways/privateEndpointConnections"),
	// 	Properties: &armnetwork.ApplicationGatewayPrivateEndpointConnectionProperties{
	// 		LinkIdentifier: to.Ptr("805319460"),
	// 		PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe1"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approval Done"),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9975d3476c05bcc6bd9535ad3dfb564e6a168fa5/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/ApplicationGatewayPrivateEndpointConnectionList.json
func ExampleApplicationGatewayPrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationGatewayPrivateEndpointConnectionsClient().NewListPager("rg1", "appgw", nil)
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
		// page.ApplicationGatewayPrivateEndpointConnectionListResult = armnetwork.ApplicationGatewayPrivateEndpointConnectionListResult{
		// 	Value: []*armnetwork.ApplicationGatewayPrivateEndpointConnection{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/privateLinkResources/connection1"),
		// 			Name: to.Ptr("coonection1"),
		// 			Type: to.Ptr("Microsoft.Network/applicationGateways/privateEndpointConnections"),
		// 			Properties: &armnetwork.ApplicationGatewayPrivateEndpointConnectionProperties{
		// 				LinkIdentifier: to.Ptr("805319460"),
		// 				PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid2/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/pe1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Approval Done"),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
