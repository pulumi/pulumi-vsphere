// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DrsVmOverride struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel pulumi.StringPtrOutput `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrOutput `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
}

// NewDrsVmOverride registers a new resource with the given unique name, arguments, and options.
func NewDrsVmOverride(ctx *pulumi.Context,
	name string, args *DrsVmOverrideArgs, opts ...pulumi.ResourceOption) (*DrsVmOverride, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	if args.VirtualMachineId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineId'")
	}
	var resource DrsVmOverride
	err := ctx.RegisterResource("vsphere:index/drsVmOverride:DrsVmOverride", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDrsVmOverride gets an existing DrsVmOverride resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDrsVmOverride(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DrsVmOverrideState, opts ...pulumi.ResourceOption) (*DrsVmOverride, error) {
	var resource DrsVmOverride
	err := ctx.ReadResource("vsphere:index/drsVmOverride:DrsVmOverride", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DrsVmOverride resources.
type drsVmOverrideState struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel *string `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled *bool `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type DrsVmOverrideState struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringPtrInput
}

func (DrsVmOverrideState) ElementType() reflect.Type {
	return reflect.TypeOf((*drsVmOverrideState)(nil)).Elem()
}

type drsVmOverrideArgs struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel *string `pulumi:"drsAutomationLevel"`
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled *bool `pulumi:"drsEnabled"`
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId string `pulumi:"virtualMachineId"`
}

// The set of arguments for constructing a DrsVmOverride resource.
type DrsVmOverrideArgs struct {
	// The managed object reference
	// ID of the cluster to put the override in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// Overrides the automation level for this virtual
	// machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
	// `fullyAutomated`. Default: `manual`.
	DrsAutomationLevel pulumi.StringPtrInput
	// Overrides the default DRS setting for this virtual
	// machine. Can be either `true` or `false`. Default: `false`.
	DrsEnabled pulumi.BoolPtrInput
	// The UUID of the virtual machine to create
	// the override for.  Forces a new resource if changed.
	VirtualMachineId pulumi.StringInput
}

func (DrsVmOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*drsVmOverrideArgs)(nil)).Elem()
}

type DrsVmOverrideInput interface {
	pulumi.Input

	ToDrsVmOverrideOutput() DrsVmOverrideOutput
	ToDrsVmOverrideOutputWithContext(ctx context.Context) DrsVmOverrideOutput
}

func (*DrsVmOverride) ElementType() reflect.Type {
	return reflect.TypeOf((**DrsVmOverride)(nil)).Elem()
}

func (i *DrsVmOverride) ToDrsVmOverrideOutput() DrsVmOverrideOutput {
	return i.ToDrsVmOverrideOutputWithContext(context.Background())
}

func (i *DrsVmOverride) ToDrsVmOverrideOutputWithContext(ctx context.Context) DrsVmOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrsVmOverrideOutput)
}

// DrsVmOverrideArrayInput is an input type that accepts DrsVmOverrideArray and DrsVmOverrideArrayOutput values.
// You can construct a concrete instance of `DrsVmOverrideArrayInput` via:
//
//	DrsVmOverrideArray{ DrsVmOverrideArgs{...} }
type DrsVmOverrideArrayInput interface {
	pulumi.Input

	ToDrsVmOverrideArrayOutput() DrsVmOverrideArrayOutput
	ToDrsVmOverrideArrayOutputWithContext(context.Context) DrsVmOverrideArrayOutput
}

type DrsVmOverrideArray []DrsVmOverrideInput

func (DrsVmOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DrsVmOverride)(nil)).Elem()
}

func (i DrsVmOverrideArray) ToDrsVmOverrideArrayOutput() DrsVmOverrideArrayOutput {
	return i.ToDrsVmOverrideArrayOutputWithContext(context.Background())
}

func (i DrsVmOverrideArray) ToDrsVmOverrideArrayOutputWithContext(ctx context.Context) DrsVmOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrsVmOverrideArrayOutput)
}

// DrsVmOverrideMapInput is an input type that accepts DrsVmOverrideMap and DrsVmOverrideMapOutput values.
// You can construct a concrete instance of `DrsVmOverrideMapInput` via:
//
//	DrsVmOverrideMap{ "key": DrsVmOverrideArgs{...} }
type DrsVmOverrideMapInput interface {
	pulumi.Input

	ToDrsVmOverrideMapOutput() DrsVmOverrideMapOutput
	ToDrsVmOverrideMapOutputWithContext(context.Context) DrsVmOverrideMapOutput
}

type DrsVmOverrideMap map[string]DrsVmOverrideInput

func (DrsVmOverrideMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DrsVmOverride)(nil)).Elem()
}

func (i DrsVmOverrideMap) ToDrsVmOverrideMapOutput() DrsVmOverrideMapOutput {
	return i.ToDrsVmOverrideMapOutputWithContext(context.Background())
}

func (i DrsVmOverrideMap) ToDrsVmOverrideMapOutputWithContext(ctx context.Context) DrsVmOverrideMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrsVmOverrideMapOutput)
}

type DrsVmOverrideOutput struct{ *pulumi.OutputState }

func (DrsVmOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DrsVmOverride)(nil)).Elem()
}

func (o DrsVmOverrideOutput) ToDrsVmOverrideOutput() DrsVmOverrideOutput {
	return o
}

func (o DrsVmOverrideOutput) ToDrsVmOverrideOutputWithContext(ctx context.Context) DrsVmOverrideOutput {
	return o
}

type DrsVmOverrideArrayOutput struct{ *pulumi.OutputState }

func (DrsVmOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DrsVmOverride)(nil)).Elem()
}

func (o DrsVmOverrideArrayOutput) ToDrsVmOverrideArrayOutput() DrsVmOverrideArrayOutput {
	return o
}

func (o DrsVmOverrideArrayOutput) ToDrsVmOverrideArrayOutputWithContext(ctx context.Context) DrsVmOverrideArrayOutput {
	return o
}

func (o DrsVmOverrideArrayOutput) Index(i pulumi.IntInput) DrsVmOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DrsVmOverride {
		return vs[0].([]*DrsVmOverride)[vs[1].(int)]
	}).(DrsVmOverrideOutput)
}

type DrsVmOverrideMapOutput struct{ *pulumi.OutputState }

func (DrsVmOverrideMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DrsVmOverride)(nil)).Elem()
}

func (o DrsVmOverrideMapOutput) ToDrsVmOverrideMapOutput() DrsVmOverrideMapOutput {
	return o
}

func (o DrsVmOverrideMapOutput) ToDrsVmOverrideMapOutputWithContext(ctx context.Context) DrsVmOverrideMapOutput {
	return o
}

func (o DrsVmOverrideMapOutput) MapIndex(k pulumi.StringInput) DrsVmOverrideOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DrsVmOverride {
		return vs[0].(map[string]*DrsVmOverride)[vs[1].(string)]
	}).(DrsVmOverrideOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DrsVmOverrideInput)(nil)).Elem(), &DrsVmOverride{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrsVmOverrideArrayInput)(nil)).Elem(), DrsVmOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DrsVmOverrideMapInput)(nil)).Elem(), DrsVmOverrideMap{})
	pulumi.RegisterOutputType(DrsVmOverrideOutput{})
	pulumi.RegisterOutputType(DrsVmOverrideArrayOutput{})
	pulumi.RegisterOutputType(DrsVmOverrideMapOutput{})
}
