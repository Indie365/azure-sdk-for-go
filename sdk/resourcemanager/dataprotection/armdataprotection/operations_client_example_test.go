//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/Operations/List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdataprotection.NewOperationsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
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
		// page.ClientDiscoveryResponse = armdataprotection.ClientDiscoveryResponse{
		// 	Value: []*armdataprotection.ClientDiscoveryValueForSingleAPI{
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/locations/getBackupStatus/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Check Backup Status for Recovery Services Vaults"),
		// 				Operation: to.Ptr("Check Backup Status for Vault"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/write"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Creates a Backup Instance"),
		// 				Operation: to.Ptr("Create a Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/delete"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Deletes the Backup Instance"),
		// 				Operation: to.Ptr("Delete Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns details of the Backup Instance"),
		// 				Operation: to.Ptr("Get Backup Instance Details"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all Backup Instances"),
		// 				Operation: to.Ptr("Get Backup Instances"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/backup/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Performs Backup on the Backup Instance"),
		// 				Operation: to.Ptr("Backup Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/sync/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Sync operation retries last failed operation on backup instance to bring it to a valid state."),
		// 				Operation: to.Ptr("Sync Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/operationResults/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Operation Result for Backup Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Result"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/stopProtection/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Stop Protection operation stops both backup and retention schedules of backup instance. Existing data will be retained forever."),
		// 				Operation: to.Ptr("Stop Protection of Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/suspendBackups/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Suspend Backups operation stops only backups of backup instance. Retention activities will continue and hence data will be ratained as per policy."),
		// 				Operation: to.Ptr("Suspend Backups of Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/resumeProtection/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Resume protection of a ProtectionStopped BI."),
		// 				Operation: to.Ptr("Resume Protection of Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/resumeBackups/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Resume Backups for a BackupsSuspended BI."),
		// 				Operation: to.Ptr("Resume Backups of Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/validateRestore/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Validates for Restore of the Backup Instance"),
		// 				Operation: to.Ptr("Validate for Restore of Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/restore/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Triggers restore on the Backup Instance"),
		// 				Operation: to.Ptr("Restore Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Instance"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupPolicies/write"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Creates Backup Policy"),
		// 				Operation: to.Ptr("Create Backup Policy"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupPolicies/delete"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Deletes the Backup Policy"),
		// 				Operation: to.Ptr("Delete Backup Policy"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupPolicies/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns details of the Backup Policy"),
		// 				Operation: to.Ptr("Get Backup Policy details"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupPolicies/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all Backup Policies"),
		// 				Operation: to.Ptr("Get Backup Policies"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Policies"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupResourceGuardProxies/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get the list of ResourceGuard proxies for a resource"),
		// 				Operation: to.Ptr("Get the list of ResourceGuard proxies for a resource"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guard Proxy"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupResourceGuardProxies/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get ResourceGuard proxy operation gets an object representing the Azure resource of type 'ResourceGuard proxy'"),
		// 				Operation: to.Ptr("Get ResourceGuard proxy"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guard Proxy"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupResourceGuardProxies/write"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Create ResourceGuard proxy operation creates an Azure resource of type 'ResourceGuard Proxy'"),
		// 				Operation: to.Ptr("Create ResourceGuard proxy"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guard Proxy"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupResourceGuardProxies/delete"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("The Delete ResourceGuard proxy operation deletes the specified Azure resource of type 'ResourceGuard proxy'"),
		// 				Operation: to.Ptr("Delete ResourceGuard proxy"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guard Proxy"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupResourceGuardProxies/unlockDelete/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Unlock delete ResourceGuard proxy operation unlocks the next delete critical operation"),
		// 				Operation: to.Ptr("Unlock delete ResourceGuard proxy operation unlocks the next delete critical operation"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guard Proxy"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/recoveryPoints/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns details of the Recovery Point"),
		// 				Operation: to.Ptr("Get Recovery Point Details"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/recoveryPoints/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns all Recovery Points"),
		// 				Operation: to.Ptr("Get Recovery Points"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Recovery Points"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupInstances/findRestorableTimeRanges/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Finds Restorable Time Ranges"),
		// 				Operation: to.Ptr("Find Restorable Time Ranges"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Restorable Time Ranges"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/write"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Create BackupVault operation creates an Azure resource of type 'Backup Vault'"),
		// 				Operation: to.Ptr("Create Backup Vault"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Create BackupVault operation creates an Azure resource of type 'Backup Vault'"),
		// 				Operation: to.Ptr("Create Backup Vault"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/delete"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Create BackupVault operation creates an Azure resource of type 'Backup Vault'"),
		// 				Operation: to.Ptr("Create Backup Vault"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/operationResults/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets Operation Result of a Patch Operation for a Backup Vault"),
		// 				Operation: to.Ptr("Get Operation Result of a Patch Operation for a Backup Vault"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/locations/checkNameAvailability/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Checks if the requested BackupVault Name is Available"),
		// 				Operation: to.Ptr("Check if the requested BackupVault Name is Available"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets list of Backup Vaults in a Resource Group"),
		// 				Operation: to.Ptr("Get Backup Vaults in a Resource Group"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets list of Backup Vaults in a Subscription"),
		// 				Operation: to.Ptr("Get Backup Vaults in a Subscription"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Vaults"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/resourceGuards/write"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Create ResourceGuard operation creates an Azure resource of type 'ResourceGuard'"),
		// 				Operation: to.Ptr("Create ResourceGuard"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/resourceGuards/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("The Get ResourceGuard operation gets an object representing the Azure resource of type 'ResourceGuard'"),
		// 				Operation: to.Ptr("Get ResourceGuard"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/resourceGuards/delete"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("The Delete ResourceGuard operation deletes the specified Azure resource of type 'ResourceGuard'"),
		// 				Operation: to.Ptr("Delete ResourceGuard"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/resourceGuards/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets list of ResourceGuards in a Resource Group"),
		// 				Operation: to.Ptr("Get ResourceGuards in a Resource Group"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/providers/resourceGuards/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets list of ResourceGuards in a Subscription"),
		// 				Operation: to.Ptr("Get ResourceGuards in a Subscription"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/resourceGuards/write"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Update ResouceGuard operation updates an Azure resource of type 'ResourceGuard'"),
		// 				Operation: to.Ptr("Update ResourceGuard"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/resourceGuards/{operationName}/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets ResourceGuard operation request info"),
		// 				Operation: to.Ptr("Get ResourceGuard operation request info"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/resourceGuards/{operationName}/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Gets ResourceGuard default operation request info"),
		// 				Operation: to.Ptr("Get ResourceGuard default operation request info"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Guards"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/providers/locations/checkFeatureSupport/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Validates if a feature is supported"),
		// 				Operation: to.Ptr("Validate if a feature is supported"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Provider Operation"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/locations/operationStatus/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Operation Status for Backup Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Status"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/operationStatus/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Operation Status for Backup Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Status"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/subscriptions/resourceGroups/providers/operationStatus/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Operation Status for Backup Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Status"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Operation Status"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/locations/operationResults/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Returns Backup Operation Result for Backup Vault."),
		// 				Operation: to.Ptr("Get Backup Operation Result"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Operation Results"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/validateForBackup/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Validates for backup of Backup Instance"),
		// 				Operation: to.Ptr("Validate for backup of Backup Instance"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Validate Backup"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/backupVaults/backupJobs/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get Jobs list"),
		// 				Operation: to.Ptr("Backup Jobs"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.RecoveryServices/Vaults/backupJobs/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Get Job details"),
		// 				Operation: to.Ptr("Backup Job Object"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Backup Jobs"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/register/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Registers subscription for given Resource Provider"),
		// 				Operation: to.Ptr("Register Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Provider Operation"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/unregister/action"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Unregisters subscription for given Resource Provider"),
		// 				Operation: to.Ptr("Unregister Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Provider Operation"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DataProtection/operations/read"),
		// 			Display: &armdataprotection.ClientDiscoveryDisplay{
		// 				Description: to.Ptr("Operation returns the list of Operations for a Resource Provider"),
		// 				Operation: to.Ptr("List of Operations"),
		// 				Provider: to.Ptr("Microsoft.DataProtection"),
		// 				Resource: to.Ptr("Resource Provider Operation"),
		// 			},
		// 			Origin: to.Ptr("user"),
		// 	}},
		// }
	}
}
