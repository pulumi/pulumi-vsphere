// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ComputeClusterHostImage {
    /**
     * List of custom components.
     */
    components?: outputs.ComputeClusterHostImageComponent[];
    /**
     * The ESXi version which the image is based on.
     */
    esxVersion?: string;
}

export interface ComputeClusterHostImageComponent {
    /**
     * The identifier for the component.
     */
    key?: string;
    /**
     * The version to use.
     */
    version?: string;
}

export interface ComputeClusterVsanDiskGroup {
    /**
     * Cache disk.
     */
    cache?: string;
    /**
     * List of storage disks.
     */
    storages?: string[];
}

export interface ComputeClusterVsanFaultDomain {
    /**
     * The configuration for single fault domain.
     */
    faultDomains?: outputs.ComputeClusterVsanFaultDomainFaultDomain[];
}

export interface ComputeClusterVsanFaultDomainFaultDomain {
    /**
     * The managed object IDs of the hosts to put in the fault domain.
     */
    hostIds: string[];
    /**
     * The name of the cluster.
     */
    name: string;
}

export interface ComputeClusterVsanStretchedCluster {
    /**
     * The managed object IDs of the hosts to put in the first fault domain.
     */
    preferredFaultDomainHostIds: string[];
    /**
     * The name of prepferred fault domain.
     */
    preferredFaultDomainName?: string;
    /**
     * The managed object IDs of the hosts to put in the second fault domain.
     */
    secondaryFaultDomainHostIds: string[];
    /**
     * The name of secondary fault domain.
     */
    secondaryFaultDomainName?: string;
    /**
     * The managed object IDs of the host selected as witness node when enable stretched cluster.
     */
    witnessNode: string;
}

export interface ContentLibraryPublication {
    /**
     * Method to authenticate users. Must be `NONE` or `BASIC`.
     */
    authenticationMethod?: string;
    /**
     * Password used by subscribers to authenticate.
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
     * Username used by subscribers to authenticate. Currently can only be `vcsp`.
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
    /**
     * The minimum VLAN to use in the range.
     */
    maxVlan: number;
    /**
     * The minimum VLAN to use in the range.
     */
    minVlan: number;
}

export interface DistributedVirtualSwitchHost {
    /**
     * Name of the physical NIC to be added to the proxy switch.
     */
    devices?: string[];
    /**
     * The managed object ID of the host this specification applies to.
     */
    hostSystemId: string;
}

export interface DistributedVirtualSwitchPvlanMapping {
    /**
     * The primary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
     */
    primaryVlanId: number;
    /**
     * The private VLAN type. Valid values are promiscuous, community and isolated.
     */
    pvlanType: string;
    /**
     * The secondary VLAN ID. The VLAN IDs of 0 and 4095 are reserved and cannot be used in this property.
     */
    secondaryVlanId: number;
}

export interface DistributedVirtualSwitchVlanRange {
    /**
     * The minimum VLAN to use in the range.
     */
    maxVlan: number;
    /**
     * The minimum VLAN to use in the range.
     */
    minVlan: number;
}

export interface EntityPermissionsPermission {
    /**
     * Whether `userOrGroup` field refers to a user or a
     * group. True for a group and false for a user.
     */
    isGroup: boolean;
    /**
     * Whether or not this permission propagates down the
     * hierarchy to sub-entities.
     */
    propagate: boolean;
    /**
     * The role id of the role to be given to the user on
     * the specified entity.
     */
    roleId: string;
    /**
     * The user/group getting the permission.
     */
    userOrGroup: string;
}

export interface GetGuestOsCustomizationSpec {
    /**
     * A list of DNS servers for a virtual network adapter with a static IP address.
     */
    dnsServerLists: string[];
    /**
     * A list of DNS search domains to add to the DNS configuration on the virtual machine.
     */
    dnsSuffixLists: string[];
    /**
     * A list of configuration options specific to Linux.
     */
    linuxOptions: outputs.GetGuestOsCustomizationSpecLinuxOption[];
    /**
     * A specification of network interface configuration options.
     */
    networkInterfaces: outputs.GetGuestOsCustomizationSpecNetworkInterface[];
    /**
     * A list of configuration options specific to Windows.
     */
    windowsOptions: outputs.GetGuestOsCustomizationSpecWindowsOption[];
    /**
     * Use this option to specify use of a Windows Sysprep file.
     */
    windowsSysprepText: string;
}

