// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package support

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/support/mgmt/2019-05-01-preview/support"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CommunicationDirection = original.CommunicationDirection

const (
	Inbound  CommunicationDirection = original.Inbound
	Outbound CommunicationDirection = original.Outbound
)

type CommunicationType = original.CommunicationType

const (
	Phone CommunicationType = original.Phone
	Web   CommunicationType = original.Web
)

type PreferredContactMethod = original.PreferredContactMethod

const (
	PreferredContactMethodEmail PreferredContactMethod = original.PreferredContactMethodEmail
	PreferredContactMethodPhone PreferredContactMethod = original.PreferredContactMethodPhone
)

type SeverityLevel = original.SeverityLevel

const (
	Critical SeverityLevel = original.Critical
	Minimal  SeverityLevel = original.Minimal
	Moderate SeverityLevel = original.Moderate
)

type Type = original.Type

const (
	MicrosoftSupportcommunications Type = original.MicrosoftSupportcommunications
	MicrosoftSupportsupportTickets Type = original.MicrosoftSupportsupportTickets
)

type BaseClient = original.BaseClient
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CommunicationDetails = original.CommunicationDetails
type CommunicationDetailsProperties = original.CommunicationDetailsProperties
type CommunicationsClient = original.CommunicationsClient
type CommunicationsListResult = original.CommunicationsListResult
type CommunicationsListResultIterator = original.CommunicationsListResultIterator
type CommunicationsListResultPage = original.CommunicationsListResultPage
type ContactProfile = original.ContactProfile
type CreateSupportTicketCommunicationFuture = original.CreateSupportTicketCommunicationFuture
type CreateSupportTicketForSubscriptionFuture = original.CreateSupportTicketForSubscriptionFuture
type Engineer = original.Engineer
type ExceptionResponse = original.ExceptionResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsListResult = original.OperationsListResult
type ProblemClassification = original.ProblemClassification
type ProblemClassificationProperties = original.ProblemClassificationProperties
type ProblemClassificationsClient = original.ProblemClassificationsClient
type ProblemClassificationsListResult = original.ProblemClassificationsListResult
type QuotaChangeRequest = original.QuotaChangeRequest
type QuotaTicketDetails = original.QuotaTicketDetails
type Service = original.Service
type ServiceError = original.ServiceError
type ServiceErrorDetail = original.ServiceErrorDetail
type ServiceLevelAgreement = original.ServiceLevelAgreement
type ServiceProperties = original.ServiceProperties
type ServicesClient = original.ServicesClient
type ServicesListResult = original.ServicesListResult
type TechnicalTicketDetails = original.TechnicalTicketDetails
type TicketDetails = original.TicketDetails
type TicketDetailsProperties = original.TicketDetailsProperties
type TicketSubscriptionClient = original.TicketSubscriptionClient
type TicketsClient = original.TicketsClient
type TicketsListResult = original.TicketsListResult
type TicketsListResultIterator = original.TicketsListResultIterator
type TicketsListResultPage = original.TicketsListResultPage
type UpdateContactProfile = original.UpdateContactProfile
type UpdateSupportTicket = original.UpdateSupportTicket

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCommunicationsClient(subscriptionID string) CommunicationsClient {
	return original.NewCommunicationsClient(subscriptionID)
}
func NewCommunicationsClientWithBaseURI(baseURI string, subscriptionID string) CommunicationsClient {
	return original.NewCommunicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewCommunicationsListResultIterator(page CommunicationsListResultPage) CommunicationsListResultIterator {
	return original.NewCommunicationsListResultIterator(page)
}
func NewCommunicationsListResultPage(getNextPage func(context.Context, CommunicationsListResult) (CommunicationsListResult, error)) CommunicationsListResultPage {
	return original.NewCommunicationsListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProblemClassificationsClient(subscriptionID string) ProblemClassificationsClient {
	return original.NewProblemClassificationsClient(subscriptionID)
}
func NewProblemClassificationsClientWithBaseURI(baseURI string, subscriptionID string) ProblemClassificationsClient {
	return original.NewProblemClassificationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTicketSubscriptionClient(subscriptionID string) TicketSubscriptionClient {
	return original.NewTicketSubscriptionClient(subscriptionID)
}
func NewTicketSubscriptionClientWithBaseURI(baseURI string, subscriptionID string) TicketSubscriptionClient {
	return original.NewTicketSubscriptionClientWithBaseURI(baseURI, subscriptionID)
}
func NewTicketsClient(subscriptionID string) TicketsClient {
	return original.NewTicketsClient(subscriptionID)
}
func NewTicketsClientWithBaseURI(baseURI string, subscriptionID string) TicketsClient {
	return original.NewTicketsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTicketsListResultIterator(page TicketsListResultPage) TicketsListResultIterator {
	return original.NewTicketsListResultIterator(page)
}
func NewTicketsListResultPage(getNextPage func(context.Context, TicketsListResult) (TicketsListResult, error)) TicketsListResultPage {
	return original.NewTicketsListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCommunicationDirectionValues() []CommunicationDirection {
	return original.PossibleCommunicationDirectionValues()
}
func PossibleCommunicationTypeValues() []CommunicationType {
	return original.PossibleCommunicationTypeValues()
}
func PossiblePreferredContactMethodValues() []PreferredContactMethod {
	return original.PossiblePreferredContactMethodValues()
}
func PossibleSeverityLevelValues() []SeverityLevel {
	return original.PossibleSeverityLevelValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
