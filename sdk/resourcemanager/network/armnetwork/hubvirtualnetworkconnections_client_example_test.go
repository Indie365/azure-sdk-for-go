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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/HubVirtualNetworkConnectionPut.json
func ExampleHubVirtualNetworkConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewHubVirtualNetworkConnectionsClient().BeginCreateOrUpdate(ctx, "rg1", "virtualHub1", "connection1", armnetwork.HubVirtualNetworkConnection{
		Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
			EnableInternetSecurity: to.Ptr(false),
			RemoteVirtualNetwork: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1"),
			},
			RoutingConfiguration: &armnetwork.RoutingConfiguration{
				AssociatedRouteTable: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
				},
				InboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
				},
				OutboundRouteMap: &armnetwork.SubResource{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
				},
				PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
					IDs: []*armnetwork.SubResource{
						{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
						}},
					Labels: []*string{
						to.Ptr("label1"),
						to.Ptr("label2")},
				},
				VnetRoutes: &armnetwork.VnetRoute{
					StaticRoutes: []*armnetwork.StaticRoute{
						{
							Name: to.Ptr("route1"),
							AddressPrefixes: []*string{
								to.Ptr("10.1.0.0/16"),
								to.Ptr("10.2.0.0/16")},
							NextHopIPAddress: to.Ptr("10.0.0.68"),
						},
						{
							Name: to.Ptr("route2"),
							AddressPrefixes: []*string{
								to.Ptr("10.3.0.0/16"),
								to.Ptr("10.4.0.0/16")},
							NextHopIPAddress: to.Ptr("10.0.0.65"),
						}},
					StaticRoutesConfig: &armnetwork.StaticRoutesConfig{
						VnetLocalRouteOverrideCriteria: to.Ptr(armnetwork.VnetLocalRouteOverrideCriteriaEqual),
					},
				},
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
	// res.HubVirtualNetworkConnection = armnetwork.HubVirtualNetworkConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubVirtualNetworkConnections/connection1"),
	// 	Name: to.Ptr("connection1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
	// 		EnableInternetSecurity: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteVirtualNetwork: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1"),
	// 		},
	// 		RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 			AssociatedRouteTable: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 			},
	// 			InboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 			},
	// 			OutboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 			},
	// 			PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 				IDs: []*armnetwork.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 				}},
	// 				Labels: []*string{
	// 					to.Ptr("label1"),
	// 					to.Ptr("label2")},
	// 				},
	// 				VnetRoutes: &armnetwork.VnetRoute{
	// 					BgpConnections: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/bgpConn1"),
	// 					}},
	// 					StaticRoutes: []*armnetwork.StaticRoute{
	// 						{
	// 							Name: to.Ptr("route1"),
	// 							AddressPrefixes: []*string{
	// 								to.Ptr("10.1.0.0/16"),
	// 								to.Ptr("10.2.0.0/16")},
	// 								NextHopIPAddress: to.Ptr("10.0.0.68"),
	// 							},
	// 							{
	// 								Name: to.Ptr("route2"),
	// 								AddressPrefixes: []*string{
	// 									to.Ptr("10.3.0.0/16"),
	// 									to.Ptr("10.4.0.0/16")},
	// 									NextHopIPAddress: to.Ptr("10.0.0.65"),
	// 							}},
	// 							StaticRoutesConfig: &armnetwork.StaticRoutesConfig{
	// 								PropagateStaticRoutes: to.Ptr(true),
	// 								VnetLocalRouteOverrideCriteria: to.Ptr(armnetwork.VnetLocalRouteOverrideCriteriaEqual),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/HubVirtualNetworkConnectionDelete.json
func ExampleHubVirtualNetworkConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewHubVirtualNetworkConnectionsClient().BeginDelete(ctx, "rg1", "virtualHub1", "connection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/HubVirtualNetworkConnectionGet.json
func ExampleHubVirtualNetworkConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewHubVirtualNetworkConnectionsClient().Get(ctx, "rg1", "virtualHub1", "connection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.HubVirtualNetworkConnection = armnetwork.HubVirtualNetworkConnection{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/virtualHubVnetConnections/connection1"),
	// 	Name: to.Ptr("connection1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.HubVirtualNetworkConnectionProperties{
	// 		EnableInternetSecurity: to.Ptr(false),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RemoteVirtualNetwork: &armnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1"),
	// 		},
	// 		RoutingConfiguration: &armnetwork.RoutingConfiguration{
	// 			AssociatedRouteTable: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 			},
	// 			InboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
	// 			},
	// 			OutboundRouteMap: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
	// 			},
	// 			PropagatedRouteTables: &armnetwork.PropagatedRouteTable{
	// 				IDs: []*armnetwork.SubResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
	// 					},
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
	// 				}},
	// 				Labels: []*string{
	// 					to.Ptr("label1"),
	// 					to.Ptr("label2")},
	// 				},
	// 				VnetRoutes: &armnetwork.VnetRoute{
	// 					BgpConnections: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/bgpConn1"),
	// 					}},
	// 					StaticRoutes: []*armnetwork.StaticRoute{
	// 						{
	// 							Name: to.Ptr("route1"),
	// 							AddressPrefixes: []*string{
	// 								to.Ptr("10.1.0.0/16"),
	// 								to.Ptr("10.2.0.0/16")},
	// 								NextHopIPAddress: to.Ptr("10.0.0.68"),
	// 							},
	// 							{
	// 								Name: to.Ptr("route2"),
	// 								AddressPrefixes: []*string{
	// 									to.Ptr("10.3.0.0/16"),
	// 									to.Ptr("10.4.0.0/16")},
	// 									NextHopIPAddress: to.Ptr("10.0.0.65"),
	// 							}},
	// 							StaticRoutesConfig: &armnetwork.StaticRoutesConfig{
	// 								PropagateStaticRoutes: to.Ptr(true),
	// 								VnetLocalRouteOverrideCriteria: to.Ptr(armnetwork.VnetLocalRouteOverrideCriteriaEqual),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a60468a0c5e2beb054680ae488fb9f92699f0a0d/specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/HubVirtualNetworkConnectionList.json
func ExampleHubVirtualNetworkConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHubVirtualNetworkConnectionsClient().NewListPager("rg1", "virtualHub1", nil)
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
		// page.ListHubVirtualNetworkConnectionsResult = armnetwork.ListHubVirtualNetworkConnectionsResult{
		// }
	}
}
