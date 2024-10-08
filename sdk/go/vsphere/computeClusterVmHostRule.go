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

// The `ComputeClusterVmHostRule` resource can be used to manage
// VM-to-host rules in a cluster, either created by the
// `ComputeCluster` resource or looked up
// by the `ComputeCluster` data source.
//
// This resource can create both _affinity rules_, where virtual machines run on
// specified hosts, or _anti-affinity_ rules, where virtual machines run on hosts
// outside of the ones specified in the rule. Virtual machines and hosts are
// supplied via groups, which can be managed via the
// `ComputeClusterVmGroup` and
// `ComputeClusterHostGroup`
// resources.
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// ## Import
//
// # An existing rule can be imported into this resource by supplying
//
// both the path to the cluster, and the name the rule. If the name or cluster is
//
// not found, or if the rule is of a different type, an error will be returned. An
//
// example is below:
//
// ```sh
// $ pulumi import vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule cluster_vm_host_rule \
// ```
//
//	'{"compute_cluster_path": "/dc1/host/cluster1", \
//
//	"name": "pulumi-test-cluster-vm-host-rule"}'
type ComputeClusterVmHostRule struct {
	pulumi.CustomResourceState

	// When this field is used, the virtual
	// machines defined in `vmGroupName` will be run on the
	// hosts defined in this host group.
	AffinityHostGroupName pulumi.StringPtrOutput `pulumi:"affinityHostGroupName"`
	// When this field is used, the
	// virtual machines defined in `vmGroupName` will _not_ be
	// run on the hosts defined in this host group.
	AntiAffinityHostGroupName pulumi.StringPtrOutput `pulumi:"antiAffinityHostGroupName"`
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** One of `affinityHostGroupName` or
	// `antiAffinityHostGroupName` must be
	// defined, but not both.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory pulumi.BoolPtrOutput `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the
	// cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the virtual machine group to use
	// with this rule.
	VmGroupName pulumi.StringOutput `pulumi:"vmGroupName"`
}

// NewComputeClusterVmHostRule registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmHostRule(ctx *pulumi.Context,
	name string, args *ComputeClusterVmHostRuleArgs, opts ...pulumi.ResourceOption) (*ComputeClusterVmHostRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	if args.VmGroupName == nil {
		return nil, errors.New("invalid value for required argument 'VmGroupName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ComputeClusterVmHostRule
	err := ctx.RegisterResource("vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterVmHostRule gets an existing ComputeClusterVmHostRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmHostRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterVmHostRuleState, opts ...pulumi.ResourceOption) (*ComputeClusterVmHostRule, error) {
	var resource ComputeClusterVmHostRule
	err := ctx.ReadResource("vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterVmHostRule resources.
type computeClusterVmHostRuleState struct {
	// When this field is used, the virtual
	// machines defined in `vmGroupName` will be run on the
	// hosts defined in this host group.
	AffinityHostGroupName *string `pulumi:"affinityHostGroupName"`
	// When this field is used, the
	// virtual machines defined in `vmGroupName` will _not_ be
	// run on the hosts defined in this host group.
	AntiAffinityHostGroupName *string `pulumi:"antiAffinityHostGroupName"`
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** One of `affinityHostGroupName` or
	// `antiAffinityHostGroupName` must be
	// defined, but not both.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the
	// cluster.
	Name *string `pulumi:"name"`
	// The name of the virtual machine group to use
	// with this rule.
	VmGroupName *string `pulumi:"vmGroupName"`
}

type ComputeClusterVmHostRuleState struct {
	// When this field is used, the virtual
	// machines defined in `vmGroupName` will be run on the
	// hosts defined in this host group.
	AffinityHostGroupName pulumi.StringPtrInput
	// When this field is used, the
	// virtual machines defined in `vmGroupName` will _not_ be
	// run on the hosts defined in this host group.
	AntiAffinityHostGroupName pulumi.StringPtrInput
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** One of `affinityHostGroupName` or
	// `antiAffinityHostGroupName` must be
	// defined, but not both.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the
	// cluster.
	Name pulumi.StringPtrInput
	// The name of the virtual machine group to use
	// with this rule.
	VmGroupName pulumi.StringPtrInput
}

func (ComputeClusterVmHostRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmHostRuleState)(nil)).Elem()
}

type computeClusterVmHostRuleArgs struct {
	// When this field is used, the virtual
	// machines defined in `vmGroupName` will be run on the
	// hosts defined in this host group.
	AffinityHostGroupName *string `pulumi:"affinityHostGroupName"`
	// When this field is used, the
	// virtual machines defined in `vmGroupName` will _not_ be
	// run on the hosts defined in this host group.
	AntiAffinityHostGroupName *string `pulumi:"antiAffinityHostGroupName"`
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** One of `affinityHostGroupName` or
	// `antiAffinityHostGroupName` must be
	// defined, but not both.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the
	// cluster.
	Name *string `pulumi:"name"`
	// The name of the virtual machine group to use
	// with this rule.
	VmGroupName string `pulumi:"vmGroupName"`
}

// The set of arguments for constructing a ComputeClusterVmHostRule resource.
type ComputeClusterVmHostRuleArgs struct {
	// When this field is used, the virtual
	// machines defined in `vmGroupName` will be run on the
	// hosts defined in this host group.
	AffinityHostGroupName pulumi.StringPtrInput
	// When this field is used, the
	// virtual machines defined in `vmGroupName` will _not_ be
	// run on the hosts defined in this host group.
	AntiAffinityHostGroupName pulumi.StringPtrInput
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** One of `affinityHostGroupName` or
	// `antiAffinityHostGroupName` must be
	// defined, but not both.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the
	// cluster.
	Name pulumi.StringPtrInput
	// The name of the virtual machine group to use
	// with this rule.
	VmGroupName pulumi.StringInput
}

func (ComputeClusterVmHostRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmHostRuleArgs)(nil)).Elem()
}

type ComputeClusterVmHostRuleInput interface {
	pulumi.Input

	ToComputeClusterVmHostRuleOutput() ComputeClusterVmHostRuleOutput
	ToComputeClusterVmHostRuleOutputWithContext(ctx context.Context) ComputeClusterVmHostRuleOutput
}

func (*ComputeClusterVmHostRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmHostRule)(nil)).Elem()
}

