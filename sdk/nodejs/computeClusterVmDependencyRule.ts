// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.ComputeClusterVmDependencyRule` resource can be used to manage
 * VM dependency rules in a cluster, either created by the
 * `vsphere.ComputeCluster` resource or looked up
 * by the `vsphere.ComputeCluster` data source.
 *
 * A virtual machine dependency rule applies to vSphere HA, and allows
 * user-defined startup orders for virtual machines in the case of host failure.
 * Virtual machines are supplied via groups, which can be managed via the
 * `vsphere.ComputeClusterVmGroup`
 * resource.
 *
 * > **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ## Example Usage
 *
 * The example below creates two virtual machine in a cluster using the
 * `vsphere.VirtualMachine` resource in a cluster
 * looked up by the `vsphere.ComputeCluster`
 * data source. It then creates a group with this virtual machine. Two groups are created, each with one of the created VMs. Finally, a rule is created to ensure that `vm1` starts before `vm2`.
 *
 * > Note how `dependencyVmGroupName` and
 * `vmGroupName` are sourced off of the `name` attributes from
 * the `vsphere.ComputeClusterVmGroup`
 * resource. This is to ensure that the rule is not created before the groups
 * exist, which may not possibly happen in the event that the names came from a
 * "static" source such as a variable.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const datastore = datacenter.then(datacenter => vsphere.getDatastore({
 *     name: "datastore1",
 *     datacenterId: datacenter.id,
 * }));
 * const cluster = datacenter.then(datacenter => vsphere.getComputeCluster({
 *     name: "cluster-01",
 *     datacenterId: datacenter.id,
 * }));
 * const network = datacenter.then(datacenter => vsphere.getNetwork({
 *     name: "network1",
 *     datacenterId: datacenter.id,
 * }));
 * const vm1 = new vsphere.VirtualMachine("vm1", {
 *     name: "test1",
 *     resourcePoolId: cluster.then(cluster => cluster.resourcePoolId),
 *     datastoreId: datastore.then(datastore => datastore.id),
 *     numCpus: 2,
 *     memory: 2048,
 *     guestId: "otherLinux64Guest",
 *     networkInterfaces: [{
 *         networkId: network.then(network => network.id),
 *     }],
 *     disks: [{
 *         label: "disk0",
 *         size: 20,
 *     }],
 * });
 * const vm2 = new vsphere.VirtualMachine("vm2", {
 *     name: "test2",
 *     resourcePoolId: cluster.then(cluster => cluster.resourcePoolId),
 *     datastoreId: datastore.then(datastore => datastore.id),
 *     numCpus: 2,
 *     memory: 2048,
 *     guestId: "otherLinux64Guest",
 *     networkInterfaces: [{
 *         networkId: network.then(network => network.id),
 *     }],
 *     disks: [{
 *         label: "disk0",
 *         size: 20,
 *     }],
 * });
 * const clusterVmGroup1 = new vsphere.ComputeClusterVmGroup("cluster_vm_group1", {
 *     name: "test-cluster-vm-group1",
 *     computeClusterId: cluster.then(cluster => cluster.id),
 *     virtualMachineIds: [vm1.id],
 * });
 * const clusterVmGroup2 = new vsphere.ComputeClusterVmGroup("cluster_vm_group2", {
 *     name: "test-cluster-vm-group2",
 *     computeClusterId: cluster.then(cluster => cluster.id),
 *     virtualMachineIds: [vm2.id],
 * });
 * const clusterVmDependencyRule = new vsphere.ComputeClusterVmDependencyRule("cluster_vm_dependency_rule", {
 *     computeClusterId: cluster.then(cluster => cluster.id),
 *     name: "test-cluster-vm-dependency-rule",
 *     dependencyVmGroupName: clusterVmGroup1.name,
 *     vmGroupName: clusterVmGroup2.name,
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
 * $ pulumi import vsphere:index/computeClusterVmDependencyRule:ComputeClusterVmDependencyRule cluster_vm_dependency_rule \
 * ```
 *
 *   '{"compute_cluster_path": "/dc1/host/cluster1", \
 *
 *   "name": "pulumi-test-cluster-vm-dependency-rule"}'
 */
export class ComputeClusterVmDependencyRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmDependencyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmDependencyRuleState, opts?: pulumi.CustomResourceOptions): ComputeClusterVmDependencyRule {
        return new ComputeClusterVmDependencyRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/computeClusterVmDependencyRule:ComputeClusterVmDependencyRule';

    /**
     * Returns true if the given object is an instance of ComputeClusterVmDependencyRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeClusterVmDependencyRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeClusterVmDependencyRule.__pulumiType;
    }

    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    public readonly computeClusterId!: pulumi.Output<string>;
    /**
     * The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vmGroupName` will not be started until the VMs in this
     * group are started.
     */
    public readonly dependencyVmGroupName!: pulumi.Output<string>;
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
     * The name of the rule. This must be unique in the
     * cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependencyVmGroupName` are started.
     */
    public readonly vmGroupName!: pulumi.Output<string>;

    /**
     * Create a ComputeClusterVmDependencyRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmDependencyRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeClusterVmDependencyRuleArgs | ComputeClusterVmDependencyRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ComputeClusterVmDependencyRuleState | undefined;
            resourceInputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            resourceInputs["dependencyVmGroupName"] = state ? state.dependencyVmGroupName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["mandatory"] = state ? state.mandatory : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["vmGroupName"] = state ? state.vmGroupName : undefined;
        } else {
            const args = argsOrState as ComputeClusterVmDependencyRuleArgs | undefined;
            if ((!args || args.computeClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if ((!args || args.dependencyVmGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dependencyVmGroupName'");
            }
            if ((!args || args.vmGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmGroupName'");
            }
            resourceInputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            resourceInputs["dependencyVmGroupName"] = args ? args.dependencyVmGroupName : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["mandatory"] = args ? args.mandatory : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["vmGroupName"] = args ? args.vmGroupName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ComputeClusterVmDependencyRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ComputeClusterVmDependencyRule resources.
 */
export interface ComputeClusterVmDependencyRuleState {
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    computeClusterId?: pulumi.Input<string>;
    /**
     * The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vmGroupName` will not be started until the VMs in this
     * group are started.
     */
    dependencyVmGroupName?: pulumi.Input<string>;
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
     * The name of the rule. This must be unique in the
     * cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependencyVmGroupName` are started.
     */
    vmGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ComputeClusterVmDependencyRule resource.
 */
export interface ComputeClusterVmDependencyRuleArgs {
    /**
     * The managed object reference
     * ID of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    computeClusterId: pulumi.Input<string>;
    /**
     * The name of the VM group that this
     * rule depends on. The VMs defined in the group specified by
     * `vmGroupName` will not be started until the VMs in this
     * group are started.
     */
    dependencyVmGroupName: pulumi.Input<string>;
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
     * The name of the rule. This must be unique in the
     * cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the VM group that is the subject of
     * this rule. The VMs defined in this group will not be started until the VMs in
     * the group specified by
     * `dependencyVmGroupName` are started.
     */
    vmGroupName: pulumi.Input<string>;
}
