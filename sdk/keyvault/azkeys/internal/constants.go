//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package internal

const (
	module  = "internal"
	version = "v0.1.0"
)

// ActionType - The type of the action.
type ActionType string

const (
	// ActionTypeRotate - Rotate the key based on the key policy.
	ActionTypeRotate ActionType = "rotate"
	// ActionTypeNotify - Trigger event grid events. For preview, the notification time is not configurable and it is default to 30 days before expiry.
	ActionTypeNotify ActionType = "notify"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeRotate,
		ActionTypeNotify,
	}
}

// ToPtr returns a *ActionType pointing to the current value.
func (c ActionType) ToPtr() *ActionType {
	return &c
}

// DataAction - Supported permissions for data actions.
type DataAction string

const (
	// DataActionBackupHsmKeys - Backup HSM keys.
	DataActionBackupHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/backup/action"
	// DataActionCreateHsmKey - Create an HSM key.
	DataActionCreateHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/create"
	// DataActionDecryptHsmKey - Decrypt using an HSM key.
	DataActionDecryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/decrypt/action"
	// DataActionDeleteHsmKey - Delete an HSM key.
	DataActionDeleteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/delete"
	// DataActionDeleteRoleAssignment - Delete role assignment.
	DataActionDeleteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/delete/action"
	// DataActionDeleteRoleDefinition - Delete role definition.
	DataActionDeleteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/delete/action"
	// DataActionDownloadHsmSecurityDomain - Download an HSM security domain.
	DataActionDownloadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/action"
	// DataActionDownloadHsmSecurityDomainStatus - Check status of HSM security domain download.
	DataActionDownloadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/download/read"
	// DataActionEncryptHsmKey - Encrypt using an HSM key.
	DataActionEncryptHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/encrypt/action"
	// DataActionExportHsmKey - Export an HSM key.
	DataActionExportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/export/action"
	// DataActionGetRoleAssignment - Get role assignment.
	DataActionGetRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/read/action"
	// DataActionImportHsmKey - Import an HSM key.
	DataActionImportHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/import/action"
	// DataActionPurgeDeletedHsmKey - Purge a deleted HSM key.
	DataActionPurgeDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/delete"
	// DataActionRandomNumbersGenerate - Generate random numbers.
	DataActionRandomNumbersGenerate DataAction = "Microsoft.KeyVault/managedHsm/rng/action"
	// DataActionReadDeletedHsmKey - Read deleted HSM key.
	DataActionReadDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/read/action"
	// DataActionReadHsmBackupStatus - Read an HSM backup status.
	DataActionReadHsmBackupStatus DataAction = "Microsoft.KeyVault/managedHsm/backup/status/action"
	// DataActionReadHsmKey - Read HSM key metadata.
	DataActionReadHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/read/action"
	// DataActionReadHsmRestoreStatus - Read an HSM restore status.
	DataActionReadHsmRestoreStatus DataAction = "Microsoft.KeyVault/managedHsm/restore/status/action"
	// DataActionReadHsmSecurityDomainStatus - Check the status of the HSM security domain exchange file.
	DataActionReadHsmSecurityDomainStatus DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/read"
	// DataActionReadHsmSecurityDomainTransferKey - Download an HSM security domain transfer key.
	DataActionReadHsmSecurityDomainTransferKey DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/transferkey/read"
	// DataActionReadRoleDefinition - Get role definition.
	DataActionReadRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/read/action"
	// DataActionRecoverDeletedHsmKey - Recover deleted HSM key.
	DataActionRecoverDeletedHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/deletedKeys/recover/action"
	// DataActionReleaseKey - Release an HSM key using Secure Key Release.
	DataActionReleaseKey DataAction = "Microsoft.KeyVault/managedHsm/keys/release/action"
	// DataActionRestoreHsmKeys - Restore HSM keys.
	DataActionRestoreHsmKeys DataAction = "Microsoft.KeyVault/managedHsm/keys/restore/action"
	// DataActionSignHsmKey - Sign using an HSM key.
	DataActionSignHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/sign/action"
	// DataActionStartHsmBackup - Start an HSM backup.
	DataActionStartHsmBackup DataAction = "Microsoft.KeyVault/managedHsm/backup/start/action"
	// DataActionStartHsmRestore - Start an HSM restore.
	DataActionStartHsmRestore DataAction = "Microsoft.KeyVault/managedHsm/restore/start/action"
	// DataActionUnwrapHsmKey - Unwrap using an HSM key.
	DataActionUnwrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/unwrap/action"
	// DataActionUploadHsmSecurityDomain - Upload an HSM security domain.
	DataActionUploadHsmSecurityDomain DataAction = "Microsoft.KeyVault/managedHsm/securitydomain/upload/action"
	// DataActionVerifyHsmKey - Verify using an HSM key.
	DataActionVerifyHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/verify/action"
	// DataActionWrapHsmKey - Wrap using an HSM key.
	DataActionWrapHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/wrap/action"
	// DataActionWriteHsmKey - Update an HSM key.
	DataActionWriteHsmKey DataAction = "Microsoft.KeyVault/managedHsm/keys/write/action"
	// DataActionWriteRoleAssignment - Create or update role assignment.
	DataActionWriteRoleAssignment DataAction = "Microsoft.KeyVault/managedHsm/roleAssignments/write/action"
	// DataActionWriteRoleDefinition - Create or update role definition.
	DataActionWriteRoleDefinition DataAction = "Microsoft.KeyVault/managedHsm/roleDefinitions/write/action"
)

