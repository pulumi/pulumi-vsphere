import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_datastore_cluster_vm_anti_affinity_rule` resource can be used to
 * manage VM anti-affinity rules in a datastore cluster, either created by the
 * [`vsphere_datastore_cluster`][tf-vsphere-datastore-cluster-resource] resource or looked up
 * by the [`vsphere_datastore_cluster`][tf-vsphere-datastore-cluster-data-source] data source.
 *
 * [tf-vsphere-datastore-cluster-resource]: /docs/providers/vsphere/r/datastore_cluster.html
 * [tf-vsphere-datastore-cluster-data-source]: /docs/providers/vsphere/d/datastore_cluster.html
 *
 * This rule can be used to tell a set to virtual machines to run on different
 * datastores within a cluster, useful for preventing single points of failure in
 * application cluster scenarios. When configured, Storage DRS will make a best effort to
 * ensure that the virtual machines run on different datastores, or prevent any
 * operation that would keep that from happening, depending on the value of the
 * `mandatory` flag.
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
 */
export declare class DatastoreClusterVmAntiAffinityRule extends pulumi.CustomResource {
    /**
     * Get an existing DatastoreClusterVmAntiAffinityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatastoreClusterVmAntiAffinityRuleState): DatastoreClusterVmAntiAffinityRule;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
     * a new resource if changed.
     */
    readonly datastoreClusterId: pulumi.Output<string>;
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
     * on different datastores from each other.
     */
    readonly virtualMachineIds: pulumi.Output<string[]>;
    /**
     * Create a DatastoreClusterVmAntiAffinityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatastoreClusterVmAntiAffinityRuleArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering DatastoreClusterVmAntiAffinityRule resources.
 */
export interface DatastoreClusterVmAntiAffinityRuleState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
     * a new resource if changed.
     */
    readonly datastoreClusterId?: pulumi.Input<string>;
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
     * on different datastores from each other.
     */
    readonly virtualMachineIds?: pulumi.Input<pulumi.Input<string>[]>;
}
/**
 * The set of arguments for constructing a DatastoreClusterVmAntiAffinityRule resource.
 */
export interface DatastoreClusterVmAntiAffinityRuleArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
     * a new resource if changed.
     */
    readonly datastoreClusterId: pulumi.Input<string>;
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
     * on different datastores from each other.
     */
    readonly virtualMachineIds: pulumi.Input<pulumi.Input<string>[]>;
}
