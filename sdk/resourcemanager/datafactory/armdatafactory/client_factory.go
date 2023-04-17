//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFactoriesClient() *FactoriesClient {
	subClient, _ := NewFactoriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExposureControlClient() *ExposureControlClient {
	subClient, _ := NewExposureControlClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationRuntimesClient() *IntegrationRuntimesClient {
	subClient, _ := NewIntegrationRuntimesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationRuntimeObjectMetadataClient() *IntegrationRuntimeObjectMetadataClient {
	subClient, _ := NewIntegrationRuntimeObjectMetadataClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIntegrationRuntimeNodesClient() *IntegrationRuntimeNodesClient {
	subClient, _ := NewIntegrationRuntimeNodesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLinkedServicesClient() *LinkedServicesClient {
	subClient, _ := NewLinkedServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatasetsClient() *DatasetsClient {
	subClient, _ := NewDatasetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPipelinesClient() *PipelinesClient {
	subClient, _ := NewPipelinesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPipelineRunsClient() *PipelineRunsClient {
	subClient, _ := NewPipelineRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewActivityRunsClient() *ActivityRunsClient {
	subClient, _ := NewActivityRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTriggersClient() *TriggersClient {
	subClient, _ := NewTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTriggerRunsClient() *TriggerRunsClient {
	subClient, _ := NewTriggerRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataFlowsClient() *DataFlowsClient {
	subClient, _ := NewDataFlowsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataFlowDebugSessionClient() *DataFlowDebugSessionClient {
	subClient, _ := NewDataFlowDebugSessionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedVirtualNetworksClient() *ManagedVirtualNetworksClient {
	subClient, _ := NewManagedVirtualNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagedPrivateEndpointsClient() *ManagedPrivateEndpointsClient {
	subClient, _ := NewManagedPrivateEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCredentialOperationsClient() *CredentialOperationsClient {
	subClient, _ := NewCredentialOperationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndPointConnectionsClient() *PrivateEndPointConnectionsClient {
	subClient, _ := NewPrivateEndPointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionClient() *PrivateEndpointConnectionClient {
	subClient, _ := NewPrivateEndpointConnectionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGlobalParametersClient() *GlobalParametersClient {
	subClient, _ := NewGlobalParametersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
