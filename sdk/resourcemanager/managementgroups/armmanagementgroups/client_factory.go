//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	credential azcore.TokenCredential
	options    *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		credential: credential,
		options:    options.Clone(),
	}, nil
}

// NewAPIClient creates a new instance of APIClient.
func (c *ClientFactory) NewAPIClient() *APIClient {
	subClient, _ := NewAPIClient(c.credential, c.options)
	return subClient
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.credential, c.options)
	return subClient
}

// NewEntitiesClient creates a new instance of EntitiesClient.
func (c *ClientFactory) NewEntitiesClient() *EntitiesClient {
	subClient, _ := NewEntitiesClient(c.credential, c.options)
	return subClient
}

// NewHierarchySettingsClient creates a new instance of HierarchySettingsClient.
func (c *ClientFactory) NewHierarchySettingsClient() *HierarchySettingsClient {
	subClient, _ := NewHierarchySettingsClient(c.credential, c.options)
	return subClient
}

// NewManagementGroupSubscriptionsClient creates a new instance of ManagementGroupSubscriptionsClient.
func (c *ClientFactory) NewManagementGroupSubscriptionsClient() *ManagementGroupSubscriptionsClient {
	subClient, _ := NewManagementGroupSubscriptionsClient(c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}
