//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicelinker

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker"
	moduleVersion = "v1.2.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AuthType - The authentication type.
type AuthType string

const (
	AuthTypeSecret                      AuthType = "secret"
	AuthTypeServicePrincipalCertificate AuthType = "servicePrincipalCertificate"
	AuthTypeServicePrincipalSecret      AuthType = "servicePrincipalSecret"
	AuthTypeSystemAssignedIdentity      AuthType = "systemAssignedIdentity"
	AuthTypeUserAssignedIdentity        AuthType = "userAssignedIdentity"
)

// PossibleAuthTypeValues returns the possible values for the AuthType const type.
func PossibleAuthTypeValues() []AuthType {
	return []AuthType{
		AuthTypeSecret,
		AuthTypeServicePrincipalCertificate,
		AuthTypeServicePrincipalSecret,
		AuthTypeSystemAssignedIdentity,
		AuthTypeUserAssignedIdentity,
	}
}

// AzureResourceType - The azure resource type.
type AzureResourceType string

const (
	AzureResourceTypeKeyVault AzureResourceType = "KeyVault"
)

// PossibleAzureResourceTypeValues returns the possible values for the AzureResourceType const type.
func PossibleAzureResourceTypeValues() []AzureResourceType {
	return []AzureResourceType{
		AzureResourceTypeKeyVault,
	}
}

// ClientType - The application client type
type ClientType string

const (
	ClientTypeDjango     ClientType = "django"
	ClientTypeDotnet     ClientType = "dotnet"
	ClientTypeGo         ClientType = "go"
	ClientTypeJava       ClientType = "java"
	ClientTypeNodejs     ClientType = "nodejs"
	ClientTypeNone       ClientType = "none"
	ClientTypePhp        ClientType = "php"
	ClientTypePython     ClientType = "python"
	ClientTypeRuby       ClientType = "ruby"
	ClientTypeSpringBoot ClientType = "springBoot"
)

// PossibleClientTypeValues returns the possible values for the ClientType const type.
func PossibleClientTypeValues() []ClientType {
	return []ClientType{
		ClientTypeDjango,
		ClientTypeDotnet,
		ClientTypeGo,
		ClientTypeJava,
		ClientTypeNodejs,
		ClientTypeNone,
		ClientTypePhp,
		ClientTypePython,
		ClientTypeRuby,
		ClientTypeSpringBoot,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// SecretType - The secret type.
type SecretType string

const (
	SecretTypeKeyVaultSecretReference SecretType = "keyVaultSecretReference"
	SecretTypeKeyVaultSecretURI       SecretType = "keyVaultSecretUri"
	SecretTypeRawValue                SecretType = "rawValue"
)

// PossibleSecretTypeValues returns the possible values for the SecretType const type.
func PossibleSecretTypeValues() []SecretType {
	return []SecretType{
		SecretTypeKeyVaultSecretReference,
		SecretTypeKeyVaultSecretURI,
		SecretTypeRawValue,
	}
}

// TargetServiceType - The target service type.
type TargetServiceType string

const (
	TargetServiceTypeAzureResource            TargetServiceType = "AzureResource"
	TargetServiceTypeConfluentBootstrapServer TargetServiceType = "ConfluentBootstrapServer"
	TargetServiceTypeConfluentSchemaRegistry  TargetServiceType = "ConfluentSchemaRegistry"
)

// PossibleTargetServiceTypeValues returns the possible values for the TargetServiceType const type.
func PossibleTargetServiceTypeValues() []TargetServiceType {
	return []TargetServiceType{
		TargetServiceTypeAzureResource,
		TargetServiceTypeConfluentBootstrapServer,
		TargetServiceTypeConfluentSchemaRegistry,
	}
}

// VNetSolutionType - Type of VNet solution.
type VNetSolutionType string

const (
	VNetSolutionTypePrivateLink     VNetSolutionType = "privateLink"
	VNetSolutionTypeServiceEndpoint VNetSolutionType = "serviceEndpoint"
)

// PossibleVNetSolutionTypeValues returns the possible values for the VNetSolutionType const type.
func PossibleVNetSolutionTypeValues() []VNetSolutionType {
	return []VNetSolutionType{
		VNetSolutionTypePrivateLink,
		VNetSolutionTypeServiceEndpoint,
	}
}

// ValidationResultStatus - The result of validation
type ValidationResultStatus string

const (
	ValidationResultStatusFailure ValidationResultStatus = "failure"
	ValidationResultStatusSuccess ValidationResultStatus = "success"
	ValidationResultStatusWarning ValidationResultStatus = "warning"
)

// PossibleValidationResultStatusValues returns the possible values for the ValidationResultStatus const type.
func PossibleValidationResultStatusValues() []ValidationResultStatus {
	return []ValidationResultStatus{
		ValidationResultStatusFailure,
		ValidationResultStatusSuccess,
		ValidationResultStatusWarning,
	}
}
