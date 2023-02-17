// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ComputeClusterVsanDiskGroup {
    /**
     * The canonical name of the disk to use for vSAN cache.
     */
    cache?: string;
    /**
     * An array of disk canonical names for vSAN storage.
     */
    storages?: string[];
}

export interface ContentLibraryPublication {
    /**
     * Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     */
    authenticationMethod?: string;
    /**
     * Password used for authentication.
     */
    password: string;
    /**
     * The URL of the published content library.
     */
    publishUrl: string;
    /**
     * Publish the content library. Default `false`.
     */
    published?: boolean;
    /**
     * Username used for authentication.
     */
    username: string;
}

export interface ContentLibrarySubscription {
    /**
     * Authentication method to connect ro a published content library. Must be `NONE` or `BASIC`.
     */
    authenticationMethod?: string;
    /**
     * Enable automatic synchronization with the published library. Default `false`.
     */
    automaticSync?: boolean;
    /**
     * Download the library from a content only when needed. Default `true`.
     */
    onDemand?: boolean;
    /**
     * Password used for authentication.
     */
    password: string;
    /**
     * URL of the published content library.
     */
    subscriptionUrl?: string;
    /**
     * Username used for authentication.
     */
    username: string;
}

export interface DistributedPortGroupVlanRange {
    maxVlan: number;
    minVlan: number;
}

export interface DistributedVirtualSwitchHost {
    /**
     * The list of NIC devices to map to uplinks on the VDS,
     * added in order they are specified.
     */
    devices?: string[];
    /**
     * The host system ID of the host to add to the
     * VDS.
     */
    hostSystemId: string;
}

export interface DistributedVirtualSwitchPvlanMapping {
    /**
     * The primary VLAN ID. The VLAN IDs of 0 and
     * 4095 are reserved and cannot be used in this property.
     */
    primaryVlanId: number;
    /**
     * The private VLAN type. Valid values are
     * promiscuous, community and isolated.
     */
    pvlanType: string;
    /**
     * The secondary VLAN ID. The VLAN IDs of 0
     * and 4095 are reserved and cannot be used in this property.
     */
    secondaryVlanId: number;
}

export interface DistributedVirtualSwitchVlanRange {
    maxVlan: number;
    minVlan: number;
}

export interface EntityPermissionsPermission {
    /**
     * Whether userOrGroup field refers to a user or a group. True for a group and false for a user.
     */
    isGroup: boolean;
    /**
     * Whether or not this permission propagates down the hierarchy to sub-entities.
     */
    propagate: boolean;
    /**
     * The role id of the role to be given to the user on the specified entity.
     */
    roleId: string;
    /**
     * The user/group getting the permission.
     */
    userOrGroup: string;
}

export interface GetVirtualMachineDisk {
    /**
     * Set to `true` if the disk has been eager zeroed.
     */
    eagerlyScrub: boolean;
    /**
     * The label for the disk.
     */
    label: string;
    /**
     * The size of the disk, in GIB.
     */
    size: number;
    /**
     * Set to `true` if the disk has been thin provisioned.
     */
    thinProvisioned: boolean;
    /**
     * The disk number on the storage bus.
     */
    unitNumber: number;
}

export interface GetVirtualMachineNetworkInterface {
    /**
     * The network interface types for each network interface found 
     * on the virtual machine, in device bus order. Will be one of `e1000`, `e1000e`,
     * `vmxnet3vrdma`, or `vmxnet3`.
     */
    adapterType: string;
    /**
     * The upper bandwidth limit of this network interface, 
     * in Mbits/sec.
     */
    bandwidthLimit?: number;
    /**
     * The bandwidth reservation of this network interface,
     * in Mbits/sec.
     */
    bandwidthReservation?: number;
    /**
     * The share count for this network interface when the
     * share level is custom.
     */
    bandwidthShareCount: number;
    /**
     * The bandwidth share allocation level for this interface.
     * Can be one of `low`, `normal`, `high`, or `custom`.
     */
    bandwidthShareLevel?: string;
    /**
     * The MAC address of this network interface.
     */
    macAddress: string;
    /**
     * The managed object reference ID of the network this interface is
     * connected to.
     */
    networkId: string;
}

export interface GetVirtualMachineVapp {
    properties?: {[key: string]: string};
}

export interface HostPortGroupPort {
    /**
     * The key for this port group as returned from the vSphere API.
     */
    key: string;
    macAddresses: string[];
    type: string;
}

