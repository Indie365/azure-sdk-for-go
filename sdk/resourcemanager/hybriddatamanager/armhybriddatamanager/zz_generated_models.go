//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

import "time"

// AvailableProviderOperation - Class represents provider operation
type AvailableProviderOperation struct {
	// REQUIRED; Gets or Sets Name of the operations
	Name *string `json:"name,omitempty"`

	// Gets or sets Display information Contains the localized display information for this particular operation/action
	Display *AvailableProviderOperationDisplay `json:"display,omitempty"`

	// Gets or sets Origin The intended executor of the operation; governs the display of the operation in the RBAC UX and the
	// audit logs UX. Default value is “user,system”
	Origin *string `json:"origin,omitempty"`

	// Gets or sets Properties Reserved for future use
	Properties interface{} `json:"properties,omitempty"`
}

// AvailableProviderOperationDisplay - Contains the localized display information for this particular operation / action.
// These value will be used by several clients for (1) custom role definitions for RBAC; (2) complex query filters for
// the event service; and (3) audit history / records for management operations.
type AvailableProviderOperationDisplay struct {
	// Gets or sets Description The localized friendly description for the operation, as it should be shown to the user. It should
	// be thorough, yet concise – it will be used in tool tips and detailed views.
	Description *string `json:"description,omitempty"`

	// Gets or sets Operation The localized friendly name for the operation, as it should be shown to the user. It should be concise
	// (to fit in drop downs) but clear (i.e. self-documenting). It should use
	// Title Casing and include the entity/resource to which it applies.
	Operation *string `json:"operation,omitempty"`

	// Gets or sets Provider The localized friendly form of the resource provider name – it is expected to also include the publisher/company
	// responsible. It should use Title Casing and begin with
	// “Microsoft” for 1st party services.
	Provider *string `json:"provider,omitempty"`

	// Gets or sets Resource The localized friendly form of the resource type related to this action/operation – it should match
	// the public documentation for the resource provider. It should use Title Casing
	// – for examples, please refer to the “name” section.
	Resource *string `json:"resource,omitempty"`
}

// AvailableProviderOperations - Class for set of operations used for discovery of available provider operations.
type AvailableProviderOperations struct {
	// Link for the next set of operations.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations.
	Value []*AvailableProviderOperation `json:"value,omitempty"`
}

// CustomerSecret - The pair of customer secret.
type CustomerSecret struct {
	// REQUIRED; The encryption algorithm used to encrypt data.
	Algorithm *SupportedAlgorithm `json:"algorithm,omitempty"`

	// REQUIRED; The identifier to the data service input object which this secret corresponds to.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`

	// REQUIRED; It contains the encrypted customer secret.
	KeyValue *string `json:"keyValue,omitempty"`
}

