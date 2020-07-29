package storage

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2016-12-01/storage"

// Account the storage account.
type Account struct {
	autorest.Response `json:"-"`
	// Sku - READ-ONLY; Gets the SKU.
	Sku *Sku `json:"sku,omitempty"`
	// Kind - READ-ONLY; Gets the Kind. Possible values include: 'Storage', 'BlobStorage'
	Kind               Kind `json:"kind,omitempty"`
	*AccountProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Tags assigned to a resource; can be used for viewing and grouping a resource (across resource groups).
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AccountProperties != nil {
		objectMap["properties"] = a.AccountProperties
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				a.Sku = &sku
			}
		case "kind":
			if v != nil {
				var kind Kind
				err = json.Unmarshal(*v, &kind)
				if err != nil {
					return err
				}
				a.Kind = kind
			}
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				a.AccountProperties = &accountProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		}
	}

	return nil
}

// AccountCheckNameAvailabilityParameters the parameters used to check the availability of the storage account
// name.
type AccountCheckNameAvailabilityParameters struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// AccountCreateParameters the parameters used when creating a storage account.
type AccountCreateParameters struct {
	// Sku - Required. Gets or sets the sku name.
	Sku *Sku `json:"sku,omitempty"`
	// Kind - Required. Indicates the type of storage account. Possible values include: 'Storage', 'BlobStorage'
	Kind Kind `json:"kind,omitempty"`
	// Location - Required. Gets or sets the location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US, East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it is created, but if an identical geo region is specified on update, the request will succeed.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used for viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key with a length no greater than 128 characters and a value with a length no greater than 256 characters.
	Tags                               map[string]*string `json:"tags"`
	*AccountPropertiesCreateParameters `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountCreateParameters.
func (acp AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.Sku != nil {
		objectMap["sku"] = acp.Sku
	}
	if acp.Kind != "" {
		objectMap["kind"] = acp.Kind
	}
	if acp.Location != nil {
		objectMap["location"] = acp.Location
	}
	if acp.Tags != nil {
		objectMap["tags"] = acp.Tags
	}
	if acp.AccountPropertiesCreateParameters != nil {
		objectMap["properties"] = acp.AccountPropertiesCreateParameters
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountCreateParameters struct.
func (acp *AccountCreateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				acp.Sku = &sku
			}
		case "kind":
			if v != nil {
				var kind Kind
				err = json.Unmarshal(*v, &kind)
				if err != nil {
					return err
				}
				acp.Kind = kind
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				acp.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				acp.Tags = tags
			}
		case "properties":
			if v != nil {
				var accountPropertiesCreateParameters AccountPropertiesCreateParameters
				err = json.Unmarshal(*v, &accountPropertiesCreateParameters)
				if err != nil {
					return err
				}
				acp.AccountPropertiesCreateParameters = &accountPropertiesCreateParameters
			}
		}
	}

	return nil
}

// AccountKey an access key for the storage account.
type AccountKey struct {
	// KeyName - READ-ONLY; Name of the key.
	KeyName *string `json:"keyName,omitempty"`
	// Value - READ-ONLY; Base 64-encoded value of the key.
	Value *string `json:"value,omitempty"`
	// Permissions - READ-ONLY; Permissions for the key -- read-only or full permissions. Possible values include: 'Read', 'Full'
	Permissions KeyPermission `json:"permissions,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountKey.
func (ak AccountKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountListKeysResult the response from the ListKeys operation.
type AccountListKeysResult struct {
	autorest.Response `json:"-"`
	// Keys - READ-ONLY; Gets the list of storage account keys and their properties for the specified storage account.
	Keys *[]AccountKey `json:"keys,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountListKeysResult.
func (alkr AccountListKeysResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountListResult the response from the List Storage Accounts operation.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Gets the list of storage accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountListResult.
func (alr AccountListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// AccountProperties properties of the storage account.
type AccountProperties struct {
	// ProvisioningState - READ-ONLY; Gets the status of the storage account at the time the operation was called. Possible values include: 'Creating', 'ResolvingDNS', 'Succeeded'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// PrimaryEndpoints - READ-ONLY; Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
	PrimaryEndpoints *Endpoints `json:"primaryEndpoints,omitempty"`
	// PrimaryLocation - READ-ONLY; Gets the location of the primary data center for the storage account.
	PrimaryLocation *string `json:"primaryLocation,omitempty"`
	// StatusOfPrimary - READ-ONLY; Gets the status indicating whether the primary location of the storage account is available or unavailable. Possible values include: 'Available', 'Unavailable'
	StatusOfPrimary AccountStatus `json:"statusOfPrimary,omitempty"`
	// LastGeoFailoverTime - READ-ONLY; Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	LastGeoFailoverTime *date.Time `json:"lastGeoFailoverTime,omitempty"`
	// SecondaryLocation - READ-ONLY; Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
	SecondaryLocation *string `json:"secondaryLocation,omitempty"`
	// StatusOfSecondary - READ-ONLY; Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS. Possible values include: 'Available', 'Unavailable'
	StatusOfSecondary AccountStatus `json:"statusOfSecondary,omitempty"`
	// CreationTime - READ-ONLY; Gets the creation date and time of the storage account in UTC.
	CreationTime *date.Time `json:"creationTime,omitempty"`
	// CustomDomain - READ-ONLY; Gets the custom domain the user assigned to this storage account.
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
	// SecondaryEndpoints - READ-ONLY; Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
	SecondaryEndpoints *Endpoints `json:"secondaryEndpoints,omitempty"`
	// Encryption - READ-ONLY; Gets the encryption settings on the account. If unspecified, the account is unencrypted.
	Encryption *Encryption `json:"encryption,omitempty"`
	// AccessTier - READ-ONLY; Required for storage accounts where kind = BlobStorage. The access tier used for billing. Possible values include: 'Hot', 'Cool'
	AccessTier AccessTier `json:"accessTier,omitempty"`
	// EnableHTTPSTrafficOnly - Allows https traffic only to storage service if sets to true.
	EnableHTTPSTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountProperties.
func (ap AccountProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ap.EnableHTTPSTrafficOnly != nil {
		objectMap["supportsHttpsTrafficOnly"] = ap.EnableHTTPSTrafficOnly
	}
	return json.Marshal(objectMap)
}

// AccountPropertiesCreateParameters the parameters used to create the storage account.
type AccountPropertiesCreateParameters struct {
	// CustomDomain - User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
	// Encryption - Provides the encryption settings on the account. If left unspecified the account encryption settings will remain the same. The default setting is unencrypted.
	Encryption *Encryption `json:"encryption,omitempty"`
	// AccessTier - Required for storage accounts where kind = BlobStorage. The access tier used for billing. Possible values include: 'Hot', 'Cool'
	AccessTier AccessTier `json:"accessTier,omitempty"`
	// EnableHTTPSTrafficOnly - Allows https traffic only to storage service if sets to true.
	EnableHTTPSTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

// AccountPropertiesUpdateParameters the parameters used when updating a storage account.
type AccountPropertiesUpdateParameters struct {
	// CustomDomain - Custom domain assigned to the storage account by the user. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
	CustomDomain *CustomDomain `json:"customDomain,omitempty"`
	// Encryption - Provides the encryption settings on the account. The default setting is unencrypted.
	Encryption *Encryption `json:"encryption,omitempty"`
	// AccessTier - Required for storage accounts where kind = BlobStorage. The access tier used for billing. Possible values include: 'Hot', 'Cool'
	AccessTier AccessTier `json:"accessTier,omitempty"`
	// EnableHTTPSTrafficOnly - Allows https traffic only to storage service if sets to true.
	EnableHTTPSTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

// AccountRegenerateKeyParameters the parameters used to regenerate the storage account key.
type AccountRegenerateKeyParameters struct {
	KeyName *string `json:"keyName,omitempty"`
}

// AccountSasParameters the parameters to list SAS credentials of a storage account.
type AccountSasParameters struct {
	// Services - The signed services accessible with the account SAS. Possible values include: Blob (b), Queue (q), Table (t), File (f). Possible values include: 'B', 'Q', 'T', 'F'
	Services Services `json:"signedServices,omitempty"`
	// ResourceTypes - The signed resource types that are accessible with the account SAS. Service (s): Access to service-level APIs; Container (c): Access to container-level APIs; Object (o): Access to object-level APIs for blobs, queue messages, table entities, and files. Possible values include: 'ResourceTypesS', 'ResourceTypesC', 'ResourceTypesO'
	ResourceTypes ResourceTypes `json:"signedResourceTypes,omitempty"`
	// Permissions - The signed permissions for the account SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p). Possible values include: 'Permissions1R', 'Permissions1D', 'Permissions1W', 'Permissions1L', 'Permissions1A', 'Permissions1C', 'Permissions1U', 'Permissions1P'
	Permissions Permissions1 `json:"signedPermission,omitempty"`
	// IPAddressOrRange - An IP address or a range of IP addresses from which to accept requests.
	IPAddressOrRange *string `json:"signedIp,omitempty"`
	// Protocols - The protocol permitted for a request made with the account SAS. Possible values include: 'Httpshttp', 'HTTPS'
	Protocols HTTPProtocol `json:"signedProtocol,omitempty"`
	// SharedAccessStartTime - The time at which the SAS becomes valid.
	SharedAccessStartTime *date.Time `json:"signedStart,omitempty"`
	// SharedAccessExpiryTime - The time at which the shared access signature becomes invalid.
	SharedAccessExpiryTime *date.Time `json:"signedExpiry,omitempty"`
	// KeyToSign - The key to sign the account SAS token with.
	KeyToSign *string `json:"keyToSign,omitempty"`
}

// AccountsCreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type AccountsCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AccountsCreateFuture) Result(client AccountsClient) (a Account, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storage.AccountsCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("storage.AccountsCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if a.Response.Response, err = future.GetResult(sender); err == nil && a.Response.Response.StatusCode != http.StatusNoContent {
		a, err = client.CreateResponder(a.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "storage.AccountsCreateFuture", "Result", a.Response.Response, "Failure responding to request")
		}
	}
	return
}

// AccountUpdateParameters the parameters that can be provided when updating the storage account properties.
type AccountUpdateParameters struct {
	// Sku - Gets or sets the SKU name. Note that the SKU name cannot be updated to Standard_ZRS or Premium_LRS, nor can accounts of those sku names be updated to any other value.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater in length than 128 characters and a value no greater in length than 256 characters.
	Tags                               map[string]*string `json:"tags"`
	*AccountPropertiesUpdateParameters `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Sku != nil {
		objectMap["sku"] = aup.Sku
	}
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	if aup.AccountPropertiesUpdateParameters != nil {
		objectMap["properties"] = aup.AccountPropertiesUpdateParameters
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountUpdateParameters struct.
func (aup *AccountUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				aup.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				aup.Tags = tags
			}
		case "properties":
			if v != nil {
				var accountPropertiesUpdateParameters AccountPropertiesUpdateParameters
				err = json.Unmarshal(*v, &accountPropertiesUpdateParameters)
				if err != nil {
					return err
				}
				aup.AccountPropertiesUpdateParameters = &accountPropertiesUpdateParameters
			}
		}
	}

	return nil
}

// CheckNameAvailabilityResult the CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; Gets a boolean value that indicates whether the name is available for you to use. If true, the name is available. If false, the name has already been taken or is invalid and cannot be used.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; Gets the reason that a storage account name could not be used. The Reason element is only returned if NameAvailable is false. Possible values include: 'AccountNameInvalid', 'AlreadyExists'
	Reason Reason `json:"reason,omitempty"`
	// Message - READ-ONLY; Gets an error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty"`
}

