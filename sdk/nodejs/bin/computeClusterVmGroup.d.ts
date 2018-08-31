import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_compute_cluster_vm_group` resource can be used to manage groups of
 * virtual machines in a cluster, either created by the
 * [`vsphere_compute_cluster`][tf-vsphere-cluster-resource] resource or looked up
 * by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source] data source.
 *
 * [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
 * [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
 *
 * This resource mainly serves as an input to the
 * [`vsphere_compute_cluster_vm_dependency_rule`][tf-vsphere-cluster-vm-dependency-rule-resource]
 * and
 * [`vsphere_compute_cluster_vm_host_rule`][tf-vsphere-cluster-vm-host-rule-resource]
 * resources. See the individual resource documentation pages for more information.
 *
 * [tf-vsphere-cluster-vm-dependency-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_dependency_rule.html
 * [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 */
export declare class ComputeClusterVmGroup extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmGroupState): ComputeClusterVmGroup;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Output<string>;
    /**
     * The name of the VM group. This must be unique in the
     * cluster. Forces a new resource if changed.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The UUIDs of the virtual machines in this
     * group.
     */
    readonly virtualMachineIds: pulumi.Output<string[] | undefined>;
    /**
     * Create a ComputeClusterVmGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmGroupArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering ComputeClusterVmGroup resources.
 */
export interface ComputeClusterVmGroupState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * The name of the VM group. This must be unique in the
     * cluster. Forces a new resource if changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The UUIDs of the virtual machines in this
     * group.
     */
    readonly virtualMachineIds?: pulumi.Input<pulumi.Input<string>[]>;
}
/**
 * The set of arguments for constructing a ComputeClusterVmGroup resource.
 */
export interface ComputeClusterVmGroupArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * The name of the VM group. This must be unique in the
     * cluster. Forces a new resource if changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The UUIDs of the virtual machines in this
     * group.
     */
    readonly virtualMachineIds?: pulumi.Input<pulumi.Input<string>[]>;
}
