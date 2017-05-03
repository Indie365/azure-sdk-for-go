package devtestlabs

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// CustomImageOsType enumerates the values for custom image os type.
type CustomImageOsType string

const (
	// Linux specifies the linux state for custom image os type.
	Linux CustomImageOsType = "Linux"
	// None specifies the none state for custom image os type.
	None CustomImageOsType = "None"
	// Windows specifies the windows state for custom image os type.
	Windows CustomImageOsType = "Windows"
)

// EnableStatus enumerates the values for enable status.
type EnableStatus string

const (
	// Disabled specifies the disabled state for enable status.
	Disabled EnableStatus = "Disabled"
	// Enabled specifies the enabled state for enable status.
	Enabled EnableStatus = "Enabled"
)

// LabCostType enumerates the values for lab cost type.
type LabCostType string

const (
	// Projected specifies the projected state for lab cost type.
	Projected LabCostType = "Projected"
	// Reported specifies the reported state for lab cost type.
	Reported LabCostType = "Reported"
	// Unavailable specifies the unavailable state for lab cost type.
	Unavailable LabCostType = "Unavailable"
)

// LabStorageType enumerates the values for lab storage type.
type LabStorageType string

const (
	// Premium specifies the premium state for lab storage type.
	Premium LabStorageType = "Premium"
	// Standard specifies the standard state for lab storage type.
	Standard LabStorageType = "Standard"
)

// LinuxOsState enumerates the values for linux os state.
type LinuxOsState string

const (
	// DeprovisionApplied specifies the deprovision applied state for linux os
	// state.
	DeprovisionApplied LinuxOsState = "DeprovisionApplied"
	// DeprovisionRequested specifies the deprovision requested state for linux
	// os state.
	DeprovisionRequested LinuxOsState = "DeprovisionRequested"
	// NonDeprovisioned specifies the non deprovisioned state for linux os
	// state.
	NonDeprovisioned LinuxOsState = "NonDeprovisioned"
)

// PolicyEvaluatorType enumerates the values for policy evaluator type.
type PolicyEvaluatorType string

const (
	// AllowedValuesPolicy specifies the allowed values policy state for policy
	// evaluator type.
	AllowedValuesPolicy PolicyEvaluatorType = "AllowedValuesPolicy"
	// MaxValuePolicy specifies the max value policy state for policy evaluator
	// type.
	MaxValuePolicy PolicyEvaluatorType = "MaxValuePolicy"
)

// PolicyFactName enumerates the values for policy fact name.
type PolicyFactName string

const (
	// PolicyFactNameGalleryImage specifies the policy fact name gallery image
	// state for policy fact name.
	PolicyFactNameGalleryImage PolicyFactName = "GalleryImage"
	// PolicyFactNameLabVMCount specifies the policy fact name lab vm count
	// state for policy fact name.
	PolicyFactNameLabVMCount PolicyFactName = "LabVmCount"
	// PolicyFactNameLabVMSize specifies the policy fact name lab vm size state
	// for policy fact name.
	PolicyFactNameLabVMSize PolicyFactName = "LabVmSize"
	// PolicyFactNameUserOwnedLabVMCount specifies the policy fact name user
	// owned lab vm count state for policy fact name.
	PolicyFactNameUserOwnedLabVMCount PolicyFactName = "UserOwnedLabVmCount"
	// PolicyFactNameUserOwnedLabVMCountInSubnet specifies the policy fact name
	// user owned lab vm count in subnet state for policy fact name.
	PolicyFactNameUserOwnedLabVMCountInSubnet PolicyFactName = "UserOwnedLabVmCountInSubnet"
)

// PolicyStatus enumerates the values for policy status.
type PolicyStatus string

const (
	// PolicyStatusDisabled specifies the policy status disabled state for
	// policy status.
	PolicyStatusDisabled PolicyStatus = "Disabled"
	// PolicyStatusEnabled specifies the policy status enabled state for policy
	// status.
	PolicyStatusEnabled PolicyStatus = "Enabled"
)

