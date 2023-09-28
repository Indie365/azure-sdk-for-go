//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewActionsClient() *ActionsClient {
	subClient, _ := NewActionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertRuleClient() *AlertRuleClient {
	subClient, _ := NewAlertRuleClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertRuleTemplatesClient() *AlertRuleTemplatesClient {
	subClient, _ := NewAlertRuleTemplatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertRulesClient() *AlertRulesClient {
	subClient, _ := NewAlertRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAutomationRulesClient() *AutomationRulesClient {
	subClient, _ := NewAutomationRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBillingStatisticsClient() *BillingStatisticsClient {
	subClient, _ := NewBillingStatisticsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBookmarkClient() *BookmarkClient {
	subClient, _ := NewBookmarkClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBookmarkRelationsClient() *BookmarkRelationsClient {
	subClient, _ := NewBookmarkRelationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBookmarksClient() *BookmarksClient {
	subClient, _ := NewBookmarksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContentPackageClient() *ContentPackageClient {
	subClient, _ := NewContentPackageClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContentPackagesClient() *ContentPackagesClient {
	subClient, _ := NewContentPackagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContentTemplateClient() *ContentTemplateClient {
	subClient, _ := NewContentTemplateClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContentTemplatesClient() *ContentTemplatesClient {
	subClient, _ := NewContentTemplatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataConnectorDefinitionsClient() *DataConnectorDefinitionsClient {
	subClient, _ := NewDataConnectorDefinitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataConnectorsCheckRequirementsClient() *DataConnectorsCheckRequirementsClient {
	subClient, _ := NewDataConnectorsCheckRequirementsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataConnectorsClient() *DataConnectorsClient {
	subClient, _ := NewDataConnectorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDomainWhoisClient() *DomainWhoisClient {
	subClient, _ := NewDomainWhoisClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEntitiesClient() *EntitiesClient {
	subClient, _ := NewEntitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEntitiesGetTimelineClient() *EntitiesGetTimelineClient {
	subClient, _ := NewEntitiesGetTimelineClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEntitiesRelationsClient() *EntitiesRelationsClient {
	subClient, _ := NewEntitiesRelationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEntityQueriesClient() *EntityQueriesClient {
	subClient, _ := NewEntityQueriesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEntityQueryTemplatesClient() *EntityQueryTemplatesClient {
	subClient, _ := NewEntityQueryTemplatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEntityRelationsClient() *EntityRelationsClient {
	subClient, _ := NewEntityRelationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFileImportsClient() *FileImportsClient {
	subClient, _ := NewFileImportsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGetClient() *GetClient {
	subClient, _ := NewGetClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGetRecommendationsClient() *GetRecommendationsClient {
	subClient, _ := NewGetRecommendationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGetTriggeredAnalyticsRuleRunsClient() *GetTriggeredAnalyticsRuleRunsClient {
	subClient, _ := NewGetTriggeredAnalyticsRuleRunsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHuntCommentsClient() *HuntCommentsClient {
	subClient, _ := NewHuntCommentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHuntRelationsClient() *HuntRelationsClient {
	subClient, _ := NewHuntRelationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHuntsClient() *HuntsClient {
	subClient, _ := NewHuntsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIPGeodataClient() *IPGeodataClient {
	subClient, _ := NewIPGeodataClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIncidentCommentsClient() *IncidentCommentsClient {
	subClient, _ := NewIncidentCommentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIncidentRelationsClient() *IncidentRelationsClient {
	subClient, _ := NewIncidentRelationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIncidentTasksClient() *IncidentTasksClient {
	subClient, _ := NewIncidentTasksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIncidentsClient() *IncidentsClient {
	subClient, _ := NewIncidentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewMetadataClient() *MetadataClient {
	subClient, _ := NewMetadataClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOfficeConsentsClient() *OfficeConsentsClient {
	subClient, _ := NewOfficeConsentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductPackageClient() *ProductPackageClient {
	subClient, _ := NewProductPackageClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductPackagesClient() *ProductPackagesClient {
	subClient, _ := NewProductPackagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductSettingsClient() *ProductSettingsClient {
	subClient, _ := NewProductSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductTemplateClient() *ProductTemplateClient {
	subClient, _ := NewProductTemplateClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductTemplatesClient() *ProductTemplatesClient {
	subClient, _ := NewProductTemplatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecurityMLAnalyticsSettingsClient() *SecurityMLAnalyticsSettingsClient {
	subClient, _ := NewSecurityMLAnalyticsSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSentinelOnboardingStatesClient() *SentinelOnboardingStatesClient {
	subClient, _ := NewSentinelOnboardingStatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSourceControlClient() *SourceControlClient {
	subClient, _ := NewSourceControlClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSourceControlsClient() *SourceControlsClient {
	subClient, _ := NewSourceControlsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewThreatIntelligenceIndicatorClient() *ThreatIntelligenceIndicatorClient {
	subClient, _ := NewThreatIntelligenceIndicatorClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewThreatIntelligenceIndicatorMetricsClient() *ThreatIntelligenceIndicatorMetricsClient {
	subClient, _ := NewThreatIntelligenceIndicatorMetricsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewThreatIntelligenceIndicatorsClient() *ThreatIntelligenceIndicatorsClient {
	subClient, _ := NewThreatIntelligenceIndicatorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTriggeredAnalyticsRuleRunClient() *TriggeredAnalyticsRuleRunClient {
	subClient, _ := NewTriggeredAnalyticsRuleRunClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUpdateClient() *UpdateClient {
	subClient, _ := NewUpdateClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWatchlistItemsClient() *WatchlistItemsClient {
	subClient, _ := NewWatchlistItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWatchlistsClient() *WatchlistsClient {
	subClient, _ := NewWatchlistsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceManagerAssignmentJobsClient() *WorkspaceManagerAssignmentJobsClient {
	subClient, _ := NewWorkspaceManagerAssignmentJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceManagerAssignmentsClient() *WorkspaceManagerAssignmentsClient {
	subClient, _ := NewWorkspaceManagerAssignmentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceManagerConfigurationsClient() *WorkspaceManagerConfigurationsClient {
	subClient, _ := NewWorkspaceManagerConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceManagerGroupsClient() *WorkspaceManagerGroupsClient {
	subClient, _ := NewWorkspaceManagerGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewWorkspaceManagerMembersClient() *WorkspaceManagerMembersClient {
	subClient, _ := NewWorkspaceManagerMembersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
