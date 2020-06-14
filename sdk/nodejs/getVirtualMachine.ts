// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere.VirtualMachine` data source can be used to find the UUID of an
 * existing virtual machine or template. Its most relevant purpose is for finding
 * the UUID of a template to be used as the source for cloning into a new
 * `vsphere.VirtualMachine` resource. It also
 * reads the guest ID so that can be supplied as well.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = pulumi.output(vsphere.getDatacenter({
 *     name: "dc1",
 * }, { async: true }));
 * const template = datacenter.apply(datacenter => vsphere.getVirtualMachine({
 *     datacenterId: datacenter.id,
 *     name: "test-vm-template",
 * }, { async: true }));
 * ```
 */
export function getVirtualMachine(args: GetVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vsphere:index/getVirtualMachine:getVirtualMachine", {
        "datacenterId": args.datacenterId,
        "name": args.name,
        "scsiControllerScanCount": args.scsiControllerScanCount,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualMachine.
 */
export interface GetVirtualMachineArgs {
    /**
     * The managed object reference
     * ID of the datacenter the virtual machine is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the `id` attribute from an empty
     * `vsphere.Datacenter` data source.
     */
    readonly datacenterId?: string;
    /**
     * The name of the virtual machine. This can be a name or
     * path.
     */
    readonly name: string;
    /**
     * The number of SCSI controllers to
     * scan for disk attributes and controller types on. Default: `1`.
     */
    readonly scsiControllerScanCount?: number;
}

/**
 * A collection of values returned by getVirtualMachine.
 */
export interface GetVirtualMachineResult {
    /**
     * The alternate guest name of the virtual machine when
     * guestId is a non-specific operating system, like `otherGuest`.
     */
    readonly alternateGuestName: string;
    readonly datacenterId?: string;
    /**
     * Information about each of the disks on this virtual machine or
     * template. These are sorted by bus and unit number so that they can be applied
     * to a `vsphere.VirtualMachine` resource in the order the resource expects
     * while cloning. This is useful for discovering certain disk settings while
     * performing a linked clone, as all settings that are output by this data
     * source must be the same on the destination virtual machine as the source.
     * Only the first number of controllers defined by `scsiControllerScanCount`
     * are scanned for disks. The sub-attributes are:
     */
    readonly disks: outputs.GetVirtualMachineDisk[];
    /**
     * The firmware type for this virtual machine. Can be `bios` or `efi`.
     */
    readonly firmware: string;
    /**
     * The guest ID of the virtual machine or template.
     */
    readonly guestId: string;
    /**
     * A list of IP addresses as reported by VMWare tools.
     */
    readonly guestIpAddresses: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The network interface types for each network
     * interface found on the virtual machine, in device bus order. Will be one of
     * `e1000`, `e1000e`, `pcnet32`, `sriov`, `vmxnet2`, or `vmxnet3`.
     */
    readonly networkInterfaceTypes: string[];
    /**
     * Mode for sharing the SCSI bus. The modes are
     * physicalSharing, virtualSharing, and noSharing. Only the first number of
     * controllers defined by `scsiControllerScanCount` are scanned.
     */
    readonly scsiBusSharing: string;
    readonly scsiControllerScanCount?: number;
    /**
     * The common type of all SCSI controllers on this virtual machine.
     * Will be one of `lsilogic` (LSI Logic Parallel), `lsilogic-sas` (LSI Logic
     * SAS), `pvscsi` (VMware Paravirtual), `buslogic` (BusLogic), or `mixed` when
     * there are multiple controller types. Only the first number of controllers
     * defined by `scsiControllerScanCount` are scanned.
     */
    readonly scsiType: string;
}
