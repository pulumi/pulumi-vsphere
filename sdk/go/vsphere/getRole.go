// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Role` data source can be used to discover the id and privileges associated
// with a role given its name or display label in vsphere UI.
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
// 		_, err := vsphere.LookupRole(ctx, &vsphere.LookupRoleArgs{
// 			Label: "Virtual machine user (sample)",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRole(ctx *pulumi.Context, args *LookupRoleArgs, opts ...pulumi.InvokeOption) (*LookupRoleResult, error) {
	var rv LookupRoleResult
	err := ctx.Invoke("vsphere:index/getRole:getRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRole.
type LookupRoleArgs struct {
	// The description of the role.
	Description *string `pulumi:"description"`
	// The label of the role.
	Label string  `pulumi:"label"`
	Name  *string `pulumi:"name"`
	// The privileges associated with the role.
	RolePrivileges []string `pulumi:"rolePrivileges"`
}

// A collection of values returned by getRole.
type LookupRoleResult struct {
	// The description of the role.
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The display label of the role.
	Label string  `pulumi:"label"`
	Name  *string `pulumi:"name"`
	// The privileges associated with the role.
	RolePrivileges []string `pulumi:"rolePrivileges"`
}