// PossibleDataActionValues returns the possible values for the DataAction const type.
func PossibleDataActionValues() []DataAction {
	return []DataAction{
		DataActionBackupHsmKeys,
		DataActionCreateHsmKey,
		DataActionDecryptHsmKey,
		DataActionDeleteHsmKey,
		DataActionDeleteRoleAssignment,
		DataActionDeleteRoleDefinition,
		DataActionDownloadHsmSecurityDomain,
		DataActionDownloadHsmSecurityDomainStatus,
		DataActionEncryptHsmKey,
		DataActionExportHsmKey,
		DataActionGetRoleAssignment,
		DataActionImportHsmKey,
		DataActionPurgeDeletedHsmKey,
		DataActionRandomNumbersGenerate,
		DataActionReadDeletedHsmKey,
		DataActionReadHsmBackupStatus,
		DataActionReadHsmKey,
		DataActionReadHsmRestoreStatus,
		DataActionReadHsmSecurityDomainStatus,
		DataActionReadHsmSecurityDomainTransferKey,
		DataActionReadRoleDefinition,
		DataActionRecoverDeletedHsmKey,
		DataActionReleaseKey,
		DataActionRestoreHsmKeys,
		DataActionSignHsmKey,
		DataActionStartHsmBackup,
		DataActionStartHsmRestore,
		DataActionUnwrapHsmKey,
		DataActionUploadHsmSecurityDomain,
		DataActionVerifyHsmKey,
		DataActionWrapHsmKey,
		DataActionWriteHsmKey,
		DataActionWriteRoleAssignment,
		DataActionWriteRoleDefinition,
	}
}

// ToPtr returns a *DataAction pointing to the current value.
func (c DataAction) ToPtr() *DataAction {
	return &c
}

// DeletionRecoveryLevel - Reflects the deletion recovery level currently in effect for keys in the current vault. If it contains 'Purgeable' the key can
// be permanently deleted by a privileged user; otherwise, only the system
// can purge the key, at the end of the retention interval.
type DeletionRecoveryLevel string

