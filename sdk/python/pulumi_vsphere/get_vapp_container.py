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
    'GetVappContainerResult',
    'AwaitableGetVappContainerResult',
    'get_vapp_container',
    'get_vapp_container_output',
]

@pulumi.output_type
class GetVappContainerResult:
    """
    A collection of values returned by getVappContainer.
    """
    def __init__(__self__, datacenter_id=None, id=None, name=None):
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

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
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetVappContainerResult(GetVappContainerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVappContainerResult(
            datacenter_id=self.datacenter_id,
            id=self.id,
            name=self.name)


def get_vapp_container(datacenter_id: Optional[str] = None,
                       name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVappContainerResult:
    """
    The `VappContainer` data source can be used to discover the ID of a
    vApp container in vSphere. This is useful to return the ID of a vApp container
    that you want to use to create virtual machines in using the
    `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    pool = vsphere.get_vapp_container(name="vapp-container-01",
        datacenter_id=datacenter.id)
    ```


    :param str datacenter_id: The managed object reference ID
           of the datacenter in which the vApp container is located.
    :param str name: The name of the vApp container. This can be a name or
           path.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getVappContainer:getVappContainer', __args__, opts=opts, typ=GetVappContainerResult).value

    return AwaitableGetVappContainerResult(
        datacenter_id=pulumi.get(__ret__, 'datacenter_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_vapp_container)
def get_vapp_container_output(datacenter_id: Optional[pulumi.Input[str]] = None,
                              name: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVappContainerResult]:
    """
    The `VappContainer` data source can be used to discover the ID of a
    vApp container in vSphere. This is useful to return the ID of a vApp container
    that you want to use to create virtual machines in using the
    `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    pool = vsphere.get_vapp_container(name="vapp-container-01",
        datacenter_id=datacenter.id)
    ```


    :param str datacenter_id: The managed object reference ID
           of the datacenter in which the vApp container is located.
    :param str name: The name of the vApp container. This can be a name or
           path.
    """
    ...
