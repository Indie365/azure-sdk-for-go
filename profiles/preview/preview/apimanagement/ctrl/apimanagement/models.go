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

package apimanagement

import original "github.com/Azure/azure-sdk-for-go/services/preview/apimanagement/ctrl/2017-03-01/apimanagement"

type DelegationSettingsClient = original.DelegationSettingsClient
type ProductSubscriptionsClient = original.ProductSubscriptionsClient
type GroupUserClient = original.GroupUserClient
type ProductAPIClient = original.ProductAPIClient
type ProductClient = original.ProductClient
type ProductGroupClient = original.ProductGroupClient
type OpenIDConnectProviderClient = original.OpenIDConnectProviderClient
type SignInSettingsClient = original.SignInSettingsClient
type UserGroupClient = original.UserGroupClient
type PolicySnippetsClient = original.PolicySnippetsClient
type AuthorizationServerClient = original.AuthorizationServerClient
type APISchemaClient = original.APISchemaClient
type APIExportClient = original.APIExportClient
type BackendClient = original.BackendClient
type EmailTemplateClient = original.EmailTemplateClient
type LoggerClient = original.LoggerClient
type APIProductClient = original.APIProductClient
type APIOperationClient = original.APIOperationClient
type PropertyClient = original.PropertyClient
type UserIdentitiesClient = original.UserIdentitiesClient
type APIType = original.APIType

const (
	HTTP APIType = original.HTTP
	Soap APIType = original.Soap
)

type AsyncOperationStatus = original.AsyncOperationStatus

const (
	Failed     AsyncOperationStatus = original.Failed
	InProgress AsyncOperationStatus = original.InProgress
	Started    AsyncOperationStatus = original.Started
	Succeeded  AsyncOperationStatus = original.Succeeded
)

type AuthorizationMethod = original.AuthorizationMethod

const (
	DELETE  AuthorizationMethod = original.DELETE
	GET     AuthorizationMethod = original.GET
	HEAD    AuthorizationMethod = original.HEAD
	OPTIONS AuthorizationMethod = original.OPTIONS
	PATCH   AuthorizationMethod = original.PATCH
	POST    AuthorizationMethod = original.POST
	PUT     AuthorizationMethod = original.PUT
	TRACE   AuthorizationMethod = original.TRACE
)

type BackendProtocol = original.BackendProtocol

const (
	BackendProtocolHTTP BackendProtocol = original.BackendProtocolHTTP
	BackendProtocolSoap BackendProtocol = original.BackendProtocolSoap
)

type BearerTokenSendingMethod = original.BearerTokenSendingMethod

const (
	AuthorizationHeader BearerTokenSendingMethod = original.AuthorizationHeader
	Query               BearerTokenSendingMethod = original.Query
)

type ClientAuthenticationMethod = original.ClientAuthenticationMethod

const (
	Basic ClientAuthenticationMethod = original.Basic
	Body  ClientAuthenticationMethod = original.Body
)

type ContentFormat = original.ContentFormat

const (
	SwaggerJSON     ContentFormat = original.SwaggerJSON
	SwaggerLinkJSON ContentFormat = original.SwaggerLinkJSON
	WadlLinkJSON    ContentFormat = original.WadlLinkJSON
	WadlXML         ContentFormat = original.WadlXML
	Wsdl            ContentFormat = original.Wsdl
	WsdlLink        ContentFormat = original.WsdlLink
)

type GrantType = original.GrantType

const (
	AuthorizationCode     GrantType = original.AuthorizationCode
	ClientCredentials     GrantType = original.ClientCredentials
	Implicit              GrantType = original.Implicit
	ResourceOwnerPassword GrantType = original.ResourceOwnerPassword
)

type GroupType = original.GroupType

const (
	Custom   GroupType = original.Custom
	External GroupType = original.External
	System   GroupType = original.System
)

type IdentityProviderType = original.IdentityProviderType

