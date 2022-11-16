// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Tag` data source can be used to reference tags that are not
// managed by this provider. Its attributes are exactly the same as the `Tag`
// resource, and, like importing, the data source takes a name and
// category to search on. The `id` and other attributes are then populated with
// the data found by the search.
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
//			category, err := vsphere.LookupTagCategory(ctx, &GetTagCategoryArgs{
//				Name: "example-category",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.LookupTag(ctx, &GetTagArgs{
//				Name:       "example-tag",
//				CategoryId: category.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupTag(ctx *pulumi.Context, args *LookupTagArgs, opts ...pulumi.InvokeOption) (*LookupTagResult, error) {
	var rv LookupTagResult
	err := ctx.Invoke("vsphere:index/getTag:getTag", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTag.
type LookupTagArgs struct {
	// The ID of the tag category in which the tag is
	// located.
	CategoryId string `pulumi:"categoryId"`
	// The name of the tag.
	Name string `pulumi:"name"`
}

// A collection of values returned by getTag.
type LookupTagResult struct {
	CategoryId  string `pulumi:"categoryId"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupTagOutput(ctx *pulumi.Context, args LookupTagOutputArgs, opts ...pulumi.InvokeOption) LookupTagResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagResult, error) {
			args := v.(LookupTagArgs)
			r, err := LookupTag(ctx, &args, opts...)
			var s LookupTagResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagResultOutput)
}

// A collection of arguments for invoking getTag.
type LookupTagOutputArgs struct {
	// The ID of the tag category in which the tag is
	// located.
	CategoryId pulumi.StringInput `pulumi:"categoryId"`
	// The name of the tag.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupTagOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagArgs)(nil)).Elem()
}

// A collection of values returned by getTag.
type LookupTagResultOutput struct{ *pulumi.OutputState }

func (LookupTagResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagResult)(nil)).Elem()
}

func (o LookupTagResultOutput) ToLookupTagResultOutput() LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) ToLookupTagResultOutputWithContext(ctx context.Context) LookupTagResultOutput {
	return o
}

func (o LookupTagResultOutput) CategoryId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.CategoryId }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupTagResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagResultOutput{})
}
