//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential azcore.TokenCredential
	options *arm.ClientOptions
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
		subscriptionID: 	subscriptionID,		credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewAPIClient() *APIClient {
	subClient, _ := NewAPIClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIDiagnosticClient() *APIDiagnosticClient {
	subClient, _ := NewAPIDiagnosticClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIExportClient() *APIExportClient {
	subClient, _ := NewAPIExportClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIIssueAttachmentClient() *APIIssueAttachmentClient {
	subClient, _ := NewAPIIssueAttachmentClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIIssueClient() *APIIssueClient {
	subClient, _ := NewAPIIssueClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIIssueCommentClient() *APIIssueCommentClient {
	subClient, _ := NewAPIIssueCommentClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIOperationClient() *APIOperationClient {
	subClient, _ := NewAPIOperationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIOperationPolicyClient() *APIOperationPolicyClient {
	subClient, _ := NewAPIOperationPolicyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIPolicyClient() *APIPolicyClient {
	subClient, _ := NewAPIPolicyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIProductClient() *APIProductClient {
	subClient, _ := NewAPIProductClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIReleaseClient() *APIReleaseClient {
	subClient, _ := NewAPIReleaseClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIRevisionClient() *APIRevisionClient {
	subClient, _ := NewAPIRevisionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPISchemaClient() *APISchemaClient {
	subClient, _ := NewAPISchemaClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPITagDescriptionClient() *APITagDescriptionClient {
	subClient, _ := NewAPITagDescriptionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIVersionSetClient() *APIVersionSetClient {
	subClient, _ := NewAPIVersionSetClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIWikiClient() *APIWikiClient {
	subClient, _ := NewAPIWikiClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAPIWikisClient() *APIWikisClient {
	subClient, _ := NewAPIWikisClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAuthorizationAccessPolicyClient() *AuthorizationAccessPolicyClient {
	subClient, _ := NewAuthorizationAccessPolicyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAuthorizationClient() *AuthorizationClient {
	subClient, _ := NewAuthorizationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAuthorizationLoginLinksClient() *AuthorizationLoginLinksClient {
	subClient, _ := NewAuthorizationLoginLinksClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAuthorizationProviderClient() *AuthorizationProviderClient {
	subClient, _ := NewAuthorizationProviderClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAuthorizationServerClient() *AuthorizationServerClient {
	subClient, _ := NewAuthorizationServerClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackendClient() *BackendClient {
	subClient, _ := NewBackendClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCacheClient() *CacheClient {
	subClient, _ := NewCacheClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCertificateClient() *CertificateClient {
	subClient, _ := NewCertificateClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContentItemClient() *ContentItemClient {
	subClient, _ := NewContentItemClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContentTypeClient() *ContentTypeClient {
	subClient, _ := NewContentTypeClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDelegationSettingsClient() *DelegationSettingsClient {
	subClient, _ := NewDelegationSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeletedServicesClient() *DeletedServicesClient {
	subClient, _ := NewDeletedServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDiagnosticClient() *DiagnosticClient {
	subClient, _ := NewDiagnosticClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDocumentationClient() *DocumentationClient {
	subClient, _ := NewDocumentationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewEmailTemplateClient() *EmailTemplateClient {
	subClient, _ := NewEmailTemplateClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGatewayAPIClient() *GatewayAPIClient {
	subClient, _ := NewGatewayAPIClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGatewayCertificateAuthorityClient() *GatewayCertificateAuthorityClient {
	subClient, _ := NewGatewayCertificateAuthorityClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGatewayClient() *GatewayClient {
	subClient, _ := NewGatewayClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGatewayHostnameConfigurationClient() *GatewayHostnameConfigurationClient {
	subClient, _ := NewGatewayHostnameConfigurationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGlobalSchemaClient() *GlobalSchemaClient {
	subClient, _ := NewGlobalSchemaClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGraphQLAPIResolverClient() *GraphQLAPIResolverClient {
	subClient, _ := NewGraphQLAPIResolverClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGraphQLAPIResolverPolicyClient() *GraphQLAPIResolverPolicyClient {
	subClient, _ := NewGraphQLAPIResolverPolicyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGroupClient() *GroupClient {
	subClient, _ := NewGroupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewGroupUserClient() *GroupUserClient {
	subClient, _ := NewGroupUserClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIdentityProviderClient() *IdentityProviderClient {
	subClient, _ := NewIdentityProviderClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewIssueClient() *IssueClient {
	subClient, _ := NewIssueClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewLoggerClient() *LoggerClient {
	subClient, _ := NewLoggerClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNamedValueClient() *NamedValueClient {
	subClient, _ := NewNamedValueClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNetworkStatusClient() *NetworkStatusClient {
	subClient, _ := NewNetworkStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNotificationClient() *NotificationClient {
	subClient, _ := NewNotificationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNotificationRecipientEmailClient() *NotificationRecipientEmailClient {
	subClient, _ := NewNotificationRecipientEmailClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNotificationRecipientUserClient() *NotificationRecipientUserClient {
	subClient, _ := NewNotificationRecipientUserClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOpenIDConnectProviderClient() *OpenIDConnectProviderClient {
	subClient, _ := NewOpenIDConnectProviderClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationClient() *OperationClient {
	subClient, _ := NewOperationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOutboundNetworkDependenciesEndpointsClient() *OutboundNetworkDependenciesEndpointsClient {
	subClient, _ := NewOutboundNetworkDependenciesEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPolicyClient() *PolicyClient {
	subClient, _ := NewPolicyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPolicyDescriptionClient() *PolicyDescriptionClient {
	subClient, _ := NewPolicyDescriptionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPolicyFragmentClient() *PolicyFragmentClient {
	subClient, _ := NewPolicyFragmentClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPortalConfigClient() *PortalConfigClient {
	subClient, _ := NewPortalConfigClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPortalRevisionClient() *PortalRevisionClient {
	subClient, _ := NewPortalRevisionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPortalSettingsClient() *PortalSettingsClient {
	subClient, _ := NewPortalSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionClient() *PrivateEndpointConnectionClient {
	subClient, _ := NewPrivateEndpointConnectionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductAPIClient() *ProductAPIClient {
	subClient, _ := NewProductAPIClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductClient() *ProductClient {
	subClient, _ := NewProductClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductGroupClient() *ProductGroupClient {
	subClient, _ := NewProductGroupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductPolicyClient() *ProductPolicyClient {
	subClient, _ := NewProductPolicyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductSubscriptionsClient() *ProductSubscriptionsClient {
	subClient, _ := NewProductSubscriptionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductWikiClient() *ProductWikiClient {
	subClient, _ := NewProductWikiClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProductWikisClient() *ProductWikisClient {
	subClient, _ := NewProductWikisClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewQuotaByCounterKeysClient() *QuotaByCounterKeysClient {
	subClient, _ := NewQuotaByCounterKeysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewQuotaByPeriodKeysClient() *QuotaByPeriodKeysClient {
	subClient, _ := NewQuotaByPeriodKeysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRegionClient() *RegionClient {
	subClient, _ := NewRegionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewReportsClient() *ReportsClient {
	subClient, _ := NewReportsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSKUsClient() *SKUsClient {
	subClient, _ := NewSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServiceClient() *ServiceClient {
	subClient, _ := NewServiceClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewServiceSKUsClient() *ServiceSKUsClient {
	subClient, _ := NewServiceSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSignInSettingsClient() *SignInSettingsClient {
	subClient, _ := NewSignInSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSignUpSettingsClient() *SignUpSettingsClient {
	subClient, _ := NewSignUpSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSubscriptionClient() *SubscriptionClient {
	subClient, _ := NewSubscriptionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTagClient() *TagClient {
	subClient, _ := NewTagClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTagResourceClient() *TagResourceClient {
	subClient, _ := NewTagResourceClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTenantAccessClient() *TenantAccessClient {
	subClient, _ := NewTenantAccessClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTenantAccessGitClient() *TenantAccessGitClient {
	subClient, _ := NewTenantAccessGitClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTenantConfigurationClient() *TenantConfigurationClient {
	subClient, _ := NewTenantConfigurationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTenantSettingsClient() *TenantSettingsClient {
	subClient, _ := NewTenantSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUserClient() *UserClient {
	subClient, _ := NewUserClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUserConfirmationPasswordClient() *UserConfirmationPasswordClient {
	subClient, _ := NewUserConfirmationPasswordClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUserGroupClient() *UserGroupClient {
	subClient, _ := NewUserGroupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUserIdentitiesClient() *UserIdentitiesClient {
	subClient, _ := NewUserIdentitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUserSubscriptionClient() *UserSubscriptionClient {
	subClient, _ := NewUserSubscriptionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

