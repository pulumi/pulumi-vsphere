// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `License` data source can be used to get the general attributes of
// a license keys from a vCenter Server instance.
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
//			_, err := vsphere.LookupLicense(ctx, &vsphere.LookupLicenseArgs{
//				LicenseKey: "00000-00000-00000-00000-00000",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLicense(ctx *pulumi.Context, args *LookupLicenseArgs, opts ...pulumi.InvokeOption) (*LookupLicenseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLicenseResult
	err := ctx.Invoke("vsphere:index/getLicense:getLicense", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLicense.
type LookupLicenseArgs struct {
	// The license key.
	LicenseKey string `pulumi:"licenseKey"`
}

// A collection of values returned by getLicense.
type LookupLicenseResult struct {
	// The product edition of the license key.
	EditionKey string `pulumi:"editionKey"`
	Id         string `pulumi:"id"`
	// A map of key/value pairs attached as labels (tags) to the license
	// key.
	Labels     map[string]string `pulumi:"labels"`
	LicenseKey string            `pulumi:"licenseKey"`
	// The display name for the license.
	Name string `pulumi:"name"`
	// Total number of units (example: CPUs) contained in the license.
	Total int `pulumi:"total"`
	// The number of units (example: CPUs) assigned to this license.
	Used int `pulumi:"used"`
}

func LookupLicenseOutput(ctx *pulumi.Context, args LookupLicenseOutputArgs, opts ...pulumi.InvokeOption) LookupLicenseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLicenseResultOutput, error) {
			args := v.(LookupLicenseArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupLicenseResult
			secret, err := ctx.InvokePackageRaw("vsphere:index/getLicense:getLicense", args, &rv, "", opts...)
			if err != nil {
				return LookupLicenseResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupLicenseResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupLicenseResultOutput), nil
			}
			return output, nil
		}).(LookupLicenseResultOutput)
}

// A collection of arguments for invoking getLicense.
type LookupLicenseOutputArgs struct {
	// The license key.
	LicenseKey pulumi.StringInput `pulumi:"licenseKey"`
}

func (LookupLicenseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLicenseArgs)(nil)).Elem()
}

// A collection of values returned by getLicense.
type LookupLicenseResultOutput struct{ *pulumi.OutputState }

func (LookupLicenseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLicenseResult)(nil)).Elem()
}

func (o LookupLicenseResultOutput) ToLookupLicenseResultOutput() LookupLicenseResultOutput {
	return o
}

func (o LookupLicenseResultOutput) ToLookupLicenseResultOutputWithContext(ctx context.Context) LookupLicenseResultOutput {
	return o
}

// The product edition of the license key.
func (o LookupLicenseResultOutput) EditionKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseResult) string { return v.EditionKey }).(pulumi.StringOutput)
}

func (o LookupLicenseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseResult) string { return v.Id }).(pulumi.StringOutput)
}

// A map of key/value pairs attached as labels (tags) to the license
// key.
func (o LookupLicenseResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLicenseResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupLicenseResultOutput) LicenseKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseResult) string { return v.LicenseKey }).(pulumi.StringOutput)
}

// The display name for the license.
func (o LookupLicenseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseResult) string { return v.Name }).(pulumi.StringOutput)
}

// Total number of units (example: CPUs) contained in the license.
func (o LookupLicenseResultOutput) Total() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLicenseResult) int { return v.Total }).(pulumi.IntOutput)
}

// The number of units (example: CPUs) assigned to this license.
func (o LookupLicenseResultOutput) Used() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLicenseResult) int { return v.Used }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLicenseResultOutput{})
}
