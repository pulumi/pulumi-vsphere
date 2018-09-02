import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_vmfs_disks` data source can be used to discover the storage
 * devices available on an ESXi host. This data source can be combined with the
 * [`vsphere_vmfs_datastore`][data-source-vmfs-datastore] resource to create VMFS
 * datastores based off a set of discovered disks.
 *
 * [data-source-vmfs-datastore]: /docs/providers/vsphere/r/vmfs_datastore.html
 */
export declare function getVmfsDisks(args: GetVmfsDisksArgs, opts?: pulumi.InvokeOptions): Promise<GetVmfsDisksResult>;
/**
 * A collection of arguments for invoking getVmfsDisks.
 */
export interface GetVmfsDisksArgs {
    /**
     * A regular expression to filter the disks against. Only
     * disks with canonical names that match will be included.
     */
    readonly filter?: string;
    /**
     * The [managed object ID][docs-about-morefs] of
     * the host to look for disks on.
     */
    readonly hostSystemId: string;
    /**
     * Whether or not to rescan storage adapters before
     * searching for disks. This may lengthen the time it takes to perform the
     * search. Default: `false`.
     */
    readonly rescan?: boolean;
}
/**
 * A collection of values returned by getVmfsDisks.
 */
export interface GetVmfsDisksResult {
    /**
     * A lexicographically sorted list of devices discovered by the
     * operation, matching the supplied `filter`, if provided.
     */
    readonly disks: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
