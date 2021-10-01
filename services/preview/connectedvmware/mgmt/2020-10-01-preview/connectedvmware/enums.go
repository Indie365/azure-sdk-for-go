package connectedvmware

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// DiskMode enumerates the values for disk mode.
type DiskMode string

const (
	// IndependentNonpersistent ...
	IndependentNonpersistent DiskMode = "independent_nonpersistent"
	// IndependentPersistent ...
	IndependentPersistent DiskMode = "independent_persistent"
	// Persistent ...
	Persistent DiskMode = "persistent"
)

// PossibleDiskModeValues returns an array of possible values for the DiskMode const type.
func PossibleDiskModeValues() []DiskMode {
	return []DiskMode{IndependentNonpersistent, IndependentPersistent, Persistent}
}

// DiskType enumerates the values for disk type.
type DiskType string

const (
	// Flat ...
	Flat DiskType = "flat"
	// Pmem ...
	Pmem DiskType = "pmem"
	// Rawphysical ...
	Rawphysical DiskType = "rawphysical"
	// Rawvirtual ...
	Rawvirtual DiskType = "rawvirtual"
	// Sesparse ...
	Sesparse DiskType = "sesparse"
	// Sparse ...
	Sparse DiskType = "sparse"
	// Unknown ...
	Unknown DiskType = "unknown"
)

// PossibleDiskTypeValues returns an array of possible values for the DiskType const type.
func PossibleDiskTypeValues() []DiskType {
	return []DiskType{Flat, Pmem, Rawphysical, Rawvirtual, Sesparse, Sparse, Unknown}
}

// FirmwareType enumerates the values for firmware type.
type FirmwareType string

const (
	// Bios ...
	Bios FirmwareType = "bios"
	// Efi ...
	Efi FirmwareType = "efi"
)

// PossibleFirmwareTypeValues returns an array of possible values for the FirmwareType const type.
func PossibleFirmwareTypeValues() []FirmwareType {
	return []FirmwareType{Bios, Efi}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// None ...
	None IdentityType = "None"
	// SystemAssigned ...
	SystemAssigned IdentityType = "SystemAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{None, SystemAssigned}
}

// InventoryType enumerates the values for inventory type.
type InventoryType string

const (
	// InventoryTypeCluster ...
	InventoryTypeCluster InventoryType = "Cluster"
	// InventoryTypeDatastore ...
	InventoryTypeDatastore InventoryType = "Datastore"
	// InventoryTypeHost ...
	InventoryTypeHost InventoryType = "Host"
	// InventoryTypeResourcePool ...
	InventoryTypeResourcePool InventoryType = "ResourcePool"
	// InventoryTypeVirtualMachine ...
	InventoryTypeVirtualMachine InventoryType = "VirtualMachine"
	// InventoryTypeVirtualMachineTemplate ...
	InventoryTypeVirtualMachineTemplate InventoryType = "VirtualMachineTemplate"
	// InventoryTypeVirtualNetwork ...
	InventoryTypeVirtualNetwork InventoryType = "VirtualNetwork"
)

// PossibleInventoryTypeValues returns an array of possible values for the InventoryType const type.
func PossibleInventoryTypeValues() []InventoryType {
	return []InventoryType{InventoryTypeCluster, InventoryTypeDatastore, InventoryTypeHost, InventoryTypeResourcePool, InventoryTypeVirtualMachine, InventoryTypeVirtualMachineTemplate, InventoryTypeVirtualNetwork}
}

// InventoryTypeBasicInventoryItemProperties enumerates the values for inventory type basic inventory item
// properties.
type InventoryTypeBasicInventoryItemProperties string

const (
	// InventoryTypeCluster1 ...
	InventoryTypeCluster1 InventoryTypeBasicInventoryItemProperties = "Cluster"
	// InventoryTypeDatastore1 ...
	InventoryTypeDatastore1 InventoryTypeBasicInventoryItemProperties = "Datastore"
	// InventoryTypeHost1 ...
	InventoryTypeHost1 InventoryTypeBasicInventoryItemProperties = "Host"
	// InventoryTypeInventoryItemProperties ...
	InventoryTypeInventoryItemProperties InventoryTypeBasicInventoryItemProperties = "InventoryItemProperties"
	// InventoryTypeResourcePool1 ...
	InventoryTypeResourcePool1 InventoryTypeBasicInventoryItemProperties = "ResourcePool"
	// InventoryTypeVirtualMachine1 ...
	InventoryTypeVirtualMachine1 InventoryTypeBasicInventoryItemProperties = "VirtualMachine"
	// InventoryTypeVirtualMachineTemplate1 ...
	InventoryTypeVirtualMachineTemplate1 InventoryTypeBasicInventoryItemProperties = "VirtualMachineTemplate"
	// InventoryTypeVirtualNetwork1 ...
	InventoryTypeVirtualNetwork1 InventoryTypeBasicInventoryItemProperties = "VirtualNetwork"
)

// PossibleInventoryTypeBasicInventoryItemPropertiesValues returns an array of possible values for the InventoryTypeBasicInventoryItemProperties const type.
func PossibleInventoryTypeBasicInventoryItemPropertiesValues() []InventoryTypeBasicInventoryItemProperties {
	return []InventoryTypeBasicInventoryItemProperties{InventoryTypeCluster1, InventoryTypeDatastore1, InventoryTypeHost1, InventoryTypeInventoryItemProperties, InventoryTypeResourcePool1, InventoryTypeVirtualMachine1, InventoryTypeVirtualMachineTemplate1, InventoryTypeVirtualNetwork1}
}

