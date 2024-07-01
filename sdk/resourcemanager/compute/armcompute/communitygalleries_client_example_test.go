//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/communityGalleryExamples/CommunityGallery_Get.json
func ExampleCommunityGalleriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunityGalleriesClient().Get(ctx, "myLocation", "publicGalleryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommunityGallery = armcompute.CommunityGallery{
	// 	Name: to.Ptr("publicGalleryName"),
	// 	Type: to.Ptr("Microsoft.Compute/locations/communityGallery"),
	// 	Identifier: &armcompute.CommunityGalleryIdentifier{
	// 		UniqueID: to.Ptr("/CommunityGalleries/publicGalleryName"),
	// 	},
	// 	Location: to.Ptr("myLocation"),
	// 	Properties: &armcompute.CommunityGalleryProperties{
	// 		ArtifactTags: map[string]*string{
	// 			"ShareTag-CommunityGallery": to.Ptr("CommunityGallery"),
	// 		},
	// 		CommunityMetadata: &armcompute.CommunityGalleryMetadata{
	// 			Eula: to.Ptr("https://test-uri.com"),
	// 			PrivacyStatementURI: to.Ptr("https://test-uri.com"),
	// 			PublicNames: []*string{
	// 				to.Ptr("prefix-xxxxx")},
	// 				PublisherContact: to.Ptr("sameple@email.com"),
	// 				PublisherURI: to.Ptr("https://test-uri.com"),
	// 			},
	// 			Disclaimer: to.Ptr("https://test-uri.com"),
	// 		},
	// 	}
}
