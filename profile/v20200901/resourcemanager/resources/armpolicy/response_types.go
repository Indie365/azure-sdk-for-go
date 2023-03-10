//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpolicy

// AssignmentsClientCreateByIDResponse contains the response from method AssignmentsClient.CreateByID.
type AssignmentsClientCreateByIDResponse struct {
	Assignment
}

// AssignmentsClientCreateResponse contains the response from method AssignmentsClient.Create.
type AssignmentsClientCreateResponse struct {
	Assignment
}

// AssignmentsClientDeleteByIDResponse contains the response from method AssignmentsClient.DeleteByID.
type AssignmentsClientDeleteByIDResponse struct {
	Assignment
}

// AssignmentsClientDeleteResponse contains the response from method AssignmentsClient.Delete.
type AssignmentsClientDeleteResponse struct {
	Assignment
}

// AssignmentsClientGetByIDResponse contains the response from method AssignmentsClient.GetByID.
type AssignmentsClientGetByIDResponse struct {
	Assignment
}

// AssignmentsClientGetResponse contains the response from method AssignmentsClient.Get.
type AssignmentsClientGetResponse struct {
	Assignment
}

// AssignmentsClientListForResourceGroupResponse contains the response from method AssignmentsClient.NewListForResourceGroupPager.
type AssignmentsClientListForResourceGroupResponse struct {
	AssignmentListResult
}

// AssignmentsClientListForResourceResponse contains the response from method AssignmentsClient.NewListForResourcePager.
type AssignmentsClientListForResourceResponse struct {
	AssignmentListResult
}

// AssignmentsClientListResponse contains the response from method AssignmentsClient.NewListPager.
type AssignmentsClientListResponse struct {
	AssignmentListResult
}

// DefinitionsClientCreateOrUpdateAtManagementGroupResponse contains the response from method DefinitionsClient.CreateOrUpdateAtManagementGroup.
type DefinitionsClientCreateOrUpdateAtManagementGroupResponse struct {
	Definition
}

// DefinitionsClientCreateOrUpdateResponse contains the response from method DefinitionsClient.CreateOrUpdate.
type DefinitionsClientCreateOrUpdateResponse struct {
	Definition
}

// DefinitionsClientDeleteAtManagementGroupResponse contains the response from method DefinitionsClient.DeleteAtManagementGroup.
type DefinitionsClientDeleteAtManagementGroupResponse struct {
	// placeholder for future response values
}

// DefinitionsClientDeleteResponse contains the response from method DefinitionsClient.Delete.
type DefinitionsClientDeleteResponse struct {
	// placeholder for future response values
}

// DefinitionsClientGetAtManagementGroupResponse contains the response from method DefinitionsClient.GetAtManagementGroup.
type DefinitionsClientGetAtManagementGroupResponse struct {
	Definition
}

// DefinitionsClientGetBuiltInResponse contains the response from method DefinitionsClient.GetBuiltIn.
type DefinitionsClientGetBuiltInResponse struct {
	Definition
}

// DefinitionsClientGetResponse contains the response from method DefinitionsClient.Get.
type DefinitionsClientGetResponse struct {
	Definition
}

// DefinitionsClientListBuiltInResponse contains the response from method DefinitionsClient.NewListBuiltInPager.
type DefinitionsClientListBuiltInResponse struct {
	DefinitionListResult
}

// DefinitionsClientListByManagementGroupResponse contains the response from method DefinitionsClient.NewListByManagementGroupPager.
type DefinitionsClientListByManagementGroupResponse struct {
	DefinitionListResult
}

// DefinitionsClientListResponse contains the response from method DefinitionsClient.NewListPager.
type DefinitionsClientListResponse struct {
	DefinitionListResult
}
