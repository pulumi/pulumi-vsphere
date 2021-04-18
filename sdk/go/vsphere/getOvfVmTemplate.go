// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getOvfVmTemplate` data source can be used to submit an OVF to vSphere and extract its hardware
// settings in a form that can be then used as inputs for a `VirtualMachine` resource.
func GetOvfVmTemplate(ctx *pulumi.Context, args *GetOvfVmTemplateArgs, opts ...pulumi.InvokeOption) (*GetOvfVmTemplateResult, error) {
	var rv GetOvfVmTemplateResult
	err := ctx.Invoke("vsphere:index/getOvfVmTemplate:getOvfVmTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOvfVmTemplate.
type GetOvfVmTemplateArgs struct {
	// Allow unverified ssl certificates while deploying ovf/ova from url.
	AllowUnverifiedSslCert *bool `pulumi:"allowUnverifiedSslCert"`
	// The ID of the virtual machine's datastore. The virtual machine configuration is placed here, along with any virtual disks that are created without datastores.
	DatastoreId *string `pulumi:"datastoreId"`
	// The key of the chosen deployment option. If empty, the default option is chosen.
	DeploymentOption *string `pulumi:"deploymentOption"`
	// The disk provisioning. If set, all the disks in the deployed OVF will have
	// the same specified disk type (accepted values {thin, flat, thick, sameAsSource}).
	DiskProvisioning *string `pulumi:"diskProvisioning"`
	// The name of the folder to locate the virtual machine in.
	Folder *string `pulumi:"folder"`
	// The ID of an optional host system to pin the virtual machine to.
	HostSystemId string `pulumi:"hostSystemId"`
	// The IP allocation policy.
	IpAllocationPolicy *string `pulumi:"ipAllocationPolicy"`
	// The IP protocol.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The absolute path to the ovf/ova file in the local system. While deploying from ovf,
	// make sure the other necessary files like the .vmdk files are also in the same directory as the given ovf file.
	LocalOvfPath *string `pulumi:"localOvfPath"`
	// Name of the virtual machine to create.
	Name string `pulumi:"name"`
	// The mapping of name of network identifiers from the ovf descriptor to network UUID in the
	// VI infrastructure.
	OvfNetworkMap map[string]string `pulumi:"ovfNetworkMap"`
	// URL to the remote ovf/ova file to be deployed.
	RemoteOvfUrl *string `pulumi:"remoteOvfUrl"`
	// The ID of a resource pool to put the virtual machine in.
	ResourcePoolId string `pulumi:"resourcePoolId"`
}

// A collection of values returned by getOvfVmTemplate.
type GetOvfVmTemplateResult struct {
	AllowUnverifiedSslCert *bool `pulumi:"allowUnverifiedSslCert"`
	// The guest name for the operating system .
	AlternateGuestName string `pulumi:"alternateGuestName"`
	// User-provided description of the virtual machine.
	Annotation string `pulumi:"annotation"`
	// Allow CPUs to be added to this virtual machine while it is running.
	CpuHotAddEnabled bool `pulumi:"cpuHotAddEnabled"`
	// Allow CPUs to be added to this virtual machine while it is running.
	CpuHotRemoveEnabled           bool    `pulumi:"cpuHotRemoveEnabled"`
	CpuPerformanceCountersEnabled bool    `pulumi:"cpuPerformanceCountersEnabled"`
	DatastoreId                   *string `pulumi:"datastoreId"`
	DeploymentOption              *string `pulumi:"deploymentOption"`
	DiskProvisioning              *string `pulumi:"diskProvisioning"`
	// The firmware interface to use on the virtual machine.
	Firmware string  `pulumi:"firmware"`
	Folder   *string `pulumi:"folder"`
	// The guest ID for the operating system
	GuestId      string `pulumi:"guestId"`
	HostSystemId string `pulumi:"hostSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string  `pulumi:"id"`
	IdeControllerCount int     `pulumi:"ideControllerCount"`
	IpAllocationPolicy *string `pulumi:"ipAllocationPolicy"`
	IpProtocol         *string `pulumi:"ipProtocol"`
	LocalOvfPath       *string `pulumi:"localOvfPath"`
	// The size of the virtual machine's memory, in MB.
	Memory int `pulumi:"memory"`
	// Allow memory to be added to this virtual machine while it is running.
	MemoryHotAddEnabled bool   `pulumi:"memoryHotAddEnabled"`
	Name                string `pulumi:"name"`
	// Enable nested hardware virtualization on this virtual machine, facilitating nested virtualization in the guest.
	NestedHvEnabled bool `pulumi:"nestedHvEnabled"`
	// The number of cores to distribute amongst the CPUs in this virtual machine.
	NumCoresPerSocket int `pulumi:"numCoresPerSocket"`
	// The number of virtual processors to assign to this virtual machine.
	NumCpus             int               `pulumi:"numCpus"`
	OvfNetworkMap       map[string]string `pulumi:"ovfNetworkMap"`
	RemoteOvfUrl        *string           `pulumi:"remoteOvfUrl"`
	ResourcePoolId      string            `pulumi:"resourcePoolId"`
	SataControllerCount int               `pulumi:"sataControllerCount"`
	ScsiControllerCount int               `pulumi:"scsiControllerCount"`
	ScsiType            string            `pulumi:"scsiType"`
	// The swap file placement policy for this virtual machine.
	SwapPlacementPolicy string `pulumi:"swapPlacementPolicy"`
}
