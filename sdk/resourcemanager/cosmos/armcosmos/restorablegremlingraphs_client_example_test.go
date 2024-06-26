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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ec7ee8842bf615c2f0354bf8b5b8725fdac9454a/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-05-15/examples/CosmosDBRestorableGremlinGraphList.json
func ExampleRestorableGremlinGraphsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorableGremlinGraphsClient().NewListPager("WestUS", "98a570f2-63db-4117-91f0-366327b7b353", &armcosmos.RestorableGremlinGraphsClientListOptions{RestorableGremlinDatabaseRid: to.Ptr("PD5DALigDgw="),
		StartTime: nil,
		EndTime:   nil,
	})
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
		// page.RestorableGremlinGraphsListResult = armcosmos.RestorableGremlinGraphsListResult{
		// 	Value: []*armcosmos.RestorableGremlinGraphGetResult{
		// 		{
		// 			Name: to.Ptr("79609a98-3394-41f8-911f-cfab0c075c86"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts/restorableGraphs"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDb/locations/westus/restorableDatabaseAccounts/98a570f2-63db-4117-91f0-366327b7b353/restorableGraphs/79609a98-3394-41f8-911f-cfab0c075c86"),
		// 			Properties: &armcosmos.RestorableGremlinGraphProperties{
		// 				Resource: &armcosmos.RestorableGremlinGraphPropertiesResource{
		// 					Rid: to.Ptr("zAyAPQAAAA=="),
		// 					CanUndelete: to.Ptr("invalid"),
		// 					EventTimestamp: to.Ptr("2020-10-13T04:56:42Z"),
		// 					OperationType: to.Ptr(armcosmos.OperationTypeCreate),
		// 					OwnerID: to.Ptr("Graph1"),
		// 					OwnerResourceID: to.Ptr("V18LoLrv-qA="),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("79609a98-3394-41f8-911f-cfab0c075c86"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/locations/restorableDatabaseAccounts/restorableGraphs"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.DocumentDb/locations/westus/restorableDatabaseAccounts/98a570f2-63db-4117-91f0-366327b7b353/restorableGraphs/79609a98-3394-41f8-911f-cfab0c075c86"),
		// 			Properties: &armcosmos.RestorableGremlinGraphProperties{
		// 				Resource: &armcosmos.RestorableGremlinGraphPropertiesResource{
		// 					Rid: to.Ptr("zAyAPQAAAA=="),
		// 					CanUndelete: to.Ptr("notRestorable"),
		// 					CanUndeleteReason: to.Ptr("Individual shared database collections restore is not supported. Please restore shared database to restore its collections that share the throughput."),
		// 					EventTimestamp: to.Ptr("2020-10-13T04:56:42Z"),
		// 					OperationType: to.Ptr(armcosmos.OperationTypeDelete),
		// 					OwnerID: to.Ptr("Graph1"),
		// 					OwnerResourceID: to.Ptr("V18LoLrv-qA="),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
