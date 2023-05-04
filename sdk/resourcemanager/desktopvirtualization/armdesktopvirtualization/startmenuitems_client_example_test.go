//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/StartMenuItem_List.json
func ExampleStartMenuItemsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStartMenuItemsClient().NewListPager("resourceGroup1", "applicationGroup1", &armdesktopvirtualization.StartMenuItemsClientListOptions{PageSize: nil,
		IsDescending: nil,
		InitialSkip:  nil,
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
		// page.StartMenuItemList = armdesktopvirtualization.StartMenuItemList{
		// 	Value: []*armdesktopvirtualization.StartMenuItem{
		// 		{
		// 			Name: to.Ptr("application1"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups/startMenuItems"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1/startMenuItem/application1"),
		// 			Properties: &armdesktopvirtualization.StartMenuItemProperties{
		// 				AppAlias: to.Ptr("word"),
		// 				CommandLineArguments: to.Ptr("arguments"),
		// 				FilePath: to.Ptr("/path/to/file"),
		// 				IconIndex: to.Ptr[int32](1),
		// 				IconPath: to.Ptr("/path/to/icon"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("application2"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/applicationGroups/startMenuItems"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/applicationGroups/applicationGroup1/startMenuItem/application2"),
		// 			Properties: &armdesktopvirtualization.StartMenuItemProperties{
		// 				AppAlias: to.Ptr("excel"),
		// 				CommandLineArguments: to.Ptr("arguments"),
		// 				FilePath: to.Ptr("/path/to/file"),
		// 				IconIndex: to.Ptr[int32](1),
		// 				IconPath: to.Ptr("/path/to/icon"),
		// 			},
		// 	}},
		// }
	}
}
