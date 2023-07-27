//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/GetAlertIncidentById.json
func ExampleAlertIncidentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAlertIncidentsClient().Get(ctx, "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", "TooManyOwnersAssignedToResource", "5cf9ee65-d22e-4784-8b17-3de1c3b7bdcc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertIncident = armauthorization.AlertIncident{
	// 	Name: to.Ptr("5cf9ee65-d22e-4784-8b17-3de1c3b7bdcc"),
	// 	Type: to.Ptr("Microsoft.Authorization/roleManagementAlerts/alertIncidents"),
	// 	ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlerts/TooManyOwnersAssignedToResource/alertIncidents/5cf9ee65-d22e-4784-8b17-3de1c3b7bdcc"),
	// 	Properties: &armauthorization.TooManyOwnersAssignedToResourceAlertIncidentProperties{
	// 		AlertIncidentType: to.Ptr("TooManyOwnersAssignedToResourceAlertIncident"),
	// 		AssigneeName: to.Ptr("test-user"),
	// 		AssigneeType: to.Ptr("User"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/GetAlertIncidents.json
func ExampleAlertIncidentsClient_NewListForScopePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAlertIncidentsClient().NewListForScopePager("subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", "TooManyOwnersAssignedToResource", nil)
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
		// page.AlertIncidentListResult = armauthorization.AlertIncidentListResult{
		// 	Value: []*armauthorization.AlertIncident{
		// 		{
		// 			Name: to.Ptr("5cf9ee65-d22e-4784-8b17-3de1c3b7bdcc"),
		// 			Type: to.Ptr("Microsoft.Authorization/roleManagementAlerts/alertIncidents"),
		// 			ID: to.Ptr("/subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleManagementAlerts/TooManyOwnersAssignedToResource/alertIncidents/5cf9ee65-d22e-4784-8b17-3de1c3b7bdcc"),
		// 			Properties: &armauthorization.TooManyOwnersAssignedToResourceAlertIncidentProperties{
		// 				AlertIncidentType: to.Ptr("TooManyOwnersAssignedToResourceAlertIncident"),
		// 				AssigneeName: to.Ptr("test-user"),
		// 				AssigneeType: to.Ptr("User"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/RemediateAlertIncident.json
func ExampleAlertIncidentsClient_Remediate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armauthorization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAlertIncidentsClient().Remediate(ctx, "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f", "DuplicateRoleCreated", "0645231d-16ba-4ebf-851a-0875df4052bd", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
