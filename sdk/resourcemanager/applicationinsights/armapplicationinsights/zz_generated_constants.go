//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

const (
	moduleName    = "armapplicationinsights"
	moduleVersion = "v0.2.1"
)

// ApplicationType - Type of application being monitored.
type ApplicationType string

const (
	ApplicationTypeOther ApplicationType = "other"
	ApplicationTypeWeb   ApplicationType = "web"
)

// PossibleApplicationTypeValues returns the possible values for the ApplicationType const type.
func PossibleApplicationTypeValues() []ApplicationType {
	return []ApplicationType{
		ApplicationTypeOther,
		ApplicationTypeWeb,
	}
}

// ToPtr returns a *ApplicationType pointing to the current value.
func (c ApplicationType) ToPtr() *ApplicationType {
	return &c
}

type CategoryType string

const (
	CategoryTypePerformance CategoryType = "performance"
	CategoryTypeRetention   CategoryType = "retention"
	CategoryTypeTSG         CategoryType = "TSG"
	CategoryTypeWorkbook    CategoryType = "workbook"
)

// PossibleCategoryTypeValues returns the possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{
		CategoryTypePerformance,
		CategoryTypeRetention,
		CategoryTypeTSG,
		CategoryTypeWorkbook,
	}
}

// ToPtr returns a *CategoryType pointing to the current value.
func (c CategoryType) ToPtr() *CategoryType {
	return &c
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

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

type FavoriteSourceType string

const (
	FavoriteSourceTypeEvents       FavoriteSourceType = "events"
	FavoriteSourceTypeFunnel       FavoriteSourceType = "funnel"
	FavoriteSourceTypeImpact       FavoriteSourceType = "impact"
	FavoriteSourceTypeNotebook     FavoriteSourceType = "notebook"
	FavoriteSourceTypeRetention    FavoriteSourceType = "retention"
	FavoriteSourceTypeSegmentation FavoriteSourceType = "segmentation"
	FavoriteSourceTypeSessions     FavoriteSourceType = "sessions"
	FavoriteSourceTypeUserflows    FavoriteSourceType = "userflows"
)

// PossibleFavoriteSourceTypeValues returns the possible values for the FavoriteSourceType const type.
func PossibleFavoriteSourceTypeValues() []FavoriteSourceType {
	return []FavoriteSourceType{
		FavoriteSourceTypeEvents,
		FavoriteSourceTypeFunnel,
		FavoriteSourceTypeImpact,
		FavoriteSourceTypeNotebook,
		FavoriteSourceTypeRetention,
		FavoriteSourceTypeSegmentation,
		FavoriteSourceTypeSessions,
		FavoriteSourceTypeUserflows,
	}
}

// ToPtr returns a *FavoriteSourceType pointing to the current value.
func (c FavoriteSourceType) ToPtr() *FavoriteSourceType {
	return &c
}

// FavoriteType - Enum indicating if this favorite definition is owned by a specific user or is shared between all users with
// access to the Application Insights component.
type FavoriteType string

const (
	FavoriteTypeShared FavoriteType = "shared"
	FavoriteTypeUser   FavoriteType = "user"
)

// PossibleFavoriteTypeValues returns the possible values for the FavoriteType const type.
func PossibleFavoriteTypeValues() []FavoriteType {
	return []FavoriteType{
		FavoriteTypeShared,
		FavoriteTypeUser,
	}
}

// ToPtr returns a *FavoriteType pointing to the current value.
func (c FavoriteType) ToPtr() *FavoriteType {
	return &c
}

// FlowType - Used by the Application Insights system to determine what kind of flow this component was created by. This is
// to be set to 'Bluefield' when creating/updating a component via the REST API.
type FlowType string

const (
	FlowTypeBluefield FlowType = "Bluefield"
)

// PossibleFlowTypeValues returns the possible values for the FlowType const type.
func PossibleFlowTypeValues() []FlowType {
	return []FlowType{
		FlowTypeBluefield,
	}
}

// ToPtr returns a *FlowType pointing to the current value.
func (c FlowType) ToPtr() *FlowType {
	return &c
}

// IngestionMode - Indicates the flow of the ingestion.
type IngestionMode string

const (
	IngestionModeApplicationInsights                       IngestionMode = "ApplicationInsights"
	IngestionModeApplicationInsightsWithDiagnosticSettings IngestionMode = "ApplicationInsightsWithDiagnosticSettings"
	IngestionModeLogAnalytics                              IngestionMode = "LogAnalytics"
)

// PossibleIngestionModeValues returns the possible values for the IngestionMode const type.
func PossibleIngestionModeValues() []IngestionMode {
	return []IngestionMode{
		IngestionModeApplicationInsights,
		IngestionModeApplicationInsightsWithDiagnosticSettings,
		IngestionModeLogAnalytics,
	}
}

// ToPtr returns a *IngestionMode pointing to the current value.
func (c IngestionMode) ToPtr() *IngestionMode {
	return &c
}

// ItemScope - Enum indicating if this item definition is owned by a specific user or is shared between all users with access
// to the Application Insights component.
type ItemScope string

const (
	ItemScopeShared ItemScope = "shared"
	ItemScopeUser   ItemScope = "user"
)

// PossibleItemScopeValues returns the possible values for the ItemScope const type.
func PossibleItemScopeValues() []ItemScope {
	return []ItemScope{
		ItemScopeShared,
		ItemScopeUser,
	}
}

// ToPtr returns a *ItemScope pointing to the current value.
func (c ItemScope) ToPtr() *ItemScope {
	return &c
}

type ItemScopePath string

const (
	ItemScopePathAnalyticsItems   ItemScopePath = "analyticsItems"
	ItemScopePathMyanalyticsItems ItemScopePath = "myanalyticsItems"
)

// PossibleItemScopePathValues returns the possible values for the ItemScopePath const type.
func PossibleItemScopePathValues() []ItemScopePath {
	return []ItemScopePath{
		ItemScopePathAnalyticsItems,
		ItemScopePathMyanalyticsItems,
	}
}

// ToPtr returns a *ItemScopePath pointing to the current value.
func (c ItemScopePath) ToPtr() *ItemScopePath {
	return &c
}

// ItemType - Enum indicating the type of the Analytics item.
type ItemType string

const (
	ItemTypeFunction ItemType = "function"
	ItemTypeNone     ItemType = "none"
	ItemTypeQuery    ItemType = "query"
	ItemTypeRecent   ItemType = "recent"
)

// PossibleItemTypeValues returns the possible values for the ItemType const type.
func PossibleItemTypeValues() []ItemType {
	return []ItemType{
		ItemTypeFunction,
		ItemTypeNone,
		ItemTypeQuery,
		ItemTypeRecent,
	}
}

// ToPtr returns a *ItemType pointing to the current value.
func (c ItemType) ToPtr() *ItemType {
	return &c
}

type ItemTypeParameter string

const (
	ItemTypeParameterFolder   ItemTypeParameter = "folder"
	ItemTypeParameterFunction ItemTypeParameter = "function"
	ItemTypeParameterNone     ItemTypeParameter = "none"
	ItemTypeParameterQuery    ItemTypeParameter = "query"
	ItemTypeParameterRecent   ItemTypeParameter = "recent"
)

// PossibleItemTypeParameterValues returns the possible values for the ItemTypeParameter const type.
func PossibleItemTypeParameterValues() []ItemTypeParameter {
	return []ItemTypeParameter{
		ItemTypeParameterFolder,
		ItemTypeParameterFunction,
		ItemTypeParameterNone,
		ItemTypeParameterQuery,
		ItemTypeParameterRecent,
	}
}

// ToPtr returns a *ItemTypeParameter pointing to the current value.
func (c ItemTypeParameter) ToPtr() *ItemTypeParameter {
	return &c
}

// Kind - The kind of workbook. Choices are user and shared.
type Kind string

const (
	KindShared Kind = "shared"
	KindUser   Kind = "user"
)

// PossibleKindValues returns the possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{
		KindShared,
		KindUser,
	}
}

