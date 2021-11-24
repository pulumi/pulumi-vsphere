// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `VirtualMachine` data source can be used to find the UUID of an
// existing virtual machine or template. Its most relevant purpose is for finding
// the UUID of a template to be used as the source for cloning into a new
// `VirtualMachine` resource. It also
// reads the guest ID so that can be supplied as well.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "dc1"
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &GetDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := datacenter.Id
// 		_, err = vsphere.LookupVirtualMachine(ctx, &GetVirtualMachineArgs{
// 			DatacenterId: &opt1,
// 			Name:         "test-vm-template",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVirtualMachine(ctx *pulumi.Context, args *LookupVirtualMachineArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineResult, error) {
	var rv LookupVirtualMachineResult
	err := ctx.Invoke("vsphere:index/getVirtualMachine:getVirtualMachine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualMachine.
type LookupVirtualMachineArgs struct {
	// The alternate guest name of the virtual machine when
	// guestId is a non-specific operating system, like `otherGuest`.
	AlternateGuestName *string `pulumi:"alternateGuestName"`
	// The user-provided description of this virtual machine.
	Annotation                    *string `pulumi:"annotation"`
	BootDelay                     *int    `pulumi:"bootDelay"`
	BootRetryDelay                *int    `pulumi:"bootRetryDelay"`
	BootRetryEnabled              *bool   `pulumi:"bootRetryEnabled"`
	CpuHotAddEnabled              *bool   `pulumi:"cpuHotAddEnabled"`
	CpuHotRemoveEnabled           *bool   `pulumi:"cpuHotRemoveEnabled"`
	CpuLimit                      *int    `pulumi:"cpuLimit"`
	CpuPerformanceCountersEnabled *bool   `pulumi:"cpuPerformanceCountersEnabled"`
	CpuReservation                *int    `pulumi:"cpuReservation"`
	CpuShareCount                 *int    `pulumi:"cpuShareCount"`
	CpuShareLevel                 *string `pulumi:"cpuShareLevel"`
	// The managed object reference
	// ID of the datacenter the virtual machine is located in.
	// This can be omitted if the search path used in `name` is an absolute path.
	// For default datacenters, use the `id` attribute from an empty
	// `Datacenter` data source.
	DatacenterId         *string           `pulumi:"datacenterId"`
	EfiSecureBootEnabled *bool             `pulumi:"efiSecureBootEnabled"`
	EnableDiskUuid       *bool             `pulumi:"enableDiskUuid"`
	EnableLogging        *bool             `pulumi:"enableLogging"`
	EptRviMode           *string           `pulumi:"eptRviMode"`
	ExtraConfig          map[string]string `pulumi:"extraConfig"`
	// The firmware type for this virtual machine. Can be `bios` or `efi`.
	Firmware *string `pulumi:"firmware"`
	// The guest ID of the virtual machine or template.
	GuestId *string `pulumi:"guestId"`
	// The hardware version number on this virtual machine.
	HardwareVersion        *int    `pulumi:"hardwareVersion"`
	HvMode                 *string `pulumi:"hvMode"`
	IdeControllerScanCount *int    `pulumi:"ideControllerScanCount"`
	LatencySensitivity     *string `pulumi:"latencySensitivity"`
	// The size of the virtual machine's memory, in MB.
	Memory              *int    `pulumi:"memory"`
	MemoryHotAddEnabled *bool   `pulumi:"memoryHotAddEnabled"`
	MemoryLimit         *int    `pulumi:"memoryLimit"`
	MemoryReservation   *int    `pulumi:"memoryReservation"`
	MemoryShareCount    *int    `pulumi:"memoryShareCount"`
	MemoryShareLevel    *string `pulumi:"memoryShareLevel"`
	// The name of the virtual machine. This can be a name or
	// path.
	Name            string `pulumi:"name"`
	NestedHvEnabled *bool  `pulumi:"nestedHvEnabled"`
	// The number of cores per socket for this virtual machine.
	NumCoresPerSocket *int `pulumi:"numCoresPerSocket"`
	// The total number of virtual processor cores assigned to this
	// virtual machine.
	NumCpus                            *int    `pulumi:"numCpus"`
	ReplaceTrigger                     *string `pulumi:"replaceTrigger"`
	RunToolsScriptsAfterPowerOn        *bool   `pulumi:"runToolsScriptsAfterPowerOn"`
	RunToolsScriptsAfterResume         *bool   `pulumi:"runToolsScriptsAfterResume"`
	RunToolsScriptsBeforeGuestReboot   *bool   `pulumi:"runToolsScriptsBeforeGuestReboot"`
	RunToolsScriptsBeforeGuestShutdown *bool   `pulumi:"runToolsScriptsBeforeGuestShutdown"`
	RunToolsScriptsBeforeGuestStandby  *bool   `pulumi:"runToolsScriptsBeforeGuestStandby"`
	SataControllerScanCount            *int    `pulumi:"sataControllerScanCount"`
	// The number of SCSI controllers to
	// scan for disk attributes and controller types on. Default: `1`.
	ScsiControllerScanCount      *int                   `pulumi:"scsiControllerScanCount"`
	StoragePolicyId              *string                `pulumi:"storagePolicyId"`
	SwapPlacementPolicy          *string                `pulumi:"swapPlacementPolicy"`
	SyncTimeWithHost             *bool                  `pulumi:"syncTimeWithHost"`
	SyncTimeWithHostPeriodically *bool                  `pulumi:"syncTimeWithHostPeriodically"`
	Vapp                         *GetVirtualMachineVapp `pulumi:"vapp"`
	VbsEnabled                   *bool                  `pulumi:"vbsEnabled"`
	VvtdEnabled                  *bool                  `pulumi:"vvtdEnabled"`
}

// A collection of values returned by getVirtualMachine.
type LookupVirtualMachineResult struct {
	// The alternate guest name of the virtual machine when
	// guestId is a non-specific operating system, like `otherGuest`.
	AlternateGuestName *string `pulumi:"alternateGuestName"`
	// The user-provided description of this virtual machine.
	Annotation                    *string `pulumi:"annotation"`
	BootDelay                     *int    `pulumi:"bootDelay"`
	BootRetryDelay                *int    `pulumi:"bootRetryDelay"`
	BootRetryEnabled              *bool   `pulumi:"bootRetryEnabled"`
	ChangeVersion                 string  `pulumi:"changeVersion"`
	CpuHotAddEnabled              *bool   `pulumi:"cpuHotAddEnabled"`
	CpuHotRemoveEnabled           *bool   `pulumi:"cpuHotRemoveEnabled"`
	CpuLimit                      *int    `pulumi:"cpuLimit"`
	CpuPerformanceCountersEnabled *bool   `pulumi:"cpuPerformanceCountersEnabled"`
	CpuReservation                *int    `pulumi:"cpuReservation"`
	CpuShareCount                 int     `pulumi:"cpuShareCount"`
	CpuShareLevel                 *string `pulumi:"cpuShareLevel"`
	DatacenterId                  *string `pulumi:"datacenterId"`
	// Information about each of the disks on this virtual machine or
	// template. These are sorted by bus and unit number so that they can be applied
	// to a `VirtualMachine` resource in the order the resource expects
	// while cloning. This is useful for discovering certain disk settings while
	// performing a linked clone, as all settings that are output by this data
	// source must be the same on the destination virtual machine as the source.
	// Only the first number of controllers defined by `scsiControllerScanCount`
	// are scanned for disks. The sub-attributes are:
	Disks                []GetVirtualMachineDisk `pulumi:"disks"`
	EfiSecureBootEnabled *bool                   `pulumi:"efiSecureBootEnabled"`
	EnableDiskUuid       *bool                   `pulumi:"enableDiskUuid"`
	EnableLogging        *bool                   `pulumi:"enableLogging"`
	EptRviMode           *string                 `pulumi:"eptRviMode"`
	ExtraConfig          map[string]string       `pulumi:"extraConfig"`
	// The firmware type for this virtual machine. Can be `bios` or `efi`.
	Firmware *string `pulumi:"firmware"`
	// The guest ID of the virtual machine or template.
	GuestId string `pulumi:"guestId"`
	// A list of IP addresses as reported by VMWare tools.
	GuestIpAddresses []string `pulumi:"guestIpAddresses"`
	// The hardware version number on this virtual machine.
	HardwareVersion int     `pulumi:"hardwareVersion"`
	HvMode          *string `pulumi:"hvMode"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string  `pulumi:"id"`
	IdeControllerScanCount *int    `pulumi:"ideControllerScanCount"`
	LatencySensitivity     *string `pulumi:"latencySensitivity"`
	// The size of the virtual machine's memory, in MB.
	Memory              *int    `pulumi:"memory"`
	MemoryHotAddEnabled *bool   `pulumi:"memoryHotAddEnabled"`
	MemoryLimit         *int    `pulumi:"memoryLimit"`
	MemoryReservation   *int    `pulumi:"memoryReservation"`
	MemoryShareCount    int     `pulumi:"memoryShareCount"`
	MemoryShareLevel    *string `pulumi:"memoryShareLevel"`
	Name                string  `pulumi:"name"`
	NestedHvEnabled     *bool   `pulumi:"nestedHvEnabled"`
	// The network interface types for each network
	// interface found on the virtual machine, in device bus order. Will be one of
	// `e1000`, `e1000e`, `pcnet32`, `sriov`, `vmxnet2`, or `vmxnet3`.
	NetworkInterfaceTypes []string `pulumi:"networkInterfaceTypes"`
	// Information about each of the network interfaces on this
	// virtual machine or template. These are sorted by device bus order so that they
	// can be applied to a `VirtualMachine` resource in the order the resource
	// expects while cloning. This is useful for discovering certain network interface
	// settings while performing a linked clone, as all settings that are output by this
	// data source must be the same on the destination virtual machine as the source.
	// The sub-attributes are:
	NetworkInterfaces []GetVirtualMachineNetworkInterface `pulumi:"networkInterfaces"`
	// The number of cores per socket for this virtual machine.
	NumCoresPerSocket *int `pulumi:"numCoresPerSocket"`
	// The total number of virtual processor cores assigned to this
	// virtual machine.
	NumCpus                            *int    `pulumi:"numCpus"`
	ReplaceTrigger                     *string `pulumi:"replaceTrigger"`
	RunToolsScriptsAfterPowerOn        *bool   `pulumi:"runToolsScriptsAfterPowerOn"`
	RunToolsScriptsAfterResume         *bool   `pulumi:"runToolsScriptsAfterResume"`
	RunToolsScriptsBeforeGuestReboot   *bool   `pulumi:"runToolsScriptsBeforeGuestReboot"`
	RunToolsScriptsBeforeGuestShutdown *bool   `pulumi:"runToolsScriptsBeforeGuestShutdown"`
	RunToolsScriptsBeforeGuestStandby  *bool   `pulumi:"runToolsScriptsBeforeGuestStandby"`
	SataControllerScanCount            *int    `pulumi:"sataControllerScanCount"`
	// Mode for sharing the SCSI bus. The modes are
	// physicalSharing, virtualSharing, and noSharing. Only the first number of
	// controllers defined by `scsiControllerScanCount` are scanned.
	ScsiBusSharing          string `pulumi:"scsiBusSharing"`
	ScsiControllerScanCount *int   `pulumi:"scsiControllerScanCount"`
	// The common type of all SCSI controllers on this virtual machine.
	// Will be one of `lsilogic` (LSI Logic Parallel), `lsilogic-sas` (LSI Logic
	// SAS), `pvscsi` (VMware Paravirtual), `buslogic` (BusLogic), or `mixed` when
	// there are multiple controller types. Only the first number of controllers
	// defined by `scsiControllerScanCount` are scanned.
	ScsiType                     string                 `pulumi:"scsiType"`
	StoragePolicyId              string                 `pulumi:"storagePolicyId"`
	SwapPlacementPolicy          *string                `pulumi:"swapPlacementPolicy"`
	SyncTimeWithHost             *bool                  `pulumi:"syncTimeWithHost"`
	SyncTimeWithHostPeriodically *bool                  `pulumi:"syncTimeWithHostPeriodically"`
	Uuid                         string                 `pulumi:"uuid"`
	Vapp                         *GetVirtualMachineVapp `pulumi:"vapp"`
	VappTransports               []string               `pulumi:"vappTransports"`
	VbsEnabled                   *bool                  `pulumi:"vbsEnabled"`
	VvtdEnabled                  *bool                  `pulumi:"vvtdEnabled"`
}

func LookupVirtualMachineOutput(ctx *pulumi.Context, args LookupVirtualMachineOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineResult, error) {
			args := v.(LookupVirtualMachineArgs)
			r, err := LookupVirtualMachine(ctx, &args, opts...)
			return *r, err
		}).(LookupVirtualMachineResultOutput)
}

// A collection of arguments for invoking getVirtualMachine.
type LookupVirtualMachineOutputArgs struct {
	// The alternate guest name of the virtual machine when
	// guestId is a non-specific operating system, like `otherGuest`.
	AlternateGuestName pulumi.StringPtrInput `pulumi:"alternateGuestName"`
	// The user-provided description of this virtual machine.
	Annotation                    pulumi.StringPtrInput `pulumi:"annotation"`
	BootDelay                     pulumi.IntPtrInput    `pulumi:"bootDelay"`
	BootRetryDelay                pulumi.IntPtrInput    `pulumi:"bootRetryDelay"`
	BootRetryEnabled              pulumi.BoolPtrInput   `pulumi:"bootRetryEnabled"`
	CpuHotAddEnabled              pulumi.BoolPtrInput   `pulumi:"cpuHotAddEnabled"`
	CpuHotRemoveEnabled           pulumi.BoolPtrInput   `pulumi:"cpuHotRemoveEnabled"`
	CpuLimit                      pulumi.IntPtrInput    `pulumi:"cpuLimit"`
	CpuPerformanceCountersEnabled pulumi.BoolPtrInput   `pulumi:"cpuPerformanceCountersEnabled"`
	CpuReservation                pulumi.IntPtrInput    `pulumi:"cpuReservation"`
	CpuShareCount                 pulumi.IntPtrInput    `pulumi:"cpuShareCount"`
	CpuShareLevel                 pulumi.StringPtrInput `pulumi:"cpuShareLevel"`
	// The managed object reference
	// ID of the datacenter the virtual machine is located in.
	// This can be omitted if the search path used in `name` is an absolute path.
	// For default datacenters, use the `id` attribute from an empty
	// `Datacenter` data source.
	DatacenterId         pulumi.StringPtrInput `pulumi:"datacenterId"`
	EfiSecureBootEnabled pulumi.BoolPtrInput   `pulumi:"efiSecureBootEnabled"`
	EnableDiskUuid       pulumi.BoolPtrInput   `pulumi:"enableDiskUuid"`
	EnableLogging        pulumi.BoolPtrInput   `pulumi:"enableLogging"`
	EptRviMode           pulumi.StringPtrInput `pulumi:"eptRviMode"`
	ExtraConfig          pulumi.StringMapInput `pulumi:"extraConfig"`
	// The firmware type for this virtual machine. Can be `bios` or `efi`.
	Firmware pulumi.StringPtrInput `pulumi:"firmware"`
	// The guest ID of the virtual machine or template.
	GuestId pulumi.StringPtrInput `pulumi:"guestId"`
	// The hardware version number on this virtual machine.
	HardwareVersion        pulumi.IntPtrInput    `pulumi:"hardwareVersion"`
	HvMode                 pulumi.StringPtrInput `pulumi:"hvMode"`
	IdeControllerScanCount pulumi.IntPtrInput    `pulumi:"ideControllerScanCount"`
	LatencySensitivity     pulumi.StringPtrInput `pulumi:"latencySensitivity"`
	// The size of the virtual machine's memory, in MB.
	Memory              pulumi.IntPtrInput    `pulumi:"memory"`
	MemoryHotAddEnabled pulumi.BoolPtrInput   `pulumi:"memoryHotAddEnabled"`
	MemoryLimit         pulumi.IntPtrInput    `pulumi:"memoryLimit"`
	MemoryReservation   pulumi.IntPtrInput    `pulumi:"memoryReservation"`
	MemoryShareCount    pulumi.IntPtrInput    `pulumi:"memoryShareCount"`
	MemoryShareLevel    pulumi.StringPtrInput `pulumi:"memoryShareLevel"`
	// The name of the virtual machine. This can be a name or
	// path.
	Name            pulumi.StringInput  `pulumi:"name"`
	NestedHvEnabled pulumi.BoolPtrInput `pulumi:"nestedHvEnabled"`
	// The number of cores per socket for this virtual machine.
	NumCoresPerSocket pulumi.IntPtrInput `pulumi:"numCoresPerSocket"`
	// The total number of virtual processor cores assigned to this
	// virtual machine.
	NumCpus                            pulumi.IntPtrInput    `pulumi:"numCpus"`
	ReplaceTrigger                     pulumi.StringPtrInput `pulumi:"replaceTrigger"`
	RunToolsScriptsAfterPowerOn        pulumi.BoolPtrInput   `pulumi:"runToolsScriptsAfterPowerOn"`
	RunToolsScriptsAfterResume         pulumi.BoolPtrInput   `pulumi:"runToolsScriptsAfterResume"`
	RunToolsScriptsBeforeGuestReboot   pulumi.BoolPtrInput   `pulumi:"runToolsScriptsBeforeGuestReboot"`
	RunToolsScriptsBeforeGuestShutdown pulumi.BoolPtrInput   `pulumi:"runToolsScriptsBeforeGuestShutdown"`
	RunToolsScriptsBeforeGuestStandby  pulumi.BoolPtrInput   `pulumi:"runToolsScriptsBeforeGuestStandby"`
	SataControllerScanCount            pulumi.IntPtrInput    `pulumi:"sataControllerScanCount"`
	// The number of SCSI controllers to
	// scan for disk attributes and controller types on. Default: `1`.
	ScsiControllerScanCount      pulumi.IntPtrInput            `pulumi:"scsiControllerScanCount"`
	StoragePolicyId              pulumi.StringPtrInput         `pulumi:"storagePolicyId"`
	SwapPlacementPolicy          pulumi.StringPtrInput         `pulumi:"swapPlacementPolicy"`
	SyncTimeWithHost             pulumi.BoolPtrInput           `pulumi:"syncTimeWithHost"`
	SyncTimeWithHostPeriodically pulumi.BoolPtrInput           `pulumi:"syncTimeWithHostPeriodically"`
	Vapp                         GetVirtualMachineVappPtrInput `pulumi:"vapp"`
	VbsEnabled                   pulumi.BoolPtrInput           `pulumi:"vbsEnabled"`
	VvtdEnabled                  pulumi.BoolPtrInput           `pulumi:"vvtdEnabled"`
}

func (LookupVirtualMachineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualMachine.
type LookupVirtualMachineResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineResult)(nil)).Elem()
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutput() LookupVirtualMachineResultOutput {
	return o
}

func (o LookupVirtualMachineResultOutput) ToLookupVirtualMachineResultOutputWithContext(ctx context.Context) LookupVirtualMachineResultOutput {
	return o
}

// The alternate guest name of the virtual machine when
// guestId is a non-specific operating system, like `otherGuest`.
func (o LookupVirtualMachineResultOutput) AlternateGuestName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.AlternateGuestName }).(pulumi.StringPtrOutput)
}

// The user-provided description of this virtual machine.
func (o LookupVirtualMachineResultOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Annotation }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) BootDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.BootDelay }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) BootRetryDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.BootRetryDelay }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) BootRetryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.BootRetryEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ChangeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ChangeVersion }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) CpuHotAddEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.CpuHotAddEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) CpuHotRemoveEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.CpuHotRemoveEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) CpuLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.CpuLimit }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) CpuPerformanceCountersEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.CpuPerformanceCountersEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) CpuReservation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.CpuReservation }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) CpuShareCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) int { return v.CpuShareCount }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineResultOutput) CpuShareLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.CpuShareLevel }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

// Information about each of the disks on this virtual machine or
// template. These are sorted by bus and unit number so that they can be applied
// to a `VirtualMachine` resource in the order the resource expects
// while cloning. This is useful for discovering certain disk settings while
// performing a linked clone, as all settings that are output by this data
// source must be the same on the destination virtual machine as the source.
// Only the first number of controllers defined by `scsiControllerScanCount`
// are scanned for disks. The sub-attributes are:
func (o LookupVirtualMachineResultOutput) Disks() GetVirtualMachineDiskArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []GetVirtualMachineDisk { return v.Disks }).(GetVirtualMachineDiskArrayOutput)
}

func (o LookupVirtualMachineResultOutput) EfiSecureBootEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.EfiSecureBootEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) EnableDiskUuid() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.EnableDiskUuid }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) EnableLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.EnableLogging }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) EptRviMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.EptRviMode }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ExtraConfig() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) map[string]string { return v.ExtraConfig }).(pulumi.StringMapOutput)
}

// The firmware type for this virtual machine. Can be `bios` or `efi`.
func (o LookupVirtualMachineResultOutput) Firmware() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.Firmware }).(pulumi.StringPtrOutput)
}

// The guest ID of the virtual machine or template.
func (o LookupVirtualMachineResultOutput) GuestId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.GuestId }).(pulumi.StringOutput)
}

// A list of IP addresses as reported by VMWare tools.
func (o LookupVirtualMachineResultOutput) GuestIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.GuestIpAddresses }).(pulumi.StringArrayOutput)
}

// The hardware version number on this virtual machine.
func (o LookupVirtualMachineResultOutput) HardwareVersion() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) int { return v.HardwareVersion }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineResultOutput) HvMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.HvMode }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualMachineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) IdeControllerScanCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.IdeControllerScanCount }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) LatencySensitivity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.LatencySensitivity }).(pulumi.StringPtrOutput)
}

// The size of the virtual machine's memory, in MB.
func (o LookupVirtualMachineResultOutput) Memory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.Memory }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) MemoryHotAddEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.MemoryHotAddEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) MemoryLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.MemoryLimit }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) MemoryReservation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.MemoryReservation }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) MemoryShareCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) int { return v.MemoryShareCount }).(pulumi.IntOutput)
}

func (o LookupVirtualMachineResultOutput) MemoryShareLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.MemoryShareLevel }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) NestedHvEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.NestedHvEnabled }).(pulumi.BoolPtrOutput)
}

// The network interface types for each network
// interface found on the virtual machine, in device bus order. Will be one of
// `e1000`, `e1000e`, `pcnet32`, `sriov`, `vmxnet2`, or `vmxnet3`.
func (o LookupVirtualMachineResultOutput) NetworkInterfaceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.NetworkInterfaceTypes }).(pulumi.StringArrayOutput)
}

// Information about each of the network interfaces on this
// virtual machine or template. These are sorted by device bus order so that they
// can be applied to a `VirtualMachine` resource in the order the resource
// expects while cloning. This is useful for discovering certain network interface
// settings while performing a linked clone, as all settings that are output by this
// data source must be the same on the destination virtual machine as the source.
// The sub-attributes are:
func (o LookupVirtualMachineResultOutput) NetworkInterfaces() GetVirtualMachineNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []GetVirtualMachineNetworkInterface { return v.NetworkInterfaces }).(GetVirtualMachineNetworkInterfaceArrayOutput)
}

// The number of cores per socket for this virtual machine.
func (o LookupVirtualMachineResultOutput) NumCoresPerSocket() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.NumCoresPerSocket }).(pulumi.IntPtrOutput)
}

// The total number of virtual processor cores assigned to this
// virtual machine.
func (o LookupVirtualMachineResultOutput) NumCpus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.NumCpus }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineResultOutput) ReplaceTrigger() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.ReplaceTrigger }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) RunToolsScriptsAfterPowerOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.RunToolsScriptsAfterPowerOn }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) RunToolsScriptsAfterResume() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.RunToolsScriptsAfterResume }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) RunToolsScriptsBeforeGuestReboot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.RunToolsScriptsBeforeGuestReboot }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) RunToolsScriptsBeforeGuestShutdown() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.RunToolsScriptsBeforeGuestShutdown }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) RunToolsScriptsBeforeGuestStandby() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.RunToolsScriptsBeforeGuestStandby }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) SataControllerScanCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.SataControllerScanCount }).(pulumi.IntPtrOutput)
}

// Mode for sharing the SCSI bus. The modes are
// physicalSharing, virtualSharing, and noSharing. Only the first number of
// controllers defined by `scsiControllerScanCount` are scanned.
func (o LookupVirtualMachineResultOutput) ScsiBusSharing() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ScsiBusSharing }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) ScsiControllerScanCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *int { return v.ScsiControllerScanCount }).(pulumi.IntPtrOutput)
}

// The common type of all SCSI controllers on this virtual machine.
// Will be one of `lsilogic` (LSI Logic Parallel), `lsilogic-sas` (LSI Logic
// SAS), `pvscsi` (VMware Paravirtual), `buslogic` (BusLogic), or `mixed` when
// there are multiple controller types. Only the first number of controllers
// defined by `scsiControllerScanCount` are scanned.
func (o LookupVirtualMachineResultOutput) ScsiType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.ScsiType }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) StoragePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.StoragePolicyId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) SwapPlacementPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *string { return v.SwapPlacementPolicy }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineResultOutput) SyncTimeWithHost() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.SyncTimeWithHost }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) SyncTimeWithHostPeriodically() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.SyncTimeWithHostPeriodically }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineResultOutput) Vapp() GetVirtualMachineVappPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *GetVirtualMachineVapp { return v.Vapp }).(GetVirtualMachineVappPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VappTransports() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) []string { return v.VappTransports }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualMachineResultOutput) VbsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.VbsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineResultOutput) VvtdEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineResult) *bool { return v.VvtdEnabled }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineResultOutput{})
}