export interface VirtualMachineCdrom {
    /**
     * Indicates whether the device should be backed by remote client device. Conflicts with `datastoreId` and `path`.
     */
    clientDevice?: boolean;
    /**
     * The managed object reference ID of the datastore in which to place the virtual machine. The virtual machine configuration files is placed here, along with any virtual disks that are created where a datastore is not explicitly specified. See the section on virtual machine migration for more information on modifying this value.
     */
    datastoreId?: string;
    deviceAddress: string;
    /**
     * The ID of the device within the virtual machine.
     */
    key: number;
    /**
     * When using `attach`, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk filename.
     */
    path?: string;
}

export interface VirtualMachineClone {
    customize?: outputs.VirtualMachineCloneCustomize;
    linkedClone?: boolean;
    ovfNetworkMap?: {[key: string]: string};
    ovfStorageMap?: {[key: string]: string};
    templateUuid: string;
    timeout?: number;
}

export interface VirtualMachineCloneCustomize {
    dnsServerLists?: string[];
    dnsSuffixLists?: string[];
    ipv4Gateway?: string;
    ipv6Gateway?: string;
    linuxOptions?: outputs.VirtualMachineCloneCustomizeLinuxOptions;
    /**
     * A specification for a virtual NIC on the virtual machine. See network interface options for more information.
     */
    networkInterfaces?: outputs.VirtualMachineCloneCustomizeNetworkInterface[];
    timeout?: number;
    windowsOptions?: outputs.VirtualMachineCloneCustomizeWindowsOptions;
    windowsSysprepText?: string;
}

export interface VirtualMachineCloneCustomizeLinuxOptions {
    domain: string;
    hostName: string;
    hwClockUtc?: boolean;
    scriptText?: string;
    timeZone?: string;
}

export interface VirtualMachineCloneCustomizeNetworkInterface {
    dnsDomain?: string;
    dnsServerLists?: string[];
    ipv4Address?: string;
    ipv4Netmask?: number;
    ipv6Address?: string;
    ipv6Netmask?: number;
}

export interface VirtualMachineCloneCustomizeWindowsOptions {
    adminPassword?: string;
    autoLogon?: boolean;
    autoLogonCount?: number;
    computerName: string;
    domainAdminPassword?: string;
    domainAdminUser?: string;
    fullName?: string;
    joinDomain?: string;
    organizationName?: string;
    productKey?: string;
    runOnceCommandLists?: string[];
    timeZone?: number;
    workgroup?: string;
}

export interface VirtualMachineDisk {
    /**
     * Attach an external disk instead of creating a new one. Implies and conflicts with `keepOnRemove`. If set, you cannot set `size`, `eagerlyScrub`, or `thinProvisioned`. Must set `path` if used.
     */
    attach?: boolean;
    /**
     * The type of storage controller to attach the  disk to. Can be `scsi`, `sata`, or `ide`. You must have the appropriate number of controllers enabled for the selected type. Default `scsi`.
     */
    controllerType?: string;
    /**
     * The managed object reference ID of the datastore in which to place the virtual machine. The virtual machine configuration files is placed here, along with any virtual disks that are created where a datastore is not explicitly specified. See the section on virtual machine migration for more information on modifying this value.
     */
    datastoreId: string;
    deviceAddress: string;
    /**
     * The mode of this this virtual disk for purposes of writes and snapshots. One of `append`, `independentNonpersistent`, `independentPersistent`, `nonpersistent`, `persistent`, or `undoable`. Default: `persistent`. For more information on these option, please refer to the [product documentation][vmware-docs-disk-mode].
     */
    diskMode?: string;
    /**
     * The sharing mode of this virtual disk. One of `sharingMultiWriter` or `sharingNone`. Default: `sharingNone`.
     */
    diskSharing?: string;
    /**
     * If set to `true`, the disk space is zeroed out when the virtual machine is created. This will delay the creation of the virtual disk. Cannot be set to `true` when `thinProvisioned` is `true`.  See the section on picking a disk type for more information.  Default: `false`.
     */
    eagerlyScrub?: boolean;
    /**
     * The upper limit of IOPS that this disk can use. The default is no limit.
     */
    ioLimit?: number;
    /**
     * The I/O reservation (guarantee) for the virtual disk has, in IOPS.  The default is no reservation.
     */
    ioReservation?: number;
    /**
     * The share count for the virtual disk when the share level is `custom`.
     */
    ioShareCount?: number;
    /**
     * The share allocation level for the virtual disk. One of `low`, `normal`, `high`, or `custom`. Default: `normal`.
     */
    ioShareLevel?: string;
    /**
     * Keep this disk when removing the device or destroying the virtual machine. Default: `false`.
     */
    keepOnRemove?: boolean;
    /**
     * The ID of the device within the virtual machine.
     */
    key: number;
    /**
     * A label for the virtual disk. Forces a new disk, if changed.
     */
    label: string;
    /**
     * When using `attach`, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk filename.
     */
    path: string;
    /**
     * The size of the disk, in GB. Must be a whole number.
     */
    size?: number;
    /**
     * The ID of the storage policy to assign to the home directory of a virtual machine.
     */
    storagePolicyId: string;
    /**
     * If `true`, the disk is thin provisioned, with space for the file being allocated on an as-needed basis. Cannot be set to `true` when `eagerlyScrub` is `true`. See the section on selecting a disk type for more information. Default: `true`.
     */
    thinProvisioned?: boolean;
    /**
     * The disk number on the storage bus. The maximum value for this setting is the value of the controller count times the controller capacity (15 for SCSI, 30 for SATA, and 2 for IDE). Duplicate unit numbers are not allowed. Default `0`, for which one disk must be set to.
     */
    unitNumber?: number;
    /**
     * The UUID of the virtual disk VMDK file. This is used to track the virtual disk on the virtual machine.
     */
    uuid: string;
    /**
     * If `true`, writes for this disk are sent directly to the filesystem immediately instead of being buffered. Default: `false`.
     */
    writeThrough?: boolean;
}

