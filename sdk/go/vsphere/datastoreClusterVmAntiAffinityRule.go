// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `vsphere_datastore_cluster_vm_anti_affinity_rule` resource can be used to
// manage VM anti-affinity rules in a datastore cluster, either created by the
// [`vsphere_datastore_cluster`][tf-vsphere-datastore-cluster-resource] resource or looked up
// by the [`vsphere_datastore_cluster`][tf-vsphere-datastore-cluster-data-source] data source.
// 
// [tf-vsphere-datastore-cluster-resource]: /docs/providers/vsphere/r/datastore_cluster.html
// [tf-vsphere-datastore-cluster-data-source]: /docs/providers/vsphere/d/datastore_cluster.html
// 
// This rule can be used to tell a set to virtual machines to run on different
// datastores within a cluster, useful for preventing single points of failure in
// application cluster scenarios. When configured, Storage DRS will make a best effort to
// ensure that the virtual machines run on different datastores, or prevent any
// operation that would keep that from happening, depending on the value of the
// `mandatory` flag.
// 
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
// 
// > **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
type DatastoreClusterVmAntiAffinityRule struct {
	s *pulumi.ResourceState
}

// NewDatastoreClusterVmAntiAffinityRule registers a new resource with the given unique name, arguments, and options.
func NewDatastoreClusterVmAntiAffinityRule(ctx *pulumi.Context,
	name string, args *DatastoreClusterVmAntiAffinityRuleArgs, opts ...pulumi.ResourceOpt) (*DatastoreClusterVmAntiAffinityRule, error) {
	if args == nil || args.DatastoreClusterId == nil {
		return nil, errors.New("missing required argument 'DatastoreClusterId'")
	}
	if args == nil || args.VirtualMachineIds == nil {
		return nil, errors.New("missing required argument 'VirtualMachineIds'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["datastoreClusterId"] = nil
		inputs["enabled"] = nil
		inputs["mandatory"] = nil
		inputs["name"] = nil
		inputs["virtualMachineIds"] = nil
	} else {
		inputs["datastoreClusterId"] = args.DatastoreClusterId
		inputs["enabled"] = args.Enabled
		inputs["mandatory"] = args.Mandatory
		inputs["name"] = args.Name
		inputs["virtualMachineIds"] = args.VirtualMachineIds
	}
	s, err := ctx.RegisterResource("vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatastoreClusterVmAntiAffinityRule{s: s}, nil
}

// GetDatastoreClusterVmAntiAffinityRule gets an existing DatastoreClusterVmAntiAffinityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatastoreClusterVmAntiAffinityRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DatastoreClusterVmAntiAffinityRuleState, opts ...pulumi.ResourceOpt) (*DatastoreClusterVmAntiAffinityRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["datastoreClusterId"] = state.DatastoreClusterId
		inputs["enabled"] = state.Enabled
		inputs["mandatory"] = state.Mandatory
		inputs["name"] = state.Name
		inputs["virtualMachineIds"] = state.VirtualMachineIds
	}
	s, err := ctx.ReadResource("vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatastoreClusterVmAntiAffinityRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DatastoreClusterVmAntiAffinityRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DatastoreClusterVmAntiAffinityRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The [managed object reference
// ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
// a new resource if changed.
func (r *DatastoreClusterVmAntiAffinityRule) DatastoreClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["datastoreClusterId"])
}

// Enable this rule in the cluster. Default: `true`.
func (r *DatastoreClusterVmAntiAffinityRule) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// When this value is `true`, prevents any virtual
// machine operations that may violate this rule. Default: `false`.
func (r *DatastoreClusterVmAntiAffinityRule) Mandatory() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["mandatory"])
}

// The name of the rule. This must be unique in the cluster.
func (r *DatastoreClusterVmAntiAffinityRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The UUIDs of the virtual machines to run
// on different datastores from each other.
func (r *DatastoreClusterVmAntiAffinityRule) VirtualMachineIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["virtualMachineIds"])
}

// Input properties used for looking up and filtering DatastoreClusterVmAntiAffinityRule resources.
type DatastoreClusterVmAntiAffinityRuleState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
	// a new resource if changed.
	DatastoreClusterId interface{}
	// Enable this rule in the cluster. Default: `true`.
	Enabled interface{}
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory interface{}
	// The name of the rule. This must be unique in the cluster.
	Name interface{}
	// The UUIDs of the virtual machines to run
	// on different datastores from each other.
	VirtualMachineIds interface{}
}

// The set of arguments for constructing a DatastoreClusterVmAntiAffinityRule resource.
type DatastoreClusterVmAntiAffinityRuleArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
	// a new resource if changed.
	DatastoreClusterId interface{}
	// Enable this rule in the cluster. Default: `true`.
	Enabled interface{}
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory interface{}
	// The name of the rule. This must be unique in the cluster.
	Name interface{}
	// The UUIDs of the virtual machines to run
	// on different datastores from each other.
	VirtualMachineIds interface{}
}
