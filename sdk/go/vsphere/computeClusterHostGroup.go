// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vsphere

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `vsphere_compute_cluster_host_group` resource can be used to manage groups
// of hosts in a cluster, either created by the
// [`vsphere_compute_cluster`][tf-vsphere-cluster-resource] resource or looked up
// by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source] data source.
// 
// [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
// [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
// 
// This resource mainly serves as an input to the
// [`vsphere_compute_cluster_vm_host_rule`][tf-vsphere-cluster-vm-host-rule-resource]
// resource - see the documentation for that resource for further details on how
// to use host groups.
// 
// [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
// 
// ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
// 
// ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
type ComputeClusterHostGroup struct {
	s *pulumi.ResourceState
}

// NewComputeClusterHostGroup registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterHostGroup(ctx *pulumi.Context,
	name string, args *ComputeClusterHostGroupArgs, opts ...pulumi.ResourceOpt) (*ComputeClusterHostGroup, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["computeClusterId"] = nil
		inputs["hostSystemIds"] = nil
		inputs["name"] = nil
	} else {
		inputs["computeClusterId"] = args.ComputeClusterId
		inputs["hostSystemIds"] = args.HostSystemIds
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ComputeClusterHostGroup{s: s}, nil
}

// GetComputeClusterHostGroup gets an existing ComputeClusterHostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterHostGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ComputeClusterHostGroupState, opts ...pulumi.ResourceOpt) (*ComputeClusterHostGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["computeClusterId"] = state.ComputeClusterId
		inputs["hostSystemIds"] = state.HostSystemIds
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("vsphere:index/computeClusterHostGroup:ComputeClusterHostGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ComputeClusterHostGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ComputeClusterHostGroup) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ComputeClusterHostGroup) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The [managed object reference
// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
// resource if changed.
func (r *ComputeClusterHostGroup) ComputeClusterId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["computeClusterId"])
}

// The [managed object IDs][docs-about-morefs] of
// the hosts to put in the cluster.
func (r *ComputeClusterHostGroup) HostSystemIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["hostSystemIds"])
}

// The name of the host group. This must be unique in the
// cluster. Forces a new resource if changed.
func (r *ComputeClusterHostGroup) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering ComputeClusterHostGroup resources.
type ComputeClusterHostGroupState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// The [managed object IDs][docs-about-morefs] of
	// the hosts to put in the cluster.
	HostSystemIds interface{}
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name interface{}
}

// The set of arguments for constructing a ComputeClusterHostGroup resource.
type ComputeClusterHostGroupArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId interface{}
	// The [managed object IDs][docs-about-morefs] of
	// the hosts to put in the cluster.
	HostSystemIds interface{}
	// The name of the host group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name interface{}
}
