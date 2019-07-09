# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class CustomAttribute(pulumi.CustomResource):
    managed_object_type: pulumi.Output[str]
    """
    The object type that this attribute may be
    applied to. If not set, the custom attribute may be applied to any object
    type. For a full list, click here. Forces a new
    resource if changed.
    """
    name: pulumi.Output[str]
    """
    The name of the custom attribute.
    """
    def __init__(__self__, resource_name, opts=None, managed_object_type=None, name=None, __name__=None, __opts__=None):
        """
        The `vsphere_custom_attribute` resource can be used to create and manage custom
        attributes, which allow users to associate user-specific meta-information with 
        vSphere managed objects. Custom attribute values must be strings and are stored 
        on the vCenter Server and not the managed object.
        
        For more information about custom attributes, click [here][ext-custom-attributes].
        
        [ext-custom-attributes]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vcenterhost.doc/GUID-73606C4C-763C-4E27-A1DA-032E4C46219D.html
        
        > **NOTE:** Custom attributes are unsupported on direct ESXi connections 
        and require vCenter.
        
        ## Managed Object Types
        
        The following table will help you determine what value you need to enter for 
        the managed object type you want the attribute to apply to.
        
        Note that if you want a attribute to apply to all objects, leave the type 
        unspecified.
        
        <table>
        <tr><th>Type</th><th>Value</th></tr>
        <tr><td>Folders</td><td>`Folder`</td></tr>
        <tr><td>Clusters</td><td>`ClusterComputeResource`</td></tr>
        <tr><td>Datacenters</td><td>`Datacenter`</td></tr>
        <tr><td>Datastores</td><td>`Datastore`</td></tr>
        <tr><td>Datastore Clusters</td><td>`StoragePod`</td></tr>
        <tr><td>DVS Portgroups</td><td>`DistributedVirtualPortgroup`</td></tr>
        <tr><td>Distributed vSwitches</td><td>`DistributedVirtualSwitch`<br>`VmwareDistributedVirtualSwitch`</td></tr>
        <tr><td>Hosts</td><td>`HostSystem`</td></tr>
        <tr><td>Content Libraries</td><td>`com.vmware.content.Library`</td></tr>
        <tr><td>Content Library Items</td><td>`com.vmware.content.library.Item`</td></tr>
        <tr><td>Networks</td><td>`HostNetwork`<br>`Network`<br>`OpaqueNetwork`</td></tr>
        <tr><td>Resource Pools</td><td>`ResourcePool`</td></tr>
        <tr><td>vApps</td><td>`VirtualApp`</td></tr>
        <tr><td>Virtual Machines</td><td>`VirtualMachine`</td></tr>
        </table>
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] managed_object_type: The object type that this attribute may be
               applied to. If not set, the custom attribute may be applied to any object
               type. For a full list, click here. Forces a new
               resource if changed.
        :param pulumi.Input[str] name: The name of the custom attribute.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/custom_attribute.html.markdown.
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

        __props__['managed_object_type'] = managed_object_type

        __props__['name'] = name

        super(CustomAttribute, __self__).__init__(
            'vsphere:index/customAttribute:CustomAttribute',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

