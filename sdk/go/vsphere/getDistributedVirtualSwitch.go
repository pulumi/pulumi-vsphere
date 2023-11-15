// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `DistributedVirtualSwitch` data source can be used to discover
// the ID and uplink data of a of a vSphere distributed switch (VDS). This
// can then be used with resources or data sources that require a VDS, such as the
// `DistributedPortGroup` resource, for which
// an example is shown below.
//
// > **NOTE:** This data source requires vCenter Server and is not available on
// direct ESXi host connections.
//
// ## Example Usage
//
// The following example locates a distributed switch named `vds-01`, in the
// datacenter `dc-01`. It then uses this distributed switch to set up a
// `DistributedPortGroup` resource that uses the first uplink as a
// primary uplink and the second uplink as a secondary.
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
//			vds, err := vsphere.LookupDistributedVirtualSwitch(ctx, &vsphere.LookupDistributedVirtualSwitchArgs{
//				Name:         "vds-01",
//				DatacenterId: pulumi.StringRef(datacenter.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.NewDistributedPortGroup(ctx, "dvpg", &vsphere.DistributedPortGroupArgs{
//				DistributedVirtualSwitchUuid: *pulumi.String(vds.Id),
//				ActiveUplinks: pulumi.StringArray{
//					*pulumi.String(vds.Uplinks[0]),
//				},
//				StandbyUplinks: pulumi.StringArray{
//					*pulumi.String(vds.Uplinks[1]),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDistributedVirtualSwitch(ctx *pulumi.Context, args *LookupDistributedVirtualSwitchArgs, opts ...pulumi.InvokeOption) (*LookupDistributedVirtualSwitchResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDistributedVirtualSwitchResult
	err := ctx.Invoke("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDistributedVirtualSwitch.
type LookupDistributedVirtualSwitchArgs struct {
	// The managed object reference ID
	// of the datacenter the VDS is located in. This can be omitted if the search
	// path used in `name` is an absolute path. For default datacenters, use the `id`
	// attribute from an empty `Datacenter` data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name of the VDS. This can be a name or path.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDistributedVirtualSwitch.
type LookupDistributedVirtualSwitchResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The list of the uplinks on this vSphere distributed switch, as per the
	// `uplinks` argument to the
	// `DistributedVirtualSwitch`
	// resource.
	Uplinks []string `pulumi:"uplinks"`
}

func LookupDistributedVirtualSwitchOutput(ctx *pulumi.Context, args LookupDistributedVirtualSwitchOutputArgs, opts ...pulumi.InvokeOption) LookupDistributedVirtualSwitchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDistributedVirtualSwitchResult, error) {
			args := v.(LookupDistributedVirtualSwitchArgs)
			r, err := LookupDistributedVirtualSwitch(ctx, &args, opts...)
			var s LookupDistributedVirtualSwitchResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDistributedVirtualSwitchResultOutput)
}

// A collection of arguments for invoking getDistributedVirtualSwitch.
type LookupDistributedVirtualSwitchOutputArgs struct {
	// The managed object reference ID
	// of the datacenter the VDS is located in. This can be omitted if the search
	// path used in `name` is an absolute path. For default datacenters, use the `id`
	// attribute from an empty `Datacenter` data source.
	DatacenterId pulumi.StringPtrInput `pulumi:"datacenterId"`
	// The name of the VDS. This can be a name or path.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDistributedVirtualSwitchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributedVirtualSwitchArgs)(nil)).Elem()
}

// A collection of values returned by getDistributedVirtualSwitch.
type LookupDistributedVirtualSwitchResultOutput struct{ *pulumi.OutputState }

func (LookupDistributedVirtualSwitchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributedVirtualSwitchResult)(nil)).Elem()
}

func (o LookupDistributedVirtualSwitchResultOutput) ToLookupDistributedVirtualSwitchResultOutput() LookupDistributedVirtualSwitchResultOutput {
	return o
}

func (o LookupDistributedVirtualSwitchResultOutput) ToLookupDistributedVirtualSwitchResultOutputWithContext(ctx context.Context) LookupDistributedVirtualSwitchResultOutput {
	return o
}

func (o LookupDistributedVirtualSwitchResultOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDistributedVirtualSwitchResult) *string { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDistributedVirtualSwitchResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedVirtualSwitchResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDistributedVirtualSwitchResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributedVirtualSwitchResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of the uplinks on this vSphere distributed switch, as per the
// `uplinks` argument to the
// `DistributedVirtualSwitch`
// resource.
func (o LookupDistributedVirtualSwitchResultOutput) Uplinks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDistributedVirtualSwitchResult) []string { return v.Uplinks }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDistributedVirtualSwitchResultOutput{})
}
