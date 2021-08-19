//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package peering

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/peering/mgmt/2020-10-01/peering"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ConnectionState = original.ConnectionState

const (
	Active                ConnectionState = original.Active
	Approved              ConnectionState = original.Approved
	None                  ConnectionState = original.None
	PendingApproval       ConnectionState = original.PendingApproval
	ProvisioningCompleted ConnectionState = original.ProvisioningCompleted
	ProvisioningFailed    ConnectionState = original.ProvisioningFailed
	ProvisioningStarted   ConnectionState = original.ProvisioningStarted
	Validating            ConnectionState = original.Validating
)

type DirectPeeringType = original.DirectPeeringType

const (
	Cdn      DirectPeeringType = original.Cdn
	Edge     DirectPeeringType = original.Edge
	Internal DirectPeeringType = original.Internal
	Ix       DirectPeeringType = original.Ix
	IxRs     DirectPeeringType = original.IxRs
	Transit  DirectPeeringType = original.Transit
)

type Family = original.Family

const (
	Direct   Family = original.Direct
	Exchange Family = original.Exchange
)

type Kind = original.Kind

const (
	KindDirect   Kind = original.KindDirect
	KindExchange Kind = original.KindExchange
)

type LearnedType = original.LearnedType

const (
	LearnedTypeNone               LearnedType = original.LearnedTypeNone
	LearnedTypeViaServiceProvider LearnedType = original.LearnedTypeViaServiceProvider
	LearnedTypeViaSession         LearnedType = original.LearnedTypeViaSession
)

type PrefixValidationState = original.PrefixValidationState

const (
	PrefixValidationStateFailed   PrefixValidationState = original.PrefixValidationStateFailed
	PrefixValidationStateInvalid  PrefixValidationState = original.PrefixValidationStateInvalid
	PrefixValidationStateNone     PrefixValidationState = original.PrefixValidationStateNone
	PrefixValidationStatePending  PrefixValidationState = original.PrefixValidationStatePending
	PrefixValidationStateUnknown  PrefixValidationState = original.PrefixValidationStateUnknown
	PrefixValidationStateVerified PrefixValidationState = original.PrefixValidationStateVerified
	PrefixValidationStateWarning  PrefixValidationState = original.PrefixValidationStateWarning
)

type ProvisioningState = original.ProvisioningState

