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
    'GetDatacenterResult',
    'AwaitableGetDatacenterResult',
    'get_datacenter',
    'get_datacenter_output',
]

@pulumi.output_type
class GetDatacenterResult:
    """
    A collection of values returned by getDatacenter.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

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


class AwaitableGetDatacenterResult(GetDatacenterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatacenterResult(
            id=self.id,
            name=self.name)


def get_datacenter(name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatacenterResult:
    """
    The `Datacenter` data source can be used to discover the ID of a
    vSphere datacenter object. This can then be used with resources or data sources
    that require a datacenter, such as the `Host`
    data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    ```


    :param str name: The name of the datacenter. This can be a name or path.
           Can be omitted if there is only one datacenter in the inventory.
           
           > **NOTE:** When used with an ESXi host, this data source *always* returns the
           host's "default" datacenter, which is a special datacenter name unrelated to the
           datacenters that exist in the vSphere inventory when managed by a vCenter Server
           instance. Hence, the `name` attribute is completely ignored.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDatacenter:getDatacenter', __args__, opts=opts, typ=GetDatacenterResult).value

    return AwaitableGetDatacenterResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_datacenter)
def get_datacenter_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatacenterResult]:
    """
    The `Datacenter` data source can be used to discover the ID of a
    vSphere datacenter object. This can then be used with resources or data sources
    that require a datacenter, such as the `Host`
    data source.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    ```


    :param str name: The name of the datacenter. This can be a name or path.
           Can be omitted if there is only one datacenter in the inventory.
           
           > **NOTE:** When used with an ESXi host, this data source *always* returns the
           host's "default" datacenter, which is a special datacenter name unrelated to the
           datacenters that exist in the vSphere inventory when managed by a vCenter Server
           instance. Hence, the `name` attribute is completely ignored.
    """
    ...
