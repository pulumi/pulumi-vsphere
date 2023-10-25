// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The `ResourcePool` data source can be used to discover the ID of a
// resource pool in vSphere. This is useful to return the ID of a resource pool
// that you want to use to create virtual machines in using the
// `VirtualMachine` resource.
func LookupResourcePool(ctx *pulumi.Context, args *LookupResourcePoolArgs, opts ...pulumi.InvokeOption) (*LookupResourcePoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourcePoolResult
	err := ctx.Invoke("vsphere:index/getResourcePool:getResourcePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourcePool.
type LookupResourcePoolArgs struct {
	// The managed object reference ID
	// of the datacenter in which the resource pool is located. This can be omitted
	// if the search path used in `name` is an absolute path. For default
	// datacenters, use the id attribute from an empty `Datacenter` data
	// source.
	//
	// > **Note:** When using ESXi without a vCenter Server instance, you do not
	// need to specify either attribute to use this data source. An empty declaration
	// will load the ESXi host's root resource pool.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name of the resource pool. This can be a name or
	// path. This is required when using vCenter.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getResourcePool.
type LookupResourcePoolResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
}

func LookupResourcePoolOutput(ctx *pulumi.Context, args LookupResourcePoolOutputArgs, opts ...pulumi.InvokeOption) LookupResourcePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourcePoolResult, error) {
			args := v.(LookupResourcePoolArgs)
			r, err := LookupResourcePool(ctx, &args, opts...)
			var s LookupResourcePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourcePoolResultOutput)
}

// A collection of arguments for invoking getResourcePool.
type LookupResourcePoolOutputArgs struct {
	// The managed object reference ID
	// of the datacenter in which the resource pool is located. This can be omitted
	// if the search path used in `name` is an absolute path. For default
	// datacenters, use the id attribute from an empty `Datacenter` data
	// source.
	//
	// > **Note:** When using ESXi without a vCenter Server instance, you do not
	// need to specify either attribute to use this data source. An empty declaration
	// will load the ESXi host's root resource pool.
	DatacenterId pulumi.StringPtrInput `pulumi:"datacenterId"`
	// The name of the resource pool. This can be a name or
	// path. This is required when using vCenter.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupResourcePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePoolArgs)(nil)).Elem()
}

// A collection of values returned by getResourcePool.
type LookupResourcePoolResultOutput struct{ *pulumi.OutputState }

func (LookupResourcePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePoolResult)(nil)).Elem()
}

func (o LookupResourcePoolResultOutput) ToLookupResourcePoolResultOutput() LookupResourcePoolResultOutput {
	return o
}

func (o LookupResourcePoolResultOutput) ToLookupResourcePoolResultOutputWithContext(ctx context.Context) LookupResourcePoolResultOutput {
	return o
}

func (o LookupResourcePoolResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupResourcePoolResult] {
	return pulumix.Output[LookupResourcePoolResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupResourcePoolResultOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupResourcePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourcePoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourcePoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourcePoolResultOutput{})
}
