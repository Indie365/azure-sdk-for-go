//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDeviceGetBySubscription.json
func ExampleDevicesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(&armdataboxedge.DevicesClientListBySubscriptionOptions{Expand: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDeviceGetByResourceGroup.json
func ExampleDevicesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armdataboxedge.DevicesClientListByResourceGroupOptions{Expand: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDeviceGetByName.json
func ExampleDevicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientGetResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDevicePut.json
func ExampleDevicesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.Device{
			Location: to.StringPtr("<location>"),
			SKU: &armdataboxedge.SKUInfo{
				Name: armdataboxedge.SKUName("Edge").ToPtr(),
				Tier: armdataboxedge.SKUTier("Standard").ToPtr(),
			},
			Tags: map[string]*string{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDeviceDelete.json
func ExampleDevicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDevicePatch.json
func ExampleDevicesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.DevicePatch{
			Properties: &armdataboxedge.DevicePropertiesPatch{
				EdgeProfile: &armdataboxedge.EdgeProfilePatch{
					Subscription: &armdataboxedge.EdgeProfileSubscriptionPatch{
						ID: to.StringPtr("<id>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientUpdateResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DownloadUpdatesPost.json
func ExampleDevicesClient_BeginDownloadUpdates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDownloadUpdates(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/GenerateCertificate.json
func ExampleDevicesClient_GenerateCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.GenerateCertificate(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientGenerateCertificateResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/ExtendedInfoPost.json
func ExampleDevicesClient_GetExtendedInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.GetExtendedInformation(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientGetExtendedInformationResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/InstallUpdatesPost.json
func ExampleDevicesClient_BeginInstallUpdates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginInstallUpdates(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/NetworkSettingsGet.json
func ExampleDevicesClient_GetNetworkSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.GetNetworkSettings(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientGetNetworkSettingsResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/ScanForUpdatesPost.json
func ExampleDevicesClient_BeginScanForUpdates() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginScanForUpdates(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/SecuritySettingsUpdatePost.json
func ExampleDevicesClient_BeginCreateOrUpdateSecuritySettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdateSecuritySettings(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.SecuritySettings{
			Properties: &armdataboxedge.SecuritySettingsProperties{
				DeviceAdminPassword: &armdataboxedge.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:      armdataboxedge.EncryptionAlgorithm("AES256").ToPtr(),
					EncryptionCertThumbprint: to.StringPtr("<encryption-cert-thumbprint>"),
					Value:                    to.StringPtr("<value>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/GetUpdateExtendedInfo.json
func ExampleDevicesClient_UpdateExtendedInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.UpdateExtendedInformation(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.DeviceExtendedInfoPatch{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientUpdateExtendedInformationResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/UpdateSummaryGet.json
func ExampleDevicesClient_GetUpdateSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.GetUpdateSummary(ctx,
		"<device-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientGetUpdateSummaryResult)
}

// x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/UploadCertificatePost.json
func ExampleDevicesClient_UploadCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataboxedge.NewDevicesClient("<subscription-id>", cred, nil)
	res, err := client.UploadCertificate(ctx,
		"<device-name>",
		"<resource-group-name>",
		armdataboxedge.UploadCertificateRequest{
			Properties: &armdataboxedge.RawCertificateData{
				Certificate: to.StringPtr("<certificate>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DevicesClientUploadCertificateResult)
}
