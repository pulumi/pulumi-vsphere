// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.ComputeClusterVmAntiAffinityRule` resource can be used to
// manage VM anti-affinity rules in a cluster, either created by the
// [`.ComputeCluster`][tf-vsphere-cluster-resource] resource or looked up
// by the [`.ComputeCluster`][tf-vsphere-cluster-data-source] data source.
// 
// [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
// [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
// 
// This rule can be used to tell a set to virtual machines to run on different
// hosts within a cluster, useful for preventing single points of failure in
// application cluster scenarios. When configured, DRS will make a best effort to
// ensure that the virtual machines run on different hosts, or prevent any
// operation that would keep that from happening, depending on the value of the
// `mandatory` flag.
// 
// > Keep in mind that this rule can only be used to tell VMs to run separately
// on _non-specific_ hosts - specific hosts cannot be specified with this rule.
// For that, see the
// [`.ComputeClusterVmHostRule`][tf-vsphere-cluster-vm-host-rule-resource]
// resource.
// 
// [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
// 
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
// 
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_vm_anti_affinity_rule.html.markdown.
type ComputeClusterVmAntiAffinityRule struct {
	pulumi.CustomResourceState

	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
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
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil || args.VirtualMachineIds == nil {
		return nil, errors.New("missing required argument 'VirtualMachineIds'")
	}
	if args == nil {
		args = &ComputeClusterVmAntiAffinityRuleArgs{}
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
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
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
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
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
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
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
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
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

