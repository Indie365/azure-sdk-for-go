//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azcertificatemanager

//Possible CertificateFileFormat values
type CertificateFileFormat string

const (
	// CertificateFileFormatApplicationPkixCert - Pkix-cert Format
	CertificateFileFormatApplicationPkixCert CertificateFileFormat = "application/pkix-cert"
	// CertificateFileFormatApplicationXPemFile - X-Pem-File Format
	CertificateFileFormatApplicationXPemFile CertificateFileFormat = "application/x-pem-file"
	// CertificateFileFormatApplicationXPkcs7Certificates - X-Pkcs7-Certificates Format
	CertificateFileFormatApplicationXPkcs7Certificates CertificateFileFormat = "application/x-pkcs7-certificates"
)

// PossibleCertificateFileFormatValues returns the possible values for the CertificateFileFormat const type.
func PossibleCertificateFileFormatValues() []CertificateFileFormat {
	return []CertificateFileFormat{
		CertificateFileFormatApplicationPkixCert,
		CertificateFileFormatApplicationXPemFile,
		CertificateFileFormatApplicationXPkcs7Certificates,
	}
}

//Possible CertificateFormat values
type CertificateFormat string

const (
	// CertificateFormatPem - PEM
	CertificateFormatPem CertificateFormat = "pem"
	// CertificateFormatPkcs7 - PKCS7
	CertificateFormatPkcs7 CertificateFormat = "pkcs7"
)

// PossibleCertificateFormatValues returns the possible values for the CertificateFormat const type.
func PossibleCertificateFormatValues() []CertificateFormat {
	return []CertificateFormat{
		CertificateFormatPem,
		CertificateFormatPkcs7,
	}
}

//Possible ChainFormat values
type ChainFormat string

const (
	// ChainFormatP7B - P7B
	ChainFormatP7B ChainFormat = "p7b"
	// ChainFormatPem - PEM
	ChainFormatPem ChainFormat = "pem"
)

// PossibleChainFormatValues returns the possible values for the ChainFormat const type.
func PossibleChainFormatValues() []ChainFormat {
	return []ChainFormat{
		ChainFormatP7B,
		ChainFormatPem,
	}
}

// ExtendedKeyUsage - Acceptable values of ExtendedKeyUsage.
type ExtendedKeyUsage string

const (
	// ExtendedKeyUsageClientAuth - ClientAuth
	ExtendedKeyUsageClientAuth ExtendedKeyUsage = "ClientAuth"
	// ExtendedKeyUsageCodeSigning - CodeSigning
	ExtendedKeyUsageCodeSigning ExtendedKeyUsage = "CodeSigning"
	// ExtendedKeyUsageEmailProtection - EmailProtection
	ExtendedKeyUsageEmailProtection ExtendedKeyUsage = "EmailProtection"
	// ExtendedKeyUsageIPSecEndSystem - IpsecEndSystem
	ExtendedKeyUsageIPSecEndSystem ExtendedKeyUsage = "IpsecEndSystem"
	// ExtendedKeyUsageIPSecTunnel - IpsecTunnel
	ExtendedKeyUsageIPSecTunnel ExtendedKeyUsage = "IpsecTunnel"
	// ExtendedKeyUsageIPSecUser - IpsecUser
	ExtendedKeyUsageIPSecUser ExtendedKeyUsage = "IpsecUser"
	// ExtendedKeyUsageMACAddress - MACAddress
	ExtendedKeyUsageMACAddress ExtendedKeyUsage = "MACAddress"
	// ExtendedKeyUsageOCSPSigning - OCSPSigning
	ExtendedKeyUsageOCSPSigning ExtendedKeyUsage = "OCSPSigning"
	// ExtendedKeyUsageServerAuth - ServerAuth
	ExtendedKeyUsageServerAuth ExtendedKeyUsage = "ServerAuth"
	// ExtendedKeyUsageSmartcardLogon - SmartcardLogon
	ExtendedKeyUsageSmartcardLogon ExtendedKeyUsage = "SmartcardLogon"
	// ExtendedKeyUsageTimeStamping - TimeStamping
	ExtendedKeyUsageTimeStamping ExtendedKeyUsage = "TimeStamping"
)

// PossibleExtendedKeyUsageValues returns the possible values for the ExtendedKeyUsage const type.
func PossibleExtendedKeyUsageValues() []ExtendedKeyUsage {
	return []ExtendedKeyUsage{
		ExtendedKeyUsageClientAuth,
		ExtendedKeyUsageCodeSigning,
		ExtendedKeyUsageEmailProtection,
		ExtendedKeyUsageIPSecEndSystem,
		ExtendedKeyUsageIPSecTunnel,
		ExtendedKeyUsageIPSecUser,
		ExtendedKeyUsageMACAddress,
		ExtendedKeyUsageOCSPSigning,
		ExtendedKeyUsageServerAuth,
		ExtendedKeyUsageSmartcardLogon,
		ExtendedKeyUsageTimeStamping,
	}
}

