//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeploymentmanager

// ArtifactSourcesClientCreateOrUpdateResponse contains the response from method ArtifactSourcesClient.CreateOrUpdate.
type ArtifactSourcesClientCreateOrUpdateResponse struct {
	// The resource that defines the source location where the artifacts are located.
	ArtifactSource
}

// ArtifactSourcesClientDeleteResponse contains the response from method ArtifactSourcesClient.Delete.
type ArtifactSourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// ArtifactSourcesClientGetResponse contains the response from method ArtifactSourcesClient.Get.
type ArtifactSourcesClientGetResponse struct {
	// The resource that defines the source location where the artifacts are located.
	ArtifactSource
}

// ArtifactSourcesClientListResponse contains the response from method ArtifactSourcesClient.List.
type ArtifactSourcesClientListResponse struct {
	// The list of artifact sources.
	ArtifactSourceArray []*ArtifactSource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// The operations response.
	OperationsList
}

// RolloutsClientCancelResponse contains the response from method RolloutsClient.Cancel.
type RolloutsClientCancelResponse struct {
	// Defines the rollout.
	Rollout
}

// RolloutsClientCreateOrUpdateResponse contains the response from method RolloutsClient.BeginCreateOrUpdate.
type RolloutsClientCreateOrUpdateResponse struct {
	// Defines the PUT rollout request body.
	RolloutRequest
}

// RolloutsClientDeleteResponse contains the response from method RolloutsClient.Delete.
type RolloutsClientDeleteResponse struct {
	// placeholder for future response values
}

// RolloutsClientGetResponse contains the response from method RolloutsClient.Get.
type RolloutsClientGetResponse struct {
	// Defines the rollout.
	Rollout
}

// RolloutsClientListResponse contains the response from method RolloutsClient.List.
type RolloutsClientListResponse struct {
	// The list of rollouts.
	RolloutArray []*Rollout
}

// RolloutsClientRestartResponse contains the response from method RolloutsClient.Restart.
type RolloutsClientRestartResponse struct {
	// Defines the rollout.
	Rollout
}

// ServiceTopologiesClientCreateOrUpdateResponse contains the response from method ServiceTopologiesClient.CreateOrUpdate.
type ServiceTopologiesClientCreateOrUpdateResponse struct {
	// The resource representation of a service topology.
	ServiceTopologyResource
}

// ServiceTopologiesClientDeleteResponse contains the response from method ServiceTopologiesClient.Delete.
type ServiceTopologiesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceTopologiesClientGetResponse contains the response from method ServiceTopologiesClient.Get.
type ServiceTopologiesClientGetResponse struct {
	// The resource representation of a service topology.
	ServiceTopologyResource
}

// ServiceTopologiesClientListResponse contains the response from method ServiceTopologiesClient.List.
type ServiceTopologiesClientListResponse struct {
	// The list of service topologies.
	ServiceTopologyResourceArray []*ServiceTopologyResource
}

// ServiceUnitsClientCreateOrUpdateResponse contains the response from method ServiceUnitsClient.BeginCreateOrUpdate.
type ServiceUnitsClientCreateOrUpdateResponse struct {
	// Represents the response of a service unit resource.
	ServiceUnitResource
}

// ServiceUnitsClientDeleteResponse contains the response from method ServiceUnitsClient.Delete.
type ServiceUnitsClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceUnitsClientGetResponse contains the response from method ServiceUnitsClient.Get.
type ServiceUnitsClientGetResponse struct {
	// Represents the response of a service unit resource.
	ServiceUnitResource
}

// ServiceUnitsClientListResponse contains the response from method ServiceUnitsClient.List.
type ServiceUnitsClientListResponse struct {
	// The list of service units.
	ServiceUnitResourceArray []*ServiceUnitResource
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.CreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// The resource representation of a service in a service topology.
	ServiceResource
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// The resource representation of a service in a service topology.
	ServiceResource
}

// ServicesClientListResponse contains the response from method ServicesClient.List.
type ServicesClientListResponse struct {
	// The list of services.
	ServiceResourceArray []*ServiceResource
}

// StepsClientCreateOrUpdateResponse contains the response from method StepsClient.CreateOrUpdate.
type StepsClientCreateOrUpdateResponse struct {
	// The resource representation of a rollout step.
	StepResource
}

// StepsClientDeleteResponse contains the response from method StepsClient.Delete.
type StepsClientDeleteResponse struct {
	// placeholder for future response values
}

// StepsClientGetResponse contains the response from method StepsClient.Get.
type StepsClientGetResponse struct {
	// The resource representation of a rollout step.
	StepResource
}

// StepsClientListResponse contains the response from method StepsClient.List.
type StepsClientListResponse struct {
	// The list of steps.
	StepResourceArray []*StepResource
}

