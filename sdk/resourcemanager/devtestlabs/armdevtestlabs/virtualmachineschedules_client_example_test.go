//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_List.json
func ExampleVirtualMachineSchedulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineSchedulesClient().NewListPager("resourceGroupName", "{labName}", "{vmName}", &armdevtestlabs.VirtualMachineSchedulesClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.ScheduleList = armdevtestlabs.ScheduleList{
		// 	Value: []*armdevtestlabs.Schedule{
		// 		{
		// 			Name: to.Ptr("LabVmsShutdown"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}/schedules/mySchedule"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.ScheduleProperties{
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T18:40:48.1739018-07:00"); return t}()),
		// 				DailyRecurrence: &armdevtestlabs.DayDetails{
		// 					Time: to.Ptr("1900"),
		// 				},
		// 				HourlyRecurrence: &armdevtestlabs.HourDetails{
		// 					Minute: to.Ptr[int32](30),
		// 				},
		// 				NotificationSettings: &armdevtestlabs.NotificationSettings{
		// 					EmailRecipient: to.Ptr("{email}"),
		// 					NotificationLocale: to.Ptr("EN"),
		// 					Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
		// 					TimeInMinutes: to.Ptr[int32](30),
		// 					WebhookURL: to.Ptr("{webhookUrl}"),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
		// 				TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/labName}/virtualMachines/{vmName}"),
		// 				TaskType: to.Ptr("LabVmsShutdownTask"),
		// 				TimeZoneID: to.Ptr("Pacific Standard Time"),
		// 				WeeklyRecurrence: &armdevtestlabs.WeekDetails{
		// 					Time: to.Ptr("1700"),
		// 					Weekdays: []*string{
		// 						to.Ptr("Friday"),
		// 						to.Ptr("Saturday"),
		// 						to.Ptr("Sunday")},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Get.json
func ExampleVirtualMachineSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineSchedulesClient().Get(ctx, "resourceGroupName", "{labName}", "{vmName}", "LabVmsShutdown", &armdevtestlabs.VirtualMachineSchedulesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armdevtestlabs.Schedule{
	// 	Name: to.Ptr("LabVmsShutdown"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}/schedules/mySchedule"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ScheduleProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T18:40:48.1739018-07:00"); return t}()),
	// 		DailyRecurrence: &armdevtestlabs.DayDetails{
	// 			Time: to.Ptr("1900"),
	// 		},
	// 		HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 			Minute: to.Ptr[int32](30),
	// 		},
	// 		NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 			EmailRecipient: to.Ptr("{email}"),
	// 			NotificationLocale: to.Ptr("EN"),
	// 			Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
	// 			TimeInMinutes: to.Ptr[int32](30),
	// 			WebhookURL: to.Ptr("{webhookUrl}"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
	// 		TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}"),
	// 		TaskType: to.Ptr("LabVmsShutdownTask"),
	// 		TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 		WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 			Time: to.Ptr("1700"),
	// 			Weekdays: []*string{
	// 				to.Ptr("Friday"),
	// 				to.Ptr("Saturday"),
	// 				to.Ptr("Sunday")},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_CreateOrUpdate.json
func ExampleVirtualMachineSchedulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineSchedulesClient().CreateOrUpdate(ctx, "resourceGroupName", "{labName}", "{vmName}", "LabVmsShutdown", armdevtestlabs.Schedule{
		Location: to.Ptr("{location}"),
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
		Properties: &armdevtestlabs.ScheduleProperties{
			DailyRecurrence: &armdevtestlabs.DayDetails{
				Time: to.Ptr("1900"),
			},
			HourlyRecurrence: &armdevtestlabs.HourDetails{
				Minute: to.Ptr[int32](30),
			},
			NotificationSettings: &armdevtestlabs.NotificationSettings{
				EmailRecipient:     to.Ptr("{email}"),
				NotificationLocale: to.Ptr("EN"),
				Status:             to.Ptr(armdevtestlabs.EnableStatusEnabled),
				TimeInMinutes:      to.Ptr[int32](30),
				WebhookURL:         to.Ptr("{webhookUrl}"),
			},
			Status:           to.Ptr(armdevtestlabs.EnableStatusEnabled),
			TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}"),
			TaskType:         to.Ptr("LabVmsShutdownTask"),
			TimeZoneID:       to.Ptr("Pacific Standard Time"),
			WeeklyRecurrence: &armdevtestlabs.WeekDetails{
				Time: to.Ptr("1700"),
				Weekdays: []*string{
					to.Ptr("Friday"),
					to.Ptr("Saturday"),
					to.Ptr("Sunday")},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armdevtestlabs.Schedule{
	// 	Name: to.Ptr("LabVmsShutdown"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}/schedules/LabVmsShutdown"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ScheduleProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T18:40:48.1739018-07:00"); return t}()),
	// 		DailyRecurrence: &armdevtestlabs.DayDetails{
	// 			Time: to.Ptr("1900"),
	// 		},
	// 		HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 			Minute: to.Ptr[int32](30),
	// 		},
	// 		NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 			EmailRecipient: to.Ptr("{email}"),
	// 			NotificationLocale: to.Ptr("EN"),
	// 			Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
	// 			TimeInMinutes: to.Ptr[int32](30),
	// 			WebhookURL: to.Ptr("{webhookUrl}"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
	// 		TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}"),
	// 		TaskType: to.Ptr("LabVmsShutdownTask"),
	// 		TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 		WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 			Time: to.Ptr("1700"),
	// 			Weekdays: []*string{
	// 				to.Ptr("Friday"),
	// 				to.Ptr("Saturday"),
	// 				to.Ptr("Sunday")},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Delete.json
func ExampleVirtualMachineSchedulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewVirtualMachineSchedulesClient().Delete(ctx, "resourceGroupName", "{labName}", "{vmName}", "LabVmsShutdown", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Update.json
func ExampleVirtualMachineSchedulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineSchedulesClient().Update(ctx, "resourceGroupName", "{labName}", "{vmName}", "LabVmsShutdown", armdevtestlabs.ScheduleFragment{
		Tags: map[string]*string{
			"tagName1": to.Ptr("tagValue1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Schedule = armdevtestlabs.Schedule{
	// 	Name: to.Ptr("LabVmsShutdown"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/virtualMachines/schedules"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}/schedules/mySchedule"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"tagName1": to.Ptr("tagValue1"),
	// 	},
	// 	Properties: &armdevtestlabs.ScheduleProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-31T18:40:48.1739018-07:00"); return t}()),
	// 		DailyRecurrence: &armdevtestlabs.DayDetails{
	// 			Time: to.Ptr("1900"),
	// 		},
	// 		HourlyRecurrence: &armdevtestlabs.HourDetails{
	// 			Minute: to.Ptr[int32](30),
	// 		},
	// 		NotificationSettings: &armdevtestlabs.NotificationSettings{
	// 			EmailRecipient: to.Ptr("{email}"),
	// 			NotificationLocale: to.Ptr("EN"),
	// 			Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
	// 			TimeInMinutes: to.Ptr[int32](30),
	// 			WebhookURL: to.Ptr("{webhookUrl}"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		Status: to.Ptr(armdevtestlabs.EnableStatusEnabled),
	// 		TargetResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}"),
	// 		TaskType: to.Ptr("LabVmsShutdownTask"),
	// 		TimeZoneID: to.Ptr("Pacific Standard Time"),
	// 		WeeklyRecurrence: &armdevtestlabs.WeekDetails{
	// 			Time: to.Ptr("1700"),
	// 			Weekdays: []*string{
	// 				to.Ptr("Friday"),
	// 				to.Ptr("Saturday"),
	// 				to.Ptr("Sunday")},
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Execute.json
func ExampleVirtualMachineSchedulesClient_BeginExecute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineSchedulesClient().BeginExecute(ctx, "resourceGroupName", "{labName}", "{vmName}", "LabVmsShutdown", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
