//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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
//   - subscriptionID - Identifier of the subscription
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

// NewFilesClient creates a new instance of FilesClient.
func (c *ClientFactory) NewFilesClient() *FilesClient {
	subClient, _ := NewFilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewProjectsClient creates a new instance of ProjectsClient.
func (c *ClientFactory) NewProjectsClient() *ProjectsClient {
	subClient, _ := NewProjectsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient.
func (c *ClientFactory) NewResourceSKUsClient() *ResourceSKUsClient {
	subClient, _ := NewResourceSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServiceTasksClient creates a new instance of ServiceTasksClient.
func (c *ClientFactory) NewServiceTasksClient() *ServiceTasksClient {
	subClient, _ := NewServiceTasksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewTasksClient creates a new instance of TasksClient.
func (c *ClientFactory) NewTasksClient() *TasksClient {
	subClient, _ := NewTasksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewUsagesClient creates a new instance of UsagesClient.
func (c *ClientFactory) NewUsagesClient() *UsagesClient {
	subClient, _ := NewUsagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