//Possible include values
type Include string

const (
	// IncludeLogs - Logs
	IncludeLogs Include = "logs"
)

// PossibleIncludeValues returns the possible values for the Include const type.
func PossibleIncludeValues() []Include {
	return []Include{
		IncludeLogs,
	}
}

// KeyUsage - Acceptable values of KeyUsage.
type KeyUsage string

const (
	// KeyUsageCrlSign - CrlSign
	KeyUsageCrlSign KeyUsage = "CrlSign"
	// KeyUsageDataEncipherment - DataEncipherment
	KeyUsageDataEncipherment KeyUsage = "DataEncipherment"
	// KeyUsageDecipherOnly - DecipherOnly
	KeyUsageDecipherOnly KeyUsage = "DecipherOnly"
	// KeyUsageDigitalSignature - DigitalSignature
	KeyUsageDigitalSignature KeyUsage = "DigitalSignature"
	// KeyUsageEncipherOnly - EncipherOnly
	KeyUsageEncipherOnly KeyUsage = "EncipherOnly"
	// KeyUsageKeyAgreement - KeyAgreement
	KeyUsageKeyAgreement KeyUsage = "KeyAgreement"
	// KeyUsageKeyCertSign - KeyCertSign
	KeyUsageKeyCertSign KeyUsage = "KeyCertSign"
	// KeyUsageKeyEncipherment - KeyEncipherment
	KeyUsageKeyEncipherment KeyUsage = "KeyEncipherment"
	// KeyUsageNonRepudiation - NonRepudiation
	KeyUsageNonRepudiation KeyUsage = "NonRepudiation"
)

// PossibleKeyUsageValues returns the possible values for the KeyUsage const type.
func PossibleKeyUsageValues() []KeyUsage {
	return []KeyUsage{
		KeyUsageCrlSign,
		KeyUsageDataEncipherment,
		KeyUsageDecipherOnly,
		KeyUsageDigitalSignature,
		KeyUsageEncipherOnly,
		KeyUsageKeyAgreement,
		KeyUsageKeyCertSign,
		KeyUsageKeyEncipherment,
		KeyUsageNonRepudiation,
	}
}

// RevocationReason - Specifies reason for revocation
type RevocationReason string

const (
	// RevocationReasonAACompromise - AACompromise
	RevocationReasonAACompromise RevocationReason = "aACompromise"
	// RevocationReasonAffiliationChanged - AffiliationChanged
	RevocationReasonAffiliationChanged RevocationReason = "affiliationChanged"
	// RevocationReasonCACompromise - CACompromise
	RevocationReasonCACompromise RevocationReason = "cACompromise"
	// RevocationReasonCertificateHold - CertificateHold
	RevocationReasonCertificateHold RevocationReason = "certificateHold"
	// RevocationReasonCessationOfOperation - CessationOfOperation
	RevocationReasonCessationOfOperation RevocationReason = "cessationOfOperation"
	// RevocationReasonKeyCompromise - KeyCompromise
	RevocationReasonKeyCompromise RevocationReason = "keyCompromise"
	// RevocationReasonPrivilegeWithdrawn - PrivilegeWithdrawn
	RevocationReasonPrivilegeWithdrawn RevocationReason = "privilegeWithdrawn"
	// RevocationReasonRemoveFromCRL - RemoveFromCRL
	RevocationReasonRemoveFromCRL RevocationReason = "removeFromCRL"
	// RevocationReasonSuperseded - Superseded
	RevocationReasonSuperseded RevocationReason = "superseded"
	// RevocationReasonUnspecified - Unspecified
	RevocationReasonUnspecified RevocationReason = "unspecified"
)

// PossibleRevocationReasonValues returns the possible values for the RevocationReason const type.
func PossibleRevocationReasonValues() []RevocationReason {
	return []RevocationReason{
		RevocationReasonAACompromise,
		RevocationReasonAffiliationChanged,
		RevocationReasonCACompromise,
		RevocationReasonCertificateHold,
		RevocationReasonCessationOfOperation,
		RevocationReasonKeyCompromise,
		RevocationReasonPrivilegeWithdrawn,
		RevocationReasonRemoveFromCRL,
		RevocationReasonSuperseded,
		RevocationReasonUnspecified,
	}
}
