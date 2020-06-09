package azcompute

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/network/mgmt/2020-03-01/aznetwork"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

const (
	vmName   = "sampleVM"
	username = "sampleuser"
	password = "SamplePassword0"
)

var (
	// Azure location where the resource will be created
	location = os.Getenv("AZURE_LOCATION")
	// Name of the Network Interface to retrieve
	nicName = os.Getenv("AZURE_NIC")
	// Azure resource group to retrieve and create resources
	resourceGroupName = os.Getenv("AZURE_RESOURCE_GROUP")
	// The subscription ID where the resource group exists
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
)

// returns a credential that can be used to authenticate with blob storage
func getCredential() azcore.Credential {
	// NewEnvironmentCredential() will read various environment vars
	// to obtain a credential.  see the documentation for more info.
	cred, err := azidentity.NewEnvironmentCredential(nil)
	if err != nil {
		panic(err)
	}
	return cred
}

func getVMClient() VirtualMachinesOperations {
	vmClient, err := NewClient("https://management.azure.com", getCredential(), nil)
	if err != nil {
		return nil
	}
	return vmClient.VirtualMachinesOperations(subscriptionID)
}

func ExampleVirtualMachinesOperations_BeginCreateOrUpdate() {
	nic, err := getNIC()
	if err != nil {
		panic(err)
	}
	client := getVMClient()
	vm, err := client.BeginCreateOrUpdate(
		context.Background(),
		resourceGroupName,
		vmName,
		VirtualMachine{
			Resource: Resource{
				Location: to.StringPtr(location),
			},
			Properties: &VirtualMachineProperties{
				HardwareProfile: &HardwareProfile{
					VMSize: VirtualMachineSizeTypesStandardA0.ToPtr(),
				},
				StorageProfile: &StorageProfile{
					ImageReference: &ImageReference{
						Publisher: to.StringPtr("MicrosoftWindowsServer"),
						Offer:     to.StringPtr("WindowsServer"),
						Sku:       to.StringPtr("2012-R2-Datacenter"),
						Version:   to.StringPtr("latest"),
					},
				},
				OSProfile: &OSProfile{
					ComputerName:  to.StringPtr(vmName),
					AdminUsername: to.StringPtr(username),
					AdminPassword: to.StringPtr(password),
				},
				NetworkProfile: &NetworkProfile{
					NetworkInterfaces: &[]NetworkInterfaceReference{
						{
							SubResource: SubResource{
								ID: nic.NetworkInterface.ID,
							},
							Properties: &NetworkInterfaceReferenceProperties{
								Primary: to.BoolPtr(true),
							},
						},
					},
				},
			},
		})
	if err != nil {
		panic(err)
	}
	result, err := vm.PollUntilDone(context.Background(), 5*time.Second)
	if err != nil {
		panic(err)
	}
	if result != nil {
		fmt.Println(*vm.VirtualMachine.Name)
	}
	// Output:
	// sampleVM
}

func ExampleVirtualMachinesOperations_Get() {
	client := getVMClient()
	vm, err := client.Get(context.Background(), resourceGroupName, vmName)
	if err != nil {
		panic(err)
	}
	if vm.VirtualMachine != nil {
		fmt.Println(*vm.VirtualMachine.Name)
	}
	// Output:
	// sampleVM
}

func ExampleVirtualMachinesOperations_List() {
	client := getVMClient()
	vmList, err := client.List(resourceGroupName)
	if err != nil {
		panic(err)
	}
	var vmListResult []VirtualMachine
	count := 0
	for vmList.NextPage(context.Background()) {
		resp := vmList.PageResponse()
		if len(*resp.VirtualMachineListResult.Value) == 0 {
			panic("missing payload")
		}
		for _, vm := range *resp.VirtualMachineListResult.Value {
			vmListResult = append(vmListResult, vm)
		}
		count++
	}
	if err = vmList.Err(); err != nil {
		panic(err)
	}
	for _, vm := range vmListResult {
		fmt.Println(*vm.Name)
	}
	// Output:
	// sampleVM
}

// client for using the operations on the NetworkInterfacesOperations
func getNetworkInterfacesClient() aznetwork.NetworkInterfacesOperations {
	nicClient, err := aznetwork.NewClient("https://management.azure.com", getCredential(), nil)
	if err != nil {
		panic(err)
	}
	return nicClient.NetworkInterfacesOperations(subscriptionID)
}

// returns the specified network interface from Azure
func getNIC() (*aznetwork.NetworkInterfaceResponse, error) {
	client := getNetworkInterfacesClient()
	return client.Get(context.Background(), resourceGroupName, nicName, nil)
}
