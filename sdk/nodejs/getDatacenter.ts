// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere..Datacenter` data source can be used to discover the ID of a
 * vSphere datacenter. This can then be used with resources or data sources that
 * require a datacenter, such as the [`vsphere..getHost`][data-source-vsphere-host]
 * data source.
 * 
 * [data-source-vsphere-host]: /docs/providers/vsphere/d/host.html
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc1",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/datacenter.html.markdown.
 */
export function getDatacenter(args?: GetDatacenterArgs, opts?: pulumi.InvokeOptions): Promise<GetDatacenterResult> & GetDatacenterResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetDatacenterResult> = pulumi.runtime.invoke("vsphere:index/getDatacenter:getDatacenter", {
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getDatacenter.
 */
export interface GetDatacenterArgs {
    /**
     * The name of the datacenter. This can be a name or path.
     * Can be omitted if there is only one datacenter in your inventory.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getDatacenter.
 */
export interface GetDatacenterResult {
    readonly name?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
