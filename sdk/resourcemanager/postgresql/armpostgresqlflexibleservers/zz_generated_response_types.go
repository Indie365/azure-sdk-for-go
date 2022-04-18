//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	NameAvailability
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	Configuration
}

// ConfigurationsClientListByServerResponse contains the response from method ConfigurationsClient.ListByServer.
type ConfigurationsClientListByServerResponse struct {
	ConfigurationListResult
}

// ConfigurationsClientPutResponse contains the response from method ConfigurationsClient.Put.
type ConfigurationsClientPutResponse struct {
	Configuration
}

// ConfigurationsClientUpdateResponse contains the response from method ConfigurationsClient.Update.
type ConfigurationsClientUpdateResponse struct {
	Configuration
}

// DatabasesClientCreateResponse contains the response from method DatabasesClient.Create.
type DatabasesClientCreateResponse struct {
	Database
}

// DatabasesClientDeleteResponse contains the response from method DatabasesClient.Delete.
type DatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabasesClientGetResponse contains the response from method DatabasesClient.Get.
type DatabasesClientGetResponse struct {
	Database
}

// DatabasesClientListByServerResponse contains the response from method DatabasesClient.ListByServer.
type DatabasesClientListByServerResponse struct {
	DatabaseListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.CreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.Delete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	FirewallRule
}

// FirewallRulesClientListByServerResponse contains the response from method FirewallRulesClient.ListByServer.
type FirewallRulesClientListByServerResponse struct {
	FirewallRuleListResult
}

// GetPrivateDNSZoneSuffixClientExecuteResponse contains the response from method GetPrivateDNSZoneSuffixClient.Execute.
type GetPrivateDNSZoneSuffixClientExecuteResponse struct {
	// Represents a resource name availability.
	Value *string
}

// LocationBasedCapabilitiesClientExecuteResponse contains the response from method LocationBasedCapabilitiesClient.Execute.
type LocationBasedCapabilitiesClientExecuteResponse struct {
	CapabilitiesListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// ServersClientCreateResponse contains the response from method ServersClient.Create.
type ServersClientCreateResponse struct {
	Server
}

// ServersClientDeleteResponse contains the response from method ServersClient.Delete.
type ServersClientDeleteResponse struct {
	// placeholder for future response values
}

// ServersClientGetResponse contains the response from method ServersClient.Get.
type ServersClientGetResponse struct {
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.ListByResourceGroup.
type ServersClientListByResourceGroupResponse struct {
	ServerListResult
}

// ServersClientListResponse contains the response from method ServersClient.List.
type ServersClientListResponse struct {
	ServerListResult
}

// ServersClientRestartResponse contains the response from method ServersClient.Restart.
type ServersClientRestartResponse struct {
	// placeholder for future response values
}

// ServersClientStartResponse contains the response from method ServersClient.Start.
type ServersClientStartResponse struct {
	// placeholder for future response values
}

// ServersClientStopResponse contains the response from method ServersClient.Stop.
type ServersClientStopResponse struct {
	// placeholder for future response values
}

// ServersClientUpdateResponse contains the response from method ServersClient.Update.
type ServersClientUpdateResponse struct {
	Server
}

// VirtualNetworkSubnetUsageClientExecuteResponse contains the response from method VirtualNetworkSubnetUsageClient.Execute.
type VirtualNetworkSubnetUsageClientExecuteResponse struct {
	VirtualNetworkSubnetUsageResult
}
