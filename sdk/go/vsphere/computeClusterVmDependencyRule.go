// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ComputeClusterVmDependencyRule struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// The name of the VM group that this
	// rule depends on. The VMs defined in the group specified by
	// `vmGroupName` will not be started until the VMs in this
	// group are started.
	DependencyVmGroupName pulumi.StringOutput `pulumi:"dependencyVmGroupName"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrOutput `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the
	// cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the VM group that is the subject of
	// this rule. The VMs defined in this group will not be started until the VMs in
	// the group specified by
	// `dependencyVmGroupName` are started.
	VmGroupName pulumi.StringOutput `pulumi:"vmGroupName"`
}

// NewComputeClusterVmDependencyRule registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmDependencyRule(ctx *pulumi.Context,
	name string, args *ComputeClusterVmDependencyRuleArgs, opts ...pulumi.ResourceOption) (*ComputeClusterVmDependencyRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	if args.DependencyVmGroupName == nil {
		return nil, errors.New("invalid value for required argument 'DependencyVmGroupName'")
	}
	if args.VmGroupName == nil {
		return nil, errors.New("invalid value for required argument 'VmGroupName'")
	}
	var resource ComputeClusterVmDependencyRule
	err := ctx.RegisterResource("vsphere:index/computeClusterVmDependencyRule:ComputeClusterVmDependencyRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterVmDependencyRule gets an existing ComputeClusterVmDependencyRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmDependencyRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterVmDependencyRuleState, opts ...pulumi.ResourceOption) (*ComputeClusterVmDependencyRule, error) {
	var resource ComputeClusterVmDependencyRule
	err := ctx.ReadResource("vsphere:index/computeClusterVmDependencyRule:ComputeClusterVmDependencyRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterVmDependencyRule resources.
type computeClusterVmDependencyRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// The name of the VM group that this
	// rule depends on. The VMs defined in the group specified by
	// `vmGroupName` will not be started until the VMs in this
	// group are started.
	DependencyVmGroupName *string `pulumi:"dependencyVmGroupName"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the
	// cluster.
	Name *string `pulumi:"name"`
	// The name of the VM group that is the subject of
	// this rule. The VMs defined in this group will not be started until the VMs in
	// the group specified by
	// `dependencyVmGroupName` are started.
	VmGroupName *string `pulumi:"vmGroupName"`
}

type ComputeClusterVmDependencyRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// The name of the VM group that this
	// rule depends on. The VMs defined in the group specified by
	// `vmGroupName` will not be started until the VMs in this
	// group are started.
	DependencyVmGroupName pulumi.StringPtrInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the
	// cluster.
	Name pulumi.StringPtrInput
	// The name of the VM group that is the subject of
	// this rule. The VMs defined in this group will not be started until the VMs in
	// the group specified by
	// `dependencyVmGroupName` are started.
	VmGroupName pulumi.StringPtrInput
}

func (ComputeClusterVmDependencyRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmDependencyRuleState)(nil)).Elem()
}

type computeClusterVmDependencyRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// The name of the VM group that this
	// rule depends on. The VMs defined in the group specified by
	// `vmGroupName` will not be started until the VMs in this
	// group are started.
	DependencyVmGroupName string `pulumi:"dependencyVmGroupName"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the
	// cluster.
	Name *string `pulumi:"name"`
	// The name of the VM group that is the subject of
	// this rule. The VMs defined in this group will not be started until the VMs in
	// the group specified by
	// `dependencyVmGroupName` are started.
	VmGroupName string `pulumi:"vmGroupName"`
}

// The set of arguments for constructing a ComputeClusterVmDependencyRule resource.
type ComputeClusterVmDependencyRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// The name of the VM group that this
	// rule depends on. The VMs defined in the group specified by
	// `vmGroupName` will not be started until the VMs in this
	// group are started.
	DependencyVmGroupName pulumi.StringInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the
	// cluster.
	Name pulumi.StringPtrInput
	// The name of the VM group that is the subject of
	// this rule. The VMs defined in this group will not be started until the VMs in
	// the group specified by
	// `dependencyVmGroupName` are started.
	VmGroupName pulumi.StringInput
}

func (ComputeClusterVmDependencyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmDependencyRuleArgs)(nil)).Elem()
}

