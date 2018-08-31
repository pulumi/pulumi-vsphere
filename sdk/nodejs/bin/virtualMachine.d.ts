import * as pulumi from "@pulumi/pulumi";
/**
 * The `vsphere_virtual_machine` resource can be used to manage the complex
 * lifecycle of a virtual machine. It supports management of disk, network
 * interface, and CDROM devices, creation from scratch or cloning from template,
 * and migration through both host and storage vMotion.
 *
 * For more details on working with virtual machines in vSphere, see [this
 * page][vmware-docs-vm-management].
 *
 * [vmware-docs-vm-management]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-55238059-912E-411F-A0E9-A7A536972A91.html
 */
export declare class VirtualMachine extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualMachineState): VirtualMachine;
    /**
     * The guest name for the operating system
     * when `guest_id` is `other` or `other-64`.
     */
    readonly alternateGuestName: pulumi.Output<string | undefined>;
    /**
     * A user-provided description of the virtual machine.
     * The default is no annotation.
     */
    readonly annotation: pulumi.Output<string | undefined>;
    /**
     * The number of milliseconds to wait before starting
     * the boot sequence. The default is no delay.
     */
    readonly bootDelay: pulumi.Output<number | undefined>;
    /**
     * The number of milliseconds to wait before
     * retrying the boot sequence. This only valid if `boot_retry_enabled` is true.
     * Default: `10000` (10 seconds).
     */
    readonly bootRetryDelay: pulumi.Output<number | undefined>;
    /**
     * If set to true, a virtual machine that
     * fails to boot will try again after the delay defined in `boot_retry_delay`.
     * Default: `false`.
     */
    readonly bootRetryEnabled: pulumi.Output<boolean | undefined>;
    /**
     * A specification for a CDROM device on this virtual
     * machine. See CDROM options below.
     */
    readonly cdrom: pulumi.Output<{
        clientDevice?: boolean;
        datastoreId?: string;
        deviceAddress: string;
        key: number;
        path?: string;
    } | undefined>;
    /**
     * A unique identifier for a given version of the last
     * configuration applied, such the timestamp of the last update to the
     * configuration.
     */
    readonly changeVersion: pulumi.Output<string>;
    /**
     * When specified, the VM will be created as a clone of a
     * specified template. Optional customization options can be submitted as well.
     * See creating a virtual machine from a
     * template for more details.
     */
    readonly clone: pulumi.Output<{
        customize?: {
            dnsServerLists?: string[];
            dnsSuffixLists?: string[];
            ipv4Gateway?: string;
            ipv6Gateway?: string;
            linuxOptions?: {
                domain: string;
                hostName: string;
                hwClockUtc?: boolean;
                timeZone?: string;
            };
            networkInterfaces?: {
                dnsDomain?: string;
                dnsServerLists?: string[];
                ipv4Address?: string;
                ipv4Netmask?: number;
                ipv6Address?: string;
                ipv6Netmask?: number;
            }[];
            timeout?: number;
            windowsOptions?: {
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
            };
            windowsSysprepText?: string;
        };
        linkedClone?: boolean;
        templateUuid: string;
        timeout?: number;
    } | undefined>;
    /**
     * Allow CPUs to be added to this virtual
     * machine while it is running.
     */
    readonly cpuHotAddEnabled: pulumi.Output<boolean | undefined>;
    /**
     * Allow CPUs to be removed to this
     * virtual machine while it is running.
     */
    readonly cpuHotRemoveEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The maximum amount of CPU (in MHz) that this virtual
     * machine can consume, regardless of available resources. The default is no
     * limit.
     */
    readonly cpuLimit: pulumi.Output<number | undefined>;
    /**
     * Enable CPU performance
     * counters on this virtual machine. Default: `false`.
     */
    readonly cpuPerformanceCountersEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The amount of CPU (in MHz) that this virtual
     * machine is guaranteed. The default is no reservation.
     */
    readonly cpuReservation: pulumi.Output<number | undefined>;
    /**
     * The number of CPU shares allocated to the
     * virtual machine when the `cpu_share_level` is `custom`.
     */
    readonly cpuShareCount: pulumi.Output<number>;
    /**
     * The allocation level for CPU resources. Can be
     * one of `high`, `low`, `normal`, or `custom`. Default: `custom`.
     */
    readonly cpuShareLevel: pulumi.Output<string | undefined>;
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for virtual machine. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster ID to use. This setting
     * applies to entire virtual machine and implies that you wish to use Storage
     * DRS with this virtual machine. See the section on virtual machine
     * migration for details on changing this value.
     */
    readonly datastoreClusterId: pulumi.Output<string | undefined>;
    /**
     * The datastore ID that the ISO is located in.
     * Requried for using a datastore ISO. Conflicts with `client_device`.
     */
    readonly datastoreId: pulumi.Output<string>;
    /**
     * The IP address selected by Terraform to be used with
     * any [provisioners][tf-docs-provisioners] configured on this resource.
     * Whenever possible, this is the first IPv4 address that is reachable through
     * the default gateway configured on the machine, then the first reachable IPv6
     * address, and then the first general discovered address if neither exist. If
     * VMware tools is not running on the virtual machine, or if the VM is powered
     * off, this value will be blank.
     */
    readonly defaultIpAddress: pulumi.Output<string>;
    /**
     * A specification for a virtual disk device on this virtual
     * machine. See disk options below.
     */
    readonly disks: pulumi.Output<{
        attach?: boolean;
        datastoreId?: string;
        deviceAddress: string;
        diskMode?: string;
        diskSharing?: string;
        eagerlyScrub?: boolean;
        ioLimit?: number;
        ioReservation?: number;
        ioShareCount?: number;
        ioShareLevel?: string;
        keepOnRemove?: boolean;
        key: number;
        label?: string;
        name?: string;
        path: string;
        size?: number;
        thinProvisioned?: boolean;
        unitNumber?: number;
        uuid: string;
        writeThrough?: boolean;
    }[]>;
    /**
     * When the `firmware` type is set to is
     * `efi`, this enables EFI secure boot. Default: `false`.
     */
    readonly efiSecureBootEnabled: pulumi.Output<boolean | undefined>;
    /**
     * Expose the UUIDs of attached virtual disks to
     * the virtual machine, allowing access to them in the guest. Default: `false`.
     */
    readonly enableDiskUuid: pulumi.Output<boolean | undefined>;
    /**
     * Enable logging of virtual machine events to a
     * log file stored in the virtual machine directory. Default: `false`.
     */
    readonly enableLogging: pulumi.Output<boolean | undefined>;
    /**
     * The EPT/RVI (hardware memory virtualization)
     * setting for this virtual machine. Can be one of `automatic`, `on`, or `off`.
     * Default: `automatic`.
     */
    readonly eptRviMode: pulumi.Output<string | undefined>;
    /**
     * Extra configuration data for this virtual
     * machine. Can be used to supply advanced parameters not normally in
     * configuration, such as data for cloud-config (under the guestinfo namespace).
     */
    readonly extraConfig: pulumi.Output<{
        [key: string]: any;
    } | undefined>;
    /**
     * The firmware interface to use on the virtual machine.
     * Can be one of `bios` or `EFI`. Default: `bios`.
     */
    readonly firmware: pulumi.Output<string | undefined>;
    /**
     * The path to the folder to put this virtual machine in,
     * relative to the datacenter that the resource pool is in.
     */
    readonly folder: pulumi.Output<string | undefined>;
    /**
     * If a guest shutdown failed or timed out while
     * updating or destroying (see
     * `shutdown_wait_timeout`), force the power-off of
     * the virtual machine. Default: `true`.
     */
    readonly forcePowerOff: pulumi.Output<boolean | undefined>;
    /**
     * The guest ID for the operating system type. For a
     * full list of possible values, see [here][vmware-docs-guest-ids]. Default: `other-64`.
     */
    readonly guestId: pulumi.Output<string | undefined>;
    /**
     * The current list of IP addresses on this machine,
     * including the value of `default_ip_address`. If VMware tools is not running
     * on the virtual machine, or if the VM is powered off, this list will be empty.
     * * `moid`: The [managed object reference ID][docs-about-morefs] of the created
     * virtual machine.
     */
    readonly guestIpAddresses: pulumi.Output<string[]>;
    /**
     * An optional [managed object reference
     * ID][docs-about-morefs] of a host to put this virtual machine on. See the
     * section on virtual machine migration for
     * details on changing this value. If a `host_system_id` is not supplied,
     * vSphere will select a host in the resource pool to place the virtual machine,
     * according to any defaults or DRS policies in place.
     */
    readonly hostSystemId: pulumi.Output<string>;
    /**
     * The (non-nested) hardware virtualization setting for
     * this virtual machine. Can be one of `hvAuto`, `hvOn`, or `hvOff`. Default:
     * `hvAuto`.
     */
    readonly hvMode: pulumi.Output<string | undefined>;
    /**
     * This is flagged if the virtual machine has been imported, or the
     * state has been migrated from a previous version of the resource. It
     * influences the behavior of the first post-import apply operation. See the
     * section on importing below.
     */
    readonly imported: pulumi.Output<boolean>;
    /**
     * Controls the scheduling delay of the
     * virtual machine. Use a higher sensitivity for applications that require lower
     * latency, such as VOIP, media player applications, or applications that
     * require frequent access to mouse or keyboard devices. Can be one of `low`,
     * `normal`, `medium`, or `high`.
     */
    readonly latencySensitivity: pulumi.Output<string | undefined>;
    /**
     * The size of the virtual machine's memory, in MB.
     * Default: `1024` (1 GB).
     */
    readonly memory: pulumi.Output<number | undefined>;
    /**
     * Allow memory to be added to this
     * virtual machine while it is running.
     */
    readonly memoryHotAddEnabled: pulumi.Output<boolean | undefined>;
    /**
     * The maximum amount of memory (in MB) that this
     * virtual machine can consume, regardless of available resources. The default
     * is no limit.
     */
    readonly memoryLimit: pulumi.Output<number | undefined>;
    /**
     * The amount of memory (in MB) that this
     * virtual machine is guaranteed. The default is no reservation.
     */
    readonly memoryReservation: pulumi.Output<number | undefined>;
    /**
     * The number of memory shares allocated to
     * the virtual machine when the `memory_share_level` is `custom`.
     */
    readonly memoryShareCount: pulumi.Output<number>;
    /**
     * The allocation level for memory resources.
     * Can be one of `high`, `low`, `normal`, or `custom`. Default: `custom`.
     */
    readonly memoryShareLevel: pulumi.Output<string | undefined>;
    /**
     * The amount of time, in minutes, to wait
     * for a virtual machine migration to complete before failing. Default: 10
     * minutes. Also see the section on virtual machine
     * migration.
     */
    readonly migrateWaitTimeout: pulumi.Output<number | undefined>;
    /**
     * The machine object ID from VMWare
     */
    readonly moid: pulumi.Output<string>;
    /**
     * An alias for both `label` and `path`, the latter when
     * using `attach`. Required if not using `label`.
     */
    readonly name: pulumi.Output<string>;
    /**
     * Enable nested hardware virtualization on
     * this virtual machine, facilitating nested virtualization in the guest.
     * Default: `false`.
     */
    readonly nestedHvEnabled: pulumi.Output<boolean | undefined>;
    /**
     * A specification for a virtual NIC on this
     * virtual machine. See network interface options
     * below.
     */
    readonly networkInterfaces: pulumi.Output<{
        adapterType?: string;
        bandwidthLimit?: number;
        bandwidthReservation?: number;
        bandwidthShareCount: number;
        bandwidthShareLevel?: string;
        deviceAddress: string;
        key: number;
        macAddress: string;
        networkId: string;
        useStaticMac?: boolean;
    }[]>;
    /**
     * The number of cores to distribute among
     * the CPUs in this virtual machine. If specified, the value supplied to
     * `num_cpus` must be evenly divisible by this value. Default: `1`.
     */
    readonly numCoresPerSocket: pulumi.Output<number | undefined>;
    /**
     * The number of virtual processors to assign to this
     * virtual machine. Default: `1`.
     */
    readonly numCpus: pulumi.Output<number | undefined>;
    /**
     * Value internal to Terraform used to determine if a
     * configuration set change requires a reboot. This value is only useful during
     * an update process and gets reset on refresh.
     */
    readonly rebootRequired: pulumi.Output<boolean>;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the resource pool to put this virtual machine in.
     * See the section on virtual machine migration
     * for details on changing this value.
     */
    readonly resourcePoolId: pulumi.Output<string>;
    /**
     * Enable the execution of
     * post-power-on scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsAfterPowerOn: pulumi.Output<boolean | undefined>;
    /**
     * Enable the execution of
     * post-resume scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsAfterResume: pulumi.Output<boolean | undefined>;
    /**
     * Enable the execution of
     * pre-reboot scripts when VMware tools is installed. Default: `false`.
     */
    readonly runToolsScriptsBeforeGuestReboot: pulumi.Output<boolean | undefined>;
    /**
     * Enable the execution
     * of pre-shutdown scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsBeforeGuestShutdown: pulumi.Output<boolean | undefined>;
    /**
     * Enable the execution of
     * pre-standby scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsBeforeGuestStandby: pulumi.Output<boolean | undefined>;
    /**
     * Mode for sharing the SCSI bus. The modes are
     * physicalSharing, virtualSharing, and noSharing. Default: `noSharing`.
     */
    readonly scsiBusSharing: pulumi.Output<string | undefined>;
    /**
     * The number of SCSI controllers that
     * Terraform manages on this virtual machine. This directly affects the amount
     * of disks you can add to the virtual machine and the maximum disk unit number.
     * Note that lowering this value does not remove controllers. Default: `1`.
     */
    readonly scsiControllerCount: pulumi.Output<number | undefined>;
    /**
     * The type of SCSI bus this virtual machine will have.
     * Can be one of lsilogic (LSI Logic Parallel), lsilogic-sas (LSI Logic SAS) or
     * pvscsi (VMware Paravirtual). Defualt: `pvscsi`.
     */
    readonly scsiType: pulumi.Output<string | undefined>;
    /**
     * The amount of time, in minutes, to wait
     * for a graceful guest shutdown when making necessary updates to the virtual
     * machine. If `force_power_off` is set to true, the VM will be force powered-off
     * after this timeout, otherwise an error is returned. Default: 3 minutes.
     */
    readonly shutdownWaitTimeout: pulumi.Output<number | undefined>;
    /**
     * The swap file placement policy for this
     * virtual machine. Can be one of `inherit`, `hostLocal`, or `vmDirectory`.
     * Default: `inherit`.
     */
    readonly swapPlacementPolicy: pulumi.Output<string | undefined>;
    /**
     * Enable guest clock synchronization with
     * the host. Requires VMware tools to be installed. Default: `false`.
     */
    readonly syncTimeWithHost: pulumi.Output<boolean | undefined>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags: pulumi.Output<string[] | undefined>;
    /**
     * The UUID of the virtual disk's VMDK file. This is used to track the
     * virtual disk on the virtual machine.
     */
    readonly uuid: pulumi.Output<string>;
    /**
     * Optional vApp configuration. The only sub-key available
     * is `properties`, which is a key/value map of properties for virtual machines
     * imported from OVF or OVA files. See Using vApp properties to supply OVF/OVA
     * configuration for
     * more details.
     */
    readonly vapp: pulumi.Output<{
        properties?: {
            [key: string]: string;
        };
    } | undefined>;
    /**
     * Computed value which is only valid for cloned virtual
     * machines. A list of vApp transport methods supported by the source virtual
     * machine or template.
     */
    readonly vappTransports: pulumi.Output<string[]>;
    /**
     * The state of VMware tools in the guest. This will
     * determine the proper course of action for some device operations.
     */
    readonly vmwareToolsStatus: pulumi.Output<string>;
    /**
     * The path of the virtual machine's configuration file in the VM's
     * datastore.
     */
    readonly vmxPath: pulumi.Output<string>;
    /**
     * Controls whether or not the guest
     * network waiter waits for a routable address. When `false`, the waiter does
     * not wait for a default gateway, nor are IP addresses checked against any
     * discovered default gateways as part of its success criteria. Default: `true`.
     */
    readonly waitForGuestNetRoutable: pulumi.Output<boolean | undefined>;
    /**
     * The amount of time, in minutes, to
     * wait for an available IP address on this virtual machine. A value less than 1
     * disables the waiter. Default: 5 minutes.
     */
    readonly waitForGuestNetTimeout: pulumi.Output<number | undefined>;
    /**
     * Create a VirtualMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineArgs, opts?: pulumi.CustomResourceOptions);
}
/**
 * Input properties used for looking up and filtering VirtualMachine resources.
 */
