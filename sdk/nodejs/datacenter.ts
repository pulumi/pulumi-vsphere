// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a VMware vSphere datacenter resource. This can be used as the primary
 * container of inventory objects such as hosts and virtual machines.
 *
 * ## Example Usage
 * ### Create datacenter on the root folder
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const prodDatacenter = new vsphere.Datacenter("prodDatacenter", {});
 * ```
 * ### Create datacenter on a subfolder
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const researchDatacenter = new vsphere.Datacenter("researchDatacenter", {folder: "/research/"});
 * ```
 */
export class Datacenter extends pulumi.CustomResource {
    /**
     * Get an existing Datacenter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatacenterState, opts?: pulumi.CustomResourceOptions): Datacenter {
        return new Datacenter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/datacenter:Datacenter';

    /**
     * Returns true if the given object is an instance of Datacenter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Datacenter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Datacenter.__pulumiType;
    }

    /**
     * Map of custom attribute ids to value 
     * strings to set for datacenter resource. See 
     * [here][docs-setting-custom-attributes] for a reference on how to set values 
     * for custom attributes.
     *
     *
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections 
     * and require vCenter.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The folder where the datacenter should be created.
     * Forces a new resource if changed.
     */
    public readonly folder!: pulumi.Output<string | undefined>;
    /**
     * Managed object ID of this datacenter.
     */
    public /*out*/ readonly moid!: pulumi.Output<string>;
    /**
     * The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource.
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Datacenter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatacenterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatacenterArgs | DatacenterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatacenterState | undefined;
            resourceInputs["customAttributes"] = state ? state.customAttributes : undefined;
            resourceInputs["folder"] = state ? state.folder : undefined;
            resourceInputs["moid"] = state ? state.moid : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DatacenterArgs | undefined;
            resourceInputs["customAttributes"] = args ? args.customAttributes : undefined;
            resourceInputs["folder"] = args ? args.folder : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["moid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Datacenter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Datacenter resources.
 */
export interface DatacenterState {
    /**
     * Map of custom attribute ids to value 
     * strings to set for datacenter resource. See 
     * [here][docs-setting-custom-attributes] for a reference on how to set values 
     * for custom attributes.
     *
     *
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections 
     * and require vCenter.
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The folder where the datacenter should be created.
     * Forces a new resource if changed.
     */
    folder?: pulumi.Input<string>;
    /**
     * Managed object ID of this datacenter.
     */
    moid?: pulumi.Input<string>;
    /**
     * The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource.
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Datacenter resource.
 */
export interface DatacenterArgs {
    /**
     * Map of custom attribute ids to value 
     * strings to set for datacenter resource. See 
     * [here][docs-setting-custom-attributes] for a reference on how to set values 
     * for custom attributes.
     *
     *
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections 
     * and require vCenter.
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The folder where the datacenter should be created.
     * Forces a new resource if changed.
     */
    folder?: pulumi.Input<string>;
    /**
     * The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource.
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
