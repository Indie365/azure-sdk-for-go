// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdevopsinfrastructure

// AgentProfileClassification provides polymorphic access to related types.
// Call the interface's GetAgentProfile() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AgentProfile, *Stateful, *StatelessAgentProfile
type AgentProfileClassification interface {
	// GetAgentProfile returns the AgentProfile content of the underlying type.
	GetAgentProfile() *AgentProfile
}

// FabricProfileClassification provides polymorphic access to related types.
// Call the interface's GetFabricProfile() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *FabricProfile, *VmssFabricProfile
type FabricProfileClassification interface {
	// GetFabricProfile returns the FabricProfile content of the underlying type.
	GetFabricProfile() *FabricProfile
}

// OrganizationProfileClassification provides polymorphic access to related types.
// Call the interface's GetOrganizationProfile() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AzureDevOpsOrganizationProfile, *GitHubOrganizationProfile, *OrganizationProfile
type OrganizationProfileClassification interface {
	// GetOrganizationProfile returns the OrganizationProfile content of the underlying type.
	GetOrganizationProfile() *OrganizationProfile
}

// ResourcePredictionsProfileClassification provides polymorphic access to related types.
// Call the interface's GetResourcePredictionsProfile() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *AutomaticResourcePredictionsProfile, *ManualResourcePredictionsProfile, *ResourcePredictionsProfile
type ResourcePredictionsProfileClassification interface {
	// GetResourcePredictionsProfile returns the ResourcePredictionsProfile content of the underlying type.
	GetResourcePredictionsProfile() *ResourcePredictionsProfile
}
