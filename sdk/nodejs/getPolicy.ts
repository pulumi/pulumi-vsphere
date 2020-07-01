// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere.getPolicy` data source can be used to discover the UUID of a
 * vSphere storage policy. This can then be used with resources or data sources that
 * require a storage policy.
 *
 * > **NOTE:** Storage policy support is unsupported on direct ESXi connections and
 * requires vCenter 6.0 or higher.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const policy = pulumi.output(vsphere.getPolicy({
 *     name: "policy1",
 * }, { async: true }));
 * ```
 */
export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
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
    readonly name: string;
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
