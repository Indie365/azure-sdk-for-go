//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysql

// AdvisorsClientGetResponse contains the response from method AdvisorsClient.Get.
type AdvisorsClientGetResponse struct {
	Advisor
}

// AdvisorsClientListByServerResponse contains the response from method AdvisorsClient.NewListByServerPager.
type AdvisorsClientListByServerResponse struct {
	AdvisorsResultList
}

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	NameAvailability
}

// ConfigurationsClientCreateOrUpdateResponse contains the response from method ConfigurationsClient.BeginCreateOrUpdate.
type ConfigurationsClientCreateOrUpdateResponse struct {
	Configuration
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	Configuration
}

// ConfigurationsClientListByServerResponse contains the response from method ConfigurationsClient.NewListByServerPager.
type ConfigurationsClientListByServerResponse struct {
	ConfigurationListResult
}

// DatabasesClientCreateOrUpdateResponse contains the response from method DatabasesClient.BeginCreateOrUpdate.
type DatabasesClientCreateOrUpdateResponse struct {
	Database
}

// DatabasesClientDeleteResponse contains the response from method DatabasesClient.BeginDelete.
type DatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabasesClientGetResponse contains the response from method DatabasesClient.Get.
type DatabasesClientGetResponse struct {
	Database
}

// DatabasesClientListByServerResponse contains the response from method DatabasesClient.NewListByServerPager.
type DatabasesClientListByServerResponse struct {
	DatabaseListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.BeginCreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.BeginDelete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	FirewallRule
}

// FirewallRulesClientListByServerResponse contains the response from method FirewallRulesClient.NewListByServerPager.
type FirewallRulesClientListByServerResponse struct {
	FirewallRuleListResult
}

// LocationBasedPerformanceTierClientListResponse contains the response from method LocationBasedPerformanceTierClient.NewListPager.
type LocationBasedPerformanceTierClientListResponse struct {
	PerformanceTierListResult
}

// LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse contains the response from method LocationBasedRecommendedActionSessionsOperationStatusClient.Get.
type LocationBasedRecommendedActionSessionsOperationStatusClientGetResponse struct {
	RecommendedActionSessionsOperationStatus
}

// LocationBasedRecommendedActionSessionsResultClientListResponse contains the response from method LocationBasedRecommendedActionSessionsResultClient.NewListPager.
type LocationBasedRecommendedActionSessionsResultClientListResponse struct {
	RecommendationActionsResultList
}

// LogFilesClientListByServerResponse contains the response from method LogFilesClient.NewListByServerPager.
type LogFilesClientListByServerResponse struct {
	LogFileListResult
}

// ManagementClientCreateRecommendedActionSessionResponse contains the response from method ManagementClient.BeginCreateRecommendedActionSession.
type ManagementClientCreateRecommendedActionSessionResponse struct {
	// placeholder for future response values
}

