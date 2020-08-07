# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import _utilities, _tables


class GetDatastoreClusterResult:
    """
    A collection of values returned by getDatastoreCluster.
    """
    def __init__(__self__, datacenter_id=None, id=None, name=None):
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


class AwaitableGetDatastoreClusterResult(GetDatastoreClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatastoreClusterResult(
            datacenter_id=self.datacenter_id,
            id=self.id,
            name=self.name)


def get_datastore_cluster(datacenter_id=None, name=None, opts=None):
    """
    The `DatastoreCluster` data source can be used to discover the ID of a
    datastore cluster in vSphere. This is useful to fetch the ID of a datastore
    cluster that you want to use to assign datastores to using the
    `NasDatastore` or
    `VmfsDatastore` resources, or create
    virtual machines in using the
    `VirtualMachine` resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vsphere as vsphere

    datacenter = vsphere.get_datacenter(name="dc1")
    datastore_cluster = vsphere.get_datastore_cluster(datacenter_id=data["vsphere_datacenter"]["dc"]["id"],
        name="datastore-cluster1")
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
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vsphere:index/getDatastoreCluster:getDatastoreCluster', __args__, opts=opts).value

    return AwaitableGetDatastoreClusterResult(
        datacenter_id=__ret__.get('datacenterId'),
        id=__ret__.get('id'),
        name=__ret__.get('name'))
