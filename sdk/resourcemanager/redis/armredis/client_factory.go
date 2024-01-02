//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

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

// NewAccessPolicyAssignmentClient creates a new instance of AccessPolicyAssignmentClient.
func (c *ClientFactory) NewAccessPolicyAssignmentClient() *AccessPolicyAssignmentClient {
	subClient, _ := NewAccessPolicyAssignmentClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAccessPolicyClient creates a new instance of AccessPolicyClient.
func (c *ClientFactory) NewAccessPolicyClient() *AccessPolicyClient {
	subClient, _ := NewAccessPolicyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAsyncOperationStatusClient creates a new instance of AsyncOperationStatusClient.
func (c *ClientFactory) NewAsyncOperationStatusClient() *AsyncOperationStatusClient {
	subClient, _ := NewAsyncOperationStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewFirewallRulesClient creates a new instance of FirewallRulesClient.
func (c *ClientFactory) NewFirewallRulesClient() *FirewallRulesClient {
	subClient, _ := NewFirewallRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLinkedServerClient creates a new instance of LinkedServerClient.
func (c *ClientFactory) NewLinkedServerClient() *LinkedServerClient {
	subClient, _ := NewLinkedServerClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPatchSchedulesClient creates a new instance of PatchSchedulesClient.
func (c *ClientFactory) NewPatchSchedulesClient() *PatchSchedulesClient {
	subClient, _ := NewPatchSchedulesClient(c.subscriptionID, c.credential, c.options)
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
