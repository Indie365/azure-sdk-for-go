//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfrontdoor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorFrontendEndpointList.json
func ExampleFrontendEndpointsClient_NewListByFrontDoorPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFrontendEndpointsClient().NewListByFrontDoorPager("rg1", "frontDoor1", nil)
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
		// page.FrontendEndpointsListResult = armfrontdoor.FrontendEndpointsListResult{
		// 	Value: []*armfrontdoor.FrontendEndpoint{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/frontendEndpoint1"),
		// 			Name: to.Ptr("frontendEndpoint1"),
		// 			Properties: &armfrontdoor.FrontendEndpointProperties{
		// 				HostName: to.Ptr("www.contoso.com"),
		// 				SessionAffinityEnabledState: to.Ptr(armfrontdoor.SessionAffinityEnabledStateEnabled),
		// 				SessionAffinityTTLSeconds: to.Ptr[int32](60),
		// 				WebApplicationFirewallPolicyLink: &armfrontdoor.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorFrontendEndpointGet.json
func ExampleFrontendEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFrontendEndpointsClient().Get(ctx, "rg1", "frontDoor1", "frontendEndpoint1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.FrontendEndpoint = armfrontdoor.FrontendEndpoint{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoors/frontDoor1/frontendEndpoints/frontendEndpoint1"),
	// 	Name: to.Ptr("frontendEndpoint1"),
	// 	Properties: &armfrontdoor.FrontendEndpointProperties{
	// 		HostName: to.Ptr("www.contoso.com"),
	// 		SessionAffinityEnabledState: to.Ptr(armfrontdoor.SessionAffinityEnabledStateEnabled),
	// 		SessionAffinityTTLSeconds: to.Ptr[int32](60),
	// 		WebApplicationFirewallPolicyLink: &armfrontdoor.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies/policy1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorEnableHttps.json
func ExampleFrontendEndpointsClient_BeginEnableHTTPS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFrontendEndpointsClient().BeginEnableHTTPS(ctx, "rg1", "frontDoor1", "frontendEndpoint1", armfrontdoor.CustomHTTPSConfiguration{
		CertificateSource: to.Ptr(armfrontdoor.FrontDoorCertificateSourceAzureKeyVault),
		KeyVaultCertificateSourceParameters: &armfrontdoor.KeyVaultCertificateSourceParameters{
			SecretName:    to.Ptr("secret1"),
			SecretVersion: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Vault: &armfrontdoor.KeyVaultCertificateSourceParametersVault{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.KeyVault/vaults/vault1"),
			},
		},
		MinimumTLSVersion: to.Ptr(armfrontdoor.MinimumTLSVersionOne0),
		ProtocolType:      to.Ptr(armfrontdoor.FrontDoorTLSProtocolTypeServerNameIndication),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorDisableHttps.json
func ExampleFrontendEndpointsClient_BeginDisableHTTPS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfrontdoor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFrontendEndpointsClient().BeginDisableHTTPS(ctx, "rg1", "frontDoor1", "frontendEndpoint1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
