//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureIaasVm/ClassicCompute_ProtectedItem_Get.json
func ExampleProtectedItemsClient_Get_getProtectedClassicVirtualMachineDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().Get(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure", "iaasvmcontainer;iaasvmcontainer;iaasvm-rg;iaasvm-1", "vm;iaasvmcontainer;iaasvm-rg;iaasvm-1", &armrecoveryservicesbackup.ProtectedItemsClientGetOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectedItemResource = armrecoveryservicesbackup.ProtectedItemResource{
	// 	Name: to.Ptr("VM;iaasvmcontainer;iaasvm-rg;iaasvm-1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainer;iaasvm-rg;iaasvm-1/protectedItems/VM;iaasvmcontainer;iaasvm-rg;iaasvm-1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSClassicComputeVMProtectedItem{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		ContainerName: to.Ptr("iaasvmcontainer;iaasvm-rg;iaasvm-1"),
	// 		LastRecoveryPoint: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-22T12:25:32.048723Z"); return t}()),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupPolicies/testPolicy1"),
	// 		ProtectedItemType: to.Ptr("Microsoft.ClassicCompute/virtualMachines"),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.ClassicCompute/virtualMachines/iaasvm-1"),
	// 		WorkloadType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	// 		FriendlyName: to.Ptr("iaasvm-1"),
	// 		HealthStatus: to.Ptr(armrecoveryservicesbackup.HealthStatusPassed),
	// 		LastBackupStatus: to.Ptr("Completed"),
	// 		LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-22T12:25:32.048723Z"); return t}()),
	// 		ProtectedItemDataID: to.Ptr("636482643132986882"),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStateProtected),
	// 		ProtectionStatus: to.Ptr("Healthy"),
	// 		VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.ClassicCompute/virtualMachines/iaasvm-1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureIaasVm/Compute_ProtectedItem_Get.json
func ExampleProtectedItemsClient_Get_getProtectedVirtualMachineDetails() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().Get(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure", "iaasvmcontainer;iaasvmcontainerv2;iaasvm-rg;iaasvm-1", "vm;iaasvmcontainerv2;iaasvm-rg;iaasvm-1", &armrecoveryservicesbackup.ProtectedItemsClientGetOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectedItemResource = armrecoveryservicesbackup.ProtectedItemResource{
	// 	Name: to.Ptr("VM;iaasvmcontainerv2;iaasvm-rg;iaasvm-1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;iaasvm-rg;iaasvm-1/protectedItems/VM;iaasvmcontainerv2;iaasvm-rg;iaasvm-1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		ContainerName: to.Ptr("iaasvmcontainerv2;iaasvm-rg;iaasvm-1"),
	// 		LastRecoveryPoint: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-22T12:25:32.048723Z"); return t}()),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupPolicies/testPolicy1"),
	// 		ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 		WorkloadType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	// 		FriendlyName: to.Ptr("iaasvm-1"),
	// 		HealthStatus: to.Ptr(armrecoveryservicesbackup.HealthStatusPassed),
	// 		LastBackupStatus: to.Ptr("Completed"),
	// 		LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-22T12:25:32.048723Z"); return t}()),
	// 		ProtectedItemDataID: to.Ptr("636482643132986882"),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStateProtected),
	// 		ProtectionStatus: to.Ptr("Healthy"),
	// 		VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/providers/Microsoft.Compute/virtualMachines/iaasvm-1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureIaasVm/ConfigureProtection.json
func ExampleProtectedItemsClient_CreateOrUpdate_enableProtectionOnAzureIaasVm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().CreateOrUpdate(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "Azure", "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", armrecoveryservicesbackup.ProtectedItemResource{
		Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
			PolicyID:          to.Ptr("/Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/SwaggerTestRg/providers/Microsoft.RecoveryServices/vaults/NetSDKTestRsVault/backupPolicies/DefaultPolicy"),
			ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
			SourceResourceID:  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectedItemResource = armrecoveryservicesbackup.ProtectedItemResource{
	// 	Name: to.Ptr("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1/protectedItems/VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		ContainerName: to.Ptr("iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupPolicies/testPolicy1"),
	// 		ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
	// 		WorkloadType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	// 		FriendlyName: to.Ptr("netvmtestv2vm1"),
	// 		HealthStatus: to.Ptr(armrecoveryservicesbackup.HealthStatusPassed),
	// 		LastBackupStatus: to.Ptr("Completed"),
	// 		LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-22T12:25:32.048723Z"); return t}()),
	// 		ProtectedItemDataID: to.Ptr("636482643132986882"),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStateProtected),
	// 		ProtectionStatus: to.Ptr("Healthy"),
	// 		VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureIaasVm/StopProtection.json
func ExampleProtectedItemsClient_CreateOrUpdate_stopProtectionWithRetainDataOnAzureIaasVm() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProtectedItemsClient().CreateOrUpdate(ctx, "NetSDKTestRsVault", "SwaggerTestRg", "Azure", "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", armrecoveryservicesbackup.ProtectedItemResource{
		Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
			ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
			SourceResourceID:  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
			ProtectionState:   to.Ptr(armrecoveryservicesbackup.ProtectionStateProtectionStopped),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProtectedItemResource = armrecoveryservicesbackup.ProtectedItemResource{
	// 	Name: to.Ptr("VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/backupFabrics/protectionContainers/protectedItems"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupFabrics/Azure/protectionContainers/IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1/protectedItems/VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 	Properties: &armrecoveryservicesbackup.AzureIaaSComputeVMProtectedItem{
	// 		BackupManagementType: to.Ptr(armrecoveryservicesbackup.BackupManagementTypeAzureIaasVM),
	// 		ContainerName: to.Ptr("iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1"),
	// 		PolicyID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/PythonSDKBackupTestRg/providers/Microsoft.RecoveryServices/vaults/PySDKBackupTestRsVault/backupPolicies/testPolicy1"),
	// 		ProtectedItemType: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
	// 		WorkloadType: to.Ptr(armrecoveryservicesbackup.DataSourceTypeVM),
	// 		FriendlyName: to.Ptr("netvmtestv2vm1"),
	// 		HealthStatus: to.Ptr(armrecoveryservicesbackup.HealthStatusPassed),
	// 		LastBackupStatus: to.Ptr("Completed"),
	// 		LastBackupTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-22T12:25:32.048723Z"); return t}()),
	// 		ProtectedItemDataID: to.Ptr("636482643132986882"),
	// 		ProtectionState: to.Ptr(armrecoveryservicesbackup.ProtectionStateProtectionStopped),
	// 		ProtectionStatus: to.Ptr("Healthy"),
	// 		VirtualMachineID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a498cae6d1a93f4c33073f0747b93b22815c09b7/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/Common/ProtectedItem_Delete.json
func ExampleProtectedItemsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewProtectedItemsClient().Delete(ctx, "PySDKBackupTestRsVault", "PythonSDKBackupTestRg", "Azure", "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
