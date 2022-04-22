// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getDatastore` data source can be used to discover the ID of a
// datastore in vSphere. This is useful to fetch the ID of a datastore that you
// want to use to create virtual machines in using the
// `VirtualMachine` resource.
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
// 			Name: pulumi.StringRef("dc1"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.GetDatastore(ctx, &GetDatastoreArgs{
// 			DatacenterId: pulumi.StringRef(datacenter.Id),
// 			Name:         "datastore1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDatastore(ctx *pulumi.Context, args *GetDatastoreArgs, opts ...pulumi.InvokeOption) (*GetDatastoreResult, error) {
	var rv GetDatastoreResult
	err := ctx.Invoke("vsphere:index/getDatastore:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatastore.
type GetDatastoreArgs struct {
	// The managed object reference
	// ID of the datacenter the datastore is located in. This
	// can be omitted if the search path used in `name` is an absolute path. For
	// default datacenters, use the id attribute from an empty `Datacenter`
	// data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// The name of the datastore. This can be a name or path.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDatastore.
type GetDatastoreResult struct {
	DatacenterId *string `pulumi:"datacenterId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func GetDatastoreOutput(ctx *pulumi.Context, args GetDatastoreOutputArgs, opts ...pulumi.InvokeOption) GetDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatastoreResult, error) {
			args := v.(GetDatastoreArgs)
			r, err := GetDatastore(ctx, &args, opts...)
			var s GetDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatastoreResultOutput)
}

// A collection of arguments for invoking getDatastore.
type GetDatastoreOutputArgs struct {
	// The managed object reference
	// ID of the datacenter the datastore is located in. This
	// can be omitted if the search path used in `name` is an absolute path. For
	// default datacenters, use the id attribute from an empty `Datacenter`
	// data source.
	DatacenterId pulumi.StringPtrInput `pulumi:"datacenterId"`
	// The name of the datastore. This can be a name or path.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatastoreArgs)(nil)).Elem()
}

// A collection of values returned by getDatastore.
type GetDatastoreResultOutput struct{ *pulumi.OutputState }

func (GetDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatastoreResult)(nil)).Elem()
}

func (o GetDatastoreResultOutput) ToGetDatastoreResultOutput() GetDatastoreResultOutput {
	return o
}

func (o GetDatastoreResultOutput) ToGetDatastoreResultOutputWithContext(ctx context.Context) GetDatastoreResultOutput {
	return o
}

func (o GetDatastoreResultOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatastoreResult) *string { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatastoreResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatastoreResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatastoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatastoreResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatastoreResultOutput{})
}
