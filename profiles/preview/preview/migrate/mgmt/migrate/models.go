// +build go1.9

// Copyright 2018 Microsoft Corporation
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

package migrate

import original "github.com/Azure/azure-sdk-for-go/services/preview/migrate/mgmt/2018-09-01-preview/migrate"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ErrorsClient = original.ErrorsClient
type MachinesClient = original.MachinesClient
type ContainerElementKind = original.ContainerElementKind

const (
	ActionImport   ContainerElementKind = original.ActionImport
	EntitySet      ContainerElementKind = original.EntitySet
	FunctionImport ContainerElementKind = original.FunctionImport
	None           ContainerElementKind = original.None
	Singleton      ContainerElementKind = original.Singleton
)

type Direction = original.Direction

const (
	Ascending  Direction = original.Ascending
	Descending Direction = original.Descending
)

type Direction1 = original.Direction1

const (
	Direction1Ascending  Direction1 = original.Direction1Ascending
	Direction1Descending Direction1 = original.Direction1Descending
)

type ExpressionKind = original.ExpressionKind

const (
	ExpressionKindAnnotationPath             ExpressionKind = original.ExpressionKindAnnotationPath
	ExpressionKindBinaryConstant             ExpressionKind = original.ExpressionKindBinaryConstant
	ExpressionKindBooleanConstant            ExpressionKind = original.ExpressionKindBooleanConstant
	ExpressionKindCast                       ExpressionKind = original.ExpressionKindCast
	ExpressionKindCollection                 ExpressionKind = original.ExpressionKindCollection
	ExpressionKindDateConstant               ExpressionKind = original.ExpressionKindDateConstant
	ExpressionKindDateTimeOffsetConstant     ExpressionKind = original.ExpressionKindDateTimeOffsetConstant
	ExpressionKindDecimalConstant            ExpressionKind = original.ExpressionKindDecimalConstant
	ExpressionKindDurationConstant           ExpressionKind = original.ExpressionKindDurationConstant
	ExpressionKindEnumMember                 ExpressionKind = original.ExpressionKindEnumMember
	ExpressionKindFloatingConstant           ExpressionKind = original.ExpressionKindFloatingConstant
	ExpressionKindFunctionApplication        ExpressionKind = original.ExpressionKindFunctionApplication
	ExpressionKindGUIDConstant               ExpressionKind = original.ExpressionKindGUIDConstant
	ExpressionKindIf                         ExpressionKind = original.ExpressionKindIf
	ExpressionKindIntegerConstant            ExpressionKind = original.ExpressionKindIntegerConstant
	ExpressionKindIsType                     ExpressionKind = original.ExpressionKindIsType
	ExpressionKindLabeled                    ExpressionKind = original.ExpressionKindLabeled
	ExpressionKindLabeledExpressionReference ExpressionKind = original.ExpressionKindLabeledExpressionReference
	ExpressionKindNavigationPropertyPath     ExpressionKind = original.ExpressionKindNavigationPropertyPath
	ExpressionKindNone                       ExpressionKind = original.ExpressionKindNone
	ExpressionKindNull                       ExpressionKind = original.ExpressionKindNull
	ExpressionKindPath                       ExpressionKind = original.ExpressionKindPath
	ExpressionKindPropertyPath               ExpressionKind = original.ExpressionKindPropertyPath
	ExpressionKindRecord                     ExpressionKind = original.ExpressionKindRecord
	ExpressionKindStringConstant             ExpressionKind = original.ExpressionKindStringConstant
	ExpressionKindTimeOfDayConstant          ExpressionKind = original.ExpressionKindTimeOfDayConstant
)

type ExpressionKind1 = original.ExpressionKind1