// MarshalJSON is the custom marshaler for CheckNameAvailabilityResult.
func (cnar CheckNameAvailabilityResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// CustomDomain the custom domain assigned to this storage account. This can be set via Update.
type CustomDomain struct {
	// Name - Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name *string `json:"name,omitempty"`
	// UseSubDomainName - Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
	UseSubDomainName *bool `json:"useSubDomainName,omitempty"`
}

// Encryption the encryption settings on the storage account.
type Encryption struct {
	// Services - List of services which support encryption.
	Services *EncryptionServices `json:"services,omitempty"`
	// KeySource - The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
	KeySource *string `json:"keySource,omitempty"`
}

// EncryptionService a service that allows server-side encryption to be used.
type EncryptionService struct {
	// Enabled - A boolean indicating whether or not the service encrypts the data as it is stored.
	Enabled *bool `json:"enabled,omitempty"`
	// LastEnabledTime - READ-ONLY; Gets a rough estimate of the date/time when the encryption was last enabled by the user. Only returned when encryption is enabled. There might be some unencrypted blobs which were written after this time, as it is just a rough estimate.
	LastEnabledTime *date.Time `json:"lastEnabledTime,omitempty"`
}

// MarshalJSON is the custom marshaler for EncryptionService.
func (es EncryptionService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if es.Enabled != nil {
		objectMap["enabled"] = es.Enabled
	}
	return json.Marshal(objectMap)
}

// EncryptionServices a list of services that support encryption.
type EncryptionServices struct {
	// Blob - The encryption function of the blob storage service.
	Blob *EncryptionService `json:"blob,omitempty"`
	// File - The encryption function of the file storage service.
	File *EncryptionService `json:"file,omitempty"`
	// Table - READ-ONLY; The encryption function of the table storage service.
	Table *EncryptionService `json:"table,omitempty"`
	// Queue - READ-ONLY; The encryption function of the queue storage service.
	Queue *EncryptionService `json:"queue,omitempty"`
}

// MarshalJSON is the custom marshaler for EncryptionServices.
func (es EncryptionServices) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if es.Blob != nil {
		objectMap["blob"] = es.Blob
	}
	if es.File != nil {
		objectMap["file"] = es.File
	}
	return json.Marshal(objectMap)
}

