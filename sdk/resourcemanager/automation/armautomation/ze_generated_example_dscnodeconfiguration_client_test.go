//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteDscNodeConfiguration.json
func ExampleDscNodeConfigurationClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewDscNodeConfigurationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<node-configuration-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getDscNodeConfiguration.json
func ExampleDscNodeConfigurationClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewDscNodeConfigurationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<node-configuration-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createOrUpdateDscNodeConfiguration.json
func ExampleDscNodeConfigurationClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewDscNodeConfigurationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<node-configuration-name>",
		armautomation.DscNodeConfigurationCreateOrUpdateParameters{
			Name: to.Ptr("<name>"),
			Properties: &armautomation.DscNodeConfigurationCreateOrUpdateParametersProperties{
				Configuration: &armautomation.DscConfigurationAssociationProperty{
					Name: to.Ptr("<name>"),
				},
				IncrementNodeConfigurationBuild: to.Ptr(true),
				Source: &armautomation.ContentSource{
					Type: to.Ptr(armautomation.ContentSourceTypeEmbeddedContent),
					Hash: &armautomation.ContentHash{
						Algorithm: to.Ptr("<algorithm>"),
						Value:     to.Ptr("<value>"),
					},
					Value:   to.Ptr("<value>"),
					Version: to.Ptr("<version>"),
				},
			},
		},
		&armautomation.DscNodeConfigurationClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listDscNodeConfigurations.json
func ExampleDscNodeConfigurationClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armautomation.NewDscNodeConfigurationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByAutomationAccountPager("<resource-group-name>",
		"<automation-account-name>",
		&armautomation.DscNodeConfigurationClientListByAutomationAccountOptions{Filter: nil,
			Skip:        nil,
			Top:         nil,
			Inlinecount: nil,
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
