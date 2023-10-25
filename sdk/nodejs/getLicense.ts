// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.License` data source can be used to get the general attributes of
 * a license keys from a vCenter Server instance.
 */
export function getLicense(args: GetLicenseArgs, opts?: pulumi.InvokeOptions): Promise<GetLicenseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getLicense:getLicense", {
        "licenseKey": args.licenseKey,
    }, opts);
}

/**
 * A collection of arguments for invoking getLicense.
 */
export interface GetLicenseArgs {
    /**
     * The license key.
     */
    licenseKey: string;
}

/**
 * A collection of values returned by getLicense.
 */
export interface GetLicenseResult {
    /**
     * The product edition of the license key.
     */
    readonly editionKey: string;
    readonly id: string;
    /**
     * A map of key/value pairs attached as labels (tags) to the license key.
     */
    readonly labels: {[key: string]: string};
    readonly licenseKey: string;
    /**
     * The display name for the license.
     */
    readonly name: string;
    /**
     * Total number of units (example: CPUs) contained in the license.
     */
    readonly total: number;
    /**
     * The number of units (example: CPUs) assigned to this license.
     */
    readonly used: number;
}
/**
 * The `vsphere.License` data source can be used to get the general attributes of
 * a license keys from a vCenter Server instance.
 */
export function getLicenseOutput(args: GetLicenseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLicenseResult> {
    return pulumi.output(args).apply((a: any) => getLicense(a, opts))
}

/**
 * A collection of arguments for invoking getLicense.
 */
export interface GetLicenseOutputArgs {
    /**
     * The license key.
     */
    licenseKey: pulumi.Input<string>;
}