export interface GetGuestOsCustomizationSpecLinuxOption {
    /**
     * The domain name for this virtual machine.
     */
    domain: string;
    /**
     * The hostname for this virtual machine.
     */
    hostName: string;
    /**
     * Specifies whether or not the hardware clock should be in UTC or not.
     */
    hwClockUtc: boolean;
    /**
     * The customization script to run before and or after guest customization.
     */
    scriptText: string;
    /**
     * Set the time zone on the guest operating system. For a list of the acceptable values for Linux customization specifications, see [List of Time Zone Database Zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) on Wikipedia.
     */
    timeZone: string;
}

export interface GetGuestOsCustomizationSpecNetworkInterface {
    /**
     * A DNS search domain to add to the DNS configuration on the virtual machine.
     */
    dnsDomain: string;
    /**
     * Network-interface specific DNS settings for Windows operating systems. Ignored on Linux.
     */
    dnsServerLists: string[];
    /**
     * The IPv4 address assigned to this network adapter. If left blank, DHCP is used.
     */
    ipv4Address: string;
    /**
     * The IPv4 CIDR netmask for the supplied IP address. Ignored if DHCP is selected.
     */
    ipv4Netmask: number;
    /**
     * The IPv6 address assigned to this network adapter. If left blank, default auto-configuration is used.
     */
    ipv6Address: string;
    /**
     * The IPv6 CIDR netmask for the supplied IP address. Ignored if auto-configuration is selected.
     */
    ipv6Netmask: number;
}

export interface GetGuestOsCustomizationSpecWindowsOption {
    /**
     * The new administrator password for this virtual machine.
     */
    adminPassword: string;
    /**
     * Specifies whether or not the guest operating system automatically logs on as Administrator.
     */
    autoLogon: boolean;
    /**
     * Specifies how many times the guest operating system should auto-logon the Administrator account when `autoLogon` is `true`.
     */
    autoLogonCount: number;
    /**
     * The hostname for this virtual machine.
     */
    computerName: string;
    /**
     * The user account used to join this virtual machine to the Active Directory domain.
     */
    domainAdminPassword?: string;
    /**
     * The user account of the domain administrator used to join this virtual machine to the domain.
     */
    domainAdminUser: string;
    /**
     * The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
     */
    domainOu: string;
    /**
     * The Active Directory domain for the virtual machine to join.
     */
    joinDomain: string;
    /**
     * A list of commands to run at first user logon, after guest customization.
     */
    runOnceCommandLists: string[];
    /**
     * The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
     */
    timeZone: number;
    /**
     * The workgroup for this virtual machine if not joining an Active Directory domain.
     */
    workgroup: string;
}

export interface GetHostVgpuProfileVgpuProfile {
    /**
     * Indicates whether the GPU plugin on this host is
     * capable of disk-only snapshots when VM is not powered off.
     */
    diskSnapshotSupported: boolean;
    /**
     * Indicates whether the GPU plugin on this host
     * is capable of memory snapshots.
     */
    memorySnapshotSupported: boolean;
    /**
     * Indicates whether the GPU plugin on this host is
     * capable of migration.
     */
    migrateSupported: boolean;
    /**
     * Indicates whether the GPU plugin on this host is
     * capable of suspend-resume.
     */
    suspendSupported: boolean;
    /**
     * Name of a particular vGPU available as a shared GPU device (vGPU
     * profile).
     */
    vgpu: string;
}