const (
	// DeletionRecoveryLevelCustomizedRecoverable - Denotes a vault state in which deletion is recoverable without the possibility for immediate and permanent
	// deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90).This level guarantees the recoverability of the deleted entity during the retention interval
	// and while the subscription is still available.
	DeletionRecoveryLevelCustomizedRecoverable DeletionRecoveryLevel = "CustomizedRecoverable"
	// DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription - Denotes a vault and subscription state in which deletion is recoverable, immediate
	// and permanent deletion (i.e. purge) is not permitted, and in which the subscription itself cannot be permanently canceled when 7<= SoftDeleteRetentionInDays
	// < 90. This level guarantees the recoverability of the deleted entity during the retention interval, and also reflects the fact that the subscription
	// itself cannot be cancelled.
	DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription DeletionRecoveryLevel = "CustomizedRecoverable+ProtectedSubscription"
	// DeletionRecoveryLevelCustomizedRecoverablePurgeable - Denotes a vault state in which deletion is recoverable, and which also permits immediate and permanent
	// deletion (i.e. purge when 7<= SoftDeleteRetentionInDays < 90). This level guarantees the recoverability of the deleted entity during the retention interval,
	// unless a Purge operation is requested, or the subscription is cancelled.
	DeletionRecoveryLevelCustomizedRecoverablePurgeable DeletionRecoveryLevel = "CustomizedRecoverable+Purgeable"
	// DeletionRecoveryLevelPurgeable - Denotes a vault state in which deletion is an irreversible operation, without the possibility for recovery. This level
	// corresponds to no protection being available against a Delete operation; the data is irretrievably lost upon accepting a Delete operation at the entity
	// level or higher (vault, resource group, subscription etc.)
	DeletionRecoveryLevelPurgeable DeletionRecoveryLevel = "Purgeable"
	// DeletionRecoveryLevelRecoverable - Denotes a vault state in which deletion is recoverable without the possibility for immediate and permanent deletion
	// (i.e. purge). This level guarantees the recoverability of the deleted entity during the retention interval(90 days) and while the subscription is still
	// available. System wil permanently delete it after 90 days, if not recovered
	DeletionRecoveryLevelRecoverable DeletionRecoveryLevel = "Recoverable"
	// DeletionRecoveryLevelRecoverableProtectedSubscription - Denotes a vault and subscription state in which deletion is recoverable within retention interval
	// (90 days), immediate and permanent deletion (i.e. purge) is not permitted, and in which the subscription itself cannot be permanently canceled. System
	// wil permanently delete it after 90 days, if not recovered
	DeletionRecoveryLevelRecoverableProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"
	// DeletionRecoveryLevelRecoverablePurgeable - Denotes a vault state in which deletion is recoverable, and which also permits immediate and permanent deletion
	// (i.e. purge). This level guarantees the recoverability of the deleted entity during the retention interval (90 days), unless a Purge operation is requested,
	// or the subscription is cancelled. System wil permanently delete it after 90 days, if not recovered
	DeletionRecoveryLevelRecoverablePurgeable DeletionRecoveryLevel = "Recoverable+Purgeable"
)

// PossibleDeletionRecoveryLevelValues returns the possible values for the DeletionRecoveryLevel const type.
func PossibleDeletionRecoveryLevelValues() []DeletionRecoveryLevel {
	return []DeletionRecoveryLevel{
		DeletionRecoveryLevelCustomizedRecoverable,
		DeletionRecoveryLevelCustomizedRecoverableProtectedSubscription,
		DeletionRecoveryLevelCustomizedRecoverablePurgeable,
		DeletionRecoveryLevelPurgeable,
		DeletionRecoveryLevelRecoverable,
		DeletionRecoveryLevelRecoverableProtectedSubscription,
		DeletionRecoveryLevelRecoverablePurgeable,
	}
}

// ToPtr returns a *DeletionRecoveryLevel pointing to the current value.
func (c DeletionRecoveryLevel) ToPtr() *DeletionRecoveryLevel {
	return &c
}

// JSONWebKeyCurveName - Elliptic curve name. For valid values, see JsonWebKeyCurveName.
type JSONWebKeyCurveName string

