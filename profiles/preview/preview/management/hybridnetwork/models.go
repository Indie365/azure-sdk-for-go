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

package hybridnetwork

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/management/2020-01-01-preview/hybridnetwork"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type DeviceType = original.DeviceType

const (
	DeviceTypeDevicePropertiesFormat DeviceType = original.DeviceTypeDevicePropertiesFormat
)

type IPAllocationMethod = original.IPAllocationMethod

const (
	Dynamic IPAllocationMethod = original.Dynamic
	Static  IPAllocationMethod = original.Static
	Unknown IPAllocationMethod = original.Unknown
)

type IPVersion = original.IPVersion

const (
	IPVersionIPv4    IPVersion = original.IPVersionIPv4
	IPVersionUnknown IPVersion = original.IPVersionUnknown
)

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = original.OperatingSystemTypesLinux
	OperatingSystemTypesUnknown OperatingSystemTypes = original.OperatingSystemTypesUnknown
	OperatingSystemTypesWindows OperatingSystemTypes = original.OperatingSystemTypesWindows
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted  ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCanceled  ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateDeleted   ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUnknown   ProvisioningState = original.ProvisioningStateUnknown
)

type SkuDeploymentMode = original.SkuDeploymentMode

const (
	SkuDeploymentModeAzure           SkuDeploymentMode = original.SkuDeploymentModeAzure
	SkuDeploymentModePrivateEdgeZone SkuDeploymentMode = original.SkuDeploymentModePrivateEdgeZone
	SkuDeploymentModeUnknown         SkuDeploymentMode = original.SkuDeploymentModeUnknown
)

type SkuType = original.SkuType

const (
	SkuTypeEvolvedPacketCore SkuType = original.SkuTypeEvolvedPacketCore
	SkuTypeSDWAN             SkuType = original.SkuTypeSDWAN
	SkuTypeUnknown           SkuType = original.SkuTypeUnknown
)

type Status = original.Status

const (
	StatusNotRegistered Status = original.StatusNotRegistered
	StatusRegistered    Status = original.StatusRegistered
	StatusUnknown       Status = original.StatusUnknown
)

type VHDType = original.VHDType

const (
	VHDTypeUnknown VHDType = original.VHDTypeUnknown
	VHDTypeVHD     VHDType = original.VHDTypeVHD
	VHDTypeVHDX    VHDType = original.VHDTypeVHDX
)

type VMSwitchType = original.VMSwitchType

const (
	VMSwitchTypeInternal   VMSwitchType = original.VMSwitchTypeInternal
	VMSwitchTypeLan        VMSwitchType = original.VMSwitchTypeLan
	VMSwitchTypeManagement VMSwitchType = original.VMSwitchTypeManagement
	VMSwitchTypeUnknown    VMSwitchType = original.VMSwitchTypeUnknown
	VMSwitchTypeWan        VMSwitchType = original.VMSwitchTypeWan
)

type VendorProvisioningState = original.VendorProvisioningState

const (
	VendorProvisioningStateDeprovisioned            VendorProvisioningState = original.VendorProvisioningStateDeprovisioned
	VendorProvisioningStateNotProvisioned           VendorProvisioningState = original.VendorProvisioningStateNotProvisioned
	VendorProvisioningStateProvisioned              VendorProvisioningState = original.VendorProvisioningStateProvisioned
	VendorProvisioningStateProvisioning             VendorProvisioningState = original.VendorProvisioningStateProvisioning
	VendorProvisioningStateUnknown                  VendorProvisioningState = original.VendorProvisioningStateUnknown
	VendorProvisioningStateUserDataValidationFailed VendorProvisioningState = original.VendorProvisioningStateUserDataValidationFailed
)

type VirtualMachineSizeTypes = original.VirtualMachineSizeTypes