export interface GetNetworkFilter {
    /**
     * This is required if you have multiple port groups with the same name. This will be one of `DistributedVirtualPortgroup` for distributed port groups, `Network` for standard (host-based) port groups, or `OpaqueNetwork` for networks managed externally, such as those managed by NSX.
     */
    networkType?: string;
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
     * on the virtual machine, in device bus order. Will be one of `e1000`,
     * `e1000e`, `vmxnet3vrdma`, or `vmxnet3`.
     */
    adapterType: string;
    /**
     * The upper bandwidth limit of this network interface,
     * in Mbits/sec.
     */
    bandwidthLimit?: number;
    /**
     * The bandwidth reservation of this network
     * interface, in Mbits/sec.
     */
    bandwidthReservation?: number;
    /**
     * The share count for this network interface when the
     * share level is custom.
     */
    bandwidthShareCount: number;
    /**
     * The bandwidth share allocation level for this
     * interface. Can be one of `low`, `normal`, `high`, or `custom`.
     */
    bandwidthShareLevel?: string;
    /**
     * The MAC address of this network interface.
     */
    macAddress: string;
    /**
     * The managed object reference ID of the network this interface
     * is connected to.
     */
    networkId: string;
    /**
     * The ID of the Physical SR-IOV NIC to attach to, e.g. '0000:d8:00.0'
     */
    physicalFunction: string;
}

export interface GetVirtualMachineVapp {
    /**
     * A map of customizable vApp properties and their values. Allows customization of VMs cloned from OVF templates which have customizable vApp properties.
     */
    properties?: {[key: string]: string};
}

export interface GuestOsCustomizationSpec {
    /**
     * The list of DNS servers for a virtual network adapter with a static IP address.
     */
    dnsServerLists?: string[];
    /**
     * A list of DNS search domains to add to the DNS configuration on the virtual machine.
     */
    dnsSuffixLists?: string[];
    /**
     * The IPv4 default gateway when using networkInterface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     */
    ipv4Gateway?: string;
    /**
     * The IPv6 default gateway when using networkInterface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     */
    ipv6Gateway?: string;
    /**
     * A list of configuration options specific to Linux virtual machines.
     */
    linuxOptions?: outputs.GuestOsCustomizationSpecLinuxOptions;
    /**
     * A specification of network interface configuration options.
     */
    networkInterfaces?: outputs.GuestOsCustomizationSpecNetworkInterface[];
    /**
     * A list of configuration options specific to Windows virtual machines.
     */
    windowsOptions?: outputs.GuestOsCustomizationSpecWindowsOptions;
    /**
     * Use this option to specify a windows sysprep file directly.
     */
    windowsSysprepText?: string;
}

export interface GuestOsCustomizationSpecLinuxOptions {
    /**
     * The domain name for this virtual machine.
     */
    domain: string;
    /**
     * The hostname for this virtual machine.
     */
    hostName: string;
    /**
     * Specifies whether or not the hardware clock should be in UTC or not.
     */
    hwClockUtc?: boolean;
    /**
     * The customization script to run before and or after guest customization
     */
    scriptText?: string;
    /**
     * Customize the time zone on the VM. This should be a time zone-style entry, like America/Los_Angeles.
     */
    timeZone?: string;
}

export interface GuestOsCustomizationSpecNetworkInterface {
    /**
     * A DNS search domain to add to the DNS configuration on the virtual machine.
     */
    dnsDomain?: string;
    /**
     * Network-interface specific DNS settings for Windows operating systems. Ignored on Linux.
     */
    dnsServerLists?: string[];
    /**
     * The IPv4 address assigned to this network adapter. If left blank, DHCP is used.
     */
    ipv4Address?: string;
    /**
     * The IPv4 CIDR netmask for the supplied IP address. Ignored if DHCP is selected.
     */
    ipv4Netmask?: number;
    /**
     * The IPv6 address assigned to this network adapter. If left blank, default auto-configuration is used.
     */
    ipv6Address?: string;
    /**
     * The IPv6 CIDR netmask for the supplied IP address. Ignored if auto-configuration is selected.
     */
    ipv6Netmask?: number;
}

