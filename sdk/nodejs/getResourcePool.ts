// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.ResourcePool` data source can be used to discover the ID of a
 * resource pool in vSphere. This is useful to return the ID of a resource pool
 * that you want to use to create virtual machines in using the
 * `vsphere.VirtualMachine` resource.
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
 * const pool = datacenter.then(datacenter => vsphere.getResourcePool({
 *     name: "resource-pool-01",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 *
 * ### Specifying the Root Resource Pool for a Standalone ESXi Host
 *
 * > **NOTE:** Returning the root resource pool for a cluster can be done
 * directly via the `vsphere.ComputeCluster`
 * data source.
 *
 * All compute resources in vSphere have a resource pool, even if one has not been
 * explicitly created. This resource pool is referred to as the *root resource
 * pool* and can be looked up by specifying the path.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const pool = vsphere.getResourcePool({
 *     name: "esxi-01.example.com/Resources",
 *     datacenterId: datacenter.id,
 * });
 * ```
 *
 * For more information on the root resource pool, see
 * [Managing Resource Pools](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-60077B40-66FF-4625-934A-641703ED7601.html) in the vSphere
 * documentation.
 */
export function getResourcePool(args?: GetResourcePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePoolResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getResourcePool:getResourcePool", {
        "datacenterId": args.datacenterId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourcePool.
 */
export interface GetResourcePoolArgs {
    /**
     * The managed object reference ID
     * of the datacenter in which the resource pool is located. This can be omitted
     * if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
     * source.
     *
     * > **Note:** When using ESXi without a vCenter Server instance, you do not
     * need to specify either attribute to use this data source. An empty declaration
     * will load the ESXi host's root resource pool.
     */
    datacenterId?: string;
    /**
     * The name of the resource pool. This can be a name or
     * path. This is required when using vCenter.
     */
    name?: string;
}

/**
 * A collection of values returned by getResourcePool.
 */
export interface GetResourcePoolResult {
    readonly datacenterId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
}
/**
 * The `vsphere.ResourcePool` data source can be used to discover the ID of a
 * resource pool in vSphere. This is useful to return the ID of a resource pool
 * that you want to use to create virtual machines in using the
 * `vsphere.VirtualMachine` resource.
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
 * const pool = datacenter.then(datacenter => vsphere.getResourcePool({
 *     name: "resource-pool-01",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 *
 * ### Specifying the Root Resource Pool for a Standalone ESXi Host
 *
 * > **NOTE:** Returning the root resource pool for a cluster can be done
 * directly via the `vsphere.ComputeCluster`
 * data source.
 *
 * All compute resources in vSphere have a resource pool, even if one has not been
 * explicitly created. This resource pool is referred to as the *root resource
 * pool* and can be looked up by specifying the path.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const pool = vsphere.getResourcePool({
 *     name: "esxi-01.example.com/Resources",
 *     datacenterId: datacenter.id,
 * });
 * ```
 *
 * For more information on the root resource pool, see
 * [Managing Resource Pools](https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-60077B40-66FF-4625-934A-641703ED7601.html) in the vSphere
 * documentation.
 */
export function getResourcePoolOutput(args?: GetResourcePoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcePoolResult> {
    return pulumi.output(args).apply((a: any) => getResourcePool(a, opts))
}

/**
 * A collection of arguments for invoking getResourcePool.
 */
export interface GetResourcePoolOutputArgs {
    /**
     * The managed object reference ID
     * of the datacenter in which the resource pool is located. This can be omitted
     * if the search path used in `name` is an absolute path. For default
     * datacenters, use the id attribute from an empty `vsphere.Datacenter` data
     * source.
     *
     * > **Note:** When using ESXi without a vCenter Server instance, you do not
     * need to specify either attribute to use this data source. An empty declaration
     * will load the ESXi host's root resource pool.
     */
    datacenterId?: pulumi.Input<string>;
    /**
     * The name of the resource pool. This can be a name or
     * path. This is required when using vCenter.
     */
    name?: pulumi.Input<string>;
}
