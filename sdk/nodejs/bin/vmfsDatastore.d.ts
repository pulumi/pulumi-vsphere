import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_vmfs_datastore` resource can be used to create and manage VMFS
 * datastores on an ESXi host or a set of hosts. The resource supports using any
 * SCSI device that can generally be used in a datastore, such as local disks, or
 * disks presented to a host or multiple hosts over Fibre Channel or iSCSI.
 * Devices can be specified manually, or discovered using the
 * [`vsphere_vmfs_disks`][data-source-vmfs-disks] data source.
 *
 * [data-source-vmfs-disks]: /docs/providers/vsphere/d/vmfs_disks.html
 */
export declare class VmfsDatastore extends pulumi.CustomResource {
    /**
     * Get an existing VmfsDatastore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VmfsDatastoreState): VmfsDatastore;
    /**
     * The connectivity status of the datastore. If this is `false`,
     * some other computed attributes may be out of date.
     */
    readonly accessible: pulumi.Output<boolean>;
    /**
     * Maximum capacity of the datastore, in megabytes.
     */
    readonly capacity: pulumi.Output<number>;
    /**
     * Map of custom attribute ids to attribute
     * value string to set on datastore resource. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The [managed object
     * ID][docs-about-morefs] of a datastore cluster to put this datastore in.
     * Conflicts with `folder`.
     */
    readonly datastoreClusterId: pulumi.Output<string | undefined>;
    /**
     * The disks to use with the datastore.
     */
    readonly disks: pulumi.Output<string[]>;
    /**
     * The relative path to a folder to put this datastore in.
     * This is a path relative to the datacenter you are deploying the datastore to.
     * Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
     * Terraform will place a datastore named `terraform-test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/terraform-test`. Conflicts with
     * `datastore_cluster_id`.
     */
    readonly folder: pulumi.Output<string | undefined>;
    /**
     * Available space of this datastore, in megabytes.
     */
    readonly freeSpace: pulumi.Output<number>;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to set the datastore up on. Note that this is not necessarily the
     * only host that the datastore will be set up on - see
     * here for more info. Forces a
     * new resource if changed.
     */
    readonly hostSystemId: pulumi.Output<string>;
    /**
     * The current maintenance mode state of the datastore.
     */
    readonly maintenanceMode: pulumi.Output<string>;
    /**
     * If `true`, more than one host in the datacenter has
     * been configured with access to the datastore.
     */
    readonly multipleHostAccess: pulumi.Output<boolean>;
    /**
     * The name of the datastore. Forces a new resource if
     * changed.
     */
    readonly name: pulumi.Output<string>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * Total additional storage space, in megabytes,
     * potentially used by all virtual machines on this datastore.
     */
    readonly uncommittedSpace: pulumi.Output<number>;
    /**
     * The unique locator for the datastore.
     */
    readonly url: pulumi.Output<string>;
    /**
     * Create a VmfsDatastore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VmfsDatastoreArgs, opts?: pulumi.CustomResourceOptions);
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
     * value string to set on datastore resource. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The [managed object
     * ID][docs-about-morefs] of a datastore cluster to put this datastore in.
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
     * Terraform will place a datastore named `terraform-test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/terraform-test`. Conflicts with
     * `datastore_cluster_id`.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * Available space of this datastore, in megabytes.
     */
    readonly freeSpace?: pulumi.Input<number>;
    /**
     * The [managed object ID][docs-about-morefs] of
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
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
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
     * value string to set on datastore resource. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The [managed object
     * ID][docs-about-morefs] of a datastore cluster to put this datastore in.
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
     * Terraform will place a datastore named `terraform-test` in a datastore folder
     * located at `/dc1/datastore/foo/bar`, with the final inventory path being
     * `/dc1/datastore/foo/bar/terraform-test`. Conflicts with
     * `datastore_cluster_id`.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * The [managed object ID][docs-about-morefs] of
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
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
