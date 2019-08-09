// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/network.html.markdown.
 */
export function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> & GetNetworkResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetNetworkResult> = pulumi.runtime.invoke("vsphere:index/getNetwork:getNetwork", {
        "datacenterId": args.datacenterId,
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getNetwork.
 */
export interface GetNetworkArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datacenter the network is located in. This can
     * be omitted if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere..Datacenter` data
     * source.
     */
    readonly datacenterId?: string;
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
    readonly name: string;
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
