# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['VappContainer']


class VappContainer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu_expandable: Optional[pulumi.Input[bool]] = None,
                 cpu_limit: Optional[pulumi.Input[int]] = None,
                 cpu_reservation: Optional[pulumi.Input[int]] = None,
                 cpu_share_level: Optional[pulumi.Input[str]] = None,
                 cpu_shares: Optional[pulumi.Input[int]] = None,
                 custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 memory_expandable: Optional[pulumi.Input[bool]] = None,
                 memory_limit: Optional[pulumi.Input[int]] = None,
                 memory_reservation: Optional[pulumi.Input[int]] = None,
                 memory_share_level: Optional[pulumi.Input[str]] = None,
                 memory_shares: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_folder_id: Optional[pulumi.Input[str]] = None,
                 parent_resource_pool_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a VappContainer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cpu_expandable: Determines if the reservation on a vApp
               container can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[int] cpu_limit: The CPU utilization of a vApp container will not
               exceed this limit, even if there are available resources. Set to `-1` for
               unlimited.
               Default: `-1`
        :param pulumi.Input[int] cpu_reservation: Amount of CPU (MHz) that is guaranteed
               available to the vApp container. Default: `0`
        :param pulumi.Input[str] cpu_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `cpu_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[int] cpu_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `cpu_share_level` must be `custom`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[bool] memory_expandable: Determines if the reservation on a vApp
               container can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[int] memory_limit: The CPU utilization of a vApp container will not
               exceed this limit, even if there are available resources. Set to `-1` for
               unlimited.
               Default: `-1`
        :param pulumi.Input[int] memory_reservation: Amount of CPU (MHz) that is guaranteed
               available to the vApp container. Default: `0`
        :param pulumi.Input[str] memory_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `memory_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[int] memory_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `memory_share_level` must be `custom`.
        :param pulumi.Input[str] name: The name of the vApp container.
        :param pulumi.Input[str] parent_folder_id: The managed object ID of
               the vApp container's parent folder.
        :param pulumi.Input[str] parent_resource_pool_id: The managed object ID
               of the parent resource pool. This can be the root resource pool for a cluster
               or standalone host, or a resource pool itself. When moving a vApp container
               from one parent resource pool to another, both must share a common root
               resource pool or the move will fail.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
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
            opts.version = _utilities.get_version()
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
            __props__['parent_folder_id'] = parent_folder_id
            if parent_resource_pool_id is None:
                raise TypeError("Missing required property 'parent_resource_pool_id'")
            __props__['parent_resource_pool_id'] = parent_resource_pool_id
            __props__['tags'] = tags
        super(VappContainer, __self__).__init__(
            'vsphere:index/vappContainer:VappContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cpu_expandable: Optional[pulumi.Input[bool]] = None,
            cpu_limit: Optional[pulumi.Input[int]] = None,
            cpu_reservation: Optional[pulumi.Input[int]] = None,
            cpu_share_level: Optional[pulumi.Input[str]] = None,
            cpu_shares: Optional[pulumi.Input[int]] = None,
            custom_attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            memory_expandable: Optional[pulumi.Input[bool]] = None,
            memory_limit: Optional[pulumi.Input[int]] = None,
            memory_reservation: Optional[pulumi.Input[int]] = None,
            memory_share_level: Optional[pulumi.Input[str]] = None,
            memory_shares: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_folder_id: Optional[pulumi.Input[str]] = None,
            parent_resource_pool_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'VappContainer':
        """
        Get an existing VappContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cpu_expandable: Determines if the reservation on a vApp
               container can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[int] cpu_limit: The CPU utilization of a vApp container will not
               exceed this limit, even if there are available resources. Set to `-1` for
               unlimited.
               Default: `-1`
        :param pulumi.Input[int] cpu_reservation: Amount of CPU (MHz) that is guaranteed
               available to the vApp container. Default: `0`
        :param pulumi.Input[str] cpu_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `cpu_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[int] cpu_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `cpu_share_level` must be `custom`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_attributes: A list of custom attributes to set on this resource.
        :param pulumi.Input[bool] memory_expandable: Determines if the reservation on a vApp
               container can grow beyond the specified value if the parent resource pool has
               unreserved resources. Default: `true`
        :param pulumi.Input[int] memory_limit: The CPU utilization of a vApp container will not
               exceed this limit, even if there are available resources. Set to `-1` for
               unlimited.
               Default: `-1`
        :param pulumi.Input[int] memory_reservation: Amount of CPU (MHz) that is guaranteed
               available to the vApp container. Default: `0`
        :param pulumi.Input[str] memory_share_level: The CPU allocation level. The level is a
               simplified view of shares. Levels map to a pre-determined set of numeric
               values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
               `low`, `normal`, or `high` are specified values in `memory_shares` will be
               ignored.  Default: `normal`
        :param pulumi.Input[int] memory_shares: The number of shares allocated for CPU. Used to
               determine resource allocation in case of resource contention. If this is set,
               `memory_share_level` must be `custom`.
        :param pulumi.Input[str] name: The name of the vApp container.
        :param pulumi.Input[str] parent_folder_id: The managed object ID of
               the vApp container's parent folder.
        :param pulumi.Input[str] parent_resource_pool_id: The managed object ID
               of the parent resource pool. This can be the root resource pool for a cluster
               or standalone host, or a resource pool itself. When moving a vApp container
               from one parent resource pool to another, both must share a common root
               resource pool or the move will fail.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The IDs of any tags to attach to this resource.
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
        __props__["parent_folder_id"] = parent_folder_id
        __props__["parent_resource_pool_id"] = parent_resource_pool_id
        __props__["tags"] = tags
        return VappContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cpuExpandable")
    def cpu_expandable(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if the reservation on a vApp
        container can grow beyond the specified value if the parent resource pool has
        unreserved resources. Default: `true`
        """
        return pulumi.get(self, "cpu_expandable")

    @property
    @pulumi.getter(name="cpuLimit")
    def cpu_limit(self) -> pulumi.Output[Optional[int]]:
        """
        The CPU utilization of a vApp container will not
        exceed this limit, even if there are available resources. Set to `-1` for
        unlimited.
        Default: `-1`
        """
        return pulumi.get(self, "cpu_limit")

    @property
    @pulumi.getter(name="cpuReservation")
    def cpu_reservation(self) -> pulumi.Output[Optional[int]]:
        """
        Amount of CPU (MHz) that is guaranteed
        available to the vApp container. Default: `0`
        """
        return pulumi.get(self, "cpu_reservation")

    @property
    @pulumi.getter(name="cpuShareLevel")
    def cpu_share_level(self) -> pulumi.Output[Optional[str]]:
        """
        The CPU allocation level. The level is a
        simplified view of shares. Levels map to a pre-determined set of numeric
        values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
        `low`, `normal`, or `high` are specified values in `cpu_shares` will be
        ignored.  Default: `normal`
        """
        return pulumi.get(self, "cpu_share_level")

    @property
    @pulumi.getter(name="cpuShares")
    def cpu_shares(self) -> pulumi.Output[int]:
        """
        The number of shares allocated for CPU. Used to
        determine resource allocation in case of resource contention. If this is set,
        `cpu_share_level` must be `custom`.
        """
        return pulumi.get(self, "cpu_shares")

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A list of custom attributes to set on this resource.
        """
        return pulumi.get(self, "custom_attributes")

    @property
    @pulumi.getter(name="memoryExpandable")
    def memory_expandable(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if the reservation on a vApp
        container can grow beyond the specified value if the parent resource pool has
        unreserved resources. Default: `true`
        """
        return pulumi.get(self, "memory_expandable")

    @property
    @pulumi.getter(name="memoryLimit")
    def memory_limit(self) -> pulumi.Output[Optional[int]]:
        """
        The CPU utilization of a vApp container will not
        exceed this limit, even if there are available resources. Set to `-1` for
        unlimited.
        Default: `-1`
        """
        return pulumi.get(self, "memory_limit")

    @property
    @pulumi.getter(name="memoryReservation")
    def memory_reservation(self) -> pulumi.Output[Optional[int]]:
        """
        Amount of CPU (MHz) that is guaranteed
        available to the vApp container. Default: `0`
        """
        return pulumi.get(self, "memory_reservation")

    @property
    @pulumi.getter(name="memoryShareLevel")
    def memory_share_level(self) -> pulumi.Output[Optional[str]]:
        """
        The CPU allocation level. The level is a
        simplified view of shares. Levels map to a pre-determined set of numeric
        values for shares. Can be one of `low`, `normal`, `high`, or `custom`. When
        `low`, `normal`, or `high` are specified values in `memory_shares` will be
        ignored.  Default: `normal`
        """
        return pulumi.get(self, "memory_share_level")

    @property
    @pulumi.getter(name="memoryShares")
    def memory_shares(self) -> pulumi.Output[int]:
        """
        The number of shares allocated for CPU. Used to
        determine resource allocation in case of resource contention. If this is set,
        `memory_share_level` must be `custom`.
        """
        return pulumi.get(self, "memory_shares")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the vApp container.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentFolderId")
    def parent_folder_id(self) -> pulumi.Output[Optional[str]]:
        """
        The managed object ID of
        the vApp container's parent folder.
        """
        return pulumi.get(self, "parent_folder_id")

    @property
    @pulumi.getter(name="parentResourcePoolId")
    def parent_resource_pool_id(self) -> pulumi.Output[str]:
        """
        The managed object ID
        of the parent resource pool. This can be the root resource pool for a cluster
        or standalone host, or a resource pool itself. When moving a vApp container
        from one parent resource pool to another, both must share a common root
        resource pool or the move will fail.
        """
        return pulumi.get(self, "parent_resource_pool_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of any tags to attach to this resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