export interface GuestOsCustomizationSpecWindowsOptions {
    /**
     * The new administrator password for this virtual machine.
     */
    adminPassword?: string;
    /**
     * Specifies whether or not the VM automatically logs on as Administrator.
     */
    autoLogon?: boolean;
    /**
     * Specifies how many times the VM should auto-logon the Administrator account when autoLogon is true.
     */
    autoLogonCount?: number;
    /**
     * The host name for this virtual machine.
     */
    computerName: string;
    /**
     * The password of the domain administrator used to join this virtual machine to the domain.
     */
    domainAdminPassword?: string;
    /**
     * The user account of the domain administrator used to join this virtual machine to the domain.
     */
    domainAdminUser?: string;
    /**
     * The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
     */
    domainOu?: string;
    /**
     * The full name of the user of this virtual machine.
     */
    fullName?: string;
    /**
     * The domain that the virtual machine should join.
     */
    joinDomain?: string;
    /**
     * The organization name this virtual machine is being installed for.
     */
    organizationName?: string;
    /**
     * The product key for this virtual machine.
     */
    productKey?: string;
    /**
     * A list of commands to run at first user logon, after guest customization.
     */
    runOnceCommandLists?: string[];
    /**
     * The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
     */
    timeZone?: number;
    /**
     * The workgroup for this virtual machine if not joining a domain.
     */
    workgroup?: string;
}

export interface HostPortGroupPort {
    /**
     * The key for this port group as returned from the vSphere API.
     */
    key: string;
    /**
     * The MAC addresses of the network service of the virtual machine connected on this port.
     */
    macAddresses: string[];
    /**
     * Type type of the entity connected on this port. Possible values are host (VMKkernel), systemManagement (service console), virtualMachine, or unknown.
     */
    type: string;
}

export interface HostService {
    /**
     * service has three settings, `enabled` sets service to running or not running, `policy` sets service based on setting of `on` which sets service to "Start and stop with host", `off` which sets service to "Start and stop manually", `automatic` which sets service to "Start and stop with port usage".
     *
     * > **NOTE:** `services` only supports ntpd service today.
     */
    ntpd?: outputs.HostServiceNtpd;
}

export interface HostServiceNtpd {
    /**
     * Whether the NTP service is enabled. Default is false.
     */
    enabled?: boolean;
    ntpServers: string[];
    /**
     * The policy for the NTP service. Valid values are 'Start and stop with host', 'Start and stop manually', 'Start and stop with port usage'.
     */
    policy?: string;
}

export interface OfflineSoftwareDepotComponent {
    /**
     * The name of the component. Useful for easier identification.
     */
    displayName: string;
    /**
     * The identifier of the component.
     */
    key: string;
    /**
     * The list of available versions of the component.
     */
    versions: string[];
}

export interface SupervisorEgressCidr {
    /**
     * Network address.
     */
    address: string;
    /**
     * Subnet prefix.
     */
    prefix: number;
}

export interface SupervisorIngressCidr {
    /**
     * Network address.
     */
    address: string;
    /**
     * Subnet prefix.
     */
    prefix: number;
}

export interface SupervisorManagementNetwork {
    /**
     * Number of addresses to allocate. Starts from 'starting_address'
     */
    addressCount: number;
    /**
     * Gateway IP address.
     */
    gateway: string;
    /**
     * ID of the network. (e.g. a distributed port group).
     */
    network: string;
    /**
     * Starting address of the management network range.
     */
    startingAddress: string;
    /**
     * Subnet mask.
     */
    subnetMask: string;
}

export interface SupervisorNamespace {
    /**
     * A list of content libraries.
     */
    contentLibraries?: string[];
    /**
     * The name of the namespace.
     */
    name: string;
    /**
     * A list of virtual machine classes.
     */
    vmClasses?: string[];
}

export interface SupervisorPodCidr {
    /**
     * Network address.
     */
    address: string;
    /**
     * Subnet prefix.
     */
    prefix: number;
}

