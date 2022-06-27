// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface ComputeClusterVsanDiskGroup {
    cache?: pulumi.Input<string>;
    storages?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ContentLibraryPublication {
    /**
     * Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     */
    authenticationMethod?: pulumi.Input<string>;
    /**
     * Password used for authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * The URL of the published content library.
     */
    publishUrl?: pulumi.Input<string>;
    /**
     * Publish the content library. Default `false`.
     */
    published?: pulumi.Input<boolean>;
    /**
     * Username used for authentication.
     */
    username?: pulumi.Input<string>;
}

export interface ContentLibrarySubscription {
    /**
     * Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     */
    authenticationMethod?: pulumi.Input<string>;
    /**
     * Enable automatic synchronization with the published library. Default `false`.
     */
    automaticSync?: pulumi.Input<boolean>;
    /**
     * Download the library from a content only when needed. Default `true`.
     */
    onDemand?: pulumi.Input<boolean>;
    /**
     * Password used for authentication.
     */
    password?: pulumi.Input<string>;
    /**
     * URL of the published content library.
     */
    subscriptionUrl?: pulumi.Input<string>;
    /**
     * Username used for authentication.
     */
    username?: pulumi.Input<string>;
}

export interface DistributedPortGroupVlanRange {
    maxVlan: pulumi.Input<number>;
    minVlan: pulumi.Input<number>;
}

export interface DistributedVirtualSwitchHost {
    /**
     * The list of NIC devices to map to uplinks on the VDS,
     * added in order they are specified.
     */
    devices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The host system ID of the host to add to the
     * VDS.
     */
    hostSystemId: pulumi.Input<string>;
}

export interface DistributedVirtualSwitchPvlanMapping {
    /**
     * The primary VLAN ID. The VLAN IDs of 0 and
     * 4095 are reserved and cannot be used in this property.
     */
    primaryVlanId: pulumi.Input<number>;
    /**
     * The private VLAN type. Valid values are
     * promiscuous, community and isolated.
     */
    pvlanType: pulumi.Input<string>;
    /**
     * The secondary VLAN ID. The VLAN IDs of 0
     * and 4095 are reserved and cannot be used in this property.
     */
    secondaryVlanId: pulumi.Input<number>;
}

export interface DistributedVirtualSwitchVlanRange {
    maxVlan: pulumi.Input<number>;
    minVlan: pulumi.Input<number>;
}

export interface EntityPermissionsPermission {
    /**
     * Whether userOrGroup field refers to a user or a group. True for a group and false for a user.
     */
    isGroup: pulumi.Input<boolean>;
    /**
     * Whether or not this permission propagates down the hierarchy to sub-entities.
     */
    propagate: pulumi.Input<boolean>;
    /**
     * The role id of the role to be given to the user on the specified entity.
     */
    roleId: pulumi.Input<string>;
    /**
     * The user/group getting the permission.
     */
    userOrGroup: pulumi.Input<string>;
}

