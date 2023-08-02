//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package service

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/filesystem"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/internal/generated_blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azdatalake/sas"
	"time"
)
import blobSAS "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"

// KeyInfo contains KeyInfo struct.
type KeyInfo = generated_blob.KeyInfo

type CreateFilesystemOptions = filesystem.CreateOptions

type DeleteFilesystemOptions = filesystem.DeleteOptions

// CORSRule - CORS is an HTTP feature that enables a web application running under one domain to access resources in another
// domain. Web browsers implement a security restriction known as same-origin policy that
// prevents a web page from calling APIs in a different domain; CORS provides a secure way to allow one domain (the origin
// domain) to call APIs in another domain.
type CORSRule = service.CORSRule

// StorageServiceProperties - Storage Service Properties.
type StorageServiceProperties = service.StorageServiceProperties

// RetentionPolicy - the retention policy which determines how long the associated data should persist.
type RetentionPolicy = service.RetentionPolicy

// Metrics - a summary of request statistics grouped by API in hour or minute aggregates
type Metrics = service.Metrics

// Logging - Azure Analytics Logging settings.
type Logging = service.Logging

// StaticWebsite - The properties that enable an account to host a static website.
type StaticWebsite = service.StaticWebsite

// SharedKeyCredential contains an account's name and its primary or secondary key.
type SharedKeyCredential = exported.SharedKeyCredential

// GetUserDelegationCredentialOptions contains optional parameters for Service.GetUserDelegationKey method.
type GetUserDelegationCredentialOptions struct {
	// placeholder for future options
}

func (o *GetUserDelegationCredentialOptions) format() *generated_blob.ServiceClientGetUserDelegationKeyOptions {
	return nil
}

// UserDelegationCredential contains an account's name and its user delegation key.
type UserDelegationCredential = exported.UserDelegationCredential

// UserDelegationKey contains UserDelegationKey.
type UserDelegationKey = exported.UserDelegationKey

// GetPropertiesOptions contains the optional parameters for the Client.GetProperties method.
type GetPropertiesOptions struct {
	// placeholder for future options
}

func (o *GetPropertiesOptions) format() *service.GetPropertiesOptions {
	if o == nil {
		return nil
	}
	return &service.GetPropertiesOptions{}
}

// SetPropertiesOptions provides set of options for Client.SetProperties
type SetPropertiesOptions struct {
	// The set of CORS rules.
	CORS []*CORSRule

	// The default version to use for requests to the Datalake service if an incoming request's version is not specified. Possible
	// values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion *string

	// the retention policy which determines how long the associated data should persist.
	DeleteRetentionPolicy *RetentionPolicy

	// a summary of request statistics grouped by API in hour or minute aggregates
	// If version is not set - we default to "1.0"
	HourMetrics *Metrics

	// Azure Analytics Logging settings.
	// If version is not set - we default to "1.0"
	Logging *Logging

	// a summary of request statistics grouped by API in hour or minute aggregates
	// If version is not set - we default to "1.0"
	MinuteMetrics *Metrics

	// The properties that enable an account to host a static website.
	StaticWebsite *StaticWebsite
}

func (o *SetPropertiesOptions) format() *service.SetPropertiesOptions {
	if o == nil {
		return nil
	}
	return &service.SetPropertiesOptions{
		CORS:                  o.CORS,
		DefaultServiceVersion: o.DefaultServiceVersion,
		DeleteRetentionPolicy: o.DeleteRetentionPolicy,
		HourMetrics:           o.HourMetrics,
		Logging:               o.Logging,
		MinuteMetrics:         o.MinuteMetrics,
		StaticWebsite:         o.StaticWebsite,
	}
}

// ListFilesystemsInclude indicates what additional information the service should return with each filesystem.
type ListFilesystemsInclude struct {
	// Tells the service whether to return metadata for each filesystem.
	Metadata bool

	// Tells the service whether to return soft-deleted filesystems.
	Deleted bool

	System bool
}

// ListFilesystemsOptions contains the optional parameters for the Client.List method.
type ListFilesystemsOptions struct {
	Include    ListFilesystemsInclude
	Marker     *string
	MaxResults *int32
	Prefix     *string
}

// GetSASURLOptions contains the optional parameters for the Client.GetSASURL method.
type GetSASURLOptions struct {
	StartTime *time.Time
}

func (o *GetSASURLOptions) format(resources sas.AccountResourceTypes, permissions sas.AccountPermissions) (blobSAS.AccountResourceTypes, blobSAS.AccountPermissions, *service.GetSASURLOptions) {
	res := blobSAS.AccountResourceTypes{
		Service:   resources.Service,
		Container: resources.Container,
		Object:    resources.Object,
	}
	perms := blobSAS.AccountPermissions{
		Read:    permissions.Read,
		Write:   permissions.Write,
		Delete:  permissions.Delete,
		List:    permissions.List,
		Add:     permissions.Add,
		Create:  permissions.Create,
		Update:  permissions.Update,
		Process: permissions.Process,
	}
	if o == nil {
		return res, perms, nil
	}

	return res, perms, &service.GetSASURLOptions{
		StartTime: o.StartTime,
	}
}