// SourceControlType enumerates the values for source control type.
type SourceControlType string

const (
	// GitHub specifies the git hub state for source control type.
	GitHub SourceControlType = "GitHub"
	// VsoGit specifies the vso git state for source control type.
	VsoGit SourceControlType = "VsoGit"
)

// SubscriptionNotificationState enumerates the values for subscription
// notification state.
type SubscriptionNotificationState string

const (
	// Deleted specifies the deleted state for subscription notification state.
	Deleted SubscriptionNotificationState = "Deleted"
	// NotDefined specifies the not defined state for subscription notification
	// state.
	NotDefined SubscriptionNotificationState = "NotDefined"
	// Registered specifies the registered state for subscription notification
	// state.
	Registered SubscriptionNotificationState = "Registered"
	// Suspended specifies the suspended state for subscription notification
	// state.
	Suspended SubscriptionNotificationState = "Suspended"
	// Unregistered specifies the unregistered state for subscription
	// notification state.
	Unregistered SubscriptionNotificationState = "Unregistered"
	// Warned specifies the warned state for subscription notification state.
	Warned SubscriptionNotificationState = "Warned"
)

// UsagePermissionType enumerates the values for usage permission type.
type UsagePermissionType string

const (
	// Allow specifies the allow state for usage permission type.
	Allow UsagePermissionType = "Allow"
	// Default specifies the default state for usage permission type.
	Default UsagePermissionType = "Default"
	// Deny specifies the deny state for usage permission type.
	Deny UsagePermissionType = "Deny"
)

// WindowsOsState enumerates the values for windows os state.
type WindowsOsState string

const (
	// NonSysprepped specifies the non sysprepped state for windows os state.
	NonSysprepped WindowsOsState = "NonSysprepped"
	// SysprepApplied specifies the sysprep applied state for windows os state.
	SysprepApplied WindowsOsState = "SysprepApplied"
	// SysprepRequested specifies the sysprep requested state for windows os
	// state.
	SysprepRequested WindowsOsState = "SysprepRequested"
)

// ApplyArtifactsRequest is request body for applying artifacts to a virtual
// machine.
type ApplyArtifactsRequest struct {
	Artifacts *[]ArtifactInstallProperties `json:"artifacts,omitempty"`
}

// ArmTemplateInfo is information about a generated ARM template.
type ArmTemplateInfo struct {
	autorest.Response `json:"-"`
	Template          *map[string]interface{} `json:"template,omitempty"`
	Parameters        *map[string]interface{} `json:"parameters,omitempty"`
}

// Artifact is an artifact.
type Artifact struct {
	autorest.Response   `json:"-"`
	*ArtifactProperties `json:"properties,omitempty"`
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Location            *string             `json:"location,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
}

// ArtifactDeploymentStatusProperties is properties of an artifact deployment.
type ArtifactDeploymentStatusProperties struct {
	DeploymentStatus *string `json:"deploymentStatus,omitempty"`
	ArtifactsApplied *int32  `json:"artifactsApplied,omitempty"`
	TotalArtifacts   *int32  `json:"totalArtifacts,omitempty"`
}

// ArtifactInstallProperties is properties of an artifact.
type ArtifactInstallProperties struct {
	ArtifactID *string                        `json:"artifactId,omitempty"`
	Parameters *[]ArtifactParameterProperties `json:"parameters,omitempty"`
}

// ArtifactParameterProperties is properties of an artifact parameter.
type ArtifactParameterProperties struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// ArtifactProperties is properties of an artifact.
type ArtifactProperties struct {
	Title        *string                 `json:"title,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	FilePath     *string                 `json:"filePath,omitempty"`
	Icon         *string                 `json:"icon,omitempty"`
	TargetOsType *string                 `json:"targetOsType,omitempty"`
	Parameters   *map[string]interface{} `json:"parameters,omitempty"`
}

