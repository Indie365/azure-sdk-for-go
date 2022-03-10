//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/EnableLinuxClusterMonitoring.json
func ExampleExtensionsClient_BeginEnableMonitoring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginEnableMonitoring(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armhdinsight.ClusterMonitoringRequest{
			PrimaryKey:  to.StringPtr("<primary-key>"),
			WorkspaceID: to.StringPtr("<workspace-id>"),
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

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxClusterMonitoringStatus.json
func ExampleExtensionsClient_GetMonitoringStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	res, err := client.GetMonitoringStatus(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExtensionsClientGetMonitoringStatusResult)
}

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DisableLinuxClusterMonitoring.json
func ExampleExtensionsClient_BeginDisableMonitoring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDisableMonitoring(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/EnableLinuxClusterAzureMonitor.json
func ExampleExtensionsClient_BeginEnableAzureMonitor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginEnableAzureMonitor(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armhdinsight.AzureMonitorRequest{
			PrimaryKey:  to.StringPtr("<primary-key>"),
			WorkspaceID: to.StringPtr("<workspace-id>"),
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

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxClusterAzureMonitorStatus.json
func ExampleExtensionsClient_GetAzureMonitorStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	res, err := client.GetAzureMonitorStatus(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExtensionsClientGetAzureMonitorStatusResult)
}

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DisableLinuxClusterAzureMonitor.json
func ExampleExtensionsClient_BeginDisableAzureMonitor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDisableAzureMonitor(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateExtension.json
func ExampleExtensionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<extension-name>",
		armhdinsight.Extension{
			PrimaryKey:  to.StringPtr("<primary-key>"),
			WorkspaceID: to.StringPtr("<workspace-id>"),
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

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetExtension.json
func ExampleExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<extension-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExtensionsClientGetResult)
}

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DeleteExtension.json
func ExampleExtensionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<extension-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetExtensionCreationAsyncOperationStatus.json
func ExampleExtensionsClient_GetAzureAsyncOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhdinsight.NewExtensionsClient("<subscription-id>", cred, nil)
	res, err := client.GetAzureAsyncOperationStatus(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<extension-name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ExtensionsClientGetAzureAsyncOperationStatusResult)
}
