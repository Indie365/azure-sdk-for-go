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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_List.json
func ExampleFilesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFilesClient().NewListPager("DmsSdkRg", "DmsSdkService", "DmsSdkProject", nil)
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
		// page.FileList = armdatamigration.FileList{
		// 	Value: []*armdatamigration.ProjectFile{
		// 		{
		// 			Name: to.Ptr("x114d023d8"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services/projects/files"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/files/x114d023d8"),
		// 			Properties: &armdatamigration.ProjectFileProperties{
		// 				Extension: to.Ptr("sql"),
		// 				FilePath: to.Ptr("SchemaInput/DmsSdkFile.sql"),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T20:01:33.000Z"); return t}()),
		// 				MediaType: to.Ptr("text/plain"),
		// 				Size: to.Ptr[int64](51835),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pfpu7fxqcpziyg2h3qj2vb7d8jpbbg7p"),
		// 			Type: to.Ptr("Microsoft.DataMigration/services/projects/files"),
		// 			ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/files/pfpu7fxqcpziyg2h3qj2vb7d8jpbbg7p"),
		// 			Properties: &armdatamigration.ProjectFileProperties{
		// 				Extension: to.Ptr("sql"),
		// 				FilePath: to.Ptr("7daf055f-82b0-483e-9da0-4244b4eaa9b0/AdventureWorks2008.sql"),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-05T16:51:03.000Z"); return t}()),
		// 				MediaType: to.Ptr("text/plain"),
		// 				Size: to.Ptr[int64](910278),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_Get.json
func ExampleFilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().Get(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectFile = armdatamigration.ProjectFile{
	// 	Name: to.Ptr("x114d023d8"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects/files"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/files/x114d023d8"),
	// 	Properties: &armdatamigration.ProjectFileProperties{
	// 		Extension: to.Ptr("sql"),
	// 		FilePath: to.Ptr("SchemaInput/DmsSdkFile.sql"),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-15T20:01:33.000Z"); return t}()),
	// 		MediaType: to.Ptr("text/plain"),
	// 		Size: to.Ptr[int64](51835),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_CreateOrUpdate.json
func ExampleFilesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().CreateOrUpdate(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", armdatamigration.ProjectFile{
		Properties: &armdatamigration.ProjectFileProperties{
			FilePath: to.Ptr("DmsSdkFilePath/DmsSdkFile.sql"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectFile = armdatamigration.ProjectFile{
	// 	Name: to.Ptr("x114d023d8"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects/files"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/files/x114d023d8"),
	// 	Etag: to.Ptr("C2WE6C3yt2I0hunjpjzffY8LhTLqrJZHJ20gkuq2ZOA="),
	// 	Properties: &armdatamigration.ProjectFileProperties{
	// 		Extension: to.Ptr("sql"),
	// 		FilePath: to.Ptr("DmsSdkFilePath/DmsSdkFile.sql"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_Delete.json
func ExampleFilesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFilesClient().Delete(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_Update.json
func ExampleFilesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().Update(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", armdatamigration.ProjectFile{
		Properties: &armdatamigration.ProjectFileProperties{
			FilePath: to.Ptr("DmsSdkFilePath/DmsSdkFile.sql"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectFile = armdatamigration.ProjectFile{
	// 	Name: to.Ptr("x114d023d8"),
	// 	Type: to.Ptr("Microsoft.DataMigration/services/projects/files"),
	// 	ID: to.Ptr("/subscriptions/fc04246f-04c5-437e-ac5e-206a19e7193f/resourceGroups/DmsSdkRg/providers/Microsoft.DataMigration/services/DmsSdkService/projects/DmsSdkProject/files/x114d023d8"),
	// 	Etag: to.Ptr("C2WE6C3yt2I0hunjpjzffY8LhTLqrJZHJ20gkuq2ZOA="),
	// 	Properties: &armdatamigration.ProjectFileProperties{
	// 		Extension: to.Ptr("sql"),
	// 		FilePath: to.Ptr("DmsSdkFilePath/DmsSdkFile.sql"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_Read.json
func ExampleFilesClient_Read() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().Read(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileStorageInfo = armdatamigration.FileStorageInfo{
	// 	Headers: map[string]*string{
	// 		"x-ms-blob-type": to.Ptr("BlockBlob"),
	// 	},
	// 	URI: to.Ptr("https://dmssdkservicestorage.blob.core.windows.net/dmssdkservicecontainer/_rpfiles/dmssdkproject/pfpu7fxqcpziyg2h3qj2vb7d8jpbbg7p?sv=2016-05-31&sr=b&sig=sassignature&se=2018-10-05T18%3A21%3A42Z&sp=r"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Files_ReadWrite.json
func ExampleFilesClient_ReadWrite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFilesClient().ReadWrite(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FileStorageInfo = armdatamigration.FileStorageInfo{
	// 	Headers: map[string]*string{
	// 		"x-ms-blob-type": to.Ptr("BlockBlob"),
	// 	},
	// 	URI: to.Ptr("https://dmssdkservicestorage.blob.core.windows.net/dmssdkservicecontainer/_rpfiles/dmssdkproject/pfpu7fxqcpziyg2h3qj2vb7d8jpbbg7p?sv=2016-05-31&sr=b&sig=sassignature&se=2018-10-05T18%3A21%3A42Z&sp=racwd"),
	// }
}
