# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetNetworkResult',
    'AwaitableGetNetworkResult',
    'get_network',
]

@pulumi.output_type
class GetNetworkResult:
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, datacenter_id=None, distributed_virtual_switch_uuid=None, id=None, name=None, type=None):
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if distributed_virtual_switch_uuid and not isinstance(distributed_virtual_switch_uuid, str):
            raise TypeError("Expected argument 'distributed_virtual_switch_uuid' to be a str")
        pulumi.set(__self__, "distributed_virtual_switch_uuid", distributed_virtual_switch_uuid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> Optional[str]:
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter(name="distributedVirtualSwitchUuid")
    def distributed_virtual_switch_uuid(self) -> Optional[str]:
        return pulumi.get(self, "distributed_virtual_switch_uuid")

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

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetNetworkResult(GetNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkResult(
            datacenter_id=self.datacenter_id,
            distributed_virtual_switch_uuid=self.distributed_virtual_switch_uuid,
            id=self.id,
            name=self.name,
            type=self.type)


def get_network(datacenter_id: Optional[str] = None,
                distributed_virtual_switch_uuid: Optional[str] = None,
                name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkResult:
    """
    The `getNetwork` data source can be used to discover the ID of a network
    in vSphere. This can be any network that can be used as the backing for a
    network interface for `VirtualMachine` or any other vSphere resource
    that requires a network. This includes standard (host-based) port groups, DVS
    port groups, or opaque networks such as those managed by NSX.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc1")
    net = vsphere.get_network(datacenter_id=datacenter.id,
        name="test-net")
    ```


    :param str datacenter_id: The managed object reference
           ID of the datacenter the network is located in. This can
           be omitted if the search path used in `name` is an absolute path. For default
           datacenters, use the id attribute from an empty `Datacenter` data
           source.
    :param str distributed_virtual_switch_uuid: For distributed port group type 
           network objects, the ID of the distributed virtual switch the given port group
           belongs to. It is useful to differentiate port groups with same name using the
           Distributed virtual switch ID.
    :param str name: The name of the network. This can be a name or path.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['distributedVirtualSwitchUuid'] = distributed_virtual_switch_uuid
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getNetwork:getNetwork', __args__, opts=opts, typ=GetNetworkResult).value

    return AwaitableGetNetworkResult(
        datacenter_id=__ret__.datacenter_id,
        distributed_virtual_switch_uuid=__ret__.distributed_virtual_switch_uuid,
        id=__ret__.id,
        name=__ret__.name,
        type=__ret__.type)
