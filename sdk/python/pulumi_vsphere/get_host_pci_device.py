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
    'GetHostPciDeviceResult',
    'AwaitableGetHostPciDeviceResult',
    'get_host_pci_device',
    'get_host_pci_device_output',
]

@pulumi.output_type
class GetHostPciDeviceResult:
    """
    A collection of values returned by getHostPciDevice.
    """
    def __init__(__self__, class_id=None, host_id=None, id=None, name=None, name_regex=None, vendor_id=None):
        if class_id and not isinstance(class_id, str):
            raise TypeError("Expected argument 'class_id' to be a str")
        pulumi.set(__self__, "class_id", class_id)
        if host_id and not isinstance(host_id, str):
            raise TypeError("Expected argument 'host_id' to be a str")
        pulumi.set(__self__, "host_id", host_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if vendor_id and not isinstance(vendor_id, str):
            raise TypeError("Expected argument 'vendor_id' to be a str")
        pulumi.set(__self__, "vendor_id", vendor_id)

    @property
    @pulumi.getter(name="classId")
    def class_id(self) -> Optional[str]:
        return pulumi.get(self, "class_id")

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> str:
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the PCI device.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="vendorId")
    def vendor_id(self) -> Optional[str]:
        return pulumi.get(self, "vendor_id")


class AwaitableGetHostPciDeviceResult(GetHostPciDeviceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostPciDeviceResult(
            class_id=self.class_id,
            host_id=self.host_id,
            id=self.id,
            name=self.name,
            name_regex=self.name_regex,
            vendor_id=self.vendor_id)


def get_host_pci_device(class_id: Optional[str] = None,
                        host_id: Optional[str] = None,
                        name_regex: Optional[str] = None,
                        vendor_id: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostPciDeviceResult:
    """
    The `get_host_pci_device` data source can be used to discover the device ID
    of a vSphere host's PCI device. This can then be used with
    `VirtualMachine`'s `pci_device_id`.

    ## Example Usage
    ### With Vendor ID And Class ID

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    dev = vsphere.get_host_pci_device(host_id=host.id,
        class_id="123",
        vendor_id="456")
    ```
    ### With Name Regular Expression


    :param str class_id: The hexadecimal PCI device class ID
           
           [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
           
           > **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
    :param str host_id: The [managed object reference ID][docs-about-morefs] of a host.
    :param str name_regex: A regular expression that will be used to match the
           host PCI device name.
    :param str vendor_id: The hexadecimal PCI device vendor ID.
    """
    __args__ = dict()
    __args__['classId'] = class_id
    __args__['hostId'] = host_id
    __args__['nameRegex'] = name_regex
    __args__['vendorId'] = vendor_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getHostPciDevice:getHostPciDevice', __args__, opts=opts, typ=GetHostPciDeviceResult).value

    return AwaitableGetHostPciDeviceResult(
        class_id=pulumi.get(__ret__, 'class_id'),
        host_id=pulumi.get(__ret__, 'host_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        vendor_id=pulumi.get(__ret__, 'vendor_id'))


@_utilities.lift_output_func(get_host_pci_device)
def get_host_pci_device_output(class_id: Optional[pulumi.Input[Optional[str]]] = None,
                               host_id: Optional[pulumi.Input[str]] = None,
                               name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                               vendor_id: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHostPciDeviceResult]:
    """
    The `get_host_pci_device` data source can be used to discover the device ID
    of a vSphere host's PCI device. This can then be used with
    `VirtualMachine`'s `pci_device_id`.

    ## Example Usage
    ### With Vendor ID And Class ID

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    dev = vsphere.get_host_pci_device(host_id=host.id,
        class_id="123",
        vendor_id="456")
    ```
    ### With Name Regular Expression


    :param str class_id: The hexadecimal PCI device class ID
           
           [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
           
           > **NOTE:** `name_regex`, `vendor_id`, and `class_id` can all be used together.
    :param str host_id: The [managed object reference ID][docs-about-morefs] of a host.
    :param str name_regex: A regular expression that will be used to match the
           host PCI device name.
    :param str vendor_id: The hexadecimal PCI device vendor ID.
    """
    ...
