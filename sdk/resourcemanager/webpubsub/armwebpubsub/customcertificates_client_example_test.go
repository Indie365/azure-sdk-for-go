//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSubCustomCertificates_List.json
func ExampleCustomCertificatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomCertificatesClient().NewListPager("myResourceGroup", "myWebPubSubService", nil)
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
		// page.CustomCertificateList = armwebpubsub.CustomCertificateList{
		// 	Value: []*armwebpubsub.CustomCertificate{
		// 		{
		// 			Name: to.Ptr("myCert"),
		// 			Type: to.Ptr("Microsoft.SignalRService/WebPubSub/customCertificates"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/customCertificates/myCert"),
		// 			Properties: &armwebpubsub.CustomCertificateProperties{
		// 				KeyVaultBaseURI: to.Ptr("https://myvault.keyvault.azure.net/"),
		// 				KeyVaultSecretName: to.Ptr("mycert"),
		// 				KeyVaultSecretVersion: to.Ptr("bb6a44b2743f47f68dad0d6cc9756432"),
		// 				ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSubCustomCertificates_Get.json
func ExampleCustomCertificatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomCertificatesClient().Get(ctx, "myResourceGroup", "myWebPubSubService", "myCert", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomCertificate = armwebpubsub.CustomCertificate{
	// 	Name: to.Ptr("myCert"),
	// 	Type: to.Ptr("Microsoft.SignalRService/WebPubSub/customCertificates"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/customCertificates/myCert"),
	// 	Properties: &armwebpubsub.CustomCertificateProperties{
	// 		KeyVaultBaseURI: to.Ptr("https://myvault.keyvault.azure.net/"),
	// 		KeyVaultSecretName: to.Ptr("mycert"),
	// 		KeyVaultSecretVersion: to.Ptr("bb6a44b2743f47f68dad0d6cc9756432"),
	// 		ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSubCustomCertificates_CreateOrUpdate.json
func ExampleCustomCertificatesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomCertificatesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myWebPubSubService", "myCert", armwebpubsub.CustomCertificate{
		Properties: &armwebpubsub.CustomCertificateProperties{
			KeyVaultBaseURI:       to.Ptr("https://myvault.keyvault.azure.net/"),
			KeyVaultSecretName:    to.Ptr("mycert"),
			KeyVaultSecretVersion: to.Ptr("bb6a44b2743f47f68dad0d6cc9756432"),
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
	// res.CustomCertificate = armwebpubsub.CustomCertificate{
	// 	Name: to.Ptr("myCert"),
	// 	Type: to.Ptr("Microsoft.SignalRService/WebPubSub/customCertificates"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/Microsoft.SignalRService/WebPubSub/myWebPubSubService/customCertificates/myCert"),
	// 	Properties: &armwebpubsub.CustomCertificateProperties{
	// 		KeyVaultBaseURI: to.Ptr("https://myvault.keyvault.azure.net/"),
	// 		KeyVaultSecretName: to.Ptr("mycert"),
	// 		KeyVaultSecretVersion: to.Ptr("bb6a44b2743f47f68dad0d6cc9756432"),
	// 		ProvisioningState: to.Ptr(armwebpubsub.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2023-02-01/examples/WebPubSubCustomCertificates_Delete.json
func ExampleCustomCertificatesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armwebpubsub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewCustomCertificatesClient().Delete(ctx, "myResourceGroup", "myWebPubSubService", "myCert", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
