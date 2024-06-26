//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_listBySubscription.json
func ExampleCloudExadataInfrastructuresClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudExadataInfrastructuresClient().NewListBySubscriptionPager(nil)
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
		// page.CloudExadataInfrastructureListResult = armoracledatabase.CloudExadataInfrastructureListResult{
		// 	Value: []*armoracledatabase.CloudExadataInfrastructure{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/cloudExadataInfrastructures"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"tagK1": to.Ptr("tagV1"),
		// 			},
		// 			Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
		// 				ActivatedStorageCount: to.Ptr[int32](1),
		// 				AdditionalStorageCount: to.Ptr[int32](1),
		// 				AvailableStorageSizeInGbs: to.Ptr[int32](1000),
		// 				ComputeCount: to.Ptr[int32](100),
		// 				CPUCount: to.Ptr[int32](10),
		// 				CustomerContacts: []*armoracledatabase.CustomerContact{
		// 					{
		// 						Email: to.Ptr("noreply@oracle.com"),
		// 				}},
		// 				DataStorageSizeInTbs: to.Ptr[float64](10),
		// 				DataStorageSizeInTbs: to.Ptr[float64](10),
		// 				DbNodeStorageSizeInGbs: to.Ptr[int32](10),
		// 				DbServerVersion: to.Ptr("19.0.0.0"),
		// 				DisplayName: to.Ptr("infra 1"),
		// 				EstimatedPatchingTime: &armoracledatabase.EstimatedPatchingTime{
		// 					EstimatedDbServerPatchingTime: to.Ptr[int32](3000),
		// 					EstimatedNetworkSwitchesPatchingTime: to.Ptr[int32](3000),
		// 					EstimatedStorageServerPatchingTime: to.Ptr[int32](3000),
		// 					TotalEstimatedPatchingTime: to.Ptr[int32](3000),
		// 				},
		// 				LastMaintenanceRunID: to.Ptr("ocid1..aaaaa"),
		// 				LifecycleDetails: to.Ptr("none"),
		// 				MaintenanceWindow: &armoracledatabase.MaintenanceWindow{
		// 					CustomActionTimeoutInMins: to.Ptr[int32](120),
		// 					DaysOfWeek: []*armoracledatabase.DayOfWeek{
		// 						{
		// 							Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
		// 					}},
		// 					HoursOfDay: []*int32{
		// 						to.Ptr[int32](0)},
		// 						IsCustomActionTimeoutEnabled: to.Ptr(true),
		// 						IsMonthlyPatchingEnabled: to.Ptr(true),
		// 						LeadTimeInWeeks: to.Ptr[int32](0),
		// 						Months: []*armoracledatabase.Month{
		// 							{
		// 								Name: to.Ptr(armoracledatabase.MonthNameJanuary),
		// 						}},
		// 						PatchingMode: to.Ptr(armoracledatabase.PatchingModeRolling),
		// 						Preference: to.Ptr(armoracledatabase.PreferenceNoPreference),
		// 						WeeksOfMonth: []*int32{
		// 							to.Ptr[int32](0)},
		// 						},
		// 						MaxCPUCount: to.Ptr[int32](100),
		// 						MaxDataStorageInTbs: to.Ptr[float64](1000),
		// 						MaxDbNodeStorageSizeInGbs: to.Ptr[int32](10),
		// 						MaxMemoryInGbs: to.Ptr[int32](1000),
		// 						MemorySizeInGbs: to.Ptr[int32](100),
		// 						MonthlyDbServerVersion: to.Ptr("aaaa"),
		// 						MonthlyStorageServerVersion: to.Ptr("aaaa"),
		// 						NextMaintenanceRunID: to.Ptr("ocid1..aaaaaa"),
		// 						OciURL: to.Ptr("https://url"),
		// 						Ocid: to.Ptr("ocid1..aaaaaa"),
		// 						ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 						Shape: to.Ptr("EXADATA.X9M"),
		// 						StorageCount: to.Ptr[int32](10),
		// 						StorageServerVersion: to.Ptr("0.0"),
		// 						TimeCreated: to.Ptr("2023-02-01T01:01:00"),
		// 						TotalStorageSizeInGbs: to.Ptr[int32](1000),
		// 					},
		// 					Zones: []*string{
		// 						to.Ptr("1")},
		// 				}},
		// 			}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_listByResourceGroup.json
func ExampleCloudExadataInfrastructuresClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudExadataInfrastructuresClient().NewListByResourceGroupPager("rg000", nil)
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
		// page.CloudExadataInfrastructureListResult = armoracledatabase.CloudExadataInfrastructureListResult{
		// 	Value: []*armoracledatabase.CloudExadataInfrastructure{
		// 		{
		// 			Type: to.Ptr("Oracle.Database/cloudExadataInfrastructures"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"tagK1": to.Ptr("tagV1"),
		// 			},
		// 			Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
		// 				ActivatedStorageCount: to.Ptr[int32](1),
		// 				AdditionalStorageCount: to.Ptr[int32](1),
		// 				AvailableStorageSizeInGbs: to.Ptr[int32](1000),
		// 				ComputeCount: to.Ptr[int32](100),
		// 				CPUCount: to.Ptr[int32](10),
		// 				CustomerContacts: []*armoracledatabase.CustomerContact{
		// 					{
		// 						Email: to.Ptr("noreply@oracle.com"),
		// 				}},
		// 				DataStorageSizeInTbs: to.Ptr[float64](10),
		// 				DataStorageSizeInTbs: to.Ptr[float64](10),
		// 				DbNodeStorageSizeInGbs: to.Ptr[int32](10),
		// 				DbServerVersion: to.Ptr("19.0.0.0"),
		// 				DisplayName: to.Ptr("infra 1"),
		// 				EstimatedPatchingTime: &armoracledatabase.EstimatedPatchingTime{
		// 					EstimatedDbServerPatchingTime: to.Ptr[int32](3000),
		// 					EstimatedNetworkSwitchesPatchingTime: to.Ptr[int32](3000),
		// 					EstimatedStorageServerPatchingTime: to.Ptr[int32](3000),
		// 					TotalEstimatedPatchingTime: to.Ptr[int32](3000),
		// 				},
		// 				LastMaintenanceRunID: to.Ptr("ocid1..aaaaa"),
		// 				LifecycleDetails: to.Ptr("none"),
		// 				MaintenanceWindow: &armoracledatabase.MaintenanceWindow{
		// 					CustomActionTimeoutInMins: to.Ptr[int32](120),
		// 					DaysOfWeek: []*armoracledatabase.DayOfWeek{
		// 						{
		// 							Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
		// 					}},
		// 					HoursOfDay: []*int32{
		// 						to.Ptr[int32](0)},
		// 						IsCustomActionTimeoutEnabled: to.Ptr(true),
		// 						IsMonthlyPatchingEnabled: to.Ptr(true),
		// 						LeadTimeInWeeks: to.Ptr[int32](0),
		// 						Months: []*armoracledatabase.Month{
		// 							{
		// 								Name: to.Ptr(armoracledatabase.MonthNameJanuary),
		// 						}},
		// 						PatchingMode: to.Ptr(armoracledatabase.PatchingModeRolling),
		// 						Preference: to.Ptr(armoracledatabase.PreferenceNoPreference),
		// 						WeeksOfMonth: []*int32{
		// 							to.Ptr[int32](0)},
		// 						},
		// 						MaxCPUCount: to.Ptr[int32](100),
		// 						MaxDataStorageInTbs: to.Ptr[float64](1000),
		// 						MaxDbNodeStorageSizeInGbs: to.Ptr[int32](10),
		// 						MaxMemoryInGbs: to.Ptr[int32](1000),
		// 						MemorySizeInGbs: to.Ptr[int32](100),
		// 						MonthlyDbServerVersion: to.Ptr("aaaa"),
		// 						MonthlyStorageServerVersion: to.Ptr("aaaa"),
		// 						NextMaintenanceRunID: to.Ptr("ocid1..aaaaaa"),
		// 						OciURL: to.Ptr("https://url"),
		// 						Ocid: to.Ptr("ocid1..aaaaaa"),
		// 						ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 						Shape: to.Ptr("EXADATA.X9M"),
		// 						StorageCount: to.Ptr[int32](10),
		// 						StorageServerVersion: to.Ptr("0.0"),
		// 						TimeCreated: to.Ptr("2023-02-01T01:01:00"),
		// 						TotalStorageSizeInGbs: to.Ptr[int32](1000),
		// 					},
		// 					Zones: []*string{
		// 						to.Ptr("1")},
		// 				}},
		// 			}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_get.json
func ExampleCloudExadataInfrastructuresClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCloudExadataInfrastructuresClient().Get(ctx, "rg000", "infra1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CloudExadataInfrastructure = armoracledatabase.CloudExadataInfrastructure{
	// 	Type: to.Ptr("Oracle.Database/cloudExadataInfrastructures"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tagK1": to.Ptr("tagV1"),
	// 	},
	// 	Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
	// 		ActivatedStorageCount: to.Ptr[int32](1),
	// 		AdditionalStorageCount: to.Ptr[int32](1),
	// 		AvailableStorageSizeInGbs: to.Ptr[int32](1000),
	// 		ComputeCount: to.Ptr[int32](100),
	// 		CPUCount: to.Ptr[int32](10),
	// 		CustomerContacts: []*armoracledatabase.CustomerContact{
	// 			{
	// 				Email: to.Ptr("noreply@oracle.com"),
	// 		}},
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 		DbServerVersion: to.Ptr("19.0.0.0"),
	// 		DisplayName: to.Ptr("infra 1"),
	// 		EstimatedPatchingTime: &armoracledatabase.EstimatedPatchingTime{
	// 			EstimatedDbServerPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedNetworkSwitchesPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedStorageServerPatchingTime: to.Ptr[int32](3000),
	// 			TotalEstimatedPatchingTime: to.Ptr[int32](3000),
	// 		},
	// 		LastMaintenanceRunID: to.Ptr("ocid1..aaaaa"),
	// 		LifecycleDetails: to.Ptr("none"),
	// 		MaintenanceWindow: &armoracledatabase.MaintenanceWindow{
	// 			CustomActionTimeoutInMins: to.Ptr[int32](120),
	// 			DaysOfWeek: []*armoracledatabase.DayOfWeek{
	// 				{
	// 					Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
	// 			}},
	// 			HoursOfDay: []*int32{
	// 				to.Ptr[int32](0)},
	// 				IsCustomActionTimeoutEnabled: to.Ptr(true),
	// 				IsMonthlyPatchingEnabled: to.Ptr(true),
	// 				LeadTimeInWeeks: to.Ptr[int32](0),
	// 				Months: []*armoracledatabase.Month{
	// 					{
	// 						Name: to.Ptr(armoracledatabase.MonthNameJanuary),
	// 				}},
	// 				PatchingMode: to.Ptr(armoracledatabase.PatchingModeRolling),
	// 				Preference: to.Ptr(armoracledatabase.PreferenceNoPreference),
	// 				WeeksOfMonth: []*int32{
	// 					to.Ptr[int32](0)},
	// 				},
	// 				MaxCPUCount: to.Ptr[int32](100),
	// 				MaxDataStorageInTbs: to.Ptr[float64](1000),
	// 				MaxDbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 				MaxMemoryInGbs: to.Ptr[int32](1000),
	// 				MemorySizeInGbs: to.Ptr[int32](100),
	// 				MonthlyDbServerVersion: to.Ptr("aaaa"),
	// 				MonthlyStorageServerVersion: to.Ptr("aaaa"),
	// 				NextMaintenanceRunID: to.Ptr("ocid1..aaaaaa"),
	// 				OciURL: to.Ptr("https://url"),
	// 				Ocid: to.Ptr("ocid1..aaaaaa"),
	// 				ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 				Shape: to.Ptr("EXADATA.X9M"),
	// 				StorageCount: to.Ptr[int32](10),
	// 				StorageServerVersion: to.Ptr("0.0"),
	// 				TimeCreated: to.Ptr("2023-02-01T01:01:00"),
	// 				TotalStorageSizeInGbs: to.Ptr[int32](1000),
	// 			},
	// 			Zones: []*string{
	// 				to.Ptr("1")},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_create.json
func ExampleCloudExadataInfrastructuresClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudExadataInfrastructuresClient().BeginCreateOrUpdate(ctx, "rg000", "infra1", armoracledatabase.CloudExadataInfrastructure{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"tagK1": to.Ptr("tagV1"),
		},
		Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
			ComputeCount: to.Ptr[int32](100),
			DisplayName:  to.Ptr("infra 1"),
			Shape:        to.Ptr("EXADATA.X9M"),
			StorageCount: to.Ptr[int32](10),
		},
		Zones: []*string{
			to.Ptr("1")},
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
	// res.CloudExadataInfrastructure = armoracledatabase.CloudExadataInfrastructure{
	// 	Type: to.Ptr("Oracle.Database/cloudExadataInfrastructures"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tagK1": to.Ptr("tagV1"),
	// 	},
	// 	Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
	// 		ActivatedStorageCount: to.Ptr[int32](1),
	// 		AdditionalStorageCount: to.Ptr[int32](1),
	// 		AvailableStorageSizeInGbs: to.Ptr[int32](1000),
	// 		ComputeCount: to.Ptr[int32](100),
	// 		CPUCount: to.Ptr[int32](10),
	// 		CustomerContacts: []*armoracledatabase.CustomerContact{
	// 			{
	// 				Email: to.Ptr("noreply@oracle.com"),
	// 		}},
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 		DbServerVersion: to.Ptr("19.0.0.0"),
	// 		DisplayName: to.Ptr("infra 1"),
	// 		EstimatedPatchingTime: &armoracledatabase.EstimatedPatchingTime{
	// 			EstimatedDbServerPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedNetworkSwitchesPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedStorageServerPatchingTime: to.Ptr[int32](3000),
	// 			TotalEstimatedPatchingTime: to.Ptr[int32](3000),
	// 		},
	// 		LastMaintenanceRunID: to.Ptr("ocid1..aaaaa"),
	// 		LifecycleDetails: to.Ptr("none"),
	// 		MaintenanceWindow: &armoracledatabase.MaintenanceWindow{
	// 			CustomActionTimeoutInMins: to.Ptr[int32](120),
	// 			DaysOfWeek: []*armoracledatabase.DayOfWeek{
	// 				{
	// 					Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
	// 			}},
	// 			HoursOfDay: []*int32{
	// 				to.Ptr[int32](0)},
	// 				IsCustomActionTimeoutEnabled: to.Ptr(true),
	// 				IsMonthlyPatchingEnabled: to.Ptr(true),
	// 				LeadTimeInWeeks: to.Ptr[int32](0),
	// 				Months: []*armoracledatabase.Month{
	// 					{
	// 						Name: to.Ptr(armoracledatabase.MonthNameJanuary),
	// 				}},
	// 				PatchingMode: to.Ptr(armoracledatabase.PatchingModeRolling),
	// 				Preference: to.Ptr(armoracledatabase.PreferenceNoPreference),
	// 				WeeksOfMonth: []*int32{
	// 					to.Ptr[int32](0)},
	// 				},
	// 				MaxCPUCount: to.Ptr[int32](100),
	// 				MaxDataStorageInTbs: to.Ptr[float64](1000),
	// 				MaxDbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 				MaxMemoryInGbs: to.Ptr[int32](1000),
	// 				MemorySizeInGbs: to.Ptr[int32](100),
	// 				MonthlyDbServerVersion: to.Ptr("aaaa"),
	// 				MonthlyStorageServerVersion: to.Ptr("aaaa"),
	// 				NextMaintenanceRunID: to.Ptr("ocid1..aaaaaa"),
	// 				OciURL: to.Ptr("https://url"),
	// 				Ocid: to.Ptr("ocid1..aaaaaa"),
	// 				ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 				Shape: to.Ptr("EXADATA.X9M"),
	// 				StorageCount: to.Ptr[int32](10),
	// 				StorageServerVersion: to.Ptr("0.0"),
	// 				TimeCreated: to.Ptr("2023-02-01T01:01:00"),
	// 				TotalStorageSizeInGbs: to.Ptr[int32](1000),
	// 			},
	// 			Zones: []*string{
	// 				to.Ptr("1")},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_patch.json
func ExampleCloudExadataInfrastructuresClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudExadataInfrastructuresClient().BeginUpdate(ctx, "rg000", "infra1", armoracledatabase.CloudExadataInfrastructureUpdate{}, nil)
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
	// res.CloudExadataInfrastructure = armoracledatabase.CloudExadataInfrastructure{
	// 	Type: to.Ptr("Oracle.Database/cloudExadataInfrastructures"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tagK1": to.Ptr("tagV1"),
	// 	},
	// 	Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
	// 		ActivatedStorageCount: to.Ptr[int32](1),
	// 		AdditionalStorageCount: to.Ptr[int32](1),
	// 		AvailableStorageSizeInGbs: to.Ptr[int32](1000),
	// 		ComputeCount: to.Ptr[int32](100),
	// 		CPUCount: to.Ptr[int32](10),
	// 		CustomerContacts: []*armoracledatabase.CustomerContact{
	// 			{
	// 				Email: to.Ptr("noreply@oracle.com"),
	// 		}},
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 		DbServerVersion: to.Ptr("19.0.0.0"),
	// 		DisplayName: to.Ptr("infra 1"),
	// 		EstimatedPatchingTime: &armoracledatabase.EstimatedPatchingTime{
	// 			EstimatedDbServerPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedNetworkSwitchesPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedStorageServerPatchingTime: to.Ptr[int32](3000),
	// 			TotalEstimatedPatchingTime: to.Ptr[int32](3000),
	// 		},
	// 		LastMaintenanceRunID: to.Ptr("ocid1..aaaaa"),
	// 		LifecycleDetails: to.Ptr("none"),
	// 		MaintenanceWindow: &armoracledatabase.MaintenanceWindow{
	// 			CustomActionTimeoutInMins: to.Ptr[int32](120),
	// 			DaysOfWeek: []*armoracledatabase.DayOfWeek{
	// 				{
	// 					Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
	// 			}},
	// 			HoursOfDay: []*int32{
	// 				to.Ptr[int32](0)},
	// 				IsCustomActionTimeoutEnabled: to.Ptr(true),
	// 				IsMonthlyPatchingEnabled: to.Ptr(true),
	// 				LeadTimeInWeeks: to.Ptr[int32](0),
	// 				Months: []*armoracledatabase.Month{
	// 					{
	// 						Name: to.Ptr(armoracledatabase.MonthNameJanuary),
	// 				}},
	// 				PatchingMode: to.Ptr(armoracledatabase.PatchingModeRolling),
	// 				Preference: to.Ptr(armoracledatabase.PreferenceNoPreference),
	// 				WeeksOfMonth: []*int32{
	// 					to.Ptr[int32](0)},
	// 				},
	// 				MaxCPUCount: to.Ptr[int32](100),
	// 				MaxDataStorageInTbs: to.Ptr[float64](1000),
	// 				MaxDbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 				MaxMemoryInGbs: to.Ptr[int32](1000),
	// 				MemorySizeInGbs: to.Ptr[int32](100),
	// 				MonthlyDbServerVersion: to.Ptr("aaaa"),
	// 				MonthlyStorageServerVersion: to.Ptr("aaaa"),
	// 				NextMaintenanceRunID: to.Ptr("ocid1..aaaaaa"),
	// 				OciURL: to.Ptr("https://url"),
	// 				Ocid: to.Ptr("ocid1..aaaaaa"),
	// 				ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 				Shape: to.Ptr("EXADATA.X9M"),
	// 				StorageCount: to.Ptr[int32](10),
	// 				StorageServerVersion: to.Ptr("0.0"),
	// 				TimeCreated: to.Ptr("2023-02-01T01:01:00"),
	// 				TotalStorageSizeInGbs: to.Ptr[int32](1000),
	// 			},
	// 			Zones: []*string{
	// 				to.Ptr("1")},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_delete.json
func ExampleCloudExadataInfrastructuresClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudExadataInfrastructuresClient().BeginDelete(ctx, "rg000", "infra1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0ce2859b2f018adcea3d14346951ff4270dcff3d/specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_addStorageCapacity.json
func ExampleCloudExadataInfrastructuresClient_BeginAddStorageCapacity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudExadataInfrastructuresClient().BeginAddStorageCapacity(ctx, "rg000", "infra1", nil)
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
	// res.CloudExadataInfrastructure = armoracledatabase.CloudExadataInfrastructure{
	// 	Type: to.Ptr("Oracle.Database/cloudExadataInfrastructures"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudExadataInfrastructures/infra1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tagK1": to.Ptr("tagV1"),
	// 	},
	// 	Properties: &armoracledatabase.CloudExadataInfrastructureProperties{
	// 		ActivatedStorageCount: to.Ptr[int32](1),
	// 		AdditionalStorageCount: to.Ptr[int32](1),
	// 		AvailableStorageSizeInGbs: to.Ptr[int32](1000),
	// 		ComputeCount: to.Ptr[int32](100),
	// 		CPUCount: to.Ptr[int32](10),
	// 		CustomerContacts: []*armoracledatabase.CustomerContact{
	// 			{
	// 				Email: to.Ptr("noreply@oracle.com"),
	// 		}},
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DataStorageSizeInTbs: to.Ptr[float64](10),
	// 		DbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 		DbServerVersion: to.Ptr("19.0.0.0"),
	// 		DisplayName: to.Ptr("infra 1"),
	// 		EstimatedPatchingTime: &armoracledatabase.EstimatedPatchingTime{
	// 			EstimatedDbServerPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedNetworkSwitchesPatchingTime: to.Ptr[int32](3000),
	// 			EstimatedStorageServerPatchingTime: to.Ptr[int32](3000),
	// 			TotalEstimatedPatchingTime: to.Ptr[int32](3000),
	// 		},
	// 		LastMaintenanceRunID: to.Ptr("ocid1..aaaaa"),
	// 		LifecycleDetails: to.Ptr("none"),
	// 		MaintenanceWindow: &armoracledatabase.MaintenanceWindow{
	// 			CustomActionTimeoutInMins: to.Ptr[int32](120),
	// 			DaysOfWeek: []*armoracledatabase.DayOfWeek{
	// 				{
	// 					Name: to.Ptr(armoracledatabase.DayOfWeekNameMonday),
	// 			}},
	// 			HoursOfDay: []*int32{
	// 				to.Ptr[int32](0)},
	// 				IsCustomActionTimeoutEnabled: to.Ptr(true),
	// 				IsMonthlyPatchingEnabled: to.Ptr(true),
	// 				LeadTimeInWeeks: to.Ptr[int32](0),
	// 				Months: []*armoracledatabase.Month{
	// 					{
	// 						Name: to.Ptr(armoracledatabase.MonthNameJanuary),
	// 				}},
	// 				PatchingMode: to.Ptr(armoracledatabase.PatchingModeRolling),
	// 				Preference: to.Ptr(armoracledatabase.PreferenceNoPreference),
	// 				WeeksOfMonth: []*int32{
	// 					to.Ptr[int32](0)},
	// 				},
	// 				MaxCPUCount: to.Ptr[int32](100),
	// 				MaxDataStorageInTbs: to.Ptr[float64](1000),
	// 				MaxDbNodeStorageSizeInGbs: to.Ptr[int32](10),
	// 				MaxMemoryInGbs: to.Ptr[int32](1000),
	// 				MemorySizeInGbs: to.Ptr[int32](100),
	// 				MonthlyDbServerVersion: to.Ptr("aaaa"),
	// 				MonthlyStorageServerVersion: to.Ptr("aaaa"),
	// 				NextMaintenanceRunID: to.Ptr("ocid1..aaaaaa"),
	// 				OciURL: to.Ptr("https://url"),
	// 				Ocid: to.Ptr("ocid1..aaaaaa"),
	// 				ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
	// 				Shape: to.Ptr("EXADATA.X9M"),
	// 				StorageCount: to.Ptr[int32](10),
	// 				StorageServerVersion: to.Ptr("0.0"),
	// 				TimeCreated: to.Ptr("2023-02-01T01:01:00"),
	// 				TotalStorageSizeInGbs: to.Ptr[int32](1000),
	// 			},
	// 			Zones: []*string{
	// 				to.Ptr("1")},
	// 			}
}
