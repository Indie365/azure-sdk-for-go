//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azcertificates

// BackupCertificateResponse contains the response from method Client.BackupCertificate.
type BackupCertificateResponse struct {
	BackupCertificateResult
}

// CreateCertificateResponse contains the response from method Client.CreateCertificate.
type CreateCertificateResponse struct {
	CertificateOperation
}

// DeleteCertificateContactsResponse contains the response from method Client.DeleteCertificateContacts.
type DeleteCertificateContactsResponse struct {
	Contacts
}

// DeleteCertificateIssuerResponse contains the response from method Client.DeleteCertificateIssuer.
type DeleteCertificateIssuerResponse struct {
	IssuerBundle
}

// DeleteCertificateOperationResponse contains the response from method Client.DeleteCertificateOperation.
type DeleteCertificateOperationResponse struct {
	CertificateOperation
}

// DeleteCertificateResponse contains the response from method Client.DeleteCertificate.
type DeleteCertificateResponse struct {
	DeletedCertificateBundle
}

// GetCertificateContactsResponse contains the response from method Client.GetCertificateContacts.
type GetCertificateContactsResponse struct {
	Contacts
}

// GetCertificateIssuerResponse contains the response from method Client.GetCertificateIssuer.
type GetCertificateIssuerResponse struct {
	IssuerBundle
}

// GetCertificateOperationResponse contains the response from method Client.GetCertificateOperation.
type GetCertificateOperationResponse struct {
	CertificateOperation
}

// GetCertificatePolicyResponse contains the response from method Client.GetCertificatePolicy.
type GetCertificatePolicyResponse struct {
	CertificatePolicy
}

// GetCertificateResponse contains the response from method Client.GetCertificate.
type GetCertificateResponse struct {
	CertificateBundle
}

// GetDeletedCertificateResponse contains the response from method Client.GetDeletedCertificate.
type GetDeletedCertificateResponse struct {
	DeletedCertificateBundle
}

// ImportCertificateResponse contains the response from method Client.ImportCertificate.
type ImportCertificateResponse struct {
	CertificateBundle
}

// ListCertificateIssuersResponse contains the response from method Client.ListCertificateIssuers.
type ListCertificateIssuersResponse struct {
	CertificateIssuerListResult
}

// ListCertificateVersionsResponse contains the response from method Client.ListCertificateVersions.
type ListCertificateVersionsResponse struct {
	CertificateListResult
}

// ListCertificatesResponse contains the response from method Client.ListCertificates.
type ListCertificatesResponse struct {
	CertificateListResult
}

// ListDeletedCertificatesResponse contains the response from method Client.ListDeletedCertificates.
type ListDeletedCertificatesResponse struct {
	DeletedCertificateListResult
}

// MergeCertificateResponse contains the response from method Client.MergeCertificate.
type MergeCertificateResponse struct {
	CertificateBundle
}

// PurgeDeletedCertificateResponse contains the response from method Client.PurgeDeletedCertificate.
type PurgeDeletedCertificateResponse struct {
	// placeholder for future response values
}

// RecoverDeletedCertificateResponse contains the response from method Client.RecoverDeletedCertificate.
type RecoverDeletedCertificateResponse struct {
	CertificateBundle
}

// RestoreCertificateResponse contains the response from method Client.RestoreCertificate.
type RestoreCertificateResponse struct {
	CertificateBundle
}

// SetCertificateContactsResponse contains the response from method Client.SetCertificateContacts.
type SetCertificateContactsResponse struct {
	Contacts
}

// SetCertificateIssuerResponse contains the response from method Client.SetCertificateIssuer.
type SetCertificateIssuerResponse struct {
	IssuerBundle
}

// UpdateCertificateIssuerResponse contains the response from method Client.UpdateCertificateIssuer.
type UpdateCertificateIssuerResponse struct {
	IssuerBundle
}

// UpdateCertificateOperationResponse contains the response from method Client.UpdateCertificateOperation.
type UpdateCertificateOperationResponse struct {
	CertificateOperation
}

// UpdateCertificatePolicyResponse contains the response from method Client.UpdateCertificatePolicy.
type UpdateCertificatePolicyResponse struct {
	CertificatePolicy
}

// UpdateCertificateResponse contains the response from method Client.UpdateCertificate.
type UpdateCertificateResponse struct {
	CertificateBundle
}