// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package vsphere

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The `.ComputeClusterVmGroup` resource can be used to manage groups of
// virtual machines in a cluster, either created by the
// [`.ComputeCluster`][tf-vsphere-cluster-resource] resource or looked up
// by the [`.ComputeCluster`][tf-vsphere-cluster-data-source] data source.
//
// [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
// [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
//
// This resource mainly serves as an input to the
// [`.ComputeClusterVmDependencyRule`][tf-vsphere-cluster-vm-dependency-rule-resource]
// and
// [`.ComputeClusterVmHostRule`][tf-vsphere-cluster-vm-host-rule-resource]
// resources. See the individual resource documentation pages for more information.
//
// [tf-vsphere-cluster-vm-dependency-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_dependency_rule.html
// [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
//
// > **NOTE:** This resource requires vCenter and is not available on direct ESXi
// connections.
//
// > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_vm_group.html.markdown.
type ComputeClusterVmGroup struct {
	pulumi.CustomResourceState

	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringOutput `pulumi:"computeClusterId"`
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds pulumi.StringArrayOutput `pulumi:"virtualMachineIds"`
}

// NewComputeClusterVmGroup registers a new resource with the given unique name, arguments, and options.
func NewComputeClusterVmGroup(ctx *pulumi.Context,
	name string, args *ComputeClusterVmGroupArgs, opts ...pulumi.ResourceOption) (*ComputeClusterVmGroup, error) {
	if args == nil || args.ComputeClusterId == nil {
		return nil, errors.New("missing required argument 'ComputeClusterId'")
	}
	if args == nil {
		args = &ComputeClusterVmGroupArgs{}
	}
	var resource ComputeClusterVmGroup
	err := ctx.RegisterResource("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeClusterVmGroup gets an existing ComputeClusterVmGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeClusterVmGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeClusterVmGroupState, opts ...pulumi.ResourceOption) (*ComputeClusterVmGroup, error) {
	var resource ComputeClusterVmGroup
	err := ctx.ReadResource("vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeClusterVmGroup resources.
type computeClusterVmGroupState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId *string `pulumi:"computeClusterId"`
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

type ComputeClusterVmGroupState struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringPtrInput
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmGroupState)(nil)).Elem()
}

type computeClusterVmGroupArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId string `pulumi:"computeClusterId"`
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name *string `pulumi:"name"`
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds []string `pulumi:"virtualMachineIds"`
}

// The set of arguments for constructing a ComputeClusterVmGroup resource.
type ComputeClusterVmGroupArgs struct {
	// The [managed object reference
	// ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
	// resource if changed.
	ComputeClusterId pulumi.StringInput
	// The name of the VM group. This must be unique in the
	// cluster. Forces a new resource if changed.
	Name pulumi.StringPtrInput
	// The UUIDs of the virtual machines in this
	// group.
	VirtualMachineIds pulumi.StringArrayInput
}

func (ComputeClusterVmGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeClusterVmGroupArgs)(nil)).Elem()
}

