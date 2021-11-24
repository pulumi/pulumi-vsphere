// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeClusterHostGroup struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds pulumi.StringArrayOutput `pulumi:"hostSystemIds"`
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewComputeClusterHostGroup registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterHostGroup(ctx *pulumi.Context,
	name string, args *ComputeClusterHostGroupArgs, opts ...pulumi.ResourceOption) (*ComputeClusterHostGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	var resource ComputeClusterHostGroup
	err := ctx.RegisterResource("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterHostGroup gets an existing ComputeClusterHostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterHostGroupState, opts ...pulumi.ResourceOption) (*ComputeClusterHostGroup, error) {
	var resource ComputeClusterHostGroup
	err := ctx.ReadResource("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterHostGroup resources.
type computeClusterHostGroupState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds []string `pulumi:"hostSystemIds"`
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
}

type ComputeClusterHostGroupState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds pulumi.StringArrayInput
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
}

func (ComputeClusterHostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterHostGroupState)(nil)).Elem()
}

type computeClusterHostGroupArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds []string `pulumi:"hostSystemIds"`
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ComputeClusterHostGroup resource.
type ComputeClusterHostGroupArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// The managed object IDs of
	// the hosts to put in the cluster.
	HostSystemIds pulumi.StringArrayInput
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
}

func (ComputeClusterHostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterHostGroupArgs)(nil)).Elem()
}

type ComputeClusterHostGroupInput interface {
	pulumi.Input

	ToComputeClusterHostGroupOutput() ComputeClusterHostGroupOutput
	ToComputeClusterHostGroupOutputWithContext(ctx context.Context) ComputeClusterHostGroupOutput
}

func (*ComputeClusterHostGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeClusterHostGroup)(nil))
}

func (i *ComputeClusterHostGroup) ToComputeClusterHostGroupOutput() ComputeClusterHostGroupOutput {
	return i.ToComputeClusterHostGroupOutputWithContext(context.Background())
}

func (i *ComputeClusterHostGroup) ToComputeClusterHostGroupOutputWithContext(ctx context.Context) ComputeClusterHostGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterHostGroupOutput)
}

func (i *ComputeClusterHostGroup) ToComputeClusterHostGroupPtrOutput() ComputeClusterHostGroupPtrOutput {
	return i.ToComputeClusterHostGroupPtrOutputWithContext(context.Background())
}

func (i *ComputeClusterHostGroup) ToComputeClusterHostGroupPtrOutputWithContext(ctx context.Context) ComputeClusterHostGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterHostGroupPtrOutput)
}

type ComputeClusterHostGroupPtrInput interface {
	pulumi.Input

	ToComputeClusterHostGroupPtrOutput() ComputeClusterHostGroupPtrOutput
	ToComputeClusterHostGroupPtrOutputWithContext(ctx context.Context) ComputeClusterHostGroupPtrOutput
}

type computeClusterHostGroupPtrType ComputeClusterHostGroupArgs

func (*computeClusterHostGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterHostGroup)(nil))
}

func (i *computeClusterHostGroupPtrType) ToComputeClusterHostGroupPtrOutput() ComputeClusterHostGroupPtrOutput {
	return i.ToComputeClusterHostGroupPtrOutputWithContext(context.Background())
}

func (i *computeClusterHostGroupPtrType) ToComputeClusterHostGroupPtrOutputWithContext(ctx context.Context) ComputeClusterHostGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterHostGroupPtrOutput)
}

// ComputeClusterHostGroupArrayInput is an input type that accepts ComputeClusterHostGroupArray and ComputeClusterHostGroupArrayOutput values.
// You can construct a concrete instance of `ComputeClusterHostGroupArrayInput` via:
//
//          ComputeClusterHostGroupArray{ ComputeClusterHostGroupArgs{...} }
type ComputeClusterHostGroupArrayInput interface {
	pulumi.Input

	ToComputeClusterHostGroupArrayOutput() ComputeClusterHostGroupArrayOutput
	ToComputeClusterHostGroupArrayOutputWithContext(context.Context) ComputeClusterHostGroupArrayOutput
}

type ComputeClusterHostGroupArray []ComputeClusterHostGroupInput

func (ComputeClusterHostGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeClusterHostGroup)(nil)).Elem()
}

func (i ComputeClusterHostGroupArray) ToComputeClusterHostGroupArrayOutput() ComputeClusterHostGroupArrayOutput {
	return i.ToComputeClusterHostGroupArrayOutputWithContext(context.Background())
}

func (i ComputeClusterHostGroupArray) ToComputeClusterHostGroupArrayOutputWithContext(ctx context.Context) ComputeClusterHostGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterHostGroupArrayOutput)
}

// ComputeClusterHostGroupMapInput is an input type that accepts ComputeClusterHostGroupMap and ComputeClusterHostGroupMapOutput values.
// You can construct a concrete instance of `ComputeClusterHostGroupMapInput` via:
//
//          ComputeClusterHostGroupMap{ "key": ComputeClusterHostGroupArgs{...} }
type ComputeClusterHostGroupMapInput interface {
	pulumi.Input

	ToComputeClusterHostGroupMapOutput() ComputeClusterHostGroupMapOutput
	ToComputeClusterHostGroupMapOutputWithContext(context.Context) ComputeClusterHostGroupMapOutput
}

