//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSubPrivateLinkResources_List.json
func ExamplePrivateLinkResourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListPager("myResourceGroup", "myWebPubSubService", nil)
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
		// page.PrivateLinkResourceList = armwebpubsub.PrivateLinkResourceList{
		// 	Value: []*armwebpubsub.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("myPrivateLink"),
		// 			Type: to.Ptr("privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/privateLinkResources/myPrivateLink"),
		// 			Properties: &armwebpubsub.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("webpubsub"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("webpubsub")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.webpubsub.azure.com")},
		// 						ShareablePrivateLinkResourceTypes: []*armwebpubsub.ShareablePrivateLinkResourceType{
		// 							{
		// 								Name: to.Ptr("site"),
		// 								Properties: &armwebpubsub.ShareablePrivateLinkResourceProperties{
		// 									Type: to.Ptr("Microsoft.Web/sites"),
		// 									Description: to.Ptr("Azure App Service can be used as an upstream"),
		// 									GroupID: to.Ptr("sites"),
		// 								},
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
