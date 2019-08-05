// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/compute_cluster.html.markdown.
 */
export function getComputeCluster(args: GetComputeClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeClusterResult> & GetComputeClusterResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetComputeClusterResult> = pulumi.runtime.invoke("vsphere:index/getComputeCluster:getComputeCluster", {
        "datacenterId": args.datacenterId,
        "name": args.name,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getComputeCluster.
 */
export interface GetComputeClusterArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datacenter the cluster is located in.  This can
     * be omitted if the search path used in `name` is an absolute path.  For
     * default datacenters, use the id attribute from an empty `vsphere_datacenter`
     * data source.
     */
    readonly datacenterId?: string;
    /**
     * The name or absolute path to the cluster.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getComputeCluster.
 */
export interface GetComputeClusterResult {
    readonly datacenterId?: string;
    readonly name: string;
    readonly resourcePoolId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
