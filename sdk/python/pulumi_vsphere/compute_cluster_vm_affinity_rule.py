# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from . import utilities, tables

class ComputeClusterVmAffinityRule(pulumi.CustomResource):
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
    
    -> Keep in mind that this rule can only be used to tell VMs to run together on
    a _non-specific_ host - it can't be used to pin VMs to a host. For that, see
    the
    [`vsphere_compute_cluster_vm_host_rule`][tf-vsphere-cluster-vm-host-rule-resource]
    resource.
    
    [tf-vsphere-cluster-vm-host-rule-resource]: /docs/providers/vsphere/r/compute_cluster_vm_host_rule.html
    
    ~> **NOTE:** This resource requires vCenter and is not available on direct ESXi
    connections.
    
    ~> **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
    """
    def __init__(__self__, __name__, __opts__=None, compute_cluster_id=None, enabled=None, mandatory=None, name=None, virtual_machine_ids=None):
        """Create a ComputeClusterVmAffinityRule resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not compute_cluster_id:
            raise TypeError('Missing required property compute_cluster_id')
        __props__['compute_cluster_id'] = compute_cluster_id

        __props__['enabled'] = enabled

        __props__['mandatory'] = mandatory

        __props__['name'] = name

        if not virtual_machine_ids:
            raise TypeError('Missing required property virtual_machine_ids')
        __props__['virtual_machine_ids'] = virtual_machine_ids

        super(ComputeClusterVmAffinityRule, __self__).__init__(
            'vsphere:index/computeClusterVmAffinityRule:ComputeClusterVmAffinityRule',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

