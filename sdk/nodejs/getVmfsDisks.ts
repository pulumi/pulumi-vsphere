// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere_vmfs_disks` data source can be used to discover the storage
 * devices available on an ESXi host. This data source can be combined with the
 * [`vsphere_vmfs_datastore`][data-source-vmfs-datastore] resource to create VMFS
 * datastores based off a set of discovered disks.
 * 
 * [data-source-vmfs-datastore]: /docs/providers/vsphere/r/vmfs_datastore.html
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const vsphere_datacenter_datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }));
 * const vsphere_host_host = pulumi.output(vsphere.getHost({
 *     datacenterId: vsphere_datacenter_datacenter.apply(__arg0 => __arg0.id),
 *     name: "esxi1",
 * }));
 * const vsphere_vmfs_disks_available = pulumi.output(vsphere.getVmfsDisks({
 *     filter: "mpx.vmhba1:C0:T[12]:L0",
 *     hostSystemId: vsphere_host_host.apply(__arg0 => __arg0.id),
 *     rescan: true,
 * }));
 * ```
 */
export function getVmfsDisks(args: GetVmfsDisksArgs, opts?: pulumi.InvokeOptions): Promise<GetVmfsDisksResult> {
    return pulumi.runtime.invoke("vsphere:index/getVmfsDisks:getVmfsDisks", {
        "filter": args.filter,
        "hostSystemId": args.hostSystemId,
        "rescan": args.rescan,
    }, opts);
}

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
