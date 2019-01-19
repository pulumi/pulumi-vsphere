# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetComputeClusterResult(object):
    """
    A collection of values returned by getComputeCluster.
    """
    def __init__(__self__, resource_pool_id=None, id=None):
        if resource_pool_id and not isinstance(resource_pool_id, str):
            raise TypeError('Expected argument resource_pool_id to be a str')
        __self__.resource_pool_id = resource_pool_id
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_compute_cluster(datacenter_id=None, name=None):
    """
    The `vsphere_compute_cluster` data source can be used to discover the ID of a
    cluster in vSphere. This is useful to fetch the ID of a cluster that you want
    to use for virtual machine placement via the
    [`vsphere_virtual_machine`][docs-virtual-machine-resource] resource, allowing
    you to specify the cluster's root resource pool directly versus using the alias
    available through the [`vsphere_resource_pool`][docs-resource-pool-data-source]
    data source.
    
    [docs-virtual-machine-resource]: /docs/providers/vsphere/r/virtual_machine.html
    [docs-resource-pool-data-source]: /docs/providers/vsphere/d/resource_pool.html
    
    -> You may also wish to see the
    [`vsphere_compute_cluster`][docs-compute-cluster-resource] resource for further
    details about clusters or how to work with them in Terraform.
    
    [docs-compute-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
    """
    __args__ = dict()

    __args__['datacenterId'] = datacenter_id
    __args__['name'] = name
    __ret__ = await pulumi.runtime.invoke('vsphere:index/getComputeCluster:getComputeCluster', __args__)

    return GetComputeClusterResult(
        resource_pool_id=__ret__.get('resourcePoolId'),
        id=__ret__.get('id'))
