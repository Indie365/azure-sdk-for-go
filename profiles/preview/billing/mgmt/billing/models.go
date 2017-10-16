// +build go1.9

// Copyright 2017 Microsoft Corporation
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
// commit ID: 714052a3359963ba46703fe033cf9dd4838bea11

package billing

import original "github.com/Azure/azure-sdk-for-go/services/billing/mgmt/2017-04-24-preview/billing"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type InvoicesClient = original.InvoicesClient
type DownloadURL = original.DownloadURL
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Invoice = original.Invoice
type InvoiceProperties = original.InvoiceProperties
type InvoicesListResult = original.InvoicesListResult
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Period = original.Period
type PeriodProperties = original.PeriodProperties
type PeriodsListResult = original.PeriodsListResult
type Resource = original.Resource
type OperationsClient = original.OperationsClient
type PeriodsClient = original.PeriodsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPeriodsClient(subscriptionID string) PeriodsClient {
	return original.NewPeriodsClient(subscriptionID)
}
func NewPeriodsClientWithBaseURI(baseURI string, subscriptionID string) PeriodsClient {
	return original.NewPeriodsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewInvoicesClient(subscriptionID string) InvoicesClient {
	return original.NewInvoicesClient(subscriptionID)
}
func NewInvoicesClientWithBaseURI(baseURI string, subscriptionID string) InvoicesClient {
	return original.NewInvoicesClientWithBaseURI(baseURI, subscriptionID)
}
