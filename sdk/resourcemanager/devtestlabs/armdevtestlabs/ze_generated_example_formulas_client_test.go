//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_List.json
func ExampleFormulasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewFormulasClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager("<resource-group-name>",
		"<lab-name>",
		&armdevtestlabs.FormulasClientListOptions{Expand: nil,
			Filter:  nil,
			Top:     nil,
			Orderby: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_Get.json
func ExampleFormulasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewFormulasClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		&armdevtestlabs.FormulasClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_CreateOrUpdate.json
func ExampleFormulasClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewFormulasClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.Formula{
			Location: to.Ptr("<location>"),
			Properties: &armdevtestlabs.FormulaProperties{
				Description: to.Ptr("<description>"),
				FormulaContent: &armdevtestlabs.LabVirtualMachineCreationParameter{
					Location: to.Ptr("<location>"),
					Properties: &armdevtestlabs.LabVirtualMachineCreationParameterProperties{
						AllowClaim: to.Ptr(false),
						Artifacts: []*armdevtestlabs.ArtifactInstallProperties{
							{
								ArtifactID: to.Ptr("<artifact-id>"),
								Parameters: []*armdevtestlabs.ArtifactParameterProperties{},
							}},
						DisallowPublicIPAddress: to.Ptr(true),
						GalleryImageReference: &armdevtestlabs.GalleryImageReference{
							Offer:     to.Ptr("<offer>"),
							OSType:    to.Ptr("<ostype>"),
							Publisher: to.Ptr("<publisher>"),
							SKU:       to.Ptr("<sku>"),
							Version:   to.Ptr("<version>"),
						},
						IsAuthenticationWithSSHKey: to.Ptr(false),
						LabSubnetName:              to.Ptr("<lab-subnet-name>"),
						LabVirtualNetworkID:        to.Ptr("<lab-virtual-network-id>"),
						NetworkInterface: &armdevtestlabs.NetworkInterfaceProperties{
							SharedPublicIPAddressConfiguration: &armdevtestlabs.SharedPublicIPAddressConfiguration{
								InboundNatRules: []*armdevtestlabs.InboundNatRule{
									{
										BackendPort:       to.Ptr[int32](22),
										TransportProtocol: to.Ptr(armdevtestlabs.TransportProtocolTCP),
									}},
							},
						},
						Notes:       to.Ptr("<notes>"),
						Size:        to.Ptr("<size>"),
						StorageType: to.Ptr("<storage-type>"),
						UserName:    to.Ptr("<user-name>"),
					},
				},
			},
		},
		&armdevtestlabs.FormulasClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_Delete.json
func ExampleFormulasClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewFormulasClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Formulas_Update.json
func ExampleFormulasClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewFormulasClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<lab-name>",
		"<name>",
		armdevtestlabs.FormulaFragment{
			Tags: map[string]*string{
				"tagName1": to.Ptr("tagValue1"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
