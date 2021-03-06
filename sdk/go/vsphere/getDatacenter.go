// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Datacenter` data source can be used to discover the ID of a
// vSphere datacenter. This can then be used with resources or data sources that
// require a datacenter, such as the `Host`
// data source.
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
// 		opt0 := "dc1"
// 		_, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
// 			Name: &opt0,
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
	// Can be omitted if there is only one datacenter in your inventory.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDatacenter.
type LookupDatacenterResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
}
