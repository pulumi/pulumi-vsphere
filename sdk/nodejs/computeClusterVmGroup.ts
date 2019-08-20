// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere..ComputeClusterVmGroup` resource can be used to manage groups of
 * virtual machines in a cluster, either created by the
 * [`vsphere..ComputeCluster`][tf-vsphere-cluster-resource] resource or looked up
 * by the [`vsphere..ComputeCluster`][tf-vsphere-cluster-data-source] data source.
 * 
 * [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
 * [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
 * 
 * This resource mainly serves as an input to the
 * [`vsphere..ComputeClusterVmDependencyRule`][tf-vsphere-cluster-vm-dependency-rule-resource]
 * and
 * [`vsphere..ComputeClusterVmHostRule`][tf-vsphere-cluster-vm-host-rule-resource]
 * resources. See the individual resource documentation pages for more information.
 * 
 * [tf-vsphere-cluster-vm-dependency-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_dependency_rule.html
 * [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
 * 
 * > **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 * 
 * > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 * 
 * ## Example Usage
 * 
 * The example below creates two virtual machines in a cluster using the
 * [`vsphere..VirtualMachine`][tf-vsphere-vm-resource] resource, creating the
 * virtual machine in the cluster looked up by the
 * [`vsphere..ComputeCluster`][tf-vsphere-cluster-data-source] data source. It
 * then creates a group from these two virtual machines.
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
 * const datastore = vsphere.getDatastore({
 *     datacenterId: dc.id,
 *     name: "datastore1",
 * });
 * const network = vsphere.getNetwork({
 *     datacenterId: dc.id,
 *     name: "network1",
 * });
 * const vm: vsphere.VirtualMachine[] = [];
 * for (let i = 0; i < 2; i++) {
 *     vm.push(new vsphere.VirtualMachine(`vm-${i}`, {
 *         datastoreId: datastore.id,
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
 * const clusterVmGroup = new vsphere.ComputeClusterVmGroup("clusterVmGroup", {
 *     computeClusterId: cluster.id,
 *     virtualMachineIds: vm.map(v => v.id),
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_vm_group.html.markdown.
 */
export class ComputeClusterVmGroup extends pulumi.CustomResource {
    /**
     * Get an existing ComputeClusterVmGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ComputeClusterVmGroupState, opts?: pulumi.CustomResourceOptions): ComputeClusterVmGroup {
        return new ComputeClusterVmGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/computeClusterVmGroup:ComputeClusterVmGroup';

    /**
     * Returns true if the given object is an instance of ComputeClusterVmGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ComputeClusterVmGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ComputeClusterVmGroup.__pulumiType;
    }

    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
     * resource if changed.
     */
    public readonly computeClusterId!: pulumi.Output<string>;
    /**
     * The name of the VM group. This must be unique in the
     * cluster. Forces a new resource if changed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The UUIDs of the virtual machines in this
     * group.
     */
    public readonly virtualMachineIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ComputeClusterVmGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComputeClusterVmGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ComputeClusterVmGroupArgs | ComputeClusterVmGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ComputeClusterVmGroupState | undefined;
            inputs["computeClusterId"] = state ? state.computeClusterId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["virtualMachineIds"] = state ? state.virtualMachineIds : undefined;
        } else {
            const args = argsOrState as ComputeClusterVmGroupArgs | undefined;
            if (!args || args.computeClusterId === undefined) {
                throw new Error("Missing required property 'computeClusterId'");
            }
            inputs["computeClusterId"] = args ? args.computeClusterId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["virtualMachineIds"] = args ? args.virtualMachineIds : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ComputeClusterVmGroup.__pulumiType, name, inputs, opts);
    }
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
