// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `TagCategory` data source can be used to reference tag categories
// that are not managed by this provider. Its attributes are exactly the same as the
// `TagCategory` resource, and, like importing,
// the data source takes a name to search on. The `id` and other attributes are
// then populated with the data found by the search.
//
// > **NOTE:** Tagging support is unsupported on direct ESXi connections and
// requires vCenter 6.0 or higher.
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
// 		_, err := vsphere.LookupTagCategory(ctx, &vsphere.LookupTagCategoryArgs{
// 			Name: "test-category",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupTagCategory(ctx *pulumi.Context, args *LookupTagCategoryArgs, opts ...pulumi.InvokeOption) (*LookupTagCategoryResult, error) {
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
