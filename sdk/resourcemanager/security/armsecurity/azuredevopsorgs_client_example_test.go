//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListAvailableAzureDevOpsOrgs_example.json
func ExampleAzureDevOpsOrgsClient_ListAvailable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureDevOpsOrgsClient().ListAvailable(ctx, "myRg", "mySecurityConnectorName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsOrgListResponse = armsecurity.AzureDevOpsOrgListResponse{
	// 	Value: []*armsecurity.AzureDevOpsOrg{
	// 		{
	// 			Name: to.Ptr("myAzDevOpsOrg"),
	// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
	// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg"),
	// 			Properties: &armsecurity.AzureDevOpsOrgProperties{
	// 				ActionableRemediation: &armsecurity.ActionableRemediation{
	// 					State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 				},
	// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 				ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("anotherOrg"),
	// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
	// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/anotherOrg"),
	// 			Properties: &armsecurity.AzureDevOpsOrgProperties{
	// 				ActionableRemediation: &armsecurity.ActionableRemediation{
	// 					State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 				},
	// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboardedByOtherConnector),
	// 				ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("notOnboardedOrg"),
	// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
	// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/notOnboardedOrg"),
	// 			Properties: &armsecurity.AzureDevOpsOrgProperties{
	// 				ActionableRemediation: &armsecurity.ActionableRemediation{
	// 					State: to.Ptr(armsecurity.ActionableRemediationStateNone),
	// 				},
	// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateNotOnboarded),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/ListAzureDevOpsOrgs_example.json
func ExampleAzureDevOpsOrgsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureDevOpsOrgsClient().NewListPager("myRg", "mySecurityConnectorName", nil)
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
		// page.AzureDevOpsOrgListResponse = armsecurity.AzureDevOpsOrgListResponse{
		// 	Value: []*armsecurity.AzureDevOpsOrg{
		// 		{
		// 			Name: to.Ptr("myAzDevOpsOrg"),
		// 			Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
		// 			ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg"),
		// 			Properties: &armsecurity.AzureDevOpsOrgProperties{
		// 				ActionableRemediation: &armsecurity.ActionableRemediation{
		// 					State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
		// 				},
		// 				OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
		// 				ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetAzureDevOpsOrgs_example.json
func ExampleAzureDevOpsOrgsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAzureDevOpsOrgsClient().Get(ctx, "myRg", "mySecurityConnectorName", "myAzDevOpsOrg", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AzureDevOpsOrg = armsecurity.AzureDevOpsOrg{
	// 	Name: to.Ptr("myAzDevOpsOrg"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg"),
	// 	Properties: &armsecurity.AzureDevOpsOrgProperties{
	// 		ActionableRemediation: &armsecurity.ActionableRemediation{
	// 			State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 		},
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/CreateOrUpdateAzureDevOpsOrgs_example.json
func ExampleAzureDevOpsOrgsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsOrgsClient().BeginCreateOrUpdate(ctx, "myRg", "mySecurityConnectorName", "myAzDevOpsOrg", armsecurity.AzureDevOpsOrg{
		Properties: &armsecurity.AzureDevOpsOrgProperties{
			ActionableRemediation: &armsecurity.ActionableRemediation{
				State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
			},
			OnboardingState: to.Ptr(armsecurity.OnboardingStateNotApplicable),
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
	// res.AzureDevOpsOrg = armsecurity.AzureDevOpsOrg{
	// 	Name: to.Ptr("myAzDevOpsOrg"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg"),
	// 	Properties: &armsecurity.AzureDevOpsOrgProperties{
	// 		ActionableRemediation: &armsecurity.ActionableRemediation{
	// 			State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 		},
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b43974e07d3204c4b6f8396627f5430994a7f7c9/specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/UpdateAzureDevOpsOrgs_example.json
func ExampleAzureDevOpsOrgsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAzureDevOpsOrgsClient().BeginUpdate(ctx, "myRg", "mySecurityConnectorName", "myAzDevOpsOrg", armsecurity.AzureDevOpsOrg{
		Properties: &armsecurity.AzureDevOpsOrgProperties{
			ActionableRemediation: &armsecurity.ActionableRemediation{
				State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
			},
			OnboardingState: to.Ptr(armsecurity.OnboardingStateNotApplicable),
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
	// res.AzureDevOpsOrg = armsecurity.AzureDevOpsOrg{
	// 	Name: to.Ptr("myAzDevOpsOrg"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors/devops/azureDevOpsOrgs"),
	// 	ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default/azureDevOpsOrgs/myAzDevOpsOrg"),
	// 	Properties: &armsecurity.AzureDevOpsOrgProperties{
	// 		ActionableRemediation: &armsecurity.ActionableRemediation{
	// 			State: to.Ptr(armsecurity.ActionableRemediationStateEnabled),
	// 		},
	// 		OnboardingState: to.Ptr(armsecurity.OnboardingStateOnboarded),
	// 		ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 	},
	// }
}
