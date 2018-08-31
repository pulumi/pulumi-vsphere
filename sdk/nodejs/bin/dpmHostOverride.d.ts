import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_dpm_host_override` resource can be used to add a DPM override to a
 * cluster for a particular host. This allows you to control the power management
 * settings for individual hosts in the cluster while leaving any unspecified ones
 * at the default power management settings.
 *
 * For more information on DPM within vSphere clusters, see [this
 * page][ref-vsphere-cluster-dpm].
 *
 * [ref-vsphere-cluster-dpm]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-5E5E349A-4644-4C9C-B434-1C0243EBDC80.html
 *
 * ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
 * connections.
 *
 * ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
 */
export declare class DpmHostOverride extends pulumi.CustomResource {
    /**
     * Get an existing DpmHostOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DpmHostOverrideState): DpmHostOverride;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Output<string>;
    /**
     * The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     */
    readonly dpmAutomationLevel: pulumi.Output<string | undefined>;
    /**
     * Enable DPM support for this host. Default:
     * `false`.
     */
    readonly dpmEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The managed object ID of the host.
     */
    readonly hostSystemId: pulumi.Output<string>;
    /**
     * Create a DpmHostOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DpmHostOverrideArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering DpmHostOverride resources.
 */
export interface DpmHostOverrideState {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId?: pulumi.Input<string>;
    /**
     * The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     */
    readonly dpmAutomationLevel?: pulumi.Input<string>;
    /**
     * Enable DPM support for this host. Default:
     * `false`.
     */
    readonly dpmEnabled?: pulumi.Input<boolean>;
    /**
     * The managed object ID of the host.
     */
    readonly hostSystemId?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a DpmHostOverride resource.
 */
export interface DpmHostOverrideArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
     * resource if changed.
     */
    readonly computeClusterId: pulumi.Input<string>;
    /**
     * The automation level for host power
     * operations on this host. Can be one of `manual` or `automated`. Default:
     * `manual`.
     */
    readonly dpmAutomationLevel?: pulumi.Input<string>;
    /**
     * Enable DPM support for this host. Default:
     * `false`.
     */
    readonly dpmEnabled?: pulumi.Input<boolean>;
    /**
     * The managed object ID of the host.
     */
    readonly hostSystemId: pulumi.Input<string>;
}
