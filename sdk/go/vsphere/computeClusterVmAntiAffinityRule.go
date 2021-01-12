// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ComputeClusterVmAntiAffinityRule struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrOutput `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on hosts different from each other.
	VirtualMachineIds pulumi.StringArrayOutput `pulumi:"virtualMachineIds"`
}

// NewComputeClusterVmAntiAffinityRule registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmAntiAffinityRule(ctx *pulumi.Context,
	name string, args *ComputeClusterVmAntiAffinityRuleArgs, opts ...pulumi.ResourceOption) (*ComputeClusterVmAntiAffinityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ComputeClusterId'")
	}
	if args.VirtualMachineIds == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineIds'")
	}
	var resource ComputeClusterVmAntiAffinityRule
	err := ctx.RegisterResource("vsphere:index/computeClusterVmAntiAffinityRule:ComputeClusterVmAntiAffinityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterVmAntiAffinityRule gets an existing ComputeClusterVmAntiAffinityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmAntiAffinityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterVmAntiAffinityRuleState, opts ...pulumi.ResourceOption) (*ComputeClusterVmAntiAffinityRule, error) {
	var resource ComputeClusterVmAntiAffinityRule
	err := ctx.ReadResource("vsphere:index/computeClusterVmAntiAffinityRule:ComputeClusterVmAntiAffinityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterVmAntiAffinityRule resources.
type computeClusterVmAntiAffinityRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on hosts different from each other.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

type ComputeClusterVmAntiAffinityRuleState struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on hosts different from each other.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmAntiAffinityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmAntiAffinityRuleState)(nil)).Elem()
}

type computeClusterVmAntiAffinityRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on hosts different from each other.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

// The set of arguments for constructing a ComputeClusterVmAntiAffinityRule resource.
type ComputeClusterVmAntiAffinityRuleArgs struct {
	// The managed object reference
	// ID of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on hosts different from each other.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmAntiAffinityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmAntiAffinityRuleArgs)(nil)).Elem()
}

type ComputeClusterVmAntiAffinityRuleInput interface {
	pulumi.Input

	ToComputeClusterVmAntiAffinityRuleOutput() ComputeClusterVmAntiAffinityRuleOutput
	ToComputeClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) ComputeClusterVmAntiAffinityRuleOutput
}

func (ComputeClusterVmAntiAffinityRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeClusterVmAntiAffinityRule)(nil)).Elem()
}

func (i ComputeClusterVmAntiAffinityRule) ToComputeClusterVmAntiAffinityRuleOutput() ComputeClusterVmAntiAffinityRuleOutput {
	return i.ToComputeClusterVmAntiAffinityRuleOutputWithContext(context.Background())
}

func (i ComputeClusterVmAntiAffinityRule) ToComputeClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) ComputeClusterVmAntiAffinityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeClusterVmAntiAffinityRuleOutput)
}

type ComputeClusterVmAntiAffinityRuleOutput struct {
	*pulumi.OutputState
}

func (ComputeClusterVmAntiAffinityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ComputeClusterVmAntiAffinityRuleOutput)(nil)).Elem()
}

func (o ComputeClusterVmAntiAffinityRuleOutput) ToComputeClusterVmAntiAffinityRuleOutput() ComputeClusterVmAntiAffinityRuleOutput {
	return o
}

func (o ComputeClusterVmAntiAffinityRuleOutput) ToComputeClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) ComputeClusterVmAntiAffinityRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ComputeClusterVmAntiAffinityRuleOutput{})
}
