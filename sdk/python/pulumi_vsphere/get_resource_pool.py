# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'GetResourcePoolResult',
    'AwaitableGetResourcePoolResult',
    'get_resource_pool',
    'get_resource_pool_output',
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
    resource pool in vSphere. This is useful to return the ID of a resource pool
    that you want to use to create virtual machines in using the
    `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    pool = vsphere.get_resource_pool(name="resource-pool-01",
        datacenter_id=datacenter.id)
    ```

    ### Specifying the Root Resource Pool for a Standalone ESXi Host

    > **NOTE:** Returning the root resource pool for a cluster can be done
    directly via the `ComputeCluster`
    data source.

    All compute resources in vSphere have a resource pool, even if one has not been
    explicitly created. This resource pool is referred to as the _root resource
    pool_ and can be looked up by specifying the path.

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    pool = vsphere.get_resource_pool(name="esxi-01.example.com/Resources",
        datacenter_id=datacenter["id"])
    ```

    For more information on the root resource pool, see
    [Managing Resource Pools][vmware-docs-resource-pools] in the vSphere
    documentation.

    [vmware-docs-resource-pools]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-60077B40-66FF-4625-934A-641703ED7601.html


    :param str datacenter_id: The managed object reference ID
           of the datacenter in which the resource pool is located. This can be omitted
           if the search path used in `name` is an absolute path. For default
           datacenters, use the id attribute from an empty `Datacenter` data
           source.
           
           > **Note:** When using ESXi without a vCenter Server instance, you do not
           need to specify either attribute to use this data source. An empty declaration
           will load the ESXi host's root resource pool.
    :param str name: The name of the resource pool. This can be a name or
           path. This is required when using vCenter.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getResourcePool:getResourcePool', __args__, opts=opts, typ=GetResourcePoolResult).value

    return AwaitableGetResourcePoolResult(
        datacenter_id=pulumi.get(__ret__, 'datacenter_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_resource_pool_output(datacenter_id: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResourcePoolResult]:
    """
    The `ResourcePool` data source can be used to discover the ID of a
    resource pool in vSphere. This is useful to return the ID of a resource pool
    that you want to use to create virtual machines in using the
    `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    pool = vsphere.get_resource_pool(name="resource-pool-01",
        datacenter_id=datacenter.id)
    ```

    ### Specifying the Root Resource Pool for a Standalone ESXi Host

    > **NOTE:** Returning the root resource pool for a cluster can be done
    directly via the `ComputeCluster`
    data source.

    All compute resources in vSphere have a resource pool, even if one has not been
    explicitly created. This resource pool is referred to as the _root resource
    pool_ and can be looked up by specifying the path.

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    pool = vsphere.get_resource_pool(name="esxi-01.example.com/Resources",
        datacenter_id=datacenter["id"])
    ```

    For more information on the root resource pool, see
    [Managing Resource Pools][vmware-docs-resource-pools] in the vSphere
    documentation.

    [vmware-docs-resource-pools]: https://docs.vmware.com/en/VMware-vSphere/8.0/vsphere-resource-management/GUID-60077B40-66FF-4625-934A-641703ED7601.html


    :param str datacenter_id: The managed object reference ID
           of the datacenter in which the resource pool is located. This can be omitted
           if the search path used in `name` is an absolute path. For default
           datacenters, use the id attribute from an empty `Datacenter` data
           source.
           
           > **Note:** When using ESXi without a vCenter Server instance, you do not
           need to specify either attribute to use this data source. An empty declaration
           will load the ESXi host's root resource pool.
    :param str name: The name of the resource pool. This can be a name or
           path. This is required when using vCenter.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vsphere:index/getResourcePool:getResourcePool', __args__, opts=opts, typ=GetResourcePoolResult)
    return __ret__.apply(lambda __response__: GetResourcePoolResult(
        datacenter_id=pulumi.get(__response__, 'datacenter_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))
