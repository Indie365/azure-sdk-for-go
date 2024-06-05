//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ItemLevelRecoveryConnectionsClient contains the methods for the ItemLevelRecoveryConnections group.
// Don't use this type directly, use NewItemLevelRecoveryConnectionsClient() instead.
type ItemLevelRecoveryConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewItemLevelRecoveryConnectionsClient creates a new instance of ItemLevelRecoveryConnectionsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewItemLevelRecoveryConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ItemLevelRecoveryConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ItemLevelRecoveryConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Provision - Provisions a script which invokes an iSCSI connection to the backup data. Executing this script opens a file
// explorer displaying all the recoverable files and folders. This is an asynchronous
// operation. To know the status of provisioning, call GetProtectedItemOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the backed up items.
//   - containerName - Container name associated with the backed up items.
//   - protectedItemName - Backed up item name whose files/folders are to be restored.
//   - recoveryPointID - Recovery point ID which represents backed up data. iSCSI connection will be provisioned for this backed
//     up data.
//   - parameters - resource ILR request
//   - options - ItemLevelRecoveryConnectionsClientProvisionOptions contains the optional parameters for the ItemLevelRecoveryConnectionsClient.Provision
//     method.
func (client *ItemLevelRecoveryConnectionsClient) Provision(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters ILRRequestResource, options *ItemLevelRecoveryConnectionsClientProvisionOptions) (ItemLevelRecoveryConnectionsClientProvisionResponse, error) {
	var err error
	const operationName = "ItemLevelRecoveryConnectionsClient.Provision"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.provisionCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters, options)
	if err != nil {
		return ItemLevelRecoveryConnectionsClientProvisionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemLevelRecoveryConnectionsClientProvisionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ItemLevelRecoveryConnectionsClientProvisionResponse{}, err
	}
	return ItemLevelRecoveryConnectionsClientProvisionResponse{}, nil
}

// provisionCreateRequest creates the Provision request.
func (client *ItemLevelRecoveryConnectionsClient) provisionCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters ILRRequestResource, options *ItemLevelRecoveryConnectionsClientProvisionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/provisionInstantItemRecovery"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	if recoveryPointID == "" {
		return nil, errors.New("parameter recoveryPointID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointId}", url.PathEscape(recoveryPointID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Revoke - Revokes an iSCSI connection which can be used to download a script. Executing this script opens a file explorer
// displaying all recoverable files and folders. This is an asynchronous operation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the backed up items.
//   - containerName - Container name associated with the backed up items.
//   - protectedItemName - Backed up item name whose files/folders are to be restored.
//   - recoveryPointID - Recovery point ID which represents backed up data. iSCSI connection will be revoked for this backed up
//     data.
//   - options - ItemLevelRecoveryConnectionsClientRevokeOptions contains the optional parameters for the ItemLevelRecoveryConnectionsClient.Revoke
//     method.
func (client *ItemLevelRecoveryConnectionsClient) Revoke(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, options *ItemLevelRecoveryConnectionsClientRevokeOptions) (ItemLevelRecoveryConnectionsClientRevokeResponse, error) {
	var err error
	const operationName = "ItemLevelRecoveryConnectionsClient.Revoke"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.revokeCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, options)
	if err != nil {
		return ItemLevelRecoveryConnectionsClientRevokeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ItemLevelRecoveryConnectionsClientRevokeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ItemLevelRecoveryConnectionsClientRevokeResponse{}, err
	}
	return ItemLevelRecoveryConnectionsClientRevokeResponse{}, nil
}

// revokeCreateRequest creates the Revoke request.
func (client *ItemLevelRecoveryConnectionsClient) revokeCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, options *ItemLevelRecoveryConnectionsClientRevokeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/revokeInstantItemRecovery"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	if recoveryPointID == "" {
		return nil, errors.New("parameter recoveryPointID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointId}", url.PathEscape(recoveryPointID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
