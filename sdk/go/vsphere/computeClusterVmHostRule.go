// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.ComputeClusterVmHostRule` resource can be used to manage
// VM-to-host rules in a cluster, either created by the
// [`.ComputeCluster`][tf-vsphere-cluster-resource] resource or looked up
// by the [`.ComputeCluster`][tf-vsphere-cluster-data-source] data source.
// 
// [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
// [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
// 
// This resource can create both _affinity rules_, where virtual machines run on
// specified hosts, or _anti-affinity_ rules, where virtual machines run on hosts
// outside of the ones specified in the rule. Virtual machines and hosts are
// supplied via groups, which can be managed via the
// [`.ComputeClusterVmGroup`][tf-vsphere-cluster-vm-group-resource] and
// [`.ComputeClusterHostGroup`][tf-vsphere-cluster-host-group-resource]
// resources.
// 
// [tf-vsphere-cluster-vm-group-resource]: /docs/providers/vsphere/r/compute_cluster_vm_group.html
// [tf-vsphere-cluster-host-group-resource]: /docs/providers/vsphere/r/compute_cluster_host_group.html
// 
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
// 
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_vm_host_rule.html.markdown.
type ComputeClusterVmHostRule struct {
	s *pulumi.ResourceState
}

// NewComputeClusterVmHostRule registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmHostRule(ctx *pulumi.Context,
	name string, args *ComputeClusterVmHostRuleArgs, opts ...pulumi.ResourceOpt) (*ComputeClusterVmHostRule, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil || args.VmGroupName == nil {
		return nil, errors.New("missing required argument 'VmGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["affinityHostGroupName"] = nil
		inputs["antiAffinityHostGroupName"] = nil
		inputs["computeClusterId"] = nil
		inputs["enabled"] = nil
		inputs["mandatory"] = nil
		inputs["name"] = nil
		inputs["vmGroupName"] = nil
	} else {
		inputs["affinityHostGroupName"] = args.AffinityHostGroupName
		inputs["antiAffinityHostGroupName"] = args.AntiAffinityHostGroupName
		inputs["computeClusterId"] = args.ComputeClusterId
		inputs["enabled"] = args.Enabled
		inputs["mandatory"] = args.Mandatory
		inputs["name"] = args.Name
		inputs["vmGroupName"] = args.VmGroupName
	}
	s, err := ctx.RegisterResource("vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ComputeClusterVmHostRule{s: s}, nil
}

// GetComputeClusterVmHostRule gets an existing ComputeClusterVmHostRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmHostRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ComputeClusterVmHostRuleState, opts ...pulumi.ResourceOpt) (*ComputeClusterVmHostRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["affinityHostGroupName"] = state.AffinityHostGroupName
		inputs["antiAffinityHostGroupName"] = state.AntiAffinityHostGroupName
		inputs["computeClusterId"] = state.ComputeClusterId
		inputs["enabled"] = state.Enabled
		inputs["mandatory"] = state.Mandatory
		inputs["name"] = state.Name
		inputs["vmGroupName"] = state.VmGroupName
	}
	s, err := ctx.ReadResource("vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ComputeClusterVmHostRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ComputeClusterVmHostRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ComputeClusterVmHostRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// When this field is used, the virtual
// machines defined in `vmGroupName` will be run on the
// hosts defined in this host group.
func (r *ComputeClusterVmHostRule) AffinityHostGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["affinityHostGroupName"])
}

// When this field is used, the
// virtual machines defined in `vmGroupName` will _not_ be
// run on the hosts defined in this host group.
func (r *ComputeClusterVmHostRule) AntiAffinityHostGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["antiAffinityHostGroupName"])
}

// The [managed object reference
// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
// resource if changed.
func (r *ComputeClusterVmHostRule) ComputeClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["computeClusterId"])
}

// Enable this rule in the cluster. Default: `true`.
func (r *ComputeClusterVmHostRule) Enabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enabled"])
}

// When this value is `true`, prevents any virtual
// machine operations that may violate this rule. Default: `false`.
func (r *ComputeClusterVmHostRule) Mandatory() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["mandatory"])
}

// The name of the rule. This must be unique in the
// cluster.
func (r *ComputeClusterVmHostRule) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the virtual machine group to use
// with this rule.
func (r *ComputeClusterVmHostRule) VmGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vmGroupName"])
}

// Input properties used for looking up and filtering ComputeClusterVmHostRule resources.
type ComputeClusterVmHostRuleState struct {
	// When this field is used, the virtual
	// machines defined in `vmGroupName` will be run on the
	// hosts defined in this host group.
	AffinityHostGroupName interface{}
	// When this field is used, the
	// virtual machines defined in `vmGroupName` will _not_ be
	// run on the hosts defined in this host group.
	AntiAffinityHostGroupName interface{}
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// Enable this rule in the cluster. Default: `true`.
	Enabled interface{}
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory interface{}
	// The name of the rule. This must be unique in the
	// cluster.
	Name interface{}
	// The name of the virtual machine group to use
	// with this rule.
	VmGroupName interface{}
}

// The set of arguments for constructing a ComputeClusterVmHostRule resource.
type ComputeClusterVmHostRuleArgs struct {
	// When this field is used, the virtual
	// machines defined in `vmGroupName` will be run on the
	// hosts defined in this host group.
	AffinityHostGroupName interface{}
	// When this field is used, the
	// virtual machines defined in `vmGroupName` will _not_ be
	// run on the hosts defined in this host group.
	AntiAffinityHostGroupName interface{}
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// Enable this rule in the cluster. Default: `true`.
	Enabled interface{}
	// When this value is `true`, prevents any virtual
	// machine operations that may violate this rule. Default: `false`.
	Mandatory interface{}
	// The name of the rule. This must be unique in the
	// cluster.
	Name interface{}
	// The name of the virtual machine group to use
	// with this rule.
	VmGroupName interface{}
}
