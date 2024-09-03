// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DpmHostOverride struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// The automation level for host power
	// operations on this host. Can be one of `manual` or `automated`. Default:
	// `manual`.
	//
	// > **NOTE:** Using this resource *always* implies an override, even if one of
	// `dpmEnabled` or `dpmAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DpmAutomationLevel pulumi.StringPtrOutput `pulumi:"dpmAutomationLevel"`
	// Enable DPM support for this host. Default:
	// `false`.
	DpmEnabled pulumi.BoolPtrOutput `pulumi:"dpmEnabled"`
	// The managed object ID of the host.
	HostSystemId pulumi.StringOutput `pulumi:"hostSystemId"`
}

// NewDpmHostOverride registers a new resource with the given unique name, arguments, and options.
func NewDpmHostOverride(ctx *pulumi.Context,
	name string, args *DpmHostOverrideArgs, opts ...pulumi.ResourceOption) (*DpmHostOverride, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	if args.HostSystemId == nil {
		return nil, errors.New("invalid value for required argument 'HostSystemId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DpmHostOverride
	err := ctx.RegisterResource("vsphere:index/dpmHostOverride:DpmHostOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDpmHostOverride gets an existing DpmHostOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDpmHostOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DpmHostOverrideState, opts ...pulumi.ResourceOption) (*DpmHostOverride, error) {
	var resource DpmHostOverride
	err := ctx.ReadResource("vsphere:index/dpmHostOverride:DpmHostOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DpmHostOverride resources.
type dpmHostOverrideState struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// The automation level for host power
	// operations on this host. Can be one of `manual` or `automated`. Default:
	// `manual`.
	//
	// > **NOTE:** Using this resource *always* implies an override, even if one of
	// `dpmEnabled` or `dpmAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DpmAutomationLevel *string `pulumi:"dpmAutomationLevel"`
	// Enable DPM support for this host. Default:
	// `false`.
	DpmEnabled *bool `pulumi:"dpmEnabled"`
	// The managed object ID of the host.
	HostSystemId *string `pulumi:"hostSystemId"`
}

type DpmHostOverrideState struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// The automation level for host power
	// operations on this host. Can be one of `manual` or `automated`. Default:
	// `manual`.
	//
	// > **NOTE:** Using this resource *always* implies an override, even if one of
	// `dpmEnabled` or `dpmAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DpmAutomationLevel pulumi.StringPtrInput
	// Enable DPM support for this host. Default:
	// `false`.
	DpmEnabled pulumi.BoolPtrInput
	// The managed object ID of the host.
	HostSystemId pulumi.StringPtrInput
}

func (DpmHostOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*dpmHostOverrideState)(nil)).Elem()
}

type dpmHostOverrideArgs struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// The automation level for host power
	// operations on this host. Can be one of `manual` or `automated`. Default:
	// `manual`.
	//
	// > **NOTE:** Using this resource *always* implies an override, even if one of
	// `dpmEnabled` or `dpmAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DpmAutomationLevel *string `pulumi:"dpmAutomationLevel"`
	// Enable DPM support for this host. Default:
	// `false`.
	DpmEnabled *bool `pulumi:"dpmEnabled"`
	// The managed object ID of the host.
	HostSystemId string `pulumi:"hostSystemId"`
}

// The set of arguments for constructing a DpmHostOverride resource.
type DpmHostOverrideArgs struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// The automation level for host power
	// operations on this host. Can be one of `manual` or `automated`. Default:
	// `manual`.
	//
	// > **NOTE:** Using this resource *always* implies an override, even if one of
	// `dpmEnabled` or `dpmAutomationLevel` is omitted. Take note of the defaults
	// for both options.
	DpmAutomationLevel pulumi.StringPtrInput
	// Enable DPM support for this host. Default:
	// `false`.
	DpmEnabled pulumi.BoolPtrInput
	// The managed object ID of the host.
	HostSystemId pulumi.StringInput
}

func (DpmHostOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dpmHostOverrideArgs)(nil)).Elem()
}

type DpmHostOverrideInput interface {
	pulumi.Input

	ToDpmHostOverrideOutput() DpmHostOverrideOutput
	ToDpmHostOverrideOutputWithContext(ctx context.Context) DpmHostOverrideOutput
}

func (*DpmHostOverride) ElementType() reflect.Type {
	return reflect.TypeOf((**DpmHostOverride)(nil)).Elem()
}

func (i *DpmHostOverride) ToDpmHostOverrideOutput() DpmHostOverrideOutput {
	return i.ToDpmHostOverrideOutputWithContext(context.Background())
}

func (i *DpmHostOverride) ToDpmHostOverrideOutputWithContext(ctx context.Context) DpmHostOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpmHostOverrideOutput)
}

// DpmHostOverrideArrayInput is an input type that accepts DpmHostOverrideArray and DpmHostOverrideArrayOutput values.
// You can construct a concrete instance of `DpmHostOverrideArrayInput` via:
//
//	DpmHostOverrideArray{ DpmHostOverrideArgs{...} }
type DpmHostOverrideArrayInput interface {
	pulumi.Input

	ToDpmHostOverrideArrayOutput() DpmHostOverrideArrayOutput
	ToDpmHostOverrideArrayOutputWithContext(context.Context) DpmHostOverrideArrayOutput
}

