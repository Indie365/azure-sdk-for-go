//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksListByDevice.json
func ExampleIscsiDisksClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiDisksClient().NewListByDevicePager("HSDK-0NZI14MDTF", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.ISCSIDiskList = armstorsimple1200series.ISCSIDiskList{
		// 	Value: []*armstorsimple1200series.ISCSIDisk{
		// 		{
		// 			Name: to.Ptr("Auto-TestIscsiDisk2"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk2"),
		// 			Properties: &armstorsimple1200series.ISCSIDiskProperties{
		// 				Description: to.Ptr("Demo IscsiDisk for SDK Test Tiered"),
		// 				AccessControlRecords: []*string{
		// 				},
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOnline),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Auto-TestIscsiDisk1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk1"),
		// 			Properties: &armstorsimple1200series.ISCSIDiskProperties{
		// 				Description: to.Ptr("Updated: Demo IscsiDisk for SDK Test Tiered"),
		// 				AccessControlRecords: []*string{
		// 				},
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOffline),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksListByIscsiServer.json
func ExampleIscsiDisksClient_NewListByIscsiServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiDisksClient().NewListByIscsiServerPager("HSDK-0NZI14MDTF", "HSDK-0NZI14MDTF", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.ISCSIDiskList = armstorsimple1200series.ISCSIDiskList{
		// 	Value: []*armstorsimple1200series.ISCSIDisk{
		// 		{
		// 			Name: to.Ptr("Auto-TestIscsiDisk2"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk2"),
		// 			Properties: &armstorsimple1200series.ISCSIDiskProperties{
		// 				Description: to.Ptr("Demo IscsiDisk for SDK Test Tiered"),
		// 				AccessControlRecords: []*string{
		// 				},
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOnline),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Auto-TestIscsiDisk1"),
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
		// 			ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk1"),
		// 			Properties: &armstorsimple1200series.ISCSIDiskProperties{
		// 				Description: to.Ptr("Updated: Demo IscsiDisk for SDK Test Tiered"),
		// 				AccessControlRecords: []*string{
		// 				},
		// 				DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
		// 				DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOffline),
		// 				LocalUsedCapacityInBytes: to.Ptr[int64](0),
		// 				MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
		// 				ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
		// 				UsedCapacityInBytes: to.Ptr[int64](0),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksGet.json
func ExampleIscsiDisksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIscsiDisksClient().Get(ctx, "HSDK-0NZI14MDTF", "HSDK-0NZI14MDTF", "Auto-TestIscsiDisk1", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ISCSIDisk = armstorsimple1200series.ISCSIDisk{
	// 	Name: to.Ptr("Auto-TestIscsiDisk1"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk1"),
	// 	Properties: &armstorsimple1200series.ISCSIDiskProperties{
	// 		Description: to.Ptr("Demo IscsiDisk for SDK Test Tiered"),
	// 		AccessControlRecords: []*string{
	// 		},
	// 		DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
	// 		DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOnline),
	// 		LocalUsedCapacityInBytes: to.Ptr[int64](0),
	// 		MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
	// 		ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
	// 		UsedCapacityInBytes: to.Ptr[int64](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksCreateOrUpdate.json
func ExampleIscsiDisksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIscsiDisksClient().BeginCreateOrUpdate(ctx, "HSDK-0NZI14MDTF", "HSDK-0NZI14MDTF", "Auto-TestIscsiDisk1", "ResourceGroupForSDKTest", "hAzureSDKOperations", armstorsimple1200series.ISCSIDisk{
		Name: to.Ptr("Auto-TestIscsiDisk1"),
		Properties: &armstorsimple1200series.ISCSIDiskProperties{
			Description:                to.Ptr("Demo IscsiDisk for SDK Test Tiered"),
			AccessControlRecords:       []*string{},
			DataPolicy:                 to.Ptr(armstorsimple1200series.DataPolicyTiered),
			DiskStatus:                 to.Ptr(armstorsimple1200series.DiskStatusOnline),
			MonitoringStatus:           to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
			ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
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
	// res.ISCSIDisk = armstorsimple1200series.ISCSIDisk{
	// 	Name: to.Ptr("Auto-TestIscsiDisk1"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks"),
	// 	ID: to.Ptr("/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-0NZI14MDTF/iscsiServers/HSDK-0NZI14MDTF/disks/Auto-TestIscsiDisk1"),
	// 	Properties: &armstorsimple1200series.ISCSIDiskProperties{
	// 		Description: to.Ptr("Demo IscsiDisk for SDK Test Tiered"),
	// 		AccessControlRecords: []*string{
	// 		},
	// 		DataPolicy: to.Ptr(armstorsimple1200series.DataPolicyTiered),
	// 		DiskStatus: to.Ptr(armstorsimple1200series.DiskStatusOnline),
	// 		LocalUsedCapacityInBytes: to.Ptr[int64](0),
	// 		MonitoringStatus: to.Ptr(armstorsimple1200series.MonitoringStatusEnabled),
	// 		ProvisionedCapacityInBytes: to.Ptr[int64](536870912000),
	// 		UsedCapacityInBytes: to.Ptr[int64](0),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksDelete.json
func ExampleIscsiDisksClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIscsiDisksClient().BeginDelete(ctx, "HSDK-UGU4PITWNI", "HSDK-UGU4PITWNI", "ClonedTieredIscsiDiskForSDKTest", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksListMetrics.json
func ExampleIscsiDisksClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiDisksClient().NewListMetricsPager("HSDK-WSJQERQW3F", "HSDK-WSJQERQW3F", "TieredIscsiDiskForSDKTest", "ResourceGroupForSDKTest", "hAzureSDKOperations", &armstorsimple1200series.IscsiDisksClientListMetricsOptions{Filter: to.Ptr("startTime%20ge%20'2018-08-10T18:30:00Z'%20and%20endTime%20le%20'2018-08-11T18:30:00Z'")})
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
		// page.MetricList = armstorsimple1200series.MetricList{
		// 	Value: []*armstorsimple1200series.Metrics{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Storage Used"),
		// 				Value: to.Ptr("HostUsedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks/metrics"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Disk"),
		// 					Value: to.Ptr("TieredIscsiDiskForSDKTest"),
		// 			}},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-11T18:30:00Z"); return t}()),
		// 			PrimaryAggregation: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374/devices/HSDK-WSJQERQW3F/iscsiservers/HSDK-WSJQERQW3F/disks/TieredIscsiDiskForSDKTest"),
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:30:00Z"); return t}()),
		// 			TimeGrain: to.Ptr("1.00:00:00"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 			Values: []*armstorsimple1200series.MetricData{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksListMetricDefinition.json
func ExampleIscsiDisksClient_NewListMetricDefinitionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple1200series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIscsiDisksClient().NewListMetricDefinitionPager("HSDK-WSJQERQW3F", "HSDK-WSJQERQW3F", "TieredIscsiDiskForSDKTest", "ResourceGroupForSDKTest", "hAzureSDKOperations", nil)
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
		// page.MetricDefinitionList = armstorsimple1200series.MetricDefinitionList{
		// 	Value: []*armstorsimple1200series.MetricDefinition{
		// 		{
		// 			Name: &armstorsimple1200series.MetricName{
		// 				LocalizedValue: to.Ptr("Primary Storage Used"),
		// 				Value: to.Ptr("HostUsedStorage"),
		// 			},
		// 			Type: to.Ptr("Microsoft.StorSimple/managers/devices/iscsiServers/disks/metricsDefinitions"),
		// 			Dimensions: []*armstorsimple1200series.MetricDimension{
		// 				{
		// 					Name: to.Ptr("Disk"),
		// 					Value: to.Ptr("TieredIscsiDiskForSDKTest"),
		// 			}},
		// 			MetricAvailabilities: []*armstorsimple1200series.MetricAvailablity{
		// 				{
		// 					Retention: to.Ptr("PT6H"),
		// 					TimeGrain: to.Ptr("PT1M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P7D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P1Y"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armstorsimple1200series.MetricAggregationTypeAverage),
		// 			ResourceID: to.Ptr("https://pod01-cis2.sea.storsimple.windowsazure.com/managers/4239154091695873374/devices/HSDK-WSJQERQW3F/iscsiservers/HSDK-WSJQERQW3F/disks/TieredIscsiDiskForSDKTest"),
		// 			Unit: to.Ptr(armstorsimple1200series.MetricUnitBytes),
		// 	}},
		// }
	}
}