export interface VirtualMachineNetworkInterface {
    /**
     * The network interface type. One of `e1000`, `e1000e`, or `vmxnet3`. Default: `vmxnet3`.
     */
    adapterType?: string;
    /**
     * The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit.
     */
    bandwidthLimit?: number;
    /**
     * The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.
     */
    bandwidthReservation?: number;
    /**
     * The share count for the network interface when the share level is `custom`.
     */
    bandwidthShareCount: number;
    /**
     * The bandwidth share allocation level for the network interface. One of `low`, `normal`, `high`, or `custom`. Default: `normal`.
     */
    bandwidthShareLevel?: string;
    deviceAddress: string;
    /**
     * The ID of the device within the virtual machine.
     */
    key: number;
    /**
     * The MAC address of the network interface. Can only be manually set if `useStaticMac` is `true`. Otherwise, the value is computed and presents the assigned MAC address for the interface.
     */
    macAddress: string;
    /**
     * The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.
     */
    networkId: string;
    /**
     * Specifies which NIC in an OVF/OVA the `networkInterface` should be associated. Only applies at creation when deploying from an OVF/OVA.
     */
    ovfMapping?: string;
    /**
     * If true, the `macAddress` field is treated as a static MAC address and set accordingly. Setting this to `true` requires `macAddress` to be set. Default: `false`.
     */
    useStaticMac?: boolean;
}

export interface VirtualMachineOvfDeploy {
    allowUnverifiedSslCert?: boolean;
    deploymentOption?: string;
    diskProvisioning?: string;
    enableHiddenProperties?: boolean;
    ipAllocationPolicy?: string;
    ipProtocol?: string;
    localOvfPath?: string;
    ovfNetworkMap?: {[key: string]: string};
    remoteOvfUrl?: string;
}

export interface VirtualMachineVapp {
    properties?: {[key: string]: string};
}

export interface VmStoragePolicyTagRule {
    /**
     * Include datastores with the given tags or exclude. Default `true`.
     */
    includeDatastoresWithTags?: boolean;
    /**
     * Name of the tag category.
     */
    tagCategory: string;
    /**
     * List of Name of tags to select from the given category.
     */
    tags: string[];
}

export interface VnicIpv4 {
    /**
     * Use DHCP to configure the interface's IPv4 stack.
     */
    dhcp?: boolean;
    /**
     * IP address of the default gateway, if DHCP is not set.
     */
    gw?: string;
    /**
     * Address of the interface, if DHCP is not set.
     */
    ip?: string;
    /**
     * Netmask of the interface, if DHCP is not set.
     */
    netmask?: string;
}

export interface VnicIpv6 {
    /**
     * List of IPv6 addresses
     */
    addresses?: string[];
    /**
     * Use IPv6 Autoconfiguration (RFC2462).
     */
    autoconfig?: boolean;
    /**
     * Use DHCP to configure the interface's IPv4 stack.
     */
    dhcp?: boolean;
    /**
     * IP address of the default gateway, if DHCP is not set.
     */
    gw?: string;
}