export interface VirtualMachineState {
    /**
     * The guest name for the operating system
     * when `guest_id` is `other` or `other-64`.
     */
    readonly alternateGuestName?: pulumi.Input<string>;
    /**
     * A user-provided description of the virtual machine.
     * The default is no annotation.
     */
    readonly annotation?: pulumi.Input<string>;
    /**
     * The number of milliseconds to wait before starting
     * the boot sequence. The default is no delay.
     */
    readonly bootDelay?: pulumi.Input<number>;
    /**
     * The number of milliseconds to wait before
     * retrying the boot sequence. This only valid if `boot_retry_enabled` is true.
     * Default: `10000` (10 seconds).
     */
    readonly bootRetryDelay?: pulumi.Input<number>;
    /**
     * If set to true, a virtual machine that
     * fails to boot will try again after the delay defined in `boot_retry_delay`.
     * Default: `false`.
     */
    readonly bootRetryEnabled?: pulumi.Input<boolean>;
    /**
     * A specification for a CDROM device on this virtual
     * machine. See CDROM options below.
     */
    readonly cdrom?: pulumi.Input<{
        clientDevice?: pulumi.Input<boolean>;
        datastoreId?: pulumi.Input<string>;
        deviceAddress?: pulumi.Input<string>;
        key?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
    }>;
    /**
     * A unique identifier for a given version of the last
     * configuration applied, such the timestamp of the last update to the
     * configuration.
     */
    readonly changeVersion?: pulumi.Input<string>;
    /**
     * When specified, the VM will be created as a clone of a
     * specified template. Optional customization options can be submitted as well.
     * See creating a virtual machine from a
     * template for more details.
     */
    readonly clone?: pulumi.Input<{
        customize?: pulumi.Input<{
            dnsServerLists?: pulumi.Input<pulumi.Input<string>[]>;
            dnsSuffixLists?: pulumi.Input<pulumi.Input<string>[]>;
            ipv4Gateway?: pulumi.Input<string>;
            ipv6Gateway?: pulumi.Input<string>;
            linuxOptions?: pulumi.Input<{
                domain: pulumi.Input<string>;
                hostName: pulumi.Input<string>;
                hwClockUtc?: pulumi.Input<boolean>;
                timeZone?: pulumi.Input<string>;
            }>;
            networkInterfaces?: pulumi.Input<pulumi.Input<{
                dnsDomain?: pulumi.Input<string>;
                dnsServerLists?: pulumi.Input<pulumi.Input<string>[]>;
                ipv4Address?: pulumi.Input<string>;
                ipv4Netmask?: pulumi.Input<number>;
                ipv6Address?: pulumi.Input<string>;
                ipv6Netmask?: pulumi.Input<number>;
            }>[]>;
            timeout?: pulumi.Input<number>;
            windowsOptions?: pulumi.Input<{
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
            }>;
            windowsSysprepText?: pulumi.Input<string>;
        }>;
        linkedClone?: pulumi.Input<boolean>;
        templateUuid: pulumi.Input<string>;
        timeout?: pulumi.Input<number>;
    }>;
    /**
     * Allow CPUs to be added to this virtual
     * machine while it is running.
     */
    readonly cpuHotAddEnabled?: pulumi.Input<boolean>;
    /**
     * Allow CPUs to be removed to this
     * virtual machine while it is running.
     */
    readonly cpuHotRemoveEnabled?: pulumi.Input<boolean>;
    /**
     * The maximum amount of CPU (in MHz) that this virtual
     * machine can consume, regardless of available resources. The default is no
     * limit.
     */
    readonly cpuLimit?: pulumi.Input<number>;
    /**
     * Enable CPU performance
     * counters on this virtual machine. Default: `false`.
     */
    readonly cpuPerformanceCountersEnabled?: pulumi.Input<boolean>;
    /**
     * The amount of CPU (in MHz) that this virtual
     * machine is guaranteed. The default is no reservation.
     */
    readonly cpuReservation?: pulumi.Input<number>;
    /**
     * The number of CPU shares allocated to the
     * virtual machine when the `cpu_share_level` is `custom`.
     */
    readonly cpuShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for CPU resources. Can be
     * one of `high`, `low`, `normal`, or `custom`. Default: `custom`.
     */
    readonly cpuShareLevel?: pulumi.Input<string>;
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for virtual machine. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster ID to use. This setting
     * applies to entire virtual machine and implies that you wish to use Storage
     * DRS with this virtual machine. See the section on virtual machine
     * migration for details on changing this value.
     */
    readonly datastoreClusterId?: pulumi.Input<string>;
    /**
     * The datastore ID that the ISO is located in.
     * Requried for using a datastore ISO. Conflicts with `client_device`.
     */
    readonly datastoreId?: pulumi.Input<string>;
    /**
     * The IP address selected by Terraform to be used with
     * any [provisioners][tf-docs-provisioners] configured on this resource.
     * Whenever possible, this is the first IPv4 address that is reachable through
     * the default gateway configured on the machine, then the first reachable IPv6
     * address, and then the first general discovered address if neither exist. If
     * VMware tools is not running on the virtual machine, or if the VM is powered
     * off, this value will be blank.
     */
    readonly defaultIpAddress?: pulumi.Input<string>;
    /**
     * A specification for a virtual disk device on this virtual
     * machine. See disk options below.
     */
    readonly disks?: pulumi.Input<pulumi.Input<{
        attach?: pulumi.Input<boolean>;
        datastoreId?: pulumi.Input<string>;
        deviceAddress?: pulumi.Input<string>;
        diskMode?: pulumi.Input<string>;
        diskSharing?: pulumi.Input<string>;
        eagerlyScrub?: pulumi.Input<boolean>;
        ioLimit?: pulumi.Input<number>;
        ioReservation?: pulumi.Input<number>;
        ioShareCount?: pulumi.Input<number>;
        ioShareLevel?: pulumi.Input<string>;
        keepOnRemove?: pulumi.Input<boolean>;
        key?: pulumi.Input<number>;
        label?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
        path?: pulumi.Input<string>;
        size?: pulumi.Input<number>;
        thinProvisioned?: pulumi.Input<boolean>;
        unitNumber?: pulumi.Input<number>;
        uuid?: pulumi.Input<string>;
        writeThrough?: pulumi.Input<boolean>;
    }>[]>;
    /**
     * When the `firmware` type is set to is
     * `efi`, this enables EFI secure boot. Default: `false`.
     */
    readonly efiSecureBootEnabled?: pulumi.Input<boolean>;
    /**
     * Expose the UUIDs of attached virtual disks to
     * the virtual machine, allowing access to them in the guest. Default: `false`.
     */
    readonly enableDiskUuid?: pulumi.Input<boolean>;
    /**
     * Enable logging of virtual machine events to a
     * log file stored in the virtual machine directory. Default: `false`.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    /**
     * The EPT/RVI (hardware memory virtualization)
     * setting for this virtual machine. Can be one of `automatic`, `on`, or `off`.
     * Default: `automatic`.
     */
    readonly eptRviMode?: pulumi.Input<string>;
    /**
     * Extra configuration data for this virtual
     * machine. Can be used to supply advanced parameters not normally in
     * configuration, such as data for cloud-config (under the guestinfo namespace).
     */
    readonly extraConfig?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The firmware interface to use on the virtual machine.
     * Can be one of `bios` or `EFI`. Default: `bios`.
     */
    readonly firmware?: pulumi.Input<string>;
    /**
     * The path to the folder to put this virtual machine in,
     * relative to the datacenter that the resource pool is in.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * If a guest shutdown failed or timed out while
     * updating or destroying (see
     * `shutdown_wait_timeout`), force the power-off of
     * the virtual machine. Default: `true`.
     */
    readonly forcePowerOff?: pulumi.Input<boolean>;
    /**
     * The guest ID for the operating system type. For a
     * full list of possible values, see [here][vmware-docs-guest-ids]. Default: `other-64`.
     */
    readonly guestId?: pulumi.Input<string>;
    /**
     * The current list of IP addresses on this machine,
     * including the value of `default_ip_address`. If VMware tools is not running
     * on the virtual machine, or if the VM is powered off, this list will be empty.
     * * `moid`: The [managed object reference ID][docs-about-morefs] of the created
     * virtual machine.
     */
    readonly guestIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional [managed object reference
     * ID][docs-about-morefs] of a host to put this virtual machine on. See the
     * section on virtual machine migration for
     * details on changing this value. If a `host_system_id` is not supplied,
     * vSphere will select a host in the resource pool to place the virtual machine,
     * according to any defaults or DRS policies in place.
     */
    readonly hostSystemId?: pulumi.Input<string>;
    /**
     * The (non-nested) hardware virtualization setting for
     * this virtual machine. Can be one of `hvAuto`, `hvOn`, or `hvOff`. Default:
     * `hvAuto`.
     */
    readonly hvMode?: pulumi.Input<string>;
    /**
     * This is flagged if the virtual machine has been imported, or the
     * state has been migrated from a previous version of the resource. It
     * influences the behavior of the first post-import apply operation. See the
     * section on importing below.
     */
    readonly imported?: pulumi.Input<boolean>;
    /**
     * Controls the scheduling delay of the
     * virtual machine. Use a higher sensitivity for applications that require lower
     * latency, such as VOIP, media player applications, or applications that
     * require frequent access to mouse or keyboard devices. Can be one of `low`,
     * `normal`, `medium`, or `high`.
     */
    readonly latencySensitivity?: pulumi.Input<string>;
    /**
     * The size of the virtual machine's memory, in MB.
     * Default: `1024` (1 GB).
     */
    readonly memory?: pulumi.Input<number>;
    /**
     * Allow memory to be added to this
     * virtual machine while it is running.
     */
    readonly memoryHotAddEnabled?: pulumi.Input<boolean>;
    /**
     * The maximum amount of memory (in MB) that this
     * virtual machine can consume, regardless of available resources. The default
     * is no limit.
     */
    readonly memoryLimit?: pulumi.Input<number>;
    /**
     * The amount of memory (in MB) that this
     * virtual machine is guaranteed. The default is no reservation.
     */
    readonly memoryReservation?: pulumi.Input<number>;
    /**
     * The number of memory shares allocated to
     * the virtual machine when the `memory_share_level` is `custom`.
     */
    readonly memoryShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for memory resources.
     * Can be one of `high`, `low`, `normal`, or `custom`. Default: `custom`.
     */
    readonly memoryShareLevel?: pulumi.Input<string>;
    /**
     * The amount of time, in minutes, to wait
     * for a virtual machine migration to complete before failing. Default: 10
     * minutes. Also see the section on virtual machine
     * migration.
     */
    readonly migrateWaitTimeout?: pulumi.Input<number>;
    /**
     * The machine object ID from VMWare
     */
    readonly moid?: pulumi.Input<string>;
    /**
     * An alias for both `label` and `path`, the latter when
     * using `attach`. Required if not using `label`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Enable nested hardware virtualization on
     * this virtual machine, facilitating nested virtualization in the guest.
     * Default: `false`.
     */
    readonly nestedHvEnabled?: pulumi.Input<boolean>;
    /**
     * A specification for a virtual NIC on this
     * virtual machine. See network interface options
     * below.
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<{
        adapterType?: pulumi.Input<string>;
        bandwidthLimit?: pulumi.Input<number>;
        bandwidthReservation?: pulumi.Input<number>;
        bandwidthShareCount?: pulumi.Input<number>;
        bandwidthShareLevel?: pulumi.Input<string>;
        deviceAddress?: pulumi.Input<string>;
        key?: pulumi.Input<number>;
        macAddress?: pulumi.Input<string>;
        networkId: pulumi.Input<string>;
        useStaticMac?: pulumi.Input<boolean>;
    }>[]>;
    /**
     * The number of cores to distribute among
     * the CPUs in this virtual machine. If specified, the value supplied to
     * `num_cpus` must be evenly divisible by this value. Default: `1`.
     */
    readonly numCoresPerSocket?: pulumi.Input<number>;
    /**
     * The number of virtual processors to assign to this
     * virtual machine. Default: `1`.
     */
    readonly numCpus?: pulumi.Input<number>;
    /**
     * Value internal to Terraform used to determine if a
     * configuration set change requires a reboot. This value is only useful during
     * an update process and gets reset on refresh.
     */
    readonly rebootRequired?: pulumi.Input<boolean>;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the resource pool to put this virtual machine in.
     * See the section on virtual machine migration
     * for details on changing this value.
     */
    readonly resourcePoolId?: pulumi.Input<string>;
    /**
     * Enable the execution of
     * post-power-on scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsAfterPowerOn?: pulumi.Input<boolean>;
    /**
     * Enable the execution of
     * post-resume scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsAfterResume?: pulumi.Input<boolean>;
    /**
     * Enable the execution of
     * pre-reboot scripts when VMware tools is installed. Default: `false`.
     */
    readonly runToolsScriptsBeforeGuestReboot?: pulumi.Input<boolean>;
    /**
     * Enable the execution
     * of pre-shutdown scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsBeforeGuestShutdown?: pulumi.Input<boolean>;
    /**
     * Enable the execution of
     * pre-standby scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsBeforeGuestStandby?: pulumi.Input<boolean>;
    /**
     * Mode for sharing the SCSI bus. The modes are
     * physicalSharing, virtualSharing, and noSharing. Default: `noSharing`.
     */
    readonly scsiBusSharing?: pulumi.Input<string>;
    /**
     * The number of SCSI controllers that
     * Terraform manages on this virtual machine. This directly affects the amount
     * of disks you can add to the virtual machine and the maximum disk unit number.
     * Note that lowering this value does not remove controllers. Default: `1`.
     */
    readonly scsiControllerCount?: pulumi.Input<number>;
    /**
     * The type of SCSI bus this virtual machine will have.
     * Can be one of lsilogic (LSI Logic Parallel), lsilogic-sas (LSI Logic SAS) or
     * pvscsi (VMware Paravirtual). Defualt: `pvscsi`.
     */
    readonly scsiType?: pulumi.Input<string>;
    /**
     * The amount of time, in minutes, to wait
     * for a graceful guest shutdown when making necessary updates to the virtual
     * machine. If `force_power_off` is set to true, the VM will be force powered-off
     * after this timeout, otherwise an error is returned. Default: 3 minutes.
     */
    readonly shutdownWaitTimeout?: pulumi.Input<number>;
    /**
     * The swap file placement policy for this
     * virtual machine. Can be one of `inherit`, `hostLocal`, or `vmDirectory`.
     * Default: `inherit`.
     */
    readonly swapPlacementPolicy?: pulumi.Input<string>;
    /**
     * Enable guest clock synchronization with
     * the host. Requires VMware tools to be installed. Default: `false`.
     */
    readonly syncTimeWithHost?: pulumi.Input<boolean>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The UUID of the virtual disk's VMDK file. This is used to track the
     * virtual disk on the virtual machine.
     */
    readonly uuid?: pulumi.Input<string>;
    /**
     * Optional vApp configuration. The only sub-key available
     * is `properties`, which is a key/value map of properties for virtual machines
     * imported from OVF or OVA files. See Using vApp properties to supply OVF/OVA
     * configuration for
     * more details.
     */
    readonly vapp?: pulumi.Input<{
        properties?: pulumi.Input<{
            [key: string]: pulumi.Input<string>;
        }>;
    }>;
    /**
     * Computed value which is only valid for cloned virtual
     * machines. A list of vApp transport methods supported by the source virtual
     * machine or template.
     */
    readonly vappTransports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The state of VMware tools in the guest. This will
     * determine the proper course of action for some device operations.
     */
    readonly vmwareToolsStatus?: pulumi.Input<string>;
    /**
     * The path of the virtual machine's configuration file in the VM's
     * datastore.
     */
    readonly vmxPath?: pulumi.Input<string>;
    /**
     * Controls whether or not the guest
     * network waiter waits for a routable address. When `false`, the waiter does
     * not wait for a default gateway, nor are IP addresses checked against any
     * discovered default gateways as part of its success criteria. Default: `true`.
     */
    readonly waitForGuestNetRoutable?: pulumi.Input<boolean>;
    /**
     * The amount of time, in minutes, to
     * wait for an available IP address on this virtual machine. A value less than 1
     * disables the waiter. Default: 5 minutes.
     */
    readonly waitForGuestNetTimeout?: pulumi.Input<number>;
}
/**
 * The set of arguments for constructing a VirtualMachine resource.
 */
