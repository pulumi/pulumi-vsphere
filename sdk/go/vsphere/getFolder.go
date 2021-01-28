// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The `Folder` data source can be used to get the general attributes of a
// vSphere inventory folder. Paths are absolute and include must include the
// datacenter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere"
// 	"github.com/pulumi/pulumi-vsphere/sdk/v2/go/vsphere/"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vsphere.LookupFolder(ctx, &vsphere.LookupFolderArgs{
// 			Path: "/dc1/datastore/folder1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	var rv LookupFolderResult
	err := ctx.Invoke("vsphere:index/getFolder:getFolder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFolder.
type LookupFolderArgs struct {
	// The absolute path of the folder. For example, given a
	// default datacenter of `default-dc`, a folder of type `vm`, and a folder name
	// of `test-folder`, the resulting path would be
	// `/default-dc/vm/test-folder`. The valid folder types to be used in
	// the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
	Path string `pulumi:"path"`
}

// A collection of values returned by getFolder.
type LookupFolderResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Path string `pulumi:"path"`
}
