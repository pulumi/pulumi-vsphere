import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_drs_vm_override` resource can be used to add a DRS override to a
 * cluster for a specific virtual machine. With this resource, one can enable or
 * disable DRS and control the automation level for a single virtual machine
 * without affecting the rest of the cluster.
 *
 * For more information on vSphere clusters and DRS, see [this
 * page][ref-vsphere-drs-clusters].
 *
 * [ref-vsphere-drs-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-8ACF3502-5314-469F-8CC9-4A9BD5925BC2.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 */
export declare class DrsVmOverride extends pulumi.CustomResource {
    /**
     * Get an existing DrsVmOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DrsVmOverrideState): DrsVmOverride;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Output<string>;
    /**
     * Overrides the automation level for this virtual
     * machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
     * `fullyAutomated`. Default: `manual`.
     */
    readonly drsAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * Overrides the default DRS setting for this virtual
     * machine. Can be either `true` or `false`. Default: `false`.
     */
    readonly drsEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId: pulumi.Output<string>;
    /**
     * Create a DrsVmOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DrsVmOverrideArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering DrsVmOverride resources.
 */
export interface DrsVmOverrideState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * Overrides the automation level for this virtual
     * machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
     * `fullyAutomated`. Default: `manual`.
     */
    readonly drsAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default DRS setting for this virtual
     * machine. Can be either `true` or `false`. Default: `false`.
     */
    readonly drsEnabled?: pulumi.Input<boolean>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a DrsVmOverride resource.
 */
export interface DrsVmOverrideArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * Overrides the automation level for this virtual
     * machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
     * `fullyAutomated`. Default: `manual`.
     */
    readonly drsAutomationLevel?: pulumi.Input<string>;
    /**
     * Overrides the default DRS setting for this virtual
     * machine. Can be either `true` or `false`. Default: `false`.
     */
    readonly drsEnabled?: pulumi.Input<boolean>;
    /**
     * The UUID of the virtual machine to create
     * the override for.  Forces a new resource if changed.
     */
    readonly virtualMachineId: pulumi.Input<string>;
}
