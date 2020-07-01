# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class DatastoreClusterVmAntiAffinityRule(pulumi.CustomResource):
    datastore_cluster_id: pulumi.Output[str]
    """
    The managed object reference
    ID of the datastore cluster to put the group in.  Forces
    a new resource if changed.
    """
    enabled: pulumi.Output[bool]
    """
    Enable this rule in the cluster. Default: `true`.
    """
    mandatory: pulumi.Output[bool]
    """
    When this value is `true`, prevents any virtual
    machine operations that may violate this rule. Default: `false`.
    """
    name: pulumi.Output[str]
    """
    The name of the rule. This must be unique in the cluster.
    """
    virtual_machine_ids: pulumi.Output[list]
    """
    The UUIDs of the virtual machines to run
    on different datastores from each other.
    """
    def __init__(__self__, resource_name, opts=None, datastore_cluster_id=None, enabled=None, mandatory=None, name=None, virtual_machine_ids=None, __props__=None, __name__=None, __opts__=None):
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
        :param pulumi.Input[list] virtual_machine_ids: The UUIDs of the virtual machines to run
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
            opts.version = utilities.get_version()
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
    def get(resource_name, id, opts=None, datastore_cluster_id=None, enabled=None, mandatory=None, name=None, virtual_machine_ids=None):
        """
        Get an existing DatastoreClusterVmAntiAffinityRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datastore_cluster_id: The managed object reference
               ID of the datastore cluster to put the group in.  Forces
               a new resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[list] virtual_machine_ids: The UUIDs of the virtual machines to run
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
