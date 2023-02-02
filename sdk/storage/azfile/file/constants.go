//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package file

import "github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/internal/generated"

// NTFSFileAttributes for Files and Directories.
// The subset of attributes is listed at: https://learn.microsoft.com/en-us/rest/api/storageservices/set-file-properties#file-system-attributes.
// Their respective values are listed at: https://learn.microsoft.com/en-us/windows/win32/fileio/file-attribute-constants.
type NTFSFileAttributes uint32

const (
	Readonly          NTFSFileAttributes = 1
	Hidden            NTFSFileAttributes = 2
	System            NTFSFileAttributes = 4
	Directory         NTFSFileAttributes = 16
	Archive           NTFSFileAttributes = 32
	None              NTFSFileAttributes = 128
	Temporary         NTFSFileAttributes = 256
	Offline           NTFSFileAttributes = 4096
	NotContentIndexed NTFSFileAttributes = 8192
	NoScrubData       NTFSFileAttributes = 131072
)

// PossibleNtfsFileAttributesValues returns the possible values for the NTFSFileAttributes const type.
func PossibleNtfsFileAttributesValues() []NTFSFileAttributes {
	return []NTFSFileAttributes{
		Readonly,
		Hidden,
		System,
		Directory,
		Archive,
		None,
		Temporary,
		Offline,
		NotContentIndexed,
		NoScrubData,
	}
}

// CopyStatusType defines the states of the copy operation.
type CopyStatusType = generated.CopyStatusType

const (
	CopyStatusTypePending CopyStatusType = generated.CopyStatusTypePending
	CopyStatusTypeSuccess CopyStatusType = generated.CopyStatusTypeSuccess
	CopyStatusTypeAborted CopyStatusType = generated.CopyStatusTypeAborted
	CopyStatusTypeFailed  CopyStatusType = generated.CopyStatusTypeFailed
)

// PossibleCopyStatusTypeValues returns the possible values for the CopyStatusType const type.
func PossibleCopyStatusTypeValues() []CopyStatusType {
	return generated.PossibleCopyStatusTypeValues()
}

// PermissionCopyModeType determines the copy behavior of the security descriptor of the file.
//   - source: The security descriptor on the destination file is copied from the source file.
//   - override: The security descriptor on the destination file is determined via the x-ms-file-permission or x-ms-file-permission-key header.
type PermissionCopyModeType = generated.PermissionCopyModeType

const (
	PermissionCopyModeTypeSource   PermissionCopyModeType = generated.PermissionCopyModeTypeSource
	PermissionCopyModeTypeOverride PermissionCopyModeType = generated.PermissionCopyModeTypeOverride
)

// PossiblePermissionCopyModeTypeValues returns the possible values for the PermissionCopyModeType const type.
func PossiblePermissionCopyModeTypeValues() []PermissionCopyModeType {
	return generated.PossiblePermissionCopyModeTypeValues()
}

// RangeWriteType represents one of the following options.
//   - update: Writes the bytes specified by the request body into the specified range. The Range and Content-Length headers must match to perform the update.
//   - clear: Clears the specified range and releases the space used in storage for that range. To clear a range, set the Content-Length header to zero,
//     and set the Range header to a value that indicates the range to clear, up to maximum file size.
type RangeWriteType = generated.FileRangeWriteType

const (
	RangeWriteTypeUpdate RangeWriteType = generated.FileRangeWriteTypeUpdate
	RangeWriteTypeClear  RangeWriteType = generated.FileRangeWriteTypeClear
)

// PossibleRangeWriteTypeValues returns the possible values for the RangeWriteType const type.
func PossibleRangeWriteTypeValues() []RangeWriteType {
	return generated.PossibleFileRangeWriteTypeValues()
}
