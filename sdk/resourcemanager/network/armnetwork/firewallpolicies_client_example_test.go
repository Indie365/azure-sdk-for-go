//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyDelete.json
func ExampleFirewallPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "firewallPolicy", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyGet.json
func ExampleFirewallPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "firewallPolicy", &armnetwork.FirewallPoliciesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyPut.json
func ExampleFirewallPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "firewallPolicy", armnetwork.FirewallPolicy{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.FirewallPolicyPropertiesFormat{
			DNSSettings: &armnetwork.DNSSettings{
				EnableProxy:                 to.Ptr(true),
				RequireProxyForNetworkRules: to.Ptr(false),
				Servers: []*string{
					to.Ptr("30.3.4.5")},
			},
			Insights: &armnetwork.FirewallPolicyInsights{
				IsEnabled: to.Ptr(true),
				LogAnalyticsResources: &armnetwork.FirewallPolicyLogAnalyticsResources{
					DefaultWorkspaceID: &armnetwork.SubResource{
						ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/defaultWorkspace"),
					},
					Workspaces: []*armnetwork.FirewallPolicyLogAnalyticsWorkspace{
						{
							Region: to.Ptr("westus"),
							WorkspaceID: &armnetwork.SubResource{
								ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace1"),
							},
						},
						{
							Region: to.Ptr("eastus"),
							WorkspaceID: &armnetwork.SubResource{
								ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/microsoft.operationalinsights/workspaces/workspace2"),
							},
						}},
				},
				RetentionDays: to.Ptr[int32](100),
			},
			IntrusionDetection: &armnetwork.FirewallPolicyIntrusionDetection{
				Configuration: &armnetwork.FirewallPolicyIntrusionDetectionConfiguration{
					BypassTrafficSettings: []*armnetwork.FirewallPolicyIntrusionDetectionBypassTrafficSpecifications{
						{
							Name:        to.Ptr("bypassRule1"),
							Description: to.Ptr("Rule 1"),
							DestinationAddresses: []*string{
								to.Ptr("5.6.7.8")},
							DestinationPorts: []*string{
								to.Ptr("*")},
							SourceAddresses: []*string{
								to.Ptr("1.2.3.4")},
							Protocol: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionProtocolTCP),
						}},
					SignatureOverrides: []*armnetwork.FirewallPolicyIntrusionDetectionSignatureSpecification{
						{
							ID:   to.Ptr("2525004"),
							Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeDeny),
						}},
				},
				Mode: to.Ptr(armnetwork.FirewallPolicyIntrusionDetectionStateTypeAlert),
			},
			SKU: &armnetwork.FirewallPolicySKU{
				Tier: to.Ptr(armnetwork.FirewallPolicySKUTierPremium),
			},
			Snat: &armnetwork.FirewallPolicySNAT{
				PrivateRanges: []*string{
					to.Ptr("IANAPrivateRanges")},
			},
			SQL: &armnetwork.FirewallPolicySQL{
				AllowSQLRedirect: to.Ptr(true),
			},
			ThreatIntelMode: to.Ptr(armnetwork.AzureFirewallThreatIntelModeAlert),
			ThreatIntelWhitelist: &armnetwork.FirewallPolicyThreatIntelWhitelist{
				Fqdns: []*string{
					to.Ptr("*.microsoft.com")},
				IPAddresses: []*string{
					to.Ptr("20.3.4.5")},
			},
			TransportSecurity: &armnetwork.FirewallPolicyTransportSecurity{
				CertificateAuthority: &armnetwork.FirewallPolicyCertificateAuthority{
					Name:             to.Ptr("clientcert"),
					KeyVaultSecretID: to.Ptr("https://kv/secret"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyPatch.json
func ExampleFirewallPoliciesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPoliciesClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UpdateTags(ctx, "myResourceGroup", "firewallPolicy", armnetwork.TagsObject{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyListByResourceGroup.json
func ExampleFirewallPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyListBySubscription.json
func ExampleFirewallPoliciesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewFirewallPoliciesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListAllPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
