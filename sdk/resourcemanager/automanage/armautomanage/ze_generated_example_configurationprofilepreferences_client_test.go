//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"
)

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/createOrUpdateConfigurationProfilePreference.json
func ExampleConfigurationProfilePreferencesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfilePreferencesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<configuration-profile-preference-name>",
		"<resource-group-name>",
		armautomanage.ConfigurationProfilePreference{
			TrackedResource: armautomanage.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"Organization": to.StringPtr("Administration"),
				},
			},
			Properties: &armautomanage.ConfigurationProfilePreferenceProperties{
				AntiMalware: &armautomanage.ConfigurationProfilePreferenceAntiMalware{
					EnableRealTimeProtection: armautomanage.EnableRealTimeProtectionTrue.ToPtr(),
				},
				VMBackup: &armautomanage.ConfigurationProfilePreferenceVMBackup{
					TimeZone: to.StringPtr("<time-zone>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationProfilePreference.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/getConfigurationProfilePreference.json
func ExampleConfigurationProfilePreferencesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfilePreferencesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<configuration-profile-preference-name>",
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationProfilePreference.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/deleteConfigurationProfilePreference.json
func ExampleConfigurationProfilePreferencesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfilePreferencesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<configuration-profile-preference-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/updateConfigurationProfilePreference.json
func ExampleConfigurationProfilePreferencesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfilePreferencesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<configuration-profile-preference-name>",
		"<resource-group-name>",
		armautomanage.ConfigurationProfilePreferenceUpdate{
			UpdateResource: armautomanage.UpdateResource{
				Tags: map[string]*string{
					"Organization": to.StringPtr("Administration"),
				},
			},
			Properties: &armautomanage.ConfigurationProfilePreferenceProperties{
				AntiMalware: &armautomanage.ConfigurationProfilePreferenceAntiMalware{
					EnableRealTimeProtection: armautomanage.EnableRealTimeProtectionTrue.ToPtr(),
				},
				VMBackup: &armautomanage.ConfigurationProfilePreferenceVMBackup{
					TimeZone: to.StringPtr("<time-zone>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationProfilePreference.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/listConfigurationProfilePreferencesByResourceGroup.json
func ExampleConfigurationProfilePreferencesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfilePreferencesClient("<subscription-id>", cred, nil)
	_, err = client.ListByResourceGroup(ctx,
		"<resource-group-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/preview/2020-06-30-preview/examples/listConfigurationProfilePreferencesBySubscription.json
func ExampleConfigurationProfilePreferencesClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomanage.NewConfigurationProfilePreferencesClient("<subscription-id>", cred, nil)
	_, err = client.ListBySubscription(ctx,
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
