import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_virtual_machine` data source can be used to find the UUID of an
 * existing virtual machine or template. Its most relevant purpose is for finding
 * the UUID of a template to be used as the source for cloning into a new
 * [`vsphere_virtual_machine`][docs-virtual-machine-resource] resource. It also
 * reads the guest ID so that can be supplied as well.
 *
 * [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
 */
export declare function Virtual_machine(args: Virtual_machineArgs, opts?: pulumi.InvokeOptions): Promise<Virtual_machineResult>;
/**
 * A collection of arguments for invoking Virtual_machine.
 */
export interface Virtual_machineArgs {
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datacenter the virtual machine is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the `id` attribute from an empty
     * `vsphere_datacenter` data source.
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
 * A collection of values returned by Virtual_machine.
 */
export interface Virtual_machineResult {
    /**
     * The alternate guest name of the virtual machine when
     * guest_id is a non-specific operating system, like `otherGuest`.
     */
    readonly alternateGuestName: string;
    /**
     * Information about each of the disks on this virtual machine or
     * template. These are sorted by bus and unit number so that they can be applied
     * to a `vsphere_virtual_machine` resource in the order the resource expects
     * while cloning. This is useful for discovering certain disk settings while
     * performing a linked clone, as all settings that are output by this data
     * source must be the same on the destination virtual machine as the source.
     * Only the first number of controllers defined by `scsi_controller_scan_count`
     * are scanned for disks. The sub-attributes are:
     */
    readonly disks: {
        eagerlyScrub: boolean;
        size: number;
        thinProvisioned: boolean;
    }[];
    /**
     * The firmware type for this virtual machine. Can be `bios` or `efi`.
     */
    readonly firmware: string;
    /**
     * The guest ID of the virtual machine or template.
     */
    readonly guestId: string;
    /**
     * The network interface types for each network
     * interface found on the virtual machine, in device bus order. Will be one of
     * `e1000`, `e1000e`, `pcnet32`, `sriov`, `vmxnet2`, or `vmxnet3`.
     */
    readonly networkInterfaceTypes: string[];
    /**
     * Mode for sharing the SCSI bus. The modes are
     * physicalSharing, virtualSharing, and noSharing. Only the first number of
     * controllers defined by `scsi_controller_scan_count` are scanned.
     */
    readonly scsiBusSharing: string;
    /**
     * The common type of all SCSI controllers on this virtual machine.
     * Will be one of `lsilogic` (LSI Logic Parallel), `lsilogic-sas` (LSI Logic
     * SAS), `pvscsi` (VMware Paravirtual), `buslogic` (BusLogic), or `mixed` when
     * there are multiple controller types. Only the first number of controllers
     * defined by `scsi_controller_scan_count` are scanned.
     */
    readonly scsiType: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
