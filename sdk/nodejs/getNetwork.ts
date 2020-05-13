// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere..getNetwork` data source can be used to discover the ID of a network
 * in vSphere. This can be any network that can be used as the backing for a
 * network interface for `vsphere..VirtualMachine` or any other vSphere resource
 * that requires a network. This includes standard (host-based) port groups, DVS
 * port groups, or opaque networks such as those managed by NSX.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }, { async: true }));
 * const net = datacenter.apply(datacenter => vsphere.getNetwork({
 *     datacenterId: datacenter.id,
 *     name: "test-net",
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/network.html.markdown.
 */
export function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vsphere:index/getNetwork:getNetwork", {
        "datacenterId": args.datacenterId,
        "distributedVirtualSwitchUuid": args.distributedVirtualSwitchUuid,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * The managed object reference
     * ID of the datacenter the network is located in. This can
     * be omitted if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere..Datacenter` data
     * source.
     */
    readonly datacenterId?: string;
    /**
     * For distributed port group type 
     * network objects, the ID of the distributed virtual switch the given port group
     * belongs to. It is useful to differentiate port groups with same name using the
     * Distributed virtual switch ID.
     */
    readonly distributedVirtualSwitchUuid?: string;
    /**
     * The name of the network. This can be a name or path.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getNetwork.
 */
export interface GetNetworkResult {
    readonly datacenterId?: string;
    readonly distributedVirtualSwitchUuid?: string;
    readonly name: string;
    readonly type: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
