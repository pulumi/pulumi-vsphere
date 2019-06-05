// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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
 * > **NOTE:** Custom attributes are unsupported on direct ESXi connections 
 * and require vCenter.
 * 
 * ## Example Usage
 * 
 * This example creates a custom attribute named `terraform-test-attribute`. The 
 * resulting custom attribute can be assigned to VMs only.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const attribute = new vsphere.CustomAttribute("attribute", {
 *     managedObjectType: "VirtualMachine",
 * });
 * ```
 * 
 * ## Managed Object Types
 * 
 * The following table will help you determine what value you need to enter for 
 * the managed object type you want the attribute to apply to.
 * 
 * Note that if you want a attribute to apply to all objects, leave the type 
 * unspecified.
 * 
 * <table>
 * <tr><th>Type</th><th>Value</th></tr>
 * <tr><td>Folders</td><td>`Folder`</td></tr>
 * <tr><td>Clusters</td><td>`ClusterComputeResource`</td></tr>
 * <tr><td>Datacenters</td><td>`Datacenter`</td></tr>
 * <tr><td>Datastores</td><td>`Datastore`</td></tr>
 * <tr><td>Datastore Clusters</td><td>`StoragePod`</td></tr>
 * <tr><td>DVS Portgroups</td><td>`DistributedVirtualPortgroup`</td></tr>
 * <tr><td>Distributed vSwitches</td><td>`DistributedVirtualSwitch`<br>`VmwareDistributedVirtualSwitch`</td></tr>
 * <tr><td>Hosts</td><td>`HostSystem`</td></tr>
 * <tr><td>Content Libraries</td><td>`com.vmware.content.Library`</td></tr>
 * <tr><td>Content Library Items</td><td>`com.vmware.content.library.Item`</td></tr>
 * <tr><td>Networks</td><td>`HostNetwork`<br>`Network`<br>`OpaqueNetwork`</td></tr>
 * <tr><td>Resource Pools</td><td>`ResourcePool`</td></tr>
 * <tr><td>vApps</td><td>`VirtualApp`</td></tr>
 * <tr><td>Virtual Machines</td><td>`VirtualMachine`</td></tr>
 * </table>
 * 
 * ## Importing
 * 
 * An existing custom attribute can be [imported][docs-import] into this resource 
 * via its name, using the following command:
 * 
 * [docs-import]: https://www.terraform.io/docs/import/index.html
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 */
export class CustomAttribute extends pulumi.CustomResource {
    /**
     * Get an existing CustomAttribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomAttributeState, opts?: pulumi.CustomResourceOptions): CustomAttribute {
        return new CustomAttribute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/customAttribute:CustomAttribute';

    /**
     * Returns true if the given object is an instance of CustomAttribute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomAttribute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomAttribute.__pulumiType;
    }

    /**
     * The object type that this attribute may be
     * applied to. If not set, the custom attribute may be applied to any object
     * type. For a full list, click here. Forces a new
     * resource if changed.
     */
    public readonly managedObjectType!: pulumi.Output<string | undefined>;
    /**
     * The name of the custom attribute.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a CustomAttribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CustomAttributeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomAttributeArgs | CustomAttributeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CustomAttributeState | undefined;
            inputs["managedObjectType"] = state ? state.managedObjectType : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as CustomAttributeArgs | undefined;
            inputs["managedObjectType"] = args ? args.managedObjectType : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        super(CustomAttribute.__pulumiType, name, inputs, opts);
    }
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
