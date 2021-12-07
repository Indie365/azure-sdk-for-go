//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad

const (
	module  = "armaad"
	version = "v0.1.1"
)

// Category - Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of Diagnostic Log categories for a resource,
// first perform a GET diagnostic settings operation.
type Category string

const (
	CategoryAuditLogs  Category = "AuditLogs"
	CategorySignInLogs Category = "SignInLogs"
)

// PossibleCategoryValues returns the possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{
		CategoryAuditLogs,
		CategorySignInLogs,
	}
}

// ToPtr returns a *Category pointing to the current value.
func (c Category) ToPtr() *Category {
	return &c
}

// CategoryType - The type of the diagnostic settings category.
type CategoryType string

const (
	CategoryTypeLogs CategoryType = "Logs"
)

// PossibleCategoryTypeValues returns the possible values for the CategoryType const type.
func PossibleCategoryTypeValues() []CategoryType {
	return []CategoryType{
		CategoryTypeLogs,
	}
}

// ToPtr returns a *CategoryType pointing to the current value.
func (c CategoryType) ToPtr() *CategoryType {
	return &c
}