// IPAddressAllocationMethod enumerates the values for ip address allocation method.
type IPAddressAllocationMethod string

const (
	// Dynamic ...
	Dynamic IPAddressAllocationMethod = "dynamic"
	// Linklayer ...
	Linklayer IPAddressAllocationMethod = "linklayer"
	// Other ...
	Other IPAddressAllocationMethod = "other"
	// Random ...
	Random IPAddressAllocationMethod = "random"
	// Static ...
	Static IPAddressAllocationMethod = "static"
	// Unset ...
	Unset IPAddressAllocationMethod = "unset"
)

// PossibleIPAddressAllocationMethodValues returns an array of possible values for the IPAddressAllocationMethod const type.
func PossibleIPAddressAllocationMethodValues() []IPAddressAllocationMethod {
	return []IPAddressAllocationMethod{Dynamic, Linklayer, Other, Random, Static, Unset}
}

// NICType enumerates the values for nic type.
type NICType string

const (
	// E1000 ...
	E1000 NICType = "e1000"
	// E1000e ...
	E1000e NICType = "e1000e"
	// Pcnet32 ...
	Pcnet32 NICType = "pcnet32"
	// Vmxnet ...
	Vmxnet NICType = "vmxnet"
	// Vmxnet2 ...
	Vmxnet2 NICType = "vmxnet2"
	// Vmxnet3 ...
	Vmxnet3 NICType = "vmxnet3"
)

// PossibleNICTypeValues returns an array of possible values for the NICType const type.
func PossibleNICTypeValues() []NICType {
	return []NICType{E1000, E1000e, Pcnet32, Vmxnet, Vmxnet2, Vmxnet3}
}

// OsType enumerates the values for os type.
type OsType string

const (
	// OsTypeLinux ...
	OsTypeLinux OsType = "Linux"
	// OsTypeOther ...
	OsTypeOther OsType = "Other"
	// OsTypeWindows ...
	OsTypeWindows OsType = "Windows"
)

// PossibleOsTypeValues returns an array of possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{OsTypeLinux, OsTypeOther, OsTypeWindows}
}

// PowerOnBootOption enumerates the values for power on boot option.
type PowerOnBootOption string

const (
	// Disabled ...
	Disabled PowerOnBootOption = "disabled"
	// Enabled ...
	Enabled PowerOnBootOption = "enabled"
)

// PossiblePowerOnBootOptionValues returns an array of possible values for the PowerOnBootOption const type.
func PossiblePowerOnBootOptionValues() []PowerOnBootOption {
	return []PowerOnBootOption{Disabled, Enabled}
}

// ProvisioningAction enumerates the values for provisioning action.
type ProvisioningAction string

const (
	// Install ...
	Install ProvisioningAction = "install"
	// Repair ...
	Repair ProvisioningAction = "repair"
	// Uninstall ...
	Uninstall ProvisioningAction = "uninstall"
)

// PossibleProvisioningActionValues returns an array of possible values for the ProvisioningAction const type.
func PossibleProvisioningActionValues() []ProvisioningAction {
	return []ProvisioningAction{Install, Repair, Uninstall}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Created ...
	Created ProvisioningState = "Created"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Created, Deleting, Failed, Provisioning, Succeeded, Updating}
}

// SCSIControllerType enumerates the values for scsi controller type.
type SCSIControllerType string

const (
	// Buslogic ...
	Buslogic SCSIControllerType = "buslogic"
	// Lsilogic ...
	Lsilogic SCSIControllerType = "lsilogic"
	// Lsilogicsas ...
	Lsilogicsas SCSIControllerType = "lsilogicsas"
	// Pvscsi ...
	Pvscsi SCSIControllerType = "pvscsi"
)

// PossibleSCSIControllerTypeValues returns an array of possible values for the SCSIControllerType const type.
func PossibleSCSIControllerTypeValues() []SCSIControllerType {
	return []SCSIControllerType{Buslogic, Lsilogic, Lsilogicsas, Pvscsi}
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

// StatusTypes enumerates the values for status types.
type StatusTypes string

const (
	// StatusTypesConnected ...
	StatusTypesConnected StatusTypes = "Connected"
	// StatusTypesDisconnected ...
	StatusTypesDisconnected StatusTypes = "Disconnected"
	// StatusTypesError ...
	StatusTypesError StatusTypes = "Error"
)

// PossibleStatusTypesValues returns an array of possible values for the StatusTypes const type.
func PossibleStatusTypesValues() []StatusTypes {
	return []StatusTypes{StatusTypesConnected, StatusTypesDisconnected, StatusTypesError}
}

// VirtualSCSISharing enumerates the values for virtual scsi sharing.
type VirtualSCSISharing string

const (
	// NoSharing ...
	NoSharing VirtualSCSISharing = "noSharing"
	// PhysicalSharing ...
	PhysicalSharing VirtualSCSISharing = "physicalSharing"
	// VirtualSharing ...
	VirtualSharing VirtualSCSISharing = "virtualSharing"
)

// PossibleVirtualSCSISharingValues returns an array of possible values for the VirtualSCSISharing const type.
func PossibleVirtualSCSISharingValues() []VirtualSCSISharing {
	return []VirtualSCSISharing{NoSharing, PhysicalSharing, VirtualSharing}
}
