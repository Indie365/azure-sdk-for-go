//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

const (
	moduleName    = "armappplatform"
	moduleVersion = "v2.0.0-beta.1"
)

// APIPortalProvisioningState - State of the API portal.
type APIPortalProvisioningState string

const (
	APIPortalProvisioningStateCreating  APIPortalProvisioningState = "Creating"
	APIPortalProvisioningStateDeleting  APIPortalProvisioningState = "Deleting"
	APIPortalProvisioningStateFailed    APIPortalProvisioningState = "Failed"
	APIPortalProvisioningStateSucceeded APIPortalProvisioningState = "Succeeded"
	APIPortalProvisioningStateUpdating  APIPortalProvisioningState = "Updating"
)

// PossibleAPIPortalProvisioningStateValues returns the possible values for the APIPortalProvisioningState const type.
func PossibleAPIPortalProvisioningStateValues() []APIPortalProvisioningState {
	return []APIPortalProvisioningState{
		APIPortalProvisioningStateCreating,
		APIPortalProvisioningStateDeleting,
		APIPortalProvisioningStateFailed,
		APIPortalProvisioningStateSucceeded,
		APIPortalProvisioningStateUpdating,
	}
}

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

// ApmType - Type of application performance monitoring
type ApmType string

const (
	ApmTypeAppDynamics         ApmType = "AppDynamics"
	ApmTypeApplicationInsights ApmType = "ApplicationInsights"
	ApmTypeDynatrace           ApmType = "Dynatrace"
	ApmTypeElasticAPM          ApmType = "ElasticAPM"
	ApmTypeNewRelic            ApmType = "NewRelic"
)

// PossibleApmTypeValues returns the possible values for the ApmType const type.
func PossibleApmTypeValues() []ApmType {
	return []ApmType{
		ApmTypeAppDynamics,
		ApmTypeApplicationInsights,
		ApmTypeDynatrace,
		ApmTypeElasticAPM,
		ApmTypeNewRelic,
	}
}

// AppResourceProvisioningState - Provisioning state of the App
type AppResourceProvisioningState string

const (
	AppResourceProvisioningStateCreating  AppResourceProvisioningState = "Creating"
	AppResourceProvisioningStateDeleting  AppResourceProvisioningState = "Deleting"
	AppResourceProvisioningStateFailed    AppResourceProvisioningState = "Failed"
	AppResourceProvisioningStateSucceeded AppResourceProvisioningState = "Succeeded"
	AppResourceProvisioningStateUpdating  AppResourceProvisioningState = "Updating"
)

// PossibleAppResourceProvisioningStateValues returns the possible values for the AppResourceProvisioningState const type.
func PossibleAppResourceProvisioningStateValues() []AppResourceProvisioningState {
	return []AppResourceProvisioningState{
		AppResourceProvisioningStateCreating,
		AppResourceProvisioningStateDeleting,
		AppResourceProvisioningStateFailed,
		AppResourceProvisioningStateSucceeded,
		AppResourceProvisioningStateUpdating,
	}
}

// ApplicationAcceleratorProvisioningState - State of the application accelerator.
type ApplicationAcceleratorProvisioningState string

const (
	ApplicationAcceleratorProvisioningStateCreating  ApplicationAcceleratorProvisioningState = "Creating"
	ApplicationAcceleratorProvisioningStateDeleting  ApplicationAcceleratorProvisioningState = "Deleting"
	ApplicationAcceleratorProvisioningStateFailed    ApplicationAcceleratorProvisioningState = "Failed"
	ApplicationAcceleratorProvisioningStateSucceeded ApplicationAcceleratorProvisioningState = "Succeeded"
	ApplicationAcceleratorProvisioningStateUpdating  ApplicationAcceleratorProvisioningState = "Updating"
)

// PossibleApplicationAcceleratorProvisioningStateValues returns the possible values for the ApplicationAcceleratorProvisioningState const type.
func PossibleApplicationAcceleratorProvisioningStateValues() []ApplicationAcceleratorProvisioningState {
	return []ApplicationAcceleratorProvisioningState{
		ApplicationAcceleratorProvisioningStateCreating,
		ApplicationAcceleratorProvisioningStateDeleting,
		ApplicationAcceleratorProvisioningStateFailed,
		ApplicationAcceleratorProvisioningStateSucceeded,
		ApplicationAcceleratorProvisioningStateUpdating,
	}
}

// ApplicationLiveViewProvisioningState - State of the Application Live View.
type ApplicationLiveViewProvisioningState string

const (
	ApplicationLiveViewProvisioningStateCanceled  ApplicationLiveViewProvisioningState = "Canceled"
	ApplicationLiveViewProvisioningStateCreating  ApplicationLiveViewProvisioningState = "Creating"
	ApplicationLiveViewProvisioningStateDeleting  ApplicationLiveViewProvisioningState = "Deleting"
	ApplicationLiveViewProvisioningStateFailed    ApplicationLiveViewProvisioningState = "Failed"
	ApplicationLiveViewProvisioningStateSucceeded ApplicationLiveViewProvisioningState = "Succeeded"
	ApplicationLiveViewProvisioningStateUpdating  ApplicationLiveViewProvisioningState = "Updating"
)

