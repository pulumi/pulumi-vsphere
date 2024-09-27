// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere.ContentLibrary` resource can be used to manage content libraries.
 *
 * > **NOTE:** This resource requires a vCenter Server instance and is not available on direct ESXi host connections.
 *
 * ## Example Usage
 *
 * The following example creates a publishing content library using the datastore named `publisher-datastore` as the storage backing.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenterA = vsphere.getDatacenter({
 *     name: "dc-01-a",
 * });
 * const publisherDatastore = datacenterA.then(datacenterA => vsphere.getDatastore({
 *     name: "publisher-datastore",
 *     datacenterId: datacenterA.id,
 * }));
 * const publisherContentLibrary = new vsphere.ContentLibrary("publisher_content_library", {
 *     name: "Publisher Content Library",
 *     description: "A publishing content library.",
 *     storageBackings: [publisherDatastore.then(publisherDatastore => publisherDatastore.id)],
 * });
 * ```
 *
 * The next example creates a subscribed content library using the URL of the publisher content library as the source and the datastore named `subscriber-datastore` as the storage backing.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenterB = vsphere.getDatacenter({
 *     name: "dc-01-b",
 * });
 * const subscriberDatastore = datacenterB.then(datacenterB => vsphere.getDatastore({
 *     name: "subscriber-datastore",
 *     datacenterId: datacenterB.id,
 * }));
 * const subscriberContentLibrary = new vsphere.ContentLibrary("subscriber_content_library", {
 *     name: "Subscriber Content Library",
 *     description: "A subscribing content library.",
 *     storageBackings: [subscriberDatastore.then(subscriberDatastore => subscriberDatastore.id)],
 *     subscription: {
 *         subscriptionUrl: "https://vc-01-a.example.com:443/cls/vcsp/lib/f42a4b25-844a-44ec-9063-a3a5e9cc88c7/lib.json",
 *         automaticSync: true,
 *         onDemand: false,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * An existing content library can be imported into this resource by supplying the content library ID. For example:
 *
 * ```sh
 * $ pulumi import vsphere:index/contentLibrary:ContentLibrary vsphere_content_library publisher_content_library f42a4b25-844a-44ec-9063-a3a5e9cc88c7
 * ```
 */
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
     * A description for the content library.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the content library.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Options to publish a local content library.
     */
    public readonly publication!: pulumi.Output<outputs.ContentLibraryPublication>;
    /**
     * The managed object reference ID of the datastore on which to store the content library items.
     */
    public readonly storageBackings!: pulumi.Output<string[]>;
    /**
     * Options subscribe to a published content library.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContentLibraryState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["publication"] = state ? state.publication : undefined;
            resourceInputs["storageBackings"] = state ? state.storageBackings : undefined;
            resourceInputs["subscription"] = state ? state.subscription : undefined;
        } else {
            const args = argsOrState as ContentLibraryArgs | undefined;
            if ((!args || args.storageBackings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageBackings'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["publication"] = args ? args.publication : undefined;
            resourceInputs["storageBackings"] = args ? args.storageBackings : undefined;
            resourceInputs["subscription"] = args ? args.subscription : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContentLibrary.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContentLibrary resources.
 */
export interface ContentLibraryState {
    /**
     * A description for the content library.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the content library.
     */
    name?: pulumi.Input<string>;
    /**
     * Options to publish a local content library.
     */
    publication?: pulumi.Input<inputs.ContentLibraryPublication>;
    /**
     * The managed object reference ID of the datastore on which to store the content library items.
     */
    storageBackings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Options subscribe to a published content library.
     */
    subscription?: pulumi.Input<inputs.ContentLibrarySubscription>;
}

/**
 * The set of arguments for constructing a ContentLibrary resource.
 */
export interface ContentLibraryArgs {
    /**
     * A description for the content library.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the content library.
     */
    name?: pulumi.Input<string>;
    /**
     * Options to publish a local content library.
     */
    publication?: pulumi.Input<inputs.ContentLibraryPublication>;
    /**
     * The managed object reference ID of the datastore on which to store the content library items.
     */
    storageBackings: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Options subscribe to a published content library.
     */
    subscription?: pulumi.Input<inputs.ContentLibrarySubscription>;
}
