package advisor

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Category enumerates the values for category.
type Category string

const (
	// Cost ...
	Cost Category = "Cost"
	// HighAvailability ...
	HighAvailability Category = "HighAvailability"
	// OperationalExcellence ...
	OperationalExcellence Category = "OperationalExcellence"
	// Performance ...
	Performance Category = "Performance"
	// Security ...
	Security Category = "Security"
)

// PossibleCategoryValues returns an array of possible values for the Category const type.
func PossibleCategoryValues() []Category {
	return []Category{Cost, HighAvailability, OperationalExcellence, Performance, Security}
}

// Impact enumerates the values for impact.
type Impact string

const (
	// High ...
	High Impact = "High"
	// Low ...
	Low Impact = "Low"
	// Medium ...
	Medium Impact = "Medium"
)

// PossibleImpactValues returns an array of possible values for the Impact const type.
func PossibleImpactValues() []Impact {
	return []Impact{High, Low, Medium}
}

// Risk enumerates the values for risk.
type Risk string

const (
	// Error ...
	Error Risk = "Error"
	// None ...
	None Risk = "None"
	// Warning ...
	Warning Risk = "Warning"
)

// PossibleRiskValues returns an array of possible values for the Risk const type.
func PossibleRiskValues() []Risk {
	return []Risk{Error, None, Warning}
}

// Scenario enumerates the values for scenario.
type Scenario string

const (
	// Alerts ...
	Alerts Scenario = "Alerts"
)

// PossibleScenarioValues returns an array of possible values for the Scenario const type.
func PossibleScenarioValues() []Scenario {
	return []Scenario{Alerts}
}
