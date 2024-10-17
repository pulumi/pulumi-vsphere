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
    'GetDatastoreClusterResult',
    'AwaitableGetDatastoreClusterResult',
    'get_datastore_cluster',
    'get_datastore_cluster_output',
]

@pulumi.output_type
class GetDatastoreClusterResult:
    """
    A collection of values returned by getDatastoreCluster.
    """
    def __init__(__self__, datacenter_id=None, datastores=None, id=None, name=None):
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if datastores and not isinstance(datastores, list):
            raise TypeError("Expected argument 'datastores' to be a list")
        pulumi.set(__self__, "datastores", datastores)
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
    def datastores(self) -> Sequence[str]:
        """
        (Optional) The names of the datastores included in the specific 
        cluster.
        """
        return pulumi.get(self, "datastores")

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


class AwaitableGetDatastoreClusterResult(GetDatastoreClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatastoreClusterResult(
            datacenter_id=self.datacenter_id,
            datastores=self.datastores,
            id=self.id,
            name=self.name)


def get_datastore_cluster(datacenter_id: Optional[str] = None,
                          name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatastoreClusterResult:
    """
    The `DatastoreCluster` data source can be used to discover the ID of a
    vSphere datastore cluster object. This can then be used with resources or data sources
    that require a datastore. For example, to assign datastores using the
    `NasDatastore` or `VmfsDatastore` resources, or to create virtual machines in using the `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    datastore_cluster = vsphere.get_datastore_cluster(name="datastore-cluster-01",
        datacenter_id=datacenter.id)
    ```


    :param str datacenter_id: The managed object reference
           ID of the datacenter the datastore cluster is located in.
           This can be omitted if the search path used in `name` is an absolute path.
           For default datacenters, use the id attribute from an empty
           `Datacenter` data source.
    :param str name: The name or absolute path to the datastore cluster.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDatastoreCluster:getDatastoreCluster', __args__, opts=opts, typ=GetDatastoreClusterResult).value

    return AwaitableGetDatastoreClusterResult(
        datacenter_id=pulumi.get(__ret__, 'datacenter_id'),
        datastores=pulumi.get(__ret__, 'datastores'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_datastore_cluster_output(datacenter_id: Optional[pulumi.Input[Optional[str]]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatastoreClusterResult]:
    """
    The `DatastoreCluster` data source can be used to discover the ID of a
    vSphere datastore cluster object. This can then be used with resources or data sources
    that require a datastore. For example, to assign datastores using the
    `NasDatastore` or `VmfsDatastore` resources, or to create virtual machines in using the `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    datastore_cluster = vsphere.get_datastore_cluster(name="datastore-cluster-01",
        datacenter_id=datacenter.id)
    ```


    :param str datacenter_id: The managed object reference
           ID of the datacenter the datastore cluster is located in.
           This can be omitted if the search path used in `name` is an absolute path.
           For default datacenters, use the id attribute from an empty
           `Datacenter` data source.
    :param str name: The name or absolute path to the datastore cluster.
    """
    __args__ = dict()
    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vsphere:index/getDatastoreCluster:getDatastoreCluster', __args__, opts=opts, typ=GetDatastoreClusterResult)
    return __ret__.apply(lambda __response__: GetDatastoreClusterResult(
        datacenter_id=pulumi.get(__response__, 'datacenter_id'),
        datastores=pulumi.get(__response__, 'datastores'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))
