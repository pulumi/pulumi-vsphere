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

// The `ComputeClusterVmAffinityRule` resource can be used to
// manage virtual machine affinity rules in a cluster, either created by the
// `ComputeCluster` resource or looked up
// by the `ComputeCluster` data source.
//
// This rule can be used to tell a set of virtual machines to run together on the
// same host within a cluster. When configured, DRS will make a best effort to
// ensure that the virtual machines run on the same host, or prevent any operation
// that would keep that from happening, depending on the value of the
// `mandatory` flag.
//
// > An affinity rule can only be used to place virtual machines on the same
// _non-specific_ hosts. It cannot be used to pin virtual machines to a host.
// To enable this capability, use the
// `ComputeClusterVmHostRule`
// resource.
//
// > **NOTE:** This resource requires vCenter Server and is not available on
// direct ESXi host connections.
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
// $ pulumi import vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule vm_affinity_rule \
// ```
//
//	'{"compute_cluster_path": "/dc-01/host/cluster-01", \
//
//	"name": "vm-affinity-rule"}'
type ComputeClusterVmAffinityRule struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory pulumi.BoolPtrOutput `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds pulumi.StringArrayOutput `pulumi:"virtualMachineIds"`
}

// NewComputeClusterVmAffinityRule registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmAffinityRule(ctx *pulumi.Context,
	name string, args *ComputeClusterVmAffinityRuleArgs, opts ...pulumi.ResourceOption) (*ComputeClusterVmAffinityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	if args.VirtualMachineIds == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ComputeClusterVmAffinityRule
	err := ctx.RegisterResource("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterVmAffinityRule gets an existing ComputeClusterVmAffinityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmAffinityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterVmAffinityRuleState, opts ...pulumi.ResourceOption) (*ComputeClusterVmAffinityRule, error) {
	var resource ComputeClusterVmAffinityRule
	err := ctx.ReadResource("vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterVmAffinityRule resources.
type computeClusterVmAffinityRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

type ComputeClusterVmAffinityRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmAffinityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmAffinityRuleState)(nil)).Elem()
}

type computeClusterVmAffinityRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

// The set of arguments for constructing a ComputeClusterVmAffinityRule resource.
type ComputeClusterVmAffinityRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	//
	// > **NOTE:** The namespace for rule names on this resource (defined by the
	// `name` argument) is shared with all rules in the cluster - consider
	// this when naming your rules.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on the same host together.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmAffinityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmAffinityRuleArgs)(nil)).Elem()
}

type ComputeClusterVmAffinityRuleInput interface {
	pulumi.Input

	ToComputeClusterVmAffinityRuleOutput() ComputeClusterVmAffinityRuleOutput
	ToComputeClusterVmAffinityRuleOutputWithContext(ctx context.Context) ComputeClusterVmAffinityRuleOutput
}

func (*ComputeClusterVmAffinityRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmAffinityRule)(nil)).Elem()
}

func (i *ComputeClusterVmAffinityRule) ToComputeClusterVmAffinityRuleOutput() ComputeClusterVmAffinityRuleOutput {
	return i.ToComputeClusterVmAffinityRuleOutputWithContext(context.Background())
}

func (i *ComputeClusterVmAffinityRule) ToComputeClusterVmAffinityRuleOutputWithContext(ctx context.Context) ComputeClusterVmAffinityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmAffinityRuleOutput)
}

// ComputeClusterVmAffinityRuleArrayInput is an input type that accepts ComputeClusterVmAffinityRuleArray and ComputeClusterVmAffinityRuleArrayOutput values.
// You can construct a concrete instance of `ComputeClusterVmAffinityRuleArrayInput` via:
//
//	ComputeClusterVmAffinityRuleArray{ ComputeClusterVmAffinityRuleArgs{...} }
type ComputeClusterVmAffinityRuleArrayInput interface {
	pulumi.Input

	ToComputeClusterVmAffinityRuleArrayOutput() ComputeClusterVmAffinityRuleArrayOutput
	ToComputeClusterVmAffinityRuleArrayOutputWithContext(context.Context) ComputeClusterVmAffinityRuleArrayOutput
}

type ComputeClusterVmAffinityRuleArray []ComputeClusterVmAffinityRuleInput

func (ComputeClusterVmAffinityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeClusterVmAffinityRule)(nil)).Elem()
}

func (i ComputeClusterVmAffinityRuleArray) ToComputeClusterVmAffinityRuleArrayOutput() ComputeClusterVmAffinityRuleArrayOutput {
	return i.ToComputeClusterVmAffinityRuleArrayOutputWithContext(context.Background())
}

func (i ComputeClusterVmAffinityRuleArray) ToComputeClusterVmAffinityRuleArrayOutputWithContext(ctx context.Context) ComputeClusterVmAffinityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmAffinityRuleArrayOutput)
}

// ComputeClusterVmAffinityRuleMapInput is an input type that accepts ComputeClusterVmAffinityRuleMap and ComputeClusterVmAffinityRuleMapOutput values.
// You can construct a concrete instance of `ComputeClusterVmAffinityRuleMapInput` via:
//
//	ComputeClusterVmAffinityRuleMap{ "key": ComputeClusterVmAffinityRuleArgs{...} }
type ComputeClusterVmAffinityRuleMapInput interface {
	pulumi.Input

	ToComputeClusterVmAffinityRuleMapOutput() ComputeClusterVmAffinityRuleMapOutput
	ToComputeClusterVmAffinityRuleMapOutputWithContext(context.Context) ComputeClusterVmAffinityRuleMapOutput
}

type ComputeClusterVmAffinityRuleMap map[string]ComputeClusterVmAffinityRuleInput

func (ComputeClusterVmAffinityRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeClusterVmAffinityRule)(nil)).Elem()
}

func (i ComputeClusterVmAffinityRuleMap) ToComputeClusterVmAffinityRuleMapOutput() ComputeClusterVmAffinityRuleMapOutput {
	return i.ToComputeClusterVmAffinityRuleMapOutputWithContext(context.Background())
}

func (i ComputeClusterVmAffinityRuleMap) ToComputeClusterVmAffinityRuleMapOutputWithContext(ctx context.Context) ComputeClusterVmAffinityRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmAffinityRuleMapOutput)
}

type ComputeClusterVmAffinityRuleOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmAffinityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeClusterVmAffinityRule)(nil)).Elem()
}

func (o ComputeClusterVmAffinityRuleOutput) ToComputeClusterVmAffinityRuleOutput() ComputeClusterVmAffinityRuleOutput {
	return o
}

func (o ComputeClusterVmAffinityRuleOutput) ToComputeClusterVmAffinityRuleOutputWithContext(ctx context.Context) ComputeClusterVmAffinityRuleOutput {
	return o
}

// The managed object reference
// ID of the cluster to put the group in.  Forces a new
// resource if changed.
func (o ComputeClusterVmAffinityRuleOutput) ComputeClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeClusterVmAffinityRule) pulumi.StringOutput { return v.ComputeClusterId }).(pulumi.StringOutput)
}

// Enable this rule in the cluster. Default: `true`.
func (o ComputeClusterVmAffinityRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeClusterVmAffinityRule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// When this value is `true`, prevents any virtual
// machine operations that may violate this rule. Default: `false`.
//
// > **NOTE:** The namespace for rule names on this resource (defined by the
// `name` argument) is shared with all rules in the cluster - consider
// this when naming your rules.
func (o ComputeClusterVmAffinityRuleOutput) Mandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputeClusterVmAffinityRule) pulumi.BoolPtrOutput { return v.Mandatory }).(pulumi.BoolPtrOutput)
}

// The name of the rule. This must be unique in the cluster.
func (o ComputeClusterVmAffinityRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeClusterVmAffinityRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The UUIDs of the virtual machines to run
// on the same host together.
func (o ComputeClusterVmAffinityRuleOutput) VirtualMachineIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeClusterVmAffinityRule) pulumi.StringArrayOutput { return v.VirtualMachineIds }).(pulumi.StringArrayOutput)
}

type ComputeClusterVmAffinityRuleArrayOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmAffinityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ComputeClusterVmAffinityRule)(nil)).Elem()
}

func (o ComputeClusterVmAffinityRuleArrayOutput) ToComputeClusterVmAffinityRuleArrayOutput() ComputeClusterVmAffinityRuleArrayOutput {
	return o
}

func (o ComputeClusterVmAffinityRuleArrayOutput) ToComputeClusterVmAffinityRuleArrayOutputWithContext(ctx context.Context) ComputeClusterVmAffinityRuleArrayOutput {
	return o
}

func (o ComputeClusterVmAffinityRuleArrayOutput) Index(i pulumi.IntInput) ComputeClusterVmAffinityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ComputeClusterVmAffinityRule {
		return vs[0].([]*ComputeClusterVmAffinityRule)[vs[1].(int)]
	}).(ComputeClusterVmAffinityRuleOutput)
}

type ComputeClusterVmAffinityRuleMapOutput struct{ *pulumi.OutputState }

func (ComputeClusterVmAffinityRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ComputeClusterVmAffinityRule)(nil)).Elem()
}

func (o ComputeClusterVmAffinityRuleMapOutput) ToComputeClusterVmAffinityRuleMapOutput() ComputeClusterVmAffinityRuleMapOutput {
	return o
}

func (o ComputeClusterVmAffinityRuleMapOutput) ToComputeClusterVmAffinityRuleMapOutputWithContext(ctx context.Context) ComputeClusterVmAffinityRuleMapOutput {
	return o
}

func (o ComputeClusterVmAffinityRuleMapOutput) MapIndex(k pulumi.StringInput) ComputeClusterVmAffinityRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ComputeClusterVmAffinityRule {
		return vs[0].(map[string]*ComputeClusterVmAffinityRule)[vs[1].(string)]
	}).(ComputeClusterVmAffinityRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmAffinityRuleInput)(nil)).Elem(), &ComputeClusterVmAffinityRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmAffinityRuleArrayInput)(nil)).Elem(), ComputeClusterVmAffinityRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeClusterVmAffinityRuleMapInput)(nil)).Elem(), ComputeClusterVmAffinityRuleMap{})
	pulumi.RegisterOutputType(ComputeClusterVmAffinityRuleOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmAffinityRuleArrayOutput{})
	pulumi.RegisterOutputType(ComputeClusterVmAffinityRuleMapOutput{})
}
