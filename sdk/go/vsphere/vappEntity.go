// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vsphere/sdk/v4/go/vsphere/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type VappEntity struct {
	pulumi.CustomResourceState

	// Managed object ID of the vApp
	// container the entity is a member of.
	ContainerId pulumi.StringOutput `pulumi:"containerId"`
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	StartAction pulumi.StringPtrOutput `pulumi:"startAction"`
	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	StartDelay pulumi.IntPtrOutput `pulumi:"startDelay"`
	// Order to start and stop target in vApp. Default: 1
	StartOrder pulumi.IntPtrOutput `pulumi:"startOrder"`
	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	StopAction pulumi.StringPtrOutput `pulumi:"stopAction"`
	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	StopDelay pulumi.IntPtrOutput `pulumi:"stopDelay"`
	// A list of tag IDs to apply to this object.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Managed object ID of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for `startDelay`. This
	// property has no effect for vApps. Default: false
	WaitForGuest pulumi.BoolPtrOutput `pulumi:"waitForGuest"`
}

// NewVappEntity registers a new resource with the given unique name, arguments, and options.
func NewVappEntity(ctx *pulumi.Context,
	name string, args *VappEntityArgs, opts ...pulumi.ResourceOption) (*VappEntity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerId == nil {
		return nil, errors.New("invalid value for required argument 'ContainerId'")
	}
	if args.TargetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VappEntity
	err := ctx.RegisterResource("vsphere:index/vappEntity:VappEntity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVappEntity gets an existing VappEntity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVappEntity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VappEntityState, opts ...pulumi.ResourceOption) (*VappEntity, error) {
	var resource VappEntity
	err := ctx.ReadResource("vsphere:index/vappEntity:VappEntity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VappEntity resources.
type vappEntityState struct {
	// Managed object ID of the vApp
	// container the entity is a member of.
	ContainerId *string `pulumi:"containerId"`
	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	StartAction *string `pulumi:"startAction"`
	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	StartDelay *int `pulumi:"startDelay"`
	// Order to start and stop target in vApp. Default: 1
	StartOrder *int `pulumi:"startOrder"`
	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	StopAction *string `pulumi:"stopAction"`
	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	StopDelay *int `pulumi:"stopDelay"`
	// A list of tag IDs to apply to this object.
	Tags []string `pulumi:"tags"`
	// Managed object ID of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	TargetId *string `pulumi:"targetId"`
	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for `startDelay`. This
	// property has no effect for vApps. Default: false
	WaitForGuest *bool `pulumi:"waitForGuest"`
}

type VappEntityState struct {
	// Managed object ID of the vApp
	// container the entity is a member of.
	ContainerId pulumi.StringPtrInput
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapInput
	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	StartAction pulumi.StringPtrInput
	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	StartDelay pulumi.IntPtrInput
	// Order to start and stop target in vApp. Default: 1
	StartOrder pulumi.IntPtrInput
	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	StopAction pulumi.StringPtrInput
	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	StopDelay pulumi.IntPtrInput
	// A list of tag IDs to apply to this object.
	Tags pulumi.StringArrayInput
	// Managed object ID of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	TargetId pulumi.StringPtrInput
	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for `startDelay`. This
	// property has no effect for vApps. Default: false
	WaitForGuest pulumi.BoolPtrInput
}

func (VappEntityState) ElementType() reflect.Type {
	return reflect.TypeOf((*vappEntityState)(nil)).Elem()
}

type vappEntityArgs struct {
	// Managed object ID of the vApp
	// container the entity is a member of.
	ContainerId string `pulumi:"containerId"`
	// A list of custom attributes to set on this resource.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	StartAction *string `pulumi:"startAction"`
	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	StartDelay *int `pulumi:"startDelay"`
	// Order to start and stop target in vApp. Default: 1
	StartOrder *int `pulumi:"startOrder"`
	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	StopAction *string `pulumi:"stopAction"`
	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	StopDelay *int `pulumi:"stopDelay"`
	// A list of tag IDs to apply to this object.
	Tags []string `pulumi:"tags"`
	// Managed object ID of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	TargetId string `pulumi:"targetId"`
	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for `startDelay`. This
	// property has no effect for vApps. Default: false
	WaitForGuest *bool `pulumi:"waitForGuest"`
}

// The set of arguments for constructing a VappEntity resource.
type VappEntityArgs struct {
	// Managed object ID of the vApp
	// container the entity is a member of.
	ContainerId pulumi.StringInput
	// A list of custom attributes to set on this resource.
	CustomAttributes pulumi.StringMapInput
	// How to start the entity. Valid settings are none
	// or powerOn. If set to none, then the entity does not participate in auto-start.
	// Default: powerOn
	StartAction pulumi.StringPtrInput
	// Delay in seconds before continuing with the next
	// entity in the order of entities to be started. Default: 120
	StartDelay pulumi.IntPtrInput
	// Order to start and stop target in vApp. Default: 1
	StartOrder pulumi.IntPtrInput
	// Defines the stop action for the entity. Can be set
	// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
	// does not participate in auto-stop. Default: powerOff
	StopAction pulumi.StringPtrInput
	// Delay in seconds before continuing with the next
	// entity in the order sequence. This is only used if the stopAction is
	// guestShutdown. Default: 120
	StopDelay pulumi.IntPtrInput
	// A list of tag IDs to apply to this object.
	Tags pulumi.StringArrayInput
	// Managed object ID of the entity
	// to power on or power off. This can be a virtual machine or a vApp.
	TargetId pulumi.StringInput
	// Determines if the VM should be marked as being
	// started when VMware Tools are ready instead of waiting for `startDelay`. This
	// property has no effect for vApps. Default: false
	WaitForGuest pulumi.BoolPtrInput
}

func (VappEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vappEntityArgs)(nil)).Elem()
}

type VappEntityInput interface {
	pulumi.Input

	ToVappEntityOutput() VappEntityOutput
	ToVappEntityOutputWithContext(ctx context.Context) VappEntityOutput
}

func (*VappEntity) ElementType() reflect.Type {
	return reflect.TypeOf((**VappEntity)(nil)).Elem()
}

func (i *VappEntity) ToVappEntityOutput() VappEntityOutput {
	return i.ToVappEntityOutputWithContext(context.Background())
}

func (i *VappEntity) ToVappEntityOutputWithContext(ctx context.Context) VappEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappEntityOutput)
}

