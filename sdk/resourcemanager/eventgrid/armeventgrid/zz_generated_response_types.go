//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

// ChannelsClientCreateOrUpdateResponse contains the response from method ChannelsClient.CreateOrUpdate.
type ChannelsClientCreateOrUpdateResponse struct {
	Channel
}

// ChannelsClientDeleteResponse contains the response from method ChannelsClient.Delete.
type ChannelsClientDeleteResponse struct {
	// placeholder for future response values
}

// ChannelsClientGetFullURLResponse contains the response from method ChannelsClient.GetFullURL.
type ChannelsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// ChannelsClientGetResponse contains the response from method ChannelsClient.Get.
type ChannelsClientGetResponse struct {
	Channel
}

// ChannelsClientListByPartnerNamespaceResponse contains the response from method ChannelsClient.ListByPartnerNamespace.
type ChannelsClientListByPartnerNamespaceResponse struct {
	ChannelsListResult
}

// ChannelsClientUpdateResponse contains the response from method ChannelsClient.Update.
type ChannelsClientUpdateResponse struct {
	// placeholder for future response values
}

// DomainEventSubscriptionsClientCreateOrUpdateResponse contains the response from method DomainEventSubscriptionsClient.CreateOrUpdate.
type DomainEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// DomainEventSubscriptionsClientDeleteResponse contains the response from method DomainEventSubscriptionsClient.Delete.
type DomainEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method DomainEventSubscriptionsClient.GetDeliveryAttributes.
type DomainEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// DomainEventSubscriptionsClientGetFullURLResponse contains the response from method DomainEventSubscriptionsClient.GetFullURL.
type DomainEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// DomainEventSubscriptionsClientGetResponse contains the response from method DomainEventSubscriptionsClient.Get.
type DomainEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// DomainEventSubscriptionsClientListResponse contains the response from method DomainEventSubscriptionsClient.List.
type DomainEventSubscriptionsClientListResponse struct {
	EventSubscriptionsListResult
}

// DomainEventSubscriptionsClientUpdateResponse contains the response from method DomainEventSubscriptionsClient.Update.
type DomainEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// DomainTopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method DomainTopicEventSubscriptionsClient.CreateOrUpdate.
type DomainTopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// DomainTopicEventSubscriptionsClientDeleteResponse contains the response from method DomainTopicEventSubscriptionsClient.Delete.
type DomainTopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainTopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method DomainTopicEventSubscriptionsClient.GetDeliveryAttributes.
type DomainTopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// DomainTopicEventSubscriptionsClientGetFullURLResponse contains the response from method DomainTopicEventSubscriptionsClient.GetFullURL.
type DomainTopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// DomainTopicEventSubscriptionsClientGetResponse contains the response from method DomainTopicEventSubscriptionsClient.Get.
type DomainTopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// DomainTopicEventSubscriptionsClientListResponse contains the response from method DomainTopicEventSubscriptionsClient.List.
type DomainTopicEventSubscriptionsClientListResponse struct {
	EventSubscriptionsListResult
}

// DomainTopicEventSubscriptionsClientUpdateResponse contains the response from method DomainTopicEventSubscriptionsClient.Update.
type DomainTopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// DomainTopicsClientCreateOrUpdateResponse contains the response from method DomainTopicsClient.CreateOrUpdate.
type DomainTopicsClientCreateOrUpdateResponse struct {
	DomainTopic
}

// DomainTopicsClientDeleteResponse contains the response from method DomainTopicsClient.Delete.
type DomainTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainTopicsClientGetResponse contains the response from method DomainTopicsClient.Get.
type DomainTopicsClientGetResponse struct {
	DomainTopic
}

// DomainTopicsClientListByDomainResponse contains the response from method DomainTopicsClient.ListByDomain.
type DomainTopicsClientListByDomainResponse struct {
	DomainTopicsListResult
}

// DomainsClientCreateOrUpdateResponse contains the response from method DomainsClient.CreateOrUpdate.
type DomainsClientCreateOrUpdateResponse struct {
	Domain
}

// DomainsClientDeleteResponse contains the response from method DomainsClient.Delete.
type DomainsClientDeleteResponse struct {
	// placeholder for future response values
}

// DomainsClientGetResponse contains the response from method DomainsClient.Get.
type DomainsClientGetResponse struct {
	Domain
}

