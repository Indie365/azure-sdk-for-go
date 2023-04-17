//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armedgeorder

// ManagementClientCancelOrderItemResponse contains the response from method ManagementClient.CancelOrderItem.
type ManagementClientCancelOrderItemResponse struct {
	// placeholder for future response values
}

// ManagementClientCreateAddressResponse contains the response from method ManagementClient.BeginCreateAddress.
type ManagementClientCreateAddressResponse struct {
	AddressResource
}

// ManagementClientCreateOrderItemResponse contains the response from method ManagementClient.BeginCreateOrderItem.
type ManagementClientCreateOrderItemResponse struct {
	OrderItemResource
}

// ManagementClientDeleteAddressByNameResponse contains the response from method ManagementClient.BeginDeleteAddressByName.
type ManagementClientDeleteAddressByNameResponse struct {
	// placeholder for future response values
}

// ManagementClientDeleteOrderItemByNameResponse contains the response from method ManagementClient.BeginDeleteOrderItemByName.
type ManagementClientDeleteOrderItemByNameResponse struct {
	// placeholder for future response values
}

// ManagementClientGetAddressByNameResponse contains the response from method ManagementClient.GetAddressByName.
type ManagementClientGetAddressByNameResponse struct {
	AddressResource
}

// ManagementClientGetOrderByNameResponse contains the response from method ManagementClient.GetOrderByName.
type ManagementClientGetOrderByNameResponse struct {
	OrderResource
}

// ManagementClientGetOrderItemByNameResponse contains the response from method ManagementClient.GetOrderItemByName.
type ManagementClientGetOrderItemByNameResponse struct {
	OrderItemResource
}

// ManagementClientListAddressesAtResourceGroupLevelResponse contains the response from method ManagementClient.NewListAddressesAtResourceGroupLevelPager.
type ManagementClientListAddressesAtResourceGroupLevelResponse struct {
	AddressResourceList
}

// ManagementClientListAddressesAtSubscriptionLevelResponse contains the response from method ManagementClient.NewListAddressesAtSubscriptionLevelPager.
type ManagementClientListAddressesAtSubscriptionLevelResponse struct {
	AddressResourceList
}

// ManagementClientListConfigurationsResponse contains the response from method ManagementClient.NewListConfigurationsPager.
type ManagementClientListConfigurationsResponse struct {
	Configurations
}

// ManagementClientListOperationsResponse contains the response from method ManagementClient.NewListOperationsPager.
type ManagementClientListOperationsResponse struct {
	OperationListResult
}

// ManagementClientListOrderAtResourceGroupLevelResponse contains the response from method ManagementClient.NewListOrderAtResourceGroupLevelPager.
type ManagementClientListOrderAtResourceGroupLevelResponse struct {
	OrderResourceList
}

// ManagementClientListOrderAtSubscriptionLevelResponse contains the response from method ManagementClient.NewListOrderAtSubscriptionLevelPager.
type ManagementClientListOrderAtSubscriptionLevelResponse struct {
	OrderResourceList
}

// ManagementClientListOrderItemsAtResourceGroupLevelResponse contains the response from method ManagementClient.NewListOrderItemsAtResourceGroupLevelPager.
type ManagementClientListOrderItemsAtResourceGroupLevelResponse struct {
	OrderItemResourceList
}

// ManagementClientListOrderItemsAtSubscriptionLevelResponse contains the response from method ManagementClient.NewListOrderItemsAtSubscriptionLevelPager.
type ManagementClientListOrderItemsAtSubscriptionLevelResponse struct {
	OrderItemResourceList
}

// ManagementClientListProductFamiliesMetadataResponse contains the response from method ManagementClient.NewListProductFamiliesMetadataPager.
type ManagementClientListProductFamiliesMetadataResponse struct {
	ProductFamiliesMetadata
}

// ManagementClientListProductFamiliesResponse contains the response from method ManagementClient.NewListProductFamiliesPager.
type ManagementClientListProductFamiliesResponse struct {
	ProductFamilies
}

// ManagementClientReturnOrderItemResponse contains the response from method ManagementClient.BeginReturnOrderItem.
type ManagementClientReturnOrderItemResponse struct {
	// placeholder for future response values
}

// ManagementClientUpdateAddressResponse contains the response from method ManagementClient.BeginUpdateAddress.
type ManagementClientUpdateAddressResponse struct {
	AddressResource
}

// ManagementClientUpdateOrderItemResponse contains the response from method ManagementClient.BeginUpdateOrderItem.
type ManagementClientUpdateOrderItemResponse struct {
	OrderItemResource
}
