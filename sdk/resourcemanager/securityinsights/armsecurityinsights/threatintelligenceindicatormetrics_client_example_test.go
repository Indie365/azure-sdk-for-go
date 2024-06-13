//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/CollectThreatIntelligenceMetrics.json
func ExampleThreatIntelligenceIndicatorMetricsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewThreatIntelligenceIndicatorMetricsClient().List(ctx, "myRg", "myWorkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThreatIntelligenceMetricsList = armsecurityinsights.ThreatIntelligenceMetricsList{
	// 	Value: []*armsecurityinsights.ThreatIntelligenceMetrics{
	// 		{
	// 			Properties: &armsecurityinsights.ThreatIntelligenceMetric{
	// 				LastUpdatedTimeUTC: to.Ptr("2021-09-01T19:44:44.117403Z"),
	// 				PatternTypeMetrics: []*armsecurityinsights.ThreatIntelligenceMetricEntity{
	// 					{
	// 						MetricName: to.Ptr("url"),
	// 						MetricValue: to.Ptr[int32](20),
	// 				}},
	// 				SourceMetrics: []*armsecurityinsights.ThreatIntelligenceMetricEntity{
	// 					{
	// 						MetricName: to.Ptr("Azure Sentinel"),
	// 						MetricValue: to.Ptr[int32](10315),
	// 					},
	// 					{
	// 						MetricName: to.Ptr("zinga"),
	// 						MetricValue: to.Ptr[int32](2),
	// 				}},
	// 				ThreatTypeMetrics: []*armsecurityinsights.ThreatIntelligenceMetricEntity{
	// 					{
	// 						MetricName: to.Ptr("compromised"),
	// 						MetricValue: to.Ptr[int32](20),
	// 				}},
	// 			},
	// 	}},
	// }
}
