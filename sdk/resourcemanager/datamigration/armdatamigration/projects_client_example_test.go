//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Projects_List.json
func ExampleProjectsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProjectsClient().NewListPager("DmsSdkRg", "DmsSdkService", nil)
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
		// page.ProjectList = armdatamigration.ProjectList{
		// 	Value: []*armdatamigration.Project{
		// 		{
		// 			Name: to.Ptr("project1"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services/projects"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/project1"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armdatamigration.ProjectProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-14T01:29:56.3049263+00:00"); return t}()),
		// 				ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
		// 				SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
		// 				TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("project2"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services/projects"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/project2"),
		// 			Location: to.Ptr("southcentralus"),
		// 			Properties: &armdatamigration.ProjectProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-14T01:30:05.6185148+00:00"); return t}()),
		// 				ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
		// 				SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
		// 				TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Projects_CreateOrUpdate.json
func ExampleProjectsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().CreateOrUpdate(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", armdatamigration.Project{
		Location: to.Ptr("southcentralus"),
		Properties: &armdatamigration.ProjectProperties{
			SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
			TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armdatamigration.Project{
	// 	Name: to.Ptr("DmsSdkProject"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Properties: &armdatamigration.ProjectProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-03T02:14:54.2458282-07:00"); return t}()),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
	// 		SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
	// 		TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Projects_Get.json
func ExampleProjectsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().Get(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armdatamigration.Project{
	// 	Name: to.Ptr("DmsSdkProject"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Properties: &armdatamigration.ProjectProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-03T02:35:11.6561001-07:00"); return t}()),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
	// 		SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
	// 		TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Projects_Delete.json
func ExampleProjectsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProjectsClient().Delete(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", &armdatamigration.ProjectsClientDeleteOptions{DeleteRunningTasks: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Projects_Update.json
func ExampleProjectsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectsClient().Update(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", armdatamigration.Project{
		Location: to.Ptr("southcentralus"),
		Properties: &armdatamigration.ProjectProperties{
			SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
			TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Project = armdatamigration.Project{
	// 	Name: to.Ptr("DmsSdkProject"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject"),
	// 	Location: to.Ptr("southcentralus"),
	// 	Properties: &armdatamigration.ProjectProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-03T02:14:54.2458282-07:00"); return t}()),
	// 		ProvisioningState: to.Ptr(armdatamigration.ProjectProvisioningStateSucceeded),
	// 		SourcePlatform: to.Ptr(armdatamigration.ProjectSourcePlatformSQL),
	// 		TargetPlatform: to.Ptr(armdatamigration.ProjectTargetPlatformSQLDB),
	// 	},
	// }
}