func (i *ComputeClusterVmHostRule) ToComputeClusterVmHostRuleOutput() ComputeClusterVmHostRuleOutput {
	return i.ToComputeClusterVmHostRuleOutputWithContext(context.Background())
}

func (i *ComputeClusterVmHostRule) ToComputeClusterVmHostRuleOutputWithContext(ctx context.Context) ComputeClusterVmHostRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmHostRuleOutput)
}

// ComputeClusterVmHostRuleArrayInput is an input type that accepts ComputeClusterVmHostRuleArray and ComputeClusterVmHostRuleArrayOutput values.
// You can construct a concrete instance of `ComputeClusterVmHostRuleArrayInput` via:
//
//	ComputeClusterVmHostRuleArray{ ComputeClusterVmHostRuleArgs{...} }
type ComputeClusterVmHostRuleArrayInput interface {
	pulumi.Input

	ToComputeClusterVmHostRuleArrayOutput() ComputeClusterVmHostRuleArrayOutput
	ToComputeClusterVmHostRuleArrayOutputWithContext(context.Context) ComputeClusterVmHostRuleArrayOutput
}

type ComputeClusterVmHostRuleArray []ComputeClusterVmHostRuleInput

func (ComputeClusterVmHostRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeClusterVmHostRule)(nil)).Elem()
}

func (i ComputeClusterVmHostRuleArray) ToComputeClusterVmHostRuleArrayOutput() ComputeClusterVmHostRuleArrayOutput {
	return i.ToComputeClusterVmHostRuleArrayOutputWithContext(context.Background())
}

func (i ComputeClusterVmHostRuleArray) ToComputeClusterVmHostRuleArrayOutputWithContext(ctx context.Context) ComputeClusterVmHostRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmHostRuleArrayOutput)
}

// ComputeClusterVmHostRuleMapInput is an input type that accepts ComputeClusterVmHostRuleMap and ComputeClusterVmHostRuleMapOutput values.
// You can construct a concrete instance of `ComputeClusterVmHostRuleMapInput` via:
//
//	ComputeClusterVmHostRuleMap{ "key": ComputeClusterVmHostRuleArgs{...} }
type ComputeClusterVmHostRuleMapInput interface {
	pulumi.Input

	ToComputeClusterVmHostRuleMapOutput() ComputeClusterVmHostRuleMapOutput
	ToComputeClusterVmHostRuleMapOutputWithContext(context.Context) ComputeClusterVmHostRuleMapOutput
}

type ComputeClusterVmHostRuleMap map[string]ComputeClusterVmHostRuleInput

func (ComputeClusterVmHostRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeClusterVmHostRule)(nil)).Elem()
}

func (i ComputeClusterVmHostRuleMap) ToComputeClusterVmHostRuleMapOutput() ComputeClusterVmHostRuleMapOutput {
	return i.ToComputeClusterVmHostRuleMapOutputWithContext(context.Background())
}

