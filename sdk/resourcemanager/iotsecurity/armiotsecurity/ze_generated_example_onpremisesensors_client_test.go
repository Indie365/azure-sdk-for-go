//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotsecurity/armiotsecurity"
)

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/OnPremiseSensors/List.json
func ExampleOnPremiseSensorsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewOnPremiseSensorsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OnPremiseSensorsClientListResult)
}

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/OnPremiseSensors/Get.json
func ExampleOnPremiseSensorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewOnPremiseSensorsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<on-premise-sensor-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OnPremiseSensorsClientGetResult)
}

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/OnPremiseSensors/Put.json
func ExampleOnPremiseSensorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewOnPremiseSensorsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<on-premise-sensor-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OnPremiseSensorsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/OnPremiseSensors/Delete.json
func ExampleOnPremiseSensorsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewOnPremiseSensorsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<on-premise-sensor-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/OnPremiseSensors/DownloadActivation.json
func ExampleOnPremiseSensorsClient_DownloadActivation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewOnPremiseSensorsClient("<subscription-id>", cred, nil)
	_, err = client.DownloadActivation(ctx,
		"<on-premise-sensor-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/iotsecurity/resource-manager/Microsoft.IoTSecurity/preview/2021-02-01-preview/examples/OnPremiseSensors/DownloadResetPassword.json
func ExampleOnPremiseSensorsClient_DownloadResetPassword() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiotsecurity.NewOnPremiseSensorsClient("<subscription-id>", cred, nil)
	_, err = client.DownloadResetPassword(ctx,
		"<on-premise-sensor-name>",
		armiotsecurity.ResetPasswordInput{
			ApplianceID: to.StringPtr("<appliance-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
