// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";

export interface DistributedPortGroupVlanRange {
    maxVlan: pulumi.Input<number>;
    minVlan: pulumi.Input<number>;
}

export interface DistributedVirtualSwitchHost {
    /**
     * The list of NIC devices to map to uplinks on the DVS,
     * added in order they are specified.
     */
    devices: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The host system ID of the host to add to the
     * DVS.
     */
    hostSystemId: pulumi.Input<string>;
}

export interface DistributedVirtualSwitchVlanRange {
    maxVlan: pulumi.Input<number>;
    minVlan: pulumi.Input<number>;
}

export interface HostPortGroupPorts {
    /**
     * The key for this port group as returned from the vSphere API.
     */
    key?: pulumi.Input<string>;
    macAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
}

export interface VirtualMachineCdrom {
    /**
     * Indicates whether the device should be backed by
     * remote client device. Conflicts with `datastoreId` and `path`.
     */
    clientDevice?: pulumi.Input<boolean>;
    /**
     * The datastore ID that the ISO is located in.
     * Requried for using a datastore ISO. Conflicts with `clientDevice`.
     */
    datastoreId?: pulumi.Input<string>;
    /**
     * An address internal to this provider that helps locate the
     * device when `key` is unavailable. This follows a convention of
     * `CONTROLLER_TYPE:BUS_NUMBER:UNIT_NUMBER`. Example: `scsi:0:1` means device
     * unit 1 on SCSI bus 0.
     */
    deviceAddress?: pulumi.Input<string>;
    /**
     * The ID of the device within the virtual machine.
     */
    key?: pulumi.Input<number>;
    /**
     * The path to the ISO file. Required for using a datastore
     * ISO. Conflicts with `clientDevice`.
     */
    path?: pulumi.Input<string>;
}

export interface VirtualMachineClone {
    customize?: pulumi.Input<inputs.VirtualMachineCloneCustomize>;
    linkedClone?: pulumi.Input<boolean>;
    ovfNetworkMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    ovfStorageMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    templateUuid: pulumi.Input<string>;
    timeout?: pulumi.Input<number>;
}

