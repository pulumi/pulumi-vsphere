// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/folder.html.markdown.
 */
export class Folder extends pulumi.CustomResource {
    /**
     * Get an existing Folder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FolderState, opts?: pulumi.CustomResourceOptions): Folder {
        return new Folder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/folder:Folder';

    /**
     * Returns true if the given object is an instance of Folder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Folder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Folder.__pulumiType;
    }

    /**
     * Map of custom attribute ids to attribute 
     * value strings to set for folder. See [here][docs-setting-custom-attributes]
     * for a reference on how to set values for custom attributes.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     */
    public readonly datacenterId!: pulumi.Output<string | undefined>;
    /**
     * The path of the folder and any parents, relative to the datacenter and folder type being defined.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of folder to create. Allowed options are
     * `datacenter` for datacenter folders, `host` for host and cluster folders,
     * `vm` for virtual machine folders, `datastore` for datastore folders, and
     * `network` for network folders. Forces a new resource if changed.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Folder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FolderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FolderArgs | FolderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FolderState | undefined;
            inputs["customAttributes"] = state ? state.customAttributes : undefined;
            inputs["datacenterId"] = state ? state.datacenterId : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as FolderArgs | undefined;
            if (!args || args.path === undefined) {
                throw new Error("Missing required property 'path'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["datacenterId"] = args ? args.datacenterId : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Folder.__pulumiType, name, inputs, opts);
    }
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
    readonly customAttributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     */
    readonly datacenterId?: pulumi.Input<string>;
    /**
     * The path of the folder and any parents, relative to the datacenter and folder type being defined.
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
    readonly customAttributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the datacenter the folder will be created in.
     * Required for all folder types except for datacenter folders. Forces a new
     * resource if changed.
     */
    readonly datacenterId?: pulumi.Input<string>;
    /**
     * The path of the folder and any parents, relative to the datacenter and folder type being defined.
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
