# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class VmfsDatastore(pulumi.CustomResource):
    accessible: pulumi.Output[bool]
    """
    The connectivity status of the datastore. If this is `false`,
    some other computed attributes may be out of date.
    """
    capacity: pulumi.Output[float]
    """
    Maximum capacity of the datastore, in megabytes.
    """
    custom_attributes: pulumi.Output[dict]
    """
    Map of custom attribute ids to attribute 
    value string to set on datastore resource.
    """
    datastore_cluster_id: pulumi.Output[str]
    """
    The managed object
    ID of a datastore cluster to put this datastore in.
    Conflicts with `folder`.
    """
    disks: pulumi.Output[list]
    """
    The disks to use with the datastore.
    """
    folder: pulumi.Output[str]
    """
    The relative path to a folder to put this datastore in.
    This is a path relative to the datacenter you are deploying the datastore to.
    Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
    The provider will place a datastore named `test` in a datastore folder
    located at `/dc1/datastore/foo/bar`, with the final inventory path being
    `/dc1/datastore/foo/bar/test`. Conflicts with
    `datastore_cluster_id`.
    """
    free_space: pulumi.Output[float]
    """
    Available space of this datastore, in megabytes.
    """
    host_system_id: pulumi.Output[str]
    """
    The managed object ID of
    the host to set the datastore up on. Note that this is not necessarily the
    only host that the datastore will be set up on - see
    here for more info. Forces a
    new resource if changed.
    """
    maintenance_mode: pulumi.Output[str]
    """
    The current maintenance mode state of the datastore.
    """
    multiple_host_access: pulumi.Output[bool]
    """
    If `true`, more than one host in the datacenter has
    been configured with access to the datastore.
    """
    name: pulumi.Output[str]
    """
    The name of the datastore. Forces a new resource if
    changed.
    """
    tags: pulumi.Output[list]
    """
    The IDs of any tags to attach to this resource.
    """
    uncommitted_space: pulumi.Output[float]
    """
    Total additional storage space, in megabytes,
    potentially used by all virtual machines on this datastore.
    """
    url: pulumi.Output[str]
    """
    The unique locator for the datastore.
    """
    def __init__(__self__, resource_name, opts=None, custom_attributes=None, datastore_cluster_id=None, disks=None, folder=None, host_system_id=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a VmfsDatastore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] custom_attributes: Map of custom attribute ids to attribute 
               value string to set on datastore resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object
               ID of a datastore cluster to put this datastore in.
               Conflicts with `folder`.
        :param pulumi.Input[list] disks: The disks to use with the datastore.
        :param pulumi.Input[str] folder: The relative path to a folder to put this datastore in.
               This is a path relative to the datacenter you are deploying the datastore to.
               Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
               The provider will place a datastore named `test` in a datastore folder
               located at `/dc1/datastore/foo/bar`, with the final inventory path being
               `/dc1/datastore/foo/bar/test`. Conflicts with
               `datastore_cluster_id`.
        :param pulumi.Input[str] host_system_id: The managed object ID of
               the host to set the datastore up on. Note that this is not necessarily the
               only host that the datastore will be set up on - see
               here for more info. Forces a
               new resource if changed.
        :param pulumi.Input[str] name: The name of the datastore. Forces a new resource if
               changed.
        :param pulumi.Input[list] tags: The IDs of any tags to attach to this resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['custom_attributes'] = custom_attributes
            __props__['datastore_cluster_id'] = datastore_cluster_id
            if disks is None:
                raise TypeError("Missing required property 'disks'")
            __props__['disks'] = disks
            __props__['folder'] = folder
            if host_system_id is None:
                raise TypeError("Missing required property 'host_system_id'")
            __props__['host_system_id'] = host_system_id
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['accessible'] = None
            __props__['capacity'] = None
            __props__['free_space'] = None
            __props__['maintenance_mode'] = None
            __props__['multiple_host_access'] = None
            __props__['uncommitted_space'] = None
            __props__['url'] = None
        super(VmfsDatastore, __self__).__init__(
            'vsphere:index/vmfsDatastore:VmfsDatastore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accessible=None, capacity=None, custom_attributes=None, datastore_cluster_id=None, disks=None, folder=None, free_space=None, host_system_id=None, maintenance_mode=None, multiple_host_access=None, name=None, tags=None, uncommitted_space=None, url=None):
        """
        Get an existing VmfsDatastore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accessible: The connectivity status of the datastore. If this is `false`,
               some other computed attributes may be out of date.
        :param pulumi.Input[float] capacity: Maximum capacity of the datastore, in megabytes.
        :param pulumi.Input[dict] custom_attributes: Map of custom attribute ids to attribute 
               value string to set on datastore resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object
               ID of a datastore cluster to put this datastore in.
               Conflicts with `folder`.
        :param pulumi.Input[list] disks: The disks to use with the datastore.
        :param pulumi.Input[str] folder: The relative path to a folder to put this datastore in.
               This is a path relative to the datacenter you are deploying the datastore to.
               Example: for the `dc1` datacenter, and a provided `folder` of `foo/bar`,
               The provider will place a datastore named `test` in a datastore folder
               located at `/dc1/datastore/foo/bar`, with the final inventory path being
               `/dc1/datastore/foo/bar/test`. Conflicts with
               `datastore_cluster_id`.
        :param pulumi.Input[float] free_space: Available space of this datastore, in megabytes.
        :param pulumi.Input[str] host_system_id: The managed object ID of
               the host to set the datastore up on. Note that this is not necessarily the
               only host that the datastore will be set up on - see
               here for more info. Forces a
               new resource if changed.
        :param pulumi.Input[str] maintenance_mode: The current maintenance mode state of the datastore.
        :param pulumi.Input[bool] multiple_host_access: If `true`, more than one host in the datacenter has
               been configured with access to the datastore.
        :param pulumi.Input[str] name: The name of the datastore. Forces a new resource if
               changed.
        :param pulumi.Input[list] tags: The IDs of any tags to attach to this resource.
        :param pulumi.Input[float] uncommitted_space: Total additional storage space, in megabytes,
               potentially used by all virtual machines on this datastore.
        :param pulumi.Input[str] url: The unique locator for the datastore.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessible"] = accessible
        __props__["capacity"] = capacity
        __props__["custom_attributes"] = custom_attributes
        __props__["datastore_cluster_id"] = datastore_cluster_id
        __props__["disks"] = disks
        __props__["folder"] = folder
        __props__["free_space"] = free_space
        __props__["host_system_id"] = host_system_id
        __props__["maintenance_mode"] = maintenance_mode
        __props__["multiple_host_access"] = multiple_host_access
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["uncommitted_space"] = uncommitted_space
        __props__["url"] = url
        return VmfsDatastore(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

