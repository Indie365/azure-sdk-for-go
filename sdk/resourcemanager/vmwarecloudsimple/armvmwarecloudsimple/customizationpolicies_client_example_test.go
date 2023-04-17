//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListCustomizationPolicies.json
func ExampleCustomizationPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomizationPoliciesClient().NewListPager("myResourceGroup", "myPrivateCloud", &armvmwarecloudsimple.CustomizationPoliciesClientListOptions{Filter: nil})
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
		// page.CustomizationPoliciesListResponse = armvmwarecloudsimple.CustomizationPoliciesListResponse{
		// 	Value: []*armvmwarecloudsimple.CustomizationPolicy{
		// 		{
		// 			Name: to.Ptr("Linux1"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/customizationPolicies"),
		// 			ID: to.Ptr("/subscriptions/b85c4986-56ae-4ebd-b5c5-a1595ca3dab1/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/foobar/customizationpolicies/Linux1"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.CustomizationPolicyProperties{
		// 				Type: to.Ptr(armvmwarecloudsimple.CustomizationPolicyPropertiesTypeLINUX),
		// 				PrivateCloudID: to.Ptr("/subscriptions/b85c4986-56ae-4ebd-b5c5-a1595ca3dab1/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/foobar"),
		// 				Specification: &armvmwarecloudsimple.CustomizationSpecification{
		// 					Identity: &armvmwarecloudsimple.CustomizationIdentity{
		// 						Type: to.Ptr(armvmwarecloudsimple.CustomizationIdentityTypeLINUX),
		// 						HostName: &armvmwarecloudsimple.CustomizationHostName{
		// 							Type: to.Ptr(armvmwarecloudsimple.CustomizationHostNameTypeUSERDEFINED),
		// 						},
		// 						UserData: &armvmwarecloudsimple.CustomizationIdentityUserData{
		// 							IsPasswordPredefined: to.Ptr(false),
		// 						},
		// 					},
		// 					NicSettings: []*armvmwarecloudsimple.CustomizationNicSetting{
		// 						{
		// 							Adapter: &armvmwarecloudsimple.CustomizationIPSettings{
		// 								IP: &armvmwarecloudsimple.CustomizationIPAddress{
		// 									Type: to.Ptr(armvmwarecloudsimple.CustomizationIPAddressTypeUSERDEFINED),
		// 								},
		// 							},
		// 					}},
		// 				},
		// 				Version: to.Ptr("1568102823"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Linux2"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/customizationPolicies"),
		// 			ID: to.Ptr("/subscriptions/b85c4986-56ae-4ebd-b5c5-a1595ca3dab1/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/foobar/customizationpolicies/Linux2"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.CustomizationPolicyProperties{
		// 				Type: to.Ptr(armvmwarecloudsimple.CustomizationPolicyPropertiesTypeLINUX),
		// 				PrivateCloudID: to.Ptr("/subscriptions/b85c4986-56ae-4ebd-b5c5-a1595ca3dab1/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/foobar"),
		// 				Specification: &armvmwarecloudsimple.CustomizationSpecification{
		// 					Identity: &armvmwarecloudsimple.CustomizationIdentity{
		// 						Type: to.Ptr(armvmwarecloudsimple.CustomizationIdentityTypeLINUX),
		// 						HostName: &armvmwarecloudsimple.CustomizationHostName{
		// 							Type: to.Ptr(armvmwarecloudsimple.CustomizationHostNameTypeVIRTUALMACHINENAME),
		// 						},
		// 						UserData: &armvmwarecloudsimple.CustomizationIdentityUserData{
		// 							IsPasswordPredefined: to.Ptr(false),
		// 						},
		// 					},
		// 					NicSettings: []*armvmwarecloudsimple.CustomizationNicSetting{
		// 						{
		// 							Adapter: &armvmwarecloudsimple.CustomizationIPSettings{
		// 								IP: &armvmwarecloudsimple.CustomizationIPAddress{
		// 									Type: to.Ptr(armvmwarecloudsimple.CustomizationIPAddressTypeDHCPIP),
		// 								},
		// 							},
		// 					}},
		// 				},
		// 				Version: to.Ptr("1568102846"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetCustomizationPolicy.json
func ExampleCustomizationPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomizationPoliciesClient().Get(ctx, "myResourceGroup", "myPrivateCloud", "Linux1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomizationPolicy = armvmwarecloudsimple.CustomizationPolicy{
	// 	Name: to.Ptr("Linux1"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/customizationPolicies"),
	// 	ID: to.Ptr("/subscriptions/b85c4986-56ae-4ebd-b5c5-a1595ca3dab1/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/foobar/customizationpolicies/Linux1"),
	// 	Location: to.Ptr("westus2"),
	// 	Properties: &armvmwarecloudsimple.CustomizationPolicyProperties{
	// 		Type: to.Ptr(armvmwarecloudsimple.CustomizationPolicyPropertiesTypeLINUX),
	// 		PrivateCloudID: to.Ptr("/subscriptions/b85c4986-56ae-4ebd-b5c5-a1595ca3dab1/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/foobar"),
	// 		Specification: &armvmwarecloudsimple.CustomizationSpecification{
	// 			Identity: &armvmwarecloudsimple.CustomizationIdentity{
	// 				Type: to.Ptr(armvmwarecloudsimple.CustomizationIdentityTypeLINUX),
	// 				HostName: &armvmwarecloudsimple.CustomizationHostName{
	// 					Type: to.Ptr(armvmwarecloudsimple.CustomizationHostNameTypeUSERDEFINED),
	// 				},
	// 				UserData: &armvmwarecloudsimple.CustomizationIdentityUserData{
	// 					IsPasswordPredefined: to.Ptr(false),
	// 				},
	// 			},
	// 			NicSettings: []*armvmwarecloudsimple.CustomizationNicSetting{
	// 				{
	// 					Adapter: &armvmwarecloudsimple.CustomizationIPSettings{
	// 						IP: &armvmwarecloudsimple.CustomizationIPAddress{
	// 							Type: to.Ptr(armvmwarecloudsimple.CustomizationIPAddressTypeUSERDEFINED),
	// 						},
	// 					},
	// 			}},
	// 		},
	// 		Version: to.Ptr("1568102823"),
	// 	},
	// }
}