export interface SupervisorServiceCidr {
    /**
     * Network address.
     */
    address: string;
    /**
     * Subnet prefix.
     */
    prefix: number;
}

export interface VirtualMachineCdrom {
    /**
     * Indicates whether the device should be mapped to a remote client device
     */
    clientDevice?: boolean;
    /**
     * The datastore ID the ISO is located on.
     */
    datastoreId?: string;
    /**
     * The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
     */
    deviceAddress: string;
    /**
     * The ID of the device within the virtual machine.
     */
    key: number;
    /**
     * The path to the ISO file on the datastore.
     */
    path?: string;
}

export interface VirtualMachineClone {
    /**
     * The customization specification for the virtual machine post-clone.
     */
    customizationSpec?: outputs.VirtualMachineCloneCustomizationSpec;
    /**
     * The customization specification for the virtual machine post-clone.
     */
    customize?: outputs.VirtualMachineCloneCustomize;
    /**
     * Whether or not to create a linked clone when cloning. When this option is used, the source VM must have a single snapshot associated with it.
     */
    linkedClone?: boolean;
    /**
     * Mapping of ovf networks to the networks to use in vSphere.
     */
    ovfNetworkMap?: {[key: string]: string};
    /**
     * Mapping of ovf storage to the datastores to use in vSphere.
     */
    ovfStorageMap?: {[key: string]: string};
    /**
     * The UUID of the source virtual machine or template.
     */
    templateUuid: string;
    /**
     * The timeout, in minutes, to wait for the virtual machine clone to complete.
     */
    timeout?: number;
}

export interface VirtualMachineCloneCustomizationSpec {
    /**
     * The UUID of the virtual machine.
     */
    id: string;
    /**
     * The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error. Setting this value to 0 or a negative value skips the waiter. Default: 10.
     */
    timeout?: number;
}

export interface VirtualMachineCloneCustomize {
    /**
     * The list of DNS servers for a virtual network adapter with a static IP address.
     */
    dnsServerLists?: string[];
    /**
     * A list of DNS search domains to add to the DNS configuration on the virtual machine.
     */
    dnsSuffixLists?: string[];
    /**
     * The IPv4 default gateway when using networkInterface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     */
    ipv4Gateway?: string;
    /**
     * The IPv6 default gateway when using networkInterface customization on the virtual machine. This address must be local to a static IPv4 address configured in an interface sub-resource.
     */
    ipv6Gateway?: string;
    /**
     * A list of configuration options specific to Linux virtual machines.
     */
    linuxOptions?: outputs.VirtualMachineCloneCustomizeLinuxOptions;
    /**
     * A specification of network interface configuration options.
     */
    networkInterfaces?: outputs.VirtualMachineCloneCustomizeNetworkInterface[];
    /**
     * The amount of time, in minutes, to wait for guest OS customization to complete before returning with an error. Setting this value to 0 or a negative value skips the waiter. Default: 10.
     */
    timeout?: number;
    /**
     * A list of configuration options specific to Windows virtual machines.
     */
    windowsOptions?: outputs.VirtualMachineCloneCustomizeWindowsOptions;
    /**
     * Use this option to specify a windows sysprep file directly.
     */
    windowsSysprepText?: string;
}

export interface VirtualMachineCloneCustomizeLinuxOptions {
    /**
     * The domain name for this virtual machine.
     */
    domain: string;
    /**
     * The hostname for this virtual machine.
     */
    hostName: string;
    /**
     * Specifies whether or not the hardware clock should be in UTC or not.
     */
    hwClockUtc?: boolean;
    /**
     * The customization script to run before and or after guest customization
     */
    scriptText?: string;
    /**
     * Customize the time zone on the VM. This should be a time zone-style entry, like America/Los_Angeles.
     */
    timeZone?: string;
}

