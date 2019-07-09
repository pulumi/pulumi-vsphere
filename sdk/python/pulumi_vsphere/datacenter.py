# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Datacenter(pulumi.CustomResource):
    custom_attributes: pulumi.Output[dict]
    """
    Map of custom attribute ids to value 
    strings to set for datacenter resource. See
    [here][docs-setting-custom-attributes] for a reference on how to set values
    for custom attributes.
    """
    folder: pulumi.Output[str]
    """
    The folder where the datacenter should be created.
    Forces a new resource if changed.
    """
    moid: pulumi.Output[str]
    """
    [Managed object ID][docs-about-morefs] of this datacenter.
    """
    name: pulumi.Output[str]
    """
    The name of the datacenter. This name needs to be unique
    within the folder. Forces a new resource if changed.
    """
    tags: pulumi.Output[list]
    """
    The IDs of any tags to attach to this resource. See
    [here][docs-applying-tags] for a reference on how to apply tags.
    """
    def __init__(__self__, resource_name, opts=None, custom_attributes=None, folder=None, name=None, tags=None, __name__=None, __opts__=None):
        """
        Provides a VMware vSphere datacenter resource. This can be used as the primary
        container of inventory objects such as hosts and virtual machines.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] custom_attributes: Map of custom attribute ids to value 
               strings to set for datacenter resource. See
               [here][docs-setting-custom-attributes] for a reference on how to set values
               for custom attributes.
        :param pulumi.Input[str] folder: The folder where the datacenter should be created.
               Forces a new resource if changed.
        :param pulumi.Input[str] name: The name of the datacenter. This name needs to be unique
               within the folder. Forces a new resource if changed.
        :param pulumi.Input[list] tags: The IDs of any tags to attach to this resource. See
               [here][docs-applying-tags] for a reference on how to apply tags.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vsphere/blob/master/website/docs/r/datacenter.html.markdown.
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

        __props__['custom_attributes'] = custom_attributes

        __props__['folder'] = folder

        __props__['name'] = name

        __props__['tags'] = tags

        __props__['moid'] = None

        super(Datacenter, __self__).__init__(
            'vsphere:index/datacenter:Datacenter',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

