// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The `vsphere.DistributedVirtualSwitch` data source can be used to discover
 * the ID and uplink data of a of a vSphere distributed virtual switch (DVS). This
 * can then be used with resources or data sources that require a DVS, such as the
 * `vsphere.DistributedPortGroup` resource, for which
 * an example is shown below.
 *
 * > **NOTE:** This data source requires vCenter and is not available on direct
 * ESXi connections.
 *
 * ## Example Usage
 *
 * The following example locates a DVS that is named `test-dvs`, in the
 * datacenter `dc1`. It then uses this DVS to set up a
 * `vsphere.DistributedPortGroup` resource that uses the first uplink as a
 * primary uplink and the second uplink as a secondary.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }, { async: true }));
 * const dvs = datacenter.apply(datacenter => vsphere.getDistributedVirtualSwitch({
 *     datacenterId: datacenter.id,
 *     name: "test-dvs",
 * }, { async: true }));
 * const pg = new vsphere.DistributedPortGroup("pg", {
 *     activeUplinks: [dvs.apply(dvs => dvs.uplinks[0])],
 *     distributedVirtualSwitchUuid: dvs.id,
 *     standbyUplinks: [dvs.apply(dvs => dvs.uplinks[1])],
 * });
 * ```
 */
export function getDistributedVirtualSwitch(args: GetDistributedVirtualSwitchArgs, opts?: pulumi.InvokeOptions): Promise<GetDistributedVirtualSwitchResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch", {
        "datacenterId": args.datacenterId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDistributedVirtualSwitch.
 */
export interface GetDistributedVirtualSwitchArgs {
    /**
     * The managed object reference
     * ID of the datacenter the DVS is located in. This can be
     * omitted if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
     * source.
     */
    readonly datacenterId?: string;
    /**
     * The name of the distributed virtual switch. This can be a
     * name or path.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getDistributedVirtualSwitch.
 */
export interface GetDistributedVirtualSwitchResult {
    readonly datacenterId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly uplinks: string[];
}
