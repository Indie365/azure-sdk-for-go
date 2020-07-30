package compute

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

// AccessLevel enumerates the values for access level.
type AccessLevel string

const (
	// None ...
	None AccessLevel = "None"
	// Read ...
	Read AccessLevel = "Read"
)

// PossibleAccessLevelValues returns an array of possible values for the AccessLevel const type.
func PossibleAccessLevelValues() []AccessLevel {
	return []AccessLevel{None, Read}
}

// CachingTypes enumerates the values for caching types.
type CachingTypes string

const (
	// CachingTypesNone ...
	CachingTypesNone CachingTypes = "None"
	// CachingTypesReadOnly ...
	CachingTypesReadOnly CachingTypes = "ReadOnly"
	// CachingTypesReadWrite ...
	CachingTypesReadWrite CachingTypes = "ReadWrite"
)

// PossibleCachingTypesValues returns an array of possible values for the CachingTypes const type.
func PossibleCachingTypesValues() []CachingTypes {
	return []CachingTypes{CachingTypesNone, CachingTypesReadOnly, CachingTypesReadWrite}
}

// ComponentNames enumerates the values for component names.
type ComponentNames string

const (
	// MicrosoftWindowsShellSetup ...
	MicrosoftWindowsShellSetup ComponentNames = "Microsoft-Windows-Shell-Setup"
)

// PossibleComponentNamesValues returns an array of possible values for the ComponentNames const type.
func PossibleComponentNamesValues() []ComponentNames {
	return []ComponentNames{MicrosoftWindowsShellSetup}
}

// DiskCreateOption enumerates the values for disk create option.
type DiskCreateOption string

const (
	// Attach ...
	Attach DiskCreateOption = "Attach"
	// Copy ...
	Copy DiskCreateOption = "Copy"
	// Empty ...
	Empty DiskCreateOption = "Empty"
	// FromImage ...
	FromImage DiskCreateOption = "FromImage"
	// Import ...
	Import DiskCreateOption = "Import"
	// Restore ...
	Restore DiskCreateOption = "Restore"
)

// PossibleDiskCreateOptionValues returns an array of possible values for the DiskCreateOption const type.
func PossibleDiskCreateOptionValues() []DiskCreateOption {
	return []DiskCreateOption{Attach, Copy, Empty, FromImage, Import, Restore}
}

// DiskCreateOptionTypes enumerates the values for disk create option types.
type DiskCreateOptionTypes string

const (
	// DiskCreateOptionTypesAttach ...
	DiskCreateOptionTypesAttach DiskCreateOptionTypes = "Attach"
	// DiskCreateOptionTypesEmpty ...
	DiskCreateOptionTypesEmpty DiskCreateOptionTypes = "Empty"
	// DiskCreateOptionTypesFromImage ...
	DiskCreateOptionTypesFromImage DiskCreateOptionTypes = "FromImage"
)

// PossibleDiskCreateOptionTypesValues returns an array of possible values for the DiskCreateOptionTypes const type.
func PossibleDiskCreateOptionTypesValues() []DiskCreateOptionTypes {
	return []DiskCreateOptionTypes{DiskCreateOptionTypesAttach, DiskCreateOptionTypesEmpty, DiskCreateOptionTypesFromImage}
}

// InstanceViewTypes enumerates the values for instance view types.
type InstanceViewTypes string

const (
	// InstanceView ...
	InstanceView InstanceViewTypes = "instanceView"
)

// PossibleInstanceViewTypesValues returns an array of possible values for the InstanceViewTypes const type.
func PossibleInstanceViewTypesValues() []InstanceViewTypes {
	return []InstanceViewTypes{InstanceView}
}

// OperatingSystemStateTypes enumerates the values for operating system state types.
type OperatingSystemStateTypes string

const (
	// Generalized ...
	Generalized OperatingSystemStateTypes = "Generalized"
	// Specialized ...
	Specialized OperatingSystemStateTypes = "Specialized"
)

// PossibleOperatingSystemStateTypesValues returns an array of possible values for the OperatingSystemStateTypes const type.
func PossibleOperatingSystemStateTypesValues() []OperatingSystemStateTypes {
	return []OperatingSystemStateTypes{Generalized, Specialized}
}

// OperatingSystemTypes enumerates the values for operating system types.
type OperatingSystemTypes string