const (
	// JSONWebKeyCurveNameP256 - The NIST P-256 elliptic curve, AKA SECG curve SECP256R1.
	JSONWebKeyCurveNameP256 JSONWebKeyCurveName = "P-256"
	// JSONWebKeyCurveNameP256K - The SECG SECP256K1 elliptic curve.
	JSONWebKeyCurveNameP256K JSONWebKeyCurveName = "P-256K"
	// JSONWebKeyCurveNameP384 - The NIST P-384 elliptic curve, AKA SECG curve SECP384R1.
	JSONWebKeyCurveNameP384 JSONWebKeyCurveName = "P-384"
	// JSONWebKeyCurveNameP521 - The NIST P-521 elliptic curve, AKA SECG curve SECP521R1.
	JSONWebKeyCurveNameP521 JSONWebKeyCurveName = "P-521"
)

// PossibleJSONWebKeyCurveNameValues returns the possible values for the JSONWebKeyCurveName const type.
func PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName {
	return []JSONWebKeyCurveName{
		JSONWebKeyCurveNameP256,
		JSONWebKeyCurveNameP256K,
		JSONWebKeyCurveNameP384,
		JSONWebKeyCurveNameP521,
	}
}

// ToPtr returns a *JSONWebKeyCurveName pointing to the current value.
func (c JSONWebKeyCurveName) ToPtr() *JSONWebKeyCurveName {
	return &c
}

// JSONWebKeyEncryptionAlgorithm - algorithm identifier
type JSONWebKeyEncryptionAlgorithm string

const (
	JSONWebKeyEncryptionAlgorithmA128CBC    JSONWebKeyEncryptionAlgorithm = "A128CBC"
	JSONWebKeyEncryptionAlgorithmA128CBCPAD JSONWebKeyEncryptionAlgorithm = "A128CBCPAD"
	JSONWebKeyEncryptionAlgorithmA128GCM    JSONWebKeyEncryptionAlgorithm = "A128GCM"
	JSONWebKeyEncryptionAlgorithmA128KW     JSONWebKeyEncryptionAlgorithm = "A128KW"
	JSONWebKeyEncryptionAlgorithmA192CBC    JSONWebKeyEncryptionAlgorithm = "A192CBC"
	JSONWebKeyEncryptionAlgorithmA192CBCPAD JSONWebKeyEncryptionAlgorithm = "A192CBCPAD"
	JSONWebKeyEncryptionAlgorithmA192GCM    JSONWebKeyEncryptionAlgorithm = "A192GCM"
	JSONWebKeyEncryptionAlgorithmA192KW     JSONWebKeyEncryptionAlgorithm = "A192KW"
	JSONWebKeyEncryptionAlgorithmA256CBC    JSONWebKeyEncryptionAlgorithm = "A256CBC"
	JSONWebKeyEncryptionAlgorithmA256CBCPAD JSONWebKeyEncryptionAlgorithm = "A256CBCPAD"
	JSONWebKeyEncryptionAlgorithmA256GCM    JSONWebKeyEncryptionAlgorithm = "A256GCM"
	JSONWebKeyEncryptionAlgorithmA256KW     JSONWebKeyEncryptionAlgorithm = "A256KW"
	JSONWebKeyEncryptionAlgorithmRSA15      JSONWebKeyEncryptionAlgorithm = "RSA1_5"
	JSONWebKeyEncryptionAlgorithmRSAOAEP    JSONWebKeyEncryptionAlgorithm = "RSA-OAEP"
	JSONWebKeyEncryptionAlgorithmRSAOAEP256 JSONWebKeyEncryptionAlgorithm = "RSA-OAEP-256"
)

