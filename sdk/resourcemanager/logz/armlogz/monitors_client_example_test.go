//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MonitoredResources_List.json
func ExampleMonitorsClient_NewListMonitoredResourcesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListMonitoredResourcesPager("myResourceGroup", "myMonitor", nil)
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
		// page.MonitoredResourceListResponse = armlogz.MonitoredResourceListResponse{
		// 	Value: []*armlogz.MonitoredResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor/listMonitoredResources"),
		// 			ReasonForLogsStatus: to.Ptr("CapturedByRules"),
		// 			ReasonForMetricsStatus: to.Ptr("CapturedByRules"),
		// 			SendingLogs: to.Ptr(true),
		// 			SendingMetrics: to.Ptr(true),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_List.json
func ExampleMonitorsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListBySubscriptionPager(nil)
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
		// page.MonitorResourceListResponse = armlogz.MonitorResourceListResponse{
		// 	Value: []*armlogz.MonitorResource{
		// 		{
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armlogz.MonitorProperties{
		// 				MonitoringStatus: to.Ptr(armlogz.MonitoringStatusEnabled),
		// 				PlanData: &armlogz.PlanData{
		// 					BillingCycle: to.Ptr("Monthly"),
		// 					EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T15:14:33+02:00"); return t}()),
		// 					PlanDetails: to.Ptr("logzapitestplan"),
		// 					UsageType: to.Ptr("Committed"),
		// 				},
		// 				ProvisioningState: to.Ptr(armlogz.ProvisioningStateSucceeded),
		// 				UserInfo: &armlogz.UserInfo{
		// 					EmailAddress: to.Ptr("alice@microsoft.com"),
		// 					FirstName: to.Ptr("Alice"),
		// 					LastName: to.Ptr("Bob"),
		// 					PhoneNumber: to.Ptr("123456"),
		// 				},
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_ListByResourceGroup.json
func ExampleMonitorsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.MonitorResourceListResponse = armlogz.MonitorResourceListResponse{
		// 	Value: []*armlogz.MonitorResource{
		// 		{
		// 			Name: to.Ptr("myMonitor"),
		// 			Type: to.Ptr("Microsoft.Logz/monitors"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armlogz.MonitorProperties{
		// 				PlanData: &armlogz.PlanData{
		// 					BillingCycle: to.Ptr("Monthly"),
		// 					EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T15:14:33+02:00"); return t}()),
		// 					PlanDetails: to.Ptr("logzapitestplan"),
		// 					UsageType: to.Ptr("Committed"),
		// 				},
		// 				ProvisioningState: to.Ptr(armlogz.ProvisioningStateSucceeded),
		// 				UserInfo: &armlogz.UserInfo{
		// 					EmailAddress: to.Ptr("alice@microsoft.com"),
		// 					FirstName: to.Ptr("Alice"),
		// 					LastName: to.Ptr("Bob"),
		// 					PhoneNumber: to.Ptr("123456"),
		// 				},
		// 			},
		// 			Tags: map[string]*string{
		// 				"Environment": to.Ptr("Dev"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Get.json
func ExampleMonitorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().Get(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitorResource = armlogz.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Logz/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armlogz.MonitorProperties{
	// 		MonitoringStatus: to.Ptr(armlogz.MonitoringStatusEnabled),
	// 		PlanData: &armlogz.PlanData{
	// 			BillingCycle: to.Ptr("Monthly"),
	// 			EffectiveDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-30T15:14:33+02:00"); return t}()),
	// 			PlanDetails: to.Ptr("logzapitestplan"),
	// 			UsageType: to.Ptr("Committed"),
	// 		},
	// 		ProvisioningState: to.Ptr(armlogz.ProvisioningStateSucceeded),
	// 		UserInfo: &armlogz.UserInfo{
	// 			EmailAddress: to.Ptr("alice@microsoft.com"),
	// 			FirstName: to.Ptr("Alice"),
	// 			LastName: to.Ptr("Bob"),
	// 			PhoneNumber: to.Ptr("123456"),
	// 		},
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Create.json
func ExampleMonitorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginCreate(ctx, "myResourceGroup", "myMonitor", &armlogz.MonitorsClientBeginCreateOptions{Body: nil})
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
	// res.MonitorResource = armlogz.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Logz/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US 2"),
	// 	Properties: &armlogz.MonitorProperties{
	// 		ProvisioningState: to.Ptr(armlogz.ProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Update.json
func ExampleMonitorsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMonitorsClient().Update(ctx, "myResourceGroup", "myMonitor", &armlogz.MonitorsClientUpdateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitorResource = armlogz.MonitorResource{
	// 	Name: to.Ptr("myMonitor"),
	// 	Type: to.Ptr("Microsoft.Logz/monitors"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/monitors/myMonitor"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armlogz.MonitorProperties{
	// 		MonitoringStatus: to.Ptr(armlogz.MonitoringStatusEnabled),
	// 		ProvisioningState: to.Ptr(armlogz.ProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Environment": to.Ptr("Dev"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Monitors_Delete.json
func ExampleMonitorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMonitorsClient().BeginDelete(ctx, "myResourceGroup", "myMonitor", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/MainAccount_listUserRoles.json
func ExampleMonitorsClient_NewListUserRolesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMonitorsClient().NewListUserRolesPager("myResourceGroup", "myMonitor", &armlogz.MonitorsClientListUserRolesOptions{Body: nil})
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
		// page.UserRoleListResponse = armlogz.UserRoleListResponse{
		// 	Value: []*armlogz.UserRoleResponse{
		// 		{
		// 			Role: to.Ptr(armlogz.UserRoleAdmin),
		// 	}},
		// }
	}
}
