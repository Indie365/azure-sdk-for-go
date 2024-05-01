//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetManagedShortTermRetentionPolicyRestorableDropped.json
func ExampleManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000", armsql.ManagedShortTermRetentionPolicyNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedBackupShortTermRetentionPolicy = armsql.ManagedBackupShortTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/restorableDroppedDatabases/backupShortTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstance/testsvr/restorableDroppedDatabases/testdb,131403269876900000/backupShortTermRetentionPolicies/default"),
	// 	Properties: &armsql.ManagedBackupShortTermRetentionPolicyProperties{
	// 		RetentionDays: to.Ptr[int32](14),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/UpdateManagedShortTermRetentionPolicyRestorableDropped.json
func ExampleManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient().BeginCreateOrUpdate(ctx, "resourceGroup", "testsvr", "testdb,131403269876900000", armsql.ManagedShortTermRetentionPolicyNameDefault, armsql.ManagedBackupShortTermRetentionPolicy{
		Properties: &armsql.ManagedBackupShortTermRetentionPolicyProperties{
			RetentionDays: to.Ptr[int32](14),
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
	// res.ManagedBackupShortTermRetentionPolicy = armsql.ManagedBackupShortTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/restorableDroppedDatabases/backupShortTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroup/providers/Microsoft.Sql/managedInstances/testsvr/restorableDroppedDatabases/testdb,131403269876900000/backupShortTermRetentionPolicies/default"),
	// 	Properties: &armsql.ManagedBackupShortTermRetentionPolicyProperties{
	// 		RetentionDays: to.Ptr[int32](14),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetListManagedShortTermRetentionPolicyRestorableDropped.json
func ExampleManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient_NewListByRestorableDroppedDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient().NewListByRestorableDroppedDatabasePager("Default-SQL-SouthEastAsia", "testsvr", "testdb,131403269876900000", nil)
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
		// page.ManagedBackupShortTermRetentionPolicyListResult = armsql.ManagedBackupShortTermRetentionPolicyListResult{
		// 	Value: []*armsql.ManagedBackupShortTermRetentionPolicy{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/restorableDroppedDatabases/backupShortTermRetentionPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/restorableDroppedDatabases/testdb,131403269876900000/backupShortTermRetentionPolicies/default"),
		// 			Properties: &armsql.ManagedBackupShortTermRetentionPolicyProperties{
		// 				RetentionDays: to.Ptr[int32](14),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/restorableDroppedDatabases/backupShortTermRetentionPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/restorableDroppedDatabases/testdb,131403269876900000/backupShortTermRetentionPolicies/default"),
		// 			Properties: &armsql.ManagedBackupShortTermRetentionPolicyProperties{
		// 				RetentionDays: to.Ptr[int32](14),
		// 			},
		// 	}},
		// }
	}
}