// PossibleJSONWebKeyEncryptionAlgorithmValues returns the possible values for the JSONWebKeyEncryptionAlgorithm const type.
func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
	return []JSONWebKeyEncryptionAlgorithm{
		JSONWebKeyEncryptionAlgorithmA128CBC,
		JSONWebKeyEncryptionAlgorithmA128CBCPAD,
		JSONWebKeyEncryptionAlgorithmA128GCM,
		JSONWebKeyEncryptionAlgorithmA128KW,
		JSONWebKeyEncryptionAlgorithmA192CBC,
		JSONWebKeyEncryptionAlgorithmA192CBCPAD,
		JSONWebKeyEncryptionAlgorithmA192GCM,
		JSONWebKeyEncryptionAlgorithmA192KW,
		JSONWebKeyEncryptionAlgorithmA256CBC,
		JSONWebKeyEncryptionAlgorithmA256CBCPAD,
		JSONWebKeyEncryptionAlgorithmA256GCM,
		JSONWebKeyEncryptionAlgorithmA256KW,
		JSONWebKeyEncryptionAlgorithmRSA15,
		JSONWebKeyEncryptionAlgorithmRSAOAEP,
		JSONWebKeyEncryptionAlgorithmRSAOAEP256,
	}
}

// ToPtr returns a *JSONWebKeyEncryptionAlgorithm pointing to the current value.
func (c JSONWebKeyEncryptionAlgorithm) ToPtr() *JSONWebKeyEncryptionAlgorithm {
	return &c
}

// JSONWebKeyOperation - JSON web key operations. For more information, see JsonWebKeyOperation.
type JSONWebKeyOperation string

const (
	JSONWebKeyOperationDecrypt   JSONWebKeyOperation = "decrypt"
	JSONWebKeyOperationEncrypt   JSONWebKeyOperation = "encrypt"
	JSONWebKeyOperationExport    JSONWebKeyOperation = "export"
	JSONWebKeyOperationImport    JSONWebKeyOperation = "import"
	JSONWebKeyOperationSign      JSONWebKeyOperation = "sign"
	JSONWebKeyOperationUnwrapKey JSONWebKeyOperation = "unwrapKey"
	JSONWebKeyOperationVerify    JSONWebKeyOperation = "verify"
	JSONWebKeyOperationWrapKey   JSONWebKeyOperation = "wrapKey"
)

// PossibleJSONWebKeyOperationValues returns the possible values for the JSONWebKeyOperation const type.
func PossibleJSONWebKeyOperationValues() []JSONWebKeyOperation {
	return []JSONWebKeyOperation{
		JSONWebKeyOperationDecrypt,
		JSONWebKeyOperationEncrypt,
		JSONWebKeyOperationExport,
		JSONWebKeyOperationImport,
		JSONWebKeyOperationSign,
		JSONWebKeyOperationUnwrapKey,
		JSONWebKeyOperationVerify,
		JSONWebKeyOperationWrapKey,
	}
}

// ToPtr returns a *JSONWebKeyOperation pointing to the current value.
func (c JSONWebKeyOperation) ToPtr() *JSONWebKeyOperation {
	return &c
}

// JSONWebKeySignatureAlgorithm - The signing/verification algorithm identifier. For more information on possible algorithm types, see JsonWebKeySignatureAlgorithm.
type JSONWebKeySignatureAlgorithm string

