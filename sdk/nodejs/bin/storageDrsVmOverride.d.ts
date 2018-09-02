import * as pulumi from "@pulumi/pulumi";
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
 */
export declare class StorageDrsVmOverride extends pulumi.CustomResource {
    /**
     * Get an existing StorageDrsVmOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StorageDrsVmOverrideState): StorageDrsVmOverride;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster to put the override in.
     * Forces a new resource if changed.
     */
    readonly datastoreClusterId: pulumi.Output<string>;
    /**
     * Overrides any Storage DRS automation
     * levels for this virtual machine. Can be one of `automated` or `manual`. When
     * not specified, the datastore cluster's settings are used according to the
     * [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels].
     */
    readonly sdrsAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * Overrides the default Storage DRS setting for
     * this virtual machine. When not specified, the datastore cluster setting is
     * used.
     */
    readonly sdrsEnabled: pulumi.Output<string | undefined>;
    /**
     * Overrides the intra-VM affinity setting
     * for this virtual machine. When `true`, all disks for this virtual machine
     * will be kept on the same datastore. When `false`, Storage DRS may locate
     * individual disks on different datastores if it helps satisfy cluster
     * requirements. When not specified, the datastore cluster's settings are used.
     */
    readonly sdrsIntraVmAffinity: pulumi.Output<string | undefined>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId: pulumi.Output<string>;
    /**
     * Create a StorageDrsVmOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StorageDrsVmOverrideArgs, opts?: pulumi.CustomResourceOptions);
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
