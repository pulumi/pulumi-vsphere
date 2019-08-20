// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere..DatastoreClusterVmAntiAffinityRule` resource can be used to
 * manage VM anti-affinity rules in a datastore cluster, either created by the
 * [`vsphere..DatastoreCluster`][tf-vsphere-datastore-cluster-resource] resource or looked up
 * by the [`vsphere..DatastoreCluster`][tf-vsphere-datastore-cluster-data-source] data source.
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
 * > **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * > **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
 * 
 * ## Example Usage
 * 
 * The example below creates two virtual machines in a cluster using the
 * [`vsphere..VirtualMachine`][tf-vsphere-vm-resource] resource, creating the
 * virtual machines in the datastore cluster looked up by the
 * [`vsphere..DatastoreCluster`][tf-vsphere-datastore-cluster-data-source] data
 * source. It then creates an anti-affinity rule for these two virtual machines,
 * ensuring they will run on different datastores whenever possible.
 * 
 * [tf-vsphere-vm-resource]: /docs/providers/vsphere/r/virtual_machine.html
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const dc = vsphere.getDatacenter({
 *     name: "dc1",
 * });
 * const cluster = vsphere.getComputeCluster({
 *     datacenterId: dc.id,
 *     name: "cluster1",
 * });
 * const datastoreCluster = vsphere.getDatastoreCluster({
 *     datacenterId: dc.id,
 *     name: "datastore-cluster1",
 * });
 * const network = vsphere.getNetwork({
 *     datacenterId: dc.id,
 *     name: "network1",
 * });
 * const vm: vsphere.VirtualMachine[] = [];
 * for (let i = 0; i < 2; i++) {
 *     vm.push(new vsphere.VirtualMachine(`vm-${i}`, {
 *         datastoreClusterId: datastoreCluster.id,
 *         disks: [{
 *             label: "disk0",
 *             size: 20,
 *         }],
 *         guestId: "other3xLinux64Guest",
 *         memory: 2048,
 *         networkInterfaces: [{
 *             networkId: network.id,
 *         }],
 *         numCpus: 2,
 *         resourcePoolId: cluster.resourcePoolId,
 *     }));
 * }
 * const clusterVmAntiAffinityRule = new vsphere.DatastoreClusterVmAntiAffinityRule("clusterVmAntiAffinityRule", {
 *     datastoreClusterId: datastoreCluster.id,
 *     virtualMachineIds: vm.map(v => v.id),
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/datastore_cluster_vm_anti_affinity_rule.html.markdown.
 */
export class DatastoreClusterVmAntiAffinityRule extends pulumi.CustomResource {
    /**
     * Get an existing DatastoreClusterVmAntiAffinityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatastoreClusterVmAntiAffinityRuleState, opts?: pulumi.CustomResourceOptions): DatastoreClusterVmAntiAffinityRule {
        return new DatastoreClusterVmAntiAffinityRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule';

    /**
     * Returns true if the given object is an instance of DatastoreClusterVmAntiAffinityRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatastoreClusterVmAntiAffinityRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatastoreClusterVmAntiAffinityRule.__pulumiType;
    }

    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
     * a new resource if changed.
     */
    public readonly datastoreClusterId!: pulumi.Output<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     */
    public readonly mandatory!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the rule. This must be unique in the cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The UUIDs of the virtual machines to run
     * on different datastores from each other.
     */
    public readonly virtualMachineIds!: pulumi.Output<string[]>;

    /**
     * Create a DatastoreClusterVmAntiAffinityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatastoreClusterVmAntiAffinityRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatastoreClusterVmAntiAffinityRuleArgs | DatastoreClusterVmAntiAffinityRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatastoreClusterVmAntiAffinityRuleState | undefined;
            inputs["datastoreClusterId"] = state ? state.datastoreClusterId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["mandatory"] = state ? state.mandatory : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["virtualMachineIds"] = state ? state.virtualMachineIds : undefined;
        } else {
            const args = argsOrState as DatastoreClusterVmAntiAffinityRuleArgs | undefined;
            if (!args || args.datastoreClusterId === undefined) {
                throw new Error("Missing required property 'datastoreClusterId'");
            }
            if (!args || args.virtualMachineIds === undefined) {
                throw new Error("Missing required property 'virtualMachineIds'");
            }
            inputs["datastoreClusterId"] = args ? args.datastoreClusterId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["mandatory"] = args ? args.mandatory : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["virtualMachineIds"] = args ? args.virtualMachineIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatastoreClusterVmAntiAffinityRule.__pulumiType, name, inputs, opts);
    }
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
