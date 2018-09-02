import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_tag_category` resource can be used to create and manage tag
 * categories, which determine how tags are grouped together and applied to
 * specific objects.
 *
 * For more information about tags, click [here][ext-tags-general]. For more
 * information about tag categories specifically, click
 * [here][ext-tag-categories].
 *
 * [ext-tags-general]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-E8E854DD-AA97-4E0C-8419-CE84F93C4058.html
 * [ext-tag-categories]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-BA3D1794-28F2-43F3-BCE9-3964CB207FB6.html
 *
 * ~> **NOTE:** Tagging support is unsupported on direct ESXi connections and
 * requires vCenter 6.0 or higher.
 */
export declare class TagCategory extends pulumi.CustomResource {
    /**
     * Get an existing TagCategory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagCategoryState): TagCategory;
    /**
     * A list object types that this category is
     * valid to be assigned to. For a full list, click
     * here.
     */
    readonly associableTypes: pulumi.Output<string[]>;
    /**
     * The number of tags that can be assigned from this
     * category to a single object at once. Can be one of `SINGLE` (object can only
     * be assigned one tag in this category), to `MULTIPLE` (object can be assigned
     * multiple tags in this category). Forces a new resource if changed.
     */
    readonly cardinality: pulumi.Output<string>;
    /**
     * A description for the category.
     */
    readonly description: pulumi.Output<string | undefined>;
    /**
     * The name of the category.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Create a TagCategory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagCategoryArgs, opts?: pulumi.CustomResourceOptions);
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
    readonly associableTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of tags that can be assigned from this
     * category to a single object at once. Can be one of `SINGLE` (object can only
     * be assigned one tag in this category), to `MULTIPLE` (object can be assigned
     * multiple tags in this category). Forces a new resource if changed.
     */
    readonly cardinality?: pulumi.Input<string>;
    /**
     * A description for the category.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the category.
     */
    readonly name?: pulumi.Input<string>;
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
    readonly associableTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of tags that can be assigned from this
     * category to a single object at once. Can be one of `SINGLE` (object can only
     * be assigned one tag in this category), to `MULTIPLE` (object can be assigned
     * multiple tags in this category). Forces a new resource if changed.
     */
    readonly cardinality: pulumi.Input<string>;
    /**
     * A description for the category.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the category.
     */
    readonly name?: pulumi.Input<string>;
}
