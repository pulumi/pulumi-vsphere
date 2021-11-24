// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.DatastoreCluster` data source can be used to discover the ID of a
 * datastore cluster in vSphere. This is useful to fetch the ID of a datastore
 * cluster that you want to use to assign datastores to using the
 * `vsphere.NasDatastore` or
 * `vsphere.VmfsDatastore` resources, or create
 * virtual machines in using the
 * `vsphere.VirtualMachine` resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }));
 * const datastoreCluster = vsphere_datacenter_dc.id.apply(id => vsphere.getDatastoreCluster({
 *     datacenterId: id,
 *     name: "datastore-cluster1",
 * }));
 * ```
 */
export function getDatastoreCluster(args: GetDatastoreClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetDatastoreClusterResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vsphere:index/getDatastoreCluster:getDatastoreCluster", {
        "datacenterId": args.datacenterId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatastoreCluster.
 */
export interface GetDatastoreClusterArgs {
    /**
     * The managed object reference
     * ID of the datacenter the datastore cluster is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the id attribute from an empty
     * `vsphere.Datacenter` data source.
     */
    datacenterId?: string;
    /**
     * The name or absolute path to the datastore cluster.
     */
    name: string;
}

/**
 * A collection of values returned by getDatastoreCluster.
 */
export interface GetDatastoreClusterResult {
    readonly datacenterId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}

export function getDatastoreClusterOutput(args: GetDatastoreClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatastoreClusterResult> {
    return pulumi.output(args).apply(a => getDatastoreCluster(a, opts))
}

/**
 * A collection of arguments for invoking getDatastoreCluster.
 */
export interface GetDatastoreClusterOutputArgs {
    /**
     * The managed object reference
     * ID of the datacenter the datastore cluster is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the id attribute from an empty
     * `vsphere.Datacenter` data source.
     */
    datacenterId?: pulumi.Input<string>;
    /**
     * The name or absolute path to the datastore cluster.
     */
    name: pulumi.Input<string>;
}