// PossibleApplicationLiveViewProvisioningStateValues returns the possible values for the ApplicationLiveViewProvisioningState const type.
func PossibleApplicationLiveViewProvisioningStateValues() []ApplicationLiveViewProvisioningState {
	return []ApplicationLiveViewProvisioningState{
		ApplicationLiveViewProvisioningStateCanceled,
		ApplicationLiveViewProvisioningStateCreating,
		ApplicationLiveViewProvisioningStateDeleting,
		ApplicationLiveViewProvisioningStateFailed,
		ApplicationLiveViewProvisioningStateSucceeded,
		ApplicationLiveViewProvisioningStateUpdating,
	}
}

// BackendProtocol - How ingress should communicate with this app backend service.
type BackendProtocol string

const (
	BackendProtocolDefault BackendProtocol = "Default"
	BackendProtocolGRPC    BackendProtocol = "GRPC"
)

// PossibleBackendProtocolValues returns the possible values for the BackendProtocol const type.
func PossibleBackendProtocolValues() []BackendProtocol {
	return []BackendProtocol{
		BackendProtocolDefault,
		BackendProtocolGRPC,
	}
}

// BindingType - Buildpack Binding Type
type BindingType string

const (
	BindingTypeApacheSkyWalking    BindingType = "ApacheSkyWalking"
	BindingTypeAppDynamics         BindingType = "AppDynamics"
	BindingTypeApplicationInsights BindingType = "ApplicationInsights"
	BindingTypeCACertificates      BindingType = "CACertificates"
	BindingTypeDynatrace           BindingType = "Dynatrace"
	BindingTypeElasticAPM          BindingType = "ElasticAPM"
	BindingTypeNewRelic            BindingType = "NewRelic"
)

// PossibleBindingTypeValues returns the possible values for the BindingType const type.
func PossibleBindingTypeValues() []BindingType {
	return []BindingType{
		BindingTypeApacheSkyWalking,
		BindingTypeAppDynamics,
		BindingTypeApplicationInsights,
		BindingTypeCACertificates,
		BindingTypeDynatrace,
		BindingTypeElasticAPM,
		BindingTypeNewRelic,
	}
}

// BuildProvisioningState - Provisioning state of the KPack build result
type BuildProvisioningState string

const (
	BuildProvisioningStateCreating  BuildProvisioningState = "Creating"
	BuildProvisioningStateDeleting  BuildProvisioningState = "Deleting"
	BuildProvisioningStateFailed    BuildProvisioningState = "Failed"
	BuildProvisioningStateSucceeded BuildProvisioningState = "Succeeded"
	BuildProvisioningStateUpdating  BuildProvisioningState = "Updating"
)

// PossibleBuildProvisioningStateValues returns the possible values for the BuildProvisioningState const type.
func PossibleBuildProvisioningStateValues() []BuildProvisioningState {
	return []BuildProvisioningState{
		BuildProvisioningStateCreating,
		BuildProvisioningStateDeleting,
		BuildProvisioningStateFailed,
		BuildProvisioningStateSucceeded,
		BuildProvisioningStateUpdating,
	}
}

// BuildResultProvisioningState - Provisioning state of the KPack build result
type BuildResultProvisioningState string

const (
	BuildResultProvisioningStateBuilding  BuildResultProvisioningState = "Building"
	BuildResultProvisioningStateDeleting  BuildResultProvisioningState = "Deleting"
	BuildResultProvisioningStateFailed    BuildResultProvisioningState = "Failed"
	BuildResultProvisioningStateQueuing   BuildResultProvisioningState = "Queuing"
	BuildResultProvisioningStateSucceeded BuildResultProvisioningState = "Succeeded"
)

// PossibleBuildResultProvisioningStateValues returns the possible values for the BuildResultProvisioningState const type.
func PossibleBuildResultProvisioningStateValues() []BuildResultProvisioningState {
	return []BuildResultProvisioningState{
		BuildResultProvisioningStateBuilding,
		BuildResultProvisioningStateDeleting,
		BuildResultProvisioningStateFailed,
		BuildResultProvisioningStateQueuing,
		BuildResultProvisioningStateSucceeded,
	}
}

// BuildServiceProvisioningState - Provisioning state of the KPack build result
type BuildServiceProvisioningState string

const (
	BuildServiceProvisioningStateCreating  BuildServiceProvisioningState = "Creating"
	BuildServiceProvisioningStateDeleting  BuildServiceProvisioningState = "Deleting"
	BuildServiceProvisioningStateFailed    BuildServiceProvisioningState = "Failed"
	BuildServiceProvisioningStateSucceeded BuildServiceProvisioningState = "Succeeded"
	BuildServiceProvisioningStateUpdating  BuildServiceProvisioningState = "Updating"
)

