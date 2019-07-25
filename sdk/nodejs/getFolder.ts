// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere_folder` data source can be used to get the general attributes of a
 * vSphere inventory folder. Paths are absolute and include must include the
 * datacenter.  
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 * 
 * const folder = pulumi.output(vsphere.getFolder({
 *     path: "/dc1/datastore/folder1",
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/folder.html.markdown.
 */
export function getFolder(args: GetFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderResult> & GetFolderResult {
    const promise: Promise<GetFolderResult> = pulumi.runtime.invoke("vsphere:index/getFolder:getFolder", {
        "path": args.path,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getFolder.
 */
export interface GetFolderArgs {
    readonly path: string;
}

/**
 * A collection of values returned by getFolder.
 */
export interface GetFolderResult {
    readonly path: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
