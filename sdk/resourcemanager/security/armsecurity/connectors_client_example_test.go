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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/GetSecurityConnectorsSubscription_example.json
func ExampleConnectorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectorsClient().NewListPager(nil)
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
		// page.ConnectorsList = armsecurity.ConnectorsList{
		// 	Value: []*armsecurity.Connector{
		// 		{
		// 			Location: to.Ptr("Central US"),
		// 			Etag: to.Ptr("etag value"),
		// 			Kind: to.Ptr(""),
		// 			Name: to.Ptr("exampleSecurityConnectorAws"),
		// 			Type: to.Ptr("Microsoft.Security/securityConnectors"),
		// 			ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup1/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorAws"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armsecurity.ConnectorProperties{
		// 				EnvironmentData: &armsecurity.AwsEnvironmentData{
		// 					EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
		// 				},
		// 				EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
		// 				HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
		// 				Offerings: []armsecurity.CloudOfferingClassification{
		// 					&armsecurity.CspmMonitorAwsOffering{
		// 						OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
		// 						NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
		// 							CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
		// 						},
		// 				}},
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Location: to.Ptr("Central US"),
		// 			Etag: to.Ptr("etag value"),
		// 			Kind: to.Ptr(""),
		// 			Name: to.Ptr("exampleSecurityConnectorAwsOrganization"),
		// 			Type: to.Ptr("Microsoft.Security/securityConnectors"),
		// 			ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup2/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorAwsOrganization"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armsecurity.ConnectorProperties{
		// 				EnvironmentData: &armsecurity.AwsEnvironmentData{
		// 					EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
		// 					OrganizationalData: &armsecurity.AwsOrganizationalDataMaster{
		// 						OrganizationMembershipType: to.Ptr(armsecurity.OrganizationMembershipTypeOrganization),
		// 						ExcludedAccountIDs: []*string{
		// 							to.Ptr("excludedAccountIdExample")},
		// 							StacksetName: to.Ptr("myStackSetName"),
		// 						},
		// 					},
		// 					EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
		// 					HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
		// 					Offerings: []armsecurity.CloudOfferingClassification{
		// 						&armsecurity.CspmMonitorAwsOffering{
		// 							OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
		// 							NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
		// 								CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
		// 							},
		// 					}},
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Location: to.Ptr("centralus"),
		// 				Etag: to.Ptr("etag value"),
		// 				Kind: to.Ptr(""),
		// 				Name: to.Ptr("githubTest"),
		// 				Type: to.Ptr("Microsoft.Security/securityconnectors"),
		// 				ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup3/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorGithub"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armsecurity.ConnectorProperties{
		// 					EnvironmentData: &armsecurity.GithubScopeEnvironmentData{
		// 						EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeGithubScope),
		// 					},
		// 					EnvironmentName: to.Ptr(armsecurity.CloudNameGithub),
		// 					HierarchyIdentifier: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup3/providers/Microsoft.SecurityDevops/githubConnectors/exampleGithubConnector"),
		// 					Offerings: []armsecurity.CloudOfferingClassification{
		// 						&armsecurity.CspmMonitorGithubOffering{
		// 							OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorGithub),
		// 					}},
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Location: to.Ptr("centralus"),
		// 				Etag: to.Ptr("etag value"),
		// 				Kind: to.Ptr(""),
		// 				Name: to.Ptr("AzureDevOpsTest"),
		// 				Type: to.Ptr("Microsoft.Security/securityconnectors"),
		// 				ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup3/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorAzureDevOpsConnectors"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armsecurity.ConnectorProperties{
		// 					EnvironmentData: &armsecurity.GithubScopeEnvironmentData{
		// 						EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeGithubScope),
		// 					},
		// 					EnvironmentName: to.Ptr(armsecurity.CloudNameGithub),
		// 					HierarchyIdentifier: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup3/providers/Microsoft.SecurityDevops/azureDevOpsConnectors/exampleAzureDevOpsConnector"),
		// 					Offerings: []armsecurity.CloudOfferingClassification{
		// 						&armsecurity.CspmMonitorGithubOffering{
		// 							OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorGithub),
		// 					}},
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Location: to.Ptr("Central US"),
		// 				Etag: to.Ptr("etag value"),
		// 				Kind: to.Ptr(""),
		// 				Name: to.Ptr("exampleSecurityConnectorGcp"),
		// 				Type: to.Ptr("Microsoft.Security/securityConnectors"),
		// 				ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup3/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorGcp"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armsecurity.ConnectorProperties{
		// 					EnvironmentData: &armsecurity.GcpProjectEnvironmentData{
		// 						EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeGcpProject),
		// 						ProjectDetails: &armsecurity.GcpProjectDetails{
		// 							ProjectID: to.Ptr("My-0GCP-Project"),
		// 							ProjectNumber: to.Ptr("exampleHierarchyId"),
		// 							WorkloadIdentityPoolID: to.Ptr("6c78da41157548d3b1d8b3c72effdf8c"),
		// 						},
		// 					},
		// 					EnvironmentName: to.Ptr(armsecurity.CloudNameGCP),
		// 					HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
		// 					Offerings: []armsecurity.CloudOfferingClassification{
		// 						&armsecurity.CspmMonitorGcpOffering{
		// 							OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorGcp),
		// 							NativeCloudConnection: &armsecurity.CspmMonitorGcpOfferingNativeCloudConnection{
		// 								ServiceAccountEmailAddress: to.Ptr("capm@projectName.com"),
		// 								WorkloadIdentityProviderID: to.Ptr("My workload identity provider Id"),
		// 							},
		// 					}},
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/GetSecurityConnectorsResourceGroup_example.json
func ExampleConnectorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectorsClient().NewListByResourceGroupPager("exampleResourceGroup", nil)
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
		// page.ConnectorsList = armsecurity.ConnectorsList{
		// 	Value: []*armsecurity.Connector{
		// 		{
		// 			Location: to.Ptr("Central US"),
		// 			Etag: to.Ptr("etag value"),
		// 			Kind: to.Ptr(""),
		// 			Name: to.Ptr("exampleSecurityConnectorAws"),
		// 			Type: to.Ptr("Microsoft.Security/securityConnectors"),
		// 			ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorAws"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armsecurity.ConnectorProperties{
		// 				EnvironmentData: &armsecurity.AwsEnvironmentData{
		// 					EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
		// 				},
		// 				EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
		// 				HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
		// 				Offerings: []armsecurity.CloudOfferingClassification{
		// 					&armsecurity.CspmMonitorAwsOffering{
		// 						OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
		// 						NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
		// 							CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
		// 						},
		// 				}},
		// 			},
		// 			SystemData: &armsecurity.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@contoso.com"),
		// 				CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Location: to.Ptr("Central US"),
		// 			Etag: to.Ptr("etag value"),
		// 			Kind: to.Ptr(""),
		// 			Name: to.Ptr("exampleSecurityConnectorAwsOrganization"),
		// 			Type: to.Ptr("Microsoft.Security/securityConnectors"),
		// 			ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorAwsOrganization"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armsecurity.ConnectorProperties{
		// 				EnvironmentData: &armsecurity.AwsEnvironmentData{
		// 					EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
		// 					OrganizationalData: &armsecurity.AwsOrganizationalDataMaster{
		// 						OrganizationMembershipType: to.Ptr(armsecurity.OrganizationMembershipTypeOrganization),
		// 						ExcludedAccountIDs: []*string{
		// 							to.Ptr("excludedAccountIdExample")},
		// 							StacksetName: to.Ptr("myStackSetName"),
		// 						},
		// 					},
		// 					EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
		// 					HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
		// 					Offerings: []armsecurity.CloudOfferingClassification{
		// 						&armsecurity.CspmMonitorAwsOffering{
		// 							OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
		// 							NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
		// 								CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
		// 							},
		// 					}},
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Location: to.Ptr("Central US"),
		// 				Etag: to.Ptr("etag value"),
		// 				Kind: to.Ptr(""),
		// 				Name: to.Ptr("exampleSecurityConnectorGcp"),
		// 				Type: to.Ptr("Microsoft.Security/securityConnectors"),
		// 				ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorGcp"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armsecurity.ConnectorProperties{
		// 					EnvironmentData: &armsecurity.GcpProjectEnvironmentData{
		// 						EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeGcpProject),
		// 						ProjectDetails: &armsecurity.GcpProjectDetails{
		// 							ProjectID: to.Ptr("My-0GCP-Project"),
		// 							ProjectNumber: to.Ptr("exampleHierarchyId"),
		// 							WorkloadIdentityPoolID: to.Ptr("My-workload-identity-federation-pool-id"),
		// 						},
		// 					},
		// 					EnvironmentName: to.Ptr(armsecurity.CloudNameGCP),
		// 					HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
		// 					Offerings: []armsecurity.CloudOfferingClassification{
		// 						&armsecurity.CspmMonitorGcpOffering{
		// 							OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorGcp),
		// 							NativeCloudConnection: &armsecurity.CspmMonitorGcpOfferingNativeCloudConnection{
		// 								ServiceAccountEmailAddress: to.Ptr("capm@projectName.com"),
		// 								WorkloadIdentityProviderID: to.Ptr("My workload identity provider Id"),
		// 							},
		// 					}},
		// 				},
		// 				SystemData: &armsecurity.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					CreatedBy: to.Ptr("user@contoso.com"),
		// 					CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-15T13:47:50.328Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/GetSecurityConnectorSingleResource_example.json
func ExampleConnectorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectorsClient().Get(ctx, "exampleResourceGroup", "exampleSecurityConnectorName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connector = armsecurity.Connector{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("etag value"),
	// 	Kind: to.Ptr(""),
	// 	Name: to.Ptr("exampleSecurityConnectorName"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorName"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.ConnectorProperties{
	// 		EnvironmentData: &armsecurity.AwsEnvironmentData{
	// 			EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
	// 		},
	// 		EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
	// 		HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
	// 		Offerings: []armsecurity.CloudOfferingClassification{
	// 			&armsecurity.CspmMonitorAwsOffering{
	// 				OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
	// 				NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
	// 					CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
	// 				},
	// 		}},
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/PutSecurityConnector_example.json
func ExampleConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectorsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleSecurityConnectorName", armsecurity.Connector{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.ConnectorProperties{
			EnvironmentData: &armsecurity.AwsEnvironmentData{
				EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
			},
			EnvironmentName:     to.Ptr(armsecurity.CloudNameAWS),
			HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
			Offerings: []armsecurity.CloudOfferingClassification{
				&armsecurity.CspmMonitorAwsOffering{
					OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
					NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
						CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connector = armsecurity.Connector{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr(""),
	// 	Kind: to.Ptr(""),
	// 	Name: to.Ptr("exampleSecurityConnectorName"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorName"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.ConnectorProperties{
	// 		EnvironmentData: &armsecurity.AwsEnvironmentData{
	// 			EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
	// 		},
	// 		EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
	// 		HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
	// 		Offerings: []armsecurity.CloudOfferingClassification{
	// 			&armsecurity.CspmMonitorAwsOffering{
	// 				OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
	// 				NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
	// 					CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
	// 				},
	// 		}},
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/PatchSecurityConnector_example.json
func ExampleConnectorsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectorsClient().Update(ctx, "exampleResourceGroup", "exampleSecurityConnectorName", armsecurity.Connector{
		Location: to.Ptr("Central US"),
		Etag:     to.Ptr("etag value (must be supplied for update)"),
		Tags:     map[string]*string{},
		Properties: &armsecurity.ConnectorProperties{
			EnvironmentData: &armsecurity.AwsEnvironmentData{
				EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
			},
			EnvironmentName:     to.Ptr(armsecurity.CloudNameAWS),
			HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
			Offerings: []armsecurity.CloudOfferingClassification{
				&armsecurity.CspmMonitorAwsOffering{
					OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
					NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
						CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Connector = armsecurity.Connector{
	// 	Location: to.Ptr("Central US"),
	// 	Etag: to.Ptr("etag value"),
	// 	Kind: to.Ptr(""),
	// 	Name: to.Ptr("exampleSecurityConnectorName"),
	// 	Type: to.Ptr("Microsoft.Security/securityConnectors"),
	// 	ID: to.Ptr("/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/exampleResourceGroup/providers/Microsoft.Security/securityConnectors/exampleSecurityConnectorName"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armsecurity.ConnectorProperties{
	// 		EnvironmentData: &armsecurity.AwsEnvironmentData{
	// 			EnvironmentType: to.Ptr(armsecurity.EnvironmentTypeAwsAccount),
	// 		},
	// 		EnvironmentName: to.Ptr(armsecurity.CloudNameAWS),
	// 		HierarchyIdentifier: to.Ptr("exampleHierarchyId"),
	// 		Offerings: []armsecurity.CloudOfferingClassification{
	// 			&armsecurity.CspmMonitorAwsOffering{
	// 				OfferingType: to.Ptr(armsecurity.OfferingTypeCspmMonitorAws),
	// 				NativeCloudConnection: &armsecurity.CspmMonitorAwsOfferingNativeCloudConnection{
	// 					CloudRoleArn: to.Ptr("arn:aws:iam::00000000:role/ASCMonitor"),
	// 				},
	// 		}},
	// 	},
	// 	SystemData: &armsecurity.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@contoso.com"),
	// 		CreatedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-31T13:47:50.328Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armsecurity.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/DeleteSecurityConnector_example.json
func ExampleConnectorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectorsClient().Delete(ctx, "myRg", "mySecurityConnectorName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
