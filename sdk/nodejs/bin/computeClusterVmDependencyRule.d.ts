import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_compute_cluster_vm_dependency_rule` resource can be used to manage
 * VM dependency rules in a cluster, either created by the
 * [`vsphere_compute_cluster`][tf-vsphere-cluster-resource] resource or looked up
 * by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source] data source.
 *
 * [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
 * [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
 *
 * A virtual machine dependency rule applies to vSphere HA, and allows
 * user-defined startup orders for virtual machines in the case of host failure.
 * Virtual machines are supplied via groups, which can be managed via the
 * [`vsphere_compute_cluster_vm_group`][tf-vsphere-cluster-vm-group-resource]
 * resource.
 *
 * [tf-vsphere-cluster-vm-group-resource]: /docs/providers/vsphere/r/compute_cluster_vm_group.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 */
export declare class ComputeClusterVmDependencyRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmDependencyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmDependencyRuleState): ComputeClusterVmDependencyRule;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Output<string>;
    /**
     * The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vm_group_name` will not be started until the VMs in this
     * group are started.
     */
    readonly dependencyVmGroupName: pulumi.Output<string>;
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
     * The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependency_vm_group_name` are started.
     */
    readonly vmGroupName: pulumi.Output<string>;
    /**
     * Create a ComputeClusterVmDependencyRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmDependencyRuleArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering ComputeClusterVmDependencyRule resources.
 */
export interface ComputeClusterVmDependencyRuleState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vm_group_name` will not be started until the VMs in this
     * group are started.
     */
    readonly dependencyVmGroupName?: pulumi.Input<string>;
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
     * The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependency_vm_group_name` are started.
     */
    readonly vmGroupName?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a ComputeClusterVmDependencyRule resource.
 */
export interface ComputeClusterVmDependencyRuleArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vm_group_name` will not be started until the VMs in this
     * group are started.
     */
    readonly dependencyVmGroupName: pulumi.Input<string>;
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
     * The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependency_vm_group_name` are started.
     */
    readonly vmGroupName: pulumi.Input<string>;
}
