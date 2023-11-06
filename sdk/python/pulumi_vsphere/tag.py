# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TagArgs', 'Tag']

@pulumi.input_type
class TagArgs:
    def __init__(__self__, *,
                 category_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Tag resource.
        :param pulumi.Input[str] category_id: The unique identifier of the parent category in
               which this tag will be created. Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the tag.
        :param pulumi.Input[str] name: The display name of the tag. The name must be unique
               within its category.
        """
        TagArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            category_id=category_id,
            description=description,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             category_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if category_id is None and 'categoryId' in kwargs:
            category_id = kwargs['categoryId']
        if category_id is None:
            raise TypeError("Missing 'category_id' argument")

        _setter("category_id", category_id)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> pulumi.Input[str]:
        """
        The unique identifier of the parent category in
        which this tag will be created. Forces a new resource if changed.
        """
        return pulumi.get(self, "category_id")

    @category_id.setter
    def category_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "category_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the tag.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the tag. The name must be unique
        within its category.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _TagState:
    def __init__(__self__, *,
                 category_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Tag resources.
        :param pulumi.Input[str] category_id: The unique identifier of the parent category in
               which this tag will be created. Forces a new resource if changed.
        :param pulumi.Input[str] description: A description for the tag.
        :param pulumi.Input[str] name: The display name of the tag. The name must be unique
               within its category.
        """
        _TagState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            category_id=category_id,
            description=description,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             category_id: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if category_id is None and 'categoryId' in kwargs:
            category_id = kwargs['categoryId']

        if category_id is not None:
            _setter("category_id", category_id)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the parent category in
        which this tag will be created. Forces a new resource if changed.
        """
        return pulumi.get(self, "category_id")

    @category_id.setter
    def category_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "category_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the tag.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the tag. The name must be unique
        within its category.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Tag(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Tag resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TagArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TagArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 category_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TagArgs.__new__(TagArgs)

            if category_id is None and not opts.urn:
                raise TypeError("Missing required property 'category_id'")
            __props__.__dict__["category_id"] = category_id
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
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

        __props__ = _TagState.__new__(_TagState)

        __props__.__dict__["category_id"] = category_id
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        return Tag(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the parent category in
        which this tag will be created. Forces a new resource if changed.
        """
        return pulumi.get(self, "category_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the tag.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The display name of the tag. The name must be unique
        within its category.
        """
        return pulumi.get(self, "name")