// ToPtr returns a *Kind pointing to the current value.
func (c Kind) ToPtr() *Kind {
	return &c
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *ManagedServiceIdentityType pointing to the current value.
func (c ManagedServiceIdentityType) ToPtr() *ManagedServiceIdentityType {
	return &c
}

// MyWorkbookManagedIdentityType - The identity type.
type MyWorkbookManagedIdentityType string

const (
	MyWorkbookManagedIdentityTypeNone         MyWorkbookManagedIdentityType = "None"
	MyWorkbookManagedIdentityTypeUserAssigned MyWorkbookManagedIdentityType = "UserAssigned"
)

// PossibleMyWorkbookManagedIdentityTypeValues returns the possible values for the MyWorkbookManagedIdentityType const type.
func PossibleMyWorkbookManagedIdentityTypeValues() []MyWorkbookManagedIdentityType {
	return []MyWorkbookManagedIdentityType{
		MyWorkbookManagedIdentityTypeNone,
		MyWorkbookManagedIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *MyWorkbookManagedIdentityType pointing to the current value.
func (c MyWorkbookManagedIdentityType) ToPtr() *MyWorkbookManagedIdentityType {
	return &c
}

// PublicNetworkAccessType - The network access type for operating on the Application Insights Component. By default it is
// Enabled
type PublicNetworkAccessType string

const (
	// PublicNetworkAccessTypeDisabled - Disables public connectivity to Application Insights through public DNS.
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	// PublicNetworkAccessTypeEnabled - Enables connectivity to Application Insights through public DNS.
	PublicNetworkAccessTypeEnabled PublicNetworkAccessType = "Enabled"
)

// PossiblePublicNetworkAccessTypeValues returns the possible values for the PublicNetworkAccessType const type.
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return []PublicNetworkAccessType{
		PublicNetworkAccessTypeDisabled,
		PublicNetworkAccessTypeEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccessType pointing to the current value.
func (c PublicNetworkAccessType) ToPtr() *PublicNetworkAccessType {
	return &c
}

// PurgeState - Status of the operation represented by the requested Id.
type PurgeState string

const (
	PurgeStateCompleted PurgeState = "completed"
	PurgeStatePending   PurgeState = "pending"
)

// PossiblePurgeStateValues returns the possible values for the PurgeState const type.
func PossiblePurgeStateValues() []PurgeState {
	return []PurgeState{
		PurgeStateCompleted,
		PurgeStatePending,
	}
}

// ToPtr returns a *PurgeState pointing to the current value.
func (c PurgeState) ToPtr() *PurgeState {
	return &c
}

// RequestSource - Describes what tool created this Application Insights component. Customers using this API should set this
// to the default 'rest'.
type RequestSource string

const (
	RequestSourceRest RequestSource = "rest"
)

// PossibleRequestSourceValues returns the possible values for the RequestSource const type.
func PossibleRequestSourceValues() []RequestSource {
	return []RequestSource{
		RequestSourceRest,
	}
}

// ToPtr returns a *RequestSource pointing to the current value.
func (c RequestSource) ToPtr() *RequestSource {
	return &c
}

// SharedTypeKind - The kind of workbook. Choices are user and shared.
type SharedTypeKind string

const (
	SharedTypeKindShared SharedTypeKind = "shared"
	SharedTypeKindUser   SharedTypeKind = "user"
)

// PossibleSharedTypeKindValues returns the possible values for the SharedTypeKind const type.
func PossibleSharedTypeKindValues() []SharedTypeKind {
	return []SharedTypeKind{
		SharedTypeKindShared,
		SharedTypeKindUser,
	}
}

// ToPtr returns a *SharedTypeKind pointing to the current value.
func (c SharedTypeKind) ToPtr() *SharedTypeKind {
	return &c
}

type StorageType string

const (
	StorageTypeServiceProfiler StorageType = "ServiceProfiler"
)

// PossibleStorageTypeValues returns the possible values for the StorageType const type.
func PossibleStorageTypeValues() []StorageType {
	return []StorageType{
		StorageTypeServiceProfiler,
	}
}

// ToPtr returns a *StorageType pointing to the current value.
func (c StorageType) ToPtr() *StorageType {
	return &c
}

// WebTestKind - The kind of web test that this web test watches. Choices are ping and multistep.
type WebTestKind string

const (
	WebTestKindPing      WebTestKind = "ping"
	WebTestKindMultistep WebTestKind = "multistep"
)

// PossibleWebTestKindValues returns the possible values for the WebTestKind const type.
func PossibleWebTestKindValues() []WebTestKind {
	return []WebTestKind{
		WebTestKindPing,
		WebTestKindMultistep,
	}
}

// ToPtr returns a *WebTestKind pointing to the current value.
func (c WebTestKind) ToPtr() *WebTestKind {
	return &c
}
