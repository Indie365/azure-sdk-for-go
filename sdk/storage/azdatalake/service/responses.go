//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package service

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/filesystem"
)

// CreateFilesystemResponse contains the response fields for the CreateFilesystem operation.
type CreateFilesystemResponse = filesystem.CreateResponse

// DeleteFilesystemResponse contains the response fields for the DeleteFilesystem operation.
type DeleteFilesystemResponse = filesystem.DeleteResponse

// SetPropertiesResponse contains the response fields for the SetProperties operation.
type SetPropertiesResponse = service.SetPropertiesResponse

// GetPropertiesResponse contains the response fields for the GetProperties operation.
type GetPropertiesResponse = service.GetPropertiesResponse

// ListFilesystemsResponse contains the response fields for the ListFilesystems operation.
type ListFilesystemsResponse = service.ListContainersResponse

// TODO: IN THIS RESPONSE THERE IS A CONTAINERITEM YOU NEED TO DESERIALIZE THIS MANUALLY
// ListContainersSegmentResponse - An enumeration of containers
type ListContainersSegmentResponse = service.ListContainersSegmentResponse

//type ListFilesystemsResponse struct {
//	ListFilesystemsSegmentResponse
//	// ClientRequestID contains the information returned from the x-ms-client-request-id header response.
//	ClientRequestID *string `xml:"ClientRequestID"`
//
//	// RequestID contains the information returned from the x-ms-request-id header response.
//	RequestID *string `xml:"RequestID"`
//
//	// Version contains the information returned from the x-ms-version header response.
//	Version *string `xml:"Version"`
//}
//
//// ListFilesystemsSegmentResponse - An enumeration of containers
//type ListFilesystemsSegmentResponse struct {
//	// REQUIRED
//	ContainerItems []*FilesystemItem `xml:"Containers>Container"`
//
//	// REQUIRED
//	ServiceEndpoint *string `xml:"ServiceEndpoint,attr"`
//	Marker          *string `xml:"Marker"`
//	MaxResults      *int32  `xml:"MaxResults"`
//	NextMarker      *string `xml:"NextMarker"`
//	Prefix          *string `xml:"Prefix"`
//}
//
//// FilesystemItem - An Azure Storage filesystem
//type FilesystemItem struct {
//	// REQUIRED
//	Name *string `xml:"Name"`
//
//	// REQUIRED; Properties of a container
//	Properties *FilesystemProperties `xml:"Properties"`
//	Deleted    *bool                 `xml:"Deleted"`
//
//	// Dictionary of
//	Metadata map[string]*string `xml:"Metadata"`
//	Version  *string            `xml:"Version"`
//}
//
//// FilesystemProperties - Properties of a filesystem
//type FilesystemProperties struct {
//	// REQUIRED
//	ETag *azcore.ETag `xml:"Etag"`
//
//	// REQUIRED
//	LastModified           *time.Time `xml:"Last-Modified"`
//	DefaultEncryptionScope *string    `xml:"DefaultEncryptionScope"`
//	DeletedTime            *time.Time `xml:"DeletedTime"`
//	HasImmutabilityPolicy  *bool      `xml:"HasImmutabilityPolicy"`
//	HasLegalHold           *bool      `xml:"HasLegalHold"`
//
//	// Indicates if version level worm is enabled on this container.
//	IsImmutableStorageWithVersioningEnabled *bool              `xml:"ImmutableStorageWithVersioningEnabled"`
//	LeaseDuration                           *LeaseDurationType `xml:"LeaseDuration"`
//	LeaseState                              *LeaseStateType    `xml:"LeaseState"`
//	LeaseStatus                             *LeaseStatusType   `xml:"LeaseStatus"`
//	PreventEncryptionScopeOverride          *bool              `xml:"DenyEncryptionScopeOverride"`
//	PublicAccess                            *PublicAccessType  `xml:"PublicAccess"`
//	RemainingRetentionDays                  *int32             `xml:"RemainingRetentionDays"`
//}
//
//type LeaseDurationType = lease.DurationType
//
//type LeaseStateType = lease.StateType
//
//type LeaseStatusType = lease.StatusType
//
//type PublicAccessType = filesystem.PublicAccessType
