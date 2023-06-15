//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/RouteFilterRuleDelete.json
func ExampleRouteFilterRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRouteFilterRulesClient().BeginDelete(ctx, "rg1", "filterName", "ruleName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/RouteFilterRuleGet.json
func ExampleRouteFilterRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRouteFilterRulesClient().Get(ctx, "rg1", "filterName", "filterName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RouteFilterRule = armnetwork.RouteFilterRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName/routeFilterRules/ruleName"),
	// 	Name: to.Ptr("ruleName"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.RouteFilterRulePropertiesFormat{
	// 		Access: to.Ptr(armnetwork.AccessAllow),
	// 		Communities: []*string{
	// 			to.Ptr("12076:5030"),
	// 			to.Ptr("12076:5040")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/RouteFilterRuleCreate.json
func ExampleRouteFilterRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRouteFilterRulesClient().BeginCreateOrUpdate(ctx, "rg1", "filterName", "ruleName", armnetwork.RouteFilterRule{
		Properties: &armnetwork.RouteFilterRulePropertiesFormat{
			Access: to.Ptr(armnetwork.AccessAllow),
			Communities: []*string{
				to.Ptr("12076:5030"),
				to.Ptr("12076:5040")},
			RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
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
	// res.RouteFilterRule = armnetwork.RouteFilterRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName/routeFilterRules/ruleName"),
	// 	Name: to.Ptr("ruleName"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.RouteFilterRulePropertiesFormat{
	// 		Access: to.Ptr(armnetwork.AccessAllow),
	// 		Communities: []*string{
	// 			to.Ptr("12076:5030"),
	// 			to.Ptr("12076:5040")},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 			RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/RouteFilterRuleListByRouteFilter.json
func ExampleRouteFilterRulesClient_NewListByRouteFilterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRouteFilterRulesClient().NewListByRouteFilterPager("rg1", "filterName", nil)
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
		// page.RouteFilterRuleListResult = armnetwork.RouteFilterRuleListResult{
		// 	Value: []*armnetwork.RouteFilterRule{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/routeFilters/filterName/routeFilterRules/ruleName"),
		// 			Name: to.Ptr("ruleName"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.RouteFilterRulePropertiesFormat{
		// 				Access: to.Ptr(armnetwork.AccessAllow),
		// 				Communities: []*string{
		// 					to.Ptr("12076:5030"),
		// 					to.Ptr("12076:5040")},
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 					RouteFilterRuleType: to.Ptr(armnetwork.RouteFilterRuleTypeCommunity),
		// 				},
		// 		}},
		// 	}
	}
}