const (
	ExpressionKind1AnnotationPath             ExpressionKind1 = original.ExpressionKind1AnnotationPath
	ExpressionKind1BinaryConstant             ExpressionKind1 = original.ExpressionKind1BinaryConstant
	ExpressionKind1BooleanConstant            ExpressionKind1 = original.ExpressionKind1BooleanConstant
	ExpressionKind1Cast                       ExpressionKind1 = original.ExpressionKind1Cast
	ExpressionKind1Collection                 ExpressionKind1 = original.ExpressionKind1Collection
	ExpressionKind1DateConstant               ExpressionKind1 = original.ExpressionKind1DateConstant
	ExpressionKind1DateTimeOffsetConstant     ExpressionKind1 = original.ExpressionKind1DateTimeOffsetConstant
	ExpressionKind1DecimalConstant            ExpressionKind1 = original.ExpressionKind1DecimalConstant
	ExpressionKind1DurationConstant           ExpressionKind1 = original.ExpressionKind1DurationConstant
	ExpressionKind1EnumMember                 ExpressionKind1 = original.ExpressionKind1EnumMember
	ExpressionKind1FloatingConstant           ExpressionKind1 = original.ExpressionKind1FloatingConstant
	ExpressionKind1FunctionApplication        ExpressionKind1 = original.ExpressionKind1FunctionApplication
	ExpressionKind1GUIDConstant               ExpressionKind1 = original.ExpressionKind1GUIDConstant
	ExpressionKind1If                         ExpressionKind1 = original.ExpressionKind1If
	ExpressionKind1IntegerConstant            ExpressionKind1 = original.ExpressionKind1IntegerConstant
	ExpressionKind1IsType                     ExpressionKind1 = original.ExpressionKind1IsType
	ExpressionKind1Labeled                    ExpressionKind1 = original.ExpressionKind1Labeled
	ExpressionKind1LabeledExpressionReference ExpressionKind1 = original.ExpressionKind1LabeledExpressionReference
	ExpressionKind1NavigationPropertyPath     ExpressionKind1 = original.ExpressionKind1NavigationPropertyPath
	ExpressionKind1None                       ExpressionKind1 = original.ExpressionKind1None
	ExpressionKind1Null                       ExpressionKind1 = original.ExpressionKind1Null
	ExpressionKind1Path                       ExpressionKind1 = original.ExpressionKind1Path
	ExpressionKind1PropertyPath               ExpressionKind1 = original.ExpressionKind1PropertyPath
	ExpressionKind1Record                     ExpressionKind1 = original.ExpressionKind1Record
	ExpressionKind1StringConstant             ExpressionKind1 = original.ExpressionKind1StringConstant
	ExpressionKind1TimeOfDayConstant          ExpressionKind1 = original.ExpressionKind1TimeOfDayConstant
)

type Kind = original.Kind

const (
	Aggregate Kind = original.Aggregate
	Compute   Kind = original.Compute
	Filter    Kind = original.Filter
	GroupBy   Kind = original.GroupBy
)

type Kind1 = original.Kind1