export interface GetVirtualMachineVappArgs {
    properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface GetVirtualMachineVapp {
    properties?: {[key: string]: string};
}

export interface HostPortGroupPort {
    /**
     * The key for this port group as returned from the vSphere API.
     */
    key?: pulumi.Input<string>;
    macAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
}

export interface VirtualMachineCdrom {
    /**
     * Indicates whether the device should be backed by remote client device. Conflicts with `datastoreId` and `path`.
     */
    clientDevice?: pulumi.Input<boolean>;
    /**
     * The datastore ID that on which the ISO is located. Required for using a datastore ISO. Conflicts with `clientDevice`.
     */
    datastoreId?: pulumi.Input<string>;
    deviceAddress?: pulumi.Input<string>;
    /**
     * The ID of the device within the virtual machine.
     */
    key?: pulumi.Input<number>;
    /**
     * The path to the ISO file. Required for using a datastore ISO. Conflicts with `clientDevice`.
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
     * A specification for a virtual NIC on the virtual machine. See network interface options for more information.
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
    scriptText?: pulumi.Input<string>;
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
     * Attach an external disk instead of creating a new one. Implies and conflicts with `keepOnRemove`. If set, you cannot set `size`, `eagerlyScrub`, or `thinProvisioned`. Must set `path` if used.
     */
    attach?: pulumi.Input<boolean>;
    /**
     * The type of storage controller to attach the  disk to. Can be `scsi`, `sata`, or `ide`. You must have the appropriate number of controllers enabled for the selected type. Default `scsi`.
     */
    controllerType?: pulumi.Input<string>;
    /**
     * The datastore ID that on which the ISO is located. Required for using a datastore ISO. Conflicts with `clientDevice`.
     */
    datastoreId?: pulumi.Input<string>;
    deviceAddress?: pulumi.Input<string>;
    /**
     * The mode of this this virtual disk for purposes of writes and snapshots. One of `append`, `independentNonpersistent`, `independentPersistent`, `nonpersistent`, `persistent`, or `undoable`. Default: `persistent`. For more information on these option, please refer to the [product documentation][vmware-docs-disk-mode].
     */
    diskMode?: pulumi.Input<string>;
    /**
     * The sharing mode of this virtual disk. One of `sharingMultiWriter` or `sharingNone`. Default: `sharingNone`.
     */
    diskSharing?: pulumi.Input<string>;
    /**
     * If set to `true`, the disk space is zeroed out when the virtual machine is created. This will delay the creation of the virtual disk. Cannot be set to `true` when `thinProvisioned` is `true`.  See the section on picking a disk type for more information.  Default: `false`.
     */
    eagerlyScrub?: pulumi.Input<boolean>;
    /**
     * The upper limit of IOPS that this disk can use. The default is no limit.
     */
    ioLimit?: pulumi.Input<number>;
    /**
     * The I/O reservation (guarantee) for the virtual disk has, in IOPS.  The default is no reservation.
     */
    ioReservation?: pulumi.Input<number>;
    /**
     * The share count for the virtual disk when the share level is `custom`.
     */
    ioShareCount?: pulumi.Input<number>;
    /**
     * The share allocation level for the virtual disk. One of `low`, `normal`, `high`, or `custom`. Default: `normal`.
     */
    ioShareLevel?: pulumi.Input<string>;
    /**
     * Keep this disk when removing the device or destroying the virtual machine. Default: `false`.
     */
    keepOnRemove?: pulumi.Input<boolean>;
    /**
     * The ID of the device within the virtual machine.
     */
    key?: pulumi.Input<number>;
    /**
     * A label for the virtual disk. Forces a new disk, if changed.
     */
    label: pulumi.Input<string>;
    /**
     * The path to the ISO file. Required for using a datastore ISO. Conflicts with `clientDevice`.
     */
    path?: pulumi.Input<string>;
    /**
     * The size of the disk, in GB. Must be a whole number.
     */
    size?: pulumi.Input<number>;
    /**
     * The UUID of the storage policy to assign to the virtual disk.
     */
    storagePolicyId?: pulumi.Input<string>;
    /**
     * If `true`, the disk is thin provisioned, with space for the file being allocated on an as-needed basis. Cannot be set to `true` when `eagerlyScrub` is `true`. See the section on selecting a disk type for more information. Default: `true`.
     */
    thinProvisioned?: pulumi.Input<boolean>;
    /**
     * The disk number on the storage bus. The maximum value for this setting is the value of the controller count times the controller capacity (15 for SCSI, 30 for SATA, and 2 for IDE). Duplicate unit numbers are not allowed. Default `0`, for which one disk must be set to.
     */
    unitNumber?: pulumi.Input<number>;
    /**
     * The UUID of the virtual disk VMDK file. This is used to track the virtual disk on the virtual machine.
     */
    uuid?: pulumi.Input<string>;
    /**
     * If `true`, writes for this disk are sent directly to the filesystem immediately instead of being buffered. Default: `false`.
     */
    writeThrough?: pulumi.Input<boolean>;
}

export interface VirtualMachineNetworkInterface {
    /**
     * The network interface type. One of `e1000`, `e1000e`, or `vmxnet3`. Default: `vmxnet3`.
     */
    adapterType?: pulumi.Input<string>;
    /**
     * The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit.
     */
    bandwidthLimit?: pulumi.Input<number>;
    /**
     * The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
     */
    bandwidthReservation?: pulumi.Input<number>;
    /**
     * The share count for the network interface when the share level is `custom`.
     */
    bandwidthShareCount?: pulumi.Input<number>;
    /**
     * The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`.
     */
    bandwidthShareLevel?: pulumi.Input<string>;
    deviceAddress?: pulumi.Input<string>;
    /**
     * The ID of the device within the virtual machine.
     */
    key?: pulumi.Input<number>;
    /**
     * The MAC address of the network interface. Can only be manually set if `useStaticMac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
     */
    macAddress?: pulumi.Input<string>;
    /**
     * The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
     */
    networkId: pulumi.Input<string>;
    /**
     * Specifies which NIC in an OVF/OVA the `networkInterface` should be associated. Only applies at creation when deploying from an OVF/OVA.
     */
    ovfMapping?: pulumi.Input<string>;
    /**
     * If true, the `macAddress` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `macAddress` to be set. Default: `false`.
     */
    useStaticMac?: pulumi.Input<boolean>;
}

export interface VirtualMachineOvfDeploy {
    allowUnverifiedSslCert?: pulumi.Input<boolean>;
    deploymentOption?: pulumi.Input<string>;
    diskProvisioning?: pulumi.Input<string>;
    enableHiddenProperties?: pulumi.Input<boolean>;
    ipAllocationPolicy?: pulumi.Input<string>;
    ipProtocol?: pulumi.Input<string>;
    localOvfPath?: pulumi.Input<string>;
    ovfNetworkMap?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    remoteOvfUrl?: pulumi.Input<string>;
}

export interface VirtualMachineVapp {
    properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

export interface VmStoragePolicyTagRule {
    /**
     * Include datastores with the given tags or exclude. Default `true`.
     */
    includeDatastoresWithTags?: pulumi.Input<boolean>;
    /**
     * Name of the tag category.
     */
    tagCategory: pulumi.Input<string>;
    /**
     * List of Name of tags to select from the given category.
     */
    tags: pulumi.Input<pulumi.Input<string>[]>;
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
