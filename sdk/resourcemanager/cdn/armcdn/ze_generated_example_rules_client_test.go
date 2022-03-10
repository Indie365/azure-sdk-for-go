//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_ListByRuleSet.json
func ExampleRulesClient_ListByRuleSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	pager := client.ListByRuleSet("<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Get.json
func ExampleRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RulesClientGetResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Create.json
func ExampleRulesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		armcdn.Rule{
			Properties: &armcdn.RuleProperties{
				Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
					&armcdn.DeliveryRuleResponseHeaderAction{
						Name: armcdn.DeliveryRuleAction("ModifyResponseHeader").ToPtr(),
						Parameters: &armcdn.HeaderActionParameters{
							HeaderAction: armcdn.HeaderAction("Overwrite").ToPtr(),
							HeaderName:   to.StringPtr("<header-name>"),
							TypeName:     armcdn.HeaderActionParametersTypeName("DeliveryRuleHeaderActionParameters").ToPtr(),
							Value:        to.StringPtr("<value>"),
						},
					}},
				Conditions: []armcdn.DeliveryRuleConditionClassification{
					&armcdn.DeliveryRuleRequestMethodCondition{
						Name: armcdn.MatchVariable("RequestMethod").ToPtr(),
						Parameters: &armcdn.RequestMethodMatchConditionParameters{
							MatchValues: []*armcdn.RequestMethodMatchConditionParametersMatchValuesItem{
								armcdn.RequestMethodMatchConditionParametersMatchValuesItem("GET").ToPtr()},
							NegateCondition: to.BoolPtr(false),
							Operator:        armcdn.RequestMethodOperator("Equal").ToPtr(),
							TypeName:        armcdn.RequestMethodMatchConditionParametersTypeName("DeliveryRuleRequestMethodConditionParameters").ToPtr(),
						},
					}},
				Order: to.Int32Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RulesClientCreateResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Update.json
func ExampleRulesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		armcdn.RuleUpdateParameters{
			Properties: &armcdn.RuleUpdatePropertiesParameters{
				Actions: []armcdn.DeliveryRuleActionAutoGeneratedClassification{
					&armcdn.DeliveryRuleResponseHeaderAction{
						Name: armcdn.DeliveryRuleAction("ModifyResponseHeader").ToPtr(),
						Parameters: &armcdn.HeaderActionParameters{
							HeaderAction: armcdn.HeaderAction("Overwrite").ToPtr(),
							HeaderName:   to.StringPtr("<header-name>"),
							TypeName:     armcdn.HeaderActionParametersTypeName("DeliveryRuleHeaderActionParameters").ToPtr(),
							Value:        to.StringPtr("<value>"),
						},
					}},
				Order: to.Int32Ptr(1),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RulesClientUpdateResult)
}

// x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Rules_Delete.json
func ExampleRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcdn.NewRulesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<rule-set-name>",
		"<rule-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
