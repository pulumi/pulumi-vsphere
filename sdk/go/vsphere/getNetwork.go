// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getNetwork` data source can be used to discover the ID of a network in
// vSphere. This can be any network that can be used as the backing for a network
// interface for `VirtualMachine` or any other vSphere resource that
// requires a network. This includes standard (host-based) port groups, distributed
// port groups, or opaque networks such as those managed by NSX.
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
//			datacenter, err := vsphere.LookupDatacenter(ctx, &vsphere.LookupDatacenterArgs{
//				Name: pulumi.StringRef("dc-01"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
//				Name:         "VM Network",
//				DatacenterId: pulumi.StringRef(datacenter.Id),
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
// ### Additional Examples
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
//			_, err = vsphere.GetNetwork(ctx, &vsphere.GetNetworkArgs{
//				DatacenterId: pulumi.StringRef(datacenter.Id),
//				Name:         "VM Network",
//				Filters: []vsphere.GetNetworkFilter{
//					{
//						NetworkType: pulumi.StringRef("Network"),
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNetwork(ctx *pulumi.Context, args *GetNetworkArgs, opts ...pulumi.InvokeOption) (*GetNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNetworkResult
	err := ctx.Invoke("vsphere:index/getNetwork:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetwork.
type GetNetworkArgs struct {
	// The managed object reference ID
	// of the datacenter the network is located in. This can be omitted if the
	// search path used in `name` is an absolute path. For default datacenters,
	// use the `id` attribute from an empty `Datacenter` data source.
	DatacenterId *string `pulumi:"datacenterId"`
	// For distributed port group type
	// network objects, the ID of the distributed virtual switch for which the port
	// group belongs. It is useful to differentiate port groups with same name using
	// the distributed virtual switch ID.
	DistributedVirtualSwitchUuid *string `pulumi:"distributedVirtualSwitchUuid"`
	// Apply a filter for the discovered network.
	Filters []GetNetworkFilter `pulumi:"filters"`
	// The name of the network. This can be a name or path.
	Name string `pulumi:"name"`
}

// A collection of values returned by getNetwork.
type GetNetworkResult struct {
	DatacenterId                 *string            `pulumi:"datacenterId"`
	DistributedVirtualSwitchUuid *string            `pulumi:"distributedVirtualSwitchUuid"`
	Filters                      []GetNetworkFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The managed object type for the discovered network. This will be one
	// of `DistributedVirtualPortgroup` for distributed port groups, `Network` for
	// standard (host-based) port groups, or `OpaqueNetwork` for networks managed
	// externally, such as those managed by NSX.
	Type string `pulumi:"type"`
}

func GetNetworkOutput(ctx *pulumi.Context, args GetNetworkOutputArgs, opts ...pulumi.InvokeOption) GetNetworkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNetworkResultOutput, error) {
			args := v.(GetNetworkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vsphere:index/getNetwork:getNetwork", args, GetNetworkResultOutput{}, options).(GetNetworkResultOutput), nil
		}).(GetNetworkResultOutput)
}

// A collection of arguments for invoking getNetwork.
type GetNetworkOutputArgs struct {
	// The managed object reference ID
	// of the datacenter the network is located in. This can be omitted if the
	// search path used in `name` is an absolute path. For default datacenters,
	// use the `id` attribute from an empty `Datacenter` data source.
	DatacenterId pulumi.StringPtrInput `pulumi:"datacenterId"`
	// For distributed port group type
	// network objects, the ID of the distributed virtual switch for which the port
	// group belongs. It is useful to differentiate port groups with same name using
	// the distributed virtual switch ID.
	DistributedVirtualSwitchUuid pulumi.StringPtrInput `pulumi:"distributedVirtualSwitchUuid"`
	// Apply a filter for the discovered network.
	Filters GetNetworkFilterArrayInput `pulumi:"filters"`
	// The name of the network. This can be a name or path.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getNetwork.
type GetNetworkResultOutput struct{ *pulumi.OutputState }

func (GetNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkResult)(nil)).Elem()
}

func (o GetNetworkResultOutput) ToGetNetworkResultOutput() GetNetworkResultOutput {
	return o
}

func (o GetNetworkResultOutput) ToGetNetworkResultOutputWithContext(ctx context.Context) GetNetworkResultOutput {
	return o
}

func (o GetNetworkResultOutput) DatacenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworkResult) *string { return v.DatacenterId }).(pulumi.StringPtrOutput)
}

func (o GetNetworkResultOutput) DistributedVirtualSwitchUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNetworkResult) *string { return v.DistributedVirtualSwitchUuid }).(pulumi.StringPtrOutput)
}

func (o GetNetworkResultOutput) Filters() GetNetworkFilterArrayOutput {
	return o.ApplyT(func(v GetNetworkResult) []GetNetworkFilter { return v.Filters }).(GetNetworkFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNetworkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Name }).(pulumi.StringOutput)
}

// The managed object type for the discovered network. This will be one
// of `DistributedVirtualPortgroup` for distributed port groups, `Network` for
// standard (host-based) port groups, or `OpaqueNetwork` for networks managed
// externally, such as those managed by NSX.
func (o GetNetworkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworkResultOutput{})
}
