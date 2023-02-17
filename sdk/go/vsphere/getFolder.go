// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `Folder` data source can be used to get the general attributes of a
// vSphere inventory folder. Paths are absolute and must include the datacenter.
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
//			_, err := vsphere.LookupFolder(ctx, &vsphere.LookupFolderArgs{
//				Path: "/dc-01/datastore-01/folder-01",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
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

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFolderResult, error) {
			args := v.(LookupFolderArgs)
			r, err := LookupFolder(ctx, &args, opts...)
			var s LookupFolderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFolderResultOutput)
}

// A collection of arguments for invoking getFolder.
type LookupFolderOutputArgs struct {
	// The absolute path of the folder. For example, given a
	// default datacenter of `default-dc`, a folder of type `vm`, and a folder name
	// of `test-folder`, the resulting path would be
	// `/default-dc/vm/test-folder`. The valid folder types to be used in
	// the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
	Path pulumi.StringInput `pulumi:"path"`
}

func (LookupFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderArgs)(nil)).Elem()
}

// A collection of values returned by getFolder.
type LookupFolderResultOutput struct{ *pulumi.OutputState }

func (LookupFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderResult)(nil)).Elem()
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutput() LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutputWithContext(ctx context.Context) LookupFolderResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFolderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Path }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderResultOutput{})
}
