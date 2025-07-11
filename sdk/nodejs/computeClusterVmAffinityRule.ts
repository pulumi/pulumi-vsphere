// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.ComputeClusterVmAffinityRule` resource can be used to
 * manage virtual machine affinity rules in a cluster, either created by the
 * `vsphere.ComputeCluster` resource or looked up
 * by the `vsphere.ComputeCluster` data source.
 *
 * This rule can be used to tell a set of virtual machines to run together on the
 * same host within a cluster. When configured, DRS will make a best effort to
 * ensure that the virtual machines run on the same host, or prevent any operation
 * that would keep that from happening, depending on the value of the
 * `mandatory` flag.
 *
 * > An affinity rule can only be used to place virtual machines on the same
 * _non-specific_ hosts. It cannot be used to pin virtual machines to a host.
 * To enable this capability, use the
 * `vsphere.ComputeClusterVmHostRule`
 * resource.
 *
 * > **NOTE:** This resource requires vCenter Server and is not available on
 * direct ESXi host connections.
 *
 * ## Example Usage
 *
 * The following example creates two virtual machines in a cluster using the
 * `vsphere.VirtualMachine` resource, creating the
 * virtual machines in the cluster looked up by the
 * `vsphere.ComputeCluster` data source. It
 * then creates an affinity rule for these two virtual machines, ensuring they
 * will run on the same host whenever possible.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const datastore = datacenter.then(datacenter => vsphere.getDatastore({
 *     name: "datastore-01",
 *     datacenterId: datacenter.id,
 * }));
 * const cluster = datacenter.then(datacenter => vsphere.getComputeCluster({
 *     name: "cluster-01",
 *     datacenterId: datacenter.id,
 * }));
 * const network = datacenter.then(datacenter => vsphere.getNetwork({
 *     name: "VM Network",
 *     datacenterId: datacenter.id,
 * }));
 * const vm: vsphere.VirtualMachine[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     vm.push(new vsphere.VirtualMachine(`vm-${range.value}`, {
 *         name: `foo-${range.value}`,
 *         resourcePoolId: cluster.then(cluster => cluster.resourcePoolId),
 *         datastoreId: datastore.then(datastore => datastore.id),
 *         numCpus: 1,
 *         memory: 1024,
 *         guestId: "otherLinux64Guest",
 *         networkInterfaces: [{
 *             networkId: network.then(network => network.id),
 *         }],
 *         disks: [{
 *             label: "disk0",
 *             size: 20,
 *         }],
 *     }));
 * }
 * const vmAffinityRule = new vsphere.ComputeClusterVmAffinityRule("vm_affinity_rule", {
 *     name: "vm-affinity-rule",
 *     computeClusterId: cluster.then(cluster => cluster.id),
 *     virtualMachineIds: vm.map((v, k) => [k, v]).map(([k, v]) => (v.id)),
 * });
 * ```
 *
 * The following example creates an affinity rule for a set of virtual machines
 * in the cluster by looking up the virtual machine UUIDs from the
 * `vsphere.VirtualMachine` data source.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const vms = [
 *     "foo-0",
 *     "foo-1",
 * ];
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const cluster = datacenter.then(datacenter => vsphere.getComputeCluster({
 *     name: "cluster-01",
 *     datacenterId: datacenter.id,
 * }));
 * const vmsGetVirtualMachine = (new Array(vms.length)).map((_, i) => i).map(__index => (vsphere.getVirtualMachine({
 *     name: vms[__index],
 *     datacenterId: _arg0_.id,
 * })));
 * const vmAffinityRule = new vsphere.ComputeClusterVmAffinityRule("vm_affinity_rule", {
 *     name: "vm-affinity-rule",
 *     enabled: true,
 *     computeClusterId: cluster.then(cluster => cluster.id),
 *     virtualMachineIds: vmsGetVirtualMachine.map(__item => __item.id),
 * });
 * ```
 *
 * ## Import
 *
 * An existing rule can be imported into this resource by supplying
 *
 * both the path to the cluster, and the name the rule. If the name or cluster is
 *
 * not found, or if the rule is of a different type, an error will be returned. An
 *
 * example is below:
 *
 * ```sh
 * $ pulumi import vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule vm_affinity_rule \
 * ```
 *
 *   '{"compute_cluster_path": "/dc-01/host/cluster-01", \
 *
 *   "name": "vm-affinity-rule"}'
 */
export class ComputeClusterVmAffinityRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmAffinityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmAffinityRuleState, opts?: pulumi.CustomResourceOptions): ComputeClusterVmAffinityRule {
        return new ComputeClusterVmAffinityRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule';

    /**
     * Returns true if the given object is an instance of ComputeClusterVmAffinityRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeClusterVmAffinityRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeClusterVmAffinityRule.__pulumiType;
    }

    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    public readonly computeClusterId!: pulumi.Output<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     *
     * > **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     */
    public readonly mandatory!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the rule. This must be unique in the cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The UUIDs of the virtual machines to run
     * on the same host together.
     */
    public readonly virtualMachineIds!: pulumi.Output<string[]>;

    /**
     * Create a ComputeClusterVmAffinityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmAffinityRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeClusterVmAffinityRuleArgs | ComputeClusterVmAffinityRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeClusterVmAffinityRuleState | undefined;
            resourceInputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["mandatory"] = state ? state.mandatory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["virtualMachineIds"] = state ? state.virtualMachineIds : undefined;
        } else {
            const args = argsOrState as ComputeClusterVmAffinityRuleArgs | undefined;
            if ((!args || args.computeClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if ((!args || args.virtualMachineIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMachineIds'");
            }
            resourceInputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["mandatory"] = args ? args.mandatory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["virtualMachineIds"] = args ? args.virtualMachineIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComputeClusterVmAffinityRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeClusterVmAffinityRule resources.
 */
export interface ComputeClusterVmAffinityRuleState {
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    computeClusterId?: pulumi.Input<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     *
     * > **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     */
    mandatory?: pulumi.Input<boolean>;
    /**
     * The name of the rule. This must be unique in the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The UUIDs of the virtual machines to run
     * on the same host together.
     */
    virtualMachineIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ComputeClusterVmAffinityRule resource.
 */
export interface ComputeClusterVmAffinityRuleArgs {
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    computeClusterId: pulumi.Input<string>;
    /**
     * Enable this rule in the cluster. Default: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When this value is `true`, prevents any virtual
     * machine operations that may violate this rule. Default: `false`.
     *
     * > **NOTE:** The namespace for rule names on this resource (defined by the
     * `name` argument) is shared with all rules in the cluster - consider
     * this when naming your rules.
     */
    mandatory?: pulumi.Input<boolean>;
    /**
     * The name of the rule. This must be unique in the cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The UUIDs of the virtual machines to run
     * on the same host together.
     */
    virtualMachineIds: pulumi.Input<pulumi.Input<string>[]>;
}