// PossibleBuildServiceProvisioningStateValues returns the possible values for the BuildServiceProvisioningState const type.
func PossibleBuildServiceProvisioningStateValues() []BuildServiceProvisioningState {
	return []BuildServiceProvisioningState{
		BuildServiceProvisioningStateCreating,
		BuildServiceProvisioningStateDeleting,
		BuildServiceProvisioningStateFailed,
		BuildServiceProvisioningStateSucceeded,
		BuildServiceProvisioningStateUpdating,
	}
}

// BuilderProvisioningState - Builder provision status.
type BuilderProvisioningState string

const (
	BuilderProvisioningStateCreating  BuilderProvisioningState = "Creating"
	BuilderProvisioningStateDeleting  BuilderProvisioningState = "Deleting"
	BuilderProvisioningStateFailed    BuilderProvisioningState = "Failed"
	BuilderProvisioningStateSucceeded BuilderProvisioningState = "Succeeded"
	BuilderProvisioningStateUpdating  BuilderProvisioningState = "Updating"
)

// PossibleBuilderProvisioningStateValues returns the possible values for the BuilderProvisioningState const type.
func PossibleBuilderProvisioningStateValues() []BuilderProvisioningState {
	return []BuilderProvisioningState{
		BuilderProvisioningStateCreating,
		BuilderProvisioningStateDeleting,
		BuilderProvisioningStateFailed,
		BuilderProvisioningStateSucceeded,
		BuilderProvisioningStateUpdating,
	}
}

// BuildpackBindingProvisioningState - State of the Buildpack Binding.
type BuildpackBindingProvisioningState string

const (
	BuildpackBindingProvisioningStateCreating  BuildpackBindingProvisioningState = "Creating"
	BuildpackBindingProvisioningStateDeleting  BuildpackBindingProvisioningState = "Deleting"
	BuildpackBindingProvisioningStateFailed    BuildpackBindingProvisioningState = "Failed"
	BuildpackBindingProvisioningStateSucceeded BuildpackBindingProvisioningState = "Succeeded"
	BuildpackBindingProvisioningStateUpdating  BuildpackBindingProvisioningState = "Updating"
)

// PossibleBuildpackBindingProvisioningStateValues returns the possible values for the BuildpackBindingProvisioningState const type.
func PossibleBuildpackBindingProvisioningStateValues() []BuildpackBindingProvisioningState {
	return []BuildpackBindingProvisioningState{
		BuildpackBindingProvisioningStateCreating,
		BuildpackBindingProvisioningStateDeleting,
		BuildpackBindingProvisioningStateFailed,
		BuildpackBindingProvisioningStateSucceeded,
		BuildpackBindingProvisioningStateUpdating,
	}
}

// CertificateResourceProvisioningState - Provisioning state of the Certificate
type CertificateResourceProvisioningState string

const (
	CertificateResourceProvisioningStateCreating  CertificateResourceProvisioningState = "Creating"
	CertificateResourceProvisioningStateDeleting  CertificateResourceProvisioningState = "Deleting"
	CertificateResourceProvisioningStateFailed    CertificateResourceProvisioningState = "Failed"
	CertificateResourceProvisioningStateSucceeded CertificateResourceProvisioningState = "Succeeded"
	CertificateResourceProvisioningStateUpdating  CertificateResourceProvisioningState = "Updating"
)

// PossibleCertificateResourceProvisioningStateValues returns the possible values for the CertificateResourceProvisioningState const type.
func PossibleCertificateResourceProvisioningStateValues() []CertificateResourceProvisioningState {
	return []CertificateResourceProvisioningState{
		CertificateResourceProvisioningStateCreating,
		CertificateResourceProvisioningStateDeleting,
		CertificateResourceProvisioningStateFailed,
		CertificateResourceProvisioningStateSucceeded,
		CertificateResourceProvisioningStateUpdating,
	}
}

// ConfigServerState - State of the config server.
type ConfigServerState string

const (
	ConfigServerStateDeleted      ConfigServerState = "Deleted"
	ConfigServerStateFailed       ConfigServerState = "Failed"
	ConfigServerStateNotAvailable ConfigServerState = "NotAvailable"
	ConfigServerStateSucceeded    ConfigServerState = "Succeeded"
	ConfigServerStateUpdating     ConfigServerState = "Updating"
)

// PossibleConfigServerStateValues returns the possible values for the ConfigServerState const type.
func PossibleConfigServerStateValues() []ConfigServerState {
	return []ConfigServerState{
		ConfigServerStateDeleted,
		ConfigServerStateFailed,
		ConfigServerStateNotAvailable,
		ConfigServerStateSucceeded,
		ConfigServerStateUpdating,
	}
}

// ConfigurationServiceProvisioningState - State of the Application Configuration Service.
type ConfigurationServiceProvisioningState string

