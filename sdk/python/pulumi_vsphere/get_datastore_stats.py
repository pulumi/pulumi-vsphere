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
    'GetDatastoreStatsResult',
    'AwaitableGetDatastoreStatsResult',
    'get_datastore_stats',
    'get_datastore_stats_output',
]

@pulumi.output_type
class GetDatastoreStatsResult:
    """
    A collection of values returned by getDatastoreStats.
    """
    def __init__(__self__, capacity=None, datacenter_id=None, free_space=None, id=None):
        if capacity and not isinstance(capacity, dict):
            raise TypeError("Expected argument 'capacity' to be a dict")
        pulumi.set(__self__, "capacity", capacity)
        if datacenter_id and not isinstance(datacenter_id, str):
            raise TypeError("Expected argument 'datacenter_id' to be a str")
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        if free_space and not isinstance(free_space, dict):
            raise TypeError("Expected argument 'free_space' to be a dict")
        pulumi.set(__self__, "free_space", free_space)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of the capacity for all datastore in the datacenter,
        where the name of the datastore is used as key and the capacity as value.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> str:
        """
        The [managed object reference ID][docs-about-morefs] of the
        datacenter the datastores are located in.
        """
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter(name="freeSpace")
    def free_space(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of the free space for each datastore in the
        datacenter, where the name of the datastore is used as key and the free space
        as value.
        """
        return pulumi.get(self, "free_space")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetDatastoreStatsResult(GetDatastoreStatsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatastoreStatsResult(
            capacity=self.capacity,
            datacenter_id=self.datacenter_id,
            free_space=self.free_space,
            id=self.id)


def get_datastore_stats(capacity: Optional[Mapping[str, str]] = None,
                        datacenter_id: Optional[str] = None,
                        free_space: Optional[Mapping[str, str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatastoreStatsResult:
    """
    The `get_datastore_stats` data source can be used to retrieve the usage
    stats of all vSphere datastore objects in a datacenter. This can then be used as
    a standalone data source to get information required as input to other data
    sources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    datastore_stats = vsphere.get_datastore_stats(datacenter_id=datacenter.id)
    ```

    A useful example of this data source would be to determine the datastore with
    the most free space. For example, in addition to the above:

    Create an `outputs.tf` like that:

    ```python
    import pulumi

    pulumi.export("maxFreeSpaceName", their_max_free_space_name)
    pulumi.export("maxFreeSpace", their_max_free_space)
    ```

    and a `locals.tf` like that:


    :param Mapping[str, str] capacity: A mapping of the capacity for all datastore in the datacenter,
           where the name of the datastore is used as key and the capacity as value.
    :param str datacenter_id: The
           [managed object reference ID][docs-about-morefs] of the datacenter the
           datastores are located in. For default datacenters, use the `id` attribute
           from an empty `Datacenter` data source.
           
           [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
    :param Mapping[str, str] free_space: A mapping of the free space for each datastore in the
           datacenter, where the name of the datastore is used as key and the free space
           as value.
    """
    __args__ = dict()
    __args__['capacity'] = capacity
    __args__['datacenterId'] = datacenter_id
    __args__['freeSpace'] = free_space
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDatastoreStats:getDatastoreStats', __args__, opts=opts, typ=GetDatastoreStatsResult).value

    return AwaitableGetDatastoreStatsResult(
        capacity=pulumi.get(__ret__, 'capacity'),
        datacenter_id=pulumi.get(__ret__, 'datacenter_id'),
        free_space=pulumi.get(__ret__, 'free_space'),
        id=pulumi.get(__ret__, 'id'))
def get_datastore_stats_output(capacity: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                               datacenter_id: Optional[pulumi.Input[str]] = None,
                               free_space: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatastoreStatsResult]:
    """
    The `get_datastore_stats` data source can be used to retrieve the usage
    stats of all vSphere datastore objects in a datacenter. This can then be used as
    a standalone data source to get information required as input to other data
    sources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc-01")
    datastore_stats = vsphere.get_datastore_stats(datacenter_id=datacenter.id)
    ```

    A useful example of this data source would be to determine the datastore with
    the most free space. For example, in addition to the above:

    Create an `outputs.tf` like that:

    ```python
    import pulumi

    pulumi.export("maxFreeSpaceName", their_max_free_space_name)
    pulumi.export("maxFreeSpace", their_max_free_space)
    ```

    and a `locals.tf` like that:


    :param Mapping[str, str] capacity: A mapping of the capacity for all datastore in the datacenter,
           where the name of the datastore is used as key and the capacity as value.
    :param str datacenter_id: The
           [managed object reference ID][docs-about-morefs] of the datacenter the
           datastores are located in. For default datacenters, use the `id` attribute
           from an empty `Datacenter` data source.
           
           [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider
    :param Mapping[str, str] free_space: A mapping of the free space for each datastore in the
           datacenter, where the name of the datastore is used as key and the free space
           as value.
    """
    __args__ = dict()
    __args__['capacity'] = capacity
    __args__['datacenterId'] = datacenter_id
    __args__['freeSpace'] = free_space
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vsphere:index/getDatastoreStats:getDatastoreStats', __args__, opts=opts, typ=GetDatastoreStatsResult)
    return __ret__.apply(lambda __response__: GetDatastoreStatsResult(
        capacity=pulumi.get(__response__, 'capacity'),
        datacenter_id=pulumi.get(__response__, 'datacenter_id'),
        free_space=pulumi.get(__response__, 'free_space'),
        id=pulumi.get(__response__, 'id')))
