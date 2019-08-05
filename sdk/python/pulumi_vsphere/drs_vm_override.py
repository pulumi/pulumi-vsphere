# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class DrsVmOverride(pulumi.CustomResource):
    compute_cluster_id: pulumi.Output[str]
    """
    The [managed object reference
    ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
    resource if changed.
    """
    drs_automation_level: pulumi.Output[str]
    """
    Overrides the automation level for this virtual
    machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
    `fullyAutomated`. Default: `manual`.
    """
    drs_enabled: pulumi.Output[bool]
    """
    Overrides the default DRS setting for this virtual
    machine. Can be either `true` or `false`. Default: `false`.
    """
    virtual_machine_id: pulumi.Output[str]
    """
    The UUID of the virtual machine to create
    the override for.  Forces a new resource if changed.
    """
    def __init__(__self__, resource_name, opts=None, compute_cluster_id=None, drs_automation_level=None, drs_enabled=None, virtual_machine_id=None, __name__=None, __opts__=None):
        """
        The `vsphere_drs_vm_override` resource can be used to add a DRS override to a
        cluster for a specific virtual machine. With this resource, one can enable or
        disable DRS and control the automation level for a single virtual machine
        without affecting the rest of the cluster.
        
        For more information on vSphere clusters and DRS, see [this
        page][ref-vsphere-drs-clusters].
        
        [ref-vsphere-drs-clusters]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-8ACF3502-5314-469F-8CC9-4A9BD5925BC2.html
        
        > **NOTE:** This resource requires vCenter and is not available on direct ESXi
        connections.
        
        > **NOTE:** vSphere DRS requires a vSphere Enterprise Plus license.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_cluster_id: The [managed object reference
               ID][docs-about-morefs] of the cluster to put the override in.  Forces a new
               resource if changed.
        :param pulumi.Input[str] drs_automation_level: Overrides the automation level for this virtual
               machine in the cluster. Can be one of `manual`, `partiallyAutomated`, or
               `fullyAutomated`. Default: `manual`.
        :param pulumi.Input[bool] drs_enabled: Overrides the default DRS setting for this virtual
               machine. Can be either `true` or `false`. Default: `false`.
        :param pulumi.Input[str] virtual_machine_id: The UUID of the virtual machine to create
               the override for.  Forces a new resource if changed.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/drs_vm_override.html.markdown.
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

        __props__['drs_automation_level'] = drs_automation_level

        __props__['drs_enabled'] = drs_enabled

        if virtual_machine_id is None:
            raise TypeError("Missing required property 'virtual_machine_id'")
        __props__['virtual_machine_id'] = virtual_machine_id

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(DrsVmOverride, __self__).__init__(
            'vsphere:index/drsVmOverride:DrsVmOverride',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

