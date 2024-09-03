// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Host` data source can be used to discover the ID of an ESXi host.
// This can then be used with resources or data sources that require an ESX
// host's managed object reference ID.
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
//			_, err = vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
//				Name:         pulumi.StringRef("esxi-01.example.com"),
//				DatacenterId: datacenter.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupHost(ctx *pulumi.Context, args *LookupHostArgs, opts ...pulumi.InvokeOption) (*LookupHostResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupHostResult
	err := ctx.Invoke("vsphere:index/getHost:getHost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHost.
type LookupHostArgs struct {
	// The managed object reference ID
	// of a vSphere datacenter object.
	DatacenterId string `pulumi:"datacenterId"`
	// The name of the ESXI host. This can be a name or path.
	// Can be omitted if there is only one host in your inventory.
	//
	// > **NOTE:** When used against an ESXi host directly, this data source *always*
	// returns the ESXi host's object ID, regardless of what is entered into `name`.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getHost.
type LookupHostResult struct {
	DatacenterId string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// The managed object ID of the ESXi
	// host's root resource pool.
	ResourcePoolId string `pulumi:"resourcePoolId"`
}

func LookupHostOutput(ctx *pulumi.Context, args LookupHostOutputArgs, opts ...pulumi.InvokeOption) LookupHostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostResult, error) {
			args := v.(LookupHostArgs)
			r, err := LookupHost(ctx, &args, opts...)
			var s LookupHostResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostResultOutput)
}

// A collection of arguments for invoking getHost.
type LookupHostOutputArgs struct {
	// The managed object reference ID
	// of a vSphere datacenter object.
	DatacenterId pulumi.StringInput `pulumi:"datacenterId"`
	// The name of the ESXI host. This can be a name or path.
	// Can be omitted if there is only one host in your inventory.
	//
	// > **NOTE:** When used against an ESXi host directly, this data source *always*
	// returns the ESXi host's object ID, regardless of what is entered into `name`.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupHostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostArgs)(nil)).Elem()
}

// A collection of values returned by getHost.
type LookupHostResultOutput struct{ *pulumi.OutputState }

func (LookupHostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostResult)(nil)).Elem()
}

func (o LookupHostResultOutput) ToLookupHostResultOutput() LookupHostResultOutput {
	return o
}

func (o LookupHostResultOutput) ToLookupHostResultOutputWithContext(ctx context.Context) LookupHostResultOutput {
	return o
}

func (o LookupHostResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupHostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The managed object ID of the ESXi
// host's root resource pool.
func (o LookupHostResultOutput) ResourcePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostResult) string { return v.ResourcePoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostResultOutput{})
}