export interface VirtualMachineCloneCustomizeNetworkInterface {
    /**
     * A DNS search domain to add to the DNS configuration on the virtual machine.
     */
    dnsDomain?: string;
    /**
     * Network-interface specific DNS settings for Windows operating systems. Ignored on Linux.
     */
    dnsServerLists?: string[];
    /**
     * The IPv4 address assigned to this network adapter. If left blank, DHCP is used.
     */
    ipv4Address?: string;
    /**
     * The IPv4 CIDR netmask for the supplied IP address. Ignored if DHCP is selected.
     */
    ipv4Netmask?: number;
    /**
     * The IPv6 address assigned to this network adapter. If left blank, default auto-configuration is used.
     */
    ipv6Address?: string;
    /**
     * The IPv6 CIDR netmask for the supplied IP address. Ignored if auto-configuration is selected.
     */
    ipv6Netmask?: number;
}

export interface VirtualMachineCloneCustomizeWindowsOptions {
    /**
     * The new administrator password for this virtual machine.
     */
    adminPassword?: string;
    /**
     * Specifies whether or not the VM automatically logs on as Administrator.
     */
    autoLogon?: boolean;
    /**
     * Specifies how many times the VM should auto-logon the Administrator account when autoLogon is true.
     */
    autoLogonCount?: number;
    /**
     * The host name for this virtual machine.
     */
    computerName: string;
    /**
     * The password of the domain administrator used to join this virtual machine to the domain.
     */
    domainAdminPassword?: string;
    /**
     * The user account of the domain administrator used to join this virtual machine to the domain.
     */
    domainAdminUser?: string;
    /**
     * The MachineObjectOU which specifies the full LDAP path name of the OU to which the virtual machine belongs.
     */
    domainOu?: string;
    /**
     * The full name of the user of this virtual machine.
     */
    fullName?: string;
    /**
     * The domain that the virtual machine should join.
     */
    joinDomain?: string;
    /**
     * The organization name this virtual machine is being installed for.
     */
    organizationName?: string;
    /**
     * The product key for this virtual machine.
     */
    productKey?: string;
    /**
     * A list of commands to run at first user logon, after guest customization.
     */
    runOnceCommandLists?: string[];
    /**
     * The new time zone for the virtual machine. This is a sysprep-dictated timezone code.
     */
    timeZone?: number;
    /**
     * The workgroup for this virtual machine if not joining a domain.
     */
    workgroup?: string;
}

export interface VirtualMachineDisk {
    /**
     * If this is true, the disk is attached instead of created. Implies keep_on_remove.
     */
    attach?: boolean;
    /**
     * The type of controller the disk should be connected to. Must be 'scsi', 'sata', or 'ide'.
     */
    controllerType?: string;
    /**
     * The datastore ID for this virtual disk, if different than the virtual machine.
     */
    datastoreId: string;
    /**
     * The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
     */
    deviceAddress: string;
    /**
     * The mode of this this virtual disk for purposes of writes and snapshotting. Can be one of append, independent_nonpersistent, independent_persistent, nonpersistent, persistent, or undoable.
     */
    diskMode?: string;
    /**
     * The sharing mode of this virtual disk. Can be one of sharingMultiWriter or sharingNone.
     */
    diskSharing?: string;
    /**
     * The virtual disk file zeroing policy when thinProvision is not true. The default is false, which lazily-zeros the disk, speeding up thick-provisioned disk creation time.
     */
    eagerlyScrub?: boolean;
    /**
     * The upper limit of IOPS that this disk can use.
     */
    ioLimit?: number;
    /**
     * The I/O guarantee that this disk has, in IOPS.
     */
    ioReservation?: number;
    /**
     * The share count for this disk when the share level is custom.
     */
    ioShareCount?: number;
    /**
     * The share allocation level for this disk. Can be one of low, normal, high, or custom.
     */
    ioShareLevel?: string;
    /**
     * Set to true to keep the underlying VMDK file when removing this virtual disk from configuration.
     */
    keepOnRemove?: boolean;
    /**
     * The ID of the device within the virtual machine.
     */
    key: number;
    /**
     * A unique label for this disk.
     */
    label: string;
    /**
     * The full path of the virtual disk. This can only be provided if attach is set to true, otherwise it is a read-only value.
     */
    path: string;
    /**
     * The size of the disk, in GB.
     */
    size?: number;
    /**
     * The ID of the storage policy to assign to the virtual disk in VM.
     */
    storagePolicyId: string;
    /**
     * If true, this disk is thin provisioned, with space for the file being allocated on an as-needed basis.
     */
    thinProvisioned?: boolean;
    /**
     * The unique device number for this disk. This number determines where on the SCSI bus this device will be attached.
     */
    unitNumber?: number;
    /**
     * The UUID of the virtual machine. Also exposed as the `id` of the resource.
     */
    uuid: string;
    /**
     * If true, writes for this disk are sent directly to the filesystem immediately instead of being buffered.
     */
    writeThrough?: boolean;
}

