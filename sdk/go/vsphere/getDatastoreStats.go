// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getDatastoreStats` data source can be used to retrieve the usage
// stats of all vSphere datastore objects in a datacenter. This can then be used as
// a standalone data source to get information required as input to other data
// sources.
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
//			_, err = vsphere.GetDatastoreStats(ctx, &vsphere.GetDatastoreStatsArgs{
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
//
// A useful example of this data source would be to determine the datastore with
// the most free space. For example, in addition to the above:
//
// Create an `outputs.tf` like that:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			ctx.Export("maxFreeSpaceName", theirMaxFreeSpaceName)
//			ctx.Export("maxFreeSpace", theirMaxFreeSpace)
//			return nil
//		})
//	}
//
// ```
//
// and a `locals.tf` like that:
func GetDatastoreStats(ctx *pulumi.Context, args *GetDatastoreStatsArgs, opts ...pulumi.InvokeOption) (*GetDatastoreStatsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatastoreStatsResult
	err := ctx.Invoke("vsphere:index/getDatastoreStats:getDatastoreStats", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatastoreStats.
type GetDatastoreStatsArgs struct {
	// A mapping of the capacity for all datastore in the datacenter,
	// where the name of the datastore is used as key and the capacity as value.
	Capacity map[string]string `pulumi:"capacity"`
	// The
	// [managed object reference ID][docs-about-morefs] of the datacenter the
	// datastores are located in. For default datacenters, use the `id` attribute
	// from an empty `Datacenter` data source.
	//
	// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
	DatacenterId string `pulumi:"datacenterId"`
	// A mapping of the free space for each datastore in the
	// datacenter, where the name of the datastore is used as key and the free space
	// as value.
	FreeSpace map[string]string `pulumi:"freeSpace"`
}

// A collection of values returned by getDatastoreStats.
type GetDatastoreStatsResult struct {
	// A mapping of the capacity for all datastore in the datacenter,
	// where the name of the datastore is used as key and the capacity as value.
	Capacity map[string]string `pulumi:"capacity"`
	// The [managed object reference ID][docs-about-morefs] of the
	// datacenter the datastores are located in.
	DatacenterId string `pulumi:"datacenterId"`
	// A mapping of the free space for each datastore in the
	// datacenter, where the name of the datastore is used as key and the free space
	// as value.
	FreeSpace map[string]string `pulumi:"freeSpace"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetDatastoreStatsOutput(ctx *pulumi.Context, args GetDatastoreStatsOutputArgs, opts ...pulumi.InvokeOption) GetDatastoreStatsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatastoreStatsResult, error) {
			args := v.(GetDatastoreStatsArgs)
			r, err := GetDatastoreStats(ctx, &args, opts...)
			var s GetDatastoreStatsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatastoreStatsResultOutput)
}

// A collection of arguments for invoking getDatastoreStats.
type GetDatastoreStatsOutputArgs struct {
	// A mapping of the capacity for all datastore in the datacenter,
	// where the name of the datastore is used as key and the capacity as value.
	Capacity pulumi.StringMapInput `pulumi:"capacity"`
	// The
	// [managed object reference ID][docs-about-morefs] of the datacenter the
	// datastores are located in. For default datacenters, use the `id` attribute
	// from an empty `Datacenter` data source.
	//
	// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
	DatacenterId pulumi.StringInput `pulumi:"datacenterId"`
	// A mapping of the free space for each datastore in the
	// datacenter, where the name of the datastore is used as key and the free space
	// as value.
	FreeSpace pulumi.StringMapInput `pulumi:"freeSpace"`
}

func (GetDatastoreStatsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatastoreStatsArgs)(nil)).Elem()
}

// A collection of values returned by getDatastoreStats.
type GetDatastoreStatsResultOutput struct{ *pulumi.OutputState }

func (GetDatastoreStatsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatastoreStatsResult)(nil)).Elem()
}

func (o GetDatastoreStatsResultOutput) ToGetDatastoreStatsResultOutput() GetDatastoreStatsResultOutput {
	return o
}

func (o GetDatastoreStatsResultOutput) ToGetDatastoreStatsResultOutputWithContext(ctx context.Context) GetDatastoreStatsResultOutput {
	return o
}

// A mapping of the capacity for all datastore in the datacenter,
// where the name of the datastore is used as key and the capacity as value.
func (o GetDatastoreStatsResultOutput) Capacity() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDatastoreStatsResult) map[string]string { return v.Capacity }).(pulumi.StringMapOutput)
}

// The [managed object reference ID][docs-about-morefs] of the
// datacenter the datastores are located in.
func (o GetDatastoreStatsResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatastoreStatsResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// A mapping of the free space for each datastore in the
// datacenter, where the name of the datastore is used as key and the free space
// as value.
func (o GetDatastoreStatsResultOutput) FreeSpace() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDatastoreStatsResult) map[string]string { return v.FreeSpace }).(pulumi.StringMapOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatastoreStatsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatastoreStatsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatastoreStatsResultOutput{})
}