const (
	ConfigurationServiceProvisioningStateCreating  ConfigurationServiceProvisioningState = "Creating"
	ConfigurationServiceProvisioningStateDeleting  ConfigurationServiceProvisioningState = "Deleting"
	ConfigurationServiceProvisioningStateFailed    ConfigurationServiceProvisioningState = "Failed"
	ConfigurationServiceProvisioningStateSucceeded ConfigurationServiceProvisioningState = "Succeeded"
	ConfigurationServiceProvisioningStateUpdating  ConfigurationServiceProvisioningState = "Updating"
)

// PossibleConfigurationServiceProvisioningStateValues returns the possible values for the ConfigurationServiceProvisioningState const type.
func PossibleConfigurationServiceProvisioningStateValues() []ConfigurationServiceProvisioningState {
	return []ConfigurationServiceProvisioningState{
		ConfigurationServiceProvisioningStateCreating,
		ConfigurationServiceProvisioningStateDeleting,
		ConfigurationServiceProvisioningStateFailed,
		ConfigurationServiceProvisioningStateSucceeded,
		ConfigurationServiceProvisioningStateUpdating,
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

// CustomDomainResourceProvisioningState - Provisioning state of the Domain
type CustomDomainResourceProvisioningState string

const (
	CustomDomainResourceProvisioningStateCreating  CustomDomainResourceProvisioningState = "Creating"
	CustomDomainResourceProvisioningStateDeleting  CustomDomainResourceProvisioningState = "Deleting"
	CustomDomainResourceProvisioningStateFailed    CustomDomainResourceProvisioningState = "Failed"
	CustomDomainResourceProvisioningStateSucceeded CustomDomainResourceProvisioningState = "Succeeded"
	CustomDomainResourceProvisioningStateUpdating  CustomDomainResourceProvisioningState = "Updating"
)

// PossibleCustomDomainResourceProvisioningStateValues returns the possible values for the CustomDomainResourceProvisioningState const type.
func PossibleCustomDomainResourceProvisioningStateValues() []CustomDomainResourceProvisioningState {
	return []CustomDomainResourceProvisioningState{
		CustomDomainResourceProvisioningStateCreating,
		CustomDomainResourceProvisioningStateDeleting,
		CustomDomainResourceProvisioningStateFailed,
		CustomDomainResourceProvisioningStateSucceeded,
		CustomDomainResourceProvisioningStateUpdating,
	}
}

// CustomizedAcceleratorProvisioningState - State of the customized accelerator.
type CustomizedAcceleratorProvisioningState string

const (
	CustomizedAcceleratorProvisioningStateCreating  CustomizedAcceleratorProvisioningState = "Creating"
	CustomizedAcceleratorProvisioningStateDeleting  CustomizedAcceleratorProvisioningState = "Deleting"
	CustomizedAcceleratorProvisioningStateFailed    CustomizedAcceleratorProvisioningState = "Failed"
	CustomizedAcceleratorProvisioningStateSucceeded CustomizedAcceleratorProvisioningState = "Succeeded"
	CustomizedAcceleratorProvisioningStateUpdating  CustomizedAcceleratorProvisioningState = "Updating"
)

// PossibleCustomizedAcceleratorProvisioningStateValues returns the possible values for the CustomizedAcceleratorProvisioningState const type.
func PossibleCustomizedAcceleratorProvisioningStateValues() []CustomizedAcceleratorProvisioningState {
	return []CustomizedAcceleratorProvisioningState{
		CustomizedAcceleratorProvisioningStateCreating,
		CustomizedAcceleratorProvisioningStateDeleting,
		CustomizedAcceleratorProvisioningStateFailed,
		CustomizedAcceleratorProvisioningStateSucceeded,
		CustomizedAcceleratorProvisioningStateUpdating,
	}
}

// CustomizedAcceleratorValidateResultState - State of the customized accelerator validation result
type CustomizedAcceleratorValidateResultState string

const (
	// CustomizedAcceleratorValidateResultStateInvalid - Customized accelerator properties are invalid.
	CustomizedAcceleratorValidateResultStateInvalid CustomizedAcceleratorValidateResultState = "Invalid"
	// CustomizedAcceleratorValidateResultStateValid - Customized accelerator properties are valid.
	CustomizedAcceleratorValidateResultStateValid CustomizedAcceleratorValidateResultState = "Valid"
)

// PossibleCustomizedAcceleratorValidateResultStateValues returns the possible values for the CustomizedAcceleratorValidateResultState const type.
func PossibleCustomizedAcceleratorValidateResultStateValues() []CustomizedAcceleratorValidateResultState {
	return []CustomizedAcceleratorValidateResultState{
		CustomizedAcceleratorValidateResultStateInvalid,
		CustomizedAcceleratorValidateResultStateValid,
	}
}

// DeploymentResourceProvisioningState - Provisioning state of the Deployment
type DeploymentResourceProvisioningState string

const (
	DeploymentResourceProvisioningStateCreating  DeploymentResourceProvisioningState = "Creating"
	DeploymentResourceProvisioningStateFailed    DeploymentResourceProvisioningState = "Failed"
	DeploymentResourceProvisioningStateSucceeded DeploymentResourceProvisioningState = "Succeeded"
	DeploymentResourceProvisioningStateUpdating  DeploymentResourceProvisioningState = "Updating"
)

// PossibleDeploymentResourceProvisioningStateValues returns the possible values for the DeploymentResourceProvisioningState const type.
func PossibleDeploymentResourceProvisioningStateValues() []DeploymentResourceProvisioningState {
	return []DeploymentResourceProvisioningState{
		DeploymentResourceProvisioningStateCreating,
		DeploymentResourceProvisioningStateFailed,
		DeploymentResourceProvisioningStateSucceeded,
		DeploymentResourceProvisioningStateUpdating,
	}
}

// DeploymentResourceStatus - Status of the Deployment
type DeploymentResourceStatus string

const (
	DeploymentResourceStatusRunning DeploymentResourceStatus = "Running"
	DeploymentResourceStatusStopped DeploymentResourceStatus = "Stopped"
)

// PossibleDeploymentResourceStatusValues returns the possible values for the DeploymentResourceStatus const type.
func PossibleDeploymentResourceStatusValues() []DeploymentResourceStatus {
	return []DeploymentResourceStatus{
		DeploymentResourceStatusRunning,
		DeploymentResourceStatusStopped,
	}
}

// DevToolPortalFeatureState - State of the plugin
type DevToolPortalFeatureState string

const (
	// DevToolPortalFeatureStateDisabled - Disable the plugin in Dev Tool Portal.
	DevToolPortalFeatureStateDisabled DevToolPortalFeatureState = "Disabled"
	// DevToolPortalFeatureStateEnabled - Enable the plugin in Dev Tool Portal.
	DevToolPortalFeatureStateEnabled DevToolPortalFeatureState = "Enabled"
)

// PossibleDevToolPortalFeatureStateValues returns the possible values for the DevToolPortalFeatureState const type.
func PossibleDevToolPortalFeatureStateValues() []DevToolPortalFeatureState {
	return []DevToolPortalFeatureState{
		DevToolPortalFeatureStateDisabled,
		DevToolPortalFeatureStateEnabled,
	}
}

// DevToolPortalProvisioningState - State of the Dev Tool Portal.
type DevToolPortalProvisioningState string

const (
	DevToolPortalProvisioningStateCanceled  DevToolPortalProvisioningState = "Canceled"
	DevToolPortalProvisioningStateCreating  DevToolPortalProvisioningState = "Creating"
	DevToolPortalProvisioningStateDeleting  DevToolPortalProvisioningState = "Deleting"
	DevToolPortalProvisioningStateFailed    DevToolPortalProvisioningState = "Failed"
	DevToolPortalProvisioningStateSucceeded DevToolPortalProvisioningState = "Succeeded"
	DevToolPortalProvisioningStateUpdating  DevToolPortalProvisioningState = "Updating"
)

// PossibleDevToolPortalProvisioningStateValues returns the possible values for the DevToolPortalProvisioningState const type.
func PossibleDevToolPortalProvisioningStateValues() []DevToolPortalProvisioningState {
	return []DevToolPortalProvisioningState{
		DevToolPortalProvisioningStateCanceled,
		DevToolPortalProvisioningStateCreating,
		DevToolPortalProvisioningStateDeleting,
		DevToolPortalProvisioningStateFailed,
		DevToolPortalProvisioningStateSucceeded,
		DevToolPortalProvisioningStateUpdating,
	}
}

// GatewayProvisioningState - State of the Spring Cloud Gateway.
type GatewayProvisioningState string

const (
	GatewayProvisioningStateCreating  GatewayProvisioningState = "Creating"
	GatewayProvisioningStateDeleting  GatewayProvisioningState = "Deleting"
	GatewayProvisioningStateFailed    GatewayProvisioningState = "Failed"
	GatewayProvisioningStateSucceeded GatewayProvisioningState = "Succeeded"
	GatewayProvisioningStateUpdating  GatewayProvisioningState = "Updating"
)

// PossibleGatewayProvisioningStateValues returns the possible values for the GatewayProvisioningState const type.
func PossibleGatewayProvisioningStateValues() []GatewayProvisioningState {
	return []GatewayProvisioningState{
		GatewayProvisioningStateCreating,
		GatewayProvisioningStateDeleting,
		GatewayProvisioningStateFailed,
		GatewayProvisioningStateSucceeded,
		GatewayProvisioningStateUpdating,
	}
}

// GatewayRouteConfigProtocol - Protocol of routed Azure Spring Apps applications.
type GatewayRouteConfigProtocol string

const (
	GatewayRouteConfigProtocolHTTP  GatewayRouteConfigProtocol = "HTTP"
	GatewayRouteConfigProtocolHTTPS GatewayRouteConfigProtocol = "HTTPS"
)

// PossibleGatewayRouteConfigProtocolValues returns the possible values for the GatewayRouteConfigProtocol const type.
func PossibleGatewayRouteConfigProtocolValues() []GatewayRouteConfigProtocol {
	return []GatewayRouteConfigProtocol{
		GatewayRouteConfigProtocolHTTP,
		GatewayRouteConfigProtocolHTTPS,
	}
}

// HTTPSchemeType - Scheme to use for connecting to the host. Defaults to HTTP.
// Possible enum values:
// * "HTTP" means that the scheme used will be http://
// * "HTTPS" means that the scheme used will be https://
type HTTPSchemeType string

const (
	HTTPSchemeTypeHTTP  HTTPSchemeType = "HTTP"
	HTTPSchemeTypeHTTPS HTTPSchemeType = "HTTPS"
)

// PossibleHTTPSchemeTypeValues returns the possible values for the HTTPSchemeType const type.
func PossibleHTTPSchemeTypeValues() []HTTPSchemeType {
	return []HTTPSchemeType{
		HTTPSchemeTypeHTTP,
		HTTPSchemeTypeHTTPS,
	}
}

// KPackBuildStageProvisioningState - The provisioning state of this build stage resource.
type KPackBuildStageProvisioningState string

const (
	KPackBuildStageProvisioningStateFailed     KPackBuildStageProvisioningState = "Failed"
	KPackBuildStageProvisioningStateNotStarted KPackBuildStageProvisioningState = "NotStarted"
	KPackBuildStageProvisioningStateRunning    KPackBuildStageProvisioningState = "Running"
	KPackBuildStageProvisioningStateSucceeded  KPackBuildStageProvisioningState = "Succeeded"
)

// PossibleKPackBuildStageProvisioningStateValues returns the possible values for the KPackBuildStageProvisioningState const type.
func PossibleKPackBuildStageProvisioningStateValues() []KPackBuildStageProvisioningState {
	return []KPackBuildStageProvisioningState{
		KPackBuildStageProvisioningStateFailed,
		KPackBuildStageProvisioningStateNotStarted,
		KPackBuildStageProvisioningStateRunning,
		KPackBuildStageProvisioningStateSucceeded,
	}
}

// LastModifiedByType - The type of identity that last modified the resource.
type LastModifiedByType string

const (
	LastModifiedByTypeApplication     LastModifiedByType = "Application"
	LastModifiedByTypeKey             LastModifiedByType = "Key"
	LastModifiedByTypeManagedIdentity LastModifiedByType = "ManagedIdentity"
	LastModifiedByTypeUser            LastModifiedByType = "User"
)

// PossibleLastModifiedByTypeValues returns the possible values for the LastModifiedByType const type.
func PossibleLastModifiedByTypeValues() []LastModifiedByType {
	return []LastModifiedByType{
		LastModifiedByTypeApplication,
		LastModifiedByTypeKey,
		LastModifiedByTypeManagedIdentity,
		LastModifiedByTypeUser,
	}
}

// ManagedIdentityType - Type of the managed identity
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                       ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAssigned             ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeSystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	ManagedIdentityTypeUserAssigned               ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeSystemAssignedUserAssigned,
		ManagedIdentityTypeUserAssigned,
	}
}