// DomainsClientListByResourceGroupResponse contains the response from method DomainsClient.ListByResourceGroup.
type DomainsClientListByResourceGroupResponse struct {
	DomainsListResult
}

// DomainsClientListBySubscriptionResponse contains the response from method DomainsClient.ListBySubscription.
type DomainsClientListBySubscriptionResponse struct {
	DomainsListResult
}

// DomainsClientListSharedAccessKeysResponse contains the response from method DomainsClient.ListSharedAccessKeys.
type DomainsClientListSharedAccessKeysResponse struct {
	DomainSharedAccessKeys
}

// DomainsClientRegenerateKeyResponse contains the response from method DomainsClient.RegenerateKey.
type DomainsClientRegenerateKeyResponse struct {
	DomainSharedAccessKeys
}

// DomainsClientUpdateResponse contains the response from method DomainsClient.Update.
type DomainsClientUpdateResponse struct {
	Domain
}

// EventSubscriptionsClientCreateOrUpdateResponse contains the response from method EventSubscriptionsClient.CreateOrUpdate.
type EventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// EventSubscriptionsClientDeleteResponse contains the response from method EventSubscriptionsClient.Delete.
type EventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// EventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method EventSubscriptionsClient.GetDeliveryAttributes.
type EventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// EventSubscriptionsClientGetFullURLResponse contains the response from method EventSubscriptionsClient.GetFullURL.
type EventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// EventSubscriptionsClientGetResponse contains the response from method EventSubscriptionsClient.Get.
type EventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// EventSubscriptionsClientListByDomainTopicResponse contains the response from method EventSubscriptionsClient.ListByDomainTopic.
type EventSubscriptionsClientListByDomainTopicResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListByResourceResponse contains the response from method EventSubscriptionsClient.ListByResource.
type EventSubscriptionsClientListByResourceResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalByResourceGroupForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListGlobalByResourceGroupForTopicType.
type EventSubscriptionsClientListGlobalByResourceGroupForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalByResourceGroupResponse contains the response from method EventSubscriptionsClient.ListGlobalByResourceGroup.
type EventSubscriptionsClientListGlobalByResourceGroupResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalBySubscriptionForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListGlobalBySubscriptionForTopicType.
type EventSubscriptionsClientListGlobalBySubscriptionForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListGlobalBySubscriptionResponse contains the response from method EventSubscriptionsClient.ListGlobalBySubscription.
type EventSubscriptionsClientListGlobalBySubscriptionResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalByResourceGroupForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListRegionalByResourceGroupForTopicType.
type EventSubscriptionsClientListRegionalByResourceGroupForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalByResourceGroupResponse contains the response from method EventSubscriptionsClient.ListRegionalByResourceGroup.
type EventSubscriptionsClientListRegionalByResourceGroupResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalBySubscriptionForTopicTypeResponse contains the response from method EventSubscriptionsClient.ListRegionalBySubscriptionForTopicType.
type EventSubscriptionsClientListRegionalBySubscriptionForTopicTypeResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientListRegionalBySubscriptionResponse contains the response from method EventSubscriptionsClient.ListRegionalBySubscription.
type EventSubscriptionsClientListRegionalBySubscriptionResponse struct {
	EventSubscriptionsListResult
}

// EventSubscriptionsClientUpdateResponse contains the response from method EventSubscriptionsClient.Update.
type EventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// ExtensionTopicsClientGetResponse contains the response from method ExtensionTopicsClient.Get.
type ExtensionTopicsClientGetResponse struct {
	ExtensionTopic
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsListResult
}

// PartnerConfigurationsClientAuthorizePartnerResponse contains the response from method PartnerConfigurationsClient.AuthorizePartner.
type PartnerConfigurationsClientAuthorizePartnerResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientCreateOrUpdateResponse contains the response from method PartnerConfigurationsClient.CreateOrUpdate.
type PartnerConfigurationsClientCreateOrUpdateResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientDeleteResponse contains the response from method PartnerConfigurationsClient.Delete.
type PartnerConfigurationsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerConfigurationsClientGetResponse contains the response from method PartnerConfigurationsClient.Get.
type PartnerConfigurationsClientGetResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientListByResourceGroupResponse contains the response from method PartnerConfigurationsClient.ListByResourceGroup.
type PartnerConfigurationsClientListByResourceGroupResponse struct {
	PartnerConfigurationsListResult
}

