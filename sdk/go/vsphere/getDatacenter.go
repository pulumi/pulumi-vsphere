// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Datacenter` data source can be used to discover the ID of a
// vSphere datacenter object. This can then be used with resources or data sources
// that require a datacenter, such as the `Host`
// data source.
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
// 		_, err := vsphere.LookupDatacenter(ctx, &GetDatacenterArgs{
// 			Name: pulumi.StringRef("dc-01"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDatacenter(ctx *pulumi.Context, args *LookupDatacenterArgs, opts ...pulumi.InvokeOption) (*LookupDatacenterResult, error) {
	var rv LookupDatacenterResult
	err := ctx.Invoke("vsphere:index/getDatacenter:getDatacenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatacenter.
type LookupDatacenterArgs struct {
	// The name of the datacenter. This can be a name or path.
	// Can be omitted if there is only one datacenter in the inventory.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDatacenter.
type LookupDatacenterResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
}

func LookupDatacenterOutput(ctx *pulumi.Context, args LookupDatacenterOutputArgs, opts ...pulumi.InvokeOption) LookupDatacenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatacenterResult, error) {
			args := v.(LookupDatacenterArgs)
			r, err := LookupDatacenter(ctx, &args, opts...)
			var s LookupDatacenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatacenterResultOutput)
}

// A collection of arguments for invoking getDatacenter.
type LookupDatacenterOutputArgs struct {
	// The name of the datacenter. This can be a name or path.
	// Can be omitted if there is only one datacenter in the inventory.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupDatacenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatacenterArgs)(nil)).Elem()
}

// A collection of values returned by getDatacenter.
type LookupDatacenterResultOutput struct{ *pulumi.OutputState }

func (LookupDatacenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatacenterResult)(nil)).Elem()
}

func (o LookupDatacenterResultOutput) ToLookupDatacenterResultOutput() LookupDatacenterResultOutput {
	return o
}

func (o LookupDatacenterResultOutput) ToLookupDatacenterResultOutputWithContext(ctx context.Context) LookupDatacenterResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDatacenterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDatacenterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatacenterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatacenterResultOutput{})
}
