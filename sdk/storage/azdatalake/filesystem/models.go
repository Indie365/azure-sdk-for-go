//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package filesystem

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake"
)

// SetAccessPolicyOptions provides set of configurations for Filesystem.SetAccessPolicy operation.
type SetAccessPolicyOptions struct {
	// Specifies whether data in the filesystem may be accessed publicly and the level of access.
	// If this header is not included in the request, filesystem data is private to the account owner.
	Access           *PublicAccessType
	AccessConditions *AccessConditions
	FilesystemACL    []*SignedIdentifier
}

func (o *SetAccessPolicyOptions) format() *container.SetAccessPolicyOptions {
	return &container.SetAccessPolicyOptions{
		Access:           o.Access,
		AccessConditions: azdatalake.FormatContainerAccessConditions(o.AccessConditions),
		ContainerACL:     o.FilesystemACL,
	}
}

// CreateOptions contains the optional parameters for the Client.Create method.
type CreateOptions struct {
	// Specifies whether data in the filesystem may be accessed publicly and the level of access.
	Access *PublicAccessType

	// Optional. Specifies a user-defined name-value pair associated with the filesystem.
	Metadata map[string]*string

	// Optional. Specifies the encryption scope settings to set on the filesystem.
	CPKScopeInfo *CPKScopeInfo
}

func (o *CreateOptions) format() *container.CreateOptions {
	return &container.CreateOptions{
		Access:       o.Access,
		Metadata:     o.Metadata,
		CPKScopeInfo: o.CPKScopeInfo,
	}
}

// DeleteOptions contains the optional parameters for the Client.Delete method.
type DeleteOptions struct {
	AccessConditions *AccessConditions
}

func (o *DeleteOptions) format() *container.DeleteOptions {
	return &container.DeleteOptions{
		AccessConditions: azdatalake.FormatContainerAccessConditions(o.AccessConditions),
	}
}

// GetPropertiesOptions contains the optional parameters for the FilesystemClient.GetProperties method.
type GetPropertiesOptions struct {
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *GetPropertiesOptions) format() *container.GetPropertiesOptions {
	return &container.GetPropertiesOptions{
		LeaseAccessConditions: &container.LeaseAccessConditions{
			LeaseID: o.LeaseAccessConditions.LeaseID,
		},
	}
}

// SetMetadataOptions contains the optional parameters for the Client.SetMetadata method.
type SetMetadataOptions struct {
	Metadata                 map[string]*string
	LeaseAccessConditions    *LeaseAccessConditions
	ModifiedAccessConditions *ModifiedAccessConditions
}

func (o *SetMetadataOptions) format() *container.SetMetadataOptions {
	return &container.SetMetadataOptions{
		Metadata: o.Metadata,
		LeaseAccessConditions: &container.LeaseAccessConditions{
			LeaseID: o.LeaseAccessConditions.LeaseID,
		},
		ModifiedAccessConditions: &container.ModifiedAccessConditions{
			IfMatch:           o.ModifiedAccessConditions.IfMatch,
			IfNoneMatch:       o.ModifiedAccessConditions.IfNoneMatch,
			IfModifiedSince:   o.ModifiedAccessConditions.IfModifiedSince,
			IfUnmodifiedSince: o.ModifiedAccessConditions.IfUnmodifiedSince,
		},
	}
}

// GetAccessPolicyOptions contains the optional parameters for the Client.GetAccessPolicy method.
type GetAccessPolicyOptions struct {
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *GetAccessPolicyOptions) format() *container.GetAccessPolicyOptions {
	return &container.GetAccessPolicyOptions{
		LeaseAccessConditions: &container.LeaseAccessConditions{
			LeaseID: o.LeaseAccessConditions.LeaseID,
		},
	}
}

// CPKScopeInfo contains a group of parameters for the FilesystemClient.Create method.
type CPKScopeInfo = container.CPKScopeInfo

// AccessConditions identifies filesystem-specific access conditions which you optionally set.
type AccessConditions = azdatalake.FilesystemAccessConditions

// LeaseAccessConditions contains optional parameters to access leased entity.
type LeaseAccessConditions = azdatalake.LeaseAccessConditions

// ModifiedAccessConditions contains a group of parameters for specifying access conditions.
type ModifiedAccessConditions = azdatalake.ModifiedAccessConditions

// AccessPolicy - An Access policy.
type AccessPolicy = container.AccessPolicy

// SignedIdentifier - signed identifier.
type SignedIdentifier = container.SignedIdentifier

// ListPathsOptions contains the optional parameters for the Filesystem.ListPaths operation.
type ListPathsOptions struct {
	Marker     *string
	MaxResults *int32
	Prefix     *string
	Upn        *bool
}

// ListDeletedPathsOptions contains the optional parameters for the Filesystem.ListDeletedPaths operation. PLACEHOLDER
type ListDeletedPathsOptions struct {
	Marker     *string
	MaxResults *int32
	Prefix     *string
	Upn        *bool
}

// UndeletePathOptions contains the optional parameters for the Filesystem.UndeletePath operation.
type UndeletePathOptions struct {
	// placeholder
}
