// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package avs

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/avs/mgmt/2020-03-20/avs"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ClusterProvisioningState = original.ClusterProvisioningState

const (
	Cancelled ClusterProvisioningState = original.Cancelled
	Deleting  ClusterProvisioningState = original.Deleting
	Failed    ClusterProvisioningState = original.Failed
	Succeeded ClusterProvisioningState = original.Succeeded
	Updating  ClusterProvisioningState = original.Updating
)

type ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningState

const (
	ExpressRouteAuthorizationProvisioningStateFailed    ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningStateFailed
	ExpressRouteAuthorizationProvisioningStateSucceeded ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningStateSucceeded
	ExpressRouteAuthorizationProvisioningStateUpdating  ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningStateUpdating
)

type HcxEnterpriseSiteStatus = original.HcxEnterpriseSiteStatus

const (
	Available   HcxEnterpriseSiteStatus = original.Available
	Consumed    HcxEnterpriseSiteStatus = original.Consumed
	Deactivated HcxEnterpriseSiteStatus = original.Deactivated
	Deleted     HcxEnterpriseSiteStatus = original.Deleted
)

type InternetEnum = original.InternetEnum

const (
	Disabled InternetEnum = original.Disabled
	Enabled  InternetEnum = original.Enabled
)

type PrivateCloudProvisioningState = original.PrivateCloudProvisioningState

const (
	PrivateCloudProvisioningStateBuilding  PrivateCloudProvisioningState = original.PrivateCloudProvisioningStateBuilding
	PrivateCloudProvisioningStateCancelled PrivateCloudProvisioningState = original.PrivateCloudProvisioningStateCancelled
	PrivateCloudProvisioningStateDeleting  PrivateCloudProvisioningState = original.PrivateCloudProvisioningStateDeleting
	PrivateCloudProvisioningStateFailed    PrivateCloudProvisioningState = original.PrivateCloudProvisioningStateFailed
	PrivateCloudProvisioningStatePending   PrivateCloudProvisioningState = original.PrivateCloudProvisioningStatePending
	PrivateCloudProvisioningStateSucceeded PrivateCloudProvisioningState = original.PrivateCloudProvisioningStateSucceeded
	PrivateCloudProvisioningStateUpdating  PrivateCloudProvisioningState = original.PrivateCloudProvisioningStateUpdating
)

type QuotaEnabled = original.QuotaEnabled

const (
	QuotaEnabledDisabled QuotaEnabled = original.QuotaEnabledDisabled
	QuotaEnabledEnabled  QuotaEnabled = original.QuotaEnabledEnabled
)

type SslEnum = original.SslEnum

const (
	SslEnumDisabled SslEnum = original.SslEnumDisabled
	SslEnumEnabled  SslEnum = original.SslEnumEnabled
)

type TrialStatus = original.TrialStatus

const (
	TrialAvailable TrialStatus = original.TrialAvailable
	TrialDisabled  TrialStatus = original.TrialDisabled
	TrialUsed      TrialStatus = original.TrialUsed
)

