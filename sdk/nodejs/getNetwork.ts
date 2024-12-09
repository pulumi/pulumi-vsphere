// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere.getNetwork` data source can be used to discover the ID of a network in
 * vSphere. This can be any network that can be used as the backing for a network
 * interface for `vsphere.VirtualMachine` or any other vSphere resource that
 * requires a network. This includes standard (host-based) port groups, distributed
 * port groups, or opaque networks such as those managed by NSX.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const network = datacenter.then(datacenter => vsphere.getNetwork({
 *     name: "VM Network",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 *
 * ### Additional Examples
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const myPortGroup = datacenter.then(datacenter => vsphere.getNetwork({
 *     datacenterId: datacenter.id,
 *     name: "VM Network",
 *     filters: [{
 *         networkType: "Network",
 *     }],
 * }));
 * ```
 */
export function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getNetwork:getNetwork", {
        "datacenterId": args.datacenterId,
        "distributedVirtualSwitchUuid": args.distributedVirtualSwitchUuid,
        "filters": args.filters,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * The managed object reference ID
     * of the datacenter the network is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters,
     * use the `id` attribute from an empty `vsphere.Datacenter` data source.
     */
    datacenterId?: string;
    /**
     * For distributed port group type
     * network objects, the ID of the distributed virtual switch for which the port
     * group belongs. It is useful to differentiate port groups with same name using
     * the distributed virtual switch ID.
     */
    distributedVirtualSwitchUuid?: string;
    /**
     * Apply a filter for the discovered network.
     */
    filters?: inputs.GetNetworkFilter[];
    /**
     * The name of the network. This can be a name or path.
     */
    name: string;
}

/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    readonly datacenterId?: string;
    readonly distributedVirtualSwitchUuid?: string;
    readonly filters?: outputs.GetNetworkFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The managed object type for the discovered network. This will be one
     * of `DistributedVirtualPortgroup` for distributed port groups, `Network` for
     * standard (host-based) port groups, or `OpaqueNetwork` for networks managed
     * externally, such as those managed by NSX.
     */
    readonly type: string;
}
/**
 * The `vsphere.getNetwork` data source can be used to discover the ID of a network in
 * vSphere. This can be any network that can be used as the backing for a network
 * interface for `vsphere.VirtualMachine` or any other vSphere resource that
 * requires a network. This includes standard (host-based) port groups, distributed
 * port groups, or opaque networks such as those managed by NSX.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const network = datacenter.then(datacenter => vsphere.getNetwork({
 *     name: "VM Network",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 *
 * ### Additional Examples
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const myPortGroup = datacenter.then(datacenter => vsphere.getNetwork({
 *     datacenterId: datacenter.id,
 *     name: "VM Network",
 *     filters: [{
 *         networkType: "Network",
 *     }],
 * }));
 * ```
 */
export function getNetworkOutput(args: GetNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vsphere:index/getNetwork:getNetwork", {
        "datacenterId": args.datacenterId,
        "distributedVirtualSwitchUuid": args.distributedVirtualSwitchUuid,
        "filters": args.filters,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkOutputArgs {
    /**
     * The managed object reference ID
     * of the datacenter the network is located in. This can be omitted if the
     * search path used in `name` is an absolute path. For default datacenters,
     * use the `id` attribute from an empty `vsphere.Datacenter` data source.
     */
    datacenterId?: pulumi.Input<string>;
    /**
     * For distributed port group type
     * network objects, the ID of the distributed virtual switch for which the port
     * group belongs. It is useful to differentiate port groups with same name using
     * the distributed virtual switch ID.
     */
    distributedVirtualSwitchUuid?: pulumi.Input<string>;
    /**
     * Apply a filter for the discovered network.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.GetNetworkFilterArgs>[]>;
    /**
     * The name of the network. This can be a name or path.
     */
    name: pulumi.Input<string>;
}
