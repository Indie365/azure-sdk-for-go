//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/TriggerGetAllInDevice.json
func ExampleTriggersClient_NewListByDataBoxEdgeDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTriggersClient().NewListByDataBoxEdgeDevicePager("testedgedevice", "GroupForEdgeAutomation", &armdataboxedge.TriggersClientListByDataBoxEdgeDeviceOptions{Filter: nil})
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
		// page.TriggerList = armdataboxedge.TriggerList{
		// 	Value: []armdataboxedge.TriggerClassification{
		// 		&armdataboxedge.FileEventTrigger{
		// 			Name: to.Ptr("trigger1"),
		// 			Type: to.Ptr("dataBoxEdgeDevices/triggers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/triggers/trigger1"),
		// 			Kind: to.Ptr(armdataboxedge.TriggerEventTypeFileEvent),
		// 			Properties: &armdataboxedge.FileTriggerProperties{
		// 				CustomContextTag: to.Ptr("CustomContextTags-1235346475"),
		// 				SinkInfo: &armdataboxedge.RoleSinkInfo{
		// 					RoleID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/role1"),
		// 				},
		// 				SourceInfo: &armdataboxedge.FileSourceInfo{
		// 					ShareID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/shares/share1"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/TriggerGet.json
func ExampleTriggersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTriggersClient().Get(ctx, "testedgedevice", "trigger1", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataboxedge.TriggersClientGetResponse{
	// 	                            TriggerClassification: &armdataboxedge.FileEventTrigger{
	// 		Name: to.Ptr("trigger1"),
	// 		Type: to.Ptr("dataBoxEdgeDevices/triggers"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/triggers/trigger1"),
	// 		Kind: to.Ptr(armdataboxedge.TriggerEventTypeFileEvent),
	// 		Properties: &armdataboxedge.FileTriggerProperties{
	// 			CustomContextTag: to.Ptr("CustomContextTags-1235346475"),
	// 			SinkInfo: &armdataboxedge.RoleSinkInfo{
	// 				RoleID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/role1"),
	// 			},
	// 			SourceInfo: &armdataboxedge.FileSourceInfo{
	// 				ShareID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/shares/share1"),
	// 			},
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/TriggerPut.json
func ExampleTriggersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTriggersClient().BeginCreateOrUpdate(ctx, "testedgedevice", "trigger1", "GroupForEdgeAutomation", &armdataboxedge.FileEventTrigger{
		Kind: to.Ptr(armdataboxedge.TriggerEventTypeFileEvent),
		Properties: &armdataboxedge.FileTriggerProperties{
			CustomContextTag: to.Ptr("CustomContextTags-1235346475"),
			SinkInfo: &armdataboxedge.RoleSinkInfo{
				RoleID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/role1"),
			},
			SourceInfo: &armdataboxedge.FileSourceInfo{
				ShareID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/shares/share1"),
			},
		},
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
	// res = armdataboxedge.TriggersClientCreateOrUpdateResponse{
	// 	                            TriggerClassification: &armdataboxedge.FileEventTrigger{
	// 		Name: to.Ptr("trigger1"),
	// 		Type: to.Ptr("dataBoxEdgeDevices/triggers"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/triggers/trigger1"),
	// 		Kind: to.Ptr(armdataboxedge.TriggerEventTypeFileEvent),
	// 		Properties: &armdataboxedge.FileTriggerProperties{
	// 			CustomContextTag: to.Ptr("CustomContextTags-1235346475"),
	// 			SinkInfo: &armdataboxedge.RoleSinkInfo{
	// 				RoleID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/roles/role1"),
	// 			},
	// 			SourceInfo: &armdataboxedge.FileSourceInfo{
	// 				ShareID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/shares/share1"),
	// 			},
	// 		},
	// 	},
	// 	                        }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/TriggerDelete.json
func ExampleTriggersClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTriggersClient().BeginDelete(ctx, "testedgedevice", "trigger1", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
