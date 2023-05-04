//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountPutEncryptionScope.json
func ExampleEncryptionScopesClient_Put_storageAccountPutEncryptionScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEncryptionScopesClient().Put(ctx, "resource-group-name", "{storage-account-name}", "{encryption-scope-name}", armstorage.EncryptionScope{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EncryptionScope = armstorage.EncryptionScope{
	// 	Name: to.Ptr("{encryption-scope-name}"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/{storage-account-name}/encryptionScopes/{encryption-scope-name}"),
	// 	EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
	// 		Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftStorage),
	// 		State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountPutEncryptionScopeWithInfrastructureEncryption.json
func ExampleEncryptionScopesClient_Put_storageAccountPutEncryptionScopeWithInfrastructureEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEncryptionScopesClient().Put(ctx, "resource-group-name", "{storage-account-name}", "{encryption-scope-name}", armstorage.EncryptionScope{
		EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
			RequireInfrastructureEncryption: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EncryptionScope = armstorage.EncryptionScope{
	// 	Name: to.Ptr("{encryption-scope-name}"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/{storage-account-name}/encryptionScopes/{encryption-scope-name}"),
	// 	EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
	// 		RequireInfrastructureEncryption: to.Ptr(true),
	// 		Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftStorage),
	// 		State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountPatchEncryptionScope.json
func ExampleEncryptionScopesClient_Patch() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEncryptionScopesClient().Patch(ctx, "resource-group-name", "{storage-account-name}", "{encryption-scope-name}", armstorage.EncryptionScope{
		EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
			KeyVaultProperties: &armstorage.EncryptionScopeKeyVaultProperties{
				KeyURI: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
			},
			Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftKeyVault),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EncryptionScope = armstorage.EncryptionScope{
	// 	Name: to.Ptr("{encryption-scope-name}"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/{storage-account-name}/encryptionScopes/{encryption-scope-name}"),
	// 	EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
	// 		KeyVaultProperties: &armstorage.EncryptionScopeKeyVaultProperties{
	// 			CurrentVersionedKeyIdentifier: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
	// 			KeyURI: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
	// 			LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-13T20:36:23.7023290Z"); return t}()),
	// 		},
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-17T06:23:14.4513306Z"); return t}()),
	// 		Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftKeyVault),
	// 		State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountGetEncryptionScope.json
func ExampleEncryptionScopesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEncryptionScopesClient().Get(ctx, "resource-group-name", "{storage-account-name}", "{encryption-scope-name}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EncryptionScope = armstorage.EncryptionScope{
	// 	Name: to.Ptr("{encyrption-scope-name}"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/{storage-account-name}/encryptionScopes/{encryption-scope-name}"),
	// 	EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
	// 		Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftStorage),
	// 		State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/StorageAccountEncryptionScopeList.json
func ExampleEncryptionScopesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEncryptionScopesClient().NewListPager("resource-group-name", "{storage-account-name}", &armstorage.EncryptionScopesClientListOptions{Maxpagesize: nil,
		Filter:  nil,
		Include: nil,
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
		// page.EncryptionScopeListResult = armstorage.EncryptionScopeListResult{
		// 	Value: []*armstorage.EncryptionScope{
		// 		{
		// 			Name: to.Ptr("scope-1"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/{storage-account-name}/encryptionScopes/scope-1"),
		// 			EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T02:42:41.7633306Z"); return t}()),
		// 				Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftStorage),
		// 				State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("scope-2"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/encryptionScopes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/{storage-account-name}/encryptionScopes/scope-2"),
		// 			EncryptionScopeProperties: &armstorage.EncryptionScopeProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-16T04:32:14.3355306Z"); return t}()),
		// 				KeyVaultProperties: &armstorage.EncryptionScopeKeyVaultProperties{
		// 					KeyURI: to.Ptr("https://testvault.vault.core.windows.net/keys/key1/863425f1358359c"),
		// 				},
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-17T06:23:14.4513306Z"); return t}()),
		// 				Source: to.Ptr(armstorage.EncryptionScopeSourceMicrosoftKeyVault),
		// 				State: to.Ptr(armstorage.EncryptionScopeStateEnabled),
		// 			},
		// 	}},
		// }
	}
}
