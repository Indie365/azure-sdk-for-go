//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListGateways.json
func ExampleGatewayClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewayClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.GatewayClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
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
		// page.GatewayCollection = armapimanagement.GatewayCollection{
		// 	Count: to.Ptr[int64](2),
		// 	Value: []*armapimanagement.GatewayContract{
		// 		{
		// 			Name: to.Ptr("a1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/gateways"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1"),
		// 			Properties: &armapimanagement.GatewayContractProperties{
		// 				Description: to.Ptr("my gateway 1"),
		// 				LocationData: &armapimanagement.ResourceLocationDataContract{
		// 					Name: to.Ptr("my location 1"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5a73933b8f27f7cc82a2d533"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/gateways"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw2"),
		// 			Properties: &armapimanagement.GatewayContractProperties{
		// 				Description: to.Ptr("my gateway 2"),
		// 				LocationData: &armapimanagement.ResourceLocationDataContract{
		// 					Name: to.Ptr("my location 2"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadGateway.json
func ExampleGatewayClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewayClient().GetEntityTag(ctx, "rg1", "apimService1", "mygateway", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetGateway.json
func ExampleGatewayClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().Get(ctx, "rg1", "apimService1", "gw1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayContract = armapimanagement.GatewayContract{
	// 	Name: to.Ptr("a1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/gateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1"),
	// 	Properties: &armapimanagement.GatewayContractProperties{
	// 		Description: to.Ptr("my gateway 1"),
	// 		LocationData: &armapimanagement.ResourceLocationDataContract{
	// 			Name: to.Ptr("my location"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGateway.json
func ExampleGatewayClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().CreateOrUpdate(ctx, "rg1", "apimService1", "gw1", armapimanagement.GatewayContract{
		Properties: &armapimanagement.GatewayContractProperties{
			Description: to.Ptr("my gateway 1"),
			LocationData: &armapimanagement.ResourceLocationDataContract{
				Name: to.Ptr("my location"),
			},
		},
	}, &armapimanagement.GatewayClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayContract = armapimanagement.GatewayContract{
	// 	Name: to.Ptr("a1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/gateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1"),
	// 	Properties: &armapimanagement.GatewayContractProperties{
	// 		Description: to.Ptr("my gateway 1"),
	// 		LocationData: &armapimanagement.ResourceLocationDataContract{
	// 			Name: to.Ptr("my location"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateGateway.json
func ExampleGatewayClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().Update(ctx, "rg1", "apimService1", "gw1", "*", armapimanagement.GatewayContract{
		Properties: &armapimanagement.GatewayContractProperties{
			Description: to.Ptr("my gateway 1"),
			LocationData: &armapimanagement.ResourceLocationDataContract{
				Name: to.Ptr("my location"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayContract = armapimanagement.GatewayContract{
	// 	Name: to.Ptr("a1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/gateways"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1"),
	// 	Properties: &armapimanagement.GatewayContractProperties{
	// 		Description: to.Ptr("my gateway 1"),
	// 		LocationData: &armapimanagement.ResourceLocationDataContract{
	// 			Name: to.Ptr("my location"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteGateway.json
func ExampleGatewayClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewayClient().Delete(ctx, "rg1", "apimService1", "gw1", "*", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayListKeys.json
func ExampleGatewayClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().ListKeys(ctx, "rg1", "apimService1", "gw1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayKeysContract = armapimanagement.GatewayKeysContract{
	// 	Primary: to.Ptr("primary_key_value"),
	// 	Secondary: to.Ptr("secondary_key_value"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayRegenerateKey.json
func ExampleGatewayClient_RegenerateKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewGatewayClient().RegenerateKey(ctx, "rg1", "apimService1", "gwId", armapimanagement.GatewayKeyRegenerationRequestContract{
		KeyType: to.Ptr(armapimanagement.KeyTypePrimary),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayGenerateToken.json
func ExampleGatewayClient_GenerateToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGatewayClient().GenerateToken(ctx, "rg1", "apimService1", "gw1", armapimanagement.GatewayTokenRequestContract{
		Expiry:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-21T00:44:24.2845269Z"); return t }()),
		KeyType: to.Ptr(armapimanagement.KeyTypePrimary),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GatewayTokenContract = armapimanagement.GatewayTokenContract{
	// 	Value: to.Ptr("gw1&201904210044&9A1GR1f5WIhFvFmzQG+xxxxxxxxxxx/kBeu87DWad3tkasUXuvPL+MgzlwUHyg=="),
	// }
}
