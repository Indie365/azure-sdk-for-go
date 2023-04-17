//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatalakestore

import "time"

// Account - Data Lake Store account information.
type Account struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The Key Vault encryption identity, if any.
	Identity *EncryptionIdentity `json:"identity,omitempty" azure:"ro"`

	// READ-ONLY; The resource location.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The Data Lake Store account properties.
	Properties *AccountProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The resource tags.
	Tags map[string]*string `json:"tags,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AccountBasic - Basic Data Lake Store account information, returned on list calls.
type AccountBasic struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource location.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The basic Data Lake Store account properties.
	Properties *AccountPropertiesBasic `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The resource tags.
	Tags map[string]*string `json:"tags,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AccountListResult - Data Lake Store account list information response.
type AccountListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; The results of the list operation.
	Value []*AccountBasic `json:"value,omitempty" azure:"ro"`
}

// AccountProperties - Data Lake Store account properties information.
type AccountProperties struct {
	// READ-ONLY; The unique identifier associated with this Data Lake Store account.
	AccountID *string `json:"accountId,omitempty" azure:"ro"`

	// READ-ONLY; The account creation time.
	CreationTime *time.Time `json:"creationTime,omitempty" azure:"ro"`

	// READ-ONLY; The commitment tier in use for the current month.
	CurrentTier *TierType `json:"currentTier,omitempty" azure:"ro"`

	// READ-ONLY; The default owner group for all new folders and files created in the Data Lake Store account.
	DefaultGroup *string `json:"defaultGroup,omitempty" azure:"ro"`

	// READ-ONLY; The Key Vault encryption configuration.
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty" azure:"ro"`

	// READ-ONLY; The current state of encryption provisioning for this Data Lake Store account.
	EncryptionProvisioningState *EncryptionProvisioningState `json:"encryptionProvisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The current state of encryption for this Data Lake Store account.
	EncryptionState *EncryptionState `json:"encryptionState,omitempty" azure:"ro"`

	// READ-ONLY; The full CName endpoint for this account.
	Endpoint *string `json:"endpoint,omitempty" azure:"ro"`

	// READ-ONLY; The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall
	// is disabled, this is not enforced.
	FirewallAllowAzureIPs *FirewallAllowAzureIPsState `json:"firewallAllowAzureIps,omitempty" azure:"ro"`

	// READ-ONLY; The list of firewall rules associated with this Data Lake Store account.
	FirewallRules []*FirewallRule `json:"firewallRules,omitempty" azure:"ro"`

	// READ-ONLY; The current state of the IP address firewall for this Data Lake Store account.
	FirewallState *FirewallState `json:"firewallState,omitempty" azure:"ro"`

	// READ-ONLY; The account last modified time.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty" azure:"ro"`

	// READ-ONLY; The commitment tier to use for next month.
	NewTier *TierType `json:"newTier,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning status of the Data Lake Store account.
	ProvisioningState *DataLakeStoreAccountStatus `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The state of the Data Lake Store account.
	State *DataLakeStoreAccountState `json:"state,omitempty" azure:"ro"`

	// READ-ONLY; The current state of the trusted identity provider feature for this Data Lake Store account.
	TrustedIDProviderState *TrustedIDProviderState `json:"trustedIdProviderState,omitempty" azure:"ro"`

	// READ-ONLY; The list of trusted identity providers associated with this Data Lake Store account.
	TrustedIDProviders []*TrustedIDProvider `json:"trustedIdProviders,omitempty" azure:"ro"`

	// READ-ONLY; The list of virtual network rules associated with this Data Lake Store account.
	VirtualNetworkRules []*VirtualNetworkRule `json:"virtualNetworkRules,omitempty" azure:"ro"`
}

