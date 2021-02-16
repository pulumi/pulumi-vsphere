// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class ContentLibrary extends pulumi.CustomResource {
    /**
     * Get an existing ContentLibrary resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContentLibraryState, opts?: pulumi.CustomResourceOptions): ContentLibrary {
        return new ContentLibrary(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/contentLibrary:ContentLibrary';

    /**
     * Returns true if the given object is an instance of ContentLibrary.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContentLibrary {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContentLibrary.__pulumiType;
    }

    /**
     * A description of the Content Library.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the Content Library.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Options to publish a local Content Library.
     */
    public readonly publication!: pulumi.Output<outputs.ContentLibraryPublication>;
    /**
     * The managed object reference ID on which to store Content Library
     * items.
     */
    public readonly storageBackings!: pulumi.Output<string[]>;
    /**
     * Options to publish a local Content Library.
     */
    public readonly subscription!: pulumi.Output<outputs.ContentLibrarySubscription | undefined>;

    /**
     * Create a ContentLibrary resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContentLibraryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContentLibraryArgs | ContentLibraryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContentLibraryState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["publication"] = state ? state.publication : undefined;
            inputs["storageBackings"] = state ? state.storageBackings : undefined;
            inputs["subscription"] = state ? state.subscription : undefined;
        } else {
            const args = argsOrState as ContentLibraryArgs | undefined;
            if ((!args || args.storageBackings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageBackings'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["publication"] = args ? args.publication : undefined;
            inputs["storageBackings"] = args ? args.storageBackings : undefined;
            inputs["subscription"] = args ? args.subscription : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ContentLibrary.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContentLibrary resources.
 */
export interface ContentLibraryState {
    /**
     * A description of the Content Library.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the Content Library.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Options to publish a local Content Library.
     */
    readonly publication?: pulumi.Input<inputs.ContentLibraryPublication>;
    /**
     * The managed object reference ID on which to store Content Library
     * items.
     */
    readonly storageBackings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Options to publish a local Content Library.
     */
    readonly subscription?: pulumi.Input<inputs.ContentLibrarySubscription>;
}

/**
 * The set of arguments for constructing a ContentLibrary resource.
 */
export interface ContentLibraryArgs {
    /**
     * A description of the Content Library.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the Content Library.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Options to publish a local Content Library.
     */
    readonly publication?: pulumi.Input<inputs.ContentLibraryPublication>;
    /**
     * The managed object reference ID on which to store Content Library
     * items.
     */
    readonly storageBackings: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Options to publish a local Content Library.
     */
    readonly subscription?: pulumi.Input<inputs.ContentLibrarySubscription>;
}
