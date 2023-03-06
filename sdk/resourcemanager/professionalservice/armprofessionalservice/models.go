//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armprofessionalservice

// AppOperation - professionalService app operations
type AppOperation struct {
	// the operation display
	Display *AppOperationDisplay `json:"display,omitempty"`

	// whether the operation is a data action or not.
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// the operation name
	Name *string `json:"name,omitempty"`

	// the operation origin
	Origin *string `json:"origin,omitempty"`
}

// AppOperationDisplay - ProfessionalService app operation display
type AppOperationDisplay struct {
	// Description of the operation for display purposes
	Description *string `json:"description,omitempty"`

	// Name of the operation for display purposes
	Operation *string `json:"operation,omitempty"`

	// Name of the provider for display purposes
	Provider *string `json:"provider,omitempty"`

	// Name of the resource type for display purposes
	Resource *string `json:"resource,omitempty"`
}

// AppOperationsResponseWithContinuation - professionalService app operation response with continuation.
type AppOperationsResponseWithContinuation struct {
	// the next link to query to get the remaining results.
	NextLink *string `json:"nextLink,omitempty"`

	// the value of response.
	Value []*AppOperation `json:"value,omitempty"`
}

// CreationProperties - properties for creation professionalService
type CreationProperties struct {
	// Whether the ProfessionalService subscription will auto renew upon term end.
	AutoRenew *bool `json:"autoRenew,omitempty"`

	// The billing period eg P1M,P1Y for monthly,yearly respectively
	BillingPeriod *string `json:"billingPeriod,omitempty"`

	// The offer id.
	OfferID *string `json:"offerId,omitempty"`

	// The publisher id.
	PublisherID *string `json:"publisherId,omitempty"`

	// The quote id which the ProfessionalService will be purchase with.
	QuoteID *string `json:"quoteId,omitempty"`

	// The plan id.
	SKUID *string `json:"skuId,omitempty"`

	// The store front which initiates the purchase.
	StoreFront *string `json:"storeFront,omitempty"`

	// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
	TermUnit *string `json:"termUnit,omitempty"`
}

// DeleteOptions - delete Options
type DeleteOptions struct {
	// the feedback
	Feedback *string `json:"feedback,omitempty"`

	// The reasonCode
	ReasonCode *float32 `json:"reasonCode,omitempty"`

	// whether it is unsubscribeOnly
	UnsubscribeOnly *bool `json:"unsubscribeOnly,omitempty"`
}

// OperationClientBeginGetOptions contains the optional parameters for the OperationClient.BeginGet method.
type OperationClientBeginGetOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PropertiesTerm - The current Term object.
type PropertiesTerm struct {
	// The end date of the current term
	EndDate *string `json:"endDate,omitempty"`

	// The start date of the current term
	StartDate *string `json:"startDate,omitempty"`

	// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
	TermUnit *string `json:"termUnit,omitempty"`
}

// Resource - ProfessionalService REST API resource definition.
type Resource struct {
	// professionalService properties
	Properties *ResourceProperties `json:"properties,omitempty"`

	// the resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource uri
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceCreation - ProfessionalService REST API resource definition for creation.
type ResourceCreation struct {
	// Resource location. Only value allowed for ProfessionalService is 'global'
	Location *string `json:"location,omitempty"`

	// The resource name
	Name *string `json:"name,omitempty"`

	// Properties of the ProfessionalService resource that are relevant for creation.
	Properties *CreationProperties `json:"properties,omitempty"`

	// the resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource uri
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceProperties - professionalService properties
type ResourceProperties struct {
	// Whether the ProfessionalService subscription will auto renew upon term end.
	AutoRenew *bool `json:"autoRenew,omitempty"`

	// The billing period eg P1M,P1Y for monthly,yearly respectively
	BillingPeriod *string `json:"billingPeriod,omitempty"`

	// Whether the current term is a Free Trial term
	IsFreeTrial *bool `json:"isFreeTrial,omitempty"`

	// The last modifier date if this resource.
	LastModified *string `json:"lastModified,omitempty"`

	// The offer id.
	OfferID *string `json:"offerId,omitempty"`

	// The metadata about the ProfessionalService subscription such as the AzureSubscriptionId and ResourceUri.
	PaymentChannelMetadata map[string]*string `json:"paymentChannelMetadata,omitempty"`

	// The Payment channel for the ProfessionalServiceSubscription.
	PaymentChannelType *PaymentChannelType `json:"paymentChannelType,omitempty"`

	// The publisher id.
	PublisherID *string `json:"publisherId,omitempty"`

	// The quote id which the ProfessionalService will be purchase with.
	QuoteID *string `json:"quoteId,omitempty"`

	// The plan id.
	SKUID *string `json:"skuId,omitempty"`

	// The ProfessionalService Subscription Status.
	Status *ProfessionalServiceResourceStatus `json:"status,omitempty"`

	// The store front which initiates the purchase.
	StoreFront *string `json:"storeFront,omitempty"`

	// The current Term object.
	Term *PropertiesTerm `json:"term,omitempty"`

	// The unit term eg P1M,P1Y,P2Y,P3Y meaning month,1year,2year,3year respectively
	TermUnit *string `json:"termUnit,omitempty"`

	// READ-ONLY; The created date of this resource.
	Created *string `json:"created,omitempty" azure:"ro"`
}

// ResourceResponseWithContinuation - professionalService resources response with continuation.
type ResourceResponseWithContinuation struct {
	// the next link to query to get the remaining results.
	NextLink *string `json:"nextLink,omitempty"`

	// the value of response.
	Value []*Resource `json:"value,omitempty"`
}

// SubscriptionLevelClientBeginCreateOrUpdateOptions contains the optional parameters for the SubscriptionLevelClient.BeginCreateOrUpdate
// method.
type SubscriptionLevelClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SubscriptionLevelClientBeginDeleteOptions contains the optional parameters for the SubscriptionLevelClient.BeginDelete
// method.
type SubscriptionLevelClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SubscriptionLevelClientBeginUpdateToUnsubscribedOptions contains the optional parameters for the SubscriptionLevelClient.BeginUpdateToUnsubscribed
// method.
type SubscriptionLevelClientBeginUpdateToUnsubscribedOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SubscriptionLevelClientGetOptions contains the optional parameters for the SubscriptionLevelClient.Get method.
type SubscriptionLevelClientGetOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientListByAzureSubscriptionOptions contains the optional parameters for the SubscriptionLevelClient.NewListByAzureSubscriptionPager
// method.
type SubscriptionLevelClientListByAzureSubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientListByResourceGroupOptions contains the optional parameters for the SubscriptionLevelClient.NewListByResourceGroupPager
// method.
type SubscriptionLevelClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}
