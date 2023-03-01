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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceGet.json
func ExamplePrivateLinkScopedResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewPrivateLinkScopedResourcesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "MyResourceGroup", "MyPrivateLinkScope", "scoped-resource-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScopedResource = armmonitor.ScopedResource{
	// 	Name: to.Ptr("scoped-resource-name"),
	// 	Type: to.Ptr("Microsoft.Insights/privateLinkScopes/scopedResources"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/scopedResources/scoped-resource-name"),
	// 	Properties: &armmonitor.ScopedResourceProperties{
	// 		LinkedResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/components/my-component"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// 	SystemData: &armmonitor.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 		CreatedBy: to.Ptr("bobby@contoso.com"),
	// 		CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bobby@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceUpdate.json
func ExamplePrivateLinkScopedResourcesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewPrivateLinkScopedResourcesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "MyResourceGroup", "MyPrivateLinkScope", "scoped-resource-name", armmonitor.ScopedResource{
		Properties: &armmonitor.ScopedResourceProperties{
			LinkedResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/components/my-component"),
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
	// res.ScopedResource = armmonitor.ScopedResource{
	// 	Name: to.Ptr("scoped-resource-name"),
	// 	Type: to.Ptr("Microsoft.Insights/privateLinkScopes/scopedResources"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/scopedResources/scoped-resource-name"),
	// 	Properties: &armmonitor.ScopedResourceProperties{
	// 		LinkedResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/components/my-component"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// 	SystemData: &armmonitor.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 		CreatedBy: to.Ptr("bobby@contoso.com"),
	// 		CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("bobby@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceDelete.json
func ExamplePrivateLinkScopedResourcesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewPrivateLinkScopedResourcesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "MyResourceGroup", "MyPrivateLinkScope", "scoped-resource-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5c9459305484e0456b4a922e3d31a61e2ddd3c99/specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceList.json
func ExamplePrivateLinkScopedResourcesClient_NewListByPrivateLinkScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmonitor.NewPrivateLinkScopedResourcesClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByPrivateLinkScopePager("MyResourceGroup", "MyPrivateLinkScope", nil)
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
		// page.ScopedResourceListResult = armmonitor.ScopedResourceListResult{
		// 	Value: []*armmonitor.ScopedResource{
		// 		{
		// 			Name: to.Ptr("scoped-resource-name"),
		// 			Type: to.Ptr("Microsoft.Insights/privateLinkScopes/scopedResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/scopedResources/scoped-resource-name"),
		// 			Properties: &armmonitor.ScopedResourceProperties{
		// 				LinkedResourceID: to.Ptr("/subscriptions/00000000-0000-2222-3333-444444444444/resourceGroups/MyComponentResourceGroup/providers/Microsoft.Insights/components/my-component"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SystemData: &armmonitor.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
		// 				CreatedBy: to.Ptr("bobby@contoso.com"),
		// 				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("bobby@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("other-scoped-resource-name"),
		// 			Type: to.Ptr("Microsoft.Insights/privateLinkScopes/scopedResources"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/privateLinkScopes/MyPrivateLinkScope/scopedResources/other-scoped-resource-name"),
		// 			Properties: &armmonitor.ScopedResourceProperties{
		// 				LinkedResourceID: to.Ptr("/subscriptions/00000000-3333-2222-5555-444444444444/resourceGroups/MyWorkspaceResourceGroup/providers/Microsoft.OperationalInsights/workspaces/my-workspace"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 			SystemData: &armmonitor.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
		// 				CreatedBy: to.Ptr("bobby@contoso.com"),
		// 				CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-16T12:59:57.051056Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("bobby@contoso.com"),
		// 				LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}
