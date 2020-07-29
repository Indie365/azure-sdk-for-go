package support

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

// CommunicationDirection enumerates the values for communication direction.
type CommunicationDirection string

const (
	// Inbound ...
	Inbound CommunicationDirection = "inbound"
	// Outbound ...
	Outbound CommunicationDirection = "outbound"
)

// PossibleCommunicationDirectionValues returns an array of possible values for the CommunicationDirection const type.
func PossibleCommunicationDirectionValues() []CommunicationDirection {
	return []CommunicationDirection{Inbound, Outbound}
}

// CommunicationType enumerates the values for communication type.
type CommunicationType string

const (
	// Phone ...
	Phone CommunicationType = "phone"
	// Web ...
	Web CommunicationType = "web"
)

// PossibleCommunicationTypeValues returns an array of possible values for the CommunicationType const type.
func PossibleCommunicationTypeValues() []CommunicationType {
	return []CommunicationType{Phone, Web}
}

// PreferredContactMethod enumerates the values for preferred contact method.
type PreferredContactMethod string

const (
	// PreferredContactMethodEmail ...
	PreferredContactMethodEmail PreferredContactMethod = "email"
	// PreferredContactMethodPhone ...
	PreferredContactMethodPhone PreferredContactMethod = "phone"
)

// PossiblePreferredContactMethodValues returns an array of possible values for the PreferredContactMethod const type.
func PossiblePreferredContactMethodValues() []PreferredContactMethod {
	return []PreferredContactMethod{PreferredContactMethodEmail, PreferredContactMethodPhone}
}

// SeverityLevel enumerates the values for severity level.
type SeverityLevel string

const (
	// Critical ...
	Critical SeverityLevel = "critical"
	// Highestcriticalimpact ...
	Highestcriticalimpact SeverityLevel = "highestcriticalimpact"
	// Minimal ...
	Minimal SeverityLevel = "minimal"
	// Moderate ...
	Moderate SeverityLevel = "moderate"
)

// PossibleSeverityLevelValues returns an array of possible values for the SeverityLevel const type.
func PossibleSeverityLevelValues() []SeverityLevel {
	return []SeverityLevel{Critical, Highestcriticalimpact, Minimal, Moderate}
}

// Status enumerates the values for status.
type Status string

const (
	// Closed ...
	Closed Status = "closed"
	// Open ...
	Open Status = "open"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Closed, Open}
}

// Type enumerates the values for type.
type Type string

const (
	// MicrosoftSupportcommunications ...
	MicrosoftSupportcommunications Type = "Microsoft.Support/communications"
	// MicrosoftSupportsupportTickets ...
	MicrosoftSupportsupportTickets Type = "Microsoft.Support/supportTickets"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{MicrosoftSupportcommunications, MicrosoftSupportsupportTickets}
}
