# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetHostResult',
    'AwaitableGetHostResult',
    'get_host',
    'get_host_output',
]

@pulumi.output_type
class GetHostResult:
    """
    A collection of values returned by getHost.
    """
    def __init__(__self__, datacenter_id=None, id=None, name=None, resource_pool_id=None):
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_pool_id and not isinstance(resource_pool_id, str):
            raise TypeError("Expected argument 'resource_pool_id' to be a str")
        pulumi.set(__self__, "resource_pool_id", resource_pool_id)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> str:
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourcePoolId")
    def resource_pool_id(self) -> str:
        """
        The managed object ID of the ESXi
        host's root resource pool.
        """
        return pulumi.get(self, "resource_pool_id")


class AwaitableGetHostResult(GetHostResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostResult(
            datacenter_id=self.datacenter_id,
            id=self.id,
            name=self.name,
            resource_pool_id=self.resource_pool_id)


def get_host(datacenter_id: Optional[str] = None,
             name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostResult:
    """
    The `Host` data source can be used to discover the ID of an ESXi host.
    This can then be used with resources or data sources that require an ESX
    host's managed object reference ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    ```


    :param str datacenter_id: The managed object reference ID
           of a vSphere datacenter object.
    :param str name: The name of the ESXI host. This can be a name or path.
           Can be omitted if there is only one host in your inventory.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getHost:getHost', __args__, opts=opts, typ=GetHostResult).value

    return AwaitableGetHostResult(
        datacenter_id=__ret__.datacenter_id,
        id=__ret__.id,
        name=__ret__.name,
        resource_pool_id=__ret__.resource_pool_id)


@_utilities.lift_output_func(get_host)
def get_host_output(datacenter_id: Optional[pulumi.Input[str]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHostResult]:
    """
    The `Host` data source can be used to discover the ID of an ESXi host.
    This can then be used with resources or data sources that require an ESX
    host's managed object reference ID.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    host = vsphere.get_host(name="esxi-01.example.com",
        datacenter_id=datacenter.id)
    ```


    :param str datacenter_id: The managed object reference ID
           of a vSphere datacenter object.
    :param str name: The name of the ESXI host. This can be a name or path.
           Can be omitted if there is only one host in your inventory.
    """
    ...