func (i ComputeClusterVmHostRuleMap) ToComputeClusterVmHostRuleMapOutputWithContext(ctx context.Context) ComputeClusterVmHostRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmHostRuleMapOutput)
}

type ComputeClusterVmHostRuleOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmHostRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmHostRule)(nil)).Elem()
}

func (o ComputeClusterVmHostRuleOutput) ToComputeClusterVmHostRuleOutput() ComputeClusterVmHostRuleOutput {
	return o
}

func (o ComputeClusterVmHostRuleOutput) ToComputeClusterVmHostRuleOutputWithContext(ctx context.Context) ComputeClusterVmHostRuleOutput {
	return o
}

// When this field is used, the virtual
// machines defined in `vmGroupName` will be run on the
// hosts defined in this host group.
func (o ComputeClusterVmHostRuleOutput) AffinityHostGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeClusterVmHostRule) pulumi.StringPtrOutput { return v.AffinityHostGroupName }).(pulumi.StringPtrOutput)
}

// When this field is used, the
// virtual machines defined in `vmGroupName` will _not_ be
// run on the hosts defined in this host group.
func (o ComputeClusterVmHostRuleOutput) AntiAffinityHostGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeClusterVmHostRule) pulumi.StringPtrOutput { return v.AntiAffinityHostGroupName }).(pulumi.StringPtrOutput)
}

// The managed object reference
// ID of the cluster to put the group in.  Forces a new
// resource if changed.
func (o ComputeClusterVmHostRuleOutput) ComputeClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeClusterVmHostRule) pulumi.StringOutput { return v.ComputeClusterId }).(pulumi.StringOutput)
}

// Enable this rule in the cluster. Default: `true`.
func (o ComputeClusterVmHostRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeClusterVmHostRule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// When this value is `true`, prevents any virtual
// machine operations that may violate this rule. Default: `false`.
//
// > **NOTE:** One of `affinityHostGroupName` or
// `antiAffinityHostGroupName` must be
// defined, but not both.
//
// > **NOTE:** The namespace for rule names on this resource (defined by the
// `name` argument) is shared with all rules in the cluster - consider
// this when naming your rules.
func (o ComputeClusterVmHostRuleOutput) Mandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeClusterVmHostRule) pulumi.BoolPtrOutput { return v.Mandatory }).(pulumi.BoolPtrOutput)
}

// The name of the rule. This must be unique in the
// cluster.
func (o ComputeClusterVmHostRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeClusterVmHostRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the virtual machine group to use
// with this rule.
func (o ComputeClusterVmHostRuleOutput) VmGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeClusterVmHostRule) pulumi.StringOutput { return v.VmGroupName }).(pulumi.StringOutput)
}

type ComputeClusterVmHostRuleArrayOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmHostRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeClusterVmHostRule)(nil)).Elem()
}

func (o ComputeClusterVmHostRuleArrayOutput) ToComputeClusterVmHostRuleArrayOutput() ComputeClusterVmHostRuleArrayOutput {
	return o
}

func (o ComputeClusterVmHostRuleArrayOutput) ToComputeClusterVmHostRuleArrayOutputWithContext(ctx context.Context) ComputeClusterVmHostRuleArrayOutput {
	return o
}

func (o ComputeClusterVmHostRuleArrayOutput) Index(i pulumi.IntInput) ComputeClusterVmHostRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ComputeClusterVmHostRule {
		return vs[0].([]*ComputeClusterVmHostRule)[vs[1].(int)]
	}).(ComputeClusterVmHostRuleOutput)
}

type ComputeClusterVmHostRuleMapOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmHostRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeClusterVmHostRule)(nil)).Elem()
}

func (o ComputeClusterVmHostRuleMapOutput) ToComputeClusterVmHostRuleMapOutput() ComputeClusterVmHostRuleMapOutput {
	return o
}

func (o ComputeClusterVmHostRuleMapOutput) ToComputeClusterVmHostRuleMapOutputWithContext(ctx context.Context) ComputeClusterVmHostRuleMapOutput {
	return o
}

func (o ComputeClusterVmHostRuleMapOutput) MapIndex(k pulumi.StringInput) ComputeClusterVmHostRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ComputeClusterVmHostRule {
		return vs[0].(map[string]*ComputeClusterVmHostRule)[vs[1].(string)]
	}).(ComputeClusterVmHostRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmHostRuleInput)(nil)).Elem(), &ComputeClusterVmHostRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmHostRuleArrayInput)(nil)).Elem(), ComputeClusterVmHostRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmHostRuleMapInput)(nil)).Elem(), ComputeClusterVmHostRuleMap{})
	pulumi.RegisterOutputType(ComputeClusterVmHostRuleOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmHostRuleArrayOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmHostRuleMapOutput{})
}
