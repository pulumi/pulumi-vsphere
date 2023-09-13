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

// The `TagCategory` data source can be used to reference tag categories
// that are not managed by this provider. Its attributes are the same as the
// `TagCategory` resource, and, like importing,
// the data source uses a name and category as search criteria. The `id` and other
// attributes are populated with the data found by the search.
//
// > **NOTE:** Tagging is not supported on direct ESXi hosts connections and
// requires vCenter Server.
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
//			_, err := vsphere.LookupTagCategory(ctx, &vsphere.LookupTagCategoryArgs{
//				Name: "example-category",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTagCategory(ctx *pulumi.Context, args *LookupTagCategoryArgs, opts ...pulumi.InvokeOption) (*LookupTagCategoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTagCategoryResult
	err := ctx.Invoke("vsphere:index/getTagCategory:getTagCategory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagCategory.
type LookupTagCategoryArgs struct {
	// The name of the tag category.
	Name string `pulumi:"name"`
}

// A collection of values returned by getTagCategory.
type LookupTagCategoryResult struct {
	AssociableTypes []string `pulumi:"associableTypes"`
	Cardinality     string   `pulumi:"cardinality"`
	Description     string   `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupTagCategoryOutput(ctx *pulumi.Context, args LookupTagCategoryOutputArgs, opts ...pulumi.InvokeOption) LookupTagCategoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagCategoryResult, error) {
			args := v.(LookupTagCategoryArgs)
			r, err := LookupTagCategory(ctx, &args, opts...)
			var s LookupTagCategoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagCategoryResultOutput)
}

// A collection of arguments for invoking getTagCategory.
type LookupTagCategoryOutputArgs struct {
	// The name of the tag category.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupTagCategoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagCategoryArgs)(nil)).Elem()
}

// A collection of values returned by getTagCategory.
type LookupTagCategoryResultOutput struct{ *pulumi.OutputState }

func (LookupTagCategoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagCategoryResult)(nil)).Elem()
}

func (o LookupTagCategoryResultOutput) ToLookupTagCategoryResultOutput() LookupTagCategoryResultOutput {
	return o
}

func (o LookupTagCategoryResultOutput) ToLookupTagCategoryResultOutputWithContext(ctx context.Context) LookupTagCategoryResultOutput {
	return o
}

func (o LookupTagCategoryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTagCategoryResult] {
	return pulumix.Output[LookupTagCategoryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupTagCategoryResultOutput) AssociableTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTagCategoryResult) []string { return v.AssociableTypes }).(pulumi.StringArrayOutput)
}

func (o LookupTagCategoryResultOutput) Cardinality() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagCategoryResult) string { return v.Cardinality }).(pulumi.StringOutput)
}

func (o LookupTagCategoryResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagCategoryResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTagCategoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagCategoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagCategoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagCategoryResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagCategoryResultOutput{})
}