const (
	// JSONWebKeySignatureAlgorithmES256 - ECDSA using P-256 and SHA-256, as described in https://tools.ietf.org/html/rfc7518.
	JSONWebKeySignatureAlgorithmES256 JSONWebKeySignatureAlgorithm = "ES256"
	// JSONWebKeySignatureAlgorithmES256K - ECDSA using P-256K and SHA-256, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmES256K JSONWebKeySignatureAlgorithm = "ES256K"
	// JSONWebKeySignatureAlgorithmES384 - ECDSA using P-384 and SHA-384, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmES384 JSONWebKeySignatureAlgorithm = "ES384"
	// JSONWebKeySignatureAlgorithmES512 - ECDSA using P-521 and SHA-512, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmES512 JSONWebKeySignatureAlgorithm = "ES512"
	// JSONWebKeySignatureAlgorithmPS256 - RSASSA-PSS using SHA-256 and MGF1 with SHA-256, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmPS256 JSONWebKeySignatureAlgorithm = "PS256"
	// JSONWebKeySignatureAlgorithmPS384 - RSASSA-PSS using SHA-384 and MGF1 with SHA-384, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmPS384 JSONWebKeySignatureAlgorithm = "PS384"
	// JSONWebKeySignatureAlgorithmPS512 - RSASSA-PSS using SHA-512 and MGF1 with SHA-512, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmPS512 JSONWebKeySignatureAlgorithm = "PS512"
	// JSONWebKeySignatureAlgorithmRS256 - RSASSA-PKCS1-v1_5 using SHA-256, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmRS256 JSONWebKeySignatureAlgorithm = "RS256"
	// JSONWebKeySignatureAlgorithmRS384 - RSASSA-PKCS1-v1_5 using SHA-384, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmRS384 JSONWebKeySignatureAlgorithm = "RS384"
	// JSONWebKeySignatureAlgorithmRS512 - RSASSA-PKCS1-v1_5 using SHA-512, as described in https://tools.ietf.org/html/rfc7518
	JSONWebKeySignatureAlgorithmRS512 JSONWebKeySignatureAlgorithm = "RS512"
	// JSONWebKeySignatureAlgorithmRSNULL - Reserved
	JSONWebKeySignatureAlgorithmRSNULL JSONWebKeySignatureAlgorithm = "RSNULL"
)

// PossibleJSONWebKeySignatureAlgorithmValues returns the possible values for the JSONWebKeySignatureAlgorithm const type.
func PossibleJSONWebKeySignatureAlgorithmValues() []JSONWebKeySignatureAlgorithm {
	return []JSONWebKeySignatureAlgorithm{
		JSONWebKeySignatureAlgorithmES256,
		JSONWebKeySignatureAlgorithmES256K,
		JSONWebKeySignatureAlgorithmES384,
		JSONWebKeySignatureAlgorithmES512,
		JSONWebKeySignatureAlgorithmPS256,
		JSONWebKeySignatureAlgorithmPS384,
		JSONWebKeySignatureAlgorithmPS512,
		JSONWebKeySignatureAlgorithmRS256,
		JSONWebKeySignatureAlgorithmRS384,
		JSONWebKeySignatureAlgorithmRS512,
		JSONWebKeySignatureAlgorithmRSNULL,
	}
}

// ToPtr returns a *JSONWebKeySignatureAlgorithm pointing to the current value.
func (c JSONWebKeySignatureAlgorithm) ToPtr() *JSONWebKeySignatureAlgorithm {
	return &c
}

// JSONWebKeyType - JsonWebKey Key Type (kty), as defined in https://tools.ietf.org/html/draft-ietf-jose-json-web-algorithms-40.
type JSONWebKeyType string

const (
	// JSONWebKeyTypeEC - Elliptic Curve.
	JSONWebKeyTypeEC JSONWebKeyType = "EC"
	// JSONWebKeyTypeECHSM - Elliptic Curve with a private key which is stored in the HSM.
	JSONWebKeyTypeECHSM JSONWebKeyType = "EC-HSM"
	// JSONWebKeyTypeOct - Octet sequence (used to represent symmetric keys)
	JSONWebKeyTypeOct JSONWebKeyType = "oct"
	// JSONWebKeyTypeOctHSM - Octet sequence (used to represent symmetric keys) which is stored the HSM.
	JSONWebKeyTypeOctHSM JSONWebKeyType = "oct-HSM"
	// JSONWebKeyTypeRSA - RSA (https://tools.ietf.org/html/rfc3447)
	JSONWebKeyTypeRSA JSONWebKeyType = "RSA"
	// JSONWebKeyTypeRSAHSM - RSA with a private key which is stored in the HSM.
	JSONWebKeyTypeRSAHSM JSONWebKeyType = "RSA-HSM"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeEC,
		JSONWebKeyTypeECHSM,
		JSONWebKeyTypeOct,
		JSONWebKeyTypeOctHSM,
		JSONWebKeyTypeRSA,
		JSONWebKeyTypeRSAHSM,
	}
}

