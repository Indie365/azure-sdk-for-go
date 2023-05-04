//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpurview_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/purview/armpurview"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListByResourceGroup.json
func ExampleAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListByResourceGroupPager("SampleResourceGroup", &armpurview.AccountsClientListByResourceGroupOptions{SkipToken: nil})
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
		// page.AccountList = armpurview.AccountList{
		// 	Value: []*armpurview.Account{
		// 		{
		// 			Name: to.Ptr("account1"),
		// 			Type: to.Ptr("Microsoft.Purview/accounts"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
		// 			Location: to.Ptr("West US 2"),
		// 			SystemData: &armpurview.TrackedResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("client-name"),
		// 				LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armpurview.AccountProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByObjectID: to.Ptr("client-objectId"),
		// 				Endpoints: &armpurview.AccountPropertiesEndpoints{
		// 					Catalog: to.Ptr("https://account1.catalog.purview.azure-test.com"),
		// 					Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
		// 					Scan: to.Ptr("https://account1.scan.purview.azure-test.com"),
		// 				},
		// 				FriendlyName: to.Ptr("friendly-account1"),
		// 				ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
		// 			},
		// 			SKU: &armpurview.AccountSKU{
		// 				Name: to.Ptr(armpurview.NameStandard),
		// 				Capacity: to.Ptr[int32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("account2"),
		// 			Type: to.Ptr("Microsoft.Purview/accounts"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account2"),
		// 			Location: to.Ptr("West US 2"),
		// 			SystemData: &armpurview.TrackedResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("client-name"),
		// 				LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armpurview.AccountProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByObjectID: to.Ptr("client-objectId"),
		// 				Endpoints: &armpurview.AccountPropertiesEndpoints{
		// 					Catalog: to.Ptr("https://account2.catalog.purview.azure-test.com"),
		// 					Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
		// 					Scan: to.Ptr("https://account2.scan.purview.azure-test.com"),
		// 				},
		// 				FriendlyName: to.Ptr("friendly-account1"),
		// 				ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
		// 			},
		// 			SKU: &armpurview.AccountSKU{
		// 				Name: to.Ptr(armpurview.NameStandard),
		// 				Capacity: to.Ptr[int32](1),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListBySubscription.json
func ExampleAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListBySubscriptionPager(&armpurview.AccountsClientListBySubscriptionOptions{SkipToken: nil})
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
		// page.AccountList = armpurview.AccountList{
		// 	Value: []*armpurview.Account{
		// 		{
		// 			Name: to.Ptr("account1"),
		// 			Type: to.Ptr("Microsoft.Purview/accounts"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
		// 			Location: to.Ptr("West US 2"),
		// 			SystemData: &armpurview.TrackedResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("client-name"),
		// 				LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armpurview.AccountProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByObjectID: to.Ptr("client-objectId"),
		// 				Endpoints: &armpurview.AccountPropertiesEndpoints{
		// 					Catalog: to.Ptr("https://account1.catalog.purview.azure-test.com"),
		// 					Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
		// 					Scan: to.Ptr("https://account1.scan.purview.azure-test.com"),
		// 				},
		// 				FriendlyName: to.Ptr("friendly-account1"),
		// 				ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
		// 			},
		// 			SKU: &armpurview.AccountSKU{
		// 				Name: to.Ptr(armpurview.NameStandard),
		// 				Capacity: to.Ptr[int32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("account2"),
		// 			Type: to.Ptr("Microsoft.Purview/accounts"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account2"),
		// 			Location: to.Ptr("West US 2"),
		// 			SystemData: &armpurview.TrackedResourceSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("client-name"),
		// 				LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armpurview.AccountProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
		// 				CreatedBy: to.Ptr("client-name"),
		// 				CreatedByObjectID: to.Ptr("client-objectId"),
		// 				Endpoints: &armpurview.AccountPropertiesEndpoints{
		// 					Catalog: to.Ptr("https://account2.catalog.purview.azure-test.com"),
		// 					Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
		// 					Scan: to.Ptr("https://account2.scan.purview.azure-test.com"),
		// 				},
		// 				FriendlyName: to.Ptr("friendly-account1"),
		// 				ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
		// 			},
		// 			SKU: &armpurview.AccountSKU{
		// 				Name: to.Ptr(armpurview.NameStandard),
		// 				Capacity: to.Ptr[int32](1),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Get.json
func ExampleAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().Get(ctx, "SampleResourceGroup", "account1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armpurview.Account{
	// 	Name: to.Ptr("account1"),
	// 	Type: to.Ptr("Microsoft.Purview/accounts"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
	// 	Location: to.Ptr("West US 2"),
	// 	SystemData: &armpurview.TrackedResourceSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("client-name"),
	// 		LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armpurview.AccountProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByObjectID: to.Ptr("client-objectId"),
	// 		Endpoints: &armpurview.AccountPropertiesEndpoints{
	// 			Catalog: to.Ptr("https://account1.catalog.purview.azure-test.com"),
	// 			Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
	// 			Scan: to.Ptr("https://account1.scan.purview.azure-test.com"),
	// 		},
	// 		FriendlyName: to.Ptr("friendly-account1"),
	// 		ManagedResourceGroupName: to.Ptr("managed-rg-mwjotkl"),
	// 		ManagedResources: &armpurview.AccountPropertiesManagedResources{
	// 			EventHubNamespace: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.EventHub/namespaces/atlas-westusdddnbtp"),
	// 			ResourceGroup: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl"),
	// 			StorageAccount: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.Storage/storageAccounts/scanwestustzaagzr"),
	// 		},
	// 		PrivateEndpointConnections: []*armpurview.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("peName-8536c337-7b36-4d67-a7ce-081655baf59e"),
	// 				Type: to.Ptr("Microsoft.Purview/accounts/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1/privateEndpointConnections/peName-8536c337-7b36-4d67-a7ce-081655baf59e"),
	// 				Properties: &armpurview.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armpurview.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/baca8a88-4527-4c35-a13e-b2775ce0d7fc/resourceGroups/nrpResourceGroupName/providers/Microsoft.Network/privateEndpoints/peName"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armpurview.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Please approve my connection, thanks."),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr(armpurview.StatusPending),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
	// 	},
	// 	SKU: &armpurview.AccountSKU{
	// 		Name: to.Ptr(armpurview.NameStandard),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_CreateOrUpdate.json
func ExampleAccountsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreateOrUpdate(ctx, "SampleResourceGroup", "account1", armpurview.Account{
		Location: to.Ptr("West US 2"),
		Properties: &armpurview.AccountProperties{
			ManagedResourceGroupName: to.Ptr("custom-rgname"),
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
	// res.Account = armpurview.Account{
	// 	Name: to.Ptr("account1"),
	// 	Type: to.Ptr("Microsoft.Purview/accounts"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
	// 	Location: to.Ptr("West US 2"),
	// 	SystemData: &armpurview.TrackedResourceSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("client-name"),
	// 		LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armpurview.AccountProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByObjectID: to.Ptr("client-objectId"),
	// 		Endpoints: &armpurview.AccountPropertiesEndpoints{
	// 			Catalog: to.Ptr("https://account1.catalog.purview.azure-test.com"),
	// 			Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
	// 			Scan: to.Ptr("https://account1.scan.purview.azure-test.com"),
	// 		},
	// 		FriendlyName: to.Ptr("friendly-account1"),
	// 		ManagedResourceGroupName: to.Ptr("custom-rgname"),
	// 		ManagedResources: &armpurview.AccountPropertiesManagedResources{
	// 			EventHubNamespace: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/custom-rgname/providers/Microsoft.EventHub/namespaces/atlas-westusdddnbtp"),
	// 			ResourceGroup: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/custom-rgname"),
	// 			StorageAccount: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/custom-rgname/providers/Microsoft.Storage/storageAccounts/scanwestustzaagzr"),
	// 		},
	// 		ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
	// 	},
	// 	SKU: &armpurview.AccountSKU{
	// 		Name: to.Ptr(armpurview.NameStandard),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Delete.json
func ExampleAccountsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginDelete(ctx, "SampleResourceGroup", "account1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_Update.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginUpdate(ctx, "SampleResourceGroup", "account1", armpurview.AccountUpdateParameters{
		Tags: map[string]*string{
			"newTag": to.Ptr("New tag value."),
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
	// res.Account = armpurview.Account{
	// 	Name: to.Ptr("account1"),
	// 	Type: to.Ptr("Microsoft.Purview/accounts"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/SampleResourceGroup/providers/Microsoft.Purview/accounts/account1"),
	// 	Location: to.Ptr("West US 2"),
	// 	SystemData: &armpurview.TrackedResourceSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByType: to.Ptr(armpurview.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-03-16T23:24:34.3430059Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("client-name"),
	// 		LastModifiedByType: to.Ptr(armpurview.LastModifiedByTypeUser),
	// 	},
	// 	Tags: map[string]*string{
	// 		"newTag": to.Ptr("New tag value."),
	// 	},
	// 	Properties: &armpurview.AccountProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-11-22T18:39:58.6929344Z"); return t}()),
	// 		CreatedBy: to.Ptr("client-name"),
	// 		CreatedByObjectID: to.Ptr("client-objectId"),
	// 		Endpoints: &armpurview.AccountPropertiesEndpoints{
	// 			Catalog: to.Ptr("https://account2.catalog.purview.azure-test.com"),
	// 			Guardian: to.Ptr("https://account1.guardian.purview.azure-test.com"),
	// 			Scan: to.Ptr("https://account2.scan.purview.azure-test.com"),
	// 		},
	// 		FriendlyName: to.Ptr("friendly-account1"),
	// 		ManagedResources: &armpurview.AccountPropertiesManagedResources{
	// 			EventHubNamespace: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.EventHub/namespaces/atlas-westusdddnbtp"),
	// 			ResourceGroup: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl"),
	// 			StorageAccount: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/managed-rg-mwjotkl/providers/Microsoft.Storage/storageAccounts/scanwestustzaagzr"),
	// 		},
	// 		ProvisioningState: to.Ptr(armpurview.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armpurview.PublicNetworkAccessEnabled),
	// 	},
	// 	SKU: &armpurview.AccountSKU{
	// 		Name: to.Ptr(armpurview.NameStandard),
	// 		Capacity: to.Ptr[int32](1),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_ListKeys.json
func ExampleAccountsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().ListKeys(ctx, "SampleResourceGroup", "account1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armpurview.AccessKeys{
	// 	AtlasKafkaPrimaryEndpoint: to.Ptr("Endpoint=sb://fake_objectId.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=ASDASasdfmasdf123412341234="),
	// 	AtlasKafkaSecondaryEndpoint: to.Ptr("Endpoint=sb://fake_objectId.servicebus.windows.net/;SharedAccessKeyName=AlternateSharedAccessKey;SharedAccessKey=BSDASasdfmasdf123412341234="),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_AddRootCollectionAdmin.json
func ExampleAccountsClient_AddRootCollectionAdmin() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAccountsClient().AddRootCollectionAdmin(ctx, "SampleResourceGroup", "account1", armpurview.CollectionAdminUpdate{
		ObjectID: to.Ptr("7e8de0e7-2bfc-4e1f-9659-2a5785e4356f"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Accounts_CheckNameAvailability.json
func ExampleAccountsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpurview.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().CheckNameAvailability(ctx, armpurview.CheckNameAvailabilityRequest{
		Name: to.Ptr("account1"),
		Type: to.Ptr("Microsoft.Purview/accounts"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityResult = armpurview.CheckNameAvailabilityResult{
	// 	NameAvailable: to.Ptr(true),
	// }
}
