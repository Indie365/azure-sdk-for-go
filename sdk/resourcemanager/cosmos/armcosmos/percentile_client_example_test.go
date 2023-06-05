//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBPercentileGetMetrics.json
func ExamplePercentileClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPercentileClient().NewListMetricsPager("rg1", "ddb1", "$filter=(name.value eq 'Probabilistic Bounded Staleness') and timeGrain eq duration'PT5M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z", nil)
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
		// page.PercentileMetricListResult = armcosmos.PercentileMetricListResult{
		// 	Value: []*armcosmos.PercentileMetric{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Probabilistic Bounded Staleness-S-West Central US-T-East US"),
		// 				Value: to.Ptr("Probabilistic Bounded Staleness-S-West Central US-T-East US"),
		// 			},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:13:55.2780000Z"); return t}()),
		// 			MetricValues: []*armcosmos.PercentileMetricValue{
		// 				{
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.2780000Z"); return t}()),
		// 					P10: to.Ptr[float64](1.11),
		// 					P25: to.Ptr[float64](2.5),
		// 					P50: to.Ptr[float64](4.34),
		// 					P75: to.Ptr[float64](5.2),
		// 					P90: to.Ptr[float64](6.77),
		// 					P95: to.Ptr[float64](7.1),
		// 					P99: to.Ptr[float64](8.3),
		// 			}},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.2780000Z"); return t}()),
		// 			TimeGrain: to.Ptr("PT5M"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeMilliseconds),
		// 		},
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Probabilistic Bounded Staleness-S-West Central US-T-West US"),
		// 				Value: to.Ptr("Probabilistic Bounded Staleness-S-West Central US-T-West US"),
		// 			},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:13:55.2780000Z"); return t}()),
		// 			MetricValues: []*armcosmos.PercentileMetricValue{
		// 				{
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.2780000Z"); return t}()),
		// 					P10: to.Ptr[float64](1.11),
		// 					P25: to.Ptr[float64](2.5),
		// 					P50: to.Ptr[float64](4.34),
		// 					P75: to.Ptr[float64](5.2),
		// 					P90: to.Ptr[float64](6.77),
		// 					P95: to.Ptr[float64](7.1),
		// 					P99: to.Ptr[float64](8.3),
		// 			}},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.2780000Z"); return t}()),
		// 			TimeGrain: to.Ptr("PT5M"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeMilliseconds),
		// 	}},
		// }
	}
}
