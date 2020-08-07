# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import _utilities, _tables


class GetDistributedVirtualSwitchResult:
    """
    A collection of values returned by getDistributedVirtualSwitch.
    """
    def __init__(__self__, datacenter_id=None, id=None, name=None, uplinks=None):
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        __self__.datacenter_id = datacenter_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if uplinks and not isinstance(uplinks, list):
            raise TypeError("Expected argument 'uplinks' to be a list")
        __self__.uplinks = uplinks


class AwaitableGetDistributedVirtualSwitchResult(GetDistributedVirtualSwitchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDistributedVirtualSwitchResult(
            datacenter_id=self.datacenter_id,
            id=self.id,
            name=self.name,
            uplinks=self.uplinks)


def get_distributed_virtual_switch(datacenter_id=None, name=None, opts=None):
    """
    The `DistributedVirtualSwitch` data source can be used to discover
    the ID and uplink data of a of a vSphere distributed virtual switch (DVS). This
    can then be used with resources or data sources that require a DVS, such as the
    `DistributedPortGroup` resource, for which
    an example is shown below.

    > **NOTE:** This data source requires vCenter and is not available on direct
    ESXi connections.

    ## Example Usage

    The following example locates a DVS that is named `test-dvs`, in the
    datacenter `dc1`. It then uses this DVS to set up a
    `DistributedPortGroup` resource that uses the first uplink as a
    primary uplink and the second uplink as a secondary.

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc1")
    dvs = vsphere.get_distributed_virtual_switch(datacenter_id=datacenter.id,
        name="test-dvs")
    pg = vsphere.DistributedPortGroup("pg",
        active_uplinks=[dvs.uplinks[0]],
        distributed_virtual_switch_uuid=dvs.id,
        standby_uplinks=[dvs.uplinks[1]])
    ```


    :param str datacenter_id: The managed object reference
           ID of the datacenter the DVS is located in. This can be
           omitted if the search path used in `name` is an absolute path. For default
           datacenters, use the id attribute from an empty `Datacenter` data
           source.
    :param str name: The name of the distributed virtual switch. This can be a
           name or path.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch', __args__, opts=opts).value

    return AwaitableGetDistributedVirtualSwitchResult(
        datacenter_id=__ret__.get('datacenterId'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        uplinks=__ret__.get('uplinks'))
