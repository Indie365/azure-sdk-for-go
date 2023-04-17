//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Update.json
func ExampleIntegrationRuntimesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().Update(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", armsynapse.UpdateIntegrationRuntimeRequest{
		AutoUpdate:        to.Ptr(armsynapse.IntegrationRuntimeAutoUpdateOff),
		UpdateDelayOffset: to.Ptr("\"PT3H\""),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeResource = armsynapse.IntegrationRuntimeResource{
	// 	Name: to.Ptr("exampleIntegrationRuntime"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/integrationruntimes"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Synapse/workspaces/exampleWorkspaceName/integrationruntimes/exampleIntegrationRuntime"),
	// 	Etag: to.Ptr("0400f1a1-0000-0000-0000-5b2188640000"),
	// 	Properties: &armsynapse.SelfHostedIntegrationRuntime{
	// 		Type: to.Ptr(armsynapse.IntegrationRuntimeTypeSelfHosted),
	// 		Description: to.Ptr("A selfhosted integration runtime"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Get.json
func ExampleIntegrationRuntimesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().Get(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", &armsynapse.IntegrationRuntimesClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeResource = armsynapse.IntegrationRuntimeResource{
	// 	Name: to.Ptr("exampleIntegrationRuntime"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/integrationruntimes"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Synapse/workspaces/exampleWorkspaceName/integrationruntimes/exampleIntegrationRuntime"),
	// 	Etag: to.Ptr("15003c4f-0000-0200-0000-5cbe090b0000"),
	// 	Properties: &armsynapse.SelfHostedIntegrationRuntime{
	// 		Type: to.Ptr(armsynapse.IntegrationRuntimeTypeSelfHosted),
	// 		Description: to.Ptr("A selfhosted integration runtime"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Create.json
func ExampleIntegrationRuntimesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginCreate(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", armsynapse.IntegrationRuntimeResource{
		Properties: &armsynapse.SelfHostedIntegrationRuntime{
			Type:        to.Ptr(armsynapse.IntegrationRuntimeTypeSelfHosted),
			Description: to.Ptr("A selfhosted integration runtime"),
		},
	}, &armsynapse.IntegrationRuntimesClientBeginCreateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeResource = armsynapse.IntegrationRuntimeResource{
	// 	Name: to.Ptr("exampleIntegrationRuntime"),
	// 	Type: to.Ptr("Microsoft.Synapse/workspaces/integrationruntimes"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Synapse/workspaces/exampleWorkspaceName/integrationruntimes/exampleIntegrationRuntime"),
	// 	Etag: to.Ptr("000046c4-0000-0000-0000-5b2198bf0000"),
	// 	Properties: &armsynapse.SelfHostedIntegrationRuntime{
	// 		Type: to.Ptr(armsynapse.IntegrationRuntimeTypeSelfHosted),
	// 		Description: to.Ptr("A selfhosted integration runtime"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Delete.json
func ExampleIntegrationRuntimesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginDelete(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Upgrade.json
func ExampleIntegrationRuntimesClient_Upgrade() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIntegrationRuntimesClient().Upgrade(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_ListByWorkspace.json
func ExampleIntegrationRuntimesClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationRuntimesClient().NewListByWorkspacePager("exampleResourceGroup", "exampleWorkspace", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IntegrationRuntimeListResponse = armsynapse.IntegrationRuntimeListResponse{
		// 	Value: []*armsynapse.IntegrationRuntimeResource{
		// 		{
		// 			Name: to.Ptr("exampleIntegrationRuntime"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/integrationruntimes"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Synapse/workspaces/exampleWorkspaceName/integrationruntimes/exampleIntegrationRuntime"),
		// 			Etag: to.Ptr("0400f1a1-0000-0000-0000-5b2188640000"),
		// 			Properties: &armsynapse.SelfHostedIntegrationRuntime{
		// 				Type: to.Ptr(armsynapse.IntegrationRuntimeTypeSelfHosted),
		// 				Description: to.Ptr("A selfhosted integration runtime"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Start.json
func ExampleIntegrationRuntimesClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginStart(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleManagedIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeStatusResponse = armsynapse.IntegrationRuntimeStatusResponse{
	// 	Name: to.Ptr("exampleManagedIntegrationRuntime"),
	// 	Properties: &armsynapse.ManagedIntegrationRuntimeStatus{
	// 		Type: to.Ptr(armsynapse.IntegrationRuntimeTypeManaged),
	// 		DataFactoryName: to.Ptr("exampleWorkspaceName"),
	// 		State: to.Ptr(armsynapse.IntegrationRuntimeStateStarted),
	// 		TypeProperties: &armsynapse.ManagedIntegrationRuntimeStatusTypeProperties{
	// 			CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T21:11:01.8695494Z"); return t}()),
	// 			Nodes: []*armsynapse.ManagedIntegrationRuntimeNode{
	// 			},
	// 			OtherErrors: []*armsynapse.ManagedIntegrationRuntimeError{
	// 			},
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Stop.json
func ExampleIntegrationRuntimesClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginStop(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleManagedIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
func ExampleIntegrationRuntimesClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().ListOutboundNetworkDependenciesEndpoints(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse = armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointsResponse{
	// 	Value: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesCategoryEndpoint{
	// 		{
	// 			Category: to.Ptr("Azure Data Factory (Management)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("wu.frontend.int.clouddatahub-int.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Azure Storage (Management)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("*.blob.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("*.table.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Event Hub (Logging)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("*.servicebus.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 		},
	// 		{
	// 			Category: to.Ptr("Microsoft Logging service (Internal Use)"),
	// 			Endpoints: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpoint{
	// 				{
	// 					DomainName: to.Ptr("gcs.prod.monitoring.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("prod.warmpath.msftcloudes.com"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 				},
	// 				{
	// 					DomainName: to.Ptr("azurewatsonanalysis-prod.core.windows.net"),
	// 					EndpointDetails: []*armsynapse.IntegrationRuntimeOutboundNetworkDependenciesEndpointDetails{
	// 						{
	// 							Port: to.Ptr[int32](443),
	// 					}},
	// 			}},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_EnableInteractiveQuery.json
func ExampleIntegrationRuntimesClient_BeginEnableInteractiveQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginEnableInteractiveQuery(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleManagedIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_DisableInteractiveQuery.json
func ExampleIntegrationRuntimesClient_BeginDisableInteractiveQuery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIntegrationRuntimesClient().BeginDisableInteractiveQuery(ctx, "exampleResourceGroup", "exampleWorkspace", "exampleManagedIntegrationRuntime", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
