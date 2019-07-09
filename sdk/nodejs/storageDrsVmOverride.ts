// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere_storage_drs_vm_override` resource can be used to add a Storage DRS
 * override to a datastore cluster for a specific virtual machine. With this
 * resource, one can enable or disable Storage DRS, and control the automation
 * level and disk affinity for a single virtual machine without affecting the rest
 * of the datastore cluster.
 * 
 * For more information on vSphere datastore clusters and Storage DRS, see [this
 * page][ref-vsphere-datastore-clusters].
 * 
 * [ref-vsphere-datastore-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-598DF695-107E-406B-9C95-0AF961FC227A.html
 * 
 * ## Example Usage
 * 
 * The example below builds on the [Storage DRS
 * example][tf-vsphere-vm-storage-drs-example] in the `vsphere_virtual_machine`
 * resource. However, rather than use the output of the
 * [`vsphere_datastore_cluster` data
 * source][tf-vsphere-datastore-cluster-data-source] for the location of the
 * virtual machine, we instead get what is assumed to be a member datastore using
 * the [`vsphere_datastore` data source][tf-vsphere-datastore-data-source] and put
 * the virtual machine there instead. We then use the
 * `vsphere_storage_drs_vm_override` resource to ensure that Storage DRS does not
 * apply to this virtual machine, and hence the VM will never be migrated off of
 * the datastore.
 * 
 * [tf-vsphere-vm-storage-drs-example]: /docs/providers/vsphere/r/virtual_machine.html#using-storage-drs
 * [tf-vsphere-datastore-cluster-data-source]: /docs/providers/vsphere/d/datastore_cluster.html
 * [tf-vsphere-datastore-data-source]: /docs/providers/vsphere/d/datastore.html
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const dc = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }));
 * const memberDatastore = dc.apply(dc => vsphere.getDatastore({
 *     datacenterId: dc.id,
 *     name: "datastore-cluster1-member1",
 * }));
 * const datastoreCluster = dc.apply(dc => vsphere.getDatastoreCluster({
 *     datacenterId: dc.id,
 *     name: "datastore-cluster1",
 * }));
 * const network = dc.apply(dc => vsphere.getNetwork({
 *     datacenterId: dc.id,
 *     name: "public",
 * }));
 * const pool = dc.apply(dc => vsphere.getResourcePool({
 *     datacenterId: dc.id,
 *     name: "cluster1/Resources",
 * }));
 * const vm = new vsphere.VirtualMachine("vm", {
 *     datastoreId: memberDatastore.id,
 *     disks: [{
 *         label: "disk0",
 *         size: 20,
 *     }],
 *     guestId: "other3xLinux64Guest",
 *     memory: 1024,
 *     networkInterfaces: [{
 *         networkId: network.id,
 *     }],
 *     numCpus: 2,
 *     resourcePoolId: pool.id,
 * });
 * const drsVmOverride = new vsphere.StorageDrsVmOverride("drs_vm_override", {
 *     datastoreClusterId: datastoreCluster.id,
 *     sdrsEnabled: "false",
 *     virtualMachineId: vm.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/storage_drs_vm_override.html.markdown.
 */
export class StorageDrsVmOverride extends pulumi.CustomResource {
    /**
     * Get an existing StorageDrsVmOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StorageDrsVmOverrideState, opts?: pulumi.CustomResourceOptions): StorageDrsVmOverride {
        return new StorageDrsVmOverride(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/storageDrsVmOverride:StorageDrsVmOverride';

    /**
     * Returns true if the given object is an instance of StorageDrsVmOverride.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StorageDrsVmOverride {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StorageDrsVmOverride.__pulumiType;
    }

    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the override in.
     * Forces a new resource if changed.
     */
    public readonly datastoreClusterId!: pulumi.Output<string>;
    /**
     * Overrides any Storage DRS automation
     * levels for this virtual machine. Can be one of `automated` or `manual`. When
     * not specified, the datastore cluster's settings are used according to the
     * [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels].
     */
    public readonly sdrsAutomationLevel!: pulumi.Output<string | undefined>;
    /**
     * Overrides the default Storage DRS setting for
     * this virtual machine. When not specified, the datastore cluster setting is
     * used.
     */
    public readonly sdrsEnabled!: pulumi.Output<string | undefined>;
    /**
     * Overrides the intra-VM affinity setting
     * for this virtual machine. When `true`, all disks for this virtual machine
     * will be kept on the same datastore. When `false`, Storage DRS may locate
     * individual disks on different datastores if it helps satisfy cluster
     * requirements. When not specified, the datastore cluster's settings are used.
     */
    public readonly sdrsIntraVmAffinity!: pulumi.Output<string | undefined>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    public readonly virtualMachineId!: pulumi.Output<string>;

    /**
     * Create a StorageDrsVmOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageDrsVmOverrideArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StorageDrsVmOverrideArgs | StorageDrsVmOverrideState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StorageDrsVmOverrideState | undefined;
            inputs["datastoreClusterId"] = state ? state.datastoreClusterId : undefined;
            inputs["sdrsAutomationLevel"] = state ? state.sdrsAutomationLevel : undefined;
            inputs["sdrsEnabled"] = state ? state.sdrsEnabled : undefined;
            inputs["sdrsIntraVmAffinity"] = state ? state.sdrsIntraVmAffinity : undefined;
            inputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as StorageDrsVmOverrideArgs | undefined;
            if (!args || args.datastoreClusterId === undefined) {
                throw new Error("Missing required property 'datastoreClusterId'");
            }
            if (!args || args.virtualMachineId === undefined) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            inputs["datastoreClusterId"] = args ? args.datastoreClusterId : undefined;
            inputs["sdrsAutomationLevel"] = args ? args.sdrsAutomationLevel : undefined;
            inputs["sdrsEnabled"] = args ? args.sdrsEnabled : undefined;
            inputs["sdrsIntraVmAffinity"] = args ? args.sdrsIntraVmAffinity : undefined;
            inputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        super(StorageDrsVmOverride.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StorageDrsVmOverride resources.
 */
export interface StorageDrsVmOverrideState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the override in.
     * Forces a new resource if changed.
     */
    readonly datastoreClusterId?: pulumi.Input<string>;
    /**
     * Overrides any Storage DRS automation
     * levels for this virtual machine. Can be one of `automated` or `manual`. When
     * not specified, the datastore cluster's settings are used according to the
     * [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels].
     */
    readonly sdrsAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default Storage DRS setting for
     * this virtual machine. When not specified, the datastore cluster setting is
     * used.
     */
    readonly sdrsEnabled?: pulumi.Input<string>;
    /**
     * Overrides the intra-VM affinity setting
     * for this virtual machine. When `true`, all disks for this virtual machine
     * will be kept on the same datastore. When `false`, Storage DRS may locate
     * individual disks on different datastores if it helps satisfy cluster
     * requirements. When not specified, the datastore cluster's settings are used.
     */
    readonly sdrsIntraVmAffinity?: pulumi.Input<string>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StorageDrsVmOverride resource.
 */
export interface StorageDrsVmOverrideArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the override in.
     * Forces a new resource if changed.
     */
    readonly datastoreClusterId: pulumi.Input<string>;
    /**
     * Overrides any Storage DRS automation
     * levels for this virtual machine. Can be one of `automated` or `manual`. When
     * not specified, the datastore cluster's settings are used according to the
     * [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels].
     */
    readonly sdrsAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default Storage DRS setting for
     * this virtual machine. When not specified, the datastore cluster setting is
     * used.
     */
    readonly sdrsEnabled?: pulumi.Input<string>;
    /**
     * Overrides the intra-VM affinity setting
     * for this virtual machine. When `true`, all disks for this virtual machine
     * will be kept on the same datastore. When `false`, Storage DRS may locate
     * individual disks on different datastores if it helps satisfy cluster
     * requirements. When not specified, the datastore cluster's settings are used.
     */
    readonly sdrsIntraVmAffinity?: pulumi.Input<string>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId: pulumi.Input<string>;
}
