//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/CalculateExchange.json
func ExampleCalculateExchangeClient_BeginPost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armreservations.NewCalculateExchangeClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginPost(ctx,
		armreservations.CalculateExchangeRequest{
			Properties: &armreservations.CalculateExchangeRequestProperties{
				ReservationsToExchange: []*armreservations.ReservationToReturn{
					{
						Quantity:      to.Ptr[int32](1),
						ReservationID: to.Ptr("<reservation-id>"),
					}},
				ReservationsToPurchase: []*armreservations.PurchaseRequest{
					{
						Location: to.Ptr("<location>"),
						Properties: &armreservations.PurchaseRequestProperties{
							AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
							BillingPlan:      to.Ptr(armreservations.ReservationBillingPlanUpfront),
							BillingScopeID:   to.Ptr("<billing-scope-id>"),
							DisplayName:      to.Ptr("<display-name>"),
							Quantity:         to.Ptr[int32](1),
							Renew:            to.Ptr(false),
							ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
								InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
							},
							ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
							Term:                 to.Ptr(armreservations.ReservationTermP1Y),
						},
						SKU: &armreservations.SKUName{
							Name: to.Ptr("<name>"),
						},
					}},
			},
		},
		&armreservations.CalculateExchangeClientBeginPostOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
