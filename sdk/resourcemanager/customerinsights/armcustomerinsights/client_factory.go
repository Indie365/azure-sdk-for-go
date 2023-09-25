//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

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
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
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

func (c *ClientFactory) NewAuthorizationPoliciesClient() *AuthorizationPoliciesClient {
	subClient, _ := NewAuthorizationPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectorMappingsClient() *ConnectorMappingsClient {
	subClient, _ := NewConnectorMappingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewConnectorsClient() *ConnectorsClient {
	subClient, _ := NewConnectorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHubsClient() *HubsClient {
	subClient, _ := NewHubsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewImagesClient() *ImagesClient {
	subClient, _ := NewImagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewInteractionsClient() *InteractionsClient {
	subClient, _ := NewInteractionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewKpiClient() *KpiClient {
	subClient, _ := NewKpiClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLinksClient() *LinksClient {
	subClient, _ := NewLinksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPredictionsClient() *PredictionsClient {
	subClient, _ := NewPredictionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProfilesClient() *ProfilesClient {
	subClient, _ := NewProfilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRelationshipLinksClient() *RelationshipLinksClient {
	subClient, _ := NewRelationshipLinksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRelationshipsClient() *RelationshipsClient {
	subClient, _ := NewRelationshipsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRoleAssignmentsClient() *RoleAssignmentsClient {
	subClient, _ := NewRoleAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRolesClient() *RolesClient {
	subClient, _ := NewRolesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewViewsClient() *ViewsClient {
	subClient, _ := NewViewsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWidgetTypesClient() *WidgetTypesClient {
	subClient, _ := NewWidgetTypesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

