//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcommunication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/emailServices/get.json
func ExampleEmailServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmailServicesClient().Get(ctx, "MyResourceGroup", "MyEmailServiceResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EmailServiceResource = armcommunication.EmailServiceResource{
	// 	Name: to.Ptr("MyEmailServiceResource"),
	// 	Type: to.Ptr("Microsoft.Communication/EmailServices"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource"),
	// 	Location: to.Ptr("Global"),
	// 	Properties: &armcommunication.EmailServiceProperties{
	// 		DataLocation: to.Ptr("United States"),
	// 		ProvisioningState: to.Ptr(armcommunication.EmailServicesProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/emailServices/createOrUpdate.json
func ExampleEmailServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEmailServicesClient().BeginCreateOrUpdate(ctx, "MyResourceGroup", "MyEmailServiceResource", armcommunication.EmailServiceResource{
		Location: to.Ptr("Global"),
		Properties: &armcommunication.EmailServiceProperties{
			DataLocation: to.Ptr("United States"),
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
	// res.EmailServiceResource = armcommunication.EmailServiceResource{
	// 	Name: to.Ptr("MyEmailServiceResource"),
	// 	Type: to.Ptr("Microsoft.Communication/EmailServices"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource"),
	// 	Location: to.Ptr("Global"),
	// 	Properties: &armcommunication.EmailServiceProperties{
	// 		DataLocation: to.Ptr("United States"),
	// 		ProvisioningState: to.Ptr(armcommunication.EmailServicesProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/emailServices/delete.json
func ExampleEmailServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEmailServicesClient().BeginDelete(ctx, "MyResourceGroup", "MyEmailServiceResource", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/emailServices/update.json
func ExampleEmailServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEmailServicesClient().BeginUpdate(ctx, "MyResourceGroup", "MyEmailServiceResource", armcommunication.EmailServiceResourceUpdate{
		Tags: map[string]*string{
			"newTag": to.Ptr("newVal"),
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
	// res.EmailServiceResource = armcommunication.EmailServiceResource{
	// 	Name: to.Ptr("MyEmailServiceResource"),
	// 	Type: to.Ptr("Microsoft.Communication/EmailServices"),
	// 	ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource"),
	// 	Location: to.Ptr("Global"),
	// 	Tags: map[string]*string{
	// 		"newTag": to.Ptr("newVal"),
	// 	},
	// 	Properties: &armcommunication.EmailServiceProperties{
	// 		DataLocation: to.Ptr("United States"),
	// 		ProvisioningState: to.Ptr(armcommunication.EmailServicesProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/emailServices/listBySubscription.json
func ExampleEmailServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmailServicesClient().NewListBySubscriptionPager(nil)
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
		// page.EmailServiceResourceList = armcommunication.EmailServiceResourceList{
		// 	Value: []*armcommunication.EmailServiceResource{
		// 		{
		// 			Name: to.Ptr("MyEmailServiceResource"),
		// 			Type: to.Ptr("Microsoft.Communication/EmailServices"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource"),
		// 			Location: to.Ptr("Global"),
		// 			Properties: &armcommunication.EmailServiceProperties{
		// 				DataLocation: to.Ptr("United States"),
		// 				ProvisioningState: to.Ptr(armcommunication.EmailServicesProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/emailServices/listByResourceGroup.json
func ExampleEmailServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEmailServicesClient().NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page.EmailServiceResourceList = armcommunication.EmailServiceResourceList{
		// 	Value: []*armcommunication.EmailServiceResource{
		// 		{
		// 			Name: to.Ptr("MyEmailServiceResource"),
		// 			Type: to.Ptr("Microsoft.Communication/EmailServices"),
		// 			ID: to.Ptr("/subscriptions/11112222-3333-4444-5555-666677778888/resourceGroups/MyResourceGroup/providers/Microsoft.Communication/EmailServices/MyEmailServiceResource"),
		// 			Location: to.Ptr("Global"),
		// 			Properties: &armcommunication.EmailServiceProperties{
		// 				DataLocation: to.Ptr("United States"),
		// 				ProvisioningState: to.Ptr(armcommunication.EmailServicesProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/06042d4ffed6a1ea8319c39cd07ea21efe5d49f7/specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/emailServices/getVerifiedExchangeOnlineDomains.json
func ExampleEmailServicesClient_ListVerifiedExchangeOnlineDomains() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcommunication.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEmailServicesClient().ListVerifiedExchangeOnlineDomains(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StringArray = []*string{
	// 	to.Ptr("example.com"),
	// 	to.Ptr("contoso.com")}
}