// MonitoringSettingState - State of the Monitoring Setting.
type MonitoringSettingState string

const (
	MonitoringSettingStateFailed       MonitoringSettingState = "Failed"
	MonitoringSettingStateNotAvailable MonitoringSettingState = "NotAvailable"
	MonitoringSettingStateSucceeded    MonitoringSettingState = "Succeeded"
	MonitoringSettingStateUpdating     MonitoringSettingState = "Updating"
)

// PossibleMonitoringSettingStateValues returns the possible values for the MonitoringSettingState const type.
func PossibleMonitoringSettingStateValues() []MonitoringSettingState {
	return []MonitoringSettingState{
		MonitoringSettingStateFailed,
		MonitoringSettingStateNotAvailable,
		MonitoringSettingStateSucceeded,
		MonitoringSettingStateUpdating,
	}
}

// PowerState - Power state of the Service
type PowerState string

const (
	PowerStateRunning PowerState = "Running"
	PowerStateStopped PowerState = "Stopped"
)

// PossiblePowerStateValues returns the possible values for the PowerState const type.
func PossiblePowerStateValues() []PowerState {
	return []PowerState{
		PowerStateRunning,
		PowerStateStopped,
	}
}

// PredefinedAcceleratorProvisioningState - Provisioning state of the predefined accelerator.
type PredefinedAcceleratorProvisioningState string

