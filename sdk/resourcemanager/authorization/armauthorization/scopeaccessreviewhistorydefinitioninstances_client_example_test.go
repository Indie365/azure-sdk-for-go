//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2021-12-01-preview/examples/GetAccessReviewHistoryDefinitionInstances.json
func ExampleScopeAccessReviewHistoryDefinitionInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScopeAccessReviewHistoryDefinitionInstancesClient().NewListPager("subscriptions/129a304b-4aea-4b86-a9f7-ba7e2b23737a", "44724910-d7a5-4c29-b28f-db73e717165a", nil)
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
		// page.AccessReviewHistoryDefinitionInstanceListResult = armauthorization.AccessReviewHistoryDefinitionInstanceListResult{
		// 	Value: []*armauthorization.AccessReviewHistoryInstance{
		// 		{
		// 			Name: to.Ptr("44724910-d7a5-4c29-b28f-db73e717165a"),
		// 			Type: to.Ptr("Microsoft.Authorization/accessReviewHistoryInstance"),
		// 			ID: to.Ptr("44724910-d7a5-4c29-b28f-db73e717165a"),
		// 			Properties: &armauthorization.AccessReviewHistoryInstanceProperties{
		// 				DisplayName: to.Ptr("Hello world name"),
		// 				Expiration: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-22T00:28:33.6905189+00:00"); return t}()),
		// 				FulfilledDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-23T00:28:33.6905189+00:00"); return t}()),
		// 				ReviewHistoryPeriodEndDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-01T00:00:00-08:00"); return t}()),
		// 				ReviewHistoryPeriodStartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00-07:00"); return t}()),
		// 				RunDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-12-23T00:27:33.6905189+00:00"); return t}()),
		// 				Status: to.Ptr(armauthorization.AccessReviewHistoryDefinitionStatusRequested),
		// 			},
		// 	}},
		// }
	}
}
