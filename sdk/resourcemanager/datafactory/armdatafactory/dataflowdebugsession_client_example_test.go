//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Create.json
func ExampleDataFlowDebugSessionClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataFlowDebugSessionClient().BeginCreate(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.CreateDataFlowDebugSessionRequest{
		IntegrationRuntime: &armdatafactory.IntegrationRuntimeDebugResource{
			Name: to.Ptr("ir1"),
			Properties: &armdatafactory.ManagedIntegrationRuntime{
				Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeManaged),
				TypeProperties: &armdatafactory.ManagedIntegrationRuntimeTypeProperties{
					ComputeProperties: &armdatafactory.IntegrationRuntimeComputeProperties{
						DataFlowProperties: &armdatafactory.IntegrationRuntimeDataFlowProperties{
							ComputeType: to.Ptr(armdatafactory.DataFlowComputeTypeGeneral),
							CoreCount:   to.Ptr[int32](48),
							TimeToLive:  to.Ptr[int32](10),
						},
						Location: to.Ptr("AutoResolve"),
					},
				},
			},
		},
		TimeToLive: to.Ptr[int32](60),
	}, nil)
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
	// res.CreateDataFlowDebugSessionResponse = armdatafactory.CreateDataFlowDebugSessionResponse{
	// 	SessionID: to.Ptr("229c688c-944c-44ac-b31a-82d50f347154"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_QueryByFactory.json
func ExampleDataFlowDebugSessionClient_NewQueryByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataFlowDebugSessionClient().NewQueryByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.QueryDataFlowDebugSessionsResponse = armdatafactory.QueryDataFlowDebugSessionsResponse{
		// 	Value: []*armdatafactory.DataFlowDebugSessionInfo{
		// 		{
		// 			AdditionalProperties: map[string]any{
		// 				"dataflowName": "DebugSession-0a7e0d6e-f2b7-48cc-8cd8-618326f5662f",
		// 				"userObjectId": "0a7e0d6e-f2b7-48cc-8cd8-618326f5662f",
		// 			},
		// 			ComputeType: to.Ptr("General"),
		// 			CoreCount: to.Ptr[int32](48),
		// 			LastActivityTime: to.Ptr("2019-09-05T18:28:00.9459674+00:00"),
		// 			SessionID: to.Ptr("229c688c-944c-44ac-b31a-82d50f347154"),
		// 			StartTime: to.Ptr("2019-09-05T18:23:20.3257799+00:00"),
		// 			TimeToLiveInMinutes: to.Ptr[int32](60),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_AddDataFlow.json
func ExampleDataFlowDebugSessionClient_AddDataFlow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataFlowDebugSessionClient().AddDataFlow(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.DataFlowDebugPackage{
		DataFlow: &armdatafactory.DataFlowDebugResource{
			Name: to.Ptr("dataflow1"),
			Properties: &armdatafactory.MappingDataFlow{
				Type: to.Ptr("MappingDataFlow"),
				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
					Script: to.Ptr("\n\nsource(output(\n		Column_1 as string\n	),\n	allowSchemaDrift: true,\n	validateSchema: false) ~> source1"),
					Sinks:  []*armdatafactory.DataFlowSink{},
					Sources: []*armdatafactory.DataFlowSource{
						{
							Name: to.Ptr("source1"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("DelimitedText2"),
							},
						}},
					Transformations: []*armdatafactory.Transformation{},
				},
			},
		},
		Datasets: []*armdatafactory.DatasetDebugResource{
			{
				Name: to.Ptr("dataset1"),
				Properties: &armdatafactory.DelimitedTextDataset{
					Type: to.Ptr("DelimitedText"),
					Schema: []any{
						map[string]any{
							"type": "String",
						},
					},
					Annotations: []any{},
					LinkedServiceName: &armdatafactory.LinkedServiceReference{
						Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
						ReferenceName: to.Ptr("linkedService5"),
					},
					TypeProperties: &armdatafactory.DelimitedTextDatasetTypeProperties{
						ColumnDelimiter:  ",",
						EscapeChar:       "\\",
						FirstRowAsHeader: true,
						Location: &armdatafactory.AzureBlobStorageLocation{
							Type:      to.Ptr("AzureBlobStorageLocation"),
							FileName:  "Ansiencoding.csv",
							Container: "dataflow-sample-data",
						},
						QuoteChar: "\"",
					},
				},
			}},
		DebugSettings: &armdatafactory.DataFlowDebugPackageDebugSettings{
			DatasetParameters: map[string]any{
				"Movies": map[string]any{
					"path": "abc",
				},
				"Output": map[string]any{
					"time": "def",
				},
			},
			Parameters: map[string]any{
				"sourcePath": "Toy",
			},
			SourceSettings: []*armdatafactory.DataFlowSourceSetting{
				{
					RowLimit:   to.Ptr[int32](1000),
					SourceName: to.Ptr("source1"),
				},
				{
					RowLimit:   to.Ptr[int32](222),
					SourceName: to.Ptr("source2"),
				}},
		},
		LinkedServices: []*armdatafactory.LinkedServiceDebugResource{
			{
				Name: to.Ptr("linkedService1"),
				Properties: &armdatafactory.AzureBlobStorageLinkedService{
					Type:        to.Ptr("AzureBlobStorage"),
					Annotations: []any{},
					TypeProperties: &armdatafactory.AzureBlobStorageLinkedServiceTypeProperties{
						ConnectionString:    "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;",
						EncryptedCredential: to.Ptr("<credential>"),
					},
				},
			}},
		SessionID: to.Ptr("f06ed247-9d07-49b2-b05e-2cb4a2fc871e"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AddDataFlowToDebugSessionResponse = armdatafactory.AddDataFlowToDebugSessionResponse{
	// 	JobVersion: to.Ptr("e5328ee7-c524-4207-8ba4-b709010db33d"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_Delete.json
func ExampleDataFlowDebugSessionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDataFlowDebugSessionClient().Delete(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.DeleteDataFlowDebugSessionRequest{
		SessionID: to.Ptr("91fb57e0-8292-47be-89ff-c8f2d2bb2a7e"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_ExecuteCommand.json
func ExampleDataFlowDebugSessionClient_BeginExecuteCommand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDataFlowDebugSessionClient().BeginExecuteCommand(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.DataFlowDebugCommandRequest{
		Command: to.Ptr(armdatafactory.DataFlowDebugCommandTypeExecutePreviewQuery),
		CommandPayload: &armdatafactory.DataFlowDebugCommandPayload{
			RowLimits:  to.Ptr[int32](100),
			StreamName: to.Ptr("source1"),
		},
		SessionID: to.Ptr("f06ed247-9d07-49b2-b05e-2cb4a2fc871e"),
	}, nil)
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
	// res.DataFlowDebugCommandResponse = armdatafactory.DataFlowDebugCommandResponse{
	// 	Data: to.Ptr("some output"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}