const (
	Kind1AggregatedCollectionPropertyNode  Kind1 = original.Kind1AggregatedCollectionPropertyNode
	Kind1All                               Kind1 = original.Kind1All
	Kind1Any                               Kind1 = original.Kind1Any
	Kind1BinaryOperator                    Kind1 = original.Kind1BinaryOperator
	Kind1CollectionComplexNode             Kind1 = original.Kind1CollectionComplexNode
	Kind1CollectionConstant                Kind1 = original.Kind1CollectionConstant
	Kind1CollectionFunctionCall            Kind1 = original.Kind1CollectionFunctionCall
	Kind1CollectionNavigationNode          Kind1 = original.Kind1CollectionNavigationNode
	Kind1CollectionOpenPropertyAccess      Kind1 = original.Kind1CollectionOpenPropertyAccess
	Kind1CollectionPropertyAccess          Kind1 = original.Kind1CollectionPropertyAccess
	Kind1CollectionPropertyNode            Kind1 = original.Kind1CollectionPropertyNode
	Kind1CollectionResourceCast            Kind1 = original.Kind1CollectionResourceCast
	Kind1CollectionResourceFunctionCall    Kind1 = original.Kind1CollectionResourceFunctionCall
	Kind1Constant                          Kind1 = original.Kind1Constant
	Kind1Convert                           Kind1 = original.Kind1Convert
	Kind1Count                             Kind1 = original.Kind1Count
	Kind1EntitySet                         Kind1 = original.Kind1EntitySet
	Kind1In                                Kind1 = original.Kind1In
	Kind1KeyLookup                         Kind1 = original.Kind1KeyLookup
	Kind1NamedFunctionParameter            Kind1 = original.Kind1NamedFunctionParameter
	Kind1None                              Kind1 = original.Kind1None
	Kind1NonResourceRangeVariableReference Kind1 = original.Kind1NonResourceRangeVariableReference
	Kind1ParameterAlias                    Kind1 = original.Kind1ParameterAlias
	Kind1ResourceRangeVariableReference    Kind1 = original.Kind1ResourceRangeVariableReference
	Kind1SearchTerm                        Kind1 = original.Kind1SearchTerm
	Kind1SingleComplexNode                 Kind1 = original.Kind1SingleComplexNode
	Kind1SingleNavigationNode              Kind1 = original.Kind1SingleNavigationNode
	Kind1SingleResourceCast                Kind1 = original.Kind1SingleResourceCast
	Kind1SingleResourceFunctionCall        Kind1 = original.Kind1SingleResourceFunctionCall
	Kind1SingleValueCast                   Kind1 = original.Kind1SingleValueCast
	Kind1SingleValueFunctionCall           Kind1 = original.Kind1SingleValueFunctionCall
	Kind1SingleValueOpenPropertyAccess     Kind1 = original.Kind1SingleValueOpenPropertyAccess
	Kind1SingleValuePropertyAccess         Kind1 = original.Kind1SingleValuePropertyAccess
	Kind1UnaryOperator                     Kind1 = original.Kind1UnaryOperator
)

type OnDelete = original.OnDelete

const (
	OnDeleteCascade OnDelete = original.OnDeleteCascade
	OnDeleteNone    OnDelete = original.OnDeleteNone
)

type PropertyKind = original.PropertyKind

const (
	PropertyKindNavigation PropertyKind = original.PropertyKindNavigation
	PropertyKindNone       PropertyKind = original.PropertyKindNone
	PropertyKindStructural PropertyKind = original.PropertyKindStructural
)

type PropertyKind1 = original.PropertyKind1

const (
	PropertyKind1Navigation PropertyKind1 = original.PropertyKind1Navigation
	PropertyKind1None       PropertyKind1 = original.PropertyKind1None
	PropertyKind1Structural PropertyKind1 = original.PropertyKind1Structural
)

type PropertyKind2 = original.PropertyKind2

const (
	PropertyKind2Navigation PropertyKind2 = original.PropertyKind2Navigation
	PropertyKind2None       PropertyKind2 = original.PropertyKind2None
	PropertyKind2Structural PropertyKind2 = original.PropertyKind2Structural
)

type SchemaElementKind = original.SchemaElementKind

const (
	SchemaElementKindAction          SchemaElementKind = original.SchemaElementKindAction
	SchemaElementKindEntityContainer SchemaElementKind = original.SchemaElementKindEntityContainer
	SchemaElementKindFunction        SchemaElementKind = original.SchemaElementKindFunction
	SchemaElementKindNone            SchemaElementKind = original.SchemaElementKindNone
	SchemaElementKindTerm            SchemaElementKind = original.SchemaElementKindTerm
	SchemaElementKindTypeDefinition  SchemaElementKind = original.SchemaElementKindTypeDefinition
)

type SchemaElementKind1 = original.SchemaElementKind1

const (
	SchemaElementKind1Action          SchemaElementKind1 = original.SchemaElementKind1Action
	SchemaElementKind1EntityContainer SchemaElementKind1 = original.SchemaElementKind1EntityContainer
	SchemaElementKind1Function        SchemaElementKind1 = original.SchemaElementKind1Function
	SchemaElementKind1None            SchemaElementKind1 = original.SchemaElementKind1None
	SchemaElementKind1Term            SchemaElementKind1 = original.SchemaElementKind1Term
	SchemaElementKind1TypeDefinition  SchemaElementKind1 = original.SchemaElementKind1TypeDefinition
)

type SchemaElementKind2 = original.SchemaElementKind2

