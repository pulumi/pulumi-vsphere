// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getHostPciDevice` data source can be used to discover the DeviceID
// of a vSphere host's PCI device. This can then be used with
// `VirtualMachine`'s `pciDeviceId`.
//
// ## Example Usage
// ### With Vendor ID And Class ID
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v3/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "dc1"
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := "esxi1"
// 		host, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
// 			Name:         &opt1,
// 			DatacenterId: datacenter.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt2 := "123"
// 		opt3 := "456"
// 		_, err = vsphere.GetHostPciDevice(ctx, &vsphere.GetHostPciDeviceArgs{
// 			HostId:   host.Id,
// 			ClassId:  &opt2,
// 			VendorId: &opt3,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### With Name Regular Expression
func GetHostPciDevice(ctx *pulumi.Context, args *GetHostPciDeviceArgs, opts ...pulumi.InvokeOption) (*GetHostPciDeviceResult, error) {
	var rv GetHostPciDeviceResult
	err := ctx.Invoke("vsphere:index/getHostPciDevice:getHostPciDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostPciDevice.
type GetHostPciDeviceArgs struct {
	// The hexadecimal PCI device class ID
	ClassId *string `pulumi:"classId"`
	// The [managed object reference
	// ID][docs-about-morefs] of a host.
	HostId string `pulumi:"hostId"`
	// A regular expression that will be used to match
	// the host PCI device name.
	NameRegex *string `pulumi:"nameRegex"`
	// The hexadecimal PCI device vendor ID.
	VendorId *string `pulumi:"vendorId"`
}

// A collection of values returned by getHostPciDevice.
type GetHostPciDeviceResult struct {
	ClassId *string `pulumi:"classId"`
	HostId  string  `pulumi:"hostId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the PCI device.
	Name      string  `pulumi:"name"`
	NameRegex *string `pulumi:"nameRegex"`
	VendorId  *string `pulumi:"vendorId"`
}