const (
	Aad       IdentityProviderType = original.Aad
	AadB2C    IdentityProviderType = original.AadB2C
	Facebook  IdentityProviderType = original.Facebook
	Google    IdentityProviderType = original.Google
	Microsoft IdentityProviderType = original.Microsoft
	Twitter   IdentityProviderType = original.Twitter
)

type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type LoggerType = original.LoggerType

const (
	AzureEventHub LoggerType = original.AzureEventHub
)

type PolicyScopeContract = original.PolicyScopeContract

const (
	All       PolicyScopeContract = original.All
	API       PolicyScopeContract = original.API
	Operation PolicyScopeContract = original.Operation
	Product   PolicyScopeContract = original.Product
	Tenant    PolicyScopeContract = original.Tenant
)

type ProductState = original.ProductState

const (
	NotPublished ProductState = original.NotPublished
	Published    ProductState = original.Published
)

type Protocol = original.Protocol

const (
	ProtocolHTTP  Protocol = original.ProtocolHTTP
	ProtocolHTTPS Protocol = original.ProtocolHTTPS
)

type SubscriptionState = original.SubscriptionState

const (
	Active    SubscriptionState = original.Active
	Cancelled SubscriptionState = original.Cancelled
	Expired   SubscriptionState = original.Expired
	Rejected  SubscriptionState = original.Rejected
	Submitted SubscriptionState = original.Submitted
	Suspended SubscriptionState = original.Suspended
)

type TemplateName = original.TemplateName

const (
	AccountClosedDeveloper                            TemplateName = original.AccountClosedDeveloper
	ApplicationApprovedNotificationMessage            TemplateName = original.ApplicationApprovedNotificationMessage
	ConfirmSignUpIdentityDefault                      TemplateName = original.ConfirmSignUpIdentityDefault
	EmailChangeIdentityDefault                        TemplateName = original.EmailChangeIdentityDefault
	InviteUserNotificationMessage                     TemplateName = original.InviteUserNotificationMessage
	NewCommentNotificationMessage                     TemplateName = original.NewCommentNotificationMessage
	NewDeveloperNotificationMessage                   TemplateName = original.NewDeveloperNotificationMessage
	NewIssueNotificationMessage                       TemplateName = original.NewIssueNotificationMessage
	PasswordResetByAdminNotificationMessage           TemplateName = original.PasswordResetByAdminNotificationMessage
	PasswordResetIdentityDefault                      TemplateName = original.PasswordResetIdentityDefault
	PurchaseDeveloperNotificationMessage              TemplateName = original.PurchaseDeveloperNotificationMessage
	QuotaLimitApproachingDeveloperNotificationMessage TemplateName = original.QuotaLimitApproachingDeveloperNotificationMessage
	RejectDeveloperNotificationMessage                TemplateName = original.RejectDeveloperNotificationMessage
	RequestDeveloperNotificationMessage               TemplateName = original.RequestDeveloperNotificationMessage
)

type UserState = original.UserState

const (
	UserStateActive  UserState = original.UserStateActive
	UserStateBlocked UserState = original.UserStateBlocked
)

