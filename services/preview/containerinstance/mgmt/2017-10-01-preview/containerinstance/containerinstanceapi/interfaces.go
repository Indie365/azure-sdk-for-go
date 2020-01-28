package containerinstanceapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/containerinstance/mgmt/2017-10-01-preview/containerinstance"
)

// ContainerGroupsClientAPI contains the set of methods on the ContainerGroupsClient type.
type ContainerGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup containerinstance.ContainerGroup) (result containerinstance.ContainerGroup, err error)
	Delete(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroup, err error)
	Get(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroup, err error)
	List(ctx context.Context) (result containerinstance.ContainerGroupListResultPage, err error)
	ListComplete(ctx context.Context) (result containerinstance.ContainerGroupListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerinstance.ContainerGroupListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerinstance.ContainerGroupListResultIterator, err error)
}

var _ ContainerGroupsClientAPI = (*containerinstance.ContainerGroupsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerinstance.OperationListResult, err error)
}

var _ OperationsClientAPI = (*containerinstance.OperationsClient)(nil)

// ContainerLogsClientAPI contains the set of methods on the ContainerLogsClient type.
type ContainerLogsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32) (result containerinstance.Logs, err error)
}

var _ ContainerLogsClientAPI = (*containerinstance.ContainerLogsClient)(nil)
