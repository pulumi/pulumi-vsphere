// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ContentLibraryItem extends pulumi.CustomResource {
    /**
     * Get an existing ContentLibraryItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContentLibraryItemState, opts?: pulumi.CustomResourceOptions): ContentLibraryItem {
        return new ContentLibraryItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/contentLibraryItem:ContentLibraryItem';

    /**
     * Returns true if the given object is an instance of ContentLibraryItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContentLibraryItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContentLibraryItem.__pulumiType;
    }

    /**
     * A description for the item.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * File to import into the Content Library item. OVFs and
     * OVAs will be parsed and associated files will also be imported.
     */
    public readonly fileUrl!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Content Library the item should be created in.
     */
    public readonly libraryId!: pulumi.Output<string>;
    /**
     * The name of the item to be created in the Content Library.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Virtual machine UUID to clone to Content Library.
     */
    public readonly sourceUuid!: pulumi.Output<string | undefined>;
    /**
     * Type of content library item.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a ContentLibraryItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContentLibraryItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContentLibraryItemArgs | ContentLibraryItemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContentLibraryItemState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fileUrl"] = state ? state.fileUrl : undefined;
            inputs["libraryId"] = state ? state.libraryId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sourceUuid"] = state ? state.sourceUuid : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ContentLibraryItemArgs | undefined;
            if ((!args || args.libraryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'libraryId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["fileUrl"] = args ? args.fileUrl : undefined;
            inputs["libraryId"] = args ? args.libraryId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sourceUuid"] = args ? args.sourceUuid : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ContentLibraryItem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContentLibraryItem resources.
 */
export interface ContentLibraryItemState {
    /**
     * A description for the item.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * File to import into the Content Library item. OVFs and
     * OVAs will be parsed and associated files will also be imported.
     */
    readonly fileUrl?: pulumi.Input<string>;
    /**
     * The ID of the Content Library the item should be created in.
     */
    readonly libraryId?: pulumi.Input<string>;
    /**
     * The name of the item to be created in the Content Library.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Virtual machine UUID to clone to Content Library.
     */
    readonly sourceUuid?: pulumi.Input<string>;
    /**
     * Type of content library item.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContentLibraryItem resource.
 */
export interface ContentLibraryItemArgs {
    /**
     * A description for the item.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * File to import into the Content Library item. OVFs and
     * OVAs will be parsed and associated files will also be imported.
     */
    readonly fileUrl?: pulumi.Input<string>;
    /**
     * The ID of the Content Library the item should be created in.
     */
    readonly libraryId: pulumi.Input<string>;
    /**
     * The name of the item to be created in the Content Library.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Virtual machine UUID to clone to Content Library.
     */
    readonly sourceUuid?: pulumi.Input<string>;
    /**
     * Type of content library item.
     */
    readonly type?: pulumi.Input<string>;
}
