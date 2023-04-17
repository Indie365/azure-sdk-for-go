//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Create.json
func ExampleServerEndpointsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerEndpointsClient().BeginCreate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", armstoragesync.ServerEndpointCreateParameters{
		Properties: &armstoragesync.ServerEndpointCreateParametersProperties{
			CloudTiering:                 to.Ptr(armstoragesync.FeatureStatusOff),
			InitialDownloadPolicy:        to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
			InitialUploadPolicy:          to.Ptr(armstoragesync.InitialUploadPolicyServerAuthoritative),
			LocalCacheMode:               to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
			OfflineDataTransfer:          to.Ptr(armstoragesync.FeatureStatusOn),
			OfflineDataTransferShareName: to.Ptr("myfileshare"),
			ServerLocalPath:              to.Ptr("D:\\SampleServerEndpoint_1"),
			ServerResourceID:             to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
			TierFilesOlderThanDays:       to.Ptr[int32](0),
			VolumeFreeSpacePercent:       to.Ptr[int32](100),
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
	// res.ServerEndpoint = armstoragesync.ServerEndpoint{
	// 	Name: to.Ptr("SampleServerEndpoint_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
	// 	ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_1"),
	// 	Properties: &armstoragesync.ServerEndpointProperties{
	// 		CloudTiering: to.Ptr(armstoragesync.FeatureStatusOff),
	// 		FriendlyName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 		InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
	// 		InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyServerAuthoritative),
	// 		LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
	// 		LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/569afb5c-7172-4cf8-a8ea-d987f727f11a"),
	// 		LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
	// 		OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOn),
	// 		OfflineDataTransferShareName: to.Ptr("myfileshare"),
	// 		OfflineDataTransferStorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/myRG/providers/Microsoft.Storage/storageAccounts/mysa"),
	// 		OfflineDataTransferStorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServerLocalPath: to.Ptr("D:\\SampleServerEndpoint_1"),
	// 		ServerName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 		ServerResourceID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
	// 		TierFilesOlderThanDays: to.Ptr[int32](0),
	// 		VolumeFreeSpacePercent: to.Ptr[int32](100),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Update.json
func ExampleServerEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerEndpointsClient().BeginUpdate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", &armstoragesync.ServerEndpointsClientBeginUpdateOptions{Parameters: &armstoragesync.ServerEndpointUpdateParameters{
		Properties: &armstoragesync.ServerEndpointUpdateProperties{
			CloudTiering:           to.Ptr(armstoragesync.FeatureStatusOff),
			LocalCacheMode:         to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
			OfflineDataTransfer:    to.Ptr(armstoragesync.FeatureStatusOff),
			TierFilesOlderThanDays: to.Ptr[int32](0),
			VolumeFreeSpacePercent: to.Ptr[int32](100),
		},
	},
	})
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
	// res.ServerEndpoint = armstoragesync.ServerEndpoint{
	// 	Name: to.Ptr("SampleServerEndpoint_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
	// 	ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_1"),
	// 	Properties: &armstoragesync.ServerEndpointProperties{
	// 		CloudTiering: to.Ptr(armstoragesync.FeatureStatusOff),
	// 		FriendlyName: to.Ptr("somemachine2.redmond.corp.microsoft.com"),
	// 		InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
	// 		InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyMerge),
	// 		LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
	// 		LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/569afb5c-7172-4cf8-a8ea-d987f727f11b"),
	// 		LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
	// 		OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOff),
	// 		OfflineDataTransferShareName: to.Ptr("myfileshare"),
	// 		OfflineDataTransferStorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/myRG/providers/Microsoft.Storage/storageAccounts/mysa"),
	// 		OfflineDataTransferStorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServerLocalPath: to.Ptr("C:\\data_quota2"),
	// 		ServerName: to.Ptr("somemachine2.redmond.corp.microsoft.com"),
	// 		ServerResourceID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/anpintDemoRG/providers/kailanitest07.one.microsoft.com/storageSyncServices/healthdemo1/registeredServers/f94e2944-b48d-4e5b-bdc7-c48ab3712659"),
	// 		SyncStatus: &armstoragesync.ServerEndpointSyncStatus{
	// 			CombinedHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 			DownloadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
	// 			DownloadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 				LastSyncPerItemErrorCount: to.Ptr[int64](0),
	// 				LastSyncResult: to.Ptr[int32](0),
	// 				LastSyncSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
	// 				LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
	// 			},
	// 			LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			OfflineDataTransferStatus: to.Ptr(armstoragesync.ServerEndpointOfflineDataTransferStateStopping),
	// 			UploadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 			UploadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 				LastSyncPerItemErrorCount: to.Ptr[int64](0),
	// 				LastSyncResult: to.Ptr[int32](-2134351810),
	// 				LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			},
	// 		},
	// 		TierFilesOlderThanDays: to.Ptr[int32](0),
	// 		VolumeFreeSpacePercent: to.Ptr[int32](20),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Get.json
func ExampleServerEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerEndpointsClient().Get(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerEndpoint = armstoragesync.ServerEndpoint{
	// 	Name: to.Ptr("SampleServerEndpoint_1"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
	// 	ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_1"),
	// 	Properties: &armstoragesync.ServerEndpointProperties{
	// 		CloudTiering: to.Ptr(armstoragesync.FeatureStatusOn),
	// 		CloudTieringStatus: &armstoragesync.ServerEndpointCloudTieringStatus{
	// 			CachePerformance: &armstoragesync.CloudTieringCachePerformance{
	// 				CacheHitBytes: to.Ptr[int64](922337203685477600),
	// 				CacheHitBytesPercent: to.Ptr[int32](50),
	// 				CacheMissBytes: to.Ptr[int64](922337203685477600),
	// 				LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
	// 			},
	// 			DatePolicyStatus: &armstoragesync.CloudTieringDatePolicyStatus{
	// 				LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
	// 				TieredFilesMostRecentAccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
	// 			},
	// 			FilesNotTiering: &armstoragesync.CloudTieringFilesNotTiering{
	// 				Errors: []*armstoragesync.FilesNotTieringError{
	// 					{
	// 						ErrorCode: to.Ptr[int32](-2134347771),
	// 						FileCount: to.Ptr[int64](10),
	// 					},
	// 					{
	// 						ErrorCode: to.Ptr[int32](-2134347770),
	// 						FileCount: to.Ptr[int64](20),
	// 					},
	// 					{
	// 						ErrorCode: to.Ptr[int32](-2134347769),
	// 						FileCount: to.Ptr[int64](30),
	// 				}},
	// 				LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
	// 				TotalFileCount: to.Ptr[int64](60),
	// 			},
	// 			Health: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 			HealthLastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			LastCloudTieringResult: to.Ptr[int32](-2134347771),
	// 			LastSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			SpaceSavings: &armstoragesync.CloudTieringSpaceSavings{
	// 				CachedSizeBytes: to.Ptr[int64](50000000),
	// 				LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
	// 				SpaceSavingsBytes: to.Ptr[int64](50000000),
	// 				SpaceSavingsPercent: to.Ptr[int32](50),
	// 				TotalSizeCloudBytes: to.Ptr[int64](100000000),
	// 				VolumeSizeBytes: to.Ptr[int64](922337203685477600),
	// 			},
	// 			VolumeFreeSpacePolicyStatus: &armstoragesync.CloudTieringVolumeFreeSpacePolicyStatus{
	// 				CurrentVolumeFreeSpacePercent: to.Ptr[int32](5),
	// 				EffectiveVolumeFreeSpacePolicy: to.Ptr[int32](30),
	// 				LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
	// 			},
	// 		},
	// 		FriendlyName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 		InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
	// 		InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyMerge),
	// 		LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
	// 		LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/569afb5c-7172-4cf8-a8ea-d987f727f11a"),
	// 		LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
	// 		OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOn),
	// 		OfflineDataTransferShareName: to.Ptr("myfileshare"),
	// 		OfflineDataTransferStorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/myRG/providers/Microsoft.Storage/storageAccounts/mysa"),
	// 		OfflineDataTransferStorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RecallStatus: &armstoragesync.ServerEndpointRecallStatus{
	// 			LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			RecallErrors: []*armstoragesync.ServerEndpointRecallError{
	// 				{
	// 					Count: to.Ptr[int64](5),
	// 					ErrorCode: to.Ptr[int32](-2134347775),
	// 				},
	// 				{
	// 					Count: to.Ptr[int64](10),
	// 					ErrorCode: to.Ptr[int32](-2134347774),
	// 				},
	// 				{
	// 					Count: to.Ptr[int64](15),
	// 					ErrorCode: to.Ptr[int32](-2134347773),
	// 			}},
	// 			TotalRecallErrorsCount: to.Ptr[int64](30),
	// 		},
	// 		ServerLocalPath: to.Ptr("D:\\SampleServerEndpoint_1"),
	// 		ServerName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
	// 		ServerResourceID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
	// 		SyncStatus: &armstoragesync.ServerEndpointSyncStatus{
	// 			CombinedHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 			DownloadActivity: &armstoragesync.ServerEndpointSyncActivityStatus{
	// 				AppliedBytes: to.Ptr[int64](94805587),
	// 				AppliedItemCount: to.Ptr[int64](100),
	// 				PerItemErrorCount: to.Ptr[int64](0),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 				TotalBytes: to.Ptr[int64](19583674),
	// 				TotalItemCount: to.Ptr[int64](300),
	// 			},
	// 			DownloadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
	// 			DownloadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 				LastSyncPerItemErrorCount: to.Ptr[int64](0),
	// 				LastSyncResult: to.Ptr[int32](0),
	// 				LastSyncSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
	// 				LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
	// 			},
	// 			LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			OfflineDataTransferStatus: to.Ptr(armstoragesync.ServerEndpointOfflineDataTransferStateInProgress),
	// 			SyncActivity: to.Ptr(armstoragesync.ServerEndpointSyncActivityStateUploadAndDownload),
	// 			UploadActivity: &armstoragesync.ServerEndpointSyncActivityStatus{
	// 				AppliedBytes: to.Ptr[int64](57348983),
	// 				AppliedItemCount: to.Ptr[int64](1000),
	// 				PerItemErrorCount: to.Ptr[int64](300),
	// 				Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 				TotalBytes: to.Ptr[int64](1958367412),
	// 				TotalItemCount: to.Ptr[int64](2300),
	// 			},
	// 			UploadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
	// 			UploadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
	// 				LastSyncPerItemErrorCount: to.Ptr[int64](1000),
	// 				LastSyncResult: to.Ptr[int32](-2134351810),
	// 				LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
	// 			},
	// 		},
	// 		TierFilesOlderThanDays: to.Ptr[int32](0),
	// 		VolumeFreeSpacePercent: to.Ptr[int32](100),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Delete.json
func ExampleServerEndpointsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerEndpointsClient().BeginDelete(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_ListBySyncGroup.json
func ExampleServerEndpointsClient_NewListBySyncGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerEndpointsClient().NewListBySyncGroupPager("SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", nil)
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
		// page.ServerEndpointArray = armstoragesync.ServerEndpointArray{
		// 	Value: []*armstoragesync.ServerEndpoint{
		// 		{
		// 			Name: to.Ptr("SampleServerEndpoint_1"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
		// 			ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_1"),
		// 			Properties: &armstoragesync.ServerEndpointProperties{
		// 				CloudTiering: to.Ptr(armstoragesync.FeatureStatusOff),
		// 				FriendlyName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
		// 				InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
		// 				InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyMerge),
		// 				LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
		// 				LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/569afb5c-7172-4cf8-a8ea-d987f727f11a"),
		// 				LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
		// 				OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOn),
		// 				OfflineDataTransferShareName: to.Ptr("myfileshare"),
		// 				OfflineDataTransferStorageAccountResourceID: to.Ptr("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/myRG/providers/Microsoft.Storage/storageAccounts/mysa"),
		// 				OfflineDataTransferStorageAccountTenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServerLocalPath: to.Ptr("D:\\SampleServerEndpoint_1"),
		// 				ServerName: to.Ptr("somemachine.redmond.corp.microsoft.com"),
		// 				ServerResourceID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
		// 				SyncStatus: &armstoragesync.ServerEndpointSyncStatus{
		// 					CombinedHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
		// 					DownloadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
		// 					DownloadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
		// 						LastSyncPerItemErrorCount: to.Ptr[int64](1000),
		// 						LastSyncResult: to.Ptr[int32](0),
		// 						LastSyncSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
		// 						LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
		// 					},
		// 					LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					OfflineDataTransferStatus: to.Ptr(armstoragesync.ServerEndpointOfflineDataTransferStateComplete),
		// 					UploadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
		// 					UploadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
		// 						LastSyncPerItemErrorCount: to.Ptr[int64](0),
		// 						LastSyncResult: to.Ptr[int32](0),
		// 						LastSyncSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 						LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					},
		// 				},
		// 				TierFilesOlderThanDays: to.Ptr[int32](0),
		// 				VolumeFreeSpacePercent: to.Ptr[int32](100),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SampleServerEndpoint_2"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/syncGroups/serverEndpoints"),
		// 			ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/syncGroups/SampleSyncGroup_1/serverEndpoints/SampleServerEndpoint_2"),
		// 			Properties: &armstoragesync.ServerEndpointProperties{
		// 				CloudTiering: to.Ptr(armstoragesync.FeatureStatusOn),
		// 				CloudTieringStatus: &armstoragesync.ServerEndpointCloudTieringStatus{
		// 					CachePerformance: &armstoragesync.CloudTieringCachePerformance{
		// 						CacheHitBytes: to.Ptr[int64](922337203685477600),
		// 						CacheHitBytesPercent: to.Ptr[int32](50),
		// 						CacheMissBytes: to.Ptr[int64](922337203685477600),
		// 						LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
		// 					},
		// 					DatePolicyStatus: &armstoragesync.CloudTieringDatePolicyStatus{
		// 						LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
		// 						TieredFilesMostRecentAccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
		// 					},
		// 					FilesNotTiering: &armstoragesync.CloudTieringFilesNotTiering{
		// 						Errors: []*armstoragesync.FilesNotTieringError{
		// 							{
		// 								ErrorCode: to.Ptr[int32](-2134347771),
		// 								FileCount: to.Ptr[int64](10),
		// 							},
		// 							{
		// 								ErrorCode: to.Ptr[int32](-2134347770),
		// 								FileCount: to.Ptr[int64](20),
		// 							},
		// 							{
		// 								ErrorCode: to.Ptr[int32](-2134347769),
		// 								FileCount: to.Ptr[int64](30),
		// 						}},
		// 						LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
		// 						TotalFileCount: to.Ptr[int64](60),
		// 					},
		// 					Health: to.Ptr(armstoragesync.ServerEndpointHealthStateError),
		// 					HealthLastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					LastCloudTieringResult: to.Ptr[int32](-2134347771),
		// 					LastSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					SpaceSavings: &armstoragesync.CloudTieringSpaceSavings{
		// 						CachedSizeBytes: to.Ptr[int64](50000000),
		// 						LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
		// 						SpaceSavingsBytes: to.Ptr[int64](50000000),
		// 						SpaceSavingsPercent: to.Ptr[int32](50),
		// 						TotalSizeCloudBytes: to.Ptr[int64](100000000),
		// 						VolumeSizeBytes: to.Ptr[int64](922337203685477600),
		// 					},
		// 					VolumeFreeSpacePolicyStatus: &armstoragesync.CloudTieringVolumeFreeSpacePolicyStatus{
		// 						CurrentVolumeFreeSpacePercent: to.Ptr[int32](5),
		// 						EffectiveVolumeFreeSpacePolicy: to.Ptr[int32](30),
		// 						LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-17T19:04:59.1959227Z"); return t}()),
		// 					},
		// 				},
		// 				FriendlyName: to.Ptr("somemachine2.redmond.corp.microsoft.com"),
		// 				InitialDownloadPolicy: to.Ptr(armstoragesync.InitialDownloadPolicyNamespaceThenModifiedFiles),
		// 				InitialUploadPolicy: to.Ptr(armstoragesync.InitialUploadPolicyMerge),
		// 				LastOperationName: to.Ptr("ICreateServerEndpointWorkflow"),
		// 				LastWorkflowID: to.Ptr("storageSyncServices/healthDemo1/workflows/40b1dc00-d7d9-4721-a1e9-ab60139b830a"),
		// 				LocalCacheMode: to.Ptr(armstoragesync.LocalCacheModeUpdateLocallyCachedFiles),
		// 				OfflineDataTransfer: to.Ptr(armstoragesync.FeatureStatusOff),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RecallStatus: &armstoragesync.ServerEndpointRecallStatus{
		// 					LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					RecallErrors: []*armstoragesync.ServerEndpointRecallError{
		// 						{
		// 							Count: to.Ptr[int64](5),
		// 							ErrorCode: to.Ptr[int32](-2134347775),
		// 						},
		// 						{
		// 							Count: to.Ptr[int64](10),
		// 							ErrorCode: to.Ptr[int32](-2134347774),
		// 						},
		// 						{
		// 							Count: to.Ptr[int64](15),
		// 							ErrorCode: to.Ptr[int32](-2134347773),
		// 					}},
		// 					TotalRecallErrorsCount: to.Ptr[int64](30),
		// 				},
		// 				ServerLocalPath: to.Ptr("D:\\SampleServerEndpoint_2"),
		// 				ServerName: to.Ptr("somemachine2.redmond.corp.microsoft.com"),
		// 				ServerResourceID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/f94e2944-b48d-4e5b-bdc7-c48ab3712659"),
		// 				SyncStatus: &armstoragesync.ServerEndpointSyncStatus{
		// 					CombinedHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
		// 					DownloadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
		// 					DownloadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
		// 						LastSyncPerItemErrorCount: to.Ptr[int64](0),
		// 						LastSyncResult: to.Ptr[int32](0),
		// 						LastSyncSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
		// 						LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
		// 					},
		// 					LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					OfflineDataTransferStatus: to.Ptr(armstoragesync.ServerEndpointOfflineDataTransferStateNotRunning),
		// 					SyncActivity: to.Ptr(armstoragesync.ServerEndpointSyncActivityStateUpload),
		// 					UploadActivity: &armstoragesync.ServerEndpointSyncActivityStatus{
		// 						AppliedBytes: to.Ptr[int64](57348983),
		// 						AppliedItemCount: to.Ptr[int64](1000),
		// 						PerItemErrorCount: to.Ptr[int64](300),
		// 						Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 						TotalBytes: to.Ptr[int64](1958367412),
		// 						TotalItemCount: to.Ptr[int64](2300),
		// 					},
		// 					UploadHealth: to.Ptr(armstoragesync.ServerEndpointHealthStateHealthy),
		// 					UploadStatus: &armstoragesync.ServerEndpointSyncSessionStatus{
		// 						LastSyncPerItemErrorCount: to.Ptr[int64](0),
		// 						LastSyncResult: to.Ptr[int32](0),
		// 						LastSyncSuccessTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:28:33.9217334Z"); return t}()),
		// 						LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T23:32:51.1057915Z"); return t}()),
		// 					},
		// 				},
		// 				TierFilesOlderThanDays: to.Ptr[int32](5),
		// 				VolumeFreeSpacePercent: to.Ptr[int32](80),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Recall.json
func ExampleServerEndpointsClient_BeginRecallAction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerEndpointsClient().BeginRecallAction(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "SampleSyncGroup_1", "SampleServerEndpoint_1", armstoragesync.RecallActionParameters{
		Pattern:    to.Ptr(""),
		RecallPath: to.Ptr(""),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
