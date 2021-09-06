//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewGalleryImageVersionsClient(con,
		"<subscription-id>")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		armcompute.GalleryImageVersion{
			Resource: armcompute.Resource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &armcompute.GalleryImageVersionProperties{
				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: armcompute.GalleryArtifactPublishingProfileBase{
						TargetRegions: []*armcompute.TargetRegion{
							{
								Name: to.StringPtr("<name>"),
								Encryption: &armcompute.EncryptionImages{
									DataDiskImages: []*armcompute.DataDiskImageEncryption{
										{
											DiskImageEncryption: armcompute.DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("<disk-encryption-set-id>"),
											},
											Lun: to.Int32Ptr(0),
										},
										{
											DiskImageEncryption: armcompute.DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("<disk-encryption-set-id>"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &armcompute.OSDiskImageEncryption{
										DiskImageEncryption: armcompute.DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("<disk-encryption-set-id>"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							{
								Name: to.StringPtr("<name>"),
								Encryption: &armcompute.EncryptionImages{
									DataDiskImages: []*armcompute.DataDiskImageEncryption{
										{
											DiskImageEncryption: armcompute.DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("<disk-encryption-set-id>"),
											},
											Lun: to.Int32Ptr(0),
										},
										{
											DiskImageEncryption: armcompute.DiskImageEncryption{
												DiskEncryptionSetID: to.StringPtr("<disk-encryption-set-id>"),
											},
											Lun: to.Int32Ptr(1),
										}},
									OSDiskImage: &armcompute.OSDiskImageEncryption{
										DiskImageEncryption: armcompute.DiskImageEncryption{
											DiskEncryptionSetID: to.StringPtr("<disk-encryption-set-id>"),
										},
									},
								},
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   armcompute.StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
					Source: &armcompute.GalleryArtifactVersionSource{
						ID: to.StringPtr("<id>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GalleryImageVersion.ID: %s\n", *res.ID)
}

func ExampleGalleryImageVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewGalleryImageVersionsClient(con,
		"<subscription-id>")
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		armcompute.GalleryImageVersionUpdate{
			Properties: &armcompute.GalleryImageVersionProperties{
				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
					GalleryArtifactPublishingProfileBase: armcompute.GalleryArtifactPublishingProfileBase{
						TargetRegions: []*armcompute.TargetRegion{
							{
								Name:                 to.StringPtr("<name>"),
								RegionalReplicaCount: to.Int32Ptr(1),
							},
							{
								Name:                 to.StringPtr("<name>"),
								RegionalReplicaCount: to.Int32Ptr(2),
								StorageAccountType:   armcompute.StorageAccountTypeStandardZRS.ToPtr(),
							}},
					},
				},
				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
					Source: &armcompute.GalleryArtifactVersionSource{
						ID: to.StringPtr("<id>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GalleryImageVersion.ID: %s\n", *res.ID)
}

func ExampleGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewGalleryImageVersionsClient(con,
		"<subscription-id>")
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		&armcompute.GalleryImageVersionsGetOptions{Expand: armcompute.ReplicationStatusTypesReplicationStatus.ToPtr()})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("GalleryImageVersion.ID: %s\n", *res.ID)
}

func ExampleGalleryImageVersionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewGalleryImageVersionsClient(con,
		"<subscription-id>")
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleGalleryImageVersionsClient_ListByGalleryImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := armcompute.NewGalleryImageVersionsClient(con,
		"<subscription-id>")
	pager := client.ListByGalleryImage("<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("GalleryImageVersion.ID: %s\n", *v.ID)
		}
	}
}
