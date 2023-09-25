//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

import "time"

// DpsCertificateClientCreateOrUpdateOptions contains the optional parameters for the DpsCertificateClient.CreateOrUpdate
// method.
type DpsCertificateClientCreateOrUpdateOptions struct {
	// ETag of the certificate. This is required to update an existing certificate, and ignored while creating a brand new certificate.
	IfMatch *string
}

// DpsCertificateClientDeleteOptions contains the optional parameters for the DpsCertificateClient.Delete method.
type DpsCertificateClientDeleteOptions struct {
	// Time the certificate is created.
	CertificateCreated *time.Time

	// Indicates if the certificate contains a private key.
	CertificateHasPrivateKey *bool

	// Indicates if certificate has been verified by owner of the private key.
	CertificateIsVerified *bool

	// Time the certificate is last updated.
	CertificateLastUpdated *time.Time

	// This is optional, and it is the Common Name of the certificate.
	CertificateName1 *string

	// Random number generated to indicate Proof of Possession.
	CertificateNonce *string

	// A description that mentions the purpose of the certificate.
	CertificatePurpose *CertificatePurpose

	// Raw data within the certificate.
	CertificateRawBytes []byte
}

// DpsCertificateClientGenerateVerificationCodeOptions contains the optional parameters for the DpsCertificateClient.GenerateVerificationCode
// method.
type DpsCertificateClientGenerateVerificationCodeOptions struct {
	// Certificate creation time.
	CertificateCreated *time.Time

	// Indicates if the certificate contains private key.
	CertificateHasPrivateKey *bool

	// Indicates if the certificate has been verified by owner of the private key.
	CertificateIsVerified *bool

	// Certificate last updated time.
	CertificateLastUpdated *time.Time

	// Common Name for the certificate.
	CertificateName1 *string

	// Random number generated to indicate Proof of Possession.
	CertificateNonce *string

	// Description mentioning the purpose of the certificate.
	CertificatePurpose *CertificatePurpose

	// Raw data of certificate.
	CertificateRawBytes []byte
}

// DpsCertificateClientGetOptions contains the optional parameters for the DpsCertificateClient.Get method.
type DpsCertificateClientGetOptions struct {
	// ETag of the certificate.
	IfMatch *string
}

// DpsCertificateClientListOptions contains the optional parameters for the DpsCertificateClient.List method.
type DpsCertificateClientListOptions struct {
	// placeholder for future optional parameters
}

// DpsCertificateClientVerifyCertificateOptions contains the optional parameters for the DpsCertificateClient.VerifyCertificate
// method.
type DpsCertificateClientVerifyCertificateOptions struct {
	// Certificate creation time.
	CertificateCreated *time.Time

	// Indicates if the certificate contains private key.
	CertificateHasPrivateKey *bool

	// Indicates if the certificate has been verified by owner of the private key.
	CertificateIsVerified *bool

	// Certificate last updated time.
	CertificateLastUpdated *time.Time

	// Common Name for the certificate.
	CertificateName1 *string

	// Random number generated to indicate Proof of Possession.
	CertificateNonce *string

	// Describe the purpose of the certificate.
	CertificatePurpose *CertificatePurpose

	// Raw data of certificate.
	CertificateRawBytes []byte
}

// IotDpsResourceClientBeginCreateOrUpdateOptions contains the optional parameters for the IotDpsResourceClient.BeginCreateOrUpdate
// method.
type IotDpsResourceClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginCreateOrUpdatePrivateEndpointConnectionOptions contains the optional parameters for the IotDpsResourceClient.BeginCreateOrUpdatePrivateEndpointConnection
// method.
type IotDpsResourceClientBeginCreateOrUpdatePrivateEndpointConnectionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginDeleteOptions contains the optional parameters for the IotDpsResourceClient.BeginDelete method.
type IotDpsResourceClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginDeletePrivateEndpointConnectionOptions contains the optional parameters for the IotDpsResourceClient.BeginDeletePrivateEndpointConnection
// method.
type IotDpsResourceClientBeginDeletePrivateEndpointConnectionOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientBeginUpdateOptions contains the optional parameters for the IotDpsResourceClient.BeginUpdate method.
type IotDpsResourceClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// IotDpsResourceClientCheckProvisioningServiceNameAvailabilityOptions contains the optional parameters for the IotDpsResourceClient.CheckProvisioningServiceNameAvailability
// method.
type IotDpsResourceClientCheckProvisioningServiceNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetOperationResultOptions contains the optional parameters for the IotDpsResourceClient.GetOperationResult
// method.
type IotDpsResourceClientGetOperationResultOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetOptions contains the optional parameters for the IotDpsResourceClient.Get method.
type IotDpsResourceClientGetOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetPrivateEndpointConnectionOptions contains the optional parameters for the IotDpsResourceClient.GetPrivateEndpointConnection
// method.
type IotDpsResourceClientGetPrivateEndpointConnectionOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientGetPrivateLinkResourcesOptions contains the optional parameters for the IotDpsResourceClient.GetPrivateLinkResources
// method.
type IotDpsResourceClientGetPrivateLinkResourcesOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListByResourceGroupOptions contains the optional parameters for the IotDpsResourceClient.NewListByResourceGroupPager
// method.
type IotDpsResourceClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListBySubscriptionOptions contains the optional parameters for the IotDpsResourceClient.NewListBySubscriptionPager
// method.
type IotDpsResourceClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListKeysForKeyNameOptions contains the optional parameters for the IotDpsResourceClient.ListKeysForKeyName
// method.
type IotDpsResourceClientListKeysForKeyNameOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListKeysOptions contains the optional parameters for the IotDpsResourceClient.NewListKeysPager method.
type IotDpsResourceClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListPrivateEndpointConnectionsOptions contains the optional parameters for the IotDpsResourceClient.ListPrivateEndpointConnections
// method.
type IotDpsResourceClientListPrivateEndpointConnectionsOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListPrivateLinkResourcesOptions contains the optional parameters for the IotDpsResourceClient.ListPrivateLinkResources
// method.
type IotDpsResourceClientListPrivateLinkResourcesOptions struct {
	// placeholder for future optional parameters
}

// IotDpsResourceClientListValidSKUsOptions contains the optional parameters for the IotDpsResourceClient.NewListValidSKUsPager
// method.
type IotDpsResourceClientListValidSKUsOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

