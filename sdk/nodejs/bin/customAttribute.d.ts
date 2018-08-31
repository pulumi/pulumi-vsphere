import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_custom_attribute` resource can be used to create and manage custom
 * attributes, which allow users to associate user-specific meta-information with
 * vSphere managed objects. Custom attribute values must be strings and are stored
 * on the vCenter Server and not the managed object.
 *
 * For more information about custom attributes, click [here][ext-custom-attributes].
 *
 * [ext-custom-attributes]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-73606C4C-763C-4E27-A1DA-032E4C46219D.html
 *
 * ~> **NOTE:** Custom attributes are unsupported on direct ESXi connections
 * and require vCenter.
 */
export declare class CustomAttribute extends pulumi.CustomResource {
    /**
     * Get an existing CustomAttribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomAttributeState): CustomAttribute;
    /**
     * The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, click here. Forces a new
     * resource if changed.
     */
    readonly managedObjectType: pulumi.Output<string | undefined>;
    /**
     * The name of the custom attribute.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Create a CustomAttribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CustomAttributeArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering CustomAttribute resources.
 */
export interface CustomAttributeState {
    /**
     * The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, click here. Forces a new
     * resource if changed.
     */
    readonly managedObjectType?: pulumi.Input<string>;
    /**
     * The name of the custom attribute.
     */
    readonly name?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a CustomAttribute resource.
 */
export interface CustomAttributeArgs {
    /**
     * The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, click here. Forces a new
     * resource if changed.
     */
    readonly managedObjectType?: pulumi.Input<string>;
    /**
     * The name of the custom attribute.
     */
    readonly name?: pulumi.Input<string>;
}
