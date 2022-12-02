//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

// DpsCertificateClientCreateOrUpdateResponse contains the response from method DpsCertificateClient.CreateOrUpdate.
type DpsCertificateClientCreateOrUpdateResponse struct {
	CertificateResponse
}

// DpsCertificateClientDeleteResponse contains the response from method DpsCertificateClient.Delete.
type DpsCertificateClientDeleteResponse struct {
	// placeholder for future response values
}

// DpsCertificateClientGenerateVerificationCodeResponse contains the response from method DpsCertificateClient.GenerateVerificationCode.
type DpsCertificateClientGenerateVerificationCodeResponse struct {
	VerificationCodeResponse
}

// DpsCertificateClientGetResponse contains the response from method DpsCertificateClient.Get.
type DpsCertificateClientGetResponse struct {
	CertificateResponse
}

// DpsCertificateClientListResponse contains the response from method DpsCertificateClient.List.
type DpsCertificateClientListResponse struct {
	CertificateListDescription
}

// DpsCertificateClientVerifyCertificateResponse contains the response from method DpsCertificateClient.VerifyCertificate.
type DpsCertificateClientVerifyCertificateResponse struct {
	CertificateResponse
}

// IotDpsResourceClientCheckProvisioningServiceNameAvailabilityResponse contains the response from method IotDpsResourceClient.CheckProvisioningServiceNameAvailability.
type IotDpsResourceClientCheckProvisioningServiceNameAvailabilityResponse struct {
	NameAvailabilityInfo
}

// IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionResponse contains the response from method IotDpsResourceClient.CreateOrUpdatePrivateEndpointConnection.
type IotDpsResourceClientCreateOrUpdatePrivateEndpointConnectionResponse struct {
	PrivateEndpointConnection
}

// IotDpsResourceClientCreateOrUpdateResponse contains the response from method IotDpsResourceClient.CreateOrUpdate.
type IotDpsResourceClientCreateOrUpdateResponse struct {
	ProvisioningServiceDescription
}

// IotDpsResourceClientDeletePrivateEndpointConnectionResponse contains the response from method IotDpsResourceClient.DeletePrivateEndpointConnection.
type IotDpsResourceClientDeletePrivateEndpointConnectionResponse struct {
	PrivateEndpointConnection
}

// IotDpsResourceClientDeleteResponse contains the response from method IotDpsResourceClient.Delete.
type IotDpsResourceClientDeleteResponse struct {
	// placeholder for future response values
}

// IotDpsResourceClientGetOperationResultResponse contains the response from method IotDpsResourceClient.GetOperationResult.
type IotDpsResourceClientGetOperationResultResponse struct {
	AsyncOperationResult
}

// IotDpsResourceClientGetPrivateEndpointConnectionResponse contains the response from method IotDpsResourceClient.GetPrivateEndpointConnection.
type IotDpsResourceClientGetPrivateEndpointConnectionResponse struct {
	PrivateEndpointConnection
}

// IotDpsResourceClientGetPrivateLinkResourcesResponse contains the response from method IotDpsResourceClient.GetPrivateLinkResources.
type IotDpsResourceClientGetPrivateLinkResourcesResponse struct {
	GroupIDInformation
}

// IotDpsResourceClientGetResponse contains the response from method IotDpsResourceClient.Get.
type IotDpsResourceClientGetResponse struct {
	ProvisioningServiceDescription
}

// IotDpsResourceClientListByResourceGroupResponse contains the response from method IotDpsResourceClient.ListByResourceGroup.
type IotDpsResourceClientListByResourceGroupResponse struct {
	ProvisioningServiceDescriptionListResult
}

// IotDpsResourceClientListBySubscriptionResponse contains the response from method IotDpsResourceClient.ListBySubscription.
type IotDpsResourceClientListBySubscriptionResponse struct {
	ProvisioningServiceDescriptionListResult
}

// IotDpsResourceClientListKeysForKeyNameResponse contains the response from method IotDpsResourceClient.ListKeysForKeyName.
type IotDpsResourceClientListKeysForKeyNameResponse struct {
	SharedAccessSignatureAuthorizationRuleAccessRightsDescription
}

// IotDpsResourceClientListKeysResponse contains the response from method IotDpsResourceClient.ListKeys.
type IotDpsResourceClientListKeysResponse struct {
	SharedAccessSignatureAuthorizationRuleListResult
}

// IotDpsResourceClientListPrivateEndpointConnectionsResponse contains the response from method IotDpsResourceClient.ListPrivateEndpointConnections.
type IotDpsResourceClientListPrivateEndpointConnectionsResponse struct {
	// The list of private endpoint connections for a provisioning service
	PrivateEndpointConnectionArray []*PrivateEndpointConnection
}

// IotDpsResourceClientListPrivateLinkResourcesResponse contains the response from method IotDpsResourceClient.ListPrivateLinkResources.
type IotDpsResourceClientListPrivateLinkResourcesResponse struct {
	PrivateLinkResources
}

// IotDpsResourceClientListValidSKUsResponse contains the response from method IotDpsResourceClient.ListValidSKUs.
type IotDpsResourceClientListValidSKUsResponse struct {
	IotDpsSKUDefinitionListResult
}

// IotDpsResourceClientUpdateResponse contains the response from method IotDpsResourceClient.Update.
type IotDpsResourceClientUpdateResponse struct {
	ProvisioningServiceDescription
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}