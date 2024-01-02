//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionGroupsListByPublisherName.json
func ExampleNetworkFunctionDefinitionGroupsClient_NewListByPublisherPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkFunctionDefinitionGroupsClient().NewListByPublisherPager("rg", "TestPublisher", nil)
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
		// page.NetworkFunctionDefinitionGroupListResult = armhybridnetwork.NetworkFunctionDefinitionGroupListResult{
		// 	Value: []*armhybridnetwork.NetworkFunctionDefinitionGroup{
		// 		{
		// 			Name: to.Ptr("TestNetworkFunctionDefinitionVersion"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/publishers/networkFunctionDefinitionGroups"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/rgproviders/Microsoft.HybridNetwork/publishers/TestPublisher/networkFunctionDefinitionGroups/TestNetworkFunctionDefinitionGroupName"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armhybridnetwork.NetworkFunctionDefinitionGroupPropertiesFormat{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionGroupDelete.json
func ExampleNetworkFunctionDefinitionGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFunctionDefinitionGroupsClient().BeginDelete(ctx, "rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionGroupCreate.json
func ExampleNetworkFunctionDefinitionGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFunctionDefinitionGroupsClient().BeginCreateOrUpdate(ctx, "rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupName", armhybridnetwork.NetworkFunctionDefinitionGroup{
		Location: to.Ptr("eastus"),
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
	// res.NetworkFunctionDefinitionGroup = armhybridnetwork.NetworkFunctionDefinitionGroup{
	// 	Name: to.Ptr("TestPublisherSkuVersion"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/publishers/networkFunctionDefinitionGroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/rgproviders/Microsoft.HybridNetwork/publishers/TestPublisher/networkFunctionDefinitionGroups/TestNetworkFunctionDefinitionGroupName"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.NetworkFunctionDefinitionGroupPropertiesFormat{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionGroupGet.json
func ExampleNetworkFunctionDefinitionGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFunctionDefinitionGroupsClient().Get(ctx, "rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFunctionDefinitionGroup = armhybridnetwork.NetworkFunctionDefinitionGroup{
	// 	Name: to.Ptr("TestNetworkFunctionDefinitionVersion"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/publishers/networkFunctionDefinitionGroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/rgproviders/Microsoft.HybridNetwork/publishers/TestPublisher/networkFunctionDefinitionGroups/TestNetworkFunctionDefinitionGroupName"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.NetworkFunctionDefinitionGroupPropertiesFormat{
	// 		Description: to.Ptr("Test NFD group"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionGroupUpdateTags.json
func ExampleNetworkFunctionDefinitionGroupsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFunctionDefinitionGroupsClient().Update(ctx, "rg", "TestPublisher", "TestNetworkFunctionDefinitionGroupName", armhybridnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFunctionDefinitionGroup = armhybridnetwork.NetworkFunctionDefinitionGroup{
	// 	Name: to.Ptr("TestPublisherSkuVersion"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/publishers/networkFunctionDefinitionGroups"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/rgproviders/Microsoft.HybridNetwork/publishers/TestPublisher/networkFunctionDefinitionGroups/TestNetworkFunctionDefinitionGroupName"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armhybridnetwork.NetworkFunctionDefinitionGroupPropertiesFormat{
	// 	},
	// }
}
