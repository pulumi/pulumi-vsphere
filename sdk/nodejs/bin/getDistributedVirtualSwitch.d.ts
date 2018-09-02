import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_distributed_virtual_switch` data source can be used to discover
 * the ID and uplink data of a of a vSphere distributed virtual switch (DVS). This
 * can then be used with resources or data sources that require a DVS, such as the
 * [`vsphere_distributed_port_group`][distributed-port-group] resource, for which
 * an example is shown below.
 *
 * [distributed-port-group]: /docs/providers/vsphere/r/distributed_port_group.html
 *
 * ~> **NOTE:** This data source requires vCenter and is not available on direct
 * ESXi connections.
 */
export declare function getDistributedVirtualSwitch(args: GetDistributedVirtualSwitchArgs, opts?: pulumi.InvokeOptions): Promise<GetDistributedVirtualSwitchResult>;
/**
 * A collection of arguments for invoking getDistributedVirtualSwitch.
 */
export interface GetDistributedVirtualSwitchArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datacenter the DVS is located in. This can be
     * omitted if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere_datacenter` data
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
    readonly uplinks: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
