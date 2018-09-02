import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_compute_cluster_vm_host_rule` resource can be used to manage
 * VM-to-host rules in a cluster, either created by the
 * [`vsphere_compute_cluster`][tf-vsphere-cluster-resource] resource or looked up
 * by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source] data source.
 *
 * [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
 * [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
 *
 * This resource can create both _affinity rules_, where virtual machines run on
 * specified hosts, or _anti-affinity_ rules, where virtual machines run on hosts
 * outside of the ones specified in the rule. Virtual machines and hosts are
 * supplied via groups, which can be managed via the
 * [`vsphere_compute_cluster_vm_group`][tf-vsphere-cluster-vm-group-resource] and
 * [`vsphere_compute_cluster_host_group`][tf-vsphere-cluster-host-group-resource]
 * resources.
 *
 * [tf-vsphere-cluster-vm-group-resource]: /docs/providers/vsphere/r/compute_cluster_vm_group.html
 * [tf-vsphere-cluster-host-group-resource]: /docs/providers/vsphere/r/compute_cluster_host_group.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 */
export declare class ComputeClusterVmHostRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmHostRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmHostRuleState): ComputeClusterVmHostRule;
    /**
     * When this field is used, the virtual
     * machines defined in `vm_group_name` will be run on the
     * hosts defined in this host group.
     */
    readonly affinityHostGroupName: pulumi.Output<string | undefined>;
    /**
     * When this field is used, the
     * virtual machines defined in `vm_group_name` will _not_ be
     * run on the hosts defined in this host group.
     */
    readonly antiAffinityHostGroupName: pulumi.Output<string | undefined>;
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
     * The name of the rule. This must be unique in the
     * cluster.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The name of the virtual machine group to use
     * with this rule.
     */
    readonly vmGroupName: pulumi.Output<string>;
    /**
     * Create a ComputeClusterVmHostRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmHostRuleArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering ComputeClusterVmHostRule resources.
 */
export interface ComputeClusterVmHostRuleState {
    /**
     * When this field is used, the virtual
     * machines defined in `vm_group_name` will be run on the
     * hosts defined in this host group.
     */
    readonly affinityHostGroupName?: pulumi.Input<string>;
    /**
     * When this field is used, the
     * virtual machines defined in `vm_group_name` will _not_ be
     * run on the hosts defined in this host group.
     */
    readonly antiAffinityHostGroupName?: pulumi.Input<string>;
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
     * The name of the rule. This must be unique in the
     * cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the virtual machine group to use
     * with this rule.
     */
    readonly vmGroupName?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a ComputeClusterVmHostRule resource.
 */
export interface ComputeClusterVmHostRuleArgs {
    /**
     * When this field is used, the virtual
     * machines defined in `vm_group_name` will be run on the
     * hosts defined in this host group.
     */
    readonly affinityHostGroupName?: pulumi.Input<string>;
    /**
     * When this field is used, the
     * virtual machines defined in `vm_group_name` will _not_ be
     * run on the hosts defined in this host group.
     */
    readonly antiAffinityHostGroupName?: pulumi.Input<string>;
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
     * The name of the rule. This must be unique in the
     * cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the virtual machine group to use
     * with this rule.
     */
    readonly vmGroupName: pulumi.Input<string>;
}
