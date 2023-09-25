//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

func (c *ClientFactory) NewAddonsClient() *AddonsClient {
	subClient, _ := NewAddonsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAuthorizationsClient() *AuthorizationsClient {
	subClient, _ := NewAuthorizationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCloudLinksClient() *CloudLinksClient {
	subClient, _ := NewCloudLinksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewClustersClient() *ClustersClient {
	subClient, _ := NewClustersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDatastoresClient() *DatastoresClient {
	subClient, _ := NewDatastoresClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGlobalReachConnectionsClient() *GlobalReachConnectionsClient {
	subClient, _ := NewGlobalReachConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHcxEnterpriseSitesClient() *HcxEnterpriseSitesClient {
	subClient, _ := NewHcxEnterpriseSitesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLocationsClient() *LocationsClient {
	subClient, _ := NewLocationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPlacementPoliciesClient() *PlacementPoliciesClient {
	subClient, _ := NewPlacementPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateCloudsClient() *PrivateCloudsClient {
	subClient, _ := NewPrivateCloudsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewScriptCmdletsClient() *ScriptCmdletsClient {
	subClient, _ := NewScriptCmdletsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewScriptExecutionsClient() *ScriptExecutionsClient {
	subClient, _ := NewScriptExecutionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewScriptPackagesClient() *ScriptPackagesClient {
	subClient, _ := NewScriptPackagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualMachinesClient() *VirtualMachinesClient {
	subClient, _ := NewVirtualMachinesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkloadNetworksClient() *WorkloadNetworksClient {
	subClient, _ := NewWorkloadNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

