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
    'GetDistributedVirtualSwitchResult',
    'AwaitableGetDistributedVirtualSwitchResult',
    'get_distributed_virtual_switch',
    'get_distributed_virtual_switch_output',
]

@pulumi.output_type
class GetDistributedVirtualSwitchResult:
    """
    A collection of values returned by getDistributedVirtualSwitch.
    """
    def __init__(__self__, datacenter_id=None, id=None, name=None, uplinks=None):
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if uplinks and not isinstance(uplinks, list):
            raise TypeError("Expected argument 'uplinks' to be a list")
        pulumi.set(__self__, "uplinks", uplinks)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> Optional[str]:
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

    @property
    @pulumi.getter
    def uplinks(self) -> Sequence[str]:
        """
        The list of the uplinks on this vSphere distributed switch, as per the
        `uplinks` argument to the
        `DistributedVirtualSwitch`
        resource.
        """
        return pulumi.get(self, "uplinks")


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


def get_distributed_virtual_switch(datacenter_id: Optional[str] = None,
                                   name: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDistributedVirtualSwitchResult:
    """
    The `DistributedVirtualSwitch` data source can be used to discover
    the ID and uplink data of a of a vSphere distributed switch (VDS). This
    can then be used with resources or data sources that require a VDS, such as the
    `DistributedPortGroup` resource, for which
    an example is shown below.

    > **NOTE:** This data source requires vCenter Server and is not available on
    direct ESXi host connections.

    ## Example Usage

    The following example locates a distributed switch named `vds-01`, in the
    datacenter `dc-01`. It then uses this distributed switch to set up a
    `DistributedPortGroup` resource that uses the first uplink as a
    primary uplink and the second uplink as a secondary.

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    vds = vsphere.get_distributed_virtual_switch(name="vds-01",
        datacenter_id=datacenter.id)
    dvpg = vsphere.DistributedPortGroup("dvpg",
        name="dvpg-01",
        distributed_virtual_switch_uuid=vds.id,
        active_uplinks=[vds.uplinks[0]],
        standby_uplinks=[vds.uplinks[1]])
    ```
    <!--End PulumiCodeChooser -->


    :param str datacenter_id: The managed object reference ID
           of the datacenter the VDS is located in. This can be omitted if the search
           path used in `name` is an absolute path. For default datacenters, use the `id`
           attribute from an empty `Datacenter` data source.
    :param str name: The name of the VDS. This can be a name or path.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDistributedVirtualSwitch:getDistributedVirtualSwitch', __args__, opts=opts, typ=GetDistributedVirtualSwitchResult).value

    return AwaitableGetDistributedVirtualSwitchResult(
        datacenter_id=pulumi.get(__ret__, 'datacenter_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        uplinks=pulumi.get(__ret__, 'uplinks'))


@_utilities.lift_output_func(get_distributed_virtual_switch)
def get_distributed_virtual_switch_output(datacenter_id: Optional[pulumi.Input[Optional[str]]] = None,
                                          name: Optional[pulumi.Input[str]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDistributedVirtualSwitchResult]:
    """
    The `DistributedVirtualSwitch` data source can be used to discover
    the ID and uplink data of a of a vSphere distributed switch (VDS). This
    can then be used with resources or data sources that require a VDS, such as the
    `DistributedPortGroup` resource, for which
    an example is shown below.

    > **NOTE:** This data source requires vCenter Server and is not available on
    direct ESXi host connections.

    ## Example Usage

    The following example locates a distributed switch named `vds-01`, in the
    datacenter `dc-01`. It then uses this distributed switch to set up a
    `DistributedPortGroup` resource that uses the first uplink as a
    primary uplink and the second uplink as a secondary.

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    vds = vsphere.get_distributed_virtual_switch(name="vds-01",
        datacenter_id=datacenter.id)
    dvpg = vsphere.DistributedPortGroup("dvpg",
        name="dvpg-01",
        distributed_virtual_switch_uuid=vds.id,
        active_uplinks=[vds.uplinks[0]],
        standby_uplinks=[vds.uplinks[1]])
    ```
    <!--End PulumiCodeChooser -->


    :param str datacenter_id: The managed object reference ID
           of the datacenter the VDS is located in. This can be omitted if the search
           path used in `name` is an absolute path. For default datacenters, use the `id`
           attribute from an empty `Datacenter` data source.
    :param str name: The name of the VDS. This can be a name or path.
    """
    ...