type ComputeClusterVmDependencyRuleInput interface {
	pulumi.Input

	ToComputeClusterVmDependencyRuleOutput() ComputeClusterVmDependencyRuleOutput
	ToComputeClusterVmDependencyRuleOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRuleOutput
}

func (*ComputeClusterVmDependencyRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeClusterVmDependencyRule)(nil))
}

func (i *ComputeClusterVmDependencyRule) ToComputeClusterVmDependencyRuleOutput() ComputeClusterVmDependencyRuleOutput {
	return i.ToComputeClusterVmDependencyRuleOutputWithContext(context.Background())
}

func (i *ComputeClusterVmDependencyRule) ToComputeClusterVmDependencyRuleOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmDependencyRuleOutput)
}

func (i *ComputeClusterVmDependencyRule) ToComputeClusterVmDependencyRulePtrOutput() ComputeClusterVmDependencyRulePtrOutput {
	return i.ToComputeClusterVmDependencyRulePtrOutputWithContext(context.Background())
}

func (i *ComputeClusterVmDependencyRule) ToComputeClusterVmDependencyRulePtrOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmDependencyRulePtrOutput)
}

type ComputeClusterVmDependencyRulePtrInput interface {
	pulumi.Input

	ToComputeClusterVmDependencyRulePtrOutput() ComputeClusterVmDependencyRulePtrOutput
	ToComputeClusterVmDependencyRulePtrOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRulePtrOutput
}

type computeClusterVmDependencyRulePtrType ComputeClusterVmDependencyRuleArgs

func (*computeClusterVmDependencyRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmDependencyRule)(nil))
}

func (i *computeClusterVmDependencyRulePtrType) ToComputeClusterVmDependencyRulePtrOutput() ComputeClusterVmDependencyRulePtrOutput {
	return i.ToComputeClusterVmDependencyRulePtrOutputWithContext(context.Background())
}

func (i *computeClusterVmDependencyRulePtrType) ToComputeClusterVmDependencyRulePtrOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmDependencyRulePtrOutput)
}

// ComputeClusterVmDependencyRuleArrayInput is an input type that accepts ComputeClusterVmDependencyRuleArray and ComputeClusterVmDependencyRuleArrayOutput values.
// You can construct a concrete instance of `ComputeClusterVmDependencyRuleArrayInput` via:
//
//          ComputeClusterVmDependencyRuleArray{ ComputeClusterVmDependencyRuleArgs{...} }
type ComputeClusterVmDependencyRuleArrayInput interface {
	pulumi.Input

	ToComputeClusterVmDependencyRuleArrayOutput() ComputeClusterVmDependencyRuleArrayOutput
	ToComputeClusterVmDependencyRuleArrayOutputWithContext(context.Context) ComputeClusterVmDependencyRuleArrayOutput
}

type ComputeClusterVmDependencyRuleArray []ComputeClusterVmDependencyRuleInput

func (ComputeClusterVmDependencyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ComputeClusterVmDependencyRule)(nil))
}

func (i ComputeClusterVmDependencyRuleArray) ToComputeClusterVmDependencyRuleArrayOutput() ComputeClusterVmDependencyRuleArrayOutput {
	return i.ToComputeClusterVmDependencyRuleArrayOutputWithContext(context.Background())
}

func (i ComputeClusterVmDependencyRuleArray) ToComputeClusterVmDependencyRuleArrayOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmDependencyRuleArrayOutput)
}

// ComputeClusterVmDependencyRuleMapInput is an input type that accepts ComputeClusterVmDependencyRuleMap and ComputeClusterVmDependencyRuleMapOutput values.
// You can construct a concrete instance of `ComputeClusterVmDependencyRuleMapInput` via:
//
//          ComputeClusterVmDependencyRuleMap{ "key": ComputeClusterVmDependencyRuleArgs{...} }
type ComputeClusterVmDependencyRuleMapInput interface {
	pulumi.Input

	ToComputeClusterVmDependencyRuleMapOutput() ComputeClusterVmDependencyRuleMapOutput
	ToComputeClusterVmDependencyRuleMapOutputWithContext(context.Context) ComputeClusterVmDependencyRuleMapOutput
}

