//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package share

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azfile/internal/generated"
	"time"
)

// SharedKeyCredential contains an account's name and its primary or secondary key.
type SharedKeyCredential = exported.SharedKeyCredential

// ---------------------------------------------------------------------------------------------------------------------

// CreateOptions contains the optional parameters for the Client.Create method.
type CreateOptions struct {
	// Specifies the access tier of the share.
	AccessTier *AccessTier
	// Protocols to enable on the share.
	EnabledProtocols *string
	// A name-value pair to associate with a file storage object.
	Metadata map[string]*string
	// Specifies the maximum size of the share, in gigabytes.
	Quota *int32
	// Root squash to set on the share. Only valid for NFS shares.
	RootSquash *RootSquash
}

func (o *CreateOptions) format() *generated.ShareClientCreateOptions {
	if o == nil {
		return nil
	}

	return &generated.ShareClientCreateOptions{
		AccessTier:       o.AccessTier,
		EnabledProtocols: o.EnabledProtocols,
		Metadata:         o.Metadata,
		Quota:            o.Quota,
		RootSquash:       o.RootSquash,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// DeleteOptions contains the optional parameters for the Client.Delete method.
type DeleteOptions struct {
	// Specifies the option include to delete the base share and all of its snapshots.
	DeleteSnapshots *DeleteSnapshotsOptionType
	// TODO: Should snapshot be removed from the option bag
	// The snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot to query.
	ShareSnapshot *string
	// LeaseAccessConditions contains optional parameters to access leased entity.
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *DeleteOptions) format() (*generated.ShareClientDeleteOptions, *LeaseAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return &generated.ShareClientDeleteOptions{
		DeleteSnapshots: o.DeleteSnapshots,
		Sharesnapshot:   o.ShareSnapshot,
	}, o.LeaseAccessConditions
}

// LeaseAccessConditions contains optional parameters to access leased entity.
type LeaseAccessConditions = generated.LeaseAccessConditions

// ---------------------------------------------------------------------------------------------------------------------

// RestoreOptions contains the optional parameters for the Client.Restore method.
type RestoreOptions struct {
	// placeholder for future options
}

// ---------------------------------------------------------------------------------------------------------------------

// GetPropertiesOptions contains the optional parameters for the Client.GetProperties method.
type GetPropertiesOptions struct {
	// TODO: Should snapshot be removed from the option bag
	// The snapshot parameter is an opaque DateTime value that, when present, specifies the share snapshot to query.
	ShareSnapshot *string
	// LeaseAccessConditions contains optional parameters to access leased entity.
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *GetPropertiesOptions) format() (*generated.ShareClientGetPropertiesOptions, *LeaseAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return &generated.ShareClientGetPropertiesOptions{
		Sharesnapshot: o.ShareSnapshot,
	}, o.LeaseAccessConditions
}

// ---------------------------------------------------------------------------------------------------------------------

// SetPropertiesOptions contains the optional parameters for the Client.SetProperties method.
type SetPropertiesOptions struct {
	// Specifies the access tier of the share.
	AccessTier *AccessTier
	// Specifies the maximum size of the share, in gigabytes.
	Quota *int32
	// Root squash to set on the share. Only valid for NFS shares.
	RootSquash *RootSquash
	// LeaseAccessConditions contains optional parameters to access leased entity.
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *SetPropertiesOptions) format() (*generated.ShareClientSetPropertiesOptions, *LeaseAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return &generated.ShareClientSetPropertiesOptions{
		AccessTier: o.AccessTier,
		Quota:      o.Quota,
		RootSquash: o.RootSquash,
	}, o.LeaseAccessConditions
}

// ---------------------------------------------------------------------------------------------------------------------

// CreateSnapshotOptions contains the optional parameters for the Client.CreateSnapshot method.
type CreateSnapshotOptions struct {
	// A name-value pair to associate with a file storage object.
	Metadata map[string]*string
}

func (o *CreateSnapshotOptions) format() *generated.ShareClientCreateSnapshotOptions {
	if o == nil {
		return nil
	}

	return &generated.ShareClientCreateSnapshotOptions{
		Metadata: o.Metadata,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// GetAccessPolicyOptions contains the optional parameters for the Client.GetAccessPolicy method.
type GetAccessPolicyOptions struct {
	// LeaseAccessConditions contains optional parameters to access leased entity.
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *GetAccessPolicyOptions) format() (*generated.ShareClientGetAccessPolicyOptions, *LeaseAccessConditions) {
	if o == nil {
		return nil, nil
	}

	return nil, o.LeaseAccessConditions
}

// SignedIdentifier - Signed identifier.
type SignedIdentifier = generated.SignedIdentifier

// AccessPolicy - An Access policy.
type AccessPolicy = generated.AccessPolicy

// ---------------------------------------------------------------------------------------------------------------------

// SetAccessPolicyOptions contains the optional parameters for the Client.SetAccessPolicy method.
type SetAccessPolicyOptions struct {
	// Specifies the ACL for the share.
	ShareACL []*SignedIdentifier
	// LeaseAccessConditions contains optional parameters to access leased entity.
	LeaseAccessConditions *LeaseAccessConditions
}

func (o *SetAccessPolicyOptions) format() (*generated.ShareClientSetAccessPolicyOptions, []*SignedIdentifier, *LeaseAccessConditions, error) {
	if o == nil {
		return nil, nil, nil, nil
	}

	if o.ShareACL != nil {
		for _, si := range o.ShareACL {
			err := formatTime(si)
			if err != nil {
				return nil, nil, nil, err
			}
		}
	}

	return nil, o.ShareACL, o.LeaseAccessConditions, nil
}

func formatTime(si *SignedIdentifier) error {
	if si.AccessPolicy == nil {
		return nil
	}

	if si.AccessPolicy.Start != nil {
		st, err := time.Parse(time.RFC3339, si.AccessPolicy.Start.UTC().Format(time.RFC3339))
		if err != nil {
			return err
		}
		si.AccessPolicy.Start = &st
	}
	if si.AccessPolicy.Expiry != nil {
		et, err := time.Parse(time.RFC3339, si.AccessPolicy.Expiry.UTC().Format(time.RFC3339))
		if err != nil {
			return err
		}
		si.AccessPolicy.Expiry = &et
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// CreatePermissionOptions contains the optional parameters for the Client.CreatePermission method.
type CreatePermissionOptions struct {
	// placeholder for future options
}

// Permission - A permission (a security descriptor) at the share level.
type Permission = generated.SharePermission

// ---------------------------------------------------------------------------------------------------------------------

// GetPermissionOptions contains the optional parameters for the Client.GetPermission method.
type GetPermissionOptions struct {
	// placeholder for future options
}

// ---------------------------------------------------------------------------------------------------------------------

// SetMetadataOptions contains the optional parameters for the Client.SetMetadata method.
type SetMetadataOptions struct {
	// A name-value pair to associate with a file storage object.
	Metadata map[string]*string
	// LeaseAccessConditions contains optional parameters to access leased entity.
	LeaseAccessConditions *LeaseAccessConditions
}

// ---------------------------------------------------------------------------------------------------------------------

// GetStatisticsOptions contains the optional parameters for the Client.GetStatistics method.
type GetStatisticsOptions struct {
	// LeaseAccessConditions contains optional parameters to access leased entity.
	LeaseAccessConditions *LeaseAccessConditions
}

// Stats - Stats for the share.
type Stats = generated.ShareStats
