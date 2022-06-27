// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ComputeCluster` data source can be used to discover the ID of a
// cluster in vSphere. This is useful to fetch the ID of a cluster that you want
// to use for virtual machine placement via the `VirtualMachine` resource, allowing to specify the cluster's root resource pool directly versus
// using the alias available through the `ResourcePool`
// data source.
//
// > You may also wish to see the `ComputeCluster`
//  resource for more information about clusters and how to managed the resource
//  in this provider.
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
// 		datacenter, err := vsphere.LookupDatacenter(ctx, &GetDatacenterArgs{
// 			Name: pulumi.StringRef("dc-01"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.LookupComputeCluster(ctx, &GetComputeClusterArgs{
// 			Name:         "cluster-01",
// 			DatacenterId: pulumi.StringRef(datacenter.Id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupComputeCluster(ctx *pulumi.Context, args *LookupComputeClusterArgs, opts ...pulumi.InvokeOption) (*LookupComputeClusterResult, error) {
	var rv LookupComputeClusterResult
	err := ctx.Invoke("vsphere:index/getComputeCluster:getComputeCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeCluster.
type LookupComputeClusterArgs struct {
	// The managed object reference ID
	// of the datacenter the cluster is located in.  This can be omitted if the
	// search path used in `name` is an absolute path. For default datacenters,
	// use the `id` attribute from an empty `Datacenter` data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name or absolute path to the cluster.
	Name string `pulumi:"name"`
}

// A collection of values returned by getComputeCluster.
type LookupComputeClusterResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	ResourcePoolId string `pulumi:"resourcePoolId"`
}

func LookupComputeClusterOutput(ctx *pulumi.Context, args LookupComputeClusterOutputArgs, opts ...pulumi.InvokeOption) LookupComputeClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComputeClusterResult, error) {
			args := v.(LookupComputeClusterArgs)
			r, err := LookupComputeCluster(ctx, &args, opts...)
			var s LookupComputeClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComputeClusterResultOutput)
}

// A collection of arguments for invoking getComputeCluster.
type LookupComputeClusterOutputArgs struct {
	// The managed object reference ID
	// of the datacenter the cluster is located in.  This can be omitted if the
	// search path used in `name` is an absolute path. For default datacenters,
	// use the `id` attribute from an empty `Datacenter` data source.
	DatacenterId pulumi.StringPtrInput `pulumi:"datacenterId"`
	// The name or absolute path to the cluster.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupComputeClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeClusterArgs)(nil)).Elem()
}

// A collection of values returned by getComputeCluster.
type LookupComputeClusterResultOutput struct{ *pulumi.OutputState }

func (LookupComputeClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeClusterResult)(nil)).Elem()
}

func (o LookupComputeClusterResultOutput) ToLookupComputeClusterResultOutput() LookupComputeClusterResultOutput {
	return o
}

func (o LookupComputeClusterResultOutput) ToLookupComputeClusterResultOutputWithContext(ctx context.Context) LookupComputeClusterResultOutput {
	return o
}

func (o LookupComputeClusterResultOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeClusterResult) *string { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupComputeClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeClusterResultOutput) ResourcePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeClusterResult) string { return v.ResourcePoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeClusterResultOutput{})
}
