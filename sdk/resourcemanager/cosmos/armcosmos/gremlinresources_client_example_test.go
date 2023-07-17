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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseList.json
func ExampleGremlinResourcesClient_NewListGremlinDatabasesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGremlinResourcesClient().NewListGremlinDatabasesPager("rgName", "ddb1", nil)
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
		// page.GremlinDatabaseListResult = armcosmos.GremlinDatabaseListResult{
		// 	Value: []*armcosmos.GremlinDatabaseGetResults{
		// 		{
		// 			Name: to.Ptr("databaseName"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.GremlinDatabaseGetProperties{
		// 				Resource: &armcosmos.GremlinDatabaseGetPropertiesResource{
		// 					Etag: to.Ptr("\"00000a00-0000-0000-0000-56672f920000\""),
		// 					Rid: to.Ptr("CqNBAA=="),
		// 					Ts: to.Ptr[float32](1449602962),
		// 					ID: to.Ptr("databaseName"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseGet.json
func ExampleGremlinResourcesClient_GetGremlinDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGremlinResourcesClient().GetGremlinDatabase(ctx, "rg1", "ddb1", "databaseName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GremlinDatabaseGetResults = armcosmos.GremlinDatabaseGetResults{
	// 	Name: to.Ptr("databaseName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.GremlinDatabaseGetProperties{
	// 		Resource: &armcosmos.GremlinDatabaseGetPropertiesResource{
	// 			Etag: to.Ptr("\"00000a00-0000-0000-0000-56672f920000\""),
	// 			Rid: to.Ptr("CqNBAA=="),
	// 			Ts: to.Ptr[float32](1449602962),
	// 			ID: to.Ptr("databaseName"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseCreateUpdate.json
func ExampleGremlinResourcesClient_BeginCreateUpdateGremlinDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginCreateUpdateGremlinDatabase(ctx, "rg1", "ddb1", "databaseName", armcosmos.GremlinDatabaseCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.GremlinDatabaseCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.GremlinDatabaseResource{
				ID: to.Ptr("databaseName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GremlinDatabaseGetResults = armcosmos.GremlinDatabaseGetResults{
	// 	Name: to.Ptr("databaseName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.GremlinDatabaseGetProperties{
	// 		Resource: &armcosmos.GremlinDatabaseGetPropertiesResource{
	// 			Etag: to.Ptr("\"00000a00-0000-0000-0000-56672f920000\""),
	// 			Rid: to.Ptr("CqNBAA=="),
	// 			Ts: to.Ptr[float32](1449602962),
	// 			ID: to.Ptr("databaseName"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseDelete.json
func ExampleGremlinResourcesClient_BeginDeleteGremlinDatabase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginDeleteGremlinDatabase(ctx, "rg1", "ddb1", "databaseName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseThroughputGet.json
func ExampleGremlinResourcesClient_GetGremlinDatabaseThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGremlinResourcesClient().GetGremlinDatabaseThroughput(ctx, "rg1", "ddb1", "databaseName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases/throughputSettings"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName/throughputSettings/default"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseThroughputUpdate.json
func ExampleGremlinResourcesClient_BeginUpdateGremlinDatabaseThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginUpdateGremlinDatabaseThroughput(ctx, "rg1", "ddb1", "databaseName", armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases/throughputSettings"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName/throughputSettings/default"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseMigrateToAutoscale.json
func ExampleGremlinResourcesClient_BeginMigrateGremlinDatabaseToAutoscale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginMigrateGremlinDatabaseToAutoscale(ctx, "rg1", "ddb1", "databaseName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			AutoscaleSettings: &armcosmos.AutoscaleSettingsResource{
	// 				MaxThroughput: to.Ptr[int32](4000),
	// 			},
	// 			MinimumThroughput: to.Ptr("4000"),
	// 			OfferReplacePending: to.Ptr("false"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinDatabaseMigrateToManualThroughput.json
func ExampleGremlinResourcesClient_BeginMigrateGremlinDatabaseToManualThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginMigrateGremlinDatabaseToManualThroughput(ctx, "rg1", "ddb1", "databaseName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphList.json
func ExampleGremlinResourcesClient_NewListGremlinGraphsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGremlinResourcesClient().NewListGremlinGraphsPager("rgName", "ddb1", "databaseName", nil)
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
		// page.GremlinGraphListResult = armcosmos.GremlinGraphListResult{
		// 	Value: []*armcosmos.GremlinGraphGetResults{
		// 		{
		// 			Name: to.Ptr("testgrf"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/apis/databases/graphs"),
		// 			ID: to.Ptr("testgrf"),
		// 			Properties: &armcosmos.GremlinGraphGetProperties{
		// 				Resource: &armcosmos.GremlinGraphGetPropertiesResource{
		// 					Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
		// 					Rid: to.Ptr("PD5DALigDgw="),
		// 					Ts: to.Ptr[float32](1459200611),
		// 					ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
		// 						ConflictResolutionPath: to.Ptr("/path"),
		// 						Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
		// 					},
		// 					DefaultTTL: to.Ptr[int32](100),
		// 					ID: to.Ptr("testgrf"),
		// 					IndexingPolicy: &armcosmos.IndexingPolicy{
		// 						Automatic: to.Ptr(true),
		// 						ExcludedPaths: []*armcosmos.ExcludedPath{
		// 						},
		// 						IncludedPaths: []*armcosmos.IncludedPath{
		// 							{
		// 								Path: to.Ptr("/*"),
		// 								Indexes: []*armcosmos.Indexes{
		// 									{
		// 										DataType: to.Ptr(armcosmos.DataTypeString),
		// 										Kind: to.Ptr(armcosmos.IndexKindRange),
		// 										Precision: to.Ptr[int32](-1),
		// 									},
		// 									{
		// 										DataType: to.Ptr(armcosmos.DataTypeNumber),
		// 										Kind: to.Ptr(armcosmos.IndexKindRange),
		// 										Precision: to.Ptr[int32](-1),
		// 								}},
		// 						}},
		// 						IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
		// 					},
		// 					PartitionKey: &armcosmos.ContainerPartitionKey{
		// 						Kind: to.Ptr(armcosmos.PartitionKindHash),
		// 						Paths: []*string{
		// 							to.Ptr("/AccountNumber")},
		// 						},
		// 						UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
		// 							UniqueKeys: []*armcosmos.UniqueKey{
		// 								{
		// 									Paths: []*string{
		// 										to.Ptr("/testPath")},
		// 								}},
		// 							},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphGet.json
func ExampleGremlinResourcesClient_GetGremlinGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGremlinResourcesClient().GetGremlinGraph(ctx, "rgName", "ddb1", "databaseName", "graphName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GremlinGraphGetResults = armcosmos.GremlinGraphGetResults{
	// 	Name: to.Ptr("graphName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/apis/databases/graphs"),
	// 	ID: to.Ptr("graphName"),
	// 	Properties: &armcosmos.GremlinGraphGetProperties{
	// 		Resource: &armcosmos.GremlinGraphGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
	// 				ConflictResolutionPath: to.Ptr("/path"),
	// 				Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
	// 			},
	// 			DefaultTTL: to.Ptr[int32](100),
	// 			ID: to.Ptr("graphName"),
	// 			IndexingPolicy: &armcosmos.IndexingPolicy{
	// 				Automatic: to.Ptr(true),
	// 				ExcludedPaths: []*armcosmos.ExcludedPath{
	// 				},
	// 				IncludedPaths: []*armcosmos.IncludedPath{
	// 					{
	// 						Path: to.Ptr("/*"),
	// 						Indexes: []*armcosmos.Indexes{
	// 							{
	// 								DataType: to.Ptr(armcosmos.DataTypeString),
	// 								Kind: to.Ptr(armcosmos.IndexKindRange),
	// 								Precision: to.Ptr[int32](-1),
	// 							},
	// 							{
	// 								DataType: to.Ptr(armcosmos.DataTypeNumber),
	// 								Kind: to.Ptr(armcosmos.IndexKindRange),
	// 								Precision: to.Ptr[int32](-1),
	// 						}},
	// 				}},
	// 				IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
	// 			},
	// 			PartitionKey: &armcosmos.ContainerPartitionKey{
	// 				Kind: to.Ptr(armcosmos.PartitionKindHash),
	// 				Paths: []*string{
	// 					to.Ptr("/AccountNumber")},
	// 				},
	// 				UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
	// 					UniqueKeys: []*armcosmos.UniqueKey{
	// 						{
	// 							Paths: []*string{
	// 								to.Ptr("/testPath")},
	// 						}},
	// 					},
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphCreateUpdate.json
func ExampleGremlinResourcesClient_BeginCreateUpdateGremlinGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginCreateUpdateGremlinGraph(ctx, "rg1", "ddb1", "databaseName", "graphName", armcosmos.GremlinGraphCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.GremlinGraphCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.GremlinGraphResource{
				ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
					ConflictResolutionPath: to.Ptr("/path"),
					Mode:                   to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
				},
				DefaultTTL: to.Ptr[int32](100),
				ID:         to.Ptr("graphName"),
				IndexingPolicy: &armcosmos.IndexingPolicy{
					Automatic:     to.Ptr(true),
					ExcludedPaths: []*armcosmos.ExcludedPath{},
					IncludedPaths: []*armcosmos.IncludedPath{
						{
							Path: to.Ptr("/*"),
							Indexes: []*armcosmos.Indexes{
								{
									DataType:  to.Ptr(armcosmos.DataTypeString),
									Kind:      to.Ptr(armcosmos.IndexKindRange),
									Precision: to.Ptr[int32](-1),
								},
								{
									DataType:  to.Ptr(armcosmos.DataTypeNumber),
									Kind:      to.Ptr(armcosmos.IndexKindRange),
									Precision: to.Ptr[int32](-1),
								}},
						}},
					IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
				},
				PartitionKey: &armcosmos.ContainerPartitionKey{
					Kind: to.Ptr(armcosmos.PartitionKindHash),
					Paths: []*string{
						to.Ptr("/AccountNumber")},
				},
				UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
					UniqueKeys: []*armcosmos.UniqueKey{
						{
							Paths: []*string{
								to.Ptr("/testPath")},
						}},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GremlinGraphGetResults = armcosmos.GremlinGraphGetResults{
	// 	Name: to.Ptr("graphName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases/gremlinGraphs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName/gremlinGraphs/graphName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.GremlinGraphGetProperties{
	// 		Resource: &armcosmos.GremlinGraphGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			ConflictResolutionPolicy: &armcosmos.ConflictResolutionPolicy{
	// 				ConflictResolutionPath: to.Ptr("/path"),
	// 				Mode: to.Ptr(armcosmos.ConflictResolutionModeLastWriterWins),
	// 			},
	// 			DefaultTTL: to.Ptr[int32](100),
	// 			ID: to.Ptr("graphName"),
	// 			IndexingPolicy: &armcosmos.IndexingPolicy{
	// 				Automatic: to.Ptr(true),
	// 				ExcludedPaths: []*armcosmos.ExcludedPath{
	// 				},
	// 				IncludedPaths: []*armcosmos.IncludedPath{
	// 					{
	// 						Path: to.Ptr("/*"),
	// 						Indexes: []*armcosmos.Indexes{
	// 							{
	// 								DataType: to.Ptr(armcosmos.DataTypeString),
	// 								Kind: to.Ptr(armcosmos.IndexKindRange),
	// 								Precision: to.Ptr[int32](-1),
	// 							},
	// 							{
	// 								DataType: to.Ptr(armcosmos.DataTypeNumber),
	// 								Kind: to.Ptr(armcosmos.IndexKindRange),
	// 								Precision: to.Ptr[int32](-1),
	// 						}},
	// 				}},
	// 				IndexingMode: to.Ptr(armcosmos.IndexingModeConsistent),
	// 			},
	// 			PartitionKey: &armcosmos.ContainerPartitionKey{
	// 				Kind: to.Ptr(armcosmos.PartitionKindHash),
	// 				Paths: []*string{
	// 					to.Ptr("/AccountNumber")},
	// 				},
	// 				UniqueKeyPolicy: &armcosmos.UniqueKeyPolicy{
	// 					UniqueKeys: []*armcosmos.UniqueKey{
	// 						{
	// 							Paths: []*string{
	// 								to.Ptr("/testPath")},
	// 						}},
	// 					},
	// 				},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphDelete.json
func ExampleGremlinResourcesClient_BeginDeleteGremlinGraph() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginDeleteGremlinGraph(ctx, "rg1", "ddb1", "databaseName", "graphName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphThroughputGet.json
func ExampleGremlinResourcesClient_GetGremlinGraphThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGremlinResourcesClient().GetGremlinGraphThroughput(ctx, "rg1", "ddb1", "databaseName", "graphName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases/throughputSettings"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName/gremlinGraphs/graphName/throughputSettings/default"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphThroughputUpdate.json
func ExampleGremlinResourcesClient_BeginUpdateGremlinGraphThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginUpdateGremlinGraphThroughput(ctx, "rg1", "ddb1", "databaseName", "graphName", armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/gremlinDatabases/throughputSettings"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/gremlinDatabases/databaseName/gremlinGraphs/graphName/throughputSettings/default"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphMigrateToAutoscale.json
func ExampleGremlinResourcesClient_BeginMigrateGremlinGraphToAutoscale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginMigrateGremlinGraphToAutoscale(ctx, "rg1", "ddb1", "databaseName", "graphName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			AutoscaleSettings: &armcosmos.AutoscaleSettingsResource{
	// 				MaxThroughput: to.Ptr[int32](4000),
	// 			},
	// 			MinimumThroughput: to.Ptr("4000"),
	// 			OfferReplacePending: to.Ptr("false"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphMigrateToManualThroughput.json
func ExampleGremlinResourcesClient_BeginMigrateGremlinGraphToManualThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginMigrateGremlinGraphToManualThroughput(ctx, "rg1", "ddb1", "databaseName", "graphName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBGremlinGraphBackupInformation.json
func ExampleGremlinResourcesClient_BeginRetrieveContinuousBackupInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGremlinResourcesClient().BeginRetrieveContinuousBackupInformation(ctx, "rgName", "ddb1", "databaseName", "graphName", armcosmos.ContinuousBackupRestoreLocation{
		Location: to.Ptr("North Europe"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupInformation = armcosmos.BackupInformation{
	// 	ContinuousBackupInformation: &armcosmos.ContinuousBackupInformation{
	// 		LatestRestorableTimestamp: to.Ptr("2021-02-05T02:40:50Z"),
	// 	},
	// }
}