const (
	PredefinedAcceleratorProvisioningStateCreating  PredefinedAcceleratorProvisioningState = "Creating"
	PredefinedAcceleratorProvisioningStateFailed    PredefinedAcceleratorProvisioningState = "Failed"
	PredefinedAcceleratorProvisioningStateSucceeded PredefinedAcceleratorProvisioningState = "Succeeded"
	PredefinedAcceleratorProvisioningStateUpdating  PredefinedAcceleratorProvisioningState = "Updating"
)

// PossiblePredefinedAcceleratorProvisioningStateValues returns the possible values for the PredefinedAcceleratorProvisioningState const type.
func PossiblePredefinedAcceleratorProvisioningStateValues() []PredefinedAcceleratorProvisioningState {
	return []PredefinedAcceleratorProvisioningState{
		PredefinedAcceleratorProvisioningStateCreating,
		PredefinedAcceleratorProvisioningStateFailed,
		PredefinedAcceleratorProvisioningStateSucceeded,
		PredefinedAcceleratorProvisioningStateUpdating,
	}
}

// PredefinedAcceleratorState - State of the predefined accelerator.
type PredefinedAcceleratorState string

const (
	// PredefinedAcceleratorStateDisabled - Disable the predefined accelerator.
	PredefinedAcceleratorStateDisabled PredefinedAcceleratorState = "Disabled"
	// PredefinedAcceleratorStateEnabled - Enable the predefined accelerator.
	PredefinedAcceleratorStateEnabled PredefinedAcceleratorState = "Enabled"
)

