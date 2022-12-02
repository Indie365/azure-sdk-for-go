//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/BillingAccountQuery.json
func ExampleQueryClient_Usage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewQueryClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Usage(ctx,
		"providers/Microsoft.Billing/billingAccounts/70664866",
		armcostmanagement.QueryDefinition{
			Type: to.Ptr(armcostmanagement.ExportTypeUsage),
			Dataset: &armcostmanagement.QueryDataset{
				Filter: &armcostmanagement.QueryFilter{
					And: []*armcostmanagement.QueryFilter{
						{
							Or: []*armcostmanagement.QueryFilter{
								{
									Dimensions: &armcostmanagement.QueryComparisonExpression{
										Name:     to.Ptr("ResourceLocation"),
										Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
										Values: []*string{
											to.Ptr("East US"),
											to.Ptr("West Europe")},
									},
								},
								{
									Tags: &armcostmanagement.QueryComparisonExpression{
										Name:     to.Ptr("Environment"),
										Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
										Values: []*string{
											to.Ptr("UAT"),
											to.Ptr("Prod")},
									},
								}},
						},
						{
							Dimensions: &armcostmanagement.QueryComparisonExpression{
								Name:     to.Ptr("ResourceGroup"),
								Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
								Values: []*string{
									to.Ptr("API")},
							},
						}},
				},
				Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
			},
			Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalBillingAccountsQuery.json
func ExampleQueryClient_UsageByExternalCloudProviderType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewQueryClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.UsageByExternalCloudProviderType(ctx,
		armcostmanagement.ExternalCloudProviderTypeExternalBillingAccounts,
		"100",
		armcostmanagement.QueryDefinition{
			Type: to.Ptr(armcostmanagement.ExportTypeUsage),
			Dataset: &armcostmanagement.QueryDataset{
				Filter: &armcostmanagement.QueryFilter{
					And: []*armcostmanagement.QueryFilter{
						{
							Or: []*armcostmanagement.QueryFilter{
								{
									Dimensions: &armcostmanagement.QueryComparisonExpression{
										Name:     to.Ptr("ResourceLocation"),
										Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
										Values: []*string{
											to.Ptr("East US"),
											to.Ptr("West Europe")},
									},
								},
								{
									Tags: &armcostmanagement.QueryComparisonExpression{
										Name:     to.Ptr("Environment"),
										Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
										Values: []*string{
											to.Ptr("UAT"),
											to.Ptr("Prod")},
									},
								}},
						},
						{
							Dimensions: &armcostmanagement.QueryComparisonExpression{
								Name:     to.Ptr("ResourceGroup"),
								Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
								Values: []*string{
									to.Ptr("API")},
							},
						}},
				},
				Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
			},
			Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}