// ManagementClientResetQueryPerformanceInsightDataResponse contains the response from method ManagementClient.ResetQueryPerformanceInsightData.
type ManagementClientResetQueryPerformanceInsightDataResponse struct {
	QueryPerformanceInsightResetDataResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByServerResponse contains the response from method PrivateEndpointConnectionsClient.NewListByServerPager.
type PrivateEndpointConnectionsClientListByServerResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateTagsResponse contains the response from method PrivateEndpointConnectionsClient.BeginUpdateTags.
type PrivateEndpointConnectionsClientUpdateTagsResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByServerResponse contains the response from method PrivateLinkResourcesClient.NewListByServerPager.
type PrivateLinkResourcesClientListByServerResponse struct {
	PrivateLinkResourceListResult
}

// QueryTextsClientGetResponse contains the response from method QueryTextsClient.Get.
type QueryTextsClientGetResponse struct {
	QueryText
}

// QueryTextsClientListByServerResponse contains the response from method QueryTextsClient.NewListByServerPager.
type QueryTextsClientListByServerResponse struct {
	QueryTextsResultList
}

// RecommendedActionsClientGetResponse contains the response from method RecommendedActionsClient.Get.
type RecommendedActionsClientGetResponse struct {
	RecommendationAction
}

// RecommendedActionsClientListByServerResponse contains the response from method RecommendedActionsClient.NewListByServerPager.
type RecommendedActionsClientListByServerResponse struct {
	RecommendationActionsResultList
}

// RecoverableServersClientGetResponse contains the response from method RecoverableServersClient.Get.
type RecoverableServersClientGetResponse struct {
	RecoverableServerResource
}

// ReplicasClientListByServerResponse contains the response from method ReplicasClient.NewListByServerPager.
type ReplicasClientListByServerResponse struct {
	ServerListResult
}

// ServerAdministratorsClientCreateOrUpdateResponse contains the response from method ServerAdministratorsClient.BeginCreateOrUpdate.
type ServerAdministratorsClientCreateOrUpdateResponse struct {
	ServerAdministratorResource
}

// ServerAdministratorsClientDeleteResponse contains the response from method ServerAdministratorsClient.BeginDelete.
type ServerAdministratorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerAdministratorsClientGetResponse contains the response from method ServerAdministratorsClient.Get.
type ServerAdministratorsClientGetResponse struct {
	ServerAdministratorResource
}

// ServerAdministratorsClientListResponse contains the response from method ServerAdministratorsClient.NewListPager.
type ServerAdministratorsClientListResponse struct {
	ServerAdministratorResourceListResult
}

// ServerBasedPerformanceTierClientListResponse contains the response from method ServerBasedPerformanceTierClient.NewListPager.
type ServerBasedPerformanceTierClientListResponse struct {
	PerformanceTierListResult
}

// ServerKeysClientCreateOrUpdateResponse contains the response from method ServerKeysClient.BeginCreateOrUpdate.
type ServerKeysClientCreateOrUpdateResponse struct {
	ServerKey
}

// ServerKeysClientDeleteResponse contains the response from method ServerKeysClient.BeginDelete.
type ServerKeysClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerKeysClientGetResponse contains the response from method ServerKeysClient.Get.
type ServerKeysClientGetResponse struct {
	ServerKey
}

// ServerKeysClientListResponse contains the response from method ServerKeysClient.NewListPager.
type ServerKeysClientListResponse struct {
	ServerKeyListResult
}

// ServerParametersClientListUpdateConfigurationsResponse contains the response from method ServerParametersClient.BeginListUpdateConfigurations.
type ServerParametersClientListUpdateConfigurationsResponse struct {
	ConfigurationListResult
}

// ServerSecurityAlertPoliciesClientCreateOrUpdateResponse contains the response from method ServerSecurityAlertPoliciesClient.BeginCreateOrUpdate.
type ServerSecurityAlertPoliciesClientCreateOrUpdateResponse struct {
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesClientGetResponse contains the response from method ServerSecurityAlertPoliciesClient.Get.
type ServerSecurityAlertPoliciesClientGetResponse struct {
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesClientListByServerResponse contains the response from method ServerSecurityAlertPoliciesClient.NewListByServerPager.
type ServerSecurityAlertPoliciesClientListByServerResponse struct {
	ServerSecurityAlertPolicyListResult
}

// ServersClientCreateResponse contains the response from method ServersClient.BeginCreate.
type ServersClientCreateResponse struct {
	Server
}

// ServersClientDeleteResponse contains the response from method ServersClient.BeginDelete.
type ServersClientDeleteResponse struct {
	// placeholder for future response values
}

// ServersClientGetResponse contains the response from method ServersClient.Get.
type ServersClientGetResponse struct {
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.NewListByResourceGroupPager.
type ServersClientListByResourceGroupResponse struct {
	ServerListResult
}

// ServersClientListResponse contains the response from method ServersClient.NewListPager.
type ServersClientListResponse struct {
	ServerListResult
}

// ServersClientRestartResponse contains the response from method ServersClient.BeginRestart.
type ServersClientRestartResponse struct {
	// placeholder for future response values
}

// ServersClientStartResponse contains the response from method ServersClient.BeginStart.
type ServersClientStartResponse struct {
	// placeholder for future response values
}

// ServersClientStopResponse contains the response from method ServersClient.BeginStop.
type ServersClientStopResponse struct {
	// placeholder for future response values
}

// ServersClientUpdateResponse contains the response from method ServersClient.BeginUpdate.
type ServersClientUpdateResponse struct {
	Server
}

// ServersClientUpgradeResponse contains the response from method ServersClient.BeginUpgrade.
type ServersClientUpgradeResponse struct {
	// placeholder for future response values
}

// TopQueryStatisticsClientGetResponse contains the response from method TopQueryStatisticsClient.Get.
type TopQueryStatisticsClientGetResponse struct {
	QueryStatistic
}

// TopQueryStatisticsClientListByServerResponse contains the response from method TopQueryStatisticsClient.NewListByServerPager.
type TopQueryStatisticsClientListByServerResponse struct {
	TopQueryStatisticsResultList
}

// VirtualNetworkRulesClientCreateOrUpdateResponse contains the response from method VirtualNetworkRulesClient.BeginCreateOrUpdate.
type VirtualNetworkRulesClientCreateOrUpdateResponse struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesClientDeleteResponse contains the response from method VirtualNetworkRulesClient.BeginDelete.
type VirtualNetworkRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworkRulesClientGetResponse contains the response from method VirtualNetworkRulesClient.Get.
type VirtualNetworkRulesClientGetResponse struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesClientListByServerResponse contains the response from method VirtualNetworkRulesClient.NewListByServerPager.
type VirtualNetworkRulesClientListByServerResponse struct {
	VirtualNetworkRuleListResult
}

// WaitStatisticsClientGetResponse contains the response from method WaitStatisticsClient.Get.
type WaitStatisticsClientGetResponse struct {
	WaitStatistic
}

// WaitStatisticsClientListByServerResponse contains the response from method WaitStatisticsClient.NewListByServerPager.
type WaitStatisticsClientListByServerResponse struct {
	WaitStatisticsResultList
}
