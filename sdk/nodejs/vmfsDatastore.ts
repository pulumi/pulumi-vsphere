// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.VmfsDatastore` resource can be used to create and manage VMFS
 * datastores on an ESXi host or a set of hosts. The resource supports using any
 * SCSI device that can generally be used in a datastore, such as local disks, or
 * disks presented to a host or multiple hosts over Fibre Channel or iSCSI.
 * Devices can be specified manually, or discovered using the
 * [`vsphere.getVmfsDisks`][data-source-vmfs-disks] data source.
 *
 * [data-source-vmfs-disks]: /docs/providers/vsphere/d/vmfs_disks.html
 *
 * ## Auto-Mounting of Datastores Within vCenter
 *
 * Note that the current behavior of this resource will auto-mount any created
 * datastores to any other host within vCenter that has access to the same disk.
 *
 * Example: You want to create a datastore with a iSCSI LUN that is visible on 3
 * hosts in a single vSphere cluster (`esxi1`, `esxi2` and `esxi3`). When you
 * create the datastore on `esxi1`, the datastore will be automatically mounted on
 * `esxi2` and `esxi3`, without the need to configure the resource on either of
 * those two hosts.
 *
 * Future versions of this resource may allow you to control the hosts that a
 * datastore is mounted to, but currently, this automatic behavior cannot be
 * changed, so keep this in mind when writing your configurations and deploying
 * your disks.
 *
 * ## Increasing Datastore Size
 *
 * To increase the size of a datastore, you must add additional disks to the
 * `disks` attribute. Expanding the size of a datastore by increasing the size of
 * an already provisioned disk is currently not supported (but may be in future
 * versions of this resource).
 *
 * > **NOTE:** You cannot decrease the size of a datastore. If the resource
 * detects disks removed from the configuration, the provider will give an error.
 *
 * [cmd-taint]: /docs/commands/taint.html
 *
 * ## Example Usage
 *
 * ### Addition of local disks on a single host
 *
 * The following example uses the default datacenter and default host to add a
 * datastore with local disks to a single ESXi server.
 *
 * > **NOTE:** There are some situations where datastore creation will not work
 * when working through vCenter (usually when trying to create a datastore on a
 * single host with local disks). If you experience trouble creating the datastore
 * you need through vCenter, break the datastore off into a different configuration
 * and deploy it using the ESXi server as the provider endpoint, using a similar
 * configuration to what is below.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({});
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *     datacenterId: datacenter.id,
 * }));
 * const datastore = new vsphere.VmfsDatastore("datastore", {
 *     name: "test",
 *     hostSystemId: esxiHost.id,
 *     disks: [
 *         "mpx.vmhba1:C0:T1:L0",
 *         "mpx.vmhba1:C0:T2:L0",
 *         "mpx.vmhba1:C0:T2:L0",
 *     ],
 * });
 * ```
 *
 * ### Auto-detection of disks via `vsphere.getVmfsDisks`
 *
 * The following example makes use of the
 * `vsphere.getVmfsDisks` data source to auto-detect
 * exported iSCSI LUNS matching a certain NAA vendor ID (in this case, LUNs
 * exported from a [NetApp][ext-netapp]). These discovered disks are then loaded
 * into `vsphere.VmfsDatastore`. The datastore is also placed in the
 * `datastore-folder` folder afterwards.
 *
 * [ext-netapp]: https://kb.netapp.com/support/s/article/ka31A0000000rLRQAY/how-to-match-a-lun-s-naa-number-to-its-serial-number?language=en_US
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *     name: "esxi-01.example.com",
 *     datacenterId: datacenter.id,
 * }));
 * const available = host.then(host => vsphere.getVmfsDisks({
 *     hostSystemId: host.id,
 *     rescan: true,
 *     filter: "naa.60a98000",
 * }));
 * const datastore = new vsphere.VmfsDatastore("datastore", {
 *     name: "test",
 *     hostSystemId: esxiHost.id,
 *     folder: "datastore-folder",
 *     disks: [available.then(available => available.disks)],
 * });
 * ```
 *
 * ## Import
 *
 * An existing VMFS datastore can be imported into this resource
 *
 * via its managed object ID, via the command below. You also need the host system
 *
 * ID.
 *
 * ```sh
 * $ pulumi import vsphere:index/vmfsDatastore:VmfsDatastore datastore datastore-123:host-10
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
 *
 * To locate host IDs, it might be a good idea to supply the `-l` flag as well so
 *
 * that you can line up the names with the IDs:
 *
 * $ govc ls -l -i /dc/host/cluster1
 *
 * ResourcePool:resgroup-10 /dc/host/cluster1/Resources
 *
 * HostSystem:host-10 /dc/host/cluster1/esxi1
 *
 * HostSystem:host-11 /dc/host/cluster1/esxi2
 *
 * HostSystem:host-12 /dc/host/cluster1/esxi3
 */
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
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VmfsDatastoreState | undefined;
            resourceInputs["accessible"] = state ? state.accessible : undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["customAttributes"] = state ? state.customAttributes : undefined;
            resourceInputs["datastoreClusterId"] = state ? state.datastoreClusterId : undefined;
            resourceInputs["disks"] = state ? state.disks : undefined;
            resourceInputs["folder"] = state ? state.folder : undefined;
            resourceInputs["freeSpace"] = state ? state.freeSpace : undefined;
            resourceInputs["hostSystemId"] = state ? state.hostSystemId : undefined;
            resourceInputs["maintenanceMode"] = state ? state.maintenanceMode : undefined;
            resourceInputs["multipleHostAccess"] = state ? state.multipleHostAccess : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["uncommittedSpace"] = state ? state.uncommittedSpace : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as VmfsDatastoreArgs | undefined;
            if ((!args || args.disks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'disks'");
            }
            if ((!args || args.hostSystemId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostSystemId'");
            }
            resourceInputs["customAttributes"] = args ? args.customAttributes : undefined;
            resourceInputs["datastoreClusterId"] = args ? args.datastoreClusterId : undefined;
            resourceInputs["disks"] = args ? args.disks : undefined;
            resourceInputs["folder"] = args ? args.folder : undefined;
            resourceInputs["hostSystemId"] = args ? args.hostSystemId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["accessible"] = undefined /*out*/;
            resourceInputs["capacity"] = undefined /*out*/;
            resourceInputs["freeSpace"] = undefined /*out*/;
            resourceInputs["maintenanceMode"] = undefined /*out*/;
            resourceInputs["multipleHostAccess"] = undefined /*out*/;
            resourceInputs["uncommittedSpace"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VmfsDatastore.__pulumiType, name, resourceInputs, opts);
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
    accessible?: pulumi.Input<boolean>;
    /**
     * Maximum capacity of the datastore, in megabytes.
     */
    capacity?: pulumi.Input<number>;
    /**
     * Map of custom attribute ids to attribute 
     * value string to set on datastore resource.
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
     * The disks to use with the datastore.
     */
    disks?: pulumi.Input<pulumi.Input<string>[]>;
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
     * The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     */
    hostSystemId?: pulumi.Input<string>;
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
     * The IDs of any tags to attach to this resource. 
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
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
 * The set of arguments for constructing a VmfsDatastore resource.
 */
export interface VmfsDatastoreArgs {
    /**
     * Map of custom attribute ids to attribute 
     * value string to set on datastore resource.
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
     * The disks to use with the datastore.
     */
    disks: pulumi.Input<pulumi.Input<string>[]>;
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
     * The managed object ID of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     */
    hostSystemId: pulumi.Input<string>;
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The IDs of any tags to attach to this resource. 
     *
     * > **NOTE:** Tagging support is unsupported on direct ESXi connections and
     * requires vCenter 6.0 or higher.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
