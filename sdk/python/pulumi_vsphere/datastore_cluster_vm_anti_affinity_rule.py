# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['DatastoreClusterVmAntiAffinityRule']


class DatastoreClusterVmAntiAffinityRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 datastore_cluster_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 mandatory: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a DatastoreClusterVmAntiAffinityRule resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object reference
               ID of the datastore cluster to put the group in.  Forces
               a new resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines to run
               on different datastores from each other.
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

            if datastore_cluster_id is None:
                raise TypeError("Missing required property 'datastore_cluster_id'")
            __props__['datastore_cluster_id'] = datastore_cluster_id
            __props__['enabled'] = enabled
            __props__['mandatory'] = mandatory
            __props__['name'] = name
            if virtual_machine_ids is None:
                raise TypeError("Missing required property 'virtual_machine_ids'")
            __props__['virtual_machine_ids'] = virtual_machine_ids
        super(DatastoreClusterVmAntiAffinityRule, __self__).__init__(
            'vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            datastore_cluster_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            mandatory: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            virtual_machine_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'DatastoreClusterVmAntiAffinityRule':
        """
        Get an existing DatastoreClusterVmAntiAffinityRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object reference
               ID of the datastore cluster to put the group in.  Forces
               a new resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] virtual_machine_ids: The UUIDs of the virtual machines to run
               on different datastores from each other.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["datastore_cluster_id"] = datastore_cluster_id
        __props__["enabled"] = enabled
        __props__["mandatory"] = mandatory
        __props__["name"] = name
        __props__["virtual_machine_ids"] = virtual_machine_ids
        return DatastoreClusterVmAntiAffinityRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="datastoreClusterId")
    def datastore_cluster_id(self) -> pulumi.Output[str]:
        """
        The managed object reference
        ID of the datastore cluster to put the group in.  Forces
        a new resource if changed.
        """
        return pulumi.get(self, "datastore_cluster_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable this rule in the cluster. Default: `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def mandatory(self) -> pulumi.Output[Optional[bool]]:
        """
        When this value is `true`, prevents any virtual
        machine operations that may violate this rule. Default: `false`.
        """
        return pulumi.get(self, "mandatory")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the rule. This must be unique in the cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="virtualMachineIds")
    def virtual_machine_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The UUIDs of the virtual machines to run
        on different datastores from each other.
        """
        return pulumi.get(self, "virtual_machine_ids")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

