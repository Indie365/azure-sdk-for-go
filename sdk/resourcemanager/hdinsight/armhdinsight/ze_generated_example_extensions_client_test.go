//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/EnableLinuxClusterMonitoring.json
func ExampleExtensionsClient_BeginEnableMonitoring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginEnableMonitoring(ctx,
		"rg1",
		"cluster1",
		armhdinsight.ClusterMonitoringRequest{
			PrimaryKey:  to.Ptr("**********"),
			WorkspaceID: to.Ptr("a2090ead-8c9f-4fba-b70e-533e3e003163"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxClusterMonitoringStatus.json
func ExampleExtensionsClient_GetMonitoringStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetMonitoringStatus(ctx,
		"rg1",
		"cluster1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DisableLinuxClusterMonitoring.json
func ExampleExtensionsClient_BeginDisableMonitoring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDisableMonitoring(ctx,
		"rg1",
		"cluster1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/EnableLinuxClusterAzureMonitor.json
func ExampleExtensionsClient_BeginEnableAzureMonitor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginEnableAzureMonitor(ctx,
		"rg1",
		"cluster1",
		armhdinsight.AzureMonitorRequest{
			PrimaryKey:  to.Ptr("**********"),
			WorkspaceID: to.Ptr("a2090ead-8c9f-4fba-b70e-533e3e003163"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetLinuxClusterAzureMonitorStatus.json
func ExampleExtensionsClient_GetAzureMonitorStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAzureMonitorStatus(ctx,
		"rg1",
		"cluster1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DisableLinuxClusterAzureMonitor.json
func ExampleExtensionsClient_BeginDisableAzureMonitor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDisableAzureMonitor(ctx,
		"rg1",
		"cluster1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/CreateExtension.json
func ExampleExtensionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cluster1",
		"clustermonitoring",
		armhdinsight.Extension{
			PrimaryKey:  to.Ptr("**********"),
			WorkspaceID: to.Ptr("a2090ead-8c9f-4fba-b70e-533e3e003163"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetExtension.json
func ExampleExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"rg1",
		"cluster1",
		"clustermonitoring",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DeleteExtension.json
func ExampleExtensionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"rg1",
		"cluster1",
		"clustermonitoring",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetExtensionCreationAsyncOperationStatus.json
func ExampleExtensionsClient_GetAzureAsyncOperationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhdinsight.NewExtensionsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAzureAsyncOperationStatus(ctx,
		"rg1",
		"cluster1",
		"azuremonitor",
		"CF938302-6B4D-44A0-A6D2-C0D67E847AEC",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}