// DataManager - The DataManager resource.
type DataManager struct {
	// REQUIRED; The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US,
	// East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it
	// is created, but if an identical geo region is specified on update the request will succeed.
	Location *string `json:"location,omitempty"`

	// Etag of the Resource.
	Etag *string `json:"etag,omitempty"`

	// The sku type.
	SKU *SKU `json:"sku,omitempty"`

	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across
	// resource groups).
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The Resource Id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The Resource Name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DataManagerList - DataManager resources Collection.
type DataManagerList struct {
	// Link for the next set of data stores.
	NextLink *string `json:"nextLink,omitempty"`

	// List of data manager resources.
	Value []*DataManager `json:"value,omitempty"`
}

// DataManagerUpdateParameter - The DataManagerUpdateParameter.
type DataManagerUpdateParameter struct {
	// The sku type.
	SKU *SKU `json:"sku,omitempty"`

	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across
	// resource groups).
	Tags map[string]*string `json:"tags,omitempty"`
}

// DataManagersClientBeginCreateOptions contains the optional parameters for the DataManagersClient.BeginCreate method.
type DataManagersClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DataManagersClientBeginDeleteOptions contains the optional parameters for the DataManagersClient.BeginDelete method.
type DataManagersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DataManagersClientBeginUpdateOptions contains the optional parameters for the DataManagersClient.BeginUpdate method.
type DataManagersClientBeginUpdateOptions struct {
	// Defines the If-Match condition. The patch will be performed only if the ETag of the data manager resource on the server
	// matches this value.
	IfMatch *string
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DataManagersClientGetOptions contains the optional parameters for the DataManagersClient.Get method.
type DataManagersClientGetOptions struct {
	// placeholder for future optional parameters
}

// DataManagersClientListByResourceGroupOptions contains the optional parameters for the DataManagersClient.ListByResourceGroup
// method.
type DataManagersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DataManagersClientListOptions contains the optional parameters for the DataManagersClient.List method.
type DataManagersClientListOptions struct {
	// placeholder for future optional parameters
}

// DataService - Data Service.
type DataService struct {
	// REQUIRED; DataService properties.
	Properties *DataServiceProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the object.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the object.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DataServiceList - Data Service Collection.
type DataServiceList struct {
	// Link for the next set of data services.
	NextLink *string `json:"nextLink,omitempty"`

	// List of data services.
	Value []*DataService `json:"value,omitempty"`
}

// DataServiceProperties - Data Service properties.
type DataServiceProperties struct {
	// REQUIRED; State of the data service.
	State *State `json:"state,omitempty"`

	// Supported data store types which can be used as a sink.
	SupportedDataSinkTypes []*string `json:"supportedDataSinkTypes,omitempty"`

	// Supported data store types which can be used as a source.
	SupportedDataSourceTypes []*string `json:"supportedDataSourceTypes,omitempty"`
}

// DataServicesClientGetOptions contains the optional parameters for the DataServicesClient.Get method.
type DataServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// DataServicesClientListByDataManagerOptions contains the optional parameters for the DataServicesClient.ListByDataManager
// method.
type DataServicesClientListByDataManagerOptions struct {
	// placeholder for future optional parameters
}

// DataStore - Data store.
type DataStore struct {
	// REQUIRED; DataStore properties.
	Properties *DataStoreProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the object.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the object.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DataStoreFilter - Contains the information about the filters for the DataStore.
type DataStoreFilter struct {
	// The data store type id.
	DataStoreTypeID *string `json:"dataStoreTypeId,omitempty"`
}

// DataStoreList - Data Store Collection.
type DataStoreList struct {
	// Link for the next set of data stores.
	NextLink *string `json:"nextLink,omitempty"`

	// List of data stores.
	Value []*DataStore `json:"value,omitempty"`
}

// DataStoreProperties - Data Store for sources and sinks
type DataStoreProperties struct {
	// REQUIRED; The arm id of the data store type.
	DataStoreTypeID *string `json:"dataStoreTypeId,omitempty"`

	// REQUIRED; State of the data source.
	State *State `json:"state,omitempty"`

	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source
	// to understand the key. Value contains customer secret encrypted by the
	// encryptionKeys.
	CustomerSecrets []*CustomerSecret `json:"customerSecrets,omitempty"`

	// A generic json used differently by each data source type.
	ExtendedProperties interface{} `json:"extendedProperties,omitempty"`

	// Arm Id for the manager resource to which the data source is associated. This is optional.
	RepositoryID *string `json:"repositoryId,omitempty"`
}

// DataStoreType - Data Store Type.
type DataStoreType struct {
	// REQUIRED; DataStoreType properties.
	Properties *DataStoreTypeProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the object.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the object.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DataStoreTypeList - Data Store Type Collection.
type DataStoreTypeList struct {
	// Link for the next set of data store types.
	NextLink *string `json:"nextLink,omitempty"`

	// List of DataStoreType.
	Value []*DataStoreType `json:"value,omitempty"`
}

// DataStoreTypeProperties - Data Store Type properties.
type DataStoreTypeProperties struct {
	// REQUIRED; State of the data store type.
	State *State `json:"state,omitempty"`

	// Arm type for the manager resource to which the data source type is associated. This is optional.
	RepositoryType *string `json:"repositoryType,omitempty"`

	// Supported data services where it can be used as a sink.
	SupportedDataServicesAsSink []*string `json:"supportedDataServicesAsSink,omitempty"`

	// Supported data services where it can be used as a source.
	SupportedDataServicesAsSource []*string `json:"supportedDataServicesAsSource,omitempty"`
}

// DataStoreTypesClientGetOptions contains the optional parameters for the DataStoreTypesClient.Get method.
type DataStoreTypesClientGetOptions struct {
	// placeholder for future optional parameters
}

// DataStoreTypesClientListByDataManagerOptions contains the optional parameters for the DataStoreTypesClient.ListByDataManager
// method.
type DataStoreTypesClientListByDataManagerOptions struct {
	// placeholder for future optional parameters
}

// DataStoresClientBeginCreateOrUpdateOptions contains the optional parameters for the DataStoresClient.BeginCreateOrUpdate
// method.
type DataStoresClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DataStoresClientBeginDeleteOptions contains the optional parameters for the DataStoresClient.BeginDelete method.
type DataStoresClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DataStoresClientGetOptions contains the optional parameters for the DataStoresClient.Get method.
type DataStoresClientGetOptions struct {
	// placeholder for future optional parameters
}

// DataStoresClientListByDataManagerOptions contains the optional parameters for the DataStoresClient.ListByDataManager method.
type DataStoresClientListByDataManagerOptions struct {
	// OData Filter options
	Filter *string
}

// DmsBaseObject - Base class for all objects under DataManager Service
type DmsBaseObject struct {
	// READ-ONLY; Id of the object.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the object.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Error - Top level error for the job.
type Error struct {
	// REQUIRED; Error code that can be used to programmatically identify the error.
	Code *string `json:"code,omitempty"`

	// Describes the error in detail and provides debugging information.
	Message *string `json:"message,omitempty"`
}

// ErrorDetails - Error Details
type ErrorDetails struct {
	// Error code.
	ErrorCode *int32 `json:"errorCode,omitempty"`

	// Error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Contains the non localized exception message
	ExceptionMessage *string `json:"exceptionMessage,omitempty"`

	// Recommended action for the error.
	RecommendedAction *string `json:"recommendedAction,omitempty"`
}

// Job - Data service job.
type Job struct {
	// REQUIRED; Job properties.
	Properties *JobProperties `json:"properties,omitempty"`

	// REQUIRED; Time at which the job was started in UTC ISO 8601 format.
	StartTime *time.Time `json:"startTime,omitempty"`

	// REQUIRED; Status of the job.
	Status *JobStatus `json:"status,omitempty"`

	// Time at which the job ended in UTC ISO 8601 format.
	EndTime *time.Time `json:"endTime,omitempty"`

	// Top level error for the job.
	Error *Error `json:"error,omitempty"`

	// READ-ONLY; Id of the object.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the object.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// JobDefinition - Job Definition.
type JobDefinition struct {
	// REQUIRED; JobDefinition properties.
	Properties *JobDefinitionProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the object.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the object.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// JobDefinitionFilter - Contains the supported job definition filters.
type JobDefinitionFilter struct {
	// REQUIRED; The state of the job definition.
	State *State `json:"state,omitempty"`

	// The data source associated with the job definition
	DataSource *string `json:"dataSource,omitempty"`

	// The last modified date time of the data source.
	LastModified *time.Time `json:"lastModified,omitempty"`
}

// JobDefinitionList - Job Definition Collection.
type JobDefinitionList struct {
	// Link for the next set of job definitions.
	NextLink *string `json:"nextLink,omitempty"`

	// List of job definitions.
	Value []*JobDefinition `json:"value,omitempty"`
}

// JobDefinitionProperties - Job Definition
type JobDefinitionProperties struct {
	// REQUIRED; Data Sink Id associated to the job definition.
	DataSinkID *string `json:"dataSinkId,omitempty"`

	// REQUIRED; Data Source Id associated to the job definition.
	DataSourceID *string `json:"dataSourceId,omitempty"`

	// REQUIRED; State of the job definition.
	State *State `json:"state,omitempty"`

	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source
	// to understand the key. Value contains customer secret encrypted by the
	// encryptionKeys.
	CustomerSecrets []*CustomerSecret `json:"customerSecrets,omitempty"`

	// A generic json used differently by each data service type.
	DataServiceInput interface{} `json:"dataServiceInput,omitempty"`

	// Last modified time of the job definition.
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`

	// This is the preferred geo location for the job to run.
	RunLocation *RunLocation `json:"runLocation,omitempty"`

	// Schedule for running the job definition
	Schedules []*Schedule `json:"schedules,omitempty"`

	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation *UserConfirmation `json:"userConfirmation,omitempty"`
}

// JobDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the JobDefinitionsClient.BeginCreateOrUpdate
// method.
type JobDefinitionsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// JobDefinitionsClientBeginDeleteOptions contains the optional parameters for the JobDefinitionsClient.BeginDelete method.
type JobDefinitionsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// JobDefinitionsClientBeginRunOptions contains the optional parameters for the JobDefinitionsClient.BeginRun method.
type JobDefinitionsClientBeginRunOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// JobDefinitionsClientGetOptions contains the optional parameters for the JobDefinitionsClient.Get method.
type JobDefinitionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// JobDefinitionsClientListByDataManagerOptions contains the optional parameters for the JobDefinitionsClient.ListByDataManager
// method.
type JobDefinitionsClientListByDataManagerOptions struct {
	// OData Filter options
	Filter *string
}

// JobDefinitionsClientListByDataServiceOptions contains the optional parameters for the JobDefinitionsClient.ListByDataService
// method.
type JobDefinitionsClientListByDataServiceOptions struct {
	// OData Filter options
	Filter *string
}

// JobDetails - Job details.
type JobDetails struct {
	// Error details for failure. This is optional.
	ErrorDetails []*ErrorDetails `json:"errorDetails,omitempty"`

	// Item Details Link to download files or see details
	ItemDetailsLink *string `json:"itemDetailsLink,omitempty"`

	// JobDefinition at the time of the run
	JobDefinition *JobDefinition `json:"jobDefinition,omitempty"`

	// List of stages that ran in the job
	JobStages []*JobStages `json:"jobStages,omitempty"`
}

// JobFilter - Contains the information about the filters for the job.
type JobFilter struct {
	// REQUIRED; The status of the job.
	Status *JobStatus `json:"status,omitempty"`

	// The start time of the job.
	StartTime *time.Time `json:"startTime,omitempty"`
}

// JobList - Job Collection.
type JobList struct {
	// Link for the next set of jobs.
	NextLink *string `json:"nextLink,omitempty"`

	// List of jobs.
	Value []*Job `json:"value,omitempty"`
}

// JobProperties - Job Properties
type JobProperties struct {
	// REQUIRED; Describes whether the job is cancellable.
	IsCancellable *IsJobCancellable `json:"isCancellable,omitempty"`

	// Number of bytes processed by the job as of now.
	BytesProcessed *int64 `json:"bytesProcessed,omitempty"`

	// Name of the data sink on which the job was triggered.
	DataSinkName *string `json:"dataSinkName,omitempty"`

	// Name of the data source on which the job was triggered.
	DataSourceName *string `json:"dataSourceName,omitempty"`

	// Details of a job run. This field will only be sent for expand details filter.
	Details *JobDetails `json:"details,omitempty"`

	// Number of items processed by the job as of now
	ItemsProcessed *int64 `json:"itemsProcessed,omitempty"`

	// Number of bytes to be processed by the job in total.
	TotalBytesToProcess *int64 `json:"totalBytesToProcess,omitempty"`

	// Number of items to be processed by the job in total
	TotalItemsToProcess *int64 `json:"totalItemsToProcess,omitempty"`
}

// JobStages - Job stages.
type JobStages struct {
	// REQUIRED; Status of the job stage.
	StageStatus *JobStatus `json:"stageStatus,omitempty"`

	// Error details for the stage. This is optional
	ErrorDetails []*ErrorDetails `json:"errorDetails,omitempty"`

	// Job Stage Details
	JobStageDetails interface{} `json:"jobStageDetails,omitempty"`

	// Name of the job stage.
	StageName *string `json:"stageName,omitempty"`
}

// JobsClientBeginCancelOptions contains the optional parameters for the JobsClient.BeginCancel method.
type JobsClientBeginCancelOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// JobsClientBeginResumeOptions contains the optional parameters for the JobsClient.BeginResume method.
type JobsClientBeginResumeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
type JobsClientGetOptions struct {
	// $expand is supported on details parameter for job, which provides details on the job stages.
	Expand *string
}

// JobsClientListByDataManagerOptions contains the optional parameters for the JobsClient.ListByDataManager method.
type JobsClientListByDataManagerOptions struct {
	// OData Filter options
	Filter *string
}

// JobsClientListByDataServiceOptions contains the optional parameters for the JobsClient.ListByDataService method.
type JobsClientListByDataServiceOptions struct {
	// OData Filter options
	Filter *string
}

// JobsClientListByJobDefinitionOptions contains the optional parameters for the JobsClient.ListByJobDefinition method.
type JobsClientListByJobDefinitionOptions struct {
	// OData Filter options
	Filter *string
}

// Key - Encryption Key.
type Key struct {
	// REQUIRED; The maximum byte size that can be encrypted by the key. For a key size larger than the size, break into chunks
	// and encrypt each chunk, append each encrypted chunk with : to mark the end of the chunk.
	EncryptionChunkSizeInBytes *int32 `json:"encryptionChunkSizeInBytes,omitempty"`

	// REQUIRED; Exponent of the encryption key.
	KeyExponent *string `json:"keyExponent,omitempty"`

	// REQUIRED; Modulus of the encryption key.
	KeyModulus *string `json:"keyModulus,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PublicKey - Public key
type PublicKey struct {
	// REQUIRED; Public key property.
	Properties *PublicKeyProperties `json:"properties,omitempty"`

	// READ-ONLY; Id of the object.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the object.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Type of the object.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PublicKeyList - PublicKey Collection
type PublicKeyList struct {
	// Link for the next set of public keys.
	NextLink *string `json:"nextLink,omitempty"`

	// List of public keys.
	Value []*PublicKey `json:"value,omitempty"`
}

// PublicKeyProperties - PublicKey Properties
type PublicKeyProperties struct {
	// REQUIRED; Level one public key for encryption
	DataServiceLevel1Key *Key `json:"dataServiceLevel1Key,omitempty"`

	// REQUIRED; Level two public key for encryption
	DataServiceLevel2Key *Key `json:"dataServiceLevel2Key,omitempty"`
}

// PublicKeysClientGetOptions contains the optional parameters for the PublicKeysClient.Get method.
type PublicKeysClientGetOptions struct {
	// placeholder for future optional parameters
}

// PublicKeysClientListByDataManagerOptions contains the optional parameters for the PublicKeysClient.ListByDataManager method.
type PublicKeysClientListByDataManagerOptions struct {
	// placeholder for future optional parameters
}

// Resource - Model of the Resource.
type Resource struct {
	// REQUIRED; The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US,
	// East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it
	// is created, but if an identical geo region is specified on update the request will succeed.
	Location *string `json:"location,omitempty"`

	// The sku type.
	SKU *SKU `json:"sku,omitempty"`

	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across
	// resource groups).
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The Resource Id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The Resource Name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RunParameters - Run parameters for a job.
type RunParameters struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source
	// to understand the key. Value contains customer secret encrypted by the
	// encryptionKeys.
	CustomerSecrets []*CustomerSecret `json:"customerSecrets,omitempty"`

	// A generic json used differently by each data service type.
	DataServiceInput interface{} `json:"dataServiceInput,omitempty"`

	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation *UserConfirmation `json:"userConfirmation,omitempty"`
}

// SKU - The sku type.
type SKU struct {
	// The sku name. Required for data manager creation, optional for update.
	Name *string `json:"name,omitempty"`

	// The sku tier. This is based on the SKU name.
	Tier *string `json:"tier,omitempty"`
}

// Schedule for the job run.
type Schedule struct {
	// Name of the schedule.
	Name *string `json:"name,omitempty"`

	// A list of repetition intervals in ISO 8601 format.
	PolicyList []*string `json:"policyList,omitempty"`
}