// PossiblePredefinedAcceleratorStateValues returns the possible values for the PredefinedAcceleratorState const type.
func PossiblePredefinedAcceleratorStateValues() []PredefinedAcceleratorState {
	return []PredefinedAcceleratorState{
		PredefinedAcceleratorStateDisabled,
		PredefinedAcceleratorStateEnabled,
	}
}

// ProbeActionType - The type of the action to take to perform the health check.
type ProbeActionType string

const (
	ProbeActionTypeExecAction      ProbeActionType = "ExecAction"
	ProbeActionTypeHTTPGetAction   ProbeActionType = "HTTPGetAction"
	ProbeActionTypeTCPSocketAction ProbeActionType = "TCPSocketAction"
)

// PossibleProbeActionTypeValues returns the possible values for the ProbeActionType const type.
func PossibleProbeActionTypeValues() []ProbeActionType {
	return []ProbeActionType{
		ProbeActionTypeExecAction,
		ProbeActionTypeHTTPGetAction,
		ProbeActionTypeTCPSocketAction,
	}
}

// ProvisioningState - Provisioning state of the Service
type ProvisioningState string

const (
	ProvisioningStateCreating   ProvisioningState = "Creating"
	ProvisioningStateDeleted    ProvisioningState = "Deleted"
	ProvisioningStateDeleting   ProvisioningState = "Deleting"
	ProvisioningStateFailed     ProvisioningState = "Failed"
	ProvisioningStateMoveFailed ProvisioningState = "MoveFailed"
	ProvisioningStateMoved      ProvisioningState = "Moved"
	ProvisioningStateMoving     ProvisioningState = "Moving"
	ProvisioningStateStarting   ProvisioningState = "Starting"
	ProvisioningStateStopping   ProvisioningState = "Stopping"
	ProvisioningStateSucceeded  ProvisioningState = "Succeeded"
	ProvisioningStateUpdating   ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoveFailed,
		ProvisioningStateMoved,
		ProvisioningStateMoving,
		ProvisioningStateStarting,
		ProvisioningStateStopping,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ResourceSKURestrictionsReasonCode - Gets the reason for restriction. Possible values include: 'QuotaId', 'NotAvailableForSubscription'
type ResourceSKURestrictionsReasonCode string

const (
	ResourceSKURestrictionsReasonCodeNotAvailableForSubscription ResourceSKURestrictionsReasonCode = "NotAvailableForSubscription"
	ResourceSKURestrictionsReasonCodeQuotaID                     ResourceSKURestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSKURestrictionsReasonCodeValues returns the possible values for the ResourceSKURestrictionsReasonCode const type.
func PossibleResourceSKURestrictionsReasonCodeValues() []ResourceSKURestrictionsReasonCode {
	return []ResourceSKURestrictionsReasonCode{
		ResourceSKURestrictionsReasonCodeNotAvailableForSubscription,
		ResourceSKURestrictionsReasonCodeQuotaID,
	}
}

// ResourceSKURestrictionsType - Gets the type of restrictions. Possible values include: 'Location', 'Zone'
type ResourceSKURestrictionsType string

const (
	ResourceSKURestrictionsTypeLocation ResourceSKURestrictionsType = "Location"
	ResourceSKURestrictionsTypeZone     ResourceSKURestrictionsType = "Zone"
)

// PossibleResourceSKURestrictionsTypeValues returns the possible values for the ResourceSKURestrictionsType const type.
func PossibleResourceSKURestrictionsTypeValues() []ResourceSKURestrictionsType {
	return []ResourceSKURestrictionsType{
		ResourceSKURestrictionsTypeLocation,
		ResourceSKURestrictionsTypeZone,
	}
}

// SKUScaleType - Gets or sets the type of the scale.
type SKUScaleType string

const (
	SKUScaleTypeAutomatic SKUScaleType = "Automatic"
	SKUScaleTypeManual    SKUScaleType = "Manual"
	SKUScaleTypeNone      SKUScaleType = "None"
)

// PossibleSKUScaleTypeValues returns the possible values for the SKUScaleType const type.
func PossibleSKUScaleTypeValues() []SKUScaleType {
	return []SKUScaleType{
		SKUScaleTypeAutomatic,
		SKUScaleTypeManual,
		SKUScaleTypeNone,
	}
}

// ServiceRegistryProvisioningState - State of the Service Registry.
type ServiceRegistryProvisioningState string

const (
	ServiceRegistryProvisioningStateCreating  ServiceRegistryProvisioningState = "Creating"
	ServiceRegistryProvisioningStateDeleting  ServiceRegistryProvisioningState = "Deleting"
	ServiceRegistryProvisioningStateFailed    ServiceRegistryProvisioningState = "Failed"
	ServiceRegistryProvisioningStateSucceeded ServiceRegistryProvisioningState = "Succeeded"
	ServiceRegistryProvisioningStateUpdating  ServiceRegistryProvisioningState = "Updating"
)

// PossibleServiceRegistryProvisioningStateValues returns the possible values for the ServiceRegistryProvisioningState const type.
func PossibleServiceRegistryProvisioningStateValues() []ServiceRegistryProvisioningState {
	return []ServiceRegistryProvisioningState{
		ServiceRegistryProvisioningStateCreating,
		ServiceRegistryProvisioningStateDeleting,
		ServiceRegistryProvisioningStateFailed,
		ServiceRegistryProvisioningStateSucceeded,
		ServiceRegistryProvisioningStateUpdating,
	}
}

// SessionAffinity - Type of the affinity, set this to Cookie to enable session affinity.
type SessionAffinity string

const (
	SessionAffinityCookie SessionAffinity = "Cookie"
	SessionAffinityNone   SessionAffinity = "None"
)

// PossibleSessionAffinityValues returns the possible values for the SessionAffinity const type.
func PossibleSessionAffinityValues() []SessionAffinity {
	return []SessionAffinity{
		SessionAffinityCookie,
		SessionAffinityNone,
	}
}

// StorageType - The type of the storage.
type StorageType string

const (
	StorageTypeStorageAccount StorageType = "StorageAccount"
)

// PossibleStorageTypeValues returns the possible values for the StorageType const type.
func PossibleStorageTypeValues() []StorageType {
	return []StorageType{
		StorageTypeStorageAccount,
	}
}

// SupportedRuntimePlatform - The platform of this runtime version (possible values: "Java" or ".NET").
type SupportedRuntimePlatform string

const (
	SupportedRuntimePlatformJava    SupportedRuntimePlatform = "Java"
	SupportedRuntimePlatformNETCore SupportedRuntimePlatform = ".NET Core"
)

// PossibleSupportedRuntimePlatformValues returns the possible values for the SupportedRuntimePlatform const type.
func PossibleSupportedRuntimePlatformValues() []SupportedRuntimePlatform {
	return []SupportedRuntimePlatform{
		SupportedRuntimePlatformJava,
		SupportedRuntimePlatformNETCore,
	}
}

// SupportedRuntimeValue - The raw value which could be passed to deployment CRUD operations.
type SupportedRuntimeValue string

const (
	SupportedRuntimeValueJava11    SupportedRuntimeValue = "Java_11"
	SupportedRuntimeValueJava17    SupportedRuntimeValue = "Java_17"
	SupportedRuntimeValueJava8     SupportedRuntimeValue = "Java_8"
	SupportedRuntimeValueNetCore31 SupportedRuntimeValue = "NetCore_31"
)

// PossibleSupportedRuntimeValueValues returns the possible values for the SupportedRuntimeValue const type.
func PossibleSupportedRuntimeValueValues() []SupportedRuntimeValue {
	return []SupportedRuntimeValue{
		SupportedRuntimeValueJava11,
		SupportedRuntimeValueJava17,
		SupportedRuntimeValueJava8,
		SupportedRuntimeValueNetCore31,
	}
}

// TestKeyType - Type of the test key
type TestKeyType string

const (
	TestKeyTypePrimary   TestKeyType = "Primary"
	TestKeyTypeSecondary TestKeyType = "Secondary"
)

// PossibleTestKeyTypeValues returns the possible values for the TestKeyType const type.
func PossibleTestKeyTypeValues() []TestKeyType {
	return []TestKeyType{
		TestKeyTypePrimary,
		TestKeyTypeSecondary,
	}
}

// TrafficDirection - The direction of required traffic
type TrafficDirection string

const (
	TrafficDirectionInbound  TrafficDirection = "Inbound"
	TrafficDirectionOutbound TrafficDirection = "Outbound"
)

// PossibleTrafficDirectionValues returns the possible values for the TrafficDirection const type.
func PossibleTrafficDirectionValues() []TrafficDirection {
	return []TrafficDirection{
		TrafficDirectionInbound,
		TrafficDirectionOutbound,
	}
}

// Type - The type of the underlying resource to mount as a persistent disk.
type Type string

const (
	TypeAzureFileVolume Type = "AzureFileVolume"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeAzureFileVolume,
	}
}