const (
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type Role = original.Role

const (
	RoleEscalation Role = original.RoleEscalation
	RoleNoc        Role = original.RoleNoc
	RoleOther      Role = original.RoleOther
	RolePolicy     Role = original.RolePolicy
	RoleService    Role = original.RoleService
	RoleTechnical  Role = original.RoleTechnical
)

type SessionAddressProvider = original.SessionAddressProvider

const (
	Microsoft SessionAddressProvider = original.Microsoft
	Peer      SessionAddressProvider = original.Peer
)

type SessionStateV4 = original.SessionStateV4

const (
	SessionStateV4Active        SessionStateV4 = original.SessionStateV4Active
	SessionStateV4Connect       SessionStateV4 = original.SessionStateV4Connect
	SessionStateV4Established   SessionStateV4 = original.SessionStateV4Established
	SessionStateV4Idle          SessionStateV4 = original.SessionStateV4Idle
	SessionStateV4None          SessionStateV4 = original.SessionStateV4None
	SessionStateV4OpenConfirm   SessionStateV4 = original.SessionStateV4OpenConfirm
	SessionStateV4OpenReceived  SessionStateV4 = original.SessionStateV4OpenReceived
	SessionStateV4OpenSent      SessionStateV4 = original.SessionStateV4OpenSent
	SessionStateV4PendingAdd    SessionStateV4 = original.SessionStateV4PendingAdd
	SessionStateV4PendingRemove SessionStateV4 = original.SessionStateV4PendingRemove
	SessionStateV4PendingUpdate SessionStateV4 = original.SessionStateV4PendingUpdate
)

type SessionStateV6 = original.SessionStateV6

const (
	SessionStateV6Active        SessionStateV6 = original.SessionStateV6Active
	SessionStateV6Connect       SessionStateV6 = original.SessionStateV6Connect
	SessionStateV6Established   SessionStateV6 = original.SessionStateV6Established
	SessionStateV6Idle          SessionStateV6 = original.SessionStateV6Idle
	SessionStateV6None          SessionStateV6 = original.SessionStateV6None
	SessionStateV6OpenConfirm   SessionStateV6 = original.SessionStateV6OpenConfirm
	SessionStateV6OpenReceived  SessionStateV6 = original.SessionStateV6OpenReceived
	SessionStateV6OpenSent      SessionStateV6 = original.SessionStateV6OpenSent
	SessionStateV6PendingAdd    SessionStateV6 = original.SessionStateV6PendingAdd
	SessionStateV6PendingRemove SessionStateV6 = original.SessionStateV6PendingRemove
	SessionStateV6PendingUpdate SessionStateV6 = original.SessionStateV6PendingUpdate
)

type Size = original.Size

const (
	Free      Size = original.Free
	Metered   Size = original.Metered
	Unlimited Size = original.Unlimited
)

type Tier = original.Tier

const (
	Basic   Tier = original.Basic
	Premium Tier = original.Premium
)

type ValidationState = original.ValidationState

const (
	ValidationStateApproved ValidationState = original.ValidationStateApproved
	ValidationStateFailed   ValidationState = original.ValidationStateFailed
	ValidationStateNone     ValidationState = original.ValidationStateNone
	ValidationStatePending  ValidationState = original.ValidationStatePending
)

type BandwidthOffer = original.BandwidthOffer
type BaseClient = original.BaseClient
type BgpSession = original.BgpSession
type CdnPeeringPrefix = original.CdnPeeringPrefix
type CdnPeeringPrefixListResult = original.CdnPeeringPrefixListResult
type CdnPeeringPrefixListResultIterator = original.CdnPeeringPrefixListResultIterator
type CdnPeeringPrefixListResultPage = original.CdnPeeringPrefixListResultPage
type CdnPeeringPrefixProperties = original.CdnPeeringPrefixProperties
type CdnPeeringPrefixesClient = original.CdnPeeringPrefixesClient
type CheckServiceProviderAvailabilityInput = original.CheckServiceProviderAvailabilityInput
type ContactDetail = original.ContactDetail
type DirectConnection = original.DirectConnection
type DirectPeeringFacility = original.DirectPeeringFacility
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type ExchangeConnection = original.ExchangeConnection
type ExchangePeeringFacility = original.ExchangePeeringFacility
type LegacyPeeringsClient = original.LegacyPeeringsClient
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type Location = original.Location
type LocationListResult = original.LocationListResult
type LocationListResultIterator = original.LocationListResultIterator
type LocationListResultPage = original.LocationListResultPage
type LocationProperties = original.LocationProperties
type LocationPropertiesDirect = original.LocationPropertiesDirect
type LocationPropertiesExchange = original.LocationPropertiesExchange
type LocationsClient = original.LocationsClient
type Model = original.Model
type Operation = original.Operation
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PeerAsn = original.PeerAsn
type PeerAsnListResult = original.PeerAsnListResult
type PeerAsnListResultIterator = original.PeerAsnListResultIterator
type PeerAsnListResultPage = original.PeerAsnListResultPage
type PeerAsnProperties = original.PeerAsnProperties
type PeerAsnsClient = original.PeerAsnsClient
type PeeringsClient = original.PeeringsClient
type PrefixesClient = original.PrefixesClient
type Properties = original.Properties
type PropertiesDirect = original.PropertiesDirect
type PropertiesExchange = original.PropertiesExchange
type ReceivedRoute = original.ReceivedRoute
type ReceivedRouteListResult = original.ReceivedRouteListResult
type ReceivedRouteListResultIterator = original.ReceivedRouteListResultIterator
type ReceivedRouteListResultPage = original.ReceivedRouteListResultPage
type ReceivedRoutesClient = original.ReceivedRoutesClient
type RegisteredAsn = original.RegisteredAsn
type RegisteredAsnListResult = original.RegisteredAsnListResult
type RegisteredAsnListResultIterator = original.RegisteredAsnListResultIterator
type RegisteredAsnListResultPage = original.RegisteredAsnListResultPage
type RegisteredAsnProperties = original.RegisteredAsnProperties
type RegisteredAsnsClient = original.RegisteredAsnsClient
type RegisteredPrefix = original.RegisteredPrefix
type RegisteredPrefixListResult = original.RegisteredPrefixListResult
type RegisteredPrefixListResultIterator = original.RegisteredPrefixListResultIterator
type RegisteredPrefixListResultPage = original.RegisteredPrefixListResultPage
type RegisteredPrefixProperties = original.RegisteredPrefixProperties
type RegisteredPrefixesClient = original.RegisteredPrefixesClient
type Resource = original.Resource
type ResourceTags = original.ResourceTags
type Service = original.Service
type ServiceCountriesClient = original.ServiceCountriesClient
type ServiceCountry = original.ServiceCountry
type ServiceCountryListResult = original.ServiceCountryListResult
type ServiceCountryListResultIterator = original.ServiceCountryListResultIterator
type ServiceCountryListResultPage = original.ServiceCountryListResultPage
type ServiceListResult = original.ServiceListResult
type ServiceListResultIterator = original.ServiceListResultIterator
type ServiceListResultPage = original.ServiceListResultPage
type ServiceLocation = original.ServiceLocation
type ServiceLocationListResult = original.ServiceLocationListResult
type ServiceLocationListResultIterator = original.ServiceLocationListResultIterator
type ServiceLocationListResultPage = original.ServiceLocationListResultPage
type ServiceLocationProperties = original.ServiceLocationProperties
type ServiceLocationsClient = original.ServiceLocationsClient
type ServicePrefix = original.ServicePrefix
type ServicePrefixEvent = original.ServicePrefixEvent
type ServicePrefixListResult = original.ServicePrefixListResult
type ServicePrefixListResultIterator = original.ServicePrefixListResultIterator
type ServicePrefixListResultPage = original.ServicePrefixListResultPage
type ServicePrefixProperties = original.ServicePrefixProperties
type ServiceProperties = original.ServiceProperties
type ServiceProvider = original.ServiceProvider
type ServiceProviderListResult = original.ServiceProviderListResult
type ServiceProviderListResultIterator = original.ServiceProviderListResultIterator
type ServiceProviderListResultPage = original.ServiceProviderListResultPage
type ServiceProviderProperties = original.ServiceProviderProperties
type ServiceProvidersClient = original.ServiceProvidersClient
type ServiceSku = original.ServiceSku
type ServicesClient = original.ServicesClient
type Sku = original.Sku
type String = original.String
type SubResource = original.SubResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCdnPeeringPrefixListResultIterator(page CdnPeeringPrefixListResultPage) CdnPeeringPrefixListResultIterator {
	return original.NewCdnPeeringPrefixListResultIterator(page)
}
func NewCdnPeeringPrefixListResultPage(cur CdnPeeringPrefixListResult, getNextPage func(context.Context, CdnPeeringPrefixListResult) (CdnPeeringPrefixListResult, error)) CdnPeeringPrefixListResultPage {
	return original.NewCdnPeeringPrefixListResultPage(cur, getNextPage)
}
func NewCdnPeeringPrefixesClient(subscriptionID string) CdnPeeringPrefixesClient {
	return original.NewCdnPeeringPrefixesClient(subscriptionID)
}
func NewCdnPeeringPrefixesClientWithBaseURI(baseURI string, subscriptionID string) CdnPeeringPrefixesClient {
	return original.NewCdnPeeringPrefixesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLegacyPeeringsClient(subscriptionID string) LegacyPeeringsClient {
	return original.NewLegacyPeeringsClient(subscriptionID)
}
func NewLegacyPeeringsClientWithBaseURI(baseURI string, subscriptionID string) LegacyPeeringsClient {
	return original.NewLegacyPeeringsClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewLocationListResultIterator(page LocationListResultPage) LocationListResultIterator {
	return original.NewLocationListResultIterator(page)
}
func NewLocationListResultPage(cur LocationListResult, getNextPage func(context.Context, LocationListResult) (LocationListResult, error)) LocationListResultPage {
	return original.NewLocationListResultPage(cur, getNextPage)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPeerAsnListResultIterator(page PeerAsnListResultPage) PeerAsnListResultIterator {
	return original.NewPeerAsnListResultIterator(page)
}
func NewPeerAsnListResultPage(cur PeerAsnListResult, getNextPage func(context.Context, PeerAsnListResult) (PeerAsnListResult, error)) PeerAsnListResultPage {
	return original.NewPeerAsnListResultPage(cur, getNextPage)
}
func NewPeerAsnsClient(subscriptionID string) PeerAsnsClient {
	return original.NewPeerAsnsClient(subscriptionID)
}
func NewPeerAsnsClientWithBaseURI(baseURI string, subscriptionID string) PeerAsnsClient {
	return original.NewPeerAsnsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPeeringsClient(subscriptionID string) PeeringsClient {
	return original.NewPeeringsClient(subscriptionID)
}
func NewPeeringsClientWithBaseURI(baseURI string, subscriptionID string) PeeringsClient {
	return original.NewPeeringsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrefixesClient(subscriptionID string) PrefixesClient {
	return original.NewPrefixesClient(subscriptionID)
}
func NewPrefixesClientWithBaseURI(baseURI string, subscriptionID string) PrefixesClient {
	return original.NewPrefixesClientWithBaseURI(baseURI, subscriptionID)
}
func NewReceivedRouteListResultIterator(page ReceivedRouteListResultPage) ReceivedRouteListResultIterator {
	return original.NewReceivedRouteListResultIterator(page)
}
func NewReceivedRouteListResultPage(cur ReceivedRouteListResult, getNextPage func(context.Context, ReceivedRouteListResult) (ReceivedRouteListResult, error)) ReceivedRouteListResultPage {
	return original.NewReceivedRouteListResultPage(cur, getNextPage)
}
func NewReceivedRoutesClient(subscriptionID string) ReceivedRoutesClient {
	return original.NewReceivedRoutesClient(subscriptionID)
}
func NewReceivedRoutesClientWithBaseURI(baseURI string, subscriptionID string) ReceivedRoutesClient {
	return original.NewReceivedRoutesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegisteredAsnListResultIterator(page RegisteredAsnListResultPage) RegisteredAsnListResultIterator {
	return original.NewRegisteredAsnListResultIterator(page)
}
func NewRegisteredAsnListResultPage(cur RegisteredAsnListResult, getNextPage func(context.Context, RegisteredAsnListResult) (RegisteredAsnListResult, error)) RegisteredAsnListResultPage {
	return original.NewRegisteredAsnListResultPage(cur, getNextPage)
}
func NewRegisteredAsnsClient(subscriptionID string) RegisteredAsnsClient {
	return original.NewRegisteredAsnsClient(subscriptionID)
}
func NewRegisteredAsnsClientWithBaseURI(baseURI string, subscriptionID string) RegisteredAsnsClient {
	return original.NewRegisteredAsnsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRegisteredPrefixListResultIterator(page RegisteredPrefixListResultPage) RegisteredPrefixListResultIterator {
	return original.NewRegisteredPrefixListResultIterator(page)
}
func NewRegisteredPrefixListResultPage(cur RegisteredPrefixListResult, getNextPage func(context.Context, RegisteredPrefixListResult) (RegisteredPrefixListResult, error)) RegisteredPrefixListResultPage {
	return original.NewRegisteredPrefixListResultPage(cur, getNextPage)
}
func NewRegisteredPrefixesClient(subscriptionID string) RegisteredPrefixesClient {
	return original.NewRegisteredPrefixesClient(subscriptionID)
}
func NewRegisteredPrefixesClientWithBaseURI(baseURI string, subscriptionID string) RegisteredPrefixesClient {
	return original.NewRegisteredPrefixesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceCountriesClient(subscriptionID string) ServiceCountriesClient {
	return original.NewServiceCountriesClient(subscriptionID)
}
func NewServiceCountriesClientWithBaseURI(baseURI string, subscriptionID string) ServiceCountriesClient {
	return original.NewServiceCountriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceCountryListResultIterator(page ServiceCountryListResultPage) ServiceCountryListResultIterator {
	return original.NewServiceCountryListResultIterator(page)
}
func NewServiceCountryListResultPage(cur ServiceCountryListResult, getNextPage func(context.Context, ServiceCountryListResult) (ServiceCountryListResult, error)) ServiceCountryListResultPage {
	return original.NewServiceCountryListResultPage(cur, getNextPage)
}
func NewServiceListResultIterator(page ServiceListResultPage) ServiceListResultIterator {
	return original.NewServiceListResultIterator(page)
}
func NewServiceListResultPage(cur ServiceListResult, getNextPage func(context.Context, ServiceListResult) (ServiceListResult, error)) ServiceListResultPage {
	return original.NewServiceListResultPage(cur, getNextPage)
}
func NewServiceLocationListResultIterator(page ServiceLocationListResultPage) ServiceLocationListResultIterator {
	return original.NewServiceLocationListResultIterator(page)
}
func NewServiceLocationListResultPage(cur ServiceLocationListResult, getNextPage func(context.Context, ServiceLocationListResult) (ServiceLocationListResult, error)) ServiceLocationListResultPage {
	return original.NewServiceLocationListResultPage(cur, getNextPage)
}
func NewServiceLocationsClient(subscriptionID string) ServiceLocationsClient {
	return original.NewServiceLocationsClient(subscriptionID)
}
func NewServiceLocationsClientWithBaseURI(baseURI string, subscriptionID string) ServiceLocationsClient {
	return original.NewServiceLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicePrefixListResultIterator(page ServicePrefixListResultPage) ServicePrefixListResultIterator {
	return original.NewServicePrefixListResultIterator(page)
}
func NewServicePrefixListResultPage(cur ServicePrefixListResult, getNextPage func(context.Context, ServicePrefixListResult) (ServicePrefixListResult, error)) ServicePrefixListResultPage {
	return original.NewServicePrefixListResultPage(cur, getNextPage)
}
func NewServiceProviderListResultIterator(page ServiceProviderListResultPage) ServiceProviderListResultIterator {
	return original.NewServiceProviderListResultIterator(page)
}
func NewServiceProviderListResultPage(cur ServiceProviderListResult, getNextPage func(context.Context, ServiceProviderListResult) (ServiceProviderListResult, error)) ServiceProviderListResultPage {
	return original.NewServiceProviderListResultPage(cur, getNextPage)
}
func NewServiceProvidersClient(subscriptionID string) ServiceProvidersClient {
	return original.NewServiceProvidersClient(subscriptionID)
}
func NewServiceProvidersClientWithBaseURI(baseURI string, subscriptionID string) ServiceProvidersClient {
	return original.NewServiceProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleConnectionStateValues() []ConnectionState {
	return original.PossibleConnectionStateValues()
}
func PossibleDirectPeeringTypeValues() []DirectPeeringType {
	return original.PossibleDirectPeeringTypeValues()
}
func PossibleFamilyValues() []Family {
	return original.PossibleFamilyValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLearnedTypeValues() []LearnedType {
	return original.PossibleLearnedTypeValues()
}
func PossiblePrefixValidationStateValues() []PrefixValidationState {
	return original.PossiblePrefixValidationStateValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRoleValues() []Role {
	return original.PossibleRoleValues()
}
func PossibleSessionAddressProviderValues() []SessionAddressProvider {
	return original.PossibleSessionAddressProviderValues()
}
func PossibleSessionStateV4Values() []SessionStateV4 {
	return original.PossibleSessionStateV4Values()
}
func PossibleSessionStateV6Values() []SessionStateV6 {
	return original.PossibleSessionStateV6Values()
}
func PossibleSizeValues() []Size {
	return original.PossibleSizeValues()
}
func PossibleTierValues() []Tier {
	return original.PossibleTierValues()
}
func PossibleValidationStateValues() []ValidationState {
	return original.PossibleValidationStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