type AccessInformationContract = original.AccessInformationContract
type AccessInformationUpdateParameters = original.AccessInformationUpdateParameters
type APICollection = original.APICollection
type APICollectionIterator = original.APICollectionIterator
type APICollectionPage = original.APICollectionPage
type APIContract = original.APIContract
type APIContractProperties = original.APIContractProperties
type APICreateOrUpdateParameter = original.APICreateOrUpdateParameter
type APICreateOrUpdateParameterWsdlSelector = original.APICreateOrUpdateParameterWsdlSelector
type APIEntityBaseContract = original.APIEntityBaseContract
type APIExportResult = original.APIExportResult
type APIUpdateContract = original.APIUpdateContract
type AuthenticationSettingsContract = original.AuthenticationSettingsContract
type AuthorizationServerCollection = original.AuthorizationServerCollection
type AuthorizationServerCollectionIterator = original.AuthorizationServerCollectionIterator
type AuthorizationServerCollectionPage = original.AuthorizationServerCollectionPage
type AuthorizationServerContract = original.AuthorizationServerContract
type AuthorizationServerContractBaseProperties = original.AuthorizationServerContractBaseProperties
type AuthorizationServerContractProperties = original.AuthorizationServerContractProperties
type AuthorizationServerUpdateContract = original.AuthorizationServerUpdateContract
type BackendAuthorizationHeaderCredentials = original.BackendAuthorizationHeaderCredentials
type BackendBaseParameters = original.BackendBaseParameters
type BackendCollection = original.BackendCollection
type BackendCollectionIterator = original.BackendCollectionIterator
type BackendCollectionPage = original.BackendCollectionPage
type BackendContract = original.BackendContract
type BackendContractProperties = original.BackendContractProperties
type BackendCredentialsContract = original.BackendCredentialsContract
type BackendProperties = original.BackendProperties
type BackendProxyContract = original.BackendProxyContract
type BackendServiceFabricClusterProperties = original.BackendServiceFabricClusterProperties
type BackendTLSProperties = original.BackendTLSProperties
type BackendUpdateParameters = original.BackendUpdateParameters
type CertificateCollection = original.CertificateCollection
type CertificateCollectionIterator = original.CertificateCollectionIterator
type CertificateCollectionPage = original.CertificateCollectionPage
type CertificateContract = original.CertificateContract
type CertificateContractProperties = original.CertificateContractProperties
type CertificateCreateOrUpdateParameters = original.CertificateCreateOrUpdateParameters
type DeployConfigurationParameters = original.DeployConfigurationParameters
type EmailTemplateCollection = original.EmailTemplateCollection
type EmailTemplateCollectionIterator = original.EmailTemplateCollectionIterator
type EmailTemplateCollectionPage = original.EmailTemplateCollectionPage
type EmailTemplateContract = original.EmailTemplateContract
type EmailTemplateContractProperties = original.EmailTemplateContractProperties
type EmailTemplateParametersContractProperties = original.EmailTemplateParametersContractProperties
type EmailTemplateUpdateParameters = original.EmailTemplateUpdateParameters
type ErrorFieldContract = original.ErrorFieldContract
type ErrorResponse = original.ErrorResponse
type GenerateSsoURLResult = original.GenerateSsoURLResult
type GroupCollection = original.GroupCollection
type GroupCollectionIterator = original.GroupCollectionIterator
type GroupCollectionPage = original.GroupCollectionPage
type GroupContract = original.GroupContract
type GroupContractProperties = original.GroupContractProperties
type GroupCreateParameters = original.GroupCreateParameters
type GroupUpdateParameters = original.GroupUpdateParameters
type IdentityProviderBaseParameters = original.IdentityProviderBaseParameters
type IdentityProviderContract = original.IdentityProviderContract
type IdentityProviderContractProperties = original.IdentityProviderContractProperties
type IdentityProviderList = original.IdentityProviderList
type IdentityProviderUpdateParameters = original.IdentityProviderUpdateParameters
type LoggerCollection = original.LoggerCollection
type LoggerCollectionIterator = original.LoggerCollectionIterator
type LoggerCollectionPage = original.LoggerCollectionPage
type LoggerContract = original.LoggerContract
type LoggerContractProperties = original.LoggerContractProperties
type LoggerUpdateContract = original.LoggerUpdateContract
type OAuth2AuthenticationSettingsContract = original.OAuth2AuthenticationSettingsContract
type OpenIDConnectProviderCollection = original.OpenIDConnectProviderCollection
type OpenIDConnectProviderCollectionIterator = original.OpenIDConnectProviderCollectionIterator
type OpenIDConnectProviderCollectionPage = original.OpenIDConnectProviderCollectionPage
type OpenidConnectProviderContract = original.OpenidConnectProviderContract
type OpenidConnectProviderContractProperties = original.OpenidConnectProviderContractProperties
type OpenidConnectProviderUpdateContract = original.OpenidConnectProviderUpdateContract
type OperationCollection = original.OperationCollection
type OperationCollectionIterator = original.OperationCollectionIterator
type OperationCollectionPage = original.OperationCollectionPage
type OperationContract = original.OperationContract
type OperationContractProperties = original.OperationContractProperties
type OperationEntityBaseContract = original.OperationEntityBaseContract
type OperationResultContract = original.OperationResultContract
type OperationResultLogItemContract = original.OperationResultLogItemContract
type OperationUpdateContract = original.OperationUpdateContract
type ParameterContract = original.ParameterContract
type PolicyCollection = original.PolicyCollection
type PolicyContract = original.PolicyContract
type PolicySnippetContract = original.PolicySnippetContract
type PolicySnippetsCollection = original.PolicySnippetsCollection
type PortalDelegationSettings = original.PortalDelegationSettings
type PortalSigninSettings = original.PortalSigninSettings
type PortalSignupSettings = original.PortalSignupSettings
type ProductCollection = original.ProductCollection
type ProductCollectionIterator = original.ProductCollectionIterator
type ProductCollectionPage = original.ProductCollectionPage
type ProductContract = original.ProductContract
type ProductContractProperties = original.ProductContractProperties
type ProductEntityBaseParameters = original.ProductEntityBaseParameters
type ProductUpdateParameters = original.ProductUpdateParameters
type PropertyCollection = original.PropertyCollection
type PropertyCollectionIterator = original.PropertyCollectionIterator
type PropertyCollectionPage = original.PropertyCollectionPage
type PropertyContract = original.PropertyContract
type PropertyContractProperties = original.PropertyContractProperties
type PropertyEntityBaseParameters = original.PropertyEntityBaseParameters
type PropertyUpdateParameters = original.PropertyUpdateParameters
type QuotaCounterCollection = original.QuotaCounterCollection
type QuotaCounterContract = original.QuotaCounterContract
type QuotaCounterValueContract = original.QuotaCounterValueContract
type QuotaCounterValueContractProperties = original.QuotaCounterValueContractProperties
type RegionContract = original.RegionContract
type RegionListResult = original.RegionListResult
type RegistrationDelegationSettingsProperties = original.RegistrationDelegationSettingsProperties
type ReportCollection = original.ReportCollection
type ReportCollectionIterator = original.ReportCollectionIterator
type ReportCollectionPage = original.ReportCollectionPage
type ReportRecordContract = original.ReportRecordContract
type RepresentationContract = original.RepresentationContract
type RequestContract = original.RequestContract
type RequestReportCollection = original.RequestReportCollection
type RequestReportRecordContract = original.RequestReportRecordContract
type ResponseContract = original.ResponseContract
type SaveConfigurationParameter = original.SaveConfigurationParameter
type SchemaCollection = original.SchemaCollection
type SchemaCollectionIterator = original.SchemaCollectionIterator
type SchemaCollectionPage = original.SchemaCollectionPage
type SchemaContract = original.SchemaContract
type SchemaContractProperties = original.SchemaContractProperties
type SchemaDocumentProperties = original.SchemaDocumentProperties
type SubscriptionCollection = original.SubscriptionCollection
type SubscriptionCollectionIterator = original.SubscriptionCollectionIterator
type SubscriptionCollectionPage = original.SubscriptionCollectionPage
type SubscriptionContract = original.SubscriptionContract
type SubscriptionContractProperties = original.SubscriptionContractProperties
type SubscriptionCreateParameters = original.SubscriptionCreateParameters
type SubscriptionKeyParameterNamesContract = original.SubscriptionKeyParameterNamesContract
type SubscriptionsDelegationSettingsProperties = original.SubscriptionsDelegationSettingsProperties
type SubscriptionUpdateParameters = original.SubscriptionUpdateParameters
type TenantConfigurationDeployFuture = original.TenantConfigurationDeployFuture
type TenantConfigurationSaveFuture = original.TenantConfigurationSaveFuture
type TenantConfigurationSyncStateContract = original.TenantConfigurationSyncStateContract
type TenantConfigurationValidateFuture = original.TenantConfigurationValidateFuture
type TermsOfServiceProperties = original.TermsOfServiceProperties
type TokenBodyParameterContract = original.TokenBodyParameterContract
type UserCollection = original.UserCollection
type UserCollectionIterator = original.UserCollectionIterator
type UserCollectionPage = original.UserCollectionPage
type UserContract = original.UserContract
type UserContractProperties = original.UserContractProperties
type UserCreateParameters = original.UserCreateParameters
type UserEntityBaseParameters = original.UserEntityBaseParameters
type UserIdentityCollection = original.UserIdentityCollection
type UserIdentityContract = original.UserIdentityContract
type UserTokenParameters = original.UserTokenParameters
type UserTokenResult = original.UserTokenResult
type UserUpdateParameters = original.UserUpdateParameters
type X509CertificateName = original.X509CertificateName
type APIPolicyClient = original.APIPolicyClient
type ReportsClient = original.ReportsClient
type BaseClient = original.BaseClient
type IdentityProviderClient = original.IdentityProviderClient
type TenantAccessClient = original.TenantAccessClient
type SubscriptionClient = original.SubscriptionClient
type QuotaByCounterKeysClient = original.QuotaByCounterKeysClient
type TenantAccessGitClient = original.TenantAccessGitClient
type PolicyClient = original.PolicyClient
type UserClient = original.UserClient
type TenantConfigurationClient = original.TenantConfigurationClient
type APIOperationPolicyClient = original.APIOperationPolicyClient
type GroupClient = original.GroupClient
type CertificateClient = original.CertificateClient
type APIClient = original.APIClient
type ProductPolicyClient = original.ProductPolicyClient
type QuotaByPeriodKeysClient = original.QuotaByPeriodKeysClient
type RegionsClient = original.RegionsClient
type SignUpSettingsClient = original.SignUpSettingsClient
type UserSubscriptionClient = original.UserSubscriptionClient