type ComputeClusterVmDependencyRuleMap map[string]ComputeClusterVmDependencyRuleInput

func (ComputeClusterVmDependencyRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ComputeClusterVmDependencyRule)(nil))
}

func (i ComputeClusterVmDependencyRuleMap) ToComputeClusterVmDependencyRuleMapOutput() ComputeClusterVmDependencyRuleMapOutput {
	return i.ToComputeClusterVmDependencyRuleMapOutputWithContext(context.Background())
}

func (i ComputeClusterVmDependencyRuleMap) ToComputeClusterVmDependencyRuleMapOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmDependencyRuleMapOutput)
}

type ComputeClusterVmDependencyRuleOutput struct {
	*pulumi.OutputState
}

func (ComputeClusterVmDependencyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeClusterVmDependencyRule)(nil))
}

func (o ComputeClusterVmDependencyRuleOutput) ToComputeClusterVmDependencyRuleOutput() ComputeClusterVmDependencyRuleOutput {
	return o
}

func (o ComputeClusterVmDependencyRuleOutput) ToComputeClusterVmDependencyRuleOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRuleOutput {
	return o
}

func (o ComputeClusterVmDependencyRuleOutput) ToComputeClusterVmDependencyRulePtrOutput() ComputeClusterVmDependencyRulePtrOutput {
	return o.ToComputeClusterVmDependencyRulePtrOutputWithContext(context.Background())
}

func (o ComputeClusterVmDependencyRuleOutput) ToComputeClusterVmDependencyRulePtrOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRulePtrOutput {
	return o.ApplyT(func(v ComputeClusterVmDependencyRule) *ComputeClusterVmDependencyRule {
		return &v
	}).(ComputeClusterVmDependencyRulePtrOutput)
}

type ComputeClusterVmDependencyRulePtrOutput struct {
	*pulumi.OutputState
}

func (ComputeClusterVmDependencyRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmDependencyRule)(nil))
}

func (o ComputeClusterVmDependencyRulePtrOutput) ToComputeClusterVmDependencyRulePtrOutput() ComputeClusterVmDependencyRulePtrOutput {
	return o
}

func (o ComputeClusterVmDependencyRulePtrOutput) ToComputeClusterVmDependencyRulePtrOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRulePtrOutput {
	return o
}

type ComputeClusterVmDependencyRuleArrayOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmDependencyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ComputeClusterVmDependencyRule)(nil))
}

func (o ComputeClusterVmDependencyRuleArrayOutput) ToComputeClusterVmDependencyRuleArrayOutput() ComputeClusterVmDependencyRuleArrayOutput {
	return o
}

func (o ComputeClusterVmDependencyRuleArrayOutput) ToComputeClusterVmDependencyRuleArrayOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRuleArrayOutput {
	return o
}

func (o ComputeClusterVmDependencyRuleArrayOutput) Index(i pulumi.IntInput) ComputeClusterVmDependencyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ComputeClusterVmDependencyRule {
		return vs[0].([]ComputeClusterVmDependencyRule)[vs[1].(int)]
	}).(ComputeClusterVmDependencyRuleOutput)
}

type ComputeClusterVmDependencyRuleMapOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmDependencyRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ComputeClusterVmDependencyRule)(nil))
}

func (o ComputeClusterVmDependencyRuleMapOutput) ToComputeClusterVmDependencyRuleMapOutput() ComputeClusterVmDependencyRuleMapOutput {
	return o
}

func (o ComputeClusterVmDependencyRuleMapOutput) ToComputeClusterVmDependencyRuleMapOutputWithContext(ctx context.Context) ComputeClusterVmDependencyRuleMapOutput {
	return o
}

func (o ComputeClusterVmDependencyRuleMapOutput) MapIndex(k pulumi.StringInput) ComputeClusterVmDependencyRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ComputeClusterVmDependencyRule {
		return vs[0].(map[string]ComputeClusterVmDependencyRule)[vs[1].(string)]
	}).(ComputeClusterVmDependencyRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(ComputeClusterVmDependencyRuleOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmDependencyRulePtrOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmDependencyRuleArrayOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmDependencyRuleMapOutput{})
}
