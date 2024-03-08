// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `GuestOsCustomization` data source can be used to discover the details about a customization specification for a guest operating system.
//
// Suggested change
// > **NOTE:** The name attribute is the unique identifier for the customization specification per vCenter Server instance.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := vsphere.LookupGuestOsCustomization(ctx, &vsphere.LookupGuestOsCustomizationArgs{
//				Name: "linux-spec",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupGuestOsCustomization(ctx *pulumi.Context, args *LookupGuestOsCustomizationArgs, opts ...pulumi.InvokeOption) (*LookupGuestOsCustomizationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGuestOsCustomizationResult
	err := ctx.Invoke("vsphere:index/getGuestOsCustomization:getGuestOsCustomization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGuestOsCustomization.
type LookupGuestOsCustomizationArgs struct {
	// The name of the customization specification is the unique identifier per vCenter Server instance.
	Name string `pulumi:"name"`
}

// A collection of values returned by getGuestOsCustomization.
type LookupGuestOsCustomizationResult struct {
	// The number of last changed version to the customization specification.
	ChangeVersion string `pulumi:"changeVersion"`
	// The description for the customization specification.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The time of last modification to the customization specification.
	LastUpdateTime string `pulumi:"lastUpdateTime"`
	Name           string `pulumi:"name"`
	// Container object for the guest operating system properties to be customized. See virtual machine customizations
	Specs []GetGuestOsCustomizationSpec `pulumi:"specs"`
	// The type of customization specification: One among: Windows, Linux.
	Type string `pulumi:"type"`
}

func LookupGuestOsCustomizationOutput(ctx *pulumi.Context, args LookupGuestOsCustomizationOutputArgs, opts ...pulumi.InvokeOption) LookupGuestOsCustomizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGuestOsCustomizationResult, error) {
			args := v.(LookupGuestOsCustomizationArgs)
			r, err := LookupGuestOsCustomization(ctx, &args, opts...)
			var s LookupGuestOsCustomizationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGuestOsCustomizationResultOutput)
}

// A collection of arguments for invoking getGuestOsCustomization.
type LookupGuestOsCustomizationOutputArgs struct {
	// The name of the customization specification is the unique identifier per vCenter Server instance.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupGuestOsCustomizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestOsCustomizationArgs)(nil)).Elem()
}

// A collection of values returned by getGuestOsCustomization.
type LookupGuestOsCustomizationResultOutput struct{ *pulumi.OutputState }

func (LookupGuestOsCustomizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGuestOsCustomizationResult)(nil)).Elem()
}

func (o LookupGuestOsCustomizationResultOutput) ToLookupGuestOsCustomizationResultOutput() LookupGuestOsCustomizationResultOutput {
	return o
}

func (o LookupGuestOsCustomizationResultOutput) ToLookupGuestOsCustomizationResultOutputWithContext(ctx context.Context) LookupGuestOsCustomizationResultOutput {
	return o
}

// The number of last changed version to the customization specification.
func (o LookupGuestOsCustomizationResultOutput) ChangeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestOsCustomizationResult) string { return v.ChangeVersion }).(pulumi.StringOutput)
}

// The description for the customization specification.
func (o LookupGuestOsCustomizationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestOsCustomizationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGuestOsCustomizationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestOsCustomizationResult) string { return v.Id }).(pulumi.StringOutput)
}

// The time of last modification to the customization specification.
func (o LookupGuestOsCustomizationResultOutput) LastUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestOsCustomizationResult) string { return v.LastUpdateTime }).(pulumi.StringOutput)
}

func (o LookupGuestOsCustomizationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestOsCustomizationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Container object for the guest operating system properties to be customized. See virtual machine customizations
func (o LookupGuestOsCustomizationResultOutput) Specs() GetGuestOsCustomizationSpecArrayOutput {
	return o.ApplyT(func(v LookupGuestOsCustomizationResult) []GetGuestOsCustomizationSpec { return v.Specs }).(GetGuestOsCustomizationSpecArrayOutput)
}

// The type of customization specification: One among: Windows, Linux.
func (o LookupGuestOsCustomizationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGuestOsCustomizationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGuestOsCustomizationResultOutput{})
}
