//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_List.json
func ExampleVirtualNetworksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworksClient().NewListPager("resourceGroupName", "{labName}", &armdevtestlabs.VirtualNetworksClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.VirtualNetworkList = armdevtestlabs.VirtualNetworkList{
		// 	Value: []*armdevtestlabs.VirtualNetwork{
		// 		{
		// 			Name: to.Ptr("{virtualNetworkName}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/virtualNetworks"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.VirtualNetworkProperties{
		// 				AllowedSubnets: []*armdevtestlabs.Subnet{
		// 					{
		// 						AllowPublicIP: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
		// 						LabSubnetName: to.Ptr("{virtualNetworkName}Subnet"),
		// 						ResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{virtualNetworkName}Subnet"),
		// 				}},
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T13:01:44.6005134-07:00"); return t}()),
		// 				ExternalProviderResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SubnetOverrides: []*armdevtestlabs.SubnetOverride{
		// 					{
		// 						LabSubnetName: to.Ptr("{virtualNetworkName}Subnet"),
		// 						ResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{virtualNetworkName}Subnet"),
		// 						SharedPublicIPAddressConfiguration: &armdevtestlabs.SubnetSharedPublicIPAddressConfiguration{
		// 							AllowedPorts: []*armdevtestlabs.Port{
		// 								{
		// 									BackendPort: to.Ptr[int32](3389),
		// 									TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
		// 								},
		// 								{
		// 									BackendPort: to.Ptr[int32](22),
		// 									TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
		// 							}},
		// 						},
		// 						UseInVMCreationPermission: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
		// 						UsePublicIPAddressPermission: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
		// 				}},
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_Get.json
func ExampleVirtualNetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworksClient().Get(ctx, "resourceGroupName", "{labName}", "{virtualNetworkName}", &armdevtestlabs.VirtualNetworksClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armdevtestlabs.VirtualNetwork{
	// 	Name: to.Ptr("{virtualNetworkName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/virtualNetworks"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.VirtualNetworkProperties{
	// 		AllowedSubnets: []*armdevtestlabs.Subnet{
	// 			{
	// 				AllowPublicIP: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
	// 				LabSubnetName: to.Ptr("{virtualNetworkName}Subnet"),
	// 				ResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{virtualNetworkName}Subnet"),
	// 		}},
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T13:01:44.6005134-07:00"); return t}()),
	// 		ExternalProviderResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SubnetOverrides: []*armdevtestlabs.SubnetOverride{
	// 			{
	// 				LabSubnetName: to.Ptr("{virtualNetworkName}Subnet"),
	// 				ResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{virtualNetworkName}Subnet"),
	// 				SharedPublicIPAddressConfiguration: &armdevtestlabs.SubnetSharedPublicIPAddressConfiguration{
	// 					AllowedPorts: []*armdevtestlabs.Port{
	// 						{
	// 							BackendPort: to.Ptr[int32](3389),
	// 							TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
	// 						},
	// 						{
	// 							BackendPort: to.Ptr[int32](22),
	// 							TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
	// 					}},
	// 				},
	// 				UseInVMCreationPermission: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
	// 				UsePublicIPAddressPermission: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
	// 		}},
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_CreateOrUpdate.json
func ExampleVirtualNetworksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworksClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "{labName}", "{virtualNetworkName}", armdevtestlabs.VirtualNetwork{
		Location: to.Ptr("{location}"),
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
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
	// res.VirtualNetwork = armdevtestlabs.VirtualNetwork{
	// 	Name: to.Ptr("{virtualNetworkName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/virtualNetworks"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.VirtualNetworkProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T13:01:44.6005134-07:00"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_Delete.json
func ExampleVirtualNetworksClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualNetworksClient().BeginDelete(ctx, "resourceGroupName", "{labName}", "{virtualNetworkName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualNetworks_Update.json
func ExampleVirtualNetworksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworksClient().Update(ctx, "resourceGroupName", "{labName}", "{virtualNetworkName}", armdevtestlabs.VirtualNetworkFragment{
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetwork = armdevtestlabs.VirtualNetwork{
	// 	Name: to.Ptr("{virtualNetworkName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/virtualNetworks"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.VirtualNetworkProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-01T13:01:44.6005134-07:00"); return t}()),
	// 		ExternalProviderResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SubnetOverrides: []*armdevtestlabs.SubnetOverride{
	// 			{
	// 				LabSubnetName: to.Ptr("{virtualNetworkName}Subnet"),
	// 				ResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{virtualNetworkName}Subnet"),
	// 				SharedPublicIPAddressConfiguration: &armdevtestlabs.SubnetSharedPublicIPAddressConfiguration{
	// 					AllowedPorts: []*armdevtestlabs.Port{
	// 						{
	// 							BackendPort: to.Ptr[int32](3389),
	// 							TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
	// 						},
	// 						{
	// 							BackendPort: to.Ptr[int32](22),
	// 							TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
	// 					}},
	// 				},
	// 				UseInVMCreationPermission: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
	// 				UsePublicIPAddressPermission: to.Ptr(armdevtestlabs.UsagePermissionTypeAllow),
	// 		}},
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 	},
	// }
}
