// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphereThumbprint` data source can be used to discover the host
 * thumbprint of an ESXi host. This can be used when adding the `vsphere.Host`
 * resource. If the ESXi host is using a certificate chain, the first one returned
 * will be used to generate the thumbprint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const thumbprint = pulumi.output(vsphere.getHostThumbprint({
 *     address: "esxi-01.example.com",
 * }));
 * ```
 */
export function getHostThumbprint(args: GetHostThumbprintArgs, opts?: pulumi.InvokeOptions): Promise<GetHostThumbprintResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vsphere:index/getHostThumbprint:getHostThumbprint", {
        "address": args.address,
        "insecure": args.insecure,
        "port": args.port,
    }, opts);
}

/**
 * A collection of arguments for invoking getHostThumbprint.
 */
export interface GetHostThumbprintArgs {
    /**
     * The address of the ESXi host to retrieve the
     * thumbprint from.
     */
    address: string;
    /**
     * Disables SSL certificate verification.
     * Default: `false`
     */
    insecure?: boolean;
    /**
     * The port to use connecting to the ESXi host. Default: 443
     */
    port?: string;
}

/**
 * A collection of values returned by getHostThumbprint.
 */
export interface GetHostThumbprintResult {
    readonly address: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly insecure?: boolean;
    readonly port?: string;
}

export function getHostThumbprintOutput(args: GetHostThumbprintOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHostThumbprintResult> {
    return pulumi.output(args).apply(a => getHostThumbprint(a, opts))
}

/**
 * A collection of arguments for invoking getHostThumbprint.
 */
export interface GetHostThumbprintOutputArgs {
    /**
     * The address of the ESXi host to retrieve the
     * thumbprint from.
     */
    address: pulumi.Input<string>;
    /**
     * Disables SSL certificate verification.
     * Default: `false`
     */
    insecure?: pulumi.Input<boolean>;
    /**
     * The port to use connecting to the ESXi host. Default: 443
     */
    port?: pulumi.Input<string>;
}
