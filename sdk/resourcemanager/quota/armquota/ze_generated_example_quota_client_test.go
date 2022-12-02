//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getComputeOneSkuQuotaLimit.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armquota.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"standardNDSFamily",
		"subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/putNetworkOneSkuQuotaRequest.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armquota.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"MinPublicIpInterNetworkPrefixLength",
		"subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Network/locations/eastus",
		armquota.CurrentQuotaLimitBase{
			Properties: &armquota.Properties{
				Name: &armquota.ResourceName{
					Value: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
				},
				Limit: &armquota.LimitObject{
					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
					Value:           to.Ptr[int32](10),
				},
				ResourceType: to.Ptr("MinPublicIpInterNetworkPrefixLength"),
			},
		},
		nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/patchComputeQuotaRequest.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armquota.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"standardFSv2Family",
		"subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus",
		armquota.CurrentQuotaLimitBase{
			Properties: &armquota.Properties{
				Name: &armquota.ResourceName{
					Value: to.Ptr("standardFSv2Family"),
				},
				Limit: &armquota.LimitObject{
					LimitObjectType: to.Ptr(armquota.LimitTypeLimitValue),
					Value:           to.Ptr[int32](10),
				},
			},
		},
		nil)
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/getComputeQuotaLimits.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armquota.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus",
		nil)
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