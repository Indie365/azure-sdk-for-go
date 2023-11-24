//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.OperationsListResult = armsupport.OperationsListResult{
		// 	Value: []*armsupport.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/services/read"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Gets all the Azure services available for support"),
		// 				Operation: to.Ptr("Reads Services"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("Service"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/problemClassifications/read"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Gets all the problem classifications available for a specific Azure service"),
		// 				Operation: to.Ptr("Reads Problem Classifications"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("Problem Classification"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/supportTickets/read"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Gets all the support tickets"),
		// 				Operation: to.Ptr("Reads Support Tickets"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("Support Ticket"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/supportTickets/write"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Updates support ticket"),
		// 				Operation: to.Ptr("Updates support ticket"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("Support Ticket"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/communications/read"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Gets all the communications"),
		// 				Operation: to.Ptr("Reads Communications"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("Communication"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/communications/write"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Creates a communication"),
		// 				Operation: to.Ptr("Creates a communication"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("Communication"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/register/action"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Registers Support Resource Provider"),
		// 				Operation: to.Ptr("Registers Support Resource Provider"),
		// 				Provider: to.Ptr("Registers Support Resource Provider"),
		// 				Resource: to.Ptr("Support Registration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/createSupportTicket/action"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Creates support ticket"),
		// 				Operation: to.Ptr("Registers Support Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("SupportTicket"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Support/addCommunication/action"),
		// 			Display: &armsupport.OperationDisplay{
		// 				Description: to.Ptr("Add communication to support ticket"),
		// 				Operation: to.Ptr("Registers Support Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Support"),
		// 				Resource: to.Ptr("Communication"),
		// 			},
		// 	}},
		// }
	}
}