export interface VirtualMachineArgs {
    /**
     * The guest name for the operating system
     * when `guest_id` is `other` or `other-64`.
     */
    readonly alternateGuestName?: pulumi.Input<string>;
    /**
     * A user-provided description of the virtual machine.
     * The default is no annotation.
     */
    readonly annotation?: pulumi.Input<string>;
    /**
     * The number of milliseconds to wait before starting
     * the boot sequence. The default is no delay.
     */
    readonly bootDelay?: pulumi.Input<number>;
    /**
     * The number of milliseconds to wait before
     * retrying the boot sequence. This only valid if `boot_retry_enabled` is true.
     * Default: `10000` (10 seconds).
     */
    readonly bootRetryDelay?: pulumi.Input<number>;
    /**
     * If set to true, a virtual machine that
     * fails to boot will try again after the delay defined in `boot_retry_delay`.
     * Default: `false`.
     */
    readonly bootRetryEnabled?: pulumi.Input<boolean>;
    /**
     * A specification for a CDROM device on this virtual
     * machine. See CDROM options below.
     */
    readonly cdrom?: pulumi.Input<{
        clientDevice?: pulumi.Input<boolean>;
        datastoreId?: pulumi.Input<string>;
        deviceAddress?: pulumi.Input<string>;
        key?: pulumi.Input<number>;
        path?: pulumi.Input<string>;
    }>;
    /**
     * When specified, the VM will be created as a clone of a
     * specified template. Optional customization options can be submitted as well.
     * See creating a virtual machine from a
     * template for more details.
     */
    readonly clone?: pulumi.Input<{
        customize?: pulumi.Input<{
            dnsServerLists?: pulumi.Input<pulumi.Input<string>[]>;
            dnsSuffixLists?: pulumi.Input<pulumi.Input<string>[]>;
            ipv4Gateway?: pulumi.Input<string>;
            ipv6Gateway?: pulumi.Input<string>;
            linuxOptions?: pulumi.Input<{
                domain: pulumi.Input<string>;
                hostName: pulumi.Input<string>;
                hwClockUtc?: pulumi.Input<boolean>;
                timeZone?: pulumi.Input<string>;
            }>;
            networkInterfaces?: pulumi.Input<pulumi.Input<{
                dnsDomain?: pulumi.Input<string>;
                dnsServerLists?: pulumi.Input<pulumi.Input<string>[]>;
                ipv4Address?: pulumi.Input<string>;
                ipv4Netmask?: pulumi.Input<number>;
                ipv6Address?: pulumi.Input<string>;
                ipv6Netmask?: pulumi.Input<number>;
            }>[]>;
            timeout?: pulumi.Input<number>;
            windowsOptions?: pulumi.Input<{
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
            }>;
            windowsSysprepText?: pulumi.Input<string>;
        }>;
        linkedClone?: pulumi.Input<boolean>;
        templateUuid: pulumi.Input<string>;
        timeout?: pulumi.Input<number>;
    }>;
    /**
     * Allow CPUs to be added to this virtual
     * machine while it is running.
     */
    readonly cpuHotAddEnabled?: pulumi.Input<boolean>;
    /**
     * Allow CPUs to be removed to this
     * virtual machine while it is running.
     */
    readonly cpuHotRemoveEnabled?: pulumi.Input<boolean>;
    /**
     * The maximum amount of CPU (in MHz) that this virtual
     * machine can consume, regardless of available resources. The default is no
     * limit.
     */
    readonly cpuLimit?: pulumi.Input<number>;
    /**
     * Enable CPU performance
     * counters on this virtual machine. Default: `false`.
     */
    readonly cpuPerformanceCountersEnabled?: pulumi.Input<boolean>;
    /**
     * The amount of CPU (in MHz) that this virtual
     * machine is guaranteed. The default is no reservation.
     */
    readonly cpuReservation?: pulumi.Input<number>;
    /**
     * The number of CPU shares allocated to the
     * virtual machine when the `cpu_share_level` is `custom`.
     */
    readonly cpuShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for CPU resources. Can be
     * one of `high`, `low`, `normal`, or `custom`. Default: `custom`.
     */
    readonly cpuShareLevel?: pulumi.Input<string>;
    /**
     * Map of custom attribute ids to attribute
     * value strings to set for virtual machine. See
     * [here][docs-setting-custom-attributes] for a reference on how to set values
     * for custom attributes.
     */
    readonly customAttributes?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the datastore cluster ID to use. This setting
     * applies to entire virtual machine and implies that you wish to use Storage
     * DRS with this virtual machine. See the section on virtual machine
     * migration for details on changing this value.
     */
    readonly datastoreClusterId?: pulumi.Input<string>;
    /**
     * The datastore ID that the ISO is located in.
     * Requried for using a datastore ISO. Conflicts with `client_device`.
     */
    readonly datastoreId?: pulumi.Input<string>;
    /**
     * A specification for a virtual disk device on this virtual
     * machine. See disk options below.
     */
    readonly disks?: pulumi.Input<pulumi.Input<{
        attach?: pulumi.Input<boolean>;
        datastoreId?: pulumi.Input<string>;
        deviceAddress?: pulumi.Input<string>;
        diskMode?: pulumi.Input<string>;
        diskSharing?: pulumi.Input<string>;
        eagerlyScrub?: pulumi.Input<boolean>;
        ioLimit?: pulumi.Input<number>;
        ioReservation?: pulumi.Input<number>;
        ioShareCount?: pulumi.Input<number>;
        ioShareLevel?: pulumi.Input<string>;
        keepOnRemove?: pulumi.Input<boolean>;
        key?: pulumi.Input<number>;
        label?: pulumi.Input<string>;
        name?: pulumi.Input<string>;
        path?: pulumi.Input<string>;
        size?: pulumi.Input<number>;
        thinProvisioned?: pulumi.Input<boolean>;
        unitNumber?: pulumi.Input<number>;
        uuid?: pulumi.Input<string>;
        writeThrough?: pulumi.Input<boolean>;
    }>[]>;
    /**
     * When the `firmware` type is set to is
     * `efi`, this enables EFI secure boot. Default: `false`.
     */
    readonly efiSecureBootEnabled?: pulumi.Input<boolean>;
    /**
     * Expose the UUIDs of attached virtual disks to
     * the virtual machine, allowing access to them in the guest. Default: `false`.
     */
    readonly enableDiskUuid?: pulumi.Input<boolean>;
    /**
     * Enable logging of virtual machine events to a
     * log file stored in the virtual machine directory. Default: `false`.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    /**
     * The EPT/RVI (hardware memory virtualization)
     * setting for this virtual machine. Can be one of `automatic`, `on`, or `off`.
     * Default: `automatic`.
     */
    readonly eptRviMode?: pulumi.Input<string>;
    /**
     * Extra configuration data for this virtual
     * machine. Can be used to supply advanced parameters not normally in
     * configuration, such as data for cloud-config (under the guestinfo namespace).
     */
    readonly extraConfig?: pulumi.Input<{
        [key: string]: any;
    }>;
    /**
     * The firmware interface to use on the virtual machine.
     * Can be one of `bios` or `EFI`. Default: `bios`.
     */
    readonly firmware?: pulumi.Input<string>;
    /**
     * The path to the folder to put this virtual machine in,
     * relative to the datacenter that the resource pool is in.
     */
    readonly folder?: pulumi.Input<string>;
    /**
     * If a guest shutdown failed or timed out while
     * updating or destroying (see
     * `shutdown_wait_timeout`), force the power-off of
     * the virtual machine. Default: `true`.
     */
    readonly forcePowerOff?: pulumi.Input<boolean>;
    /**
     * The guest ID for the operating system type. For a
     * full list of possible values, see [here][vmware-docs-guest-ids]. Default: `other-64`.
     */
    readonly guestId?: pulumi.Input<string>;
    /**
     * An optional [managed object reference
     * ID][docs-about-morefs] of a host to put this virtual machine on. See the
     * section on virtual machine migration for
     * details on changing this value. If a `host_system_id` is not supplied,
     * vSphere will select a host in the resource pool to place the virtual machine,
     * according to any defaults or DRS policies in place.
     */
    readonly hostSystemId?: pulumi.Input<string>;
    /**
     * The (non-nested) hardware virtualization setting for
     * this virtual machine. Can be one of `hvAuto`, `hvOn`, or `hvOff`. Default:
     * `hvAuto`.
     */
    readonly hvMode?: pulumi.Input<string>;
    /**
     * Controls the scheduling delay of the
     * virtual machine. Use a higher sensitivity for applications that require lower
     * latency, such as VOIP, media player applications, or applications that
     * require frequent access to mouse or keyboard devices. Can be one of `low`,
     * `normal`, `medium`, or `high`.
     */
    readonly latencySensitivity?: pulumi.Input<string>;
    /**
     * The size of the virtual machine's memory, in MB.
     * Default: `1024` (1 GB).
     */
    readonly memory?: pulumi.Input<number>;
    /**
     * Allow memory to be added to this
     * virtual machine while it is running.
     */
    readonly memoryHotAddEnabled?: pulumi.Input<boolean>;
    /**
     * The maximum amount of memory (in MB) that this
     * virtual machine can consume, regardless of available resources. The default
     * is no limit.
     */
    readonly memoryLimit?: pulumi.Input<number>;
    /**
     * The amount of memory (in MB) that this
     * virtual machine is guaranteed. The default is no reservation.
     */
    readonly memoryReservation?: pulumi.Input<number>;
    /**
     * The number of memory shares allocated to
     * the virtual machine when the `memory_share_level` is `custom`.
     */
    readonly memoryShareCount?: pulumi.Input<number>;
    /**
     * The allocation level for memory resources.
     * Can be one of `high`, `low`, `normal`, or `custom`. Default: `custom`.
     */
    readonly memoryShareLevel?: pulumi.Input<string>;
    /**
     * The amount of time, in minutes, to wait
     * for a virtual machine migration to complete before failing. Default: 10
     * minutes. Also see the section on virtual machine
     * migration.
     */
    readonly migrateWaitTimeout?: pulumi.Input<number>;
    /**
     * An alias for both `label` and `path`, the latter when
     * using `attach`. Required if not using `label`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Enable nested hardware virtualization on
     * this virtual machine, facilitating nested virtualization in the guest.
     * Default: `false`.
     */
    readonly nestedHvEnabled?: pulumi.Input<boolean>;
    /**
     * A specification for a virtual NIC on this
     * virtual machine. See network interface options
     * below.
     */
    readonly networkInterfaces: pulumi.Input<pulumi.Input<{
        adapterType?: pulumi.Input<string>;
        bandwidthLimit?: pulumi.Input<number>;
        bandwidthReservation?: pulumi.Input<number>;
        bandwidthShareCount?: pulumi.Input<number>;
        bandwidthShareLevel?: pulumi.Input<string>;
        deviceAddress?: pulumi.Input<string>;
        key?: pulumi.Input<number>;
        macAddress?: pulumi.Input<string>;
        networkId: pulumi.Input<string>;
        useStaticMac?: pulumi.Input<boolean>;
    }>[]>;
    /**
     * The number of cores to distribute among
     * the CPUs in this virtual machine. If specified, the value supplied to
     * `num_cpus` must be evenly divisible by this value. Default: `1`.
     */
    readonly numCoresPerSocket?: pulumi.Input<number>;
    /**
     * The number of virtual processors to assign to this
     * virtual machine. Default: `1`.
     */
    readonly numCpus?: pulumi.Input<number>;
    /**
     * The [managed object reference
     * ID][docs-about-morefs] of the resource pool to put this virtual machine in.
     * See the section on virtual machine migration
     * for details on changing this value.
     */
    readonly resourcePoolId: pulumi.Input<string>;
    /**
     * Enable the execution of
     * post-power-on scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsAfterPowerOn?: pulumi.Input<boolean>;
    /**
     * Enable the execution of
     * post-resume scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsAfterResume?: pulumi.Input<boolean>;
    /**
     * Enable the execution of
     * pre-reboot scripts when VMware tools is installed. Default: `false`.
     */
    readonly runToolsScriptsBeforeGuestReboot?: pulumi.Input<boolean>;
    /**
     * Enable the execution
     * of pre-shutdown scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsBeforeGuestShutdown?: pulumi.Input<boolean>;
    /**
     * Enable the execution of
     * pre-standby scripts when VMware tools is installed. Default: `true`.
     */
    readonly runToolsScriptsBeforeGuestStandby?: pulumi.Input<boolean>;
    /**
     * Mode for sharing the SCSI bus. The modes are
     * physicalSharing, virtualSharing, and noSharing. Default: `noSharing`.
     */
    readonly scsiBusSharing?: pulumi.Input<string>;
    /**
     * The number of SCSI controllers that
     * Terraform manages on this virtual machine. This directly affects the amount
     * of disks you can add to the virtual machine and the maximum disk unit number.
     * Note that lowering this value does not remove controllers. Default: `1`.
     */
    readonly scsiControllerCount?: pulumi.Input<number>;
    /**
     * The type of SCSI bus this virtual machine will have.
     * Can be one of lsilogic (LSI Logic Parallel), lsilogic-sas (LSI Logic SAS) or
     * pvscsi (VMware Paravirtual). Defualt: `pvscsi`.
     */
    readonly scsiType?: pulumi.Input<string>;
    /**
     * The amount of time, in minutes, to wait
     * for a graceful guest shutdown when making necessary updates to the virtual
     * machine. If `force_power_off` is set to true, the VM will be force powered-off
     * after this timeout, otherwise an error is returned. Default: 3 minutes.
     */
    readonly shutdownWaitTimeout?: pulumi.Input<number>;
    /**
     * The swap file placement policy for this
     * virtual machine. Can be one of `inherit`, `hostLocal`, or `vmDirectory`.
     * Default: `inherit`.
     */
    readonly swapPlacementPolicy?: pulumi.Input<string>;
    /**
     * Enable guest clock synchronization with
     * the host. Requires VMware tools to be installed. Default: `false`.
     */
    readonly syncTimeWithHost?: pulumi.Input<boolean>;
    /**
     * The IDs of any tags to attach to this resource. See
     * [here][docs-applying-tags] for a reference on how to apply tags.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional vApp configuration. The only sub-key available
     * is `properties`, which is a key/value map of properties for virtual machines
     * imported from OVF or OVA files. See Using vApp properties to supply OVF/OVA
     * configuration for
     * more details.
     */
    readonly vapp?: pulumi.Input<{
        properties?: pulumi.Input<{
            [key: string]: pulumi.Input<string>;
        }>;
    }>;
    /**
     * Controls whether or not the guest
     * network waiter waits for a routable address. When `false`, the waiter does
     * not wait for a default gateway, nor are IP addresses checked against any
     * discovered default gateways as part of its success criteria. Default: `true`.
     */
    readonly waitForGuestNetRoutable?: pulumi.Input<boolean>;
    /**
     * The amount of time, in minutes, to
     * wait for an available IP address on this virtual machine. A value less than 1
     * disables the waiter. Default: 5 minutes.
     */
    readonly waitForGuestNetTimeout?: pulumi.Input<number>;
}
