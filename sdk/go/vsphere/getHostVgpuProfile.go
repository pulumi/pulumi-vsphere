// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getHostVgpuProfile` data source can be used to discover the
// available vGPU profiles of a vSphere host.
//
// ## Example Usage
//
// ### To Return All VGPU Profiles
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
//			datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
//				Name: pulumi.StringRef("dc-01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			host, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
//				Name:         pulumi.StringRef("esxi-01.example.com"),
//				DatacenterId: datacenter.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.GetHostVgpuProfile(ctx, &vsphere.GetHostVgpuProfileArgs{
//				HostId: host.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### With VGPU Profile Name_regex
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
//			datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
//				Name: pulumi.StringRef("dc-01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			host, err := vsphere.LookupHost(ctx, &vsphere.LookupHostArgs{
//				Name:         pulumi.StringRef("esxi-01.example.com"),
//				DatacenterId: datacenter.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.GetHostVgpuProfile(ctx, &vsphere.GetHostVgpuProfileArgs{
//				HostId:    host.Id,
//				NameRegex: pulumi.StringRef("a100"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetHostVgpuProfile(ctx *pulumi.Context, args *GetHostVgpuProfileArgs, opts ...pulumi.InvokeOption) (*GetHostVgpuProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetHostVgpuProfileResult
	err := ctx.Invoke("vsphere:index/getHostVgpuProfile:getHostVgpuProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostVgpuProfile.
type GetHostVgpuProfileArgs struct {
	// The [managed object reference ID][docs-about-morefs] of
	// a host.
	HostId string `pulumi:"hostId"`
	// A regular expression that will be used to match the
	// host vGPU profile name.
	//
	// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
	NameRegex *string `pulumi:"nameRegex"`
}

// A collection of values returned by getHostVgpuProfile.
type GetHostVgpuProfileResult struct {
	// The [managed objectID][docs-about-morefs] of the ESXi host.
	HostId string `pulumi:"hostId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Optional) A regular expression that will be used to match the
	// host vGPU profile name.
	NameRegex *string `pulumi:"nameRegex"`
	// The list of available vGPU profiles on the ESXi host.
	// This may be and empty array if no vGPU profile are identified.
	VgpuProfiles []GetHostVgpuProfileVgpuProfile `pulumi:"vgpuProfiles"`
}

func GetHostVgpuProfileOutput(ctx *pulumi.Context, args GetHostVgpuProfileOutputArgs, opts ...pulumi.InvokeOption) GetHostVgpuProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetHostVgpuProfileResult, error) {
			args := v.(GetHostVgpuProfileArgs)
			r, err := GetHostVgpuProfile(ctx, &args, opts...)
			var s GetHostVgpuProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetHostVgpuProfileResultOutput)
}

// A collection of arguments for invoking getHostVgpuProfile.
type GetHostVgpuProfileOutputArgs struct {
	// The [managed object reference ID][docs-about-morefs] of
	// a host.
	HostId pulumi.StringInput `pulumi:"hostId"`
	// A regular expression that will be used to match the
	// host vGPU profile name.
	//
	// [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
}

func (GetHostVgpuProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostVgpuProfileArgs)(nil)).Elem()
}

// A collection of values returned by getHostVgpuProfile.
type GetHostVgpuProfileResultOutput struct{ *pulumi.OutputState }

func (GetHostVgpuProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetHostVgpuProfileResult)(nil)).Elem()
}

func (o GetHostVgpuProfileResultOutput) ToGetHostVgpuProfileResultOutput() GetHostVgpuProfileResultOutput {
	return o
}

func (o GetHostVgpuProfileResultOutput) ToGetHostVgpuProfileResultOutputWithContext(ctx context.Context) GetHostVgpuProfileResultOutput {
	return o
}

// The [managed objectID][docs-about-morefs] of the ESXi host.
func (o GetHostVgpuProfileResultOutput) HostId() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostVgpuProfileResult) string { return v.HostId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetHostVgpuProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetHostVgpuProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// (Optional) A regular expression that will be used to match the
// host vGPU profile name.
func (o GetHostVgpuProfileResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetHostVgpuProfileResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The list of available vGPU profiles on the ESXi host.
// This may be and empty array if no vGPU profile are identified.
func (o GetHostVgpuProfileResultOutput) VgpuProfiles() GetHostVgpuProfileVgpuProfileArrayOutput {
	return o.ApplyT(func(v GetHostVgpuProfileResult) []GetHostVgpuProfileVgpuProfile { return v.VgpuProfiles }).(GetHostVgpuProfileVgpuProfileArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetHostVgpuProfileResultOutput{})
}
