//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azkeys

// BackupKeyResponse contains the response from method Client.BackupKey.
type BackupKeyResponse struct {
	BackupKeyResult
}

// CreateKeyResponse contains the response from method Client.CreateKey.
type CreateKeyResponse struct {
	KeyBundle
}

// DecryptResponse contains the response from method Client.Decrypt.
type DecryptResponse struct {
	KeyOperationResult
}

// DeleteKeyResponse contains the response from method Client.DeleteKey.
type DeleteKeyResponse struct {
	DeletedKeyBundle
}

// EncryptResponse contains the response from method Client.Encrypt.
type EncryptResponse struct {
	KeyOperationResult
}

// GetDeletedKeyResponse contains the response from method Client.GetDeletedKey.
type GetDeletedKeyResponse struct {
	DeletedKeyBundle
}

// GetKeyResponse contains the response from method Client.GetKey.
type GetKeyResponse struct {
	KeyBundle
}

// GetKeyRotationPolicyResponse contains the response from method Client.GetKeyRotationPolicy.
type GetKeyRotationPolicyResponse struct {
	KeyRotationPolicy
}

// GetRandomBytesResponse contains the response from method Client.GetRandomBytes.
type GetRandomBytesResponse struct {
	RandomBytes
}

// ImportKeyResponse contains the response from method Client.ImportKey.
type ImportKeyResponse struct {
	KeyBundle
}

// ListDeletedKeysResponse contains the response from method Client.ListDeletedKeys.
type ListDeletedKeysResponse struct {
	DeletedKeyListResult
}

// ListKeyVersionsResponse contains the response from method Client.ListKeyVersions.
type ListKeyVersionsResponse struct {
	KeyListResult
}

// ListKeysResponse contains the response from method Client.ListKeys.
type ListKeysResponse struct {
	KeyListResult
}

// PurgeDeletedKeyResponse contains the response from method Client.PurgeDeletedKey.
type PurgeDeletedKeyResponse struct {
	// placeholder for future response values
}

// RecoverDeletedKeyResponse contains the response from method Client.RecoverDeletedKey.
type RecoverDeletedKeyResponse struct {
	KeyBundle
}

// ReleaseResponse contains the response from method Client.Release.
type ReleaseResponse struct {
	KeyReleaseResult
}

// RestoreKeyResponse contains the response from method Client.RestoreKey.
type RestoreKeyResponse struct {
	KeyBundle
}

// RotateKeyResponse contains the response from method Client.RotateKey.
type RotateKeyResponse struct {
	KeyBundle
}

// SignResponse contains the response from method Client.Sign.
type SignResponse struct {
	KeyOperationResult
}

// UnwrapKeyResponse contains the response from method Client.UnwrapKey.
type UnwrapKeyResponse struct {
	KeyOperationResult
}

// UpdateKeyResponse contains the response from method Client.UpdateKey.
type UpdateKeyResponse struct {
	KeyBundle
}

// UpdateKeyRotationPolicyResponse contains the response from method Client.UpdateKeyRotationPolicy.
type UpdateKeyRotationPolicyResponse struct {
	KeyRotationPolicy
}

// VerifyResponse contains the response from method Client.Verify.
type VerifyResponse struct {
	KeyVerifyResult
}

// WrapKeyResponse contains the response from method Client.WrapKey.
type WrapKeyResponse struct {
	KeyOperationResult
}