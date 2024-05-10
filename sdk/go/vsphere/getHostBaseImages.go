// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getHostBaseImages` data source can be used to get the list of ESXi base images available
// for cluster software management.
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
//			_, err := vsphere.GetHostBaseImages(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetHostBaseImages(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetHostBaseImagesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHostBaseImagesResult
	err := ctx.Invoke("vsphere:index/getHostBaseImages:getHostBaseImages", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getHostBaseImages.
type GetHostBaseImagesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ESXi version identifier for the image
	Versions []string `pulumi:"versions"`
}

func GetHostBaseImagesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetHostBaseImagesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetHostBaseImagesResult, error) {
		r, err := GetHostBaseImages(ctx, opts...)
		var s GetHostBaseImagesResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetHostBaseImagesResultOutput)
}

// A collection of values returned by getHostBaseImages.
type GetHostBaseImagesResultOutput struct{ *pulumi.OutputState }

func (GetHostBaseImagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostBaseImagesResult)(nil)).Elem()
}

func (o GetHostBaseImagesResultOutput) ToGetHostBaseImagesResultOutput() GetHostBaseImagesResultOutput {
	return o
}

func (o GetHostBaseImagesResultOutput) ToGetHostBaseImagesResultOutputWithContext(ctx context.Context) GetHostBaseImagesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetHostBaseImagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostBaseImagesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ESXi version identifier for the image
func (o GetHostBaseImagesResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetHostBaseImagesResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHostBaseImagesResultOutput{})
}