// PartnerConfigurationsClientListBySubscriptionResponse contains the response from method PartnerConfigurationsClient.ListBySubscription.
type PartnerConfigurationsClientListBySubscriptionResponse struct {
	PartnerConfigurationsListResult
}

// PartnerConfigurationsClientUnauthorizePartnerResponse contains the response from method PartnerConfigurationsClient.UnauthorizePartner.
type PartnerConfigurationsClientUnauthorizePartnerResponse struct {
	PartnerConfiguration
}

// PartnerConfigurationsClientUpdateResponse contains the response from method PartnerConfigurationsClient.Update.
type PartnerConfigurationsClientUpdateResponse struct {
	PartnerConfiguration
}

// PartnerNamespacesClientCreateOrUpdateResponse contains the response from method PartnerNamespacesClient.CreateOrUpdate.
type PartnerNamespacesClientCreateOrUpdateResponse struct {
	PartnerNamespace
}

// PartnerNamespacesClientDeleteResponse contains the response from method PartnerNamespacesClient.Delete.
type PartnerNamespacesClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerNamespacesClientGetResponse contains the response from method PartnerNamespacesClient.Get.
type PartnerNamespacesClientGetResponse struct {
	PartnerNamespace
}

// PartnerNamespacesClientListByResourceGroupResponse contains the response from method PartnerNamespacesClient.ListByResourceGroup.
type PartnerNamespacesClientListByResourceGroupResponse struct {
	PartnerNamespacesListResult
}

// PartnerNamespacesClientListBySubscriptionResponse contains the response from method PartnerNamespacesClient.ListBySubscription.
type PartnerNamespacesClientListBySubscriptionResponse struct {
	PartnerNamespacesListResult
}

// PartnerNamespacesClientListSharedAccessKeysResponse contains the response from method PartnerNamespacesClient.ListSharedAccessKeys.
type PartnerNamespacesClientListSharedAccessKeysResponse struct {
	PartnerNamespaceSharedAccessKeys
}

// PartnerNamespacesClientRegenerateKeyResponse contains the response from method PartnerNamespacesClient.RegenerateKey.
type PartnerNamespacesClientRegenerateKeyResponse struct {
	PartnerNamespaceSharedAccessKeys
}

// PartnerNamespacesClientUpdateResponse contains the response from method PartnerNamespacesClient.Update.
type PartnerNamespacesClientUpdateResponse struct {
	PartnerNamespace
}

// PartnerRegistrationsClientCreateOrUpdateResponse contains the response from method PartnerRegistrationsClient.CreateOrUpdate.
type PartnerRegistrationsClientCreateOrUpdateResponse struct {
	PartnerRegistration
}

// PartnerRegistrationsClientDeleteResponse contains the response from method PartnerRegistrationsClient.Delete.
type PartnerRegistrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerRegistrationsClientGetResponse contains the response from method PartnerRegistrationsClient.Get.
type PartnerRegistrationsClientGetResponse struct {
	PartnerRegistration
}

// PartnerRegistrationsClientListByResourceGroupResponse contains the response from method PartnerRegistrationsClient.ListByResourceGroup.
type PartnerRegistrationsClientListByResourceGroupResponse struct {
	PartnerRegistrationsListResult
}

// PartnerRegistrationsClientListBySubscriptionResponse contains the response from method PartnerRegistrationsClient.ListBySubscription.
type PartnerRegistrationsClientListBySubscriptionResponse struct {
	PartnerRegistrationsListResult
}

// PartnerRegistrationsClientUpdateResponse contains the response from method PartnerRegistrationsClient.Update.
type PartnerRegistrationsClientUpdateResponse struct {
	PartnerRegistration
}

// PartnerTopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method PartnerTopicEventSubscriptionsClient.CreateOrUpdate.
type PartnerTopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// PartnerTopicEventSubscriptionsClientDeleteResponse contains the response from method PartnerTopicEventSubscriptionsClient.Delete.
type PartnerTopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method PartnerTopicEventSubscriptionsClient.GetDeliveryAttributes.
type PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// PartnerTopicEventSubscriptionsClientGetFullURLResponse contains the response from method PartnerTopicEventSubscriptionsClient.GetFullURL.
type PartnerTopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// PartnerTopicEventSubscriptionsClientGetResponse contains the response from method PartnerTopicEventSubscriptionsClient.Get.
type PartnerTopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse contains the response from method PartnerTopicEventSubscriptionsClient.ListByPartnerTopic.
type PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse struct {
	EventSubscriptionsListResult
}

