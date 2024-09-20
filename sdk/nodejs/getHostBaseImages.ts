// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.getHostBaseImages` data source can be used to get the list of ESXi
 * base images available for cluster software management.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const baseImages = vsphere.getHostBaseImages({});
 * ```
 */
export function getHostBaseImages(opts?: pulumi.InvokeOptions): Promise<GetHostBaseImagesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getHostBaseImages:getHostBaseImages", {
    }, opts);
}

/**
 * A collection of values returned by getHostBaseImages.
 */
export interface GetHostBaseImagesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ESXi version identifier for the image
     */
    readonly versions: string[];
}
/**
 * The `vsphere.getHostBaseImages` data source can be used to get the list of ESXi
 * base images available for cluster software management.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const baseImages = vsphere.getHostBaseImages({});
 * ```
 */
export function getHostBaseImagesOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetHostBaseImagesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vsphere:index/getHostBaseImages:getHostBaseImages", {
    }, opts);
}
