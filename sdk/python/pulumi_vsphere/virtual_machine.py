# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities

class VirtualMachine(pulumi.CustomResource):
    """
    The `vsphere_virtual_machine` resource can be used to manage the complex
    lifecycle of a virtual machine. It supports management of disk, network
    interface, and CDROM devices, creation from scratch or cloning from template,
    and migration through both host and storage vMotion.
    
    For more details on working with virtual machines in vSphere, see [this
    page][vmware-docs-vm-management].
    
    [vmware-docs-vm-management]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-55238059-912E-411F-A0E9-A7A536972A91.html
    """
    def __init__(__self__, __name__, __opts__=None, alternate_guest_name=None, annotation=None, boot_delay=None, boot_retry_delay=None, boot_retry_enabled=None, cdrom=None, clone=None, cpu_hot_add_enabled=None, cpu_hot_remove_enabled=None, cpu_limit=None, cpu_performance_counters_enabled=None, cpu_reservation=None, cpu_share_count=None, cpu_share_level=None, custom_attributes=None, datastore_cluster_id=None, datastore_id=None, disks=None, efi_secure_boot_enabled=None, enable_disk_uuid=None, enable_logging=None, ept_rvi_mode=None, extra_config=None, firmware=None, folder=None, force_power_off=None, guest_id=None, host_system_id=None, hv_mode=None, latency_sensitivity=None, memory=None, memory_hot_add_enabled=None, memory_limit=None, memory_reservation=None, memory_share_count=None, memory_share_level=None, migrate_wait_timeout=None, name=None, nested_hv_enabled=None, network_interfaces=None, num_cores_per_socket=None, num_cpus=None, resource_pool_id=None, run_tools_scripts_after_power_on=None, run_tools_scripts_after_resume=None, run_tools_scripts_before_guest_reboot=None, run_tools_scripts_before_guest_shutdown=None, run_tools_scripts_before_guest_standby=None, scsi_bus_sharing=None, scsi_controller_count=None, scsi_type=None, shutdown_wait_timeout=None, swap_placement_policy=None, sync_time_with_host=None, tags=None, vapp=None, wait_for_guest_net_routable=None, wait_for_guest_net_timeout=None):
        """Create a VirtualMachine resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['alternateGuestName'] = alternate_guest_name

        __props__['annotation'] = annotation

        __props__['bootDelay'] = boot_delay

        __props__['bootRetryDelay'] = boot_retry_delay

        __props__['bootRetryEnabled'] = boot_retry_enabled

        __props__['cdrom'] = cdrom

        __props__['clone'] = clone

        __props__['cpuHotAddEnabled'] = cpu_hot_add_enabled

        __props__['cpuHotRemoveEnabled'] = cpu_hot_remove_enabled

        __props__['cpuLimit'] = cpu_limit

        __props__['cpuPerformanceCountersEnabled'] = cpu_performance_counters_enabled

        __props__['cpuReservation'] = cpu_reservation

        __props__['cpuShareCount'] = cpu_share_count

        __props__['cpuShareLevel'] = cpu_share_level

        __props__['customAttributes'] = custom_attributes

        __props__['datastoreClusterId'] = datastore_cluster_id

        __props__['datastoreId'] = datastore_id

        __props__['disks'] = disks

        __props__['efiSecureBootEnabled'] = efi_secure_boot_enabled

        __props__['enableDiskUuid'] = enable_disk_uuid

        __props__['enableLogging'] = enable_logging

        __props__['eptRviMode'] = ept_rvi_mode

        __props__['extraConfig'] = extra_config

        __props__['firmware'] = firmware

        __props__['folder'] = folder

        __props__['forcePowerOff'] = force_power_off

        __props__['guestId'] = guest_id

        __props__['hostSystemId'] = host_system_id

        __props__['hvMode'] = hv_mode

        __props__['latencySensitivity'] = latency_sensitivity

        __props__['memory'] = memory

        __props__['memoryHotAddEnabled'] = memory_hot_add_enabled

        __props__['memoryLimit'] = memory_limit

        __props__['memoryReservation'] = memory_reservation

        __props__['memoryShareCount'] = memory_share_count

        __props__['memoryShareLevel'] = memory_share_level

        __props__['migrateWaitTimeout'] = migrate_wait_timeout

        __props__['name'] = name

        __props__['nestedHvEnabled'] = nested_hv_enabled

        if not network_interfaces:
            raise TypeError('Missing required property network_interfaces')
        __props__['networkInterfaces'] = network_interfaces

        __props__['numCoresPerSocket'] = num_cores_per_socket

        __props__['numCpus'] = num_cpus

        if not resource_pool_id:
            raise TypeError('Missing required property resource_pool_id')
        __props__['resourcePoolId'] = resource_pool_id

        __props__['runToolsScriptsAfterPowerOn'] = run_tools_scripts_after_power_on

        __props__['runToolsScriptsAfterResume'] = run_tools_scripts_after_resume

        __props__['runToolsScriptsBeforeGuestReboot'] = run_tools_scripts_before_guest_reboot

        __props__['runToolsScriptsBeforeGuestShutdown'] = run_tools_scripts_before_guest_shutdown

        __props__['runToolsScriptsBeforeGuestStandby'] = run_tools_scripts_before_guest_standby

        __props__['scsiBusSharing'] = scsi_bus_sharing

        __props__['scsiControllerCount'] = scsi_controller_count

        __props__['scsiType'] = scsi_type

        __props__['shutdownWaitTimeout'] = shutdown_wait_timeout

        __props__['swapPlacementPolicy'] = swap_placement_policy

        __props__['syncTimeWithHost'] = sync_time_with_host

        __props__['tags'] = tags

        __props__['vapp'] = vapp

        __props__['waitForGuestNetRoutable'] = wait_for_guest_net_routable

        __props__['waitForGuestNetTimeout'] = wait_for_guest_net_timeout

        __props__['change_version'] = None
        __props__['default_ip_address'] = None
        __props__['guest_ip_addresses'] = None
        __props__['imported'] = None
        __props__['moid'] = None
        __props__['reboot_required'] = None
        __props__['uuid'] = None
        __props__['vapp_transports'] = None
        __props__['vmware_tools_status'] = None
        __props__['vmx_path'] = None

        super(VirtualMachine, __self__).__init__(
            'vsphere:index/virtualMachine:VirtualMachine',
            __name__,
            __props__,
            __opts__)

