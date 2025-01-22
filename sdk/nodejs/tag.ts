// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.Tag` resource can be used to create and manage tags, which allow
 * you to attach metadata to objects in the vSphere inventory to make these
 * objects more sortable and searchable.
 *
 * For more information about tags, click [here][ext-tags-general].
 *
 * [ext-tags-general]: https://techdocs.broadcom.com/us/en/vmware-cis/vsphere/vsphere/8-0/vsphere-tags-and-attributes.html
 *
 * ## Example Usage
 *
 * This example creates a tag named `test-tag`. This tag is assigned the
 * `test-category` category, which was created by the
 * `vsphere.TagCategory` resource. The resulting
 * tag can be assigned to VMs and datastores only, and can be the only value in
 * the category that can be assigned, as per the restrictions defined by the
 * category.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const category = new vsphere.TagCategory("category", {
 *     name: "test-category",
 *     cardinality: "SINGLE",
 *     description: "Managed by Pulumi",
 *     associableTypes: [
 *         "VirtualMachine",
 *         "Datastore",
 *     ],
 * });
 * const tag = new vsphere.Tag("tag", {
 *     name: "test-tag",
 *     categoryId: category.id,
 *     description: "Managed by Pulumi",
 * });
 * ```
 *
 * ### Using Tags in a Supported Resource
 *
 * Tags can be applied to vSphere resources via the `tags` argument
 * in any supported resource.
 *
 * The following example builds on the above example by creating a
 * `vsphere.VirtualMachine` and applying the
 * created tag to it:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const category = new vsphere.TagCategory("category", {
 *     name: "test-category",
 *     cardinality: "SINGLE",
 *     description: "Managed by Pulumi",
 *     associableTypes: [
 *         "VirtualMachine",
 *         "Datastore",
 *     ],
 * });
 * const tag = new vsphere.Tag("tag", {
 *     name: "test-tag",
 *     categoryId: category.id,
 *     description: "Managed by Pulumi",
 * });
 * const web = new vsphere.VirtualMachine("web", {tags: [tag.id]});
 * ```
 *
 * ## Import
 *
 * An existing tag can be imported into this resource by supplying
 *
 * both the tag's category name and the name of the tag as a JSON string to
 *
 * `pulumi import`, as per the example below:
 *
 * ```sh
 * $ pulumi import vsphere:index/tag:Tag tag \
 * ```
 *
 *   '{"category_name": "pulumi-test-category", "tag_name": "pulumi-test-tag"}'
 */
export class Tag extends pulumi.CustomResource {
    /**
     * Get an existing Tag resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagState, opts?: pulumi.CustomResourceOptions): Tag {
        return new Tag(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/tag:Tag';

    /**
     * Returns true if the given object is an instance of Tag.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tag {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tag.__pulumiType;
    }

    /**
     * The unique identifier of the parent category in
     * which this tag will be created. Forces a new resource if changed.
     */
    public readonly categoryId!: pulumi.Output<string>;
    /**
     * A description for the tag.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The display name of the tag. The name must be unique
     * within its category.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Tag resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagArgs | TagState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagState | undefined;
            resourceInputs["categoryId"] = state ? state.categoryId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as TagArgs | undefined;
            if ((!args || args.categoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'categoryId'");
            }
            resourceInputs["categoryId"] = args ? args.categoryId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tag.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tag resources.
 */
export interface TagState {
    /**
     * The unique identifier of the parent category in
     * which this tag will be created. Forces a new resource if changed.
     */
    categoryId?: pulumi.Input<string>;
    /**
     * A description for the tag.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the tag. The name must be unique
     * within its category.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tag resource.
 */
export interface TagArgs {
    /**
     * The unique identifier of the parent category in
     * which this tag will be created. Forces a new resource if changed.
     */
    categoryId: pulumi.Input<string>;
    /**
     * A description for the tag.
     */
    description?: pulumi.Input<string>;
    /**
     * The display name of the tag. The name must be unique
     * within its category.
     */
    name?: pulumi.Input<string>;
}
