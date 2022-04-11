//go:build go1.16
// +build go1.16

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineList.json
func ExampleRulesEnginesClient_ListByFrontDoor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByFrontDoor("<resource-group-name>",
		"<front-door-name>",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineGet.json
func ExampleRulesEnginesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<rules-engine-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineCreate.json
func ExampleRulesEnginesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<rules-engine-name>",
		armfrontdoor.RulesEngine{
			Properties: &armfrontdoor.RulesEngineProperties{
				Rules: []*armfrontdoor.RulesEngineRule{
					{
						Name: to.Ptr("<name>"),
						Action: &armfrontdoor.RulesEngineAction{
							RouteConfigurationOverride: &armfrontdoor.RedirectConfiguration{
								ODataType:         to.Ptr("<odata-type>"),
								CustomFragment:    to.Ptr("<custom-fragment>"),
								CustomHost:        to.Ptr("<custom-host>"),
								CustomPath:        to.Ptr("<custom-path>"),
								CustomQueryString: to.Ptr("<custom-query-string>"),
								RedirectProtocol:  to.Ptr(armfrontdoor.FrontDoorRedirectProtocolHTTPSOnly),
								RedirectType:      to.Ptr(armfrontdoor.FrontDoorRedirectTypeMoved),
							},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								RulesEngineMatchValue: []*string{
									to.Ptr("CH")},
								RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRemoteAddr),
								RulesEngineOperator:      to.Ptr(armfrontdoor.RulesEngineOperatorGeoMatch),
							}},
						MatchProcessingBehavior: to.Ptr(armfrontdoor.MatchProcessingBehaviorStop),
						Priority:                to.Ptr[int32](1),
					},
					{
						Name: to.Ptr("<name>"),
						Action: &armfrontdoor.RulesEngineAction{
							ResponseHeaderActions: []*armfrontdoor.HeaderAction{
								{
									HeaderActionType: to.Ptr(armfrontdoor.HeaderActionTypeOverwrite),
									HeaderName:       to.Ptr("<header-name>"),
									Value:            to.Ptr("<value>"),
								}},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								RulesEngineMatchValue: []*string{
									to.Ptr("jpg")},
								RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRequestFilenameExtension),
								RulesEngineOperator:      to.Ptr(armfrontdoor.RulesEngineOperatorEqual),
								Transforms: []*armfrontdoor.Transform{
									to.Ptr(armfrontdoor.TransformLowercase)},
							}},
						Priority: to.Ptr[int32](2),
					},
					{
						Name: to.Ptr("<name>"),
						Action: &armfrontdoor.RulesEngineAction{
							RouteConfigurationOverride: &armfrontdoor.ForwardingConfiguration{
								ODataType: to.Ptr("<odata-type>"),
								BackendPool: &armfrontdoor.SubResource{
									ID: to.Ptr("<id>"),
								},
								CacheConfiguration: &armfrontdoor.CacheConfiguration{
									CacheDuration:                to.Ptr("<cache-duration>"),
									DynamicCompression:           to.Ptr(armfrontdoor.DynamicCompressionEnabledDisabled),
									QueryParameterStripDirective: to.Ptr(armfrontdoor.FrontDoorQueryStripOnly),
									QueryParameters:              to.Ptr("<query-parameters>"),
								},
								ForwardingProtocol: to.Ptr(armfrontdoor.FrontDoorForwardingProtocolHTTPSOnly),
							},
						},
						MatchConditions: []*armfrontdoor.RulesEngineMatchCondition{
							{
								NegateCondition: to.Ptr(false),
								RulesEngineMatchValue: []*string{
									to.Ptr("allowoverride")},
								RulesEngineMatchVariable: to.Ptr(armfrontdoor.RulesEngineMatchVariableRequestHeader),
								RulesEngineOperator:      to.Ptr(armfrontdoor.RulesEngineOperatorEqual),
								Selector:                 to.Ptr("<selector>"),
								Transforms: []*armfrontdoor.Transform{
									to.Ptr(armfrontdoor.TransformLowercase)},
							}},
						Priority: to.Ptr[int32](3),
					}},
			},
		},
		&armfrontdoor.RulesEnginesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorRulesEngineDelete.json
func ExampleRulesEnginesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewRulesEnginesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		"<rules-engine-name>",
		&armfrontdoor.RulesEnginesClientBeginDeleteOptions{ResumeToken: ""})
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
