//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

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
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewDicomServicesClient creates a new instance of DicomServicesClient.
func (c *ClientFactory) NewDicomServicesClient() *DicomServicesClient {
	subClient, _ := NewDicomServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewFhirDestinationsClient creates a new instance of FhirDestinationsClient.
func (c *ClientFactory) NewFhirDestinationsClient() *FhirDestinationsClient {
	subClient, _ := NewFhirDestinationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewFhirServicesClient creates a new instance of FhirServicesClient.
func (c *ClientFactory) NewFhirServicesClient() *FhirServicesClient {
	subClient, _ := NewFhirServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewIotConnectorFhirDestinationClient creates a new instance of IotConnectorFhirDestinationClient.
func (c *ClientFactory) NewIotConnectorFhirDestinationClient() *IotConnectorFhirDestinationClient {
	subClient, _ := NewIotConnectorFhirDestinationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewIotConnectorsClient creates a new instance of IotConnectorsClient.
func (c *ClientFactory) NewIotConnectorsClient() *IotConnectorsClient {
	subClient, _ := NewIotConnectorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationResultsClient creates a new instance of OperationResultsClient.
func (c *ClientFactory) NewOperationResultsClient() *OperationResultsClient {
	subClient, _ := NewOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkspacePrivateEndpointConnectionsClient creates a new instance of WorkspacePrivateEndpointConnectionsClient.
func (c *ClientFactory) NewWorkspacePrivateEndpointConnectionsClient() *WorkspacePrivateEndpointConnectionsClient {
	subClient, _ := NewWorkspacePrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkspacePrivateLinkResourcesClient creates a new instance of WorkspacePrivateLinkResourcesClient.
func (c *ClientFactory) NewWorkspacePrivateLinkResourcesClient() *WorkspacePrivateLinkResourcesClient {
	subClient, _ := NewWorkspacePrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkspacesClient creates a new instance of WorkspacesClient.
func (c *ClientFactory) NewWorkspacesClient() *WorkspacesClient {
	subClient, _ := NewWorkspacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
