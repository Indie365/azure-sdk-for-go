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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetLatencyScorecard.json
func ExampleReportsClient_GetLatencyScorecards() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewReportsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetLatencyScorecards(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<experiment-name>",
		armfrontdoor.LatencyScorecardAggregationIntervalDaily,
		&armfrontdoor.ReportsClientGetLatencyScorecardsOptions{EndDateTimeUTC: nil,
			Country: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetTimeseries.json
func ExampleReportsClient_GetTimeseries() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armfrontdoor.NewReportsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetTimeseries(ctx,
		"<resource-group-name>",
		"<profile-name>",
		"<experiment-name>",
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-21T17:32:28Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-21T17:32:28Z"); return t }(),
		armfrontdoor.TimeseriesAggregationIntervalHourly,
		armfrontdoor.TimeseriesTypeMeasurementCounts,
		&armfrontdoor.ReportsClientGetTimeseriesOptions{Endpoint: nil,
			Country: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
