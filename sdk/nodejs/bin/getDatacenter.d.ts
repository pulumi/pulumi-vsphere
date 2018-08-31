import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_datacenter` data source can be used to discover the ID of a
 * vSphere datacenter. This can then be used with resources or data sources that
 * require a datacenter, such as the [`vsphere_host`][data-source-vsphere-host]
 * data source.
 *
 * [data-source-vsphere-host]: /docs/providers/vsphere/d/host.html
 */
export declare function getDatacenter(args?: GetDatacenterArgs, opts?: pulumi.InvokeOptions): Promise<GetDatacenterResult>;
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
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
