//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/BackupPolicies_List.json
func ExampleBackupPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewBackupPoliciesClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("myRG", "account1", nil)
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
		// page.BackupPoliciesList = armnetapp.BackupPoliciesList{
		// 	Value: []*armnetapp.BackupPolicy{
		// 		{
		// 			Name: to.Ptr("account1/backupPolicy1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicy1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetapp.BackupPolicyProperties{
		// 				DailyBackupsToKeep: to.Ptr[int32](10),
		// 				Enabled: to.Ptr(true),
		// 				MonthlyBackupsToKeep: to.Ptr[int32](10),
		// 				VolumesAssigned: to.Ptr[int32](0),
		// 				WeeklyBackupsToKeep: to.Ptr[int32](10),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/BackupPolicies_Get.json
func ExampleBackupPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewBackupPoliciesClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "myRG", "account1", "backupPolicyName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupPolicy = armnetapp.BackupPolicy{
	// 	Name: to.Ptr("account1/backupPolicyName"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicyName"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetapp.BackupPolicyProperties{
	// 		DailyBackupsToKeep: to.Ptr[int32](10),
	// 		Enabled: to.Ptr(true),
	// 		MonthlyBackupsToKeep: to.Ptr[int32](10),
	// 		VolumeBackups: []*armnetapp.VolumeBackups{
	// 			{
	// 				BackupsCount: to.Ptr[int32](5),
	// 				PolicyEnabled: to.Ptr(true),
	// 				VolumeName: to.Ptr("volume 1"),
	// 		}},
	// 		VolumesAssigned: to.Ptr[int32](0),
	// 		WeeklyBackupsToKeep: to.Ptr[int32](10),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/BackupPolicies_Create.json
func ExampleBackupPoliciesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewBackupPoliciesClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "myRG", "account1", "backupPolicyName", armnetapp.BackupPolicy{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.BackupPolicyProperties{
			DailyBackupsToKeep:   to.Ptr[int32](10),
			Enabled:              to.Ptr(true),
			MonthlyBackupsToKeep: to.Ptr[int32](10),
			WeeklyBackupsToKeep:  to.Ptr[int32](10),
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
	// res.BackupPolicy = armnetapp.BackupPolicy{
	// 	Name: to.Ptr("account1/backupPolicyName"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicyName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetapp.BackupPolicyProperties{
	// 		DailyBackupsToKeep: to.Ptr[int32](10),
	// 		Enabled: to.Ptr(true),
	// 		MonthlyBackupsToKeep: to.Ptr[int32](10),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		WeeklyBackupsToKeep: to.Ptr[int32](10),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/BackupPolicies_Update.json
func ExampleBackupPoliciesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewBackupPoliciesClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myRG", "account1", "backupPolicyName", armnetapp.BackupPolicyPatch{
		Location: to.Ptr("westus"),
		Properties: &armnetapp.BackupPolicyProperties{
			DailyBackupsToKeep:   to.Ptr[int32](5),
			Enabled:              to.Ptr(false),
			MonthlyBackupsToKeep: to.Ptr[int32](10),
			WeeklyBackupsToKeep:  to.Ptr[int32](10),
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
	// res.BackupPolicy = armnetapp.BackupPolicy{
	// 	Name: to.Ptr("account1/backupPolicyName"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/backupPolicies"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/backupPolocies/backupPolicyName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetapp.BackupPolicyProperties{
	// 		DailyBackupsToKeep: to.Ptr[int32](5),
	// 		Enabled: to.Ptr(false),
	// 		MonthlyBackupsToKeep: to.Ptr[int32](10),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		VolumeBackups: []*armnetapp.VolumeBackups{
	// 			{
	// 				BackupsCount: to.Ptr[int32](5),
	// 				PolicyEnabled: to.Ptr(true),
	// 				VolumeName: to.Ptr("volume 1"),
	// 		}},
	// 		VolumesAssigned: to.Ptr[int32](1),
	// 		WeeklyBackupsToKeep: to.Ptr[int32](10),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/BackupPolicies_Delete.json
func ExampleBackupPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetapp.NewBackupPoliciesClient("D633CC2E-722B-4AE1-B636-BBD9E4C60ED9", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "resourceGroup", "accountName", "backupPolicyName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
