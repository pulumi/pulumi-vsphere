// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.Folder` data source can be used to get the general attributes of a
 * vSphere inventory folder. Paths are absolute and must include the datacenter.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const folder = vsphere.getFolder({
 *     path: "/dc-01/datastore-01/folder-01",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getFolder(args: GetFolderArgs, opts?: pulumi.InvokeOptions): Promise<GetFolderResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
     * of `test-folder`, the resulting path would be
     * `/default-dc/vm/test-folder`. The valid folder types to be used in
     * the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
     */
    path: string;
}

/**
 * A collection of values returned by getFolder.
 */
export interface GetFolderResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly path: string;
}
/**
 * The `vsphere.Folder` data source can be used to get the general attributes of a
 * vSphere inventory folder. Paths are absolute and must include the datacenter.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const folder = vsphere.getFolder({
 *     path: "/dc-01/datastore-01/folder-01",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getFolderOutput(args: GetFolderOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFolderResult> {
    return pulumi.output(args).apply((a: any) => getFolder(a, opts))
}

/**
 * A collection of arguments for invoking getFolder.
 */
export interface GetFolderOutputArgs {
    /**
     * The absolute path of the folder. For example, given a
     * default datacenter of `default-dc`, a folder of type `vm`, and a folder name
     * of `test-folder`, the resulting path would be
     * `/default-dc/vm/test-folder`. The valid folder types to be used in
     * the path are: `vm`, `host`, `datacenter`, `datastore`, or `network`.
     */
    path: pulumi.Input<string>;
}
