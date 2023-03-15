//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package backup

import "time"

// BeginFullBackupOptions contains the optional parameters for the Client.BeginFullBackup method.
type BeginFullBackupOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BeginFullRestoreOptions contains the optional parameters for the Client.BeginFullRestore method.
type BeginFullRestoreOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BeginSelectiveKeyRestoreOptions contains the optional parameters for the Client.BeginSelectiveKeyRestore method.
type BeginSelectiveKeyRestoreOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FullBackupOperation - Full backup operation
type FullBackupOperation struct {
	// The Azure blob storage container Uri which contains the full backup
	AzureStorageBlobContainerURI *string `json:"azureStorageBlobContainerUri,omitempty"`

	// The end time of the backup operation in UTC
	EndTime *time.Time `json:"endTime,omitempty"`

	// Error encountered, if any, during the full backup operation.
	Error *ServerError `json:"error,omitempty"`

	// Identifier for the full backup operation.
	JobID *string `json:"jobId,omitempty"`

	// The start time of the backup operation in UTC
	StartTime *time.Time `json:"startTime,omitempty"`

	// Status of the backup operation.
	Status *string `json:"status,omitempty"`

	// The status details of backup operation.
	StatusDetails *string `json:"statusDetails,omitempty"`
}

// RestoreOperation - Restore operation
type RestoreOperation struct {
	// The end time of the restore operation
	EndTime *time.Time `json:"endTime,omitempty"`

	// Error encountered, if any, during the restore operation.
	Error *ServerError `json:"error,omitempty"`

	// Identifier for the restore operation.
	JobID *string `json:"jobId,omitempty"`

	// The start time of the restore operation
	StartTime *time.Time `json:"startTime,omitempty"`

	// Status of the restore operation.
	Status *string `json:"status,omitempty"`

	// The status details of restore operation.
	StatusDetails *string `json:"statusDetails,omitempty"`
}

// RestoreOperationParameters - Parameters for the restore operation
type RestoreOperationParameters struct {
	// REQUIRED; The Folder name of the blob where the previous successful full backup was stored
	FolderToRestore *string `json:"folderToRestore,omitempty"`

	// REQUIRED; Contains the information required to access blob storage.
	SasTokenParameters *SASTokenParameter `json:"sasTokenParameters,omitempty"`
}

// SASTokenParameter - Contains the information required to access blob storage.
type SASTokenParameter struct {
	// REQUIRED; Azure Blob storage container Uri
	StorageResourceURI *string `json:"storageResourceUri,omitempty"`

	// REQUIRED; The SAS token pointing to an Azure Blob storage container
	Token *string `json:"token,omitempty"`
}

// SelectiveKeyRestoreOperation - Selective Key Restore operation
type SelectiveKeyRestoreOperation struct {
	// The end time of the restore operation
	EndTime *time.Time `json:"endTime,omitempty"`

	// Error encountered, if any, during the selective key restore operation.
	Error *ServerError `json:"error,omitempty"`

	// Identifier for the selective key restore operation.
	JobID *string `json:"jobId,omitempty"`

	// The start time of the restore operation
	StartTime *time.Time `json:"startTime,omitempty"`

	// Status of the restore operation.
	Status *string `json:"status,omitempty"`

	// The status details of restore operation.
	StatusDetails *string `json:"statusDetails,omitempty"`
}

// SelectiveKeyRestoreOperationParameters - Parameters for the selective restore operation
type SelectiveKeyRestoreOperationParameters struct {
	// REQUIRED; The Folder name of the blob where the previous successful full backup was stored
	Folder *string `json:"folder,omitempty"`

	// REQUIRED; Contains the information required to access blob storage.
	SasTokenParameters *SASTokenParameter `json:"sasTokenParameters,omitempty"`
}
