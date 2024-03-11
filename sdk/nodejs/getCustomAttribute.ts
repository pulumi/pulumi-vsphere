// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.CustomAttribute` data source can be used to reference custom
 * attributes that are not managed by this provider. Its attributes are exactly the
 * same as the `vsphere.CustomAttribute` resource,
 * and, like importing, the data source takes a name argument for the search. The
 * `id` and other attributes are then populated with the data found by the search.
 *
 * > **NOTE:** Custom attributes are unsupported on direct ESXi host connections
 * and require vCenter Server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const attribute = vsphere.getCustomAttribute({
 *     name: "test-attribute",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCustomAttribute(args: GetCustomAttributeArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomAttributeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getCustomAttribute:getCustomAttribute", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustomAttribute.
 */
export interface GetCustomAttributeArgs {
    /**
     * The name of the custom attribute.
     */
    name: string;
}

/**
 * A collection of values returned by getCustomAttribute.
 */
export interface GetCustomAttributeResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly managedObjectType: string;
    readonly name: string;
}
/**
 * The `vsphere.CustomAttribute` data source can be used to reference custom
 * attributes that are not managed by this provider. Its attributes are exactly the
 * same as the `vsphere.CustomAttribute` resource,
 * and, like importing, the data source takes a name argument for the search. The
 * `id` and other attributes are then populated with the data found by the search.
 *
 * > **NOTE:** Custom attributes are unsupported on direct ESXi host connections
 * and require vCenter Server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const attribute = vsphere.getCustomAttribute({
 *     name: "test-attribute",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCustomAttributeOutput(args: GetCustomAttributeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomAttributeResult> {
    return pulumi.output(args).apply((a: any) => getCustomAttribute(a, opts))
}

/**
 * A collection of arguments for invoking getCustomAttribute.
 */
export interface GetCustomAttributeOutputArgs {
    /**
     * The name of the custom attribute.
     */
    name: pulumi.Input<string>;
}