func (i *VappEntity) ToOutput(ctx context.Context) pulumix.Output[*VappEntity] {
	return pulumix.Output[*VappEntity]{
		OutputState: i.ToVappEntityOutputWithContext(ctx).OutputState,
	}
}

// VappEntityArrayInput is an input type that accepts VappEntityArray and VappEntityArrayOutput values.
// You can construct a concrete instance of `VappEntityArrayInput` via:
//
//	VappEntityArray{ VappEntityArgs{...} }
type VappEntityArrayInput interface {
	pulumi.Input

	ToVappEntityArrayOutput() VappEntityArrayOutput
	ToVappEntityArrayOutputWithContext(context.Context) VappEntityArrayOutput
}

type VappEntityArray []VappEntityInput

func (VappEntityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappEntity)(nil)).Elem()
}

func (i VappEntityArray) ToVappEntityArrayOutput() VappEntityArrayOutput {
	return i.ToVappEntityArrayOutputWithContext(context.Background())
}

func (i VappEntityArray) ToVappEntityArrayOutputWithContext(ctx context.Context) VappEntityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappEntityArrayOutput)
}

func (i VappEntityArray) ToOutput(ctx context.Context) pulumix.Output[[]*VappEntity] {
	return pulumix.Output[[]*VappEntity]{
		OutputState: i.ToVappEntityArrayOutputWithContext(ctx).OutputState,
	}
}

// VappEntityMapInput is an input type that accepts VappEntityMap and VappEntityMapOutput values.
// You can construct a concrete instance of `VappEntityMapInput` via:
//
//	VappEntityMap{ "key": VappEntityArgs{...} }
type VappEntityMapInput interface {
	pulumi.Input

	ToVappEntityMapOutput() VappEntityMapOutput
	ToVappEntityMapOutputWithContext(context.Context) VappEntityMapOutput
}

type VappEntityMap map[string]VappEntityInput

func (VappEntityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappEntity)(nil)).Elem()
}

func (i VappEntityMap) ToVappEntityMapOutput() VappEntityMapOutput {
	return i.ToVappEntityMapOutputWithContext(context.Background())
}

func (i VappEntityMap) ToVappEntityMapOutputWithContext(ctx context.Context) VappEntityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappEntityMapOutput)
}

func (i VappEntityMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VappEntity] {
	return pulumix.Output[map[string]*VappEntity]{
		OutputState: i.ToVappEntityMapOutputWithContext(ctx).OutputState,
	}
}

type VappEntityOutput struct{ *pulumi.OutputState }

func (VappEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VappEntity)(nil)).Elem()
}

func (o VappEntityOutput) ToVappEntityOutput() VappEntityOutput {
	return o
}

