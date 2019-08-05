# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class DatastoreClusterVmAntiAffinityRule(pulumi.CustomResource):
    datastore_cluster_id: pulumi.Output[str]
    """
    The [managed object reference
    ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
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
    def __init__(__self__, resource_name, opts=None, datastore_cluster_id=None, enabled=None, mandatory=None, name=None, virtual_machine_ids=None, __name__=None, __opts__=None):
        """
        The `vsphere_datastore_cluster_vm_anti_affinity_rule` resource can be used to
        manage VM anti-affinity rules in a datastore cluster, either created by the
        [`vsphere_datastore_cluster`][tf-vsphere-datastore-cluster-resource] resource or looked up
        by the [`vsphere_datastore_cluster`][tf-vsphere-datastore-cluster-data-source] data source.
        
        [tf-vsphere-datastore-cluster-resource]: /docs/providers/vsphere/r/datastore_cluster.html
        [tf-vsphere-datastore-cluster-data-source]: /docs/providers/vsphere/d/datastore_cluster.html
        
        This rule can be used to tell a set to virtual machines to run on different
        datastores within a cluster, useful for preventing single points of failure in
        application cluster scenarios. When configured, Storage DRS will make a best effort to
        ensure that the virtual machines run on different datastores, or prevent any
        operation that would keep that from happening, depending on the value of the
        `mandatory` flag.
        
        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.
        
        > **NOTE:** Storage DRS requires a vSphere Enterprise Plus license.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] datastore_cluster_id: The [managed object reference
               ID][docs-about-morefs] of the datastore cluster to put the group in.  Forces
               a new resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[list] virtual_machine_ids: The UUIDs of the virtual machines to run
               on different datastores from each other.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/datastore_cluster_vm_anti_affinity_rule.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

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

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(DatastoreClusterVmAntiAffinityRule, __self__).__init__(
            'vsphere:index/datastoreClusterVmAntiAffinityRule:DatastoreClusterVmAntiAffinityRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

