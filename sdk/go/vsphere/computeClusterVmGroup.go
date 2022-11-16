// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeClusterVmGroup struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds pulumi.StringArrayOutput `pulumi:"virtualMachineIds"`
}

// NewComputeClusterVmGroup registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmGroup(ctx *pulumi.Context,
	name string, args *ComputeClusterVmGroupArgs, opts ...pulumi.ResourceOption) (*ComputeClusterVmGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	var resource ComputeClusterVmGroup
	err := ctx.RegisterResource("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterVmGroup gets an existing ComputeClusterVmGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterVmGroupState, opts ...pulumi.ResourceOption) (*ComputeClusterVmGroup, error) {
	var resource ComputeClusterVmGroup
	err := ctx.ReadResource("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterVmGroup resources.
type computeClusterVmGroupState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

type ComputeClusterVmGroupState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmGroupState)(nil)).Elem()
}

type computeClusterVmGroupArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

// The set of arguments for constructing a ComputeClusterVmGroup resource.
type ComputeClusterVmGroupArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmGroupArgs)(nil)).Elem()
}

type ComputeClusterVmGroupInput interface {
	pulumi.Input

	ToComputeClusterVmGroupOutput() ComputeClusterVmGroupOutput
	ToComputeClusterVmGroupOutputWithContext(ctx context.Context) ComputeClusterVmGroupOutput
}

func (*ComputeClusterVmGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmGroup)(nil)).Elem()
}

func (i *ComputeClusterVmGroup) ToComputeClusterVmGroupOutput() ComputeClusterVmGroupOutput {
	return i.ToComputeClusterVmGroupOutputWithContext(context.Background())
}

func (i *ComputeClusterVmGroup) ToComputeClusterVmGroupOutputWithContext(ctx context.Context) ComputeClusterVmGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmGroupOutput)
}

// ComputeClusterVmGroupArrayInput is an input type that accepts ComputeClusterVmGroupArray and ComputeClusterVmGroupArrayOutput values.
// You can construct a concrete instance of `ComputeClusterVmGroupArrayInput` via:
//
//	ComputeClusterVmGroupArray{ ComputeClusterVmGroupArgs{...} }
type ComputeClusterVmGroupArrayInput interface {
	pulumi.Input

	ToComputeClusterVmGroupArrayOutput() ComputeClusterVmGroupArrayOutput
	ToComputeClusterVmGroupArrayOutputWithContext(context.Context) ComputeClusterVmGroupArrayOutput
}

type ComputeClusterVmGroupArray []ComputeClusterVmGroupInput

func (ComputeClusterVmGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeClusterVmGroup)(nil)).Elem()
}

func (i ComputeClusterVmGroupArray) ToComputeClusterVmGroupArrayOutput() ComputeClusterVmGroupArrayOutput {
	return i.ToComputeClusterVmGroupArrayOutputWithContext(context.Background())
}

func (i ComputeClusterVmGroupArray) ToComputeClusterVmGroupArrayOutputWithContext(ctx context.Context) ComputeClusterVmGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmGroupArrayOutput)
}

// ComputeClusterVmGroupMapInput is an input type that accepts ComputeClusterVmGroupMap and ComputeClusterVmGroupMapOutput values.
// You can construct a concrete instance of `ComputeClusterVmGroupMapInput` via:
//
//	ComputeClusterVmGroupMap{ "key": ComputeClusterVmGroupArgs{...} }
type ComputeClusterVmGroupMapInput interface {
	pulumi.Input

	ToComputeClusterVmGroupMapOutput() ComputeClusterVmGroupMapOutput
	ToComputeClusterVmGroupMapOutputWithContext(context.Context) ComputeClusterVmGroupMapOutput
}

type ComputeClusterVmGroupMap map[string]ComputeClusterVmGroupInput

func (ComputeClusterVmGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeClusterVmGroup)(nil)).Elem()
}

func (i ComputeClusterVmGroupMap) ToComputeClusterVmGroupMapOutput() ComputeClusterVmGroupMapOutput {
	return i.ToComputeClusterVmGroupMapOutputWithContext(context.Background())
}

func (i ComputeClusterVmGroupMap) ToComputeClusterVmGroupMapOutputWithContext(ctx context.Context) ComputeClusterVmGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmGroupMapOutput)
}

type ComputeClusterVmGroupOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmGroup)(nil)).Elem()
}

func (o ComputeClusterVmGroupOutput) ToComputeClusterVmGroupOutput() ComputeClusterVmGroupOutput {
	return o
}

func (o ComputeClusterVmGroupOutput) ToComputeClusterVmGroupOutputWithContext(ctx context.Context) ComputeClusterVmGroupOutput {
	return o
}

type ComputeClusterVmGroupArrayOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeClusterVmGroup)(nil)).Elem()
}

func (o ComputeClusterVmGroupArrayOutput) ToComputeClusterVmGroupArrayOutput() ComputeClusterVmGroupArrayOutput {
	return o
}

func (o ComputeClusterVmGroupArrayOutput) ToComputeClusterVmGroupArrayOutputWithContext(ctx context.Context) ComputeClusterVmGroupArrayOutput {
	return o
}

func (o ComputeClusterVmGroupArrayOutput) Index(i pulumi.IntInput) ComputeClusterVmGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ComputeClusterVmGroup {
		return vs[0].([]*ComputeClusterVmGroup)[vs[1].(int)]
	}).(ComputeClusterVmGroupOutput)
}

type ComputeClusterVmGroupMapOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeClusterVmGroup)(nil)).Elem()
}

func (o ComputeClusterVmGroupMapOutput) ToComputeClusterVmGroupMapOutput() ComputeClusterVmGroupMapOutput {
	return o
}

func (o ComputeClusterVmGroupMapOutput) ToComputeClusterVmGroupMapOutputWithContext(ctx context.Context) ComputeClusterVmGroupMapOutput {
	return o
}

func (o ComputeClusterVmGroupMapOutput) MapIndex(k pulumi.StringInput) ComputeClusterVmGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ComputeClusterVmGroup {
		return vs[0].(map[string]*ComputeClusterVmGroup)[vs[1].(string)]
	}).(ComputeClusterVmGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmGroupInput)(nil)).Elem(), &ComputeClusterVmGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmGroupArrayInput)(nil)).Elem(), ComputeClusterVmGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmGroupMapInput)(nil)).Elem(), ComputeClusterVmGroupMap{})
	pulumi.RegisterOutputType(ComputeClusterVmGroupOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmGroupArrayOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmGroupMapOutput{})
}
