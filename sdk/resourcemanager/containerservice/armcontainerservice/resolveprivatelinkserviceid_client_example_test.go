//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c183bb012de8e9e1d0d2e67a0994748df4747d2c/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-04-02-preview/examples/ResolvePrivateLinkServiceId.json
func ExampleResolvePrivateLinkServiceIDClient_POST() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResolvePrivateLinkServiceIDClient().POST(ctx, "rg1", "clustername1", armcontainerservice.PrivateLinkResource{
		Name: to.Ptr("management"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armcontainerservice.PrivateLinkResource{
	// 	PrivateLinkServiceID: to.Ptr("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Network/privateLinkServices/plsName"),
	// }
}
