//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_List.json
func ExampleTasksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		&armdatamigration.TasksClientListOptions{TaskType: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_CreateOrUpdate.json
func ExampleTasksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		"DmsSdkTask",
		armdatamigration.ProjectTask{
			Properties: &armdatamigration.ConnectToTargetSQLDbTaskProperties{
				TaskType: to.Ptr("ConnectToTarget.SqlDb"),
				Input: &armdatamigration.ConnectToTargetSQLDbTaskInput{
					TargetConnectionInfo: &armdatamigration.SQLConnectionInfo{
						Type:                   to.Ptr("SqlConnectionInfo"),
						Password:               to.Ptr("testpassword"),
						UserName:               to.Ptr("testuser"),
						Authentication:         to.Ptr(armdatamigration.AuthenticationTypeSQLAuthentication),
						DataSource:             to.Ptr("ssma-test-server.database.windows.net"),
						EncryptConnection:      to.Ptr(true),
						TrustServerCertificate: to.Ptr(true),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_Get.json
func ExampleTasksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		"DmsSdkTask",
		&armdatamigration.TasksClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_Delete.json
func ExampleTasksClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		"DmsSdkTask",
		&armdatamigration.TasksClientDeleteOptions{DeleteRunningTasks: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_Update.json
func ExampleTasksClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		"DmsSdkTask",
		armdatamigration.ProjectTask{
			Properties: &armdatamigration.ConnectToTargetSQLDbTaskProperties{
				TaskType: to.Ptr("ConnectToTarget.SqlDb"),
				Input: &armdatamigration.ConnectToTargetSQLDbTaskInput{
					TargetConnectionInfo: &armdatamigration.SQLConnectionInfo{
						Type:                   to.Ptr("SqlConnectionInfo"),
						Password:               to.Ptr("testpassword"),
						UserName:               to.Ptr("testuser"),
						Authentication:         to.Ptr(armdatamigration.AuthenticationTypeSQLAuthentication),
						DataSource:             to.Ptr("ssma-test-server.database.windows.net"),
						EncryptConnection:      to.Ptr(true),
						TrustServerCertificate: to.Ptr(true),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_Cancel.json
func ExampleTasksClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Cancel(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		"DmsSdkTask",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_Command.json
func ExampleTasksClient_Command() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Command(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		"DmsSdkTask",
		&armdatamigration.MigrateSyncCompleteCommandProperties{
			CommandType: to.Ptr("Migrate.Sync.Complete.Database"),
			Input: &armdatamigration.MigrateSyncCompleteCommandInput{
				DatabaseName: to.Ptr("TestDatabase"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}