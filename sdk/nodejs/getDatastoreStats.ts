// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.getDatastoreStats` data source can be used to retrieve the usage stats
 * of all vSphere datastore objects in a datacenter. This can then be used as a
 * standalone datasource to get information required as input to other data sources.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const datastoreStats = datacenter.then(datacenter => vsphere.getDatastoreStats({
 *     datacenterId: datacenter.id,
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * A usefull example of this datasource would be to determine the
 * datastore with the most free space. For example, in addition to
 * the above:
 *
 * Create an `outputs.tf` like that:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 *
 * export const maxFreeSpaceName = local.max_free_space_name;
 * export const maxFreeSpace = local.max_free_space;
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * and a `locals.tf` like that:
 */
export function getDatastoreStats(args: GetDatastoreStatsArgs, opts?: pulumi.InvokeOptions): Promise<GetDatastoreStatsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getDatastoreStats:getDatastoreStats", {
        "capacity": args.capacity,
        "datacenterId": args.datacenterId,
        "freeSpace": args.freeSpace,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatastoreStats.
 */
export interface GetDatastoreStatsArgs {
    /**
     * A mapping of the capacity for all datastore in the datacenter
     * , where the name of the datastore is used as key and the capacity as value.
     */
    capacity?: {[key: string]: any};
    /**
     * The [managed object reference ID][docs-about-morefs]
     * of the datacenter the datastores are located in. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     *
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     */
    datacenterId: string;
    /**
     * A mapping of the free space for each datastore in the
     * datacenter, where the name of the datastore is used as key and the free
     * space as value.
     */
    freeSpace?: {[key: string]: any};
}

/**
 * A collection of values returned by getDatastoreStats.
 */
export interface GetDatastoreStatsResult {
    /**
     * A mapping of the capacity for all datastore in the datacenter
     * , where the name of the datastore is used as key and the capacity as value.
     */
    readonly capacity?: {[key: string]: any};
    /**
     * The [managed object reference ID][docs-about-morefs]
     * of the datacenter the datastores are located in.
     */
    readonly datacenterId: string;
    /**
     * A mapping of the free space for each datastore in the
     * datacenter, where the name of the datastore is used as key and the free
     * space as value.
     */
    readonly freeSpace?: {[key: string]: any};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * The `vsphere.getDatastoreStats` data source can be used to retrieve the usage stats
 * of all vSphere datastore objects in a datacenter. This can then be used as a
 * standalone datasource to get information required as input to other data sources.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const datastoreStats = datacenter.then(datacenter => vsphere.getDatastoreStats({
 *     datacenterId: datacenter.id,
 * }));
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * A usefull example of this datasource would be to determine the
 * datastore with the most free space. For example, in addition to
 * the above:
 *
 * Create an `outputs.tf` like that:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 *
 * export const maxFreeSpaceName = local.max_free_space_name;
 * export const maxFreeSpace = local.max_free_space;
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * and a `locals.tf` like that:
 */
export function getDatastoreStatsOutput(args: GetDatastoreStatsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatastoreStatsResult> {
    return pulumi.output(args).apply((a: any) => getDatastoreStats(a, opts))
}

/**
 * A collection of arguments for invoking getDatastoreStats.
 */
export interface GetDatastoreStatsOutputArgs {
    /**
     * A mapping of the capacity for all datastore in the datacenter
     * , where the name of the datastore is used as key and the capacity as value.
     */
    capacity?: pulumi.Input<{[key: string]: any}>;
    /**
     * The [managed object reference ID][docs-about-morefs]
     * of the datacenter the datastores are located in. For default datacenters, use
     * the `id` attribute from an empty `vsphere.Datacenter` data source.
     *
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     */
    datacenterId: pulumi.Input<string>;
    /**
     * A mapping of the free space for each datastore in the
     * datacenter, where the name of the datastore is used as key and the free
     * space as value.
     */
    freeSpace?: pulumi.Input<{[key: string]: any}>;
}