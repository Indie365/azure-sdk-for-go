//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential azcore.TokenCredential
	options *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: 	subscriptionID,		credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewCassandraClustersClient() *CassandraClustersClient {
	subClient, _ := NewCassandraClustersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCassandraDataCentersClient() *CassandraDataCentersClient {
	subClient, _ := NewCassandraDataCentersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCassandraResourcesClient() *CassandraResourcesClient {
	subClient, _ := NewCassandraResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCollectionClient() *CollectionClient {
	subClient, _ := NewCollectionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCollectionPartitionClient() *CollectionPartitionClient {
	subClient, _ := NewCollectionPartitionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCollectionPartitionRegionClient() *CollectionPartitionRegionClient {
	subClient, _ := NewCollectionPartitionRegionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCollectionRegionClient() *CollectionRegionClient {
	subClient, _ := NewCollectionRegionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataTransferJobsClient() *DataTransferJobsClient {
	subClient, _ := NewDataTransferJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatabaseAccountRegionClient() *DatabaseAccountRegionClient {
	subClient, _ := NewDatabaseAccountRegionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatabaseAccountsClient() *DatabaseAccountsClient {
	subClient, _ := NewDatabaseAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatabaseClient() *DatabaseClient {
	subClient, _ := NewDatabaseClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGraphResourcesClient() *GraphResourcesClient {
	subClient, _ := NewGraphResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGremlinResourcesClient() *GremlinResourcesClient {
	subClient, _ := NewGremlinResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLocationsClient() *LocationsClient {
	subClient, _ := NewLocationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMongoClustersClient() *MongoClustersClient {
	subClient, _ := NewMongoClustersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMongoDBResourcesClient() *MongoDBResourcesClient {
	subClient, _ := NewMongoDBResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNotebookWorkspacesClient() *NotebookWorkspacesClient {
	subClient, _ := NewNotebookWorkspacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPartitionKeyRangeIDClient() *PartitionKeyRangeIDClient {
	subClient, _ := NewPartitionKeyRangeIDClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPartitionKeyRangeIDRegionClient() *PartitionKeyRangeIDRegionClient {
	subClient, _ := NewPartitionKeyRangeIDRegionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPercentileClient() *PercentileClient {
	subClient, _ := NewPercentileClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPercentileSourceTargetClient() *PercentileSourceTargetClient {
	subClient, _ := NewPercentileSourceTargetClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPercentileTargetClient() *PercentileTargetClient {
	subClient, _ := NewPercentileTargetClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableDatabaseAccountsClient() *RestorableDatabaseAccountsClient {
	subClient, _ := NewRestorableDatabaseAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableGremlinDatabasesClient() *RestorableGremlinDatabasesClient {
	subClient, _ := NewRestorableGremlinDatabasesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableGremlinGraphsClient() *RestorableGremlinGraphsClient {
	subClient, _ := NewRestorableGremlinGraphsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableGremlinResourcesClient() *RestorableGremlinResourcesClient {
	subClient, _ := NewRestorableGremlinResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableMongodbCollectionsClient() *RestorableMongodbCollectionsClient {
	subClient, _ := NewRestorableMongodbCollectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableMongodbDatabasesClient() *RestorableMongodbDatabasesClient {
	subClient, _ := NewRestorableMongodbDatabasesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableMongodbResourcesClient() *RestorableMongodbResourcesClient {
	subClient, _ := NewRestorableMongodbResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableSQLContainersClient() *RestorableSQLContainersClient {
	subClient, _ := NewRestorableSQLContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableSQLDatabasesClient() *RestorableSQLDatabasesClient {
	subClient, _ := NewRestorableSQLDatabasesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableSQLResourcesClient() *RestorableSQLResourcesClient {
	subClient, _ := NewRestorableSQLResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableTableResourcesClient() *RestorableTableResourcesClient {
	subClient, _ := NewRestorableTableResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestorableTablesClient() *RestorableTablesClient {
	subClient, _ := NewRestorableTablesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSQLResourcesClient() *SQLResourcesClient {
	subClient, _ := NewSQLResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServiceClient() *ServiceClient {
	subClient, _ := NewServiceClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTableResourcesClient() *TableResourcesClient {
	subClient, _ := NewTableResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

