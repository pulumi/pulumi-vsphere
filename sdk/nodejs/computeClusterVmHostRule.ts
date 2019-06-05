// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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
 * > **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 * 
 * ## Example Usage
 * 
 * The example below creates a virtual machine in a cluster using the
 * [`vsphere_virtual_machine`][tf-vsphere-vm-resource] resource in a cluster
 * looked up by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source]
 * data source. It then creates a group with this virtual machine. It also creates
 * a host group off of the host looked up via the
 * [`vsphere_host`][tf-vsphere-host-data-source] data source. Finally, this
 * virtual machine is configured to run specifically on that host via a
 * `vsphere_compute_cluster_vm_host_rule` resource.
 * 
 * [tf-vsphere-vm-resource]: /docs/providers/vsphere/r/virtual_machine.html
 * [tf-vsphere-host-data-source]: /docs/providers/vsphere/d/host.html
 * 
 * > Note how `vm_group_name` and
 * `affinity_host_group_name` are sourced off of the
 * `name` attributes from the
 * [`vsphere_compute_cluster_vm_group`][tf-vsphere-cluster-vm-group-resource] and
 * [`vsphere_compute_cluster_host_group`][tf-vsphere-cluster-host-group-resource]
 * resources. This is to ensure that the rule is not created before the groups
 * exist, which may not possibly happen in the event that the names came from a
 * "static" source such as a variable.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const dc = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }));
 * const cluster = dc.apply(dc => vsphere.getComputeCluster({
 *     datacenterId: dc.id,
 *     name: "cluster1",
 * }));
 * const datastore = dc.apply(dc => vsphere.getDatastore({
 *     datacenterId: dc.id,
 *     name: "datastore1",
 * }));
 * const host = dc.apply(dc => vsphere.getHost({
 *     datacenterId: dc.id,
 *     name: "esxi1",
 * }));
 * const network = dc.apply(dc => vsphere.getNetwork({
 *     datacenterId: dc.id,
 *     name: "network1",
 * }));
 * const clusterHostGroup = new vsphere.ComputeClusterHostGroup("cluster_host_group", {
 *     computeClusterId: cluster.id,
 *     hostSystemIds: [host.id],
 * });
 * const vm = new vsphere.VirtualMachine("vm", {
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
 * const clusterVmGroup = new vsphere.ComputeClusterVmGroup("cluster_vm_group", {
 *     computeClusterId: cluster.id,
 *     virtualMachineIds: [vm.id],
 * });
 * const clusterVmHostRule = new vsphere.ComputeClusterVmHostRule("cluster_vm_host_rule", {
 *     affinityHostGroupName: clusterHostGroup.name,
 *     computeClusterId: cluster.id,
 *     vmGroupName: clusterVmGroup.name,
 * });
 * ```
 */
export class ComputeClusterVmHostRule extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmHostRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmHostRuleState, opts?: pulumi.CustomResourceOptions): ComputeClusterVmHostRule {
        return new ComputeClusterVmHostRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/computeClusterVmHostRule:ComputeClusterVmHostRule';

    /**
     * Returns true if the given object is an instance of ComputeClusterVmHostRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeClusterVmHostRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeClusterVmHostRule.__pulumiType;
    }

    /**
     * When this field is used, the virtual
     * machines defined in `vm_group_name` will be run on the
     * hosts defined in this host group.
     */
    public readonly affinityHostGroupName!: pulumi.Output<string | undefined>;
    /**
     * When this field is used, the
     * virtual machines defined in `vm_group_name` will _not_ be
     * run on the hosts defined in this host group.
     */
    public readonly antiAffinityHostGroupName!: pulumi.Output<string | undefined>;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
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
     */
    public readonly mandatory!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the rule. This must be unique in the
     * cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the virtual machine group to use
     * with this rule.
     */
    public readonly vmGroupName!: pulumi.Output<string>;

    /**
     * Create a ComputeClusterVmHostRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmHostRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeClusterVmHostRuleArgs | ComputeClusterVmHostRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ComputeClusterVmHostRuleState | undefined;
            inputs["affinityHostGroupName"] = state ? state.affinityHostGroupName : undefined;
            inputs["antiAffinityHostGroupName"] = state ? state.antiAffinityHostGroupName : undefined;
            inputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["mandatory"] = state ? state.mandatory : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["vmGroupName"] = state ? state.vmGroupName : undefined;
        } else {
            const args = argsOrState as ComputeClusterVmHostRuleArgs | undefined;
            if (!args || args.computeClusterId === undefined) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            if (!args || args.vmGroupName === undefined) {
                throw new Error("Missing required property 'vmGroupName'");
            }
            inputs["affinityHostGroupName"] = args ? args.affinityHostGroupName : undefined;
            inputs["antiAffinityHostGroupName"] = args ? args.antiAffinityHostGroupName : undefined;
            inputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["mandatory"] = args ? args.mandatory : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["vmGroupName"] = args ? args.vmGroupName : undefined;
        }
        super(ComputeClusterVmHostRule.__pulumiType, name, inputs, opts);
    }
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
