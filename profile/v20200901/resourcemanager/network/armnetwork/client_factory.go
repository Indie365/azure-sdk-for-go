//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

import (
	"github.com/Azure/azure-sdk-for-go/profile/v20200901/internal"
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
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(internal.ModuleName+".ClientFactory", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewVirtualNetworkGatewaysClient() *VirtualNetworkGatewaysClient {
	subClient, _ := NewVirtualNetworkGatewaysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualNetworkGatewayConnectionsClient() *VirtualNetworkGatewayConnectionsClient {
	subClient, _ := NewVirtualNetworkGatewayConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLocalNetworkGatewaysClient() *LocalNetworkGatewaysClient {
	subClient, _ := NewLocalNetworkGatewaysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoadBalancersClient() *LoadBalancersClient {
	subClient, _ := NewLoadBalancersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoadBalancerBackendAddressPoolsClient() *LoadBalancerBackendAddressPoolsClient {
	subClient, _ := NewLoadBalancerBackendAddressPoolsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoadBalancerFrontendIPConfigurationsClient() *LoadBalancerFrontendIPConfigurationsClient {
	subClient, _ := NewLoadBalancerFrontendIPConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInboundNatRulesClient() *InboundNatRulesClient {
	subClient, _ := NewInboundNatRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoadBalancerLoadBalancingRulesClient() *LoadBalancerLoadBalancingRulesClient {
	subClient, _ := NewLoadBalancerLoadBalancingRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoadBalancerOutboundRulesClient() *LoadBalancerOutboundRulesClient {
	subClient, _ := NewLoadBalancerOutboundRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoadBalancerNetworkInterfacesClient() *LoadBalancerNetworkInterfacesClient {
	subClient, _ := NewLoadBalancerNetworkInterfacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoadBalancerProbesClient() *LoadBalancerProbesClient {
	subClient, _ := NewLoadBalancerProbesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInterfacesClient() *InterfacesClient {
	subClient, _ := NewInterfacesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInterfaceIPConfigurationsClient() *InterfaceIPConfigurationsClient {
	subClient, _ := NewInterfaceIPConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInterfaceLoadBalancersClient() *InterfaceLoadBalancersClient {
	subClient, _ := NewInterfaceLoadBalancersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInterfaceTapConfigurationsClient() *InterfaceTapConfigurationsClient {
	subClient, _ := NewInterfaceTapConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecurityGroupsClient() *SecurityGroupsClient {
	subClient, _ := NewSecurityGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecurityRulesClient() *SecurityRulesClient {
	subClient, _ := NewSecurityRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDefaultSecurityRulesClient() *DefaultSecurityRulesClient {
	subClient, _ := NewDefaultSecurityRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPublicIPAddressesClient() *PublicIPAddressesClient {
	subClient, _ := NewPublicIPAddressesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRouteTablesClient() *RouteTablesClient {
	subClient, _ := NewRouteTablesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRoutesClient() *RoutesClient {
	subClient, _ := NewRoutesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualNetworksClient() *VirtualNetworksClient {
	subClient, _ := NewVirtualNetworksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSubnetsClient() *SubnetsClient {
	subClient, _ := NewSubnetsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVirtualNetworkPeeringsClient() *VirtualNetworkPeeringsClient {
	subClient, _ := NewVirtualNetworkPeeringsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
