// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `vsphere.getHostPciDevice` data source can be used to discover the device ID
 * of a vSphere host's PCI device. This can then be used with
 * `vsphere.VirtualMachine`'s `pciDeviceId`.
 *
 * ## Example Usage
 *
 * ### With Vendor ID And Class ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *     name: "esxi-01.example.com",
 *     datacenterId: datacenter.id,
 * }));
 * const dev = host.then(host => vsphere.getHostPciDevice({
 *     hostId: host.id,
 *     classId: "123",
 *     vendorId: "456",
 * }));
 * ```
 *
 * ### With Name Regular Expression
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *    name: "dc-01",
 * });
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *    name: "esxi-01.example.com",
 *    datacenterId: datacenter.id,
 * }));
 * const dev = host.then(host => vsphere.getHostPciDevice({
 *    hostId: host.id,
 *    nameRegex: "MMC",
 * }));
 * ```
 */
export function getHostPciDevice(args: GetHostPciDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetHostPciDeviceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getHostPciDevice:getHostPciDevice", {
        "classId": args.classId,
        "hostId": args.hostId,
        "nameRegex": args.nameRegex,
        "vendorId": args.vendorId,
    }, opts);
}

/**
 * A collection of arguments for invoking getHostPciDevice.
 */
export interface GetHostPciDeviceArgs {
    /**
     * The hexadecimal PCI device class ID
     *
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     *
     * > **NOTE:** `nameRegex`, `vendorId`, and `classId` can all be used together.
     */
    classId?: string;
    /**
     * The [managed object reference ID][docs-about-morefs] of
     * a host.
     */
    hostId: string;
    /**
     * A regular expression that will be used to match the
     * host PCI device name.
     */
    nameRegex?: string;
    /**
     * The hexadecimal PCI device vendor ID.
     */
    vendorId?: string;
}

/**
 * A collection of values returned by getHostPciDevice.
 */
export interface GetHostPciDeviceResult {
    readonly classId?: string;
    readonly hostId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the PCI device.
     */
    readonly name: string;
    readonly nameRegex?: string;
    readonly vendorId?: string;
}
/**
 * The `vsphere.getHostPciDevice` data source can be used to discover the device ID
 * of a vSphere host's PCI device. This can then be used with
 * `vsphere.VirtualMachine`'s `pciDeviceId`.
 *
 * ## Example Usage
 *
 * ### With Vendor ID And Class ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *     name: "esxi-01.example.com",
 *     datacenterId: datacenter.id,
 * }));
 * const dev = host.then(host => vsphere.getHostPciDevice({
 *     hostId: host.id,
 *     classId: "123",
 *     vendorId: "456",
 * }));
 * ```
 *
 * ### With Name Regular Expression
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *    name: "dc-01",
 * });
 * const host = datacenter.then(datacenter => vsphere.getHost({
 *    name: "esxi-01.example.com",
 *    datacenterId: datacenter.id,
 * }));
 * const dev = host.then(host => vsphere.getHostPciDevice({
 *    hostId: host.id,
 *    nameRegex: "MMC",
 * }));
 * ```
 */
export function getHostPciDeviceOutput(args: GetHostPciDeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHostPciDeviceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vsphere:index/getHostPciDevice:getHostPciDevice", {
        "classId": args.classId,
        "hostId": args.hostId,
        "nameRegex": args.nameRegex,
        "vendorId": args.vendorId,
    }, opts);
}

/**
 * A collection of arguments for invoking getHostPciDevice.
 */
export interface GetHostPciDeviceOutputArgs {
    /**
     * The hexadecimal PCI device class ID
     *
     * [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
     *
     * > **NOTE:** `nameRegex`, `vendorId`, and `classId` can all be used together.
     */
    classId?: pulumi.Input<string>;
    /**
     * The [managed object reference ID][docs-about-morefs] of
     * a host.
     */
    hostId: pulumi.Input<string>;
    /**
     * A regular expression that will be used to match the
     * host PCI device name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * The hexadecimal PCI device vendor ID.
     */
    vendorId?: pulumi.Input<string>;
}
