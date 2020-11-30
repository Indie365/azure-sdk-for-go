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

	original "github.com/Azure/azure-sdk-for-go/services/preview/avs/mgmt/2020-07-17-preview/avs"
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

type DhcpType = original.DhcpType

const (
	DhcpTypeRELAY                     DhcpType = original.DhcpTypeRELAY
	DhcpTypeSERVER                    DhcpType = original.DhcpTypeSERVER
	DhcpTypeWorkloadNetworkDhcpEntity DhcpType = original.DhcpTypeWorkloadNetworkDhcpEntity
)

type ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningState

const (
	ExpressRouteAuthorizationProvisioningStateFailed    ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningStateFailed
	ExpressRouteAuthorizationProvisioningStateSucceeded ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningStateSucceeded
	ExpressRouteAuthorizationProvisioningStateUpdating  ExpressRouteAuthorizationProvisioningState = original.ExpressRouteAuthorizationProvisioningStateUpdating
)

type GlobalReachConnectionProvisioningState = original.GlobalReachConnectionProvisioningState

const (
	GlobalReachConnectionProvisioningStateFailed    GlobalReachConnectionProvisioningState = original.GlobalReachConnectionProvisioningStateFailed
	GlobalReachConnectionProvisioningStateSucceeded GlobalReachConnectionProvisioningState = original.GlobalReachConnectionProvisioningStateSucceeded
	GlobalReachConnectionProvisioningStateUpdating  GlobalReachConnectionProvisioningState = original.GlobalReachConnectionProvisioningStateUpdating
)

type GlobalReachConnectionStatus = original.GlobalReachConnectionStatus

