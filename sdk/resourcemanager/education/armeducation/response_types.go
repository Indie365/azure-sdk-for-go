//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeducation

// GrantsClientGetResponse contains the response from method GrantsClient.Get.
type GrantsClientGetResponse struct {
	// Grant details.
	GrantDetails
}

// GrantsClientListAllResponse contains the response from method GrantsClient.NewListAllPager.
type GrantsClientListAllResponse struct {
	// List of Grants info.
	GrantListResponse
}

// GrantsClientListResponse contains the response from method GrantsClient.NewListPager.
type GrantsClientListResponse struct {
	// List of Grants info.
	GrantListResponse
}

// JoinRequestsClientApproveResponse contains the response from method JoinRequestsClient.Approve.
type JoinRequestsClientApproveResponse struct {
	// placeholder for future response values
}

// JoinRequestsClientDenyResponse contains the response from method JoinRequestsClient.Deny.
type JoinRequestsClientDenyResponse struct {
	// placeholder for future response values
}

// JoinRequestsClientGetResponse contains the response from method JoinRequestsClient.Get.
type JoinRequestsClientGetResponse struct {
	// join requests.
	JoinRequestDetails
}

// JoinRequestsClientListResponse contains the response from method JoinRequestsClient.NewListPager.
type JoinRequestsClientListResponse struct {
	// list of join requests.
	JoinRequestList
}

// LabsClientCreateOrUpdateResponse contains the response from method LabsClient.CreateOrUpdate.
type LabsClientCreateOrUpdateResponse struct {
	// Lab details.
	LabDetails
}

// LabsClientDeleteResponse contains the response from method LabsClient.Delete.
type LabsClientDeleteResponse struct {
	// placeholder for future response values
}

// LabsClientGenerateInviteCodeResponse contains the response from method LabsClient.GenerateInviteCode.
type LabsClientGenerateInviteCodeResponse struct {
	// Lab details.
	LabDetails
}

// LabsClientGetResponse contains the response from method LabsClient.Get.
type LabsClientGetResponse struct {
	// Lab details.
	LabDetails
}

// LabsClientListAllResponse contains the response from method LabsClient.NewListAllPager.
type LabsClientListAllResponse struct {
	// List of labs.
	LabListResult
}

// LabsClientListResponse contains the response from method LabsClient.NewListPager.
type LabsClientListResponse struct {
	// List of labs.
	LabListResult
}

// ManagementClientRedeemInvitationCodeResponse contains the response from method ManagementClient.RedeemInvitationCode.
type ManagementClientRedeemInvitationCodeResponse struct {
	// placeholder for future response values
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// StudentLabsClientGetResponse contains the response from method StudentLabsClient.Get.
type StudentLabsClientGetResponse struct {
	// Student lab details.
	StudentLabDetails
}

// StudentLabsClientListAllResponse contains the response from method StudentLabsClient.NewListAllPager.
type StudentLabsClientListAllResponse struct {
	// List of labs.
	StudentLabListResult
}

// StudentsClientCreateOrUpdateResponse contains the response from method StudentsClient.CreateOrUpdate.
type StudentsClientCreateOrUpdateResponse struct {
	// Student details.
	StudentDetails
}

// StudentsClientDeleteResponse contains the response from method StudentsClient.Delete.
type StudentsClientDeleteResponse struct {
	// placeholder for future response values
}

// StudentsClientGetResponse contains the response from method StudentsClient.Get.
type StudentsClientGetResponse struct {
	// Student details.
	StudentDetails
}

// StudentsClientListResponse contains the response from method StudentsClient.NewListPager.
type StudentsClientListResponse struct {
	// List of students.
	StudentListResult
}