func NewAPIOperationPolicyClient() APIOperationPolicyClient {
	return original.NewAPIOperationPolicyClient()
}
func NewGroupClient() GroupClient {
	return original.NewGroupClient()
}
func NewUserClient() UserClient {
	return original.NewUserClient()
}
func NewTenantConfigurationClient() TenantConfigurationClient {
	return original.NewTenantConfigurationClient()
}
func NewSignUpSettingsClient() SignUpSettingsClient {
	return original.NewSignUpSettingsClient()
}
func NewUserSubscriptionClient() UserSubscriptionClient {
	return original.NewUserSubscriptionClient()
}
func NewCertificateClient() CertificateClient {
	return original.NewCertificateClient()
}
func NewAPIClient() APIClient {
	return original.NewAPIClient()
}
func NewProductPolicyClient() ProductPolicyClient {
	return original.NewProductPolicyClient()
}
func NewQuotaByPeriodKeysClient() QuotaByPeriodKeysClient {
	return original.NewQuotaByPeriodKeysClient()
}
func NewRegionsClient() RegionsClient {
	return original.NewRegionsClient()
}
func NewProductClient() ProductClient {
	return original.NewProductClient()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewDelegationSettingsClient() DelegationSettingsClient {
	return original.NewDelegationSettingsClient()
}
func NewProductSubscriptionsClient() ProductSubscriptionsClient {
	return original.NewProductSubscriptionsClient()
}
func NewGroupUserClient() GroupUserClient {
	return original.NewGroupUserClient()
}
func NewProductAPIClient() ProductAPIClient {
	return original.NewProductAPIClient()
}
func NewPolicySnippetsClient() PolicySnippetsClient {
	return original.NewPolicySnippetsClient()
}
func NewAuthorizationServerClient() AuthorizationServerClient {
	return original.NewAuthorizationServerClient()
}
func NewProductGroupClient() ProductGroupClient {
	return original.NewProductGroupClient()
}
func NewOpenIDConnectProviderClient() OpenIDConnectProviderClient {
	return original.NewOpenIDConnectProviderClient()
}
func NewSignInSettingsClient() SignInSettingsClient {
	return original.NewSignInSettingsClient()
}
func NewUserGroupClient() UserGroupClient {
	return original.NewUserGroupClient()
}
func NewEmailTemplateClient() EmailTemplateClient {
	return original.NewEmailTemplateClient()
}
func NewLoggerClient() LoggerClient {
	return original.NewLoggerClient()
}
func NewAPISchemaClient() APISchemaClient {
	return original.NewAPISchemaClient()
}
func NewAPIExportClient() APIExportClient {
	return original.NewAPIExportClient()
}
func NewBackendClient() BackendClient {
	return original.NewBackendClient()
}
func NewUserIdentitiesClient() UserIdentitiesClient {
	return original.NewUserIdentitiesClient()
}
func PossibleAPITypeValues() []APIType {
	return original.PossibleAPITypeValues()
}
func PossibleAsyncOperationStatusValues() []AsyncOperationStatus {
	return original.PossibleAsyncOperationStatusValues()
}
func PossibleAuthorizationMethodValues() []AuthorizationMethod {
	return original.PossibleAuthorizationMethodValues()
}
func PossibleBackendProtocolValues() []BackendProtocol {
	return original.PossibleBackendProtocolValues()
}
func PossibleBearerTokenSendingMethodValues() []BearerTokenSendingMethod {
	return original.PossibleBearerTokenSendingMethodValues()
}
func PossibleClientAuthenticationMethodValues() []ClientAuthenticationMethod {
	return original.PossibleClientAuthenticationMethodValues()
}
func PossibleContentFormatValues() []ContentFormat {
	return original.PossibleContentFormatValues()
}
func PossibleGrantTypeValues() []GrantType {
	return original.PossibleGrantTypeValues()
}
func PossibleGroupTypeValues() []GroupType {
	return original.PossibleGroupTypeValues()
}
func PossibleIdentityProviderTypeValues() []IdentityProviderType {
	return original.PossibleIdentityProviderTypeValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleLoggerTypeValues() []LoggerType {
	return original.PossibleLoggerTypeValues()
}
func PossiblePolicyScopeContractValues() []PolicyScopeContract {
	return original.PossiblePolicyScopeContractValues()
}
func PossibleProductStateValues() []ProductState {
	return original.PossibleProductStateValues()
}
func PossibleProtocolValues() []Protocol {
	return original.PossibleProtocolValues()
}
func PossibleSubscriptionStateValues() []SubscriptionState {
	return original.PossibleSubscriptionStateValues()
}
func PossibleTemplateNameValues() []TemplateName {
	return original.PossibleTemplateNameValues()
}
func PossibleUserStateValues() []UserState {
	return original.PossibleUserStateValues()
}
func NewAPIProductClient() APIProductClient {
	return original.NewAPIProductClient()
}
func NewAPIOperationClient() APIOperationClient {
	return original.NewAPIOperationClient()
}
func NewPropertyClient() PropertyClient {
	return original.NewPropertyClient()
}
func NewTenantAccessClient() TenantAccessClient {
	return original.NewTenantAccessClient()
}
func NewSubscriptionClient() SubscriptionClient {
	return original.NewSubscriptionClient()
}
func NewAPIPolicyClient() APIPolicyClient {
	return original.NewAPIPolicyClient()
}
func NewReportsClient() ReportsClient {
	return original.NewReportsClient()
}
func New() BaseClient {
	return original.New()
}
func NewWithoutDefaults() BaseClient {
	return original.NewWithoutDefaults()
}
func NewIdentityProviderClient() IdentityProviderClient {
	return original.NewIdentityProviderClient()
}
func NewTenantAccessGitClient() TenantAccessGitClient {
	return original.NewTenantAccessGitClient()
}
func NewPolicyClient() PolicyClient {
	return original.NewPolicyClient()
}
func NewQuotaByCounterKeysClient() QuotaByCounterKeysClient {
	return original.NewQuotaByCounterKeysClient()
}