const (
	// Linux ...
	Linux OperatingSystemTypes = "Linux"
	// Windows ...
	Windows OperatingSystemTypes = "Windows"
)

// PossibleOperatingSystemTypesValues returns an array of possible values for the OperatingSystemTypes const type.
func PossibleOperatingSystemTypesValues() []OperatingSystemTypes {
	return []OperatingSystemTypes{Linux, Windows}
}

// PassNames enumerates the values for pass names.
type PassNames string

const (
	// OobeSystem ...
	OobeSystem PassNames = "OobeSystem"
)

// PossiblePassNamesValues returns an array of possible values for the PassNames const type.
func PossiblePassNamesValues() []PassNames {
	return []PassNames{OobeSystem}
}

// ProtocolTypes enumerates the values for protocol types.
type ProtocolTypes string

const (
	// HTTP ...
	HTTP ProtocolTypes = "Http"
	// HTTPS ...
	HTTPS ProtocolTypes = "Https"
)

// PossibleProtocolTypesValues returns an array of possible values for the ProtocolTypes const type.
func PossibleProtocolTypesValues() []ProtocolTypes {
	return []ProtocolTypes{HTTP, HTTPS}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{SystemAssigned}
}

// SettingNames enumerates the values for setting names.
type SettingNames string

const (
	// AutoLogon ...
	AutoLogon SettingNames = "AutoLogon"
	// FirstLogonCommands ...
	FirstLogonCommands SettingNames = "FirstLogonCommands"
)

// PossibleSettingNamesValues returns an array of possible values for the SettingNames const type.
func PossibleSettingNamesValues() []SettingNames {
	return []SettingNames{AutoLogon, FirstLogonCommands}
}

// StatusLevelTypes enumerates the values for status level types.
type StatusLevelTypes string

const (
	// Error ...
	Error StatusLevelTypes = "Error"
	// Info ...
	Info StatusLevelTypes = "Info"
	// Warning ...
	Warning StatusLevelTypes = "Warning"
)

// PossibleStatusLevelTypesValues returns an array of possible values for the StatusLevelTypes const type.
func PossibleStatusLevelTypesValues() []StatusLevelTypes {
	return []StatusLevelTypes{Error, Info, Warning}
}

// StorageAccountTypes enumerates the values for storage account types.
type StorageAccountTypes string

const (
	// PremiumLRS ...
	PremiumLRS StorageAccountTypes = "Premium_LRS"
	// StandardLRS ...
	StandardLRS StorageAccountTypes = "Standard_LRS"
)

// PossibleStorageAccountTypesValues returns an array of possible values for the StorageAccountTypes const type.
func PossibleStorageAccountTypesValues() []StorageAccountTypes {
	return []StorageAccountTypes{PremiumLRS, StandardLRS}
}

// UpgradeMode enumerates the values for upgrade mode.
type UpgradeMode string

const (
	// Automatic ...
	Automatic UpgradeMode = "Automatic"
	// Manual ...
	Manual UpgradeMode = "Manual"
)

// PossibleUpgradeModeValues returns an array of possible values for the UpgradeMode const type.
func PossibleUpgradeModeValues() []UpgradeMode {
	return []UpgradeMode{Automatic, Manual}
}

// VirtualMachineScaleSetSkuScaleType enumerates the values for virtual machine scale set sku scale type.
type VirtualMachineScaleSetSkuScaleType string

const (
	// VirtualMachineScaleSetSkuScaleTypeAutomatic ...
	VirtualMachineScaleSetSkuScaleTypeAutomatic VirtualMachineScaleSetSkuScaleType = "Automatic"
	// VirtualMachineScaleSetSkuScaleTypeNone ...
	VirtualMachineScaleSetSkuScaleTypeNone VirtualMachineScaleSetSkuScaleType = "None"
)

// PossibleVirtualMachineScaleSetSkuScaleTypeValues returns an array of possible values for the VirtualMachineScaleSetSkuScaleType const type.
func PossibleVirtualMachineScaleSetSkuScaleTypeValues() []VirtualMachineScaleSetSkuScaleType {
	return []VirtualMachineScaleSetSkuScaleType{VirtualMachineScaleSetSkuScaleTypeAutomatic, VirtualMachineScaleSetSkuScaleTypeNone}
}

