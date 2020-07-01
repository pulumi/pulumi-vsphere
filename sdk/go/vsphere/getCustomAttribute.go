// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `CustomAttribute` data source can be used to reference custom
// attributes that are not managed by this provider. Its attributes are exactly the
// same as the `CustomAttribute` resource,
// and, like importing, the data source takes a name to search on. The `id` and
// other attributes are then populated with the data found by the search.
//
// > **NOTE:** Custom attributes are unsupported on direct ESXi connections
// and require vCenter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vsphere.LookupCustomAttribute(ctx, &vsphere.LookupCustomAttributeArgs{
// 			Name: "test-attribute",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCustomAttribute(ctx *pulumi.Context, args *LookupCustomAttributeArgs, opts ...pulumi.InvokeOption) (*LookupCustomAttributeResult, error) {
	var rv LookupCustomAttributeResult
	err := ctx.Invoke("vsphere:index/getCustomAttribute:getCustomAttribute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomAttribute.
type LookupCustomAttributeArgs struct {
	// The name of the custom attribute.
	Name string `pulumi:"name"`
}

// A collection of values returned by getCustomAttribute.
type LookupCustomAttributeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	ManagedObjectType string `pulumi:"managedObjectType"`
	Name              string `pulumi:"name"`
}
