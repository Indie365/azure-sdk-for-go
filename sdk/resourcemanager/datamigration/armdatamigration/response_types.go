//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

// FilesClientCreateOrUpdateResponse contains the response from method FilesClient.CreateOrUpdate.
type FilesClientCreateOrUpdateResponse struct {
	// A file resource
	ProjectFile
}

// FilesClientDeleteResponse contains the response from method FilesClient.Delete.
type FilesClientDeleteResponse struct {
	// placeholder for future response values
}

// FilesClientGetResponse contains the response from method FilesClient.Get.
type FilesClientGetResponse struct {
	// A file resource
	ProjectFile
}

// FilesClientListResponse contains the response from method FilesClient.NewListPager.
type FilesClientListResponse struct {
	// OData page of files
	FileList
}

// FilesClientReadResponse contains the response from method FilesClient.Read.
type FilesClientReadResponse struct {
	// File storage information.
	FileStorageInfo
}

// FilesClientReadWriteResponse contains the response from method FilesClient.ReadWrite.
type FilesClientReadWriteResponse struct {
	// File storage information.
	FileStorageInfo
}

// FilesClientUpdateResponse contains the response from method FilesClient.Update.
type FilesClientUpdateResponse struct {
	// A file resource
	ProjectFile
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// OData page of action (operation) objects
	ServiceOperationList
}

// ProjectsClientCreateOrUpdateResponse contains the response from method ProjectsClient.CreateOrUpdate.
type ProjectsClientCreateOrUpdateResponse struct {
	// A project resource
	Project
}

// ProjectsClientDeleteResponse contains the response from method ProjectsClient.Delete.
type ProjectsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProjectsClientGetResponse contains the response from method ProjectsClient.Get.
type ProjectsClientGetResponse struct {
	// A project resource
	Project
}

// ProjectsClientListResponse contains the response from method ProjectsClient.NewListPager.
type ProjectsClientListResponse struct {
	// OData page of project resources
	ProjectList
}

// ProjectsClientUpdateResponse contains the response from method ProjectsClient.Update.
type ProjectsClientUpdateResponse struct {
	// A project resource
	Project
}

// ResourceSKUsClientListSKUsResponse contains the response from method ResourceSKUsClient.NewListSKUsPager.
type ResourceSKUsClientListSKUsResponse struct {
	// The DMS List SKUs operation response.
	ResourceSKUsResult
}

// ServiceTasksClientCancelResponse contains the response from method ServiceTasksClient.Cancel.
type ServiceTasksClientCancelResponse struct {
	// A task resource
	ProjectTask
}

// ServiceTasksClientCreateOrUpdateResponse contains the response from method ServiceTasksClient.CreateOrUpdate.
type ServiceTasksClientCreateOrUpdateResponse struct {
	// A task resource
	ProjectTask
}

// ServiceTasksClientDeleteResponse contains the response from method ServiceTasksClient.Delete.
type ServiceTasksClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceTasksClientGetResponse contains the response from method ServiceTasksClient.Get.
type ServiceTasksClientGetResponse struct {
	// A task resource
	ProjectTask
}

// ServiceTasksClientListResponse contains the response from method ServiceTasksClient.NewListPager.
type ServiceTasksClientListResponse struct {
	// OData page of tasks
	TaskList
}

// ServiceTasksClientUpdateResponse contains the response from method ServiceTasksClient.Update.
type ServiceTasksClientUpdateResponse struct {
	// A task resource
	ProjectTask
}

// ServicesClientCheckChildrenNameAvailabilityResponse contains the response from method ServicesClient.CheckChildrenNameAvailability.
type ServicesClientCheckChildrenNameAvailabilityResponse struct {
	// Indicates whether a proposed resource name is available
	NameAvailabilityResponse
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	// Indicates whether a proposed resource name is available
	NameAvailabilityResponse
}

// ServicesClientCheckStatusResponse contains the response from method ServicesClient.CheckStatus.
type ServicesClientCheckStatusResponse struct {
	// Service health status
	ServiceStatusResponse
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// A Database Migration Service resource
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// A Database Migration Service resource
	Service
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// OData page of service objects
	ServiceList
}

// ServicesClientListResponse contains the response from method ServicesClient.NewListPager.
type ServicesClientListResponse struct {
	// OData page of service objects
	ServiceList
}

// ServicesClientListSKUsResponse contains the response from method ServicesClient.NewListSKUsPager.
type ServicesClientListSKUsResponse struct {
	// OData page of available SKUs
	ServiceSKUList
}

// ServicesClientStartResponse contains the response from method ServicesClient.BeginStart.
type ServicesClientStartResponse struct {
	// placeholder for future response values
}

// ServicesClientStopResponse contains the response from method ServicesClient.BeginStop.
type ServicesClientStopResponse struct {
	// placeholder for future response values
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.BeginUpdate.
type ServicesClientUpdateResponse struct {
	// A Database Migration Service resource
	Service
}

// TasksClientCancelResponse contains the response from method TasksClient.Cancel.
type TasksClientCancelResponse struct {
	// A task resource
	ProjectTask
}

// TasksClientCommandResponse contains the response from method TasksClient.Command.
type TasksClientCommandResponse struct {
	// Base class for all types of DMS command properties. If command is not supported by current client, this object is returned.
	CommandPropertiesClassification
}

	// UnmarshalJSON implements the json.Unmarshaller interface for type TasksClientCommandResponse.
func (t *TasksClientCommandResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalCommandPropertiesClassification(data)
	if err != nil {
		return err
	}
	t.CommandPropertiesClassification = res
	return nil
}

// TasksClientCreateOrUpdateResponse contains the response from method TasksClient.CreateOrUpdate.
type TasksClientCreateOrUpdateResponse struct {
	// A task resource
	ProjectTask
}

// TasksClientDeleteResponse contains the response from method TasksClient.Delete.
type TasksClientDeleteResponse struct {
	// placeholder for future response values
}

// TasksClientGetResponse contains the response from method TasksClient.Get.
type TasksClientGetResponse struct {
	// A task resource
	ProjectTask
}

// TasksClientListResponse contains the response from method TasksClient.NewListPager.
type TasksClientListResponse struct {
	// OData page of tasks
	TaskList
}

// TasksClientUpdateResponse contains the response from method TasksClient.Update.
type TasksClientUpdateResponse struct {
	// A task resource
	ProjectTask
}

// UsagesClientListResponse contains the response from method UsagesClient.NewListPager.
type UsagesClientListResponse struct {
	// OData page of quota objects
	QuotaList
}