type AdminCredentials = original.AdminCredentials
type AuthorizationsClient = original.AuthorizationsClient
type AuthorizationsCreateOrUpdateFuture = original.AuthorizationsCreateOrUpdateFuture
type AuthorizationsDeleteFuture = original.AuthorizationsDeleteFuture
type BaseClient = original.BaseClient
type Circuit = original.Circuit
type CloudError = original.CloudError
type Cluster = original.Cluster
type ClusterList = original.ClusterList
type ClusterListIterator = original.ClusterListIterator
type ClusterListPage = original.ClusterListPage
type ClusterProperties = original.ClusterProperties
type ClusterUpdate = original.ClusterUpdate
type ClusterUpdateProperties = original.ClusterUpdateProperties
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type Endpoints = original.Endpoints
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ExpressRouteAuthorization = original.ExpressRouteAuthorization
type ExpressRouteAuthorizationList = original.ExpressRouteAuthorizationList
type ExpressRouteAuthorizationListIterator = original.ExpressRouteAuthorizationListIterator
type ExpressRouteAuthorizationListPage = original.ExpressRouteAuthorizationListPage
type ExpressRouteAuthorizationProperties = original.ExpressRouteAuthorizationProperties
type HcxEnterpriseSite = original.HcxEnterpriseSite
type HcxEnterpriseSiteList = original.HcxEnterpriseSiteList
type HcxEnterpriseSiteListIterator = original.HcxEnterpriseSiteListIterator
type HcxEnterpriseSiteListPage = original.HcxEnterpriseSiteListPage
type HcxEnterpriseSiteProperties = original.HcxEnterpriseSiteProperties
type HcxEnterpriseSitesClient = original.HcxEnterpriseSitesClient
type IdentitySource = original.IdentitySource
type LocationsClient = original.LocationsClient
type ManagementCluster = original.ManagementCluster
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type PrivateCloud = original.PrivateCloud
type PrivateCloudList = original.PrivateCloudList
type PrivateCloudListIterator = original.PrivateCloudListIterator
type PrivateCloudListPage = original.PrivateCloudListPage
type PrivateCloudProperties = original.PrivateCloudProperties
type PrivateCloudUpdate = original.PrivateCloudUpdate
type PrivateCloudUpdateProperties = original.PrivateCloudUpdateProperties
type PrivateCloudsClient = original.PrivateCloudsClient
type PrivateCloudsCreateOrUpdateFuture = original.PrivateCloudsCreateOrUpdateFuture
type PrivateCloudsDeleteFuture = original.PrivateCloudsDeleteFuture
type PrivateCloudsUpdateFuture = original.PrivateCloudsUpdateFuture
type Quota = original.Quota
type Resource = original.Resource
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type Trial = original.Trial

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAuthorizationsClient(subscriptionID string) AuthorizationsClient {
	return original.NewAuthorizationsClient(subscriptionID)
}
func NewAuthorizationsClientWithBaseURI(baseURI string, subscriptionID string) AuthorizationsClient {
	return original.NewAuthorizationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterListIterator(page ClusterListPage) ClusterListIterator {
	return original.NewClusterListIterator(page)
}
func NewClusterListPage(cur ClusterList, getNextPage func(context.Context, ClusterList) (ClusterList, error)) ClusterListPage {
	return original.NewClusterListPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewExpressRouteAuthorizationListIterator(page ExpressRouteAuthorizationListPage) ExpressRouteAuthorizationListIterator {
	return original.NewExpressRouteAuthorizationListIterator(page)
}
func NewExpressRouteAuthorizationListPage(cur ExpressRouteAuthorizationList, getNextPage func(context.Context, ExpressRouteAuthorizationList) (ExpressRouteAuthorizationList, error)) ExpressRouteAuthorizationListPage {
	return original.NewExpressRouteAuthorizationListPage(cur, getNextPage)
}
func NewHcxEnterpriseSiteListIterator(page HcxEnterpriseSiteListPage) HcxEnterpriseSiteListIterator {
	return original.NewHcxEnterpriseSiteListIterator(page)
}
func NewHcxEnterpriseSiteListPage(cur HcxEnterpriseSiteList, getNextPage func(context.Context, HcxEnterpriseSiteList) (HcxEnterpriseSiteList, error)) HcxEnterpriseSiteListPage {
	return original.NewHcxEnterpriseSiteListPage(cur, getNextPage)
}
func NewHcxEnterpriseSitesClient(subscriptionID string) HcxEnterpriseSitesClient {
	return original.NewHcxEnterpriseSitesClient(subscriptionID)
}
func NewHcxEnterpriseSitesClientWithBaseURI(baseURI string, subscriptionID string) HcxEnterpriseSitesClient {
	return original.NewHcxEnterpriseSitesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateCloudListIterator(page PrivateCloudListPage) PrivateCloudListIterator {
	return original.NewPrivateCloudListIterator(page)
}
func NewPrivateCloudsClient(subscriptionID string) PrivateCloudsClient {
	return original.NewPrivateCloudsClient(subscriptionID)
}
func NewPrivateCloudsClientWithBaseURI(baseURI string, subscriptionID string) PrivateCloudsClient {
	return original.NewPrivateCloudsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return original.PossibleClusterProvisioningStateValues()
}
func PossibleExpressRouteAuthorizationProvisioningStateValues() []ExpressRouteAuthorizationProvisioningState {
	return original.PossibleExpressRouteAuthorizationProvisioningStateValues()
}
func PossibleHcxEnterpriseSiteStatusValues() []HcxEnterpriseSiteStatus {
	return original.PossibleHcxEnterpriseSiteStatusValues()
}
func PossibleInternetEnumValues() []InternetEnum {
	return original.PossibleInternetEnumValues()
}
func PossiblePrivateCloudProvisioningStateValues() []PrivateCloudProvisioningState {
	return original.PossiblePrivateCloudProvisioningStateValues()
}
func PossibleQuotaEnabledValues() []QuotaEnabled {
	return original.PossibleQuotaEnabledValues()
}
func PossibleSslEnumValues() []SslEnum {
	return original.PossibleSslEnumValues()
}
func PossibleTrialStatusValues() []TrialStatus {
	return original.PossibleTrialStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
