// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ContentLibrary` data source can be used to discover the ID of a content library.
//
// > **NOTE:** This resource requires vCenter Server and is not available on direct ESXi host connections.
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
//			_, err := vsphere.LookupContentLibrary(ctx, &GetContentLibraryArgs{
//				Name: "Content Library",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
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
	// The name of the content library.
	Name string `pulumi:"name"`
}

// A collection of values returned by getContentLibrary.
type LookupContentLibraryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupContentLibraryOutput(ctx *pulumi.Context, args LookupContentLibraryOutputArgs, opts ...pulumi.InvokeOption) LookupContentLibraryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentLibraryResult, error) {
			args := v.(LookupContentLibraryArgs)
			r, err := LookupContentLibrary(ctx, &args, opts...)
			var s LookupContentLibraryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentLibraryResultOutput)
}

// A collection of arguments for invoking getContentLibrary.
type LookupContentLibraryOutputArgs struct {
	// The name of the content library.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupContentLibraryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentLibraryArgs)(nil)).Elem()
}

// A collection of values returned by getContentLibrary.
type LookupContentLibraryResultOutput struct{ *pulumi.OutputState }

func (LookupContentLibraryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentLibraryResult)(nil)).Elem()
}

func (o LookupContentLibraryResultOutput) ToLookupContentLibraryResultOutput() LookupContentLibraryResultOutput {
	return o
}

func (o LookupContentLibraryResultOutput) ToLookupContentLibraryResultOutputWithContext(ctx context.Context) LookupContentLibraryResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContentLibraryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentLibraryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContentLibraryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentLibraryResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentLibraryResultOutput{})
}