export interface VirtualMachineCloneCustomize {
    dnsServerLists?: pulumi.Input<pulumi.Input<string>[]>;
    dnsSuffixLists?: pulumi.Input<pulumi.Input<string>[]>;
    ipv4Gateway?: pulumi.Input<string>;
    ipv6Gateway?: pulumi.Input<string>;
    linuxOptions?: pulumi.Input<inputs.VirtualMachineCloneCustomizeLinuxOptions>;
    /**
     * A specification for a virtual NIC on this
     * virtual machine. See network interface options
     * below.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.VirtualMachineCloneCustomizeNetworkInterface>[]>;
    timeout?: pulumi.Input<number>;
    windowsOptions?: pulumi.Input<inputs.VirtualMachineCloneCustomizeWindowsOptions>;
    windowsSysprepText?: pulumi.Input<string>;
}

export interface VirtualMachineCloneCustomizeLinuxOptions {
    domain: pulumi.Input<string>;
    hostName: pulumi.Input<string>;
    hwClockUtc?: pulumi.Input<boolean>;
    timeZone?: pulumi.Input<string>;
}

export interface VirtualMachineCloneCustomizeNetworkInterface {
    dnsDomain?: pulumi.Input<string>;
    dnsServerLists?: pulumi.Input<pulumi.Input<string>[]>;
    ipv4Address?: pulumi.Input<string>;
    ipv4Netmask?: pulumi.Input<number>;
    ipv6Address?: pulumi.Input<string>;
    ipv6Netmask?: pulumi.Input<number>;
}

export interface VirtualMachineCloneCustomizeWindowsOptions {
    adminPassword?: pulumi.Input<string>;
    autoLogon?: pulumi.Input<boolean>;
    autoLogonCount?: pulumi.Input<number>;
    computerName: pulumi.Input<string>;
    domainAdminPassword?: pulumi.Input<string>;
    domainAdminUser?: pulumi.Input<string>;
    fullName?: pulumi.Input<string>;
    joinDomain?: pulumi.Input<string>;
    organizationName?: pulumi.Input<string>;
    productKey?: pulumi.Input<string>;
    runOnceCommandLists?: pulumi.Input<pulumi.Input<string>[]>;
    timeZone?: pulumi.Input<number>;
    workgroup?: pulumi.Input<string>;
}

export interface VirtualMachineDisk {
    /**
     * Attach an external disk instead of creating a new one.
     * Implies and conflicts with `keepOnRemove`. If set, you cannot set `size`,
     * `eagerlyScrub`, or `thinProvisioned`. Must set `path` if used.
     */
    attach?: pulumi.Input<boolean>;
    /**
     * The datastore ID that the ISO is located in.
     * Requried for using a datastore ISO. Conflicts with `clientDevice`.
     */
    datastoreId?: pulumi.Input<string>;
    /**
     * An address internal to this provider that helps locate the
     * device when `key` is unavailable. This follows a convention of
     * `CONTROLLER_TYPE:BUS_NUMBER:UNIT_NUMBER`. Example: `scsi:0:1` means device
     * unit 1 on SCSI bus 0.
     */
    deviceAddress?: pulumi.Input<string>;
    /**
     * The mode of this this virtual disk for purposes of
     * writes and snapshotting. Can be one of `append`, `independentNonpersistent`,
     * `independentPersistent`, `nonpersistent`, `persistent`, or `undoable`.
     * Default: `persistent`. For an explanation of options, click
     * [here][vmware-docs-disk-mode].
     */
    diskMode?: pulumi.Input<string>;
    /**
     * The sharing mode of this virtual disk. Can be one
     * of `sharingMultiWriter` or `sharingNone`. Default: `sharingNone`.
     */
    diskSharing?: pulumi.Input<string>;
    /**
     * If set to `true`, the disk space is zeroed out
     * on VM creation. This will delay the creation of the disk or virtual machine.
     * Cannot be set to `true` when `thinProvisioned` is `true`.  See the section
     * on picking a disk type.  Default: `false`.
     */
    eagerlyScrub?: pulumi.Input<boolean>;
    /**
     * The upper limit of IOPS that this disk can use. The
     * default is no limit.
     */
    ioLimit?: pulumi.Input<number>;
    /**
     * The I/O reservation (guarantee) that this disk
     * has, in IOPS.  The default is no reservation.
     */
    ioReservation?: pulumi.Input<number>;
    /**
     * The share count for this disk when the share
     * level is `custom`.
     */
    ioShareCount?: pulumi.Input<number>;
    /**
     * The share allocation level for this disk. Can
     * be one of `low`, `normal`, `high`, or `custom`. Default: `normal`.
     */
    ioShareLevel?: pulumi.Input<string>;
    /**
     * Keep this disk when removing the device or
     * destroying the virtual machine. Default: `false`.
     */
    keepOnRemove?: pulumi.Input<boolean>;
    /**
     * The ID of the device within the virtual machine.
     */
    key?: pulumi.Input<number>;
    /**
     * A label for the disk. Forces a new disk if changed.
     */
    label?: pulumi.Input<string>;
    /**
     * An alias for both `label` and `path`, the latter when
     * using `attach`. Required if not using `label`.
     */
    name?: pulumi.Input<string>;
    /**
     * The path to the ISO file. Required for using a datastore
     * ISO. Conflicts with `clientDevice`.
     */
    path?: pulumi.Input<string>;
    /**
     * The size of the disk, in GB.
     */
    size?: pulumi.Input<number>;
    /**
     * The UUID of the storage policy to assign to this disk.
     */
    storagePolicyId?: pulumi.Input<string>;
    /**
     * If `true`, this disk is thin provisioned,
     * with space for the file being allocated on an as-needed basis. Cannot be set
     * to `true` when `eagerlyScrub` is `true`. See the section on picking a disk
     * type. Default: `true`.
     */
    thinProvisioned?: pulumi.Input<boolean>;
    /**
     * The disk number on the SCSI bus. The maximum value
     * for this setting is the value of
     * `scsiControllerCount` times 15, minus 1 (so `14`,
     * `29`, `44`, and `59`, for 1-4 controllers respectively). The default is `0`,
     * for which one disk must be set to. Duplicate unit numbers are not allowed.
     */
    unitNumber?: pulumi.Input<number>;
    /**
     * The UUID of the virtual disk's VMDK file. This is used to track the
     * virtual disk on the virtual machine.
     */
    uuid?: pulumi.Input<string>;
    /**
     * If `true`, writes for this disk are sent
     * directly to the filesystem immediately instead of being buffered. Default:
     * `false`.
     */
    writeThrough?: pulumi.Input<boolean>;
}

