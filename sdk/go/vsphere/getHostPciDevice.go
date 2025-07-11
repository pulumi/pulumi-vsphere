// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getHostPciDevice` data source can be used to discover the device ID
// of a vSphere host's PCI device. This can then be used with
// `VirtualMachine`'s `pciDeviceId`.
//
// ## Example Usage
//
// ### With Vendor ID And Class ID
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
//				Name: pulumi.StringRef("dc-01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			host, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
//				Name:         pulumi.StringRef("esxi-01.example.com"),
//				DatacenterId: datacenter.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.GetHostPciDevice(ctx, &vsphere.GetHostPciDeviceArgs{
//				HostId:   host.Id,
//				ClassId:  pulumi.StringRef("123"),
//				VendorId: pulumi.StringRef("456"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### With Name Regular Expression
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//	   pulumi.Run(func(ctx *pulumi.Context) error {
//	       datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
//	           Name: pulumi.StringRef("dc-01"),
//	       }, nil)
//	       if err != nil {
//	           return err
//	       }
//	       host, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
//	           Name:         pulumi.StringRef("esxi-01.example.com"),
//	           DatacenterId: datacenter.Id,
//	       }, nil)
//	       if err != nil {
//	           return err
//	       }
//	       _, err = vsphere.GetHostPciDevice(ctx, &vsphere.GetHostPciDeviceArgs{
//	           HostId:    host.Id,
//	           NameRegex: pulumi.StringRef("MMC"),
//	       }, nil)
//	       if err != nil {
//	           return err
//	       }
//	       return nil
//	   })
//	}
//
// ```
func GetHostPciDevice(ctx *pulumi.Context, args *GetHostPciDeviceArgs, opts ...pulumi.InvokeOption) (*GetHostPciDeviceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
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
	//
	// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
	//
	// > **NOTE:** `nameRegex`, `vendorId`, and `classId` can all be used together.
	ClassId *string `pulumi:"classId"`
	// The [managed object reference ID][docs-about-morefs] of
	// a host.
	HostId string `pulumi:"hostId"`
	// A regular expression that will be used to match the
	// host PCI device name.
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

func GetHostPciDeviceOutput(ctx *pulumi.Context, args GetHostPciDeviceOutputArgs, opts ...pulumi.InvokeOption) GetHostPciDeviceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetHostPciDeviceResultOutput, error) {
			args := v.(GetHostPciDeviceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vsphere:index/getHostPciDevice:getHostPciDevice", args, GetHostPciDeviceResultOutput{}, options).(GetHostPciDeviceResultOutput), nil
		}).(GetHostPciDeviceResultOutput)
}

// A collection of arguments for invoking getHostPciDevice.
type GetHostPciDeviceOutputArgs struct {
	// The hexadecimal PCI device class ID
	//
	// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
	//
	// > **NOTE:** `nameRegex`, `vendorId`, and `classId` can all be used together.
	ClassId pulumi.StringPtrInput `pulumi:"classId"`
	// The [managed object reference ID][docs-about-morefs] of
	// a host.
	HostId pulumi.StringInput `pulumi:"hostId"`
	// A regular expression that will be used to match the
	// host PCI device name.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The hexadecimal PCI device vendor ID.
	VendorId pulumi.StringPtrInput `pulumi:"vendorId"`
}

func (GetHostPciDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostPciDeviceArgs)(nil)).Elem()
}

// A collection of values returned by getHostPciDevice.
type GetHostPciDeviceResultOutput struct{ *pulumi.OutputState }

func (GetHostPciDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostPciDeviceResult)(nil)).Elem()
}

func (o GetHostPciDeviceResultOutput) ToGetHostPciDeviceResultOutput() GetHostPciDeviceResultOutput {
	return o
}

func (o GetHostPciDeviceResultOutput) ToGetHostPciDeviceResultOutputWithContext(ctx context.Context) GetHostPciDeviceResultOutput {
	return o
}

func (o GetHostPciDeviceResultOutput) ClassId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostPciDeviceResult) *string { return v.ClassId }).(pulumi.StringPtrOutput)
}

func (o GetHostPciDeviceResultOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostPciDeviceResult) string { return v.HostId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHostPciDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostPciDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the PCI device.
func (o GetHostPciDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostPciDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetHostPciDeviceResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostPciDeviceResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetHostPciDeviceResultOutput) VendorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostPciDeviceResult) *string { return v.VendorId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHostPciDeviceResultOutput{})
}
