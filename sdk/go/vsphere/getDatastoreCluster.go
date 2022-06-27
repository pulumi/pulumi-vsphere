// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `DatastoreCluster` data source can be used to discover the ID of a
// vSphere datastore cluster object. This can then be used with resources or data sources
// that require a datastore. For example, to assign datastores using the
// `NasDatastore` or `VmfsDatastore` resources, or to create virtual machines in using the `VirtualMachine` resource.
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
// 		_, err = vsphere.LookupDatastoreCluster(ctx, &GetDatastoreClusterArgs{
// 			Name:         "datastore-cluster-01",
// 			DatacenterId: pulumi.StringRef(datacenter.Id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDatastoreCluster(ctx *pulumi.Context, args *LookupDatastoreClusterArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreClusterResult, error) {
	var rv LookupDatastoreClusterResult
	err := ctx.Invoke("vsphere:index/getDatastoreCluster:getDatastoreCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatastoreCluster.
type LookupDatastoreClusterArgs struct {
	// The managed object reference
	// ID of the datacenter the datastore cluster is located in.
	// This can be omitted if the search path used in `name` is an absolute path.
	// For default datacenters, use the id attribute from an empty
	// `Datacenter` data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name or absolute path to the datastore cluster.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDatastoreCluster.
type LookupDatastoreClusterResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupDatastoreClusterOutput(ctx *pulumi.Context, args LookupDatastoreClusterOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreClusterResult, error) {
			args := v.(LookupDatastoreClusterArgs)
			r, err := LookupDatastoreCluster(ctx, &args, opts...)
			var s LookupDatastoreClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreClusterResultOutput)
}

// A collection of arguments for invoking getDatastoreCluster.
type LookupDatastoreClusterOutputArgs struct {
	// The managed object reference
	// ID of the datacenter the datastore cluster is located in.
	// This can be omitted if the search path used in `name` is an absolute path.
	// For default datacenters, use the id attribute from an empty
	// `Datacenter` data source.
	DatacenterId pulumi.StringPtrInput `pulumi:"datacenterId"`
	// The name or absolute path to the datastore cluster.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDatastoreClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreClusterArgs)(nil)).Elem()
}

// A collection of values returned by getDatastoreCluster.
type LookupDatastoreClusterResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreClusterResult)(nil)).Elem()
}

func (o LookupDatastoreClusterResultOutput) ToLookupDatastoreClusterResultOutput() LookupDatastoreClusterResultOutput {
	return o
}

func (o LookupDatastoreClusterResultOutput) ToLookupDatastoreClusterResultOutputWithContext(ctx context.Context) LookupDatastoreClusterResultOutput {
	return o
}

func (o LookupDatastoreClusterResultOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreClusterResult) *string { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatastoreClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatastoreClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatastoreClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreClusterResultOutput{})
}
