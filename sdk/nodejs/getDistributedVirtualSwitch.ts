// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.DistributedVirtualSwitch` data source can be used to discover
 * the ID and uplink data of a of a vSphere distributed switch (VDS). This
 * can then be used with resources or data sources that require a VDS, such as the
 * `vsphere.DistributedPortGroup` resource, for which
 * an example is shown below.
 *
 * > **NOTE:** This data source requires vCenter Server and is not available on
 * direct ESXi host connections.
 *
 * ## Example Usage
 *
 * The following example locates a distributed switch named `vds-01`, in the
 * datacenter `dc-01`. It then uses this distributed switch to set up a
 * `vsphere.DistributedPortGroup` resource that uses the first uplink as a
 * primary uplink and the second uplink as a secondary.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const vds = datacenter.then(datacenter => vsphere.getDistributedVirtualSwitch({
 *     name: "vds-01",
 *     datacenterId: datacenter.id,
 * }));
 * const dvpg = new vsphere.DistributedPortGroup("dvpg", {
 *     distributedVirtualSwitchUuid: vds.then(vds => vds.id),
 *     activeUplinks: [vds.then(vds => vds.uplinks?.[0])],
 *     standbyUplinks: [vds.then(vds => vds.uplinks?.[1])],
 * });
 * ```
 */
export function getDistributedVirtualSwitch(args: GetDistributedVirtualSwitchArgs, opts?: pulumi.InvokeOptions): Promise<GetDistributedVirtualSwitchResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
     * The managed object reference ID
     * of the datacenter the VDS is located in. This can be omitted if the search
     * path used in `name` is an absolute path. For default datacenters, use the `id`
     * attribute from an empty `vsphere.Datacenter` data source.
     */
    datacenterId?: string;
    /**
     * The name of the VDS. This can be a name or path.
     */
    name: string;
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
    /**
     * The list of the uplinks on this vSphere distributed switch, as per the
     * `uplinks` argument to the
     * `vsphere.DistributedVirtualSwitch`
     * resource.
     */
    readonly uplinks: string[];
}
/**
 * The `vsphere.DistributedVirtualSwitch` data source can be used to discover
 * the ID and uplink data of a of a vSphere distributed switch (VDS). This
 * can then be used with resources or data sources that require a VDS, such as the
 * `vsphere.DistributedPortGroup` resource, for which
 * an example is shown below.
 *
 * > **NOTE:** This data source requires vCenter Server and is not available on
 * direct ESXi host connections.
 *
 * ## Example Usage
 *
 * The following example locates a distributed switch named `vds-01`, in the
 * datacenter `dc-01`. It then uses this distributed switch to set up a
 * `vsphere.DistributedPortGroup` resource that uses the first uplink as a
 * primary uplink and the second uplink as a secondary.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const vds = datacenter.then(datacenter => vsphere.getDistributedVirtualSwitch({
 *     name: "vds-01",
 *     datacenterId: datacenter.id,
 * }));
 * const dvpg = new vsphere.DistributedPortGroup("dvpg", {
 *     distributedVirtualSwitchUuid: vds.then(vds => vds.id),
 *     activeUplinks: [vds.then(vds => vds.uplinks?.[0])],
 *     standbyUplinks: [vds.then(vds => vds.uplinks?.[1])],
 * });
 * ```
 */
export function getDistributedVirtualSwitchOutput(args: GetDistributedVirtualSwitchOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDistributedVirtualSwitchResult> {
    return pulumi.output(args).apply((a: any) => getDistributedVirtualSwitch(a, opts))
}

/**
 * A collection of arguments for invoking getDistributedVirtualSwitch.
 */
export interface GetDistributedVirtualSwitchOutputArgs {
    /**
     * The managed object reference ID
     * of the datacenter the VDS is located in. This can be omitted if the search
     * path used in `name` is an absolute path. For default datacenters, use the `id`
     * attribute from an empty `vsphere.Datacenter` data source.
     */
    datacenterId?: pulumi.Input<string>;
    /**
     * The name of the VDS. This can be a name or path.
     */
    name: pulumi.Input<string>;
}