// PartnerTopicEventSubscriptionsClientUpdateResponse contains the response from method PartnerTopicEventSubscriptionsClient.Update.
type PartnerTopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// PartnerTopicsClientActivateResponse contains the response from method PartnerTopicsClient.Activate.
type PartnerTopicsClientActivateResponse struct {
	PartnerTopic
}

// PartnerTopicsClientCreateOrUpdateResponse contains the response from method PartnerTopicsClient.CreateOrUpdate.
type PartnerTopicsClientCreateOrUpdateResponse struct {
	PartnerTopic
}

// PartnerTopicsClientDeactivateResponse contains the response from method PartnerTopicsClient.Deactivate.
type PartnerTopicsClientDeactivateResponse struct {
	PartnerTopic
}

// PartnerTopicsClientDeleteResponse contains the response from method PartnerTopicsClient.Delete.
type PartnerTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// PartnerTopicsClientGetResponse contains the response from method PartnerTopicsClient.Get.
type PartnerTopicsClientGetResponse struct {
	PartnerTopic
}

// PartnerTopicsClientListByResourceGroupResponse contains the response from method PartnerTopicsClient.ListByResourceGroup.
type PartnerTopicsClientListByResourceGroupResponse struct {
	PartnerTopicsListResult
}

// PartnerTopicsClientListBySubscriptionResponse contains the response from method PartnerTopicsClient.ListBySubscription.
type PartnerTopicsClientListBySubscriptionResponse struct {
	PartnerTopicsListResult
}

