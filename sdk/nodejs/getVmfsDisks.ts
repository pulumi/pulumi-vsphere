// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The `vsphere.getVmfsDisks` data source can be used to discover the storage
 * devices available on an ESXi host. This data source can be combined with the
 * `vsphere.VmfsDatastore` resource to create VMFS
 * datastores based off a set of discovered disks.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }, { async: true }));
 * const host = datacenter.apply(datacenter => vsphere.getHost({
 *     datacenterId: datacenter.id,
 *     name: "esxi1",
 * }, { async: true }));
 * const available = host.apply(host => vsphere.getVmfsDisks({
 *     filter: "mpx.vmhba1:C0:T[12]:L0",
 *     hostSystemId: host.id,
 *     rescan: true,
 * }, { async: true }));
 * ```
 */
export function getVmfsDisks(args: GetVmfsDisksArgs, opts?: pulumi.InvokeOptions): Promise<GetVmfsDisksResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
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
     * The managed object ID of
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
    readonly filter?: string;
    readonly hostSystemId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly rescan?: boolean;
}
