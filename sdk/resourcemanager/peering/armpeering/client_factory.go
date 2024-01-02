//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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
//   - subscriptionID - The Azure subscription ID.
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

// NewCdnPeeringPrefixesClient creates a new instance of CdnPeeringPrefixesClient.
func (c *ClientFactory) NewCdnPeeringPrefixesClient() *CdnPeeringPrefixesClient {
	subClient, _ := NewCdnPeeringPrefixesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConnectionMonitorTestsClient creates a new instance of ConnectionMonitorTestsClient.
func (c *ClientFactory) NewConnectionMonitorTestsClient() *ConnectionMonitorTestsClient {
	subClient, _ := NewConnectionMonitorTestsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLegacyPeeringsClient creates a new instance of LegacyPeeringsClient.
func (c *ClientFactory) NewLegacyPeeringsClient() *LegacyPeeringsClient {
	subClient, _ := NewLegacyPeeringsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLocationsClient creates a new instance of LocationsClient.
func (c *ClientFactory) NewLocationsClient() *LocationsClient {
	subClient, _ := NewLocationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLookingGlassClient creates a new instance of LookingGlassClient.
func (c *ClientFactory) NewLookingGlassClient() *LookingGlassClient {
	subClient, _ := NewLookingGlassClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewManagementClient creates a new instance of ManagementClient.
func (c *ClientFactory) NewManagementClient() *ManagementClient {
	subClient, _ := NewManagementClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPeerAsnsClient creates a new instance of PeerAsnsClient.
func (c *ClientFactory) NewPeerAsnsClient() *PeerAsnsClient {
	subClient, _ := NewPeerAsnsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPeeringsClient creates a new instance of PeeringsClient.
func (c *ClientFactory) NewPeeringsClient() *PeeringsClient {
	subClient, _ := NewPeeringsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrefixesClient creates a new instance of PrefixesClient.
func (c *ClientFactory) NewPrefixesClient() *PrefixesClient {
	subClient, _ := NewPrefixesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewReceivedRoutesClient creates a new instance of ReceivedRoutesClient.
func (c *ClientFactory) NewReceivedRoutesClient() *ReceivedRoutesClient {
	subClient, _ := NewReceivedRoutesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewRegisteredAsnsClient creates a new instance of RegisteredAsnsClient.
func (c *ClientFactory) NewRegisteredAsnsClient() *RegisteredAsnsClient {
	subClient, _ := NewRegisteredAsnsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewRegisteredPrefixesClient creates a new instance of RegisteredPrefixesClient.
func (c *ClientFactory) NewRegisteredPrefixesClient() *RegisteredPrefixesClient {
	subClient, _ := NewRegisteredPrefixesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServiceCountriesClient creates a new instance of ServiceCountriesClient.
func (c *ClientFactory) NewServiceCountriesClient() *ServiceCountriesClient {
	subClient, _ := NewServiceCountriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServiceLocationsClient creates a new instance of ServiceLocationsClient.
func (c *ClientFactory) NewServiceLocationsClient() *ServiceLocationsClient {
	subClient, _ := NewServiceLocationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServiceProvidersClient creates a new instance of ServiceProvidersClient.
func (c *ClientFactory) NewServiceProvidersClient() *ServiceProvidersClient {
	subClient, _ := NewServiceProvidersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
