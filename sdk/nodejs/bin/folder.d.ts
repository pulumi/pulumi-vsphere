import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_folder` resource can be used to manage vSphere inventory folders.
 * The resource supports creating folders of the 5 major types - datacenter
 * folders, host and cluster folders, virtual machine folders, datastore folders,
 * and network folders.
 *
 * Paths are always relative to the specific type of folder you are creating.
 * Subfolders are discovered by parsing the relative path specified in `path`, so
 * `foo/bar` will create a folder named `bar` in the parent folder `foo`, as long
 * as that folder exists.
 */
export declare class Folder extends pulumi.CustomResource {
    /**
     * Get an existing Folder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FolderState): Folder;
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for folder. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     */
    readonly datacenterId: pulumi.Output<string | undefined>;
    /**
     * The path of the folder to be created. This is relative to
     * the root of the type of folder you are creating, and the supplied datacenter.
     * For example, given a default datacenter of `default-dc`, a folder of type
     * `vm` (denoting a virtual machine folder), and a supplied folder of
     * `terraform-test-folder`, the resulting path would be
     * `/default-dc/vm/terraform-test-folder`.
     */
    readonly path: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * The type of folder to create. Allowed options are
     * `datacenter` for datacenter folders, `host` for host and cluster folders,
     * `vm` for virtual machine folders, `datastore` for datastore folders, and
     * `network` for network folders. Forces a new resource if changed.
     */
    readonly type: pulumi.Output<string>;
    /**
     * Create a Folder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering Folder resources.
 */
export interface FolderState {
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for folder. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     */
    readonly datacenterId?: pulumi.Input<string>;
    /**
     * The path of the folder to be created. This is relative to
     * the root of the type of folder you are creating, and the supplied datacenter.
     * For example, given a default datacenter of `default-dc`, a folder of type
     * `vm` (denoting a virtual machine folder), and a supplied folder of
     * `terraform-test-folder`, the resulting path would be
     * `/default-dc/vm/terraform-test-folder`.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of folder to create. Allowed options are
     * `datacenter` for datacenter folders, `host` for host and cluster folders,
     * `vm` for virtual machine folders, `datastore` for datastore folders, and
     * `network` for network folders. Forces a new resource if changed.
     */
    readonly type?: pulumi.Input<string>;
}
/**
 * The set of arguments for constructing a Folder resource.
 */
export interface FolderArgs {
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for folder. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     */
    readonly datacenterId?: pulumi.Input<string>;
    /**
     * The path of the folder to be created. This is relative to
     * the root of the type of folder you are creating, and the supplied datacenter.
     * For example, given a default datacenter of `default-dc`, a folder of type
     * `vm` (denoting a virtual machine folder), and a supplied folder of
     * `terraform-test-folder`, the resulting path would be
     * `/default-dc/vm/terraform-test-folder`.
     */
    readonly path: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of folder to create. Allowed options are
     * `datacenter` for datacenter folders, `host` for host and cluster folders,
     * `vm` for virtual machine folders, `datastore` for datastore folders, and
     * `network` for network folders. Forces a new resource if changed.
     */
    readonly type: pulumi.Input<string>;
}
