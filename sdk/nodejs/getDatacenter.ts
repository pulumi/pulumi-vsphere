// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.Datacenter` data source can be used to discover the ID of a
 * vSphere datacenter object. This can then be used with resources or data sources
 * that require a datacenter, such as the `vsphere.Host`
 * data source.
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
 * ```
 */
export function getDatacenter(args?: GetDatacenterArgs, opts?: pulumi.InvokeOptions): Promise<GetDatacenterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getDatacenter:getDatacenter", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatacenter.
 */
export interface GetDatacenterArgs {
    /**
     * The name of the datacenter. This can be a name or path.
     * Can be omitted if there is only one datacenter in the inventory.
     *
     * > **NOTE:** When used with an ESXi host, this data source _always_ returns the
     * host's "default" datacenter, which is a special datacenter name unrelated to the
     * datacenters that exist in the vSphere inventory when managed by a vCenter Server
     * instance. Hence, the `name` attribute is completely ignored.
     */
    name?: string;
}

/**
 * A collection of values returned by getDatacenter.
 */
export interface GetDatacenterResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
}
/**
 * The `vsphere.Datacenter` data source can be used to discover the ID of a
 * vSphere datacenter object. This can then be used with resources or data sources
 * that require a datacenter, such as the `vsphere.Host`
 * data source.
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
 * ```
 */
export function getDatacenterOutput(args?: GetDatacenterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatacenterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vsphere:index/getDatacenter:getDatacenter", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatacenter.
 */
export interface GetDatacenterOutputArgs {
    /**
     * The name of the datacenter. This can be a name or path.
     * Can be omitted if there is only one datacenter in the inventory.
     *
     * > **NOTE:** When used with an ESXi host, this data source _always_ returns the
     * host's "default" datacenter, which is a special datacenter name unrelated to the
     * datacenters that exist in the vSphere inventory when managed by a vCenter Server
     * instance. Hence, the `name` attribute is completely ignored.
     */
    name?: pulumi.Input<string>;
}