const (
	SchemaElementKind2Action          SchemaElementKind2 = original.SchemaElementKind2Action
	SchemaElementKind2EntityContainer SchemaElementKind2 = original.SchemaElementKind2EntityContainer
	SchemaElementKind2Function        SchemaElementKind2 = original.SchemaElementKind2Function
	SchemaElementKind2None            SchemaElementKind2 = original.SchemaElementKind2None
	SchemaElementKind2Term            SchemaElementKind2 = original.SchemaElementKind2Term
	SchemaElementKind2TypeDefinition  SchemaElementKind2 = original.SchemaElementKind2TypeDefinition
)

type TypeKind = original.TypeKind

const (
	TypeKindCollection      TypeKind = original.TypeKindCollection
	TypeKindComplex         TypeKind = original.TypeKindComplex
	TypeKindEntity          TypeKind = original.TypeKindEntity
	TypeKindEntityReference TypeKind = original.TypeKindEntityReference
	TypeKindEnum            TypeKind = original.TypeKindEnum
	TypeKindNone            TypeKind = original.TypeKindNone
	TypeKindPath            TypeKind = original.TypeKindPath
	TypeKindPrimitive       TypeKind = original.TypeKindPrimitive
	TypeKindTypeDefinition  TypeKind = original.TypeKindTypeDefinition
	TypeKindUntyped         TypeKind = original.TypeKindUntyped
)

type TypeKind1 = original.TypeKind1

const (
	TypeKind1Collection      TypeKind1 = original.TypeKind1Collection
	TypeKind1Complex         TypeKind1 = original.TypeKind1Complex
	TypeKind1Entity          TypeKind1 = original.TypeKind1Entity
	TypeKind1EntityReference TypeKind1 = original.TypeKind1EntityReference
	TypeKind1Enum            TypeKind1 = original.TypeKind1Enum
	TypeKind1None            TypeKind1 = original.TypeKind1None
	TypeKind1Path            TypeKind1 = original.TypeKind1Path
	TypeKind1Primitive       TypeKind1 = original.TypeKind1Primitive
	TypeKind1TypeDefinition  TypeKind1 = original.TypeKind1TypeDefinition
	TypeKind1Untyped         TypeKind1 = original.TypeKind1Untyped
)

type ApplyClause = original.ApplyClause
type ApplyQueryOption = original.ApplyQueryOption
type AssessmentDetails = original.AssessmentDetails
type CountQueryOption = original.CountQueryOption
type DefaultQuerySettings = original.DefaultQuerySettings
type DiscoveryDetails = original.DiscoveryDetails
type EdmReferentialConstraintPropertyPair = original.EdmReferentialConstraintPropertyPair
type Error = original.Error
type ErrorCollection = original.ErrorCollection
type ErrorProperties = original.ErrorProperties
type FilterClause = original.FilterClause
type FilterQueryOption = original.FilterQueryOption
type GoalSummary = original.GoalSummary
type IEdmEntityContainer = original.IEdmEntityContainer
type IEdmEntityContainerElement = original.IEdmEntityContainerElement
type IEdmExpression = original.IEdmExpression
type IEdmModel = original.IEdmModel
type IEdmNavigationProperty = original.IEdmNavigationProperty
type IEdmNavigationPropertyBinding = original.IEdmNavigationPropertyBinding
type IEdmNavigationSource = original.IEdmNavigationSource
type IEdmPathExpression = original.IEdmPathExpression
type IEdmProperty = original.IEdmProperty
type IEdmReferentialConstraint = original.IEdmReferentialConstraint
type IEdmSchemaElement = original.IEdmSchemaElement
type IEdmStructuralProperty = original.IEdmStructuralProperty
type IEdmStructuredType = original.IEdmStructuredType
type IEdmTerm = original.IEdmTerm
type IEdmType = original.IEdmType
type IEdmTypeReference = original.IEdmTypeReference
type IEdmVocabularyAnnotation = original.IEdmVocabularyAnnotation
type Machine = original.Machine
type MachineCollection = original.MachineCollection
type MachineProperties = original.MachineProperties
type MigrationDetails = original.MigrationDetails
type ODataPath = original.ODataPath
type ODataPathSegment = original.ODataPathSegment
type ODataQueryContext = original.ODataQueryContext
type ODataQueryOptions1 = original.ODataQueryOptions1
type ODataRawQueryOptions = original.ODataRawQueryOptions
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationResultList = original.OperationResultList
type OrderByClause = original.OrderByClause
type OrderByNode = original.OrderByNode
type OrderByQueryOption = original.OrderByQueryOption
type Project = original.Project
type ProjectProperties = original.ProjectProperties
type RangeVariable = original.RangeVariable
type RefreshSummaryResult = original.RefreshSummaryResult
type RegisterToolInput = original.RegisterToolInput
type RegistrationResult = original.RegistrationResult
type SelectExpandClause = original.SelectExpandClause
type SelectExpandQueryOption = original.SelectExpandQueryOption
type SingleValueNode = original.SingleValueNode
type SkipQueryOption = original.SkipQueryOption
type Solution = original.Solution
type SolutionConfig = original.SolutionConfig
type SolutionDetails = original.SolutionDetails
type SolutionProperties = original.SolutionProperties
type SolutionsCollection = original.SolutionsCollection
type SolutionSummary = original.SolutionSummary
type TopQueryOption = original.TopQueryOption
type TransformationNode = original.TransformationNode
type OperationsClient = original.OperationsClient
type ProjectsClient = original.ProjectsClient
type SolutionsClient = original.SolutionsClient

