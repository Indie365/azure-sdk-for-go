//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerEncryptionProtector.json
func ExampleWorkspaceManagedSQLServerEncryptionProtectorClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceManagedSQLServerEncryptionProtectorClient().Get(ctx, "wsg-7398", "testWorkspace", armsynapse.EncryptionProtectorNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EncryptionProtector = armsynapse.EncryptionProtector{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/encryptionProtector"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/encryptionProtector/current"),
	// 	Kind: to.Ptr("azurekeyvault"),
	// 	Location: to.Ptr("Japan East"),
	// 	Properties: &armsynapse.EncryptionProtectorProperties{
	// 		ServerKeyName: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
	// 		ServerKeyType: to.Ptr(armsynapse.ServerKeyTypeAzureKeyVault),
	// 		URI: to.Ptr("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateKeyVault.json
func ExampleWorkspaceManagedSQLServerEncryptionProtectorClient_BeginCreateOrUpdate_updateTheEncryptionProtectorToKeyVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceManagedSQLServerEncryptionProtectorClient().BeginCreateOrUpdate(ctx, "wsg-7398", "testWorkspace", armsynapse.EncryptionProtectorNameCurrent, armsynapse.EncryptionProtector{
		Properties: &armsynapse.EncryptionProtectorProperties{
			ServerKeyName: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
			ServerKeyType: to.Ptr(armsynapse.ServerKeyTypeAzureKeyVault),
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
	// res.EncryptionProtector = armsynapse.EncryptionProtector{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/encryptionProtector"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/encryptionProtector/current"),
	// 	Kind: to.Ptr("azurekeyvault"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armsynapse.EncryptionProtectorProperties{
	// 		ServerKeyName: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
	// 		ServerKeyType: to.Ptr(armsynapse.ServerKeyTypeAzureKeyVault),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorCreateOrUpdateServiceManaged.json
func ExampleWorkspaceManagedSQLServerEncryptionProtectorClient_BeginCreateOrUpdate_updateTheEncryptionProtectorToServiceManaged() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceManagedSQLServerEncryptionProtectorClient().BeginCreateOrUpdate(ctx, "wsg-7398", "testWorkspace", armsynapse.EncryptionProtectorNameCurrent, armsynapse.EncryptionProtector{
		Properties: &armsynapse.EncryptionProtectorProperties{
			ServerKeyName: to.Ptr("ServiceManaged"),
			ServerKeyType: to.Ptr(armsynapse.ServerKeyTypeServiceManaged),
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
	// res.EncryptionProtector = armsynapse.EncryptionProtector{
	// 	Name: to.Ptr("current"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/encryptionProtector"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/encryptionProtector/current"),
	// 	Kind: to.Ptr("servicemanaged"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armsynapse.EncryptionProtectorProperties{
	// 		ServerKeyName: to.Ptr("ServiceManaged"),
	// 		ServerKeyType: to.Ptr(armsynapse.ServerKeyTypeServiceManaged),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerEncryptionProtectorList.json
func ExampleWorkspaceManagedSQLServerEncryptionProtectorClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceManagedSQLServerEncryptionProtectorClient().NewListPager("wsg-7398", "testWorkspace", nil)
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
		// page.EncryptionProtectorListResult = armsynapse.EncryptionProtectorListResult{
		// 	Value: []*armsynapse.EncryptionProtector{
		// 		{
		// 			Name: to.Ptr("current"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/encryptionProtector"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/wsg-7398/providers/Microsoft.Synapse/workspaces/testWorkspace/encryptionProtector/current"),
		// 			Kind: to.Ptr("azurekeyvault"),
		// 			Location: to.Ptr("Japan East"),
		// 			Properties: &armsynapse.EncryptionProtectorProperties{
		// 				ServerKeyName: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
		// 				ServerKeyType: to.Ptr(armsynapse.ServerKeyTypeAzureKeyVault),
		// 				URI: to.Ptr("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/WorkspaceManagedSqlServerEncryptionProtectorRevalidate.json
func ExampleWorkspaceManagedSQLServerEncryptionProtectorClient_BeginRevalidate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspaceManagedSQLServerEncryptionProtectorClient().BeginRevalidate(ctx, "wsg-7398", "testWorkspace", armsynapse.EncryptionProtectorNameCurrent, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
