# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class ComputeClusterVmAffinityRule(pulumi.CustomResource):
    compute_cluster_id: pulumi.Output[str]
    """
    The [managed object reference
    ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
    resource if changed.
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
    on the same host together.
    """
    def __init__(__self__, resource_name, opts=None, compute_cluster_id=None, enabled=None, mandatory=None, name=None, virtual_machine_ids=None, __name__=None, __opts__=None):
        """
        The `vsphere_compute_cluster_vm_affinity_rule` resource can be used to manage
        VM affinity rules in a cluster, either created by the
        [`vsphere_compute_cluster`][tf-vsphere-cluster-resource] resource or looked up
        by the [`vsphere_compute_cluster`][tf-vsphere-cluster-data-source] data source.
        
        [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html
        [tf-vsphere-cluster-data-source]: /docs/providers/vsphere/d/compute_cluster.html
        
        This rule can be used to tell a set to virtual machines to run together on a
        single host within a cluster. When configured, DRS will make a best effort to
        ensure that the virtual machines run on the same host, or prevent any operation
        that would keep that from happening, depending on the value of the
        `mandatory` flag.
        
        > Keep in mind that this rule can only be used to tell VMs to run together on
        a _non-specific_ host - it can't be used to pin VMs to a host. For that, see
        the
        [`vsphere_compute_cluster_vm_host_rule`][tf-vsphere-cluster-vm-host-rule-resource]
        resource.
        
        [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
        
        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.
        
        > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The [managed object reference
               ID][docs-about-morefs] of the cluster to put the group in.  Forces a new
               resource if changed.
        :param pulumi.Input[bool] enabled: Enable this rule in the cluster. Default: `true`.
        :param pulumi.Input[bool] mandatory: When this value is `true`, prevents any virtual
               machine operations that may violate this rule. Default: `false`.
        :param pulumi.Input[str] name: The name of the rule. This must be unique in the cluster.
        :param pulumi.Input[list] virtual_machine_ids: The UUIDs of the virtual machines to run
               on the same host together.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/compute_cluster_vm_affinity_rule.html.markdown.
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

        if compute_cluster_id is None:
            raise TypeError("Missing required property 'compute_cluster_id'")
        __props__['compute_cluster_id'] = compute_cluster_id

        __props__['enabled'] = enabled

        __props__['mandatory'] = mandatory

        __props__['name'] = name

        if virtual_machine_ids is None:
            raise TypeError("Missing required property 'virtual_machine_ids'")
        __props__['virtual_machine_ids'] = virtual_machine_ids

        super(ComputeClusterVmAffinityRule, __self__).__init__(
            'vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

