//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsqlvirtualmachine_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/GetAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvailabilityGroupListenersClient().Get(ctx, "testrg", "testvmgroup", "agl-test", &armsqlvirtualmachine.AvailabilityGroupListenersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvailabilityGroupListener = armsqlvirtualmachine.AvailabilityGroupListener{
	// 	Name: to.Ptr("agl-test"),
	// 	Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/availabilityGroupListeners"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup/availabilityGroupListeners/agl-test"),
	// 	Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
	// 		AvailabilityGroupName: to.Ptr("ag-test"),
	// 		LoadBalancerConfigurations: []*armsqlvirtualmachine.LoadBalancerConfiguration{
	// 			{
	// 				LoadBalancerResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb-test"),
	// 				PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
	// 					IPAddress: to.Ptr("10.1.0.112"),
	// 					SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
	// 				},
	// 				ProbePort: to.Ptr[int32](59983),
	// 				SQLVirtualMachineInstances: []*string{
	// 					to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3"),
	// 					to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2")},
	// 			}},
	// 			Port: to.Ptr[int32](1433),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_BeginCreateOrUpdate_createsOrUpdatesAnAvailabilityGroupListenerUsingLoadBalancerThisIsUsedForVMsPresentInSingleSubnet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvailabilityGroupListenersClient().BeginCreateOrUpdate(ctx, "testrg", "testvmgroup", "agl-test", armsqlvirtualmachine.AvailabilityGroupListener{
		Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
			AvailabilityGroupName: to.Ptr("ag-test"),
			LoadBalancerConfigurations: []*armsqlvirtualmachine.LoadBalancerConfiguration{
				{
					LoadBalancerResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb-test"),
					PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
						IPAddress:        to.Ptr("10.1.0.112"),
						SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
					},
					ProbePort: to.Ptr[int32](59983),
					SQLVirtualMachineInstances: []*string{
						to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
						to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3")},
				}},
			Port: to.Ptr[int32](1433),
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
	// res.AvailabilityGroupListener = armsqlvirtualmachine.AvailabilityGroupListener{
	// 	Name: to.Ptr("agl-test"),
	// 	Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/availabilityGroupListeners"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup/availabilityGroupListeners/agl-test"),
	// 	Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
	// 		LoadBalancerConfigurations: []*armsqlvirtualmachine.LoadBalancerConfiguration{
	// 			{
	// 				PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
	// 					IPAddress: to.Ptr("10.1.0.113"),
	// 					SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
	// 				},
	// 				SQLVirtualMachineInstances: []*string{
	// 					to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
	// 					to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3")},
	// 			}},
	// 			Port: to.Ptr[int32](1433),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/CreateOrUpdateAvailabilityGroupListenerWithMultiSubnet.json
func ExampleAvailabilityGroupListenersClient_BeginCreateOrUpdate_createsOrUpdatesAnAvailabilityGroupListenerThisIsUsedForVMsPresentInMultiSubnet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvailabilityGroupListenersClient().BeginCreateOrUpdate(ctx, "testrg", "testvmgroup", "agl-test", armsqlvirtualmachine.AvailabilityGroupListener{
		Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
			AvailabilityGroupName: to.Ptr("ag-test"),
			MultiSubnetIPConfigurations: []*armsqlvirtualmachine.MultiSubnetIPConfiguration{
				{
					PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
						IPAddress:        to.Ptr("10.0.0.112"),
						SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
					},
					SQLVirtualMachineInstance: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
				},
				{
					PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
						IPAddress:        to.Ptr("10.0.1.112"),
						SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/alternate"),
					},
					SQLVirtualMachineInstance: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm1"),
				}},
			Port: to.Ptr[int32](1433),
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
	// res.AvailabilityGroupListener = armsqlvirtualmachine.AvailabilityGroupListener{
	// 	Name: to.Ptr("agl-test"),
	// 	Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/availabilityGroupListeners"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup/availabilityGroupListeners/agl-test"),
	// 	Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
	// 		MultiSubnetIPConfigurations: []*armsqlvirtualmachine.MultiSubnetIPConfiguration{
	// 			{
	// 				PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
	// 					IPAddress: to.Ptr("10.0.0.112"),
	// 					SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
	// 				},
	// 				SQLVirtualMachineInstance: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2"),
	// 			},
	// 			{
	// 				PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
	// 					IPAddress: to.Ptr("10.0.1.112"),
	// 					SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/alternate"),
	// 				},
	// 				SQLVirtualMachineInstance: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm1"),
	// 		}},
	// 		Port: to.Ptr[int32](1433),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/DeleteAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAvailabilityGroupListenersClient().BeginDelete(ctx, "testrg", "testvmgroup", "agl-test", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/ListByGroupAvailabilityGroupListener.json
func ExampleAvailabilityGroupListenersClient_NewListByGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsqlvirtualmachine.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvailabilityGroupListenersClient().NewListByGroupPager("testrg", "testvmgroup", nil)
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
		// page.AvailabilityGroupListenerListResult = armsqlvirtualmachine.AvailabilityGroupListenerListResult{
		// 	Value: []*armsqlvirtualmachine.AvailabilityGroupListener{
		// 		{
		// 			Name: to.Ptr("agl-test"),
		// 			Type: to.Ptr("Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/availabilityGroupListeners"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups/testvmgroup/availabilityGroupListeners/agl-test"),
		// 			Properties: &armsqlvirtualmachine.AvailabilityGroupListenerProperties{
		// 				AvailabilityGroupName: to.Ptr("ag-test"),
		// 				LoadBalancerConfigurations: []*armsqlvirtualmachine.LoadBalancerConfiguration{
		// 					{
		// 						LoadBalancerResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lb-test"),
		// 						PrivateIPAddress: &armsqlvirtualmachine.PrivateIPAddress{
		// 							IPAddress: to.Ptr("10.1.0.112"),
		// 							SubnetResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/default"),
		// 						},
		// 						ProbePort: to.Ptr[int32](59983),
		// 						SQLVirtualMachineInstances: []*string{
		// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm3"),
		// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/testvm2")},
		// 					}},
		// 					Port: to.Ptr[int32](1433),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 		}},
		// 	}
	}
}
