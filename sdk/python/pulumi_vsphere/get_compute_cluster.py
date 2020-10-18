# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = [
    'GetComputeClusterResult',
    'AwaitableGetComputeClusterResult',
    'get_compute_cluster',
]

@pulumi.output_type
class GetComputeClusterResult:
    """
    A collection of values returned by getComputeCluster.
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
    @pulumi.getter(name="resourcePoolId")
    def resource_pool_id(self) -> str:
        return pulumi.get(self, "resource_pool_id")


class AwaitableGetComputeClusterResult(GetComputeClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetComputeClusterResult(
            datacenter_id=self.datacenter_id,
            id=self.id,
            name=self.name,
            resource_pool_id=self.resource_pool_id)


def get_compute_cluster(datacenter_id: Optional[str] = None,
                        name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetComputeClusterResult:
    """
    The `ComputeCluster` data source can be used to discover the ID of a
    cluster in vSphere. This is useful to fetch the ID of a cluster that you want
    to use for virtual machine placement via the
    `VirtualMachine` resource, allowing
    you to specify the cluster's root resource pool directly versus using the alias
    available through the `ResourcePool`
    data source.

    > You may also wish to see the
    `ComputeCluster` resource for further
    details about clusters or how to work with them.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc1")
    compute_cluster = vsphere.get_compute_cluster(datacenter_id=data["vsphere_datacenter"]["dc"]["id"],
        name="compute-cluster1")
    ```


    :param str datacenter_id: The managed object reference
           ID of the datacenter the cluster is located in.  This can
           be omitted if the search path used in `name` is an absolute path.  For
           default datacenters, use the id attribute from an empty `Datacenter`
           data source.
    :param str name: The name or absolute path to the cluster.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getComputeCluster:getComputeCluster', __args__, opts=opts, typ=GetComputeClusterResult).value

    return AwaitableGetComputeClusterResult(
        datacenter_id=__ret__.datacenter_id,
        id=__ret__.id,
        name=__ret__.name,
        resource_pool_id=__ret__.resource_pool_id)