// VirtualMachineSizeTypes enumerates the values for virtual machine size types.
type VirtualMachineSizeTypes string

const (
	// BasicA0 ...
	BasicA0 VirtualMachineSizeTypes = "Basic_A0"
	// BasicA1 ...
	BasicA1 VirtualMachineSizeTypes = "Basic_A1"
	// BasicA2 ...
	BasicA2 VirtualMachineSizeTypes = "Basic_A2"
	// BasicA3 ...
	BasicA3 VirtualMachineSizeTypes = "Basic_A3"
	// BasicA4 ...
	BasicA4 VirtualMachineSizeTypes = "Basic_A4"
	// StandardA0 ...
	StandardA0 VirtualMachineSizeTypes = "Standard_A0"
	// StandardA1 ...
	StandardA1 VirtualMachineSizeTypes = "Standard_A1"
	// StandardA10 ...
	StandardA10 VirtualMachineSizeTypes = "Standard_A10"
	// StandardA11 ...
	StandardA11 VirtualMachineSizeTypes = "Standard_A11"
	// StandardA2 ...
	StandardA2 VirtualMachineSizeTypes = "Standard_A2"
	// StandardA3 ...
	StandardA3 VirtualMachineSizeTypes = "Standard_A3"
	// StandardA4 ...
	StandardA4 VirtualMachineSizeTypes = "Standard_A4"
	// StandardA5 ...
	StandardA5 VirtualMachineSizeTypes = "Standard_A5"
	// StandardA6 ...
	StandardA6 VirtualMachineSizeTypes = "Standard_A6"
	// StandardA7 ...
	StandardA7 VirtualMachineSizeTypes = "Standard_A7"
	// StandardA8 ...
	StandardA8 VirtualMachineSizeTypes = "Standard_A8"
	// StandardA9 ...
	StandardA9 VirtualMachineSizeTypes = "Standard_A9"
	// StandardD1 ...
	StandardD1 VirtualMachineSizeTypes = "Standard_D1"
	// StandardD11 ...
	StandardD11 VirtualMachineSizeTypes = "Standard_D11"
	// StandardD11V2 ...
	StandardD11V2 VirtualMachineSizeTypes = "Standard_D11_v2"
	// StandardD12 ...
	StandardD12 VirtualMachineSizeTypes = "Standard_D12"
	// StandardD12V2 ...
	StandardD12V2 VirtualMachineSizeTypes = "Standard_D12_v2"
	// StandardD13 ...
	StandardD13 VirtualMachineSizeTypes = "Standard_D13"
	// StandardD13V2 ...
	StandardD13V2 VirtualMachineSizeTypes = "Standard_D13_v2"
	// StandardD14 ...
	StandardD14 VirtualMachineSizeTypes = "Standard_D14"
	// StandardD14V2 ...
	StandardD14V2 VirtualMachineSizeTypes = "Standard_D14_v2"
	// StandardD15V2 ...
	StandardD15V2 VirtualMachineSizeTypes = "Standard_D15_v2"
	// StandardD1V2 ...
	StandardD1V2 VirtualMachineSizeTypes = "Standard_D1_v2"
	// StandardD2 ...
	StandardD2 VirtualMachineSizeTypes = "Standard_D2"
	// StandardD2V2 ...
	StandardD2V2 VirtualMachineSizeTypes = "Standard_D2_v2"
	// StandardD3 ...
	StandardD3 VirtualMachineSizeTypes = "Standard_D3"
	// StandardD3V2 ...
	StandardD3V2 VirtualMachineSizeTypes = "Standard_D3_v2"
	// StandardD4 ...
	StandardD4 VirtualMachineSizeTypes = "Standard_D4"
	// StandardD4V2 ...
	StandardD4V2 VirtualMachineSizeTypes = "Standard_D4_v2"
	// StandardD5V2 ...
	StandardD5V2 VirtualMachineSizeTypes = "Standard_D5_v2"
	// StandardDS1 ...
	StandardDS1 VirtualMachineSizeTypes = "Standard_DS1"
	// StandardDS11 ...
	StandardDS11 VirtualMachineSizeTypes = "Standard_DS11"
	// StandardDS11V2 ...
	StandardDS11V2 VirtualMachineSizeTypes = "Standard_DS11_v2"
	// StandardDS12 ...
	StandardDS12 VirtualMachineSizeTypes = "Standard_DS12"
	// StandardDS12V2 ...
	StandardDS12V2 VirtualMachineSizeTypes = "Standard_DS12_v2"
	// StandardDS13 ...
	StandardDS13 VirtualMachineSizeTypes = "Standard_DS13"
	// StandardDS13V2 ...
	StandardDS13V2 VirtualMachineSizeTypes = "Standard_DS13_v2"
	// StandardDS14 ...
	StandardDS14 VirtualMachineSizeTypes = "Standard_DS14"
	// StandardDS14V2 ...
	StandardDS14V2 VirtualMachineSizeTypes = "Standard_DS14_v2"
	// StandardDS15V2 ...
	StandardDS15V2 VirtualMachineSizeTypes = "Standard_DS15_v2"
	// StandardDS1V2 ...
	StandardDS1V2 VirtualMachineSizeTypes = "Standard_DS1_v2"
	// StandardDS2 ...
	StandardDS2 VirtualMachineSizeTypes = "Standard_DS2"
	// StandardDS2V2 ...
	StandardDS2V2 VirtualMachineSizeTypes = "Standard_DS2_v2"
	// StandardDS3 ...
	StandardDS3 VirtualMachineSizeTypes = "Standard_DS3"
	// StandardDS3V2 ...
	StandardDS3V2 VirtualMachineSizeTypes = "Standard_DS3_v2"
	// StandardDS4 ...
	StandardDS4 VirtualMachineSizeTypes = "Standard_DS4"
	// StandardDS4V2 ...
	StandardDS4V2 VirtualMachineSizeTypes = "Standard_DS4_v2"
	// StandardDS5V2 ...
	StandardDS5V2 VirtualMachineSizeTypes = "Standard_DS5_v2"
	// StandardG1 ...
	StandardG1 VirtualMachineSizeTypes = "Standard_G1"
	// StandardG2 ...
	StandardG2 VirtualMachineSizeTypes = "Standard_G2"
	// StandardG3 ...
	StandardG3 VirtualMachineSizeTypes = "Standard_G3"
	// StandardG4 ...
	StandardG4 VirtualMachineSizeTypes = "Standard_G4"
	// StandardG5 ...
	StandardG5 VirtualMachineSizeTypes = "Standard_G5"
	// StandardGS1 ...
	StandardGS1 VirtualMachineSizeTypes = "Standard_GS1"
	// StandardGS2 ...
	StandardGS2 VirtualMachineSizeTypes = "Standard_GS2"
	// StandardGS3 ...
	StandardGS3 VirtualMachineSizeTypes = "Standard_GS3"
	// StandardGS4 ...
	StandardGS4 VirtualMachineSizeTypes = "Standard_GS4"
	// StandardGS5 ...
	StandardGS5 VirtualMachineSizeTypes = "Standard_GS5"
)

