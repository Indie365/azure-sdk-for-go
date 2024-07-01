//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyDelete.json
func ExampleServiceEndpointPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceEndpointPoliciesClient().BeginDelete(ctx, "rg1", "serviceEndpointPolicy1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyGet.json
func ExampleServiceEndpointPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceEndpointPoliciesClient().Get(ctx, "rg1", "testServiceEndpointPolicy", &armnetwork.ServiceEndpointPoliciesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceEndpointPolicy = armnetwork.ServiceEndpointPolicy{
	// 	Name: to.Ptr("testServiceEndpointPolicy"),
	// 	Type: to.Ptr("Microsoft.Network/serviceEndpointPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
	// 		ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
	// 			{
	// 				Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
	// 				Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
	// 					Description: to.Ptr("Storage Service EndpointPolicy Definition"),
	// 					Service: to.Ptr("Microsoft.Storage"),
	// 					ServiceResources: []*string{
	// 						to.Ptr("/subscriptions/subid1"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
	// 					},
	// 			}},
	// 			Subnets: []*armnetwork.Subnet{
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyCreate.json
func ExampleServiceEndpointPoliciesClient_BeginCreateOrUpdate_createServiceEndpointPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceEndpointPoliciesClient().BeginCreateOrUpdate(ctx, "rg1", "testPolicy", armnetwork.ServiceEndpointPolicy{
		Location: to.Ptr("westus"),
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
	// res.ServiceEndpointPolicy = armnetwork.ServiceEndpointPolicy{
	// 	Name: to.Ptr("testnsg"),
	// 	Type: to.Ptr("Microsoft.Network/ServiceEndpointPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ServiceEndpointPolicies/testpolicy"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
	// 		ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
	// 		},
	// 		Subnets: []*armnetwork.Subnet{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyCreateWithDefinition.json
func ExampleServiceEndpointPoliciesClient_BeginCreateOrUpdate_createServiceEndpointPolicyWithDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceEndpointPoliciesClient().BeginCreateOrUpdate(ctx, "rg1", "testPolicy", armnetwork.ServiceEndpointPolicy{
		Location: to.Ptr("westus"),
		Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
			ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
				{
					Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
					Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
						Description: to.Ptr("Storage Service EndpointPolicy Definition"),
						Service:     to.Ptr("Microsoft.Storage"),
						ServiceResources: []*string{
							to.Ptr("/subscriptions/subid1"),
							to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
							to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
					},
				}},
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
	// res.ServiceEndpointPolicy = armnetwork.ServiceEndpointPolicy{
	// 	Name: to.Ptr("testnsg"),
	// 	Type: to.Ptr("Microsoft.Network/ServiceEndpointPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ServiceEndpointPolicies/testpolicy"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
	// 		ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
	// 			{
	// 				Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
	// 				Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
	// 					Description: to.Ptr("Storage Service EndpointPolicy Definition"),
	// 					Service: to.Ptr("Microsoft.Storage"),
	// 					ServiceResources: []*string{
	// 						to.Ptr("/subscriptions/subid1"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
	// 					},
	// 			}},
	// 			Subnets: []*armnetwork.Subnet{
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyUpdateTags.json
func ExampleServiceEndpointPoliciesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceEndpointPoliciesClient().UpdateTags(ctx, "rg1", "testServiceEndpointPolicy", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceEndpointPolicy = armnetwork.ServiceEndpointPolicy{
	// 	Name: to.Ptr("testServiceEndpointPolicy"),
	// 	Type: to.Ptr("Microsoft.Network/serviceEndpointPolicies"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
	// 			{
	// 				Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
	// 				Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
	// 					Description: to.Ptr("Storage Service EndpointPolicy Definition"),
	// 					Service: to.Ptr("Microsoft.Storage"),
	// 					ServiceResources: []*string{
	// 						to.Ptr("/subscriptions/subid1"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
	// 						to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
	// 					},
	// 			}},
	// 			Subnets: []*armnetwork.Subnet{
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyListAll.json
func ExampleServiceEndpointPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceEndpointPoliciesClient().NewListPager(nil)
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
		// page.ServiceEndpointPolicyListResult = armnetwork.ServiceEndpointPolicyListResult{
		// 	Value: []*armnetwork.ServiceEndpointPolicy{
		// 		{
		// 			Name: to.Ptr("testPolicy"),
		// 			Type: to.Ptr("Microsoft.Network/serviceEndpointPolicies"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testPolicy"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 				ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
		// 					{
		// 						Name: to.Ptr("StorageServiceEndpointPolicyDefinition1"),
		// 						Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
		// 							Description: to.Ptr("Storage Service EndpointPolicy Definition"),
		// 							Service: to.Ptr("Microsoft.Storage"),
		// 							ServiceResources: []*string{
		// 								to.Ptr("/subscriptions/subid1"),
		// 								to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
		// 								to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
		// 							},
		// 					}},
		// 					Subnets: []*armnetwork.Subnet{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testPolicy1"),
		// 				Type: to.Ptr("Microsoft.Network/serviceEndpointPolicies"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/serviceEndpointPolicies/testPolicy2"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 					ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
		// 						{
		// 							Name: to.Ptr("StorageServiceEndpointPolicyDefinition2"),
		// 							Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
		// 								Description: to.Ptr("Storage Service EndpointPolicy Definition"),
		// 								Service: to.Ptr("Microsoft.Storage"),
		// 								ServiceResources: []*string{
		// 									to.Ptr("/subscriptions/subid1"),
		// 									to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
		// 									to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
		// 								},
		// 						}},
		// 						Subnets: []*armnetwork.Subnet{
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/ServiceEndpointPolicyList.json
func ExampleServiceEndpointPoliciesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceEndpointPoliciesClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ServiceEndpointPolicyListResult = armnetwork.ServiceEndpointPolicyListResult{
		// 	Value: []*armnetwork.ServiceEndpointPolicy{
		// 		{
		// 			Name: to.Ptr("testServiceEndpointPolicy"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 				ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
		// 					{
		// 						Name: to.Ptr("StorageServiceEndpointPolicyDefinition"),
		// 						Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
		// 							Description: to.Ptr("Storage Service EndpointPolicy Definition"),
		// 							Service: to.Ptr("Microsoft.Storage"),
		// 							ServiceResources: []*string{
		// 								to.Ptr("/subscriptions/subid1"),
		// 								to.Ptr("/subscriptions/subid1resourceGroups/storageRg"),
		// 								to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
		// 							},
		// 					}},
		// 					Subnets: []*armnetwork.Subnet{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("testServiceEndpointPolicy1"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/serviceEndpointPolicies/testServiceEndpointPolicy1"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armnetwork.ServiceEndpointPolicyPropertiesFormat{
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					ResourceGUID: to.Ptr("6A7C139D-8B8D-499B-B7CB-4F3F02A8A44F"),
		// 					ServiceEndpointPolicyDefinitions: []*armnetwork.ServiceEndpointPolicyDefinition{
		// 						{
		// 							Name: to.Ptr("StorageServiceEndpointPolicyDefinition1"),
		// 							Properties: &armnetwork.ServiceEndpointPolicyDefinitionPropertiesFormat{
		// 								Description: to.Ptr("Storage Service EndpointPolicy Definition"),
		// 								Service: to.Ptr("Microsoft.Storage"),
		// 								ServiceResources: []*string{
		// 									to.Ptr("/subscriptions/subid1"),
		// 									to.Ptr("/subscriptions/subid1/resourceGroups/storageRg"),
		// 									to.Ptr("/subscriptions/subid1/resourceGroups/storageRg/providers/Microsoft.Storage/storageAccounts/stAccount")},
		// 								},
		// 						}},
		// 						Subnets: []*armnetwork.Subnet{
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
