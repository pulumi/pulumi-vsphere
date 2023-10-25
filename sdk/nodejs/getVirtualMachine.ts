// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `vsphere.VirtualMachine` data source can be used to find the UUID of an
 * existing virtual machine or template. The most common purpose is for finding
 * the UUID of a template to be used as the source for cloning to a new
 * `vsphere.VirtualMachine` resource. It also
 * reads the guest ID so that can be supplied as well.
 *
 * ## Example Usage
 *
 * In the following example, a virtual machine template is returned by its
 * unique name within the `vsphere.Datacenter`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const template = datacenter.then(datacenter => vsphere.getVirtualMachine({
 *     name: "ubuntu-server-template",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 * In the following example, each virtual machine template is returned by its
 * unique full path within the `vsphere.Datacenter`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const productionTemplate = datacenter.then(datacenter => vsphere.getVirtualMachine({
 *     name: "production/templates/ubuntu-server-template",
 *     datacenterId: datacenter.id,
 * }));
 * const developmentTemplate = datacenter.then(datacenter => vsphere.getVirtualMachine({
 *     name: "development/templates/ubuntu-server-template",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 */
export function getVirtualMachine(args?: GetVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vsphere:index/getVirtualMachine:getVirtualMachine", {
        "alternateGuestName": args.alternateGuestName,
        "annotation": args.annotation,
        "bootDelay": args.bootDelay,
        "bootRetryDelay": args.bootRetryDelay,
        "bootRetryEnabled": args.bootRetryEnabled,
        "cpuHotAddEnabled": args.cpuHotAddEnabled,
        "cpuHotRemoveEnabled": args.cpuHotRemoveEnabled,
        "cpuLimit": args.cpuLimit,
        "cpuPerformanceCountersEnabled": args.cpuPerformanceCountersEnabled,
        "cpuReservation": args.cpuReservation,
        "cpuShareCount": args.cpuShareCount,
        "cpuShareLevel": args.cpuShareLevel,
        "datacenterId": args.datacenterId,
        "efiSecureBootEnabled": args.efiSecureBootEnabled,
        "enableDiskUuid": args.enableDiskUuid,
        "enableLogging": args.enableLogging,
        "eptRviMode": args.eptRviMode,
        "extraConfig": args.extraConfig,
        "extraConfigRebootRequired": args.extraConfigRebootRequired,
        "firmware": args.firmware,
        "guestId": args.guestId,
        "hardwareVersion": args.hardwareVersion,
        "hvMode": args.hvMode,
        "ideControllerScanCount": args.ideControllerScanCount,
        "latencySensitivity": args.latencySensitivity,
        "memory": args.memory,
        "memoryHotAddEnabled": args.memoryHotAddEnabled,
        "memoryLimit": args.memoryLimit,
        "memoryReservation": args.memoryReservation,
        "memoryShareCount": args.memoryShareCount,
        "memoryShareLevel": args.memoryShareLevel,
        "moid": args.moid,
        "name": args.name,
        "nestedHvEnabled": args.nestedHvEnabled,
        "numCoresPerSocket": args.numCoresPerSocket,
        "numCpus": args.numCpus,
        "replaceTrigger": args.replaceTrigger,
        "runToolsScriptsAfterPowerOn": args.runToolsScriptsAfterPowerOn,
        "runToolsScriptsAfterResume": args.runToolsScriptsAfterResume,
        "runToolsScriptsBeforeGuestReboot": args.runToolsScriptsBeforeGuestReboot,
        "runToolsScriptsBeforeGuestShutdown": args.runToolsScriptsBeforeGuestShutdown,
        "runToolsScriptsBeforeGuestStandby": args.runToolsScriptsBeforeGuestStandby,
        "sataControllerScanCount": args.sataControllerScanCount,
        "scsiControllerScanCount": args.scsiControllerScanCount,
        "storagePolicyId": args.storagePolicyId,
        "swapPlacementPolicy": args.swapPlacementPolicy,
        "syncTimeWithHost": args.syncTimeWithHost,
        "syncTimeWithHostPeriodically": args.syncTimeWithHostPeriodically,
        "toolsUpgradePolicy": args.toolsUpgradePolicy,
        "uuid": args.uuid,
        "vapp": args.vapp,
        "vbsEnabled": args.vbsEnabled,
        "vvtdEnabled": args.vvtdEnabled,
    }, opts);
}

/**
 * A collection of arguments for invoking getVirtualMachine.
 */
export interface GetVirtualMachineArgs {
    /**
     * The alternate guest name of the virtual machine when
     * `guestId` is a non-specific operating system, like `otherGuest` or `otherGuest64`.
     */
    alternateGuestName?: string;
    /**
     * The user-provided description of this virtual machine.
     */
    annotation?: string;
    bootDelay?: number;
    bootRetryDelay?: number;
    bootRetryEnabled?: boolean;
    cpuHotAddEnabled?: boolean;
    cpuHotRemoveEnabled?: boolean;
    cpuLimit?: number;
    cpuPerformanceCountersEnabled?: boolean;
    cpuReservation?: number;
    cpuShareCount?: number;
    cpuShareLevel?: string;
    /**
     * The managed object reference
     * ID of the datacenter the virtual machine is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the `id` attribute from an empty
     * `vsphere.Datacenter` data source.
     */
    datacenterId?: string;
    efiSecureBootEnabled?: boolean;
    enableDiskUuid?: boolean;
    enableLogging?: boolean;
    eptRviMode?: string;
    extraConfig?: {[key: string]: string};
    extraConfigRebootRequired?: boolean;
    /**
     * The firmware type for this virtual machine. Can be `bios` or `efi`.
     */
    firmware?: string;
    /**
     * The guest ID of the virtual machine or template.
     */
    guestId?: string;
    /**
     * The hardware version number on this virtual machine.
     */
    hardwareVersion?: number;
    hvMode?: string;
    ideControllerScanCount?: number;
    latencySensitivity?: string;
    /**
     * The size of the virtual machine's memory, in MB.
     */
    memory?: number;
    memoryHotAddEnabled?: boolean;
    memoryLimit?: number;
    memoryReservation?: number;
    memoryShareCount?: number;
    memoryShareLevel?: string;
    moid?: string;
    /**
     * The name of the virtual machine. This can be a name or
     * the full path relative to the datacenter. This is required if a UUID lookup
     * is not performed.
     */
    name?: string;
    nestedHvEnabled?: boolean;
    /**
     * The number of cores per socket for this virtual machine.
     */
    numCoresPerSocket?: number;
    /**
     * The total number of virtual processor cores assigned to this
     * virtual machine.
     */
    numCpus?: number;
    replaceTrigger?: string;
    runToolsScriptsAfterPowerOn?: boolean;
    runToolsScriptsAfterResume?: boolean;
    runToolsScriptsBeforeGuestReboot?: boolean;
    runToolsScriptsBeforeGuestShutdown?: boolean;
    runToolsScriptsBeforeGuestStandby?: boolean;
    sataControllerScanCount?: number;
    /**
     * The number of SCSI controllers to
     * scan for disk attributes and controller types on. Default: `1`.
     *
     * > **NOTE:** For best results, ensure that all the disks on any templates you
     * use with this data source reside on the primary controller, and leave this
     * value at the default. See the `vsphere.VirtualMachine`
     * resource documentation for the significance of this setting, specifically the
     * additional requirements and notes for cloning section.
     */
    scsiControllerScanCount?: number;
    storagePolicyId?: string;
    swapPlacementPolicy?: string;
    syncTimeWithHost?: boolean;
    syncTimeWithHostPeriodically?: boolean;
    toolsUpgradePolicy?: string;
    /**
     * Specify this field for a UUID lookup, `name` and `datacenterId`
     * are not required if this is specified.
     */
    uuid?: string;
    vapp?: inputs.GetVirtualMachineVapp;
    vbsEnabled?: boolean;
    vvtdEnabled?: boolean;
}

/**
 * A collection of values returned by getVirtualMachine.
 */
export interface GetVirtualMachineResult {
    /**
     * The alternate guest name of the virtual machine when
     * `guestId` is a non-specific operating system, like `otherGuest` or `otherGuest64`.
     */
    readonly alternateGuestName?: string;
    /**
     * The user-provided description of this virtual machine.
     */
    readonly annotation: string;
    readonly bootDelay?: number;
    readonly bootRetryDelay?: number;
    readonly bootRetryEnabled?: boolean;
    readonly changeVersion: string;
    readonly cpuHotAddEnabled?: boolean;
    readonly cpuHotRemoveEnabled?: boolean;
    readonly cpuLimit?: number;
    readonly cpuPerformanceCountersEnabled?: boolean;
    readonly cpuReservation?: number;
    readonly cpuShareCount: number;
    readonly cpuShareLevel?: string;
    readonly datacenterId?: string;
    /**
     * Whenever possible, this is the first IPv4 address that is reachable through
     * the default gateway configured on the machine, then the first reachable IPv6
     * address, and then the first general discovered address if neither exist. If
     * VMware Tools is not running on the virtual machine, or if the VM is powered
     * off, this value will be blank.
     */
    readonly defaultIpAddress: string;
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
    readonly efiSecureBootEnabled?: boolean;
    readonly enableDiskUuid?: boolean;
    readonly enableLogging?: boolean;
    readonly eptRviMode?: string;
    readonly extraConfig?: {[key: string]: string};
    readonly extraConfigRebootRequired?: boolean;
    /**
     * The firmware type for this virtual machine. Can be `bios` or `efi`.
     */
    readonly firmware?: string;
    /**
     * The guest ID of the virtual machine or template.
     */
    readonly guestId: string;
    /**
     * A list of IP addresses as reported by VMware Tools.
     */
    readonly guestIpAddresses: string[];
    /**
     * The hardware version number on this virtual machine.
     */
    readonly hardwareVersion: number;
    readonly hvMode?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ideControllerScanCount?: number;
    readonly latencySensitivity?: string;
    /**
     * The size of the virtual machine's memory, in MB.
     */
    readonly memory?: number;
    readonly memoryHotAddEnabled?: boolean;
    readonly memoryLimit?: number;
    readonly memoryReservation?: number;
    readonly memoryShareCount: number;
    readonly memoryShareLevel?: string;
    readonly moid: string;
    readonly name?: string;
    readonly nestedHvEnabled?: boolean;
    /**
     * The network interface types for each network
     * interface found on the virtual machine, in device bus order. Will be one of
     * `e1000`, `e1000e`, `pcnet32`, `sriov`, `vmxnet2`, `vmxnet3vrdma`, or `vmxnet3`.
     */
    readonly networkInterfaceTypes: string[];
    /**
     * Information about each of the network interfaces on this 
     * virtual machine or template. These are sorted by device bus order so that they
     * can be applied to a `vsphere.VirtualMachine` resource in the order the resource
     * expects while cloning. This is useful for discovering certain network interface
     * settings while performing a linked clone, as all settings that are output by this
     * data source must be the same on the destination virtual machine as the source.
     * The sub-attributes are:
     */
    readonly networkInterfaces: outputs.GetVirtualMachineNetworkInterface[];
    /**
     * The number of cores per socket for this virtual machine.
     */
    readonly numCoresPerSocket?: number;
    /**
     * The total number of virtual processor cores assigned to this
     * virtual machine.
     */
    readonly numCpus?: number;
    readonly replaceTrigger?: string;
    readonly runToolsScriptsAfterPowerOn?: boolean;
    readonly runToolsScriptsAfterResume?: boolean;
    readonly runToolsScriptsBeforeGuestReboot?: boolean;
    readonly runToolsScriptsBeforeGuestShutdown?: boolean;
    readonly runToolsScriptsBeforeGuestStandby?: boolean;
    readonly sataControllerScanCount?: number;
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
    readonly storagePolicyId: string;
    readonly swapPlacementPolicy?: string;
    readonly syncTimeWithHost?: boolean;
    readonly syncTimeWithHostPeriodically?: boolean;
    readonly toolsUpgradePolicy?: string;
    readonly uuid: string;
    readonly vapp?: outputs.GetVirtualMachineVapp;
    readonly vappTransports: string[];
    readonly vbsEnabled?: boolean;
    readonly vvtdEnabled?: boolean;
}
/**
 * The `vsphere.VirtualMachine` data source can be used to find the UUID of an
 * existing virtual machine or template. The most common purpose is for finding
 * the UUID of a template to be used as the source for cloning to a new
 * `vsphere.VirtualMachine` resource. It also
 * reads the guest ID so that can be supplied as well.
 *
 * ## Example Usage
 *
 * In the following example, a virtual machine template is returned by its
 * unique name within the `vsphere.Datacenter`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const template = datacenter.then(datacenter => vsphere.getVirtualMachine({
 *     name: "ubuntu-server-template",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 * In the following example, each virtual machine template is returned by its
 * unique full path within the `vsphere.Datacenter`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vsphere from "@pulumi/vsphere";
 *
 * const datacenter = vsphere.getDatacenter({
 *     name: "dc-01",
 * });
 * const productionTemplate = datacenter.then(datacenter => vsphere.getVirtualMachine({
 *     name: "production/templates/ubuntu-server-template",
 *     datacenterId: datacenter.id,
 * }));
 * const developmentTemplate = datacenter.then(datacenter => vsphere.getVirtualMachine({
 *     name: "development/templates/ubuntu-server-template",
 *     datacenterId: datacenter.id,
 * }));
 * ```
 */
export function getVirtualMachineOutput(args?: GetVirtualMachineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualMachineResult> {
    return pulumi.output(args).apply((a: any) => getVirtualMachine(a, opts))
}

/**
 * A collection of arguments for invoking getVirtualMachine.
 */
export interface GetVirtualMachineOutputArgs {
    /**
     * The alternate guest name of the virtual machine when
     * `guestId` is a non-specific operating system, like `otherGuest` or `otherGuest64`.
     */
    alternateGuestName?: pulumi.Input<string>;
    /**
     * The user-provided description of this virtual machine.
     */
    annotation?: pulumi.Input<string>;
    bootDelay?: pulumi.Input<number>;
    bootRetryDelay?: pulumi.Input<number>;
    bootRetryEnabled?: pulumi.Input<boolean>;
    cpuHotAddEnabled?: pulumi.Input<boolean>;
    cpuHotRemoveEnabled?: pulumi.Input<boolean>;
    cpuLimit?: pulumi.Input<number>;
    cpuPerformanceCountersEnabled?: pulumi.Input<boolean>;
    cpuReservation?: pulumi.Input<number>;
    cpuShareCount?: pulumi.Input<number>;
    cpuShareLevel?: pulumi.Input<string>;
    /**
     * The managed object reference
     * ID of the datacenter the virtual machine is located in.
     * This can be omitted if the search path used in `name` is an absolute path.
     * For default datacenters, use the `id` attribute from an empty
     * `vsphere.Datacenter` data source.
     */
    datacenterId?: pulumi.Input<string>;
    efiSecureBootEnabled?: pulumi.Input<boolean>;
    enableDiskUuid?: pulumi.Input<boolean>;
    enableLogging?: pulumi.Input<boolean>;
    eptRviMode?: pulumi.Input<string>;
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    extraConfigRebootRequired?: pulumi.Input<boolean>;
    /**
     * The firmware type for this virtual machine. Can be `bios` or `efi`.
     */
    firmware?: pulumi.Input<string>;
    /**
     * The guest ID of the virtual machine or template.
     */
    guestId?: pulumi.Input<string>;
    /**
     * The hardware version number on this virtual machine.
     */
    hardwareVersion?: pulumi.Input<number>;
    hvMode?: pulumi.Input<string>;
    ideControllerScanCount?: pulumi.Input<number>;
    latencySensitivity?: pulumi.Input<string>;
    /**
     * The size of the virtual machine's memory, in MB.
     */
    memory?: pulumi.Input<number>;
    memoryHotAddEnabled?: pulumi.Input<boolean>;
    memoryLimit?: pulumi.Input<number>;
    memoryReservation?: pulumi.Input<number>;
    memoryShareCount?: pulumi.Input<number>;
    memoryShareLevel?: pulumi.Input<string>;
    moid?: pulumi.Input<string>;
    /**
     * The name of the virtual machine. This can be a name or
     * the full path relative to the datacenter. This is required if a UUID lookup
     * is not performed.
     */
    name?: pulumi.Input<string>;
    nestedHvEnabled?: pulumi.Input<boolean>;
    /**
     * The number of cores per socket for this virtual machine.
     */
    numCoresPerSocket?: pulumi.Input<number>;
    /**
     * The total number of virtual processor cores assigned to this
     * virtual machine.
     */
    numCpus?: pulumi.Input<number>;
    replaceTrigger?: pulumi.Input<string>;
    runToolsScriptsAfterPowerOn?: pulumi.Input<boolean>;
    runToolsScriptsAfterResume?: pulumi.Input<boolean>;
    runToolsScriptsBeforeGuestReboot?: pulumi.Input<boolean>;
    runToolsScriptsBeforeGuestShutdown?: pulumi.Input<boolean>;
    runToolsScriptsBeforeGuestStandby?: pulumi.Input<boolean>;
    sataControllerScanCount?: pulumi.Input<number>;
    /**
     * The number of SCSI controllers to
     * scan for disk attributes and controller types on. Default: `1`.
     *
     * > **NOTE:** For best results, ensure that all the disks on any templates you
     * use with this data source reside on the primary controller, and leave this
     * value at the default. See the `vsphere.VirtualMachine`
     * resource documentation for the significance of this setting, specifically the
     * additional requirements and notes for cloning section.
     */
    scsiControllerScanCount?: pulumi.Input<number>;
    storagePolicyId?: pulumi.Input<string>;
    swapPlacementPolicy?: pulumi.Input<string>;
    syncTimeWithHost?: pulumi.Input<boolean>;
    syncTimeWithHostPeriodically?: pulumi.Input<boolean>;
    toolsUpgradePolicy?: pulumi.Input<string>;
    /**
     * Specify this field for a UUID lookup, `name` and `datacenterId`
     * are not required if this is specified.
     */
    uuid?: pulumi.Input<string>;
    vapp?: pulumi.Input<inputs.GetVirtualMachineVappArgs>;
    vbsEnabled?: pulumi.Input<boolean>;
    vvtdEnabled?: pulumi.Input<boolean>;
}
