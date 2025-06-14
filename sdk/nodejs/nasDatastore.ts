// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.NasDatastore` resource can be used to create and manage NAS
 * datastores on an ESXi host or a set of hosts. The resource supports mounting
 * NFS v3 and v4.1 shares to be used as datastores.
 *
 * > **NOTE:** Unlike `vsphere.VmfsDatastore`, a NAS
 * datastore is only mounted on the hosts you choose to mount it on. To mount on
 * multiple hosts, you must specify each host that you want to add in the
 * `hostSystemIds` argument.
 *
 * ## Example Usage
 *
 * The following example would set up a NFS v3 share on 3 hosts connected through
 * vCenter in the same datacenter - `esxi1`, `esxi2`, and `esxi3`. The remote host
 * is named `nfs` and has `/export/test` exported.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const config = new pulumi.Config();
 * const hosts = config.getObject<any>("hosts") || [
 *     "esxi-01.example.com",
 *     "esxi-02.example.com",
 *     "esxi-03.example.com",
 * ];
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const hostsGetHost = (new Array(hosts.length)).map((_, i) => i).map(__index => (vsphere.getHost({
 *     name: hosts[__index],
 *     datacenterId: _arg0_.id,
 * })));
 * const datastore = new vsphere.NasDatastore("datastore", {
 *     name: "test",
 *     hostSystemIds: [esxiHosts.map(__item => __item.id)],
 *     type: "NFS",
 *     remoteHosts: ["nfs"],
 *     remotePath: "/export/test",
 * });
 * ```
 *
 * ## Import
 *
 * An existing NAS datastore can be imported into this resource via
 *
 * its managed object ID, via the following command:
 *
 * [docs-import]: https://developer.hashicorp.com/terraform/cli/import
 *
 * ```sh
 * $ pulumi import vsphere:index/nasDatastore:NasDatastore datastore datastore-123
 * ```
 *
 * You need a tool like [`govc`][ext-govc] that can display managed object IDs.
 *
 * [ext-govc]: https://github.com/vmware/govmomi/tree/master/govc
 *
 * In the case of govc, you can locate a managed object ID from an inventory path
 *
 * by doing the following:
 *
 * $ govc ls -i /dc/datastore/terraform-test
 *
 * Datastore:datastore-123
 */
export class NasDatastore extends pulumi.CustomResource {
    /**
     * Get an existing NasDatastore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NasDatastoreState, opts?: pulumi.CustomResourceOptions): NasDatastore {
        return new NasDatastore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vsphere:index/nasDatastore:NasDatastore';

    /**
     * Returns true if the given object is an instance of NasDatastore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NasDatastore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NasDatastore.__pulumiType;
    }

    /**
     * Access mode for the mount point. Can be one of
     * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
     * that the datastore will be read-write depending on the permissions of the
     * actual share. Default: `readWrite`. Forces a new resource if changed.
     */
    public readonly accessMode!: pulumi.Output<string | undefined>;
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
     * value strings to set on datasource resource.
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     */
    public readonly customAttributes!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     */
    public readonly datastoreClusterId!: pulumi.Output<string | undefined>;
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
     * The managed object IDs of
     * the hosts to mount the datastore on.
     */
    public readonly hostSystemIds!: pulumi.Output<string[]>;
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
     * Indicates that this NAS volume is a protocol endpoint.
     * This field is only populated if the host supports virtual datastores.
     */
    public /*out*/ readonly protocolEndpoint!: pulumi.Output<boolean>;
    /**
     * The hostnames or IP addresses of the remote
     * servers. Only one element should be present for NFS v3 but multiple
     * can be present for NFS v4.1. Forces a new resource if changed.
     */
    public readonly remoteHosts!: pulumi.Output<string[]>;
    /**
     * The remote path of the mount point. Forces a new
     * resource if changed.
     */
    public readonly remotePath!: pulumi.Output<string>;
    /**
     * The security type to use when using NFS v4.1.
     * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
     * if changed.
     */
    public readonly securityType!: pulumi.Output<string | undefined>;
    /**
     * The IDs of any tags to attach to this resource. 
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The type of NAS volume. Can be one of `NFS` (to denote
     * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
     * changed.
     */
    public readonly type!: pulumi.Output<string | undefined>;
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
     * Create a NasDatastore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NasDatastoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NasDatastoreArgs | NasDatastoreState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NasDatastoreState | undefined;
            resourceInputs["accessMode"] = state ? state.accessMode : undefined;
            resourceInputs["accessible"] = state ? state.accessible : undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["customAttributes"] = state ? state.customAttributes : undefined;
            resourceInputs["datastoreClusterId"] = state ? state.datastoreClusterId : undefined;
            resourceInputs["folder"] = state ? state.folder : undefined;
            resourceInputs["freeSpace"] = state ? state.freeSpace : undefined;
            resourceInputs["hostSystemIds"] = state ? state.hostSystemIds : undefined;
            resourceInputs["maintenanceMode"] = state ? state.maintenanceMode : undefined;
            resourceInputs["multipleHostAccess"] = state ? state.multipleHostAccess : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocolEndpoint"] = state ? state.protocolEndpoint : undefined;
            resourceInputs["remoteHosts"] = state ? state.remoteHosts : undefined;
            resourceInputs["remotePath"] = state ? state.remotePath : undefined;
            resourceInputs["securityType"] = state ? state.securityType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uncommittedSpace"] = state ? state.uncommittedSpace : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as NasDatastoreArgs | undefined;
            if ((!args || args.hostSystemIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostSystemIds'");
            }
            if ((!args || args.remoteHosts === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteHosts'");
            }
            if ((!args || args.remotePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remotePath'");
            }
            resourceInputs["accessMode"] = args ? args.accessMode : undefined;
            resourceInputs["customAttributes"] = args ? args.customAttributes : undefined;
            resourceInputs["datastoreClusterId"] = args ? args.datastoreClusterId : undefined;
            resourceInputs["folder"] = args ? args.folder : undefined;
            resourceInputs["hostSystemIds"] = args ? args.hostSystemIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["remoteHosts"] = args ? args.remoteHosts : undefined;
            resourceInputs["remotePath"] = args ? args.remotePath : undefined;
            resourceInputs["securityType"] = args ? args.securityType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["accessible"] = undefined /*out*/;
            resourceInputs["capacity"] = undefined /*out*/;
            resourceInputs["freeSpace"] = undefined /*out*/;
            resourceInputs["maintenanceMode"] = undefined /*out*/;
            resourceInputs["multipleHostAccess"] = undefined /*out*/;
            resourceInputs["protocolEndpoint"] = undefined /*out*/;
            resourceInputs["uncommittedSpace"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NasDatastore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NasDatastore resources.
 */
export interface NasDatastoreState {
    /**
     * Access mode for the mount point. Can be one of
     * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
     * that the datastore will be read-write depending on the permissions of the
     * actual share. Default: `readWrite`. Forces a new resource if changed.
     */
    accessMode?: pulumi.Input<string>;
    /**
     * The connectivity status of the datastore. If this is `false`,
     * some other computed attributes may be out of date.
     */
    accessible?: pulumi.Input<boolean>;
    /**
     * Maximum capacity of the datastore, in megabytes.
     */
    capacity?: pulumi.Input<number>;
    /**
     * Map of custom attribute ids to attribute 
     * value strings to set on datasource resource.
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     */
    datastoreClusterId?: pulumi.Input<string>;
    /**
     * The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * The provider will place a datastore named `test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/test`. Conflicts with
     * `datastoreClusterId`.
     */
    folder?: pulumi.Input<string>;
    /**
     * Available space of this datastore, in megabytes.
     */
    freeSpace?: pulumi.Input<number>;
    /**
     * The managed object IDs of
     * the hosts to mount the datastore on.
     */
    hostSystemIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The current maintenance mode state of the datastore.
     */
    maintenanceMode?: pulumi.Input<string>;
    /**
     * If `true`, more than one host in the datacenter has
     * been configured with access to the datastore.
     */
    multipleHostAccess?: pulumi.Input<boolean>;
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     */
    name?: pulumi.Input<string>;
    /**
     * Indicates that this NAS volume is a protocol endpoint.
     * This field is only populated if the host supports virtual datastores.
     */
    protocolEndpoint?: pulumi.Input<boolean>;
    /**
     * The hostnames or IP addresses of the remote
     * servers. Only one element should be present for NFS v3 but multiple
     * can be present for NFS v4.1. Forces a new resource if changed.
     */
    remoteHosts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The remote path of the mount point. Forces a new
     * resource if changed.
     */
    remotePath?: pulumi.Input<string>;
    /**
     * The security type to use when using NFS v4.1.
     * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
     * if changed.
     */
    securityType?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. 
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of NAS volume. Can be one of `NFS` (to denote
     * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
     * changed.
     */
    type?: pulumi.Input<string>;
    /**
     * Total additional storage space, in megabytes,
     * potentially used by all virtual machines on this datastore.
     */
    uncommittedSpace?: pulumi.Input<number>;
    /**
     * The unique locator for the datastore.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NasDatastore resource.
 */
export interface NasDatastoreArgs {
    /**
     * Access mode for the mount point. Can be one of
     * `readOnly` or `readWrite`. Note that `readWrite` does not necessarily mean
     * that the datastore will be read-write depending on the permissions of the
     * actual share. Default: `readWrite`. Forces a new resource if changed.
     */
    accessMode?: pulumi.Input<string>;
    /**
     * Map of custom attribute ids to attribute 
     * value strings to set on datasource resource.
     *
     * > **NOTE:** Custom attributes are unsupported on direct ESXi connections
     * and require vCenter.
     */
    customAttributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The managed object
     * ID of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     */
    datastoreClusterId?: pulumi.Input<string>;
    /**
     * The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * The provider will place a datastore named `test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/test`. Conflicts with
     * `datastoreClusterId`.
     */
    folder?: pulumi.Input<string>;
    /**
     * The managed object IDs of
     * the hosts to mount the datastore on.
     */
    hostSystemIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The hostnames or IP addresses of the remote
     * servers. Only one element should be present for NFS v3 but multiple
     * can be present for NFS v4.1. Forces a new resource if changed.
     */
    remoteHosts: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The remote path of the mount point. Forces a new
     * resource if changed.
     */
    remotePath: pulumi.Input<string>;
    /**
     * The security type to use when using NFS v4.1.
     * Can be one of `AUTH_SYS`, `SEC_KRB5`, or `SEC_KRB5I`. Forces a new resource
     * if changed.
     */
    securityType?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. 
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of NAS volume. Can be one of `NFS` (to denote
     * v3) or `NFS41` (to denote NFS v4.1). Default: `NFS`. Forces a new resource if
     * changed.
     */
    type?: pulumi.Input<string>;
}