// ToPtr returns a *JSONWebKeyType pointing to the current value.
func (c JSONWebKeyType) ToPtr() *JSONWebKeyType {
	return &c
}

// KeyEncryptionAlgorithm - The encryption algorithm to use to protected the exported key material
type KeyEncryptionAlgorithm string

const (
	KeyEncryptionAlgorithmCKMRSAAESKEYWRAP KeyEncryptionAlgorithm = "CKM_RSA_AES_KEY_WRAP"
	KeyEncryptionAlgorithmRSAAESKEYWRAP256 KeyEncryptionAlgorithm = "RSA_AES_KEY_WRAP_256"
	KeyEncryptionAlgorithmRSAAESKEYWRAP384 KeyEncryptionAlgorithm = "RSA_AES_KEY_WRAP_384"
)

// PossibleKeyEncryptionAlgorithmValues returns the possible values for the KeyEncryptionAlgorithm const type.
func PossibleKeyEncryptionAlgorithmValues() []KeyEncryptionAlgorithm {
	return []KeyEncryptionAlgorithm{
		KeyEncryptionAlgorithmCKMRSAAESKEYWRAP,
		KeyEncryptionAlgorithmRSAAESKEYWRAP256,
		KeyEncryptionAlgorithmRSAAESKEYWRAP384,
	}
}

// ToPtr returns a *KeyEncryptionAlgorithm pointing to the current value.
func (c KeyEncryptionAlgorithm) ToPtr() *KeyEncryptionAlgorithm {
	return &c
}

// OperationStatus - operation status
type OperationStatus string

const (
	OperationStatusSuccess    OperationStatus = "Success"
	OperationStatusInProgress OperationStatus = "InProgress"
	OperationStatusFailed     OperationStatus = "Failed"
)

// PossibleOperationStatusValues returns the possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{
		OperationStatusSuccess,
		OperationStatusInProgress,
		OperationStatusFailed,
	}
}

// ToPtr returns a *OperationStatus pointing to the current value.
func (c OperationStatus) ToPtr() *OperationStatus {
	return &c
}

// RoleDefinitionType - The role definition type.
type RoleDefinitionType string

const (
	RoleDefinitionTypeMicrosoftAuthorizationRoleDefinitions RoleDefinitionType = "Microsoft.Authorization/roleDefinitions"
)

// PossibleRoleDefinitionTypeValues returns the possible values for the RoleDefinitionType const type.
func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
	return []RoleDefinitionType{
		RoleDefinitionTypeMicrosoftAuthorizationRoleDefinitions,
	}
}

// ToPtr returns a *RoleDefinitionType pointing to the current value.
func (c RoleDefinitionType) ToPtr() *RoleDefinitionType {
	return &c
}

// RoleScope - The role scope.
type RoleScope string

const (
	// RoleScopeGlobal - Global scope
	RoleScopeGlobal RoleScope = "/"
	// RoleScopeKeys - Keys scope
	RoleScopeKeys RoleScope = "/keys"
)

// PossibleRoleScopeValues returns the possible values for the RoleScope const type.
func PossibleRoleScopeValues() []RoleScope {
	return []RoleScope{
		RoleScopeGlobal,
		RoleScopeKeys,
	}
}

// ToPtr returns a *RoleScope pointing to the current value.
func (c RoleScope) ToPtr() *RoleScope {
	return &c
}

// RoleType - The role type.
type RoleType string

const (
	// RoleTypeBuiltInRole - Built in role.
	RoleTypeBuiltInRole RoleType = "AKVBuiltInRole"
	// RoleTypeCustomRole - Custom role.
	RoleTypeCustomRole RoleType = "CustomRole"
)

// PossibleRoleTypeValues returns the possible values for the RoleType const type.
func PossibleRoleTypeValues() []RoleType {
	return []RoleType{
		RoleTypeBuiltInRole,
		RoleTypeCustomRole,
	}
}

// ToPtr returns a *RoleType pointing to the current value.
func (c RoleType) ToPtr() *RoleType {
	return &c
}