// AccountPropertiesBasic - The basic account specific properties that are associated with an underlying Data Lake Store account.
type AccountPropertiesBasic struct {
	// READ-ONLY; The unique identifier associated with this Data Lake Store account.
	AccountID *string `json:"accountId,omitempty" azure:"ro"`

	// READ-ONLY; The account creation time.
	CreationTime *time.Time `json:"creationTime,omitempty" azure:"ro"`

	// READ-ONLY; The full CName endpoint for this account.
	Endpoint *string `json:"endpoint,omitempty" azure:"ro"`

	// READ-ONLY; The account last modified time.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty" azure:"ro"`

	// READ-ONLY; The provisioning status of the Data Lake Store account.
	ProvisioningState *DataLakeStoreAccountStatus `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The state of the Data Lake Store account.
	State *DataLakeStoreAccountState `json:"state,omitempty" azure:"ro"`
}

// AccountsClientBeginCreateOptions contains the optional parameters for the AccountsClient.BeginCreate method.
type AccountsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginDeleteOptions contains the optional parameters for the AccountsClient.BeginDelete method.
type AccountsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientBeginUpdateOptions contains the optional parameters for the AccountsClient.BeginUpdate method.
type AccountsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AccountsClientCheckNameAvailabilityOptions contains the optional parameters for the AccountsClient.CheckNameAvailability
// method.
type AccountsClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientEnableKeyVaultOptions contains the optional parameters for the AccountsClient.EnableKeyVault method.
type AccountsClientEnableKeyVaultOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.NewListByResourceGroupPager
// method.
type AccountsClientListByResourceGroupOptions struct {
	// A Boolean value of true or false to request a count of the matching resources included with the resources in the response,
	// e.g. Categories?$count=true. Optional.
	Count *bool
	// OData filter. Optional.
	Filter *string
	// OrderBy clause. One or more comma-separated expressions with an optional "asc" (the default) or "desc" depending on the
	// order you'd like the values sorted, e.g. Categories?$orderby=CategoryName desc.
	// Optional.
	Orderby *string
	// OData Select statement. Limits the properties on each entry to just those requested, e.g. Categories?$select=CategoryName,Description.
	// Optional.
	Select *string
	// The number of items to skip over before returning elements. Optional.
	Skip *int32
	// The number of items to return. Optional.
	Top *int32
}

// AccountsClientListOptions contains the optional parameters for the AccountsClient.NewListPager method.
type AccountsClientListOptions struct {
	// The Boolean value of true or false to request a count of the matching resources included with the resources in the response,
	// e.g. Categories?$count=true. Optional.
	Count *bool
	// OData filter. Optional.
	Filter *string
	// OrderBy clause. One or more comma-separated expressions with an optional "asc" (the default) or "desc" depending on the
	// order you'd like the values sorted, e.g. Categories?$orderby=CategoryName desc.
	// Optional.
	Orderby *string
	// OData Select statement. Limits the properties on each entry to just those requested, e.g. Categories?$select=CategoryName,Description.
	// Optional.
	Select *string
	// The number of items to skip over before returning elements. Optional.
	Skip *int32
	// The number of items to return. Optional.
	Top *int32
}

// CapabilityInformation - Subscription-level properties and limits for Data Lake Store.
type CapabilityInformation struct {
	// READ-ONLY; The current number of accounts under this subscription.
	AccountCount *int32 `json:"accountCount,omitempty" azure:"ro"`

	// READ-ONLY; The maximum supported number of accounts under this subscription.
	MaxAccountCount *int32 `json:"maxAccountCount,omitempty" azure:"ro"`

	// READ-ONLY; The Boolean value of true or false to indicate the maintenance state.
	MigrationState *bool `json:"migrationState,omitempty" azure:"ro"`

	// READ-ONLY; The subscription state.
	State *SubscriptionState `json:"state,omitempty" azure:"ro"`

	// READ-ONLY; The subscription credentials that uniquely identifies the subscription.
	SubscriptionID *string `json:"subscriptionId,omitempty" azure:"ro"`
}

// CheckNameAvailabilityParameters - Data Lake Store account name availability check parameters.
type CheckNameAvailabilityParameters struct {
	// REQUIRED; The Data Lake Store name to check availability for.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The resource type. Note: This should not be set by the user, as the constant value is Microsoft.DataLakeStore/accounts
	Type *CheckNameAvailabilityParametersType `json:"type,omitempty"`
}

type CreateDataLakeStoreAccountParameters struct {
	// REQUIRED; The resource location.
	Location *string `json:"location,omitempty"`

	// The Key Vault encryption identity, if any.
	Identity *EncryptionIdentity `json:"identity,omitempty"`

	// The Data Lake Store account properties to use for creating.
	Properties *CreateDataLakeStoreAccountProperties `json:"properties,omitempty"`

	// The resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

type CreateDataLakeStoreAccountProperties struct {
	// The default owner group for all new folders and files created in the Data Lake Store account.
	DefaultGroup *string `json:"defaultGroup,omitempty"`

	// The Key Vault encryption configuration.
	EncryptionConfig *EncryptionConfig `json:"encryptionConfig,omitempty"`

	// The current state of encryption for this Data Lake Store account.
	EncryptionState *EncryptionState `json:"encryptionState,omitempty"`

	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled,
	// this is not enforced.
	FirewallAllowAzureIPs *FirewallAllowAzureIPsState `json:"firewallAllowAzureIps,omitempty"`

	// The list of firewall rules associated with this Data Lake Store account.
	FirewallRules []*CreateFirewallRuleWithAccountParameters `json:"firewallRules,omitempty"`

	// The current state of the IP address firewall for this Data Lake Store account.
	FirewallState *FirewallState `json:"firewallState,omitempty"`

	// The commitment tier to use for next month.
	NewTier *TierType `json:"newTier,omitempty"`

	// The current state of the trusted identity provider feature for this Data Lake Store account.
	TrustedIDProviderState *TrustedIDProviderState `json:"trustedIdProviderState,omitempty"`

	// The list of trusted identity providers associated with this Data Lake Store account.
	TrustedIDProviders []*CreateTrustedIDProviderWithAccountParameters `json:"trustedIdProviders,omitempty"`

	// The list of virtual network rules associated with this Data Lake Store account.
	VirtualNetworkRules []*CreateVirtualNetworkRuleWithAccountParameters `json:"virtualNetworkRules,omitempty"`
}

// CreateFirewallRuleWithAccountParameters - The parameters used to create a new firewall rule while creating a new Data Lake
// Store account.
type CreateFirewallRuleWithAccountParameters struct {
	// REQUIRED; The unique name of the firewall rule to create.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The firewall rule properties to use when creating a new firewall rule.
	Properties *CreateOrUpdateFirewallRuleProperties `json:"properties,omitempty"`
}

// CreateOrUpdateFirewallRuleParameters - The parameters used to create a new firewall rule.
type CreateOrUpdateFirewallRuleParameters struct {
	// REQUIRED; The firewall rule properties to use when creating a new firewall rule.
	Properties *CreateOrUpdateFirewallRuleProperties `json:"properties,omitempty"`
}

// CreateOrUpdateFirewallRuleProperties - The firewall rule properties to use when creating a new firewall rule.
type CreateOrUpdateFirewallRuleProperties struct {
	// REQUIRED; The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same
	// protocol.
	EndIPAddress *string `json:"endIpAddress,omitempty"`

	// REQUIRED; The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same
	// protocol.
	StartIPAddress *string `json:"startIpAddress,omitempty"`
}

// CreateOrUpdateTrustedIDProviderParameters - The parameters used to create a new trusted identity provider.
type CreateOrUpdateTrustedIDProviderParameters struct {
	// REQUIRED; The trusted identity provider properties to use when creating a new trusted identity provider.
	Properties *CreateOrUpdateTrustedIDProviderProperties `json:"properties,omitempty"`
}

// CreateOrUpdateTrustedIDProviderProperties - The trusted identity provider properties to use when creating a new trusted
// identity provider.
type CreateOrUpdateTrustedIDProviderProperties struct {
	// REQUIRED; The URL of this trusted identity provider.
	IDProvider *string `json:"idProvider,omitempty"`
}

// CreateOrUpdateVirtualNetworkRuleParameters - The parameters used to create a new virtual network rule.
type CreateOrUpdateVirtualNetworkRuleParameters struct {
	// REQUIRED; The virtual network rule properties to use when creating a new virtual network rule.
	Properties *CreateOrUpdateVirtualNetworkRuleProperties `json:"properties,omitempty"`
}

// CreateOrUpdateVirtualNetworkRuleProperties - The virtual network rule properties to use when creating a new virtual network
// rule.
type CreateOrUpdateVirtualNetworkRuleProperties struct {
	// REQUIRED; The resource identifier for the subnet.
	SubnetID *string `json:"subnetId,omitempty"`
}

// CreateTrustedIDProviderWithAccountParameters - The parameters used to create a new trusted identity provider while creating
// a new Data Lake Store account.
type CreateTrustedIDProviderWithAccountParameters struct {
	// REQUIRED; The unique name of the trusted identity provider to create.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The trusted identity provider properties to use when creating a new trusted identity provider.
	Properties *CreateOrUpdateTrustedIDProviderProperties `json:"properties,omitempty"`
}

// CreateVirtualNetworkRuleWithAccountParameters - The parameters used to create a new virtual network rule while creating
// a new Data Lake Store account.
type CreateVirtualNetworkRuleWithAccountParameters struct {
	// REQUIRED; The unique name of the virtual network rule to create.
	Name *string `json:"name,omitempty"`

	// REQUIRED; The virtual network rule properties to use when creating a new virtual network rule.
	Properties *CreateOrUpdateVirtualNetworkRuleProperties `json:"properties,omitempty"`
}

// EncryptionConfig - The encryption configuration for the account.
type EncryptionConfig struct {
	// REQUIRED; The type of encryption configuration being used. Currently the only supported types are 'UserManaged' and 'ServiceManaged'.
	Type *EncryptionConfigType `json:"type,omitempty"`

	// The Key Vault information for connecting to user managed encryption keys.
	KeyVaultMetaInfo *KeyVaultMetaInfo `json:"keyVaultMetaInfo,omitempty"`
}

// EncryptionIdentity - The encryption identity properties.
type EncryptionIdentity struct {
	// CONSTANT; The type of encryption being used. Currently the only supported type is 'SystemAssigned'.
	// Field has constant value "SystemAssigned", any specified value is ignored.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; The principal identifier associated with the encryption.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant identifier associated with the encryption.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// FirewallRule - Data Lake Store firewall rule information.
type FirewallRule struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The firewall rule properties.
	Properties *FirewallRuleProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// FirewallRuleListResult - Data Lake Store firewall rule list information.
type FirewallRuleListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; The results of the list operation.
	Value []*FirewallRule `json:"value,omitempty" azure:"ro"`
}

// FirewallRuleProperties - The firewall rule properties.
type FirewallRuleProperties struct {
	// READ-ONLY; The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same
	// protocol.
	EndIPAddress *string `json:"endIpAddress,omitempty" azure:"ro"`

	// READ-ONLY; The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the
	// same protocol.
	StartIPAddress *string `json:"startIpAddress,omitempty" azure:"ro"`
}

// FirewallRulesClientCreateOrUpdateOptions contains the optional parameters for the FirewallRulesClient.CreateOrUpdate method.
type FirewallRulesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientDeleteOptions contains the optional parameters for the FirewallRulesClient.Delete method.
type FirewallRulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientGetOptions contains the optional parameters for the FirewallRulesClient.Get method.
type FirewallRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientListByAccountOptions contains the optional parameters for the FirewallRulesClient.NewListByAccountPager
// method.
type FirewallRulesClientListByAccountOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientUpdateOptions contains the optional parameters for the FirewallRulesClient.Update method.
type FirewallRulesClientUpdateOptions struct {
	// Parameters supplied to update the firewall rule.
	Parameters *UpdateFirewallRuleParameters
}

// KeyVaultMetaInfo - Metadata information used by account encryption.
type KeyVaultMetaInfo struct {
	// REQUIRED; The name of the user managed encryption key.
	EncryptionKeyName *string `json:"encryptionKeyName,omitempty"`

	// REQUIRED; The version of the user managed encryption key.
	EncryptionKeyVersion *string `json:"encryptionKeyVersion,omitempty"`

	// REQUIRED; The resource identifier for the user managed Key Vault being used to encrypt.
	KeyVaultResourceID *string `json:"keyVaultResourceId,omitempty"`
}

// LocationsClientGetCapabilityOptions contains the optional parameters for the LocationsClient.GetCapability method.
type LocationsClientGetCapabilityOptions struct {
	// placeholder for future optional parameters
}

// LocationsClientGetUsageOptions contains the optional parameters for the LocationsClient.NewGetUsagePager method.
type LocationsClientGetUsageOptions struct {
	// placeholder for future optional parameters
}

// NameAvailabilityInformation - Data Lake Store account name availability result information.
type NameAvailabilityInformation struct {
	// READ-ONLY; The message describing why the Data Lake Store account name is not available, if nameAvailable is false.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The Boolean value of true or false to indicate whether the Data Lake Store account name is available or not.
	NameAvailable *bool `json:"nameAvailable,omitempty" azure:"ro"`

	// READ-ONLY; The reason why the Data Lake Store account name is not available, if nameAvailable is false.
	Reason *string `json:"reason,omitempty" azure:"ro"`
}

// Operation - An available operation for Data Lake Store.
type Operation struct {
	// The display information for the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; The name of the operation.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation.
	Origin *OperationOrigin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - The display information for a particular operation.
type OperationDisplay struct {
	// READ-ONLY; A friendly description of the operation.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; A friendly name of the operation.
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The resource provider of the operation.
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The resource type of the operation.
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - The list of available operations for Data Lake Store.
type OperationListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; The results of the list operation.
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - The resource model definition.
type Resource struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource location.
	Location *string `json:"location,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource tags.
	Tags map[string]*string `json:"tags,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SubResource - The resource model definition for a nested resource.
type SubResource struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TrustedIDProvider - Data Lake Store trusted identity provider information.
type TrustedIDProvider struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The trusted identity provider properties.
	Properties *TrustedIDProviderProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TrustedIDProviderListResult - Data Lake Store trusted identity provider list information.
type TrustedIDProviderListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; The results of the list operation.
	Value []*TrustedIDProvider `json:"value,omitempty" azure:"ro"`
}

// TrustedIDProviderProperties - The trusted identity provider properties.
type TrustedIDProviderProperties struct {
	// READ-ONLY; The URL of this trusted identity provider.
	IDProvider *string `json:"idProvider,omitempty" azure:"ro"`
}

// TrustedIDProvidersClientCreateOrUpdateOptions contains the optional parameters for the TrustedIDProvidersClient.CreateOrUpdate
// method.
type TrustedIDProvidersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TrustedIDProvidersClientDeleteOptions contains the optional parameters for the TrustedIDProvidersClient.Delete method.
type TrustedIDProvidersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// TrustedIDProvidersClientGetOptions contains the optional parameters for the TrustedIDProvidersClient.Get method.
type TrustedIDProvidersClientGetOptions struct {
	// placeholder for future optional parameters
}

// TrustedIDProvidersClientListByAccountOptions contains the optional parameters for the TrustedIDProvidersClient.NewListByAccountPager
// method.
type TrustedIDProvidersClientListByAccountOptions struct {
	// placeholder for future optional parameters
}

// TrustedIDProvidersClientUpdateOptions contains the optional parameters for the TrustedIDProvidersClient.Update method.
type TrustedIDProvidersClientUpdateOptions struct {
	// Parameters supplied to update the trusted identity provider.
	Parameters *UpdateTrustedIDProviderParameters
}

// UpdateDataLakeStoreAccountParameters - Data Lake Store account information to update.
type UpdateDataLakeStoreAccountParameters struct {
	// The Data Lake Store account properties to update.
	Properties *UpdateDataLakeStoreAccountProperties `json:"properties,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}

// UpdateDataLakeStoreAccountProperties - Data Lake Store account properties information to be updated.
type UpdateDataLakeStoreAccountProperties struct {
	// The default owner group for all new folders and files created in the Data Lake Store account.
	DefaultGroup *string `json:"defaultGroup,omitempty"`

	// Used for rotation of user managed Key Vault keys. Can only be used to rotate a user managed encryption Key Vault key.
	EncryptionConfig *UpdateEncryptionConfig `json:"encryptionConfig,omitempty"`

	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled,
	// this is not enforced.
	FirewallAllowAzureIPs *FirewallAllowAzureIPsState `json:"firewallAllowAzureIps,omitempty"`

	// The list of firewall rules associated with this Data Lake Store account.
	FirewallRules []*UpdateFirewallRuleWithAccountParameters `json:"firewallRules,omitempty"`

	// The current state of the IP address firewall for this Data Lake Store account. Disabling the firewall does not remove existing
	// rules, they will just be ignored until the firewall is re-enabled.
	FirewallState *FirewallState `json:"firewallState,omitempty"`

	// The commitment tier to use for next month.
	NewTier *TierType `json:"newTier,omitempty"`

	// The current state of the trusted identity provider feature for this Data Lake Store account. Disabling trusted identity
	// provider functionality does not remove the providers, they will just be ignored
	// until this feature is re-enabled.
	TrustedIDProviderState *TrustedIDProviderState `json:"trustedIdProviderState,omitempty"`

	// The list of trusted identity providers associated with this Data Lake Store account.
	TrustedIDProviders []*UpdateTrustedIDProviderWithAccountParameters `json:"trustedIdProviders,omitempty"`

	// The list of virtual network rules associated with this Data Lake Store account.
	VirtualNetworkRules []*UpdateVirtualNetworkRuleWithAccountParameters `json:"virtualNetworkRules,omitempty"`
}

// UpdateEncryptionConfig - The encryption configuration used to update a user managed Key Vault key.
type UpdateEncryptionConfig struct {
	// The updated Key Vault key to use in user managed key rotation.
	KeyVaultMetaInfo *UpdateKeyVaultMetaInfo `json:"keyVaultMetaInfo,omitempty"`
}

// UpdateFirewallRuleParameters - The parameters used to update a firewall rule.
type UpdateFirewallRuleParameters struct {
	// The firewall rule properties to use when updating a firewall rule.
	Properties *UpdateFirewallRuleProperties `json:"properties,omitempty"`
}

// UpdateFirewallRuleProperties - The firewall rule properties to use when updating a firewall rule.
type UpdateFirewallRuleProperties struct {
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIPAddress *string `json:"endIpAddress,omitempty"`

	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIPAddress *string `json:"startIpAddress,omitempty"`
}

// UpdateFirewallRuleWithAccountParameters - The parameters used to update a firewall rule while updating a Data Lake Store
// account.
type UpdateFirewallRuleWithAccountParameters struct {
	// REQUIRED; The unique name of the firewall rule to update.
	Name *string `json:"name,omitempty"`

	// The firewall rule properties to use when updating a firewall rule.
	Properties *UpdateFirewallRuleProperties `json:"properties,omitempty"`
}

// UpdateKeyVaultMetaInfo - The Key Vault update information used for user managed key rotation.
type UpdateKeyVaultMetaInfo struct {
	// The version of the user managed encryption key to update through a key rotation.
	EncryptionKeyVersion *string `json:"encryptionKeyVersion,omitempty"`
}

// UpdateTrustedIDProviderParameters - The parameters used to update a trusted identity provider.
type UpdateTrustedIDProviderParameters struct {
	// The trusted identity provider properties to use when updating a trusted identity provider.
	Properties *UpdateTrustedIDProviderProperties `json:"properties,omitempty"`
}

// UpdateTrustedIDProviderProperties - The trusted identity provider properties to use when updating a trusted identity provider.
type UpdateTrustedIDProviderProperties struct {
	// The URL of this trusted identity provider.
	IDProvider *string `json:"idProvider,omitempty"`
}

// UpdateTrustedIDProviderWithAccountParameters - The parameters used to update a trusted identity provider while updating
// a Data Lake Store account.
type UpdateTrustedIDProviderWithAccountParameters struct {
	// REQUIRED; The unique name of the trusted identity provider to update.
	Name *string `json:"name,omitempty"`

	// The trusted identity provider properties to use when updating a trusted identity provider.
	Properties *UpdateTrustedIDProviderProperties `json:"properties,omitempty"`
}

// UpdateVirtualNetworkRuleParameters - The parameters used to update a virtual network rule.
type UpdateVirtualNetworkRuleParameters struct {
	// The virtual network rule properties to use when updating a virtual network rule.
	Properties *UpdateVirtualNetworkRuleProperties `json:"properties,omitempty"`
}

// UpdateVirtualNetworkRuleProperties - The virtual network rule properties to use when updating a virtual network rule.
type UpdateVirtualNetworkRuleProperties struct {
	// The resource identifier for the subnet.
	SubnetID *string `json:"subnetId,omitempty"`
}

// UpdateVirtualNetworkRuleWithAccountParameters - The parameters used to update a virtual network rule while updating a Data
// Lake Store account.
type UpdateVirtualNetworkRuleWithAccountParameters struct {
	// REQUIRED; The unique name of the virtual network rule to update.
	Name *string `json:"name,omitempty"`

	// The virtual network rule properties to use when updating a virtual network rule.
	Properties *UpdateVirtualNetworkRuleProperties `json:"properties,omitempty"`
}

// Usage - Describes the Resource Usage.
type Usage struct {
	// READ-ONLY; Gets the current count of the allocated resources in the subscription.
	CurrentValue *int32 `json:"currentValue,omitempty" azure:"ro"`

	// READ-ONLY; Resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Gets the maximum count of the resources that can be allocated in the subscription.
	Limit *int32 `json:"limit,omitempty" azure:"ro"`

	// READ-ONLY; Gets the name of the type of usage.
	Name *UsageName `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Gets the unit of measurement.
	Unit *UsageUnit `json:"unit,omitempty" azure:"ro"`
}

// UsageListResult - The response from the List Usages operation.
type UsageListResult struct {
	// Gets or sets the list of Storage Resource Usages.
	Value []*Usage `json:"value,omitempty"`
}

// UsageName - The usage names that can be used.
type UsageName struct {
	// READ-ONLY; Gets a localized string describing the resource name.
	LocalizedValue *string `json:"localizedValue,omitempty" azure:"ro"`

	// READ-ONLY; Gets a string describing the resource name.
	Value *string `json:"value,omitempty" azure:"ro"`
}

// VirtualNetworkRule - Data Lake Store virtual network rule information.
type VirtualNetworkRule struct {
	// READ-ONLY; The resource identifier.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The virtual network rule properties.
	Properties *VirtualNetworkRuleProperties `json:"properties,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// VirtualNetworkRuleListResult - Data Lake Store virtual network rule list information.
type VirtualNetworkRuleListResult struct {
	// READ-ONLY; The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; The results of the list operation.
	Value []*VirtualNetworkRule `json:"value,omitempty" azure:"ro"`
}

// VirtualNetworkRuleProperties - The virtual network rule properties.
type VirtualNetworkRuleProperties struct {
	// READ-ONLY; The resource identifier for the subnet.
	SubnetID *string `json:"subnetId,omitempty" azure:"ro"`
}

// VirtualNetworkRulesClientCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkRulesClient.CreateOrUpdate
// method.
type VirtualNetworkRulesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkRulesClientDeleteOptions contains the optional parameters for the VirtualNetworkRulesClient.Delete method.
type VirtualNetworkRulesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkRulesClientGetOptions contains the optional parameters for the VirtualNetworkRulesClient.Get method.
type VirtualNetworkRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkRulesClientListByAccountOptions contains the optional parameters for the VirtualNetworkRulesClient.NewListByAccountPager
// method.
type VirtualNetworkRulesClientListByAccountOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkRulesClientUpdateOptions contains the optional parameters for the VirtualNetworkRulesClient.Update method.
type VirtualNetworkRulesClientUpdateOptions struct {
	// Parameters supplied to update the virtual network rule.
	Parameters *UpdateVirtualNetworkRuleParameters
}
