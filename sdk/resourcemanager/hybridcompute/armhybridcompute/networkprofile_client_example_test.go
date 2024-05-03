//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71a0c7adf2a6e169ab9a33c7cf36bb93db083e86/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/NetworkProfile_Get.json
func ExampleNetworkProfileClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkProfileClient().Get(ctx, "myResourceGroup", "myMachine", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkProfile = armhybridcompute.NetworkProfile{
	// 	NetworkInterfaces: []*armhybridcompute.NetworkInterface{
	// 		{
	// 			IPAddresses: []*armhybridcompute.IPAddress{
	// 				{
	// 					Address: to.Ptr("192.168.12.345"),
	// 					IPAddressVersion: to.Ptr("IPv4"),
	// 					Subnet: &armhybridcompute.Subnet{
	// 						AddressPrefix: to.Ptr("192.168.12.0/24"),
	// 					},
	// 			}},
	// 		},
	// 		{
	// 			IPAddresses: []*armhybridcompute.IPAddress{
	// 				{
	// 					Address: to.Ptr("1001:0:34aa:5000:1234:aaaa:bbbb:cccc"),
	// 					IPAddressVersion: to.Ptr("IPv6"),
	// 					Subnet: &armhybridcompute.Subnet{
	// 						AddressPrefix: to.Ptr("1001:0:34aa:5000::/64"),
	// 					},
	// 			}},
	// 	}},
	// }
}