func New(subscriptionID string, acceptLanguage string) BaseClient {
	return original.New(subscriptionID, acceptLanguage)
}
func NewWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewErrorsClient(subscriptionID string, acceptLanguage string) ErrorsClient {
	return original.NewErrorsClient(subscriptionID, acceptLanguage)
}
func NewErrorsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) ErrorsClient {
	return original.NewErrorsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewMachinesClient(subscriptionID string, acceptLanguage string) MachinesClient {
	return original.NewMachinesClient(subscriptionID, acceptLanguage)
}
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) MachinesClient {
	return original.NewMachinesClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func PossibleContainerElementKindValues() []ContainerElementKind {
	return original.PossibleContainerElementKindValues()
}
func PossibleDirectionValues() []Direction {
	return original.PossibleDirectionValues()
}
func PossibleDirection1Values() []Direction1 {
	return original.PossibleDirection1Values()
}
func PossibleExpressionKindValues() []ExpressionKind {
	return original.PossibleExpressionKindValues()
}
func PossibleExpressionKind1Values() []ExpressionKind1 {
	return original.PossibleExpressionKind1Values()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleKind1Values() []Kind1 {
	return original.PossibleKind1Values()
}
func PossibleOnDeleteValues() []OnDelete {
	return original.PossibleOnDeleteValues()
}
func PossiblePropertyKindValues() []PropertyKind {
	return original.PossiblePropertyKindValues()
}
func PossiblePropertyKind1Values() []PropertyKind1 {
	return original.PossiblePropertyKind1Values()
}
func PossiblePropertyKind2Values() []PropertyKind2 {
	return original.PossiblePropertyKind2Values()
}
func PossibleSchemaElementKindValues() []SchemaElementKind {
	return original.PossibleSchemaElementKindValues()
}
func PossibleSchemaElementKind1Values() []SchemaElementKind1 {
	return original.PossibleSchemaElementKind1Values()
}
func PossibleSchemaElementKind2Values() []SchemaElementKind2 {
	return original.PossibleSchemaElementKind2Values()
}
func PossibleTypeKindValues() []TypeKind {
	return original.PossibleTypeKindValues()
}
func PossibleTypeKind1Values() []TypeKind1 {
	return original.PossibleTypeKind1Values()
}
func NewOperationsClient(subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, acceptLanguage)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewProjectsClient(subscriptionID string, acceptLanguage string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID, acceptLanguage)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewSolutionsClient(subscriptionID string, acceptLanguage string) SolutionsClient {
	return original.NewSolutionsClient(subscriptionID, acceptLanguage)
}
func NewSolutionsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) SolutionsClient {
	return original.NewSolutionsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
