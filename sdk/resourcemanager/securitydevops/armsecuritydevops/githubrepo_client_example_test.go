//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecuritydevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoListByConnector.json
func ExampleGitHubRepoClient_NewListByConnectorPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGitHubRepoClient().NewListByConnectorPager("westusrg", "testconnector", nil)
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
		// page.GitHubRepoListResponse = armsecuritydevops.GitHubRepoListResponse{
		// 	Value: []*armsecuritydevops.GitHubRepo{
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/githubconnectors/owners/repos"),
		// 			ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector/owners/Azure/repos/azure-rest-api-specs"),
		// 			Properties: &armsecuritydevops.GitHubRepoProperties{
		// 				AccountID: to.Ptr[int64](6844498),
		// 				RepoURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("microsoft.securitydevops/githubconnectors/owners/repos"),
		// 			ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector/owners/Azure-Samples/repos/another-repo"),
		// 			Properties: &armsecuritydevops.GitHubRepoProperties{
		// 				AccountID: to.Ptr[int64](6844499),
		// 				RepoURL: to.Ptr("https://github.com/Azure-Samples/another-repo"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoList.json
func ExampleGitHubRepoClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGitHubRepoClient().NewListPager("westusrg", "testconnector", "Azure", nil)
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
		// page.GitHubRepoListResponse = armsecuritydevops.GitHubRepoListResponse{
		// 	Value: []*armsecuritydevops.GitHubRepo{
		// 		{
		// 			Name: to.Ptr("azure-rest-api-specs"),
		// 			Type: to.Ptr("microsoft.securitydevops/githubconnectors/owners/repos"),
		// 			ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector/owners/Azure/repos/azure-rest-api-specs"),
		// 			Properties: &armsecuritydevops.GitHubRepoProperties{
		// 				AccountID: to.Ptr[int64](6844498),
		// 				RepoURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("azure-rest-api-specs"),
		// 			Type: to.Ptr("microsoft.securitydevops/githubconnectors/owners/repos"),
		// 			ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector/owners/Azure/repos/another-repo"),
		// 			Properties: &armsecuritydevops.GitHubRepoProperties{
		// 				AccountID: to.Ptr[int64](6844499),
		// 				RepoURL: to.Ptr("https://github.com/Azure/another-repo"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoGet.json
func ExampleGitHubRepoClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGitHubRepoClient().Get(ctx, "westusrg", "testconnector", "Azure", "39093389", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GitHubRepo = armsecuritydevops.GitHubRepo{
	// 	Name: to.Ptr("azure-rest-api-specs"),
	// 	Type: to.Ptr("Microsoft.SecurityDevOps/owners/repos"),
	// 	ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector/owners/Azure/repos/azure-rest-api-specs"),
	// 	Properties: &armsecuritydevops.GitHubRepoProperties{
	// 		AccountID: to.Ptr[int64](6844498),
	// 		ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 		RepoURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoCreateOrUpdate.json
func ExampleGitHubRepoClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitHubRepoClient().BeginCreateOrUpdate(ctx, "westusrg", "testconnector", "Azure", "azure-rest-api-specs", armsecuritydevops.GitHubRepo{}, nil)
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
	// res.GitHubRepo = armsecuritydevops.GitHubRepo{
	// 	Name: to.Ptr("azure-rest-api-specs"),
	// 	Type: to.Ptr("Microsoft.SecurityDevOps/owners/repos"),
	// 	ID: to.Ptr("/subscriptions/e7032cc6-7422-4ddd-9022-bfbf23b05332/resourceGroups/westusrg/providers/Microsoft.SecurityDevOps/gitHubConnectors/testconnector/owners/myOrg/repos/azure-rest-api-specs"),
	// 	Properties: &armsecuritydevops.GitHubRepoProperties{
	// 		AccountID: to.Ptr[int64](6844498),
	// 		OwnerName: to.Ptr("Azure"),
	// 		ProvisioningState: to.Ptr(armsecuritydevops.ProvisioningStateSucceeded),
	// 		RepoURL: to.Ptr("https://github.com/Azure/azure-rest-api-specs"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoUpdate.json
func ExampleGitHubRepoClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecuritydevops.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGitHubRepoClient().BeginUpdate(ctx, "westusrg", "testconnector", "Azure", "azure-rest-api-specs", armsecuritydevops.GitHubRepo{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