// PossibleVirtualMachineSizeTypesValues returns an array of possible values for the VirtualMachineSizeTypes const type.
func PossibleVirtualMachineSizeTypesValues() []VirtualMachineSizeTypes {
	return []VirtualMachineSizeTypes{BasicA0, BasicA1, BasicA2, BasicA3, BasicA4, StandardA0, StandardA1, StandardA10, StandardA11, StandardA2, StandardA3, StandardA4, StandardA5, StandardA6, StandardA7, StandardA8, StandardA9, StandardD1, StandardD11, StandardD11V2, StandardD12, StandardD12V2, StandardD13, StandardD13V2, StandardD14, StandardD14V2, StandardD15V2, StandardD1V2, StandardD2, StandardD2V2, StandardD3, StandardD3V2, StandardD4, StandardD4V2, StandardD5V2, StandardDS1, StandardDS11, StandardDS11V2, StandardDS12, StandardDS12V2, StandardDS13, StandardDS13V2, StandardDS14, StandardDS14V2, StandardDS15V2, StandardDS1V2, StandardDS2, StandardDS2V2, StandardDS3, StandardDS3V2, StandardDS4, StandardDS4V2, StandardDS5V2, StandardG1, StandardG2, StandardG3, StandardG4, StandardG5, StandardGS1, StandardGS2, StandardGS3, StandardGS4, StandardGS5}
}
