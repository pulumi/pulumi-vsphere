// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.TagCategory` resource can be used to create and manage tag
 * categories, which determine how tags are grouped together and applied to
 * specific objects.
 *
 * For more information about tags, click [here][ext-tags-general].
 *
 * [ext-tags-general]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-vcenter-esxi-management/GUID-E8E854DD-AA97-4E0C-8419-CE84F93C4058.html
 *
 * ## Example Usage
 *
 * This example creates a tag category named `test-category`, with
 * single cardinality (meaning that only one tag in this category can be assigned
 * to an object at any given time). Tags in this category can only be assigned to
 * VMs and datastores.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const category = new vsphere.TagCategory("category", {
 *     name: "test-category",
 *     description: "Managed by Pulumi",
 *     cardinality: "SINGLE",
 *     associableTypes: [
 *         "VirtualMachine",
 *         "Datastore",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * An existing tag category can be imported into this resource via
 *
 * its name, using the following command:
 *
 * ```sh
 * $ pulumi import vsphere:index/tagCategory:TagCategory category terraform-test-category
 * ```
 */
export class TagCategory extends pulumi.CustomResource {
    /**
     * Get an existing TagCategory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagCategoryState, opts?: pulumi.CustomResourceOptions): TagCategory {
        return new TagCategory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/tagCategory:TagCategory';

    /**
     * Returns true if the given object is an instance of TagCategory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagCategory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagCategory.__pulumiType;
    }

    /**
     * A list object types that this category is
     * valid to be assigned to. For a full list, click
     * here.
     */
    public readonly associableTypes!: pulumi.Output<string[]>;
    /**
     * The number of tags that can be assigned from this
     * category to a single object at once. Can be one of `SINGLE` (object can only
     * be assigned one tag in this category), to `MULTIPLE` (object can be assigned
     * multiple tags in this category). Forces a new resource if changed.
     */
    public readonly cardinality!: pulumi.Output<string>;
    /**
     * A description for the category.
     *
     * > **NOTE:** You can add associable types to a category, but you cannot remove
     * them. Attempting to do so will result in an error.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the category.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a TagCategory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagCategoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagCategoryArgs | TagCategoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagCategoryState | undefined;
            resourceInputs["associableTypes"] = state ? state.associableTypes : undefined;
            resourceInputs["cardinality"] = state ? state.cardinality : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as TagCategoryArgs | undefined;
            if ((!args || args.associableTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'associableTypes'");
            }
            if ((!args || args.cardinality === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cardinality'");
            }
            resourceInputs["associableTypes"] = args ? args.associableTypes : undefined;
            resourceInputs["cardinality"] = args ? args.cardinality : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TagCategory.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagCategory resources.
 */
export interface TagCategoryState {
    /**
     * A list object types that this category is
     * valid to be assigned to. For a full list, click
     * here.
     */
    associableTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of tags that can be assigned from this
     * category to a single object at once. Can be one of `SINGLE` (object can only
     * be assigned one tag in this category), to `MULTIPLE` (object can be assigned
     * multiple tags in this category). Forces a new resource if changed.
     */
    cardinality?: pulumi.Input<string>;
    /**
     * A description for the category.
     *
     * > **NOTE:** You can add associable types to a category, but you cannot remove
     * them. Attempting to do so will result in an error.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the category.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TagCategory resource.
 */
export interface TagCategoryArgs {
    /**
     * A list object types that this category is
     * valid to be assigned to. For a full list, click
     * here.
     */
    associableTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of tags that can be assigned from this
     * category to a single object at once. Can be one of `SINGLE` (object can only
     * be assigned one tag in this category), to `MULTIPLE` (object can be assigned
     * multiple tags in this category). Forces a new resource if changed.
     */
    cardinality: pulumi.Input<string>;
    /**
     * A description for the category.
     *
     * > **NOTE:** You can add associable types to a category, but you cannot remove
     * them. Attempting to do so will result in an error.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the category.
     */
    name?: pulumi.Input<string>;
}