export interface VirtualMachineNetworkInterface {
    /**
     * The network interface type. Can be one of
     * `e1000`, `e1000e`, or `vmxnet3`. Default: `vmxnet3`.
     */
    adapterType?: pulumi.Input<string>;
    /**
     * The upper bandwidth limit of this network
     * interface, in Mbits/sec. The default is no limit.
     */
    bandwidthLimit?: pulumi.Input<number>;
    /**
     * The bandwidth reservation of this
     * network interface, in Mbits/sec. The default is no reservation.
     */
    bandwidthReservation?: pulumi.Input<number>;
    /**
     * The share count for this network
     * interface when the share level is `custom`.
     */
    bandwidthShareCount?: pulumi.Input<number>;
    /**
     * The bandwidth share allocation level for
     * this interface. Can be one of `low`, `normal`, `high`, or `custom`. Default:
     * `normal`.
     */
    bandwidthShareLevel?: pulumi.Input<string>;
    /**
     * An address internal to this provider that helps locate the
     * device when `key` is unavailable. This follows a convention of
     * `CONTROLLER_TYPE:BUS_NUMBER:UNIT_NUMBER`. Example: `scsi:0:1` means device
     * unit 1 on SCSI bus 0.
     */
    deviceAddress?: pulumi.Input<string>;
    /**
     * The ID of the device within the virtual machine.
     */
    key?: pulumi.Input<number>;
    /**
     * The MAC address of this network interface. Can
     * only be manually set if `useStaticMac` is true, otherwise this is a
     * computed value that gives the current MAC address of this interface.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The managed object reference
     * ID of the network to connect this interface to.
     */
    networkId: pulumi.Input<string>;
    /**
     * Specifies which OVF NIC the `networkInterface`
     * should be associated with. Only applies at creation and only when deploying
     * from an OVF source.
     */
    ovfMapping?: pulumi.Input<string>;
    /**
     * If true, the `macAddress` field is treated as
     * a static MAC address and set accordingly. Setting this to `true` requires
     * `macAddress` to be set. Default: `false`.
     */
    useStaticMac?: pulumi.Input<boolean>;
}

export interface VirtualMachineOvfDeploy {
    diskProvisioning?: pulumi.Input<string>;
    ipAllocationPolicy?: pulumi.Input<string>;
    ipProtocol?: pulumi.Input<string>;
    localOvfPath?: pulumi.Input<string>;
    ovfNetworkMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    remoteOvfUrl?: pulumi.Input<string>;
}

export interface VirtualMachineVapp {
    properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface VnicIpv4 {
    /**
     * Use DHCP to configure the interface's IPv4 stack.
     */
    dhcp?: pulumi.Input<boolean>;
    /**
     * IP address of the default gateway, if DHCP or autoconfig is not set.
     */
    gw?: pulumi.Input<string>;
    /**
     * Address of the interface, if DHCP is not set.
     */
    ip?: pulumi.Input<string>;
    /**
     * Netmask of the interface, if DHCP is not set.
     */
    netmask?: pulumi.Input<string>;
}

export interface VnicIpv6 {
    /**
     * List of IPv6 addresses
     */
    addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Use IPv6 Autoconfiguration (RFC2462).
     */
    autoconfig?: pulumi.Input<boolean>;
    /**
     * Use DHCP to configure the interface's IPv4 stack.
     */
    dhcp?: pulumi.Input<boolean>;
    /**
     * IP address of the default gateway, if DHCP or autoconfig is not set.
     */
    gw?: pulumi.Input<string>;
}