// ArtifactSource is properties of an artifact source.
type ArtifactSource struct {
	autorest.Response         `json:"-"`
	*ArtifactSourceProperties `json:"properties,omitempty"`
	ID                        *string             `json:"id,omitempty"`
	Name                      *string             `json:"name,omitempty"`
	Type                      *string             `json:"type,omitempty"`
	Location                  *string             `json:"location,omitempty"`
	Tags                      *map[string]*string `json:"tags,omitempty"`
}

// ArtifactSourceProperties is properties of an artifact source.
type ArtifactSourceProperties struct {
	DisplayName       *string           `json:"displayName,omitempty"`
	URI               *string           `json:"uri,omitempty"`
	SourceType        SourceControlType `json:"sourceType,omitempty"`
	FolderPath        *string           `json:"folderPath,omitempty"`
	BranchRef         *string           `json:"branchRef,omitempty"`
	SecurityToken     *string           `json:"securityToken,omitempty"`
	Status            EnableStatus      `json:"status,omitempty"`
	ProvisioningState *string           `json:"provisioningState,omitempty"`
	UniqueIdentifier  *string           `json:"uniqueIdentifier,omitempty"`
}

// CloudError is
type CloudError struct {
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody is
type CloudErrorBody struct {
	Code    *string           `json:"code,omitempty"`
	Message *string           `json:"message,omitempty"`
	Target  *string           `json:"target,omitempty"`
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// Cost is a cost item.
type Cost struct {
	autorest.Response `json:"-"`
	*CostProperties   `json:"properties,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
}

// CostPerDayProperties is the properties of a lab cost item.
type CostPerDayProperties struct {
	Date     *date.Time  `json:"date,omitempty"`
	Cost     *float64    `json:"cost,omitempty"`
	CostType LabCostType `json:"costType,omitempty"`
}

// CostProperties is properties of a cost item.
type CostProperties struct {
	CurrencyCode  *string                   `json:"currencyCode,omitempty"`
	Costs         *[]CostPerDayProperties   `json:"costs,omitempty"`
	ResourceCosts *[]ResourceCostProperties `json:"resourceCosts,omitempty"`
}

// CustomImage is a custom image.
type CustomImage struct {
	autorest.Response      `json:"-"`
	*CustomImageProperties `json:"properties,omitempty"`
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Location               *string             `json:"location,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
}

// CustomImageProperties is properties of a custom image.
type CustomImageProperties struct {
	VM                *CustomImagePropertiesFromVM `json:"vm,omitempty"`
	Vhd               *CustomImagePropertiesCustom `json:"vhd,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	Author            *string                      `json:"author,omitempty"`
	CreationDate      *date.Time                   `json:"creationDate,omitempty"`
	ProvisioningState *string                      `json:"provisioningState,omitempty"`
	UniqueIdentifier  *string                      `json:"uniqueIdentifier,omitempty"`
}

// CustomImagePropertiesCustom is properties for creating a custom image from a
// VHD.
type CustomImagePropertiesCustom struct {
	ImageName *string           `json:"imageName,omitempty"`
	SysPrep   *bool             `json:"sysPrep,omitempty"`
	OsType    CustomImageOsType `json:"osType,omitempty"`
}

// CustomImagePropertiesFromVM is properties for creating a custom image from a
// virtual machine.
type CustomImagePropertiesFromVM struct {
	SourceVMID    *string        `json:"sourceVmId,omitempty"`
	WindowsOsInfo *WindowsOsInfo `json:"windowsOsInfo,omitempty"`
	LinuxOsInfo   *LinuxOsInfo   `json:"linuxOsInfo,omitempty"`
}

// DayDetails is properties of a daily schedule.
type DayDetails struct {
	Time *string `json:"time,omitempty"`
}

// EvaluatePoliciesProperties is properties for evaluating a policy set.
type EvaluatePoliciesProperties struct {
	FactName    *string `json:"factName,omitempty"`
	FactData    *string `json:"factData,omitempty"`
	ValueOffset *string `json:"valueOffset,omitempty"`
}

// EvaluatePoliciesRequest is request body for evaluating a policy set.
type EvaluatePoliciesRequest struct {
	Policies *[]EvaluatePoliciesProperties `json:"policies,omitempty"`
}

// EvaluatePoliciesResponse is response body for evaluating a policy set.
type EvaluatePoliciesResponse struct {
	autorest.Response `json:"-"`
	Results           *[]PolicySetResult `json:"results,omitempty"`
}

// Formula is a formula.
type Formula struct {
	autorest.Response  `json:"-"`
	*FormulaProperties `json:"properties,omitempty"`
	ID                 *string             `json:"id,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Type               *string             `json:"type,omitempty"`
	Location           *string             `json:"location,omitempty"`
	Tags               *map[string]*string `json:"tags,omitempty"`
}

// FormulaProperties is properties of a formula.
type FormulaProperties struct {
	Description       *string                  `json:"description,omitempty"`
	Author            *string                  `json:"author,omitempty"`
	OsType            *string                  `json:"osType,omitempty"`
	CreationDate      *date.Time               `json:"creationDate,omitempty"`
	FormulaContent    *LabVirtualMachine       `json:"formulaContent,omitempty"`
	VM                *FormulaPropertiesFromVM `json:"vm,omitempty"`
	ProvisioningState *string                  `json:"provisioningState,omitempty"`
	UniqueIdentifier  *string                  `json:"uniqueIdentifier,omitempty"`
}

// FormulaPropertiesFromVM is information about a VM from which a formula is to
// be created.
type FormulaPropertiesFromVM struct {
	LabVMID *string `json:"labVmId,omitempty"`
}

// GalleryImage is a gallery image.
type GalleryImage struct {
	*GalleryImageProperties `json:"properties,omitempty"`
	ID                      *string             `json:"id,omitempty"`
	Name                    *string             `json:"name,omitempty"`
	Type                    *string             `json:"type,omitempty"`
	Location                *string             `json:"location,omitempty"`
	Tags                    *map[string]*string `json:"tags,omitempty"`
}

// GalleryImageProperties is properties of a gallery image.
type GalleryImageProperties struct {
	Author         *string                `json:"author,omitempty"`
	CreatedDate    *date.Time             `json:"createdDate,omitempty"`
	Description    *string                `json:"description,omitempty"`
	ImageReference *GalleryImageReference `json:"imageReference,omitempty"`
	Icon           *string                `json:"icon,omitempty"`
	Enabled        *bool                  `json:"enabled,omitempty"`
}

// GalleryImageReference is the reference information for an Azure Marketplace
// image.
type GalleryImageReference struct {
	Offer     *string `json:"offer,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	OsType    *string `json:"osType,omitempty"`
	Version   *string `json:"version,omitempty"`
}

// GenerateArmTemplateRequest is parameters for generating an ARM template for
// deploying artifacts.
type GenerateArmTemplateRequest struct {
	VirtualMachineName *string          `json:"virtualMachineName,omitempty"`
	Parameters         *[]ParameterInfo `json:"parameters,omitempty"`
	Location           *string          `json:"location,omitempty"`
}

// GenerateUploadURIParameter is properties for generating an upload URI.
type GenerateUploadURIParameter struct {
	BlobName *string `json:"blobName,omitempty"`
}

// GenerateUploadURIResponse is reponse body for generating an upload URI.
type GenerateUploadURIResponse struct {
	autorest.Response `json:"-"`
	UploadURI         *string `json:"uploadUri,omitempty"`
}

// HourDetails is properties of an hourly schedule.
type HourDetails struct {
	Minute *int32 `json:"minute,omitempty"`
}

// Lab is a lab.
type Lab struct {
	autorest.Response `json:"-"`
	*LabProperties    `json:"properties,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
}

// LabProperties is properties of a lab.
type LabProperties struct {
	DefaultStorageAccount   *string        `json:"defaultStorageAccount,omitempty"`
	ArtifactsStorageAccount *string        `json:"artifactsStorageAccount,omitempty"`
	VaultName               *string        `json:"vaultName,omitempty"`
	LabStorageType          LabStorageType `json:"labStorageType,omitempty"`
	CreatedDate             *date.Time     `json:"createdDate,omitempty"`
	ProvisioningState       *string        `json:"provisioningState,omitempty"`
	UniqueIdentifier        *string        `json:"uniqueIdentifier,omitempty"`
}

// LabVhd is properties of a VHD in the lab.
type LabVhd struct {
	ID *string `json:"id,omitempty"`
}

// LabVirtualMachine is a virtual machine.
type LabVirtualMachine struct {
	autorest.Response            `json:"-"`
	*LabVirtualMachineProperties `json:"properties,omitempty"`
	ID                           *string             `json:"id,omitempty"`
	Name                         *string             `json:"name,omitempty"`
	Type                         *string             `json:"type,omitempty"`
	Location                     *string             `json:"location,omitempty"`
	Tags                         *map[string]*string `json:"tags,omitempty"`
}

// LabVirtualMachineProperties is properties of a virtual machine.
type LabVirtualMachineProperties struct {
	Notes                      *string                             `json:"notes,omitempty"`
	OwnerObjectID              *string                             `json:"ownerObjectId,omitempty"`
	CreatedByUserID            *string                             `json:"createdByUserId,omitempty"`
	CreatedByUser              *string                             `json:"createdByUser,omitempty"`
	ComputeID                  *string                             `json:"computeId,omitempty"`
	CustomImageID              *string                             `json:"customImageId,omitempty"`
	OsType                     *string                             `json:"osType,omitempty"`
	Size                       *string                             `json:"size,omitempty"`
	UserName                   *string                             `json:"userName,omitempty"`
	Password                   *string                             `json:"password,omitempty"`
	SSHKey                     *string                             `json:"sshKey,omitempty"`
	IsAuthenticationWithSSHKey *bool                               `json:"isAuthenticationWithSshKey,omitempty"`
	Fqdn                       *string                             `json:"fqdn,omitempty"`
	LabSubnetName              *string                             `json:"labSubnetName,omitempty"`
	LabVirtualNetworkID        *string                             `json:"labVirtualNetworkId,omitempty"`
	DisallowPublicIPAddress    *bool                               `json:"disallowPublicIpAddress,omitempty"`
	Artifacts                  *[]ArtifactInstallProperties        `json:"artifacts,omitempty"`
	ArtifactDeploymentStatus   *ArtifactDeploymentStatusProperties `json:"artifactDeploymentStatus,omitempty"`
	GalleryImageReference      *GalleryImageReference              `json:"galleryImageReference,omitempty"`
	ProvisioningState          *string                             `json:"provisioningState,omitempty"`
	UniqueIdentifier           *string                             `json:"uniqueIdentifier,omitempty"`
}

// LinuxOsInfo is information about a Linux OS.
type LinuxOsInfo struct {
	LinuxOsState LinuxOsState `json:"linuxOsState,omitempty"`
}

// ParameterInfo is
type ParameterInfo struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Policy is a Policy.
type Policy struct {
	autorest.Response `json:"-"`
	*PolicyProperties `json:"properties,omitempty"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
}

// PolicyProperties is properties of a Policy.
type PolicyProperties struct {
	Description       *string             `json:"description,omitempty"`
	Status            PolicyStatus        `json:"status,omitempty"`
	FactName          PolicyFactName      `json:"factName,omitempty"`
	FactData          *string             `json:"factData,omitempty"`
	Threshold         *string             `json:"threshold,omitempty"`
	EvaluatorType     PolicyEvaluatorType `json:"evaluatorType,omitempty"`
	ProvisioningState *string             `json:"provisioningState,omitempty"`
	UniqueIdentifier  *string             `json:"uniqueIdentifier,omitempty"`
}

// PolicySetResult is result of a policy set evaluation.
type PolicySetResult struct {
	HasError         *bool              `json:"hasError,omitempty"`
	PolicyViolations *[]PolicyViolation `json:"policyViolations,omitempty"`
}

// PolicyViolation is policy violation.
type PolicyViolation struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ResourceCostProperties is the properties of a resource cost item.
type ResourceCostProperties struct {
	Resourcename      *string  `json:"resourcename,omitempty"`
	ResourceGroupName *string  `json:"resourceGroupName,omitempty"`
	ResourceCost      *float64 `json:"resourceCost,omitempty"`
	Owner             *string  `json:"owner,omitempty"`
	Category          *string  `json:"category,omitempty"`
	Exists            *bool    `json:"exists,omitempty"`
	ResourceType      *string  `json:"resourceType,omitempty"`
}

// ResponseWithContinuationArtifact is the response of a list operation.
type ResponseWithContinuationArtifact struct {
	autorest.Response `json:"-"`
	Value             *[]Artifact `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// ResponseWithContinuationArtifactPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationArtifact) ResponseWithContinuationArtifactPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationArtifactSource is the response of a list operation.
type ResponseWithContinuationArtifactSource struct {
	autorest.Response `json:"-"`
	Value             *[]ArtifactSource `json:"value,omitempty"`
	NextLink          *string           `json:"nextLink,omitempty"`
}

// ResponseWithContinuationArtifactSourcePreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationArtifactSource) ResponseWithContinuationArtifactSourcePreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationCustomImage is the response of a list operation.
type ResponseWithContinuationCustomImage struct {
	autorest.Response `json:"-"`
	Value             *[]CustomImage `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
}

// ResponseWithContinuationCustomImagePreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationCustomImage) ResponseWithContinuationCustomImagePreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationFormula is the response of a list operation.
type ResponseWithContinuationFormula struct {
	autorest.Response `json:"-"`
	Value             *[]Formula `json:"value,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// ResponseWithContinuationFormulaPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationFormula) ResponseWithContinuationFormulaPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationGalleryImage is the response of a list operation.
type ResponseWithContinuationGalleryImage struct {
	autorest.Response `json:"-"`
	Value             *[]GalleryImage `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
}

// ResponseWithContinuationGalleryImagePreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationGalleryImage) ResponseWithContinuationGalleryImagePreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationLab is the response of a list operation.
type ResponseWithContinuationLab struct {
	autorest.Response `json:"-"`
	Value             *[]Lab  `json:"value,omitempty"`
	NextLink          *string `json:"nextLink,omitempty"`
}

// ResponseWithContinuationLabPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationLab) ResponseWithContinuationLabPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationLabVhd is the response of a list operation.
type ResponseWithContinuationLabVhd struct {
	autorest.Response `json:"-"`
	Value             *[]LabVhd `json:"value,omitempty"`
	NextLink          *string   `json:"nextLink,omitempty"`
}

// ResponseWithContinuationLabVhdPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationLabVhd) ResponseWithContinuationLabVhdPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationLabVirtualMachine is the response of a list
// operation.
type ResponseWithContinuationLabVirtualMachine struct {
	autorest.Response `json:"-"`
	Value             *[]LabVirtualMachine `json:"value,omitempty"`
	NextLink          *string              `json:"nextLink,omitempty"`
}

// ResponseWithContinuationLabVirtualMachinePreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationLabVirtualMachine) ResponseWithContinuationLabVirtualMachinePreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationPolicy is the response of a list operation.
type ResponseWithContinuationPolicy struct {
	autorest.Response `json:"-"`
	Value             *[]Policy `json:"value,omitempty"`
	NextLink          *string   `json:"nextLink,omitempty"`
}

// ResponseWithContinuationPolicyPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationPolicy) ResponseWithContinuationPolicyPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationSchedule is the response of a list operation.
type ResponseWithContinuationSchedule struct {
	autorest.Response `json:"-"`
	Value             *[]Schedule `json:"value,omitempty"`
	NextLink          *string     `json:"nextLink,omitempty"`
}

// ResponseWithContinuationSchedulePreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationSchedule) ResponseWithContinuationSchedulePreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ResponseWithContinuationVirtualNetwork is the response of a list operation.
type ResponseWithContinuationVirtualNetwork struct {
	autorest.Response `json:"-"`
	Value             *[]VirtualNetwork `json:"value,omitempty"`
	NextLink          *string           `json:"nextLink,omitempty"`
}

// ResponseWithContinuationVirtualNetworkPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ResponseWithContinuationVirtualNetwork) ResponseWithContinuationVirtualNetworkPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Schedule is a schedule.
type Schedule struct {
	autorest.Response   `json:"-"`
	*ScheduleProperties `json:"properties,omitempty"`
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Location            *string             `json:"location,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
}

// ScheduleProperties is properties of a schedule.
type ScheduleProperties struct {
	Status            EnableStatus `json:"status,omitempty"`
	TaskType          *string      `json:"taskType,omitempty"`
	WeeklyRecurrence  *WeekDetails `json:"weeklyRecurrence,omitempty"`
	DailyRecurrence   *DayDetails  `json:"dailyRecurrence,omitempty"`
	HourlyRecurrence  *HourDetails `json:"hourlyRecurrence,omitempty"`
	TimeZoneID        *string      `json:"timeZoneId,omitempty"`
	ProvisioningState *string      `json:"provisioningState,omitempty"`
	UniqueIdentifier  *string      `json:"uniqueIdentifier,omitempty"`
}

// Subnet is
type Subnet struct {
	ResourceID    *string             `json:"resourceId,omitempty"`
	LabSubnetName *string             `json:"labSubnetName,omitempty"`
	AllowPublicIP UsagePermissionType `json:"allowPublicIp,omitempty"`
}

// SubnetOverride is property overrides on a subnet of a virtual network.
type SubnetOverride struct {
	ResourceID                   *string             `json:"resourceId,omitempty"`
	LabSubnetName                *string             `json:"labSubnetName,omitempty"`
	UseInVMCreationPermission    UsagePermissionType `json:"useInVmCreationPermission,omitempty"`
	UsePublicIPAddressPermission UsagePermissionType `json:"usePublicIpAddressPermission,omitempty"`
}

// SubscriptionNotification is
type SubscriptionNotification struct {
	RegistrationDate *string                             `json:"registrationDate,omitempty"`
	State            SubscriptionNotificationState       `json:"state,omitempty"`
	Properties       *SubscriptionNotificationProperties `json:"properties,omitempty"`
}

// SubscriptionNotificationProperties is
type SubscriptionNotificationProperties struct {
	TenantID *string `json:"tenantId,omitempty"`
}

// VirtualNetwork is a virtual network.
type VirtualNetwork struct {
	autorest.Response         `json:"-"`
	*VirtualNetworkProperties `json:"properties,omitempty"`
	ID                        *string             `json:"id,omitempty"`
	Name                      *string             `json:"name,omitempty"`
	Type                      *string             `json:"type,omitempty"`
	Location                  *string             `json:"location,omitempty"`
	Tags                      *map[string]*string `json:"tags,omitempty"`
}

// VirtualNetworkProperties is properties of a virtual network.
type VirtualNetworkProperties struct {
	AllowedSubnets             *[]Subnet         `json:"allowedSubnets,omitempty"`
	Description                *string           `json:"description,omitempty"`
	ExternalProviderResourceID *string           `json:"externalProviderResourceId,omitempty"`
	SubnetOverrides            *[]SubnetOverride `json:"subnetOverrides,omitempty"`
	ProvisioningState          *string           `json:"provisioningState,omitempty"`
	UniqueIdentifier           *string           `json:"uniqueIdentifier,omitempty"`
}

// WeekDetails is properties of a weekly schedule.
type WeekDetails struct {
	Weekdays *[]string `json:"weekdays,omitempty"`
	Time     *string   `json:"time,omitempty"`
}

// WindowsOsInfo is information about a Windows OS.
type WindowsOsInfo struct {
	WindowsOsState WindowsOsState `json:"windowsOsState,omitempty"`
}
