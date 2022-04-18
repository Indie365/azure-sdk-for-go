//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorListAll.json
func ExampleFrontDoorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorList.json
func ExampleFrontDoorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByResourceGroupPager("<resource-group-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorGet.json
func ExampleFrontDoorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorCreate.json
func ExampleFrontDoorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		armfrontdoor.FrontDoor{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armfrontdoor.Properties{
				BackendPools: []*armfrontdoor.BackendPool{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.BackendPoolProperties{
							Backends: []*armfrontdoor.Backend{
								{
									Address:   to.Ptr("<address>"),
									HTTPPort:  to.Ptr[int32](80),
									HTTPSPort: to.Ptr[int32](443),
									Priority:  to.Ptr[int32](2),
									Weight:    to.Ptr[int32](1),
								},
								{
									Address:                    to.Ptr("<address>"),
									HTTPPort:                   to.Ptr[int32](80),
									HTTPSPort:                  to.Ptr[int32](443),
									Priority:                   to.Ptr[int32](1),
									PrivateLinkApprovalMessage: to.Ptr("<private-link-approval-message>"),
									PrivateLinkLocation:        to.Ptr("<private-link-location>"),
									PrivateLinkResourceID:      to.Ptr("<private-link-resource-id>"),
									Weight:                     to.Ptr[int32](2),
								},
								{
									Address:                    to.Ptr("<address>"),
									HTTPPort:                   to.Ptr[int32](80),
									HTTPSPort:                  to.Ptr[int32](443),
									Priority:                   to.Ptr[int32](1),
									PrivateLinkAlias:           to.Ptr("<private-link-alias>"),
									PrivateLinkApprovalMessage: to.Ptr("<private-link-approval-message>"),
									Weight:                     to.Ptr[int32](1),
								}},
							HealthProbeSettings: &armfrontdoor.SubResource{
								ID: to.Ptr("<id>"),
							},
							LoadBalancingSettings: &armfrontdoor.SubResource{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				BackendPoolsSettings: &armfrontdoor.BackendPoolsSettings{
					EnforceCertificateNameCheck: to.Ptr(armfrontdoor.EnforceCertificateNameCheckEnabledStateEnabled),
					SendRecvTimeoutSeconds:      to.Ptr[int32](60),
				},
				EnabledState: to.Ptr(armfrontdoor.FrontDoorEnabledStateEnabled),
				FrontendEndpoints: []*armfrontdoor.FrontendEndpoint{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName:                    to.Ptr("<host-name>"),
							SessionAffinityEnabledState: to.Ptr(armfrontdoor.SessionAffinityEnabledStateEnabled),
							SessionAffinityTTLSeconds:   to.Ptr[int32](60),
							WebApplicationFirewallPolicyLink: &armfrontdoor.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.Ptr("<id>"),
							},
						},
					},
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName: to.Ptr("<host-name>"),
						},
					}},
				HealthProbeSettings: []*armfrontdoor.HealthProbeSettingsModel{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.HealthProbeSettingsProperties{
							Path:              to.Ptr("<path>"),
							EnabledState:      to.Ptr(armfrontdoor.HealthProbeEnabledEnabled),
							HealthProbeMethod: to.Ptr(armfrontdoor.FrontDoorHealthProbeMethodHEAD),
							IntervalInSeconds: to.Ptr[int32](120),
							Protocol:          to.Ptr(armfrontdoor.FrontDoorProtocolHTTP),
						},
					}},
				LoadBalancingSettings: []*armfrontdoor.LoadBalancingSettingsModel{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.LoadBalancingSettingsProperties{
							SampleSize:                to.Ptr[int32](4),
							SuccessfulSamplesRequired: to.Ptr[int32](2),
						},
					}},
				RoutingRules: []*armfrontdoor.RoutingRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armfrontdoor.RoutingRuleProperties{
							AcceptedProtocols: []*armfrontdoor.FrontDoorProtocol{
								to.Ptr(armfrontdoor.FrontDoorProtocolHTTP)},
							EnabledState: to.Ptr(armfrontdoor.RoutingRuleEnabledStateEnabled),
							FrontendEndpoints: []*armfrontdoor.SubResource{
								{
									ID: to.Ptr("<id>"),
								},
								{
									ID: to.Ptr("<id>"),
								}},
							PatternsToMatch: []*string{
								to.Ptr("/*")},
							RouteConfiguration: &armfrontdoor.ForwardingConfiguration{
								ODataType: to.Ptr("<odata-type>"),
								BackendPool: &armfrontdoor.SubResource{
									ID: to.Ptr("<id>"),
								},
							},
							RulesEngine: &armfrontdoor.SubResource{
								ID: to.Ptr("<id>"),
							},
							WebApplicationFirewallPolicyLink: &armfrontdoor.RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.Ptr("<id>"),
							},
						},
					}},
			},
		},
		&armfrontdoor.FrontDoorsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorDelete.json
func ExampleFrontDoorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		&armfrontdoor.FrontDoorsClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorValidateCustomDomain.json
func ExampleFrontDoorsClient_ValidateCustomDomain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.ValidateCustomDomain(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		armfrontdoor.ValidateCustomDomainInput{
			HostName: to.Ptr("<host-name>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
