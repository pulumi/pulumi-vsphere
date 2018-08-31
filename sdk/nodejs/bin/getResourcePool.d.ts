import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_resource_pool` data source can be used to discover the ID of a
 * resource pool in vSphere. This is useful to fetch the ID of a resource pool
 * that you want to use to create virtual machines in using the
 * [`vsphere_virtual_machine`][docs-virtual-machine-resource] resource.
 *
 * [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
 */
export declare function getResourcePool(args?: GetResourcePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePoolResult>;
/**
 * A collection of arguments for invoking getResourcePool.
 */
export interface GetResourcePoolArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datacenter the resource pool is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the id attribute from an empty
     * `vsphere_datacenter` data source.
     */
    readonly datacenterId?: string;
    /**
     * The name of the resource pool. This can be a name or
     * path. This is required when using vCenter.
     */
    readonly name?: string;
}
/**
 * A collection of values returned by getResourcePool.
 */
export interface GetResourcePoolResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