export interface VirtualMachineNetworkInterface {
    /**
     * The controller type. Can be one of e1000, e1000e, sriov, vmxnet3, or vrdma.
     */
    adapterType?: string;
    /**
     * The upper bandwidth limit of this network interface, in Mbits/sec.
     */
    bandwidthLimit?: number;
    /**
     * The bandwidth reservation of this network interface, in Mbits/sec.
     */
    bandwidthReservation?: number;
    /**
     * The share count for this network interface when the share level is custom.
     */
    bandwidthShareCount: number;
    /**
     * The bandwidth share allocation level for this interface. Can be one of low, normal, high, or custom.
     */
    bandwidthShareLevel?: string;
    /**
     * The internally-computed address of this device, such as scsi:0:1, denoting scsi bus #0 and device unit 1.
     */
    deviceAddress: string;
    /**
     * The ID of the device within the virtual machine.
     */
    key: number;
    /**
     * The MAC address of this network interface. Can only be manually set if useStaticMac is true.
     */
    macAddress: string;
    /**
     * The ID of the network to connect this network interface to.
     */
    networkId: string;
    /**
     * Mapping of network interface to OVF network.
     */
    ovfMapping?: string;
    /**
     * The ID of the Physical SR-IOV NIC to attach to, e.g. '0000:d8:00.0'
     */
    physicalFunction?: string;
    /**
     * If true, the macAddress field is treated as a static MAC address and set accordingly.
     */
    useStaticMac?: boolean;
}

export interface VirtualMachineOvfDeploy {
    /**
     * Allow unverified ssl certificates while deploying ovf/ova from url.
     */
    allowUnverifiedSslCert?: boolean;
    /**
     * The Deployment option to be chosen. If empty, the default option is used.
     */
    deploymentOption?: string;
    /**
     * An optional disk provisioning. If set, all the disks in the deployed ovf will have the same specified disk type (e.g., thin provisioned).
     */
    diskProvisioning?: string;
    /**
     * Allow properties with ovf:userConfigurable=false to be set.
     */
    enableHiddenProperties?: boolean;
    /**
     * The IP allocation policy.
     */
    ipAllocationPolicy?: string;
    /**
     * The IP protocol.
     */
    ipProtocol?: string;
    /**
     * The absolute path to the ovf/ova file in the local system.
     */
    localOvfPath?: string;
    /**
     * The mapping of name of network identifiers from the ovf descriptor to network UUID in the VI infrastructure.
     */
    ovfNetworkMap?: {[key: string]: string};
    /**
     * URL to the remote ovf/ova file to be deployed.
     */
    remoteOvfUrl?: string;
}

export interface VirtualMachineVapp {
    /**
     * A map of customizable vApp properties and their values. Allows customization of VMs cloned from OVF templates which have customizable vApp properties.
     */
    properties?: {[key: string]: string};
}

export interface VirtualMachineVtpm {
    /**
     * The version of the TPM device. Default is 2.0.
     */
    version?: string;
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
     * address of the interface, if DHCP is not set.
     */
    ip?: string;
    /**
     * netmask of the interface, if DHCP is not set.
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
     * IP address of the default gateway, if DHCP or autoconfig is not set.
     */
    gw?: string;
}

