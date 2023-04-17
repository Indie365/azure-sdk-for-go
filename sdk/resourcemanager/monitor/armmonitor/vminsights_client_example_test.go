//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusVMScaleSet.json
func ExampleVMInsightsClient_GetOnboardingStatus_getStatusForAVmScaleSetThatIsActivelyReportingData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMInsightsClient().GetOnboardingStatus(ctx, "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/my-service-cluster/providers/Microsoft.Compute/virtualMachineScaleSets/scale-set-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMInsightsOnboardingStatus = armmonitor.VMInsightsOnboardingStatus{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Insights/vmInsightsOnboardingStatuses"),
	// 	ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/my-service-cluster/providers/Microsoft.Compute/virtualMachineScaleSets/scale-set-01/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"),
	// 	Properties: &armmonitor.VMInsightsOnboardingStatusProperties{
	// 		Data: []*armmonitor.DataContainer{
	// 			{
	// 				Workspace: &armmonitor.WorkspaceInfo{
	// 					ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourcegroups/monitoring/providers/microsoft.operationalinsights/workspaces/vm-monitoring"),
	// 					Location: to.Ptr("eastus"),
	// 					Properties: &armmonitor.WorkspaceInfoProperties{
	// 						CustomerID: to.Ptr("c7f8f44d-d8ee-4b79-9d9a-4d8a1f2a112a"),
	// 					},
	// 				},
	// 		}},
	// 		DataStatus: to.Ptr(armmonitor.DataStatusPresent),
	// 		OnboardingStatus: to.Ptr(armmonitor.OnboardingStatusOnboarded),
	// 		ResourceID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/my-service-cluster/providers/Microsoft.Compute/virtualMachineScaleSets/scale-set-01"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusSingleVMUnknown.json
func ExampleVMInsightsClient_GetOnboardingStatus_getStatusForAVmThatHasNotYetReportedData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMInsightsClient().GetOnboardingStatus(ctx, "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMInsightsOnboardingStatus = armmonitor.VMInsightsOnboardingStatus{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Insights/vmInsightsOnboardingStatuses"),
	// 	ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"),
	// 	Properties: &armmonitor.VMInsightsOnboardingStatusProperties{
	// 		Data: []*armmonitor.DataContainer{
	// 		},
	// 		DataStatus: to.Ptr(armmonitor.DataStatusNotPresent),
	// 		OnboardingStatus: to.Ptr(armmonitor.OnboardingStatusUnknown),
	// 		ResourceID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusSingleVM.json
func ExampleVMInsightsClient_GetOnboardingStatus_getStatusForAVmThatIsActivelyReportingData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMInsightsClient().GetOnboardingStatus(ctx, "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMInsightsOnboardingStatus = armmonitor.VMInsightsOnboardingStatus{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Insights/vmInsightsOnboardingStatuses"),
	// 	ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"),
	// 	Properties: &armmonitor.VMInsightsOnboardingStatusProperties{
	// 		Data: []*armmonitor.DataContainer{
	// 			{
	// 				Workspace: &armmonitor.WorkspaceInfo{
	// 					ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourcegroups/monitoring/providers/microsoft.operationalinsights/workspaces/vm-monitoring"),
	// 					Location: to.Ptr("eastus"),
	// 					Properties: &armmonitor.WorkspaceInfoProperties{
	// 						CustomerID: to.Ptr("c7f8f44d-d8ee-4b79-9d9a-4d8a1f2a112a"),
	// 					},
	// 				},
	// 		}},
	// 		DataStatus: to.Ptr(armmonitor.DataStatusPresent),
	// 		OnboardingStatus: to.Ptr(armmonitor.OnboardingStatusOnboarded),
	// 		ResourceID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusResourceGroup.json
func ExampleVMInsightsClient_GetOnboardingStatus_getStatusForAResourceGroupThatHasAtLeastOneVmThatIsActivelyReportingData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMInsightsClient().GetOnboardingStatus(ctx, "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/resource-group-with-vms", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMInsightsOnboardingStatus = armmonitor.VMInsightsOnboardingStatus{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Insights/vmInsightsOnboardingStatuses"),
	// 	ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/resource-group-with-vms/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"),
	// 	Properties: &armmonitor.VMInsightsOnboardingStatusProperties{
	// 		Data: []*armmonitor.DataContainer{
	// 			{
	// 				Workspace: &armmonitor.WorkspaceInfo{
	// 					ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourcegroups/monitoring/providers/microsoft.operationalinsights/workspaces/vm-monitoring"),
	// 					Location: to.Ptr("eastus"),
	// 					Properties: &armmonitor.WorkspaceInfoProperties{
	// 						CustomerID: to.Ptr("c7f8f44d-d8ee-4b79-9d9a-4d8a1f2a112a"),
	// 					},
	// 				},
	// 		}},
	// 		DataStatus: to.Ptr(armmonitor.DataStatusPresent),
	// 		OnboardingStatus: to.Ptr(armmonitor.OnboardingStatusOnboarded),
	// 		ResourceID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/resource-group-with-vms"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/getOnboardingStatusSubscription.json
func ExampleVMInsightsClient_GetOnboardingStatus_getStatusForASubscriptionThatHasAtLeastOneVmThatIsActivelyReportingData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVMInsightsClient().GetOnboardingStatus(ctx, "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VMInsightsOnboardingStatus = armmonitor.VMInsightsOnboardingStatus{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Insights/vmInsightsOnboardingStatuses"),
	// 	ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default"),
	// 	Properties: &armmonitor.VMInsightsOnboardingStatusProperties{
	// 		Data: []*armmonitor.DataContainer{
	// 			{
	// 				Workspace: &armmonitor.WorkspaceInfo{
	// 					ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourcegroups/monitoring/providers/microsoft.operationalinsights/workspaces/vm-monitoring-secondary"),
	// 					Location: to.Ptr("eastus"),
	// 					Properties: &armmonitor.WorkspaceInfoProperties{
	// 						CustomerID: to.Ptr("f096d163-206e-4abf-9db3-2c62af003d68"),
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Workspace: &armmonitor.WorkspaceInfo{
	// 					ID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourcegroups/monitoring/providers/microsoft.operationalinsights/workspaces/vm-monitoring"),
	// 					Location: to.Ptr("eastus"),
	// 					Properties: &armmonitor.WorkspaceInfoProperties{
	// 						CustomerID: to.Ptr("c7f8f44d-d8ee-4b79-9d9a-4d8a1f2a112a"),
	// 					},
	// 				},
	// 		}},
	// 		DataStatus: to.Ptr(armmonitor.DataStatusPresent),
	// 		OnboardingStatus: to.Ptr(armmonitor.OnboardingStatusOnboarded),
	// 		ResourceID: to.Ptr("/subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87"),
	// 	},
	// }
}
