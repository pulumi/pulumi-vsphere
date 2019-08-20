// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere..ComputeClusterVmDependencyRule` resource can be used to manage
 * VM dependency rules in a cluster, either created by the
 * [`vsphere..ComputeCluster`][tf-vsphere-cluster-resource] resource or looked up
 * by the [`vsphere..ComputeCluster`][tf-vsphere-cluster-data-source] data source.
 * 
 * [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
 * [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
 * 
 * A virtual machine dependency rule applies to vSphere HA, and allows
 * user-defined startup orders for virtual machines in the case of host failure.
 * Virtual machines are supplied via groups, which can be managed via the
 * [`vsphere..ComputeClusterVmGroup`][tf-vsphere-cluster-vm-group-resource]
 * resource.
 * 
 * [tf-vsphere-cluster-vm-group-resource]: /docs/providers/vsphere/r/compute_cluster_vm_group.html
 * 
 * > **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * ## Example Usage
 * 
 * The example below creates two virtual machine in a cluster using the
 * [`vsphere..VirtualMachine`][tf-vsphere-vm-resource] resource in a cluster
 * looked up by the [`vsphere..ComputeCluster`][tf-vsphere-cluster-data-source]
 * data source. It then creates a group with this virtual machine. Two groups are created, each with one of the created VMs. Finally, a rule is created to ensure that `vm1` starts before `vm2`.
 * 
 * [tf-vsphere-vm-resource]: /docs/providers/vsphere/r/virtual_machine.html
 * 
 * > Note how `dependencyVmGroupName` and
 * `vmGroupName` are sourced off of the `name` attributes from
 * the [`vsphere..ComputeClusterVmGroup`][tf-vsphere-cluster-vm-group-resource]
 * resource. This is to ensure that the rule is not created before the groups
 * exist, which may not possibly happen in the event that the names came from a
 * "static" source such as a variable.
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
 * const datastore = vsphere.getDatastore({
 *     datacenterId: dc.id,
 *     name: "datastore1",
 * });
 * const network = vsphere.getNetwork({
 *     datacenterId: dc.id,
 *     name: "network1",
 * });
 * const vm1 = new vsphere.VirtualMachine("vm1", {
 *     datastoreId: datastore.id,
 *     disks: [{
 *         label: "disk0",
 *         size: 20,
 *     }],
 *     guestId: "other3xLinux64Guest",
 *     memory: 2048,
 *     networkInterfaces: [{
 *         networkId: network.id,
 *     }],
 *     numCpus: 2,
 *     resourcePoolId: cluster.resourcePoolId,
 * });
 * const clusterVmGroup1 = new vsphere.ComputeClusterVmGroup("clusterVmGroup1", {
 *     computeClusterId: cluster.id,
 *     virtualMachineIds: [vm1.id],
 * });
 * const vm2 = new vsphere.VirtualMachine("vm2", {
 *     datastoreId: datastore.id,
 *     disks: [{
 *         label: "disk0",
 *         size: 20,
 *     }],
 *     guestId: "other3xLinux64Guest",
 *     memory: 2048,
 *     networkInterfaces: [{
 *         networkId: network.id,
 *     }],
 *     numCpus: 2,
 *     resourcePoolId: cluster.resourcePoolId,
 * });
 * const clusterVmGroup2 = new vsphere.ComputeClusterVmGroup("clusterVmGroup2", {
 *     computeClusterId: cluster.id,
 *     virtualMachineIds: [vm2.id],
 * });
 * const clusterVmDependencyRule = new vsphere.ComputeClusterVmDependencyRule("clusterVmDependencyRule", {
 *     computeClusterId: cluster.id,
 *     dependencyVmGroupName: clusterVmGroup1.name,
 *     vmGroupName: clusterVmGroup2.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_vm_dependency_rule.html.markdown.
 */
export class ComputeClusterVmDependencyRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmDependencyRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ComputeClusterVmDependencyRuleState | undefined;
            inputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            inputs["dependencyVmGroupName"] = state ? state.dependencyVmGroupName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["mandatory"] = state ? state.mandatory : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vmGroupName"] = state ? state.vmGroupName : undefined;
        } else {
            const args = argsOrState as ComputeClusterVmDependencyRuleArgs | undefined;
            if (!args || args.computeClusterId === undefined) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if (!args || args.dependencyVmGroupName === undefined) {
                throw new Error("Missing required property 'dependencyVmGroupName'");
            }
            if (!args || args.vmGroupName === undefined) {
                throw new Error("Missing required property 'vmGroupName'");
            }
            inputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            inputs["dependencyVmGroupName"] = args ? args.dependencyVmGroupName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["mandatory"] = args ? args.mandatory : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vmGroupName"] = args ? args.vmGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ComputeClusterVmDependencyRule.__pulumiType, name, inputs, opts);
    }
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
     * `vmGroupName` will not be started until the VMs in this
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
     * `dependencyVmGroupName` are started.
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
     * `vmGroupName` will not be started until the VMs in this
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
     * `dependencyVmGroupName` are started.
     */
    readonly vmGroupName: pulumi.Input<string>;
}
