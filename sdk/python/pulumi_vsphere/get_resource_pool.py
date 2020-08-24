# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'GetResourcePoolResult',
    'AwaitableGetResourcePoolResult',
    'get_resource_pool',
]

@pulumi.output_type
class GetResourcePoolResult:
    """
    A collection of values returned by getResourcePool.
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
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")


class AwaitableGetResourcePoolResult(GetResourcePoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourcePoolResult(
            datacenter_id=self.datacenter_id,
            id=self.id,
            name=self.name)


def get_resource_pool(datacenter_id: Optional[str] = None,
                      name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourcePoolResult:
    """
    The `ResourcePool` data source can be used to discover the ID of a
    resource pool in vSphere. This is useful to fetch the ID of a resource pool
    that you want to use to create virtual machines in using the
    `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc1")
    pool = vsphere.get_resource_pool(datacenter_id=datacenter.id,
        name="resource-pool-1")
    ```
    ### Specifying the root resource pool for a standalone host

    > **NOTE:** Fetching the root resource pool for a cluster can now be done
    directly via the `ComputeCluster`
    data source.

    All compute resources in vSphere (clusters, standalone hosts, and standalone
    ESXi) have a resource pool, even if one has not been explicitly created. This
    resource pool is referred to as the _root resource pool_ and can be looked up
    by specifying the path as per the example below:

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    pool = vsphere.get_resource_pool(datacenter_id=data["vsphere_datacenter"]["dc"]["id"],
        name="esxi1/Resources")
    ```

    For more information on the root resource pool, see [Managing Resource
    Pools][vmware-docs-resource-pools] in the vSphere documentation.

    [vmware-docs-resource-pools]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-60077B40-66FF-4625-934A-641703ED7601.html


    :param str datacenter_id: The managed object reference
           ID of the datacenter the resource pool is located in.
           This can be omitted if the search path used in `name` is an absolute path.
           For default datacenters, use the id attribute from an empty
           `Datacenter` data source.
    :param str name: The name of the resource pool. This can be a name or
           path. This is required when using vCenter.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getResourcePool:getResourcePool', __args__, opts=opts, typ=GetResourcePoolResult).value

    return AwaitableGetResourcePoolResult(
        datacenter_id=__ret__.datacenter_id,
        id=__ret__.id,
        name=__ret__.name)
