package contentmoderator

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

// StatusEnum enumerates the values for status enum.
type StatusEnum string

const (
	// Complete ...
	Complete StatusEnum = "Complete"
	// Pending ...
	Pending StatusEnum = "Pending"
	// Unpublished ...
	Unpublished StatusEnum = "Unpublished"
)

// PossibleStatusEnumValues returns an array of possible values for the StatusEnum const type.
func PossibleStatusEnumValues() []StatusEnum {
	return []StatusEnum{Complete, Pending, Unpublished}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeImage ...
	TypeImage Type = "Image"
	// TypeText ...
	TypeText Type = "Text"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeImage, TypeText}
}
