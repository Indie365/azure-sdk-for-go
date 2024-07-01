//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	internal       *arm.Client
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	internal, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID,
		internal:       internal,
	}, nil
}

// NewBackupInstancesClient creates a new instance of BackupInstancesClient.
func (c *ClientFactory) NewBackupInstancesClient() *BackupInstancesClient {
	return &BackupInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupInstancesExtensionRoutingClient creates a new instance of BackupInstancesExtensionRoutingClient.
func (c *ClientFactory) NewBackupInstancesExtensionRoutingClient() *BackupInstancesExtensionRoutingClient {
	return &BackupInstancesExtensionRoutingClient{
		internal: c.internal,
	}
}

// NewBackupPoliciesClient creates a new instance of BackupPoliciesClient.
func (c *ClientFactory) NewBackupPoliciesClient() *BackupPoliciesClient {
	return &BackupPoliciesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupVaultOperationResultsClient creates a new instance of BackupVaultOperationResultsClient.
func (c *ClientFactory) NewBackupVaultOperationResultsClient() *BackupVaultOperationResultsClient {
	return &BackupVaultOperationResultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewBackupVaultsClient creates a new instance of BackupVaultsClient.
func (c *ClientFactory) NewBackupVaultsClient() *BackupVaultsClient {
	return &BackupVaultsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewClient creates a new instance of Client.
func (c *ClientFactory) NewClient() *Client {
	return &Client{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDeletedBackupInstancesClient creates a new instance of DeletedBackupInstancesClient.
func (c *ClientFactory) NewDeletedBackupInstancesClient() *DeletedBackupInstancesClient {
	return &DeletedBackupInstancesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewDppResourceGuardProxyClient creates a new instance of DppResourceGuardProxyClient.
func (c *ClientFactory) NewDppResourceGuardProxyClient() *DppResourceGuardProxyClient {
	return &DppResourceGuardProxyClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExportJobsClient creates a new instance of ExportJobsClient.
func (c *ClientFactory) NewExportJobsClient() *ExportJobsClient {
	return &ExportJobsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewExportJobsOperationResultClient creates a new instance of ExportJobsOperationResultClient.
func (c *ClientFactory) NewExportJobsOperationResultClient() *ExportJobsOperationResultClient {
	return &ExportJobsOperationResultClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFetchCrossRegionRestoreJobClient creates a new instance of FetchCrossRegionRestoreJobClient.
func (c *ClientFactory) NewFetchCrossRegionRestoreJobClient() *FetchCrossRegionRestoreJobClient {
	return &FetchCrossRegionRestoreJobClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFetchCrossRegionRestoreJobsClient creates a new instance of FetchCrossRegionRestoreJobsClient.
func (c *ClientFactory) NewFetchCrossRegionRestoreJobsClient() *FetchCrossRegionRestoreJobsClient {
	return &FetchCrossRegionRestoreJobsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewFetchSecondaryRecoveryPointsClient creates a new instance of FetchSecondaryRecoveryPointsClient.
func (c *ClientFactory) NewFetchSecondaryRecoveryPointsClient() *FetchSecondaryRecoveryPointsClient {
	return &FetchSecondaryRecoveryPointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewJobsClient creates a new instance of JobsClient.
func (c *ClientFactory) NewJobsClient() *JobsClient {
	return &JobsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationResultClient creates a new instance of OperationResultClient.
func (c *ClientFactory) NewOperationResultClient() *OperationResultClient {
	return &OperationResultClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationStatusBackupVaultContextClient creates a new instance of OperationStatusBackupVaultContextClient.
func (c *ClientFactory) NewOperationStatusBackupVaultContextClient() *OperationStatusBackupVaultContextClient {
	return &OperationStatusBackupVaultContextClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationStatusClient creates a new instance of OperationStatusClient.
func (c *ClientFactory) NewOperationStatusClient() *OperationStatusClient {
	return &OperationStatusClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationStatusResourceGroupContextClient creates a new instance of OperationStatusResourceGroupContextClient.
func (c *ClientFactory) NewOperationStatusResourceGroupContextClient() *OperationStatusResourceGroupContextClient {
	return &OperationStatusResourceGroupContextClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal: c.internal,
	}
}

// NewRecoveryPointsClient creates a new instance of RecoveryPointsClient.
func (c *ClientFactory) NewRecoveryPointsClient() *RecoveryPointsClient {
	return &RecoveryPointsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewResourceGuardsClient creates a new instance of ResourceGuardsClient.
func (c *ClientFactory) NewResourceGuardsClient() *ResourceGuardsClient {
	return &ResourceGuardsClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}

// NewRestorableTimeRangesClient creates a new instance of RestorableTimeRangesClient.
func (c *ClientFactory) NewRestorableTimeRangesClient() *RestorableTimeRangesClient {
	return &RestorableTimeRangesClient{
		subscriptionID: c.subscriptionID,
		internal:       c.internal,
	}
}
