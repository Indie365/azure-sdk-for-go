//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e680f9bf4f154ec427ba7feac432ed8efd9665d0/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/RackSkus_ListBySubscription.json
func ExampleRackSKUsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRackSKUsClient().NewListBySubscriptionPager(nil)
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
		// page.RackSKUList = armnetworkcloud.RackSKUList{
		// 	Value: []*armnetworkcloud.RackSKU{
		// 		{
		// 			Name: to.Ptr("rackSkuName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/rackSkus"),
		// 			ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Properties: &armnetworkcloud.RackSKUProperties{
		// 				Description: to.Ptr("Sample SKU for the single rack."),
		// 				ComputeMachines: []*armnetworkcloud.MachineSKUSlot{
		// 					{
		// 						Properties: &armnetworkcloud.MachineSKUProperties{
		// 							BootstrapProtocol: to.Ptr(armnetworkcloud.BootstrapProtocolPXE),
		// 							CPUCores: to.Ptr[int64](28),
		// 							CPUSockets: to.Ptr[int64](2),
		// 							Disks: []*armnetworkcloud.MachineDisk{
		// 								{
		// 									Type: to.Ptr(armnetworkcloud.DiskTypeSSD),
		// 									CapacityGB: to.Ptr[int64](893),
		// 									Connection: to.Ptr(armnetworkcloud.MachineSKUDiskConnectionTypeSAS),
		// 							}},
		// 							Generation: to.Ptr("8"),
		// 							HardwareVersion: to.Ptr("3"),
		// 							MemoryCapacityGB: to.Ptr[int64](512),
		// 							Model: to.Ptr("PowerEdge R750"),
		// 							NetworkInterfaces: []*armnetworkcloud.NetworkInterface{
		// 								{
		// 									Address: to.Ptr("04:00.0"),
		// 									DeviceConnectionType: to.Ptr(armnetworkcloud.DeviceConnectionTypePCI),
		// 									Model: to.Ptr("Broadcom Gigabit Ethernet BCM5720"),
		// 									PhysicalSlot: to.Ptr[int64](1),
		// 									PortCount: to.Ptr[int64](2),
		// 									PortSpeed: to.Ptr[int64](1),
		// 									Vendor: to.Ptr("Broadcom Corp"),
		// 							}},
		// 							TotalThreads: to.Ptr[int64](112),
		// 							Vendor: to.Ptr("Dell"),
		// 						},
		// 						RackSlot: to.Ptr[int64](1),
		// 				}},
		// 				ControllerMachines: []*armnetworkcloud.MachineSKUSlot{
		// 					{
		// 						Properties: &armnetworkcloud.MachineSKUProperties{
		// 							BootstrapProtocol: to.Ptr(armnetworkcloud.BootstrapProtocolPXE),
		// 							CPUCores: to.Ptr[int64](8),
		// 							CPUSockets: to.Ptr[int64](1),
		// 							Disks: []*armnetworkcloud.MachineDisk{
		// 								{
		// 									Type: to.Ptr(armnetworkcloud.DiskTypeSSD),
		// 									CapacityGB: to.Ptr[int64](893),
		// 									Connection: to.Ptr(armnetworkcloud.MachineSKUDiskConnectionTypeSAS),
		// 							}},
		// 							Generation: to.Ptr("8"),
		// 							HardwareVersion: to.Ptr("3"),
		// 							MemoryCapacityGB: to.Ptr[int64](128),
		// 							Model: to.Ptr("PowerEdge R650"),
		// 							NetworkInterfaces: []*armnetworkcloud.NetworkInterface{
		// 								{
		// 									Address: to.Ptr("04:00.0"),
		// 									DeviceConnectionType: to.Ptr(armnetworkcloud.DeviceConnectionTypePCI),
		// 									Model: to.Ptr("Broadcom Gigabit Ethernet BCM5720"),
		// 									PhysicalSlot: to.Ptr[int64](2),
		// 									PortCount: to.Ptr[int64](2),
		// 									PortSpeed: to.Ptr[int64](1),
		// 									Vendor: to.Ptr("Broadcom Corp"),
		// 							}},
		// 							TotalThreads: to.Ptr[int64](16),
		// 							Vendor: to.Ptr("Dell"),
		// 						},
		// 						RackSlot: to.Ptr[int64](5),
		// 				}},
		// 				MaxClusterSlots: to.Ptr[int64](0),
		// 				ProvisioningState: to.Ptr(armnetworkcloud.RackSKUProvisioningStateSucceeded),
		// 				RackType: to.Ptr(armnetworkcloud.RackSKUTypeSingle),
		// 				StorageAppliances: []*armnetworkcloud.StorageApplianceSKUSlot{
		// 					{
		// 						Properties: &armnetworkcloud.StorageApplianceSKUProperties{
		// 							CapacityGB: to.Ptr[int64](9100),
		// 							Model: to.Ptr("x70r3-9"),
		// 						},
		// 						RackSlot: to.Ptr[int64](2),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e680f9bf4f154ec427ba7feac432ed8efd9665d0/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/RackSkus_Get.json
func ExampleRackSKUsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRackSKUsClient().Get(ctx, "rackSkuName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RackSKU = armnetworkcloud.RackSKU{
	// 	Name: to.Ptr("rackSkuName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/rackSkus"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/providers/Microsoft.NetworkCloud/rackSkus/rackSkuName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Properties: &armnetworkcloud.RackSKUProperties{
	// 		Description: to.Ptr("Sample SKU for the single rack."),
	// 		ComputeMachines: []*armnetworkcloud.MachineSKUSlot{
	// 			{
	// 				Properties: &armnetworkcloud.MachineSKUProperties{
	// 					BootstrapProtocol: to.Ptr(armnetworkcloud.BootstrapProtocolPXE),
	// 					CPUCores: to.Ptr[int64](28),
	// 					CPUSockets: to.Ptr[int64](2),
	// 					Disks: []*armnetworkcloud.MachineDisk{
	// 						{
	// 							Type: to.Ptr(armnetworkcloud.DiskTypeSSD),
	// 							CapacityGB: to.Ptr[int64](893),
	// 							Connection: to.Ptr(armnetworkcloud.MachineSKUDiskConnectionTypeSAS),
	// 					}},
	// 					Generation: to.Ptr("8"),
	// 					HardwareVersion: to.Ptr("3"),
	// 					MemoryCapacityGB: to.Ptr[int64](512),
	// 					Model: to.Ptr("PowerEdge R750"),
	// 					NetworkInterfaces: []*armnetworkcloud.NetworkInterface{
	// 						{
	// 							Address: to.Ptr("04:00.0"),
	// 							DeviceConnectionType: to.Ptr(armnetworkcloud.DeviceConnectionTypePCI),
	// 							Model: to.Ptr("Broadcom Gigabit Ethernet BCM5720"),
	// 							PhysicalSlot: to.Ptr[int64](1),
	// 							PortCount: to.Ptr[int64](2),
	// 							PortSpeed: to.Ptr[int64](1),
	// 							Vendor: to.Ptr("Broadcom Corp"),
	// 					}},
	// 					TotalThreads: to.Ptr[int64](112),
	// 					Vendor: to.Ptr("Dell"),
	// 				},
	// 				RackSlot: to.Ptr[int64](1),
	// 		}},
	// 		ControllerMachines: []*armnetworkcloud.MachineSKUSlot{
	// 			{
	// 				Properties: &armnetworkcloud.MachineSKUProperties{
	// 					BootstrapProtocol: to.Ptr(armnetworkcloud.BootstrapProtocolPXE),
	// 					CPUCores: to.Ptr[int64](8),
	// 					CPUSockets: to.Ptr[int64](1),
	// 					Disks: []*armnetworkcloud.MachineDisk{
	// 						{
	// 							Type: to.Ptr(armnetworkcloud.DiskTypeSSD),
	// 							CapacityGB: to.Ptr[int64](893),
	// 							Connection: to.Ptr(armnetworkcloud.MachineSKUDiskConnectionTypeSAS),
	// 					}},
	// 					Generation: to.Ptr("8"),
	// 					HardwareVersion: to.Ptr("3"),
	// 					MemoryCapacityGB: to.Ptr[int64](128),
	// 					Model: to.Ptr("PowerEdge R650"),
	// 					NetworkInterfaces: []*armnetworkcloud.NetworkInterface{
	// 						{
	// 							Address: to.Ptr("04:00.0"),
	// 							DeviceConnectionType: to.Ptr(armnetworkcloud.DeviceConnectionTypePCI),
	// 							Model: to.Ptr("Broadcom Gigabit Ethernet BCM5720"),
	// 							PhysicalSlot: to.Ptr[int64](2),
	// 							PortCount: to.Ptr[int64](2),
	// 							PortSpeed: to.Ptr[int64](1),
	// 							Vendor: to.Ptr("Broadcom Corp"),
	// 					}},
	// 					TotalThreads: to.Ptr[int64](16),
	// 					Vendor: to.Ptr("Dell"),
	// 				},
	// 				RackSlot: to.Ptr[int64](5),
	// 		}},
	// 		MaxClusterSlots: to.Ptr[int64](0),
	// 		ProvisioningState: to.Ptr(armnetworkcloud.RackSKUProvisioningStateSucceeded),
	// 		RackType: to.Ptr(armnetworkcloud.RackSKUTypeSingle),
	// 		StorageAppliances: []*armnetworkcloud.StorageApplianceSKUSlot{
	// 			{
	// 				Properties: &armnetworkcloud.StorageApplianceSKUProperties{
	// 					CapacityGB: to.Ptr[int64](9100),
	// 					Model: to.Ptr("x70r3-9"),
	// 				},
	// 				RackSlot: to.Ptr[int64](2),
	// 		}},
	// 	},
	// }
}