const (
	Connected    GlobalReachConnectionStatus = original.Connected
	Connecting   GlobalReachConnectionStatus = original.Connecting
	Disconnected GlobalReachConnectionStatus = original.Disconnected
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

type NsxtAdminRotateEnum = original.NsxtAdminRotateEnum

const (
	OnetimeRotate NsxtAdminRotateEnum = original.OnetimeRotate
)

type PortMirroringDirectionEnum = original.PortMirroringDirectionEnum

const (
	INGRESSEGRESSBIDIRECTIONAL PortMirroringDirectionEnum = original.INGRESSEGRESSBIDIRECTIONAL
)

type PortMirroringStatusEnum = original.PortMirroringStatusEnum

const (
	SUCCESSFAILURE PortMirroringStatusEnum = original.SUCCESSFAILURE
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

type SegmentStatusEnum = original.SegmentStatusEnum

const (
	SegmentStatusEnumSUCCESSFAILURE SegmentStatusEnum = original.SegmentStatusEnumSUCCESSFAILURE
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

type VMGroupStatusEnum = original.VMGroupStatusEnum

const (
	VMGroupStatusEnumSUCCESSFAILURE VMGroupStatusEnum = original.VMGroupStatusEnumSUCCESSFAILURE
)

type VMTypeEnum = original.VMTypeEnum

const (
	REGULAREDGESERVICE VMTypeEnum = original.REGULAREDGESERVICE
)

type VcsaAdminRotateEnum = original.VcsaAdminRotateEnum

const (
	VcsaAdminRotateEnumOnetimeRotate VcsaAdminRotateEnum = original.VcsaAdminRotateEnumOnetimeRotate
)

type WorkloadNetworkDhcpProvisioningState = original.WorkloadNetworkDhcpProvisioningState

const (
	WorkloadNetworkDhcpProvisioningStateBuilding  WorkloadNetworkDhcpProvisioningState = original.WorkloadNetworkDhcpProvisioningStateBuilding
	WorkloadNetworkDhcpProvisioningStateDeleting  WorkloadNetworkDhcpProvisioningState = original.WorkloadNetworkDhcpProvisioningStateDeleting
	WorkloadNetworkDhcpProvisioningStateFailed    WorkloadNetworkDhcpProvisioningState = original.WorkloadNetworkDhcpProvisioningStateFailed
	WorkloadNetworkDhcpProvisioningStateSucceeded WorkloadNetworkDhcpProvisioningState = original.WorkloadNetworkDhcpProvisioningStateSucceeded
	WorkloadNetworkDhcpProvisioningStateUpdating  WorkloadNetworkDhcpProvisioningState = original.WorkloadNetworkDhcpProvisioningStateUpdating
)

type WorkloadNetworkPortMirroringProvisioningState = original.WorkloadNetworkPortMirroringProvisioningState

const (
	WorkloadNetworkPortMirroringProvisioningStateBuilding  WorkloadNetworkPortMirroringProvisioningState = original.WorkloadNetworkPortMirroringProvisioningStateBuilding
	WorkloadNetworkPortMirroringProvisioningStateDeleting  WorkloadNetworkPortMirroringProvisioningState = original.WorkloadNetworkPortMirroringProvisioningStateDeleting
	WorkloadNetworkPortMirroringProvisioningStateFailed    WorkloadNetworkPortMirroringProvisioningState = original.WorkloadNetworkPortMirroringProvisioningStateFailed
	WorkloadNetworkPortMirroringProvisioningStateSucceeded WorkloadNetworkPortMirroringProvisioningState = original.WorkloadNetworkPortMirroringProvisioningStateSucceeded
	WorkloadNetworkPortMirroringProvisioningStateUpdating  WorkloadNetworkPortMirroringProvisioningState = original.WorkloadNetworkPortMirroringProvisioningStateUpdating
)

type WorkloadNetworkSegmentProvisioningState = original.WorkloadNetworkSegmentProvisioningState

const (
	WorkloadNetworkSegmentProvisioningStateBuilding  WorkloadNetworkSegmentProvisioningState = original.WorkloadNetworkSegmentProvisioningStateBuilding
	WorkloadNetworkSegmentProvisioningStateDeleting  WorkloadNetworkSegmentProvisioningState = original.WorkloadNetworkSegmentProvisioningStateDeleting
	WorkloadNetworkSegmentProvisioningStateFailed    WorkloadNetworkSegmentProvisioningState = original.WorkloadNetworkSegmentProvisioningStateFailed
	WorkloadNetworkSegmentProvisioningStateSucceeded WorkloadNetworkSegmentProvisioningState = original.WorkloadNetworkSegmentProvisioningStateSucceeded
	WorkloadNetworkSegmentProvisioningStateUpdating  WorkloadNetworkSegmentProvisioningState = original.WorkloadNetworkSegmentProvisioningStateUpdating
)

type WorkloadNetworkVMGroupProvisioningState = original.WorkloadNetworkVMGroupProvisioningState

const (
	WorkloadNetworkVMGroupProvisioningStateBuilding  WorkloadNetworkVMGroupProvisioningState = original.WorkloadNetworkVMGroupProvisioningStateBuilding
	WorkloadNetworkVMGroupProvisioningStateDeleting  WorkloadNetworkVMGroupProvisioningState = original.WorkloadNetworkVMGroupProvisioningStateDeleting
	WorkloadNetworkVMGroupProvisioningStateFailed    WorkloadNetworkVMGroupProvisioningState = original.WorkloadNetworkVMGroupProvisioningStateFailed
	WorkloadNetworkVMGroupProvisioningStateSucceeded WorkloadNetworkVMGroupProvisioningState = original.WorkloadNetworkVMGroupProvisioningStateSucceeded
	WorkloadNetworkVMGroupProvisioningStateUpdating  WorkloadNetworkVMGroupProvisioningState = original.WorkloadNetworkVMGroupProvisioningStateUpdating
)

type AdminCredentials = original.AdminCredentials
type AuthorizationsClient = original.AuthorizationsClient
type AuthorizationsCreateOrUpdateFuture = original.AuthorizationsCreateOrUpdateFuture
type AuthorizationsDeleteFuture = original.AuthorizationsDeleteFuture
type BaseClient = original.BaseClient
type BasicWorkloadNetworkDhcpEntity = original.BasicWorkloadNetworkDhcpEntity
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
type CommonClusterProperties = original.CommonClusterProperties
type Endpoints = original.Endpoints
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorResponse = original.ErrorResponse
type ExpressRouteAuthorization = original.ExpressRouteAuthorization
type ExpressRouteAuthorizationList = original.ExpressRouteAuthorizationList
type ExpressRouteAuthorizationListIterator = original.ExpressRouteAuthorizationListIterator
type ExpressRouteAuthorizationListPage = original.ExpressRouteAuthorizationListPage
type ExpressRouteAuthorizationProperties = original.ExpressRouteAuthorizationProperties
type GlobalReachConnection = original.GlobalReachConnection
type GlobalReachConnectionList = original.GlobalReachConnectionList
type GlobalReachConnectionListIterator = original.GlobalReachConnectionListIterator
type GlobalReachConnectionListPage = original.GlobalReachConnectionListPage
type GlobalReachConnectionProperties = original.GlobalReachConnectionProperties
type GlobalReachConnectionsClient = original.GlobalReachConnectionsClient
type GlobalReachConnectionsCreateOrUpdateFuture = original.GlobalReachConnectionsCreateOrUpdateFuture
type GlobalReachConnectionsDeleteFuture = original.GlobalReachConnectionsDeleteFuture
type HcxEnterpriseSite = original.HcxEnterpriseSite
type HcxEnterpriseSiteList = original.HcxEnterpriseSiteList
type HcxEnterpriseSiteListIterator = original.HcxEnterpriseSiteListIterator
type HcxEnterpriseSiteListPage = original.HcxEnterpriseSiteListPage
type HcxEnterpriseSiteProperties = original.HcxEnterpriseSiteProperties
type HcxEnterpriseSitesClient = original.HcxEnterpriseSitesClient
type IdentitySource = original.IdentitySource
type LocationsClient = original.LocationsClient
type LogSpecification = original.LogSpecification
type ManagementCluster = original.ManagementCluster
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationProperties = original.OperationProperties
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
type ProxyResource = original.ProxyResource
type Quota = original.Quota
type Resource = original.Resource
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type TrackedResource = original.TrackedResource
type Trial = original.Trial
type WorkloadNetworkDhcp = original.WorkloadNetworkDhcp
type WorkloadNetworkDhcpEntity = original.WorkloadNetworkDhcpEntity
type WorkloadNetworkDhcpList = original.WorkloadNetworkDhcpList
type WorkloadNetworkDhcpListIterator = original.WorkloadNetworkDhcpListIterator
type WorkloadNetworkDhcpListPage = original.WorkloadNetworkDhcpListPage
type WorkloadNetworkDhcpRelay = original.WorkloadNetworkDhcpRelay
type WorkloadNetworkDhcpServer = original.WorkloadNetworkDhcpServer
type WorkloadNetworkGateway = original.WorkloadNetworkGateway
type WorkloadNetworkGatewayList = original.WorkloadNetworkGatewayList
type WorkloadNetworkGatewayListIterator = original.WorkloadNetworkGatewayListIterator
type WorkloadNetworkGatewayListPage = original.WorkloadNetworkGatewayListPage
type WorkloadNetworkGatewayProperties = original.WorkloadNetworkGatewayProperties
type WorkloadNetworkPortMirroring = original.WorkloadNetworkPortMirroring
type WorkloadNetworkPortMirroringList = original.WorkloadNetworkPortMirroringList
type WorkloadNetworkPortMirroringListIterator = original.WorkloadNetworkPortMirroringListIterator
type WorkloadNetworkPortMirroringListPage = original.WorkloadNetworkPortMirroringListPage
type WorkloadNetworkPortMirroringProperties = original.WorkloadNetworkPortMirroringProperties
type WorkloadNetworkSegment = original.WorkloadNetworkSegment
type WorkloadNetworkSegmentPortVif = original.WorkloadNetworkSegmentPortVif
type WorkloadNetworkSegmentProperties = original.WorkloadNetworkSegmentProperties
type WorkloadNetworkSegmentSubnet = original.WorkloadNetworkSegmentSubnet
type WorkloadNetworkSegmentsList = original.WorkloadNetworkSegmentsList
type WorkloadNetworkSegmentsListIterator = original.WorkloadNetworkSegmentsListIterator
type WorkloadNetworkSegmentsListPage = original.WorkloadNetworkSegmentsListPage
type WorkloadNetworkVMGroup = original.WorkloadNetworkVMGroup
type WorkloadNetworkVMGroupProperties = original.WorkloadNetworkVMGroupProperties
type WorkloadNetworkVMGroupsList = original.WorkloadNetworkVMGroupsList
type WorkloadNetworkVMGroupsListIterator = original.WorkloadNetworkVMGroupsListIterator
type WorkloadNetworkVMGroupsListPage = original.WorkloadNetworkVMGroupsListPage
type WorkloadNetworkVirtualMachine = original.WorkloadNetworkVirtualMachine
type WorkloadNetworkVirtualMachineProperties = original.WorkloadNetworkVirtualMachineProperties
type WorkloadNetworkVirtualMachinesList = original.WorkloadNetworkVirtualMachinesList
type WorkloadNetworkVirtualMachinesListIterator = original.WorkloadNetworkVirtualMachinesListIterator
type WorkloadNetworkVirtualMachinesListPage = original.WorkloadNetworkVirtualMachinesListPage
type WorkloadNetworksClient = original.WorkloadNetworksClient
type WorkloadNetworksCreateDhcpFuture = original.WorkloadNetworksCreateDhcpFuture
type WorkloadNetworksCreatePortMirroringFuture = original.WorkloadNetworksCreatePortMirroringFuture
type WorkloadNetworksCreateSegmentsFuture = original.WorkloadNetworksCreateSegmentsFuture
type WorkloadNetworksCreateVMGroupFuture = original.WorkloadNetworksCreateVMGroupFuture
type WorkloadNetworksDeleteDhcpFuture = original.WorkloadNetworksDeleteDhcpFuture
type WorkloadNetworksDeletePortMirroringFuture = original.WorkloadNetworksDeletePortMirroringFuture
type WorkloadNetworksDeleteSegmentFuture = original.WorkloadNetworksDeleteSegmentFuture
type WorkloadNetworksDeleteVMGroupFuture = original.WorkloadNetworksDeleteVMGroupFuture
type WorkloadNetworksUpdateDhcpFuture = original.WorkloadNetworksUpdateDhcpFuture
type WorkloadNetworksUpdatePortMirroringFuture = original.WorkloadNetworksUpdatePortMirroringFuture
type WorkloadNetworksUpdateSegmentsFuture = original.WorkloadNetworksUpdateSegmentsFuture
type WorkloadNetworksUpdateVMGroupFuture = original.WorkloadNetworksUpdateVMGroupFuture

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
func NewGlobalReachConnectionListIterator(page GlobalReachConnectionListPage) GlobalReachConnectionListIterator {
	return original.NewGlobalReachConnectionListIterator(page)
}
func NewGlobalReachConnectionListPage(cur GlobalReachConnectionList, getNextPage func(context.Context, GlobalReachConnectionList) (GlobalReachConnectionList, error)) GlobalReachConnectionListPage {
	return original.NewGlobalReachConnectionListPage(cur, getNextPage)
}
func NewGlobalReachConnectionsClient(subscriptionID string) GlobalReachConnectionsClient {
	return original.NewGlobalReachConnectionsClient(subscriptionID)
}
func NewGlobalReachConnectionsClientWithBaseURI(baseURI string, subscriptionID string) GlobalReachConnectionsClient {
	return original.NewGlobalReachConnectionsClientWithBaseURI(baseURI, subscriptionID)
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
func NewPrivateCloudListPage(cur PrivateCloudList, getNextPage func(context.Context, PrivateCloudList) (PrivateCloudList, error)) PrivateCloudListPage {
	return original.NewPrivateCloudListPage(cur, getNextPage)
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
func NewWorkloadNetworkDhcpListIterator(page WorkloadNetworkDhcpListPage) WorkloadNetworkDhcpListIterator {
	return original.NewWorkloadNetworkDhcpListIterator(page)
}
func NewWorkloadNetworkDhcpListPage(cur WorkloadNetworkDhcpList, getNextPage func(context.Context, WorkloadNetworkDhcpList) (WorkloadNetworkDhcpList, error)) WorkloadNetworkDhcpListPage {
	return original.NewWorkloadNetworkDhcpListPage(cur, getNextPage)
}
func NewWorkloadNetworkGatewayListIterator(page WorkloadNetworkGatewayListPage) WorkloadNetworkGatewayListIterator {
	return original.NewWorkloadNetworkGatewayListIterator(page)
}
func NewWorkloadNetworkGatewayListPage(cur WorkloadNetworkGatewayList, getNextPage func(context.Context, WorkloadNetworkGatewayList) (WorkloadNetworkGatewayList, error)) WorkloadNetworkGatewayListPage {
	return original.NewWorkloadNetworkGatewayListPage(cur, getNextPage)
}
func NewWorkloadNetworkPortMirroringListIterator(page WorkloadNetworkPortMirroringListPage) WorkloadNetworkPortMirroringListIterator {
	return original.NewWorkloadNetworkPortMirroringListIterator(page)
}
func NewWorkloadNetworkPortMirroringListPage(cur WorkloadNetworkPortMirroringList, getNextPage func(context.Context, WorkloadNetworkPortMirroringList) (WorkloadNetworkPortMirroringList, error)) WorkloadNetworkPortMirroringListPage {
	return original.NewWorkloadNetworkPortMirroringListPage(cur, getNextPage)
}
func NewWorkloadNetworkSegmentsListIterator(page WorkloadNetworkSegmentsListPage) WorkloadNetworkSegmentsListIterator {
	return original.NewWorkloadNetworkSegmentsListIterator(page)
}
func NewWorkloadNetworkSegmentsListPage(cur WorkloadNetworkSegmentsList, getNextPage func(context.Context, WorkloadNetworkSegmentsList) (WorkloadNetworkSegmentsList, error)) WorkloadNetworkSegmentsListPage {
	return original.NewWorkloadNetworkSegmentsListPage(cur, getNextPage)
}
func NewWorkloadNetworkVMGroupsListIterator(page WorkloadNetworkVMGroupsListPage) WorkloadNetworkVMGroupsListIterator {
	return original.NewWorkloadNetworkVMGroupsListIterator(page)
}
func NewWorkloadNetworkVMGroupsListPage(cur WorkloadNetworkVMGroupsList, getNextPage func(context.Context, WorkloadNetworkVMGroupsList) (WorkloadNetworkVMGroupsList, error)) WorkloadNetworkVMGroupsListPage {
	return original.NewWorkloadNetworkVMGroupsListPage(cur, getNextPage)
}
func NewWorkloadNetworkVirtualMachinesListIterator(page WorkloadNetworkVirtualMachinesListPage) WorkloadNetworkVirtualMachinesListIterator {
	return original.NewWorkloadNetworkVirtualMachinesListIterator(page)
}
func NewWorkloadNetworkVirtualMachinesListPage(cur WorkloadNetworkVirtualMachinesList, getNextPage func(context.Context, WorkloadNetworkVirtualMachinesList) (WorkloadNetworkVirtualMachinesList, error)) WorkloadNetworkVirtualMachinesListPage {
	return original.NewWorkloadNetworkVirtualMachinesListPage(cur, getNextPage)
}
func NewWorkloadNetworksClient(subscriptionID string) WorkloadNetworksClient {
	return original.NewWorkloadNetworksClient(subscriptionID)
}
func NewWorkloadNetworksClientWithBaseURI(baseURI string, subscriptionID string) WorkloadNetworksClient {
	return original.NewWorkloadNetworksClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return original.PossibleClusterProvisioningStateValues()
}
func PossibleDhcpTypeValues() []DhcpType {
	return original.PossibleDhcpTypeValues()
}
func PossibleExpressRouteAuthorizationProvisioningStateValues() []ExpressRouteAuthorizationProvisioningState {
	return original.PossibleExpressRouteAuthorizationProvisioningStateValues()
}
func PossibleGlobalReachConnectionProvisioningStateValues() []GlobalReachConnectionProvisioningState {
	return original.PossibleGlobalReachConnectionProvisioningStateValues()
}
func PossibleGlobalReachConnectionStatusValues() []GlobalReachConnectionStatus {
	return original.PossibleGlobalReachConnectionStatusValues()
}
func PossibleHcxEnterpriseSiteStatusValues() []HcxEnterpriseSiteStatus {
	return original.PossibleHcxEnterpriseSiteStatusValues()
}
func PossibleInternetEnumValues() []InternetEnum {
	return original.PossibleInternetEnumValues()
}
func PossibleNsxtAdminRotateEnumValues() []NsxtAdminRotateEnum {
	return original.PossibleNsxtAdminRotateEnumValues()
}
func PossiblePortMirroringDirectionEnumValues() []PortMirroringDirectionEnum {
	return original.PossiblePortMirroringDirectionEnumValues()
}
func PossiblePortMirroringStatusEnumValues() []PortMirroringStatusEnum {
	return original.PossiblePortMirroringStatusEnumValues()
}
func PossiblePrivateCloudProvisioningStateValues() []PrivateCloudProvisioningState {
	return original.PossiblePrivateCloudProvisioningStateValues()
}
func PossibleQuotaEnabledValues() []QuotaEnabled {
	return original.PossibleQuotaEnabledValues()
}
func PossibleSegmentStatusEnumValues() []SegmentStatusEnum {
	return original.PossibleSegmentStatusEnumValues()
}
func PossibleSslEnumValues() []SslEnum {
	return original.PossibleSslEnumValues()
}
func PossibleTrialStatusValues() []TrialStatus {
	return original.PossibleTrialStatusValues()
}
func PossibleVMGroupStatusEnumValues() []VMGroupStatusEnum {
	return original.PossibleVMGroupStatusEnumValues()
}
func PossibleVMTypeEnumValues() []VMTypeEnum {
	return original.PossibleVMTypeEnumValues()
}
func PossibleVcsaAdminRotateEnumValues() []VcsaAdminRotateEnum {
	return original.PossibleVcsaAdminRotateEnumValues()
}
func PossibleWorkloadNetworkDhcpProvisioningStateValues() []WorkloadNetworkDhcpProvisioningState {
	return original.PossibleWorkloadNetworkDhcpProvisioningStateValues()
}
func PossibleWorkloadNetworkPortMirroringProvisioningStateValues() []WorkloadNetworkPortMirroringProvisioningState {
	return original.PossibleWorkloadNetworkPortMirroringProvisioningStateValues()
}
func PossibleWorkloadNetworkSegmentProvisioningStateValues() []WorkloadNetworkSegmentProvisioningState {
	return original.PossibleWorkloadNetworkSegmentProvisioningStateValues()
}
func PossibleWorkloadNetworkVMGroupProvisioningStateValues() []WorkloadNetworkVMGroupProvisioningState {
	return original.PossibleWorkloadNetworkVMGroupProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
