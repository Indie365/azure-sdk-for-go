// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package trafficmanager

import original "github.com/Azure/azure-sdk-for-go/services/preview/trafficmanager/mgmt/2017-09-01-preview/trafficmanager"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type EndpointMonitorStatus = original.EndpointMonitorStatus

const (
	CheckingEndpoint EndpointMonitorStatus = original.CheckingEndpoint
	Degraded         EndpointMonitorStatus = original.Degraded
	Disabled         EndpointMonitorStatus = original.Disabled
	Inactive         EndpointMonitorStatus = original.Inactive
	Online           EndpointMonitorStatus = original.Online
	Stopped          EndpointMonitorStatus = original.Stopped
)

type EndpointStatus = original.EndpointStatus

const (
	EndpointStatusDisabled EndpointStatus = original.EndpointStatusDisabled
	EndpointStatusEnabled  EndpointStatus = original.EndpointStatusEnabled
)

type MonitorProtocol = original.MonitorProtocol

const (
	HTTP  MonitorProtocol = original.HTTP
	HTTPS MonitorProtocol = original.HTTPS
	TCP   MonitorProtocol = original.TCP
)

type ProfileMonitorStatus = original.ProfileMonitorStatus

const (
	ProfileMonitorStatusCheckingEndpoints ProfileMonitorStatus = original.ProfileMonitorStatusCheckingEndpoints
	ProfileMonitorStatusDegraded          ProfileMonitorStatus = original.ProfileMonitorStatusDegraded
	ProfileMonitorStatusDisabled          ProfileMonitorStatus = original.ProfileMonitorStatusDisabled
	ProfileMonitorStatusInactive          ProfileMonitorStatus = original.ProfileMonitorStatusInactive
	ProfileMonitorStatusOnline            ProfileMonitorStatus = original.ProfileMonitorStatusOnline
)

type ProfileStatus = original.ProfileStatus

const (
	ProfileStatusDisabled ProfileStatus = original.ProfileStatusDisabled
	ProfileStatusEnabled  ProfileStatus = original.ProfileStatusEnabled
)

type TrafficRoutingMethod = original.TrafficRoutingMethod

const (
	Geographic  TrafficRoutingMethod = original.Geographic
	Performance TrafficRoutingMethod = original.Performance
	Priority    TrafficRoutingMethod = original.Priority
	Weighted    TrafficRoutingMethod = original.Weighted
)

type BaseClient = original.BaseClient
type CheckTrafficManagerRelativeDNSNameAvailabilityParameters = original.CheckTrafficManagerRelativeDNSNameAvailabilityParameters
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type DNSConfig = original.DNSConfig
type DeleteOperationResult = original.DeleteOperationResult
type Endpoint = original.Endpoint
type EndpointProperties = original.EndpointProperties
type EndpointsClient = original.EndpointsClient
type GeographicHierarchiesClient = original.GeographicHierarchiesClient
type GeographicHierarchy = original.GeographicHierarchy
type GeographicHierarchyProperties = original.GeographicHierarchyProperties
type HeatMapClient = original.HeatMapClient
type HeatMapEndpoint = original.HeatMapEndpoint
type HeatMapModel = original.HeatMapModel
type HeatMapProperties = original.HeatMapProperties
type MonitorConfig = original.MonitorConfig
type NameAvailability = original.NameAvailability
type Profile = original.Profile
type ProfileListResult = original.ProfileListResult
type ProfileProperties = original.ProfileProperties
type ProfilesClient = original.ProfilesClient
type ProxyResource = original.ProxyResource
type QueryExperience = original.QueryExperience
type Region = original.Region
type Resource = original.Resource
type TrackedResource = original.TrackedResource
type TrafficFlow = original.TrafficFlow
type UserMetricsKeyModel = original.UserMetricsKeyModel
type UserMetricsKeysClient = original.UserMetricsKeysClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewGeographicHierarchiesClient(subscriptionID string) GeographicHierarchiesClient {
	return original.NewGeographicHierarchiesClient(subscriptionID)
}
func NewGeographicHierarchiesClientWithBaseURI(baseURI string, subscriptionID string) GeographicHierarchiesClient {
	return original.NewGeographicHierarchiesClientWithBaseURI(baseURI, subscriptionID)
}
func NewHeatMapClient(subscriptionID string) HeatMapClient {
	return original.NewHeatMapClient(subscriptionID)
}
func NewHeatMapClientWithBaseURI(baseURI string, subscriptionID string) HeatMapClient {
	return original.NewHeatMapClientWithBaseURI(baseURI, subscriptionID)
}
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return original.NewProfilesClient(subscriptionID)
}
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return original.NewProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUserMetricsKeysClient(subscriptionID string) UserMetricsKeysClient {
	return original.NewUserMetricsKeysClient(subscriptionID)
}
func NewUserMetricsKeysClientWithBaseURI(baseURI string, subscriptionID string) UserMetricsKeysClient {
	return original.NewUserMetricsKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleEndpointMonitorStatusValues() []EndpointMonitorStatus {
	return original.PossibleEndpointMonitorStatusValues()
}
func PossibleEndpointStatusValues() []EndpointStatus {
	return original.PossibleEndpointStatusValues()
}
func PossibleMonitorProtocolValues() []MonitorProtocol {
	return original.PossibleMonitorProtocolValues()
}
func PossibleProfileMonitorStatusValues() []ProfileMonitorStatus {
	return original.PossibleProfileMonitorStatusValues()
}
func PossibleProfileStatusValues() []ProfileStatus {
	return original.PossibleProfileStatusValues()
}
func PossibleTrafficRoutingMethodValues() []TrafficRoutingMethod {
	return original.PossibleTrafficRoutingMethodValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
