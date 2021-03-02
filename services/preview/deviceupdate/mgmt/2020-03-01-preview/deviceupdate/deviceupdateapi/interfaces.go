package deviceupdateapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/deviceupdate/mgmt/2020-03-01-preview/deviceupdate"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, account deviceupdate.Account) (result deviceupdate.AccountsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result deviceupdate.AccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result deviceupdate.Account, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result deviceupdate.AccountListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result deviceupdate.AccountListIterator, err error)
	ListBySubscription(ctx context.Context) (result deviceupdate.AccountListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result deviceupdate.AccountListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, accountUpdatePayload deviceupdate.AccountUpdate) (result deviceupdate.AccountsUpdateFuture, err error)
}

var _ AccountsClientAPI = (*deviceupdate.AccountsClient)(nil)

// InstancesClientAPI contains the set of methods on the InstancesClient type.
type InstancesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, accountName string, instanceName string, instance deviceupdate.Instance) (result deviceupdate.InstancesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, instanceName string) (result deviceupdate.InstancesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, instanceName string) (result deviceupdate.Instance, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result deviceupdate.InstanceListPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result deviceupdate.InstanceListIterator, err error)
	ListBySubscription(ctx context.Context, accountName string) (result deviceupdate.InstanceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, accountName string) (result deviceupdate.InstanceListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, instanceName string, tagUpdatePayload deviceupdate.TagUpdate) (result deviceupdate.Instance, err error)
}

var _ InstancesClientAPI = (*deviceupdate.InstancesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result deviceupdate.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result deviceupdate.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*deviceupdate.OperationsClient)(nil)
