# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class ResourcePool(pulumi.CustomResource):
    cpu_expandable: pulumi.Output[bool]
    """
    Determines if the reservation on a resource
    pool can grow beyond the specified value if the parent resource pool has
    unreserved resources. Default: `true`
    """
    cpu_limit: pulumi.Output[float]
    """
    The CPU utilization of a resource pool will not exceed
    this limit, even if there are available resources. Set to `-1` for unlimited.
    Default: `-1`
    """
    cpu_reservation: pulumi.Output[float]
    """
    Amount of CPU (MHz) that is guaranteed
    available to the resource pool. Default: `0`
    """
    cpu_share_level: pulumi.Output[str]
    """
    The CPU allocation level. The level is a
    simplified view of shares. Levels map to a pre-determined set of numeric
    values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
    `low`, `normal`, or `high` are specified values in `cpu_shares` will be
    ignored.  Default: `normal`
    """
    cpu_shares: pulumi.Output[float]
    """
    The number of shares allocated for CPU. Used to
    determine resource allocation in case of resource contention. If this is set,
    `cpu_share_level` must be `custom`.
    """
    custom_attributes: pulumi.Output[dict]
    """
    A list of custom attributes to set on this resource.
    """
    memory_expandable: pulumi.Output[bool]
    """
    Determines if the reservation on a resource
    pool can grow beyond the specified value if the parent resource pool has
    unreserved resources. Default: `true`
    """
    memory_limit: pulumi.Output[float]
    """
    The CPU utilization of a resource pool will not exceed
    this limit, even if there are available resources. Set to `-1` for unlimited.
    Default: `-1`
    """
    memory_reservation: pulumi.Output[float]
    """
    Amount of CPU (MHz) that is guaranteed
    available to the resource pool. Default: `0`
    """
    memory_share_level: pulumi.Output[str]
    """
    The CPU allocation level. The level is a
    simplified view of shares. Levels map to a pre-determined set of numeric
    values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
    `low`, `normal`, or `high` are specified values in `memory_shares` will be
    ignored.  Default: `normal`
    """
    memory_shares: pulumi.Output[float]
    """
    The number of shares allocated for CPU. Used to
    determine resource allocation in case of resource contention. If this is set,
    `memory_share_level` must be `custom`.
    """
    name: pulumi.Output[str]
    """
    The name of the resource pool.
    """
    parent_resource_pool_id: pulumi.Output[str]
    """
    The managed object ID
    of the parent resource pool. This can be the root resource pool for a cluster
    or standalone host, or a resource pool itself. When moving a resource pool
    from one parent resource pool to another, both must share a common root
    resource pool or the move will fail.
    """
    tags: pulumi.Output[list]
    """
    The IDs of any tags to attach to this resource.
    """
    def __init__(__self__, resource_name, opts=None, cpu_expandable=None, cpu_limit=None, cpu_reservation=None, cpu_share_level=None, cpu_shares=None, custom_attributes=None, memory_expandable=None, memory_limit=None, memory_reservation=None, memory_share_level=None, memory_shares=None, name=None, parent_resource_pool_id=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a ResourcePool resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cpu_expandable: Determines if the reservation on a resource
               pool can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[float] cpu_limit: The CPU utilization of a resource pool will not exceed
               this limit, even if there are available resources. Set to `-1` for unlimited.
               Default: `-1`
        :param pulumi.Input[float] cpu_reservation: Amount of CPU (MHz) that is guaranteed
               available to the resource pool. Default: `0`
        :param pulumi.Input[str] cpu_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `cpu_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[float] cpu_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `cpu_share_level` must be `custom`.
        :param pulumi.Input[dict] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[bool] memory_expandable: Determines if the reservation on a resource
               pool can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[float] memory_limit: The CPU utilization of a resource pool will not exceed
               this limit, even if there are available resources. Set to `-1` for unlimited.
               Default: `-1`
        :param pulumi.Input[float] memory_reservation: Amount of CPU (MHz) that is guaranteed
               available to the resource pool. Default: `0`
        :param pulumi.Input[str] memory_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `memory_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[float] memory_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `memory_share_level` must be `custom`.
        :param pulumi.Input[str] name: The name of the resource pool.
        :param pulumi.Input[str] parent_resource_pool_id: The managed object ID
               of the parent resource pool. This can be the root resource pool for a cluster
               or standalone host, or a resource pool itself. When moving a resource pool
               from one parent resource pool to another, both must share a common root
               resource pool or the move will fail.
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

            __props__['cpu_expandable'] = cpu_expandable
            __props__['cpu_limit'] = cpu_limit
            __props__['cpu_reservation'] = cpu_reservation
            __props__['cpu_share_level'] = cpu_share_level
            __props__['cpu_shares'] = cpu_shares
            __props__['custom_attributes'] = custom_attributes
            __props__['memory_expandable'] = memory_expandable
            __props__['memory_limit'] = memory_limit
            __props__['memory_reservation'] = memory_reservation
            __props__['memory_share_level'] = memory_share_level
            __props__['memory_shares'] = memory_shares
            __props__['name'] = name
            if parent_resource_pool_id is None:
                raise TypeError("Missing required property 'parent_resource_pool_id'")
            __props__['parent_resource_pool_id'] = parent_resource_pool_id
            __props__['tags'] = tags
        super(ResourcePool, __self__).__init__(
            'vsphere:index/resourcePool:ResourcePool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, cpu_expandable=None, cpu_limit=None, cpu_reservation=None, cpu_share_level=None, cpu_shares=None, custom_attributes=None, memory_expandable=None, memory_limit=None, memory_reservation=None, memory_share_level=None, memory_shares=None, name=None, parent_resource_pool_id=None, tags=None):
        """
        Get an existing ResourcePool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cpu_expandable: Determines if the reservation on a resource
               pool can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[float] cpu_limit: The CPU utilization of a resource pool will not exceed
               this limit, even if there are available resources. Set to `-1` for unlimited.
               Default: `-1`
        :param pulumi.Input[float] cpu_reservation: Amount of CPU (MHz) that is guaranteed
               available to the resource pool. Default: `0`
        :param pulumi.Input[str] cpu_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `cpu_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[float] cpu_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `cpu_share_level` must be `custom`.
        :param pulumi.Input[dict] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[bool] memory_expandable: Determines if the reservation on a resource
               pool can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[float] memory_limit: The CPU utilization of a resource pool will not exceed
               this limit, even if there are available resources. Set to `-1` for unlimited.
               Default: `-1`
        :param pulumi.Input[float] memory_reservation: Amount of CPU (MHz) that is guaranteed
               available to the resource pool. Default: `0`
        :param pulumi.Input[str] memory_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `memory_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[float] memory_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `memory_share_level` must be `custom`.
        :param pulumi.Input[str] name: The name of the resource pool.
        :param pulumi.Input[str] parent_resource_pool_id: The managed object ID
               of the parent resource pool. This can be the root resource pool for a cluster
               or standalone host, or a resource pool itself. When moving a resource pool
               from one parent resource pool to another, both must share a common root
               resource pool or the move will fail.
        :param pulumi.Input[list] tags: The IDs of any tags to attach to this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cpu_expandable"] = cpu_expandable
        __props__["cpu_limit"] = cpu_limit
        __props__["cpu_reservation"] = cpu_reservation
        __props__["cpu_share_level"] = cpu_share_level
        __props__["cpu_shares"] = cpu_shares
        __props__["custom_attributes"] = custom_attributes
        __props__["memory_expandable"] = memory_expandable
        __props__["memory_limit"] = memory_limit
        __props__["memory_reservation"] = memory_reservation
        __props__["memory_share_level"] = memory_share_level
        __props__["memory_shares"] = memory_shares
        __props__["name"] = name
        __props__["parent_resource_pool_id"] = parent_resource_pool_id
        __props__["tags"] = tags
        return ResourcePool(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
