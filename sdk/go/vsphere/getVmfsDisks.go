// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getVmfsDisks` data source can be used to discover the storage
// devices available on an ESXi host. This data source can be combined with the
// `VmfsDatastore` resource to create VMFS
// datastores based off a set of discovered disks.
//
// ## Example Usage
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
//			_, err = vsphere.GetVmfsDisks(ctx, &vsphere.GetVmfsDisksArgs{
//				HostSystemId: host.Id,
//				Rescan:       pulumi.BoolRef(true),
//				Filter:       pulumi.StringRef("mpx.vmhba1:C0:T[12]:L0"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVmfsDisks(ctx *pulumi.Context, args *GetVmfsDisksArgs, opts ...pulumi.InvokeOption) (*GetVmfsDisksResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVmfsDisksResult
	err := ctx.Invoke("vsphere:index/getVmfsDisks:getVmfsDisks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmfsDisks.
type GetVmfsDisksArgs struct {
	// A regular expression to filter the disks against. Only
	// disks with canonical names that match will be included.
	//
	// > **NOTE:** Using a `filter` is recommended if there is any chance the host
	// will have any specific storage devices added to it that may affect the order of
	// the output `disks` attribute below, which is lexicographically sorted.
	Filter *string `pulumi:"filter"`
	// The managed object ID of
	// the host to look for disks on.
	HostSystemId string `pulumi:"hostSystemId"`
	// Whether or not to rescan storage adapters before
	// searching for disks. This may lengthen the time it takes to perform the
	// search. Default: `false`.
	Rescan *bool `pulumi:"rescan"`
}

// A collection of values returned by getVmfsDisks.
type GetVmfsDisksResult struct {
	// A lexicographically sorted list of devices discovered by the
	// operation, matching the supplied `filter`, if provided.
	Disks        []string `pulumi:"disks"`
	Filter       *string  `pulumi:"filter"`
	HostSystemId string   `pulumi:"hostSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Rescan *bool  `pulumi:"rescan"`
}

func GetVmfsDisksOutput(ctx *pulumi.Context, args GetVmfsDisksOutputArgs, opts ...pulumi.InvokeOption) GetVmfsDisksResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVmfsDisksResultOutput, error) {
			args := v.(GetVmfsDisksArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vsphere:index/getVmfsDisks:getVmfsDisks", args, GetVmfsDisksResultOutput{}, options).(GetVmfsDisksResultOutput), nil
		}).(GetVmfsDisksResultOutput)
}

// A collection of arguments for invoking getVmfsDisks.
type GetVmfsDisksOutputArgs struct {
	// A regular expression to filter the disks against. Only
	// disks with canonical names that match will be included.
	//
	// > **NOTE:** Using a `filter` is recommended if there is any chance the host
	// will have any specific storage devices added to it that may affect the order of
	// the output `disks` attribute below, which is lexicographically sorted.
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// The managed object ID of
	// the host to look for disks on.
	HostSystemId pulumi.StringInput `pulumi:"hostSystemId"`
	// Whether or not to rescan storage adapters before
	// searching for disks. This may lengthen the time it takes to perform the
	// search. Default: `false`.
	Rescan pulumi.BoolPtrInput `pulumi:"rescan"`
}

func (GetVmfsDisksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVmfsDisksArgs)(nil)).Elem()
}

// A collection of values returned by getVmfsDisks.
type GetVmfsDisksResultOutput struct{ *pulumi.OutputState }

func (GetVmfsDisksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVmfsDisksResult)(nil)).Elem()
}

func (o GetVmfsDisksResultOutput) ToGetVmfsDisksResultOutput() GetVmfsDisksResultOutput {
	return o
}

func (o GetVmfsDisksResultOutput) ToGetVmfsDisksResultOutputWithContext(ctx context.Context) GetVmfsDisksResultOutput {
	return o
}

// A lexicographically sorted list of devices discovered by the
// operation, matching the supplied `filter`, if provided.
func (o GetVmfsDisksResultOutput) Disks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVmfsDisksResult) []string { return v.Disks }).(pulumi.StringArrayOutput)
}

func (o GetVmfsDisksResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVmfsDisksResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o GetVmfsDisksResultOutput) HostSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmfsDisksResult) string { return v.HostSystemId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVmfsDisksResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVmfsDisksResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVmfsDisksResultOutput) Rescan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVmfsDisksResult) *bool { return v.Rescan }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVmfsDisksResultOutput{})
}
