// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `ContentLibrary` data source can be used to discover the ID of a Content Library.
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
func LookupContentLibrary(ctx *pulumi.Context, args *LookupContentLibraryArgs, opts ...pulumi.InvokeOption) (*LookupContentLibraryResult, error) {
	var rv LookupContentLibraryResult
	err := ctx.Invoke("vsphere:index/getContentLibrary:getContentLibrary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContentLibrary.
type LookupContentLibraryArgs struct {
	// The name of the Content Library.
	Name string `pulumi:"name"`
}

// A collection of values returned by getContentLibrary.
type LookupContentLibraryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