// Endpoints the URIs that are used to perform a retrieval of a public blob, queue, or table object.
type Endpoints struct {
	// Blob - READ-ONLY; Gets the blob endpoint.
	Blob *string `json:"blob,omitempty"`
	// Queue - READ-ONLY; Gets the queue endpoint.
	Queue *string `json:"queue,omitempty"`
	// Table - READ-ONLY; Gets the table endpoint.
	Table *string `json:"table,omitempty"`
	// File - READ-ONLY; Gets the file endpoint.
	File *string `json:"file,omitempty"`
}

// MarshalJSON is the custom marshaler for Endpoints.
func (e Endpoints) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ListAccountSasResponse the List SAS credentials operation response.
type ListAccountSasResponse struct {
	autorest.Response `json:"-"`
	// AccountSasToken - READ-ONLY; List SAS credentials of storage account.
	AccountSasToken *string `json:"accountSasToken,omitempty"`
}

// MarshalJSON is the custom marshaler for ListAccountSasResponse.
func (lasr ListAccountSasResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ListServiceSasResponse the List service SAS credentials operation response.
type ListServiceSasResponse struct {
	autorest.Response `json:"-"`
	// ServiceSasToken - READ-ONLY; List service SAS credentials of specific resource.
	ServiceSasToken *string `json:"serviceSasToken,omitempty"`
}

// MarshalJSON is the custom marshaler for ListServiceSasResponse.
func (lssr ListServiceSasResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Resource describes a storage resource.
type Resource struct {
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Tags assigned to a resource; can be used for viewing and grouping a resource (across resource groups).
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ServiceSasParameters the parameters to list service SAS credentials of a specific resource.
type ServiceSasParameters struct {
	// CanonicalizedResource - The canonical path to the signed resource.
	CanonicalizedResource *string `json:"canonicalizedResource,omitempty"`
	// Resource - The signed services accessible with the service SAS. Possible values include: Blob (b), Container (c), File (f), Share (s). Possible values include: 'SignedResourceB', 'SignedResourceC', 'SignedResourceF', 'SignedResourceS'
	Resource SignedResource `json:"signedResource,omitempty"`
	// Permissions - The signed permissions for the service SAS. Possible values include: Read (r), Write (w), Delete (d), List (l), Add (a), Create (c), Update (u) and Process (p). Possible values include: 'R', 'D', 'W', 'L', 'A', 'C', 'U', 'P'
	Permissions Permissions `json:"signedPermission,omitempty"`
	// IPAddressOrRange - An IP address or a range of IP addresses from which to accept requests.
	IPAddressOrRange *string `json:"signedIp,omitempty"`
	// Protocols - The protocol permitted for a request made with the account SAS. Possible values include: 'Httpshttp', 'HTTPS'
	Protocols HTTPProtocol `json:"signedProtocol,omitempty"`
	// SharedAccessStartTime - The time at which the SAS becomes valid.
	SharedAccessStartTime *date.Time `json:"signedStart,omitempty"`
	// SharedAccessExpiryTime - The time at which the shared access signature becomes invalid.
	SharedAccessExpiryTime *date.Time `json:"signedExpiry,omitempty"`
	// Identifier - A unique value up to 64 characters in length that correlates to an access policy specified for the container, queue, or table.
	Identifier *string `json:"signedIdentifier,omitempty"`
	// PartitionKeyStart - The start of partition key.
	PartitionKeyStart *string `json:"startPk,omitempty"`
	// PartitionKeyEnd - The end of partition key.
	PartitionKeyEnd *string `json:"endPk,omitempty"`
	// RowKeyStart - The start of row key.
	RowKeyStart *string `json:"startRk,omitempty"`
	// RowKeyEnd - The end of row key.
	RowKeyEnd *string `json:"endRk,omitempty"`
	// KeyToSign - The key to sign the account SAS token with.
	KeyToSign *string `json:"keyToSign,omitempty"`
	// CacheControl - The response header override for cache control.
	CacheControl *string `json:"rscc,omitempty"`
	// ContentDisposition - The response header override for content disposition.
	ContentDisposition *string `json:"rscd,omitempty"`
	// ContentEncoding - The response header override for content encoding.
	ContentEncoding *string `json:"rsce,omitempty"`
	// ContentLanguage - The response header override for content language.
	ContentLanguage *string `json:"rscl,omitempty"`
	// ContentType - The response header override for content type.
	ContentType *string `json:"rsct,omitempty"`
}

// Sku the SKU of the storage account.
type Sku struct {
	// Name - Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType. Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS'
	Name SkuName `json:"name,omitempty"`
	// Tier - READ-ONLY; Gets the sku tier. This is based on the SKU name. Possible values include: 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
}

// MarshalJSON is the custom marshaler for Sku.
func (s Sku) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.Name != "" {
		objectMap["name"] = s.Name
	}
	return json.Marshal(objectMap)
}

// Usage describes Storage Resource Usage.
type Usage struct {
	// Unit - READ-ONLY; Gets the unit of measurement. Possible values include: 'Count', 'Bytes', 'Seconds', 'Percent', 'CountsPerSecond', 'BytesPerSecond'
	Unit UsageUnit `json:"unit,omitempty"`
	// CurrentValue - READ-ONLY; Gets the current count of the allocated resources in the subscription.
	CurrentValue *int32 `json:"currentValue,omitempty"`
	// Limit - READ-ONLY; Gets the maximum count of the resources that can be allocated in the subscription.
	Limit *int32 `json:"limit,omitempty"`
	// Name - READ-ONLY; Gets the name of the type of usage.
	Name *UsageName `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for Usage.
func (u Usage) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// UsageListResult the response from the List Usages operation.
type UsageListResult struct {
	autorest.Response `json:"-"`
	// Value - Gets or sets the list of Storage Resource Usages.
	Value *[]Usage `json:"value,omitempty"`
}

// UsageName the usage names that can be used; currently limited to StorageAccount.
type UsageName struct {
	// Value - READ-ONLY; Gets a string describing the resource name.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - READ-ONLY; Gets a localized string describing the resource name.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// MarshalJSON is the custom marshaler for UsageName.
func (un UsageName) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}
