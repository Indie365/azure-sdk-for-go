//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/watchlists/GetWatchlistItems.json
func ExampleWatchlistItemsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWatchlistItemsClient().NewListPager("myRg", "myWorkspace", "highValueAsset", &armsecurityinsights.WatchlistItemsClientListOptions{SkipToken: nil})
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
		// page.WatchlistItemList = armsecurityinsights.WatchlistItemList{
		// 	Value: []*armsecurityinsights.WatchlistItem{
		// 		{
		// 			Name: to.Ptr("fd37d325-7090-47fe-851a-5b5a00c3f576"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/Watchlists/WatchlistItems"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Watchlists/highValueAsset/WatchlistItems/fd37d325-7090-47fe-851a-5b5a00c3f576"),
		// 			Etag: to.Ptr("\"f2089bfa-0000-0d00-0000-601c58b42021\""),
		// 			Properties: &armsecurityinsights.WatchlistItemProperties{
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-04T12:27:32.3783333-08:00"); return t}()),
		// 				CreatedBy: &armsecurityinsights.UserInfo{
		// 					Name: to.Ptr("john doe"),
		// 					Email: to.Ptr("john@contoso.com"),
		// 					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 				},
		// 				EntityMapping: map[string]any{
		// 				},
		// 				IsDeleted: to.Ptr(false),
		// 				ItemsKeyValue: map[string]any{
		// 					"Header-1": "v1_1",
		// 					"Header-2": "v1_2",
		// 					"Header-3": "v1_3",
		// 					"Header-4": "v1_4",
		// 					"Header-5": "v1_5",
		// 				},
		// 				TenantID: to.Ptr("3f8901fe-63d9-4875-9ad5-9fb3b8105797"),
		// 				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-04T12:27:32.3783333-08:00"); return t}()),
		// 				UpdatedBy: &armsecurityinsights.UserInfo{
		// 					Name: to.Ptr("john doe"),
		// 					Email: to.Ptr("john@contoso.com"),
		// 					ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 				},
		// 				WatchlistItemID: to.Ptr("fd37d325-7090-47fe-851a-5b5a00c3f576"),
		// 				WatchlistItemType: to.Ptr("watchlist-item"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/watchlists/GetWatchlistItemById.json
func ExampleWatchlistItemsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistItemsClient().Get(ctx, "myRg", "myWorkspace", "highValueAsset", "3f8901fe-63d9-4875-9ad5-9fb3b8105797", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WatchlistItem = armsecurityinsights.WatchlistItem{
	// 	Name: to.Ptr("fd37d325-7090-47fe-851a-5b5a00c3f576"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists/WatchlistItems"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Watchlists/highValueAsset/WatchlistItems/fd37d325-7090-47fe-851a-5b5a00c3f576"),
	// 	Etag: to.Ptr("\"f2089bfa-0000-0d00-0000-601c58b42021\""),
	// 	Properties: &armsecurityinsights.WatchlistItemProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-04T12:27:32.3783333-08:00"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		EntityMapping: map[string]any{
	// 		},
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsKeyValue: map[string]any{
	// 			"Header-1": "v1_1",
	// 			"Header-2": "v1_2",
	// 			"Header-3": "v1_3",
	// 			"Header-4": "v1_4",
	// 			"Header-5": "v1_5",
	// 		},
	// 		TenantID: to.Ptr("3f8901fe-63d9-4875-9ad5-9fb3b8105797"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-02-04T12:27:32.3783333-08:00"); return t}()),
	// 		UpdatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		WatchlistItemID: to.Ptr("fd37d325-7090-47fe-851a-5b5a00c3f576"),
	// 		WatchlistItemType: to.Ptr("watchlist-item"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/watchlists/DeleteWatchlistItem.json
func ExampleWatchlistItemsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWatchlistItemsClient().Delete(ctx, "myRg", "myWorkspace", "highValueAsset", "4008512e-1d30-48b2-9ee2-d3612ed9d3ea", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/watchlists/CreateWatchlistItem.json
func ExampleWatchlistItemsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWatchlistItemsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "highValueAsset", "82ba292c-dc97-4dfc-969d-d4dd9e666842", armsecurityinsights.WatchlistItem{
		Etag: to.Ptr("0300bf09-0000-0000-0000-5c37296e0000"),
		Properties: &armsecurityinsights.WatchlistItemProperties{
			ItemsKeyValue: map[string]any{
				"Business tier":  "10.0.2.0/24",
				"Data tier":      "10.0.2.0/24",
				"Gateway subnet": "10.0.255.224/27",
				"Private DMZ in": "10.0.0.0/27",
				"Public DMZ out": "10.0.0.96/27",
				"Web Tier":       "10.0.1.0/24",
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WatchlistItem = armsecurityinsights.WatchlistItem{
	// 	Type: to.Ptr("Microsoft.SecurityInsights/Watchlists/WatchlistItems"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Watchlists/highValueAsset/WatchlistItems/82ba292c-dc97-4dfc-969d-d4dd9e666842"),
	// 	Etag: to.Ptr("0300bf09-0000-0000-0000-5c37296e0000"),
	// 	Properties: &armsecurityinsights.WatchlistItemProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-15T04:58:56.0748363+00:00"); return t}()),
	// 		CreatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		IsDeleted: to.Ptr(false),
	// 		ItemsKeyValue: map[string]any{
	// 			"Business tier": "10.0.2.0/24",
	// 			"Data tier": "10.0.2.0/24",
	// 			"Gateway subnet": "10.0.255.224/27",
	// 			"Private DMZ in": "10.0.0.0/27",
	// 			"Public DMZ out": "10.0.0.96/27",
	// 			"Web Tier": "10.0.1.0/24",
	// 		},
	// 		TenantID: to.Ptr("4008512e-1d30-48b2-9ee2-d3612ed9d3ea"),
	// 		Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-16T16:05:20+00:00"); return t}()),
	// 		UpdatedBy: &armsecurityinsights.UserInfo{
	// 			Name: to.Ptr("john doe"),
	// 			Email: to.Ptr("john@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 		},
	// 		WatchlistItemID: to.Ptr("82ba292c-dc97-4dfc-969d-d4dd9e666842"),
	// 		WatchlistItemType: to.Ptr("watchlist-item"),
	// 	},
	// }
}
