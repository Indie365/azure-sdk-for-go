//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

// ArcSettingsClientCreateIdentityResponse contains the response from method ArcSettingsClient.BeginCreateIdentity.
type ArcSettingsClientCreateIdentityResponse struct {
	// ArcIdentity details.
	ArcIdentityResponse
}

// ArcSettingsClientCreateResponse contains the response from method ArcSettingsClient.Create.
type ArcSettingsClientCreateResponse struct {
	// ArcSetting details.
	ArcSetting
}

// ArcSettingsClientDeleteResponse contains the response from method ArcSettingsClient.BeginDelete.
type ArcSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ArcSettingsClientGeneratePasswordResponse contains the response from method ArcSettingsClient.GeneratePassword.
type ArcSettingsClientGeneratePasswordResponse struct {
	PasswordCredential
}

// ArcSettingsClientGetResponse contains the response from method ArcSettingsClient.Get.
type ArcSettingsClientGetResponse struct {
	// ArcSetting details.
	ArcSetting
}

// ArcSettingsClientListByClusterResponse contains the response from method ArcSettingsClient.NewListByClusterPager.
type ArcSettingsClientListByClusterResponse struct {
	// List of ArcSetting proxy resources for the HCI cluster.
	ArcSettingList
}

// ArcSettingsClientUpdateResponse contains the response from method ArcSettingsClient.Update.
type ArcSettingsClientUpdateResponse struct {
	// ArcSetting details.
	ArcSetting
}

// ClustersClientCreateIdentityResponse contains the response from method ClustersClient.BeginCreateIdentity.
type ClustersClientCreateIdentityResponse struct {
	// Cluster Identity details.
	ClusterIdentityResponse
}

// ClustersClientCreateResponse contains the response from method ClustersClient.Create.
type ClustersClientCreateResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.BeginDelete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.NewListByResourceGroupPager.
type ClustersClientListByResourceGroupResponse struct {
	// List of clusters.
	ClusterList
}

// ClustersClientListBySubscriptionResponse contains the response from method ClustersClient.NewListBySubscriptionPager.
type ClustersClientListBySubscriptionResponse struct {
	// List of clusters.
	ClusterList
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.Update.
type ClustersClientUpdateResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientUploadCertificateResponse contains the response from method ClustersClient.BeginUploadCertificate.
type ClustersClientUploadCertificateResponse struct {
	// placeholder for future response values
}

// ExtensionsClientCreateResponse contains the response from method ExtensionsClient.BeginCreate.
type ExtensionsClientCreateResponse struct {
	// Details of a particular extension in HCI Cluster.
	Extension
}

// ExtensionsClientDeleteResponse contains the response from method ExtensionsClient.BeginDelete.
type ExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExtensionsClientGetResponse contains the response from method ExtensionsClient.Get.
type ExtensionsClientGetResponse struct {
	// Details of a particular extension in HCI Cluster.
	Extension
}

// ExtensionsClientListByArcSettingResponse contains the response from method ExtensionsClient.NewListByArcSettingPager.
type ExtensionsClientListByArcSettingResponse struct {
	// List of Extensions in HCI cluster.
	ExtensionList
}

// ExtensionsClientUpdateResponse contains the response from method ExtensionsClient.BeginUpdate.
type ExtensionsClientUpdateResponse struct {
	// Details of a particular extension in HCI Cluster.
	Extension
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