// PartnerTopicsClientUpdateResponse contains the response from method PartnerTopicsClient.Update.
type PartnerTopicsClientUpdateResponse struct {
	PartnerTopic
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByResourceResponse contains the response from method PrivateEndpointConnectionsClient.ListByResource.
type PrivateEndpointConnectionsClientListByResourceResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByResourceResponse contains the response from method PrivateLinkResourcesClient.ListByResource.
type PrivateLinkResourcesClientListByResourceResponse struct {
	PrivateLinkResourcesListResult
}

// SystemTopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method SystemTopicEventSubscriptionsClient.CreateOrUpdate.
type SystemTopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// SystemTopicEventSubscriptionsClientDeleteResponse contains the response from method SystemTopicEventSubscriptionsClient.Delete.
type SystemTopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method SystemTopicEventSubscriptionsClient.GetDeliveryAttributes.
type SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// SystemTopicEventSubscriptionsClientGetFullURLResponse contains the response from method SystemTopicEventSubscriptionsClient.GetFullURL.
type SystemTopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// SystemTopicEventSubscriptionsClientGetResponse contains the response from method SystemTopicEventSubscriptionsClient.Get.
type SystemTopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// SystemTopicEventSubscriptionsClientListBySystemTopicResponse contains the response from method SystemTopicEventSubscriptionsClient.ListBySystemTopic.
type SystemTopicEventSubscriptionsClientListBySystemTopicResponse struct {
	EventSubscriptionsListResult
}

// SystemTopicEventSubscriptionsClientUpdateResponse contains the response from method SystemTopicEventSubscriptionsClient.Update.
type SystemTopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// SystemTopicsClientCreateOrUpdateResponse contains the response from method SystemTopicsClient.CreateOrUpdate.
type SystemTopicsClientCreateOrUpdateResponse struct {
	SystemTopic
}

// SystemTopicsClientDeleteResponse contains the response from method SystemTopicsClient.Delete.
type SystemTopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// SystemTopicsClientGetResponse contains the response from method SystemTopicsClient.Get.
type SystemTopicsClientGetResponse struct {
	SystemTopic
}

// SystemTopicsClientListByResourceGroupResponse contains the response from method SystemTopicsClient.ListByResourceGroup.
type SystemTopicsClientListByResourceGroupResponse struct {
	SystemTopicsListResult
}

// SystemTopicsClientListBySubscriptionResponse contains the response from method SystemTopicsClient.ListBySubscription.
type SystemTopicsClientListBySubscriptionResponse struct {
	SystemTopicsListResult
}

// SystemTopicsClientUpdateResponse contains the response from method SystemTopicsClient.Update.
type SystemTopicsClientUpdateResponse struct {
	SystemTopic
}

// TopicEventSubscriptionsClientCreateOrUpdateResponse contains the response from method TopicEventSubscriptionsClient.CreateOrUpdate.
type TopicEventSubscriptionsClientCreateOrUpdateResponse struct {
	EventSubscription
}

// TopicEventSubscriptionsClientDeleteResponse contains the response from method TopicEventSubscriptionsClient.Delete.
type TopicEventSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// TopicEventSubscriptionsClientGetDeliveryAttributesResponse contains the response from method TopicEventSubscriptionsClient.GetDeliveryAttributes.
type TopicEventSubscriptionsClientGetDeliveryAttributesResponse struct {
	DeliveryAttributeListResult
}

// TopicEventSubscriptionsClientGetFullURLResponse contains the response from method TopicEventSubscriptionsClient.GetFullURL.
type TopicEventSubscriptionsClientGetFullURLResponse struct {
	EventSubscriptionFullURL
}

// TopicEventSubscriptionsClientGetResponse contains the response from method TopicEventSubscriptionsClient.Get.
type TopicEventSubscriptionsClientGetResponse struct {
	EventSubscription
}

// TopicEventSubscriptionsClientListResponse contains the response from method TopicEventSubscriptionsClient.List.
type TopicEventSubscriptionsClientListResponse struct {
	EventSubscriptionsListResult
}

// TopicEventSubscriptionsClientUpdateResponse contains the response from method TopicEventSubscriptionsClient.Update.
type TopicEventSubscriptionsClientUpdateResponse struct {
	EventSubscription
}

// TopicTypesClientGetResponse contains the response from method TopicTypesClient.Get.
type TopicTypesClientGetResponse struct {
	TopicTypeInfo
}

// TopicTypesClientListEventTypesResponse contains the response from method TopicTypesClient.ListEventTypes.
type TopicTypesClientListEventTypesResponse struct {
	EventTypesListResult
}

// TopicTypesClientListResponse contains the response from method TopicTypesClient.List.
type TopicTypesClientListResponse struct {
	TopicTypesListResult
}

// TopicsClientCreateOrUpdateResponse contains the response from method TopicsClient.CreateOrUpdate.
type TopicsClientCreateOrUpdateResponse struct {
	Topic
}

// TopicsClientDeleteResponse contains the response from method TopicsClient.Delete.
type TopicsClientDeleteResponse struct {
	// placeholder for future response values
}

// TopicsClientGetResponse contains the response from method TopicsClient.Get.
type TopicsClientGetResponse struct {
	Topic
}

// TopicsClientListByResourceGroupResponse contains the response from method TopicsClient.ListByResourceGroup.
type TopicsClientListByResourceGroupResponse struct {
	TopicsListResult
}

// TopicsClientListBySubscriptionResponse contains the response from method TopicsClient.ListBySubscription.
type TopicsClientListBySubscriptionResponse struct {
	TopicsListResult
}

// TopicsClientListEventTypesResponse contains the response from method TopicsClient.ListEventTypes.
type TopicsClientListEventTypesResponse struct {
	EventTypesListResult
}

// TopicsClientListSharedAccessKeysResponse contains the response from method TopicsClient.ListSharedAccessKeys.
type TopicsClientListSharedAccessKeysResponse struct {
	TopicSharedAccessKeys
}

// TopicsClientRegenerateKeyResponse contains the response from method TopicsClient.RegenerateKey.
type TopicsClientRegenerateKeyResponse struct {
	TopicSharedAccessKeys
}

// TopicsClientUpdateResponse contains the response from method TopicsClient.Update.
type TopicsClientUpdateResponse struct {
	Topic
}

// VerifiedPartnersClientGetResponse contains the response from method VerifiedPartnersClient.Get.
type VerifiedPartnersClientGetResponse struct {
	VerifiedPartner
}

// VerifiedPartnersClientListResponse contains the response from method VerifiedPartnersClient.List.
type VerifiedPartnersClientListResponse struct {
	VerifiedPartnersListResult
}