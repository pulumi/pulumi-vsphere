# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOvfVmTemplateResult',
    'AwaitableGetOvfVmTemplateResult',
    'get_ovf_vm_template',
    'get_ovf_vm_template_output',
]

@pulumi.output_type
class GetOvfVmTemplateResult:
    """
    A collection of values returned by getOvfVmTemplate.
    """
    def __init__(__self__, allow_unverified_ssl_cert=None, alternate_guest_name=None, annotation=None, cpu_hot_add_enabled=None, cpu_hot_remove_enabled=None, cpu_performance_counters_enabled=None, datastore_id=None, deployment_option=None, disk_provisioning=None, enable_hidden_properties=None, firmware=None, folder=None, guest_id=None, host_system_id=None, id=None, ide_controller_count=None, ip_allocation_policy=None, ip_protocol=None, local_ovf_path=None, memory=None, memory_hot_add_enabled=None, name=None, nested_hv_enabled=None, num_cores_per_socket=None, num_cpus=None, ovf_network_map=None, remote_ovf_url=None, resource_pool_id=None, sata_controller_count=None, scsi_controller_count=None, scsi_type=None, swap_placement_policy=None):
        if allow_unverified_ssl_cert and not isinstance(allow_unverified_ssl_cert, bool):
            raise TypeError("Expected argument 'allow_unverified_ssl_cert' to be a bool")
        pulumi.set(__self__, "allow_unverified_ssl_cert", allow_unverified_ssl_cert)
        if alternate_guest_name and not isinstance(alternate_guest_name, str):
            raise TypeError("Expected argument 'alternate_guest_name' to be a str")
        pulumi.set(__self__, "alternate_guest_name", alternate_guest_name)
        if annotation and not isinstance(annotation, str):
            raise TypeError("Expected argument 'annotation' to be a str")
        pulumi.set(__self__, "annotation", annotation)
        if cpu_hot_add_enabled and not isinstance(cpu_hot_add_enabled, bool):
            raise TypeError("Expected argument 'cpu_hot_add_enabled' to be a bool")
        pulumi.set(__self__, "cpu_hot_add_enabled", cpu_hot_add_enabled)
        if cpu_hot_remove_enabled and not isinstance(cpu_hot_remove_enabled, bool):
            raise TypeError("Expected argument 'cpu_hot_remove_enabled' to be a bool")
        pulumi.set(__self__, "cpu_hot_remove_enabled", cpu_hot_remove_enabled)
        if cpu_performance_counters_enabled and not isinstance(cpu_performance_counters_enabled, bool):
            raise TypeError("Expected argument 'cpu_performance_counters_enabled' to be a bool")
        pulumi.set(__self__, "cpu_performance_counters_enabled", cpu_performance_counters_enabled)
        if datastore_id and not isinstance(datastore_id, str):
            raise TypeError("Expected argument 'datastore_id' to be a str")
        pulumi.set(__self__, "datastore_id", datastore_id)
        if deployment_option and not isinstance(deployment_option, str):
            raise TypeError("Expected argument 'deployment_option' to be a str")
        pulumi.set(__self__, "deployment_option", deployment_option)
        if disk_provisioning and not isinstance(disk_provisioning, str):
            raise TypeError("Expected argument 'disk_provisioning' to be a str")
        pulumi.set(__self__, "disk_provisioning", disk_provisioning)
        if enable_hidden_properties and not isinstance(enable_hidden_properties, bool):
            raise TypeError("Expected argument 'enable_hidden_properties' to be a bool")
        pulumi.set(__self__, "enable_hidden_properties", enable_hidden_properties)
        if firmware and not isinstance(firmware, str):
            raise TypeError("Expected argument 'firmware' to be a str")
        pulumi.set(__self__, "firmware", firmware)
        if folder and not isinstance(folder, str):
            raise TypeError("Expected argument 'folder' to be a str")
        pulumi.set(__self__, "folder", folder)
        if guest_id and not isinstance(guest_id, str):
            raise TypeError("Expected argument 'guest_id' to be a str")
        pulumi.set(__self__, "guest_id", guest_id)
        if host_system_id and not isinstance(host_system_id, str):
            raise TypeError("Expected argument 'host_system_id' to be a str")
        pulumi.set(__self__, "host_system_id", host_system_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ide_controller_count and not isinstance(ide_controller_count, int):
            raise TypeError("Expected argument 'ide_controller_count' to be a int")
        pulumi.set(__self__, "ide_controller_count", ide_controller_count)
        if ip_allocation_policy and not isinstance(ip_allocation_policy, str):
            raise TypeError("Expected argument 'ip_allocation_policy' to be a str")
        pulumi.set(__self__, "ip_allocation_policy", ip_allocation_policy)
        if ip_protocol and not isinstance(ip_protocol, str):
            raise TypeError("Expected argument 'ip_protocol' to be a str")
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        if local_ovf_path and not isinstance(local_ovf_path, str):
            raise TypeError("Expected argument 'local_ovf_path' to be a str")
        pulumi.set(__self__, "local_ovf_path", local_ovf_path)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if memory_hot_add_enabled and not isinstance(memory_hot_add_enabled, bool):
            raise TypeError("Expected argument 'memory_hot_add_enabled' to be a bool")
        pulumi.set(__self__, "memory_hot_add_enabled", memory_hot_add_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nested_hv_enabled and not isinstance(nested_hv_enabled, bool):
            raise TypeError("Expected argument 'nested_hv_enabled' to be a bool")
        pulumi.set(__self__, "nested_hv_enabled", nested_hv_enabled)
        if num_cores_per_socket and not isinstance(num_cores_per_socket, int):
            raise TypeError("Expected argument 'num_cores_per_socket' to be a int")
        pulumi.set(__self__, "num_cores_per_socket", num_cores_per_socket)
        if num_cpus and not isinstance(num_cpus, int):
            raise TypeError("Expected argument 'num_cpus' to be a int")
        pulumi.set(__self__, "num_cpus", num_cpus)
        if ovf_network_map and not isinstance(ovf_network_map, dict):
            raise TypeError("Expected argument 'ovf_network_map' to be a dict")
        pulumi.set(__self__, "ovf_network_map", ovf_network_map)
        if remote_ovf_url and not isinstance(remote_ovf_url, str):
            raise TypeError("Expected argument 'remote_ovf_url' to be a str")
        pulumi.set(__self__, "remote_ovf_url", remote_ovf_url)
        if resource_pool_id and not isinstance(resource_pool_id, str):
            raise TypeError("Expected argument 'resource_pool_id' to be a str")
        pulumi.set(__self__, "resource_pool_id", resource_pool_id)
        if sata_controller_count and not isinstance(sata_controller_count, int):
            raise TypeError("Expected argument 'sata_controller_count' to be a int")
        pulumi.set(__self__, "sata_controller_count", sata_controller_count)
        if scsi_controller_count and not isinstance(scsi_controller_count, int):
            raise TypeError("Expected argument 'scsi_controller_count' to be a int")
        pulumi.set(__self__, "scsi_controller_count", scsi_controller_count)
        if scsi_type and not isinstance(scsi_type, str):
            raise TypeError("Expected argument 'scsi_type' to be a str")
        pulumi.set(__self__, "scsi_type", scsi_type)
        if swap_placement_policy and not isinstance(swap_placement_policy, str):
            raise TypeError("Expected argument 'swap_placement_policy' to be a str")
        pulumi.set(__self__, "swap_placement_policy", swap_placement_policy)

    @property
    @pulumi.getter(name="allowUnverifiedSslCert")
    def allow_unverified_ssl_cert(self) -> Optional[bool]:
        return pulumi.get(self, "allow_unverified_ssl_cert")

    @property
    @pulumi.getter(name="alternateGuestName")
    def alternate_guest_name(self) -> str:
        """
        An alternate guest operating system name.
        """
        return pulumi.get(self, "alternate_guest_name")

    @property
    @pulumi.getter
    def annotation(self) -> str:
        """
        A description of the virtual machine.
        """
        return pulumi.get(self, "annotation")

    @property
    @pulumi.getter(name="cpuHotAddEnabled")
    def cpu_hot_add_enabled(self) -> bool:
        """
        Allow CPUs to be added to the virtual machine while
        powered on.
        """
        return pulumi.get(self, "cpu_hot_add_enabled")

    @property
    @pulumi.getter(name="cpuHotRemoveEnabled")
    def cpu_hot_remove_enabled(self) -> bool:
        """
        Allow CPUs to be removed from the virtual machine
        while powered on.
        """
        return pulumi.get(self, "cpu_hot_remove_enabled")

    @property
    @pulumi.getter(name="cpuPerformanceCountersEnabled")
    def cpu_performance_counters_enabled(self) -> bool:
        return pulumi.get(self, "cpu_performance_counters_enabled")

    @property
    @pulumi.getter(name="datastoreId")
    def datastore_id(self) -> Optional[str]:
        return pulumi.get(self, "datastore_id")

    @property
    @pulumi.getter(name="deploymentOption")
    def deployment_option(self) -> Optional[str]:
        return pulumi.get(self, "deployment_option")

    @property
    @pulumi.getter(name="diskProvisioning")
    def disk_provisioning(self) -> Optional[str]:
        return pulumi.get(self, "disk_provisioning")

    @property
    @pulumi.getter(name="enableHiddenProperties")
    def enable_hidden_properties(self) -> Optional[bool]:
        return pulumi.get(self, "enable_hidden_properties")

    @property
    @pulumi.getter
    def firmware(self) -> str:
        """
        The firmware to use on the virtual machine.
        """
        return pulumi.get(self, "firmware")

    @property
    @pulumi.getter
    def folder(self) -> Optional[str]:
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="guestId")
    def guest_id(self) -> str:
        """
        The ID for the guest operating system
        """
        return pulumi.get(self, "guest_id")

    @property
    @pulumi.getter(name="hostSystemId")
    def host_system_id(self) -> str:
        return pulumi.get(self, "host_system_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ideControllerCount")
    def ide_controller_count(self) -> int:
        return pulumi.get(self, "ide_controller_count")

    @property
    @pulumi.getter(name="ipAllocationPolicy")
    def ip_allocation_policy(self) -> Optional[str]:
        return pulumi.get(self, "ip_allocation_policy")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[str]:
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="localOvfPath")
    def local_ovf_path(self) -> Optional[str]:
        return pulumi.get(self, "local_ovf_path")

    @property
    @pulumi.getter
    def memory(self) -> int:
        """
        The size of the virtual machine memory, in MB.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="memoryHotAddEnabled")
    def memory_hot_add_enabled(self) -> bool:
        """
        Allow memory to be added to the virtual machine
        while powered on.
        """
        return pulumi.get(self, "memory_hot_add_enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nestedHvEnabled")
    def nested_hv_enabled(self) -> bool:
        """
        Enable nested hardware virtualization on the virtual
        machine, facilitating nested virtualization in the guest.
        """
        return pulumi.get(self, "nested_hv_enabled")

    @property
    @pulumi.getter(name="numCoresPerSocket")
    def num_cores_per_socket(self) -> int:
        """
        The number of cores per virtual CPU in the virtual
        machine.
        """
        return pulumi.get(self, "num_cores_per_socket")

    @property
    @pulumi.getter(name="numCpus")
    def num_cpus(self) -> int:
        """
        The number of virtual CPUs to assign to the virtual machine.
        """
        return pulumi.get(self, "num_cpus")

    @property
    @pulumi.getter(name="ovfNetworkMap")
    def ovf_network_map(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "ovf_network_map")

    @property
    @pulumi.getter(name="remoteOvfUrl")
    def remote_ovf_url(self) -> Optional[str]:
        return pulumi.get(self, "remote_ovf_url")

    @property
    @pulumi.getter(name="resourcePoolId")
    def resource_pool_id(self) -> str:
        return pulumi.get(self, "resource_pool_id")

    @property
    @pulumi.getter(name="sataControllerCount")
    def sata_controller_count(self) -> int:
        return pulumi.get(self, "sata_controller_count")

    @property
    @pulumi.getter(name="scsiControllerCount")
    def scsi_controller_count(self) -> int:
        return pulumi.get(self, "scsi_controller_count")

    @property
    @pulumi.getter(name="scsiType")
    def scsi_type(self) -> str:
        return pulumi.get(self, "scsi_type")

    @property
    @pulumi.getter(name="swapPlacementPolicy")
    def swap_placement_policy(self) -> str:
        """
        The swap file placement policy for the virtual
        machine.
        """
        return pulumi.get(self, "swap_placement_policy")


class AwaitableGetOvfVmTemplateResult(GetOvfVmTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOvfVmTemplateResult(
            allow_unverified_ssl_cert=self.allow_unverified_ssl_cert,
            alternate_guest_name=self.alternate_guest_name,
            annotation=self.annotation,
            cpu_hot_add_enabled=self.cpu_hot_add_enabled,
            cpu_hot_remove_enabled=self.cpu_hot_remove_enabled,
            cpu_performance_counters_enabled=self.cpu_performance_counters_enabled,
            datastore_id=self.datastore_id,
            deployment_option=self.deployment_option,
            disk_provisioning=self.disk_provisioning,
            enable_hidden_properties=self.enable_hidden_properties,
            firmware=self.firmware,
            folder=self.folder,
            guest_id=self.guest_id,
            host_system_id=self.host_system_id,
            id=self.id,
            ide_controller_count=self.ide_controller_count,
            ip_allocation_policy=self.ip_allocation_policy,
            ip_protocol=self.ip_protocol,
            local_ovf_path=self.local_ovf_path,
            memory=self.memory,
            memory_hot_add_enabled=self.memory_hot_add_enabled,
            name=self.name,
            nested_hv_enabled=self.nested_hv_enabled,
            num_cores_per_socket=self.num_cores_per_socket,
            num_cpus=self.num_cpus,
            ovf_network_map=self.ovf_network_map,
            remote_ovf_url=self.remote_ovf_url,
            resource_pool_id=self.resource_pool_id,
            sata_controller_count=self.sata_controller_count,
            scsi_controller_count=self.scsi_controller_count,
            scsi_type=self.scsi_type,
            swap_placement_policy=self.swap_placement_policy)


def get_ovf_vm_template(allow_unverified_ssl_cert: Optional[bool] = None,
                        datastore_id: Optional[str] = None,
                        deployment_option: Optional[str] = None,
                        disk_provisioning: Optional[str] = None,
                        enable_hidden_properties: Optional[bool] = None,
                        folder: Optional[str] = None,
                        host_system_id: Optional[str] = None,
                        ip_allocation_policy: Optional[str] = None,
                        ip_protocol: Optional[str] = None,
                        local_ovf_path: Optional[str] = None,
                        name: Optional[str] = None,
                        ovf_network_map: Optional[Mapping[str, str]] = None,
                        remote_ovf_url: Optional[str] = None,
                        resource_pool_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOvfVmTemplateResult:
    """
    The `get_ovf_vm_template` data source can be used to submit an OVF to
    vSphere and extract its hardware settings in a form that can be then used as
    inputs for a `VirtualMachine` resource.


    :param bool allow_unverified_ssl_cert: Allow unverified SSL certificates
           when deploying OVF/OVA from a URL.
    :param str datastore_id: The ID of the virtual machine's datastore. The
           virtual machine configuration is placed here, along with any virtual disks
           that are created without datastores.
    :param str deployment_option: The key of the chosen deployment option. If
           empty, the default option is chosen.
    :param str disk_provisioning: The disk provisioning type. If set, all the
           disks in the deployed OVA/OVF will have the same specified disk type. Can be
           one of `thin`, `flat`, `thick` or `sameAsSource`.
    :param bool enable_hidden_properties: Allow properties with
           `ovf:userConfigurable=false` to be set.
    :param str folder: The name of the folder in which to place the virtual
           machine.
    :param str host_system_id: The ID of the ESXi host system to deploy the
           virtual machine.
    :param str ip_allocation_policy: The IP allocation policy.
    :param str ip_protocol: The IP protocol.
    :param str local_ovf_path: The absolute path to the OVF/OVA file on the
           local system. When deploying from an OVF, ensure all necessary files such as
           the `.vmdk` files are present in the same directory as the OVF.
    :param str name: Name of the virtual machine to create.
    :param Mapping[str, str] ovf_network_map: The mapping of name of network identifiers
           from the OVF descriptor to network UUID in the environment.
    :param str remote_ovf_url: URL of the remote OVF/OVA file to be deployed.
           
           > **NOTE:** Either `local_ovf_path` or `remote_ovf_url` is required, both can
           not be empty.
    :param str resource_pool_id: The ID of a resource pool in which to place
           the virtual machine.
    """
    __args__ = dict()
    __args__['allowUnverifiedSslCert'] = allow_unverified_ssl_cert
    __args__['datastoreId'] = datastore_id
    __args__['deploymentOption'] = deployment_option
    __args__['diskProvisioning'] = disk_provisioning
    __args__['enableHiddenProperties'] = enable_hidden_properties
    __args__['folder'] = folder
    __args__['hostSystemId'] = host_system_id
    __args__['ipAllocationPolicy'] = ip_allocation_policy
    __args__['ipProtocol'] = ip_protocol
    __args__['localOvfPath'] = local_ovf_path
    __args__['name'] = name
    __args__['ovfNetworkMap'] = ovf_network_map
    __args__['remoteOvfUrl'] = remote_ovf_url
    __args__['resourcePoolId'] = resource_pool_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getOvfVmTemplate:getOvfVmTemplate', __args__, opts=opts, typ=GetOvfVmTemplateResult).value

    return AwaitableGetOvfVmTemplateResult(
        allow_unverified_ssl_cert=pulumi.get(__ret__, 'allow_unverified_ssl_cert'),
        alternate_guest_name=pulumi.get(__ret__, 'alternate_guest_name'),
        annotation=pulumi.get(__ret__, 'annotation'),
        cpu_hot_add_enabled=pulumi.get(__ret__, 'cpu_hot_add_enabled'),
        cpu_hot_remove_enabled=pulumi.get(__ret__, 'cpu_hot_remove_enabled'),
        cpu_performance_counters_enabled=pulumi.get(__ret__, 'cpu_performance_counters_enabled'),
        datastore_id=pulumi.get(__ret__, 'datastore_id'),
        deployment_option=pulumi.get(__ret__, 'deployment_option'),
        disk_provisioning=pulumi.get(__ret__, 'disk_provisioning'),
        enable_hidden_properties=pulumi.get(__ret__, 'enable_hidden_properties'),
        firmware=pulumi.get(__ret__, 'firmware'),
        folder=pulumi.get(__ret__, 'folder'),
        guest_id=pulumi.get(__ret__, 'guest_id'),
        host_system_id=pulumi.get(__ret__, 'host_system_id'),
        id=pulumi.get(__ret__, 'id'),
        ide_controller_count=pulumi.get(__ret__, 'ide_controller_count'),
        ip_allocation_policy=pulumi.get(__ret__, 'ip_allocation_policy'),
        ip_protocol=pulumi.get(__ret__, 'ip_protocol'),
        local_ovf_path=pulumi.get(__ret__, 'local_ovf_path'),
        memory=pulumi.get(__ret__, 'memory'),
        memory_hot_add_enabled=pulumi.get(__ret__, 'memory_hot_add_enabled'),
        name=pulumi.get(__ret__, 'name'),
        nested_hv_enabled=pulumi.get(__ret__, 'nested_hv_enabled'),
        num_cores_per_socket=pulumi.get(__ret__, 'num_cores_per_socket'),
        num_cpus=pulumi.get(__ret__, 'num_cpus'),
        ovf_network_map=pulumi.get(__ret__, 'ovf_network_map'),
        remote_ovf_url=pulumi.get(__ret__, 'remote_ovf_url'),
        resource_pool_id=pulumi.get(__ret__, 'resource_pool_id'),
        sata_controller_count=pulumi.get(__ret__, 'sata_controller_count'),
        scsi_controller_count=pulumi.get(__ret__, 'scsi_controller_count'),
        scsi_type=pulumi.get(__ret__, 'scsi_type'),
        swap_placement_policy=pulumi.get(__ret__, 'swap_placement_policy'))


@_utilities.lift_output_func(get_ovf_vm_template)
def get_ovf_vm_template_output(allow_unverified_ssl_cert: Optional[pulumi.Input[Optional[bool]]] = None,
                               datastore_id: Optional[pulumi.Input[Optional[str]]] = None,
                               deployment_option: Optional[pulumi.Input[Optional[str]]] = None,
                               disk_provisioning: Optional[pulumi.Input[Optional[str]]] = None,
                               enable_hidden_properties: Optional[pulumi.Input[Optional[bool]]] = None,
                               folder: Optional[pulumi.Input[Optional[str]]] = None,
                               host_system_id: Optional[pulumi.Input[str]] = None,
                               ip_allocation_policy: Optional[pulumi.Input[Optional[str]]] = None,
                               ip_protocol: Optional[pulumi.Input[Optional[str]]] = None,
                               local_ovf_path: Optional[pulumi.Input[Optional[str]]] = None,
                               name: Optional[pulumi.Input[str]] = None,
                               ovf_network_map: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                               remote_ovf_url: Optional[pulumi.Input[Optional[str]]] = None,
                               resource_pool_id: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOvfVmTemplateResult]:
    """
    The `get_ovf_vm_template` data source can be used to submit an OVF to
    vSphere and extract its hardware settings in a form that can be then used as
    inputs for a `VirtualMachine` resource.


    :param bool allow_unverified_ssl_cert: Allow unverified SSL certificates
           when deploying OVF/OVA from a URL.
    :param str datastore_id: The ID of the virtual machine's datastore. The
           virtual machine configuration is placed here, along with any virtual disks
           that are created without datastores.
    :param str deployment_option: The key of the chosen deployment option. If
           empty, the default option is chosen.
    :param str disk_provisioning: The disk provisioning type. If set, all the
           disks in the deployed OVA/OVF will have the same specified disk type. Can be
           one of `thin`, `flat`, `thick` or `sameAsSource`.
    :param bool enable_hidden_properties: Allow properties with
           `ovf:userConfigurable=false` to be set.
    :param str folder: The name of the folder in which to place the virtual
           machine.
    :param str host_system_id: The ID of the ESXi host system to deploy the
           virtual machine.
    :param str ip_allocation_policy: The IP allocation policy.
    :param str ip_protocol: The IP protocol.
    :param str local_ovf_path: The absolute path to the OVF/OVA file on the
           local system. When deploying from an OVF, ensure all necessary files such as
           the `.vmdk` files are present in the same directory as the OVF.
    :param str name: Name of the virtual machine to create.
    :param Mapping[str, str] ovf_network_map: The mapping of name of network identifiers
           from the OVF descriptor to network UUID in the environment.
    :param str remote_ovf_url: URL of the remote OVF/OVA file to be deployed.
           
           > **NOTE:** Either `local_ovf_path` or `remote_ovf_url` is required, both can
           not be empty.
    :param str resource_pool_id: The ID of a resource pool in which to place
           the virtual machine.
    """
    ...
