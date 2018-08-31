import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_compute_cluster_vm_affinity_rule` resource can be used to manage
 * VM affinity rules in a cluster, either created by the
 * [`vsphere_compute_cluster`][tf-vsphere-cluster-resource] resource or looked up
 * by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source] data source.
 *
 * [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
 * [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
 *
 * This rule can be used to tell a set to virtual machines to run together on a
 * single host within a cluster. When configured, DRS will make a best effort to
 * ensure that the virtual machines run on the same host, or prevent any operation
 * that would keep that from happening, depending on the value of the
 * `mandatory` flag.
 *
 * -> Keep in mind that this rule can only be used to tell VMs to run together on
 * a _non-specific_ host - it can't be used to pin VMs to a host. For that, see
 * the
 * [`vsphere_compute_cluster_vm_host_rule`][tf-vsphere-cluster-vm-host-rule-resource]
 * resource.
 *
 * [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 */
export declare class ComputeClusterVmAffinityRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmAffinityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmAffinityRuleState): ComputeClusterVmAffinityRule;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Output<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    readonly enabled: pulumi.Output<boolean | undefined>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     */
    readonly mandatory: pulumi.Output<boolean | undefined>;
    /**
     * The name of the rule. This must be unique in the cluster.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The UUIDs of the virtual machines to run
     * on the same host together.
     */
    readonly virtualMachineIds: pulumi.Output<string[]>;
    /**
     * Create a ComputeClusterVmAffinityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmAffinityRuleArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering ComputeClusterVmAffinityRule resources.
 */
export interface ComputeClusterVmAffinityRuleState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     */
    readonly mandatory?: pulumi.Input<boolean>;
    /**
     * The name of the rule. This must be unique in the cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The UUIDs of the virtual machines to run
     * on the same host together.
     */
    readonly virtualMachineIds?: pulumi.Input<pulumi.Input<string>[]>;
}
/**
 * The set of arguments for constructing a ComputeClusterVmAffinityRule resource.
 */
export interface ComputeClusterVmAffinityRuleArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     */
    readonly mandatory?: pulumi.Input<boolean>;
    /**
     * The name of the rule. This must be unique in the cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The UUIDs of the virtual machines to run
     * on the same host together.
     */
    readonly virtualMachineIds: pulumi.Input<pulumi.Input<string>[]>;
}
