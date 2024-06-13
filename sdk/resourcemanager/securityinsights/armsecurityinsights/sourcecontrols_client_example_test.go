//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/sourcecontrols/GetSourceControls.json
func ExampleSourceControlsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSourceControlsClient().NewListPager("myRg", "myWorkspace", nil)
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
		// page.SourceControlList = armsecurityinsights.SourceControlList{
		// 	Value: []*armsecurityinsights.SourceControl{
		// 		{
		// 			Name: to.Ptr("789e0c1f-4a3d-43ad-809c-e713b677b04a"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/SourceControls"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/sourcecontrols/789e0c1f-4a3d-43ad-809c-e713b677b04a"),
		// 			SystemData: &armsecurityinsights.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armsecurityinsights.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armsecurityinsights.CreatedByTypeUser),
		// 			},
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Properties: &armsecurityinsights.SourceControlProperties{
		// 				Description: to.Ptr("this is a source control"),
		// 				ContentTypes: []*armsecurityinsights.ContentType{
		// 					to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
		// 					to.Ptr(armsecurityinsights.ContentTypeWorkbook)},
		// 					DisplayName: to.Ptr("My Source Control"),
		// 					ID: to.Ptr("789e0c1f-4a3d-43ad-809c-e713b677b04a"),
		// 					LastDeploymentInfo: &armsecurityinsights.DeploymentInfo{
		// 						Deployment: &armsecurityinsights.Deployment{
		// 							DeploymentID: to.Ptr("4985046420"),
		// 							DeploymentLogsURL: to.Ptr("https://github.com/user/repo/actions"),
		// 							DeploymentResult: to.Ptr(armsecurityinsights.DeploymentResultSuccess),
		// 							DeploymentState: to.Ptr(armsecurityinsights.DeploymentStateCompleted),
		// 							DeploymentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
		// 						},
		// 						DeploymentFetchStatus: to.Ptr(armsecurityinsights.DeploymentFetchStatusSuccess),
		// 						Message: to.Ptr("Successful deployment"),
		// 					},
		// 					RepoType: to.Ptr(armsecurityinsights.RepoTypeGithub),
		// 					Repository: &armsecurityinsights.Repository{
		// 						Branch: to.Ptr("master"),
		// 						DeploymentLogsURL: to.Ptr("https://github.com/user/repo/actions"),
		// 						DisplayURL: to.Ptr("https://github.com/user/repo"),
		// 						PathMapping: []*armsecurityinsights.ContentPathMap{
		// 							{
		// 								Path: to.Ptr("path/to/rules"),
		// 								ContentType: to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
		// 							},
		// 							{
		// 								Path: to.Ptr("path/to/workbooks"),
		// 								ContentType: to.Ptr(armsecurityinsights.ContentTypeWorkbook),
		// 						}},
		// 						URL: to.Ptr("https://github.com/user/repo"),
		// 					},
		// 					RepositoryResourceInfo: &armsecurityinsights.RepositoryResourceInfo{
		// 						GitHubResourceInfo: &armsecurityinsights.GitHubResourceInfo{
		// 							AppInstallationID: to.Ptr("123"),
		// 						},
		// 						Webhook: &armsecurityinsights.Webhook{
		// 							WebhookID: to.Ptr("342768323"),
		// 							WebhookSecretUpdateTime: to.Ptr("2021-01-01T17:18:19.1234567Z"),
		// 							WebhookURL: to.Ptr("https://cac.sentinel.azure.com/workspaces/b7c525e9-1bfa-4435-88c0-817e13abb088/webhooks/ado/sourceControl/789e0c1f-4a3d-43ad-809c-e713b677b04a"),
		// 						},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/sourcecontrols/GetSourceControlById.json
func ExampleSourceControlsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSourceControlsClient().Get(ctx, "myRg", "myWorkspace", "789e0c1f-4a3d-43ad-809c-e713b677b04a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SourceControl = armsecurityinsights.SourceControl{
	// 	Name: to.Ptr("789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/SourceControls"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/sourcecontrols/789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 	SystemData: &armsecurityinsights.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armsecurityinsights.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armsecurityinsights.CreatedByTypeUser),
	// 	},
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.SourceControlProperties{
	// 		Description: to.Ptr("this is a source control"),
	// 		ContentTypes: []*armsecurityinsights.ContentType{
	// 			to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
	// 			to.Ptr(armsecurityinsights.ContentTypeWorkbook)},
	// 			DisplayName: to.Ptr("My Source Control"),
	// 			ID: to.Ptr("789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 			LastDeploymentInfo: &armsecurityinsights.DeploymentInfo{
	// 				Deployment: &armsecurityinsights.Deployment{
	// 					DeploymentID: to.Ptr("4985046420"),
	// 					DeploymentLogsURL: to.Ptr("https://github.com/user/repo/actions"),
	// 					DeploymentResult: to.Ptr(armsecurityinsights.DeploymentResultSuccess),
	// 					DeploymentState: to.Ptr(armsecurityinsights.DeploymentStateCompleted),
	// 					DeploymentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 				},
	// 				DeploymentFetchStatus: to.Ptr(armsecurityinsights.DeploymentFetchStatusSuccess),
	// 				Message: to.Ptr("Successful deployment"),
	// 			},
	// 			RepoType: to.Ptr(armsecurityinsights.RepoTypeGithub),
	// 			Repository: &armsecurityinsights.Repository{
	// 				Branch: to.Ptr("master"),
	// 				DeploymentLogsURL: to.Ptr("https://github.com/user/repo/actions"),
	// 				DisplayURL: to.Ptr("https://github.com/user/repo"),
	// 				PathMapping: []*armsecurityinsights.ContentPathMap{
	// 					{
	// 						Path: to.Ptr("path/to/rules"),
	// 						ContentType: to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
	// 					},
	// 					{
	// 						Path: to.Ptr("path/to/workbooks"),
	// 						ContentType: to.Ptr(armsecurityinsights.ContentTypeWorkbook),
	// 				}},
	// 				URL: to.Ptr("https://github.com/user/repo"),
	// 			},
	// 			RepositoryResourceInfo: &armsecurityinsights.RepositoryResourceInfo{
	// 				GitHubResourceInfo: &armsecurityinsights.GitHubResourceInfo{
	// 					AppInstallationID: to.Ptr("123"),
	// 				},
	// 				Webhook: &armsecurityinsights.Webhook{
	// 					WebhookID: to.Ptr("342768323"),
	// 					WebhookSecretUpdateTime: to.Ptr("2021-01-01T17:18:19.1234567Z"),
	// 					WebhookURL: to.Ptr("https://cac.sentinel.azure.com/workspaces/b7c525e9-1bfa-4435-88c0-817e13abb088/webhooks/ado/sourceControl/789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 				},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/sourcecontrols/DeleteSourceControl.json
func ExampleSourceControlsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSourceControlsClient().Delete(ctx, "myRg", "myWorkspace", "789e0c1f-4a3d-43ad-809c-e713b677b04a", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/sourcecontrols/CreateSourceControl.json
func ExampleSourceControlsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSourceControlsClient().Create(ctx, "myRg", "myWorkspace", "789e0c1f-4a3d-43ad-809c-e713b677b04a", armsecurityinsights.SourceControl{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Properties: &armsecurityinsights.SourceControlProperties{
			Description: to.Ptr("This is a source control"),
			ContentTypes: []*armsecurityinsights.ContentType{
				to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
				to.Ptr(armsecurityinsights.ContentTypeWorkbook)},
			DisplayName: to.Ptr("My Source Control"),
			RepoType:    to.Ptr(armsecurityinsights.RepoTypeGithub),
			Repository: &armsecurityinsights.Repository{
				Branch:     to.Ptr("master"),
				DisplayURL: to.Ptr("https://github.com/user/repo"),
				PathMapping: []*armsecurityinsights.ContentPathMap{
					{
						Path:        to.Ptr("path/to/rules"),
						ContentType: to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
					},
					{
						Path:        to.Ptr("path/to/workbooks"),
						ContentType: to.Ptr(armsecurityinsights.ContentTypeWorkbook),
					}},
				URL: to.Ptr("https://github.com/user/repo"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SourceControl = armsecurityinsights.SourceControl{
	// 	Name: to.Ptr("789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/SourceControls"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalIinsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/sourcecontrols/789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 	SystemData: &armsecurityinsights.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armsecurityinsights.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armsecurityinsights.CreatedByTypeUser),
	// 	},
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.SourceControlProperties{
	// 		Description: to.Ptr("this is a source control"),
	// 		ContentTypes: []*armsecurityinsights.ContentType{
	// 			to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
	// 			to.Ptr(armsecurityinsights.ContentTypeWorkbook)},
	// 			DisplayName: to.Ptr("My Source Control"),
	// 			ID: to.Ptr("789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 			LastDeploymentInfo: &armsecurityinsights.DeploymentInfo{
	// 				Deployment: &armsecurityinsights.Deployment{
	// 					DeploymentID: to.Ptr("4985046420"),
	// 					DeploymentLogsURL: to.Ptr("https://github.com/user/repo/actions"),
	// 					DeploymentResult: to.Ptr(armsecurityinsights.DeploymentResultSuccess),
	// 					DeploymentState: to.Ptr(armsecurityinsights.DeploymentStateCompleted),
	// 					DeploymentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 				},
	// 				DeploymentFetchStatus: to.Ptr(armsecurityinsights.DeploymentFetchStatusSuccess),
	// 				Message: to.Ptr("Successful deployment"),
	// 			},
	// 			RepoType: to.Ptr(armsecurityinsights.RepoTypeGithub),
	// 			Repository: &armsecurityinsights.Repository{
	// 				Branch: to.Ptr("master"),
	// 				DeploymentLogsURL: to.Ptr("https://github.com/user/repo/actions"),
	// 				DisplayURL: to.Ptr("https://github.com/user/repo"),
	// 				PathMapping: []*armsecurityinsights.ContentPathMap{
	// 					{
	// 						Path: to.Ptr("path/to/rules"),
	// 						ContentType: to.Ptr(armsecurityinsights.ContentType("AnalyticRules")),
	// 					},
	// 					{
	// 						Path: to.Ptr("path/to/workbooks"),
	// 						ContentType: to.Ptr(armsecurityinsights.ContentTypeWorkbook),
	// 				}},
	// 				URL: to.Ptr("https://github.com/user/repo"),
	// 			},
	// 			RepositoryResourceInfo: &armsecurityinsights.RepositoryResourceInfo{
	// 				GitHubResourceInfo: &armsecurityinsights.GitHubResourceInfo{
	// 					AppInstallationID: to.Ptr("123"),
	// 				},
	// 				Webhook: &armsecurityinsights.Webhook{
	// 					WebhookID: to.Ptr("342768323"),
	// 					WebhookSecretUpdateTime: to.Ptr("2021-01-01T17:18:19.1234567Z"),
	// 					WebhookURL: to.Ptr("https://cac.sentinel.azure.com/workspaces/b7c525e9-1bfa-4435-88c0-817e13abb088/webhooks/ado/sourceControl/789e0c1f-4a3d-43ad-809c-e713b677b04a"),
	// 				},
	// 			},
	// 		},
	// 	}
}
