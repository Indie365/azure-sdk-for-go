//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchangeanalysis

// ChangesClientListChangesByResourceGroupResponse contains the response from method ChangesClient.NewListChangesByResourceGroupPager.
type ChangesClientListChangesByResourceGroupResponse struct {
	ChangeList
}

// ChangesClientListChangesBySubscriptionResponse contains the response from method ChangesClient.NewListChangesBySubscriptionPager.
type ChangesClientListChangesBySubscriptionResponse struct {
	ChangeList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	ResourceProviderOperationList
}

// ResourceChangesClientListResponse contains the response from method ResourceChangesClient.NewListPager.
type ResourceChangesClientListResponse struct {
	ChangeList
}
