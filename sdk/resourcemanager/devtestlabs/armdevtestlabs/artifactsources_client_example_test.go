//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_List.json
func ExampleArtifactSourcesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArtifactSourcesClient().NewListPager("resourceGroupName", "{labName}", &armdevtestlabs.ArtifactSourcesClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
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
		// page.ArtifactSourceList = armdevtestlabs.ArtifactSourceList{
		// 	Value: []*armdevtestlabs.ArtifactSource{
		// 		{
		// 			Name: to.Ptr("{artifactSourceName}"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/artifactsources"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.ArtifactSourceProperties{
		// 				ArmTemplateFolderPath: to.Ptr("{armTemplateFolderPath}"),
		// 				BranchRef: to.Ptr("{branchRef}"),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-16T23:53:02.483Z"); return t}()),
		// 				DisplayName: to.Ptr("{displayName}"),
		// 				FolderPath: to.Ptr("{folderPath}"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SecurityToken: to.Ptr("{securityToken}"),
		// 				SourceType: to.Ptr(armdevtestlabs.SourceControlType("{VsoGit|GitHub|StorageAccount}")),
		// 				Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				URI: to.Ptr("{artifactSourceUri}"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_Get.json
func ExampleArtifactSourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactSourcesClient().Get(ctx, "resourceGroupName", "{labName}", "{artifactSourceName}", &armdevtestlabs.ArtifactSourcesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArtifactSource = armdevtestlabs.ArtifactSource{
	// 	Name: to.Ptr("{artifactSourceName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/artifactsources"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ArtifactSourceProperties{
	// 		ArmTemplateFolderPath: to.Ptr("{armTemplateFolderPath}"),
	// 		BranchRef: to.Ptr("{branchRef}"),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-16T23:53:02.483Z"); return t}()),
	// 		DisplayName: to.Ptr("{displayName}"),
	// 		FolderPath: to.Ptr("{folderPath}"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SecurityToken: to.Ptr("{securityToken}"),
	// 		SourceType: to.Ptr(armdevtestlabs.SourceControlType("{VsoGit|GitHub|StorageAccount}")),
	// 		Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 		URI: to.Ptr("{artifactSourceUri}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_CreateOrUpdate.json
func ExampleArtifactSourcesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactSourcesClient().CreateOrUpdate(ctx, "resourceGroupName", "{labName}", "{artifactSourceName}", armdevtestlabs.ArtifactSource{
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
		Properties: &armdevtestlabs.ArtifactSourceProperties{
			ArmTemplateFolderPath: to.Ptr("{armTemplateFolderPath}"),
			BranchRef:             to.Ptr("{branchRef}"),
			DisplayName:           to.Ptr("{displayName}"),
			FolderPath:            to.Ptr("{folderPath}"),
			SecurityToken:         to.Ptr("{securityToken}"),
			SourceType:            to.Ptr(armdevtestlabs.SourceControlType("{VsoGit|GitHub|StorageAccount}")),
			Status:                to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
			URI:                   to.Ptr("{artifactSourceUri}"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArtifactSource = armdevtestlabs.ArtifactSource{
	// 	Name: to.Ptr("{artifactSourceName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/artifactsources"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ArtifactSourceProperties{
	// 		ArmTemplateFolderPath: to.Ptr("{armTemplateFolderPath}"),
	// 		BranchRef: to.Ptr("{branchRef}"),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-16T23:53:02.483Z"); return t}()),
	// 		DisplayName: to.Ptr("{displayName}"),
	// 		FolderPath: to.Ptr("{folderPath}"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SecurityToken: to.Ptr("{securityToken}"),
	// 		SourceType: to.Ptr(armdevtestlabs.SourceControlType("{VsoGit|GitHub|StorageAccount}")),
	// 		Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 		URI: to.Ptr("{artifactSourceUri}"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_Delete.json
func ExampleArtifactSourcesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewArtifactSourcesClient().Delete(ctx, "resourceGroupName", "{labName}", "{artifactSourceName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_Update.json
func ExampleArtifactSourcesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArtifactSourcesClient().Update(ctx, "resourceGroupName", "{labName}", "{artifactSourceName}", armdevtestlabs.ArtifactSourceFragment{
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArtifactSource = armdevtestlabs.ArtifactSource{
	// 	Name: to.Ptr("{artifactSourceName}"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/artifactsources"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactsources/{artifactSourceName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ArtifactSourceProperties{
	// 		ArmTemplateFolderPath: to.Ptr("{armTemplateFolderPath}"),
	// 		BranchRef: to.Ptr("{branchRef}"),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-16T23:53:02.483Z"); return t}()),
	// 		DisplayName: to.Ptr("{displayName}"),
	// 		FolderPath: to.Ptr("{folderPath}"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SecurityToken: to.Ptr("{securityToken}"),
	// 		SourceType: to.Ptr(armdevtestlabs.SourceControlType("{VsoGit|GitHub|StorageAccount}")),
	// 		Status: to.Ptr(armdevtestlabs.EnableStatus("{Enabled|Disabled}")),
	// 		UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
	// 		URI: to.Ptr("{artifactSourceUri}"),
	// 	},
	// }
}
