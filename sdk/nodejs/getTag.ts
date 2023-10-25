// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.Tag` data source can be used to reference tags that are not
 * managed by this provider. Its attributes are exactly the same as the `vsphere.Tag`
 * resource, and, like importing, the data source takes a name and
 * category to search on. The `id` and other attributes are then populated with
 * the data found by the search.
 *
 * > **NOTE:** Tagging is not supported on direct ESXi hosts connections and
 * requires vCenter Server.
 */
export function getTag(args: GetTagArgs, opts?: pulumi.InvokeOptions): Promise<GetTagResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getTag:getTag", {
        "categoryId": args.categoryId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTag.
 */
export interface GetTagArgs {
    /**
     * The ID of the tag category in which the tag is
     * located.
     */
    categoryId: string;
    /**
     * The name of the tag.
     */
    name: string;
}

/**
 * A collection of values returned by getTag.
 */
export interface GetTagResult {
    readonly categoryId: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
/**
 * The `vsphere.Tag` data source can be used to reference tags that are not
 * managed by this provider. Its attributes are exactly the same as the `vsphere.Tag`
 * resource, and, like importing, the data source takes a name and
 * category to search on. The `id` and other attributes are then populated with
 * the data found by the search.
 *
 * > **NOTE:** Tagging is not supported on direct ESXi hosts connections and
 * requires vCenter Server.
 */
export function getTagOutput(args: GetTagOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTagResult> {
    return pulumi.output(args).apply((a: any) => getTag(a, opts))
}

/**
 * A collection of arguments for invoking getTag.
 */
export interface GetTagOutputArgs {
    /**
     * The ID of the tag category in which the tag is
     * located.
     */
    categoryId: pulumi.Input<string>;
    /**
     * The name of the tag.
     */
    name: pulumi.Input<string>;
}
