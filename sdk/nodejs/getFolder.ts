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
 */
export function getFolder(args: GetFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderResult> {
    return pulumi.runtime.invoke("vsphere:index/getFolder:getFolder", {
        "path": args.path,
    }, opts);
}

/**
 * A collection of arguments for invoking getFolder.
 */
export interface GetFolderArgs {
    /**
     * The absolute path of the folder. For example, given a
     * default datacenter of `default-dc`, a folder of type `vm`, and a folder name
     * of `terraform-test-folder`, the resulting path would be
     * `/default-dc/vm/terraform-test-folder`. The valid folder types to be used in
     * the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
     */
    readonly path: string;
}

/**
 * A collection of values returned by getFolder.
 */
export interface GetFolderResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