const (
	VirtualMachineSizeTypesStandardD11V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD11V2
	VirtualMachineSizeTypesStandardD12V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD12V2
	VirtualMachineSizeTypesStandardD13V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD13V2
	VirtualMachineSizeTypesStandardD1V2   VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD1V2
	VirtualMachineSizeTypesStandardD2V2   VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD2V2
	VirtualMachineSizeTypesStandardD3V2   VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD3V2
	VirtualMachineSizeTypesStandardD4V2   VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD4V2
	VirtualMachineSizeTypesStandardD5V2   VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardD5V2
	VirtualMachineSizeTypesStandardDS11V2 VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS11V2
	VirtualMachineSizeTypesStandardDS12V2 VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS12V2
	VirtualMachineSizeTypesStandardDS13V2 VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS13V2
	VirtualMachineSizeTypesStandardDS1V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS1V2
	VirtualMachineSizeTypesStandardDS2V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS2V2
	VirtualMachineSizeTypesStandardDS3V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS3V2
	VirtualMachineSizeTypesStandardDS4V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS4V2
	VirtualMachineSizeTypesStandardDS5V2  VirtualMachineSizeTypes = original.VirtualMachineSizeTypesStandardDS5V2
	VirtualMachineSizeTypesUnknown        VirtualMachineSizeTypes = original.VirtualMachineSizeTypesUnknown
)

type VirtualNetworkFunctionRoleConfigurationType = original.VirtualNetworkFunctionRoleConfigurationType

const (
	VirtualNetworkFunctionRoleConfigurationTypeUnknown        VirtualNetworkFunctionRoleConfigurationType = original.VirtualNetworkFunctionRoleConfigurationTypeUnknown
	VirtualNetworkFunctionRoleConfigurationTypeVirtualMachine VirtualNetworkFunctionRoleConfigurationType = original.VirtualNetworkFunctionRoleConfigurationTypeVirtualMachine
)

