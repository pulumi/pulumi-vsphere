// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type DatastoreClusterVmAntiAffinityRule struct {
	pulumi.CustomResourceState

	// The managed object reference
	// ID of the datastore cluster to put the group in.  Forces
	// a new resource if changed.
	DatastoreClusterId pulumi.StringOutput `pulumi:"datastoreClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrOutput `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on different datastores from each other.
	VirtualMachineIds pulumi.StringArrayOutput `pulumi:"virtualMachineIds"`
}

// NewDatastoreClusterVmAntiAffinityRule registers a new resource with the given unique name, arguments, and options.
func NewDatastoreClusterVmAntiAffinityRule(ctx *pulumi.Context,
	name string, args *DatastoreClusterVmAntiAffinityRuleArgs, opts ...pulumi.ResourceOption) (*DatastoreClusterVmAntiAffinityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatastoreClusterId == nil {
		return nil, errors.New("invalid value for required argument 'DatastoreClusterId'")
	}
	if args.VirtualMachineIds == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineIds'")
	}
	var resource DatastoreClusterVmAntiAffinityRule
	err := ctx.RegisterResource("vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatastoreClusterVmAntiAffinityRule gets an existing DatastoreClusterVmAntiAffinityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatastoreClusterVmAntiAffinityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatastoreClusterVmAntiAffinityRuleState, opts ...pulumi.ResourceOption) (*DatastoreClusterVmAntiAffinityRule, error) {
	var resource DatastoreClusterVmAntiAffinityRule
	err := ctx.ReadResource("vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatastoreClusterVmAntiAffinityRule resources.
type datastoreClusterVmAntiAffinityRuleState struct {
	// The managed object reference
	// ID of the datastore cluster to put the group in.  Forces
	// a new resource if changed.
	DatastoreClusterId *string `pulumi:"datastoreClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on different datastores from each other.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

type DatastoreClusterVmAntiAffinityRuleState struct {
	// The managed object reference
	// ID of the datastore cluster to put the group in.  Forces
	// a new resource if changed.
	DatastoreClusterId pulumi.StringPtrInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on different datastores from each other.
	VirtualMachineIds pulumi.StringArrayInput
}

func (DatastoreClusterVmAntiAffinityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreClusterVmAntiAffinityRuleState)(nil)).Elem()
}

type datastoreClusterVmAntiAffinityRuleArgs struct {
	// The managed object reference
	// ID of the datastore cluster to put the group in.  Forces
	// a new resource if changed.
	DatastoreClusterId string `pulumi:"datastoreClusterId"`
	// Enable this rule in the cluster. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory *bool `pulumi:"mandatory"`
	// The name of the rule. This must be unique in the cluster.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines to run
	// on different datastores from each other.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

// The set of arguments for constructing a DatastoreClusterVmAntiAffinityRule resource.
type DatastoreClusterVmAntiAffinityRuleArgs struct {
	// The managed object reference
	// ID of the datastore cluster to put the group in.  Forces
	// a new resource if changed.
	DatastoreClusterId pulumi.StringInput
	// Enable this rule in the cluster. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory pulumi.BoolPtrInput
	// The name of the rule. This must be unique in the cluster.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines to run
	// on different datastores from each other.
	VirtualMachineIds pulumi.StringArrayInput
}

func (DatastoreClusterVmAntiAffinityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datastoreClusterVmAntiAffinityRuleArgs)(nil)).Elem()
}

type DatastoreClusterVmAntiAffinityRuleInput interface {
	pulumi.Input

	ToDatastoreClusterVmAntiAffinityRuleOutput() DatastoreClusterVmAntiAffinityRuleOutput
	ToDatastoreClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleOutput
}

func (DatastoreClusterVmAntiAffinityRule) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreClusterVmAntiAffinityRule)(nil)).Elem()
}

func (i DatastoreClusterVmAntiAffinityRule) ToDatastoreClusterVmAntiAffinityRuleOutput() DatastoreClusterVmAntiAffinityRuleOutput {
	return i.ToDatastoreClusterVmAntiAffinityRuleOutputWithContext(context.Background())
}

func (i DatastoreClusterVmAntiAffinityRule) ToDatastoreClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreClusterVmAntiAffinityRuleOutput)
}

type DatastoreClusterVmAntiAffinityRuleOutput struct {
	*pulumi.OutputState
}

func (DatastoreClusterVmAntiAffinityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatastoreClusterVmAntiAffinityRuleOutput)(nil)).Elem()
}

func (o DatastoreClusterVmAntiAffinityRuleOutput) ToDatastoreClusterVmAntiAffinityRuleOutput() DatastoreClusterVmAntiAffinityRuleOutput {
	return o
}

func (o DatastoreClusterVmAntiAffinityRuleOutput) ToDatastoreClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatastoreClusterVmAntiAffinityRuleOutput{})
}
