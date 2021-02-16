// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VmfsDatastore extends pulumi.CustomResource {
    /**
     * Get an existing VmfsDatastore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VmfsDatastoreState, opts?: pulumi.CustomResourceOptions): VmfsDatastore {
        return new VmfsDatastore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/vmfsDatastore:VmfsDatastore';

    /**
     * Returns true if the given object is an instance of VmfsDatastore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VmfsDatastore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VmfsDatastore.__pulumiType;
    }

    /**
     * The connectivity status of the datastore. If this is `false`,
     * some other computed attributes may be out of date.
     */
    public /*out*/ readonly accessible!: pulumi.Output<boolean>;
    /**
     * Maximum capacity of the datastore, in megabytes.
     */
    public /*out*/ readonly capacity!: pulumi.Output<number>;
    /**
     * Map of custom attribute ids to attribute 
     * value string to set on datastore resource.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     */
    public readonly datastoreClusterId!: pulumi.Output<string | undefined>;
    /**
     * The disks to use with the datastore.
     */
    public readonly disks!: pulumi.Output<string[]>;
    /**
     * The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * The provider will place a datastore named `test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/test`. Conflicts with
     * `datastoreClusterId`.
     */
    public readonly folder!: pulumi.Output<string | undefined>;
    /**
     * Available space of this datastore, in megabytes.
     */
    public /*out*/ readonly freeSpace!: pulumi.Output<number>;
    /**
     * The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     */
    public readonly hostSystemId!: pulumi.Output<string>;
    /**
     * The current maintenance mode state of the datastore.
     */
    public /*out*/ readonly maintenanceMode!: pulumi.Output<string>;
    /**
     * If `true`, more than one host in the datacenter has
     * been configured with access to the datastore.
     */
    public /*out*/ readonly multipleHostAccess!: pulumi.Output<boolean>;
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Total additional storage space, in megabytes,
     * potentially used by all virtual machines on this datastore.
     */
    public /*out*/ readonly uncommittedSpace!: pulumi.Output<number>;
    /**
     * The unique locator for the datastore.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a VmfsDatastore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VmfsDatastoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VmfsDatastoreArgs | VmfsDatastoreState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VmfsDatastoreState | undefined;
            inputs["accessible"] = state ? state.accessible : undefined;
            inputs["capacity"] = state ? state.capacity : undefined;
            inputs["customAttributes"] = state ? state.customAttributes : undefined;
            inputs["datastoreClusterId"] = state ? state.datastoreClusterId : undefined;
            inputs["disks"] = state ? state.disks : undefined;
            inputs["folder"] = state ? state.folder : undefined;
            inputs["freeSpace"] = state ? state.freeSpace : undefined;
            inputs["hostSystemId"] = state ? state.hostSystemId : undefined;
            inputs["maintenanceMode"] = state ? state.maintenanceMode : undefined;
            inputs["multipleHostAccess"] = state ? state.multipleHostAccess : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uncommittedSpace"] = state ? state.uncommittedSpace : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as VmfsDatastoreArgs | undefined;
            if ((!args || args.disks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'disks'");
            }
            if ((!args || args.hostSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostSystemId'");
            }
            inputs["customAttributes"] = args ? args.customAttributes : undefined;
            inputs["datastoreClusterId"] = args ? args.datastoreClusterId : undefined;
            inputs["disks"] = args ? args.disks : undefined;
            inputs["folder"] = args ? args.folder : undefined;
            inputs["hostSystemId"] = args ? args.hostSystemId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["accessible"] = undefined /*out*/;
            inputs["capacity"] = undefined /*out*/;
            inputs["freeSpace"] = undefined /*out*/;
            inputs["maintenanceMode"] = undefined /*out*/;
            inputs["multipleHostAccess"] = undefined /*out*/;
            inputs["uncommittedSpace"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VmfsDatastore.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VmfsDatastore resources.
 */
export interface VmfsDatastoreState {
    /**
     * The connectivity status of the datastore. If this is `false`,
     * some other computed attributes may be out of date.
     */
    readonly accessible?: pulumi.Input<boolean>;
    /**
     * Maximum capacity of the datastore, in megabytes.
     */
    readonly capacity?: pulumi.Input<number>;
    /**
     * Map of custom attribute ids to attribute 
     * value string to set on datastore resource.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     */
    readonly datastoreClusterId?: pulumi.Input<string>;
    /**
     * The disks to use with the datastore.
     */
    readonly disks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * The provider will place a datastore named `test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/test`. Conflicts with
     * `datastoreClusterId`.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * Available space of this datastore, in megabytes.
     */
    readonly freeSpace?: pulumi.Input<number>;
    /**
     * The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     */
    readonly hostSystemId?: pulumi.Input<string>;
    /**
     * The current maintenance mode state of the datastore.
     */
    readonly maintenanceMode?: pulumi.Input<string>;
    /**
     * If `true`, more than one host in the datacenter has
     * been configured with access to the datastore.
     */
    readonly multipleHostAccess?: pulumi.Input<boolean>;
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Total additional storage space, in megabytes,
     * potentially used by all virtual machines on this datastore.
     */
    readonly uncommittedSpace?: pulumi.Input<number>;
    /**
     * The unique locator for the datastore.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VmfsDatastore resource.
 */
export interface VmfsDatastoreArgs {
    /**
     * Map of custom attribute ids to attribute 
     * value string to set on datastore resource.
     */
    readonly customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     */
    readonly datastoreClusterId?: pulumi.Input<string>;
    /**
     * The disks to use with the datastore.
     */
    readonly disks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * The provider will place a datastore named `test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/test`. Conflicts with
     * `datastoreClusterId`.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     */
    readonly hostSystemId: pulumi.Input<string>;
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
