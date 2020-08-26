# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['Tag']


class Tag(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Tag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category_id: The unique identifier of the parent category in
               which this tag will be created. Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the tag.
        :param pulumi.Input[str] name: The display name of the tag. The name must be unique
               within its category.
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

            if category_id is None:
                raise TypeError("Missing required property 'category_id'")
            __props__['category_id'] = category_id
            __props__['description'] = description
            __props__['name'] = name
        super(Tag, __self__).__init__(
            'vsphere:index/tag:Tag',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            category_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Tag':
        """
        Get an existing Tag resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] category_id: The unique identifier of the parent category in
               which this tag will be created. Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the tag.
        :param pulumi.Input[str] name: The display name of the tag. The name must be unique
               within its category.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["category_id"] = category_id
        __props__["description"] = description
        __props__["name"] = name
        return Tag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> str:
        """
        The unique identifier of the parent category in
        which this tag will be created. Forces a new resource if changed.
        """
        return pulumi.get(self, "category_id")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description for the tag.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The display name of the tag. The name must be unique
        within its category.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