type BaseClient = original.BaseClient
type BasicDevicePropertiesFormat = original.BasicDevicePropertiesFormat
type Device = original.Device
type DeviceListResult = original.DeviceListResult
type DeviceListResultIterator = original.DeviceListResultIterator
type DeviceListResultPage = original.DeviceListResultPage
type DevicePropertiesFormat = original.DevicePropertiesFormat
type DeviceRegistrationKey = original.DeviceRegistrationKey
type DevicesClient = original.DevicesClient
type DevicesCreateOrUpdateFuture = original.DevicesCreateOrUpdateFuture
type DevicesDeleteFuture = original.DevicesDeleteFuture
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type ImageReference = original.ImageReference
type NetworkInterface = original.NetworkInterface
type NetworkInterfaceIPConfiguration = original.NetworkInterfaceIPConfiguration
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type OsProfile = original.OsProfile
type PreviewSubscription = original.PreviewSubscription
type PreviewSubscriptionsList = original.PreviewSubscriptionsList
type PreviewSubscriptionsListIterator = original.PreviewSubscriptionsListIterator
type PreviewSubscriptionsListPage = original.PreviewSubscriptionsListPage
type Resource = original.Resource
type SkuOverview = original.SkuOverview
type SubResource = original.SubResource
type TagsObject = original.TagsObject
type Vendor = original.Vendor
type VendorDetails = original.VendorDetails
type VendorListResult = original.VendorListResult
type VendorListResultIterator = original.VendorListResultIterator
type VendorListResultPage = original.VendorListResultPage
type VendorPropertiesFormat = original.VendorPropertiesFormat
type VendorSku = original.VendorSku
type VendorSkuListResult = original.VendorSkuListResult
type VendorSkuListResultIterator = original.VendorSkuListResultIterator
type VendorSkuListResultPage = original.VendorSkuListResultPage
type VendorSkuPreviewClient = original.VendorSkuPreviewClient
type VendorSkuPreviewCreateOrUpdateFuture = original.VendorSkuPreviewCreateOrUpdateFuture
type VendorSkuPreviewDeleteFuture = original.VendorSkuPreviewDeleteFuture
type VendorSkuPropertiesFormat = original.VendorSkuPropertiesFormat
type VendorSkusClient = original.VendorSkusClient
type VendorSkusCreateOrUpdateFuture = original.VendorSkusCreateOrUpdateFuture
type VendorSkusDeleteFuture = original.VendorSkusDeleteFuture
type VendorVirtualNetworkFunction = original.VendorVirtualNetworkFunction
type VendorVirtualNetworkFunctionListResult = original.VendorVirtualNetworkFunctionListResult
type VendorVirtualNetworkFunctionListResultIterator = original.VendorVirtualNetworkFunctionListResultIterator
type VendorVirtualNetworkFunctionListResultPage = original.VendorVirtualNetworkFunctionListResultPage
type VendorVirtualNetworkFunctionPropertiesFormat = original.VendorVirtualNetworkFunctionPropertiesFormat
type VendorVirtualNetworkFunctionsClient = original.VendorVirtualNetworkFunctionsClient
type VendorVirtualNetworkFunctionsCreateOrUpdateFuture = original.VendorVirtualNetworkFunctionsCreateOrUpdateFuture
type VendorsClient = original.VendorsClient
type VendorsCreateOrUpdateFuture = original.VendorsCreateOrUpdateFuture
type VendorsDeleteFuture = original.VendorsDeleteFuture
type VirtualNetworkFunction = original.VirtualNetworkFunction
type VirtualNetworkFunctionListResult = original.VirtualNetworkFunctionListResult
type VirtualNetworkFunctionListResultIterator = original.VirtualNetworkFunctionListResultIterator
type VirtualNetworkFunctionListResultPage = original.VirtualNetworkFunctionListResultPage
type VirtualNetworkFunctionPropertiesFormat = original.VirtualNetworkFunctionPropertiesFormat
type VirtualNetworkFunctionRoleConfiguration = original.VirtualNetworkFunctionRoleConfiguration
type VirtualNetworkFunctionSkuDetails = original.VirtualNetworkFunctionSkuDetails
type VirtualNetworkFunctionSkuDetailsIterator = original.VirtualNetworkFunctionSkuDetailsIterator
type VirtualNetworkFunctionSkuDetailsPage = original.VirtualNetworkFunctionSkuDetailsPage
type VirtualNetworkFunctionSkuListResult = original.VirtualNetworkFunctionSkuListResult
type VirtualNetworkFunctionSkuListResultIterator = original.VirtualNetworkFunctionSkuListResultIterator
type VirtualNetworkFunctionSkuListResultPage = original.VirtualNetworkFunctionSkuListResultPage
type VirtualNetworkFunctionSkuRoleDetails = original.VirtualNetworkFunctionSkuRoleDetails
type VirtualNetworkFunctionTemplate = original.VirtualNetworkFunctionTemplate
type VirtualNetworkFunctionUserConfiguration = original.VirtualNetworkFunctionUserConfiguration
type VirtualNetworkFunctionVendor = original.VirtualNetworkFunctionVendor
type VirtualNetworkFunctionVendorConfiguration = original.VirtualNetworkFunctionVendorConfiguration
type VirtualNetworkFunctionVendorListResult = original.VirtualNetworkFunctionVendorListResult
type VirtualNetworkFunctionVendorListResultIterator = original.VirtualNetworkFunctionVendorListResultIterator
type VirtualNetworkFunctionVendorListResultPage = original.VirtualNetworkFunctionVendorListResultPage
type VirtualNetworkFunctionVendorSkusClient = original.VirtualNetworkFunctionVendorSkusClient
type VirtualNetworkFunctionVendorsClient = original.VirtualNetworkFunctionVendorsClient
type VirtualNetworkFunctionsClient = original.VirtualNetworkFunctionsClient
type VirtualNetworkFunctionsCreateOrUpdateFuture = original.VirtualNetworkFunctionsCreateOrUpdateFuture
type VirtualNetworkFunctionsDeleteFuture = original.VirtualNetworkFunctionsDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDeviceListResultIterator(page DeviceListResultPage) DeviceListResultIterator {
	return original.NewDeviceListResultIterator(page)
}
func NewDeviceListResultPage(getNextPage func(context.Context, DeviceListResult) (DeviceListResult, error)) DeviceListResultPage {
	return original.NewDeviceListResultPage(getNextPage)
}
func NewDevicesClient(subscriptionID string) DevicesClient {
	return original.NewDevicesClient(subscriptionID)
}
func NewDevicesClientWithBaseURI(baseURI string, subscriptionID string) DevicesClient {
	return original.NewDevicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPreviewSubscriptionsListIterator(page PreviewSubscriptionsListPage) PreviewSubscriptionsListIterator {
	return original.NewPreviewSubscriptionsListIterator(page)
}
func NewPreviewSubscriptionsListPage(getNextPage func(context.Context, PreviewSubscriptionsList) (PreviewSubscriptionsList, error)) PreviewSubscriptionsListPage {
	return original.NewPreviewSubscriptionsListPage(getNextPage)
}
func NewVendorListResultIterator(page VendorListResultPage) VendorListResultIterator {
	return original.NewVendorListResultIterator(page)
}
func NewVendorListResultPage(getNextPage func(context.Context, VendorListResult) (VendorListResult, error)) VendorListResultPage {
	return original.NewVendorListResultPage(getNextPage)
}
func NewVendorSkuListResultIterator(page VendorSkuListResultPage) VendorSkuListResultIterator {
	return original.NewVendorSkuListResultIterator(page)
}
func NewVendorSkuListResultPage(getNextPage func(context.Context, VendorSkuListResult) (VendorSkuListResult, error)) VendorSkuListResultPage {
	return original.NewVendorSkuListResultPage(getNextPage)
}
func NewVendorSkuPreviewClient(subscriptionID string) VendorSkuPreviewClient {
	return original.NewVendorSkuPreviewClient(subscriptionID)
}
func NewVendorSkuPreviewClientWithBaseURI(baseURI string, subscriptionID string) VendorSkuPreviewClient {
	return original.NewVendorSkuPreviewClientWithBaseURI(baseURI, subscriptionID)
}
func NewVendorSkusClient(subscriptionID string) VendorSkusClient {
	return original.NewVendorSkusClient(subscriptionID)
}
func NewVendorSkusClientWithBaseURI(baseURI string, subscriptionID string) VendorSkusClient {
	return original.NewVendorSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewVendorVirtualNetworkFunctionListResultIterator(page VendorVirtualNetworkFunctionListResultPage) VendorVirtualNetworkFunctionListResultIterator {
	return original.NewVendorVirtualNetworkFunctionListResultIterator(page)
}
func NewVendorVirtualNetworkFunctionListResultPage(getNextPage func(context.Context, VendorVirtualNetworkFunctionListResult) (VendorVirtualNetworkFunctionListResult, error)) VendorVirtualNetworkFunctionListResultPage {
	return original.NewVendorVirtualNetworkFunctionListResultPage(getNextPage)
}
func NewVendorVirtualNetworkFunctionsClient(subscriptionID string) VendorVirtualNetworkFunctionsClient {
	return original.NewVendorVirtualNetworkFunctionsClient(subscriptionID)
}
func NewVendorVirtualNetworkFunctionsClientWithBaseURI(baseURI string, subscriptionID string) VendorVirtualNetworkFunctionsClient {
	return original.NewVendorVirtualNetworkFunctionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVendorsClient(subscriptionID string) VendorsClient {
	return original.NewVendorsClient(subscriptionID)
}
func NewVendorsClientWithBaseURI(baseURI string, subscriptionID string) VendorsClient {
	return original.NewVendorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkFunctionListResultIterator(page VirtualNetworkFunctionListResultPage) VirtualNetworkFunctionListResultIterator {
	return original.NewVirtualNetworkFunctionListResultIterator(page)
}
func NewVirtualNetworkFunctionListResultPage(getNextPage func(context.Context, VirtualNetworkFunctionListResult) (VirtualNetworkFunctionListResult, error)) VirtualNetworkFunctionListResultPage {
	return original.NewVirtualNetworkFunctionListResultPage(getNextPage)
}
func NewVirtualNetworkFunctionSkuDetailsIterator(page VirtualNetworkFunctionSkuDetailsPage) VirtualNetworkFunctionSkuDetailsIterator {
	return original.NewVirtualNetworkFunctionSkuDetailsIterator(page)
}
func NewVirtualNetworkFunctionSkuDetailsPage(getNextPage func(context.Context, VirtualNetworkFunctionSkuDetails) (VirtualNetworkFunctionSkuDetails, error)) VirtualNetworkFunctionSkuDetailsPage {
	return original.NewVirtualNetworkFunctionSkuDetailsPage(getNextPage)
}
func NewVirtualNetworkFunctionSkuListResultIterator(page VirtualNetworkFunctionSkuListResultPage) VirtualNetworkFunctionSkuListResultIterator {
	return original.NewVirtualNetworkFunctionSkuListResultIterator(page)
}
func NewVirtualNetworkFunctionSkuListResultPage(getNextPage func(context.Context, VirtualNetworkFunctionSkuListResult) (VirtualNetworkFunctionSkuListResult, error)) VirtualNetworkFunctionSkuListResultPage {
	return original.NewVirtualNetworkFunctionSkuListResultPage(getNextPage)
}
func NewVirtualNetworkFunctionVendorListResultIterator(page VirtualNetworkFunctionVendorListResultPage) VirtualNetworkFunctionVendorListResultIterator {
	return original.NewVirtualNetworkFunctionVendorListResultIterator(page)
}
func NewVirtualNetworkFunctionVendorListResultPage(getNextPage func(context.Context, VirtualNetworkFunctionVendorListResult) (VirtualNetworkFunctionVendorListResult, error)) VirtualNetworkFunctionVendorListResultPage {
	return original.NewVirtualNetworkFunctionVendorListResultPage(getNextPage)
}
func NewVirtualNetworkFunctionVendorSkusClient(subscriptionID string) VirtualNetworkFunctionVendorSkusClient {
	return original.NewVirtualNetworkFunctionVendorSkusClient(subscriptionID)
}
func NewVirtualNetworkFunctionVendorSkusClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkFunctionVendorSkusClient {
	return original.NewVirtualNetworkFunctionVendorSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkFunctionVendorsClient(subscriptionID string) VirtualNetworkFunctionVendorsClient {
	return original.NewVirtualNetworkFunctionVendorsClient(subscriptionID)
}
func NewVirtualNetworkFunctionVendorsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkFunctionVendorsClient {
	return original.NewVirtualNetworkFunctionVendorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualNetworkFunctionsClient(subscriptionID string) VirtualNetworkFunctionsClient {
	return original.NewVirtualNetworkFunctionsClient(subscriptionID)
}
func NewVirtualNetworkFunctionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkFunctionsClient {
	return original.NewVirtualNetworkFunctionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleDeviceTypeValues() []DeviceType {
	return original.PossibleDeviceTypeValues()
}
func PossibleIPAllocationMethodValues() []IPAllocationMethod {
	return original.PossibleIPAllocationMethodValues()
}
func PossibleIPVersionValues() []IPVersion {
	return original.PossibleIPVersionValues()
}
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return original.PossibleOperatingSystemTypesValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSkuDeploymentModeValues() []SkuDeploymentMode {
	return original.PossibleSkuDeploymentModeValues()
}
func PossibleSkuTypeValues() []SkuType {
	return original.PossibleSkuTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleVHDTypeValues() []VHDType {
	return original.PossibleVHDTypeValues()
}
func PossibleVMSwitchTypeValues() []VMSwitchType {
	return original.PossibleVMSwitchTypeValues()
}
func PossibleVendorProvisioningStateValues() []VendorProvisioningState {
	return original.PossibleVendorProvisioningStateValues()
}
func PossibleVirtualMachineSizeTypesValues() []VirtualMachineSizeTypes {
	return original.PossibleVirtualMachineSizeTypesValues()
}
func PossibleVirtualNetworkFunctionRoleConfigurationTypeValues() []VirtualNetworkFunctionRoleConfigurationType {
	return original.PossibleVirtualNetworkFunctionRoleConfigurationTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
