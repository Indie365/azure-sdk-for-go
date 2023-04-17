//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureStorage/ProtectableContainers_List.json
func ExampleProtectableContainersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProtectableContainersClient().NewListPager("testvault", "testRg", "Azure", &armrecoveryservicesbackup.ProtectableContainersClientListOptions{Filter: to.Ptr("backupManagementType eq 'AzureStorage' and workloadType eq 'AzureFileShare'")})
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
		// page.ProtectableContainerResourceList = armrecoveryservicesbackup.ProtectableContainerResourceList{
		// 	Value: []*armrecoveryservicesbackup.ProtectableContainerResource{
		// 		{
		// 			Name: to.Ptr("StorageContainer;storage;test-rg;testst"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectableContainers"),
		// 			ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.RecoveryServices/vaults/testvault/backupFabrics/Azure/protectableContainers/StorageContainer;storage;test-rg;teststorage"),
		// 			Properties: &armrecoveryservicesbackup.AzureStorageProtectableContainer{
		// 				BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
		// 				ContainerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorage"),
		// 				FriendlyName: to.Ptr("teststorage"),
		// 				HealthStatus: to.Ptr("Healthy"),
		// 				ProtectableContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("StorageContainer;ClassicStorage;test-rg;teststorage"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectableContainers"),
		// 			ID: to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.RecoveryServices/vaults/testvault/backupFabrics/Azure/protectableContainers/StorageContainer;ClassicStorage;test-rg;teststorage"),
		// 			Properties: &armrecoveryservicesbackup.AzureStorageProtectableContainer{
		// 				BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureStorage),
		// 				ContainerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.ClassicStorage/storageAccounts/teststorage"),
		// 				FriendlyName: to.Ptr("teststorage"),
		// 				HealthStatus: to.Ptr("Healthy"),
		// 				ProtectableContainerType: to.Ptr(armrecoveryservicesbackup.ProtectableContainerTypeStorageContainer),
		// 			},
		// 	}},
		// }
	}
}
