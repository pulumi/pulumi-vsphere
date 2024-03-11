// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.getPolicy` data source can be used to discover the UUID of a
 * storage policy. This can then be used with other resources or data sources that
 * use a storage policy.
 *
 * > **NOTE:** Storage policies are not supported on direct ESXi hosts and
 * requires vCenter Server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const prodPlatinumReplicated = vsphere.getPolicy({
 *     name: "prod_platinum_replicated",
 * });
 * const devSilverNonreplicated = vsphere.getPolicy({
 *     name: "dev_silver_nonreplicated",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getPolicy:getPolicy", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    /**
     * The name of the storage policy.
     */
    name: string;
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
/**
 * The `vsphere.getPolicy` data source can be used to discover the UUID of a
 * storage policy. This can then be used with other resources or data sources that
 * use a storage policy.
 *
 * > **NOTE:** Storage policies are not supported on direct ESXi hosts and
 * requires vCenter Server.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const prodPlatinumReplicated = vsphere.getPolicy({
 *     name: "prod_platinum_replicated",
 * });
 * const devSilverNonreplicated = vsphere.getPolicy({
 *     name: "dev_silver_nonreplicated",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPolicyOutput(args: GetPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPolicyResult> {
    return pulumi.output(args).apply((a: any) => getPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyOutputArgs {
    /**
     * The name of the storage policy.
     */
    name: pulumi.Input<string>;
}
