// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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

func (*DatastoreClusterVmAntiAffinityRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastoreClusterVmAntiAffinityRule)(nil)).Elem()
}

func (i *DatastoreClusterVmAntiAffinityRule) ToDatastoreClusterVmAntiAffinityRuleOutput() DatastoreClusterVmAntiAffinityRuleOutput {
	return i.ToDatastoreClusterVmAntiAffinityRuleOutputWithContext(context.Background())
}

func (i *DatastoreClusterVmAntiAffinityRule) ToDatastoreClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreClusterVmAntiAffinityRuleOutput)
}

// DatastoreClusterVmAntiAffinityRuleArrayInput is an input type that accepts DatastoreClusterVmAntiAffinityRuleArray and DatastoreClusterVmAntiAffinityRuleArrayOutput values.
// You can construct a concrete instance of `DatastoreClusterVmAntiAffinityRuleArrayInput` via:
//
//	DatastoreClusterVmAntiAffinityRuleArray{ DatastoreClusterVmAntiAffinityRuleArgs{...} }
type DatastoreClusterVmAntiAffinityRuleArrayInput interface {
	pulumi.Input

	ToDatastoreClusterVmAntiAffinityRuleArrayOutput() DatastoreClusterVmAntiAffinityRuleArrayOutput
	ToDatastoreClusterVmAntiAffinityRuleArrayOutputWithContext(context.Context) DatastoreClusterVmAntiAffinityRuleArrayOutput
}

type DatastoreClusterVmAntiAffinityRuleArray []DatastoreClusterVmAntiAffinityRuleInput

func (DatastoreClusterVmAntiAffinityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatastoreClusterVmAntiAffinityRule)(nil)).Elem()
}

func (i DatastoreClusterVmAntiAffinityRuleArray) ToDatastoreClusterVmAntiAffinityRuleArrayOutput() DatastoreClusterVmAntiAffinityRuleArrayOutput {
	return i.ToDatastoreClusterVmAntiAffinityRuleArrayOutputWithContext(context.Background())
}

func (i DatastoreClusterVmAntiAffinityRuleArray) ToDatastoreClusterVmAntiAffinityRuleArrayOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreClusterVmAntiAffinityRuleArrayOutput)
}

// DatastoreClusterVmAntiAffinityRuleMapInput is an input type that accepts DatastoreClusterVmAntiAffinityRuleMap and DatastoreClusterVmAntiAffinityRuleMapOutput values.
// You can construct a concrete instance of `DatastoreClusterVmAntiAffinityRuleMapInput` via:
//
//	DatastoreClusterVmAntiAffinityRuleMap{ "key": DatastoreClusterVmAntiAffinityRuleArgs{...} }
type DatastoreClusterVmAntiAffinityRuleMapInput interface {
	pulumi.Input

	ToDatastoreClusterVmAntiAffinityRuleMapOutput() DatastoreClusterVmAntiAffinityRuleMapOutput
	ToDatastoreClusterVmAntiAffinityRuleMapOutputWithContext(context.Context) DatastoreClusterVmAntiAffinityRuleMapOutput
}

type DatastoreClusterVmAntiAffinityRuleMap map[string]DatastoreClusterVmAntiAffinityRuleInput

func (DatastoreClusterVmAntiAffinityRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatastoreClusterVmAntiAffinityRule)(nil)).Elem()
}

func (i DatastoreClusterVmAntiAffinityRuleMap) ToDatastoreClusterVmAntiAffinityRuleMapOutput() DatastoreClusterVmAntiAffinityRuleMapOutput {
	return i.ToDatastoreClusterVmAntiAffinityRuleMapOutputWithContext(context.Background())
}

func (i DatastoreClusterVmAntiAffinityRuleMap) ToDatastoreClusterVmAntiAffinityRuleMapOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatastoreClusterVmAntiAffinityRuleMapOutput)
}

type DatastoreClusterVmAntiAffinityRuleOutput struct{ *pulumi.OutputState }

func (DatastoreClusterVmAntiAffinityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatastoreClusterVmAntiAffinityRule)(nil)).Elem()
}

func (o DatastoreClusterVmAntiAffinityRuleOutput) ToDatastoreClusterVmAntiAffinityRuleOutput() DatastoreClusterVmAntiAffinityRuleOutput {
	return o
}

func (o DatastoreClusterVmAntiAffinityRuleOutput) ToDatastoreClusterVmAntiAffinityRuleOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleOutput {
	return o
}

type DatastoreClusterVmAntiAffinityRuleArrayOutput struct{ *pulumi.OutputState }

func (DatastoreClusterVmAntiAffinityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatastoreClusterVmAntiAffinityRule)(nil)).Elem()
}

func (o DatastoreClusterVmAntiAffinityRuleArrayOutput) ToDatastoreClusterVmAntiAffinityRuleArrayOutput() DatastoreClusterVmAntiAffinityRuleArrayOutput {
	return o
}

func (o DatastoreClusterVmAntiAffinityRuleArrayOutput) ToDatastoreClusterVmAntiAffinityRuleArrayOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleArrayOutput {
	return o
}

func (o DatastoreClusterVmAntiAffinityRuleArrayOutput) Index(i pulumi.IntInput) DatastoreClusterVmAntiAffinityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatastoreClusterVmAntiAffinityRule {
		return vs[0].([]*DatastoreClusterVmAntiAffinityRule)[vs[1].(int)]
	}).(DatastoreClusterVmAntiAffinityRuleOutput)
}

type DatastoreClusterVmAntiAffinityRuleMapOutput struct{ *pulumi.OutputState }

func (DatastoreClusterVmAntiAffinityRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatastoreClusterVmAntiAffinityRule)(nil)).Elem()
}

func (o DatastoreClusterVmAntiAffinityRuleMapOutput) ToDatastoreClusterVmAntiAffinityRuleMapOutput() DatastoreClusterVmAntiAffinityRuleMapOutput {
	return o
}

func (o DatastoreClusterVmAntiAffinityRuleMapOutput) ToDatastoreClusterVmAntiAffinityRuleMapOutputWithContext(ctx context.Context) DatastoreClusterVmAntiAffinityRuleMapOutput {
	return o
}

func (o DatastoreClusterVmAntiAffinityRuleMapOutput) MapIndex(k pulumi.StringInput) DatastoreClusterVmAntiAffinityRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatastoreClusterVmAntiAffinityRule {
		return vs[0].(map[string]*DatastoreClusterVmAntiAffinityRule)[vs[1].(string)]
	}).(DatastoreClusterVmAntiAffinityRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreClusterVmAntiAffinityRuleInput)(nil)).Elem(), &DatastoreClusterVmAntiAffinityRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreClusterVmAntiAffinityRuleArrayInput)(nil)).Elem(), DatastoreClusterVmAntiAffinityRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatastoreClusterVmAntiAffinityRuleMapInput)(nil)).Elem(), DatastoreClusterVmAntiAffinityRuleMap{})
	pulumi.RegisterOutputType(DatastoreClusterVmAntiAffinityRuleOutput{})
	pulumi.RegisterOutputType(DatastoreClusterVmAntiAffinityRuleArrayOutput{})
	pulumi.RegisterOutputType(DatastoreClusterVmAntiAffinityRuleMapOutput{})
}