func (o VappEntityOutput) ToVappEntityOutputWithContext(ctx context.Context) VappEntityOutput {
	return o
}

func (o VappEntityOutput) ToOutput(ctx context.Context) pulumix.Output[*VappEntity] {
	return pulumix.Output[*VappEntity]{
		OutputState: o.OutputState,
	}
}

// Managed object ID of the vApp
// container the entity is a member of.
func (o VappEntityOutput) ContainerId() pulumi.StringOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.StringOutput { return v.ContainerId }).(pulumi.StringOutput)
}

// A list of custom attributes to set on this resource.
func (o VappEntityOutput) CustomAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.StringMapOutput { return v.CustomAttributes }).(pulumi.StringMapOutput)
}

// How to start the entity. Valid settings are none
// or powerOn. If set to none, then the entity does not participate in auto-start.
// Default: powerOn
func (o VappEntityOutput) StartAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.StringPtrOutput { return v.StartAction }).(pulumi.StringPtrOutput)
}

// Delay in seconds before continuing with the next
// entity in the order of entities to be started. Default: 120
func (o VappEntityOutput) StartDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.IntPtrOutput { return v.StartDelay }).(pulumi.IntPtrOutput)
}

// Order to start and stop target in vApp. Default: 1
func (o VappEntityOutput) StartOrder() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.IntPtrOutput { return v.StartOrder }).(pulumi.IntPtrOutput)
}

// Defines the stop action for the entity. Can be set
// to none, powerOff, guestShutdown, or suspend. If set to none, then the entity
// does not participate in auto-stop. Default: powerOff
func (o VappEntityOutput) StopAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.StringPtrOutput { return v.StopAction }).(pulumi.StringPtrOutput)
}

// Delay in seconds before continuing with the next
// entity in the order sequence. This is only used if the stopAction is
// guestShutdown. Default: 120
func (o VappEntityOutput) StopDelay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.IntPtrOutput { return v.StopDelay }).(pulumi.IntPtrOutput)
}

// A list of tag IDs to apply to this object.
func (o VappEntityOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Managed object ID of the entity
// to power on or power off. This can be a virtual machine or a vApp.
func (o VappEntityOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.StringOutput { return v.TargetId }).(pulumi.StringOutput)
}

// Determines if the VM should be marked as being
// started when VMware Tools are ready instead of waiting for `startDelay`. This
// property has no effect for vApps. Default: false
func (o VappEntityOutput) WaitForGuest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VappEntity) pulumi.BoolPtrOutput { return v.WaitForGuest }).(pulumi.BoolPtrOutput)
}

type VappEntityArrayOutput struct{ *pulumi.OutputState }

func (VappEntityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappEntity)(nil)).Elem()
}

func (o VappEntityArrayOutput) ToVappEntityArrayOutput() VappEntityArrayOutput {
	return o
}

func (o VappEntityArrayOutput) ToVappEntityArrayOutputWithContext(ctx context.Context) VappEntityArrayOutput {
	return o
}

func (o VappEntityArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VappEntity] {
	return pulumix.Output[[]*VappEntity]{
		OutputState: o.OutputState,
	}
}

func (o VappEntityArrayOutput) Index(i pulumi.IntInput) VappEntityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VappEntity {
		return vs[0].([]*VappEntity)[vs[1].(int)]
	}).(VappEntityOutput)
}

type VappEntityMapOutput struct{ *pulumi.OutputState }

func (VappEntityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappEntity)(nil)).Elem()
}

func (o VappEntityMapOutput) ToVappEntityMapOutput() VappEntityMapOutput {
	return o
}

func (o VappEntityMapOutput) ToVappEntityMapOutputWithContext(ctx context.Context) VappEntityMapOutput {
	return o
}

func (o VappEntityMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VappEntity] {
	return pulumix.Output[map[string]*VappEntity]{
		OutputState: o.OutputState,
	}
}

func (o VappEntityMapOutput) MapIndex(k pulumi.StringInput) VappEntityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VappEntity {
		return vs[0].(map[string]*VappEntity)[vs[1].(string)]
	}).(VappEntityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VappEntityInput)(nil)).Elem(), &VappEntity{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappEntityArrayInput)(nil)).Elem(), VappEntityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappEntityMapInput)(nil)).Elem(), VappEntityMap{})
	pulumi.RegisterOutputType(VappEntityOutput{})
	pulumi.RegisterOutputType(VappEntityArrayOutput{})
	pulumi.RegisterOutputType(VappEntityMapOutput{})
}
