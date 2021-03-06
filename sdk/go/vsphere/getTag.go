// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Tag` data source can be used to reference tags that are not
// managed by this provider. Its attributes are exactly the same as the `Tag`
// resource, and, like importing, the data source takes a name and
// category to search on. The `id` and other attributes are then populated with
// the data found by the search.
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
// 	"github.com/pulumi/pulumi-vsphere/sdk/v3/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		category, err := vsphere.LookupTagCategory(ctx, &vsphere.LookupTagCategoryArgs{
// 			Name: "test-category",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vsphere.LookupTag(ctx, &vsphere.LookupTagArgs{
// 			CategoryId: category.Id,
// 			Name:       "test-tag",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
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
	// The ID of the tag category the tag is located in.
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
