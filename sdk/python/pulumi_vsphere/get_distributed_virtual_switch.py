# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetDistributedVirtualSwitchResult:
    """
    A collection of values returned by getDistributedVirtualSwitch.
    """
    def __init__(__self__, datacenter_id=None, name=None, uplinks=None, id=None):
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        __self__.datacenter_id = datacenter_id
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if uplinks and not isinstance(uplinks, list):
            raise TypeError("Expected argument 'uplinks' to be a list")
        __self__.uplinks = uplinks
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return self

    __iter__ = __await__

def get_distributed_virtual_switch(datacenter_id=None,name=None,opts=None):
    """
    The `vsphere_distributed_virtual_switch` data source can be used to discover
    the ID and uplink data of a of a vSphere distributed virtual switch (DVS). This
    can then be used with resources or data sources that require a DVS, such as the
    [`vsphere_distributed_port_group`][distributed-port-group] resource, for which
    an example is shown below.
    
    [distributed-port-group]: /docs/providers/vsphere/r/distributed_port_group.html
    
    > **NOTE:** This data source requires vCenter and is not available on direct
    ESXi connections.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/d/distributed_virtual_switch.html.markdown.
    """
    __args__ = dict()

    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch', __args__, opts=opts).value

    return GetDistributedVirtualSwitchResult(
        datacenter_id=__ret__.get('datacenterId'),
        name=__ret__.get('name'),
        uplinks=__ret__.get('uplinks'),
        id=__ret__.get('id'))
