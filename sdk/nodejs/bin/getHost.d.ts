import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_host` data source can be used to discover the ID of a vSphere
 * host. This can then be used with resources or data sources that require a host
 * managed object reference ID.
 */
export declare function getHost(args: GetHostArgs, opts?: pulumi.InvokeOptions): Promise<GetHostResult>;
/**
 * A collection of arguments for invoking getHost.
 */
export interface GetHostArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of a datacenter.
     */
    readonly datacenterId: string;
    /**
     * The name of the host. This can be a name or path. Can be
     * omitted if there is only one host in your inventory.
     */
    readonly name?: string;
}
/**
 * A collection of values returned by getHost.
 */
export interface GetHostResult {
    /**
     * The [managed object ID][docs-about-morefs] of the host's
     * root resource pool.
     */
    readonly resourcePoolId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
