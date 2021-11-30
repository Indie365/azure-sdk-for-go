//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AuthorizationProfile - Authorization Profile
type AuthorizationProfile struct {
	// READ-ONLY; The approved time
	ApprovedTime *time.Time `json:"approvedTime,omitempty" azure:"ro"`

	// READ-ONLY; The approver
	Approver *string `json:"approver,omitempty" azure:"ro"`

	// READ-ONLY; The requested time
	RequestedTime *time.Time `json:"requestedTime,omitempty" azure:"ro"`

	// READ-ONLY; The requester
	Requester *string `json:"requester,omitempty" azure:"ro"`

	// READ-ONLY; The requester object id
	RequesterObjectID *string `json:"requesterObjectId,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type AuthorizationProfile.
func (a AuthorizationProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "approvedTime", a.ApprovedTime)
	populate(objectMap, "approver", a.Approver)
	populateTimeRFC3339(objectMap, "requestedTime", a.RequestedTime)
	populate(objectMap, "requester", a.Requester)
	populate(objectMap, "requesterObjectId", a.RequesterObjectID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AuthorizationProfile.
func (a *AuthorizationProfile) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvedTime":
			err = unpopulateTimeRFC3339(val, &a.ApprovedTime)
			delete(rawMsg, key)
		case "approver":
			err = unpopulate(val, &a.Approver)
			delete(rawMsg, key)
		case "requestedTime":
			err = unpopulateTimeRFC3339(val, &a.RequestedTime)
			delete(rawMsg, key)
		case "requester":
			err = unpopulate(val, &a.Requester)
			delete(rawMsg, key)
		case "requesterObjectId":
			err = unpopulate(val, &a.RequesterObjectID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// Internal error details.
	Details []*ErrorDefinition `json:"details,omitempty"`

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDefinition.
func (e ErrorDefinition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// ErrorResponse - Error response indicates that the service is not able to process the incoming request.
// Implements the error and azcore.HTTPResponse interfaces.
type ErrorResponse struct {
	raw string
	// The error details.
	InnerError *ErrorDefinition `json:"error,omitempty"`
}

// Error implements the error interface for type ErrorResponse.
// The contents of the error text are not contractual and subject to change.
func (e ErrorResponse) Error() string {
	return e.raw
}

// FeatureClientListOperationsOptions contains the optional parameters for the FeatureClient.ListOperations method.
type FeatureClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// FeatureOperationsListResult - List of previewed features.
type FeatureOperationsListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The array of features.
	Value []*FeatureResult `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type FeatureOperationsListResult.
func (f FeatureOperationsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", f.NextLink)
	populate(objectMap, "value", f.Value)
	return json.Marshal(objectMap)
}

// FeatureProperties - Information about feature.
type FeatureProperties struct {
	// The registration state of the feature for the subscription.
	State *string `json:"state,omitempty"`
}

// FeatureResult - Previewed feature information.
type FeatureResult struct {
	// The resource ID of the feature.
	ID *string `json:"id,omitempty"`

	// The name of the feature.
	Name *string `json:"name,omitempty"`

	// Properties of the previewed feature.
	Properties *FeatureProperties `json:"properties,omitempty"`

	// The resource type of the feature.
	Type *string `json:"type,omitempty"`
}

// FeaturesGetOptions contains the optional parameters for the Features.Get method.
type FeaturesGetOptions struct {
	// placeholder for future optional parameters
}

// FeaturesListAllOptions contains the optional parameters for the Features.ListAll method.
type FeaturesListAllOptions struct {
	// placeholder for future optional parameters
}

// FeaturesListOptions contains the optional parameters for the Features.List method.
type FeaturesListOptions struct {
	// placeholder for future optional parameters
}

// FeaturesRegisterOptions contains the optional parameters for the Features.Register method.
type FeaturesRegisterOptions struct {
	// placeholder for future optional parameters
}

// FeaturesUnregisterOptions contains the optional parameters for the Features.Unregister method.
type FeaturesUnregisterOptions struct {
	// placeholder for future optional parameters
}

// Operation - Microsoft.Features operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.Features
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of the request to list Microsoft.Features operations. It contains a list of operations and a URL link to get the next set
// of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Microsoft.Features operations.
	Value []*Operation `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// ProxyResource - An Azure proxy resource.
type ProxyResource struct {
	// READ-ONLY; Azure resource Id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SubscriptionFeatureRegistration - Subscription feature registration details
type SubscriptionFeatureRegistration struct {
	ProxyResource
	Properties *SubscriptionFeatureRegistrationProperties `json:"properties,omitempty"`
}

// SubscriptionFeatureRegistrationList - The list of subscription feature registrations.
type SubscriptionFeatureRegistrationList struct {
	// The link used to get the next page of subscription feature registrations list.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of subscription feature registrations.
	Value []*SubscriptionFeatureRegistration `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionFeatureRegistrationList.
func (s SubscriptionFeatureRegistrationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

type SubscriptionFeatureRegistrationProperties struct {
	// Authorization Profile
	AuthorizationProfile *AuthorizationProfile `json:"authorizationProfile,omitempty"`

	// The feature description.
	Description *string `json:"description,omitempty"`

	// Key-value pairs for meta data.
	Metadata map[string]*string `json:"metadata,omitempty"`

	// Indicates whether feature should be displayed in Portal.
	ShouldFeatureDisplayInPortal *bool `json:"shouldFeatureDisplayInPortal,omitempty"`

	// The state.
	State *SubscriptionFeatureRegistrationState `json:"state,omitempty"`

	// READ-ONLY; The feature approval type.
	ApprovalType *SubscriptionFeatureRegistrationApprovalType `json:"approvalType,omitempty" azure:"ro"`

	// READ-ONLY; The featureDisplayName.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The feature documentation link.
	DocumentationLink *string `json:"documentationLink,omitempty" azure:"ro"`

	// READ-ONLY; The featureName.
	FeatureName *string `json:"featureName,omitempty" azure:"ro"`

	// READ-ONLY; The providerNamespace.
	ProviderNamespace *string `json:"providerNamespace,omitempty" azure:"ro"`

	// READ-ONLY; The feature registration date.
	RegistrationDate *time.Time `json:"registrationDate,omitempty" azure:"ro"`

	// READ-ONLY; The feature release date.
	ReleaseDate *time.Time `json:"releaseDate,omitempty" azure:"ro"`

	// READ-ONLY; The subscriptionId.
	SubscriptionID *string `json:"subscriptionId,omitempty" azure:"ro"`

	// READ-ONLY; The tenantId.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionFeatureRegistrationProperties.
func (s SubscriptionFeatureRegistrationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "approvalType", s.ApprovalType)
	populate(objectMap, "authorizationProfile", s.AuthorizationProfile)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "displayName", s.DisplayName)
	populate(objectMap, "documentationLink", s.DocumentationLink)
	populate(objectMap, "featureName", s.FeatureName)
	populate(objectMap, "metadata", s.Metadata)
	populate(objectMap, "providerNamespace", s.ProviderNamespace)
	populateTimeRFC3339(objectMap, "registrationDate", s.RegistrationDate)
	populateTimeRFC3339(objectMap, "releaseDate", s.ReleaseDate)
	populate(objectMap, "shouldFeatureDisplayInPortal", s.ShouldFeatureDisplayInPortal)
	populate(objectMap, "state", s.State)
	populate(objectMap, "subscriptionId", s.SubscriptionID)
	populate(objectMap, "tenantId", s.TenantID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubscriptionFeatureRegistrationProperties.
func (s *SubscriptionFeatureRegistrationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "approvalType":
			err = unpopulate(val, &s.ApprovalType)
			delete(rawMsg, key)
		case "authorizationProfile":
			err = unpopulate(val, &s.AuthorizationProfile)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, &s.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, &s.DisplayName)
			delete(rawMsg, key)
		case "documentationLink":
			err = unpopulate(val, &s.DocumentationLink)
			delete(rawMsg, key)
		case "featureName":
			err = unpopulate(val, &s.FeatureName)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, &s.Metadata)
			delete(rawMsg, key)
		case "providerNamespace":
			err = unpopulate(val, &s.ProviderNamespace)
			delete(rawMsg, key)
		case "registrationDate":
			err = unpopulateTimeRFC3339(val, &s.RegistrationDate)
			delete(rawMsg, key)
		case "releaseDate":
			err = unpopulateTimeRFC3339(val, &s.ReleaseDate)
			delete(rawMsg, key)
		case "shouldFeatureDisplayInPortal":
			err = unpopulate(val, &s.ShouldFeatureDisplayInPortal)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, &s.State)
			delete(rawMsg, key)
		case "subscriptionId":
			err = unpopulate(val, &s.SubscriptionID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, &s.TenantID)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// SubscriptionFeatureRegistrationsCreateOrUpdateOptions contains the optional parameters for the SubscriptionFeatureRegistrations.CreateOrUpdate method.
type SubscriptionFeatureRegistrationsCreateOrUpdateOptions struct {
	// Subscription Feature Registration Type details.
	SubscriptionFeatureRegistrationType *SubscriptionFeatureRegistration
}

// SubscriptionFeatureRegistrationsDeleteOptions contains the optional parameters for the SubscriptionFeatureRegistrations.Delete method.
type SubscriptionFeatureRegistrationsDeleteOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionFeatureRegistrationsGetOptions contains the optional parameters for the SubscriptionFeatureRegistrations.Get method.
type SubscriptionFeatureRegistrationsGetOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionFeatureRegistrationsListAllBySubscriptionOptions contains the optional parameters for the SubscriptionFeatureRegistrations.ListAllBySubscription
// method.
type SubscriptionFeatureRegistrationsListAllBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionFeatureRegistrationsListBySubscriptionOptions contains the optional parameters for the SubscriptionFeatureRegistrations.ListBySubscription
// method.
type SubscriptionFeatureRegistrationsListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