type ComputeClusterHostGroupMap map[string]ComputeClusterHostGroupInput

func (ComputeClusterHostGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeClusterHostGroup)(nil)).Elem()
}

func (i ComputeClusterHostGroupMap) ToComputeClusterHostGroupMapOutput() ComputeClusterHostGroupMapOutput {
	return i.ToComputeClusterHostGroupMapOutputWithContext(context.Background())
}

func (i ComputeClusterHostGroupMap) ToComputeClusterHostGroupMapOutputWithContext(ctx context.Context) ComputeClusterHostGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterHostGroupMapOutput)
}

type ComputeClusterHostGroupOutput struct{ *pulumi.OutputState }

func (ComputeClusterHostGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeClusterHostGroup)(nil))
}

func (o ComputeClusterHostGroupOutput) ToComputeClusterHostGroupOutput() ComputeClusterHostGroupOutput {
	return o
}

func (o ComputeClusterHostGroupOutput) ToComputeClusterHostGroupOutputWithContext(ctx context.Context) ComputeClusterHostGroupOutput {
	return o
}

func (o ComputeClusterHostGroupOutput) ToComputeClusterHostGroupPtrOutput() ComputeClusterHostGroupPtrOutput {
	return o.ToComputeClusterHostGroupPtrOutputWithContext(context.Background())
}

func (o ComputeClusterHostGroupOutput) ToComputeClusterHostGroupPtrOutputWithContext(ctx context.Context) ComputeClusterHostGroupPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ComputeClusterHostGroup) *ComputeClusterHostGroup {
		return &v
	}).(ComputeClusterHostGroupPtrOutput)
}

type ComputeClusterHostGroupPtrOutput struct{ *pulumi.OutputState }

func (ComputeClusterHostGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterHostGroup)(nil))
}

func (o ComputeClusterHostGroupPtrOutput) ToComputeClusterHostGroupPtrOutput() ComputeClusterHostGroupPtrOutput {
	return o
}

func (o ComputeClusterHostGroupPtrOutput) ToComputeClusterHostGroupPtrOutputWithContext(ctx context.Context) ComputeClusterHostGroupPtrOutput {
	return o
}

func (o ComputeClusterHostGroupPtrOutput) Elem() ComputeClusterHostGroupOutput {
	return o.ApplyT(func(v *ComputeClusterHostGroup) ComputeClusterHostGroup {
		if v != nil {
			return *v
		}
		var ret ComputeClusterHostGroup
		return ret
	}).(ComputeClusterHostGroupOutput)
}

type ComputeClusterHostGroupArrayOutput struct{ *pulumi.OutputState }

func (ComputeClusterHostGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeClusterHostGroup)(nil))
}

func (o ComputeClusterHostGroupArrayOutput) ToComputeClusterHostGroupArrayOutput() ComputeClusterHostGroupArrayOutput {
	return o
}

func (o ComputeClusterHostGroupArrayOutput) ToComputeClusterHostGroupArrayOutputWithContext(ctx context.Context) ComputeClusterHostGroupArrayOutput {
	return o
}

func (o ComputeClusterHostGroupArrayOutput) Index(i pulumi.IntInput) ComputeClusterHostGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputeClusterHostGroup {
		return vs[0].([]ComputeClusterHostGroup)[vs[1].(int)]
	}).(ComputeClusterHostGroupOutput)
}

type ComputeClusterHostGroupMapOutput struct{ *pulumi.OutputState }

func (ComputeClusterHostGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ComputeClusterHostGroup)(nil))
}

func (o ComputeClusterHostGroupMapOutput) ToComputeClusterHostGroupMapOutput() ComputeClusterHostGroupMapOutput {
	return o
}

func (o ComputeClusterHostGroupMapOutput) ToComputeClusterHostGroupMapOutputWithContext(ctx context.Context) ComputeClusterHostGroupMapOutput {
	return o
}

func (o ComputeClusterHostGroupMapOutput) MapIndex(k pulumi.StringInput) ComputeClusterHostGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ComputeClusterHostGroup {
		return vs[0].(map[string]ComputeClusterHostGroup)[vs[1].(string)]
	}).(ComputeClusterHostGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterHostGroupInput)(nil)).Elem(), &ComputeClusterHostGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterHostGroupPtrInput)(nil)).Elem(), &ComputeClusterHostGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterHostGroupArrayInput)(nil)).Elem(), ComputeClusterHostGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterHostGroupMapInput)(nil)).Elem(), ComputeClusterHostGroupMap{})
	pulumi.RegisterOutputType(ComputeClusterHostGroupOutput{})
	pulumi.RegisterOutputType(ComputeClusterHostGroupPtrOutput{})
	pulumi.RegisterOutputType(ComputeClusterHostGroupArrayOutput{})
	pulumi.RegisterOutputType(ComputeClusterHostGroupMapOutput{})
}