type DpmHostOverrideArray []DpmHostOverrideInput

func (DpmHostOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DpmHostOverride)(nil)).Elem()
}

func (i DpmHostOverrideArray) ToDpmHostOverrideArrayOutput() DpmHostOverrideArrayOutput {
	return i.ToDpmHostOverrideArrayOutputWithContext(context.Background())
}

func (i DpmHostOverrideArray) ToDpmHostOverrideArrayOutputWithContext(ctx context.Context) DpmHostOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpmHostOverrideArrayOutput)
}

// DpmHostOverrideMapInput is an input type that accepts DpmHostOverrideMap and DpmHostOverrideMapOutput values.
// You can construct a concrete instance of `DpmHostOverrideMapInput` via:
//
//	DpmHostOverrideMap{ "key": DpmHostOverrideArgs{...} }
type DpmHostOverrideMapInput interface {
	pulumi.Input

	ToDpmHostOverrideMapOutput() DpmHostOverrideMapOutput
	ToDpmHostOverrideMapOutputWithContext(context.Context) DpmHostOverrideMapOutput
}

type DpmHostOverrideMap map[string]DpmHostOverrideInput

func (DpmHostOverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DpmHostOverride)(nil)).Elem()
}

func (i DpmHostOverrideMap) ToDpmHostOverrideMapOutput() DpmHostOverrideMapOutput {
	return i.ToDpmHostOverrideMapOutputWithContext(context.Background())
}

func (i DpmHostOverrideMap) ToDpmHostOverrideMapOutputWithContext(ctx context.Context) DpmHostOverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DpmHostOverrideMapOutput)
}

type DpmHostOverrideOutput struct{ *pulumi.OutputState }

func (DpmHostOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DpmHostOverride)(nil)).Elem()
}

func (o DpmHostOverrideOutput) ToDpmHostOverrideOutput() DpmHostOverrideOutput {
	return o
}

func (o DpmHostOverrideOutput) ToDpmHostOverrideOutputWithContext(ctx context.Context) DpmHostOverrideOutput {
	return o
}

// The managed object reference
// ID of the cluster to put the override in.  Forces a new
// resource if changed.
func (o DpmHostOverrideOutput) ComputeClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DpmHostOverride) pulumi.StringOutput { return v.ComputeClusterId }).(pulumi.StringOutput)
}

// The automation level for host power
// operations on this host. Can be one of `manual` or `automated`. Default:
// `manual`.
//
// > **NOTE:** Using this resource *always* implies an override, even if one of
// `dpmEnabled` or `dpmAutomationLevel` is omitted. Take note of the defaults
// for both options.
func (o DpmHostOverrideOutput) DpmAutomationLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DpmHostOverride) pulumi.StringPtrOutput { return v.DpmAutomationLevel }).(pulumi.StringPtrOutput)
}

// Enable DPM support for this host. Default:
// `false`.
func (o DpmHostOverrideOutput) DpmEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DpmHostOverride) pulumi.BoolPtrOutput { return v.DpmEnabled }).(pulumi.BoolPtrOutput)
}

// The managed object ID of the host.
func (o DpmHostOverrideOutput) HostSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *DpmHostOverride) pulumi.StringOutput { return v.HostSystemId }).(pulumi.StringOutput)
}

type DpmHostOverrideArrayOutput struct{ *pulumi.OutputState }

func (DpmHostOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DpmHostOverride)(nil)).Elem()
}

func (o DpmHostOverrideArrayOutput) ToDpmHostOverrideArrayOutput() DpmHostOverrideArrayOutput {
	return o
}

func (o DpmHostOverrideArrayOutput) ToDpmHostOverrideArrayOutputWithContext(ctx context.Context) DpmHostOverrideArrayOutput {
	return o
}

func (o DpmHostOverrideArrayOutput) Index(i pulumi.IntInput) DpmHostOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DpmHostOverride {
		return vs[0].([]*DpmHostOverride)[vs[1].(int)]
	}).(DpmHostOverrideOutput)
}

type DpmHostOverrideMapOutput struct{ *pulumi.OutputState }

func (DpmHostOverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DpmHostOverride)(nil)).Elem()
}

func (o DpmHostOverrideMapOutput) ToDpmHostOverrideMapOutput() DpmHostOverrideMapOutput {
	return o
}

func (o DpmHostOverrideMapOutput) ToDpmHostOverrideMapOutputWithContext(ctx context.Context) DpmHostOverrideMapOutput {
	return o
}

func (o DpmHostOverrideMapOutput) MapIndex(k pulumi.StringInput) DpmHostOverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DpmHostOverride {
		return vs[0].(map[string]*DpmHostOverride)[vs[1].(string)]
	}).(DpmHostOverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DpmHostOverrideInput)(nil)).Elem(), &DpmHostOverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*DpmHostOverrideArrayInput)(nil)).Elem(), DpmHostOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DpmHostOverrideMapInput)(nil)).Elem(), DpmHostOverrideMap{})
	pulumi.RegisterOutputType(DpmHostOverrideOutput{})
	pulumi.RegisterOutputType(DpmHostOverrideArrayOutput{})
	pulumi.RegisterOutputType(DpmHostOverrideMapOutput{})
}
