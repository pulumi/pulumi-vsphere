import * as pulumi from "@pulumi/pulumi";
/**
 * Provides a VMware vSphere datacenter resource. This can be used as the primary
 * container of inventory objects such as hosts and virtual machines.
 */
export declare class Datacenter extends pulumi.CustomResource {
    /**
     * Get an existing Datacenter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatacenterState): Datacenter;
    /**
     * Map of custom attribute ids to value
     * strings to set for datacenter resource. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The folder where the datacenter should be created.
     * Forces a new resource if changed.
     */
    readonly folder: pulumi.Output<string | undefined>;
    /**
     * [Managed object ID][docs-about-morefs] of this datacenter.
     */
    readonly moid: pulumi.Output<string>;
    /**
     * The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * Create a Datacenter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatacenterArgs, opts?: pulumi.CustomResourceOptions);
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
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The folder where the datacenter should be created.
     * Forces a new resource if changed.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * [Managed object ID][docs-about-morefs] of this datacenter.
     */
    readonly moid?: pulumi.Input<string>;
    /**
     * The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
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
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The folder where the datacenter should be created.
     * Forces a new resource if changed.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The name of the datacenter. This name needs to be unique
     * within the folder. Forces a new resource if changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
